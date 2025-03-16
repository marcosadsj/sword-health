package manager

import (
	"sword-health-assessment/entities"

	"gorm.io/gorm"
)

type IRepository struct {
	db *gorm.DB
}

type IManagerRepository interface {
	Create(entities.Manager) error
	Read(entities.Manager) (entities.Manager, error)
	Update(entities.Manager) error
	Delete(ids []int) error
}

type ManagerRepository struct {
	db *gorm.DB
}

func (mr ManagerRepository) Create(manager entities.Manager) error {

	tx := mr.db.Create(manager)

	return tx.Error

}

func (mr ManagerRepository) Read(manager entities.Manager) (managers entities.Manager, err error) {

	tx := mr.db.Where(manager).Find(&managers)

	return managers, tx.Error

}

func (mr ManagerRepository) Update(manager entities.Manager) error {

	tx := mr.db.Save(manager)

	return tx.Error

}

func (mr ManagerRepository) Delete(ids []int) error {

	tx := mr.db.Delete(&entities.Manager{}, ids)

	return tx.Error

}
