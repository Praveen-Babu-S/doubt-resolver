package main

import (
	"fmt"
	"log"
	"time"

	auth "github.com/backend-ids/authentication"
	auth_client "github.com/backend-ids/src/client/auth_client"
	ids_client "github.com/backend-ids/src/client/ids_client"
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
	//client2 --->auth_client
	client2 := auth_client.NewAuthClient(conn, "user-4", "asdqw#$f123@G")
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
	//client1 ---> ids_client
	client1 := pb.NewIdsCRUDClient(clientConnection)
	fmt.Println(client1)
	// ids_client.CreateQuestion(client1, "subject-1", "new question", 6)
	// ids_client.EditQuestion(client1, 1, "subject-1", "modified description", 1, 2, 1)
	// ids_client.CreateSolution(client1, "slution 3", 3, 3)
	// ids_client.EditSolution(client1, 2, "another approach", 2, 3)
	// ids_client.CreateComment(client1, "ok will do", 1, 3)
	// ids_client.GetQuestions(client1, 1)
	ids_client.GetQuestionById(client1, 3, 4)
	// ids_client.CreateUser(client1, "user-6", "email-6@email.com", "123456", "student", "")
	// ids_client.UpdateUserDetails(client1, 5, "user-5", "email-5@email.com", "123123", "mentor", "")
}
