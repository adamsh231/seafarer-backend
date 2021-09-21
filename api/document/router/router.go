package router

import (
	"seafarer-backend/api"

	"github.com/gofiber/fiber/v2"

	"seafarer-backend/server/http/middlewares"
)

type DocumentRoute struct {
	RouteGroup fiber.Router
	Handler    api.Handler
}

func NewDocumentRoute(routeGroup fiber.Router, handler api.Handler) DocumentRoute {
	return DocumentRoute{RouteGroup: routeGroup, Handler: handler}
}

func (route DocumentRoute) RegisterRoute() {

	// init
	handler := NewAfeHandler(route.Handler)
	jwtMiddleware := middlewares.NewJWTMiddleware(route.Handler.Contract)

	// --------------------- AFE ---------------------- //

	afeRoute := route.RouteGroup.Group("/afe")
	afeRoute.Use(jwtMiddleware.VerifiedOnly)
	afeRoute.Get("", handler.Get)
	afeRoute.Post("", handler.Save)
	afeRoute.Get("/download", handler.Download)

	// ------------------------------------------------ //

}
