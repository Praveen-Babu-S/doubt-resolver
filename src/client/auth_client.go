package client

import (
	"context"
	"time"

	pb "github.com/backend-ids/src/proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}
func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Name:     client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}

// func Login(client pb.AuthServiceClient, username string, password string) {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	token, err := client.Login(ctx, &pb.LoginRequest{Name: username, Password: password})
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 	}
// 	fmt.Println(token)
// }
