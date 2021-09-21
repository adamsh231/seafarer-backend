package router

import (
	"seafarer-backend/api"

	"github.com/gofiber/fiber/v2"

	"seafarer-backend/server/http/middlewares"
)

type FileRoute struct {
	RouteGroup fiber.Router
	Handler    api.Handler
}

func NewFileRoute(routeGroup fiber.Router, handler api.Handler) FileRoute {
	return FileRoute{RouteGroup: routeGroup, Handler: handler}
}

func (route FileRoute) RegisterRoute() {

	// init
	fileHandler := NewFileHandler(route.Handler)
	fileRoute := route.RouteGroup.Group("/storage")
	jwtMiddleware := middlewares.NewJWTMiddleware(route.Handler.Contract)

	// verified user
	fileRoute.Use(jwtMiddleware.VerifiedOnly)
	fileRoute.Post("", fileHandler.Add)
	fileRoute.Get("", fileHandler.BrowsePresignedKey)
	fileRoute.Get("/:id", fileHandler.GetPresignedKey)
	fileRoute.Delete("/:id", fileHandler.Delete)
}
