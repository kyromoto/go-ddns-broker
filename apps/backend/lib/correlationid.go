package lib

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

const CorrelationIdKey = "correaltionID"

func GetCorrelationId(ctx *context.Context) string {
	c := *ctx

	cid := c.Value(CorrelationIdKey)

	if cid == nil {
		return "<not set>"
	}

	return cid.(string)
}

func NewContextWithCorrelationId(cid string) *context.Context {
	ctx := context.WithValue(context.TODO(), CorrelationIdKey, cid)
	return &ctx
}

func NewContextWithCorrelationIdFromFiber(c *fiber.Ctx) *context.Context {
	cid := (*c).Locals(CorrelationIdKey).(string)
	ctx := context.WithValue(context.TODO(), CorrelationIdKey, cid)
	return &ctx
}
