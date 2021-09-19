package router

import (
	"net/http"
	"seafarer-backend/api"
	"seafarer-backend/api/document/usecase"
	"seafarer-backend/domain/constants/messages"
	"seafarer-backend/domain/models"

	"github.com/gofiber/fiber/v2"
)

type AfeHandler struct {
	api.Handler
}

func NewAfeHandler(handler api.Handler) AfeHandler {
	return AfeHandler{handler}
}

func (handler AfeHandler) Save(ctx *fiber.Ctx) error {

	// get & validate param
	input := new(models.AFE)
	if err := ctx.BodyParser(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}
	if err := handler.Contract.Validator.Struct(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// database processing
	uc := usecase.NewAFEUseCase(handler.Contract)
	if err := uc.Save(input); err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, nil, http.StatusOK)
}

func (handler AfeHandler) Get(ctx *fiber.Ctx) error {

	// database processing
	uc := usecase.NewAFEUseCase(handler.Contract)
	afe, err := uc.Get()
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, afe, http.StatusOK)
}

func (handler AfeHandler) Download(ctx *fiber.Ctx) error {

	// database processing
	uc := usecase.NewAFEUseCase(handler.Contract)
	url, err := uc.DownloadAFE()
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, url, http.StatusOK)
}
