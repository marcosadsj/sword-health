package manager

import (
	"net/http"
	"sword-health-assessment/entities"
	managerService "sword-health-assessment/services/manager"

	"github.com/gin-gonic/gin"
)

type IController interface {
	Controller(*gin.Engine)
}

type ManagerController struct {
	service *managerService.ManagerService
}

func (mc ManagerController) Controller(httpServer *gin.Engine, service *managerService.ManagerService) {

	mc.service = service

	routeGroup := httpServer.Group("/manager")

	routeGroup.POST("/create", mc.Create)
	routeGroup.GET("/findById/:id", mc.Read)
	routeGroup.PUT("/update/:id", mc.Update)
	routeGroup.DELETE("/delete/:id", mc.Delete)

}

func (mc ManagerController) Create(ctx *gin.Context) {

	manager := entities.Manager{}

	if err := ctx.ShouldBindBodyWithJSON(&manager); err == nil {

		mc.service.Create(&manager)

		ctx.String(http.StatusOK, "Manager created")

		return
	}

	ctx.String(http.StatusInternalServerError, "Manager not created")

}

func (mc ManagerController) Read(ctx *gin.Context) {

}

func (mc ManagerController) Update(ctx *gin.Context) {

}

func (mc ManagerController) Delete(ctx *gin.Context) {

}
