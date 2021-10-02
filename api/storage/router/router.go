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
	jwtMiddleware := middlewares.NewJWTMiddleware(route.Handler.Contract)

	// --------------------- File ---------------------- //

	// file route
	fileRoute := route.RouteGroup.Group("/file")
	fileHandler := NewFileHandler(route.Handler)

	// admin only
	fileRoute.Get("/filter/:id", jwtMiddleware.AdminOnly, fileHandler.BrowseByUser)

	// verified user
	fileRoute.Use(jwtMiddleware.VerifiedOnly)
	fileRoute.Get("", fileHandler.BrowsePresignedKey)
	fileRoute.Post("", fileHandler.Add)
	fileRoute.Get("/:id", fileHandler.GetPresignedKey)
	fileRoute.Delete("/:id", fileHandler.Delete)

	// ------------------------------------------------- //

}
