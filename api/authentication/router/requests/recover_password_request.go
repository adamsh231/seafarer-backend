package requests

type RecoverPasswordRequest struct {
	Password string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
