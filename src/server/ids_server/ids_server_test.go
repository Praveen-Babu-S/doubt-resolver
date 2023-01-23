package server

import (
	"context"
	"reflect"
	"testing"

	db "github.com/backend-ids/src/db"
	pb "github.com/backend-ids/src/proto"
	utils "github.com/backend-ids/utils"
	"github.com/golang/mock/gomock"
)

func TestGetQuestionById(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ctx := context.Background()
	mockDb := db.NewMockDbOperations(controller)
	IdsDbServer := IdsDbServer{Db: mockDb}
	expected := &pb.QuestionById{
		Q: &pb.Question{Subject: "subject-1", Desc: "question description", StudentId: 4, AssigneeId: 3, Id: 3},
		S: &pb.Solution{Desc: "another approach", MentorId: 3, QuestionId: 3, Id: 2},
	}
	mockDb.EXPECT().GetQuestionById(gomock.Any()).Return(expected)
	got, err := IdsDbServer.GetQuestionById(ctx, &pb.Id{
		Id: 1,
	})
	utils.CheckErr(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Failed test case. Expected %v Got %v ", expected, got)
	}
}
