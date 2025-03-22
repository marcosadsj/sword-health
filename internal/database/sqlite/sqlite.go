package sqlite

import (
	"os"
	"path/filepath"
	"sword-health/internal/database/logger"
	"sword-health/internal/entities"
	"sword-health/internal/utils"

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

	var relativePath string
	var err error

	switch enviroment {
	case utils.PRODUCTION:
		relativePath = "./resources/database.db"
	case utils.DEVELOPMENT:
		relativePath = "../resources/database.db"
	case utils.TESTING:
		relativePath = "./testing.db"

	}

	path, err := filepath.Abs(relativePath)

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
	os.Remove(d.filepath)
}

func (d *SQLite) GetDB() *gorm.DB {
	return d.db
}

func (d *SQLite) Migrate() {
	d.db.AutoMigrate(&entities.Manager{}, &entities.Technician{}, &entities.Task{})
}
