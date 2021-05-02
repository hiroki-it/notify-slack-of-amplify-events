package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	if os.Getenv("APP_ENV") == "dev" {
		godotenv.Load()
	}
}
