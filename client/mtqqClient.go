package client

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Create options for MQTT client
func newMqttClientOptions(clientId string, connectionUrl string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(connectionUrl)
	opts.SetClientID(clientId)
	opts.SetDefaultPublishHandler(handleReceivedMessage)
	opts.OnConnect = handleConnect
	opts.OnConnectionLost = handleConnectionLost
	return opts
}

// Create a new MQTT client with the given options and return it
func ConnectToMqttBroker(clientId string, connectionUrl string) mqtt.Client {
	mqttClientOpts := newMqttClientOptions(clientId, connectionUrl)
	mqttClient := mqtt.NewClient(mqttClientOpts)

	// panics if connection fails
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return mqttClient
}

// Subscribe client to a topic
func SubscribeToTopic(client mqtt.Client, topic string) {
	// panics if subscription fails
	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)
}

// Callback ran on message receive
func handleReceivedMessage(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// Callback whic his ran when the client connects to the broker
func handleConnect(client mqtt.Client) {
	fmt.Println("Connected")
}

// Callback which is ran when the client loses connection to the broker
func handleConnectionLost(client mqtt.Client, err error) {
	fmt.Printf("Connection lost due to: %+v", err)
}
