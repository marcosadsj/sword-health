package technician

import (
	"fmt"
	"net/http"
	"strconv"
	"sword-health/internal/entities"
	"sword-health/internal/notification"
	taskService "sword-health/internal/services/task"
	technicianService "sword-health/internal/services/technician"

	"github.com/gin-gonic/gin"
)

func (mc TechnicianController) Controller(
	httpServer *gin.Engine,
	service *technicianService.TechnicianService,
	taskService *taskService.TaskService,
	notificationChan chan<- notification.Notification) {

	mc.service = service
	mc.taskService = taskService
	mc.notificationChan = notificationChan

	routeGroup := httpServer.Group("/technician")

	routeGroup.POST("/create", mc.create)
	routeGroup.GET("/findById/:id", mc.read)
	routeGroup.PUT("/update", mc.update)
	routeGroup.DELETE("/delete/:id", mc.delete)

	routeGroup.GET("/:id/task/list", mc.listTasks)
	routeGroup.POST("/:id/task/create", mc.createTask)
	routeGroup.PATCH("/:id/task/update", mc.updateTask)

}

func (mc TechnicianController) create(ctx *gin.Context) {

	technician := entities.Technician{}

	if err := ctx.ShouldBindBodyWithJSON(&technician); err == nil {

		mc.service.Create(&technician)

		ctx.String(http.StatusOK, "Technician created")

		return
	}

	ctx.String(http.StatusInternalServerError, "Technician not created")

}

func (mc TechnicianController) read(ctx *gin.Context) {

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

	technicians, err := mc.service.Read([]int{int(id)})

	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error to find technician(s): %v", err))

		return
	}

	ctx.JSON(http.StatusOK, technicians)
}

func (mc TechnicianController) update(ctx *gin.Context) {

	technician := entities.Technician{}

	if err := ctx.ShouldBindBodyWithJSON(&technician); err == nil {

		mc.service.Update(&technician)

		ctx.String(http.StatusOK, "Technician updated")

		return
	}

	ctx.String(http.StatusInternalServerError, "Technician not updated")

}

func (mc TechnicianController) delete(ctx *gin.Context) {

	idString := ctx.Param("id")

	if idString != "" {

		id, err := strconv.ParseInt(idString, 10, 32)

		if err != nil {

			ctx.String(http.StatusBadRequest, fmt.Sprintf("Wrong request param: %v", id))

			return
		}

		technicians, err := mc.service.Read([]int{int(id)})

		if err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error to find technician %d: %v", id, err))

			return
		}

		if len(technicians) == 0 {
			ctx.String(http.StatusNotFound, fmt.Sprintf("Technician does not exists: %d", id))

			return
		}

		err = mc.service.Delete([]int{int(id)})

		if err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error on delete Technician: %v", err))
			return
		}

		ctx.String(http.StatusOK, fmt.Sprintf("Technician deleted id: %d", id))

	}

}
