// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.5
// source: api/investor/investor.proto

package investor

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListInvestorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	PageNum   uint32 `protobuf:"varint,2,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize  uint32 `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListInvestorRequest) Reset() {
	*x = ListInvestorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_investor_investor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvestorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvestorRequest) ProtoMessage() {}

func (x *ListInvestorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_investor_investor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvestorRequest.ProtoReflect.Descriptor instead.
func (*ListInvestorRequest) Descriptor() ([]byte, []int) {
	return file_api_investor_investor_proto_rawDescGZIP(), []int{0}
}

func (x *ListInvestorRequest) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *ListInvestorRequest) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListInvestorRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListInvestorReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Investors []*ListInvestorReply_Investor `protobuf:"bytes,1,rep,name=investors,proto3" json:"investors,omitempty"`
	Total     int32                         `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListInvestorReply) Reset() {
	*x = ListInvestorReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_investor_investor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvestorReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvestorReply) ProtoMessage() {}

func (x *ListInvestorReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_investor_investor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvestorReply.ProtoReflect.Descriptor instead.
func (*ListInvestorReply) Descriptor() ([]byte, []int) {
	return file_api_investor_investor_proto_rawDescGZIP(), []int{1}
}

func (x *ListInvestorReply) GetInvestors() []*ListInvestorReply_Investor {
	if x != nil {
		return x.Investors
	}
	return nil
}

func (x *ListInvestorReply) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateInvestorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	UserId   uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateInvestorRequest) Reset() {
	*x = CreateInvestorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_investor_investor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvestorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvestorRequest) ProtoMessage() {}

func (x *CreateInvestorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_investor_investor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvestorRequest.ProtoReflect.Descriptor instead.
func (*CreateInvestorRequest) Descriptor() ([]byte, []int) {
	return file_api_investor_investor_proto_rawDescGZIP(), []int{2}
}

func (x *CreateInvestorRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *CreateInvestorRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AddMoneyInvestorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Money  int32  `protobuf:"varint,1,opt,name=money,proto3" json:"money,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *AddMoneyInvestorRequest) Reset() {
	*x = AddMoneyInvestorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_investor_investor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddMoneyInvestorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMoneyInvestorRequest) ProtoMessage() {}

func (x *AddMoneyInvestorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_investor_investor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMoneyInvestorRequest.ProtoReflect.Descriptor instead.
func (*AddMoneyInvestorRequest) Descriptor() ([]byte, []int) {
	return file_api_investor_investor_proto_rawDescGZIP(), []int{3}
}

func (x *AddMoneyInvestorRequest) GetMoney() int32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *AddMoneyInvestorRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RemoveMoneyInvestorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Money  int32  `protobuf:"varint,1,opt,name=money,proto3" json:"money,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RemoveMoneyInvestorRequest) Reset() {
	*x = RemoveMoneyInvestorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_investor_investor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveMoneyInvestorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveMoneyInvestorRequest) ProtoMessage() {}

func (x *RemoveMoneyInvestorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_investor_investor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveMoneyInvestorRequest.ProtoReflect.Descriptor instead.
func (*RemoveMoneyInvestorRequest) Descriptor() ([]byte, []int) {
	return file_api_investor_investor_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveMoneyInvestorRequest) GetMoney() int32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *RemoveMoneyInvestorRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type SetMoneyInvestorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Money  int32  `protobuf:"varint,1,opt,name=money,proto3" json:"money,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SetMoneyInvestorRequest) Reset() {
	*x = SetMoneyInvestorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_investor_investor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMoneyInvestorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMoneyInvestorRequest) ProtoMessage() {}

func (x *SetMoneyInvestorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_investor_investor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMoneyInvestorRequest.ProtoReflect.Descriptor instead.
func (*SetMoneyInvestorRequest) Descriptor() ([]byte, []int) {
	return file_api_investor_investor_proto_rawDescGZIP(), []int{5}
}

func (x *SetMoneyInvestorRequest) GetMoney() int32 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *SetMoneyInvestorRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteInvestorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteInvestorRequest) Reset() {
	*x = DeleteInvestorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_investor_investor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInvestorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInvestorRequest) ProtoMessage() {}

