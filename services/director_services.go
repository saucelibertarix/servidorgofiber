package services

import (
	"errors"
	database "github.com/saucelibertarix/servidorgofiber/database"
	"github.com/saucelibertarix/servidorgofiber/models"
)

func GetAllDirectors(list *[]models.Director) error{
	queryResult := database.GormDB.Find(&list)

	return queryResult.Error
}

func GetDirectorById(id uint) (*models.Director, error){
	findDirector := new(models.Director)

	queryResult := database.GormDB.Where("id = ?", id).First(&findDirector)

	return findDirector, queryResult.Error
}

func CreateDirector(receviedDirector *models.Director) error{
	database.GormDB.Where("name = ?", receviedDirector.Name).Find(&receviedDirector)

	if receviedDirector.ID != 0{
		return errors.New("El Director ya existe")
	}

	createResult := database.GormDB.Create(&receviedDirector)

	return createResult.Error
}

func UpdateDirector(id uint, director *models.Director) (*models.Director, error){
	queryDirector := new(models.Director)
	database.GormDB.Where("name = ?", director.Name).Find(&queryDirector)
	if queryDirector.ID == 0{
		return nil, errors.New("No existe el director que quiere actualizar ")
	}

	updateResult := database.GormDB.Updates(&director)

	return director, updateResult.Error
}

func DeleteDirector(id uint) error{
	queryDirector := new(models.Director)
	database.GormDB.Where("id = ?", id).Find(&queryDirector)
	if queryDirector.ID == 0{
		return errors.New("No existe el director que quiere borrar")
	}
	deleteResult := *database.GormDB.Delete(&queryDirector)
	return deleteResult.Error
}
