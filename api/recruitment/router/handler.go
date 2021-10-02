package router

import (
	"net/http"
	"seafarer-backend/api"
	"seafarer-backend/api/recruitment/router/requests"
	"seafarer-backend/api/recruitment/usecase"
	"seafarer-backend/domain/constants/messages"

	"github.com/gofiber/fiber/v2"
)

type RecruitmentsHandler struct {
	api.Handler
}

func NewRecruitmentsHandler(handler api.Handler) RecruitmentsHandler {
	return RecruitmentsHandler{handler}
}

func (handler RecruitmentsHandler) CreateCandidate(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(requests.CandidateRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// database processing
	handler.Contract.PostgresTX = handler.Contract.Postgres.Begin()
	uc := usecase.NewRecruitmentsUseCase(handler.Contract)
	err := uc.AddCandidate(input)
	if err != nil {
		handler.Contract.PostgresTX.Rollback()
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	handler.Contract.PostgresTX.Commit()

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, nil, http.StatusCreated)
}
