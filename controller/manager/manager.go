package manager

import (
	"fmt"
	"net/http"
	"strconv"
	"sword-health-assessment/entities"
	managerService "sword-health-assessment/services/manager"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	service *managerService.ManagerService
}

func (mc ManagerController) Controller(httpServer *gin.Engine, service *managerService.ManagerService) {

	mc.service = service

	routeGroup := httpServer.Group("/manager")

	routeGroup.POST("/create", mc.create)
	routeGroup.GET("/findById/:id", mc.read)
	routeGroup.PUT("/update", mc.update)
	routeGroup.DELETE("/delete/:id", mc.delete)

}

func (mc ManagerController) create(ctx *gin.Context) {

	manager := entities.Manager{}

	if err := ctx.ShouldBindBodyWithJSON(&manager); err == nil {

		mc.service.Create(&manager)

		ctx.String(http.StatusOK, "Manager created")

		return
	}

	ctx.String(http.StatusInternalServerError, "Manager not created")

}

func (mc ManagerController) read(ctx *gin.Context) {

	idString := ctx.Param("id")

	if idString == "" {

		ctx.String(http.StatusBadRequest, "Empty request param")

		return
	}

	id, err := strconv.ParseInt(idString, 10, 32)

	if err != nil {

		ctx.String(http.StatusBadRequest, fmt.Sprintf("Wrong request param: %v", id))

		return
	}

	managers, err := mc.service.Read([]int{int(id)})

	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error to find manager(s): %v", err))

		return
	}

	ctx.JSON(http.StatusOK, managers)
}

func (mc ManagerController) update(ctx *gin.Context) {

	manager := entities.Manager{}

	if err := ctx.ShouldBindBodyWithJSON(&manager); err == nil {

		mc.service.Update(&manager)

		ctx.String(http.StatusOK, "Manager updated")

		return
	}

	ctx.String(http.StatusInternalServerError, "Manager not updated")

}

func (mc ManagerController) delete(ctx *gin.Context) {

	idString := ctx.Param("id")

	if idString != "" {

		id, err := strconv.ParseInt(idString, 10, 32)

		if err != nil {

			ctx.String(http.StatusBadRequest, fmt.Sprintf("Wrong request param: %v", id))

			return
		}

		err = mc.service.Delete([]int{int(id)})

		if err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error on delete Manager: %v", err))
			return
		}

		ctx.String(http.StatusOK, fmt.Sprintf("Manager deleted id: %d", id))

	}

}
