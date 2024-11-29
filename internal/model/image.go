package model

type Image struct {
	ImageID   int    `json:"image_id"`
	ProductID int    `json:"product_id"`
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}
