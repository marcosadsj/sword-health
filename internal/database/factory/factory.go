package factory

import (
	databases "sword-health-assessment/internal/database"

	"sword-health-assessment/internal/database/mysql"
	"sword-health-assessment/internal/database/sqlite"
)

func Create(databaseType string) databases.IDatabase {
	switch databaseType {
	case "SQLITE":
		return &sqlite.SQLite{}
	case "MYSQL":
		return &mysql.MySQL{}
	}

	return &sqlite.SQLite{}
}
