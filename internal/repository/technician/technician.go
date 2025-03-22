package technician

import (
	"errors"
	"sword-health-assessment/internal/entities"

	"gorm.io/gorm"
)

type ITechnicianRepository interface {
	Create(*entities.Technician) error
	Read(ids []int) ([]*entities.Technician, error)
	Update(*entities.Technician) error
	Delete(ids []int) error
}

type TechnicianRepository struct {
	DB *gorm.DB
}

func (mr TechnicianRepository) Create(technician *entities.Technician) error {

	if technician.Name == "" {
		return errors.New("name cannot be empty")
	}
	tx := mr.DB.Create(&technician)

	return tx.Error

}

func (mr TechnicianRepository) Read(ids []int) (technicians []*entities.Technician, err error) {

	tx := mr.DB.Find(&technicians, ids)

	return technicians, tx.Error

}

func (mr TechnicianRepository) Update(technician *entities.Technician) error {

	tx := mr.DB.Updates(technician)

	return tx.Error

}

func (mr TechnicianRepository) Delete(ids []int) error {

	tx := mr.DB.Delete(&entities.Technician{}, ids)

	return tx.Error

}
