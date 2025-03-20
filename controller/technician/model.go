package technician

import (
	"sword-health-assessment/notification"
	taskService "sword-health-assessment/services/task"
	technicianService "sword-health-assessment/services/technician"
)

type TechnicianController struct {
	service          *technicianService.TechnicianService
	taskService      *taskService.TaskService
	notificationChan chan<- notification.Notification
}
