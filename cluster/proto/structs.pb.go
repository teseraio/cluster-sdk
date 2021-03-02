// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: cluster/proto/structs.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StreamObj_Start_Channel int32

const (
	StreamObj_Start_UNKNOWN StreamObj_Start_Channel = 0
	StreamObj_Start_SERF    StreamObj_Start_Channel = 1
	StreamObj_Start_RAFT    StreamObj_Start_Channel = 2
)

// Enum value maps for StreamObj_Start_Channel.
var (
	StreamObj_Start_Channel_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERF",
		2: "RAFT",
	}
	StreamObj_Start_Channel_value = map[string]int32{
		"UNKNOWN": 0,
		"SERF":    1,
		"RAFT":    2,
	}
)

func (x StreamObj_Start_Channel) Enum() *StreamObj_Start_Channel {
	p := new(StreamObj_Start_Channel)
	*p = x
	return p
}

func (x StreamObj_Start_Channel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamObj_Start_Channel) Descriptor() protoreflect.EnumDescriptor {
	return file_cluster_proto_structs_proto_enumTypes[0].Descriptor()
}

func (StreamObj_Start_Channel) Type() protoreflect.EnumType {
	return &file_cluster_proto_structs_proto_enumTypes[0]
}

func (x StreamObj_Start_Channel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamObj_Start_Channel.Descriptor instead.
func (StreamObj_Start_Channel) EnumDescriptor() ([]byte, []int) {
	return file_cluster_proto_structs_proto_rawDescGZIP(), []int{1, 0, 0}
}

type IsLeaderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLeader bool `protobuf:"varint,1,opt,name=isLeader,proto3" json:"isLeader,omitempty"`
}

func (x *IsLeaderResp) Reset() {
	*x = IsLeaderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_structs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsLeaderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsLeaderResp) ProtoMessage() {}

func (x *IsLeaderResp) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_structs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsLeaderResp.ProtoReflect.Descriptor instead.
func (*IsLeaderResp) Descriptor() ([]byte, []int) {
	return file_cluster_proto_structs_proto_rawDescGZIP(), []int{0}
}

func (x *IsLeaderResp) GetIsLeader() bool {
	if x != nil {
		return x.IsLeader
	}
	return false
}

type StreamObj struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*StreamObj_Start_
	//	*StreamObj_Open_
	//	*StreamObj_Input_
	Event isStreamObj_Event `protobuf_oneof:"event"`
}

func (x *StreamObj) Reset() {
	*x = StreamObj{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_structs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamObj) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamObj) ProtoMessage() {}

func (x *StreamObj) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_structs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamObj.ProtoReflect.Descriptor instead.
func (*StreamObj) Descriptor() ([]byte, []int) {
	return file_cluster_proto_structs_proto_rawDescGZIP(), []int{1}
}

func (m *StreamObj) GetEvent() isStreamObj_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *StreamObj) GetStart() *StreamObj_Start {
	if x, ok := x.GetEvent().(*StreamObj_Start_); ok {
		return x.Start
	}
	return nil
}

func (x *StreamObj) GetOpen() *StreamObj_Open {
	if x, ok := x.GetEvent().(*StreamObj_Open_); ok {
		return x.Open
	}
	return nil
}

func (x *StreamObj) GetInput() *StreamObj_Input {
	if x, ok := x.GetEvent().(*StreamObj_Input_); ok {
		return x.Input
	}
	return nil
}

type isStreamObj_Event interface {
	isStreamObj_Event()
}

