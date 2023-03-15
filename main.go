package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	env "github.com/joho/godotenv"
	api "github.com/skokcmd/Abax-transport/api"
	client "github.com/skokcmd/Abax-transport/client"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTT broker information
const (
	MQTT_BROKER_URL = "mqtts://mqtt.hsl.fi:8883"
	SUB_TOPIC       = "/hfp/v2/journey/ongoing/vp/bus/#"
)

// Function to get the HTTP port from the environment variables
func getHttpPort() string {
	err := env.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}
	return os.Getenv("HTTP_PORT")
}

func main() {
	// The program will run until it receives an interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// create a new database connection
	// Connection closes when the program exits
	db := api.ConnectToDB()

	// Create a new MQTT client and subscribe to the topic
	mqttClient, err := client.ConnectToMqttBroker("abx-trans", MQTT_BROKER_URL)
	if err != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", err)
	}
	// Subscribe to the topic
	err = client.SubscribeToTopic(
		mqttClient,
		SUB_TOPIC,
		client.HandleBusMessage(db),
	)
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	// Start the HTTP server in a separate goroutine (similiar to a thread)
	httpListenAddress := fmt.Sprintf("0.0.0.0:%s", getHttpPort())
	httpServer := api.NewServer(httpListenAddress, db)
	go httpServer.Start()

	<-c
	cleanUp(mqttClient)
	return
}

// Disconnect the MQTT client and unsubscribe from the topic
func cleanUp(mqttClient mqtt.Client) {
	mqttClient.Disconnect(250)
	mqttClient.Unsubscribe(SUB_TOPIC)
}
