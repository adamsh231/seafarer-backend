package usecase

import (
	"seafarer-backend/api"
	"seafarer-backend/api/recruitment/interfaces"
	"seafarer-backend/api/recruitment/repositories"
	"seafarer-backend/api/recruitment/router/requests"
	repositoriesUser "seafarer-backend/api/user/repositories"
	"seafarer-backend/api/user/router/presenters"
	"seafarer-backend/domain/constants"
	"seafarer-backend/domain/models"
	"time"

	"github.com/google/uuid"
)

type RecruitmentsUseCase struct {
	*api.Contract
}

func NewRecruitmentsUseCase(ucContract *api.Contract) interfaces.IRecruitmentsUseCase {
	return &RecruitmentsUseCase{ucContract}
}

func (uc RecruitmentsUseCase) AddCandidate(input *requests.CandidateRequest) (err error) {

	// init
	now := time.Now()
	id := uuid.New().String()
	model := &models.Recruitments{
		ID:           id,
		UserID:       input.UserID,
		CreatedAt:    now,
		UpdatedAt:    now,
		ExpectSalary: input.ExpectSalary,
		Position:     input.Position,
		Status:       constants.StatusCandidate,
	}

	// save new candidate
	repo := repositories.NewRecruitmentsRepository(uc.Postgres)
	if err = repo.Add(model, uc.PostgresTX); err != nil {
		api.NewErrorLog("RecruitmentsUseCase.AddCandidate", "repositories.Add", err.Error())
		return err
	}

	// set recruitment_id in user
	repoUser := repositoriesUser.NewUserRepository(uc.Postgres)
	modelUser := models.User{
		RecruitmentID: id,
	}
	if err = repoUser.Update(input.UserID, modelUser, uc.PostgresTX); err != nil {
		api.NewErrorLog("RecruitmentsUseCase.AddCandidate", "repositories.Add", err.Error())
		return err
	}

	return err
}

func (uc RecruitmentsUseCase) AddEmployee(input *requests.EmployeeRequest) (err error) {

	// init
	now := time.Now()
	model := models.Recruitments{
		UpdatedAt: now,
		Salary:    input.Salary,
		SignOn:    input.SignOn,
		Status:    constants.StatusEmployee,
	}

	// save not verified user
	repo := repositories.NewRecruitmentsRepository(uc.Postgres)
	if err = repo.UpdateByIDUser(input.UserID, model, uc.PostgresTX); err != nil {
		api.NewErrorLog("RecruitmentsUseCase.AddEmployee", "repositories.UpdateByIDUser", err.Error())
		return err
	}

	return err
}

func (uc RecruitmentsUseCase) AddStandByLetter(input *requests.StandByLetterRequest) (err error) {

	// init
	now := time.Now()
	model := models.Recruitments{
		UpdatedAt: now,
		UserID:    input.Ship,
		Status:    constants.StatusStandbyLetter,
	}

	// save not verified user
	repo := repositories.NewRecruitmentsRepository(uc.Postgres)
	if err = repo.UpdateByIDUser(input.UserID, model, uc.PostgresTX); err != nil {
		api.NewErrorLog("RecruitmentsUseCase.AddEmployee", "repositories.UpdateByIDUser", err.Error())
		return err
	}

	return err
}

func (uc RecruitmentsUseCase) AddLetter(input *requests.LetterRequest) (err error) {

	// init
	now := time.Now()
	model := models.Recruitments{
		UpdatedAt: now,
		Letter:    input.Letter,
		Status:    constants.StatusLetter,
	}

	// save not verified user
	repo := repositories.NewRecruitmentsRepository(uc.Postgres)
	if err = repo.UpdateByIDUser(input.UserID, model, uc.PostgresTX); err != nil {
		api.NewErrorLog("RecruitmentsUseCase.AddEmployee", "repositories.UpdateByIDUser", err.Error())
		return err
	}

	return err
}

