package storageNotification

import (
	message "Notify-storage-service/internal/handler/model/msg"
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

func (s service) GetCurrent(ctx context.Context, id int) ([]message.Notify, error) {
	return s.storage.GetCurrent(ctx, id)
}

func (s service) GetById(ctx context.Context, Id int) (notify.Notification, error) {
	return s.storage.GetById(ctx, Id)
}

func (s service) Delete(ctx context.Context, userId int, ids []int) error {
	return s.storage.Delete(ctx, userId, ids)
}

func (s service) GetOld(ctx context.Context, userId int) ([]message.Notify, error) {
	return s.storage.GetOld(ctx, userId)
}
