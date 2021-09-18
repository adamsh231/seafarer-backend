package router

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"seafarer-backend/api"
	"seafarer-backend/api/admin/usecase"
	"seafarer-backend/domain/constants/messages"
)

type AdminHandler struct {
	api.Handler
}

func NewAdminHandler(handler api.Handler) AdminHandler {
	return AdminHandler{handler}
}

func (handler AdminHandler) GetCurrent(ctx *fiber.Ctx) error {

	// database processing
	uc := usecase.NewAdminUseCase(handler.Contract)
	user, err := uc.ReadByEmail(handler.Contract.UserEmail)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, user, http.StatusOK)
}

