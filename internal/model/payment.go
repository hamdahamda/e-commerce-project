package model

import "time"

type Payment struct {
	PaymentID     int       `json:"payment_id"`
	OrderID       int       `json:"order_id"`
	PaymentDate   time.Time `json:"payment_date"`
	PaymentAmount float64   `json:"payment_amount"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status"`
}
