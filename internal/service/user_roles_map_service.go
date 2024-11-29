package service

import (
	"ecommerce-app/internal/database"
	"ecommerce-app/internal/model"
)

func FetchUsersRoleMap() ([]model.UserRoleMap, error) {
	return database.GetAllUsersRoleMap()
}

func InsertUserRoleMap(userRoleMap *model.UserRoleMap) error {
	return database.InsertUserRoleMap(userRoleMap)
}

func UpdateUserRoleMap(id string, userRoleMap *model.UserRoleMap) error {
	return database.UpdateUserRoleMap(id, userRoleMap)
}

func DeleteUserRoleMap(id string) error {
	return database.DeleteUserRoleMap(id)
}
