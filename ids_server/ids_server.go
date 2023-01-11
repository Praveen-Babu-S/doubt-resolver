package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/backend-ids/dbconfig"
	pb "github.com/backend-ids/ids_proto"
	"github.com/backend-ids/models"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const port = ":3030"

type idsServer struct {
	pb.UnimplementedIdsCRUDServer
	db *gorm.DB
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listes: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterIdsCRUDServer(s, &idsServer{db: dbconfig.DBSetup()})
	log.Printf("server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faile due to server: %v", err)
	}
}

func (s *idsServer) CreateQuestion(ctx context.Context, in *pb.Question) (*pb.Status, error) {
	q := models.Question{Desc: in.Desc, Subject: in.Subject, Topic: in.Topic, StudentId: uint(in.Studentid)}
	fmt.Println(q)
	s.db.Create(&q)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}
func (s *idsServer) CreateComment(ctx context.Context, in *pb.Comment) (*pb.Status, error) {
	c := models.Comment{Msg: in.Desc}
	s.db.Create(&c)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}
func (s *idsServer) CreateSolution(ctx context.Context, in *pb.Solution) (*pb.Status, error) {
	q := models.Solution{Explanation: in.Explanation}
	s.db.Create(&q)
	res := pb.Status{}
	res.Id = "1"
	return &res, nil
}
