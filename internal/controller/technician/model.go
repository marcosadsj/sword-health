package technician

import (
	"sword-health-assessment/internal/notification"
	taskService "sword-health-assessment/internal/services/task"
	technicianService "sword-health-assessment/internal/services/technician"
)

type TechnicianController struct {
	service          *technicianService.TechnicianService
	taskService      *taskService.TaskService
	notificationChan chan<- notification.Notification
}
