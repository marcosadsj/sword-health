package manager

import (
	"sword-health-assessment/entities"

	"gorm.io/gorm"
)

type IManagerRepository interface {
	Create(*entities.Manager) error
	Read(ids []int) ([]*entities.Manager, error)
	Update(*entities.Manager) error
	Delete(ids []int) error
}

type ManagerRepository struct {
	DB *gorm.DB
}

func (mr ManagerRepository) Create(manager *entities.Manager) error {

	tx := mr.DB.Create(&manager)

	return tx.Error

}

func (mr ManagerRepository) Read(ids []int) (managers []*entities.Manager, err error) {

	tx := mr.DB.Raw("SELECT * FROM `managers` WHERE `managers`.`id` IN ? AND `managers`.`deleted_at` IS NULL", ids).Scan(&managers)

	return managers, tx.Error

}

func (mr ManagerRepository) Update(manager *entities.Manager) error {

	tx := mr.DB.Updates(manager)

	return tx.Error

}

func (mr ManagerRepository) Delete(ids []int) error {

	tx := mr.DB.Delete(&entities.Manager{}, ids)

	return tx.Error

}
