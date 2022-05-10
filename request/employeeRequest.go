package request

type EmployeeRequest struct{
	Name string `json:"name" validate:"required"`
	Salary int32 `json:"salary" validate:"required, numeric"`
	Position string `json:"position" validate:"required"`
}
