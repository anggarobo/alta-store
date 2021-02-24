package models

import (
	"gorm.io/gorm"
)

type Carts struct {
	gorm.Model
	User_id     int `json: "user_id" form:"user_id"`
	Total_price int `json: "total_price" form:"total_price"`
	Status      int `json: "status" form:"status"`
}
