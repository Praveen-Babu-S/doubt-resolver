package main

import (
	"context"
	"log"
	"net"

	"github.com/backend-ids/dbconfig"
	pb "github.com/backend-ids/ids_proto"
	"github.com/backend-ids/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const port = ":3031"

type idsServer struct {
	pb.UnimplementedIdsCRUDServer
	db *gorm.DB
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listes: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterIdsCRUDServer(s, &idsServer{db: dbconfig.DBSetup()})
	log.Printf("server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faile due to server: %v", err)
	}
}

func (s *idsServer) CreateUser(ctx context.Context, in *pb.User) (*pb.Status, error) {
	u := models.User{Name: in.Name, Email: in.Email, Password: in.Password, Role: in.Role, Subject: in.Subject}
	s.db.Create(&u)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

func (s *idsServer) EditUser(ctx context.Context, in *pb.User) (*pb.Status, error) {
	u := models.User{Name: in.Name, Email: in.Email, Password: in.Password, Role: in.Role, Subject: in.Subject}
	s.db.Model(&models.User{}).Where("id = ?", in.Id).Updates(u)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

func (s *idsServer) CreateQuestion(ctx context.Context, in *pb.Question) (*pb.Status, error) {
	q := models.Question{Desc: in.Desc, Subject: in.Subject, StudentId: in.StudentId}
	u := models.User{}
	s.db.Raw("SELECT id FROM users WHERE role=? and subject=? ORDER BY RANDOM() LIMIT 1", "mentor", q.Subject).Scan(&u)
	q.AssigneeId = uint64(u.ID)
	s.db.Create(&q)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

func (s *idsServer) EditQuestion(ctx context.Context, in *pb.Question) (*pb.Status, error) {
	res := pb.Status{}
	q := models.Question{Desc: in.Desc, Subject: in.Subject, StudentId: in.StudentId, AssigneeId: in.AssigneeId}
	s.db.Model(&models.Question{}).Where("id = ?", in.Id).Updates(q)
	res.Id = "1"
	return &res, nil
}

func (s *idsServer) CreateComment(ctx context.Context, in *pb.Comment) (*pb.Status, error) {
	c := models.Comment{Msg: in.Msg, SolutionId: in.SolutionId, UserId: in.UserId}
	s.db.Create(&c)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

func (s *idsServer) CreateSolution(ctx context.Context, in *pb.Solution) (*pb.Status, error) {
	q := models.Solution{Desc: in.Desc, MentorId: in.MentorId, QuestionID: in.QuestionId}
	s.db.Create(&q)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

func (s *idsServer) EditSolution(ctx context.Context, in *pb.Solution) (*pb.Status, error) {
	res := pb.Status{}
	q := models.Solution{Desc: in.Desc, MentorId: in.MentorId, QuestionID: in.QuestionId}
	s.db.Model(&models.Solution{}).Where("id = ?", in.Id).Updates(q)
	res.Id = "1"
	return &res, nil
}

// authorised to student whose student_id matches with id
func (s *idsServer) GetQuestions(in *pb.Id, stream pb.IdsCRUD_GetQuestionsServer) error {
	questions := []pb.Question{}
	s.db.Model(&models.Question{}).Where("student_id = ?", in.Id).Find(&questions)
	for _, question := range questions {
		if err := stream.Send(&question); err != nil {
			return err
		}
	}
	return nil
}

func (s *idsServer) GetQuestionById(ctx context.Context, in *pb.Id) (*pb.QuestionById, error) {
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

func (s *idsServer) FindIDs(ctx context.Context, in *pb.Id) (*pb.Ids, error) {
	res := &pb.Ids{}
	q := models.Question{}
	s.db.Model(&models.Question{}).Where("id=?", in.Id).Find(&q)
	res.Sid = q.StudentId
	res.Aid = q.AssigneeId
	return res, nil
}

func (s *idsServer) FindQID(ctx context.Context, in *pb.Id) (*pb.Id, error) {
	res := &pb.Id{}
	sol := models.Solution{}
	s.db.Model(&models.Solution{}).Where("id=?", in.Id).Find(&sol)
	res.Id = sol.QuestionID
	return res, nil
}
