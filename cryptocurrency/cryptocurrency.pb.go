// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: cryptocurrency/cryptocurrency.proto

package cryptocurrency

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

type CryptocurrencyListMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cryptocurrency []*CryptocurrencyMessage `protobuf:"bytes,1,rep,name=cryptocurrency,proto3" json:"cryptocurrency,omitempty"`
}

func (x *CryptocurrencyListMessage) Reset() {
	*x = CryptocurrencyListMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptocurrencyListMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptocurrencyListMessage) ProtoMessage() {}

func (x *CryptocurrencyListMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptocurrencyListMessage.ProtoReflect.Descriptor instead.
func (*CryptocurrencyListMessage) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_cryptocurrency_proto_rawDescGZIP(), []int{0}
}

func (x *CryptocurrencyListMessage) GetCryptocurrency() []*CryptocurrencyMessage {
	if x != nil {
		return x.Cryptocurrency
	}
	return nil
}

type CryptocurrencyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol  string `protobuf:"bytes,1,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	IconURL string `protobuf:"bytes,3,opt,name=IconURL,proto3" json:"IconURL,omitempty"`
	Votes   int32  `protobuf:"varint,4,opt,name=votes,proto3" json:"votes,omitempty"`
}

func (x *CryptocurrencyMessage) Reset() {
	*x = CryptocurrencyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptocurrencyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptocurrencyMessage) ProtoMessage() {}

func (x *CryptocurrencyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptocurrencyMessage.ProtoReflect.Descriptor instead.
func (*CryptocurrencyMessage) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_cryptocurrency_proto_rawDescGZIP(), []int{1}
}

func (x *CryptocurrencyMessage) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *CryptocurrencyMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CryptocurrencyMessage) GetIconURL() string {
	if x != nil {
		return x.IconURL
	}
	return ""
}

func (x *CryptocurrencyMessage) GetVotes() int32 {
	if x != nil {
		return x.Votes
	}
	return 0
}

type CryptocurrencySymbol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol string `protobuf:"bytes,1,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
}

func (x *CryptocurrencySymbol) Reset() {
	*x = CryptocurrencySymbol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptocurrencySymbol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptocurrencySymbol) ProtoMessage() {}

func (x *CryptocurrencySymbol) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptocurrencySymbol.ProtoReflect.Descriptor instead.
func (*CryptocurrencySymbol) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_cryptocurrency_proto_rawDescGZIP(), []int{2}
}

func (x *CryptocurrencySymbol) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type UpdateCryptocurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldSymbol         string                 `protobuf:"bytes,1,opt,name=OldSymbol,proto3" json:"OldSymbol,omitempty"`
	NewCryptocurrency *CryptocurrencyMessage `protobuf:"bytes,2,opt,name=newCryptocurrency,proto3" json:"newCryptocurrency,omitempty"`
}

func (x *UpdateCryptocurrencyRequest) Reset() {
	*x = UpdateCryptocurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCryptocurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCryptocurrencyRequest) ProtoMessage() {}

func (x *UpdateCryptocurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCryptocurrencyRequest.ProtoReflect.Descriptor instead.
func (*UpdateCryptocurrencyRequest) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_cryptocurrency_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateCryptocurrencyRequest) GetOldSymbol() string {
	if x != nil {
		return x.OldSymbol
	}
	return ""
}

func (x *UpdateCryptocurrencyRequest) GetNewCryptocurrency() *CryptocurrencyMessage {
	if x != nil {
		return x.NewCryptocurrency
	}
	return nil
}

type CryptocurrencyNewVoteNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Votes int32 `protobuf:"varint,1,opt,name=votes,proto3" json:"votes,omitempty"`
}

func (x *CryptocurrencyNewVoteNotification) Reset() {
	*x = CryptocurrencyNewVoteNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptocurrencyNewVoteNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptocurrencyNewVoteNotification) ProtoMessage() {}

func (x *CryptocurrencyNewVoteNotification) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptocurrencyNewVoteNotification.ProtoReflect.Descriptor instead.
func (*CryptocurrencyNewVoteNotification) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_cryptocurrency_proto_rawDescGZIP(), []int{4}
}

func (x *CryptocurrencyNewVoteNotification) GetVotes() int32 {
	if x != nil {
		return x.Votes
	}
	return 0
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_cryptocurrency_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_cryptocurrency_proto_rawDescGZIP(), []int{5}
}

var File_cryptocurrency_cryptocurrency_proto protoreflect.FileDescriptor

var file_cryptocurrency_cryptocurrency_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x6a, 0x0a, 0x19, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x4d, 0x0a, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x22, 0x73, 0x0a, 0x15, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x63, 0x6f, 0x6e, 0x55, 0x52,
	0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x63, 0x6f, 0x6e, 0x55, 0x52, 0x4c,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x14, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x90, 0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x6c, 0x64, 0x53, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4f, 0x6c, 0x64, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x53, 0x0a, 0x11, 0x6e, 0x65, 0x77, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x11, 0x6e, 0x65, 0x77, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x39, 0x0a, 0x21, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e, 0x65, 0x77, 0x56, 0x6f,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x6f, 0x74, 0x65, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xef, 0x05, 0x0a, 0x15, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x25, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x24, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x1a, 0x25, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x51,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x29, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x5e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x4e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x1a, 0x1c, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x12, 0x57, 0x0a, 0x06, 0x55, 0x70, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x1a, 0x25, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x08, 0x44, 0x6f,
	0x77, 0x6e, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x1a, 0x25, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x6f, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x24, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x1a,
	0x31, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4e,
	0x65, 0x77, 0x56, 0x6f, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x42, 0x17, 0x5a, 0x15, 0x6b, 0x6c, 0x65, 0x76, 0x65, 0x72,
	0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cryptocurrency_cryptocurrency_proto_rawDescOnce sync.Once
	file_cryptocurrency_cryptocurrency_proto_rawDescData = file_cryptocurrency_cryptocurrency_proto_rawDesc
)

