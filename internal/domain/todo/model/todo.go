package model

import "time"

type Todo struct {
	ID int

	CreationDate time.Time
	Description  string
	DueDate      time.Time
	Title        string
}
