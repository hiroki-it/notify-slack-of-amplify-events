package config

import (
	"os"

	"github.com/joho/godotenv"
)

func ReadEnvFile() {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}
}
