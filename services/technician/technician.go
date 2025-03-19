package technician

import (
	"sword-health-assessment/entities"
	"sword-health-assessment/repository/technician"
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

func (s TechnicianService) Read(technician *entities.Technician) (technicians *entities.Technician, err error) {

	return s.repository.Read(technician)

}

func (s TechnicianService) Update(technician *entities.Technician) error {

	return s.repository.Update(technician)

}

func (s TechnicianService) Delete(ids []int) error {

	return s.repository.Delete(ids)
}
