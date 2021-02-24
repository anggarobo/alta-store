package routes

import (
	"alta-store/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/carts", controller.GetCartController)
	e.POST("/carts", controller.AddToCartController)
	e.DELETE("/carts", controller.DeleteCartOnController)

	return e
}
