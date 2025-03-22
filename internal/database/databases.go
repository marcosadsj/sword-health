package database

import (
	"sword-health-assessment/internal/database/mysql"
	"sword-health-assessment/internal/database/sqlite"

	"gorm.io/gorm"
)

type IDatabase interface {
	New(enviroment string)
	Connect()
	Close()
	Migrate()
	GetDB() *gorm.DB
	SetEnviroment(string)
}

func Create(databaseType string) IDatabase {
	switch databaseType {
	case "SQLITE":
		return &sqlite.SQLite{}
	case "MYSQL":
		return &mysql.MySQL{}
	}

	return &sqlite.SQLite{}
}
