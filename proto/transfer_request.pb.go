// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.3
// source: transfer_request.proto

package proto

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum representing the types of documents.
type DocumentType int32

const (
	DocumentType_DOCUMENT_TYPE_UNSPECIFIED DocumentType = 0 // Unspecified document type
	DocumentType_DOCUMENT_TYPE_LEGAL       DocumentType = 1 // Record Sheet of the Unified State Register of Legal Entities (EGRUL)
	DocumentType_DOCUMENT_TYPE_INHERITANCE DocumentType = 2 // Certificate of Inheritance
	DocumentType_DOCUMENT_TYPE_JUDGMENT    DocumentType = 3 // Writ of Execution
)

// Enum value maps for DocumentType.
var (
	DocumentType_name = map[int32]string{
		0: "DOCUMENT_TYPE_UNSPECIFIED",
		1: "DOCUMENT_TYPE_LEGAL",
		2: "DOCUMENT_TYPE_INHERITANCE",
		3: "DOCUMENT_TYPE_JUDGMENT",
	}
	DocumentType_value = map[string]int32{
		"DOCUMENT_TYPE_UNSPECIFIED": 0,
		"DOCUMENT_TYPE_LEGAL":       1,
		"DOCUMENT_TYPE_INHERITANCE": 2,
		"DOCUMENT_TYPE_JUDGMENT":    3,
	}
)

func (x DocumentType) Enum() *DocumentType {
	p := new(DocumentType)
	*p = x
	return p
}

func (x DocumentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DocumentType) Descriptor() protoreflect.EnumDescriptor {
	return file_transfer_request_proto_enumTypes[0].Descriptor()
}

func (DocumentType) Type() protoreflect.EnumType {
	return &file_transfer_request_proto_enumTypes[0]
}

func (x DocumentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DocumentType.Descriptor instead.
func (DocumentType) EnumDescriptor() ([]byte, []int) {
	return file_transfer_request_proto_rawDescGZIP(), []int{0}
}

// Enum representing the basis for the transfer.
type TransferBasis int32

const (
	TransferBasis_TRANSFER_BASIS_UNSPECIFIED    TransferBasis = 0 // Unspecified basis
	TransferBasis_TRANSFER_BASIS_REORGANIZATION TransferBasis = 1 // Reorganization of a legal entity
	TransferBasis_TRANSFER_BASIS_INHERITANCE    TransferBasis = 2 // Inheritance
	TransferBasis_TRANSFER_BASIS_COURT_DECISION TransferBasis = 3 // Court decision
)

// Enum value maps for TransferBasis.
var (
	TransferBasis_name = map[int32]string{
		0: "TRANSFER_BASIS_UNSPECIFIED",
		1: "TRANSFER_BASIS_REORGANIZATION",
		2: "TRANSFER_BASIS_INHERITANCE",
		3: "TRANSFER_BASIS_COURT_DECISION",
	}
	TransferBasis_value = map[string]int32{
		"TRANSFER_BASIS_UNSPECIFIED":    0,
		"TRANSFER_BASIS_REORGANIZATION": 1,
		"TRANSFER_BASIS_INHERITANCE":    2,
		"TRANSFER_BASIS_COURT_DECISION": 3,
	}
)

func (x TransferBasis) Enum() *TransferBasis {
	p := new(TransferBasis)
	*p = x
	return p
}

func (x TransferBasis) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferBasis) Descriptor() protoreflect.EnumDescriptor {
	return file_transfer_request_proto_enumTypes[1].Descriptor()
}

func (TransferBasis) Type() protoreflect.EnumType {
	return &file_transfer_request_proto_enumTypes[1]
}

func (x TransferBasis) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferBasis.Descriptor instead.
func (TransferBasis) EnumDescriptor() ([]byte, []int) {
	return file_transfer_request_proto_rawDescGZIP(), []int{1}
}

// Enum representing the types of balances.
type BalanceType int32

const (
	BalanceType_BALANCE_TYPE_UNSPECIFIED           BalanceType = 0  // Unspecified balance type
	BalanceType_BALANCE_TYPE_TOKEN                 BalanceType = 43 // 0x2b
	BalanceType_BALANCE_TYPE_TOKEN_EXTERNAL_LOCKED BalanceType = 50 // 0x32
)

