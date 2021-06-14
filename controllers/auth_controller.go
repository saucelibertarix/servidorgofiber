package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saucelibertarix/servidorgofiber/models"
	"github.com/saucelibertarix/servidorgofiber/services"
	"github.com/saucelibertarix/servidorgofiber/utils"
)

func SignupClient(context *fiber.Ctx) error {
	// Creamos un objeto vacío para meter dentro la información de la petición
	signupClientRequest := new(models.SignupClientRequest)

	// Decodificamos los datos de la petición y los metemos dentro del objeto
	// creado anteriormente
	err := context.BodyParser(signupClientRequest)

	// si el error existe
	if err != nil {
		// devolvemos un mensaje de error usando context
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	// Validamos la estructura y el valor de los datos de los datos
	err = utils.ValidateRequestPayload(signupClientRequest)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	// llamamos al servicio
	err = services.SignUpClient(signupClientRequest)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	return utils.ReturnSuccessResponse(context)
}
func LoginClient(context *fiber.Ctx) error {
	clientLoginRequest := new(models.LoginClientRequest)

	err := context.BodyParser(clientLoginRequest)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	// Validamos la estructura y el valor de los datos de los datos
	err = utils.ValidateRequestPayload(clientLoginRequest)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	clientLoginResponse, err := services.LoginClient(clientLoginRequest)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusNotFound, err, context)
	}

	return context.JSON(clientLoginResponse)
}
