package models

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Category_id int    `json: "category_id" form:"category_id"`
	Name        string `json: "name" form:"name"`
	Stock       int    `json: "stock" form:"stock"`
	Price       int    `json: "price" form:"price"`
	Description string `json: "description" form:"description"`
}
