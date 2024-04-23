package main

import (
	"crud-app/database"
	"crud-app/tasks"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db, err := database.Connect()

	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&tasks.Task{}); err != nil {
		panic("failed to auto migrate database: " + err.Error())
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to crud application")
	})

	tasks.TaskController(app)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
