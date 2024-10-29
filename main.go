package main

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// Configure
	// TODO: Add env var integration
	brokerAddr := "tcp://JaysDesktop.local:1883"
	timeout := time.Duration(5 * time.Second)
	subject := "arduino/basement/temperature"

	// Connect
	options := mqtt.NewClientOptions()
	options.AddBroker(brokerAddr)
	client := mqtt.NewClient(options)
	connectToken := client.Connect()
	if !connectToken.WaitTimeout(timeout) {
		log.Fatalf("Unable to connect to %s", brokerAddr)
	}
	if connectToken.Error() != nil {
		log.Fatal(connectToken.Error())
	}
	log.Printf("Connected to %s", brokerAddr)
	defer func() {
		client.Disconnect(1)
		log.Printf("Disconnected from %s", brokerAddr)
	}()

	// Subscribe
	// TODO: Query subscriptions from API
	subscribeToken := client.Subscribe(
		subject,
		0,
		func(c mqtt.Client, m mqtt.Message) {
			payload := string(m.Payload())
			// TODO: Save into local memory map
			log.Print(payload)
		},
	)
	if !subscribeToken.WaitTimeout(timeout) {
		log.Fatalf("Unable to subscribe to %s", brokerAddr)
	}
	if subscribeToken.Error() != nil {
		log.Fatal(subscribeToken.Error())
	}
	log.Printf("Subscribed to %s", subject)
	defer func() {
		client.Unsubscribe(subject)
		log.Printf("Unsubscribed from %s", subject)
	}()

	// TODO: Periodically write history into API

	// Wait & print
	// TODO: Improve to add actual Service handling
	time.Sleep(timeout)
}
