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

	FilterCandidate(filter *requests.UsersFilterRequest) (presenter presenters.ArrayFilterUsersPresenter, meta api.MetaResponsePresenter, err error)

	FilterEmployee(filter *requests.UsersFilterRequest) (presenter presenters.ArrayFilterUsersPresenter, meta api.MetaResponsePresenter, err error)

	FilterLetter(filter *requests.UsersFilterRequest) (presenter presenters.ArrayFilterUsersPresenter, meta api.MetaResponsePresenter, err error)

	FilterUserAvailable(filter *requests.UsersFilterRequest) (presenter presenters.ArrayFilterUsersPresenter, meta api.MetaResponsePresenter, err error)
}
