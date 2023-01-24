package main

import (
	"log"
	"net"
	"time"

	auth "github.com/backend-ids/authentication"
	pb "github.com/backend-ids/proto"
	"github.com/backend-ids/src/schema/dbconfig"
	auth_server "github.com/backend-ids/src/server/auth_server"
	ids_server "github.com/backend-ids/src/server/ids_server"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

// /home/praveenbabu/Desktop/backend-ids/src/server/ids_server

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
		idsServerPath + "EditSolution":    {"mentor"},
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
	authServer := auth_server.NewAuthServer(jwtManager, dbconfig.DBSetup())
	idsDbServer := ids_server.NewIdsDbserver(dbconfig.DBSetup())
	interceptor := auth.NewAuthInterceptor(jwtManager, accessibleRoles())
	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	pb.RegisterAuthServiceServer(s, authServer)
	pb.RegisterIdsCRUDServer(s, idsDbServer)

	log.Printf("server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faile due to server: %v", err)
	}
}
