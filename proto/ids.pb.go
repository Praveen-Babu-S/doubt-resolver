// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/ids.proto

package backend_ids

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject    string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Desc       string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	StudentId  uint64 `protobuf:"varint,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	AssigneeId uint64 `protobuf:"varint,4,opt,name=assignee_id,json=assigneeId,proto3" json:"assignee_id,omitempty"`
	Id         uint64 `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ids_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ids_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_proto_ids_proto_rawDescGZIP(), []int{0}
}

func (x *Question) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Question) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Question) GetStudentId() uint64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *Question) GetAssigneeId() uint64 {
	if x != nil {
		return x.AssigneeId
	}
	return 0
}

func (x *Question) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Solution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Desc       string `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	MentorId   uint64 `protobuf:"varint,2,opt,name=mentor_id,json=mentorId,proto3" json:"mentor_id,omitempty"`
	QuestionId uint64 `protobuf:"varint,3,opt,name=question_id,json=questionId,proto3" json:"question_id,omitempty"`
	Id         uint64 `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Solution) Reset() {
	*x = Solution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ids_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Solution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solution) ProtoMessage() {}

func (x *Solution) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ids_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solution.ProtoReflect.Descriptor instead.
func (*Solution) Descriptor() ([]byte, []int) {
	return file_proto_ids_proto_rawDescGZIP(), []int{1}
}

func (x *Solution) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Solution) GetMentorId() uint64 {
	if x != nil {
		return x.MentorId
	}
	return 0
}

func (x *Solution) GetQuestionId() uint64 {
	if x != nil {
		return x.QuestionId
	}
	return 0
}

func (x *Solution) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg        string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	SolutionId uint64 `protobuf:"varint,2,opt,name=solution_id,json=solutionId,proto3" json:"solution_id,omitempty"`
	UserId     uint64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ids_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ids_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_proto_ids_proto_rawDescGZIP(), []int{2}
}

func (x *Comment) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Comment) GetSolutionId() uint64 {
	if x != nil {
		return x.SolutionId
	}
	return 0
}

func (x *Comment) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ids_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ids_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_ids_proto_rawDescGZIP(), []int{3}
}

func (x *Status) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ids_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ids_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_proto_ids_proto_rawDescGZIP(), []int{4}
}

func (x *Id) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Ids struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sid uint64 `protobuf:"varint,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Aid uint64 `protobuf:"varint,2,opt,name=aid,proto3" json:"aid,omitempty"`
}

func (x *Ids) Reset() {
	*x = Ids{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ids_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ids) ProtoMessage() {}

func (x *Ids) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ids_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ids.ProtoReflect.Descriptor instead.
func (*Ids) Descriptor() ([]byte, []int) {
	return file_proto_ids_proto_rawDescGZIP(), []int{5}
}

func (x *Ids) GetSid() uint64 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *Ids) GetAid() uint64 {
	if x != nil {
		return x.Aid
	}
	return 0
}

type QuestionById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Q *Question  `protobuf:"bytes,1,opt,name=Q,proto3" json:"Q,omitempty"`
	S *Solution  `protobuf:"bytes,2,opt,name=S,proto3" json:"S,omitempty"`
	C []*Comment `protobuf:"bytes,3,rep,name=C,proto3" json:"C,omitempty"`
}

func (x *QuestionById) Reset() {
	*x = QuestionById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ids_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionById) ProtoMessage() {}

func (x *QuestionById) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ids_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionById.ProtoReflect.Descriptor instead.
func (*QuestionById) Descriptor() ([]byte, []int) {
	return file_proto_ids_proto_rawDescGZIP(), []int{6}
}

func (x *QuestionById) GetQ() *Question {
	if x != nil {
		return x.Q
	}
	return nil
}

func (x *QuestionById) GetS() *Solution {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *QuestionById) GetC() []*Comment {
	if x != nil {
		return x.C
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Role     string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Subject  string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	Id       uint64 `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ids_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ids_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_ids_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_proto_ids_proto protoreflect.FileDescriptor

