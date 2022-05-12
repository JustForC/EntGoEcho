package validation

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {

	if err := cv.Validator.Struct(i); err != nil {
		errorMessage := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessage = append(errorMessage, message)
		}
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}
	return nil
}
