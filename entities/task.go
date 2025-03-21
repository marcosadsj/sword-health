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

func (t Task) Equals(task Task) bool {

	if t.ID != task.ID {
		return false
	}

	if t.Name != task.Name {
		return false
	}

	if t.Description != task.Description {
		return false
	}

	if t.TechnicianID != task.TechnicianID {
		return false
	}

	return true
}
