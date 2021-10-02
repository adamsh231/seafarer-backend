package presenters

import (
	"seafarer-backend/domain/models"
)

type ArrayFilterUsersPresenter struct {
	FilterUsersPresenter []FilterUsersPresenter
}

type FilterUsersPresenter struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewArrayFilterUsersPresenter() ArrayFilterUsersPresenter {
	return ArrayFilterUsersPresenter{}
}

func (presenter ArrayFilterUsersPresenter) Build(model []models.User) (list ArrayFilterUsersPresenter) {
	for _, row := range model {
		list.FilterUsersPresenter = append(list.FilterUsersPresenter, FilterUsersPresenter{
			ID:    row.ID,
			Name:  row.Name,
			Email: row.Email,
		})
	}
	return list
}
