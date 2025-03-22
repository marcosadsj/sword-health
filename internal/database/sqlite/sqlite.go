package sqlite

import (
	"path/filepath"
	"sword-health-assessment/internal/database/logger"
	"sword-health-assessment/internal/entities"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	db       *gorm.DB
	filepath string
}

func (d *SQLite) New(enviroment string) {

	d.SetEnviroment(enviroment)
}

func (d *SQLite) SetEnviroment(enviroment string) {
	if enviroment == "PRODUCTION" {
		path, err := filepath.Abs("./resources/database.db")

		if err != nil {
			panic("Error to create absolute path")
		}

		d.filepath = path

		return
	}

	path, err := filepath.Abs("../resources/database.db")

	if err != nil {
		panic("Error to create absolute path")
	}

	d.filepath = path

}

func (d *SQLite) Connect() {
	db, err := gorm.Open(sqlite.Open(d.filepath), &gorm.Config{Logger: logger.GetLogger()})

	if err != nil {
		panic("failed to connect database")
	}

	d.db = db

	d.Migrate()

}

func (d SQLite) Close() {

}

func (d SQLite) Repository() {

}

func (d *SQLite) GetDB() *gorm.DB {
	return d.db
}

func (d *SQLite) Migrate() {
	d.db.AutoMigrate(&entities.Manager{}, &entities.Technician{}, &entities.Task{})
}
