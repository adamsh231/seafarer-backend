package api

import "github.com/gofiber/fiber/v2"

type Handler struct {
	Contract *Contract
}

type ResponsePresenter struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type MetaResponsePresenter struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
}

func NewHandler(ucContract *Contract) Handler {
	return Handler{Contract: ucContract}
}

func (handler Handler) SendResponseWithoutMeta(ctx *fiber.Ctx, message string, data interface{}, httpStatus int) error {
	return ctx.Status(httpStatus).JSON(ResponsePresenter{
		Message: message,
		Data:    data,
	})
}

func (handler Handler) SendResponseWithMeta(ctx *fiber.Ctx, data interface{}, message string, meta interface{}, httpStatus int) error {
	return ctx.Status(httpStatus).JSON(ResponsePresenter{
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}
