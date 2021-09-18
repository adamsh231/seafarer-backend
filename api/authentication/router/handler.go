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

func (handler AuthenticationHandler) Register(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(requests.RegisterRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// database processing
	handler.Contract.PostgresTX = handler.Contract.Postgres.Begin()
	uc := usecase.NewAuthenticationUseCase(handler.Contract)
	jwt, err := uc.Register(input)
	if err != nil {
		handler.Contract.PostgresTX.Rollback()
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	handler.Contract.PostgresTX.Commit()

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, jwt, http.StatusCreated)
}

func (handler AuthenticationHandler) SendEmailOTPVerify(ctx *fiber.Ctx) error {

	// database processing
	uc := usecase.NewAuthenticationUseCase(handler.Contract)
	err := uc.SendEmailOTPVerify()
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, nil, http.StatusOK)
}

func (handler AuthenticationHandler) VerifyOTP(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(requests.OTPVerify)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// database processing
	handler.Contract.PostgresTX = handler.Contract.Postgres.Begin()
	uc := usecase.NewAuthenticationUseCase(handler.Contract)
	jwt, err := uc.OTPVerify(input)
	if err != nil {
		handler.Contract.PostgresTX.Rollback()
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	handler.Contract.PostgresTX.Commit()

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, jwt, http.StatusOK)
}

func (handler AuthenticationHandler) SendEmailOTPRecover(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(requests.OTPEmailRecoverRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// email processing
	uc := usecase.NewAuthenticationUseCase(handler.Contract)
	if err := uc.SendEmailOTPRecover(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, nil, http.StatusOK)

}

func (handler AuthenticationHandler) RecoverOTP(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(requests.OTPRecoverRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// recover processing
	uc := usecase.NewAuthenticationUseCase(handler.Contract)
	jwt, err := uc.OTPRecover(input)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, jwt, http.StatusOK)

}

func (handler AuthenticationHandler) RecoverChangePassword(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(requests.RecoverPasswordRequest)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// password vs confirm password
	if input.Password != input.ConfirmPassword {
		return handler.SendResponseWithoutMeta(ctx, messages.PasswordAndConfirmPasswordNotSameMessage, nil, http.StatusUnprocessableEntity)
	}

	// database processing
	handler.Contract.PostgresTX = handler.Contract.Postgres.Begin()
	uc := usecase.NewAuthenticationUseCase(handler.Contract)
	err := uc.ChangePasswordRecover(input)
	if err != nil {
		handler.Contract.PostgresTX.Rollback()
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	handler.Contract.PostgresTX.Commit()

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, nil, http.StatusOK)
}

func (handler AuthenticationHandler) LoginAdmin(ctx *fiber.Ctx) error {

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
	jwt, err := uc.LoginAdmin(input)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, jwt, http.StatusOK)
}
