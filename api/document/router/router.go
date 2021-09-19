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
	authRoute := route.RouteGroup.Group("/document")
	fileRoute := authRoute.Group("/afe")
	jwtMiddleware := middlewares.NewJWTMiddleware(route.Handler.Contract)

	// verified user
	fileRoute.Use(jwtMiddleware.VerifiedOnly)
	fileRoute.Post("", handler.Save)
	fileRoute.Get("", handler.Get)
	fileRoute.Get("/download", handler.Download)
}
