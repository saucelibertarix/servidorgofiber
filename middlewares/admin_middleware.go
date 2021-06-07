package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(context *fiber.Ctx) error {
	requestApiKeyOrginal := "1234"
	requestApiKey := context.Get("key")

	if requestApiKeyOrginal != requestApiKey {
		return context.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"error": "No tienes permisos",
		})
	}

	return context.Next()
}