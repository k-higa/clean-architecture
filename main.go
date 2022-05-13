package main

import (
	"clean-architecture/adapters/controllers"
	gateways2 "clean-architecture/adapters/gateways"
	"clean-architecture/database"
	"clean-architecture/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	//Logger
	Logger echo.Logger
)

func main() {
	e := echo.New()
	router(e)
	Logger = e.Logger
	db := new(database.DBManager)
	db.Open()
	e.Logger.Fatal(e.Start(":1323"))
}

func router(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/emoloyee:id", func(c echo.Context) error {
		controller := controllers.NewEmployeeController(usecases.NewEmployeeUsecase(gateways2.NewEmoloyeeRepository()))
		return controller.Handle(c)
	})
}
