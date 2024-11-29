package service

import (
	"ecommerce-app/internal/database"
	"ecommerce-app/internal/model"
)

func FetchCategories() ([]model.Category, error) {
	return database.GetAllCategory()
}

func InsertCategory(category *model.Category) error {
	return database.InsertCategory(category)
}

func UpdateCategory(id string, category *model.Category) error {
	return database.UpdateCategory(id, category)
}

func DeleteCategory(id string) error {
	return database.DeleteCategory(id)
}
