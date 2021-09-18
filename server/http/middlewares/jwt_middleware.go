package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"seafarer-backend/api"
	"seafarer-backend/domain/constants"
	"seafarer-backend/domain/constants/messages"
	"seafarer-backend/libraries"
	"strings"
	"time"
)

type JWTMiddleware struct {
	*api.Contract
}

func NewJWTMiddleware(ucContract *api.Contract) JWTMiddleware {
	return JWTMiddleware{ucContract}
}

func (middleware JWTMiddleware) New(ctx *fiber.Ctx) error {

	// validate
	if err := middleware.validate(ctx); err != nil {
		return api.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnauthorized)
	}

	return ctx.Next()
}

func (middleware JWTMiddleware) NonVerifiedOnly(ctx *fiber.Ctx) error {

	// validate
	if err := middleware.validate(ctx); err != nil {
		return api.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnauthorized)
	}

	// user must be not verified yet
	if middleware.Contract.UserIsVerified {
		return api.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, messages.AccountIsVerifiedMessage, nil, http.StatusForbidden)
	}

	// user must be not admin
	if middleware.Contract.UserIsAdmin {
		return api.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, messages.AccountIsNotUserRole, nil, http.StatusForbidden)
	}

	return ctx.Next()
}

func (middleware JWTMiddleware) VerifiedOnly(ctx *fiber.Ctx) error {

	// validate
	if err := middleware.validate(ctx); err != nil {
		return api.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnauthorized)
	}

	// user must be verified
	if !middleware.Contract.UserIsVerified {
		return api.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, messages.AccountIsNotVerifiedMessage, nil, http.StatusForbidden)
	}

	// user must be not admin
	if middleware.Contract.UserIsAdmin {
		return api.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, messages.AccountIsNotUserRole, nil, http.StatusForbidden)
	}

	return ctx.Next()
}

func (middleware JWTMiddleware) validate(ctx *fiber.Ctx) (err error) {

	// check header
	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		return errors.New(messages.TokenIsNotProvidedMessage)
	}

	// check token is valid
	token := strings.Replace(header, "Bearer ", "", -1)
	claims, IsValid := libraries.NewJWTLibrary().ValidateToken(token)
	if !IsValid {
		return errors.New(messages.TokenIsNotValidMessage)
	}

	// check live time
	if expInt, ok := claims[constants.JWTPayloadTokenLiveTime].(float64); ok {
		now := time.Now().Unix()
		if now > int64(expInt) {
			return errors.New(messages.TokenIsExpiredMessage)
		}
	} else {
		return errors.New(messages.InterfaceConversionErrorMessage)
	}

	// insert payload
	if err = middleware.insertPayload(claims); err != nil {
		return err
	}

	return err
}

func (middleware JWTMiddleware) insertPayload(claims jwt.MapClaims) (err error) {

	// id
	if payloadID, ok := claims[constants.JWTPayloadId].(string); ok {
		middleware.Contract.UserID = payloadID
	} else {
		return errors.New(messages.InterfaceConversionErrorMessage)
	}

	// name
	if payloadName, ok := claims[constants.JWTPayloadName].(string); ok {
		middleware.Contract.UserName = payloadName
	} else {
		return errors.New(messages.InterfaceConversionErrorMessage)
	}

	// email
	if payloadEmail, ok := claims[constants.JWTPayloadEmail].(string); ok {
		middleware.Contract.UserEmail = payloadEmail
	} else {
		return errors.New(messages.InterfaceConversionErrorMessage)
	}

	// is verified
	if payloadIsVerified, ok := claims[constants.JWTPayloadIsVerified].(bool); ok {
		middleware.Contract.UserIsVerified = payloadIsVerified
	} else {
		return errors.New(messages.InterfaceConversionErrorMessage)
	}

	// is admin
	if payloadIsAdmin, ok := claims[constants.JWTPayloadIsAdmin].(bool); ok {
		middleware.Contract.UserIsAdmin = payloadIsAdmin
	} else {
		return errors.New(messages.InterfaceConversionErrorMessage)
	}

	return err
}

func (middleware JWTMiddleware) AdminOnly(ctx *fiber.Ctx) error {

	// validate
	if err := middleware.validate(ctx); err != nil {
		return api.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, err.Error(), nil, http.StatusUnauthorized)
	}

	// user must be admin
	if !middleware.Contract.UserIsAdmin {
		return api.NewHandler(middleware.Contract).SendResponseWithoutMeta(ctx, messages.AccountIsNotAdminRole, nil, http.StatusForbidden)
	}

	return ctx.Next()
}

