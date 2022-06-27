package main

import (
	"github.com/Huyvtph13755/go/config"
	"github.com/Huyvtph13755/go/controllers"
	"github.com/Huyvtph13755/go/middlewares"
	"github.com/Huyvtph13755/go/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.ConnectDB()
	config.Database.Migrator().AutoMigrate(models.User{}, models.Product{}, models.Order{}, models.OrderDetail{})
	config.InitSessionStore()
	order_api := app.Group("/orders")
	product_api := app.Group("/products")
	orderDetail_api := app.Group("/order_details")
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", middlewares.Auth, controllers.Logout)
	app.Get("/me", middlewares.Auth, controllers.GetCurrentUser)
	app.Post("/cart", middlewares.Auth, controllers.CreateOrderDetail)
	order_api.Delete("/:id", controllers.DeleteOrder)
	product_api.Get("/", controllers.GetProduct)
	product_api.Get("/:id", controllers.GetProductByParam)
	product_api.Post("/", controllers.CreateProduct)
	product_api.Put("/:id", controllers.UpdateProduct)
	product_api.Delete("/:id", controllers.DeleteProduct)
	orderDetail_api.Get("/", controllers.GetOrderDetail)
	orderDetail_api.Get("/:id", controllers.GetOrderDetailByParam)
	orderDetail_api.Put("/:id", controllers.UpdateOrderDetail)
	orderDetail_api.Delete("/:id", controllers.DeleteOrderDetail)
	app.Listen(":3000")
}
