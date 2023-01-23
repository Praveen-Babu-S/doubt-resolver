package server

import (
	"context"

	"github.com/backend-ids/src/models"
	pb "github.com/backend-ids/src/proto"
	_ "github.com/lib/pq"
)

// create new user
func (s *IdsServer) CreateUser(ctx context.Context, in *pb.User) (*pb.Status, error) {
	u := models.User{Name: in.Name, Email: in.Email, Password: in.Password, Role: in.Role, Subject: in.Subject}
	s.db.Create(&u)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}

// edit an user
func (s *IdsServer) EditUser(ctx context.Context, in *pb.User) (*pb.Status, error) {
	u := models.User{Name: in.Name, Email: in.Email, Password: in.Password, Role: in.Role, Subject: in.Subject}
	s.db.Model(&models.User{}).Where("id = ?", in.Id).Updates(u)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}
