// Code generated by MockGen. DO NOT EDIT.
// Source: src/db/dbOperations.go

// Package Mocks is a generated GoMock package.
package dbOperations

import (
	reflect "reflect"

	models "github.com/backend-ids/src/models"
	proto "github.com/backend-ids/src/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockDbOperations is a mock of DbOperations interface.
type MockDbOperations struct {
	ctrl     *gomock.Controller
	recorder *MockDbOperationsMockRecorder
}

// MockDbOperationsMockRecorder is the mock recorder for MockDbOperations.
type MockDbOperationsMockRecorder struct {
	mock *MockDbOperations
}

// NewMockDbOperations creates a new mock instance.
func NewMockDbOperations(ctrl *gomock.Controller) *MockDbOperations {
	mock := &MockDbOperations{ctrl: ctrl}
	mock.recorder = &MockDbOperationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDbOperations) EXPECT() *MockDbOperationsMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockDbOperations) CreateComment(arg0 *models.Comment) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateComment", arg0)
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockDbOperationsMockRecorder) CreateComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockDbOperations)(nil).CreateComment), arg0)
}

// CreateQuestion mocks base method.
func (m *MockDbOperations) CreateQuestion(arg0 *models.Question) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateQuestion", arg0)
}

// CreateQuestion indicates an expected call of CreateQuestion.
func (mr *MockDbOperationsMockRecorder) CreateQuestion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuestion", reflect.TypeOf((*MockDbOperations)(nil).CreateQuestion), arg0)
}

// CreateSolution mocks base method.
func (m *MockDbOperations) CreateSolution(arg0 *models.Solution) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateSolution", arg0)
}

// CreateSolution indicates an expected call of CreateSolution.
func (mr *MockDbOperationsMockRecorder) CreateSolution(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSolution", reflect.TypeOf((*MockDbOperations)(nil).CreateSolution), arg0)
}

// CreateUser mocks base method.
func (m *MockDbOperations) CreateUser(arg0 *models.User) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateUser", arg0)
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockDbOperationsMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDbOperations)(nil).CreateUser), arg0)
}

// EditQuestion mocks base method.
func (m *MockDbOperations) EditQuestion(arg0 *models.Question, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EditQuestion", arg0, arg1)
}

// EditQuestion indicates an expected call of EditQuestion.
func (mr *MockDbOperationsMockRecorder) EditQuestion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditQuestion", reflect.TypeOf((*MockDbOperations)(nil).EditQuestion), arg0, arg1)
}

// EditSolution mocks base method.
func (m *MockDbOperations) EditSolution(arg0 *models.Solution, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EditSolution", arg0, arg1)
}

// EditSolution indicates an expected call of EditSolution.
func (mr *MockDbOperationsMockRecorder) EditSolution(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditSolution", reflect.TypeOf((*MockDbOperations)(nil).EditSolution), arg0, arg1)
}

// EditUser mocks base method.
func (m *MockDbOperations) EditUser(arg0 *models.User, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EditUser", arg0, arg1)
}

// EditUser indicates an expected call of EditUser.
func (mr *MockDbOperationsMockRecorder) EditUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditUser", reflect.TypeOf((*MockDbOperations)(nil).EditUser), arg0, arg1)
}

// FindIDs mocks base method.
func (m *MockDbOperations) FindIDs(arg0 uint64) *proto.Ids {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIDs", arg0)
	ret0, _ := ret[0].(*proto.Ids)
	return ret0
}

// FindIDs indicates an expected call of FindIDs.
func (mr *MockDbOperationsMockRecorder) FindIDs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIDs", reflect.TypeOf((*MockDbOperations)(nil).FindIDs), arg0)
}

// FindQID mocks base method.
func (m *MockDbOperations) FindQID(arg0 uint64) *proto.Id {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindQID", arg0)
	ret0, _ := ret[0].(*proto.Id)
	return ret0
}

// FindQID indicates an expected call of FindQID.
func (mr *MockDbOperationsMockRecorder) FindQID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindQID", reflect.TypeOf((*MockDbOperations)(nil).FindQID), arg0)
}

// GetQuestionById mocks base method.
func (m *MockDbOperations) GetQuestionById(arg0 uint64) *proto.QuestionById {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestionById", arg0)
	ret0, _ := ret[0].(*proto.QuestionById)
	return ret0
}

// GetQuestionById indicates an expected call of GetQuestionById.
func (mr *MockDbOperationsMockRecorder) GetQuestionById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestionById", reflect.TypeOf((*MockDbOperations)(nil).GetQuestionById), arg0)
}

// GetQuestions mocks base method.
func (m *MockDbOperations) GetQuestions(arg0 uint64) []proto.Question {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestions", arg0)
	ret0, _ := ret[0].([]proto.Question)
	return ret0
}

// GetQuestions indicates an expected call of GetQuestions.
func (mr *MockDbOperationsMockRecorder) GetQuestions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestions", reflect.TypeOf((*MockDbOperations)(nil).GetQuestions), arg0)
}
