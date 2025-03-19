package entities

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Id           int64  `goorm:"column:id;primaryKey"`
	Name         string `goorm:"column:name"`
	Description  string `goorm:"column:description"`
	TechnicianID uint
}

func (Task) TableName() string {
	return "tasks"
}
