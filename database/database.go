package controllers

import (
	"gorm.io/gorm"
	"github.com/saucelibertarix/servidorgofiber/models/base"
)

var GormDB *gorm.DB

func Migrate() {

	// DROP
	GormDB.Migrator().DropTable(&models.Movie{})

	// CREATE
	GormDB.AutoMigrate(&models.Movie{})
}

