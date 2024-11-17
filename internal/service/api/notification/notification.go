package notification

import (
	message "Notify-storage-service/internal/handler/model/msg"
	"Notify-storage-service/internal/model/notification"
	"context"
)

type BrokerNotification interface {
	Add(ctx context.Context, msg message.MSG) error

	Send(ctx context.Context, msg []byte) error
}

type StorageNotification interface {
	GetCurrent(ctx context.Context, id int) ([]message.Notify, error)
	GetById(ctx context.Context, id int) (notification.Notification, error)
	Delete(ctx context.Context, userId int, ids []int) error
	GetOld(ctx context.Context, userId int) ([]message.Notify, error)
}
