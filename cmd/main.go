package main

import (
	controller "sword-health/internal/controller"
	databaseFactory "sword-health/internal/database/factory"
	"sword-health/internal/utils"

	"sword-health/internal/notification"

	"github.com/gin-gonic/gin"
)

func main() {

	httpServer := gin.Default()

	envs := utils.LoadEnv()

	database := databaseFactory.Create(envs.DATABASE_TYPE)

	database.New(envs.SW_ENVIRONMENT)
	database.Connect()

	//buffer size can be ajusted based on demand, to avoid blocking
	notificationChan := make(chan notification.Notification, 10000)

	go notification.Notify(notificationChan)

	controller.Init(httpServer, database.GetDB(), notificationChan)

	httpServer.Run(":" + envs.GIN_PORT)

}
