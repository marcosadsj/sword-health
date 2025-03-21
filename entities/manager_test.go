package entities

import (
	"testing"

	"gorm.io/gorm"
)

func TestEquals(t *testing.T) {
	cases := []struct {
		in   *Manager
		want *Manager
	}{
		{
			in:   &Manager{Model: gorm.Model{ID: uint(1)}, Name: "Marcos 1"},
			want: &Manager{Model: gorm.Model{ID: uint(1)}, Name: "Marcos 1"},
		},
		{
			in:   &Manager{Model: gorm.Model{ID: uint(2)}, Name: "Marcos 2"},
			want: &Manager{Model: gorm.Model{ID: uint(2)}, Name: "Marcos 2"},
		},
	}

	for _, c := range cases {

		if !c.in.Equals(*c.want) {
			t.Errorf("Expected: %v and got: %v", c.in, c.want)
		}
	}
}
