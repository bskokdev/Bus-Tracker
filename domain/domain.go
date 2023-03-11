package domain

import (
	"github.com/google/uuid"
)

// Bus information
// this would be stored in a database
type Bus struct {
	Id                  uuid.UUID
	Number              int
	Lon                 float64
	Lat                 float64
	NextStop            string
	NextStopArrivalTime string
}

func NewBus(number int, lon, lat float64, nextStop, nextStopArrivalTime string) *Bus {
	return &Bus{
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
	NearestBuses []Bus
}

func NewGetBusesResponse(id uuid.UUID, nearestBuses []Bus) *GetBusesResponse {
	return &GetBusesResponse{
		Id:           id,
		NearestBuses: nearestBuses,
	}
}
