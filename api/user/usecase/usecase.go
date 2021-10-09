package usecase

import (
	"seafarer-backend/api"
	"seafarer-backend/api/user/interfaces"
	"seafarer-backend/api/user/repositories"
	"seafarer-backend/api/user/router/presenters"
	"seafarer-backend/api/user/router/requests"
	"seafarer-backend/domain/constants"
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

func (uc UserUseCase) FilterUserAvailable(filter *requests.UsersFilterRequest) (presenter presenters.ArrayFilterUsersPresenter, meta api.MetaResponsePresenter, err error) {

	//init repo
	repoUsers := repositories.NewUserRepository(uc.Postgres)

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(filter.Page, filter.PerPage, filter.Order, filter.Sort)

	//get data filter
	modelUsers, total, err := repoUsers.FilterUserAvailable(offset, limit, orderBy, sort, filter.Search)
	if err != nil {
		api.NewErrorLog("UserUseCase.FilterUserAvailable", "repoUsers.FilterUserAvailable", err.Error())
		return presenter, meta, err
	}

	//build presenter
	presenter = presenters.NewArrayFilterUsersPresenter().Build(modelUsers)

	//set pagination
	meta = uc.Contract.SetPaginationResponse(page, limit, int(total))

	return presenter, meta, err
}

func (uc UserUseCase) FilterCandidate(filter *requests.UsersFilterRequest) (presenter presenters.ArrayFilterUsersPresenter, meta api.MetaResponsePresenter, err error) {

	//init repo
	repoUsers := repositories.NewUserRepository(uc.Postgres)

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(filter.Page, filter.PerPage, filter.Order, filter.Sort)

	//get data filter
	modelUsers, total, err := repoUsers.FilterByStatusRecruitment(offset, limit, orderBy, sort, filter.Search, constants.StatusCandidate)
	if err != nil {
		api.NewErrorLog("UserUseCase.FilterCandidate", "repoUsers.FilterByStatusRecruitment", err.Error())
		return presenter, meta, err
	}

	//build presenter
	presenter = presenters.NewArrayFilterUsersPresenter().Build(modelUsers)

	//set pagination
	meta = uc.Contract.SetPaginationResponse(page, limit, int(total))

	return presenter, meta, err
}

func (uc UserUseCase) FilterEmployee(filter *requests.UsersFilterRequest) (presenter presenters.ArrayFilterUsersPresenter, meta api.MetaResponsePresenter, err error) {

	//init repo
	repoUsers := repositories.NewUserRepository(uc.Postgres)

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(filter.Page, filter.PerPage, filter.Order, filter.Sort)

	//get data filter
	modelUsers, total, err := repoUsers.FilterByStatusRecruitment(offset, limit, orderBy, sort, filter.Search, constants.StatusEmployee)
	if err != nil {
		api.NewErrorLog("UserUseCase.FilterCandidate", "repoUsers.FilterByStatusRecruitment", err.Error())
		return presenter, meta, err
	}

	//build presenter
	presenter = presenters.NewArrayFilterUsersPresenter().Build(modelUsers)

	//set pagination
	meta = uc.Contract.SetPaginationResponse(page, limit, int(total))

	return presenter, meta, err
}

func (uc UserUseCase) FilterLetter(filter *requests.UsersFilterRequest) (presenter presenters.ArrayFilterUsersPresenter, meta api.MetaResponsePresenter, err error) {

	//init repo
	repoUsers := repositories.NewUserRepository(uc.Postgres)

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(filter.Page, filter.PerPage, filter.Order, filter.Sort)

	//get data filter
	modelUsers, total, err := repoUsers.FilterByStatusRecruitment(offset, limit, orderBy, sort, filter.Search, constants.StatusLetter)
	if err != nil {
		api.NewErrorLog("UserUseCase.FilterCandidate", "repoUsers.FilterByStatusRecruitment", err.Error())
		return presenter, meta, err
	}

	//build presenter
	presenter = presenters.NewArrayFilterUsersPresenter().Build(modelUsers)

	//set pagination
	meta = uc.Contract.SetPaginationResponse(page, limit, int(total))

	return presenter, meta, err
}
