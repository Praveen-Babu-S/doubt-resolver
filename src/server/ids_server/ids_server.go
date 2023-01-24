package server

import (
	"context"

	pb "github.com/backend-ids/proto"
	DB "github.com/backend-ids/src/schema/db"
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

type IdsDbServer struct {
	pb.UnimplementedIdsCRUDServer
	Db DB.DbOperations
}

func NewIdsDbserver(db *gorm.DB) *IdsDbServer {
	return &IdsDbServer{Db: DB.DbClient{Db: db}}
}

// get questions by student_id
func (s *IdsDbServer) GetQuestions(in *pb.Id, stream pb.IdsCRUD_GetQuestionsServer) error {
	questions := s.Db.GetQuestions(in.Id)
	for _, question := range questions {
		if err := stream.Send(&question); err != nil {
			return err
		}
	}
	return nil
}

// get question by question_id
func (s *IdsDbServer) GetQuestionById(ctx context.Context, in *pb.Id) (*pb.QuestionById, error) {
	q := s.Db.GetQuestionById(in.Id)
	return q, nil
}

// takes q_id and returns student_id and mentor_id
func (s *IdsDbServer) FindIDs(ctx context.Context, in *pb.Id) (*pb.Ids, error) {
	res := s.Db.FindIDs(in.Id)
	return res, nil
}

// takes solution_id returns question_id
func (s *IdsDbServer) FindQID(ctx context.Context, in *pb.Id) (*pb.Id, error) {
	res := s.Db.FindQID(in.Id)
	return res, nil
}
