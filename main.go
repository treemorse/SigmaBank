package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jintonick/SigmaBank/database"
	"github.com/jintonick/SigmaBank/routes"
)

func main() {

	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))

	routes.Setup(app)

	err := app.Listen(":8080")
	if err != nil {
		panic("could not start server")
	}
}
