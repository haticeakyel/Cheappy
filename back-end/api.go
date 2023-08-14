package main

import (
	"github.com/gofiber/fiber/v2"
	model "github.com/haticeakyel/back-end/models"
)

type Api struct {
	Service *Service
}

func NewApi(service *Service) *Api {
	return &Api{
		Service: service,
	}
}

func (a *Api) HandleAddProduct(c *fiber.Ctx) error {

	productDTO := model.ProductDTO{}
	err := c.BodyParser(&productDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	eventCreate, err := a.Service.CreateProduct(productDTO)
	switch err {
	case nil:
		c.JSON(eventCreate)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{"error": err.Error()})
	}
	return nil
}