// Enum value maps for BalanceType.
var (
	BalanceType_name = map[int32]string{
		0:  "BALANCE_TYPE_UNSPECIFIED",
		43: "BALANCE_TYPE_TOKEN",
		50: "BALANCE_TYPE_TOKEN_EXTERNAL_LOCKED",
	}
	BalanceType_value = map[string]int32{
		"BALANCE_TYPE_UNSPECIFIED":           0,
		"BALANCE_TYPE_TOKEN":                 43,
		"BALANCE_TYPE_TOKEN_EXTERNAL_LOCKED": 50,
	}
)

func (x BalanceType) Enum() *BalanceType {
	p := new(BalanceType)
	*p = x
	return p
}

func (x BalanceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BalanceType) Descriptor() protoreflect.EnumDescriptor {
	return file_transfer_request_proto_enumTypes[2].Descriptor()
}

func (BalanceType) Type() protoreflect.EnumType {
	return &file_transfer_request_proto_enumTypes[2]
}

func (x BalanceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BalanceType.Descriptor instead.
func (BalanceType) EnumDescriptor() ([]byte, []int) {
	return file_transfer_request_proto_rawDescGZIP(), []int{2}
}

// Message representing a transfer request.
type TransferRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Transfer request ID
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Basis for the transfer
	Basis TransferBasis `protobuf:"varint,2,opt,name=basis,proto3,enum=proto.TransferBasis" json:"basis,omitempty"`
	// Administrator ID
	AdministratorId string `protobuf:"bytes,3,opt,name=administrator_id,json=administratorId,proto3" json:"administrator_id,omitempty"`
	// Document type
	DocumentType DocumentType `protobuf:"varint,4,opt,name=document_type,json=documentType,proto3,enum=proto.DocumentType" json:"document_type,omitempty"`
	// Document number
	DocumentNumber string `protobuf:"bytes,5,opt,name=document_number,json=documentNumber,proto3" json:"document_number,omitempty"`
	// Document date
	DocumentDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=document_date,json=documentDate,proto3" json:"document_date,omitempty"`
	// Hashes of attached documents
	DocumentHashes []string `protobuf:"bytes,7,rep,name=document_hashes,json=documentHashes,proto3" json:"document_hashes,omitempty"`
	// Address from which the transfer is made
	FromAddress string `protobuf:"bytes,8,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// Address to which the transfer is made
	ToAddress string `protobuf:"bytes,9,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	// Token being transferred
	Token string `protobuf:"bytes,10,opt,name=token,proto3" json:"token,omitempty"`
	// Amount being transferred
	Amount string `protobuf:"bytes,11,opt,name=amount,proto3" json:"amount,omitempty"`
	// Reason for the transfer
	Reason string `protobuf:"bytes,12,opt,name=reason,proto3" json:"reason,omitempty"`
	// Balance type from which the transfer is made
	BalanceType BalanceType `protobuf:"varint,13,opt,name=balance_type,json=balanceType,proto3,enum=proto.BalanceType" json:"balance_type,omitempty"`
	// Optional additional information
	AdditionalInfo *anypb.Any `protobuf:"bytes,15,opt,name=additional_info,json=additionalInfo,proto3" json:"additional_info,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	mi := &file_transfer_request_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfer_request_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_transfer_request_proto_rawDescGZIP(), []int{0}
}

func (x *TransferRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *TransferRequest) GetBasis() TransferBasis {
	if x != nil {
		return x.Basis
	}
	return TransferBasis_TRANSFER_BASIS_UNSPECIFIED
}

func (x *TransferRequest) GetAdministratorId() string {
	if x != nil {
		return x.AdministratorId
	}
	return ""
}

func (x *TransferRequest) GetDocumentType() DocumentType {
	if x != nil {
		return x.DocumentType
	}
	return DocumentType_DOCUMENT_TYPE_UNSPECIFIED
}

func (x *TransferRequest) GetDocumentNumber() string {
	if x != nil {
		return x.DocumentNumber
	}
	return ""
}

func (x *TransferRequest) GetDocumentDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DocumentDate
	}
	return nil
}

func (x *TransferRequest) GetDocumentHashes() []string {
	if x != nil {
		return x.DocumentHashes
	}
	return nil
}

func (x *TransferRequest) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *TransferRequest) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *TransferRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TransferRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *TransferRequest) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *TransferRequest) GetBalanceType() BalanceType {
	if x != nil {
		return x.BalanceType
	}
	return BalanceType_BALANCE_TYPE_UNSPECIFIED
}

func (x *TransferRequest) GetAdditionalInfo() *anypb.Any {
	if x != nil {
		return x.AdditionalInfo
	}
	return nil
}

var File_transfer_request_proto protoreflect.FileDescriptor

var file_transfer_request_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x05, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x61, 0x73, 0x69, 0x73, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x73, 0x12, 0x32, 0x0a,
	0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x0f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x42, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x0d, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xb2,
	0x01, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x31, 0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xfa, 0x42, 0x1b,
	0x72, 0x19, 0x32, 0x17, 0x5e, 0x5b, 0x31, 0x2d, 0x39, 0x41, 0x2d, 0x48, 0x4a, 0x2d, 0x4e, 0x50,
	0x2d, 0x5a, 0x61, 0x2d, 0x6b, 0x6d, 0x2d, 0x7a, 0x5d, 0x2b, 0x24, 0x52, 0x0b, 0x66, 0x72, 0x6f,
	0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xfa, 0x42,
	0x1b, 0x72, 0x19, 0x32, 0x17, 0x5e, 0x5b, 0x31, 0x2d, 0x39, 0x41, 0x2d, 0x48, 0x4a, 0x2d, 0x4e,
	0x50, 0x2d, 0x5a, 0x61, 0x2d, 0x6b, 0x6d, 0x2d, 0x7a, 0x5d, 0x2b, 0x24, 0x52, 0x09, 0x74, 0x6f,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x3f, 0x0a, 0x0c, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x0b, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x3d, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2a,
	0x81, 0x01, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x19, 0x44, 0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x17, 0x0a, 0x13, 0x44, 0x4f, 0x43, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x4f, 0x43, 0x55,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x48, 0x45, 0x52, 0x49,
	0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x4f, 0x43, 0x55, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x55, 0x44, 0x47, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x03, 0x2a, 0x95, 0x01, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x42, 0x61, 0x73, 0x69, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45,
	0x52, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45,
	0x52, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x52, 0x45, 0x4f, 0x52, 0x47, 0x41, 0x4e, 0x49,
	0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x46, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x49, 0x4e, 0x48, 0x45, 0x52,
	0x49, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x46, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x53, 0x49, 0x53, 0x5f, 0x43, 0x4f, 0x55, 0x52, 0x54,
	0x5f, 0x44, 0x45, 0x43, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x2a, 0x6b, 0x0a, 0x0b, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x42, 0x41,
	0x4c, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x41, 0x4c, 0x41,
	0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x2b,
	0x12, 0x26, 0x0a, 0x22, 0x42, 0x41, 0x4c, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x32, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x6f, 0x69, 0x64, 0x65, 0x61, 0x6f, 0x70,
	0x65, 0x6e, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transfer_request_proto_rawDescOnce sync.Once
	file_transfer_request_proto_rawDescData = file_transfer_request_proto_rawDesc
)

func file_transfer_request_proto_rawDescGZIP() []byte {
	file_transfer_request_proto_rawDescOnce.Do(func() {
		file_transfer_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfer_request_proto_rawDescData)
	})
	return file_transfer_request_proto_rawDescData
}

var file_transfer_request_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_transfer_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_transfer_request_proto_goTypes = []any{
	(DocumentType)(0),             // 0: proto.DocumentType
	(TransferBasis)(0),            // 1: proto.TransferBasis
	(BalanceType)(0),              // 2: proto.BalanceType
	(*TransferRequest)(nil),       // 3: proto.TransferRequest
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 5: google.protobuf.Any
}
var file_transfer_request_proto_depIdxs = []int32{
	1, // 0: proto.TransferRequest.basis:type_name -> proto.TransferBasis
	0, // 1: proto.TransferRequest.document_type:type_name -> proto.DocumentType
	4, // 2: proto.TransferRequest.document_date:type_name -> google.protobuf.Timestamp
	2, // 3: proto.TransferRequest.balance_type:type_name -> proto.BalanceType
	5, // 4: proto.TransferRequest.additional_info:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_transfer_request_proto_init() }
func file_transfer_request_proto_init() {
	if File_transfer_request_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transfer_request_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transfer_request_proto_goTypes,
		DependencyIndexes: file_transfer_request_proto_depIdxs,
		EnumInfos:         file_transfer_request_proto_enumTypes,
		MessageInfos:      file_transfer_request_proto_msgTypes,
	}.Build()
	File_transfer_request_proto = out.File
	file_transfer_request_proto_rawDesc = nil
	file_transfer_request_proto_goTypes = nil
	file_transfer_request_proto_depIdxs = nil
}
