package controllers

import (
	"errors"
	"fmt"

	"github.com/Huyvtph13755/go/config"
	"github.com/Huyvtph13755/go/models"
	"github.com/gofiber/fiber/v2"
)

func GetOrder(c *fiber.Ctx) error {
	var orders []*models.Order
	config.Database.Find(&orders)

	fmt.Println(orders)

	return c.JSON(orders)
}

func GetOrderByParam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("invalid ID")
	}
	var orders *models.Order
	config.Database.First(&orders, id)

	return c.JSON(orders)
}

func DeleteOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("invalid ID")
	}

	var orders *models.Order
	config.Database.Where("id = ?", id).Delete(&orders)
	return c.Status(201).JSON(orders)
}
