package router

import (
	"github.com/gofiber/fiber/v2"
	"seafarer-backend/api"
	userHandlers "seafarer-backend/api/user/router"
	adminHandlers "seafarer-backend/api/admin/router"

	"seafarer-backend/server/http/middlewares"
)

type AuthenticationRoute struct {
	RouteGroup fiber.Router
	Handler    api.Handler
}

func NewAuthenticationRoute(routeGroup fiber.Router, handler api.Handler) AuthenticationRoute {
	return AuthenticationRoute{RouteGroup: routeGroup, Handler: handler}
}

func (route AuthenticationRoute) RegisterRoute() {

	// init
	authRoute := route.RouteGroup.Group("/authentication")
	adminRoute := authRoute.Group("/admin")
	jwtMiddleware := middlewares.NewJWTMiddleware(route.Handler.Contract)

	// handler
	handler := NewAuthHandler(route.Handler)
	userHandler := userHandlers.NewUserHandler(route.Handler)
	adminHandler:=adminHandlers.NewAdminHandler(route.Handler)

	// Auth Route
	authRoute.Post("/login", handler.Login)
	authRoute.Post("/register", handler.Register)

	// Auth Admin Route
	adminRoute.Post("/login", handler.LoginAdmin)

	// recover
	authRecoverRoute := authRoute.Group("/recover")
	authRecoverRoute.Post("/otp", handler.RecoverOTP)
	authRecoverRoute.Post("/email/otp", handler.SendEmailOTPRecover)

	authRecoverRoute.Use(jwtMiddleware.New)
	authRecoverRoute.Post("/password", handler.RecoverChangePassword)

	// non verified user
	authNonVerifiedRoute := authRoute.Group("/verify").Use(jwtMiddleware.NonVerifiedOnly)
	authNonVerifiedRoute.Post("/otp", handler.VerifyOTP)
	authNonVerifiedRoute.Post("/email/otp", handler.SendEmailOTPVerify)

	// verified user
	authVerifiedRoute := authRoute.Group("/verified").Use(jwtMiddleware.VerifiedOnly)
	authVerifiedRoute.Get("/current", userHandler.GetCurrentUser)

	// admin only
	adminRoute.Use(jwtMiddleware.AdminOnly)
	adminRoute.Get("/current", adminHandler.GetCurrent)
}
