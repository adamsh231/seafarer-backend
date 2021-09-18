package requests

type RegisterRequest struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	CompanyID string `json:"company_id" validate:"required"`
}