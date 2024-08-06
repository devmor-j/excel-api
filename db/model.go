package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	StudentId int                `json:"studentId" bson:"student_id"`
	FirstName string             `json:"firstName" bson:"first_name"`
	LastName  string             `json:"lastName" bson:"last_name"`
	Age       int                `json:"age" bson:"age"`
	Email     string             `json:"email" bson:"email"`
	Country   string             `json:"country" bson:"country"`
	Major     string             `json:"major" bson:"major"`
	GPA       float64            `json:"gpa" bson:"gpa"`
	EntryDate time.Time          `json:"entryDate" bson:"entry_date"`
}
