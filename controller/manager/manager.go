package manager

import (
	managerService "sword-health-assessment/services/manager"

	"github.com/gin-gonic/gin"
)

type IController interface {
	Controller(*gin.Engine)
}

type ManagerController struct {
	Service *managerService.ManagerService
}

func (mc ManagerController) Controller(httpServer *gin.Engine) {

	routeGroup := httpServer.Group("/manager")

	routeGroup.POST("/create", Create)
	routeGroup.GET("/findById/:id", Read)
	routeGroup.PUT("/update/:id", Update)
	routeGroup.DELETE("/delete/:id", Delete)

}

func Create(ctx *gin.Context) {

	managerService.Create()

}

func Read(ctx *gin.Context) {

}

func Update(ctx *gin.Context) {

}

func Delete(ctx *gin.Context) {

}
