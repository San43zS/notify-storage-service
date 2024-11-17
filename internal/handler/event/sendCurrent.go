package event

import (
	message "Notify-storage-service/internal/handler/model/msg/parser/msgParser"
	"Notify-storage-service/internal/handler/model/msg/parser/notifyParser"
	"context"
)

func (h handler) SendCurrent(ctx context.Context, msg []byte) error {
	m, err := message.New().Parse(msg)
	if err != nil {
		return err
	}

	current, err := h.srv.SNotification().GetCurrent(ctx, m.UserId)
	if err != nil {
		return err
	}

	mv, err := notifyParser.New().Unparse(current)
	if err != nil {
		return err
	}

	err = h.srv.BNotification().Send(ctx, mv)
	if err != nil {
		return err
	}

	return nil
}
