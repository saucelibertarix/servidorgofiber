package controllers

import (
	"github.com/saucelibertarix/servidorgofiber/models"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Migrate() {
	// DROP
	GormDB.Migrator().DropTable(&models.Movie{})
	GormDB.Migrator().DropTable(&models.Director{})

	// CREATE
	GormDB.AutoMigrate(&models.Movie{})
	GormDB.AutoMigrate(&models.Director{})
}

