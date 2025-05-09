// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/delivery/proto/delivery.proto

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

type DeliveryStatus int32

const (
	DeliveryStatus_PENDING    DeliveryStatus = 0
	DeliveryStatus_SCHEDULED  DeliveryStatus = 1
	DeliveryStatus_IN_TRANSIT DeliveryStatus = 2
	DeliveryStatus_DELIVERED  DeliveryStatus = 3
	DeliveryStatus_CANCELLED  DeliveryStatus = 4
	DeliveryStatus_DELAYED    DeliveryStatus = 5
)

// Enum value maps for DeliveryStatus.
var (
	DeliveryStatus_name = map[int32]string{
		0: "PENDING",
		1: "SCHEDULED",
		2: "IN_TRANSIT",
		3: "DELIVERED",
		4: "CANCELLED",
		5: "DELAYED",
	}
	DeliveryStatus_value = map[string]int32{
		"PENDING":    0,
		"SCHEDULED":  1,
		"IN_TRANSIT": 2,
		"DELIVERED":  3,
		"CANCELLED":  4,
		"DELAYED":    5,
	}
)

func (x DeliveryStatus) Enum() *DeliveryStatus {
	p := new(DeliveryStatus)
	*p = x
	return p
}

func (x DeliveryStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeliveryStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_services_delivery_proto_delivery_proto_enumTypes[0].Descriptor()
}

func (DeliveryStatus) Type() protoreflect.EnumType {
	return &file_services_delivery_proto_delivery_proto_enumTypes[0]
}

func (x DeliveryStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeliveryStatus.Descriptor instead.
func (DeliveryStatus) EnumDescriptor() ([]byte, []int) {
	return file_services_delivery_proto_delivery_proto_rawDescGZIP(), []int{0}
}

type DeliveryRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	OrderId         string     `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ScheduleTime    *Timestamp `protobuf:"bytes,2,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	DeliveryAddress string     `protobuf:"bytes,3,opt,name=delivery_address,json=deliveryAddress,proto3" json:"delivery_address,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *DeliveryRequest) Reset() {
	*x = DeliveryRequest{}
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeliveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryRequest) ProtoMessage() {}

func (x *DeliveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryRequest.ProtoReflect.Descriptor instead.
func (*DeliveryRequest) Descriptor() ([]byte, []int) {
	return file_services_delivery_proto_delivery_proto_rawDescGZIP(), []int{0}
}

func (x *DeliveryRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *DeliveryRequest) GetScheduleTime() *Timestamp {
	if x != nil {
		return x.ScheduleTime
	}
	return nil
}

func (x *DeliveryRequest) GetDeliveryAddress() string {
	if x != nil {
		return x.DeliveryAddress
	}
	return ""
}

type StatusUpdate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeliveryId    string                 `protobuf:"bytes,1,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id,omitempty"` // From path
	Status        DeliveryStatus `protobuf:"varint,2,opt,name=status,proto3,enum=delivery.DeliveryStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatusUpdate) Reset() {
	*x = StatusUpdate{}
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusUpdate) ProtoMessage() {}

func (x *StatusUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusUpdate.ProtoReflect.Descriptor instead.
func (*StatusUpdate) Descriptor() ([]byte, []int) {
	return file_services_delivery_proto_delivery_proto_rawDescGZIP(), []int{1}
}

func (x *StatusUpdate) GetDeliveryId() string {
	if x != nil {
		return x.DeliveryId
	}
	return ""
}

func (x *StatusUpdate) GetStatus() DeliveryStatus {
	if x != nil {
		return x.Status
	}
	return DeliveryStatus_PENDING
}

type DeliveryQuery struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeliveryId    string                 `protobuf:"bytes,1,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id,omitempty"` // From path
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeliveryQuery) Reset() {
	*x = DeliveryQuery{}
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeliveryQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryQuery) ProtoMessage() {}

func (x *DeliveryQuery) ProtoReflect() protoreflect.Message {
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryQuery.ProtoReflect.Descriptor instead.
func (*DeliveryQuery) Descriptor() ([]byte, []int) {
	return file_services_delivery_proto_delivery_proto_rawDescGZIP(), []int{2}
}

func (x *DeliveryQuery) GetDeliveryId() string {
	if x != nil {
		return x.DeliveryId
	}
	return ""
}

type DeliveryConfirmation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeliveryId    string                 `protobuf:"bytes,1,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id,omitempty"` // From path
	Signature     string                 `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeliveryConfirmation) Reset() {
	*x = DeliveryConfirmation{}
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeliveryConfirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryConfirmation) ProtoMessage() {}

