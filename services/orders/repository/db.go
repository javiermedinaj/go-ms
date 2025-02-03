package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Driver para PostgreSQL
)

type Order struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status"`
}

var db *sql.DB

func InitDB() *sql.DB {
	var err error
	connStr := "host=postgres port=5432 user=postgres password=admin dbname=ecommerce sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Verifica la conexión
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	// Crear tabla si no existe
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			product_id INT NOT NULL,
			quantity INT NOT NULL,
			status TEXT NOT NULL
		)
	`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	fmt.Println("Database connected and ready!")
	return db
}

func CreateOrder(productID int, quantity int, status string) (int, error) {
	var id int
	err := db.QueryRow(`
		INSERT INTO orders (product_id, quantity, status) 
		VALUES ($1, $2, $3) RETURNING id
	`, productID, quantity, status).Scan(&id)
	return id, err
}

// get orders
func FetchOrders(db *sql.DB) ([]Order, error) {
	rows, err := db.Query(`SELECT id, product_id, quantity, status FROM orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.ProductID, &order.Quantity, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
