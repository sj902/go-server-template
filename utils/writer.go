package utils

import (
	"github.com/gofiber/fiber/v2"
)

func WriteResponse(ctx *fiber.Ctx, resp []byte, code int) error {
	ctx.Set("Content-Type", "application/json; charset=utf-8")
	ctx.Status(code)
	return ctx.Send(resp)
}