func file_cryptocurrency_cryptocurrency_proto_rawDescGZIP() []byte {
	file_cryptocurrency_cryptocurrency_proto_rawDescOnce.Do(func() {
		file_cryptocurrency_cryptocurrency_proto_rawDescData = protoimpl.X.CompressGZIP(file_cryptocurrency_cryptocurrency_proto_rawDescData)
	})
	return file_cryptocurrency_cryptocurrency_proto_rawDescData
}

var file_cryptocurrency_cryptocurrency_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cryptocurrency_cryptocurrency_proto_goTypes = []interface{}{
	(*CryptocurrencyListMessage)(nil),         // 0: cryptocurrency.CryptocurrencyListMessage
	(*CryptocurrencyMessage)(nil),             // 1: cryptocurrency.CryptocurrencyMessage
	(*CryptocurrencySymbol)(nil),              // 2: cryptocurrency.CryptocurrencySymbol
	(*UpdateCryptocurrencyRequest)(nil),       // 3: cryptocurrency.UpdateCryptocurrencyRequest
	(*CryptocurrencyNewVoteNotification)(nil), // 4: cryptocurrency.CryptocurrencyNewVoteNotification
	(*EmptyMessage)(nil),                      // 5: cryptocurrency.EmptyMessage
}
var file_cryptocurrency_cryptocurrency_proto_depIdxs = []int32{
	1,  // 0: cryptocurrency.CryptocurrencyListMessage.cryptocurrency:type_name -> cryptocurrency.CryptocurrencyMessage
	1,  // 1: cryptocurrency.UpdateCryptocurrencyRequest.newCryptocurrency:type_name -> cryptocurrency.CryptocurrencyMessage
	1,  // 2: cryptocurrency.CryptocurrencyService.Create:input_type -> cryptocurrency.CryptocurrencyMessage
	2,  // 3: cryptocurrency.CryptocurrencyService.Get:input_type -> cryptocurrency.CryptocurrencySymbol
	5,  // 4: cryptocurrency.CryptocurrencyService.List:input_type -> cryptocurrency.EmptyMessage
	3,  // 5: cryptocurrency.CryptocurrencyService.Update:input_type -> cryptocurrency.UpdateCryptocurrencyRequest
	2,  // 6: cryptocurrency.CryptocurrencyService.Delete:input_type -> cryptocurrency.CryptocurrencySymbol
	2,  // 7: cryptocurrency.CryptocurrencyService.UpVote:input_type -> cryptocurrency.CryptocurrencySymbol
	2,  // 8: cryptocurrency.CryptocurrencyService.DownVote:input_type -> cryptocurrency.CryptocurrencySymbol
	2,  // 9: cryptocurrency.CryptocurrencyService.CreateVoteStream:input_type -> cryptocurrency.CryptocurrencySymbol
	1,  // 10: cryptocurrency.CryptocurrencyService.Create:output_type -> cryptocurrency.CryptocurrencyMessage
	1,  // 11: cryptocurrency.CryptocurrencyService.Get:output_type -> cryptocurrency.CryptocurrencyMessage
	0,  // 12: cryptocurrency.CryptocurrencyService.List:output_type -> cryptocurrency.CryptocurrencyListMessage
	1,  // 13: cryptocurrency.CryptocurrencyService.Update:output_type -> cryptocurrency.CryptocurrencyMessage
	5,  // 14: cryptocurrency.CryptocurrencyService.Delete:output_type -> cryptocurrency.EmptyMessage
	1,  // 15: cryptocurrency.CryptocurrencyService.UpVote:output_type -> cryptocurrency.CryptocurrencyMessage
	1,  // 16: cryptocurrency.CryptocurrencyService.DownVote:output_type -> cryptocurrency.CryptocurrencyMessage
	4,  // 17: cryptocurrency.CryptocurrencyService.CreateVoteStream:output_type -> cryptocurrency.CryptocurrencyNewVoteNotification
	10, // [10:18] is the sub-list for method output_type
	2,  // [2:10] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_cryptocurrency_cryptocurrency_proto_init() }
func file_cryptocurrency_cryptocurrency_proto_init() {
	if File_cryptocurrency_cryptocurrency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cryptocurrency_cryptocurrency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptocurrencyListMessage); i {
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
		file_cryptocurrency_cryptocurrency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptocurrencyMessage); i {
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
		file_cryptocurrency_cryptocurrency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptocurrencySymbol); i {
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
		file_cryptocurrency_cryptocurrency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCryptocurrencyRequest); i {
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
		file_cryptocurrency_cryptocurrency_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptocurrencyNewVoteNotification); i {
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
		file_cryptocurrency_cryptocurrency_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
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
			RawDescriptor: file_cryptocurrency_cryptocurrency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cryptocurrency_cryptocurrency_proto_goTypes,
		DependencyIndexes: file_cryptocurrency_cryptocurrency_proto_depIdxs,
		MessageInfos:      file_cryptocurrency_cryptocurrency_proto_msgTypes,
	}.Build()
	File_cryptocurrency_cryptocurrency_proto = out.File
	file_cryptocurrency_cryptocurrency_proto_rawDesc = nil
	file_cryptocurrency_cryptocurrency_proto_goTypes = nil
	file_cryptocurrency_cryptocurrency_proto_depIdxs = nil
}