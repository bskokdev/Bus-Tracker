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
	// Database information from environment variables
	var (
		host     = os.Getenv("DB_HOST")
		port     = 5432
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}
	// connection string for GORM PostgreSQL driver
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, password, dbname, port,
	)

	// Open database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	log.Println("Connected to database")

	// Auto migrate the schema
	db.AutoMigrate(&domain.BusTelemetry{})

	return db
}
