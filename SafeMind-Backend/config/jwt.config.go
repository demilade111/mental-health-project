package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JWT_SECRET string

func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Read JWT_SECRET
	JWT_SECRET = os.Getenv("JWT_SECRET")
	if JWT_SECRET == "" {
		log.Fatal("JWT_SECRET is not set in the environment variables")
	}
}

func GetJWTSecret() string {
	return JWT_SECRET
}
