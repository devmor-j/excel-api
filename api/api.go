package api

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/devmor-j/excel-api/db"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var pascalToNormalRE = regexp.MustCompile("([a-z0–9])([A-Z])")

func HealthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "ok"})
}

func ExportExcelHandler(c *fiber.Ctx) error {
	database := db.GetMongoDatabase()
	studnetsCollection := database.Collection(db.CollectionStudnets)

	cur, err := studnetsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	var students []db.Student
	if err := cur.All(context.TODO(), &students); err != nil {
		return err
	}

	f := excelize.NewFile()
	defer f.Close()
	sheet := f.GetSheetName(0)
	f.SetSheetName("Students", sheet)

	// to omit object_id we need to decrement max column length
	columnLength := reflect.TypeOf(students[0]).NumField() - 1

	headTitles := make([]string, 0, columnLength)

	for i := range columnLength {
		// TODO: use struct tag to allow custom naming for xlsx
		title := pascalToNormalRE.ReplaceAllString(students[0].FieldName(i+1), "$1 $2")
		headTitles = append(headTitles, title)
	}
	f.SetSheetRow(sheet, "A1", &headTitles)

	for rowIndex := range len(students) {
		var row = make([]any, 0, columnLength)
		for i := range columnLength {
			row = append(row, students[rowIndex].FieldValue(i+1))
			f.SetSheetRow(sheet, "A"+strconv.Itoa(rowIndex+2), &row)
		}
	}

	if err := f.SaveAs("students.xlsx"); err != nil {
		fmt.Println(err)
	}

	return nil
}
