package controllers

import (
	"github.com/saucelibertarix/servidorgofiber/models"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Migrate() {
	// DROP
	GormDB.Migrator().DropTable(&models.Movie{})

	// CREATE
	GormDB.AutoMigrate(&models.Movie{})
}

