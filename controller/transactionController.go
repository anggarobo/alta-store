package controller

import (
	"alta-store/entity"
	"alta-store/lib/database"
	"alta-store/middlewares"
	"alta-store/models"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

const errorMessage string = "Error"
const errorMessageFailCheckout string = "You don't have any active cart"
const successMessage string = "Success"
const successMessageCheckout string = "Checkout successfully"
const transactionNoFound string = "Transaction is not found"
const transactionCodeNoFound string = "Transaction Code is not found"
const nominalNotMatch string = "Nominal is not match the bill"
const paymentFailed string = "Fail to process your payment"
const paymentSuccess string = "Process your payment successfully"
const completeLeftCheckout string = "Cannot create checkout, you still have any active checkout"

func AddCheckoutController(c echo.Context) error {

	userID := middlewares.ExtractTokenUserId(c)
	activeCart, e := database.GetActiveCart(userID)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  errorMessage,
			"message": errorMessageFailCheckout,
		})
	}

	carts, e := database.GetDetailActiveCart(activeCart)

	reg, _ := regexp.Compile("[^0-9]+")
	currentTime := time.Now()
	time := currentTime.String()
	uniqueCode := time[:19]
	uniqueCode = reg.ReplaceAllString(uniqueCode, "")

	_, e = database.CloseCart(carts.ID)
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  errorMessage,
			"message": e.Error(),
		})
	}

	activeCheckout, e := database.SearchActiveCheckout(userID)

	if activeCheckout.ID != 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  errorMessage,
			"message": completeLeftCheckout,
		})
	}

	data := models.Transactions{
		User_id:                   userID,
		Cart_id:                   carts.ID,
		Total_price:               carts.Total_price,
		Transaction_code:          uniqueCode,
		Status:                    0,
		Owner_bank_name:           "BCA",
		Owner_account_number:      "1456-9876-879",
		Owner_account_number_name: "ALTA STORE INDONESIA",
	}

	lastInsertData, e := database.CreateCheckouts(data)

	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  errorMessage,
			"message": errorMessageFailCheckout,
		})
	}

	response := entity.CheckoutResponse{
		User_id:                   userID,
		Cart_id:                   lastInsertData.Cart_id,
		Transaction_code:          lastInsertData.Transaction_code,
		Total_price:               lastInsertData.Total_price,
		Owner_bank_name:           lastInsertData.Owner_bank_name,
		Owner_account_number:      lastInsertData.Owner_account_number,
		Owner_account_number_name: lastInsertData.Owner_account_number_name,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  successMessage,
		"message": successMessageCheckout,
		"data":    response,
	})

}

func ProcessPayment(c echo.Context) error {

	userID := middlewares.ExtractTokenUserId(c)

	transactionID := c.Param("transactionID")

	transaction, _ := strconv.Atoi(transactionID)

	dataTransaction, err := database.GetCodeTransaction(transaction)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  errorMessage,
			"message": transactionNoFound,
		})
	}

	if dataTransaction.Transaction_code != c.FormValue("transaction_code") {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  errorMessage,
			"message": transactionCodeNoFound,
		})
	}

	total_price_input, _ := strconv.Atoi(c.FormValue("total_price"))
	if dataTransaction.Total_price != total_price_input {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  errorMessage,
			"message": nominalNotMatch,
		})
	}

	data := models.Transactions{}
	c.Bind(&data)

	_, err = database.CreatePayments(transaction, data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  errorMessage,
			"message": paymentFailed,
		})
	}

	detailTransaction, _ := database.GetDetailTransaction(transaction)

	response := entity.PaymentResponse{
		User_id:                      userID,
		ID:                           detailTransaction.ID,
		Total_price:                  detailTransaction.Total_price,
		Transaction_code:             detailTransaction.Transaction_code,
		Status:                       "COMPLETE",
		Owner_bank_name:              detailTransaction.Owner_bank_name,
		Owner_account_number:         detailTransaction.Owner_account_number,
		Owner_account_number_name:    detailTransaction.Owner_account_number_name,
		Customer_bank_name:           detailTransaction.Customer_bank_name,
		Customer_account_number:      detailTransaction.Customer_account_number_name,
		Customer_account_number_name: detailTransaction.Customer_account_number_name,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  successMessage,
		"message": paymentSuccess,
		"data":    response,
	})

}
