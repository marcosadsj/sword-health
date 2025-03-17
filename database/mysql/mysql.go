package mysql

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MySQL struct {
	db *gorm.DB
}

func (d MySQL) Connect() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	d.db = db
}

func (d MySQL) Close() {

}

func (d MySQL) Repository() {

}
