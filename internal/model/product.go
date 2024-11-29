package model

import "time"

type Product struct {
	ProductID   int       `json:"product_id"`
	CategoryID  *int      `json:"category_id"`
	ProductName string    `json:"product_name"`
	SKU         string    `json:"sku"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Discount    float64   `json:"discount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
