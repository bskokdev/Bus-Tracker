package api

import (
	"encoding/json"
	"fmt"
	"log"
	"main/domain"
	"main/parser"
	"main/util"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

// default values for pagination
var (
	pageSize int = 10
	page     int = 1
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
	// Define routes
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

		offset := util.GetPageOffset(page, pageSize)

		var telemetries []domain.BusTelemetry
		res := db.Limit(pageSize).Offset(offset).Find(&telemetries)
		if res.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(res.Error.Error()))
			return
		}

		// Parse telemetries to JSON
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

// Function to handle GET requests to /api/v1/buses/nearest
// Endpoint accepts query parameters lat and lon and page
// pageSize is fixed via default global variable
// Returns the 10 closest buses to the user
func handleGetNearestBuses(db *gorm.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		// get lon and lat from query parameters and cast to f64
		// this would be the user's location
		lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)

		// get page of bus telemetries from the database
		telemetries := []domain.BusTelemetry{}
		offset := util.GetPageOffset(page, pageSize)
		res := db.Limit(5 * pageSize).Offset(offset).Find(&telemetries)
		if res.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(res.Error.Error()))
			return
		}

		// Parse telemetries to BusDTOs
		buses := []domain.BusDTO{}
		for _, telemetry := range telemetries {
			buses = append(buses, parser.NewBusDTOFromTelemetry(telemetry, lat, lon))
		}

		// Parse buses to JSON and return
		jsonData, err := json.Marshal(buses)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