func (uc RecruitmentsUseCase) FilterCandidate(filter *requests.FilterRequest) (presenter presenters.ArrayFilterRecruimentPresenter, meta api.MetaResponsePresenter, err error) {

	//init repo
	repoRecruitment := repositories.NewRecruitmentsRepository(uc.Postgres)

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(filter.Page, filter.PerPage, filter.Order, filter.Sort)

	//get data filter
	modelRecruitment, total, err := repoRecruitment.FilterByStatusRecruitment(offset, limit, orderBy, sort, filter.Search, constants.StatusCandidate)
	if err != nil {
		api.NewErrorLog("UserUseCase.FilterCandidate", "repoRecruitment.FilterByStatusRecruitment", err.Error())
		return presenter, meta, err
	}

	//build presenter
	presenter = presenters.NewArrayFilterRecruimentPresenter().Build(modelRecruitment)

	//set pagination
	meta = uc.Contract.SetPaginationResponse(page, limit, int(total))

	return presenter, meta, err
}

func (uc RecruitmentsUseCase) FilterEmployee(filter *requests.FilterRequest) (presenter presenters.ArrayFilterRecruimentPresenter, meta api.MetaResponsePresenter, err error) {

	//init repo
	repoRecruitment := repositories.NewRecruitmentsRepository(uc.Postgres)

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(filter.Page, filter.PerPage, filter.Order, filter.Sort)

	//get data filter
	modelRecruitment, total, err := repoRecruitment.FilterByStatusRecruitment(offset, limit, orderBy, sort, filter.Search, constants.StatusEmployee)
	if err != nil {
		api.NewErrorLog("UserUseCase.FilterEmployee", "repoRecruitment.FilterByStatusRecruitment", err.Error())
		return presenter, meta, err
	}

	//build presenter
	presenter = presenters.NewArrayFilterRecruimentPresenter().Build(modelRecruitment)

	//set pagination
	meta = uc.Contract.SetPaginationResponse(page, limit, int(total))

	return presenter, meta, err
}

func (uc RecruitmentsUseCase) FilterLetter(filter *requests.FilterRequest) (presenter presenters.ArrayFilterRecruimentPresenter, meta api.MetaResponsePresenter, err error) {

	//init repo
	repoRecruitment := repositories.NewRecruitmentsRepository(uc.Postgres)

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(filter.Page, filter.PerPage, filter.Order, filter.Sort)

	//get data filter
	modelRecruitment, total, err := repoRecruitment.FilterByStatusRecruitment(offset, limit, orderBy, sort, filter.Search, constants.StatusLetter)
	if err != nil {
		api.NewErrorLog("UserUseCase.FilterLetter", "repoRecruitment.FilterByStatusRecruitment", err.Error())
		return presenter, meta, err
	}

	//build presenter
	presenter = presenters.NewArrayFilterRecruimentPresenter().Build(modelRecruitment)

	//set pagination
	meta = uc.Contract.SetPaginationResponse(page, limit, int(total))

	return presenter, meta, err
}

func (uc RecruitmentsUseCase) FilterStandbyLetter(filter *requests.FilterRequest) (presenter presenters.ArrayFilterRecruimentPresenter, meta api.MetaResponsePresenter, err error) {

	//init repo
	repoRecruitment := repositories.NewRecruitmentsRepository(uc.Postgres)

	//set pagination
	offset, limit, page, orderBy, sort := uc.SetPaginationParameter(filter.Page, filter.PerPage, filter.Order, filter.Sort)

	//get data filter
	modelRecruitment, total, err := repoRecruitment.FilterByStatusRecruitment(offset, limit, orderBy, sort, filter.Search, constants.StatusStandbyLetter)
	if err != nil {
		api.NewErrorLog("UserUseCase.FilterStandbyLetter", "repoRecruitment.FilterByStatusRecruitment", err.Error())
		return presenter, meta, err
	}

	//build presenter
	presenter = presenters.NewArrayFilterRecruimentPresenter().Build(modelRecruitment)

	//set pagination
	meta = uc.Contract.SetPaginationResponse(page, limit, int(total))

	return presenter, meta, err
}
