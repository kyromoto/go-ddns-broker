package httpcontroller

import "github.com/gofiber/fiber/v2"

func GetStatus() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("ok")
	}
}
