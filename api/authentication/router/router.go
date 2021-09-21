package router

import (
	"github.com/gofiber/fiber/v2"
	"seafarer-backend/api"
	adminHandlers "seafarer-backend/api/admin/router"
	userHandlers "seafarer-backend/api/user/router"

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
	handler := NewAuthHandler(route.Handler)
	jwtMiddleware := middlewares.NewJWTMiddleware(route.Handler.Contract)

	// --------------------- User ---------------------- //

	// user route
	userRoute := route.RouteGroup.Group("/user")
	userHandler := userHandlers.NewUserHandler(route.Handler)

	// Auth Route
	userRoute.Post("/login", handler.Login)
	userRoute.Post("/register", handler.Register)

	// recover
	authRecoverRoute := userRoute.Group("/recover")
	authRecoverRoute.Post("/otp", handler.RecoverOTP)
	authRecoverRoute.Post("/email/otp", handler.SendEmailOTPRecover)
	authRecoverRoute.Post("/password", jwtMiddleware.New, handler.RecoverChangePassword)

	// non verified user
	authNonVerifiedRoute := userRoute.Group("/verify").Use(jwtMiddleware.NonVerifiedOnly)
	authNonVerifiedRoute.Post("/otp", handler.VerifyOTP)
	authNonVerifiedRoute.Post("/email/otp", handler.SendEmailOTPVerify)

	// verified user
	authVerifiedRoute := userRoute.Group("/verified").Use(jwtMiddleware.VerifiedOnly)
	authVerifiedRoute.Get("/current", userHandler.GetCurrentUser)

	// -------------------------------------------------- //

	// --------------------- Admin ---------------------- //

	// admin route
	adminRoute := route.RouteGroup.Group("/admin")
	adminHandler:=adminHandlers.NewAdminHandler(route.Handler)

	// login
	adminRoute.Post("/login", handler.LoginAdmin)

	// admin only
	adminRoute.Use(jwtMiddleware.AdminOnly)
	adminRoute.Get("/current", adminHandler.GetCurrent)

	// -------------------------------------------------- //

}
