package router

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"seafarer-backend/api"
	"seafarer-backend/api/user/usecase"
	"seafarer-backend/domain/constants/messages"
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
	if err != nil{
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, user, http.StatusOK)
}


