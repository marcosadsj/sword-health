package api

import (
	managerController "sword-health/internal/controller/manager"
	techController "sword-health/internal/controller/technician"
	"sword-health/internal/notification"

	managerRepository "sword-health/internal/repository/manager"
	managerService "sword-health/internal/services/manager"

	technicianRepository "sword-health/internal/repository/technician"
	technicianService "sword-health/internal/services/technician"

	taskRepository "sword-health/internal/repository/task"
	taskService "sword-health/internal/services/task"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(httpServer *gin.Engine, db *gorm.DB, notificationChan chan<- notification.Notification) {

	managerRepository := &managerRepository.ManagerRepository{DB: db}

	techRepository := &technicianRepository.TechnicianRepository{DB: db}

	taskRepository := &taskRepository.TaskRepository{DB: db}

	managerService := &managerService.ManagerService{}

	managerService.New(managerRepository)

	techService := &technicianService.TechnicianService{}

	techService.New(techRepository)

	taskService := &taskService.TaskService{}

	taskService.New(taskRepository)

	managerController := &managerController.ManagerController{}

	techController := &techController.TechnicianController{}

	managerController.Controller(httpServer, managerService, taskService)
	techController.Controller(httpServer, techService, taskService, notificationChan)
}
