package router

import (
	"github.com/gofiber/fiber/v2"
	"seafarer-backend/api"
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

	// handler
	handler := NewAuthHandler(route.Handler)

	// Auth Route
	authRoute.Post("/login", handler.Login)

}
