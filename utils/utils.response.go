package utils

import (
	"github.com/gofiber/fiber/v2"
)


type Response struct {
	Message string `json:"message"`
	Data    *fiber.Map `json:"data"`
	Error   *error        `json:"error"`
	Status  int `json:"status"`
}