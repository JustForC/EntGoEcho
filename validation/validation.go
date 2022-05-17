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

func MessageError(m validator.FieldError) string {
	var message string

	switch m.Tag() {
	case "required":
		message = fmt.Sprintf("Field %s can not be null!", m.Field())
	case "number":
		message = fmt.Sprintf("Input field %s must be number!", m.Field())
	case "email":
		message = fmt.Sprintf("Input must be email type!")
	case "min":
		message = fmt.Sprintf("Input %s minimal %s characters!", m.Field(), m.Param())
	case "max":
		message = fmt.Sprintf("Input %s maximal %s characters!", m.Field(), m.Param())
	case "alpanumeric":
		message = fmt.Sprintf("Input field %s must be alphanumeric!", m.Field())
	case "unique":
		message = fmt.Sprintf("%s already exist!", m.Field())
	}

	return message
}

func (cv *CustomValidator) Validate(i interface{}) error {

	if err := cv.Validator.Struct(i); err != nil {
		errorMessage := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			customMessage := MessageError(e)
			errorMessage = append(errorMessage, customMessage)
		}
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}
	return nil
}
