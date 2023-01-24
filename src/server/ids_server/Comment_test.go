package server

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/backend-ids/proto"
	db "github.com/backend-ids/src/schema/db"
	utils "github.com/backend-ids/utils"
	"github.com/golang/mock/gomock"
)

func TestCreateComment(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ctx := context.Background()
	mockDb := db.NewMockDbOperations(controller)
	IdsDbServer := IdsDbServer{Db: mockDb}
	mockDb.EXPECT().CreateComment(gomock.Any())
	expected := &pb.Status{
		Id: "1",
	}
	got, err := IdsDbServer.CreateComment(ctx, &pb.Comment{
		Msg:        "Comment",
		UserId:     1,
		SolutionId: 1,
	})
	utils.CheckErr(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Failed test case. Expected %v Got %v ", expected, got)
	}
}
