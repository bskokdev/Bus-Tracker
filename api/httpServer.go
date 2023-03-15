package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"

	domain "github.com/skokcmd/Abax-transport/domain"
	parser "github.com/skokcmd/Abax-transport/parser"
	util "github.com/skokcmd/Abax-transport/util"

	"gorm.io/gorm"
)

// Struct representing the HTTP Server
// Contains address and database connection
type Server struct {
	address string
	db      *gorm.DB
}

// Creates server with address localhost:{address} and database connection
func NewServer(address string, db *gorm.DB) *Server {
	return &Server{
		address: address,
		db:      db,
	}
}

// Handler is a function that handles HTTP requests
type Handler = func(http.ResponseWriter, *http.Request)

// Start starts the HTTP server and listens for requests
func (s *Server) Start() error {
	log.Println("Starting HTTP server on address " + s.address)
	// Define routes
	http.HandleFunc("/api/v1/buses/nearest", handleGetNearestBuses(s.db))
	http.HandleFunc("/api/v1/telemetries", handleGetAllTelemetries(s.db))

	return http.ListenAndServe(s.address, nil)
}

// Function to get a page of given size of bus telemetries from the database
func getTelemetriesForPage(db *gorm.DB, page, pageSize int, orderField string, telemetries *[]domain.BusTelemetry) error {
	offset := util.GetPageOffset(page, pageSize)
	res := db.Limit(pageSize).Offset(offset).Order(orderField).Find(&telemetries)
	return res.Error // returns nil if no error
}

// ----------------------------
// ROUTE HANDLERS
// ----------------------------

// Function to handle GET requests to /api/v1/telemetries
// Endpoint accepts query parameters page and pageSize
// Returns all telemetries from the database
// example: http://{host}:{port}/api/v1/telemetries?page=1&pageSize=10
func handleGetAllTelemetries(db *gorm.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /api/v1/telemetries")
		// default page size & page
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

		telemetries := make([]domain.BusTelemetry, pageSize)
		err := getTelemetriesForPage(db, page, pageSize, "created_at", &telemetries)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error getting telemetries: %v", err)
			return
		}

		// Parse telemetries to JSON
		jsonData, err := json.Marshal(telemetries)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error while parsing telemetries to JSON: %s", err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

// Function to handle GET requests to /api/v1/buses/nearest
// Endpoint accepts query parameters lat and lon and page
// Returns the 20 nearest buses away from the users' location
func handleGetNearestBuses(db *gorm.DB) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET /api/v1/buses/nearest")
		// default page size & page
		pageSize := 20
		page := 1
		// get lon and lat from query parameters and cast to f64
		// THIS WOULD BE THE USER'S LOCATION
		lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)

		// get page of bus telemetries from the database
		telemetries := make([]domain.BusTelemetry, pageSize)
		err := getTelemetriesForPage(db, page, pageSize, "created_at", &telemetries)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error getting telemetries: %v", err)
			return
		}

		// Parse telemetries to BusDTOs
		buses := make([]domain.BusDTO, 0, pageSize)
		for _, telemetry := range telemetries {
			buses = append(buses, parser.NewBusDTOFromTelemetry(telemetry, lat, lon))
		}

		// sort buses by distance from user
		sort.Slice(buses, func(i, j int) bool {
			return buses[i].DistanceFromUser < buses[j].DistanceFromUser
		})

		// Parse buses to JSON and return
		jsonData, err := json.Marshal(buses)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("Error while parsing to JSON: %s", err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
