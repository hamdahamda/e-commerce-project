package model

import "time"

type Review struct {
	ReviewID   int       `json:"review_id"`
	ProductID  int       `json:"product_id"`
	UserID     *int      `json:"user_id"`
	Rating     int       `json:"rating"`
	ReviewText string    `json:"review_text"`
	CreatedAt  time.Time `json:"created_at"`
}
