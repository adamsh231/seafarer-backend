package usecase

import (
	"seafarer-backend/api"
	"seafarer-backend/api/user/interfaces"
	"seafarer-backend/api/user/repositories"
	"seafarer-backend/api/user/router/presenters"
	"seafarer-backend/domain/models"
)

type UserUseCase struct {
	*api.Contract
}

func NewUserUseCase(ucContract *api.Contract) interfaces.IUserUseCase {
	return &UserUseCase{ucContract}
}

func (uc UserUseCase) ReadByEmail(email string) (presenter presenters.UserDetailPresenter, err error) {

	// read email
	model := models.NewUser()
	repo := repositories.NewUserRepository(uc.Postgres)
	if err = repo.ReadByEmail(email, model); err != nil {
		api.NewErrorLog("UserUseCase.GetUserByEmail", "repo.GetUserByEmail", err.Error())
		return presenter, err
	}

	presenter = presenters.NewUserDetailPresenter().Build(model)

	return presenter, err
}

func (uc UserUseCase) ChangePassword(password string) (err error) {

	// read email
	repo := repositories.NewUserRepository(uc.Postgres)
	if err = repo.UpdatePasswordByEmail(uc.UserEmail, password, uc.PostgresTX); err != nil {
		api.NewErrorLog("UserUseCase.ChangePassword", "repo.UpdatePasswordByEmail", err.Error())
		return err
	}

	return err
}