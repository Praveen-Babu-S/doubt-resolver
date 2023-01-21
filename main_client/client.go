package main

import (
	"fmt"
	"log"
	"time"

	auth "github.com/backend-ids/authentication"
	"github.com/backend-ids/src/client"
	pb "github.com/backend-ids/src/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":3031"

func accessibleMethods() map[string]bool {
	const idsServerPath = "/proto.IdsCRUD/"
	return map[string]bool{
		idsServerPath + "CreateUser":      true,
		idsServerPath + "EditUser":        true,
		idsServerPath + "CreateQuestion":  true,
		idsServerPath + "EditQuestion":    true,
		idsServerPath + "CreateSolution":  true,
		idsServerPath + "EditSolution":    true,
		idsServerPath + "CreateComment":   true,
		idsServerPath + "GetQuestionById": true,
		idsServerPath + "GetQuestions":    true,
	}
}
func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client2 := client.NewAuthClient(conn, "user-1", "123456")
	token, err := client2.Login()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(token)

	interceptor, err := auth.NewClientInterceptor(client2, accessibleMethods(), 30*time.Second)
	clientConnection, err := grpc.Dial(
		"localhost"+port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Println("cannot create auth interceptor:", err)
	}
	client1 := pb.NewIdsCRUDClient(clientConnection)
	// client.CreateQuestion(client1, "subject-1", "question description", 1)
	client.EditQuestion(client1, 1, "subject-1", "modified description", 1, 2, 1)
	// client.CreateSolution(client1, "slution 3", 3, 3)
	// client.EditSolution(client1, 2, "another approach", 2, 3)
	// client.CreateComment(client1, "ok will do", 1, 3)
	// client.GetQuestions(client1, 1)
	// client.GetQuestionById(client1, 1, 2)
	// client.CreateUser(client1, "user-5", "email-5@email.com", "asdqw#$f123@G", "mentor", "")
	// client.UpdateUserDetails(client1, 5, "user-5", "email-5@email.com", "123456", "mentor", "")
}
