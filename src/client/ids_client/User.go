package client

import (
	"context"
	"log"
	"time"

	pb "github.com/backend-ids/proto"
	_ "github.com/lib/pq"
)

// create an user ->client
func CreateUser(client pb.IdsCRUDClient, name string, email string, password string, role string, subject string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	u := &pb.User{Name: name, Email: email, Password: password, Role: role, Subject: subject}
	res, err := client.CreateUser(ctx, u)
	if err != nil {
		log.Fatalf("Unable to create user in client %v", err)
	}
	if res.Id != "1" {
		log.Fatalf("Unable to create user in server %v", err)
	} else {
		log.Fatalln("Successfully created user!")
	}
}

// update user details ->client
func UpdateUserDetails(client pb.IdsCRUDClient, id uint64, name string, email string, password string, role string, subject string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	u := &pb.User{Id: id, Name: name, Email: email, Password: password, Role: role, Subject: subject}
	res, err := client.EditUser(ctx, u)
	if err != nil {
		log.Fatalf("Unable to update user in client %v", err)
	}
	if res.Id != "1" {
		log.Fatalf("Unable to update user in server %v", err)
	} else {
		log.Fatalln("Successfully updated user!")
	}
}
