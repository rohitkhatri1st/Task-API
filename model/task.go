package model

type Task struct {
	Id   string `json:"id" gorm:"autoIncrement"`
	Name string `json:"name" gorm:"column:name"`
}
