package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	repository, err := NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	defer repository.client.Disconnect(context.Background())
	service := NewService(repository)
	api := NewApi(&service)
	app := SetupApp(api)
	app.Listen(":3001")
}

func SetupApp(api *Api) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	//products
	app.Post("/addProduct", api.HandleAddProduct)
	app.Get("/products", api.HandleGetProducts)
	app.Get("/products/:id", api.HandleGetProduct)

	//brands
	app.Post("/addBrand", api.HandleAddBrand)
	app.Get("/brands", api.HandleGetBrands)
	app.Get("/brands/:id", api.HandleGetBrand)

	//websites
	app.Post("/addWebsite", api.HandleAddWebsite)
	app.Get("/websites", api.HandleGetWebsites)
	app.Get("/websites/:id", api.HandleGetWebsite)

	return app
}
