package api

import (
	managerController "sword-health-assessment/internal/controller/manager"
	techController "sword-health-assessment/internal/controller/technician"
	"sword-health-assessment/internal/notification"

	managerRepository "sword-health-assessment/internal/repository/manager"
	managerService "sword-health-assessment/internal/services/manager"

	technicianRepository "sword-health-assessment/internal/repository/technician"
	technicianService "sword-health-assessment/internal/services/technician"

	taskRepository "sword-health-assessment/internal/repository/task"
	taskService "sword-health-assessment/internal/services/task"

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
