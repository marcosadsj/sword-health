package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ENVS struct {
	SW_ENVIRONMENT string
	DATABASE_TYPE  string
	GIN_PORT       string
}

func LoadEnv() ENVS {
	SW_ENVIRONMENT := os.Getenv("SW_ENVIRONMENT")
	DATABASE_TYPE := os.Getenv("DATABASE_TYPE")
	GIN_PORT := os.Getenv("GIN_PORT")

	if SW_ENVIRONMENT != "PRODUCTION" {

		err := godotenv.Load("../resources/.env")

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	if GIN_PORT == "" {
		GIN_PORT = "8080"
	}

	return ENVS{
		SW_ENVIRONMENT: SW_ENVIRONMENT,
		DATABASE_TYPE:  DATABASE_TYPE,
		GIN_PORT:       GIN_PORT,
	}
}
