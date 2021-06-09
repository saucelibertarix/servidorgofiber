package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saucelibertarix/servidorgofiber/controllers"
	"github.com/saucelibertarix/servidorgofiber/middlewares"
)

func DirectorRoutes(router fiber.Router){
	// /api/v1/director
	directorRouter := router.Group("/director")

	// /api/v1/director/ READ
	directorRouter.Get("/", controllers.GetAllDirectors)

	// /api/v1/director/:id READ
	directorRouter.Get("/:id", controllers.GetDirectorByID)

	// /api/v1/director/ CREATE
	directorRouter.Post("/", controllers.CreateDirector)

	// /api/v1/director/:id UPDATE
	directorRouter.Patch("/:id", controllers.UpdateDirector)

	directorRouter.Use(middlewares.AdminMiddleware)

	// /api/v1/director/:id
	directorRouter.Delete(":id", controllers.DeleteDirector)


}
