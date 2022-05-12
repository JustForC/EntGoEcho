package request

type CompanyEmployeeRequest struct {
	CompanyName      string `json:"company_name" validate:"required"`
	CompanyService   string `json:"company_service" validate:"required"`
	CompanyAddress   string `json:"company_address" validate:"required"`
	EmployeeName     string `json:"employee_name" validate:"required"`
	EmployeeSalary   int32  `json:"employee_salary" validate:"required,number"`
	EmployeePosition string `json:"employee_position" validate:"required"`
}
