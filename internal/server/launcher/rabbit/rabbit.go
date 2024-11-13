package rabbit

import (
	"Notify-storage-service/internal/broker/rabbit"
	"Notify-storage-service/internal/server/launcher"
	"Notify-storage-service/pkg/msghandler"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
)

type server struct {
	handler msghandler.MsgResolver
	broker  rabbit.Service
}

func New(broker rabbit.Service, handler msghandler.MsgResolver) launcher.Server {
	return &server{
		handler: handler,
		broker:  broker,
	}
}

func (s server) Serve(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(1)

	gr, grCtx := errgroup.WithContext(ctx)

	gr.Go(func() error {
		defer wg.Done()
		return s.serve(grCtx)
	})

	wg.Wait()

	return nil
}

func (s server) serve(ctx context.Context) error {
	c := s.broker.Consumer()

	for {
		if err := ctx.Err(); err != nil {
			return fmt.Errorf("rabbit listener stopped error: %v", err)
		}

		_, err := c.UConsume(ctx)
		if err != nil {
			//fmt.Println(m)
			// TODO: add logger
			//log.Errorf("failed to consume message error: %v", err)
			continue
		}
		//
		//go func() {
		//	_, err := s.handler.ServeMSG(ctx, m)
		//	if err != nil {
		//		fmt.Errorf("failed to handle message: %v", err)
		//		return
		//	}
		//}()
	}
}
