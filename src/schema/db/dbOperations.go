package dbOperations

import (
	notify "github.com/backend-ids/notification"
	pb "github.com/backend-ids/proto"
	"github.com/backend-ids/src/schema/models"
	"github.com/jinzhu/gorm"
)

type DbOperations interface {
	CreateUser(*models.User)
	EditUser(*models.User, uint64)
	CreateSolution(*models.Solution)
	EditSolution(*models.Solution, uint64)
	CreateQuestion(*models.Question)
	EditQuestion(*models.Question, uint64)
	CreateComment(*models.Comment)
	GetQuestionById(uint64) *pb.QuestionById
	GetQuestions(uint64) []pb.Question
	FindIDs(uint64) *pb.Ids
	FindQID(uint64) *pb.Id
}

type DbClient struct {
	Db *gorm.DB
}

func (Db DbClient) CreateUser(u *models.User) {
	Db.Db.Create(&u)
	msg := notify.Mailinfo(u.Email, "Signed up Successfully!", "You have successfully signed-up for ids.")
	notify.Notify(msg)
}

func (Db DbClient) EditUser(u *models.User, id uint64) {
	Db.Db.Model(&models.User{}).Where("id = ?", id).Updates(u)
	msg := notify.Mailinfo(u.Email, "Profile updated Successfully!", "You have successfully updated you profile details.")
	go notify.Notify(msg)
}

func (Db DbClient) CreateSolution(sol *models.Solution) {
	Db.Db.Create(&sol)
	q := models.Question{}
	Db.Db.Model(&models.Question{}).Where("id = ?", sol.QuestionID).Find(&q)
	u := models.User{}
	Db.Db.Model(&models.User{}).Where("id=?", q.StudentId).Find(&u)
	msg := notify.Mailinfo(u.Email, "Posted Solution!", "Your mentor has posted solution for your question!")
	go notify.Notify(msg)
}

func (Db DbClient) EditSolution(sol *models.Solution, sid uint64) {
	Db.Db.Model(&models.Solution{}).Where("id = ?", sid).Updates(sol)
	Db.Db.Find(&sol, sid)
	q := models.Question{}
	Db.Db.Model(&models.Question{}).Where("id = ?", sol.QuestionID).Find(&q)
	u := models.User{}
	Db.Db.Model(&models.User{}).Where("id=?", q.StudentId).Find(&u)
	msg := notify.Mailinfo(u.Email, "Edited Solution!", "Your mentor has updated solution for your question!")
	// fmt.Println("User : ", u)
	go notify.Notify(msg)
}

func (Db DbClient) CreateQuestion(q *models.Question) {
	u := models.User{}
	Db.Db.Raw("SELECT id FROM users WHERE role=? and subject=? ORDER BY RANDOM() LIMIT 1", "mentor", q.Subject).Scan(&u)
	q.AssigneeId = uint64(u.ID)
	Db.Db.Model(&models.User{}).Where("id = ?", q.AssigneeId).Find(&u)
	msg := notify.Mailinfo(u.Email, "New Question!", "You have been allocated new question!")
	go notify.Notify(msg)
	Db.Db.Create(&q)
}

func (Db DbClient) EditQuestion(q *models.Question, id uint64) {
	Db.Db.Model(&models.Question{}).Where("id = ?", id).Updates(q)
	u := models.User{}
	Db.Db.Model(&models.User{}).Where("id = ?", q.AssigneeId).Find(&u)
	msg := notify.Mailinfo(u.Email, "Question updated!", "The question allocated to you has been modified, please check it!")
	go notify.Notify(msg)
}

func (Db DbClient) CreateComment(c *models.Comment) {
	Db.Db.Create(&c)
}

func (Db DbClient) GetQuestionById(qId uint64) *pb.QuestionById {
	Q := models.Question{}
	S := models.Solution{}
	C := []models.Comment{}
	Db.Db.Model(&models.Question{}).Where("id=?", qId).Find(&Q)
	Db.Db.Model(&models.Solution{}).Where("question_id=?", qId).Find(&S)
	Db.Db.Model(&models.Comment{}).Where("solution_id=?", S.ID).Find(&C)
	c := []*pb.Comment{}
	for _, comment := range C {
		c = append(c, &pb.Comment{
			SolutionId: comment.SolutionId,
			UserId:     comment.UserId,
			Msg:        comment.Msg,
		})
	}
	q := &pb.QuestionById{Q: &pb.Question{
		Subject:    Q.Subject,
		Desc:       Q.Desc,
		StudentId:  Q.StudentId,
		AssigneeId: Q.AssigneeId,
		Id:         uint64(Q.ID),
	}, S: &pb.Solution{
		Desc:       S.Desc,
		QuestionId: S.QuestionID,
		MentorId:   S.MentorId,
		Id:         uint64(S.ID),
	}, C: c,
	}
	return q
}

func (Db DbClient) GetQuestions(u_id uint64) []pb.Question {
	questions := []pb.Question{}
	Db.Db.Model(&models.Question{}).Where("student_id = ?", u_id).Find(&questions)
	return questions
}

func (Db DbClient) FindIDs(qId uint64) *pb.Ids {
	res := &pb.Ids{}
	q := models.Question{}
	Db.Db.Model(&models.Question{}).Where("id=?", qId).Find(&q)
	res.Sid = q.StudentId
	res.Aid = q.AssigneeId
	return res
}

func (Db DbClient) FindQID(sId uint64) *pb.Id {
	res := &pb.Id{}
	sol := models.Solution{}
	Db.Db.Model(&models.Solution{}).Where("id=?", sId).Find(&sol)
	res.Id = sol.QuestionID
	return res
}
