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
	StudentId uint64   `json:"student_id,omitempty"`
	Solution  Solution `json:"solution,omitempty"`
}

type Solution struct {
	gorm.Model
	Explanation string    `json:"explanation,omitempty"`
	MentorId    uint64    `json:"mentor_id,omitempty"`
	QuestionID  uint64    `json:"question_id,omitempty"`
	Comments    []Comment `json:"comments,omitempty"`
}

type Comment struct {
	gorm.Model
	Msg        string `json:"msg,omitempty"`
	SolutionId uint64 `json:"solution_id,omitempty"`
}
