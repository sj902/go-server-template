package middleware

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type CustomMiddlewareConfig struct {
}

func NewCustomMiddleware(config CustomMiddlewareConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Printf("this is middleware")
		return c.Next()
	}
}
