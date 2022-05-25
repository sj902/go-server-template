package router

import (
	v1 "server/handlers/v1"
	"server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App, customMiddleware fiber.Handler, server v1.Server) {

	app.Get("/health", func(c *fiber.Ctx) error {
		return utils.WriteResponse(c, []byte{}, fiber.StatusOK)
	})

	// Middleware
	login := app.Group("/v1", logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${body} ${reqHeader}\n",
	}))
	login.Post("/hello", customMiddleware, server.HelloHandler)
}
