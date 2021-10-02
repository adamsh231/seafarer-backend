package router

import (
	"seafarer-backend/api"

	"github.com/gofiber/fiber/v2"

	"seafarer-backend/server/http/middlewares"
)

type UserRoute struct {
	RouteGroup fiber.Router
	Handler    api.Handler
}

func NewUserRoute(routeGroup fiber.Router, handler api.Handler) UserRoute {
	return UserRoute{RouteGroup: routeGroup, Handler: handler}
}

func (route UserRoute) RegisterRoute() {

	// init
	handler := NewUserHandler(route.Handler)
	jwtMiddleware := middlewares.NewJWTMiddleware(route.Handler.Contract)

	// --------------------- User ---------------------- //

	// user route
	userRoute := route.RouteGroup.Group("/")

	// Auth Route
	userRoute.Use(jwtMiddleware.AdminOnly)
	userRoute.Get("/filter", handler.Filter)
	userRoute.Get("/candidate/filter", handler.FilterCandidate)
	userRoute.Get("/employee/filter", handler.FilterEmployee)
	userRoute.Get("/letter/filter", handler.FilterLetter)

}
