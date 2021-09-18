package requests

type OTPRecoverRequest struct {
	Email string `json:"email" validate:"required"`
	OTP   string `json:"otp" validate:"required"`
}
