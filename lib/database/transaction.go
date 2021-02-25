package database

import (
	"alta-store/config"
	"alta-store/models"
)

func CreateCheckouts(transaction models.Transactions) (models.Transactions, error) {
	if e := config.DB.Create(&transaction).Error; e != nil {
		return transaction, e
	}

	return transaction, nil
}

func GetDetailTransaction(transactionID int) (models.Transactions, error) {
	var transaction models.Transactions
	if e := config.DB.First(&transaction, transactionID).Error; e != nil {
		return transaction, e
	}

	return transaction, nil
}

func GetCodeTransaction(transactionID int) (models.Transactions, error) {
	var transaction models.Transactions
	err := config.DB.Where("id = ?", transactionID).Where("status", 0).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func CreatePayments(transactionID int, transactions models.Transactions) (models.Transactions, error) {
	var transaction models.Transactions

	if err := config.DB.Model(models.Transactions{}).
		Where("id = ?", transactionID).
		Updates(models.Transactions{
			Status:                       1,
			Total_price:                  transactions.Total_price,
			Customer_bank_name:           transactions.Customer_bank_name,
			Customer_account_number:      transactions.Customer_account_number,
			Customer_account_number_name: transactions.Customer_account_number_name,
		}).Error; err != nil {
		return transaction, err
	}

	return transaction, nil

}
