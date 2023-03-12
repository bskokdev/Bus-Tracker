package api

import (
	"encoding/json"
	"fmt"
	"log"
	"main/domain"
	"net/http"

	"gorm.io/gorm"
)

type Server struct {
	address string
	db      *gorm.DB
}

type Handler = func(http.ResponseWriter, *http.Request)

// Creates server with address localhost:{address} and database connection
func NewServer(address string, db *gorm.DB) *Server {
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
	http.HandleFunc("/api/v1/telemetries", handleGetAllTelemetries(s.db))

	return http.ListenAndServe(s.address, nil)
}

// Function to handle GET requests to /api/v1/telemetries
// Returns all telemetries from the database
func handleGetAllTelemetries(db *gorm.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: add pagination
		var telemetries []domain.BusTelemetry
		res := db.Limit(10).Find(&telemetries)
		if res.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(res.Error.Error()))
			return
		}
		jsonData, err := json.Marshal(telemetries)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

func handleGetNearestBuses(db *gorm.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("List of closest buses"))
	}
}
