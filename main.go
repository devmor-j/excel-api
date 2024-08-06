package main

import (
	"log"
	"os"

	"github.com/devmor-j/excel-api/api"
	"github.com/devmor-j/excel-api/db"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Seed()

	app := fiber.New()

	appV1 := app.Group("/api/v1")

	appV1.Get("/health-check", api.HealthCheckHandler)
	appV1.Get("/export-excel", api.ExportExcelHandler)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
