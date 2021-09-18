package presenters

import "seafarer-backend/domain/models"

type UserDetailPresenter struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	IsVerified bool   `json:"is_verified"`
}

func NewUserDetailPresenter() UserDetailPresenter {
	return UserDetailPresenter{}
}

func (presenter UserDetailPresenter) Build(model *models.User) UserDetailPresenter {
	return UserDetailPresenter{
		Name:       model.Name,
		Email:      model.Email,
		IsVerified: model.IsVerified,
	}
}
