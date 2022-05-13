package test

import (
	"clean-architecture/adapters/controllers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	reqJson = `{"id":1}`
	resJson = `{"id":1,"name":"Sam","age":29}`
)

func setUp() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/emoloyee", strings.NewReader(reqJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func TestOK(t *testing.T) {
	c, rec := setUp()
	controller := controllers.NewEmployeeController()
	// Assertions
	if assert.NoError(t, controller.Handle(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, resJson, rec.Body.String())
	}

}
