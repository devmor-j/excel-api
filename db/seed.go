package db

import (
	"context"
	"log"
	"time"

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

func SeedStudentsColl() error {
	database := GetMongoDatabase()
	studnetsColl := database.Collection(CollStudnets)

	var students = make([]any, 0, len(seedStudents))

	for _, v := range seedStudents {
		var s Student

		s.Age = v.Age
		s.Country = v.Country
		s.Email = v.Email
		s.EntryDate = ParseShortDate(v.EntryDate)
		s.FirstName = v.FirstName
		s.GPA = v.GPA
		s.ID = primitive.NewObjectID()
		s.LastName = v.LastName
		s.Major = v.Major
		s.StudentId = v.StudentId

		students = append(students, s)
	}

	if _, err := studnetsColl.InsertMany(context.TODO(), students); err != nil {
		return err
	}

	return nil
}
