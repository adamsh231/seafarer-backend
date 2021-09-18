package interfaces

import "seafarer-backend/api/user/router/presenters"

type IUserUseCase interface {
	ReadByEmail(email string) (presenter presenters.UserDetailPresenter, err error)

	ChangePassword(password string) (err error)
}
