package notification

import (
	"testing"
	"time"
)

func TestFormatNotification(t *testing.T) {

	cases := []struct {
		in   Notification
		want string
	}{
		{
			in:   Notification{TechnicianID: 1, TaskID: 1, Date: time.Date(2025, 03, 19, 23, 59, 59, 0, time.UTC)},
			want: "The tech 1 performed the task 1 on date 2025-03-19 23:59:59\n",
		},
		{
			in:   Notification{TechnicianID: 1, TaskID: 2, Date: time.Date(2025, 04, 20, 00, 00, 00, 0, time.UTC)},
			want: "The tech 1 performed the task 2 on date 2025-04-20 00:00:00\n",
		},
	}

	for _, c := range cases {

		got := FormatNotification(c.in)

		if got != c.want {
			t.Errorf("Expected: %s and got: %s", c.want, got)
		}
	}
}
