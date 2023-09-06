// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: monitization/monit.proto

package monitization

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitization_monit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_monitization_monit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_monitization_monit_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitization_monit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_monitization_monit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_monitization_monit_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type ParticipationRewardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID         string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ConferenceID   string `protobuf:"bytes,2,opt,name=ConferenceID,proto3" json:"ConferenceID,omitempty"`
	ConferenceType string `protobuf:"bytes,3,opt,name=ConferenceType,proto3" json:"ConferenceType,omitempty"`
	Minutes        int32  `protobuf:"varint,4,opt,name=Minutes,proto3" json:"Minutes,omitempty"`
}

func (x *ParticipationRewardRequest) Reset() {
	*x = ParticipationRewardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitization_monit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipationRewardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipationRewardRequest) ProtoMessage() {}

func (x *ParticipationRewardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitization_monit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipationRewardRequest.ProtoReflect.Descriptor instead.
func (*ParticipationRewardRequest) Descriptor() ([]byte, []int) {
	return file_monitization_monit_proto_rawDescGZIP(), []int{2}
}

func (x *ParticipationRewardRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ParticipationRewardRequest) GetConferenceID() string {
	if x != nil {
		return x.ConferenceID
	}
	return ""
}

func (x *ParticipationRewardRequest) GetConferenceType() string {
	if x != nil {
		return x.ConferenceType
	}
	return ""
}

func (x *ParticipationRewardRequest) GetMinutes() int32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

type ParticipationRewardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result    string `protobuf:"bytes,1,opt,name=Result,proto3" json:"Result,omitempty"`
	CoinCount int32  `protobuf:"varint,2,opt,name=CoinCount,proto3" json:"CoinCount,omitempty"`
}

func (x *ParticipationRewardResponse) Reset() {
	*x = ParticipationRewardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitization_monit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParticipationRewardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParticipationRewardResponse) ProtoMessage() {}

func (x *ParticipationRewardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitization_monit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParticipationRewardResponse.ProtoReflect.Descriptor instead.
func (*ParticipationRewardResponse) Descriptor() ([]byte, []int) {
	return file_monitization_monit_proto_rawDescGZIP(), []int{3}
}

func (x *ParticipationRewardResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *ParticipationRewardResponse) GetCoinCount() int32 {
	if x != nil {
		return x.CoinCount
	}
	return 0
}

type UserRewardHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Sort   string `protobuf:"bytes,2,opt,name=Sort,proto3" json:"Sort,omitempty"`
}

func (x *UserRewardHistoryRequest) Reset() {
	*x = UserRewardHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitization_monit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRewardHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRewardHistoryRequest) ProtoMessage() {}

