package model

import "time"

type Order struct {
	OrderID         int       `json:"order_id"`
	UserID          *int      `json:"user_id"`
	OrderStatus     string    `json:"order_status"`
	OrderDate       time.Time `json:"order_date"`
	TotalPrice      float64   `json:"total_price"`
	PaymentStatus   string    `json:"payment_status"`
	ShippingAddress string    `json:"shipping_address"`
	BillingAddress  string    `json:"billing_address"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
