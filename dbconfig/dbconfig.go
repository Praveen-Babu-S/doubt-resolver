package dbconfig

import (
	"log"

	models "github.com/backend-ids/models"
	"github.com/jinzhu/gorm"
)

// setup call for db
func DBSetup() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=backend-pet sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

// starting db
func DBstart() *gorm.DB {
	db := DBSetup()

	//create students table
	db.DropTable(&models.Student{})
	db.CreateTable(&models.Student{})

	//create mentors table
	db.DropTable(&models.Mentor{})
	db.CreateTable(&models.Mentor{})

	//create quenstions table
	db.DropTable(&models.Question{})
	db.CreateTable(&models.Question{})

	//create solutions table
	db.DropTable(&models.Solution{})
	db.CreateTable(&models.Solution{})

	// create comments table
	db.DropTable(&models.Comment{})
	db.CreateTable(&models.Comment{})

	//set foreign key for question table (student_id is foreign key here)
	db.Debug().Model(&models.Question{}).AddForeignKey("student_id", "students(id)", "CASCADE", "CASCADE")

	//set foreign key for solution table (mentor_id is foreign key here)
	db.Debug().Model(&models.Solution{}).AddForeignKey("mentor_id", "mentors(id)", "CASCADE", "CASCADE")

	//set foreign key for comment table (solution_id is foreign key)
	db.Debug().Model(&models.Comment{}).AddForeignKey("solution_id", "solutions(id)", "CASCADE", "CASCADE")

	//set foreign key for solution table (question_id is foreign key)
	db.Debug().Model(&models.Solution{}).AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")

	s := []models.Student{
		{
			Name:  "Praveen",
			Email: "praveen@gmail.com",
			Questions: []models.Question{
				{
					Subject: "Subject 2",
					Topic:   "Topic 2",
					Desc:    "Desc 2",
				},
			},
		},
		{
			Name:  "Rohit",
			Email: "rohit@gmail.com",
		},
	}
	for _, S := range s {
		db.Create(&S)
	}
	m := models.Mentor{
		Name:    "Mentor 1",
		Email:   "Email 1",
		Subject: "Science",
	}
	db.Create(&m)
	q := models.Question{
		Subject: "Subject 1",
		Topic:   "Topic 1",
		Desc:    "Desc 1",
	}
	db.Create(&q)
	s1 := models.Solution{
		Explanation: "Explanation 1",
	}
	db.Create(&s1)
	c := models.Comment{
		Msg: "Hey,I need some better explanation here!",
	}
	db.Create(&c)

	return db
}
