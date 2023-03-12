package client

import (
	"log"
	"main/domain"
	"main/parser"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

// Create options for MQTT client
func newMqttClientOptions(clientId string, connectionUrl string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(connectionUrl)
	opts.SetClientID(clientId)
	opts.OnConnect = handleMqttConnect
	opts.OnConnectionLost = handleMqttDisconnect
	return opts
}

// Create a new MQTT client with the given options and return it
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
func SubscribeToTopic(client mqtt.Client, topic string, callback mqtt.MessageHandler) error {
	// return error if subscription fails
	if token := client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	log.Printf("Subscribed to topic: %s\n", topic)
	return nil // no error
}

// Saves the bus telemetry to the database using GORM
func storeBusTelemetry(db *gorm.DB, telemetry *domain.BusTelemetry) error {
	err := db.Create(telemetry).Error
	return err
}

// Callback which is ran when a message is received from the broker
func HandleBusMessage(db *gorm.DB) func(client mqtt.Client, msg mqtt.Message) {
	return func(client mqtt.Client, msg mqtt.Message) {
		telemetry, err := parser.ParseMessageToBusTelemetry(msg)
		if err != nil {
			log.Printf("Failed to parse message: %s\n", err)
		}
		err = storeBusTelemetry(db, telemetry)
		if err != nil {
			log.Printf("Failed to store bus telemetry: %s\n", err)
		}
		log.Printf("Received telemetry: %v,\n", telemetry)
	}
}

// Callback whic his ran when the client connects to the broker
func handleMqttConnect(client mqtt.Client) {
	log.Println("Connected")
}

// Callback which is ran when the client loses connection to the broker
func handleMqttDisconnect(client mqtt.Client, err error) {
	log.Printf("Connection lost due to: %+v\n", err)
}
