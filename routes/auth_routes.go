package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saucelibertarix/servidorgofiber/controllers"
)

func AuthRoutes(router fiber.Router){
	// /api/v1/auth
	authRouter := router.Group("/auth")

	authRouter.Post("/login", controllers.LoginClient)
	authRouter.Post("/signup", controllers.SignupClient)
}
