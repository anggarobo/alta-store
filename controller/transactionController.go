package controller

import (
	"alta-store/lib/database"
	"alta-store/middlewares"
	"net/http"

	"github.com/labstack/echo"
)

func AddCheckoutController(c echo.Context) error {

	userID := middlewares.ExtractTokenUserId(c)
	_, e := database.GetCartId(userID)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return nil
}
