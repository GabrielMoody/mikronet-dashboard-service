package main

import (
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/handler"
	"github.com/GabrielMoody/mikronet-dashboard-service/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		ServerHeader:  "mikronet.systems",
	})

	app.Use(func(c *fiber.Ctx) error {
		// c.Set("Content-Security-Policy", "default-src 'self'; img-src 'self'; script-src 'self'")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Frame-Options", "deny")
		c.Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")
		return c.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	app.Use(logger.New(logger.Config{
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Singapore",
	}))

	db := models.DatabaseInit()

	api := app.Group("/")

	handler.DashboardHandler(api, db)

	err := app.Listen("0.0.0.0:8030")
	if err != nil {
		return
	}
}
