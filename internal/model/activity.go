package model

import (
    "time"
)

type Activity struct {
    ID           int       `json:"id"`
    Title        string    `json:"title"`
    Category     string    `json:"category" validate:"required,oneof=TASK EVENT"`
    Description  string    `json:"description" validate:"required"`
    ActivityDate time.Time `json:"activity_date"`
    Status       string    `json:"status"`
    CreatedAt    time.Time `json:"created_at"`
}
