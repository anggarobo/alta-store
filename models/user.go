package models

import (
	"time"

	"gorm.io/gorm"
)

// User Model
type User struct {
	gorm.Model
	fullName   string
	email      string
	password   string
	address    string
	phone      string
	role       string
	token      string
	created_at time.Time
	updated_at time.Time
}
