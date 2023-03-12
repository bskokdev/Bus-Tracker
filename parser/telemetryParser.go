package parser

import (
	"main/domain"
	"strconv"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

// Converts MQTT message to BusTelemetry struct
func ParseMessageToBusTelemetry(msg mqtt.Message) (*domain.BusTelemetry, error) {
	var telemetry domain.BusTelemetry
	// parse msg topic to get topic values
	topicParts := strings.Split(msg.Topic(), "/")

	telemetry.ID = uuid.New()
	telemetry.Prefix = topicParts[1]
	telemetry.Version = topicParts[2]
	telemetry.JourneyType = topicParts[3]
	telemetry.TemporalType = topicParts[4]
	telemetry.EventType = topicParts[5]
	telemetry.TransportMode = topicParts[6]
	operatorId, err := strconv.Atoi(topicParts[7])
	if err == nil {
		telemetry.OperatorId = int(operatorId)
	}
	vehicleNumber, err := strconv.Atoi(topicParts[8])
	if err == nil {
		telemetry.VehicleNumber = int(vehicleNumber)
	}
	routeId, err := strconv.Atoi(topicParts[9])
	if err == nil {
		telemetry.RouteId = int(routeId)
	}
	directionId, err := strconv.Atoi(topicParts[10])
	if err == nil {
		telemetry.DirectionId = int(directionId)
	}
	telemetry.Headsign = topicParts[11]
	telemetry.StartTime = topicParts[12]
	nextStop, err := strconv.Atoi(topicParts[13])
	if err == nil {
		telemetry.NextStop = int(nextStop)
	}
	geohashLevel, err := strconv.Atoi(topicParts[14])
	if err == nil {
		telemetry.GeohashLevel = int(geohashLevel)
	}
	telemetry.Geohash = topicParts[15]
	sid, err := strconv.Atoi(topicParts[16])
	if err == nil {
		telemetry.Sid = int(sid)
	}

	// return BusTelemetry struct
	return &telemetry, nil
}

// Converts BusTelemetry struct to BusDTO struct
func ParseBusTelemetryToBusDTO(telemetry *domain.BusTelemetry) *domain.BusDTO {
	// TODO: implement & check if needed
	return nil
}
