package client

import (
	"context"
	"log"
	"time"

	pb "github.com/backend-ids/proto"
	_ "github.com/lib/pq"
)

// create comment-client
func CreateComment(client pb.IdsCRUDClient, msg string, solution_id uint64, user_id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	q_id, _ := client.FindQID(ctx, &pb.Id{Id: solution_id})
	ids, _ := client.FindIDs(ctx, &pb.Id{Id: q_id.Id})
	if ids.Sid != user_id && ids.Aid != user_id {
		log.Fatalln("You are not authorised for this question!")
	}
	comment := &pb.Comment{Msg: msg, SolutionId: solution_id, UserId: user_id}
	res, err := client.CreateComment(ctx, comment)
	if err != nil {
		log.Fatalf("Unable to create comment in client %v", err)
	}
	if res.Id != "1" {
		log.Fatalf("Unable to create comment in server %v", err)
	} else {
		log.Fatalln("Successfully created comment!")
	}
}
