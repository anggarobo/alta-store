package models

import (
	"time"

	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	ID         string
	FullName   string `json:"fullname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
	Token      string `json:"token"`
	Created_at time.Time
	Updated_at time.Time
}
