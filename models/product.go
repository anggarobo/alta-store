package models

import (
	"gorm.io/gorm"
)

// Product Model
type Product struct {
	gorm.Model
	Name        string `gorm:"size:100" json: "name" form: "name"`
	Stock       int    `json: "email" form: "email"`
	Price       int    `json: "price" form: "price"`
	Description string `json: "description" form: "description"`
	CategoryID  uint   `json: "category_id" form: "category_id"`
	// Category    Category `gorm:"foreignKey:ID"`
}

// func (products *Product) TableName() string {
// 	return "products"
// }

// func (products *Product) toString() string {
// 	return fmt.Sprintf("id: %d\nname: %s", products.ID, products.Name)
// }
