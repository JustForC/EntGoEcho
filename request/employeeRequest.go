package request

type EmployeeRequest struct {
	Name        string `json:"employee_name" validate:"required"`
	Salary      int32  `json:"employee_salary" validate:"required,number"`
	Position    string `json:"employee_position" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
}
