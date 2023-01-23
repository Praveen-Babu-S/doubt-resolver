package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/backend-ids/src/proto"
	_ "github.com/lib/pq"
)

// fetch questions by user id ->client
func GetQuestions(client pb.IdsCRUDClient, user_id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	Id := &pb.Id{Id: user_id}
	stream, err := client.GetQuestions(ctx, Id)
	if err != nil {
		log.Fatalf("%v.GetEmps(_)=_,%v", client, err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetEmps(_)=_,%v", client, err)
		}
		log.Printf("EmpInfo:%v", row)
	}
}

// fetch question by question id (Authorised to student and assignee) ->client
func GetQuestionById(client pb.IdsCRUDClient, q_id uint64, user_id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ids, _ := client.FindIDs(ctx, &pb.Id{Id: q_id})
	if ids.Sid != user_id && ids.Aid != user_id {
		log.Fatalln("You are not authorised for this question!")
	}
	Id := &pb.Id{Id: q_id}
	q, err := client.GetQuestionById(ctx, Id)
	if err != nil {
		log.Fatalf("%v.GetEmps(_)=_,%v", client, err)
	}
	fmt.Println(q)
}
