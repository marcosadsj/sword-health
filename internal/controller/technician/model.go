package technician

import (
	"sword-health/internal/notification"
	taskService "sword-health/internal/services/task"
	technicianService "sword-health/internal/services/technician"
)

type TechnicianController struct {
	service          *technicianService.TechnicianService
	taskService      *taskService.TaskService
	notificationChan chan<- notification.Notification
}
