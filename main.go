package main

import (
	"flag"
	"log"
	"main/api"
	"main/client"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// MQTT broker information
const (
	MQTT_BROKER_URL = "mqtts://mqtt.hsl.fi:8883"
	SUB_TOPIC       = "/hfp/v2/journey/ongoing/vp/bus/#"
)

// Get the HTTP listen address from the command line arguments
// example: go run main.go -listen 8080
func getHttpListenAddress() string {
	httpListenAddress := flag.String("listen", "8080", "HTTP listen address")
	flag.Parse()
	return *httpListenAddress
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
	err = client.SubscribeToTopic(mqttClient, SUB_TOPIC, client.HandleBusMessage(db))
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	// Start the HTTP server in a separate goroutine
	httpListenAddress := getHttpListenAddress()
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
