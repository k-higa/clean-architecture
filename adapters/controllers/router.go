package controllers

import (
	"clean-architecture/adapters/gateways/clients"
	"clean-architecture/adapters/gateways/rdb"
	"clean-architecture/usecases"
	"clean-architecture/usecases/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/emoloyee:id", func(c echo.Context) error {
		controller := NewEmployeeController(usecases.NewEmployeeUseCase(repository.NewEmployeeRepository(rdb.NewEmployeeGateway())))
		return controller.Get(c)
	})
	e.POST("/emoloyee", func(c echo.Context) error {
		controller := NewEmployeeController(usecases.NewEmployeeUseCase(repository.NewEmployeeRepository(rdb.NewEmployeeGateway())))
		return controller.Create(c)
	})

	e.POST("/payment/entry", func(c echo.Context) error {
		controller := NewPaymentController(usecases.NewPaymentTransactionUseCase(repository.NewPaymentTransactionRepository(rdb.NewAPIKeyGateway(), clients.NewAmazonPayClient())))
		return controller.Create(c)
	})
}
