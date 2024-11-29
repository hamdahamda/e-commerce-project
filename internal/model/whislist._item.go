package model

import "time"

type WishlistItem struct {
    WishlistItemID int       `json:"wishlist_item_id"`
    WishlistID     int       `json:"wishlist_id"`
    ProductID      int       `json:"product_id"`
    AddedAt        time.Time `json:"added_at"`
}
