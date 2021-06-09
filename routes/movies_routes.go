package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/saucelibertarix/servidorgofiber/controllers"
	"github.com/saucelibertarix/servidorgofiber/middlewares"
)

func MovieRoutes(router fiber.Router) {

	// /api/v1/movie
	movieRouter := router.Group("/movie")


	// /api/v1/movie/ | READ
	movieRouter.Get("/", controllers.GetAllMovies)

	// /api/v1/movie/:id | READ
	movieRouter.Get("/:id", controllers.GetMovie)

	// /api/v1/movie | CREATE
	movieRouter.Post("/", controllers.CreateMovie)

	// /api/v1/movie/:id | UPDATE
	movieRouter.Patch("/:id", controllers.UpdateMovie)

	// /api/v1/movie/:id | UPDATE
	movieRouter.Delete("/:id", controllers.DeleteMovieByID)

	// aplicamos middleware de admin
	movieRouter.Use(middlewares.AdminMiddleware)

	// /api/v1/movie/:id | READ
	movieRouter.Get("/", controllers.GetAllMovies)
}
