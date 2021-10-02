package interfaces

import (
	"seafarer-backend/api/recruitment/router/requests"
)

type IRecruitmentsUseCase interface {
	AddCandidate(input *requests.CandidateRequest) (err error)
}
