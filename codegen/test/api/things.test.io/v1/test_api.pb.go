// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: github.com/solo-io/skv2/codegen/test/test_api.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/cue/encoding/protobuf/cue"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AcrylicType_Body int32

const (
	AcrylicType_Light  AcrylicType_Body = 0
	AcrylicType_Medium AcrylicType_Body = 1
	AcrylicType_Heavy  AcrylicType_Body = 2
)

// Enum value maps for AcrylicType_Body.
var (
	AcrylicType_Body_name = map[int32]string{
		0: "Light",
		1: "Medium",
		2: "Heavy",
	}
	AcrylicType_Body_value = map[string]int32{
		"Light":  0,
		"Medium": 1,
		"Heavy":  2,
	}
)

func (x AcrylicType_Body) Enum() *AcrylicType_Body {
	p := new(AcrylicType_Body)
	*p = x
	return p
}

func (x AcrylicType_Body) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AcrylicType_Body) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_enumTypes[0].Descriptor()
}

func (AcrylicType_Body) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_skv2_codegen_test_test_api_proto_enumTypes[0]
}

func (x AcrylicType_Body) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AcrylicType_Body.Descriptor instead.
func (AcrylicType_Body) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescGZIP(), []int{2, 0}
}

type PaintSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Color *PaintColor `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	// Types that are assignable to PaintType:
	//
	//	*PaintSpec_Acrylic
	//	*PaintSpec_Oil
	PaintType  isPaintSpec_PaintType `protobuf_oneof:"paintType"`
	MyFavorite *anypb.Any            `protobuf:"bytes,4,opt,name=my_favorite,json=myFavorite,proto3" json:"my_favorite,omitempty"`
	// OpenAPI gen test for recursive fields
	RecursiveType *PaintSpec_RecursiveType `protobuf:"bytes,5,opt,name=recursive_type,json=recursiveType,proto3" json:"recursive_type,omitempty"`
}

func (x *PaintSpec) Reset() {
	*x = PaintSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaintSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaintSpec) ProtoMessage() {}

func (x *PaintSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaintSpec.ProtoReflect.Descriptor instead.
func (*PaintSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescGZIP(), []int{0}
}

func (x *PaintSpec) GetColor() *PaintColor {
	if x != nil {
		return x.Color
	}
	return nil
}

func (m *PaintSpec) GetPaintType() isPaintSpec_PaintType {
	if m != nil {
		return m.PaintType
	}
	return nil
}

func (x *PaintSpec) GetAcrylic() *AcrylicType {
	if x, ok := x.GetPaintType().(*PaintSpec_Acrylic); ok {
		return x.Acrylic
	}
	return nil
}

func (x *PaintSpec) GetOil() *OilType {
	if x, ok := x.GetPaintType().(*PaintSpec_Oil); ok {
		return x.Oil
	}
	return nil
}

func (x *PaintSpec) GetMyFavorite() *anypb.Any {
	if x != nil {
		return x.MyFavorite
	}
	return nil
}

func (x *PaintSpec) GetRecursiveType() *PaintSpec_RecursiveType {
	if x != nil {
		return x.RecursiveType
	}
	return nil
}

type isPaintSpec_PaintType interface {
	isPaintSpec_PaintType()
}

type PaintSpec_Acrylic struct {
	Acrylic *AcrylicType `protobuf:"bytes,2,opt,name=acrylic,proto3,oneof"`
}

type PaintSpec_Oil struct {
	Oil *OilType `protobuf:"bytes,3,opt,name=oil,proto3,oneof"`
}

func (*PaintSpec_Acrylic) isPaintSpec_PaintType() {}

func (*PaintSpec_Oil) isPaintSpec_PaintType() {}

type PaintColor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hue   string  `protobuf:"bytes,1,opt,name=hue,proto3" json:"hue,omitempty"`
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PaintColor) Reset() {
	*x = PaintColor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaintColor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaintColor) ProtoMessage() {}

func (x *PaintColor) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaintColor.ProtoReflect.Descriptor instead.
func (*PaintColor) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescGZIP(), []int{1}
}

func (x *PaintColor) GetHue() string {
	if x != nil {
		return x.Hue
	}
	return ""
}

func (x *PaintColor) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type AcrylicType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body AcrylicType_Body `protobuf:"varint,3,opt,name=body,proto3,enum=things.test.io.AcrylicType_Body" json:"body,omitempty"`
}

func (x *AcrylicType) Reset() {
	*x = AcrylicType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcrylicType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcrylicType) ProtoMessage() {}

func (x *AcrylicType) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcrylicType.ProtoReflect.Descriptor instead.
func (*AcrylicType) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescGZIP(), []int{2}
}

func (x *AcrylicType) GetBody() AcrylicType_Body {
	if x != nil {
		return x.Body
	}
	return AcrylicType_Light
}

type OilType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WaterMixable bool `protobuf:"varint,1,opt,name=waterMixable,proto3" json:"waterMixable,omitempty"`
	// Types that are assignable to PigmentType:
	//
	//	*OilType_Powder
	//	*OilType_Fluid
	PigmentType isOilType_PigmentType `protobuf_oneof:"pigmentType"`
}

func (x *OilType) Reset() {
	*x = OilType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OilType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OilType) ProtoMessage() {}

func (x *OilType) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OilType.ProtoReflect.Descriptor instead.
func (*OilType) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescGZIP(), []int{3}
}

func (x *OilType) GetWaterMixable() bool {
	if x != nil {
		return x.WaterMixable
	}
	return false
}

func (m *OilType) GetPigmentType() isOilType_PigmentType {
	if m != nil {
		return m.PigmentType
	}
	return nil
}

func (x *OilType) GetPowder() string {
	if x, ok := x.GetPigmentType().(*OilType_Powder); ok {
		return x.Powder
	}
	return ""
}

func (x *OilType) GetFluid() string {
	if x, ok := x.GetPigmentType().(*OilType_Fluid); ok {
		return x.Fluid
	}
	return ""
}

type isOilType_PigmentType interface {
	isOilType_PigmentType()
}

type OilType_Powder struct {
	Powder string `protobuf:"bytes,2,opt,name=powder,proto3,oneof"`
}

type OilType_Fluid struct {
	Fluid string `protobuf:"bytes,3,opt,name=fluid,proto3,oneof"`
}

func (*OilType_Powder) isOilType_PigmentType() {}

func (*OilType_Fluid) isOilType_PigmentType() {}

type PaintStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObservedGeneration int64                            `protobuf:"varint,1,opt,name=observedGeneration,proto3" json:"observedGeneration,omitempty"`
	PercentRemaining   int64                            `protobuf:"varint,2,opt,name=percentRemaining,proto3" json:"percentRemaining,omitempty"`
	NearbyPaints       map[string]*PaintStatus_Location `protobuf:"bytes,3,rep,name=nearbyPaints,proto3" json:"nearbyPaints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PaintStatus) Reset() {
	*x = PaintStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaintStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaintStatus) ProtoMessage() {}

