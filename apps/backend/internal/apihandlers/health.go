package apihandlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Health() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("healthy")
	}
}
