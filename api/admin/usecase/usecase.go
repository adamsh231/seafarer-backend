package usecase

import (
	"seafarer-backend/api"
	"seafarer-backend/api/admin/interfaces"
	"seafarer-backend/api/admin/repositories"
	"seafarer-backend/api/admin/router/presenters"
	"seafarer-backend/domain/models"
)

type AdminUseCase struct {
	*api.Contract
}

func NewAdminUseCase(ucContract *api.Contract) interfaces.IAdminUseCase {
	return &AdminUseCase{ucContract}
}

func (uc AdminUseCase) ReadByEmail(email string) (presenter presenters.AdminDetailPresenter, err error) {

	// read email
	model := models.NewAdmin()
	repo := repositories.NewAdminRepository(uc.Postgres)
	if err = repo.ReadByEmail(email, model); err != nil {
		api.NewErrorLog("AdminUseCase.ReadByEmail", "repo.ReadByEmail", err.Error())
		return presenter, err
	}

	presenter = presenters.NewAdminDetailPresenter().Build(model)

	return presenter, err
}

