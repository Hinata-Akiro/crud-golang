package main

import (
	"crud-app/database"
	"crud-app/tasks"
	"crud-app/users"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func main() {
	app := fiber.New()
    err := database.Connect()

	if err != nil {
		panic(err)
	}

	if err := database.Database.AutoMigrate(&tasks.Task{}, &users.User{}); err != nil {
		panic("failed to auto migrate database: " + err.Error())
	}

	app.Use(helmet.New())
	app.Use(cors.New())


	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to crud application")
	})


	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:      3,
	}))

	tasks.TaskController(app)
	users.UserController(app)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
