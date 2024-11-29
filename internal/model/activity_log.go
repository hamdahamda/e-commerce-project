package model

import "time"

type ActivityLog struct {
	LogID        int       `json:"log_id"`
	UserID       *int      `json:"user_id"`
	ActivityType string    `json:"activity_type"`
	Description  string    `json:"description"`
	ActivityTime time.Time `json:"activity_time"`
}
