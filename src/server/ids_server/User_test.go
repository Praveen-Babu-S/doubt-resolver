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

func TestCreateUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ctx := context.Background()
	mockDb := db.NewMockDbOperations(controller)
	IdsDbServer := IdsDbServer{Db: mockDb}
	mockDb.EXPECT().CreateUser(gomock.Any())
	expected := &pb.Status{
		Id: "1",
	}
	got, err := IdsDbServer.CreateUser(ctx, &pb.User{
		Name:     "user",
		Email:    "email@email.com",
		Password: "123456",
		Role:     "Student",
		Subject:  "subject-1",
	})
	utils.CheckErr(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Failed test case. Expected %v Got %v ", expected, got)
	}
}

func TestEditUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	ctx := context.Background()
	mockDb := db.NewMockDbOperations(controller)
	IdsDbServer := IdsDbServer{Db: mockDb}
	mockDb.EXPECT().EditUser(gomock.Any(), uint64(1))
	expected := &pb.Status{
		Id: "1",
	}
	got, err := IdsDbServer.EditUser(ctx, &pb.User{
		Id:       1,
		Name:     "user",
		Email:    "email@email.com",
		Password: "123456",
		Role:     "Student",
		Subject:  "subject-1",
	})
	utils.CheckErr(err)
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Failed test case. Expected %v Got %v ", expected, got)
	}
}
