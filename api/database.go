package api

import (
	"fmt"
	"log"
	"os"

	domain "github.com/skokcmd/Abax-transport/domain"

	env "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectToDB connects to the database
// Returns a pointer to the database connection
func ConnectToDB() *gorm.DB {
	err := env.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}
	// Database information from environment variables
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Open database connection
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	log.Println("Connected to database")

	// Auto migrate the schema
	db.AutoMigrate(&domain.BusTelemetry{})

	return db
}
