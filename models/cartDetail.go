package models

import (
	"gorm.io/gorm"
)

type Cart_details struct {
	gorm.Model
	Cart_id    int `json: "cart_id" form:"cart_id"`
	Product_id int `json: "product_id" form:"product_id"`
	Quantity   int `json: "quantity" form:"quantity"`
	Price      int `json: "price" form:"price"`
}
