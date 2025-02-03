package messaging

import (
	"log"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel

func InitRabbitMQ() {
	var err error
	conn, err := amqp.Dial("amqp://user:password@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	channel, err = conn.Channel()

	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	//declare the queue
	_, err = channel.QueueDeclare(
		"order_queue", // name the order_queue
		true,          // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
}

func PublishOrder(orderID string) {
	err := channel.Publish(
		"",            // exchange
		"order_queue", // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(orderID),
		})
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}
}
