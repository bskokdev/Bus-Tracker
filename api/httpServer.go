package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	address string
	db      *sql.DB
}

type Handler = func(http.ResponseWriter, *http.Request)

// Creates server with address localhost:{address} and database connection
func NewServer(address string, db *sql.DB) *Server {
	return &Server{
		address: fmt.Sprintf(":%s", address),
		db:      db,
	}
}

// Start starts the HTTP server and listens for requests
func (s *Server) Start() error {
	log.Println("Starting HTTP server on port " + s.address)
	// Routes
	http.HandleFunc("/api/v1/buses/nearest", handleGetNearestBuses(s.db))
	return http.ListenAndServe(s.address, nil)
}

func handleGetNearestBuses(db *sql.DB) Handler {
	// TODO: query the database for the closest buses and save in a variable (slice)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List of closest buses"))
	}
}
