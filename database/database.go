package database

import "sword-health-assessment/database/sqlite"

type IDatabase interface {
	Connect()
	Close()
	Repository()
}

type Database struct {
	sqlite sqlite.SQLite
}
