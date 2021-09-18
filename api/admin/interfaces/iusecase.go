package interfaces

import "seafarer-backend/api/admin/router/presenters"

type IAdminUseCase interface {
	ReadByEmail(email string) (presenter presenters.AdminDetailPresenter, err error)
}
