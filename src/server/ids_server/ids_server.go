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
