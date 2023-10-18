package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kyromoto/go-ddns/internal/services/client"
)

func HandleCmdClientDelete(clientDeleteService client.DeleteService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := NewCtxWithCorrelationId(c)

		var props client.DeleteServiceProps

		if err := c.BodyParser(&props); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponseBody{
				error: "bad request body",
			})
		}

		result, err := clientDeleteService(ctx, props)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResponseBody{
				error: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(result)
	}
}
