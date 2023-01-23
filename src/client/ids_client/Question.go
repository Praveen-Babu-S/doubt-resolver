package client

import (
	"context"
	"log"
	"time"

	pb "github.com/backend-ids/src/proto"
	_ "github.com/lib/pq"
)

// create question -> client
func CreateQuestion(client pb.IdsCRUDClient, subject string, desc string, student_id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	question := &pb.Question{Subject: subject, Desc: desc, StudentId: student_id}
	res, err := client.CreateQuestion(ctx, question)
	if err != nil {
		log.Fatalf("Unable to create question in client %v", err)
	}
	if res.Id != "1" {
		log.Fatalf("Unable to create question in server %v", err)
	} else {
		log.Fatalln("Successfully created Question!")
	}
}

// update question -> client
func EditQuestion(client pb.IdsCRUDClient, id uint64, subject string, desc string, student_id uint64, assignee_id uint64, user_id uint64) {
	//authorisation(only student who created question can edit)
	if user_id != student_id {
		log.Fatalln("Not Authorised to edit question!")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	question := &pb.Question{Id: id, Subject: subject, Desc: desc, StudentId: student_id, AssigneeId: assignee_id}
	res, err := client.EditQuestion(ctx, question)
	if err != nil {
		log.Fatalf("Unable to edit question in client %v", err)
	}
	if res.Id != "1" {
		log.Fatalf("Unable to edit question in server %v", err)
	} else {
		log.Fatalln("Successfully updated Question!")
	}
}
