package main

import (
	"crud-app/database"
	"crud-app/tasks"
	"time"

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

	app.Post("/tasks", func(c *fiber.Ctx) error {
		var input struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Status      string    `json:"status"`
			Priority    string    `json:"priority"`
			DueDate     time.Time `json:"due_date"`
			Assignee    string    `json:"assignee"`
		}

		if err := c.BodyParser(&input); err!= nil {
            return c.Status(400).SendString(err.Error())
        }

		if err := tasks.ValidateTaskInputs(input.Title, input.Description, input.Status, input.Priority, input.DueDate, input.Assignee); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		task, err := tasks.NewTask(input.Title, input.Description, input.Status, input.Priority, input.DueDate, input.Assignee)
		if err != nil {
			return err
		}

		if err := db.Create(&task).Error; err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(task)
	} )

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
