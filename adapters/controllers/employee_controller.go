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

func (e *EmployeeController) Get(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result, err := e.uc.FindEmployee(input_port.Emoployee{ID: id})
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "internal servaer error ")
	}
	return c.JSON(http.StatusOK, result)
}

func (e *EmployeeController) Create(c echo.Context) error {
	body := &input_port.Emoployee{}
	if err := c.Bind(body); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	result, err := e.uc.CreatedEmployee(*body)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "internal servaer error")
	}
	return c.JSON(http.StatusOK, result)
}
