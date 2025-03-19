package api

import (
	managerController "sword-health-assessment/controller/manager"
	taskController "sword-health-assessment/controller/task"
	technicianController "sword-health-assessment/controller/technician"

	managerRepository "sword-health-assessment/repository/manager"
	managerService "sword-health-assessment/services/manager"

	technicianRepository "sword-health-assessment/repository/technician"
	technicianService "sword-health-assessment/services/technician"

	taskRepository "sword-health-assessment/repository/task"
	taskService "sword-health-assessment/services/task"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(httpServer *gin.Engine, db *gorm.DB) {

	managerRepository := &managerRepository.ManagerRepository{DB: db}

	techRepository := &technicianRepository.TechnicianRepository{DB: db}

	tRepository := &taskRepository.TaskRepository{DB: db}

	managerService := &managerService.ManagerService{}

	managerService.New(managerRepository)

	techService := &technicianService.TechnicianService{}

	techService.New(techRepository)

	tService := &taskService.TaskService{}

	tService.New(tRepository)

	mc := &managerController.ManagerController{}

	techC := &technicianController.TechnicianController{}

	taskC := &taskController.TaskController{}

	mc.Controller(httpServer, managerService)
	techC.Controller(httpServer, techService)
	taskC.Controller(httpServer, tService)

}
