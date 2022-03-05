package middlewares

import (
	"strings"

	"github.com/breeders-zone/morphs-service/internal/handlers/http/errors"
	"github.com/breeders-zone/morphs-service/utils"
	"github.com/gofiber/fiber/v2"
)

func Role(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqRolesStr := c.Get("X-User-Roles")

		reqRoles := strings.Split(reqRolesStr, ",")

		accepted := true;

		for _, v := range roles {
			if !utils.SliceContainsString(reqRoles, v) {
				accepted = false
				break
			}
		}

		if accepted {
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(&errors.ErrorResponse{fiber.StatusForbidden, "Forbidden"})
	}
}