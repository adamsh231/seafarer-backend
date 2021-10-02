package requests

type CandidateRequest struct {
	UserID       string  `json:"user_id" validate:"required"`
	ExpectSalary float64 `json:"expect_salary" validate:"required"`
}

type EmployeeRequest struct {
	UserID   string  `json:"user_id" validate:"required"`
	Salary   float64 `json:"salary" validate:"required"`
	Position string  `json:"position" validate:"required"`
}
