package middleware

import (
	"crud-app/users"
	"crud-app/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
				Message: "Authorization header is missing",
				Error:   "Authorization header is missing",
				Status:  fiber.StatusUnauthorized,
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
				Message: "Invalid authorization header format",
				Error:   "Invalid authorization header format",
				Status:  fiber.StatusUnauthorized,
			})
		}

		tokenString := parts[1]

		claims, err := utils.VerifyToken(tokenString)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.Response{
				Message: "unauthorized",
				Error:   err.Error(),
				Status:  fiber.StatusUnauthorized,
			})
		}

		user, err := users.FindUserByEmail(claims["email"].(string))

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(utils.Response{
				Message: "user does not exist",
				Error:   err.Error(),
				Status:  fiber.StatusNotFound,
			})
		}

		c.Locals("user", user)

		return c.Next()
	}
}
