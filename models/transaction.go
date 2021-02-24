package models

import (
	"time"
)

type Transaction struct {
	Total_price        int                `json:"total_price"`
	Transaction_code   string             `json:"transaction_code"`
	Status             int                `json:"status"`
	AccountNumberOwner AccountNumberOwner `json:"account_number_owner_id"`
	Created_at         time.Time
	Updated_at         time.Time
}
