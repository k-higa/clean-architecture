package main

import (
	"clean-architecture/adapters/controllers"
	"clean-architecture/external_interfaces"
	"github.com/labstack/echo/v4"
)

var (
	//Logger
	Logger echo.Logger
)

func main() {
	e := echo.New()
	controllers.SetRoute(e)
	Logger = e.Logger
	db := new(external_interfaces.DBManager)
	db.Open()
	e.Logger.Fatal(e.Start(":1323"))
}
