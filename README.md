# Ecommerce System

This project is a simple practice to create my first microservice in Go. It consists of multiple services that communicate with each other using RabbitMQ.

## Services

- **RabbitMQ**: A message broker used for communication between services.
- **Order Service**: Manages orders.
- **Inventory Service**: Manages inventory.
- **Notification Service**: Sends notifications.

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Running the Services

To start all the services, run the following command:

```sh
docker-compose up
```

This will start the RabbitMQ service along with the Order, Inventory, and Notification services.

### Accessing the Services

- **RabbitMQ Management**: [http://localhost:15672](http://localhost:15672)
- **Order Service**: [http://localhost:8081](http://localhost:8081)
- **Inventory Service**: [http://localhost:8082](http://localhost:8082)
- **Notification Service**: [http://localhost:8083](http://localhost:8083)

## Project Structure

- `services/orders`: Contains the Order Service.
- `services/inventory`: Contains the Inventory Service.
- `services/notifications`: Contains the Notification Service.

## Built With

- [Go](https://golang.org/) - The programming language used.
- [Docker](https://www.docker.com/) - Containerization platform.
- [RabbitMQ](https://www.rabbitmq.com/) - Message broker.

