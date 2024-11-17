package service

import (
	"Notify-storage-service/internal/broker"
	"Notify-storage-service/internal/service/api/notification"
	notification2 "Notify-storage-service/internal/service/notification/brokerNotification"
	"Notify-storage-service/internal/storage"
)

type Service interface {
	BNotification() notification.BrokerNotification
	SNotification() notification.StorageNotification
}

type service struct {
	storage      storage.Storage
	notification notification.BrokerNotification
	store        notification.StorageNotification
}

func New(repos storage.Storage, broker broker.Broker) Service {
	return &service{
		storage:      repos,
		notification: notification2.New(broker),
	}
}

func (s *service) BNotification() notification.BrokerNotification {
	return s.notification
}

func (s *service) SNotification() notification.StorageNotification {
	return s.store
}
