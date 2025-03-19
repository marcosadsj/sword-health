package sqlite

import (
	"sword-health-assessment/database"
	"sword-health-assessment/entities"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	db       *gorm.DB
	Pathname string
}

func (d *SQLite) Connect() {
	db, err := gorm.Open(sqlite.Open(d.Pathname), &gorm.Config{Logger: database.GetLogger()})

	if err != nil {
		panic("failed to connect database")
	}

	d.db = db
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
