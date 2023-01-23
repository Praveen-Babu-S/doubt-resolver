package server

import (
	"context"

	"github.com/backend-ids/src/models"
	pb "github.com/backend-ids/src/proto"
	_ "github.com/lib/pq"
)

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