func (x *PaintStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaintStatus.ProtoReflect.Descriptor instead.
func (*PaintStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescGZIP(), []int{4}
}

func (x *PaintStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *PaintStatus) GetPercentRemaining() int64 {
	if x != nil {
		return x.PercentRemaining
	}
	return 0
}

func (x *PaintStatus) GetNearbyPaints() map[string]*PaintStatus_Location {
	if x != nil {
		return x.NearbyPaints
	}
	return nil
}

type ClusterResourceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imported *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=imported,proto3" json:"imported,omitempty"`
}

func (x *ClusterResourceSpec) Reset() {
	*x = ClusterResourceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterResourceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterResourceSpec) ProtoMessage() {}

func (x *ClusterResourceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterResourceSpec.ProtoReflect.Descriptor instead.
func (*ClusterResourceSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescGZIP(), []int{5}
}

func (x *ClusterResourceSpec) GetImported() *wrapperspb.StringValue {
	if x != nil {
		return x.Imported
	}
	return nil
}

// tests OpenAPI schema gen for Recursive types
type PaintSpec_RecursiveType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtobufValue          *structpb.Value            `protobuf:"bytes,4,opt,name=protobuf_value,json=protobufValue,proto3" json:"protobuf_value,omitempty"`
	RecursiveField         *PaintSpec_RecursiveType   `protobuf:"bytes,1,opt,name=recursive_field,json=recursiveField,proto3" json:"recursive_field,omitempty"`
	RepeatedRecursiveField []*PaintSpec_RecursiveType `protobuf:"bytes,2,rep,name=repeated_recursive_field,json=repeatedRecursiveField,proto3" json:"repeated_recursive_field,omitempty"`
	// Ensure that FieldOptions can be defined using package name resolution that starts from the
	// outermost scope: https://developers.google.com/protocol-buffers/docs/proto3#packages_and_name_resolution
	RecursiveFieldOutermostScope *PaintSpec_RecursiveType `protobuf:"bytes,3,opt,name=recursive_field_outermost_scope,json=recursiveFieldOutermostScope,proto3" json:"recursive_field_outermost_scope,omitempty"`
}

func (x *PaintSpec_RecursiveType) Reset() {
	*x = PaintSpec_RecursiveType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaintSpec_RecursiveType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaintSpec_RecursiveType) ProtoMessage() {}

