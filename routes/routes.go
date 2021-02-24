package routes

import (
	"alta-store/constraints"
	"alta-store/controller"

	"alta-store/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controller.LoginUsersController)
	e.POST("/users", controller.CreateUsersController)

	jwt := e.Group("/jwt")
	jwt.Use(middleware.JWT([]byte(constraints.SECRET_JWT)))

	jwt.GET("/carts", controller.GetCartController)
	jwt.POST("/carts", controller.AddToCartController)
	jwt.DELETE("/carts", controller.DeleteCartOnController)

	e.GET("/products", controllers.GetProductsController)
	e.GET("/products", controllers.GetProductByCategoryController)

	e.GET("/categories", controllers.GetCategoriesController)

	return e
}
