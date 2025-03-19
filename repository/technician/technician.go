package technician

import (
	"sword-health-assessment/entities"

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

	tx := mr.DB.Create(&technician)

	return tx.Error

}

func (mr TechnicianRepository) Read(ids []int) (technicians []*entities.Technician, err error) {

	tx := mr.DB.Raw("SELECT * FROM `technicians` WHERE `technicians`.`id` IN ? AND `technicians`.`deleted_at` IS NULL", ids).Scan(&technicians)

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
