package controllers

import (
	"clean-architecture/adapters/gateways/clients"
	"clean-architecture/adapters/gateways/rdb"
	"clean-architecture/adapters/gateways/repository"
	"clean-architecture/external_interfaces"
	"clean-architecture/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/emoloyee:id", func(c echo.Context) error {
		useCase := usecases.NewEmployeeUseCase(repository.NewEmployeeRepository(rdb.NewEmployeeGateway()), repository.NewTransactionManger(external_interfaces.DB))
		controller := NewEmployeeController(useCase)
		return controller.Get(c)
	})
	e.POST("/emoloyee", func(c echo.Context) error {
		useCase := usecases.NewEmployeeUseCase(repository.NewEmployeeRepository(rdb.NewEmployeeGateway()), repository.NewTransactionManger(external_interfaces.DB))
		controller := NewEmployeeController(useCase)
		return controller.Create(c)
	})

	e.POST("/payment/entry", func(c echo.Context) error {
		repo := repository.NewPaymentTransactionRepository(rdb.NewAPIKeyGateway(), clients.NewAmazonPayClient())
		useCase := usecases.NewPaymentTransactionUseCase(repo, repository.NewTransactionManger(external_interfaces.DB))
		controller := NewPaymentController(useCase)
		return controller.Entry(c)
	})
}
