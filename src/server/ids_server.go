package server

import (
	"context"

	"github.com/backend-ids/src/models"
	pb "github.com/backend-ids/src/proto"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type IdsServer struct {
	pb.UnimplementedIdsCRUDServer
	db *gorm.DB
}

func NewIdsServer(db *gorm.DB) *IdsServer {
	return &IdsServer{db: db}
}

// create new user
func (s *IdsServer) CreateUser(ctx context.Context, in *pb.User) (*pb.Status, error) {
	u := models.User{Name: in.Name, Email: in.Email, Password: in.Password, Role: in.Role, Subject: in.Subject}
	s.db.Create(&u)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

// edit an user
func (s *IdsServer) EditUser(ctx context.Context, in *pb.User) (*pb.Status, error) {
	u := models.User{Name: in.Name, Email: in.Email, Password: in.Password, Role: in.Role, Subject: in.Subject}
	s.db.Model(&models.User{}).Where("id = ?", in.Id).Updates(u)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

// create new question
func (s *IdsServer) CreateQuestion(ctx context.Context, in *pb.Question) (*pb.Status, error) {
	q := models.Question{Desc: in.Desc, Subject: in.Subject, StudentId: in.StudentId}
	u := models.User{}
	s.db.Raw("SELECT id FROM users WHERE role=? and subject=? ORDER BY RANDOM() LIMIT 1", "mentor", q.Subject).Scan(&u)
	q.AssigneeId = uint64(u.ID)
	s.db.Create(&q)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

// edit a question
func (s *IdsServer) EditQuestion(ctx context.Context, in *pb.Question) (*pb.Status, error) {
	res := pb.Status{}
	q := models.Question{Desc: in.Desc, Subject: in.Subject, StudentId: in.StudentId, AssigneeId: in.AssigneeId}
	s.db.Model(&models.Question{}).Where("id = ?", in.Id).Updates(q)
	res.Id = "1"
	return &res, nil
}

// create new comment
func (s *IdsServer) CreateComment(ctx context.Context, in *pb.Comment) (*pb.Status, error) {
	c := models.Comment{Msg: in.Msg, SolutionId: in.SolutionId, UserId: in.UserId}
	s.db.Create(&c)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

// create new solution
func (s *IdsServer) CreateSolution(ctx context.Context, in *pb.Solution) (*pb.Status, error) {
	q := models.Solution{Desc: in.Desc, MentorId: in.MentorId, QuestionID: in.QuestionId}
	s.db.Create(&q)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

// edit solution
func (s *IdsServer) EditSolution(ctx context.Context, in *pb.Solution) (*pb.Status, error) {
	res := pb.Status{}
	q := models.Solution{Desc: in.Desc, MentorId: in.MentorId, QuestionID: in.QuestionId}
	s.db.Model(&models.Solution{}).Where("id = ?", in.Id).Updates(q)
	res.Id = "1"
	return &res, nil
}

// get questions by student_id
func (s *IdsServer) GetQuestions(in *pb.Id, stream pb.IdsCRUD_GetQuestionsServer) error {
	questions := []pb.Question{}
	s.db.Model(&models.Question{}).Where("student_id = ?", in.Id).Find(&questions)
	for _, question := range questions {
		if err := stream.Send(&question); err != nil {
			return err
		}
	}
	return nil
}

// get question by question_id
func (s *IdsServer) GetQuestionById(ctx context.Context, in *pb.Id) (*pb.QuestionById, error) {
	Q := models.Question{}
	S := models.Solution{}
	C := []models.Comment{}
	s.db.Model(&models.Question{}).Where("id=?", in.Id).Find(&Q)
	s.db.Model(&models.Solution{}).Where("question_id=?", in.Id).Find(&S)
	s.db.Model(&models.Comment{}).Where("solution_id=?", S.ID).Find(&C)
	// fmt.Println(que, sol, com)
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
	return q, nil
}

// takes q_id and returns student_id and mentor_id
func (s *IdsServer) FindIDs(ctx context.Context, in *pb.Id) (*pb.Ids, error) {
	res := &pb.Ids{}
	q := models.Question{}
	s.db.Model(&models.Question{}).Where("id=?", in.Id).Find(&q)
	res.Sid = q.StudentId
	res.Aid = q.AssigneeId
	return res, nil
}

// takes solution_id returns question_id
func (s *IdsServer) FindQID(ctx context.Context, in *pb.Id) (*pb.Id, error) {
	res := &pb.Id{}
	sol := models.Solution{}
	s.db.Model(&models.Solution{}).Where("id=?", in.Id).Find(&sol)
	res.Id = sol.QuestionID
	return res, nil
}
