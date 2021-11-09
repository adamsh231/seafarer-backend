package requests

import "time"

type CandidateRequest struct {
	UserID       string  `json:"user_id" validate:"required"`
	ExpectSalary float64 `json:"expect_salary" validate:"required"`
	Position     string  `json:"position" validate:"required"`
}

type EmployeeRequest struct {
	UserID string    `json:"user_id" validate:"required"`
	Salary float64   `json:"salary" validate:"required"`
	SignOn time.Time `json:"sign_on" validate:"required"`
}

type StandByLetterRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Ship   string `json:"ship" validate:"required"`
}

type LetterRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Letter string `json:"letter" validate:"required"`
}

type FilterRequest struct {
	PerPage int    `query:"per_page, omitempty"`
	Page    int    `query:"page, omitempty"`
	Search  string `query:"search, omitempty"`
	Order   string `query:"order, omitempty"`
	Sort    string `query:"sort, omitempty"`
}
