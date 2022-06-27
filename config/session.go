package config

import "github.com/gofiber/fiber/v2/middleware/session"

var SessionStore *session.Store

func InitSessionStore() {
	SessionStore = session.New(session.Config{})
}
