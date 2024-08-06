package db

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID        primitive.ObjectID `json:"id" bson:"_id" xlsx:"-"`
	StudentId int                `json:"studentId" bson:"student_id" xlsx:"Student ID"`
	FirstName string             `json:"firstName" bson:"first_name" xlsx:"First Name"`
	LastName  string             `json:"lastName" bson:"last_name" xlsx:"Last Name"`
	Age       int                `json:"age" bson:"age" xlsx:"Age"`
	Email     string             `json:"email" bson:"email" xlsx:"Email"`
	Country   string             `json:"country" bson:"country" xlsx:"Country"`
	Major     string             `json:"major" bson:"major" xlsx:"Major"`
	GPA       float64            `json:"gpa" bson:"gpa" xlsx:"GPA"`
	EntryDate time.Time          `json:"entryDate" bson:"entry_date" xlsx:"Entry Date"`
}

func (s *Student) TagValue(index int) string {
	return reflect.TypeOf(*s).Field(index).Tag.Get("xlsx")
}

func (s *Student) FieldValue(index int) any {
	return reflect.ValueOf(*s).Field(index)
}
