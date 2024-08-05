package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func healthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "ok"})
}

func exportExcelHandler(c *fiber.Ctx) error {
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	appV1 := app.Group("/api/v1")

	appV1.Get("/health-check", healthCheckHandler)
	appV1.Get("/export-excel", exportExcelHandler)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
