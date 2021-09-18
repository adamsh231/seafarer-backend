package router

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"seafarer-backend/api"
	"seafarer-backend/api/authentication/router/requests"
	"seafarer-backend/api/authentication/usecase"
	"seafarer-backend/domain/constants/messages"
)

type AuthenticationHandler struct {
	api.Handler
}

func NewAuthHandler(handler api.Handler) AuthenticationHandler {
	return AuthenticationHandler{handler}
}

func (handler AuthenticationHandler) Login(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(requests.LoginRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// login and get jwt
	uc := usecase.NewAuthenticationUseCase(handler.Contract)
	jwt, err := uc.Login(input)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, jwt, http.StatusOK)
}