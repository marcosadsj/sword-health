package main

import (
	controller "sword-health-assessment/internal/controller"
	sqlite "sword-health-assessment/internal/database/sqlite"
	"sword-health-assessment/internal/notification"

	"github.com/gin-gonic/gin"
)

func main() {

	httpServer := gin.Default()

	database := &sqlite.SQLite{
		Pathname: "/Users/marcosadsj/Documents/Github/sword-health-assessment/test.db",
	}

	database.Connect()

	database.Migrate()

	//buffer size can be ajusted based on demand, to avoid blocking
	notificationChan := make(chan notification.Notification, 10000)

	go notification.Notify(notificationChan)

	controller.Init(httpServer, database.GetDB(), notificationChan)

	httpServer.Run(":8080")

}
