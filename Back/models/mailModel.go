package models

import "time"

type Mail struct {
	ID        int
	UUID      string
	Topic     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
