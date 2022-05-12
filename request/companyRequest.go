package request

type CompanyRequest struct {
	Name    string `json:"company_name" validate:"required"`
	Service string `json:"company_service" validate:"required"`
	Address string `json:"company_address" validate:"required"`
}
