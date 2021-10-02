package interfaces

import (
	"seafarer-backend/api"
	"seafarer-backend/api/user/router/presenters"
	"seafarer-backend/api/user/router/requests"
)

type IUserUseCase interface {
	ReadByEmail(email string) (presenter presenters.UserDetailPresenter, err error)

	ChangePassword(password string) (err error)

	Filter(filter *requests.UsersFilterRequest) (presenter presenters.ArrayFilterUsersPresenter, meta api.MetaResponsePresenter, err error)
}
