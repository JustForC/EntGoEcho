package request

type EmployeeRequest struct {
	Name        string `json:"name" validate:"required"`
	Salary      int32  `json:"salary" validate:"required,number"`
	Position    string `json:"position" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
}
