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

func (m Technician) Equals(technician Technician) bool {

	if m.ID != technician.ID {
		return false
	}

	if m.Name != technician.Name {
		return false
	}

	for i, t := range m.Tasks {
		if !t.Equals(technician.Tasks[i]) {
			return false
		}
	}

	return true
}
