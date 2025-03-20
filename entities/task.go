package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name         string `goorm:"column:name"`
	Description  string `goorm:"column:description;size:2500"`
	TechnicianID uint   `gorm:"<-:create"`
}

func (Task) TableName() string {
	return "tasks"
}
