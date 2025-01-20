// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.2
// source: crypto.proto

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

type CoinType int32

const (
	CoinType_XMR CoinType = 0
	CoinType_BTC CoinType = 1
	CoinType_LTC CoinType = 2
	CoinType_ETH CoinType = 3
	CoinType_TON CoinType = 4
)

// Enum value maps for CoinType.
var (
	CoinType_name = map[int32]string{
		0: "XMR",
		1: "BTC",
		2: "LTC",
		3: "ETH",
		4: "TON",
	}
	CoinType_value = map[string]int32{
		"XMR": 0,
		"BTC": 1,
		"LTC": 2,
		"ETH": 3,
		"TON": 4,
	}
)

func (x CoinType) Enum() *CoinType {
	p := new(CoinType)
	*p = x
	return p
}

func (x CoinType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoinType) Descriptor() protoreflect.EnumDescriptor {
	return file_crypto_proto_enumTypes[0].Descriptor()
}

func (CoinType) Type() protoreflect.EnumType {
	return &file_crypto_proto_enumTypes[0]
}

func (x CoinType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoinType.Descriptor instead.
func (CoinType) EnumDescriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{0}
}

type XmrKeysUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivViewKey string `protobuf:"bytes,1,opt,name=privViewKey,proto3" json:"privViewKey,omitempty"`
	PubSpendKey string `protobuf:"bytes,2,opt,name=pubSpendKey,proto3" json:"pubSpendKey,omitempty"`
}

func (x *XmrKeysUpdateRequest) Reset() {
	*x = XmrKeysUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XmrKeysUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XmrKeysUpdateRequest) ProtoMessage() {}

func (x *XmrKeysUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XmrKeysUpdateRequest.ProtoReflect.Descriptor instead.
func (*XmrKeysUpdateRequest) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *XmrKeysUpdateRequest) GetPrivViewKey() string {
	if x != nil {
		return x.PrivViewKey
	}
	return ""
}

func (x *XmrKeysUpdateRequest) GetPubSpendKey() string {
	if x != nil {
		return x.PubSpendKey
	}
	return ""
}

type BtcKeysUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterPubKey string `protobuf:"bytes,1,opt,name=masterPubKey,proto3" json:"masterPubKey,omitempty"`
}

func (x *BtcKeysUpdateRequest) Reset() {
	*x = BtcKeysUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BtcKeysUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BtcKeysUpdateRequest) ProtoMessage() {}

func (x *BtcKeysUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BtcKeysUpdateRequest.ProtoReflect.Descriptor instead.
func (*BtcKeysUpdateRequest) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *BtcKeysUpdateRequest) GetMasterPubKey() string {
	if x != nil {
		return x.MasterPubKey
	}
	return ""
}

type LtcKeysUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterPubKey string `protobuf:"bytes,1,opt,name=masterPubKey,proto3" json:"masterPubKey,omitempty"`
}

func (x *LtcKeysUpdateRequest) Reset() {
	*x = LtcKeysUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LtcKeysUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LtcKeysUpdateRequest) ProtoMessage() {}

func (x *LtcKeysUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LtcKeysUpdateRequest.ProtoReflect.Descriptor instead.
func (*LtcKeysUpdateRequest) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{2}
}

func (x *LtcKeysUpdateRequest) GetMasterPubKey() string {
	if x != nil {
		return x.MasterPubKey
	}
	return ""
}

type EthKeysUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MasterPubKey string `protobuf:"bytes,1,opt,name=masterPubKey,proto3" json:"masterPubKey,omitempty"`
}

func (x *EthKeysUpdateRequest) Reset() {
	*x = EthKeysUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crypto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthKeysUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthKeysUpdateRequest) ProtoMessage() {}

func (x *EthKeysUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_crypto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthKeysUpdateRequest.ProtoReflect.Descriptor instead.
func (*EthKeysUpdateRequest) Descriptor() ([]byte, []int) {
	return file_crypto_proto_rawDescGZIP(), []int{3}
}

func (x *EthKeysUpdateRequest) GetMasterPubKey() string {
	if x != nil {
		return x.MasterPubKey
	}
	return ""
}

var File_crypto_proto protoreflect.FileDescriptor

var file_crypto_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x5a, 0x0a, 0x14, 0x58, 0x6d, 0x72,
	0x4b, 0x65, 0x79, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x56, 0x69, 0x65, 0x77, 0x4b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x56, 0x69, 0x65, 0x77,
	0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x75, 0x62, 0x53, 0x70, 0x65, 0x6e, 0x64, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x53, 0x70, 0x65,
	0x6e, 0x64, 0x4b, 0x65, 0x79, 0x22, 0x3a, 0x0a, 0x14, 0x42, 0x74, 0x63, 0x4b, 0x65, 0x79, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4b, 0x65,
	0x79, 0x22, 0x3a, 0x0a, 0x14, 0x4c, 0x74, 0x63, 0x4b, 0x65, 0x79, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x22, 0x3a, 0x0a,
	0x14, 0x45, 0x74, 0x68, 0x4b, 0x65, 0x79, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x2a, 0x37, 0x0a, 0x08, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x58, 0x4d, 0x52, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x42, 0x54, 0x43, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x54, 0x43, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x45, 0x54, 0x48, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x4f, 0x4e,
	0x10, 0x04, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crypto_proto_rawDescOnce sync.Once
	file_crypto_proto_rawDescData = file_crypto_proto_rawDesc
)

func file_crypto_proto_rawDescGZIP() []byte {
	file_crypto_proto_rawDescOnce.Do(func() {
		file_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_crypto_proto_rawDescData)
	})
	return file_crypto_proto_rawDescData
}

var file_crypto_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_crypto_proto_goTypes = []any{
	(CoinType)(0),                // 0: crypto.v1.CoinType
	(*XmrKeysUpdateRequest)(nil), // 1: crypto.v1.XmrKeysUpdateRequest
	(*BtcKeysUpdateRequest)(nil), // 2: crypto.v1.BtcKeysUpdateRequest
	(*LtcKeysUpdateRequest)(nil), // 3: crypto.v1.LtcKeysUpdateRequest
	(*EthKeysUpdateRequest)(nil), // 4: crypto.v1.EthKeysUpdateRequest
}
var file_crypto_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crypto_proto_init() }
func file_crypto_proto_init() {
	if File_crypto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crypto_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*XmrKeysUpdateRequest); i {
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
		file_crypto_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BtcKeysUpdateRequest); i {
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
		file_crypto_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LtcKeysUpdateRequest); i {
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
		file_crypto_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EthKeysUpdateRequest); i {
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
			RawDescriptor: file_crypto_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crypto_proto_goTypes,
		DependencyIndexes: file_crypto_proto_depIdxs,
		EnumInfos:         file_crypto_proto_enumTypes,
		MessageInfos:      file_crypto_proto_msgTypes,
	}.Build()
	File_crypto_proto = out.File
	file_crypto_proto_rawDesc = nil
	file_crypto_proto_goTypes = nil
	file_crypto_proto_depIdxs = nil
}