func (x *UserRewardHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitization_monit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRewardHistoryRequest.ProtoReflect.Descriptor instead.
func (*UserRewardHistoryRequest) Descriptor() ([]byte, []int) {
	return file_monitization_monit_proto_rawDescGZIP(), []int{4}
}

func (x *UserRewardHistoryRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserRewardHistoryRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type UserRewardHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID          string                 `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	RewardType      string                 `protobuf:"bytes,2,opt,name=reward_type,json=rewardType,proto3" json:"reward_type,omitempty"`
	TransactionType string                 `protobuf:"bytes,3,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"`
	CoinCount       uint32                 `protobuf:"varint,4,opt,name=coin_count,json=coinCount,proto3" json:"coin_count,omitempty"`
	Time            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *UserRewardHistory) Reset() {
	*x = UserRewardHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitization_monit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRewardHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRewardHistory) ProtoMessage() {}

func (x *UserRewardHistory) ProtoReflect() protoreflect.Message {
	mi := &file_monitization_monit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRewardHistory.ProtoReflect.Descriptor instead.
func (*UserRewardHistory) Descriptor() ([]byte, []int) {
	return file_monitization_monit_proto_rawDescGZIP(), []int{5}
}

func (x *UserRewardHistory) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserRewardHistory) GetRewardType() string {
	if x != nil {
		return x.RewardType
	}
	return ""
}

func (x *UserRewardHistory) GetTransactionType() string {
	if x != nil {
		return x.TransactionType
	}
	return ""
}

func (x *UserRewardHistory) GetCoinCount() uint32 {
	if x != nil {
		return x.CoinCount
	}
	return 0
}

func (x *UserRewardHistory) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type UserRewardHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*UserRewardHistory `protobuf:"bytes,1,rep,name=Result,proto3" json:"Result,omitempty"`
}

func (x *UserRewardHistoryResponse) Reset() {
	*x = UserRewardHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitization_monit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRewardHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRewardHistoryResponse) ProtoMessage() {}

func (x *UserRewardHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitization_monit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRewardHistoryResponse.ProtoReflect.Descriptor instead.
func (*UserRewardHistoryResponse) Descriptor() ([]byte, []int) {
	return file_monitization_monit_proto_rawDescGZIP(), []int{6}
}

func (x *UserRewardHistoryResponse) GetResult() []*UserRewardHistory {
	if x != nil {
		return x.Result
	}
	return nil
}

type GroupRewardHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID string `protobuf:"bytes,1,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	Sort    string `protobuf:"bytes,2,opt,name=Sort,proto3" json:"Sort,omitempty"`
}

func (x *GroupRewardHistoryRequest) Reset() {
	*x = GroupRewardHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitization_monit_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRewardHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRewardHistoryRequest) ProtoMessage() {}

func (x *GroupRewardHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_monitization_monit_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRewardHistoryRequest.ProtoReflect.Descriptor instead.
func (*GroupRewardHistoryRequest) Descriptor() ([]byte, []int) {
	return file_monitization_monit_proto_rawDescGZIP(), []int{7}
}

func (x *GroupRewardHistoryRequest) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *GroupRewardHistoryRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

type GroupRewardHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupID         string                 `protobuf:"bytes,1,opt,name=GroupID,proto3" json:"GroupID,omitempty"`
	RewardType      string                 `protobuf:"bytes,2,opt,name=reward_type,json=rewardType,proto3" json:"reward_type,omitempty"`
	TransactionType string                 `protobuf:"bytes,3,opt,name=transaction_type,json=transactionType,proto3" json:"transaction_type,omitempty"`
	CoinCount       uint32                 `protobuf:"varint,4,opt,name=coin_count,json=coinCount,proto3" json:"coin_count,omitempty"`
	Time            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *GroupRewardHistory) Reset() {
	*x = GroupRewardHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitization_monit_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRewardHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRewardHistory) ProtoMessage() {}

func (x *GroupRewardHistory) ProtoReflect() protoreflect.Message {
	mi := &file_monitization_monit_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRewardHistory.ProtoReflect.Descriptor instead.
func (*GroupRewardHistory) Descriptor() ([]byte, []int) {
	return file_monitization_monit_proto_rawDescGZIP(), []int{8}
}

func (x *GroupRewardHistory) GetGroupID() string {
	if x != nil {
		return x.GroupID
	}
	return ""
}

func (x *GroupRewardHistory) GetRewardType() string {
	if x != nil {
		return x.RewardType
	}
	return ""
}

func (x *GroupRewardHistory) GetTransactionType() string {
	if x != nil {
		return x.TransactionType
	}
	return ""
}

func (x *GroupRewardHistory) GetCoinCount() uint32 {
	if x != nil {
		return x.CoinCount
	}
	return 0
}

func (x *GroupRewardHistory) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type GroupRewardHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*GroupRewardHistory `protobuf:"bytes,1,rep,name=Result,proto3" json:"Result,omitempty"`
}

func (x *GroupRewardHistoryResponse) Reset() {
	*x = GroupRewardHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_monitization_monit_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRewardHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRewardHistoryResponse) ProtoMessage() {}

func (x *GroupRewardHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_monitization_monit_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRewardHistoryResponse.ProtoReflect.Descriptor instead.
func (*GroupRewardHistoryResponse) Descriptor() ([]byte, []int) {
	return file_monitization_monit_proto_rawDescGZIP(), []int{9}
}

func (x *GroupRewardHistoryResponse) GetResult() []*GroupRewardHistory {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_monitization_monit_proto protoreflect.FileDescriptor

var file_monitization_monit_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x9a, 0x01, 0x0a,
	0x1a, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x22, 0x53, 0x0a, 0x1b, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46,
	0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22,
	0x54, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x49, 0x0a, 0x19, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x53, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74,
	0x22, 0xc9, 0x01, 0x0a, 0x12, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x44, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x1a,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0x99, 0x03, 0x0a, 0x0c, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x28, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66, 0x0a,
	0x11, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x26, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x12, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x27, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x11, 0x5a, 0x0f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_monitization_monit_proto_rawDescOnce sync.Once
	file_monitization_monit_proto_rawDescData = file_monitization_monit_proto_rawDesc
)

func file_monitization_monit_proto_rawDescGZIP() []byte {
	file_monitization_monit_proto_rawDescOnce.Do(func() {
		file_monitization_monit_proto_rawDescData = protoimpl.X.CompressGZIP(file_monitization_monit_proto_rawDescData)
	})
	return file_monitization_monit_proto_rawDescData
}

var file_monitization_monit_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_monitization_monit_proto_goTypes = []interface{}{
	(*Request)(nil),                     // 0: monitization.Request
	(*Response)(nil),                    // 1: monitization.Response
	(*ParticipationRewardRequest)(nil),  // 2: monitization.ParticipationRewardRequest
	(*ParticipationRewardResponse)(nil), // 3: monitization.ParticipationRewardResponse
	(*UserRewardHistoryRequest)(nil),    // 4: monitization.UserRewardHistoryRequest
	(*UserRewardHistory)(nil),           // 5: monitization.UserRewardHistory
	(*UserRewardHistoryResponse)(nil),   // 6: monitization.UserRewardHistoryResponse
	(*GroupRewardHistoryRequest)(nil),   // 7: monitization.GroupRewardHistoryRequest
	(*GroupRewardHistory)(nil),          // 8: monitization.GroupRewardHistory
	(*GroupRewardHistoryResponse)(nil),  // 9: monitization.GroupRewardHistoryResponse
	(*timestamppb.Timestamp)(nil),       // 10: google.protobuf.Timestamp
}
var file_monitization_monit_proto_depIdxs = []int32{
	10, // 0: monitization.UserRewardHistory.time:type_name -> google.protobuf.Timestamp
	5,  // 1: monitization.UserRewardHistoryResponse.Result:type_name -> monitization.UserRewardHistory
	10, // 2: monitization.GroupRewardHistory.time:type_name -> google.protobuf.Timestamp
	8,  // 3: monitization.GroupRewardHistoryResponse.Result:type_name -> monitization.GroupRewardHistory
	0,  // 4: monitization.Monitization.HealthCheck:input_type -> monitization.Request
	2,  // 5: monitization.Monitization.ConferenceParticipationReward:input_type -> monitization.ParticipationRewardRequest
	4,  // 6: monitization.Monitization.UserRewardHistory:input_type -> monitization.UserRewardHistoryRequest
	7,  // 7: monitization.Monitization.GroupRewardHistory:input_type -> monitization.GroupRewardHistoryRequest
	1,  // 8: monitization.Monitization.HealthCheck:output_type -> monitization.Response
	3,  // 9: monitization.Monitization.ConferenceParticipationReward:output_type -> monitization.ParticipationRewardResponse
	6,  // 10: monitization.Monitization.UserRewardHistory:output_type -> monitization.UserRewardHistoryResponse
	9,  // 11: monitization.Monitization.GroupRewardHistory:output_type -> monitization.GroupRewardHistoryResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_monitization_monit_proto_init() }
func file_monitization_monit_proto_init() {
	if File_monitization_monit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_monitization_monit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_monitization_monit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_monitization_monit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipationRewardRequest); i {
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
		file_monitization_monit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParticipationRewardResponse); i {
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
		file_monitization_monit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRewardHistoryRequest); i {
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
		file_monitization_monit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRewardHistory); i {
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
		file_monitization_monit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRewardHistoryResponse); i {
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
		file_monitization_monit_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRewardHistoryRequest); i {
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
		file_monitization_monit_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRewardHistory); i {
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
		file_monitization_monit_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRewardHistoryResponse); i {
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
			RawDescriptor: file_monitization_monit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_monitization_monit_proto_goTypes,
		DependencyIndexes: file_monitization_monit_proto_depIdxs,
		MessageInfos:      file_monitization_monit_proto_msgTypes,
	}.Build()
	File_monitization_monit_proto = out.File
	file_monitization_monit_proto_rawDesc = nil
	file_monitization_monit_proto_goTypes = nil
	file_monitization_monit_proto_depIdxs = nil
}
