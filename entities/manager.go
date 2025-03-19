package entities

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Name string `goorm:"column:name"`
}

func (Manager) TableName() string {
	return "managers"
}
