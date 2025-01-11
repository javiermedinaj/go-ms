package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

func main() {
	// connect to rabbit mq

	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ : %s", err)
	}

	defer conn.Close()

	fmt.Println("Successfully connected to RabbitMQ")

	// create a handler for the inventory service
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Inventory service is healthy"))
	})

	// listen port 8081
	fmt.Println("Inventory service is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
