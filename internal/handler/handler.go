package handler

import (
	"Notify-storage-service/internal/broker"
	"Notify-storage-service/internal/handler/event"
	"Notify-storage-service/internal/service"
	"Notify-storage-service/pkg/msghandler"
)

type Handler struct {
	Event msghandler.MsgResolver
}

func New(srv service.Service, brk broker.Broker) *Handler {
	return &Handler{
		Event: event.New(srv),
	}
}
