package server

import (
	"Notify-storage-service/internal/handler"
	"Notify-storage-service/internal/server/launcher"
	"Notify-storage-service/internal/server/launcher/rabbit"
	"context"
	"github.com/op/go-logging"
	"golang.org/x/sync/errgroup"
	"sync"

	"Notify-storage-service/internal/broker"

	"Notify-storage-service/internal/service"
)

var log = logging.MustGetLogger("server")

type server struct {
	servers []launcher.Server
}

func New(srv service.Service, broker broker.Broker) (launcher.Server, error) {
	h := handler.New(srv, broker)
	server := &server{
		servers: []launcher.Server{
			rabbit.New(broker.RabbitMQ, h.Event),
		},
	}

	return server, nil
}

func (s *server) Serve(ctx context.Context) error {
	gr, grCtx := errgroup.WithContext(ctx)

	gr.Go(func() error {
		return s.serve(grCtx)
	})

	var err error

	if err = gr.Wait(); err != nil {
		log.Criticalf("server stopped with error: %v", err)
	}

	log.Infof("app: shutting down the server")

	return err
}

func (s *server) serve(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(len(s.servers))

	gr, grCtx := errgroup.WithContext(ctx)

	for _, s := range s.servers {
		s := s

		gr.Go(func() error {
			defer wg.Done()

			return s.Serve(grCtx)
		})
	}

	wg.Wait()

	return gr.Wait()
}
