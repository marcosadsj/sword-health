package technician

import (
	"errors"
	"os"
	databases "sword-health-assessment/internal/database"
	databaseFactory "sword-health-assessment/internal/database/factory"

	"sword-health-assessment/internal/entities"

	technicianRepository "sword-health-assessment/internal/repository/technician"

	"testing"

	"gorm.io/gorm"
)

func BeforeEach() databases.IDatabase {

	database := databaseFactory.Create("SQLITE")

	database.Connect()

	database.Migrate()

	return database
}

func AfterEach() {
	os.Remove("test_technician.db")
}

func TestCreate(t *testing.T) {

	database := BeforeEach()

	technicianRepository := technicianRepository.TechnicianRepository{DB: database.GetDB()}

	technicianService := TechnicianService{repository: technicianRepository}

	cases := []struct {
		in   entities.Technician
		want error
	}{
		{
			in:   entities.Technician{Name: "Marcos"},
			want: nil,
		},
		{
			in:   entities.Technician{Name: ""},
			want: errors.New("name cannot be empty"),
		},
		{
			in: entities.Technician{Name: "Marcos", Tasks: []entities.Task{
				{Model: gorm.Model{ID: uint(1)}, Name: "Task 1", Description: "Descrição 1", TechnicianID: 1},
			}},
			want: nil,
		},
	}

	for _, c := range cases {

		got := technicianService.Create(&c.in)

		if got != nil && c.want != nil {

			if !errors.Is(errors.Unwrap(got), errors.Unwrap(c.want)) {
				t.Errorf("Expected: %v and got %v", c.want, got)
			}

			continue
		}

		if got != c.want {
			t.Errorf("Expected: %v and got %v", c.want, got)
		}

	}

	AfterEach()
}

func TestRead(t *testing.T) {

	database := BeforeEach()

	technicianRepository := technicianRepository.TechnicianRepository{DB: database.GetDB()}

	technicianService := TechnicianService{repository: technicianRepository}

	cases := []struct {
		in   *entities.Technician
		want *entities.Technician
	}{
		{
			in:   &entities.Technician{Model: gorm.Model{ID: 1}, Name: "Marcos 1"},
			want: &entities.Technician{Model: gorm.Model{ID: 1}, Name: "Marcos 1"},
		},
		{
			in: &entities.Technician{
				Model: gorm.Model{ID: 2},
				Name:  "Marcos 2",
				Tasks: []entities.Task{
					{Model: gorm.Model{ID: uint(1)}, Name: "Task 1", Description: "Descrição 1", TechnicianID: 1},
				}},
			want: &entities.Technician{
				Model: gorm.Model{ID: 2},
				Name:  "Marcos 2",
				Tasks: []entities.Task{
					{Model: gorm.Model{ID: uint(1)}, Name: "Task 1", Description: "Descrição 1", TechnicianID: 1},
				}},
		},
	}

	for i, c := range cases {

		_ = technicianService.Create(c.in)

		got, err := technicianService.Read([]int{i + 1})

		if err != nil {
			t.Errorf("Error reading technician %v", err)
		}

		for _, g := range got {

			if !g.Equals(*c.want) {
				t.Errorf("Expected: %v but got: %v", g, c.want)
			}
		}
	}

	AfterEach()

}

func TestUpdate(t *testing.T) {

	database := BeforeEach()

	technicianRepository := technicianRepository.TechnicianRepository{DB: database.GetDB()}

	technicianService := TechnicianService{repository: technicianRepository}

	cases := []struct {
		in   *entities.Technician
		want *entities.Technician
	}{
		{
			in:   &entities.Technician{Model: gorm.Model{ID: 1}, Name: "Marcos 1"},
			want: &entities.Technician{Model: gorm.Model{ID: 1}, Name: "Marcos 11"},
		},
		{
			in: &entities.Technician{
				Model: gorm.Model{ID: 2},
				Name:  "Marcos 2",
				Tasks: []entities.Task{
					{Model: gorm.Model{ID: uint(1)}, Name: "Task 1", Description: "Descrição 1", TechnicianID: 1},
				}},
			want: &entities.Technician{
				Model: gorm.Model{ID: 2},
				Name:  "Marcos 22",
				Tasks: []entities.Task{
					{Model: gorm.Model{ID: uint(1)}, Name: "Task 1", Description: "Descrição 1", TechnicianID: 1},
				}},
		},
	}

	for i, c := range cases {

		_ = technicianService.Create(c.in)

		err := technicianService.Update(c.want)

		if err != nil {
			t.Errorf("Error updating technician %v", err)
		}

		got, err := technicianService.Read([]int{i + 1})

		if err != nil {
			t.Errorf("Error reading technician %v", err)
		}

		for _, g := range got {

			if !g.Equals(*c.want) {
				t.Errorf("Expected: %v but got: %v", g, c.want)
			}
		}
	}

	AfterEach()
}

func TestDelete(t *testing.T) {

	database := BeforeEach()

	technicianRepository := technicianRepository.TechnicianRepository{DB: database.GetDB()}

	technicianService := TechnicianService{repository: technicianRepository}

	cases := []struct {
		in   *entities.Technician
		want error
	}{
		{
			in:   &entities.Technician{Model: gorm.Model{ID: 1}, Name: "Marcos 1"},
			want: nil,
		},
		{
			in: &entities.Technician{
				Model: gorm.Model{ID: 2},
				Name:  "Marcos 2",
				Tasks: []entities.Task{
					{Model: gorm.Model{ID: uint(1)}, Name: "Task 1", Description: "Descrição 1", TechnicianID: 1},
				}},
			want: nil,
		},
	}

	for _, c := range cases {

		_ = technicianService.Create(c.in)

		got := technicianService.Delete([]int{int(c.in.ID)})

		if c.want != got {
			t.Errorf("Expected: %v want: %v", c.want, got)
		}
	}

	AfterEach()
}
