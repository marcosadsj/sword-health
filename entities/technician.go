package entities

import "gorm.io/gorm"

type Technician struct {
	gorm.Model
	Id   int64  `goorm:"column:id;primaryKey"`
	Name string `goorm:"column:name"`
}

func (Technician) TableName() string {
	return "technicians"
}
