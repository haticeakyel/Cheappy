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

	productCreate, err := a.Service.CreateProduct(productDTO)
	switch err {
	case nil:
		c.JSON(productCreate)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{"error": err.Error()})
	}
	return nil
}

func (a *Api) HandleGetProducts(c *fiber.Ctx) error {

	products, err := a.Service.GetProducts()

	switch err {
	case nil:
		c.JSON(products)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}

func (a *Api) HandleGetProduct(c *fiber.Ctx) error {
	ID := c.Params("id")

	product, err := a.Service.GetProduct(ID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	c.JSON(product)
	c.Status(fiber.StatusOK)

	return nil
}

func (a *Api) HandleAddBrand(c *fiber.Ctx) error {

	brandDTO := model.BrandDTO{}
	err := c.BodyParser(&brandDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	brandCreate, err := a.Service.CreateBrand(brandDTO)
	switch err {
	case nil:
		c.JSON(brandCreate)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{"error": err.Error()})
	}
	return nil
}

func (a *Api) HandleGetBrands(c *fiber.Ctx) error {

	brands, err := a.Service.GetBrands()

	switch err {
	case nil:
		c.JSON(brands)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return c.JSON(brands)
}

func (a *Api) HandleGetBrand(c *fiber.Ctx) error {
	ID := c.Params("id")

	brand, err := a.Service.GetBrand(ID)

	switch err {
	case nil:
		c.JSON(brand)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (a *Api) HandleAddWebsite(c *fiber.Ctx) error {

	websiteDTO := model.WebsiteDTO{}
	err := c.BodyParser(&websiteDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	websiteCreate, err := a.Service.CreateWebsite(websiteDTO)
	switch err {
	case nil:
		c.JSON(websiteCreate)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{"error": err.Error()})
	}
	return nil
}

func (a *Api) HandleGetWebsites(c *fiber.Ctx) error {

	websites, err := a.Service.GetWebsites()

	switch err {
	case nil:
		c.JSON(websites)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}
	return nil
}

func (a *Api) HandleGetWebsite(c *fiber.Ctx) error {
	ID := c.Params("id")

	website, err := a.Service.GetWebsite(ID)

	switch err {
	case nil:
		c.JSON(website)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (a *Api) HandleAddCategory(c *fiber.Ctx) error {

	categoryDTO := model.CategoryDTO{}
	err := c.BodyParser(&categoryDTO)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	categoryCreate, err := a.Service.CreateCategory(categoryDTO)
	switch err {
	case nil:
		c.JSON(categoryCreate)
		c.Status(fiber.StatusCreated)
	default:
		c.Status(fiber.StatusInternalServerError)
		c.JSON(fiber.Map{"error": err.Error()})
	}
	return nil
}

func (a *Api) HandleGetCategories(c *fiber.Ctx) error {

	categories, err := a.Service.GetCategories()

	switch err {
	case nil:
		c.JSON(categories)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return c.JSON(categories)
}

func (a *Api) HandleGetCategory(c *fiber.Ctx) error {
	ID := c.Params("id")

	category, err := a.Service.GetCategory(ID)

	switch err {
	case nil:
		c.JSON(category)
		c.Status(fiber.StatusOK)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}

func (a *Api) HandleDeleteProduct(c *fiber.Ctx) error {
	ID := c.Params("id")

	err := a.Service.DeleteProduct(ID)

	switch err {
	case nil:
		c.Status(fiber.StatusNoContent)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}


func (a *Api) HandleDeleteBrand(c *fiber.Ctx) error {
	ID := c.Params("id")

	err := a.Service.DeleteBrand(ID)

	switch err {
	case nil:
		c.Status(fiber.StatusNoContent)
	default:
		c.Status(fiber.StatusInternalServerError)
	}

	return nil
}