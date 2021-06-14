package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/saucelibertarix/servidorgofiber/controllers"
	"github.com/saucelibertarix/servidorgofiber/middlewares"
	"github.com/saucelibertarix/servidorgofiber/utils"
)

func DirectorRoutes(router fiber.Router){



	// /api/v1/director
	directorRouter := router.Group("/director")

	directorRouter.Use(jwtware.New(jwtware.Config{SigningKey: []byte(utils.GetEnvVariable("JWT_CLIENT_KEY"))}))

	// /api/v1/director/ READ
	directorRouter.Get("/", controllers.GetAllDirectors)

	// /api/v1/director/:id READ
	directorRouter.Get("/:id", controllers.GetDirectorByID)

	// /api/v1/director/ CREATE
	directorRouter.Post("/", controllers.CreateDirector)

	directorRouter.Use(middlewares.AdminMiddleware)

	// /api/v1/director/:id UPDATE
	directorRouter.Patch("/:id", controllers.UpdateDirector)

	// /api/v1/director/:id
	directorRouter.Delete(":id", controllers.DeleteDirector)
}
