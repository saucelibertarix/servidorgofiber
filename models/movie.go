package models

import (
	"github.com/saucelibertarix/servidorgofiber/models/base"
)

type Movie struct {
	base.CustomGormModel

	Name     string `gorm:"index; unique; type:varchar(64) not null" json:"name" xml:"name" form:"name"`
	Director string `gorm:"type:varchar(128) not null;" json:"director" xml:"director" form:"director"`
	Year     int    `gorm:"not null" json:"year" xml:"year" form:"year"`
}
