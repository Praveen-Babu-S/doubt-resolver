package server

import (
	"context"
	"fmt"

	pb "github.com/backend-ids/proto"
	"github.com/backend-ids/src/schema/models"
	_ "github.com/lib/pq"
)

// create new question
func (s *IdsDbServer) CreateQuestion(ctx context.Context, in *pb.Question) (*pb.Status, error) {
	fmt.Println("Called")
	q := models.Question{Desc: in.Desc, Subject: in.Subject, StudentId: in.StudentId}
	s.Db.CreateQuestion(&q)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

// edit a question
func (s *IdsDbServer) EditQuestion(ctx context.Context, in *pb.Question) (*pb.Status, error) {
	res := pb.Status{}
	q := models.Question{Desc: in.Desc, Subject: in.Subject, StudentId: in.StudentId, AssigneeId: in.AssigneeId}
	s.Db.EditQuestion(&q, in.Id)
	res.Id = "1"
	return &res, nil
}
