package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/saucelibertarix/servidorgofiber/models"
	"github.com/saucelibertarix/servidorgofiber/services"
	"github.com/saucelibertarix/servidorgofiber/utils"
	"strconv"
)

func GetAllMovies(context *fiber.Ctx) error {
	// parse request
	var list = make([]models.Movie, 0)

	// validation
	err := services.GetAllMovies(&list)
	if err != nil {
		// devolvemos un mensaje de error usando context
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	return context.JSON(list)
}

func GetMovie(context *fiber.Ctx) error {
	// parse request
	movieID, _ := strconv.ParseUint(context.Params("id"), 10, 64)
	fmt.Println(movieID)

	// consultamos en la db
	movie, err := services.GetMovie(uint(movieID))
	if err != nil {
		// devolvemos un mensaje de error usando context
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	// create db

	return context.JSON(movie)
}

func CreateMovie(context *fiber.Ctx) error {
	// Creamos un objeto vacío para meter dentro la información de la petición
	movie := new(models.Movie)

	// Decodificamos los datos de la petición y los metemos dentro del objeto
	// creado anteriormente
	err := context.BodyParser(movie)

	// si el error existe
	if err != nil {
		// devolvemos un mensaje de error usando context
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	// Validamos la estructura y el valor de los datos de los datos
	err = utils.ValidateRequestPayload(movie)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	// llamamos al servicio
	err = services.CreateMovie(movie)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	return context.JSON(movie)
}

func UpdateMovie(context *fiber.Ctx) error {
	movieId, _ := strconv.ParseUint(context.Params("id"), 10, 64)
	// parse request
	movie := new(models.Movie)
	err := context.BodyParser(movie)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	// validation
	err = utils.ValidateRequestPayload(movie)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	// create db
	createMovieResult, err := services.UpdateMovie(uint(movieId), movie)
	if err != nil {
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	return context.JSON(createMovieResult)
}

func DeleteMovieByID(context *fiber.Ctx) error {
	// parse request
	movieId, _ := strconv.ParseUint(context.Params("id"), 10, 64)

	err := services.DeleteMovieById(uint(movieId))
	if err != nil{
		return utils.ReturnErrorResponse(fiber.StatusBadRequest, err, context)
	}

	return context.JSON("Movie delete successfully")
}
