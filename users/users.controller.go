package users

import (
    "github.com/gofiber/fiber/v2"
	"crud-app/utils"
)



func UserController(app *fiber.App) {
	userController := app.Group("/api/v1/users")
	userController.Post("/", newUser)
}

//create a new user 
func newUser(c *fiber.Ctx) error {
    var userInput User
    var createdUser *User // Rename the variable here
    if err := c.BodyParser(&userInput); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    if err := userInput.Validate(); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
            Message: "Validation failed",
            Error:   &err,
            Status:  fiber.StatusBadRequest,
        })
    }

    createdUser, err := createNewUser(&userInput)

    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
            Message: "Error creating user",
            Error:   &err,
            Status:  fiber.StatusBadRequest,
        })
    }

    return c.Status(fiber.StatusCreated).JSON(utils.Response{
        Message: "User created successfully",
        Data:    &fiber.Map{"task": createdUser},
        Status:  fiber.StatusCreated,
    })
}