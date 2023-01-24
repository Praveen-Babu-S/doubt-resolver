package server

import (
	"context"

	pb "github.com/backend-ids/proto"
	"github.com/backend-ids/src/schema/models"
	_ "github.com/lib/pq"
)

// create new comment
func (s *IdsDbServer) CreateComment(ctx context.Context, in *pb.Comment) (*pb.Status, error) {
	c := models.Comment{Msg: in.Msg, SolutionId: in.SolutionId, UserId: in.UserId}
	s.Db.CreateComment(&c)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}
