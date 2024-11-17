package app

import (
	"Notify-storage-service/internal/broker"
	"Notify-storage-service/internal/server"
	"Notify-storage-service/internal/service"
	"Notify-storage-service/internal/storage/config"
	"Notify-storage-service/internal/storage/db/psql"
	"context"
	"fmt"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("app")

type App struct {
	server service.Service
	broker broker.Broker
}

func New() (*App, error) {
	storage, err := psql.New(config.NewConfig())
	if err != nil {
		log.Criticalf("failed to create psql storage: %v", err)
		return &App{}, err
	}

	brkr, err := broker.New()
	if err != nil {
		log.Criticalf("failed to create rabbit broker: %v", err)
		return &App{}, err
	}

	srv := service.New(storage, brkr)

	app := &App{
		server: srv,
		broker: brkr,
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	srv, err := server.New(a.server, a.broker)
	if err != nil {
		return err
	}

	if err := srv.Serve(ctx); err != nil {
		return fmt.Errorf("server stopped with error: %w\n", err)
	}

	log.Infof("app: server stopped")
	return nil
}
