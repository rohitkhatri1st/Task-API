package model

import "time"

type Task struct {
	ID          uint      `json:"id" gorm:"primary key;autoIncrement;not null;column:id" validate:"required"`
	Title       string    `json:"title" gorm:"column:title" validate:"required"`
	Description string    `json:"description" gorm:"column:description"`
	Priority    uint      `json:"priority" gorm:"column:priority" validate:"oneof=0 1 2 3 4"`
	DueAt       time.Time `json:"due_at" gorm:"column:due_at"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
