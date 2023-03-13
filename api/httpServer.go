package api

import (
	"encoding/json"
	"fmt"
	"log"
	"main/domain"
	"net/http"
	"strconv"

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
// Endpoint accepts query parameters page and pageSize
// Returns all telemetries from the database
// example: http://{host}:{port}/api/v1/telemetries?page=1&pageSize=10
func handleGetAllTelemetries(db *gorm.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		pageSize := 10
		page := 1

		// Get the page number from the query parameters
		pageParam := r.URL.Query().Get("page")
		if pageParam != "" {
			page, _ = strconv.Atoi(pageParam)
		}

		// Get the page size from the query parameters
		pageSizeParam := r.URL.Query().Get("pageSize")
		if pageSizeParam != "" {
			pageSize, _ = strconv.Atoi(pageSizeParam)
		}

		offset := (page - 1) * pageSize

		var telemetries []domain.BusTelemetry
		res := db.Limit(pageSize).Offset(offset).Find(&telemetries)
		if res.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(res.Error.Error()))
			return
		}

		// Parse to JSON
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
		// get lon and lat from query parameters
		// lon := r.URL.Query().Get("lon")

		w.Write([]byte("List of closest buses"))
	}
}
