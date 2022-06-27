package controllers

import (
	"errors"
	"fmt"

	"github.com/Huyvtph13755/go/config"
	"github.com/Huyvtph13755/go/models"
	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {
	var products []*models.Product
	config.Database.Preload("OrderDetail").Find(&products)

	fmt.Println(products)

	return c.JSON(products)
}

func GetProductByParam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("invalid ID")
	}
	var products *models.Product
	config.Database.Preload("OrderDetail").First(&products, id)

	return c.JSON(products)
}

func CreateProduct(c *fiber.Ctx) error {
	type Params struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Discount int64   `json:"discount"`
	}

	params := new(Params)
	if err := c.BodyParser(params); err != nil {
		return err
	}

	products := &models.Product{
		Name:     params.Name,
		Price:    params.Price,
		Discount: params.Discount,
	}
	config.Database.Create(products)

	return c.Status(201).JSON(products)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("invalid ID")
	}

	var products *models.Product
	config.Database.First(&products, id)

	type Params struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Discount int64   `json:"discount"`
	}

	params := new(Params)
	if err := c.BodyParser(params); err != nil {
		return err
	}

	config.Database.Model(&products).Updates(models.Product{
		Name:     params.Name,
		Price:    params.Price,
		Discount: params.Discount,
	})

	return c.Status(201).JSON(products)
}
func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return errors.New("invalid ID")
	}

	var products *models.Product
	config.Database.Where("id = ?", id).Delete(&products)
	return c.Status(201).JSON(products)
}
