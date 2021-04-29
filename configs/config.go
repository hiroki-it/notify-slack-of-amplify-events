package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}
}
