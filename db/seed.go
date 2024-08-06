package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// https://www.kaggle.com/datasets/adeliayuliarni/dummy-student-dataset
var students = []any{
	Student{primitive.NewObjectID(), 1, "Maurits", "Katte", 20, "mkatte0@behance.net", "Senegal", "Computer Science", 2.29, "04/05/2019"},
}

func SeedStudentsColl() error {
	database := GetMongoDatabase()
	studnetsColl := database.Collection(CollStudnets)

	_, err := studnetsColl.InsertMany(context.TODO(), students)
	if err != nil {
		return err
	}

	return nil
}
