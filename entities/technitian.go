package entities

import "gorm.io/gorm"

type Technitian struct {
	gorm.Model
	Id   int64  `goorm:"column:id;primaryKey"`
	Name string `goorm:"column:name"`
}
