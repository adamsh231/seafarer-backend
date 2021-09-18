package requests

type OTPEmailRecoverRequest struct {
	Email string `json:"email" validate:"required"`
}
