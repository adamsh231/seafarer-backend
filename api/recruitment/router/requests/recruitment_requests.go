package requests

type CandidateRequest struct {
	UserID        string  `json:"user_id" validate:"required"`
	ExpectSallary float64 `json:"expect_sallary" validate:"required"`
}
