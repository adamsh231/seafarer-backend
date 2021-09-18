package presenters

import "seafarer-backend/domain/models"

type AdminDetailPresenter struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewAdminDetailPresenter() AdminDetailPresenter {
	return AdminDetailPresenter{}
}

func (presenter AdminDetailPresenter) Build(model *models.Admin) AdminDetailPresenter {
	return AdminDetailPresenter{
		Name:  model.Name,
		Email: model.Email,
	}
}
