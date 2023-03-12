package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Struct to store the entire bus telemetry topic
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
	Lat           float64
	Lon           float64
}

// Struct to store the bus telemetry payload
type BusTelemetryPayload struct {
	Vp struct {
		Desi  string    `json:"desi"`
		Dir   string    `json:"dir"`
		Oper  int       `json:"oper"`
		Veh   int       `json:"veh"`
		Tst   time.Time `json:"tst"`
		Tsi   int       `json:"tsi"`
		Spd   float64   `json:"spd"`
		Hdg   int       `json:"hdg"`
		Lat   float64   `json:"lat"`
		Lon   float64   `json:"long"`
		Acc   float64   `json:"acc"`
		Dl    int       `json:"dl"`
		Odo   int       `json:"odo"`
		Drst  int       `json:"drst"`
		Oday  string    `json:"oday"`
		Jrn   int       `json:"jrn"`
		Line  int       `json:"line"`
		Start string    `json:"start"`
		Loc   string    `json:"loc"`
		Stop  any       `json:"stop"`
		Route string    `json:"route"`
		Occu  int       `json:"occu"`
	} `json:"VP"`
}

// BusTelemetry DTO for API
type BusDTO struct {
	Id       uuid.UUID
	RouteId  int
	HeadSign string
	NextStop string
	Lat      float64
	Lon      float64
}

func NewBusDTO(number int, lon, lat float64, nextStop, nextStopArrivalTime string) *BusDTO {
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
