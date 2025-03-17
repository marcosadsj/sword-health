package api

import (
	managerController "sword-health-assessment/controller/manager"
	managerRepository "sword-health-assessment/repository/manager"
	managerService "sword-health-assessment/services/manager"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(httpServer *gin.Engine, db *gorm.DB) {

	mRepository := &managerRepository.ManagerRepository{
		DB: db,
	}

	mService := &managerService.ManagerService{}

	mService.New(mRepository)

	mc := &managerController.ManagerController{}

	mc.Controller(httpServer)
}
