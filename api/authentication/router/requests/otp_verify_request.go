package requests

type OTPVerify struct {
	OTP string `json:"otp" validate:"required"`
}
