package model

import (
	"time"
)

type Stock struct {
	StockID           int       `json:"stock_id"`
	ProductID         int       `json:"product_id"`
	WarehouseLocation string    `json:"warehouse_location"`
	Quantity          int       `json:"quantity"`
	UpdatedAt         time.Time `json:"updated_at"`
}
