package notification

import (
	"Notify-storage-service/internal/model/notification"
	"context"
)

type Notification interface {
	Add(ctx context.Context, notification notification.Notification) error
	Get(ctx context.Context, id int) ([]notification.Notification, error)
	GetById(ctx context.Context, id int) (notification.Notification, error)
	Delete(ctx context.Context, ids []int) error
}
