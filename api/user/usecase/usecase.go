package usecase

import (
	"seafarer-backend/api"
	"seafarer-backend/api/user/interfaces"
	"seafarer-backend/api/user/repositories"
	"seafarer-backend/api/user/router/presenters"
	"seafarer-backend/api/user/router/requests"
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

func (uc UserUseCase) Filter(filter *requests.UsersFilterRequest) (presenter presenters.ArrayFilterUsersPresenter, meta api.MetaResponsePresenter, err error) {

	//init repo
	repoUsers := repositories.NewUserRepository(uc.Postgres)

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(filter.Page, filter.PerPage, filter.Order, filter.Sort)

	//get data filter
	modelUsers, total, err := repoUsers.Filter(offset, limit, orderBy, sort, filter.Search)
	if err != nil {
		api.NewErrorLog("UserUseCase.Filter", "repoUsers.Filter", err.Error())
		return presenter, meta, err
	}

	//build presenter
	presenter = presenters.NewArrayFilterUsersPresenter().Build(modelUsers)

	//set pagination
	meta = uc.Contract.SetPaginationResponse(page, limit, int(total))

	return presenter, meta, err
}
