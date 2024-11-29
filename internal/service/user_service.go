package service

import (
	"ecommerce-app/internal/database"
	"ecommerce-app/internal/model"
)

func FetchUser() ([]model.User, error) {
	return database.GetAllUsers()
}

func InsertUser(user *model.User) error {
	return database.InsertUser(user)
}

func UpdateUser(id string, user *model.User) error {
	return database.UpdateUser(id, user)
}

func DeleteUser(id string) error {
	return database.DeleteUser(id)
}
