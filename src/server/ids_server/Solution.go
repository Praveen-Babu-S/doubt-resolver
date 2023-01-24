package server

import (
	"context"

	pb "github.com/backend-ids/proto"
	"github.com/backend-ids/src/schema/models"
	_ "github.com/lib/pq"
)

// create new solution
func (s *IdsDbServer) CreateSolution(ctx context.Context, in *pb.Solution) (*pb.Status, error) {
	sol := models.Solution{Desc: in.Desc, MentorId: in.MentorId, QuestionID: in.QuestionId}
	s.Db.CreateSolution(&sol)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

// edit solution
func (s *IdsDbServer) EditSolution(ctx context.Context, in *pb.Solution) (*pb.Status, error) {
	res := pb.Status{}
	sol := models.Solution{Desc: in.Desc, MentorId: in.MentorId, QuestionID: in.QuestionId}
	s.Db.EditSolution(&sol, in.Id)
	res.Id = "1"
	return &res, nil
}
