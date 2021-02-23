package controllers

import (
	"alta-store/lib/database"
	"alta-store/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCategoriesController(c echo.Context) error {
	categories, e := database.GetCategories()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"categories": categories,
	})
}

func CreateCategoryController(c echo.Context) error {
	category := models.Category{}
	c.Bind(&category)
	categories, e := database.CreateCategory(&category)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":     "success",
		"categories": categories,
	})
}
