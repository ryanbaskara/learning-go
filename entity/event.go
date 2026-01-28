package entity

import "time"

type Event struct {
	ID        int64                  `json:"id" db:"id"`
	UserID    string                 `json:"user_id" db:"user_id"`
	Type      string                 `json:"type" db:"type"`
	MetaData  map[string]interface{} `json:"metadata" db:"metadata"`
	CreatedAt time.Time              `json:"created_at" db:"created_at"`
}
