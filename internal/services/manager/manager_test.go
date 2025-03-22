package manager

import (
	"errors"
	"os"
	databases "sword-health-assessment/internal/database"
	databaseFactory "sword-health-assessment/internal/database/factory"

	"sword-health-assessment/internal/entities"
	managerRepository "sword-health-assessment/internal/repository/manager"

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
	os.Remove("test_manager.db")
}

func TestCreate(t *testing.T) {

	database := BeforeEach()

	managerRepository := managerRepository.ManagerRepository{DB: database.GetDB()}

	managerService := ManagerService{repository: managerRepository}

	cases := []struct {
		in   entities.Manager
		want error
	}{
		{
			in:   entities.Manager{Name: "Marcos"},
			want: nil,
		},
		{
			in:   entities.Manager{Name: ""},
			want: errors.New("name cannot be empty"),
		},
	}

	for _, c := range cases {

		got := managerService.Create(&c.in)

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

	managerRepository := managerRepository.ManagerRepository{DB: database.GetDB()}

	managerService := ManagerService{repository: managerRepository}

	cases := []struct {
		in   *entities.Manager
		want *entities.Manager
	}{
		{
			in:   &entities.Manager{Model: gorm.Model{ID: 1}, Name: "Marcos 1"},
			want: &entities.Manager{Model: gorm.Model{ID: 1}, Name: "Marcos 1"},
		},
		{
			in:   &entities.Manager{Model: gorm.Model{ID: 2}, Name: "Marcos 2"},
			want: &entities.Manager{Model: gorm.Model{ID: 2}, Name: "Marcos 2"},
		},
	}

	for i, c := range cases {

		_ = managerService.Create(c.in)

		got, err := managerService.Read([]int{i + 1})

		if err != nil {
			t.Errorf("Error reading manager %v", err)
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

	managerRepository := managerRepository.ManagerRepository{DB: database.GetDB()}

	managerService := ManagerService{repository: managerRepository}

	cases := []struct {
		in   *entities.Manager
		want *entities.Manager
	}{
		{
			in:   &entities.Manager{Model: gorm.Model{ID: 1}, Name: "Marcos 1"},
			want: &entities.Manager{Model: gorm.Model{ID: 1}, Name: "Marcos 11"},
		},
		{
			in:   &entities.Manager{Model: gorm.Model{ID: 2}, Name: "Marcos 2"},
			want: &entities.Manager{Model: gorm.Model{ID: 2}, Name: "Marcos 22"},
		},
	}

	for i, c := range cases {

		_ = managerService.Create(c.in)

		err := managerService.Update(c.want)

		if err != nil {
			t.Errorf("Error updating manager %v", err)
		}

		got, err := managerService.Read([]int{i + 1})

		if err != nil {
			t.Errorf("Error reading manager %v", err)
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

	managerRepository := managerRepository.ManagerRepository{DB: database.GetDB()}

	managerService := ManagerService{repository: managerRepository}

	cases := []struct {
		in   *entities.Manager
		want error
	}{
		{
			in:   &entities.Manager{Model: gorm.Model{ID: 1}, Name: "Marcos 1"},
			want: nil,
		},
		{
			in:   &entities.Manager{Model: gorm.Model{ID: 2}, Name: "Marcos 2"},
			want: nil,
		},
	}

	for _, c := range cases {

		_ = managerService.Create(c.in)

		got := managerService.Delete([]int{int(c.in.ID)})

		if c.want != got {
			t.Errorf("Expected: %v want: %v", c.want, got)
		}
	}

	AfterEach()
}
