package v1

import (
	"server/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (s Server) HelloHandler(c *fiber.Ctx) error {

	return utils.WriteResponse(c, []byte((time.Now().String())), fiber.StatusOK)
}
