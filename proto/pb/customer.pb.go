// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: proto/customer.proto

package pb

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

// Define the Customer message
type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`          // Unique identifier for the customer
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`       // Name of the customer
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`     // Email of the customer
	Phone   string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`     // Phone number of the customer
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"` // Address of the customer
}

func (x *Customer) Reset() {
	*x = Customer{}
	mi := &file_proto_customer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Customer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Request for CreateCustomer
type CreateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`       // Name of the customer
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`     // Email of the customer
	Phone   string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`     // Phone number of the customer
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"` // Customer details
}

func (x *CreateCustomerRequest) Reset() {
	*x = CreateCustomerRequest{}
	mi := &file_proto_customer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerRequest) ProtoMessage() {}

func (x *CreateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCustomerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCustomerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateCustomerRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateCustomerRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

// Response for CreateCustomer
type CreateCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`           // ID of the created customer
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // Status message
}

func (x *CreateCustomerResponse) Reset() {
	*x = CreateCustomerResponse{}
	mi := &file_proto_customer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCustomerResponse) ProtoMessage() {}

func (x *CreateCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCustomerResponse.ProtoReflect.Descriptor instead.
func (*CreateCustomerResponse) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCustomerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateCustomerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Request for GetCustomerByEmail
type GetCustomerByEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"` // Email of the customer to retrieve
}

func (x *GetCustomerByEmailRequest) Reset() {
	*x = GetCustomerByEmailRequest{}
	mi := &file_proto_customer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByEmailRequest) ProtoMessage() {}

func (x *GetCustomerByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByEmailRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerByEmailRequest) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{3}
}

func (x *GetCustomerByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Response for GetCustomerByEmail
type GetCustomerByEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"` // Customer details
}

func (x *GetCustomerByEmailResponse) Reset() {
	*x = GetCustomerByEmailResponse{}
	mi := &file_proto_customer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerByEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByEmailResponse) ProtoMessage() {}

func (x *GetCustomerByEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByEmailResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerByEmailResponse) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{4}
}

func (x *GetCustomerByEmailResponse) GetCustomer() *Customer {
	if x != nil {
		return x.Customer
	}
	return nil
}

type GetCustomerByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCustomerByIDRequest) Reset() {
	*x = GetCustomerByIDRequest{}
	mi := &file_proto_customer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCustomerByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerByIDRequest) ProtoMessage() {}

func (x *GetCustomerByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_customer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerByIDRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_customer_proto_rawDescGZIP(), []int{5}
}

func (x *GetCustomerByIDRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_customer_proto protoreflect.FileDescriptor

var file_proto_customer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x69, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x08, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x71, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x42, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x52, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x69, 0x74,
	0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x28,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xcc, 0x02, 0x0a, 0x0f, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x25,
	0x2e, 0x67, 0x69, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67, 0x69, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x29, 0x2e, 0x67, 0x69, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x67, 0x69, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x26, 0x2e,
	0x67, 0x69, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x69, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22,
	0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_customer_proto_rawDescOnce sync.Once
	file_proto_customer_proto_rawDescData = file_proto_customer_proto_rawDesc
)

func file_proto_customer_proto_rawDescGZIP() []byte {
	file_proto_customer_proto_rawDescOnce.Do(func() {
		file_proto_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_customer_proto_rawDescData)
	})
	return file_proto_customer_proto_rawDescData
}

var file_proto_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_customer_proto_goTypes = []any{
	(*Customer)(nil),                   // 0: git_test.proto.Customer
	(*CreateCustomerRequest)(nil),      // 1: git_test.proto.CreateCustomerRequest
	(*CreateCustomerResponse)(nil),     // 2: git_test.proto.CreateCustomerResponse
	(*GetCustomerByEmailRequest)(nil),  // 3: git_test.proto.GetCustomerByEmailRequest
	(*GetCustomerByEmailResponse)(nil), // 4: git_test.proto.GetCustomerByEmailResponse
	(*GetCustomerByIDRequest)(nil),     // 5: git_test.proto.GetCustomerByIDRequest
}
var file_proto_customer_proto_depIdxs = []int32{
	0, // 0: git_test.proto.GetCustomerByEmailResponse.customer:type_name -> git_test.proto.Customer
	1, // 1: git_test.proto.CustomerService.CreateCustomer:input_type -> git_test.proto.CreateCustomerRequest
	3, // 2: git_test.proto.CustomerService.GetCustomerByEmail:input_type -> git_test.proto.GetCustomerByEmailRequest
	5, // 3: git_test.proto.CustomerService.GetCustomerByID:input_type -> git_test.proto.GetCustomerByIDRequest
	2, // 4: git_test.proto.CustomerService.CreateCustomer:output_type -> git_test.proto.CreateCustomerResponse
	4, // 5: git_test.proto.CustomerService.GetCustomerByEmail:output_type -> git_test.proto.GetCustomerByEmailResponse
	0, // 6: git_test.proto.CustomerService.GetCustomerByID:output_type -> git_test.proto.Customer
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_customer_proto_init() }
func file_proto_customer_proto_init() {
	if File_proto_customer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_customer_proto_goTypes,
		DependencyIndexes: file_proto_customer_proto_depIdxs,
		MessageInfos:      file_proto_customer_proto_msgTypes,
	}.Build()
	File_proto_customer_proto = out.File
	file_proto_customer_proto_rawDesc = nil
	file_proto_customer_proto_goTypes = nil
	file_proto_customer_proto_depIdxs = nil
}
