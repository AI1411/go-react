package routes

import (
	"github.com/gofiber/fiber"
	"goreact/controllers"
	"goreact/middlewares"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Post("/api/users", controllers.CreateUser)
}
