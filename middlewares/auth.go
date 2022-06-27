package middlewares

import (
	"github.com/Huyvtph13755/go/config"
	"github.com/Huyvtph13755/go/models"
	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) error {
	session, err := config.SessionStore.Get(c)
	if err != nil {
		return err
	}
	email := session.Get("email")
	if email == nil {
		return c.Status(401).JSON("Bạn chưa đăng nhập")
	}
	var user *models.User
	config.Database.First(&user, "email = ?", email)
	c.Locals("user", user)
	return c.Next()
}
