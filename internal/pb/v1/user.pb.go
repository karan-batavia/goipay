// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.2
// source: user.proto

package v1

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

type RegisterUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId *string `protobuf:"bytes,1,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
}

func (x *RegisterUserRequest) Reset() {
	*x = RegisterUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRequest) ProtoMessage() {}

func (x *RegisterUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterUserRequest) GetUserId() string {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return ""
}

type RegisterUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RegisterUserResponse) Reset() {
	*x = RegisterUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserResponse) ProtoMessage() {}

func (x *RegisterUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateCryptoKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string                `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XmrReq *XmrKeysUpdateRequest `protobuf:"bytes,2,opt,name=xmrReq,proto3,oneof" json:"xmrReq,omitempty"`
	BtcReq *BtcKeysUpdateRequest `protobuf:"bytes,3,opt,name=btcReq,proto3,oneof" json:"btcReq,omitempty"`
	LtcReq *LtcKeysUpdateRequest `protobuf:"bytes,4,opt,name=ltcReq,proto3,oneof" json:"ltcReq,omitempty"`
	EthReq *EthKeysUpdateRequest `protobuf:"bytes,5,opt,name=ethReq,proto3,oneof" json:"ethReq,omitempty"`
}

func (x *UpdateCryptoKeysRequest) Reset() {
	*x = UpdateCryptoKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCryptoKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCryptoKeysRequest) ProtoMessage() {}

func (x *UpdateCryptoKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCryptoKeysRequest.ProtoReflect.Descriptor instead.
func (*UpdateCryptoKeysRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCryptoKeysRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateCryptoKeysRequest) GetXmrReq() *XmrKeysUpdateRequest {
	if x != nil {
		return x.XmrReq
	}
	return nil
}

func (x *UpdateCryptoKeysRequest) GetBtcReq() *BtcKeysUpdateRequest {
	if x != nil {
		return x.BtcReq
	}
	return nil
}

func (x *UpdateCryptoKeysRequest) GetLtcReq() *LtcKeysUpdateRequest {
	if x != nil {
		return x.LtcReq
	}
	return nil
}

func (x *UpdateCryptoKeysRequest) GetEthReq() *EthKeysUpdateRequest {
	if x != nil {
		return x.EthReq
	}
	return nil
}

type UpdateCryptoKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCryptoKeysResponse) Reset() {
	*x = UpdateCryptoKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCryptoKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCryptoKeysResponse) ProtoMessage() {}

func (x *UpdateCryptoKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCryptoKeysResponse.ProtoReflect.Descriptor instead.
func (*UpdateCryptoKeysResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x2e, 0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0xd5, 0x02, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x78, 0x6d, 0x72, 0x52, 0x65, 0x71,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x58, 0x6d, 0x72, 0x4b, 0x65, 0x79, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x78, 0x6d, 0x72, 0x52, 0x65,
	0x71, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x06, 0x62, 0x74, 0x63, 0x52, 0x65, 0x71, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x74, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x01, 0x52, 0x06, 0x62, 0x74, 0x63, 0x52, 0x65, 0x71, 0x88,
	0x01, 0x01, 0x12, 0x3c, 0x0a, 0x06, 0x6c, 0x74, 0x63, 0x52, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x74, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x48, 0x02, 0x52, 0x06, 0x6c, 0x74, 0x63, 0x52, 0x65, 0x71, 0x88, 0x01, 0x01,
	0x12, 0x3c, 0x0a, 0x06, 0x65, 0x74, 0x68, 0x52, 0x65, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x74, 0x68,
	0x4b, 0x65, 0x79, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x03, 0x52, 0x06, 0x65, 0x74, 0x68, 0x52, 0x65, 0x71, 0x88, 0x01, 0x01, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x78, 0x6d, 0x72, 0x52, 0x65, 0x71, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62, 0x74,
	0x63, 0x52, 0x65, 0x71, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6c, 0x74, 0x63, 0x52, 0x65, 0x71, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x65, 0x74, 0x68, 0x52, 0x65, 0x71, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xb3, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65,
	0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_proto_goTypes = []any{
	(*RegisterUserRequest)(nil),      // 0: user.v1.RegisterUserRequest
	(*RegisterUserResponse)(nil),     // 1: user.v1.RegisterUserResponse
	(*UpdateCryptoKeysRequest)(nil),  // 2: user.v1.UpdateCryptoKeysRequest
	(*UpdateCryptoKeysResponse)(nil), // 3: user.v1.UpdateCryptoKeysResponse
	(*XmrKeysUpdateRequest)(nil),     // 4: crypto.v1.XmrKeysUpdateRequest
	(*BtcKeysUpdateRequest)(nil),     // 5: crypto.v1.BtcKeysUpdateRequest
	(*LtcKeysUpdateRequest)(nil),     // 6: crypto.v1.LtcKeysUpdateRequest
	(*EthKeysUpdateRequest)(nil),     // 7: crypto.v1.EthKeysUpdateRequest
}
var file_user_proto_depIdxs = []int32{
	4, // 0: user.v1.UpdateCryptoKeysRequest.xmrReq:type_name -> crypto.v1.XmrKeysUpdateRequest
	5, // 1: user.v1.UpdateCryptoKeysRequest.btcReq:type_name -> crypto.v1.BtcKeysUpdateRequest
	6, // 2: user.v1.UpdateCryptoKeysRequest.ltcReq:type_name -> crypto.v1.LtcKeysUpdateRequest
	7, // 3: user.v1.UpdateCryptoKeysRequest.ethReq:type_name -> crypto.v1.EthKeysUpdateRequest
	0, // 4: user.v1.UserService.RegisterUser:input_type -> user.v1.RegisterUserRequest
	2, // 5: user.v1.UserService.UpdateCryptoKeys:input_type -> user.v1.UpdateCryptoKeysRequest
	1, // 6: user.v1.UserService.RegisterUser:output_type -> user.v1.RegisterUserResponse
	3, // 7: user.v1.UserService.UpdateCryptoKeys:output_type -> user.v1.UpdateCryptoKeysResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_crypto_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterUserRequest); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterUserResponse); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCryptoKeysRequest); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateCryptoKeysResponse); i {
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
	file_user_proto_msgTypes[0].OneofWrappers = []any{}
	file_user_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
