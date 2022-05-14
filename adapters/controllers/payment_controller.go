package controllers

import (
	"clean-architecture/usecases"
	"clean-architecture/usecases/input_port"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PaymentController struct {
	uc usecases.PaymentTransactionUseCase
}

func NewPaymentController(uc usecases.PaymentTransactionUseCase) *PaymentController {
	return &PaymentController{uc: uc}
}

func (p *PaymentController) Create(c echo.Context) error {
	body := &input_port.PaymentEntry{}
	if err := c.Bind(body); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	result, err := p.uc.Entry(*body)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "internal servaer error")
	}
	return c.JSON(http.StatusOK, result)
}
