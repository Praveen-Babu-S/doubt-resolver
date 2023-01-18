package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type User struct {
	gorm.Model
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Role      string     `json:"role,omitempty" gorm:"not null"`
	Subject   string     `json:"subject,omitempty"`
	Questions []Question `json:"questions,omitempty"`
	Solutions []Solution `json:"solutions,omitempty"`
	Comments  []Comment  `json:"comments,omitempty"`
}
type Question struct {
	gorm.Model
	Subject    string   `json:"subject,omitempty"`
	Desc       string   `json:"desc,omitempty"`
	StudentId  uint64   `json:"student_id,omitempty"`
	AssigneeId uint64   `json:"assignee_id,omitempty"`
	Solution   Solution `json:"solution,omitempty"`
}

type Solution struct {
	gorm.Model
	Desc       string    `json:"explanation,omitempty"`
	MentorId   uint64    `json:"mentor_id,omitempty"`
	QuestionID uint64    `json:"question_id,omitempty"`
	Comments   []Comment `json:"comments,omitempty"`
}

type Comment struct {
	gorm.Model
	Msg        string `json:"msg,omitempty"`
	SolutionId uint64 `json:"solution_id,omitempty"`
	UserId     uint64 `json:"user_id"`
}

type QuestionById struct {
	Question Question
	Solution Solution
	Comments []Comment
}
