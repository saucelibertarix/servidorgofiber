package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saucelibertarix/servidorgofiber/models"
	"github.com/saucelibertarix/servidorgofiber/services"
	"github.com/saucelibertarix/servidorgofiber/utils"
	"strconv"
)

func GetAllDirectors(context *fiber.Ctx) error{
	var list = make([]models.Director, 0)

	err := services.GetAllDirectors(&list)
	if err != nil{
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}
	return context.JSON(&list)
}

func GetDirectorByID(context *fiber.Ctx) error{
	directorID, _ := strconv.ParseUint(context.Params("id"), 10, 64)

	movie, err := services.GetDirectorById(uint(directorID))
	if err !=nil{
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}
	return context.JSON(movie)
}

func CreateDirector(context *fiber.Ctx) error{
	//Creamos objeto
	director := new(models.Director)
	//Parseamos
	err := context.BodyParser(director)
	if err != nil{
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}
	//Validamos
	err = utils.ValidateRequestPayload(director, context)
	if err != nil{
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}
	//Llamamos al Servicio
	err = services.CreateDirector(director)
	if err != nil{
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}
	//devolvemos creado
	return context.JSON(director)
}

func UpdateDirector(context *fiber.Ctx) error{
	directorId, _ := strconv.ParseUint(context.Params("id"), 10, 64)
	director := new(models.Director)
	err := context.BodyParser(director)
	if err != nil{
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	err = utils.ValidateRequestPayload(director, context)
	if err != nil{
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	director, err = services.UpdateDirector(uint(directorId), director)
	if err != nil{
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}
	return context.JSON(director)
}
func DeleteDirector(context *fiber.Ctx) error{
	directorId, _ := strconv.ParseUint(context.Params("id"), 10, 64)

	err := services.DeleteDirector(uint(directorId))
	if err != nil{
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}
	return context.JSON("Director borrado adecuadamente")
}
