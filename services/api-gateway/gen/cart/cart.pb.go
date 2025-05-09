// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/cart/proto/cart.proto

package gen

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

type AddToCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // From path
	ProductId     string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity      int32                  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddToCartRequest) Reset() {
	*x = AddToCartRequest{}
	mi := &file_services_cart_proto_cart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToCartRequest) ProtoMessage() {}

func (x *AddToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cart_proto_cart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToCartRequest.ProtoReflect.Descriptor instead.
func (*AddToCartRequest) Descriptor() ([]byte, []int) {
	return file_services_cart_proto_cart_proto_rawDescGZIP(), []int{0}
}

func (x *AddToCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddToCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddToCartRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type RemoveFromCartRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // From path
	ProductId     string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"` // From path
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveFromCartRequest) Reset() {
	*x = RemoveFromCartRequest{}
	mi := &file_services_cart_proto_cart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveFromCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFromCartRequest) ProtoMessage() {}

func (x *RemoveFromCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cart_proto_cart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFromCartRequest.ProtoReflect.Descriptor instead.
func (*RemoveFromCartRequest) Descriptor() ([]byte, []int) {
	return file_services_cart_proto_cart_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveFromCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemoveFromCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type PaymentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // From path
	PaymentMethod string                 `protobuf:"bytes,2,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	Amount        float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	mi := &file_services_cart_proto_cart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_cart_proto_cart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_services_cart_proto_cart_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PaymentRequest) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *PaymentRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PaymentRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

type PaymentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TransactionId string        `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	PaymentTime   *Timestamp    `protobuf:"bytes,2,opt,name=payment_time,json=paymentTime,proto3" json:"payment_time,omitempty"`
	Base          *BaseResponse `protobuf:"bytes,3,opt,name=base,proto3" json:"base,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	mi := &file_services_cart_proto_cart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_cart_proto_cart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_services_cart_proto_cart_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentResponse) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PaymentResponse) GetPaymentTime() *Timestamp {
	if x != nil {
		return x.PaymentTime
	}
	return nil
}

func (x *PaymentResponse) GetBase() *BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

type BaseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	mi := &file_services_cart_proto_cart_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_cart_proto_cart_proto_msgTypes[4]
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
	return file_services_cart_proto_cart_proto_rawDescGZIP(), []int{4}
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
	mi := &file_services_cart_proto_cart_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_services_cart_proto_cart_proto_msgTypes[5]
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
	return file_services_cart_proto_cart_proto_rawDescGZIP(), []int{5}
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
	mi := &file_services_cart_proto_cart_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_services_cart_proto_cart_proto_msgTypes[6]
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
	return file_services_cart_proto_cart_proto_rawDescGZIP(), []int{6}
}

var File_services_cart_proto_cart_proto protoreflect.FileDescriptor

const file_services_cart_proto_cart_proto_rawDesc = "" +
	"\n" +
	"\x1eservices/cart/proto/cart.proto\x12\x04cart\x1a\x1cgoogle/api/annotations.proto\"f\n" +
	"\x10AddToCartRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"product_id\x18\x02 \x01(\tR\tproductId\x12\x1a\n" +
	"\bquantity\x18\x03 \x01(\x05R\bquantity\"O\n" +
	"\x15RemoveFromCartRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"product_id\x18\x02 \x01(\tR\tproductId\"\x84\x01\n" +
	"\x0ePaymentRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12%\n" +
	"\x0epayment_method\x18\x02 \x01(\tR\rpaymentMethod\x12\x16\n" +
	"\x06amount\x18\x03 \x01(\x01R\x06amount\x12\x1a\n" +
	"\bcurrency\x18\x04 \x01(\tR\bcurrency\"\x94\x01\n" +
	"\x0fPaymentResponse\x12%\n" +
	"\x0etransaction_id\x18\x01 \x01(\tR\rtransactionId\x122\n" +
	"\fpayment_time\x18\x02 \x01(\v2\x0f.cart.TimestampR\vpaymentTime\x12&\n" +
	"\x04base\x18\x03 \x01(\v2\x12.cart.BaseResponseR\x04base\"<\n" +
	"\fBaseResponse\x12\x12\n" +
	"\x04code\x18\x01 \x01(\x05R\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\";\n" +
	"\tTimestamp\x12\x18\n" +
	"\aseconds\x18\x01 \x01(\x03R\aseconds\x12\x14\n" +
	"\x05nanos\x18\x02 \x01(\x05R\x05nanos\"\a\n" +
	"\x05Empty2\xba\x02\n" +
	"\vCartService\x12Y\n" +
	"\tAddToCart\x12\x16.cart.AddToCartRequest\x1a\x12.cart.BaseResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*\"\x15/cart/{user_id}/items\x12m\n" +
	"\x0eRemoveFromCart\x12\x1b.cart.RemoveFromCartRequest\x1a\x12.cart.BaseResponse\"*\x82\xd3\xe4\x93\x02$*\"/cart/{user_id}/items/{product_id}\x12a\n" +
	"\x0eProcessPayment\x12\x14.cart.PaymentRequest\x1a\x15.cart.PaymentResponse\"\"\x82\xd3\xe4\x93\x02\x1c:\x01*\"\x17/cart/{user_id}/paymentB\tZ\a./;cartb\x06proto3"

var (
	file_services_cart_proto_cart_proto_rawDescOnce sync.Once
	file_services_cart_proto_cart_proto_rawDescData []byte
)

func file_services_cart_proto_cart_proto_rawDescGZIP() []byte {
	file_services_cart_proto_cart_proto_rawDescOnce.Do(func() {
		file_services_cart_proto_cart_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_cart_proto_cart_proto_rawDesc), len(file_services_cart_proto_cart_proto_rawDesc)))
	})
	return file_services_cart_proto_cart_proto_rawDescData
}

var file_services_cart_proto_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_services_cart_proto_cart_proto_goTypes = []any{
	(*AddToCartRequest)(nil),      // 0: cart.AddToCartRequest
	(*RemoveFromCartRequest)(nil), // 1: cart.RemoveFromCartRequest
	(*PaymentRequest)(nil),        // 2: cart.PaymentRequest
	(*PaymentResponse)(nil),       // 3: cart.PaymentResponse
	(*BaseResponse)(nil),          // 4: cart.BaseResponse
	(*Timestamp)(nil),             // 5: cart.Timestamp
	(*Empty)(nil),                 // 6: cart.Empty
}
var file_services_cart_proto_cart_proto_depIdxs = []int32{
	5, // 0: cart.PaymentResponse.payment_time:type_name -> cart.Timestamp
	4, // 1: cart.PaymentResponse.base:type_name -> cart.BaseResponse
	0, // 2: cart.CartService.AddToCart:input_type -> cart.AddToCartRequest
	1, // 3: cart.CartService.RemoveFromCart:input_type -> cart.RemoveFromCartRequest
	2, // 4: cart.CartService.ProcessPayment:input_type -> cart.PaymentRequest
	4, // 5: cart.CartService.AddToCart:output_type -> cart.BaseResponse
	4, // 6: cart.CartService.RemoveFromCart:output_type -> cart.BaseResponse
	3, // 7: cart.CartService.ProcessPayment:output_type -> cart.PaymentResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_services_cart_proto_cart_proto_init() }
func file_services_cart_proto_cart_proto_init() {
	if File_services_cart_proto_cart_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_cart_proto_cart_proto_rawDesc), len(file_services_cart_proto_cart_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_cart_proto_cart_proto_goTypes,
		DependencyIndexes: file_services_cart_proto_cart_proto_depIdxs,
		MessageInfos:      file_services_cart_proto_cart_proto_msgTypes,
	}.Build()
	File_services_cart_proto_cart_proto = out.File
	file_services_cart_proto_cart_proto_goTypes = nil
	file_services_cart_proto_cart_proto_depIdxs = nil
}
