package controllers

import (
	"clean-architecture/adapters/gateways/clients"
	"clean-architecture/adapters/gateways/stores"
	"clean-architecture/external_interfaces"
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
		useCase := usecases.NewEmployeeUseCase(repository.NewEmployeeRepository(stores.NewEmployeeGateway()), stores.NewTransactionManger(external_interfaces.DB))
		controller := NewEmployeeController(useCase)
		return controller.Get(c)
	})
	e.POST("/emoloyee", func(c echo.Context) error {
		useCase := usecases.NewEmployeeUseCase(repository.NewEmployeeRepository(stores.NewEmployeeGateway()), stores.NewTransactionManger(external_interfaces.DB))
		controller := NewEmployeeController(useCase)
		return controller.Create(c)
	})

	e.POST("/payment/entry", func(c echo.Context) error {
		ptRepo := repository.NewPaymentTransactionRepository(stores.NewPaymentTransactionGateway(), stores.NewAPIKeyGateway(), clients.NewAmazonPayClient())
		logRepo := repository.NewPaymentLogReoisitory(stores.NewPaymentLogGateway())
		useCase := usecases.NewPaymentTransactionUseCase(ptRepo, logRepo, stores.NewTransactionManger(external_interfaces.DB))
		controller := NewPaymentController(useCase)
		return controller.Entry(c)
	})
}
