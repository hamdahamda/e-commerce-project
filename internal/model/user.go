package model

import (
	"time";
)

type User struct {
    UserID      int       `json:"user_id"`
    Username    string    `json:"username"`
    Email       string    `json:"email"`
    Password    string    `json:"password"`
    FullName    string    `json:"full_name"`
    PhoneNumber string    `json:"phone_number"`
    Address     string    `json:"address"`
    Status      string    `json:"status"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
