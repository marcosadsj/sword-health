package manager

import (
	"sword-health-assessment/entities"
	"sword-health-assessment/repository/manager"
)

type ManagerService struct {
	repository manager.IManagerRepository
}

func (s *ManagerService) New(repository manager.IManagerRepository) {
	s.repository = repository
}

func (s ManagerService) Create(manager *entities.Manager) error {

	return s.repository.Create(manager)

}

func (s ManagerService) Read(ids []int) (managers []*entities.Manager, err error) {

	return s.repository.Read(ids)

}

func (s ManagerService) Update(manager *entities.Manager) error {

	return s.repository.Update(manager)

}

func (s ManagerService) Delete(ids []int) error {

	return s.repository.Delete(ids)
}
