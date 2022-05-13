package controllers

import (
	"clean-architecture/usecases"
	"clean-architecture/usecases/input_port"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type EmployeeController struct {
	uc usecases.EmployeeUseCase
}

func NewEmployeeController(uc usecases.EmployeeUseCase) *EmployeeController {
	return &EmployeeController{uc: uc}
}

func (e *EmployeeController) Handle(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := e.uc.FindEmployee(input_port.Emoployee{ID: id})
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "internal servaer error ")
	}
	return c.JSON(http.StatusInternalServerError, result)
}

func createUser(c echo.Context) error {
	return nil
}
