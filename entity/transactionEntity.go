package entity

type CheckoutResponse struct {
	User_id                   int
	Cart_id                   int
	Total_price               int
	Transaction_code          string
	Owner_bank_name           string
	Owner_account_number      string
	Owner_account_number_name string
}

type PaymentResponse struct {
	User_id                      int
	ID                           int
	Total_price                  int
	Transaction_code             string
	Status                       string
	Owner_bank_name              string
	Owner_account_number         string
	Owner_account_number_name    string
	Customer_bank_name           string
	Customer_account_number      string
	Customer_account_number_name string
}
