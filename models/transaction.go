package models

import (
	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	User_id                      int
	Cart_id                      int
	ID                           int
	Total_price                  int    `json: "total_price" form:"total_price"`
	Transaction_code             string `json: "transaction_code" form:"transaction_code"`
	Status                       int
	Owner_bank_name              string
	Owner_account_number         string
	Owner_account_number_name    string
	Customer_bank_name           string `json: "customer_bank_name" form:"customer_bank_name"`
	Customer_account_number      string `json: "customer_account_number" form:"customer_account_number"`
	Customer_account_number_name string `json: "customer_account_number_name" form:"customer_account_number_name"`
}
