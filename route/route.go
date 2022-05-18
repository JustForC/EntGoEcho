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
	"github.com/labstack/echo/v4/middleware"
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
	authHand := handler.NewAuthenticationHandler(db)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello, World!")
	})

	r := e.Group("/company")
	r.Use(middleware.JWTWithConfig(handler.Config()))
	// Company
	r.POST("", companyHand.Create)
	r.GET("", companyHand.Read)
	r.GET("/:id", companyHand.ReadByID)
	r.PUT("/:id", companyHand.Update)
	r.DELETE("/:id", companyHand.Delete)

	// Join Company With Employee
	r.GET("/employee", companyHand.CompanyWithEmployee)
	r.GET("/:id/employee", companyHand.CompanyIDWithEmployee)
	r.GET("/employee/:id", companyHand.CompanyWithEmployeeID)
	r.GET("/:companyid/employee/:employeeid", companyHand.CompanyIDWithEmployeeID)

	r.PUT("/:id/employee", companyHand.UpdateCompanyWithEmployee)
	r.PUT("/employee/:id", companyHand.CompanyWithUpdateEmployee)
	r.PUT("/:companyid/employee/:employeeid", companyHand.UpdateCompanyWithUpdateEmployee)

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

	e.POST("/test/request", companyHand.TestRequestHandler)

	// Authentication
	e.POST("/auth/register", authHand.Register)
	e.POST("/auth/login", authHand.Login)

	return e
}
