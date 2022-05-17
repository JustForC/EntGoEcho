package validation

import (
	"fmt"
	"net/http"
	"strings"

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
		errorMessage := ErrorMessage(err)
		return echo.NewHTTPError(http.StatusBadRequest, errorMessage)
	}
	return nil
}

func ErrorMessage(check error) map[string]string {
	var message map[string]string

	message = map[string]string{}

	for _, e := range check.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			message[strings.ToLower(e.Field())] = fmt.Sprintf("Field %s can not be null!", e.Field())
		case "number":
			message[strings.ToLower(e.Field())] = fmt.Sprintf("Input field %s must be number!", e.Field())
		case "email":
			message[strings.ToLower(e.Field())] = fmt.Sprintf("Input must be email type!")
		case "min":
			message[strings.ToLower(e.Field())] = fmt.Sprintf("Input %s minimal %s characters!", e.Field(), e.Param())
		case "max":
			message[strings.ToLower(e.Field())] = fmt.Sprintf("Input %s maximal %s characters!", e.Field(), e.Param())
		case "alpanumeric":
			message[strings.ToLower(e.Field())] = fmt.Sprintf("Input field %s must be alphanumeric!", e.Field())
		case "unique":
			message[strings.ToLower(e.Field())] = fmt.Sprintf("%s already exist!", e.Field())
		}
	}

	return message
}
