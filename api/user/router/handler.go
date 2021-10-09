package router

import (
	"net/http"
	"seafarer-backend/api"
	"seafarer-backend/api/user/router/requests"
	"seafarer-backend/api/user/usecase"
	"seafarer-backend/domain/constants/messages"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	api.Handler
}

func NewUserHandler(handler api.Handler) UserHandler {
	return UserHandler{handler}
}

func (handler UserHandler) GetCurrentUser(ctx *fiber.Ctx) error {

	// database processing
	uc := usecase.NewUserUseCase(handler.Contract)
	user, err := uc.ReadByEmail(handler.Contract.UserEmail)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, user, http.StatusOK)
}

func (handler UserHandler) Filter(ctx *fiber.Ctx) error {
	filter := new(requests.UsersFilterRequest)

	//data parsing
	if ctx.QueryParser(filter) != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedBindQuery, nil, http.StatusBadRequest)
	}

	//data validation
	if err := handler.Contract.Validator.Struct(filter); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), filter, http.StatusBadRequest)
	}

	//database proccesing
	uc := usecase.NewUserUseCase(handler.Contract)
	presenter, meta, err := uc.Filter(filter)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedMessage, nil, http.StatusBadRequest)
	}

	return handler.SendResponseWithMeta(ctx, presenter.FilterUsersPresenter, messages.SuccessMessage, meta, http.StatusOK)
}

func (handler UserHandler) FilterUserAvailable(ctx *fiber.Ctx) error {
	filter := new(requests.UsersFilterRequest)

	//data parsing
	if ctx.QueryParser(filter) != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedBindQuery, nil, http.StatusBadRequest)
	}

	//data validation
	if err := handler.Contract.Validator.Struct(filter); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), filter, http.StatusBadRequest)
	}

	//database proccesing
	uc := usecase.NewUserUseCase(handler.Contract)
	presenter, meta, err := uc.FilterUserAvailable(filter)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedMessage, nil, http.StatusBadRequest)
	}

	return handler.SendResponseWithMeta(ctx, presenter.FilterUsersPresenter, messages.SuccessMessage, meta, http.StatusOK)
}

func (handler UserHandler) FilterCandidate(ctx *fiber.Ctx) error {
	filter := new(requests.UsersFilterRequest)

	//data parsing
	if ctx.QueryParser(filter) != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedBindQuery, nil, http.StatusBadRequest)
	}

	//data validation
	if err := handler.Contract.Validator.Struct(filter); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), filter, http.StatusBadRequest)
	}

	//database proccesing
	uc := usecase.NewUserUseCase(handler.Contract)
	presenter, meta, err := uc.FilterCandidate(filter)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedMessage, nil, http.StatusBadRequest)
	}

	return handler.SendResponseWithMeta(ctx, presenter.FilterUsersPresenter, messages.SuccessMessage, meta, http.StatusOK)
}

func (handler UserHandler) FilterEmployee(ctx *fiber.Ctx) error {
	filter := new(requests.UsersFilterRequest)

	//data parsing
	if ctx.QueryParser(filter) != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedBindQuery, nil, http.StatusBadRequest)
	}

	//data validation
	if err := handler.Contract.Validator.Struct(filter); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), filter, http.StatusBadRequest)
	}

	//database proccesing
	uc := usecase.NewUserUseCase(handler.Contract)
	presenter, meta, err := uc.FilterEmployee(filter)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedMessage, nil, http.StatusBadRequest)
	}

	return handler.SendResponseWithMeta(ctx, presenter.FilterUsersPresenter, messages.SuccessMessage, meta, http.StatusOK)
}

func (handler UserHandler) FilterLetter(ctx *fiber.Ctx) error {
	filter := new(requests.UsersFilterRequest)

	//data parsing
	if ctx.QueryParser(filter) != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedBindQuery, nil, http.StatusBadRequest)
	}

	//data validation
	if err := handler.Contract.Validator.Struct(filter); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), filter, http.StatusBadRequest)
	}

	//database proccesing
	uc := usecase.NewUserUseCase(handler.Contract)
	presenter, meta, err := uc.FilterLetter(filter)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, messages.FailedMessage, nil, http.StatusBadRequest)
	}

	return handler.SendResponseWithMeta(ctx, presenter.FilterUsersPresenter, messages.SuccessMessage, meta, http.StatusOK)
}
