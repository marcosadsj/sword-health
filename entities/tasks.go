package entities

import "gorm.io/gorm"

type Tasks struct {
	gorm.Model
	Id          int64  `goorm:"column:id;primaryKey"`
	Name        string `goorm:"column:name"`
	Description string `goorm:"column:description"`
}
