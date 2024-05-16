package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Struct Message
type Message struct {
	// Title
	Title string `json:"title"`
	// Message
	Message string `json:"message"`
	// IDUser
	IDUser string `json:"id_user"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func send(title, body, id_user string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"notificationtest", // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Create a message
	message := Message{
		Title:   title,
		Message: body,
		IDUser:  id_user,
	}
	// Convert the message to a JSON string
	jsonBody, err := json.Marshal(message)

	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBody,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
	return nil
}

// func main() {
// 	// Send a message
// 	err := send("Title", "Body", "1")
// 	if err != nil {
// 		log.Panicf("Error: %s", err)
// 	}

// }
