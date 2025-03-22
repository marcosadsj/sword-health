package factory

import (
	databases "sword-health/internal/database"

	"sword-health/internal/database/mysql"
	"sword-health/internal/database/sqlite"
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
