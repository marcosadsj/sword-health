package entities

import "gorm.io/gorm"

type Technician struct {
	gorm.Model
	Name  string `goorm:"column:name"`
	Tasks []Task `gorm:"foreignkey:TechnicianID;references:ID"`
}

func (Technician) TableName() string {
	return "technicians"
}
