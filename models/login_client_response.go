package models

type LoginClientResponse struct {
	Name  string `json:"name" xml:"name" form:"name"`
	Email string `json:"email" xml:"email" form:"email"`
	Token     string `json:"token" xml:"token" form:"token"`
}
