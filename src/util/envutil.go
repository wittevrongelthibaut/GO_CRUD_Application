package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}