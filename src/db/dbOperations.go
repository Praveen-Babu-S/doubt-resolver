package dbOperations

import (
	"github.com/backend-ids/src/models"
	"github.com/jinzhu/gorm"
)

type dbOperations interface {
	CreateUser(*models.User)
	EditUser(*models.User, uint64)
	CreateSolution(*models.Solution) string
	EditSolution(*models.Solution, uint64) string
	CreateQuestion(*models.Question) string
	EditQuestion(*models.Question, uint64) string
	CreateComment(*models.Comment) string
	GetQuestionById()
}

type DbClient struct {
	db *gorm.DB
}

func (Db DbClient) CreateUser(u *models.User) {
	Db.db.Create(&u)
}

func (Db DbClient) EditUser(u *models.User, id uint64) {
	Db.db.Model(&models.User{}).Where("id = ?", id).Updates(u)
}

func (Db DbClient) CreateSolution(sol *models.Solution) {
	Db.db.Create(&sol)
}

func (Db DbClient) EditSolution(sol *models.Solution, sid uint64) {
	Db.db.Model(&models.Solution{}).Where("id = ?", sid).Updates(sol)
}

func (Db DbClient) CreateQuestion(q *models.Question) {
	u := models.User{}
	Db.db.Raw("SELECT id FROM users WHERE role=? and subject=? ORDER BY RANDOM() LIMIT 1", "mentor", q.Subject).Scan(&u)
	q.AssigneeId = uint64(u.ID)
	Db.db.Create(&q)
}

func (Db DbClient) EditQuestion(q *models.Question, id uint64) {
	Db.db.Model(&models.Question{}).Where("id = ?", id).Updates(q)
}

func (Db DbClient) CreateComment(c *models.Comment) {
	Db.db.Create(&c)
}

func (Db DbClient) GetQuestionById() {

}
