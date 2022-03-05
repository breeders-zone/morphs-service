package middlewares

import (
	"github.com/breeders-zone/morphs-service/internal/handlers/http/errors"
	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error  {
	user_id := ctx.Get("X-User-ID")
	if user_id == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&errors.ErrorResponse{fiber.StatusUnauthorized, "Unauthorized"})
	}

	return ctx.Next()
}