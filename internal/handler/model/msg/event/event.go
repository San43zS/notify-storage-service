package event

import "time"

const (
	SendOld     string = "send-notify"
	SendCurrent string = "send-current-notify"
)

const (
	TTL     time.Duration = 15 * time.Second
	User_ID int           = 15
)
