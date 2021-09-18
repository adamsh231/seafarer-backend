package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"seafarer-backend/api"
	authenticationRouter "seafarer-backend/api/authentication/router"
	"time"
)

type Router struct {
	api.Handler
}

func NewRouter(handler api.Handler) Router {
	return Router{handler}
}

var (
	logFormat = `{"host":"${host}","pid":"${pid}","time":"${time}","request_id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}",` +
		`"user-agent":"${ua}","bytes_in":"${bytesReceived}","bytes_out":"${bytesSent}"}`
)

func (router Router) RegisterRoutes() {

	app := router.Contract.App

	// middleware
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	app.Use(logger.New(logger.Config{
		Format:     logFormat + "\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Jakarta",
	}))

	// Route for check health
	app.Get("", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("I am Fine Thanks!")
	})

	// Api Group - v1
	apiV1Group := app.Group("/v1")

	// Auth route
	authenticationRouter.NewAuthenticationRoute(apiV1Group, router.Handler).RegisterRoute()

}
