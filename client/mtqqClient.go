package client

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Create options for MQTT client
func newTransportMqttOptions(clientId string, connectionUrl string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()

	opts.AddBroker(connectionUrl)
	opts.SetClientID(clientId)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	return opts
}

// Create a new MQTT client with the given options and return it
func ConnetToMqttBroker(clientId string, connectionUrl string) mqtt.Client {
	mqttClientOpts := newTransportMqttOptions(clientId, connectionUrl)
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

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Topic: %s | %s\n", msg.Topic(), msg.Payload())
}

// Callback whic his ran when the client connects to the broker
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

// Callback which is ran when the client loses connection to the broker
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %+v", err)
}
