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
