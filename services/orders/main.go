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

	// health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Iniciar el servidor
	log.Fatal(http.ListenAndServe(":8000", router))
}
