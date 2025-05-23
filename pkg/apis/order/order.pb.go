// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: pkg/apis/order/order.proto

package order

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateUserAccountCommand struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LoginName     string                 `protobuf:"bytes,2,opt,name=login_name,json=loginName,proto3" json:"login_name,omitempty"`
	Active        bool                   `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserAccountCommand) Reset() {
	*x = CreateUserAccountCommand{}
	mi := &file_pkg_apis_order_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserAccountCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserAccountCommand) ProtoMessage() {}

func (x *CreateUserAccountCommand) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_order_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserAccountCommand.ProtoReflect.Descriptor instead.
func (*CreateUserAccountCommand) Descriptor() ([]byte, []int) {
	return file_pkg_apis_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserAccountCommand) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateUserAccountCommand) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

func (x *CreateUserAccountCommand) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type CreateUserAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserAccountResponse) Reset() {
	*x = CreateUserAccountResponse{}
	mi := &file_pkg_apis_order_order_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserAccountResponse) ProtoMessage() {}

func (x *CreateUserAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_order_order_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateUserAccountResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_order_order_proto_rawDescGZIP(), []int{1}
}

var File_pkg_apis_order_order_proto protoreflect.FileDescriptor

const file_pkg_apis_order_order_proto_rawDesc = "" +
	"\n" +
	"\x1apkg/apis/order/order.proto\x12\x02v1\"j\n" +
	"\x18CreateUserAccountCommand\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"login_name\x18\x02 \x01(\tR\tloginName\x12\x16\n" +
	"\x06active\x18\x03 \x01(\bR\x06active\"\x1b\n" +
	"\x19CreateUserAccountResponse2[\n" +
	"\x05Order\x12R\n" +
	"\x11CreateUserAccount\x12\x1c.v1.CreateUserAccountCommand\x1a\x1d.v1.CreateUserAccountResponse\"\x00B<Z:github.com/dothiphuc81299/coffeeShop-server/pkg/apis/orderb\x06proto3"

var (
	file_pkg_apis_order_order_proto_rawDescOnce sync.Once
	file_pkg_apis_order_order_proto_rawDescData []byte
)

func file_pkg_apis_order_order_proto_rawDescGZIP() []byte {
	file_pkg_apis_order_order_proto_rawDescOnce.Do(func() {
		file_pkg_apis_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_apis_order_order_proto_rawDesc), len(file_pkg_apis_order_order_proto_rawDesc)))
	})
	return file_pkg_apis_order_order_proto_rawDescData
}

var file_pkg_apis_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_apis_order_order_proto_goTypes = []any{
	(*CreateUserAccountCommand)(nil),  // 0: v1.CreateUserAccountCommand
	(*CreateUserAccountResponse)(nil), // 1: v1.CreateUserAccountResponse
}
var file_pkg_apis_order_order_proto_depIdxs = []int32{
	0, // 0: v1.Order.CreateUserAccount:input_type -> v1.CreateUserAccountCommand
	1, // 1: v1.Order.CreateUserAccount:output_type -> v1.CreateUserAccountResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_apis_order_order_proto_init() }
func file_pkg_apis_order_order_proto_init() {
	if File_pkg_apis_order_order_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_apis_order_order_proto_rawDesc), len(file_pkg_apis_order_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_order_order_proto_goTypes,
		DependencyIndexes: file_pkg_apis_order_order_proto_depIdxs,
		MessageInfos:      file_pkg_apis_order_order_proto_msgTypes,
	}.Build()
	File_pkg_apis_order_order_proto = out.File
	file_pkg_apis_order_order_proto_goTypes = nil
	file_pkg_apis_order_order_proto_depIdxs = nil
}
