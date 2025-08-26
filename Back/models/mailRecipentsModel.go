package models

type MailRecipents struct {
	ID       int
	UUID     string
	TaskUUID string
	Recipent string
	Active   bool
}
