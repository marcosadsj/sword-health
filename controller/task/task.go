package task

import (
	"fmt"
	"net/http"
	"strconv"
	"sword-health-assessment/entities"
	taskService "sword-health-assessment/services/task"

	"github.com/gin-gonic/gin"
)

type IController interface {
	Controller(*gin.Engine)
}

type TaskController struct {
	service *taskService.TaskService
}

func (mc TaskController) Controller(httpServer *gin.Engine, service *taskService.TaskService) {

	mc.service = service

	routeGroup := httpServer.Group("/task")

	routeGroup.POST("/create", mc.create)
	routeGroup.GET("/findById/:id", mc.read)
	routeGroup.PUT("/update", mc.update)
	routeGroup.DELETE("/delete/:id", mc.delete)

}

func (mc TaskController) create(ctx *gin.Context) {

	task := entities.Task{}

	if err := ctx.ShouldBindBodyWithJSON(&task); err == nil {

		mc.service.Create(&task)

		ctx.String(http.StatusOK, "Task created")

		return
	}

	ctx.String(http.StatusInternalServerError, "Task not created")

}

func (mc TaskController) read(ctx *gin.Context) {

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

	tasks, err := mc.service.Read([]int{int(id)})

	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error to find task(s): %v", err))

		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (mc TaskController) update(ctx *gin.Context) {

	task := entities.Task{}

	if err := ctx.ShouldBindBodyWithJSON(&task); err == nil {

		mc.service.Update(&task)

		ctx.String(http.StatusOK, "Task updated")

		return
	}

	ctx.String(http.StatusInternalServerError, "Task not updated")

}

func (mc TaskController) delete(ctx *gin.Context) {

	idString := ctx.Param("id")

	if idString != "" {

		id, err := strconv.ParseInt(idString, 10, 32)

		if err != nil {

			ctx.String(http.StatusBadRequest, fmt.Sprintf("Wrong request param: %v", id))

			return
		}

		err = mc.service.Delete([]int{int(id)})

		if err != nil {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error on delete Task: %v", err))
			return
		}

		ctx.String(http.StatusOK, fmt.Sprintf("Task deleted id: %d", id))

	}

}
