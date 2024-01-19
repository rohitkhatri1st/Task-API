package model

import "time"

type Task struct {
	ID          uint       `json:"id" gorm:"primary key;autoIncrement;not null;column:id"`
	Title       string     `json:"title" gorm:"column:title" validate:"required"`
	Description *string    `json:"description,omitempty" gorm:"column:description"`
	Priority    *uint      `json:"priority,omitempty" gorm:"column:priority"`
	DueAt       *time.Time `json:"due_at,omitempty" gorm:"column:due_at"`
	CreatedAt   *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}
