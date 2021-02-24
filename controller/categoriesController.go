package controller

import (
	"alta-store/lib/database"
	"alta-store/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCategoriesController(c echo.Context) error {
	categories, err := database.GetCategories()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"categories": categories,
	})
}

func CreateCategoryController(c echo.Context) error {
	category := models.Category{}
	c.Bind(&category)
	categories, err := database.CreateCategory(&category)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"categories": categories,
	})
}
