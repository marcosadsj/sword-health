package manager

import (
	"errors"
	"os"
	"sword-health-assessment/database/sqlite"
	"sword-health-assessment/entities"
	"testing"

	"gorm.io/gorm"
)

func BeforeEach() *sqlite.SQLite {
	database := &sqlite.SQLite{
		Pathname: "test_manager.db",
	}

	database.Connect()

	database.Migrate()

	return database
}

func AfterEach() {
	os.Remove("test_manager.db")
}

func TestCreate(t *testing.T) {

	database := BeforeEach()

	managerRepository := ManagerRepository{DB: database.GetDB()}

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

		got := managerRepository.Create(&c.in)

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

	managerRepository := ManagerRepository{DB: database.GetDB()}

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

		_ = managerRepository.Create(c.in)

		got, err := managerRepository.Read([]int{i + 1})

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

	managerRepository := ManagerRepository{DB: database.GetDB()}

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

		_ = managerRepository.Create(c.in)

		err := managerRepository.Update(c.want)

		if err != nil {
			t.Errorf("Error updating manager %v", err)
		}

		got, err := managerRepository.Read([]int{i + 1})

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

	managerRepository := ManagerRepository{DB: database.GetDB()}

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

		_ = managerRepository.Create(c.in)

		got := managerRepository.Delete([]int{int(c.in.ID)})

		if c.want != got {
			t.Errorf("Expected: %v want: %v", c.want, got)
		}
	}

	AfterEach()
}
