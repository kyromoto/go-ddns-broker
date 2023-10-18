package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kyromoto/go-ddns/internal/services/client"
)

func HandleCmdCreateClient(clientCreateService client.CreateService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := NewCtxWithCorrelationId(c)
		var props client.CreateServiceProps

		if err := c.BodyParser(&props); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponseBody{
				error: "bad request body",
			})
		}

		client, err := clientCreateService(ctx, props)

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(ErrorResponseBody{
				error: err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(client)
	}
}
