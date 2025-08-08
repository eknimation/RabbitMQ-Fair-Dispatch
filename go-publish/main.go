package main

import (
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	msg := "Hello from publisher!"
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	amqpURL := "amqp://rabbituser:rabbitpassword@localhost:5672/"
	conn, err := amqp091.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"liverpool", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}
	log.Printf("Published to [liverpool]: %s", msg)
}
