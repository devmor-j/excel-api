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

func ParseShortDate(timeValue string) time.Time {
	t, err := time.Parse("02/01/2006", timeValue)
	if err != nil {
		log.Fatal(err)
	}

	return t
}

func SeedStudentsCollection() error {
	// https://www.kaggle.com/datasets/adeliayuliarni/dummy-student-dataset
	seedStudents := []SeedStudent{
		{1, "Maurits", "Katte", 20, "mkatte0@behance.net", "Senegal", "Computer Science", 2.29, "04/05/2019"},
		{2, "Adoree", "Letchford", 30, "aletchford1@geocities.jp", "Brazil", "Physics", 1.95, "05/12/2019"},
		{3, "Darelle", "Heckner", 20, "dheckner2@github.io", "Canada", "Marketing", 3.24, "02/09/2015"},
		{4, "Maureen", "Ockwell", 30, "mockwell3@msn.com", "Russia", "Journalism", 0.91, "05/03/2004"},
		{5, "Abbie", "Hrus", 29, "ahrus4@narod.ru", "Venezuela", "Business Administration", 2.84, "11/02/2016"},
		{6, "Iorgo", "Carmichael", 30, "icarmichael5@myspace.com", "France", "Economics", 3.07, "09/04/2014"},
		{7, "Janetta", "Vassall", 23, "jvassall6@linkedin.com", "Argentina", "Physics", 1.38, "27/08/2012"},
		{8, "Shani", "Curedell", 26, "scuredell7@eventbrite.com", "China", "Mathematics", 2.73, "20/04/2001"},
		{9, "Bobbie", "Joao", 30, "bjoao8@tripod.com", "Indonesia", "Biology", 2.15, "12/01/2014"},
		{10, "Amy", "Pollok", 22, "apollok9@barnesandnoble.com", "Mexico", "Physics", 2.24, "10/06/2014"},
		{11, "Charil", "Whitcombe", 28, "cwhitcombea@examiner.com", "Sweden", "Civil Engineering", 3.64, "29/12/2004"},
		{12, "Denver", "Bernardoux", 20, "dbernardouxb@dmoz.org", "Botswana", "Biology", 1.97, "07/05/2009"},
		{13, "Pepito", "Alp", 28, "palpc@ca.gov", "China", "Marketing", 3.87, "11/07/2001"},
		{14, "Caye", "Fend", 18, "cfendd@china.com.cn", "Czech Republic", "Computer Science", 3.69, "29/03/2004"},
		{15, "Dacy", "Heskins", 21, "dheskinse@wunderground.com", "China", "Civil Engineering", 1.73, "10/06/2018"},
		{16, "Kennett", "Dack", 19, "kdackf@pen.io", "Russia", "Computer Science", 3.56, "20/07/2019"},
		{17, "Calla", "Newlan", 27, "cnewlang@auda.org.au", "Ecuador", "Computer Science", 2.12, "10/06/2005"},
		{18, "Mikaela", "Harget", 21, "mhargeth@php.net", "Bangladesh", "Business Administration", 3.26, "06/05/2006"},
		{19, "Gray", "Gadson", 24, "ggadsoni@multiply.com", "Malaysia", "Anthropology", 2.95, "23/11/2018"},
		{20, "Deidre", "Gillebride", 25, "dgillebridej@google.it", "Portugal", "Philosophy", 0.49, "06/10/2005"},
		{21, "Jamie", "Hyde", 28, "jhydek@ebay.com", "Egypt", "Music", 1.58, "10/09/2012"},
		{22, "Marcelline", "Pye", 24, "mpyel@cdbaby.com", "China", "Physics", 2.39, "15/12/2014"},
		{23, "Dolf", "Gelland", 30, "dgellandm@twitter.com", "Panama", "Civil Engineering", 1.23, "16/08/2011"},
		{24, "Cecilio", "Magor", 25, "cmagorn@foxnews.com", "Brazil", "English Literature", 0.6, "01/08/2005"},
		{25, "Ninetta", "Bryning", 29, "nbryningo@icq.com", "Canada", "Journalism", 2.21, "05/06/2011"},
		{26, "Mitzi", "Briand", 29, "mbriandp@washington.edu", "Argentina", "Philosophy", 1.33, "24/10/2020"},
		{27, "Auria", "Allibone", 25, "aalliboneq@soundcloud.com", "Marshall Islands", "Business Administration", 1.84, "19/01/2012"},
		{28, "Lanita", "Wackly", 28, "lwacklyr@springer.com", "Indonesia", "Physics", 3.76, "27/05/2002"},
		{30, "Rollie", "Chimes", 24, "rchimest@virginia.edu", "Ukraine", "Biology", 3.85, "02/09/2016"},
		{31, "Alana", "Joselevitz", 28, "ajoselevitzu@tripod.com", "Indonesia", "Political Science", 2.79, "13/12/2005"},
		{32, "Bil", "Patkin", 24, "bpatkinv@wp.com", "China", "Biology", 3.62, "08/07/2008"},
		{33, "Kati", "Kerley", 30, "kkerleyw@vkontakte.ru", "China", "Civil Engineering", 2.05, "12/10/2009"},
		{34, "Vida", "Nelius", 25, "vneliusx@ftc.gov", "Mexico", "Economics", 2.05, "03/01/2008"},
		{35, "Lowe", "McDermott-Row", 27, "lmcdermottrowy@businessweek.com", "Pakistan", "Environmental Science", 1.73, "17/10/2021"},
		{36, "Chrisse", "Jaskiewicz", 23, "cjaskiewiczz@yandex.ru", "Thailand", "Chemistry", 0.54, "21/09/2009"},
		{37, "Jeramie", "Aggus", 23, "jaggus10@youtube.com", "China", "Computer Science", 1.07, "11/02/2014"},
		{38, "Doralyn", "Cossington", 30, "dcossington11@typepad.com", "Thailand", "Chemistry", 3.69, "18/10/2022"},
		{39, "Rayna", "Elstub", 18, "relstub12@csmonitor.com", "Greece", "Economics", 3.51, "26/10/2011"},
		{40, "Stavro", "Capel", 18, "scapel13@vk.com", "Venezuela", "Marketing", 2.27, "23/05/2004"},
		{41, "Ebeneser", "Climar", 27, "eclimar14@soup.io", "United States", "English Literature", 2.73, "04/09/2009"},
	}

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
