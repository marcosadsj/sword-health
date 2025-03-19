package main

import (
	controller "sword-health-assessment/controller"
	sqlite "sword-health-assessment/database/sqlite"

	"github.com/gin-gonic/gin"
)

func main() {

	httpServer := gin.Default()

	database := &sqlite.SQLite{
		Pathname: "/Users/marcosadsj/Documents/Github/sword-health-assessment/test.db",
	}

	database.Connect()

	database.Migrate()

	controller.Init(httpServer, database.GetDB())

	httpServer.Run(":8080")

}
