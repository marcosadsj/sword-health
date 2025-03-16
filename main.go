package main

import (
	"sword-health-assessment/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func main() {

	httpServer := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entities.Manager{}, &entities.Technitian{}, &entities.Tasks{})

	db.Create(&entities.Manager{Name: "Marcos"})

	db.Create(&entities.Technitian{Name: "Júnior"})

	db.Create(&entities.Tasks{Name: "Tarefa 1"})

	httpServer.Run(":8080")

}
