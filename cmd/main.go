package main

import (
	"os"
	controller "sword-health-assessment/internal/controller"
	databases "sword-health-assessment/internal/database"

	"sword-health-assessment/internal/notification"

	"github.com/gin-gonic/gin"
)

func main() {

	SW_ENVIRONMENT := os.Getenv("SW_ENVIRONMENT")
	DATABASE_TYPE := os.Getenv("DATABASE_TYPE")
	GIN_PORT := os.Getenv("GIN_PORT")

	httpServer := gin.Default()

	database := databases.Create(DATABASE_TYPE)
	database.New(SW_ENVIRONMENT)
	database.Connect()

	//buffer size can be ajusted based on demand, to avoid blocking
	notificationChan := make(chan notification.Notification, 10000)

	go notification.Notify(notificationChan)

	controller.Init(httpServer, database.GetDB(), notificationChan)

	httpServer.Run(":" + GIN_PORT)

}
