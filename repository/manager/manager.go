package manager

import (
	"sword-health-assessment/entities"

	"gorm.io/gorm"
)

type IManagerRepository interface {
	Create(entities.Manager) error
	Read(entities.Manager) (entities.Manager, error)
	Update(entities.Manager) error
	Delete(ids []int) error
}

type ManagerRepository struct {
	DB *gorm.DB
}

func (mr ManagerRepository) Create(manager entities.Manager) error {

	tx := mr.DB.Create(manager)

	return tx.Error

}

func (mr ManagerRepository) Read(manager entities.Manager) (managers entities.Manager, err error) {

	tx := mr.DB.Where(manager).Find(&managers)

	return managers, tx.Error

}

func (mr ManagerRepository) Update(manager entities.Manager) error {

	tx := mr.DB.Save(manager)

	return tx.Error

}

func (mr ManagerRepository) Delete(ids []int) error {

	tx := mr.DB.Delete(&entities.Manager{}, ids)

	return tx.Error

}
