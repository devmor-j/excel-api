package api

import (
	"context"
	"fmt"

	"github.com/devmor-j/excel-api/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func HealthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "ok"})
}

func ExportExcelHandler(c *fiber.Ctx) error {
	database := db.GetMongoDatabase()
	studnetsColl := database.Collection(db.CollStudnets)

	cur, err := studnetsColl.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	var students []db.Student
	if err := cur.All(context.TODO(), &students); err != nil {
		return err
	}

	fmt.Println(len(students), students)

	return nil
}
