package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jintonick/SigmaBank/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/new_point", controllers.CreatePoint)
	app.Post("/api/new_tasks", controllers.CreateTasks)
	app.Post("/api/assign_tasks", controllers.DistributeTasks)
	app.Get("/api/tasks", controllers.GetUserTasks)
}
