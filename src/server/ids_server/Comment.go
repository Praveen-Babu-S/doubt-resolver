package server

import (
	"context"

	"github.com/backend-ids/src/models"
	pb "github.com/backend-ids/src/proto"
	_ "github.com/lib/pq"
)

// create new comment
func (s *IdsServer) CreateComment(ctx context.Context, in *pb.Comment) (*pb.Status, error) {
	c := models.Comment{Msg: in.Msg, SolutionId: in.SolutionId, UserId: in.UserId}
	s.db.Create(&c)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}