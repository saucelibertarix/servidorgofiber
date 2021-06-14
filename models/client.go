package models

import "github.com/saucelibertarix/servidorgofiber/models/base"

type Client struct {
	base.CustomGormModel
	Name     string `gorm:"type:varchar(64) not null" json:"name" xml:"name" form:"name" validate:"required"`
	Email    string `json:"email" xml:"email" form:"email"`
	Password string `json:"password" xml:"password" form:"password"`
	Role     string `json:"rol" xml:"rol" form:"rol"`
}
