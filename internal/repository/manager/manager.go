package manager

import (
	"errors"
	"sword-health-assessment/internal/entities"

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

	if manager.Name == "" {
		return errors.New("name cannot be empty")
	}

	tx := mr.DB.Create(&manager)

	return tx.Error

}

func (mr ManagerRepository) Read(ids []int) (managers []*entities.Manager, err error) {

	tx := mr.DB.Find(&managers, ids)

	return managers, tx.Error

}

func (mr ManagerRepository) Update(manager *entities.Manager) error {

	tx := mr.DB.UpdateColumns(manager)

	return tx.Error

}

func (mr ManagerRepository) Delete(ids []int) error {

	tx := mr.DB.Delete(&entities.Manager{}, ids)

	return tx.Error

}
