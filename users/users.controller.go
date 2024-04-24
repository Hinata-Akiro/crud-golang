package users

import (
	"crud-app/utils"

	"github.com/gofiber/fiber/v2"
)

func UserController(app *fiber.App) {
	userController := app.Group("/api/v1/users")
	userController.Post("/", newUser)
	userController.Post("/login", userSignIn)
}

// create a new user
func newUser(c *fiber.Ctx) error {
	var userInput User
	var createdUser *User // Rename the variable here
	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := userInput.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Message: "Validation failed",
			Error:   err.Error(),
			Status:  fiber.StatusBadRequest,
		})
	}

	createdUser, err := createNewUser(&userInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Message: "Error creating user",
			Error:   err.Error(),
			Status:  fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(utils.Response{
		Message: "User created successfully",
		Data:    &fiber.Map{"task": createdUser},
		Status:  fiber.StatusCreated,
	})
}

func userSignIn(c *fiber.Ctx) error {
	var loginInput LoginInput
	if err := c.BodyParser(&loginInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := loginInput.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Message: "Validation failed",
			Error:   err.Error(),
			Status:  fiber.StatusBadRequest,
		})
	}

	token, err := loginUser(loginInput.Email, loginInput.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
			Message: "Error logging in",
			Error:   err.Error(),
			Status:  fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.Response{
		Message: "User logged in successfully",
		Data:    &fiber.Map{"token": token},
		Status:  fiber.StatusOK,
	})

}
