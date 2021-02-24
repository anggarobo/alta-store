package routes

import (
	"alta-store/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/products/category/:category", controllers.GetCategoryIDProductsController)
	e.GET("/products", controllers.GetProductsController)
	e.GET("/products", controllers.GetProductByCategoryController)

	e.GET("/categories", controllers.GetCategoriesController)
	e.POST("/category", controllers.CreateCategoryController)

	return e
}
