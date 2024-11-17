package event

import (
	"Notify-storage-service/internal/handler/model/msg/event"
	"Notify-storage-service/internal/service"
	"Notify-storage-service/pkg/msghandler"
)

type handler struct {
	srv    service.Service
	router msghandler.MsgHandler
}

func New(srv service.Service) msghandler.MsgHandler {
	endPointParser := func(msg []byte) (string, error) {
		return "", nil
	}

	handler := &handler{
		srv:    srv,
		router: msghandler.New(endPointParser),
	}

	handler.initHandler()

	return handler.router
}

func (h handler) initHandler() {
	h.router.Add(event.SendOld, h.SendOld)
	h.router.Add(event.SendCurrent, h.SendCurrent)
}
