// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: pipeline/pipeline.proto

package pipeline

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Producto    string `protobuf:"bytes,2,opt,name=Producto,proto3" json:"Producto,omitempty"`
	Valor       int32  `protobuf:"varint,3,opt,name=Valor,proto3" json:"Valor,omitempty"`
	Tienda      string `protobuf:"bytes,4,opt,name=Tienda,proto3" json:"Tienda,omitempty"`
	Destino     string `protobuf:"bytes,5,opt,name=Destino,proto3" json:"Destino,omitempty"`
	Prioridad   int32  `protobuf:"varint,6,opt,name=Prioridad,proto3" json:"Prioridad,omitempty"`
	Seguimiento int32  `protobuf:"varint,7,opt,name=Seguimiento,proto3" json:"Seguimiento,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pipeline_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *Message) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Message) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *Message) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *Message) GetPrioridad() int32 {
	if x != nil {
		return x.Prioridad
	}
	return 0
}

func (x *Message) GetSeguimiento() int32 {
	if x != nil {
		return x.Seguimiento
	}
	return 0
}

type RespuestaCon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Producto      string `protobuf:"bytes,2,opt,name=Producto,proto3" json:"Producto,omitempty"`
	Valor         int32  `protobuf:"varint,3,opt,name=Valor,proto3" json:"Valor,omitempty"`
	Tienda        string `protobuf:"bytes,4,opt,name=Tienda,proto3" json:"Tienda,omitempty"`
	Destino       string `protobuf:"bytes,5,opt,name=Destino,proto3" json:"Destino,omitempty"`
	Prioridad     int32  `protobuf:"varint,6,opt,name=Prioridad,proto3" json:"Prioridad,omitempty"`
	Seguimiento   int32  `protobuf:"varint,7,opt,name=Seguimiento,proto3" json:"Seguimiento,omitempty"`
	Intentos      int32  `protobuf:"varint,8,opt,name=Intentos,proto3" json:"Intentos,omitempty"`
	Estado        int32  `protobuf:"varint,9,opt,name=Estado,proto3" json:"Estado,omitempty"`
	IdCamion      int32  `protobuf:"varint,10,opt,name=IdCamion,proto3" json:"IdCamion,omitempty"`
	TiempoEntrega string `protobuf:"bytes,11,opt,name=TiempoEntrega,proto3" json:"TiempoEntrega,omitempty"`
}

func (x *RespuestaCon) Reset() {
	*x = RespuestaCon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaCon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaCon) ProtoMessage() {}

func (x *RespuestaCon) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaCon.ProtoReflect.Descriptor instead.
func (*RespuestaCon) Descriptor() ([]byte, []int) {
	return file_pipeline_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *RespuestaCon) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RespuestaCon) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *RespuestaCon) GetValor() int32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *RespuestaCon) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *RespuestaCon) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *RespuestaCon) GetPrioridad() int32 {
	if x != nil {
		return x.Prioridad
	}
	return 0
}

func (x *RespuestaCon) GetSeguimiento() int32 {
	if x != nil {
		return x.Seguimiento
	}
	return 0
}

func (x *RespuestaCon) GetIntentos() int32 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *RespuestaCon) GetEstado() int32 {
	if x != nil {
		return x.Estado
	}
	return 0
}

func (x *RespuestaCon) GetIdCamion() int32 {
	if x != nil {
		return x.IdCamion
	}
	return 0
}

func (x *RespuestaCon) GetTiempoEntrega() string {
	if x != nil {
		return x.TiempoEntrega
	}
	return ""
}

type ActCamion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seguimiento int32 `protobuf:"varint,1,opt,name=Seguimiento,proto3" json:"Seguimiento,omitempty"`
	Exito       int32 `protobuf:"varint,2,opt,name=Exito,proto3" json:"Exito,omitempty"`
}

func (x *ActCamion) Reset() {
	*x = ActCamion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_pipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActCamion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActCamion) ProtoMessage() {}

func (x *ActCamion) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_pipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActCamion.ProtoReflect.Descriptor instead.
func (*ActCamion) Descriptor() ([]byte, []int) {
	return file_pipeline_pipeline_proto_rawDescGZIP(), []int{2}
}

func (x *ActCamion) GetSeguimiento() int32 {
	if x != nil {
		return x.Seguimiento
	}
	return 0
}

func (x *ActCamion) GetExito() int32 {
	if x != nil {
		return x.Exito
	}
	return 0
}

type ConsultaEstado struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seguimiento int32 `protobuf:"varint,1,opt,name=Seguimiento,proto3" json:"Seguimiento,omitempty"`
}

func (x *ConsultaEstado) Reset() {
	*x = ConsultaEstado{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_pipeline_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultaEstado) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultaEstado) ProtoMessage() {}

func (x *ConsultaEstado) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_pipeline_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultaEstado.ProtoReflect.Descriptor instead.
func (*ConsultaEstado) Descriptor() ([]byte, []int) {
	return file_pipeline_pipeline_proto_rawDescGZIP(), []int{3}
}

func (x *ConsultaEstado) GetSeguimiento() int32 {
	if x != nil {
		return x.Seguimiento
	}
	return 0
}

type Solcamion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdCamion int32 `protobuf:"varint,1,opt,name=IdCamion,proto3" json:"IdCamion,omitempty"`
}

func (x *Solcamion) Reset() {
	*x = Solcamion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_pipeline_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Solcamion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Solcamion) ProtoMessage() {}

func (x *Solcamion) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_pipeline_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Solcamion.ProtoReflect.Descriptor instead.
func (*Solcamion) Descriptor() ([]byte, []int) {
	return file_pipeline_pipeline_proto_rawDescGZIP(), []int{4}
}

func (x *Solcamion) GetIdCamion() int32 {
	if x != nil {
		return x.IdCamion
	}
	return 0
}

var File_pipeline_pipeline_proto protoreflect.FileDescriptor

var file_pipeline_pipeline_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x6f,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x54, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x64, 0x61, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x64, 0x61,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65,
	0x6e, 0x74, 0x6f, 0x22, 0xb8, 0x02, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x43, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x69, 0x65, 0x6e, 0x64, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x64, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x64, 0x61, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d,
	0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x53, 0x65, 0x67,
	0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x6f, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x49, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x49, 0x64, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x49, 0x64, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x69, 0x65, 0x6d,
	0x70, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x54, 0x69, 0x65, 0x6d, 0x70, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x22, 0x43,
	0x0a, 0x09, 0x41, 0x63, 0x74, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x53,
	0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x78, 0x69, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x45, 0x78,
	0x69, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x45,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69,
	0x65, 0x6e, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x53, 0x65, 0x67, 0x75,
	0x69, 0x6d, 0x69, 0x65, 0x6e, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x09, 0x53, 0x6f, 0x6c, 0x63, 0x61,
	0x6d, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x64, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x49, 0x64, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e,
	0x32, 0xf9, 0x01, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x08,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x11, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x18, 0x2e,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74,
	0x61, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x1a, 0x16, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x09, 0x53, 0x6f, 0x6c, 0x70, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x12, 0x13,
	0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x6f, 0x6c, 0x63, 0x61, 0x6d,
	0x69, 0x6f, 0x6e, 0x1a, 0x16, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0a, 0x41, 0x63, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x12, 0x13, 0x2e, 0x70, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x43, 0x61, 0x6d, 0x69, 0x6f, 0x6e,
	0x1a, 0x18, 0x2e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x74, 0x61, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10,
	0x4c, 0x61, 0x62, 0x31, 0x2f, 0x53, 0x44, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pipeline_pipeline_proto_rawDescOnce sync.Once
	file_pipeline_pipeline_proto_rawDescData = file_pipeline_pipeline_proto_rawDesc
)

func file_pipeline_pipeline_proto_rawDescGZIP() []byte {
	file_pipeline_pipeline_proto_rawDescOnce.Do(func() {
		file_pipeline_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_pipeline_pipeline_proto_rawDescData)
	})
	return file_pipeline_pipeline_proto_rawDescData
}

var file_pipeline_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pipeline_pipeline_proto_goTypes = []interface{}{
	(*Message)(nil),        // 0: pipeline.Message
	(*RespuestaCon)(nil),   // 1: pipeline.RespuestaCon
	(*ActCamion)(nil),      // 2: pipeline.ActCamion
	(*ConsultaEstado)(nil), // 3: pipeline.ConsultaEstado
	(*Solcamion)(nil),      // 4: pipeline.Solcamion
}
var file_pipeline_pipeline_proto_depIdxs = []int32{
	0, // 0: pipeline.Greeter.SayHello:input_type -> pipeline.Message
	3, // 1: pipeline.Greeter.ConEstado:input_type -> pipeline.ConsultaEstado
	4, // 2: pipeline.Greeter.Solpedido:input_type -> pipeline.Solcamion
	2, // 3: pipeline.Greeter.ActEntrega:input_type -> pipeline.ActCamion
	0, // 4: pipeline.Greeter.SayHello:output_type -> pipeline.Message
	1, // 5: pipeline.Greeter.ConEstado:output_type -> pipeline.RespuestaCon
	1, // 6: pipeline.Greeter.Solpedido:output_type -> pipeline.RespuestaCon
	3, // 7: pipeline.Greeter.ActEntrega:output_type -> pipeline.ConsultaEstado
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pipeline_pipeline_proto_init() }
func file_pipeline_pipeline_proto_init() {
	if File_pipeline_pipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pipeline_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_pipeline_pipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaCon); i {
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
		file_pipeline_pipeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActCamion); i {
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
		file_pipeline_pipeline_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultaEstado); i {
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
		file_pipeline_pipeline_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Solcamion); i {
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
			RawDescriptor: file_pipeline_pipeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pipeline_pipeline_proto_goTypes,
		DependencyIndexes: file_pipeline_pipeline_proto_depIdxs,
		MessageInfos:      file_pipeline_pipeline_proto_msgTypes,
	}.Build()
	File_pipeline_pipeline_proto = out.File
	file_pipeline_pipeline_proto_rawDesc = nil
	file_pipeline_pipeline_proto_goTypes = nil
	file_pipeline_pipeline_proto_depIdxs = nil
}
