package handler

import (
	"CompanyAPI/ent"
	"CompanyAPI/ent/company"
	"CompanyAPI/ent/employee"
	"CompanyAPI/request"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type companyHandler struct {
	db *ent.Client
}

func NewCompanyHandler(db *ent.Client) *companyHandler {
	return &companyHandler{db}
}

func (compHand *companyHandler) Create(c echo.Context) error {
	comp := new(request.CompanyRequest)
	if err := c.Bind(comp); err != nil {
		return err
	}
	if err := c.Validate(comp); err != nil {
		return err
	}
	ctx := context.Background()
	newComp := compHand.db.Company.Create().SetName(comp.Name).SetAddress(comp.Address).SetService(comp.Service).SaveX(ctx)
	return c.JSON(http.StatusOK, newComp)
}

func (compHand *companyHandler) Read(c echo.Context) error {
	ctx := context.Background()
	comps := compHand.db.Company.Query().AllX(ctx)
	return c.JSON(http.StatusOK, comps)
}

func (compHand *companyHandler) ReadByID(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))
	comp := compHand.db.Company.Query().Where(company.ID(id)).AllX(ctx)

	return c.JSON(http.StatusOK, comp)
}

func (compHand *companyHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request.CompanyRequest)
	ctx := context.Background()
	c.Bind(req)

	name := req.Name
	if name == "" {
		name = compHand.db.Company.Query().Where(company.ID(id)).Select(company.FieldName).StringX(ctx)
	}

	address := req.Address
	if address == "" {
		address = compHand.db.Company.Query().Where(company.ID(id)).Select(company.FieldAddress).StringX(ctx)
	}

	service := req.Service
	if service == "" {
		service = compHand.db.Company.Query().Where(company.ID(id)).Select(company.FieldService).StringX(ctx)
	}

	comp := compHand.db.Company.UpdateOneID(id).SetAddress(address).SetName(name).SetService(service).SaveX(ctx)

	return c.JSON(http.StatusOK, comp)
}

func (compHand *companyHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))

	comp := compHand.db.Company.Query().Where(company.ID(id)).AllX(ctx)
	del := compHand.db.Company.DeleteOneID(id).Exec(ctx)

	if del != nil {
		log.Fatal(del)
	}

	return c.JSON(http.StatusOK, comp)
}

// Join Edges With Employee

// Read Edges
func (compHand *companyHandler) CompanyWithEmployee(c echo.Context) error {
	ctx := context.Background()

	comps := compHand.db.Company.Query().WithEmployees().AllX(ctx)

	return c.JSON(http.StatusOK, comps)
}

func (compHand *companyHandler) CompanyIDWithEmployee(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))

	comp := compHand.db.Company.Query().Where(company.ID(id)).WithEmployees().AllX(ctx)

	return c.JSON(http.StatusOK, comp)
}

func (compHand *companyHandler) CompanyWithEmployeeID(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))

	comp := compHand.db.Company.Query().WithEmployees(func(emp *ent.EmployeeQuery) {
		emp.Where(employee.ID(id)).AllX(ctx)
	}).AllX(ctx)

	return c.JSON(http.StatusOK, comp)
}

func (compHand *companyHandler) CompanyIDWithEmployeeID(c echo.Context) error {
	ctx := context.Background()
	companyid, _ := strconv.Atoi(c.Param("companyid"))
	employeeid, _ := strconv.Atoi(c.Param("employeeid"))

	comp := compHand.db.Company.Query().Where(company.ID(companyid)).WithEmployees(func(emp *ent.EmployeeQuery) {
		emp.Where(employee.ID(employeeid)).AllX(ctx)
	}).AllX(ctx)

	return c.JSON(http.StatusOK, comp)
}

// Update Edges
func (compHand *companyHandler) UpdateCompanyWithEmployee(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request.CompanyRequest)

	c.Bind(req)

	exist := compHand.db.Company.Query().Where(company.ID(id)).ExistX(ctx)

	if exist == false {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Record Can Not Be Found!",
		})
	}

	name := req.Name
	if name == "" {
		name = compHand.db.Company.Query().Where(company.ID(id)).Select(company.FieldName).StringX(ctx)
	}

	address := req.Address
	if address == "" {
		address = compHand.db.Company.Query().Where(company.ID(id)).Select(company.FieldAddress).StringX(ctx)
	}

	service := req.Service
	if service == "" {
		service = compHand.db.Company.Query().Where(company.ID(id)).Select(company.FieldService).StringX(ctx)
	}

	comp := compHand.db.Company.UpdateOneID(id).SetAddress(address).SetName(name).SetService(service).SaveX(ctx).QueryEmployees().QueryCompanies().WithEmployees().AllX(ctx)

	return c.JSON(http.StatusOK, comp)
}

