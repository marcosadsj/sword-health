package task

import (
	"sword-health-assessment/internal/entities"
	"sword-health-assessment/internal/repository/task"
)

type TaskService struct {
	repository task.ITaskRepository
}

func (s *TaskService) New(repository task.ITaskRepository) {
	s.repository = repository
}

func (s TaskService) Create(task *entities.Task) error {

	return s.repository.Create(task)

}

func (s TaskService) Read(ids []int) (tasks []*entities.Task, err error) {

	return s.repository.Read(ids)

}

func (s TaskService) FindByTechnicianId(id int) (tasks []*entities.Task, err error) {

	return s.repository.FindByTechnicianId(id)

}

func (s TaskService) Update(task *entities.Task) error {

	return s.repository.Update(task)

}

func (s TaskService) Delete(ids []int) error {

	return s.repository.Delete(ids)
}
