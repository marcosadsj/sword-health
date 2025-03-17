package sqlite

import (
	"sword-health-assessment/entities"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	db       *gorm.DB
	Pathname string
}

func (d *SQLite) Connect() {
	db, err := gorm.Open(sqlite.Open(d.Pathname), &gorm.Config{})

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
	d.db.AutoMigrate(&entities.Manager{}, &entities.Technitian{}, &entities.Tasks{})

	d.db.Create(&entities.Manager{Name: "Marcos"})

	d.db.Create(&entities.Technitian{Name: "Júnior"})

	d.db.Create(&entities.Tasks{Name: "Tarefa 1"})
}
