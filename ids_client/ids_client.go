package main

import (
	"context"
	"log"
	"time"

	pb "github.com/backend-ids/ids_proto"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":3030"

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewIdsCRUDClient(conn)
	// runCreateQuestion(client, "Science", "asdfg", "Question4", 2)
	// runCreateSolution(client, "Solution1", 1, 1)
	runCreateComment(client, "Comment 1", 2)
}

// create question
func runCreateQuestion(client pb.IdsCRUDClient, subject string, topic string, desc string, student_id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	question := &pb.Question{Subject: subject, Topic: topic, Desc: desc, Studentid: student_id}
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

// create solution
func runCreateSolution(client pb.IdsCRUDClient, explanation string, mentor_id uint64, question_id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	solution := &pb.Solution{Explanation: explanation, Mentorid: mentor_id, Questionid: question_id}
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

// create comment
func runCreateComment(client pb.IdsCRUDClient, msg string, solution_id uint64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	comment := &pb.Comment{Desc: msg, Solutionid: solution_id}
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
