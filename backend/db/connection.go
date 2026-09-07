package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// ConnectDB initializes and configures the GORM connection pool.
func ConnectDB() *gorm.DB {
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Extract standard *sql.DB instance to configure connection pool limits
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Unable to retrieve underlying sql.DB instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Database connection established (GORM)")
	return DB
}

// CloseDB gracefully shuts down the GORM database connection pool.
func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Printf("Error retrieving sql.DB during closure: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
			return
		}
		fmt.Println("Database connection closed")
	}
}
