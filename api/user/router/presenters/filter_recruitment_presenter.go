package presenters

import (
	"seafarer-backend/domain/models"
	"time"
)

type ArrayFilterRecruimentPresenter struct {
	FilterRecruimentPresenter []FilterRecruimentPresenter
}

type FilterRecruimentPresenter struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserName     string    `json:"name"`
	ExpectSalary float64   `json:"expect_salary"`
	Salary       float64   `json:"salary"`
	Position     string    `json:"position"`
	AvailableOn  time.Time `json:"available_on"`
	SignOn       time.Time `json:"sign_on"`
	Ship         string    `json:"ship"`
	Letter       string    `json:"letter"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsFailed     bool      `json:"is_failed"`
}

func NewArrayFilterRecruimentPresenter() ArrayFilterRecruimentPresenter {
	return ArrayFilterRecruimentPresenter{}
}

func (presenter ArrayFilterRecruimentPresenter) Build(model []models.RecruitmentsDetail) (list ArrayFilterRecruimentPresenter) {
	for _, row := range model {
		list.FilterRecruimentPresenter = append(list.FilterRecruimentPresenter, FilterRecruimentPresenter{
			ID:           row.ID,
			UserID:       row.UserID,
			UserName:     row.UserName,
			ExpectSalary: row.ExpectSalary,
			Salary:       row.Salary,
			Position:     row.Position,
			AvailableOn:  row.AvailableOn,
			SignOn:       row.SignOn,
			Ship:         row.Ship,
			Letter:       row.Letter,
			Status:       row.Status,
			CreatedAt:    row.CreatedAt,
			UpdatedAt:    row.UpdatedAt,
			IsFailed:     row.IsFailed,
		})
	}
	return list
}
