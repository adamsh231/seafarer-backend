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

func (handler RecruitmentsHandler) CreateEmployee(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(requests.EmployeeRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// database processing
	handler.Contract.PostgresTX = handler.Contract.Postgres.Begin()
	uc := usecase.NewRecruitmentsUseCase(handler.Contract)
	err := uc.AddEmployee(input)
	if err != nil {
		handler.Contract.PostgresTX.Rollback()
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	handler.Contract.PostgresTX.Commit()

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, nil, http.StatusCreated)
}

func (handler RecruitmentsHandler) CreateStandByLetter(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(requests.StandByLetterRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// database processing
	handler.Contract.PostgresTX = handler.Contract.Postgres.Begin()
	uc := usecase.NewRecruitmentsUseCase(handler.Contract)
	err := uc.AddStandByLetter(input)
	if err != nil {
		handler.Contract.PostgresTX.Rollback()
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	handler.Contract.PostgresTX.Commit()

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, nil, http.StatusCreated)
}

func (handler RecruitmentsHandler) CreateLetter(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(requests.LetterRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// database processing
	handler.Contract.PostgresTX = handler.Contract.Postgres.Begin()
	uc := usecase.NewRecruitmentsUseCase(handler.Contract)
	err := uc.AddLetter(input)
	if err != nil {
		handler.Contract.PostgresTX.Rollback()
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	handler.Contract.PostgresTX.Commit()

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, nil, http.StatusCreated)
}

func (handler RecruitmentsHandler) FilterCandidate(ctx *fiber.Ctx) error {
	filter := new(requests.FilterRequest)

	//data parsing
	if ctx.QueryParser(filter) != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedBindQuery, nil, http.StatusBadRequest)
	}

	//data validation
	if err := handler.Contract.Validator.Struct(filter); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), filter, http.StatusBadRequest)
	}

	//database proccesing
	uc := usecase.NewRecruitmentsUseCase(handler.Contract)
	presenter, meta, err := uc.FilterCandidate(filter)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedMessage, nil, http.StatusBadRequest)
	}

	return handler.SendResponseWithMeta(ctx, presenter.FilterRecruimentPresenter, messages.SuccessMessage, meta, http.StatusOK)
}

func (handler RecruitmentsHandler) FilterEmployee(ctx *fiber.Ctx) error {
	filter := new(requests.FilterRequest)

	//data parsing
	if ctx.QueryParser(filter) != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedBindQuery, nil, http.StatusBadRequest)
	}

	//data validation
	if err := handler.Contract.Validator.Struct(filter); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), filter, http.StatusBadRequest)
	}

	//database proccesing
	uc := usecase.NewRecruitmentsUseCase(handler.Contract)
	presenter, meta, err := uc.FilterEmployee(filter)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedMessage, nil, http.StatusBadRequest)
	}

	return handler.SendResponseWithMeta(ctx, presenter.FilterRecruimentPresenter, messages.SuccessMessage, meta, http.StatusOK)
}

func (handler RecruitmentsHandler) FilterLetter(ctx *fiber.Ctx) error {
	filter := new(requests.FilterRequest)

	//data parsing
	if ctx.QueryParser(filter) != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedBindQuery, nil, http.StatusBadRequest)
	}

	//data validation
	if err := handler.Contract.Validator.Struct(filter); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), filter, http.StatusBadRequest)
	}

	//database proccesing
	uc := usecase.NewRecruitmentsUseCase(handler.Contract)
	presenter, meta, err := uc.FilterLetter(filter)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedMessage, nil, http.StatusBadRequest)
	}

	return handler.SendResponseWithMeta(ctx, presenter.FilterRecruimentPresenter, messages.SuccessMessage, meta, http.StatusOK)
}

func (handler RecruitmentsHandler) FilterStandByLetter(ctx *fiber.Ctx) error {
	filter := new(requests.FilterRequest)

	//data parsing
	if ctx.QueryParser(filter) != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedBindQuery, nil, http.StatusBadRequest)
	}

	//data validation
	if err := handler.Contract.Validator.Struct(filter); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), filter, http.StatusBadRequest)
	}

	//database proccesing
	uc := usecase.NewRecruitmentsUseCase(handler.Contract)
	presenter, meta, err := uc.FilterStandbyLetter(filter)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedMessage, nil, http.StatusBadRequest)
	}

	return handler.SendResponseWithMeta(ctx, presenter.FilterRecruimentPresenter, messages.SuccessMessage, meta, http.StatusOK)
}
