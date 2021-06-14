package utils

import "github.com/gofiber/fiber/v2"
import "github.com/go-playground/validator/v10"

func ReturnErrorResponse(status int, err error, context *fiber.Ctx) error {
	return context.Status(status).JSON(&fiber.Map{
		"error": err.Error(),
	})
}

func ReturnSuccessResponse(context *fiber.Ctx) error {
	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func ValidateRequestPayload(payload interface{}) error {
	v := validator.New()
	err := v.Struct(payload)
	if err != nil {
		for i := 0; i < len(err.(validator.ValidationErrors)); i++ {
			if err.(validator.ValidationErrors)[i] != nil {
				return err.(validator.ValidationErrors)[i]
			}
		}
	}
	return nil
}
