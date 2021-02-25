package models

import (
	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	Cart_id                 int
	Account_number_owner_id int
	Total_price             string
	Transaction_code        string
	Status                  string
}
