package models

import (
	"gorm.io/gorm"
)

// Product Model
type Product struct {
	gorm.Model
	Name        string `gorm:"size:100" json: "name" form: "name"`
	Stock       int    `json: "stock" form: "stock"`
	Price       int    `json: "price" form: "price"`
	Description string `json: "description" form: "description"`
	CategoryID  uint   `json: "category_id" form: "category_id"`
}
