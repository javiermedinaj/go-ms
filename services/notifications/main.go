package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to Rabbit", err)
	}
	defer conn.Close()

	fmt.Println("Connect to Rabbit")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("notifications service is running"))
	})

	log.Println("Notifications service is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
