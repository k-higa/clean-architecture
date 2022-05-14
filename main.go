package main

import (
	"clean-architecture/adapters/controllers"
	"clean-architecture/app"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controllers.SetRoute(e)
	app.Logger = e.Logger
	db := new(app.DBManager)
	db.Open()
	e.Logger.Fatal(e.Start(":1323"))
}
