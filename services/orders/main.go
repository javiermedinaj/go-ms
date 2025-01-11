package main

import (
	"ecommerce-system/services/orders/handlers"
	"ecommerce-system/services/orders/messaging"
	"ecommerce-system/services/orders/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := repository.InitDB()

	messaging.InitRabbitMQ()

	// Configuración del router
	router := gin.Default()
	router.POST("/orders", func(c *gin.Context) {
		handlers.CreateOrder(c, db) // db es de tipo *sql.DB
	})

	// Agrega la ruta para obtener órdenes
	router.GET("/orders", func(c *gin.Context) {
		handlers.GetOrders(c, db)
	})

	// Iniciar el servidor
	log.Fatal(http.ListenAndServe(":8080", router))
}