func (x *DeliveryConfirmation) ProtoReflect() protoreflect.Message {
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryConfirmation.ProtoReflect.Descriptor instead.
func (*DeliveryConfirmation) Descriptor() ([]byte, []int) {
	return file_services_delivery_proto_delivery_proto_rawDescGZIP(), []int{3}
}

func (x *DeliveryConfirmation) GetDeliveryId() string {
	if x != nil {
		return x.DeliveryId
	}
	return ""
}

func (x *DeliveryConfirmation) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type DeliveryResponse struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	DeliveryId         string                 `protobuf:"bytes,1,opt,name=delivery_id,json=deliveryId,proto3" json:"delivery_id,omitempty"`
	OrderId            string         `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status             DeliveryStatus `protobuf:"varint,3,opt,name=status,proto3,enum=delivery.DeliveryStatus" json:"status,omitempty"`
	ScheduledTime      *Timestamp     `protobuf:"bytes,4,opt,name=scheduled_time,json=scheduledTime,proto3" json:"scheduled_time,omitempty"`
	ActualDeliveryTime *Timestamp     `protobuf:"bytes,5,opt,name=actual_delivery_time,json=actualDeliveryTime,proto3" json:"actual_delivery_time,omitempty"`
	Base               *BaseResponse  `protobuf:"bytes,6,opt,name=base,proto3" json:"base,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *DeliveryResponse) Reset() {
	*x = DeliveryResponse{}
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeliveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryResponse) ProtoMessage() {}

func (x *DeliveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryResponse.ProtoReflect.Descriptor instead.
func (*DeliveryResponse) Descriptor() ([]byte, []int) {
	return file_services_delivery_proto_delivery_proto_rawDescGZIP(), []int{4}
}

func (x *DeliveryResponse) GetDeliveryId() string {
	if x != nil {
		return x.DeliveryId
	}
	return ""
}

func (x *DeliveryResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *DeliveryResponse) GetStatus() DeliveryStatus {
	if x != nil {
		return x.Status
	}
	return DeliveryStatus_PENDING
}

func (x *DeliveryResponse) GetScheduledTime() *Timestamp {
	if x != nil {
		return x.ScheduledTime
	}
	return nil
}

func (x *DeliveryResponse) GetActualDeliveryTime() *Timestamp {
	if x != nil {
		return x.ActualDeliveryTime
	}
	return nil
}

func (x *DeliveryResponse) GetBase() *BaseResponse {
	if x != nil {
		return x.Base
	}
	return nil
}

// ----------------------
// Common Messages
// ----------------------
type BaseResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BaseResponse) Reset() {
	*x = BaseResponse{}
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResponse) ProtoMessage() {}

func (x *BaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[5]
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
	return file_services_delivery_proto_delivery_proto_rawDescGZIP(), []int{5}
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
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[6]
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
	return file_services_delivery_proto_delivery_proto_rawDescGZIP(), []int{6}
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
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_services_delivery_proto_delivery_proto_msgTypes[7]
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
	return file_services_delivery_proto_delivery_proto_rawDescGZIP(), []int{7}
}

var File_services_delivery_proto_delivery_proto protoreflect.FileDescriptor

