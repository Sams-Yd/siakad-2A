package config

import (
	"log"

	"github.com/joho/godotenv"
)

// InitConfig loads environment variables from .env file
func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
}
