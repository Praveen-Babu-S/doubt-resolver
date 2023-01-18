package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/backend-ids/ids_proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":3031"

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewIdsCRUDClient(conn)
	// runCreateQuestion(client, "Subject-2", "desc 4", 1, 4)
	// runEditQuestion(client, 10, "Subject-9", "Desc-9", 1, 4)
	// runCreateSolution(client, "this is solution", 4, 6)
	// runEditSolution(client, 3, "This is new method", 4, 6)
	// runCreateComment(client, "some more explanation", 3, 1)
	// GetQuestions(client, 3)
	GetQuestionById(client, 5, 4)
}

// create question
func runCreateQuestion(client pb.IdsCRUDClient, subject string, desc string, student_id uint64, assignee_id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	question := &pb.Question{Subject: subject, Desc: desc, StudentId: student_id, AssigneeId: assignee_id}
	res, err := client.CreateQuestion(ctx, question)
	if err != nil {
		log.Fatalf("Unable to create question in client %v", err)
	}
	if res.Id != "1" {
		log.Fatalf("Unable to create question in server %v", err)
	} else {
		log.Fatalln("Successfully created Question from Client!")
	}
}

// update question
func runEditQuestion(client pb.IdsCRUDClient, id uint64, subject string, desc string, student_id uint64, assignee_id uint64, user_id uint64) {
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
		log.Fatalln("Successfully updated Question from Client!")
	}
}

// create solution
func runCreateSolution(client pb.IdsCRUDClient, explanation string, user_id uint64, question_id uint64) {
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
		log.Fatalln("Successfully created solution from Client!")
	}
}

// update Solution
func runEditSolution(client pb.IdsCRUDClient, sol_id uint64, explanation string, user_id uint64, question_id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//authorisation(only mentor who got assigned this question can edit)
	ids, _ := client.FindIDs(ctx, &pb.Id{Id: question_id})
	if user_id != ids.Aid {
		log.Fatalln("Not Authorised to create solution!")
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
		log.Fatalln("Successfully updated solution from Client!")
	}

}

// create comment
func runCreateComment(client pb.IdsCRUDClient, msg string, solution_id uint64, user_id uint64) {
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
		log.Fatalln("Successfully created comment from Client!")
	}
}

// fetch questions by user id
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

// fetch question by question id (Authorised to student and assignee)
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
