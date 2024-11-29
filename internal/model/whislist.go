package model

import (
	"time"
)

type Wishlist struct {
	WishlistID int       `json:"wishlist_id"`
	UserID     int       `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
}
