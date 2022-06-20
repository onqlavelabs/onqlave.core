package utils

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"onqlave.io/onqlave.core/pkg/common"
)

type (
	ServiceValidator struct {
		validator *validator.Validate
	}
)

func (cv *ServiceValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, common.NewBaseErrorResponse(int32(http.StatusBadRequest), err.Error()))
	}
	return nil
}

func NewServiceValidator() *ServiceValidator {
	return &ServiceValidator{validator: validator.New()}
}
