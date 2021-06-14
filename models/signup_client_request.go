package models

type SignupClientRequest struct {
	Name  string `json:"name" xml:"name" form:"name" validate:"required"`
	Email    string `json:"email" xml:"email" form:"email" validate:"required"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
}