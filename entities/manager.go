package entities

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Name string `goorm:"column:name"`
}

func (Manager) TableName() string {
	return "managers"
}

func (m Manager) Equals(manager Manager) bool {

	if m.ID != manager.ID {
		return false
	}

	if m.Name != manager.Name {
		return false
	}

	return true
}
