package api

import "github.com/gofiber/fiber/v2"

func HealthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "ok"})
}

func ExportExcelHandler(c *fiber.Ctx) error {
	return nil
}
