package task

import (
	"errors"
	"sword-health-assessment/internal/entities"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	Create(*entities.Task) error
	Read(ids []int) ([]*entities.Task, error)
	Update(*entities.Task) error
	Delete(ids []int) error
	FindByTechnicianId(id int) ([]*entities.Task, error)
}

type TaskRepository struct {
	DB *gorm.DB
}

func (mr TaskRepository) Create(task *entities.Task) error {

	if task.Name == "" {
		return errors.New("name cannot be empty")
	}
	tx := mr.DB.Create(&task)

	return tx.Error

}

func (mr TaskRepository) Read(ids []int) (tasks []*entities.Task, err error) {

	tx := mr.DB.Find(&tasks, ids)

	return tasks, tx.Error

}

func (mr TaskRepository) FindByTechnicianId(id int) (tasks []*entities.Task, err error) {

	tx := mr.DB.Find(&tasks, id)

	return tasks, tx.Error

}

func (mr TaskRepository) Update(task *entities.Task) error {

	tx := mr.DB.Updates(task)

	return tx.Error

}

func (mr TaskRepository) Delete(ids []int) error {

	tx := mr.DB.Delete(&entities.Task{}, ids)

	return tx.Error

}
