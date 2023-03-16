package client

import (
	"encoding/json"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/lib/pq"
	domain "github.com/skokcmd/Abax-transport/domain"
	parser "github.com/skokcmd/Abax-transport/parser"
	"gorm.io/gorm"
)

// Create options for MQTT client
// Returns a pointer to the options
func newMqttClientOptions(clientId string, connectionUrl string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(connectionUrl)
	opts.SetClientID(clientId)
	opts.SetKeepAlive(60)
	opts.OnConnect = handleMqttConnect
	opts.OnConnectionLost = handleMqttDisconnect
	return opts
}

// Connects to the MQTT broker at the given URL
// Returns a pointer to the client and an error
func ConnectToMqttBroker(clientId string, connectionUrl string) (mqtt.Client, error) {
	mqttClientOpts := newMqttClientOptions(clientId, connectionUrl)
	mqttClient := mqtt.NewClient(mqttClientOpts)

	// if connection fails, return error
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return mqttClient, nil
}

// Subscribe client to a topic
// Runs the callback function when a message is received
// Returns an error if subscription fails
func SubscribeToTopic(client mqtt.Client, topic string, callback mqtt.MessageHandler) error {
	// return error if subscription fails
	if token := client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	log.Printf("Subscribed to topic: %s\n", topic)
	return nil // no error
}

// Saves the bus telemetry to the database using GORM
// Returns an error if saving fails or nil if saving succeeds
func storeBusTelemetry(db *gorm.DB, telemetry *domain.BusTelemetry) error {
	err := db.Create(telemetry).Error
	return err
}

// Callback which is ran when a message is received from the broker
func HandleBusMessage(db *gorm.DB) func(client mqtt.Client, msg mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		// Unpack payload into BusTelemetryPayload struct
		var payload domain.BusTelemetryPayload
		err := json.Unmarshal(msg.Payload(), &payload)
		if err != nil {
			log.Printf("Failed to unmarshal message: %s\n", err)
		}

		// Parse the message into a BusTelemetry struct
		telemetry, err := parser.ParseMessageToBusTelemetry(msg.Topic(), payload)
		if err != nil {
			log.Printf("Failed to parse message: %s\n", err)
		}

		log.Printf("Received telemetry: %v\n", telemetry)

		// Save the telemetry to the database
		err = storeBusTelemetry(db, telemetry)
		if err != nil {
			log.Printf("Failed to store bus telemetry: %s\n", err)
		}
	}
}

// Callback whic his ran when the client connects to the broker
// Prints a message on connection success
func handleMqttConnect(client mqtt.Client) {
	log.Println("Connected")
}

// Callback which is ran when the client loses connection to the broker
// Prints the error on connection loss
func handleMqttDisconnect(client mqtt.Client, err error) {
	log.Printf("Connection lost due to: %+v\n", err)
}
