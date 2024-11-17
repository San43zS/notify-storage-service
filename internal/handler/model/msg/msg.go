package msg

import "time"

type MSG struct {
	Type   string `json:"type"`
	UserId int    `json:"user_id"`
}

type Notify struct {
	Id        string        `json:"id"`
	UserId    int           `json:"user_id"`
	Status    string        `json:"status"`
	Data      string        `json:"data"`
	TTL       time.Duration `json:"ttl"`
	CreatedAt time.Time     `json:"created_at"`
	ExpiredAt time.Time     `json:"expired_at"`
}
