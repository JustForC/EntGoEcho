package request

type InputRequest struct {
	Check_one string `json:"check_one" validate:"required"`
}
