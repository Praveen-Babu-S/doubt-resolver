package dbconfig

import (
	"log"

	"github.com/jinzhu/gorm"
)

// setup call for db
func DBSetup() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=backend-pet sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	// defer db.Close()
	return db
}

// starting db
func DBstart() *gorm.DB {
	db := DBSetup()

	//create users table
	// db.DropTable(&models.User{})
	// db.CreateTable(&models.User{})

	//create quenstions table
	// db.DropTable(&models.Question{})
	// db.CreateTable(&models.Question{})

	//create solutions table
	// db.DropTable(&models.Solution{})
	// db.CreateTable(&models.Solution{})

	// create comments table
	// db.DropTable(&models.Comment{})
	// db.CreateTable(&models.Comment{})

	// set foreign key for question table (student_id is foreign key here)
	// db.Debug().Model(&models.Question{}).AddForeignKey("student_id", "users(id)", "CASCADE", "CASCADE")
	// db.Debug().Model(&models.Question{}).AddForeignKey("assignee_id", "users(id)", "CASCADE", "CASCADE")

	// //set foreign key for comments table(user_id is foreign key here)
	// db.Debug().Model(&models.Comment{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	// //set foreign key for solution table (mentor_id is foreign key here)
	// db.Debug().Model(&models.Solution{}).AddForeignKey("mentor_id", "users(id)", "CASCADE", "CASCADE")

	// //set foreign key for comment table (solution_id is foreign key)
	// db.Debug().Model(&models.Comment{}).AddForeignKey("solution_id", "solutions(id)", "CASCADE", "CASCADE")

	// //set foreign key for solution table (question_id is foreign key)
	// db.Debug().Model(&models.Solution{}).AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")

	// s := []models.User{
	// 	{
	// 		Name:  "User 1",
	// 		Email: "mail 1@gmail.com",
	// 		Role:  "Student",
	// 		Questions: []models.Question{
	// 			{
	// 				Subject: "Subject 1",
	// 				Desc:    "Desc 1",
	// 			},
	// 		},
	// 	},
	// 	{
	// 		Name:  "user 2",
	// 		Email: "mail@gmail.com",
	// 		Role:  "Mentor",
	// 	},
	// }
	// for _, S := range s {
	// 	db.Create(&S)
	// }

	// q := models.Question{
	// 	Subject:    "Subject 2",
	// 	Desc:       "Desc 2",
	// 	StudentId:  1,
	// 	AssigneeId: 2,
	// }
	// db.Create(&q)
	// s1 := models.Solution{
	// 	Desc:       "Explanation 1",
	// 	QuestionID: 2,
	// 	MentorId:   2,
	// }
	// db.Create(&s1)
	// c := models.Comment{
	// 	Msg:        "Hey,I need some better explanation here!",
	// 	SolutionId: 1,
	// 	UserId:     1,
	// }
	// db.Create(&c)

	return db
}
