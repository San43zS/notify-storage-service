package storageNotification

import (
	notify "Notify-storage-service/internal/model/notification"
	"Notify-storage-service/internal/service/api/notification"
	store "Notify-storage-service/internal/storage/api/notification"
	"context"
)

type service struct {
	storage store.Notification
}

func New(storage notification.StorageNotification) notification.StorageNotification {
	return &service{
		storage: storage,
	}
}

func (s service) GetCurrent(ctx context.Context, id int) ([]notify.Notification, error) {
	return s.storage.GetCurrent(ctx, id)
}

func (s service) GetById(ctx context.Context, Id int) (notify.Notification, error) {
	return s.storage.GetById(ctx, Id)
}

func (s service) Delete(ctx context.Context, ids []int) error {
	return s.storage.Delete(ctx, ids)
}
