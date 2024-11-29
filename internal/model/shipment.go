package model

import "time"

type Shipment struct {
	ShipmentID      int       `json:"shipment_id"`
	OrderID         int       `json:"order_id"`
	ShipmentDate    time.Time `json:"shipment_date"`
	ShippingAddress string    `json:"shipping_address"`
	TrackingNumber  string    `json:"tracking_number"`
	ShipmentStatus  string    `json:"shipment_status"`
	CourierName     string    `json:"courier_name"`
}
