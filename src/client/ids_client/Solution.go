package client

import (
	"context"
	"log"
	"time"

	pb "github.com/backend-ids/src/proto"
	_ "github.com/lib/pq"
)

// create solution ->client
func CreateSolution(client pb.IdsCRUDClient, explanation string, user_id uint64, question_id uint64) {
	//authorisation(only mentor who got assigned this question can create)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//authorisation(only mentor who got assigned this question can edit)
	ids, _ := client.FindIDs(ctx, &pb.Id{Id: question_id})
	if user_id != ids.Aid {
		log.Fatalln("Not Authorised to create solution!")
		return
	}
	solution := &pb.Solution{Desc: explanation, QuestionId: question_id, MentorId: user_id}
	res, err := client.CreateSolution(ctx, solution)
	if err != nil {
		log.Fatalf("Unable to create solution in client %v", err)
	}
	if res.Id != "1" {
		log.Fatalf("Unable to create solution in server %v", err)
	} else {
		log.Fatalln("Successfully created solution!")
	}
}

// update Solution ->client
func EditSolution(client pb.IdsCRUDClient, sol_id uint64, explanation string, user_id uint64, question_id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//authorisation(only mentor who got assigned this question can edit)
	ids, _ := client.FindIDs(ctx, &pb.Id{Id: question_id})
	if user_id != ids.Aid {
		log.Fatalln("Not Authorised to edit solution!")
		return
	}
	solution := &pb.Solution{Id: sol_id, Desc: explanation, QuestionId: question_id, MentorId: user_id}
	res, err := client.EditSolution(ctx, solution)
	if err != nil {
		log.Fatalf("Unable to edit solution in client %v", err)
	}
	if res.Id != "1" {
		log.Fatalf("Unable to edit solution in server %v", err)
	} else {
		log.Fatalln("Successfully updated solution!")
	}

}