func (x *PaintSpec_RecursiveType) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaintSpec_RecursiveType.ProtoReflect.Descriptor instead.
func (*PaintSpec_RecursiveType) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PaintSpec_RecursiveType) GetProtobufValue() *structpb.Value {
	if x != nil {
		return x.ProtobufValue
	}
	return nil
}

func (x *PaintSpec_RecursiveType) GetRecursiveField() *PaintSpec_RecursiveType {
	if x != nil {
		return x.RecursiveField
	}
	return nil
}

func (x *PaintSpec_RecursiveType) GetRepeatedRecursiveField() []*PaintSpec_RecursiveType {
	if x != nil {
		return x.RepeatedRecursiveField
	}
	return nil
}

func (x *PaintSpec_RecursiveType) GetRecursiveFieldOutermostScope() *PaintSpec_RecursiveType {
	if x != nil {
		return x.RecursiveFieldOutermostScope
	}
	return nil
}

type PaintStatus_Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X string `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
	Y string `protobuf:"bytes,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *PaintStatus_Location) Reset() {
	*x = PaintStatus_Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaintStatus_Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaintStatus_Location) ProtoMessage() {}

func (x *PaintStatus_Location) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaintStatus_Location.ProtoReflect.Descriptor instead.
func (*PaintStatus_Location) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescGZIP(), []int{4, 1}
}

func (x *PaintStatus_Location) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

func (x *PaintStatus_Location) GetY() string {
	if x != nil {
		return x.Y
	}
	return ""
}

var File_github_com_solo_io_skv2_codegen_test_test_api_proto protoreflect.FileDescriptor

var file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65,
	0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x69, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x63, 0x75, 0x65, 0x2f, 0x63, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc9, 0x05, 0x0a, 0x09, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x30, 0x0a,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x61,
	0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x37, 0x0a, 0x07, 0x61, 0x63, 0x72, 0x79, 0x6c, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x69,
	0x6f, 0x2e, 0x41, 0x63, 0x72, 0x79, 0x6c, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52,
	0x07, 0x61, 0x63, 0x72, 0x79, 0x6c, 0x69, 0x63, 0x12, 0x2b, 0x0a, 0x03, 0x6f, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00,
	0x52, 0x03, 0x6f, 0x69, 0x6c, 0x12, 0x35, 0x0a, 0x0b, 0x6d, 0x79, 0x5f, 0x66, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x0a, 0x6d, 0x79, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x4e, 0x0a, 0x0e,
	0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x2e,
	0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x72,
	0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x8f, 0x03, 0x0a,
	0x0d, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x44,
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x05,
	0xea, 0x42, 0x02, 0x20, 0x01, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x57, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x69, 0x6f, 0x2e, 0x50,
	0x61, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x05, 0xea, 0x42, 0x02, 0x10, 0x01, 0x52, 0x0e, 0x72,
	0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x68, 0x0a,
	0x18, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73,
	0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x69, 0x6f,
	0x2e, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x52, 0x65, 0x63, 0x75, 0x72,
	0x73, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x05, 0xea, 0x42, 0x02, 0x10, 0x01, 0x52,
	0x16, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69,
	0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x75, 0x0a, 0x1f, 0x72, 0x65, 0x63, 0x75, 0x72,
	0x73, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x69,
	0x6f, 0x2e, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x52, 0x65, 0x63, 0x75,
	0x72, 0x73, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x05, 0xea, 0x42, 0x02, 0x10, 0x01,
	0x52, 0x1c, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x75, 0x74, 0x65, 0x72, 0x6d, 0x6f, 0x73, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x42, 0x0b,
	0x0a, 0x09, 0x70, 0x61, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x34, 0x0a, 0x0a, 0x50,
	0x61, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x68, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x6d, 0x0a, 0x0b, 0x41, 0x63, 0x72, 0x79, 0x6c, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x34, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x69, 0x6f, 0x2e,
	0x41, 0x63, 0x72, 0x79, 0x6c, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x42, 0x6f, 0x64, 0x79,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x28, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x09,
	0x0a, 0x05, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x65, 0x64,
	0x69, 0x75, 0x6d, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x65, 0x61, 0x76, 0x79, 0x10, 0x02,
	0x22, 0x6e, 0x0a, 0x07, 0x4f, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x77,
	0x61, 0x74, 0x65, 0x72, 0x4d, 0x69, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x77, 0x61, 0x74, 0x65, 0x72, 0x4d, 0x69, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x06, 0x70, 0x6f, 0x77, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x70, 0x6f, 0x77, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x05, 0x66, 0x6c, 0x75,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x66, 0x6c, 0x75, 0x69,
	0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x70, 0x69, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x22, 0xd9, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2e, 0x0a, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x0a, 0x10, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x51, 0x0a, 0x0c,
	0x6e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0c, 0x6e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x1a,
	0x65, 0x0a, 0x11, 0x4e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x61, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x34, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xea,
	0x42, 0x02, 0x10, 0x01, 0x52, 0x01, 0x78, 0x12, 0x13, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x05, 0xea, 0x42, 0x02, 0x10, 0x01, 0x52, 0x01, 0x79, 0x22, 0x4f, 0x0a, 0x13,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x38, 0x0a, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x08, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x42, 0x3c, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescOnce sync.Once
	file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescData = file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDesc
)

