package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/robertd2000/fiber-api/db"
	"github.com/robertd2000/fiber-api/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Running!!!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Post("/api/users", routes.CreateUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
}

func main() {
	db.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
