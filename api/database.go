package api

import (
	"fmt"
	"log"
	"main/domain"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database information
// JUST FOR TESTING
// TODO: move this to .env file !!!
const (
	host     = "containers-us-west-23.railway.app"
	port     = 7129
	user     = "postgres"
	password = "Az3TTvjEbxlX3IAhcPJM"
	dbname   = "railway"
)

// ConnectToDB connects to the database and returns a pointer to the database connection
func ConnectToDB() *gorm.DB {
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
