package handler

import (
	"CompanyAPI/ent"
	"CompanyAPI/ent/company"
	"CompanyAPI/ent/employee"
	"CompanyAPI/request"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type employeeHandler struct {
	db *ent.Client
}

func NewEmployeeHandler(db *ent.Client) *employeeHandler {
	return &employeeHandler{db}
}

func (empHand *employeeHandler) Create(c echo.Context) error {
	emp := new(request.EmployeeRequest)
	if err := c.Bind(emp); err != nil {
		return err
	}
	if err := c.Validate(emp); err != nil {
		return err
	}
	ctx := context.Background()
	comp := empHand.db.Company.Query().Where(company.Name(emp.CompanyName)).Select(company.FieldID).IntX(ctx)
	newEmp := empHand.db.Employee.Create().SetCompaniesID(comp).SetName(emp.Name).SetPosition(emp.Position).SetSalary(emp.Salary).SaveX(ctx)
	return c.JSON(http.StatusOK, newEmp)
}

func (empHand *employeeHandler) Read(c echo.Context) error {
	ctx := context.Background()
	emps := empHand.db.Employee.Query().AllX(ctx)
	return c.JSON(http.StatusOK, emps)
}

func (empHand *employeeHandler) ReadByID(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))
	emp := empHand.db.Employee.Query().Where(employee.ID(id)).AllX(ctx)
	return c.JSON(http.StatusOK, emp)
}

func (empHand *employeeHandler) Update(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request.EmployeeRequest)
	c.Bind(req)

	company_id := req.CompanyName
	if company_id == "" {
		company_id = empHand.db.Employee.Query().Where(employee.ID(id)).Select(employee.ForeignKeys[0]).StringX(ctx)
	} else {
		company_id = empHand.db.Company.Query().Where(company.Name(req.CompanyName)).Select(company.FieldID).StringX(ctx)
	}
	comp_id, _ := strconv.Atoi(company_id)

	name := req.Name
	if name == "" {
		name = empHand.db.Employee.Query().Where(employee.ID(id)).Select(employee.FieldName).StringX(ctx)
	}

	salary := req.Salary
	if salary == 0 {
		salary = int32(empHand.db.Employee.Query().Where(employee.ID(id)).Select(employee.FieldSalary).IntX(ctx))
	}

	position := req.Position
	if position == "" {
		position = empHand.db.Employee.Query().Where(employee.ID(id)).Select(employee.FieldPosition).StringX(ctx)
	}

	emp := empHand.db.Employee.UpdateOneID(id).SetName(name).SetCompaniesID(comp_id).SetPosition(position).SetSalary(salary).SaveX(ctx)

	return c.JSON(http.StatusOK, emp)
}

func (empHand *employeeHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))
	emp := empHand.db.Employee.Query().Where(employee.ID(id)).AllX(ctx)
	empHand.db.Employee.DeleteOneID(id).ExecX(ctx)

	return c.JSON(http.StatusOK, emp)
}

// Join Edges With Company
func (empHand *employeeHandler) EmployeeWithCompany(c echo.Context) error {
	ctx := context.Background()

	emps := empHand.db.Employee.Query().WithCompanies().AllX(ctx)

	return c.JSON(http.StatusOK, emps)
}

func (empHand *employeeHandler) EmployeeIDWithCompany(c echo.Context) error {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	emp := empHand.db.Employee.Query().Where(employee.ID(id)).WithCompanies().AllX(ctx)

	return c.JSON(http.StatusOK, emp)
}

func (empHand *employeeHandler) EmployeeWithCompanyID(c echo.Context) error {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))

	emp := empHand.db.Employee.Query().WithCompanies(func(com *ent.CompanyQuery) {
		com.Where(company.ID(id)).AllX(ctx)
	}).AllX(ctx)

	return c.JSON(http.StatusOK, emp)
}

func (empHand *employeeHandler) EmployeeIDWithCompanyID(c echo.Context) error {
	ctx := context.Background()

	employeeid, _ := strconv.Atoi(c.Param("employeeid"))
	companyid, _ := strconv.Atoi(c.Param("companyid"))

	emp := empHand.db.Employee.Query().Where(employee.ID(employeeid)).WithCompanies(func(com *ent.CompanyQuery) {
		com.Where(company.ID(companyid)).AllX(ctx)
	}).AllX(ctx)

	return c.JSON(http.StatusOK, emp)
}
