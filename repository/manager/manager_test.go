package manager

import (
	"errors"
	"os"
	"sword-health-assessment/database/sqlite"
	"sword-health-assessment/entities"
	"testing"
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
