package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/swaggo/fiber-swagger"

	_ "github.com/Weburz/burzpress/server/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	app := fiber.New(fiber.Config{
		AppName:       "BurzPress API",
		StrictRouting: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Failed to load the server: %w", err)
	}
}
