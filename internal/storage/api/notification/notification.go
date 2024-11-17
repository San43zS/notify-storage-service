package notification

import (
	msg "Notify-storage-service/internal/handler/model/msg"
	"Notify-storage-service/internal/model/notification"
	"context"
)

type Notification interface {
	GetOld(ctx context.Context, userID int) ([]msg.Notify, error)
	GetCurrent(ctx context.Context, userID int) ([]msg.Notify, error)
	GetById(ctx context.Context, id int) (notification.Notification, error)
	Delete(ctx context.Context, userID int, ids []int) error
}
