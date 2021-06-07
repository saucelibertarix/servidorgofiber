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

func UpdateMovie(movie *models.Movie) (*models.Movie, error) {

	/*
		updateResul, err := dbgorm.update(movie)
		if err != nill
			return err
	*/

	return nil, nil
}

func DeleteMovie(id uint) error {
	/*
		comprobar que existe
		result = findById(id)
		if result === existe
			dbgorm.delete(result)
		else {
		retirm erro("no existe")

	*/

	return nil
}
