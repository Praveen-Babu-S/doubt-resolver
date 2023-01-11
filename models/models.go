package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Student struct {
	gorm.Model
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Questions []Question `json:"questions,omitempty"`
}

type Mentor struct {
	gorm.Model
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Subject   string     `json:"subject,omitempty"`
	Solutions []Solution `json:"solutions,omitempty"`
}

type Question struct {
	gorm.Model
	Subject   string   `json:"subject,omitempty"`
	Topic     string   `json:"topic,omitempty"`
	Desc      string   `json:"desc,omitempty"`
	Status    bool     `json:"status,omitempty"`
	StudentId uint     `json:"student_id,omitempty"`
	Solution  Solution `json:"solution,omitempty"`
}

type Solution struct {
	gorm.Model
	Explanation string    `json:"explanation,omitempty"`
	MentorId    uint      `json:"mentor_id,omitempty"`
	QuestionID  uint      `json:"question_id,omitempty"`
	Comments    []Comment `json:"comments,omitempty"`
}

type Comment struct {
	gorm.Model
	Msg        string `json:"msg,omitempty"`
	SolutionId uint   `json:"solution_id,omitempty"`
}
