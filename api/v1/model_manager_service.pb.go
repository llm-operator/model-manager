// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: api/v1/model_manager_service.proto

package v1

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

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created int64  `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Object  string `protobuf:"bytes,3,opt,name=object,proto3" json:"object,omitempty"`
	OwnedBy string `protobuf:"bytes,4,opt,name=owned_by,json=ownedBy,proto3" json:"owned_by,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_model_manager_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_model_manager_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_api_v1_model_manager_service_proto_rawDescGZIP(), []int{0}
}

func (x *Model) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Model) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Model) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *Model) GetOwnedBy() string {
	if x != nil {
		return x.OwnedBy
	}
	return ""
}

type ListModelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object string   `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Model  []*Model `protobuf:"bytes,2,rep,name=model,proto3" json:"model,omitempty"`
}

func (x *ListModelsRequest) Reset() {
	*x = ListModelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_model_manager_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModelsRequest) ProtoMessage() {}

func (x *ListModelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_model_manager_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModelsRequest.ProtoReflect.Descriptor instead.
func (*ListModelsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_model_manager_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListModelsRequest) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ListModelsRequest) GetModel() []*Model {
	if x != nil {
		return x.Model
	}
	return nil
}

type ListModelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object string   `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Data   []*Model `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListModelsResponse) Reset() {
	*x = ListModelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_model_manager_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListModelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListModelsResponse) ProtoMessage() {}

func (x *ListModelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_model_manager_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListModelsResponse.ProtoReflect.Descriptor instead.
func (*ListModelsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_model_manager_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListModelsResponse) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ListModelsResponse) GetData() []*Model {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetModelRequest) Reset() {
	*x = GetModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_model_manager_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetModelRequest) ProtoMessage() {}

func (x *GetModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_model_manager_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetModelRequest.ProtoReflect.Descriptor instead.
func (*GetModelRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_model_manager_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetModelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteModelRequest) Reset() {
	*x = DeleteModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_model_manager_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteModelRequest) ProtoMessage() {}

func (x *DeleteModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_model_manager_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteModelRequest.ProtoReflect.Descriptor instead.
func (*DeleteModelRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_model_manager_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteModelRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Object  string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Deleted bool   `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteModelResponse) Reset() {
	*x = DeleteModelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_model_manager_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteModelResponse) ProtoMessage() {}

