package controllers

import (
	"github.com/saucelibertarix/servidorgofiber/models"
	"github.com/saucelibertarix/servidorgofiber/utils"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Migrate() {
	// DROP
	GormDB.Migrator().DropTable(&models.Movie{})
	GormDB.Migrator().DropTable(&models.Director{})
	GormDB.Migrator().DropTable(&models.Client{})

	// CREATE
	GormDB.AutoMigrate(&models.Client{})
	GormDB.AutoMigrate(&models.Movie{})
	GormDB.AutoMigrate(&models.Director{})
}

func CreateFakeData() {
	client := &models.Client{
		Name:            "pepe",
		Email:           "pepe@gmail.com",
		Password:        utils.HashPassword("pepe"),
		Role:            "Client",
	}

	GormDB.Create(client)
}
