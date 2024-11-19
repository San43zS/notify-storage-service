package event

import (
	message "Notify-storage-service/internal/handler/model/msg/parser/msgParser"
	"Notify-storage-service/internal/handler/model/msg/parser/notifyParser"
	"context"
)

func (h handler) SendOld(ctx context.Context, msg []byte) error {
	m, err := message.New().Parse(msg)
	if err != nil {
		return err
	}

	old, err := h.srv.SNotification().GetOld(ctx, m.UserId)
	if err != nil {
		return err
	}

	mv, err := notifyParser.New().Unparse(old)
	if err != nil {
		return err
	}

	err = h.srv.BNotification().Send(ctx, mv)
	if err != nil {
		return err
	}
	return nil
}
