package psql

import (
	"Notify-storage-service/internal/storage"
	"Notify-storage-service/internal/storage/api/notification"
	"Notify-storage-service/internal/storage/config"
	notification2 "Notify-storage-service/internal/storage/db/psql/repo/notification"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	notification notification.Notification
}

func New(config *config.Config) (storage.Storage, error) {
	db, err := sqlx.Connect(config.Driver, config.URL)
	if err != nil {
		return nil, err
	}

	return &Store{
		notification: notification2.New(db),
	}, nil
}

func (s Store) Notification() notification.Notification {
	return s.notification
}
