package routes

import (
	"alta-store/constraints"
	"alta-store/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func New() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())

	// POST
	e.POST("/login", controller.LoginUsersController)
	e.POST("/users", controller.CreateUsersController)

	jwt := e.Group("/jwt")
	jwt.Use(middleware.JWT([]byte(constraints.SECRET_JWT)))

	jwt.GET("/carts", controller.GetCartController)
	jwt.POST("/carts", controller.AddToCartController)
	jwt.DELETE("/carts", controller.DeleteCartOnController)

	e.GET("/products", controller.GetProductsController)
	e.GET("/product", controller.GetProductByCategoryController)

	e.GET("/categories", controller.GetCategoriesController)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
