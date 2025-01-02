package handlers

import (
	"database/sql" // Importa database/sql
	"ecommerce-system/services/orders/messaging"
	"ecommerce-system/services/orders/repository"
	"net/http"
	"strconv" // type convertion

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required"`
}

func CreateOrder(c *gin.Context, db *sql.DB) {
	var req CreateOrderRequest

	// Valida la entrada
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Crea el pedido en la base de datos
	orderID, err := repository.CreateOrder(req.ProductID, req.Quantity, "pending")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Convierte orderID de int a string
	orderIDStr := strconv.Itoa(orderID)

	// Publica el pedido en RabbitMQ
	messaging.PublishOrder(orderIDStr)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Order created successfully",
		"order_id": orderID,
	})
}
