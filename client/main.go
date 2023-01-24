package main

import (
	"fmt"
	"log"
	"time"

	auth "github.com/backend-ids/authentication"
	pb "github.com/backend-ids/proto"
	auth_client "github.com/backend-ids/src/client/auth_client"
	ids_client "github.com/backend-ids/src/client/ids_client"
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
	client2 := auth_client.NewAuthClient(conn, "user-2", "123456")
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
	// ids_client.CreateQuestion(client1, "subject-1", "question by Praveen", 7)
	// ids_client.EditQuestion(client1, 7, "subject-1", "modified description", 7, 2, 7)
	// ids_client.CreateSolution(client1, "new solution", 2, 7)
	ids_client.EditSolution(client1, 3, "solution a approach", 2, 7)
	// ids_client.CreateComment(client1, "ok will do", 1, 3)
	// ids_client.GetQuestions(client1, 1)
	// ids_client.GetQuestionById(client1, 3, 4)
	// ids_client.CreateUser(client1, "user-7", "praveen@beautifulcode.in", "123456", "student", "")
	// ids_client.UpdateUserDetails(client1, 5, "Praveen", "praveen@beautifulcode.in", "123456", "", "")
}
