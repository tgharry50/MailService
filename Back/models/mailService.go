package models

import "time"

type MailService struct {
	ID        int
	UUID      string
	Name      string
	Address   string
	Port      string
	Login     string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
