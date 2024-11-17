package rabbit

import (
	"Notify-storage-service/internal/broker/rabbit"
	"Notify-storage-service/internal/server/launcher"
	"Notify-storage-service/pkg/msghandler"
	"context"
	"fmt"
	"github.com/op/go-logging"
	"golang.org/x/sync/errgroup"
	"sync"
)

var log = logging.MustGetLogger("rabbit")

type server struct {
	handler msghandler.MsgResolver
	broker  rabbit.Service

	config Config
}

func New(broker rabbit.Service, handler msghandler.MsgResolver) launcher.Server {
	return &server{
		handler: handler,
		broker:  broker,

		config: NewCfg(),
	}
}

func (s server) Serve(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(len(s.config.Consumers))

	gr, grCtx := errgroup.WithContext(ctx)

	for _, c := range s.config.Consumers {
		c := c

		gr.Go(func() error {
			defer wg.Done()
			return s.serve(grCtx, c)
		})
	}

	wg.Wait()

	return nil
}

func (s server) serve(ctx context.Context, consumer Consumer) error {
	c := s.broker.Consumer()

	log.Infof("starting rabbit consumer: %s", consumer.QueueName)
	for {
		if err := ctx.Err(); err != nil {
			log.Criticalf("rabbit listener stopped error: %v", err)
			return fmt.Errorf("rabbit listener stopped error: %v", err)
		}

		m, err := c.Consume(ctx, consumer.QueueName)
		if err != nil {
			log.Infof("rabbit listener stopped error: %v", err)
			continue
		}

		go func() {
			err := s.handler.ServeMSG(ctx, m)
			if err != nil {
				log.Criticalf("failed to handle message: %v", err)
				return
			}
		}()
	}
}
