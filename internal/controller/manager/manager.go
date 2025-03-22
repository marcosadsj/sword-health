package manager

import (
	"fmt"
	"net/http"
	"strconv"
	"sword-health/internal/entities"
	managerService "sword-health/internal/services/manager"
	taskService "sword-health/internal/services/task"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	service     *managerService.ManagerService
	taskService *taskService.TaskService
}

func (mc ManagerController) Controller(
	httpServer *gin.Engine,
	service *managerService.ManagerService,
	taskService *taskService.TaskService) {

	mc.service = service
	mc.taskService = taskService

	routeGroup := httpServer.Group("/manager")

	routeGroup.POST("/create", mc.create)
	routeGroup.GET("/findById/:id", mc.read)
	routeGroup.PUT("/update", mc.update)
	routeGroup.DELETE("/delete/:id", mc.delete)

	routeGroup.GET("/:id/task/listByTechnicianId/:technicianId", mc.listByTechnicianId)
	routeGroup.DELETE("/:id/task/deleteByTaskId/:taskId", mc.deleteByTaskId)

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

func (mc ManagerController) listByTechnicianId(ctx *gin.Context) {

	idString := ctx.Param("id")

	technicianIdString := ctx.Param("technicianId")

	if idString == "" || technicianIdString == "" {

		ctx.String(http.StatusBadRequest, fmt.Sprintf("Empty request param. id: %s technicianId: %s", idString, technicianIdString))

		return
	}

	id, err := strconv.ParseInt(idString, 10, 32)

	if err != nil {

		ctx.String(http.StatusBadRequest, fmt.Sprintf("Wrong request param: %v", err))

		return
	}

	technicianId, err := strconv.ParseInt(technicianIdString, 10, 32)

	if err != nil {

		ctx.String(http.StatusBadRequest, fmt.Sprintf("Wrong request param: %v", err))

		return
	}

	manager, err := mc.service.Read([]int{int(id)})

	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error to find manager %d: %v", id, err))

		return
	}

	if len(manager) == 0 {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("Error to find manager: %d", id))

		return
	}

	tasks, err := mc.taskService.FindByTechnicianId(int(technicianId))

	if err != nil {

		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error to find tasks(s) of technician %d: %v", technicianId, err))

		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (mc ManagerController) deleteByTaskId(ctx *gin.Context) {

	idString := ctx.Param("id")

	taskIdString := ctx.Param("taskId")

	if idString == "" || taskIdString == "" {

		ctx.String(http.StatusBadRequest, fmt.Sprintf("Empty request param. id: %s taskId: %s", idString, taskIdString))

		return
	}

	id, err := strconv.ParseInt(idString, 10, 32)

	if err != nil {

		ctx.String(http.StatusBadRequest, fmt.Sprintf("Wrong request param: %v", err))

		return
	}

	taskId, err := strconv.ParseInt(taskIdString, 10, 32)

	if err != nil {

		ctx.String(http.StatusBadRequest, fmt.Sprintf("Wrong request param: %v", err))

		return
	}

	manager, err := mc.service.Read([]int{int(id)})

	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error to find manager %d: %v", id, err))

		return
	}

	if len(manager) == 0 {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("Error to find manager: %d", id))

		return
	}

	tasks, err := mc.taskService.Read([]int{int(taskId)})

	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error to find taks %d: %v", taskId, err))

		return
	}

	if len(tasks) == 0 {
		ctx.String(http.StatusNotFound, fmt.Sprintf("Task does not exists: %d", taskId))

		return
	}

	err = mc.taskService.Delete([]int{int(taskId)})

	if err != nil {

		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error deleting task %d: %v", taskId, err))

		return
	}

	ctx.String(http.StatusOK, "Task deleted")
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

		managers, err := mc.service.Read([]int{int(id)})

		if err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error to find manager %d: %v", id, err))

			return
		}

		if len(managers) == 0 {
			ctx.String(http.StatusNotFound, fmt.Sprintf("manager does not exists: %d", id))

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
