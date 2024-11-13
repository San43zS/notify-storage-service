package notification

import (
	msg2 "Notify-storage-service/internal/handler/model/msg"
	"Notify-storage-service/internal/model/notification"
	"context"
)

type BrokerNotification interface {
	Add(ctx context.Context, msg msg2.MSG) error

	GetOld(ctx context.Context) ([]byte, error)

	GetCurrent(ctx context.Context) ([]byte, error)
}

type StorageNotification interface {
	Add(ctx context.Context, notification notification.Notification) error
	Get(ctx context.Context, id int) ([]notification.Notification, error)
	GetById(ctx context.Context, id int) (notification.Notification, error)
	Delete(ctx context.Context, ids []int) error
}
