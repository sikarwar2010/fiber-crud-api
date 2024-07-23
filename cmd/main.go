package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sikarwar2010/doc-fiber-app/database"
	"github.com/sikarwar2010/doc-fiber-app/routes"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World india!")
	})

	log.Fatal(app.Listen(":4000"))
}

func setupRoutes(app *fiber.App) {
	// user endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUserById)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// product endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProductById)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	// order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("api/orders/:id", routes.GetOrderById)

	// 

}
