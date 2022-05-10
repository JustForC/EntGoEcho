package request

type CompanyRequest struct {
	Name    string `json:"name" validate:"required"`
	Service string `json:"service" validate:"required"`
	Address string `json:"address" validate:"required"`
}
