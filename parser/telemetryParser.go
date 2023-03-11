package parser

import (
	"encoding/json"
	"main/domain"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Converts MQTT message to BusTelemetry struct
func ParseMessageToBusTelemetry(msg mqtt.Message) (*domain.BusTelemetry, error) {
	// parse JSON to Message struct
	var mess domain.Message
	err := json.Unmarshal(msg.Payload(), &mess)
	if err != nil {
		return nil, err
	}
	// return BusTelemetry struct from Message struct
	return &mess.BusTelemetry, nil
}

// Converts BusTelemetry struct to BusDTO struct
func ParseBusTelemetryToBusDTO(telemetry *domain.BusTelemetry) *domain.BusDTO {
	// TODO: implement & check if needed
	return nil
}
