package entities

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Id   int64  `goorm:"column:id;primaryKey"`
	Name string `goorm:"column:name"`
}