const file_services_delivery_proto_delivery_proto_rawDesc = "" +
	"\n" +
	"&services/delivery/proto/delivery.proto\x12\bdelivery\x1a\x1cgoogle/api/annotations.proto\"\x91\x01\n" +
	"\x0fDeliveryRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x128\n" +
	"\rschedule_time\x18\x02 \x01(\v2\x13.delivery.TimestampR\fscheduleTime\x12)\n" +
	"\x10delivery_address\x18\x03 \x01(\tR\x0fdeliveryAddress\"a\n" +
	"\fStatusUpdate\x12\x1f\n" +
	"\vdelivery_id\x18\x01 \x01(\tR\n" +
	"deliveryId\x120\n" +
	"\x06status\x18\x02 \x01(\x0e2\x18.delivery.DeliveryStatusR\x06status\"0\n" +
	"\rDeliveryQuery\x12\x1f\n" +
	"\vdelivery_id\x18\x01 \x01(\tR\n" +
	"deliveryId\"U\n" +
	"\x14DeliveryConfirmation\x12\x1f\n" +
	"\vdelivery_id\x18\x01 \x01(\tR\n" +
	"deliveryId\x12\x1c\n" +
	"\tsignature\x18\x02 \x01(\tR\tsignature\"\xaf\x02\n" +
	"\x10DeliveryResponse\x12\x1f\n" +
	"\vdelivery_id\x18\x01 \x01(\tR\n" +
	"deliveryId\x12\x19\n" +
	"\border_id\x18\x02 \x01(\tR\aorderId\x120\n" +
	"\x06status\x18\x03 \x01(\x0e2\x18.delivery.DeliveryStatusR\x06status\x12:\n" +
	"\x0escheduled_time\x18\x04 \x01(\v2\x13.delivery.TimestampR\rscheduledTime\x12E\n" +
	"\x14actual_delivery_time\x18\x05 \x01(\v2\x13.delivery.TimestampR\x12actualDeliveryTime\x12*\n" +
	"\x04base\x18\x06 \x01(\v2\x16.delivery.BaseResponseR\x04base\"<\n" +
	"\fBaseResponse\x12\x12\n" +
	"\x04code\x18\x01 \x01(\x05R\x04code\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\";\n" +
	"\tTimestamp\x12\x18\n" +
	"\aseconds\x18\x01 \x01(\x03R\aseconds\x12\x14\n" +
	"\x05nanos\x18\x02 \x01(\x05R\x05nanos\"\a\n" +
	"\x05Empty*g\n" +
	"\x0eDeliveryStatus\x12\v\n" +
	"\aPENDING\x10\x00\x12\r\n" +
	"\tSCHEDULED\x10\x01\x12\x0e\n" +
	"\n" +
	"IN_TRANSIT\x10\x02\x12\r\n" +
	"\tDELIVERED\x10\x03\x12\r\n" +
	"\tCANCELLED\x10\x04\x12\v\n" +
	"\aDELAYED\x10\x052\xcc\x03\n" +
	"\x0fDeliveryService\x12a\n" +
	"\x10ScheduleDelivery\x12\x19.delivery.DeliveryRequest\x1a\x1a.delivery.DeliveryResponse\"\x16\x82\xd3\xe4\x93\x02\x10:\x01*\"\v/deliveries\x12w\n" +
	"\x14UpdateDeliveryStatus\x12\x16.delivery.StatusUpdate\x1a\x1a.delivery.DeliveryResponse\"+\x82\xd3\xe4\x93\x02%:\x01*\x1a /deliveries/{delivery_id}/status\x12g\n" +
	"\rTrackDelivery\x12\x17.delivery.DeliveryQuery\x1a\x1a.delivery.DeliveryResponse\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/deliveries/{delivery_id}\x12t\n" +
	"\x0fConfirmDelivery\x12\x1e.delivery.DeliveryConfirmation\x1a\x16.delivery.BaseResponse\")\x82\xd3\xe4\x93\x02#\"!/deliveries/{delivery_id}/confirmB\rZ\v./;deliveryb\x06proto3"

