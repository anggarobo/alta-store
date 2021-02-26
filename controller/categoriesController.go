package controller

import (
	"alta-store/lib/database"
	"alta-store/models"
	"net/http"
	"strconv"

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

func GetCategoryController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	category, err := database.GetCategory(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"products": category,
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
