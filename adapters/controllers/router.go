package controllers

import (
	"clean-architecture/adapters/gateways"
	"clean-architecture/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/emoloyee:id", func(c echo.Context) error {
		controller := NewEmployeeController(usecases.NewEmployeeUsecase(gateways.NewEmoloyeeRepository()))
		return controller.Get(c)
	})
	e.POST("/emoloyee", func(c echo.Context) error {
		controller := NewEmployeeController(usecases.NewEmployeeUsecase(gateways.NewEmoloyeeRepository()))
		return controller.Create(c)
	})
}
