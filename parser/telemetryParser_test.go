package parser

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	domain "github.com/skokcmd/Abax-transport/domain"
	"github.com/stretchr/testify/assert"
)

// Test for ParseMessageToBusTelemetry function
func TestParseMessageToBusTelemetry(t *testing.T) {
	// create a sample topic
	topic := "hfp/v2/journey/ongoing/vp/bus/123/124/125/126/headsign/2020-01-01T00:00:00+00:00/127/"

	// create a sample payload
	payload := domain.BusTelemetryPayload{}

	// Create a sample payload.Vp struct
	payload.Vp = struct {
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
	}{
		Desi:  "ABC123",
		Dir:   "East",
		Oper:  1,
		Veh:   1234,
		Tst:   time.Now(),
		Tsi:   123456,
		Spd:   25.5,
		Hdg:   90,
		Lat:   37.7749,
		Lon:   -122.4194,
		Acc:   0.5,
		Dl:    1,
		Odo:   12345,
		Drst:  1,
		Oday:  "2022-03-16",
		Jrn:   1,
		Line:  1,
		Start: "San Francisco",
		Loc:   "Mission District",
		Stop:  nil,
		Route: "1",
		Occu:  25,
	}

	// call the function being tested
	telemetry, err := ParseMessageToBusTelemetry(topic, payload)
	// set the ID to a known value for testing
	telemetry.ID, _ = uuid.Parse("00000000-0000-0000-0000-000000000000")
	if err != nil {
		t.Fatalf("ParseMessageToBusTelemetry failed: %v", err)
	}

	// verify the result
	expectedID, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	expectedTelemetry := domain.BusTelemetry{
		ID:            expectedID,
		Prefix:        telemetry.Prefix,
		Version:       telemetry.Version,
		JourneyType:   telemetry.JourneyType,
		TemporalType:  telemetry.TemporalType,
		EventType:     telemetry.EventType,
		TransportMode: telemetry.TransportMode,
		OperatorId:    telemetry.OperatorId,
		VehicleNumber: telemetry.VehicleNumber,
		RouteId:       telemetry.RouteId,
		DirectionId:   telemetry.DirectionId,
		Headsign:      telemetry.Headsign,
		StartTime:     telemetry.StartTime,
		NextStop:      telemetry.NextStop,
		Lat:           payload.Vp.Lat,
		Lon:           payload.Vp.Lon,
	}
	if !reflect.DeepEqual(*telemetry, expectedTelemetry) {
		// pretty print the result for debugging
		str, _ := json.MarshalIndent(*telemetry, "", "\t")
		fmt.Println("returned:" + string(str))

		str2, _ := json.MarshalIndent(expectedTelemetry, "", "\t")
		fmt.Println("returned:" + string(str2))
		t.Errorf("ParseMessageToBusTelemetry returned %+v, want %+v", *telemetry, expectedTelemetry)
	}
}

// test for NewBusDTOFromTelemetry function
func TestNewBusDTOFromTelemetry(t *testing.T) {
	// Create a new BusTelemetry struct with needed data
	telemetry := domain.BusTelemetry{
		ID:       uuid.New(),
		RouteId:  123,
		Headsign: "Test HeadSign",
		NextStop: 4,
		Lat:      60.2244,
		Lon:      24.7570,
	}

	userLat := 60.22
	userLon := 24.75
	busDTO := NewBusDTOFromTelemetry(telemetry, userLat, userLon)

	assert.Equal(t, telemetry.ID, busDTO.TelemetryId)
	assert.Equal(t, telemetry.RouteId, busDTO.RouteId)
	assert.Equal(t, telemetry.Headsign, busDTO.HeadSign)
	assert.Equal(t, telemetry.NextStop, busDTO.NextStop)
	assert.Equal(t, telemetry.Lat, busDTO.Lat)
	assert.Equal(t, telemetry.Lon, busDTO.Lon)
}
