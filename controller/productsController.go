package controller

import (
	"alta-store/lib/database"
	"alta-store/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	products, err := database.GetProducts()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": products,
	})
}

func GetCategoryIDProductsController(c echo.Context) error {
	catID := c.Param("category_id")
	products, err := database.GetCategoryID(catID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": products,
	})
}

func GetProductByCategoryController(c echo.Context) error {
	categoryID := c.QueryParam("category")
	if categoryID == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Failed",
		})
	}

	categoryIDExist, err := database.GetCategoryID(categoryID)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if categoryIDExist != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "Failed",
		})
	}

	category, e := database.GetProductByCategory(categoryID)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  category,
	})
}

func CreateProductController(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)
	products, err := database.CreateProduct(&product)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": products,
	})
}
