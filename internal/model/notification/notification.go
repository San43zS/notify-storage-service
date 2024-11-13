package notification

import "time"

type Notification struct {
	UserId    int
	Data      string
	CreatedAt time.Time
}