func file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescData)
	})
	return file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDescData
}

var file_github_com_solo_io_skv2_codegen_test_test_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_github_com_solo_io_skv2_codegen_test_test_api_proto_goTypes = []interface{}{
	(AcrylicType_Body)(0),           // 0: things.test.io.AcrylicType.Body
	(*PaintSpec)(nil),               // 1: things.test.io.PaintSpec
	(*PaintColor)(nil),              // 2: things.test.io.PaintColor
	(*AcrylicType)(nil),             // 3: things.test.io.AcrylicType
	(*OilType)(nil),                 // 4: things.test.io.OilType
	(*PaintStatus)(nil),             // 5: things.test.io.PaintStatus
	(*ClusterResourceSpec)(nil),     // 6: things.test.io.ClusterResourceSpec
	(*PaintSpec_RecursiveType)(nil), // 7: things.test.io.PaintSpec.RecursiveType
	nil,                             // 8: things.test.io.PaintStatus.NearbyPaintsEntry
	(*PaintStatus_Location)(nil),    // 9: things.test.io.PaintStatus.Location
	(*anypb.Any)(nil),               // 10: google.protobuf.Any
	(*wrapperspb.StringValue)(nil),  // 11: google.protobuf.StringValue
	(*structpb.Value)(nil),          // 12: google.protobuf.Value
}
var file_github_com_solo_io_skv2_codegen_test_test_api_proto_depIdxs = []int32{
	2,  // 0: things.test.io.PaintSpec.color:type_name -> things.test.io.PaintColor
	3,  // 1: things.test.io.PaintSpec.acrylic:type_name -> things.test.io.AcrylicType
	4,  // 2: things.test.io.PaintSpec.oil:type_name -> things.test.io.OilType
	10, // 3: things.test.io.PaintSpec.my_favorite:type_name -> google.protobuf.Any
	7,  // 4: things.test.io.PaintSpec.recursive_type:type_name -> things.test.io.PaintSpec.RecursiveType
	0,  // 5: things.test.io.AcrylicType.body:type_name -> things.test.io.AcrylicType.Body
	8,  // 6: things.test.io.PaintStatus.nearbyPaints:type_name -> things.test.io.PaintStatus.NearbyPaintsEntry
	11, // 7: things.test.io.ClusterResourceSpec.imported:type_name -> google.protobuf.StringValue
	12, // 8: things.test.io.PaintSpec.RecursiveType.protobuf_value:type_name -> google.protobuf.Value
	7,  // 9: things.test.io.PaintSpec.RecursiveType.recursive_field:type_name -> things.test.io.PaintSpec.RecursiveType
	7,  // 10: things.test.io.PaintSpec.RecursiveType.repeated_recursive_field:type_name -> things.test.io.PaintSpec.RecursiveType
	7,  // 11: things.test.io.PaintSpec.RecursiveType.recursive_field_outermost_scope:type_name -> things.test.io.PaintSpec.RecursiveType
	9,  // 12: things.test.io.PaintStatus.NearbyPaintsEntry.value:type_name -> things.test.io.PaintStatus.Location
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_skv2_codegen_test_test_api_proto_init() }
func file_github_com_solo_io_skv2_codegen_test_test_api_proto_init() {
	if File_github_com_solo_io_skv2_codegen_test_test_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaintSpec); i {
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
		file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaintColor); i {
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
		file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcrylicType); i {
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
		file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OilType); i {
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
		file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaintStatus); i {
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
		file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterResourceSpec); i {
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
		file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaintSpec_RecursiveType); i {
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
		file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaintStatus_Location); i {
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
	file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PaintSpec_Acrylic)(nil),
		(*PaintSpec_Oil)(nil),
	}
	file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*OilType_Powder)(nil),
		(*OilType_Fluid)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_skv2_codegen_test_test_api_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_skv2_codegen_test_test_api_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_skv2_codegen_test_test_api_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_skv2_codegen_test_test_api_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_skv2_codegen_test_test_api_proto = out.File
	file_github_com_solo_io_skv2_codegen_test_test_api_proto_rawDesc = nil
	file_github_com_solo_io_skv2_codegen_test_test_api_proto_goTypes = nil
	file_github_com_solo_io_skv2_codegen_test_test_api_proto_depIdxs = nil
}
