// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        v3.19.4
// source: service.proto

package proto

import (
	context "context"
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

type AddEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AddEntryRequest) Reset() {
	*x = AddEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEntryRequest) ProtoMessage() {}

func (x *AddEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEntryRequest.ProtoReflect.Descriptor instead.
func (*AddEntryRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddEntryRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *AddEntryRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetValueRequest) Reset() {
	*x = SetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetValueRequest) ProtoMessage() {}

func (x *SetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetValueRequest.ProtoReflect.Descriptor instead.
func (*SetValueRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *SetValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetValueRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetValueRequest) Reset() {
	*x = GetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueRequest) ProtoMessage() {}

func (x *GetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueRequest.ProtoReflect.Descriptor instead.
func (*GetValueRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteKeyRequest) Reset() {
	*x = DeleteKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKeyRequest) ProtoMessage() {}

func (x *DeleteKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteKeyRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type KeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyResponse) Reset() {
	*x = KeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyResponse) ProtoMessage() {}

func (x *KeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyResponse.ProtoReflect.Descriptor instead.
func (*KeyResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *KeyResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x39, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x39, 0x0a, 0x0f, 0x53, 0x65,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x24, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x35, 0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xd0, 0x01, 0x0a, 0x0f, 0x4b, 0x65, 0x79, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x41,
	0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x53, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x68, 0x61, 0x63, 0x68, 0x74,
	0x65, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_proto_goTypes = []interface{}{
	(*AddEntryRequest)(nil),  // 0: AddEntryRequest
	(*SetValueRequest)(nil),  // 1: SetValueRequest
	(*GetValueRequest)(nil),  // 2: GetValueRequest
	(*DeleteKeyRequest)(nil), // 3: DeleteKeyRequest
	(*KeyResponse)(nil),      // 4: KeyResponse
}
var file_service_proto_depIdxs = []int32{
	0, // 0: KeyValueService.AddEntry:input_type -> AddEntryRequest
	1, // 1: KeyValueService.UpdateEntry:input_type -> SetValueRequest
	3, // 2: KeyValueService.DeleteEntry:input_type -> DeleteKeyRequest
	2, // 3: KeyValueService.GetEntry:input_type -> GetValueRequest
	4, // 4: KeyValueService.AddEntry:output_type -> KeyResponse
	4, // 5: KeyValueService.UpdateEntry:output_type -> KeyResponse
	4, // 6: KeyValueService.DeleteEntry:output_type -> KeyResponse
	4, // 7: KeyValueService.GetEntry:output_type -> KeyResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEntryRequest); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetValueRequest); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueRequest); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKeyRequest); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyResponse); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeyValueServiceClient is the client API for KeyValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyValueServiceClient interface {
	AddEntry(ctx context.Context, in *AddEntryRequest, opts ...grpc.CallOption) (*KeyResponse, error)
	UpdateEntry(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*KeyResponse, error)
	DeleteEntry(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*KeyResponse, error)
	GetEntry(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*KeyResponse, error)
}

type keyValueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyValueServiceClient(cc grpc.ClientConnInterface) KeyValueServiceClient {
	return &keyValueServiceClient{cc}
}

func (c *keyValueServiceClient) AddEntry(ctx context.Context, in *AddEntryRequest, opts ...grpc.CallOption) (*KeyResponse, error) {
	out := new(KeyResponse)
	err := c.cc.Invoke(ctx, "/KeyValueService/AddEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) UpdateEntry(ctx context.Context, in *SetValueRequest, opts ...grpc.CallOption) (*KeyResponse, error) {
	out := new(KeyResponse)
	err := c.cc.Invoke(ctx, "/KeyValueService/UpdateEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) DeleteEntry(ctx context.Context, in *DeleteKeyRequest, opts ...grpc.CallOption) (*KeyResponse, error) {
	out := new(KeyResponse)
	err := c.cc.Invoke(ctx, "/KeyValueService/DeleteEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueServiceClient) GetEntry(ctx context.Context, in *GetValueRequest, opts ...grpc.CallOption) (*KeyResponse, error) {
	out := new(KeyResponse)
	err := c.cc.Invoke(ctx, "/KeyValueService/GetEntry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueServiceServer is the server API for KeyValueService service.
type KeyValueServiceServer interface {
	AddEntry(context.Context, *AddEntryRequest) (*KeyResponse, error)
	UpdateEntry(context.Context, *SetValueRequest) (*KeyResponse, error)
	DeleteEntry(context.Context, *DeleteKeyRequest) (*KeyResponse, error)
	GetEntry(context.Context, *GetValueRequest) (*KeyResponse, error)
}

// UnimplementedKeyValueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeyValueServiceServer struct {
}

func (*UnimplementedKeyValueServiceServer) AddEntry(context.Context, *AddEntryRequest) (*KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEntry not implemented")
}
func (*UnimplementedKeyValueServiceServer) UpdateEntry(context.Context, *SetValueRequest) (*KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEntry not implemented")
}
func (*UnimplementedKeyValueServiceServer) DeleteEntry(context.Context, *DeleteKeyRequest) (*KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEntry not implemented")
}
func (*UnimplementedKeyValueServiceServer) GetEntry(context.Context, *GetValueRequest) (*KeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEntry not implemented")
}

func RegisterKeyValueServiceServer(s *grpc.Server, srv KeyValueServiceServer) {
	s.RegisterService(&_KeyValueService_serviceDesc, srv)
}

func _KeyValueService_AddEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).AddEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeyValueService/AddEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).AddEntry(ctx, req.(*AddEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_UpdateEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).UpdateEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeyValueService/UpdateEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).UpdateEntry(ctx, req.(*SetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_DeleteEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).DeleteEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeyValueService/DeleteEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).DeleteEntry(ctx, req.(*DeleteKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueService_GetEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServiceServer).GetEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KeyValueService/GetEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServiceServer).GetEntry(ctx, req.(*GetValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyValueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "KeyValueService",
	HandlerType: (*KeyValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEntry",
			Handler:    _KeyValueService_AddEntry_Handler,
		},
		{
			MethodName: "UpdateEntry",
			Handler:    _KeyValueService_UpdateEntry_Handler,
		},
		{
			MethodName: "DeleteEntry",
			Handler:    _KeyValueService_DeleteEntry_Handler,
		},
		{
			MethodName: "GetEntry",
			Handler:    _KeyValueService_GetEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
