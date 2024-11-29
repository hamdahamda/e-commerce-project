package service

import (
	"ecommerce-app/internal/database"
	"ecommerce-app/internal/model"
)

func FetchActivities() ([]model.Activity, error) {
	return database.GetAllActivities()
}

func InsertActivity(activity *model.Activity) error {
	return database.InsertActivity(activity)
}

func UpdateActivity(id string, activity *model.Activity) error {
	return database.UpdateActivity(id, activity)
}

func DeleteActivity(id string) error {
	return database.DeleteActivity(id)
}
