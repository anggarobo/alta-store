package models

import "time"

type AccountNumberOwner struct {
	ID                  string
	Account_number      string `json:"account_number "`
	Account_bank        string `json:"account_bank"`
	Account_number_name string `json:"account_number_name"`
	Created_at          time.Time
}