type StreamObj_Start_ struct {
	Start *StreamObj_Start `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type StreamObj_Open_ struct {
	Open *StreamObj_Open `protobuf:"bytes,2,opt,name=open,proto3,oneof"`
}

type StreamObj_Input_ struct {
	Input *StreamObj_Input `protobuf:"bytes,3,opt,name=input,proto3,oneof"`
}

func (*StreamObj_Start_) isStreamObj_Event() {}

func (*StreamObj_Open_) isStreamObj_Event() {}

func (*StreamObj_Input_) isStreamObj_Event() {}

type StreamObj_Start struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel StreamObj_Start_Channel `protobuf:"varint,1,opt,name=channel,proto3,enum=proto.StreamObj_Start_Channel" json:"channel,omitempty"`
}

func (x *StreamObj_Start) Reset() {
	*x = StreamObj_Start{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_structs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamObj_Start) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamObj_Start) ProtoMessage() {}

func (x *StreamObj_Start) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_structs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamObj_Start.ProtoReflect.Descriptor instead.
func (*StreamObj_Start) Descriptor() ([]byte, []int) {
	return file_cluster_proto_structs_proto_rawDescGZIP(), []int{1, 0}
}

func (x *StreamObj_Start) GetChannel() StreamObj_Start_Channel {
	if x != nil {
		return x.Channel
	}
	return StreamObj_Start_UNKNOWN
}

type StreamObj_Open struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamObj_Open) Reset() {
	*x = StreamObj_Open{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_structs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamObj_Open) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamObj_Open) ProtoMessage() {}

func (x *StreamObj_Open) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_structs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamObj_Open.ProtoReflect.Descriptor instead.
func (*StreamObj_Open) Descriptor() ([]byte, []int) {
	return file_cluster_proto_structs_proto_rawDescGZIP(), []int{1, 1}
}

type StreamObj_Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamObj_Input) Reset() {
	*x = StreamObj_Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_structs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamObj_Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamObj_Input) ProtoMessage() {}

func (x *StreamObj_Input) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_structs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamObj_Input.ProtoReflect.Descriptor instead.
func (*StreamObj_Input) Descriptor() ([]byte, []int) {
	return file_cluster_proto_structs_proto_rawDescGZIP(), []int{1, 2}
}

func (x *StreamObj_Input) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_cluster_proto_structs_proto protoreflect.FileDescriptor

var file_cluster_proto_structs_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2a, 0x0a, 0x0c, 0x49, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0xb5, 0x02,
	0x0a, 0x09, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x62, 0x6a, 0x12, 0x2e, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x62, 0x6a, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6f,
	0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x62, 0x6a, 0x2e, 0x4f, 0x70, 0x65, 0x6e,
	0x48, 0x00, 0x52, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x12, 0x2e, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f, 0x62, 0x6a, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x48,
	0x00, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x6d, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4f, 0x62, 0x6a, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x2a, 0x0a, 0x07, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x52, 0x46, 0x10, 0x01, 0x12, 0x08, 0x0a,
	0x04, 0x52, 0x41, 0x46, 0x54, 0x10, 0x02, 0x1a, 0x06, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x1a,
	0x1b, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x07, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0x7a, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4f,
	0x62, 0x6a, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4f, 0x62, 0x6a, 0x28, 0x01, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x08, 0x49, 0x73, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cluster_proto_structs_proto_rawDescOnce sync.Once
	file_cluster_proto_structs_proto_rawDescData = file_cluster_proto_structs_proto_rawDesc
)

func file_cluster_proto_structs_proto_rawDescGZIP() []byte {
	file_cluster_proto_structs_proto_rawDescOnce.Do(func() {
		file_cluster_proto_structs_proto_rawDescData = protoimpl.X.CompressGZIP(file_cluster_proto_structs_proto_rawDescData)
	})
	return file_cluster_proto_structs_proto_rawDescData
}

var file_cluster_proto_structs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cluster_proto_structs_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cluster_proto_structs_proto_goTypes = []interface{}{
	(StreamObj_Start_Channel)(0), // 0: proto.StreamObj.Start.Channel
	(*IsLeaderResp)(nil),         // 1: proto.IsLeaderResp
	(*StreamObj)(nil),            // 2: proto.StreamObj
	(*StreamObj_Start)(nil),      // 3: proto.StreamObj.Start
	(*StreamObj_Open)(nil),       // 4: proto.StreamObj.Open
	(*StreamObj_Input)(nil),      // 5: proto.StreamObj.Input
	(*empty.Empty)(nil),          // 6: google.protobuf.Empty
}
var file_cluster_proto_structs_proto_depIdxs = []int32{
	3, // 0: proto.StreamObj.start:type_name -> proto.StreamObj.Start
	4, // 1: proto.StreamObj.open:type_name -> proto.StreamObj.Open
	5, // 2: proto.StreamObj.input:type_name -> proto.StreamObj.Input
	0, // 3: proto.StreamObj.Start.channel:type_name -> proto.StreamObj.Start.Channel
	2, // 4: proto.SystemService.Stream:input_type -> proto.StreamObj
	6, // 5: proto.SystemService.IsLeader:input_type -> google.protobuf.Empty
	2, // 6: proto.SystemService.Stream:output_type -> proto.StreamObj
	1, // 7: proto.SystemService.IsLeader:output_type -> proto.IsLeaderResp
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cluster_proto_structs_proto_init() }
func file_cluster_proto_structs_proto_init() {
	if File_cluster_proto_structs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cluster_proto_structs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsLeaderResp); i {
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
		file_cluster_proto_structs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamObj); i {
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
		file_cluster_proto_structs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamObj_Start); i {
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
		file_cluster_proto_structs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamObj_Open); i {
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
		file_cluster_proto_structs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamObj_Input); i {
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
	file_cluster_proto_structs_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StreamObj_Start_)(nil),
		(*StreamObj_Open_)(nil),
		(*StreamObj_Input_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cluster_proto_structs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cluster_proto_structs_proto_goTypes,
		DependencyIndexes: file_cluster_proto_structs_proto_depIdxs,
		EnumInfos:         file_cluster_proto_structs_proto_enumTypes,
		MessageInfos:      file_cluster_proto_structs_proto_msgTypes,
	}.Build()
	File_cluster_proto_structs_proto = out.File
	file_cluster_proto_structs_proto_rawDesc = nil
	file_cluster_proto_structs_proto_goTypes = nil
	file_cluster_proto_structs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SystemServiceClient is the client API for SystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SystemServiceClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (SystemService_StreamClient, error)
	IsLeader(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IsLeaderResp, error)
}

type systemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemServiceClient(cc grpc.ClientConnInterface) SystemServiceClient {
	return &systemServiceClient{cc}
}

func (c *systemServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (SystemService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SystemService_serviceDesc.Streams[0], "/proto.SystemService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &systemServiceStreamClient{stream}
	return x, nil
}

type SystemService_StreamClient interface {
	Send(*StreamObj) error
	Recv() (*StreamObj, error)
	grpc.ClientStream
}

type systemServiceStreamClient struct {
	grpc.ClientStream
}

func (x *systemServiceStreamClient) Send(m *StreamObj) error {
	return x.ClientStream.SendMsg(m)
}

func (x *systemServiceStreamClient) Recv() (*StreamObj, error) {
	m := new(StreamObj)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *systemServiceClient) IsLeader(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IsLeaderResp, error) {
	out := new(IsLeaderResp)
	err := c.cc.Invoke(ctx, "/proto.SystemService/IsLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServiceServer is the server API for SystemService service.
type SystemServiceServer interface {
	Stream(SystemService_StreamServer) error
	IsLeader(context.Context, *empty.Empty) (*IsLeaderResp, error)
}

// UnimplementedSystemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSystemServiceServer struct {
}

func (*UnimplementedSystemServiceServer) Stream(SystemService_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (*UnimplementedSystemServiceServer) IsLeader(context.Context, *empty.Empty) (*IsLeaderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLeader not implemented")
}

func RegisterSystemServiceServer(s *grpc.Server, srv SystemServiceServer) {
	s.RegisterService(&_SystemService_serviceDesc, srv)
}

func _SystemService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SystemServiceServer).Stream(&systemServiceStreamServer{stream})
}

type SystemService_StreamServer interface {
	Send(*StreamObj) error
	Recv() (*StreamObj, error)
	grpc.ServerStream
}

type systemServiceStreamServer struct {
	grpc.ServerStream
}

func (x *systemServiceStreamServer) Send(m *StreamObj) error {
	return x.ServerStream.SendMsg(m)
}

func (x *systemServiceStreamServer) Recv() (*StreamObj, error) {
	m := new(StreamObj)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SystemService_IsLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).IsLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SystemService/IsLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).IsLeader(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SystemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SystemService",
	HandlerType: (*SystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsLeader",
			Handler:    _SystemService_IsLeader_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _SystemService_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cluster/proto/structs.proto",
}
