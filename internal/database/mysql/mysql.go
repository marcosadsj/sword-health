package mysql

import (
	"fmt"
	"os"
	"sword-health/internal/database/logger"
	"sword-health/internal/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	db  *gorm.DB
	dsn string
}

func (d *MySQL) New(enviroment string) {

	d.SetEnviroment(enviroment)
}

func (d *MySQL) SetEnviroment(enviroment string) {

	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")

	d.SetDSN(MYSQL_HOST, MYSQL_PORT, MYSQL_DATABASE, MYSQL_USER, MYSQL_PASSWORD)

}

func (d *MySQL) SetDSN(host, port, databasename, user, password string) {

	d.dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, databasename)
}

func (d *MySQL) Connect() {

	db, err := gorm.Open(mysql.Open(d.dsn), &gorm.Config{Logger: logger.GetLogger()})

	if err != nil {
		panic("failed to connect database")
	}

	d.db = db

	d.Migrate()

}

func (d MySQL) Close() {

}

func (d *MySQL) GetDB() *gorm.DB {
	return d.db
}

func (d *MySQL) Migrate() {
	d.db.AutoMigrate(&entities.Manager{}, &entities.Technician{}, &entities.Task{})
}
