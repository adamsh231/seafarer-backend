package router

import (
	"net/http"
	"seafarer-backend/api"
	"seafarer-backend/api/storage/router/requests"
	"seafarer-backend/api/storage/usecase"
	"seafarer-backend/domain/constants/messages"

	"github.com/gofiber/fiber/v2"
)

type FileHandler struct {
	api.Handler
}

func NewFileHandler(handler api.Handler) FileHandler {
	return FileHandler{handler}
}

func (handler FileHandler) Add(ctx *fiber.Ctx) error {

	//get body form file and form value
	file, err := ctx.FormFile("file")
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// database processing
	handler.Contract.PostgresTX = handler.Contract.Postgres.Begin()
	uc := usecase.NewFileUseCase(handler.Contract)
	if err := uc.Add(file); err != nil {
		handler.Contract.PostgresTX.Rollback()
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}
	handler.Contract.PostgresTX.Commit()

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, nil, http.StatusCreated)
}

func (handler FileHandler) GetPresignedKey(ctx *fiber.Ctx) error {

	// get param
	fileID := ctx.Params("id")

	// database processing
	uc := usecase.NewFileUseCase(handler.Contract)
	file, err := uc.GetWithPresignedKey(fileID)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, file, http.StatusCreated)
}

func (handler FileHandler) BrowsePresignedKey(ctx *fiber.Ctx) error {

	// get param
	req := new(requests.BrowseFilesRequest)
	if err := ctx.QueryParser(req); err != nil {
		return err
	}

	// database processing
	uc := usecase.NewFileUseCase(handler.Contract)
	presenter, meta, err := uc.BrowseWithPresignedKey(req)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithMeta(ctx, presenter.ListData, messages.SuccessMessage, meta, http.StatusOK)
}

func (handler FileHandler) Delete(ctx *fiber.Ctx) error {

	// get param
	fileID := ctx.Params("id")

	// database processing
	handler.Contract.PostgresTX = handler.Contract.Postgres.Begin()
	uc := usecase.NewFileUseCase(handler.Contract)
	err := uc.Delete(fileID)
	if err != nil {
		handler.Contract.PostgresTX.Rollback()
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	handler.Contract.PostgresTX.Commit()

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, nil, http.StatusOK)
}

func (handler FileHandler) Upload(ctx *fiber.Ctx) error {
	userID := ctx.Query("userID")
	handler.Contract.UserID = userID
	//get body form file and form value
	file, err := ctx.FormFile("file")
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusBadRequest)
	}

	// database processing
	uc := usecase.NewFileUseCase(handler.Contract)
	urlPresigned, err := uc.PrivateUploadAndGetPresignedKey(file)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithoutMeta(ctx, messages.SuccessMessage, urlPresigned, http.StatusCreated)
}

func (handler FileHandler) BrowseByUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")

	// get param
	req := new(requests.BrowseFilesRequest)
	if err := ctx.QueryParser(req); err != nil {
		return err
	}

	// database processing
	uc := usecase.NewFileUseCase(handler.Contract)
	presenter, meta, err := uc.BrowseByUserID(userID, req)
	if err != nil {
		return handler.SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnprocessableEntity)
	}

	return handler.SendResponseWithMeta(ctx, presenter.ListData, messages.SuccessMessage, meta, http.StatusOK)
}
