package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusTelemetry struct {
	gorm.Model
	ID            uuid.UUID
	Prefix        string
	Version       string
	JourneyType   string
	TemporalType  string
	EventType     string
	TransportMode string
	OperatorId    int
	VehicleNumber int
	RouteId       int
	DirectionId   int
	Headsign      string
	StartTime     string
	NextStop      int
	GeohashLevel  int
	Geohash       string
	Sid           int
}

// BusTelemetry DTO for API
type BusDTO struct {
	Id                  uuid.UUID
	Number              int
	Lon                 float64
	Lat                 float64
	NextStop            string
	NextStopArrivalTime string
}

func NewBus(number int, lon, lat float64, nextStop, nextStopArrivalTime string) *BusDTO {
	return &BusDTO{
		Id:                  uuid.New(),
		Number:              number,
		Lon:                 lon,
		Lat:                 lat,
		NextStop:            nextStop,
		NextStopArrivalTime: nextStopArrivalTime,
	}
}

// Request to get the nearest buses
type GetBusesRequest struct {
	Id  uuid.UUID
	Lon float64
	Lat float64
}

func NewGetBusesRequest(lon, lat float64) *GetBusesRequest {
	return &GetBusesRequest{
		Id:  uuid.New(),
		Lon: lon,
		Lat: lat,
	}
}

// Response with nearest buses data
type GetBusesResponse struct {
	Id           uuid.UUID
	NearestBuses []BusDTO
}

func NewGetBusesResponse(id uuid.UUID, nearestBuses []BusDTO) *GetBusesResponse {
	return &GetBusesResponse{
		Id:           id,
		NearestBuses: nearestBuses,
	}
}
