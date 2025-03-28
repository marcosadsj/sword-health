package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ENVS struct {
	SW_ENVIRONMENT string
	DATABASE_TYPE  string
	GIN_PORT       string
}

func GetMYSQLDatabaseType() string {
	return "MYSQL"
}

func GetSQLITEDatabaseType() string {
	return "SQLITE"
}

const PRODUCTION = "PRODUCTION"
const DEVELOPMENT = "DEVELOPMENT"
const TESTING = "TESTING"

func LoadEnv() ENVS {

	SW_ENVIRONMENT := os.Getenv("SW_ENVIRONMENT")

	var ENV_FILENAME string
	var err error

	switch SW_ENVIRONMENT {
	case PRODUCTION:
		ENV_FILENAME = "prod"

	case DEVELOPMENT:
		ENV_FILENAME = "dev"

	default:
		ENV_FILENAME = "local"

		err = godotenv.Load(fmt.Sprintf("../resources/%s.env", ENV_FILENAME))

		log.Println("Local enviroment loaded")

	}

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE_TYPE := os.Getenv("DATABASE_TYPE")
	GIN_PORT := os.Getenv("GIN_PORT")

	if GIN_PORT == "" {
		GIN_PORT = "8080"
	}

	return ENVS{
		SW_ENVIRONMENT: SW_ENVIRONMENT,
		DATABASE_TYPE:  DATABASE_TYPE,
		GIN_PORT:       GIN_PORT,
	}
}
