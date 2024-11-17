package notifyParser

import (
	message "Notify-storage-service/internal/handler/model/msg"
	"encoding/json"
	"fmt"
)

type Parser interface {
	Parse([]byte) ([]message.Notify, error)
	Unparse([]message.Notify) ([]byte, error)
}

type parser struct {
}

func New() Parser {
	return &parser{}
}

func (p parser) Parse(m []byte) ([]message.Notify, error) {
	var msg []message.Notify
	test := string(m)
	if err := json.Unmarshal([]byte(test), &msg); err != nil {
		return []message.Notify{}, fmt.Errorf("error while parsing(unmarshal) msg: %w", err)
	}

	return msg, nil
}

func (p parser) Unparse(m []message.Notify) ([]byte, error) {
	arr, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error while parsing(marshal) msg: %w", err)
	}
	return arr, nil
}
