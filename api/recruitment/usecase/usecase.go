package usecase

import (
	"seafarer-backend/api"
	"seafarer-backend/api/recruitment/interfaces"
	"seafarer-backend/api/recruitment/repositories"
	"seafarer-backend/api/recruitment/router/requests"
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
		IsFailed:     false,
		Status:       constants.StatusCandidate,
	}

	// save not verified user
	repo := repositories.NewRecruitmentsRepository(uc.Postgres)
	if err = repo.Add(model, uc.PostgresTX); err != nil {
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
