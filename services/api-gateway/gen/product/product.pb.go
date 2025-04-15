// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/product/proto/product.proto

package product

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// common
type BaseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	mi := &file_services_product_proto_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResponse.ProtoReflect.Descriptor instead.
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return file_services_product_proto_product_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Timestamp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Seconds       int64                  `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos         int32                  `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
	mi := &file_services_product_proto_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timestamp.ProtoReflect.Descriptor instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_services_product_proto_product_proto_rawDescGZIP(), []int{1}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_services_product_proto_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_services_product_proto_product_proto_rawDescGZIP(), []int{2}
}

type ProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductRequest) Reset() {
	*x = ProductRequest{}
	mi := &file_services_product_proto_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRequest) ProtoMessage() {}

func (x *ProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_product_proto_msgTypes[3]
	if x != nil {
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
	return file_services_product_proto_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ProductID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductID) Reset() {
	*x = ProductID{}
	mi := &file_services_product_proto_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductID) ProtoMessage() {}

func (x *ProductID) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductID.ProtoReflect.Descriptor instead.
func (*ProductID) Descriptor() ([]byte, []int) {
	return file_services_product_proto_product_proto_rawDescGZIP(), []int{4}
}

func (x *ProductID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ProductUpdate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductUpdate) Reset() {
	*x = ProductUpdate{}
	mi := &file_services_product_proto_product_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductUpdate) ProtoMessage() {}

func (x *ProductUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_product_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductUpdate.ProtoReflect.Descriptor instead.
func (*ProductUpdate) Descriptor() ([]byte, []int) {
	return file_services_product_proto_product_proto_rawDescGZIP(), []int{5}
}

func (x *ProductUpdate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductUpdate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductUpdate) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt     *Timestamp             `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Base          *BaseResponse          `protobuf:"bytes,6,opt,name=base,proto3" json:"base,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	mi := &file_services_product_proto_product_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_product_proto_msgTypes[6]
	if x != nil {
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
	return file_services_product_proto_product_proto_rawDescGZIP(), []int{6}
}

func (x *ProductResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductResponse) GetCreatedAt() *Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ProductResponse) GetBase() *BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

type ProductListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*ProductResponse     `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Total         int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Base          *BaseResponse          `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductListResponse) Reset() {
	*x = ProductListResponse{}
	mi := &file_services_product_proto_product_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductListResponse) ProtoMessage() {}

func (x *ProductListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_product_proto_product_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductListResponse.ProtoReflect.Descriptor instead.
func (*ProductListResponse) Descriptor() ([]byte, []int) {
	return file_services_product_proto_product_proto_rawDescGZIP(), []int{7}
}

func (x *ProductListResponse) GetProducts() []*ProductResponse {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *ProductListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ProductListResponse) GetBase() *BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

var File_services_product_proto_product_proto protoreflect.FileDescriptor

const file_services_product_proto_product_proto_rawDesc = "" +
	"\n" +
	"$services/product/proto/product.proto\x12\aproduct\x1a\x1cgoogle/api/annotations.proto\x1a\x15google/api/http.proto\"<\n" +
	"\fBaseResponse\x12\x12\n" +
	"\x04code\x18\x01 \x01(\x05R\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\";\n" +
	"\tTimestamp\x12\x18\n" +
	"\aseconds\x18\x01 \x01(\x03R\aseconds\x12\x14\n" +
	"\x05nanos\x18\x02 \x01(\x05R\x05nanos\"\a\n" +
	"\x05Empty\"l\n" +
	"\x0eProductRequest\x12\x0e\n" +
	"\x02id\x18\x04 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x01R\x05price\"\x1b\n" +
	"\tProductID\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"k\n" +
	"\rProductUpdate\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\"\xcb\x01\n" +
	"\x0fProductResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x14\n" +
	"\x05price\x18\x04 \x01(\x01R\x05price\x121\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x12.product.TimestampR\tcreatedAt\x12)\n" +
	"\x04base\x18\x06 \x01(\v2\x15.product.BaseResponseR\x04base\"\x8c\x01\n" +
	"\x13ProductListResponse\x124\n" +
	"\bproducts\x18\x01 \x03(\v2\x18.product.ProductResponseR\bproducts\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x05R\x05total\x12)\n" +
	"\x04base\x18\x03 \x01(\v2\x15.product.BaseResponseR\x04base2\xc6\x03\n" +
	"\x0eProductService\x12]\n" +
	"\rCreateProduct\x12\x17.product.ProductRequest\x1a\x18.product.ProductResponse\"\x19\x82\xd3\xe4\x93\x02\x13:\x01*\"\x0e/v1/items/{id}\x12R\n" +
	"\n" +
	"GetProduct\x12\x12.product.ProductID\x1a\x18.product.ProductResponse\"\x16\x82\xd3\xe4\x93\x02\x10\x12\x0e/v1/items/{id}\x12O\n" +
	"\fListProducts\x12\x0e.product.Empty\x1a\x1c.product.ProductListResponse\"\x11\x82\xd3\xe4\x93\x02\v\x12\t/v1/items\x12\\\n" +
	"\rUpdateProduct\x12\x16.product.ProductUpdate\x1a\x18.product.ProductResponse\"\x19\x82\xd3\xe4\x93\x02\x13:\x01*\x1a\x0e/v1/items/{id}\x12R\n" +
	"\rDeleteProduct\x12\x12.product.ProductID\x1a\x15.product.BaseResponse\"\x16\x82\xd3\xe4\x93\x02\x10*\x0e/v1/items/{id}B\fZ\n" +
	"./;productb\x06proto3"

var (
	file_services_product_proto_product_proto_rawDescOnce sync.Once
	file_services_product_proto_product_proto_rawDescData []byte
)

func file_services_product_proto_product_proto_rawDescGZIP() []byte {
	file_services_product_proto_product_proto_rawDescOnce.Do(func() {
		file_services_product_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_product_proto_product_proto_rawDesc), len(file_services_product_proto_product_proto_rawDesc)))
	})
	return file_services_product_proto_product_proto_rawDescData
}

