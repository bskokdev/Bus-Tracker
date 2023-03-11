package api

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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

func ConnectToDB() *sql.DB {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Connected to database")

	return db
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
	log.Println("Closed database connection")
}
