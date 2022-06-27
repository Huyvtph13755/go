package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetCurrentUser(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(user)
}