func (compHand *companyHandler) CompanyWithUpdateEmployee(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))

	req := new(request.EmployeeRequest)

	c.Bind(req)

	exist := compHand.db.Employee.Query().Where(employee.ID(id)).ExistX(ctx)

	if exist == false {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Record Can Not Be Found!",
		})
	}

	company_id := req.CompanyName
	if company_id == "" {
		company_id = compHand.db.Employee.Query().Where(employee.ID(id)).Select(employee.ForeignKeys[0]).StringX(ctx)
	} else {
		company_id = compHand.db.Company.Query().Where(company.Name(req.CompanyName)).Select(company.FieldID).StringX(ctx)
	}
	comp_id, _ := strconv.Atoi(company_id)

	name := req.Name
	if name == "" {
		name = compHand.db.Employee.Query().Where(employee.ID(id)).Select(employee.FieldName).StringX(ctx)
	}

	salary := req.Salary
	if salary == 0 {
		salary = int32(compHand.db.Employee.Query().Where(employee.ID(id)).Select(employee.FieldSalary).IntX(ctx))
	}

	position := req.Position
	if position == "" {
		position = compHand.db.Employee.Query().Where(employee.ID(id)).Select(employee.FieldPosition).StringX(ctx)
	}

	compHand.db.Employee.UpdateOneID(id).SetName(name).SetCompaniesID(comp_id).SetPosition(position).SetSalary(salary).SaveX(ctx)

	comp := compHand.db.Company.Query().WithEmployees(func(emp *ent.EmployeeQuery) {
		emp.Where(employee.ID(id)).AllX(ctx)
	}).AllX(ctx)

	return c.JSON(http.StatusOK, comp)
}

func (compHand *companyHandler) UpdateCompanyWithUpdateEmployee(c echo.Context) error {
	ctx := context.Background()
	company_id, _ := strconv.Atoi(c.Param("companyid"))
	employee_id, _ := strconv.Atoi(c.Param("employeeid"))

	exist := compHand.db.Company.Query().Where(company.ID(company_id)).ExistX(ctx)

	if exist == false {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Company Can Not Be Found!",
		})
	}

	exist = compHand.db.Employee.Query().Where(employee.ID(employee_id)).ExistX(ctx)

	if exist == false {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Employee Can Not Be Found!",
		})
	}

	req := new(request.CompanyEmployeeRequest)
	c.Bind(req)

	employee_name := req.EmployeeName
	if employee_name == "" {
		employee_name = compHand.db.Employee.Query().Where(employee.ID(employee_id)).Select(employee.FieldName).StringX(ctx)
	}

	employee_salary := req.EmployeeSalary
	if employee_salary == 0 {
		employee_salary = int32(compHand.db.Employee.Query().Where(employee.ID(employee_id)).Select(employee.FieldSalary).IntX(ctx))
	}

	employee_position := req.EmployeePosition
	if employee_position == "" {
		employee_position = compHand.db.Employee.Query().Where(employee.ID(employee_id)).Select(employee.FieldPosition).StringX(ctx)
	}

	compHand.db.Employee.UpdateOneID(employee_id).SetCompaniesID(company_id).SetName(employee_name).SetPosition(employee_position).SetSalary(employee_salary).SaveX(ctx)

	company_name := req.CompanyName
	if company_name == "" {
		company_name = compHand.db.Company.Query().Where(company.ID(company_id)).Select(company.FieldName).StringX(ctx)
	}

	company_address := req.CompanyAddress
	if company_address == "" {
		company_address = compHand.db.Company.Query().Where(company.ID(company_id)).Select(company.FieldAddress).StringX(ctx)
	}

	company_service := req.CompanyService
	if company_service == "" {
		company_service = compHand.db.Company.Query().Where(company.ID(company_id)).Select(company.FieldService).StringX(ctx)
	}

	comp := compHand.db.Company.UpdateOneID(company_id).SetAddress(company_address).SetName(company_name).SetService(company_service).SaveX(ctx).QueryEmployees().QueryCompanies().WithEmployees(func(emp *ent.EmployeeQuery) {
		emp.Where(employee.ID(employee_id)).AllX(ctx)
	}).AllX(ctx)

	return c.JSON(http.StatusOK, comp)
}

func (compHand *companyHandler) TestRequestHandler(c echo.Context) error {
	json_map := make(map[string]interface{})
	req := new(request.CompanyRequest)
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	} else {
		req.Name = json_map["company_name"].(string)
		req.Address = json_map["company_service"].(string)
		req.Service = json_map["company_address"].(string)

		if err := c.Validate(req); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, req)
	}
}
