package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SeedStudent struct {
	StudentId int
	FirstName string
	LastName  string
	Age       int
	Email     string
	Country   string
	Major     string
	GPA       float64
	EntryDate string
}

// https://www.kaggle.com/datasets/adeliayuliarni/dummy-student-dataset
var seedStudents = []SeedStudent{
	{1, "Maurits", "Katte", 20, "mkatte0@behance.net", "Senegal", "Computer Science", 2.29, "04/05/2019"},
}

func ParseShortDate(timeValue string) time.Time {
	t, err := time.Parse("02/01/2006", timeValue)
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func SeedStudentsCollection() error {
	database := GetMongoDatabase()
	studnetsCollection := database.Collection(CollectionStudnets)

	cur, err := studnetsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}
	var dbStudents []Student
	cur.All(context.TODO(), &dbStudents)

	var students = make([]any, 0, len(seedStudents))

uniqueSeed:
	for _, ss := range seedStudents {
		for _, ds := range dbStudents {
			if ds.StudentId == ss.StudentId {
				continue uniqueSeed
			}
		}

		var s Student

		s.Age = ss.Age
		s.Country = ss.Country
		s.Email = ss.Email
		s.EntryDate = ParseShortDate(ss.EntryDate)
		s.FirstName = ss.FirstName
		s.GPA = ss.GPA
		s.ID = primitive.NewObjectID()
		s.LastName = ss.LastName
		s.Major = ss.Major
		s.StudentId = ss.StudentId

		students = append(students, s)
	}

	if len(students) != 0 {
		res, err := studnetsCollection.InsertMany(context.TODO(), students)
		if err != nil {
			return err
		}
		fmt.Println("SEED: Inserted", len(res.InsertedIDs), "Unique Students")
	}

	return nil
}
