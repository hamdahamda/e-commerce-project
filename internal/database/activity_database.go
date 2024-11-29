package database

import (
    "ecommerce-app/config"
    "ecommerce-app/internal/model"
)

func GetAllActivities() ([]model.Activity, error) {
    rows, err := config.DB.Query("SELECT * FROM activities")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var activities []model.Activity
    for rows.Next() {
        var activity model.Activity
        if err := rows.Scan(
            &activity.ID,
            &activity.Title,
            &activity.Category,
            &activity.Description,
            &activity.ActivityDate,
            &activity.Status,
            &activity.CreatedAt,
        ); err != nil {
            return nil, err
        }
        activities = append(activities, activity)
    }
    return activities, nil
}

func InsertActivity(activity *model.Activity) error {
    sqlStatement := `INSERT INTO activities (title, category, description, activity_date, status)
                    VALUES ($1, $2, $3, $4, $5) RETURNING id`
    return config.DB.QueryRow(sqlStatement, activity.Title, activity.Category, activity.Description, activity.ActivityDate, "NEW").Scan(&activity.ID)
}

func UpdateActivity(id string, activity *model.Activity) error {
    sqlStatement := `UPDATE activities SET title=$1, category=$2, description=$3, activity_date=$4 WHERE id=$5 RETURNING id`
    return config.DB.QueryRow(sqlStatement, activity.Title, activity.Category, activity.Description, activity.ActivityDate, id).Scan(&activity.ID)
}

func DeleteActivity(id string) error {
    sqlStatement := `DELETE FROM activities WHERE id=$1`
    _, err := config.DB.Exec(sqlStatement, id)
    return err
}
