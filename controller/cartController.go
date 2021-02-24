package controller

import (
	"alta-store/lib/database"
	"alta-store/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetCartController(c echo.Context) error {
	req := c.Request().Header
	userID, e := strconv.Atoi(req.Get("userid"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "Not found User ID on Header",
		})
	}

	carts, e := database.GetCarts(userID)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"carts":  carts,
	})
}

func AddToCartController(c echo.Context) error {

	req := c.Request().Header
	userID, e := strconv.Atoi(req.Get("userid"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "User ID not found in Header",
		})
	}

	cartID, e := database.GetCartId(userID)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	data := models.Cart_details{}
	c.Bind(&data)
	if data.Product_id == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "Product ID not found",
		})
	}
	if data.Quantity == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "Quantity not found",
		})
	}

	productPrice, e := database.GetProductPrice(data.Product_id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	data.Price = productPrice * data.Quantity
	data.Cart_id = int(cartID)
	result, e := database.AddToCart(data)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": result,
	})
}

func DeleteCartOnController(c echo.Context) error {
	req := c.Request().Header
	userID, e := strconv.Atoi(req.Get("userid"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Failed",
			"message": "User ID not found in Header",
		})
	}

	cartDetailID := c.QueryParam("product")
	if cartDetailID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Failed",
			"message": "Cart Detail ID not found",
		})
	}

	isCartDetailAvailable, e := database.GetCartDetailID(cartDetailID)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	if !isCartDetailAvailable {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "Failed",
			"message": "Cart Detail ID not found",
		})
	}

	result, e := database.DeleteCart(cartDetailID)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "Success",
		"ID":      userID,
		"message": result,
	})
}