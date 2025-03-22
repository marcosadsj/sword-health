package database

import (
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
