package config

import (
	"fmt"
	"log"
	"mental-health-backend/models" // Import your models package
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitializeDB initializes the database connection
func InitializeDB() (*gorm.DB, error) {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	} else {
		log.Println(".env file loaded successfully")
	}

	// Set default values if environment variables are not set
	dbHost := getEnvWithDefault("DB_HOST", "localhost")
	dbUser := getEnvWithDefault("DB_USER", "postgres")
	dbPassword := getEnvWithDefault("DB_PASSWORD", "")
	dbName := getEnvWithDefault("DB_NAME", "safemind")
	dbPort := getEnvWithDefault("DB_PORT", "5432")

	// Create connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %v", err)
	}

	// Test the connection
	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	DB = db
	log.Println("Database connection successfully established!")
	return db, nil
}

// RunMigrations runs database migrations
func RunMigrations() error {
	if DB == nil {
		log.Fatalln("Database connection not initialized")
	}

	log.Println("Running database migrations...")

	// Add all models for migration
	err := DB.AutoMigrate(
		&models.User{},     
		&models.Therapist{}, 
	)
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
		return err
	}

	log.Println("Database migrations completed successfully!")
	return nil
}

// Helper function to get environment variable with default value
func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
