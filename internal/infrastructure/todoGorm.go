package infrastructure

import "time"

type todoGorm struct {
	ID           int       `gorm:"primary_key"`
	Title        string    `gorm:"Column:title;size:255"`
	Description  *string   `gorm:"Column:description;size:255"`
	CreationDate time.Time `gorm:"Column:created_at"`
	DueDate      time.Time `gorm:"Column:due_date"`
}

func (todoGorm) TableName() string {
	return "Todo"
}