func (x *DeleteModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_model_manager_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteModelResponse.ProtoReflect.Descriptor instead.
func (*DeleteModelResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_model_manager_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteModelResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteModelResponse) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *DeleteModelResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type CreateModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseModel string `protobuf:"bytes,1,opt,name=base_model,json=baseModel,proto3" json:"base_model,omitempty"`
	Suffix    string `protobuf:"bytes,2,opt,name=suffix,proto3" json:"suffix,omitempty"`
	// TODO(kenji): Fix this.
	ModelPath string `protobuf:"bytes,3,opt,name=model_path,json=modelPath,proto3" json:"model_path,omitempty"`
	TenantId  string `protobuf:"bytes,4,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
}

func (x *CreateModelRequest) Reset() {
	*x = CreateModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_model_manager_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateModelRequest) ProtoMessage() {}

func (x *CreateModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_model_manager_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateModelRequest.ProtoReflect.Descriptor instead.
func (*CreateModelRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_model_manager_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateModelRequest) GetBaseModel() string {
	if x != nil {
		return x.BaseModel
	}
	return ""
}

func (x *CreateModelRequest) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

func (x *CreateModelRequest) GetModelPath() string {
	if x != nil {
		return x.ModelPath
	}
	return ""
}

func (x *CreateModelRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

var File_api_v1_model_manager_service_proto protoreflect.FileDescriptor

var file_api_v1_model_manager_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x64, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x77, 0x6e, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x22, 0x66, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x39, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x65,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6c, 0x6c, 0x6d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x57,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x32, 0x9c, 0x03, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x12, 0x2f, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x12, 0x77, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x2d, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x12, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x8b, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x30, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x32, 0x7f, 0x0a, 0x15, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x30, 0x2e, 0x6c, 0x6c, 0x6d, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6c, 0x6c, 0x6d,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x6c, 0x6d, 0x2d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_model_manager_service_proto_rawDescOnce sync.Once
	file_api_v1_model_manager_service_proto_rawDescData = file_api_v1_model_manager_service_proto_rawDesc
)

func file_api_v1_model_manager_service_proto_rawDescGZIP() []byte {
	file_api_v1_model_manager_service_proto_rawDescOnce.Do(func() {
		file_api_v1_model_manager_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_model_manager_service_proto_rawDescData)
	})
	return file_api_v1_model_manager_service_proto_rawDescData
}

var file_api_v1_model_manager_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_v1_model_manager_service_proto_goTypes = []interface{}{
	(*Model)(nil),               // 0: llmoperator.models.server.v1.Model
	(*ListModelsRequest)(nil),   // 1: llmoperator.models.server.v1.ListModelsRequest
	(*ListModelsResponse)(nil),  // 2: llmoperator.models.server.v1.ListModelsResponse
	(*GetModelRequest)(nil),     // 3: llmoperator.models.server.v1.GetModelRequest
	(*DeleteModelRequest)(nil),  // 4: llmoperator.models.server.v1.DeleteModelRequest
	(*DeleteModelResponse)(nil), // 5: llmoperator.models.server.v1.DeleteModelResponse
	(*CreateModelRequest)(nil),  // 6: llmoperator.models.server.v1.CreateModelRequest
}
var file_api_v1_model_manager_service_proto_depIdxs = []int32{
	0, // 0: llmoperator.models.server.v1.ListModelsRequest.model:type_name -> llmoperator.models.server.v1.Model
	0, // 1: llmoperator.models.server.v1.ListModelsResponse.data:type_name -> llmoperator.models.server.v1.Model
	1, // 2: llmoperator.models.server.v1.ModelsService.ListModels:input_type -> llmoperator.models.server.v1.ListModelsRequest
	3, // 3: llmoperator.models.server.v1.ModelsService.GetModel:input_type -> llmoperator.models.server.v1.GetModelRequest
	4, // 4: llmoperator.models.server.v1.ModelsService.DeleteModel:input_type -> llmoperator.models.server.v1.DeleteModelRequest
	6, // 5: llmoperator.models.server.v1.ModelsInternalService.CreateModel:input_type -> llmoperator.models.server.v1.CreateModelRequest
	2, // 6: llmoperator.models.server.v1.ModelsService.ListModels:output_type -> llmoperator.models.server.v1.ListModelsResponse
	0, // 7: llmoperator.models.server.v1.ModelsService.GetModel:output_type -> llmoperator.models.server.v1.Model
	5, // 8: llmoperator.models.server.v1.ModelsService.DeleteModel:output_type -> llmoperator.models.server.v1.DeleteModelResponse
	0, // 9: llmoperator.models.server.v1.ModelsInternalService.CreateModel:output_type -> llmoperator.models.server.v1.Model
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_v1_model_manager_service_proto_init() }
func file_api_v1_model_manager_service_proto_init() {
	if File_api_v1_model_manager_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_model_manager_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_api_v1_model_manager_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModelsRequest); i {
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
		file_api_v1_model_manager_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListModelsResponse); i {
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
		file_api_v1_model_manager_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetModelRequest); i {
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
		file_api_v1_model_manager_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteModelRequest); i {
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
		file_api_v1_model_manager_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteModelResponse); i {
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
		file_api_v1_model_manager_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateModelRequest); i {
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
			RawDescriptor: file_api_v1_model_manager_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_v1_model_manager_service_proto_goTypes,
		DependencyIndexes: file_api_v1_model_manager_service_proto_depIdxs,
		MessageInfos:      file_api_v1_model_manager_service_proto_msgTypes,
	}.Build()
	File_api_v1_model_manager_service_proto = out.File
	file_api_v1_model_manager_service_proto_rawDesc = nil
	file_api_v1_model_manager_service_proto_goTypes = nil
	file_api_v1_model_manager_service_proto_depIdxs = nil
}
