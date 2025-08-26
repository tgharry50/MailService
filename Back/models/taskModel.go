package models

import (
	"time"
)

type Task struct {
	ID              int
	UUID            string
	MailUUID        string
	MailServiceUUID string
	WeekDays        string // Specific days by: [1...7]
	EveryDay        int    // <
	EveryWeek       int
	EveryYear       int // ^ By how many
	SpecificDay     time.Time
	Repeat          bool
	Active          bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
