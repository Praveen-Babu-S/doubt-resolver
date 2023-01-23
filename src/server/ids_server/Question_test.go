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

func TestCreateQuestion(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ctx := context.Background()
	mockDb := db.NewMockDbOperations(controller)
	IdsDbServer := IdsDbServer{Db: mockDb}
	mockDb.EXPECT().CreateQuestion(gomock.Any())
	expected := &pb.Status{
		Id: "1",
	}
	got, err := IdsDbServer.CreateQuestion(ctx, &pb.Question{
		Desc:       "solution",
		StudentId:  1,
		AssigneeId: 2,
	})
	utils.CheckErr(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Failed test case. Expected %v Got %v ", expected, got)
	}
}

func TestEditQuestion(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ctx := context.Background()
	mockDb := db.NewMockDbOperations(controller)
	IdsDbServer := IdsDbServer{Db: mockDb}
	mockDb.EXPECT().EditQuestion(gomock.Any(), uint64(1))
	expected := &pb.Status{
		Id: "1",
	}
	got, err := IdsDbServer.EditQuestion(ctx, &pb.Question{
		Id:         1,
		Desc:       "question",
		Subject:    "subject-1",
		StudentId:  1,
		AssigneeId: 2,
	})
	utils.CheckErr(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Failed test case. Expected %v Got %v ", expected, got)
	}
}
