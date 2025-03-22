package entities

import (
	"testing"

	"gorm.io/gorm"
)

func TestTechnician(t *testing.T) {
	cases := []struct {
		in   *Technician
		want *Technician
	}{
		{
			in: &Technician{Model: gorm.Model{ID: uint(1)}, Name: "Marcos 1", Tasks: []Task{
				{Model: gorm.Model{ID: uint(1)}, Name: "Task 1", Description: "Descrição 1", TechnicianID: 1},
			}},
			want: &Technician{Model: gorm.Model{ID: uint(1)}, Name: "Marcos 1", Tasks: []Task{
				{Model: gorm.Model{ID: uint(1)}, Name: "Task 1", Description: "Descrição 1", TechnicianID: 1},
			}},
		},
		{
			in: &Technician{Model: gorm.Model{ID: uint(2)}, Name: "Marcos 2", Tasks: []Task{
				{Model: gorm.Model{ID: uint(2)}, Name: "Task 2", Description: "Descrição 2", TechnicianID: 2},
			}},
			want: &Technician{Model: gorm.Model{ID: uint(2)}, Name: "Marcos 2", Tasks: []Task{
				{Model: gorm.Model{ID: uint(2)}, Name: "Task 2", Description: "Descrição 2", TechnicianID: 2},
			}},
		},
	}

	for _, c := range cases {

		if !c.in.Equals(*c.want) {
			t.Errorf("Expected: %v and got: %v", c.in, c.want)
		}
	}
}
