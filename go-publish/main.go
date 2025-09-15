package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	baseMsg := "Hello from publisher!"
	if len(os.Args) > 1 {
		baseMsg = os.Args[1]
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

	// Publish messages from 1 to 24
	for i := 1; i <= 24; i++ {
		msg := fmt.Sprintf("%s - Message %d", baseMsg, i)

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
			log.Fatalf("Failed to publish message %d: %v", i, err)
		}
		log.Printf("Published to [liverpool]: %s", msg)
	}
}
