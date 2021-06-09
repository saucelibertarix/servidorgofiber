package models

import "github.com/saucelibertarix/servidorgofiber/models/base"

type Director struct{
	base.CustomGormModel

	Name string `gorm:"index; unique; type:varchar(64) not null" json:"name" xml:"name" form:"name" validate:"required"`
	Since string `gorm:"type:varchar(10)" json:"since" xml:"since" form:"since"`
	Nationality string `gorm:"type:varchar(40)" json:"nationality" xml:"nationality" form:"nationality"`
}