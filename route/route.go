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
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	db := configs.Open()

	ctx := context.Background()
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	companyHand := handler.NewCompanyHandler(db)
	employeeHand := handler.NewEmployeeHandler(db)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	// Company
	e.POST("/company", companyHand.Create)
	e.GET("/company", companyHand.Read)
	e.GET("/company/:id", companyHand.ReadByID)
	e.PUT("/company/:id", companyHand.Update)
	e.DELETE("/company/:id", companyHand.Delete)

	// Join Company With Employee
	e.GET("/company/employee", companyHand.CompanyWithEmployee)
	e.GET("/company/:id/employee", companyHand.CompanyIDWithEmployee)
	e.GET("/company/employee/:id", companyHand.CompanyWithEmployeeID)
	e.GET("/company/:companyid/employee/:employeeid", companyHand.CompanyIDWithEmployeeID)

	// Employee
	e.POST("/employee", employeeHand.Create)
	e.GET("/employee", employeeHand.Read)
	e.GET("/employee/:id", employeeHand.ReadByID)
	e.PUT("/employee/:id", employeeHand.Update)
	e.DELETE("/employee/:id", employeeHand.Delete)

	// Join Employee With Company
	e.GET("/employee/company", employeeHand.EmployeeWithCompany)
	e.GET("/employee/:id/company", employeeHand.EmployeeIDWithCompany)
	e.GET("/employee/company/:id", employeeHand.EmployeeWithCompanyID)
	e.GET("/employee/:employeeid/company/:companyid", employeeHand.EmployeeIDWithCompanyID)

	return e
}
