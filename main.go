package main

import (
	"main/client"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// The program will run until it receives an interrupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Create a new MQTT client and subscribe to the topic
	mqttClient := client.ConnetToMqttBroker("abx-trans", "mqtts://mqtt.hsl.fi:8883")
	client.SubscribeToTopic(mqttClient, "/hfp/v2/journey/ongoing/vp/bus/#")

	<-c
	mqttClient.Disconnect(251)
	return
}
