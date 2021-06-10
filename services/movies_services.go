package services

import (
	"errors"
	database "github.com/saucelibertarix/servidorgofiber/database"
	"github.com/saucelibertarix/servidorgofiber/models"
)

func GetAllMovies(list *[]models.Movie) error {
	// buscamos si existe en la db
	queryResult := database.GormDB.Find(&list)

	return queryResult.Error
}

func GetMovie(id uint) (*models.Movie, error) {
	findMovie := new(models.Movie)

	// buscamos si existe en la db
	queryResult := database.GormDB.Where("id = ?", id).First(&findMovie)

	return findMovie, queryResult.Error
}

func CreateMovie(receivedMovie *models.Movie) error {
	// buscamos si existe en la db
	database.GormDB.Where("name = ?", receivedMovie.Name).Find(&receivedMovie)

	// si tiene id > 0 significa que existe
	if receivedMovie.ID > 0 {
		return errors.New("Pelicula ya existe")
	}

	// insertamos el objeto la base de datos
	createResult := database.GormDB.Create(&receivedMovie)

	// devolvemos el resultando de la operación de insert
	return createResult.Error
}

func UpdateMovie(id uint, movie *models.Movie) (*models.Movie, error) {
	queryMovie := new(models.Movie)
	database.GormDB.Where("id = ?", id).Find(&queryMovie)
	if queryMovie.ID == 0{
		return nil, errors.New("No existe la película que quierwe modificar")
	}
	updatedMovie := database.GormDB.Where("id = ?", &queryMovie.ID).Updates(&movie)
	return movie , updatedMovie.Error
}

func DeleteMovieById(id uint) error {
	requestMovie := new(models.Movie)
	database.GormDB.Where("id = ?", id).Find(&requestMovie)

	if requestMovie.ID == 0{
		return errors.New("No existe la película que quiere borrar")
	}
	deleteResult := database.GormDB.Delete(requestMovie)

	return deleteResult.Error
}