var file_services_product_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_product_proto_product_proto_goTypes = []any{
	(*BaseResponse)(nil),        // 0: product.BaseResponse
	(*Timestamp)(nil),           // 1: product.Timestamp
	(*Empty)(nil),               // 2: product.Empty
	(*ProductRequest)(nil),      // 3: product.ProductRequest
	(*ProductID)(nil),           // 4: product.ProductID
	(*ProductUpdate)(nil),       // 5: product.ProductUpdate
	(*ProductResponse)(nil),     // 6: product.ProductResponse
	(*ProductListResponse)(nil), // 7: product.ProductListResponse
}
var file_services_product_proto_product_proto_depIdxs = []int32{
	1, // 0: product.ProductResponse.created_at:type_name -> product.Timestamp
	0, // 1: product.ProductResponse.base:type_name -> product.BaseResponse
	6, // 2: product.ProductListResponse.products:type_name -> product.ProductResponse
	0, // 3: product.ProductListResponse.base:type_name -> product.BaseResponse
	3, // 4: product.ProductService.CreateProduct:input_type -> product.ProductRequest
	4, // 5: product.ProductService.GetProduct:input_type -> product.ProductID
	2, // 6: product.ProductService.ListProducts:input_type -> product.Empty
	5, // 7: product.ProductService.UpdateProduct:input_type -> product.ProductUpdate
	4, // 8: product.ProductService.DeleteProduct:input_type -> product.ProductID
	6, // 9: product.ProductService.CreateProduct:output_type -> product.ProductResponse
	6, // 10: product.ProductService.GetProduct:output_type -> product.ProductResponse
	7, // 11: product.ProductService.ListProducts:output_type -> product.ProductListResponse
	6, // 12: product.ProductService.UpdateProduct:output_type -> product.ProductResponse
	0, // 13: product.ProductService.DeleteProduct:output_type -> product.BaseResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_product_proto_product_proto_init() }
func file_services_product_proto_product_proto_init() {
	if File_services_product_proto_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_product_proto_product_proto_rawDesc), len(file_services_product_proto_product_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_product_proto_product_proto_goTypes,
		DependencyIndexes: file_services_product_proto_product_proto_depIdxs,
		MessageInfos:      file_services_product_proto_product_proto_msgTypes,
	}.Build()
	File_services_product_proto_product_proto = out.File
	file_services_product_proto_product_proto_goTypes = nil
	file_services_product_proto_product_proto_depIdxs = nil
}
