package models

import "time"

type User struct {
	ID        int
	UUID      string
	Firstname string
	Lastname  string
	Username  string
	Password  string
	Email     string
	Phone     string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
