package service

import (
	"ecommerce-app/internal/database"
	"ecommerce-app/internal/model"
)

func FetchUsersRole() ([]model.UserRole, error) {
	return database.GetAllUsersRole()
}

func InsertUserRole(userRole *model.UserRole) error {
	return database.InsertUserRole(userRole)
}

func UpdateUserRole(id string, userRole *model.UserRole) error {
	return database.UpdateUserRole(id, userRole)
}

func DeleteUserRole(id string) error {
	return database.DeleteUserRole(id)
}
