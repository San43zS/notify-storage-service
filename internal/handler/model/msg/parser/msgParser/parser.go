package msgParser

import (
	message "Notify-storage-service/internal/handler/model/msg"
	"encoding/json"
	"fmt"
)

type Parser interface {
	Parse([]byte) (message.MSG, error)
	Unparse(message.MSG) ([]byte, error)
}

type parser struct {
}

func New() Parser {
	return &parser{}
}

func (p parser) Parse(m []byte) (message.MSG, error) {
	var msg message.MSG
	test := string(m)
	if err := json.Unmarshal([]byte(test), &msg); err != nil {
		return message.MSG{}, fmt.Errorf("error while parsing(unmarshal) msg: %w", err)
	}

	return msg, nil
}

func (p parser) Unparse(m message.MSG) ([]byte, error) {
	arr, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("error while parsing(marshal) msg: %w", err)
	}
	return arr, nil
}
