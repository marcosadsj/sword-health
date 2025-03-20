package technician

import (
	"fmt"
	"net/http"
	"strconv"
	"sword-health-assessment/entities"
	"sword-health-assessment/notification"

	"github.com/gin-gonic/gin"
)

func (mc TechnicianController) listTasks(ctx *gin.Context) {

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

	tasks, err := mc.taskService.FindByTechnicianId(int(id))

	if err != nil {
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error to find technician(s): %v", err))

		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (mc TechnicianController) createTask(ctx *gin.Context) {

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

	task := entities.Task{}

	if err := ctx.ShouldBindBodyWithJSON(&task); err == nil {

		mc.taskService.Create(&task)

		mc.notificationChan <- notification.Notification{TechnicianID: int(id), TaskID: int(task.ID), Date: task.CreatedAt}

		ctx.String(http.StatusOK, "Task created")

		return
	}

	ctx.String(http.StatusInternalServerError, "Task not created")
}

func (mc TechnicianController) updateTask(ctx *gin.Context) {

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

	task := entities.Task{}

	if err := ctx.ShouldBindBodyWithJSON(&task); err == nil {

		err = mc.taskService.Update(&task)

		if err != nil {

			ctx.String(http.StatusInternalServerError, "Task not updated: %v", err)

			return
		}

		mc.notificationChan <- notification.Notification{TechnicianID: int(id), TaskID: int(task.ID), Date: task.CreatedAt}

		ctx.String(http.StatusOK, "Task updated")

		return
	}

	ctx.String(http.StatusInternalServerError, "Task not updated")
}
