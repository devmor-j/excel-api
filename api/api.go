package api

import (
	"context"
	"reflect"
	"strconv"

	"github.com/devmor-j/excel-api/db"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HealthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "ok"})
}

func ExportExcelHandler(c *fiber.Ctx) error {
	database := db.GetMongoDatabase()
	studnetsCollection := database.Collection(db.CollectionStudnets)

	filters := bson.D{}
	opts := options.Find()

	limit := c.QueryInt("limit")
	if limit > 0 {
		opts.SetLimit(int64(limit))
	}

	cur, err := studnetsCollection.Find(context.TODO(), filters, opts)
	if err != nil {
		return err
	}

	var students []db.Student
	if err := cur.All(context.TODO(), &students); err != nil || len(students) == 0 {
		return err
	}

	f := excelize.NewFile()
	defer f.Close()
	sheet := f.GetSheetName(0)
	f.SetSheetName("Students", sheet)

	// TODO: find a better way to omit object_id
	columnLength := reflect.TypeOf(students[0]).NumField() - 1
	headTitles := make([]string, 0, columnLength)

	for i := range columnLength {
		headTitles = append(headTitles, students[0].TagValue(i+1))
	}
	f.SetSheetRow(sheet, "A1", &headTitles)

	for rowIndex := range len(students) {
		var row = make([]any, 0, columnLength)
		for i := range columnLength {
			row = append(row, students[rowIndex].FieldValue(i+1))
			f.SetSheetRow(sheet, "A"+strconv.Itoa(rowIndex+2), &row)
		}
	}

	// if err := f.SaveAs("export.xlsx"); err != nil {
	// 	fmt.Println(err)
	// }

	c.Response().Header.Set("Content-Type", "application/octet-stream")
	c.Response().Header.Set("Content-Disposition", "attachment; filename=export.xlsx")
	c.Response().Header.Set("Content-Transfer-Encoding", "binary")

	return f.Write(c.Response().BodyWriter())
}
