package parser

import (
	"encoding/json"
	"main/domain"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Converts MQTT message to BusTelemetry struct
func ParseMessageToBusTelemetry(msg mqtt.Message) (*domain.BusTelemetry, error) {
	var telemetry domain.BusTelemetry
	err := json.Unmarshal(msg.Payload(), &telemetry)
	if err != nil {
		return nil, err
	}
	return &telemetry, nil
}

// Converts BusTelemetry struct to BusDTO struct
func ParseBusTelemetryToBusDTO(telemetry *domain.BusTelemetry) *domain.BusDTO {
	// TODO: implement & check if needed
	return nil
}