func (x *DeleteInvestorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_investor_investor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInvestorRequest.ProtoReflect.Descriptor instead.
func (*DeleteInvestorRequest) Descriptor() ([]byte, []int) {
	return file_api_investor_investor_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteInvestorRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListInvestorReply_Investor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName string `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
}

func (x *ListInvestorReply_Investor) Reset() {
	*x = ListInvestorReply_Investor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_investor_investor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInvestorReply_Investor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInvestorReply_Investor) ProtoMessage() {}

func (x *ListInvestorReply_Investor) ProtoReflect() protoreflect.Message {
	mi := &file_api_investor_investor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInvestorReply_Investor.ProtoReflect.Descriptor instead.
func (*ListInvestorReply_Investor) Descriptor() ([]byte, []int) {
	return file_api_investor_investor_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListInvestorReply_Investor) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

var File_api_investor_investor_proto protoreflect.FileDescriptor

var file_api_investor_investor_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x69,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x73, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x09,
	0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x2e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x1a, 0x27, 0x0a, 0x08, 0x49, 0x6e,
	0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x17, 0x41,
	0x64, 0x64, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x05,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x1a, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x6d,
	0x6f, 0x6e, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x4d, 0x6f, 0x6e,
	0x65, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x30, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x32, 0xfb, 0x03, 0x0a, 0x08, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x4d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f,
	0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x51, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x57, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d, 0x6f, 0x6e, 0x65,
	0x79, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x51, 0x0a, 0x10, 0x53,
	0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x12,
	0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x53,
	0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x12, 0x21,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x37, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x50, 0x01, 0x5a, 0x25, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x3b, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_investor_investor_proto_rawDescOnce sync.Once
	file_api_investor_investor_proto_rawDescData = file_api_investor_investor_proto_rawDesc
)

func file_api_investor_investor_proto_rawDescGZIP() []byte {
	file_api_investor_investor_proto_rawDescOnce.Do(func() {
		file_api_investor_investor_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_investor_investor_proto_rawDescData)
	})
	return file_api_investor_investor_proto_rawDescData
}

var file_api_investor_investor_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_investor_investor_proto_goTypes = []interface{}{
	(*ListInvestorRequest)(nil),        // 0: api.investor.ListInvestorRequest
	(*ListInvestorReply)(nil),          // 1: api.investor.ListInvestorReply
	(*CreateInvestorRequest)(nil),      // 2: api.investor.CreateInvestorRequest
	(*AddMoneyInvestorRequest)(nil),    // 3: api.investor.AddMoneyInvestorRequest
	(*RemoveMoneyInvestorRequest)(nil), // 4: api.investor.RemoveMoneyInvestorRequest
	(*SetMoneyInvestorRequest)(nil),    // 5: api.investor.SetMoneyInvestorRequest
	(*DeleteInvestorRequest)(nil),      // 6: api.investor.DeleteInvestorRequest
	(*ListInvestorReply_Investor)(nil), // 7: api.investor.ListInvestorReply.Investor
	(*emptypb.Empty)(nil),              // 8: google.protobuf.Empty
}
var file_api_investor_investor_proto_depIdxs = []int32{
	7, // 0: api.investor.ListInvestorReply.investors:type_name -> api.investor.ListInvestorReply.Investor
	2, // 1: api.investor.Investor.CreateInvestor:input_type -> api.investor.CreateInvestorRequest
	6, // 2: api.investor.Investor.DeleteInvestor:input_type -> api.investor.DeleteInvestorRequest
	3, // 3: api.investor.Investor.AddMoneyInvestor:input_type -> api.investor.AddMoneyInvestorRequest
	4, // 4: api.investor.Investor.RemoveMoneyInvestor:input_type -> api.investor.RemoveMoneyInvestorRequest
	5, // 5: api.investor.Investor.SetMoneyInvestor:input_type -> api.investor.SetMoneyInvestorRequest
	0, // 6: api.investor.Investor.ListInvestor:input_type -> api.investor.ListInvestorRequest
	8, // 7: api.investor.Investor.CreateInvestor:output_type -> google.protobuf.Empty
	8, // 8: api.investor.Investor.DeleteInvestor:output_type -> google.protobuf.Empty
	8, // 9: api.investor.Investor.AddMoneyInvestor:output_type -> google.protobuf.Empty
	8, // 10: api.investor.Investor.RemoveMoneyInvestor:output_type -> google.protobuf.Empty
	8, // 11: api.investor.Investor.SetMoneyInvestor:output_type -> google.protobuf.Empty
	1, // 12: api.investor.Investor.ListInvestor:output_type -> api.investor.ListInvestorReply
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_investor_investor_proto_init() }
func file_api_investor_investor_proto_init() {
	if File_api_investor_investor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_investor_investor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvestorRequest); i {
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
		file_api_investor_investor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvestorReply); i {
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
		file_api_investor_investor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateInvestorRequest); i {
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
		file_api_investor_investor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddMoneyInvestorRequest); i {
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
		file_api_investor_investor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveMoneyInvestorRequest); i {
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
		file_api_investor_investor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMoneyInvestorRequest); i {
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
		file_api_investor_investor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInvestorRequest); i {
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
		file_api_investor_investor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInvestorReply_Investor); i {
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
			RawDescriptor: file_api_investor_investor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_investor_investor_proto_goTypes,
		DependencyIndexes: file_api_investor_investor_proto_depIdxs,
		MessageInfos:      file_api_investor_investor_proto_msgTypes,
	}.Build()
	File_api_investor_investor_proto = out.File
	file_api_investor_investor_proto_rawDesc = nil
	file_api_investor_investor_proto_goTypes = nil
	file_api_investor_investor_proto_depIdxs = nil
}