var (
	file_services_delivery_proto_delivery_proto_rawDescOnce sync.Once
	file_services_delivery_proto_delivery_proto_rawDescData []byte
)

func file_services_delivery_proto_delivery_proto_rawDescGZIP() []byte {
	file_services_delivery_proto_delivery_proto_rawDescOnce.Do(func() {
		file_services_delivery_proto_delivery_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_services_delivery_proto_delivery_proto_rawDesc), len(file_services_delivery_proto_delivery_proto_rawDesc)))
	})
	return file_services_delivery_proto_delivery_proto_rawDescData
}

var file_services_delivery_proto_delivery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_delivery_proto_delivery_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_delivery_proto_delivery_proto_goTypes = []any{
	(DeliveryStatus)(0),          // 0: delivery.DeliveryStatus
	(*DeliveryRequest)(nil),      // 1: delivery.DeliveryRequest
	(*StatusUpdate)(nil),         // 2: delivery.StatusUpdate
	(*DeliveryQuery)(nil),        // 3: delivery.DeliveryQuery
	(*DeliveryConfirmation)(nil), // 4: delivery.DeliveryConfirmation
	(*DeliveryResponse)(nil),     // 5: delivery.DeliveryResponse
	(*BaseResponse)(nil),         // 6: delivery.BaseResponse
	(*Timestamp)(nil),            // 7: delivery.Timestamp
	(*Empty)(nil),                // 8: delivery.Empty
}
var file_services_delivery_proto_delivery_proto_depIdxs = []int32{
	7,  // 0: delivery.DeliveryRequest.schedule_time:type_name -> delivery.Timestamp
	0,  // 1: delivery.StatusUpdate.status:type_name -> delivery.DeliveryStatus
	0,  // 2: delivery.DeliveryResponse.status:type_name -> delivery.DeliveryStatus
	7,  // 3: delivery.DeliveryResponse.scheduled_time:type_name -> delivery.Timestamp
	7,  // 4: delivery.DeliveryResponse.actual_delivery_time:type_name -> delivery.Timestamp
	6,  // 5: delivery.DeliveryResponse.base:type_name -> delivery.BaseResponse
	1,  // 6: delivery.DeliveryService.ScheduleDelivery:input_type -> delivery.DeliveryRequest
	2,  // 7: delivery.DeliveryService.UpdateDeliveryStatus:input_type -> delivery.StatusUpdate
	3,  // 8: delivery.DeliveryService.TrackDelivery:input_type -> delivery.DeliveryQuery
	4,  // 9: delivery.DeliveryService.ConfirmDelivery:input_type -> delivery.DeliveryConfirmation
	5,  // 10: delivery.DeliveryService.ScheduleDelivery:output_type -> delivery.DeliveryResponse
	5,  // 11: delivery.DeliveryService.UpdateDeliveryStatus:output_type -> delivery.DeliveryResponse
	5,  // 12: delivery.DeliveryService.TrackDelivery:output_type -> delivery.DeliveryResponse
	6,  // 13: delivery.DeliveryService.ConfirmDelivery:output_type -> delivery.BaseResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_services_delivery_proto_delivery_proto_init() }
func file_services_delivery_proto_delivery_proto_init() {
	if File_services_delivery_proto_delivery_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_delivery_proto_delivery_proto_rawDesc), len(file_services_delivery_proto_delivery_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_delivery_proto_delivery_proto_goTypes,
		DependencyIndexes: file_services_delivery_proto_delivery_proto_depIdxs,
		EnumInfos:         file_services_delivery_proto_delivery_proto_enumTypes,
		MessageInfos:      file_services_delivery_proto_delivery_proto_msgTypes,
	}.Build()
	File_services_delivery_proto_delivery_proto = out.File
	file_services_delivery_proto_delivery_proto_goTypes = nil
	file_services_delivery_proto_delivery_proto_depIdxs = nil
}
