package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadConfig() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatalf("Error loading .env.%s file", env)
	}
}
