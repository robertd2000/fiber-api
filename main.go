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
	app.Post("/api/users", routes.CreateUser)
}

func main() {
	db.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
