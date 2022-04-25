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

//**************************
//REQUESTS
//*************************
type AddReminderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Completed   bool   `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
}

func (x *AddReminderRequest) Reset() {
	*x = AddReminderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReminderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReminderRequest) ProtoMessage() {}

func (x *AddReminderRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddReminderRequest.ProtoReflect.Descriptor instead.
func (*AddReminderRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddReminderRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AddReminderRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddReminderRequest) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type GetLatestReminderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLatestReminderRequest) Reset() {
	*x = GetLatestReminderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestReminderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestReminderRequest) ProtoMessage() {}

func (x *GetLatestReminderRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetLatestReminderRequest.ProtoReflect.Descriptor instead.
func (*GetLatestReminderRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

type GetAllRemindersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRemindersRequest) Reset() {
	*x = GetAllRemindersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRemindersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRemindersRequest) ProtoMessage() {}

func (x *GetAllRemindersRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetAllRemindersRequest.ProtoReflect.Descriptor instead.
func (*GetAllRemindersRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

//**************************
//RESPONSES
//*************************
type AddReminderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitIndex uint64 `protobuf:"varint,1,opt,name=commit_index,json=commitIndex,proto3" json:"commit_index,omitempty"`
}

func (x *AddReminderResponse) Reset() {
	*x = AddReminderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReminderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReminderResponse) ProtoMessage() {}

func (x *AddReminderResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddReminderResponse.ProtoReflect.Descriptor instead.
func (*AddReminderResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *AddReminderResponse) GetCommitIndex() uint64 {
	if x != nil {
		return x.CommitIndex
	}
	return 0
}

type GetLatestReminderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReadAtIndex uint64 `protobuf:"varint,1,opt,name=read_at_index,json=readAtIndex,proto3" json:"read_at_index,omitempty"`
	Reminder    string `protobuf:"bytes,2,opt,name=reminder,proto3" json:"reminder,omitempty"`
}

func (x *GetLatestReminderResponse) Reset() {
	*x = GetLatestReminderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestReminderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestReminderResponse) ProtoMessage() {}

func (x *GetLatestReminderResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetLatestReminderResponse.ProtoReflect.Descriptor instead.
func (*GetLatestReminderResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetLatestReminderResponse) GetReadAtIndex() uint64 {
	if x != nil {
		return x.ReadAtIndex
	}
	return 0
}

func (x *GetLatestReminderResponse) GetReminder() string {
	if x != nil {
		return x.Reminder
	}
	return ""
}

type GetAllRemindersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reminders []string `protobuf:"bytes,1,rep,name=reminders,proto3" json:"reminders,omitempty"`
}

func (x *GetAllRemindersResponse) Reset() {
	*x = GetAllRemindersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRemindersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRemindersResponse) ProtoMessage() {}

func (x *GetAllRemindersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRemindersResponse.ProtoReflect.Descriptor instead.
func (*GetAllRemindersResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllRemindersResponse) GetReminders() []string {
	if x != nil {
		return x.Reminders
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6a, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x38, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x5b, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x61, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x72, 0x65, 0x61, 0x64, 0x41, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72,
	0x73, 0x32, 0xef, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x16, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x4c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x12, 0x17,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x6d, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x63, 0x68, 0x61, 0x63, 0x68, 0x74, 0x65, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_proto_goTypes = []interface{}{
	(*AddReminderRequest)(nil),        // 0: AddReminderRequest
	(*GetLatestReminderRequest)(nil),  // 1: GetLatestReminderRequest
	(*GetAllRemindersRequest)(nil),    // 2: GetAllRemindersRequest
	(*AddReminderResponse)(nil),       // 3: AddReminderResponse
	(*GetLatestReminderResponse)(nil), // 4: GetLatestReminderResponse
	(*GetAllRemindersResponse)(nil),   // 5: GetAllRemindersResponse
}
var file_service_proto_depIdxs = []int32{
	0, // 0: ReminderService.WriteReminder:input_type -> AddReminderRequest
	1, // 1: ReminderService.RetrieveLatestReminder:input_type -> GetLatestReminderRequest
	2, // 2: ReminderService.RetrieveAllReminders:input_type -> GetAllRemindersRequest
	3, // 3: ReminderService.WriteReminder:output_type -> AddReminderResponse
	4, // 4: ReminderService.RetrieveLatestReminder:output_type -> GetLatestReminderResponse
	5, // 5: ReminderService.RetrieveAllReminders:output_type -> GetAllRemindersResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
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
			switch v := v.(*AddReminderRequest); i {
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
			switch v := v.(*GetLatestReminderRequest); i {
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
			switch v := v.(*GetAllRemindersRequest); i {
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
			switch v := v.(*AddReminderResponse); i {
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
			switch v := v.(*GetLatestReminderResponse); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRemindersResponse); i {
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
			NumMessages:   6,
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

// ReminderServiceClient is the client API for ReminderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReminderServiceClient interface {
	WriteReminder(ctx context.Context, in *AddReminderRequest, opts ...grpc.CallOption) (*AddReminderResponse, error)
	RetrieveLatestReminder(ctx context.Context, in *GetLatestReminderRequest, opts ...grpc.CallOption) (*GetLatestReminderResponse, error)
	RetrieveAllReminders(ctx context.Context, in *GetAllRemindersRequest, opts ...grpc.CallOption) (*GetAllRemindersResponse, error)
}

type reminderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReminderServiceClient(cc grpc.ClientConnInterface) ReminderServiceClient {
	return &reminderServiceClient{cc}
}

func (c *reminderServiceClient) WriteReminder(ctx context.Context, in *AddReminderRequest, opts ...grpc.CallOption) (*AddReminderResponse, error) {
	out := new(AddReminderResponse)
	err := c.cc.Invoke(ctx, "/ReminderService/WriteReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reminderServiceClient) RetrieveLatestReminder(ctx context.Context, in *GetLatestReminderRequest, opts ...grpc.CallOption) (*GetLatestReminderResponse, error) {
	out := new(GetLatestReminderResponse)
	err := c.cc.Invoke(ctx, "/ReminderService/RetrieveLatestReminder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reminderServiceClient) RetrieveAllReminders(ctx context.Context, in *GetAllRemindersRequest, opts ...grpc.CallOption) (*GetAllRemindersResponse, error) {
	out := new(GetAllRemindersResponse)
	err := c.cc.Invoke(ctx, "/ReminderService/RetrieveAllReminders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReminderServiceServer is the server API for ReminderService service.
type ReminderServiceServer interface {
	WriteReminder(context.Context, *AddReminderRequest) (*AddReminderResponse, error)
	RetrieveLatestReminder(context.Context, *GetLatestReminderRequest) (*GetLatestReminderResponse, error)
	RetrieveAllReminders(context.Context, *GetAllRemindersRequest) (*GetAllRemindersResponse, error)
}

// UnimplementedReminderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReminderServiceServer struct {
}

func (*UnimplementedReminderServiceServer) WriteReminder(context.Context, *AddReminderRequest) (*AddReminderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteReminder not implemented")
}
func (*UnimplementedReminderServiceServer) RetrieveLatestReminder(context.Context, *GetLatestReminderRequest) (*GetLatestReminderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveLatestReminder not implemented")
}
func (*UnimplementedReminderServiceServer) RetrieveAllReminders(context.Context, *GetAllRemindersRequest) (*GetAllRemindersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveAllReminders not implemented")
}

func RegisterReminderServiceServer(s *grpc.Server, srv ReminderServiceServer) {
	s.RegisterService(&_ReminderService_serviceDesc, srv)
}

func _ReminderService_WriteReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReminderServiceServer).WriteReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReminderService/WriteReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReminderServiceServer).WriteReminder(ctx, req.(*AddReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReminderService_RetrieveLatestReminder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestReminderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReminderServiceServer).RetrieveLatestReminder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReminderService/RetrieveLatestReminder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReminderServiceServer).RetrieveLatestReminder(ctx, req.(*GetLatestReminderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReminderService_RetrieveAllReminders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRemindersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReminderServiceServer).RetrieveAllReminders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ReminderService/RetrieveAllReminders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReminderServiceServer).RetrieveAllReminders(ctx, req.(*GetAllRemindersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReminderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ReminderService",
	HandlerType: (*ReminderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteReminder",
			Handler:    _ReminderService_WriteReminder_Handler,
		},
		{
			MethodName: "RetrieveLatestReminder",
			Handler:    _ReminderService_RetrieveLatestReminder_Handler,
		},
		{
			MethodName: "RetrieveAllReminders",
			Handler:    _ReminderService_RetrieveAllReminders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}