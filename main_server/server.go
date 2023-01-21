package main

import (
	"log"
	"net"
	"time"

	auth "github.com/backend-ids/authentication"
	"github.com/backend-ids/src/dbconfig"
	pb "github.com/backend-ids/src/proto"
	server "github.com/backend-ids/src/server"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const port = ":3031"

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

func accessibleRoles() map[string][]string {
	const idsServerPath = "/proto.IdsCRUD/"
	return map[string][]string{
		idsServerPath + "CreateUser":      {"student", "mentor"},
		idsServerPath + "EditUser":        {"student", "mentor"},
		idsServerPath + "CreateQuestion":  {"student"},
		idsServerPath + "EditQuestion":    {"student"},
		idsServerPath + "CreateSolution":  {"mentor"},
		idsServerPath + "EditSolution":    {"student"},
		idsServerPath + "CreateComment":   {"student", "mentor"},
		idsServerPath + "GetQuestionById": {"student"},
		idsServerPath + "GetQuestions":    {"student"},
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	jwtManager := auth.NewJWTManager(secretKey, tokenDuration)
	authServer := server.NewAuthServer(jwtManager, dbconfig.DBSetup())
	idsServer := server.NewIdsServer(dbconfig.DBSetup())
	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())
	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	pb.RegisterAuthServiceServer(s, authServer)
	pb.RegisterIdsCRUDServer(s, idsServer)

	log.Printf("server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faile due to server: %v", err)
	}
}
