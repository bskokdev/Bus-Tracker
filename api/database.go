package api

import (
	"fmt"
	"log"

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

func ConnectToDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	// replace this with ORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	log.Println("Connected to database")

	return db
}
