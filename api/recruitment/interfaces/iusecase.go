package interfaces

import (
	"seafarer-backend/api"
	"seafarer-backend/api/recruitment/router/requests"
	"seafarer-backend/api/user/router/presenters"
)

type IRecruitmentsUseCase interface {
	AddCandidate(input *requests.CandidateRequest) (err error)

	AddEmployee(input *requests.EmployeeRequest) (err error)

	AddStandByLetter(input *requests.StandByLetterRequest) (err error)

	AddLetter(input *requests.LetterRequest) (err error)

	FilterCandidate(filter *requests.FilterRequest) (presenter presenters.ArrayFilterRecruimentPresenter, meta api.MetaResponsePresenter, err error)

	FilterEmployee(filter *requests.FilterRequest) (presenter presenters.ArrayFilterRecruimentPresenter, meta api.MetaResponsePresenter, err error)

	FilterLetter(filter *requests.FilterRequest) (presenter presenters.ArrayFilterRecruimentPresenter, meta api.MetaResponsePresenter, err error)

	FilterStandbyLetter(filter *requests.FilterRequest) (presenter presenters.ArrayFilterRecruimentPresenter, meta api.MetaResponsePresenter, err error)
}
