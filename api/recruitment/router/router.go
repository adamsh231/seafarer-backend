package router

import (
	"seafarer-backend/api"

	"github.com/gofiber/fiber/v2"

	"seafarer-backend/server/http/middlewares"
)

type RecruitmentsRoute struct {
	RouteGroup fiber.Router
	Handler    api.Handler
}

func NewRecruitmentsRoute(routeGroup fiber.Router, handler api.Handler) RecruitmentsRoute {
	return RecruitmentsRoute{RouteGroup: routeGroup, Handler: handler}
}

func (route RecruitmentsRoute) RegisterRoute() {

	// init
	handler := NewRecruitmentsHandler(route.Handler)
	jwtMiddleware := middlewares.NewJWTMiddleware(route.Handler.Contract)

	// --------------------- Recruitments ---------------------- //

	// user route
	userRoute := route.RouteGroup.Group("/")

	// Auth Route
	userRoute.Use(jwtMiddleware.AdminOnly)
	userRoute.Post("/candidate", handler.CreateCandidate)

}
