package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func extractRequestId(c *fiber.Ctx) uuid.UUID {
	requestIdAsStr := c.Locals("requestid").(string)

	requestID, err := uuid.Parse(requestIdAsStr)

	if err != nil {
		return uuid.New()
	}

	return requestID
}

func NewCtxWithCorrelationId(c *fiber.Ctx) context.Context {
	return context.WithValue(context.Background(), "correlationID", extractRequestId(c))
}
