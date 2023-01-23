package server

import (
	"context"

	"github.com/backend-ids/src/models"
	pb "github.com/backend-ids/src/proto"
	_ "github.com/lib/pq"
)

// create new solution
func (s *IdsServer) CreateSolution(ctx context.Context, in *pb.Solution) (*pb.Status, error) {
	sol := models.Solution{Desc: in.Desc, MentorId: in.MentorId, QuestionID: in.QuestionId}
	s.db.Create(&sol)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

// edit solution
func (s *IdsServer) EditSolution(ctx context.Context, in *pb.Solution) (*pb.Status, error) {
	res := pb.Status{}
	sol := models.Solution{Desc: in.Desc, MentorId: in.MentorId, QuestionID: in.QuestionId}
	s.db.Model(&models.Solution{}).Where("id = ?", in.Id).Updates(sol)
	res.Id = "1"
	return &res, nil
}
