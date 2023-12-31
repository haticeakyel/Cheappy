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
	app.Delete("/products/:id", api.HandleDeleteProduct)
	app.Put("/products/:id", api.HandleEditProduct)

	//brands
	app.Post("/addBrand", api.HandleAddBrand)
	app.Get("/brands", api.HandleGetBrands)
	app.Get("/brands/:id", api.HandleGetBrand)
	app.Delete("/brands/:id", api.HandleDeleteBrand)

	//websites
	app.Post("/addWebsite", api.HandleAddWebsite)
	app.Get("/websites", api.HandleGetWebsites)
	app.Get("/websites/:id", api.HandleGetWebsite)

	//categories
	app.Post("/addCategory", api.HandleAddCategory)
	app.Get("/categories", api.HandleGetCategories)
	app.Get("/categories/:id", api.HandleGetCategory)

	//search
	//app.Get("/search",api.SearchProduct())

	return app
}
