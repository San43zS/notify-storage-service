package storage

import "Notify-storage-service/internal/storage/api/notification"

type Storage interface {
	Notification() notification.Notification
}
