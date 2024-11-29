package service

import (
	"ecommerce-app/internal/database"
	"ecommerce-app/internal/model"
)

func FetchCategories() ([]model.Category, error) {
	return database.GetAllUsers()
}

func InsertCategoryr(user *model.Category) error {
	return database.InsertUser(user)
}

func UpdateCategory(id string, user *model.Category) error {
	return database.UpdateUser(id, user)
}

func DeleteCategory(id string) error {
	return database.DeleteUser(id)
}
