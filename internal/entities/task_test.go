package entities

import (
	"testing"

	"gorm.io/gorm"
)

func TestTask(t *testing.T) {
	cases := []struct {
		in   *Task
		want *Task
	}{
		{
			in:   &Task{Model: gorm.Model{ID: uint(1)}, Name: "Marcos 1", Description: "Descrição 1", TechnicianID: 1},
			want: &Task{Model: gorm.Model{ID: uint(1)}, Name: "Marcos 1", Description: "Descrição 1", TechnicianID: 1},
		},
		{
			in:   &Task{Model: gorm.Model{ID: uint(2)}, Name: "Marcos 2", Description: "Descrição 2", TechnicianID: 2},
			want: &Task{Model: gorm.Model{ID: uint(2)}, Name: "Marcos 2", Description: "Descrição 2", TechnicianID: 2},
		},
	}

	for _, c := range cases {

		if !c.in.Equals(*c.want) {
			t.Errorf("Expected: %v and got: %v", c.in, c.want)
		}
	}
}
