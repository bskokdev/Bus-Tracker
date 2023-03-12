package parser

import (
	"encoding/json"
	"log"
	"main/domain"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

// Converts MQTT message to BusTelemetry struct
func ParseMessageToBusTelemetry(msg mqtt.Message) (*domain.BusTelemetry, error) {
	var telemetry domain.BusTelemetry
	var payload domain.BusTelemetryPayload

	topicParts := strings.Split(msg.Topic(), "/")

	// Unpack payload into BusTelemetryPayload struct
	err := json.Unmarshal(msg.Payload(), &payload)
	if err != nil {
		log.Printf("Failed to unmarshal message: %s\n", err)
	}

	// parse msg topic to get topic values
	telemetry.ID = uuid.New()
	telemetry.Prefix = topicParts[1]
	telemetry.Version = topicParts[2]
	telemetry.JourneyType = topicParts[3]
	telemetry.TemporalType = topicParts[4]
	telemetry.EventType = topicParts[5]
	telemetry.TransportMode = topicParts[6]
	telemetry.OperatorId, _ = strconv.Atoi(topicParts[7])
	telemetry.VehicleNumber, _ = strconv.Atoi(topicParts[8])
	telemetry.RouteId, _ = strconv.Atoi(topicParts[9])
	telemetry.DirectionId, _ = strconv.Atoi(topicParts[10])
	telemetry.Headsign = topicParts[11]
	telemetry.StartTime = topicParts[12]
	telemetry.NextStop, _ = strconv.Atoi(topicParts[13])

	// lat & lon are in payload
	telemetry.Lat, telemetry.Lon = getLatLonFromPayload(payload)

	// return BusTelemetry struct
	return &telemetry, nil
}

// Extracts lat & lon from telemetry payload
func getLatLonFromPayload(payload domain.BusTelemetryPayload) (float64, float64) {
	return payload.Vp.Lat, payload.Vp.Lon
}

// Converts BusTelemetry struct to BusDTO struct
func ParseBusTelemetryToBusDTO(telemetry *domain.BusTelemetry) *domain.BusDTO {
	return &domain.BusDTO{
		Id:       uuid.New(),
		RouteId:  telemetry.RouteId,
		HeadSign: telemetry.Headsign,
		Lat:      telemetry.Lat,
		Lon:      telemetry.Lon,
	}
}