var file_proto_ids_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x08, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x6c, 0x0a, 0x08, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x55, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x29, 0x0a, 0x03, 0x49, 0x64, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x61, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x01, 0x51, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x01, 0x51, 0x12, 0x1d, 0x0a, 0x01, 0x53, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x01,
	0x53, 0x12, 0x1c, 0x0a, 0x01, 0x43, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x01, 0x43, 0x22,
	0x8a, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0xf3, 0x03, 0x0a,
	0x07, 0x49, 0x64, 0x73, 0x43, 0x52, 0x55, 0x44, 0x12, 0x28, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x26, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x1a,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e,
	0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x12,
	0x1f, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x51, 0x49, 0x44, 0x12, 0x09, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x44, 0x73, 0x12, 0x09, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49,
	0x64, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x69, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ids_proto_rawDescOnce sync.Once
	file_proto_ids_proto_rawDescData = file_proto_ids_proto_rawDesc
)

func file_proto_ids_proto_rawDescGZIP() []byte {
	file_proto_ids_proto_rawDescOnce.Do(func() {
		file_proto_ids_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ids_proto_rawDescData)
	})
	return file_proto_ids_proto_rawDescData
}

var file_proto_ids_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_ids_proto_goTypes = []interface{}{
	(*Question)(nil),     // 0: proto.Question
	(*Solution)(nil),     // 1: proto.Solution
	(*Comment)(nil),      // 2: proto.Comment
	(*Status)(nil),       // 3: proto.Status
	(*Id)(nil),           // 4: proto.Id
	(*Ids)(nil),          // 5: proto.Ids
	(*QuestionById)(nil), // 6: proto.QuestionById
	(*User)(nil),         // 7: proto.User
}
var file_proto_ids_proto_depIdxs = []int32{
	0,  // 0: proto.QuestionById.Q:type_name -> proto.Question
	1,  // 1: proto.QuestionById.S:type_name -> proto.Solution
	2,  // 2: proto.QuestionById.C:type_name -> proto.Comment
	7,  // 3: proto.IdsCRUD.CreateUser:input_type -> proto.User
	7,  // 4: proto.IdsCRUD.EditUser:input_type -> proto.User
	0,  // 5: proto.IdsCRUD.CreateQuestion:input_type -> proto.Question
	1,  // 6: proto.IdsCRUD.CreateSolution:input_type -> proto.Solution
	2,  // 7: proto.IdsCRUD.CreateComment:input_type -> proto.Comment
	0,  // 8: proto.IdsCRUD.EditQuestion:input_type -> proto.Question
	4,  // 9: proto.IdsCRUD.GetQuestionById:input_type -> proto.Id
	1,  // 10: proto.IdsCRUD.EditSolution:input_type -> proto.Solution
	4,  // 11: proto.IdsCRUD.GetQuestions:input_type -> proto.Id
	4,  // 12: proto.IdsCRUD.FindQID:input_type -> proto.Id
	4,  // 13: proto.IdsCRUD.FindIDs:input_type -> proto.Id
	3,  // 14: proto.IdsCRUD.CreateUser:output_type -> proto.Status
	3,  // 15: proto.IdsCRUD.EditUser:output_type -> proto.Status
	3,  // 16: proto.IdsCRUD.CreateQuestion:output_type -> proto.Status
	3,  // 17: proto.IdsCRUD.CreateSolution:output_type -> proto.Status
	3,  // 18: proto.IdsCRUD.CreateComment:output_type -> proto.Status
	3,  // 19: proto.IdsCRUD.EditQuestion:output_type -> proto.Status
	6,  // 20: proto.IdsCRUD.GetQuestionById:output_type -> proto.QuestionById
	3,  // 21: proto.IdsCRUD.EditSolution:output_type -> proto.Status
	0,  // 22: proto.IdsCRUD.GetQuestions:output_type -> proto.Question
	4,  // 23: proto.IdsCRUD.FindQID:output_type -> proto.Id
	5,  // 24: proto.IdsCRUD.FindIDs:output_type -> proto.Ids
	14, // [14:25] is the sub-list for method output_type
	3,  // [3:14] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_ids_proto_init() }
func file_proto_ids_proto_init() {
	if File_proto_ids_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ids_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Question); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ids_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Solution); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ids_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ids_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ids_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ids_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ids); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ids_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestionById); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_ids_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ids_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ids_proto_goTypes,
		DependencyIndexes: file_proto_ids_proto_depIdxs,
		MessageInfos:      file_proto_ids_proto_msgTypes,
	}.Build()
	File_proto_ids_proto = out.File
	file_proto_ids_proto_rawDesc = nil
	file_proto_ids_proto_goTypes = nil
	file_proto_ids_proto_depIdxs = nil
}
