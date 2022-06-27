package controllers

import (
	"errors"
	"fmt"

	"github.com/Huyvtph13755/go/config"
	"github.com/Huyvtph13755/go/models"
	"github.com/gofiber/fiber/v2"
)

func GetOrderDetail(c *fiber.Ctx) error {
	var order_details []*models.OrderDetail
	config.Database.Find(&order_details)

	fmt.Println(order_details)

	return c.JSON(order_details)
}

func GetOrderDetailByParam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("invalid ID")
	}
	var order_details *models.OrderDetail
	config.Database.First(&order_details, id)

	return c.JSON(order_details)
}

func CreateOrderDetail(c *fiber.Ctx) error {
	type CartProduct struct {
		ProductID int64 `json:"product_id"`
		Quantity  int64 `json:"quantity"`
	}
	type Params struct {
		Products []CartProduct `json:"products"`
	}

	params := new(Params)
	if err := c.BodyParser(params); err != nil {
		return err
	}
	var user = c.Locals("user").(*models.User)

	order := &models.Order{
		UserID: user.ID,
	}
	config.Database.Create(order)
	for _, product := range params.Products {
		order_details := &models.OrderDetail{
			OrderID:   order.ID,
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		}
		config.Database.Create(order_details)
	}

	order.CalcTotal()
	return c.Status(201).JSON(order)
}

func UpdateOrderDetail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("invalid ID")
	}

	var order_details *models.OrderDetail
	config.Database.First(&order_details, id)

	type Params struct {
		ProductID int64 `json:"product_id"`
		Quantity  int64 `json:"quantity"`
	}

	params := new(Params)
	if err := c.BodyParser(params); err != nil {
		return err
	}

	config.Database.Model(&order_details).Updates(models.OrderDetail{
		ProductID: params.ProductID,
		Quantity:  params.Quantity,
	})

	return c.Status(201).JSON(order_details)
}
func DeleteOrderDetail(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("invalid ID")
	}

	var order_details *models.OrderDetail
	config.Database.Where("id = ?", id).Delete(&order_details)
	return c.Status(201).JSON(order_details)
}
