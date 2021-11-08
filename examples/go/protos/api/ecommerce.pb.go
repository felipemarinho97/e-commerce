// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: protos/api/ecommerce.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity int64 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_ecommerce_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_ecommerce_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRequest.ProtoReflect.Descriptor instead.
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return file_protos_api_ecommerce_proto_rawDescGZIP(), []int{0}
}

func (x *ProductRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity    int64 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UnitAmount  int64 `protobuf:"varint,3,opt,name=unit_amount,json=unitAmount,proto3" json:"unit_amount,omitempty"`
	TotalAmount int64 `protobuf:"varint,4,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	Discount    int64 `protobuf:"varint,5,opt,name=discount,proto3" json:"discount,omitempty"`
	IsGift      bool  `protobuf:"varint,6,opt,name=is_gift,json=isGift,proto3" json:"is_gift,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_ecommerce_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_ecommerce_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_protos_api_ecommerce_proto_rawDescGZIP(), []int{1}
}

func (x *ProductResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductResponse) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ProductResponse) GetUnitAmount() int64 {
	if x != nil {
		return x.UnitAmount
	}
	return 0
}

func (x *ProductResponse) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *ProductResponse) GetDiscount() int64 {
	if x != nil {
		return x.Discount
	}
	return 0
}

func (x *ProductResponse) GetIsGift() bool {
	if x != nil {
		return x.IsGift
	}
	return false
}

type CheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ProductRequest `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *CheckoutRequest) Reset() {
	*x = CheckoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_ecommerce_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequest) ProtoMessage() {}

func (x *CheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_ecommerce_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequest.ProtoReflect.Descriptor instead.
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return file_protos_api_ecommerce_proto_rawDescGZIP(), []int{2}
}

func (x *CheckoutRequest) GetProducts() []*ProductRequest {
	if x != nil {
		return x.Products
	}
	return nil
}

type CheckoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalAmount             int64              `protobuf:"varint,1,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	TotalAmountWithDiscount int64              `protobuf:"varint,2,opt,name=total_amount_with_discount,json=totalAmountWithDiscount,proto3" json:"total_amount_with_discount,omitempty"`
	TotalDiscount           int64              `protobuf:"varint,3,opt,name=total_discount,json=totalDiscount,proto3" json:"total_discount,omitempty"`
	Products                []*ProductResponse `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *CheckoutResponse) Reset() {
	*x = CheckoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_api_ecommerce_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutResponse) ProtoMessage() {}

func (x *CheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_api_ecommerce_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutResponse.ProtoReflect.Descriptor instead.
func (*CheckoutResponse) Descriptor() ([]byte, []int) {
	return file_protos_api_ecommerce_proto_rawDescGZIP(), []int{3}
}

func (x *CheckoutResponse) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *CheckoutResponse) GetTotalAmountWithDiscount() int64 {
	if x != nil {
		return x.TotalAmountWithDiscount
	}
	return 0
}

func (x *CheckoutResponse) GetTotalDiscount() int64 {
	if x != nil {
		return x.TotalDiscount
	}
	return 0
}

func (x *CheckoutResponse) GetProducts() []*ProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_protos_api_ecommerce_proto protoreflect.FileDescriptor

var file_protos_api_ecommerce_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0xb6, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x6e, 0x69, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x67, 0x69, 0x66, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x47, 0x69, 0x66, 0x74, 0x22, 0x48, 0x0a, 0x0f,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3b,
	0x0a, 0x1a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x17, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x70, 0x0a, 0x10, 0x45, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c,
	0x0a, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x2e, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x65, 0x6c, 0x69, 0x70,
	0x65, 0x6d, 0x61, 0x72, 0x69, 0x6e, 0x68, 0x6f, 0x39, 0x37, 0x2f, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_api_ecommerce_proto_rawDescOnce sync.Once
	file_protos_api_ecommerce_proto_rawDescData = file_protos_api_ecommerce_proto_rawDesc
)

func file_protos_api_ecommerce_proto_rawDescGZIP() []byte {
	file_protos_api_ecommerce_proto_rawDescOnce.Do(func() {
		file_protos_api_ecommerce_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_api_ecommerce_proto_rawDescData)
	})
	return file_protos_api_ecommerce_proto_rawDescData
}

var file_protos_api_ecommerce_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_api_ecommerce_proto_goTypes = []interface{}{
	(*ProductRequest)(nil),   // 0: ecommerce.ProductRequest
	(*ProductResponse)(nil),  // 1: ecommerce.ProductResponse
	(*CheckoutRequest)(nil),  // 2: ecommerce.CheckoutRequest
	(*CheckoutResponse)(nil), // 3: ecommerce.CheckoutResponse
}
var file_protos_api_ecommerce_proto_depIdxs = []int32{
	0, // 0: ecommerce.CheckoutRequest.products:type_name -> ecommerce.ProductRequest
	1, // 1: ecommerce.CheckoutResponse.products:type_name -> ecommerce.ProductResponse
	2, // 2: ecommerce.EcommerceService.Checkout:input_type -> ecommerce.CheckoutRequest
	3, // 3: ecommerce.EcommerceService.Checkout:output_type -> ecommerce.CheckoutResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_api_ecommerce_proto_init() }
func file_protos_api_ecommerce_proto_init() {
	if File_protos_api_ecommerce_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_api_ecommerce_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRequest); i {
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
		file_protos_api_ecommerce_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
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
		file_protos_api_ecommerce_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutRequest); i {
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
		file_protos_api_ecommerce_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutResponse); i {
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
			RawDescriptor: file_protos_api_ecommerce_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_api_ecommerce_proto_goTypes,
		DependencyIndexes: file_protos_api_ecommerce_proto_depIdxs,
		MessageInfos:      file_protos_api_ecommerce_proto_msgTypes,
	}.Build()
	File_protos_api_ecommerce_proto = out.File
	file_protos_api_ecommerce_proto_rawDesc = nil
	file_protos_api_ecommerce_proto_goTypes = nil
	file_protos_api_ecommerce_proto_depIdxs = nil
}