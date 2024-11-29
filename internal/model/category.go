package model

type Category struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	ParentID     *int   `json:"parent_id"`
	Description  string `json:"description"`
}
