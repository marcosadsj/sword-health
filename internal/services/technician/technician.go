package technician

import (
	"sword-health/internal/entities"
	"sword-health/internal/repository/technician"
)

type TechnicianService struct {
	repository technician.ITechnicianRepository
}

func (s *TechnicianService) New(repository technician.ITechnicianRepository) {
	s.repository = repository
}

func (s TechnicianService) Create(technician *entities.Technician) error {

	return s.repository.Create(technician)

}

func (s TechnicianService) Read(ids []int) (technicians []*entities.Technician, err error) {

	return s.repository.Read(ids)

}

func (s TechnicianService) Update(technician *entities.Technician) error {

	return s.repository.Update(technician)

}

func (s TechnicianService) Delete(ids []int) error {

	return s.repository.Delete(ids)
}
