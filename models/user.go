package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Full_name string `gorm:"size:50" json: "full_name" form:"full_name"`
	Email     string `json: "email" form:"email"`
	Password  string `json: "password" form:"password"`
	Address   string `json: "address" form:"address"`
	Phone     string `json: "phone" form:"phone"`
	Role      string `json: "role" form:"role"`
	Token     string `json: "token" form:"token"`
}
