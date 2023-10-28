package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/Uwusie/dinozalr2-bot/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	onceConn sync.Once
	onceChan sync.Once
	conn     *amqp.Connection
	channel  *amqp.Channel
	err      error
)

func getRabbitConnection() *amqp.Connection {
	onceConn.Do(func() {
		username := "guest"
		password := "guest"
		host := "localhost"
		port := "5672"
		vhost := "/"

		amqpURI := "amqp://" + username + ":" + password + "@" + host + ":" + port + vhost

		conn, err = amqp.Dial(amqpURI)
		failOnError(err, "Failed to connect to RabbitMQ")
	})
	return conn
}

func GetChannel() *amqp.Channel {
	onceChan.Do(func() {
		connection := getRabbitConnection()
		channel, err = connection.Channel()
		failOnError(err, "Failed to open a channel")
	})
	return channel
}

func ListenForMessages() {
	ch := GetChannel()

	msgs, err := ch.Consume(
		"meow", // queue name
		"",     // consumer tag
		true,   // auto-ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %s", err)
	}

	for msg := range msgs {
		handleMessage(msg)
	}
}

func handleMessage(msg amqp.Delivery) {
	log.Printf("Received message with body: %s", msg.Body)
	handleConfigUpdate(msg)
}

func handleConfigUpdate(msg amqp.Delivery) {

	type MeowConfig struct {
		Count int `json:"count"`
	}

	var meowConfig MeowConfig

	err := json.Unmarshal(msg.Body, &meowConfig)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return
	}

	config.UpdateMeowCount(meowConfig.Count)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func CloseChannel() {
	fmt.Println("Closing channel")
	if channel != nil {
		channel.Close()
	}
}

func CloseConnection() {
	fmt.Println("Closing connection")

	if conn != nil {
		conn.Close()
	}
}
