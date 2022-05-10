package route

import (
	"CompanyAPI/configs"
	"CompanyAPI/handler"
	"CompanyAPI/validation"
	"context"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	db := configs.Open()

	ctx := context.Background()
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	companyHand := handler.NewCompanyHandler(db)
	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	e.POST("/company", companyHand.Create)
	e.GET("/company", companyHand.Read)
	e.GET("/company/:id", companyHand.ReadByID)
	e.PUT("/company/:id", companyHand.Update)
	e.DELETE("/company/:id", companyHand.Delete)

	return e
}
