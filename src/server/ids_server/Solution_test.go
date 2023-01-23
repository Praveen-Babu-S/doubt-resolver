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

func TestCreateSolution(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ctx := context.Background()
	mockDb := db.NewMockDbOperations(controller)
	IdsDbServer := IdsDbServer{Db: mockDb}
	mockDb.EXPECT().CreateSolution(gomock.Any())
	expected := &pb.Status{
		Id: "1",
	}
	got, err := IdsDbServer.CreateSolution(ctx, &pb.Solution{
		Desc:       "solution",
		QuestionId: 1,
		MentorId:   2,
	})
	utils.CheckErr(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Failed test case. Expected %v Got %v ", expected, got)
	}
}

func TestEditSolution(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ctx := context.Background()
	mockDb := db.NewMockDbOperations(controller)
	IdsDbServer := IdsDbServer{Db: mockDb}
	mockDb.EXPECT().EditSolution(gomock.Any(), uint64(1))
	expected := &pb.Status{
		Id: "1",
	}
	got, err := IdsDbServer.EditSolution(ctx, &pb.Solution{
		Id:         1,
		Desc:       "solution",
		QuestionId: 1,
		MentorId:   2,
	})
	utils.CheckErr(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Failed test case. Expected %v Got %v ", expected, got)
	}
}
