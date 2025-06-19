package entity

import "time"

type EventUser struct {
	ID        int64     `json:"id"`
	Status    UserState `json:"status"`
	EventTime time.Time `json:"event_time"`
}
