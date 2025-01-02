package handlers

import (
	"database/sql"
	"net/http"

	"ecommerce-system/services/orders/repository"

	"github.com/gin-gonic/gin"
)

// Order representa una orden en el sistema
type Order struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status"`
}

// GetOrders maneja la solicitud para obtener todas las órdenes
func GetOrders(c *gin.Context, db *sql.DB) {
	// Obtiene las órdenes de la base de datos
	orders, err := repository.FetchOrders(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	// Responde con las órdenes en formato JSON
	c.JSON(http.StatusOK, orders)
}
