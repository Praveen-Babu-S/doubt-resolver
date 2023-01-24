package server

import (
	"context"

	pb "github.com/backend-ids/proto"
	"github.com/backend-ids/src/schema/models"
	_ "github.com/lib/pq"
)

// create new user
func (s *IdsDbServer) CreateUser(ctx context.Context, in *pb.User) (*pb.Status, error) {
	u := models.User{Name: in.Name, Email: in.Email, Password: in.Password, Role: in.Role, Subject: in.Subject}
	s.Db.CreateUser(&u)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

// edit an user
func (s *IdsDbServer) EditUser(ctx context.Context, in *pb.User) (*pb.Status, error) {
	u := models.User{Name: in.Name, Email: in.Email, Password: in.Password, Role: in.Role, Subject: in.Subject}
	s.Db.EditUser(&u, in.Id)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}
