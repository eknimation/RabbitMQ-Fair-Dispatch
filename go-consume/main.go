package main

import (
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
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

	// Set Fair Dispatch - only send one message at a time to each consumer
	err = ch.Qos(
		1,     // prefetch count - ส่งข้อความครั้งละ 1 ข้อความต่อ consumer
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		log.Fatalf("Failed to set QoS: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Println("Consuming from queue 'liverpool'. Waiting for messages...")
	for d := range msgs {
		log.Printf("Received on [liverpool]: %s", string(d.Body))
		for i := 1; i <= 60; i++ {
			log.Printf("Loop: %d - %s", i, string(d.Body))
			time.Sleep(1 * time.Second)
		}
		log.Println("Processing complete.")
	}
}
