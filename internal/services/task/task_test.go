package task

import (
	"errors"
	"os"
	databases "sword-health/internal/database"
	databaseFactory "sword-health/internal/database/factory"
	"sword-health/internal/utils"

	"sword-health/internal/entities"
	taskRepository "sword-health/internal/repository/task"

	"testing"

	"gorm.io/gorm"
)

func BeforeEach() databases.IDatabase {

	database := databaseFactory.Create(utils.GetSQLITEDatabaseType())

	database.New(utils.TESTING)

	database.Connect()

	database.Migrate()

	return database
}

func AfterEach() {
	os.Remove("testing.db")
}

func TestCreate(t *testing.T) {

	database := BeforeEach()

	taskRepository := taskRepository.TaskRepository{DB: database.GetDB()}

	taskService := TaskService{repository: taskRepository}

	cases := []struct {
		in   entities.Task
		want error
	}{
		{
			in:   entities.Task{Name: "Marcos", Description: "Descrição", TechnicianID: 1},
			want: nil,
		},
		{
			in:   entities.Task{Name: "", TechnicianID: 1},
			want: errors.New("name cannot be empty"),
		},
	}

	for _, c := range cases {

		got := taskService.Create(&c.in)

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

	taskRepository := taskRepository.TaskRepository{DB: database.GetDB()}

	taskService := TaskService{repository: taskRepository}

	cases := []struct {
		in   *entities.Task
		want *entities.Task
	}{
		{
			in:   &entities.Task{Model: gorm.Model{ID: 1}, Name: "Marcos 1", TechnicianID: 1},
			want: &entities.Task{Model: gorm.Model{ID: 1}, Name: "Marcos 1", TechnicianID: 1},
		},
		{
			in:   &entities.Task{Model: gorm.Model{ID: 2}, Name: "Marcos 2", TechnicianID: 2},
			want: &entities.Task{Model: gorm.Model{ID: 2}, Name: "Marcos 2", TechnicianID: 2},
		},
	}

	for i, c := range cases {

		_ = taskService.Create(c.in)

		got, err := taskService.Read([]int{i + 1})

		if err != nil {
			t.Errorf("Error reading task %v", err)
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

	taskRepository := taskRepository.TaskRepository{DB: database.GetDB()}

	taskService := TaskService{repository: taskRepository}

	cases := []struct {
		in   *entities.Task
		want *entities.Task
	}{
		{
			in:   &entities.Task{Model: gorm.Model{ID: 1}, Name: "Marcos 1", TechnicianID: 1},
			want: &entities.Task{Model: gorm.Model{ID: 1}, Name: "Marcos 11", TechnicianID: 1},
		},
		{
			in:   &entities.Task{Model: gorm.Model{ID: 2}, Name: "Marcos 2", TechnicianID: 2},
			want: &entities.Task{Model: gorm.Model{ID: 2}, Name: "Marcos 22", TechnicianID: 2},
		},
	}

	for i, c := range cases {

		_ = taskService.Create(c.in)

		err := taskService.Update(c.want)

		if err != nil {
			t.Errorf("Error updating task %v", err)
		}

		got, err := taskService.Read([]int{i + 1})

		if err != nil {
			t.Errorf("Error reading task %v", err)
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

	taskRepository := taskRepository.TaskRepository{DB: database.GetDB()}

	taskService := TaskService{repository: taskRepository}

	cases := []struct {
		in   *entities.Task
		want error
	}{
		{
			in:   &entities.Task{Model: gorm.Model{ID: 1}, Name: "Marcos 1", TechnicianID: 1},
			want: nil,
		},
		{
			in:   &entities.Task{Model: gorm.Model{ID: 2}, Name: "Marcos 2", TechnicianID: 2},
			want: nil,
		},
	}

	for _, c := range cases {

		_ = taskService.Create(c.in)

		got := taskService.Delete([]int{int(c.in.ID)})

		if c.want != got {
			t.Errorf("Expected: %v want: %v", c.want, got)
		}
	}

	AfterEach()
}
