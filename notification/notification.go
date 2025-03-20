package notification

import (
	"fmt"
	"time"
)

// The tech X performed the task Y on date Z”);
type Notification struct {
	TechnicianID int
	TaskID       int
	Date         time.Time
}

func Notify(notificationChan <-chan Notification) {

	for n := range notificationChan {
		fmt.Print(FormatNotification(n))
	}
}

func FormatNotification(n Notification) string {
	return fmt.Sprintf("The tech %d performed the task %d on date %v\n", n.TechnicianID, n.TaskID, n.Date.Format(time.DateTime))
}
