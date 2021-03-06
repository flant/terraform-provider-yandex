// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/serverless/triggers/v1/predicate.proto

package triggers

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Predicate struct {
	// Types that are valid to be assigned to Predicate:
	//	*Predicate_AndPredicate
	//	*Predicate_FieldValuePredicate
	Predicate            isPredicate_Predicate `protobuf_oneof:"predicate"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Predicate) Reset()         { *m = Predicate{} }
func (m *Predicate) String() string { return proto.CompactTextString(m) }
func (*Predicate) ProtoMessage()    {}
func (*Predicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7702d63b0613d55f, []int{0}
}

func (m *Predicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Predicate.Unmarshal(m, b)
}
func (m *Predicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Predicate.Marshal(b, m, deterministic)
}
func (m *Predicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Predicate.Merge(m, src)
}
func (m *Predicate) XXX_Size() int {
	return xxx_messageInfo_Predicate.Size(m)
}
func (m *Predicate) XXX_DiscardUnknown() {
	xxx_messageInfo_Predicate.DiscardUnknown(m)
}

var xxx_messageInfo_Predicate proto.InternalMessageInfo

type isPredicate_Predicate interface {
	isPredicate_Predicate()
}

type Predicate_AndPredicate struct {
	AndPredicate *AndPredicate `protobuf:"bytes,2,opt,name=and_predicate,json=andPredicate,proto3,oneof"`
}

type Predicate_FieldValuePredicate struct {
	FieldValuePredicate *FieldValuePredicate `protobuf:"bytes,4,opt,name=field_value_predicate,json=fieldValuePredicate,proto3,oneof"`
}

func (*Predicate_AndPredicate) isPredicate_Predicate() {}

func (*Predicate_FieldValuePredicate) isPredicate_Predicate() {}

func (m *Predicate) GetPredicate() isPredicate_Predicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func (m *Predicate) GetAndPredicate() *AndPredicate {
	if x, ok := m.GetPredicate().(*Predicate_AndPredicate); ok {
		return x.AndPredicate
	}
	return nil
}

func (m *Predicate) GetFieldValuePredicate() *FieldValuePredicate {
	if x, ok := m.GetPredicate().(*Predicate_FieldValuePredicate); ok {
		return x.FieldValuePredicate
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Predicate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Predicate_AndPredicate)(nil),
		(*Predicate_FieldValuePredicate)(nil),
	}
}

type AndPredicate struct {
	Predicate            []*Predicate `protobuf:"bytes,1,rep,name=predicate,proto3" json:"predicate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AndPredicate) Reset()         { *m = AndPredicate{} }
func (m *AndPredicate) String() string { return proto.CompactTextString(m) }
func (*AndPredicate) ProtoMessage()    {}
func (*AndPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7702d63b0613d55f, []int{1}
}

func (m *AndPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AndPredicate.Unmarshal(m, b)
}
func (m *AndPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AndPredicate.Marshal(b, m, deterministic)
}
func (m *AndPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AndPredicate.Merge(m, src)
}
func (m *AndPredicate) XXX_Size() int {
	return xxx_messageInfo_AndPredicate.Size(m)
}
func (m *AndPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_AndPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_AndPredicate proto.InternalMessageInfo

func (m *AndPredicate) GetPredicate() []*Predicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

type FieldValuePredicate struct {
	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*FieldValuePredicate_Exact
	//	*FieldValuePredicate_Prefix
	//	*FieldValuePredicate_Suffix
	Value                isFieldValuePredicate_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *FieldValuePredicate) Reset()         { *m = FieldValuePredicate{} }
func (m *FieldValuePredicate) String() string { return proto.CompactTextString(m) }
func (*FieldValuePredicate) ProtoMessage()    {}
func (*FieldValuePredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7702d63b0613d55f, []int{2}
}

func (m *FieldValuePredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldValuePredicate.Unmarshal(m, b)
}
func (m *FieldValuePredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldValuePredicate.Marshal(b, m, deterministic)
}
func (m *FieldValuePredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValuePredicate.Merge(m, src)
}
func (m *FieldValuePredicate) XXX_Size() int {
	return xxx_messageInfo_FieldValuePredicate.Size(m)
}
func (m *FieldValuePredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValuePredicate.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValuePredicate proto.InternalMessageInfo

func (m *FieldValuePredicate) GetFieldPath() string {
	if m != nil {
		return m.FieldPath
	}
	return ""
}

type isFieldValuePredicate_Value interface {
	isFieldValuePredicate_Value()
}

type FieldValuePredicate_Exact struct {
	Exact string `protobuf:"bytes,3,opt,name=exact,proto3,oneof"`
}

type FieldValuePredicate_Prefix struct {
	Prefix string `protobuf:"bytes,8,opt,name=prefix,proto3,oneof"`
}

type FieldValuePredicate_Suffix struct {
	Suffix string `protobuf:"bytes,9,opt,name=suffix,proto3,oneof"`
}

func (*FieldValuePredicate_Exact) isFieldValuePredicate_Value() {}

func (*FieldValuePredicate_Prefix) isFieldValuePredicate_Value() {}

func (*FieldValuePredicate_Suffix) isFieldValuePredicate_Value() {}

func (m *FieldValuePredicate) GetValue() isFieldValuePredicate_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *FieldValuePredicate) GetExact() string {
	if x, ok := m.GetValue().(*FieldValuePredicate_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *FieldValuePredicate) GetPrefix() string {
	if x, ok := m.GetValue().(*FieldValuePredicate_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *FieldValuePredicate) GetSuffix() string {
	if x, ok := m.GetValue().(*FieldValuePredicate_Suffix); ok {
		return x.Suffix
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FieldValuePredicate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FieldValuePredicate_Exact)(nil),
		(*FieldValuePredicate_Prefix)(nil),
		(*FieldValuePredicate_Suffix)(nil),
	}
}

func init() {
	proto.RegisterType((*Predicate)(nil), "yandex.cloud.serverless.triggers.v1.Predicate")
	proto.RegisterType((*AndPredicate)(nil), "yandex.cloud.serverless.triggers.v1.AndPredicate")
	proto.RegisterType((*FieldValuePredicate)(nil), "yandex.cloud.serverless.triggers.v1.FieldValuePredicate")
}

func init() {
	proto.RegisterFile("yandex/cloud/serverless/triggers/v1/predicate.proto", fileDescriptor_7702d63b0613d55f)
}

var fileDescriptor_7702d63b0613d55f = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xbb, 0xf6, 0x0f, 0x66, 0xdb, 0x1e, 0xdc, 0xa2, 0x04, 0x41, 0x28, 0xed, 0xc1, 0x5e,
	0xba, 0x21, 0xed, 0x45, 0xf0, 0x64, 0x0e, 0xd2, 0x83, 0x87, 0x92, 0x83, 0x88, 0x08, 0x65, 0x9b,
	0xdd, 0xa4, 0x0b, 0x31, 0x09, 0x9b, 0x4d, 0xa8, 0xf8, 0x1c, 0xbe, 0x8f, 0x9e, 0x7c, 0x95, 0x3e,
	0x86, 0x24, 0xdb, 0xa4, 0x29, 0x2a, 0xe4, 0x96, 0x99, 0x2f, 0xdf, 0xef, 0x9b, 0x59, 0x06, 0xce,
	0xdf, 0x48, 0x40, 0xd9, 0xd6, 0x70, 0xfc, 0x30, 0xa1, 0x46, 0xcc, 0x44, 0xca, 0x84, 0xcf, 0xe2,
	0xd8, 0x90, 0x82, 0x7b, 0x1e, 0x13, 0xb1, 0x91, 0x9a, 0x46, 0x24, 0x18, 0xe5, 0x0e, 0x91, 0x0c,
	0x47, 0x22, 0x94, 0x21, 0x1a, 0x2b, 0x13, 0xce, 0x4d, 0xf8, 0x60, 0xc2, 0x85, 0x09, 0xa7, 0xe6,
	0xe5, 0xd5, 0x11, 0x39, 0x25, 0x3e, 0xa7, 0x44, 0xf2, 0x30, 0x50, 0x8c, 0xd1, 0x0e, 0x40, 0x6d,
	0x59, 0x70, 0xd1, 0x13, 0xec, 0x93, 0x80, 0xae, 0xca, 0x20, 0xfd, 0x64, 0x08, 0x26, 0xdd, 0x99,
	0x89, 0x6b, 0x24, 0xe1, 0xbb, 0x80, 0x96, 0xa4, 0x45, 0xc3, 0xee, 0x91, 0x4a, 0x8d, 0x02, 0x78,
	0xee, 0x72, 0xe6, 0xd3, 0x55, 0x4a, 0xfc, 0x84, 0x55, 0x12, 0x5a, 0x79, 0xc2, 0x4d, 0xad, 0x84,
	0xfb, 0x8c, 0xf0, 0x98, 0x01, 0xaa, 0x41, 0x03, 0xf7, 0x77, 0xdb, 0x3a, 0x83, 0x5a, 0x99, 0x81,
	0x5a, 0x9f, 0x5f, 0x26, 0x18, 0xbd, 0xc0, 0x5e, 0x75, 0x44, 0xf4, 0x50, 0xf9, 0x45, 0x07, 0xc3,
	0xe6, 0xa4, 0x3b, 0xc3, 0xb5, 0xc6, 0x28, 0x11, 0xf6, 0x01, 0x30, 0xfa, 0x00, 0x70, 0xf0, 0xc7,
	0x7c, 0x68, 0x0c, 0xa1, 0x5a, 0x3c, 0x22, 0x72, 0xa3, 0x83, 0x21, 0x98, 0x68, 0x56, 0x6b, 0xf7,
	0x6d, 0x02, 0x5b, 0xcb, 0xfb, 0x4b, 0x22, 0x37, 0xe8, 0x02, 0xb6, 0xd9, 0x96, 0x38, 0x52, 0x6f,
	0x66, 0xfa, 0xa2, 0x61, 0xab, 0x12, 0xe9, 0xb0, 0x13, 0x09, 0xe6, 0xf2, 0xad, 0x7e, 0xba, 0x17,
	0xf6, 0x75, 0xa6, 0xc4, 0x89, 0x9b, 0x29, 0x5a, 0xa1, 0xa8, 0xda, 0xea, 0xc3, 0x76, 0xfe, 0xc6,
	0x6a, 0x6b, 0xeb, 0x1d, 0x5e, 0x1f, 0xed, 0x44, 0x22, 0xfe, 0xcf, 0x5e, 0xcf, 0x4b, 0x8f, 0xcb,
	0x4d, 0xb2, 0xc6, 0x4e, 0xf8, 0x6a, 0x28, 0xcf, 0x54, 0x5d, 0x8d, 0x17, 0x4e, 0x3d, 0x16, 0xe4,
	0x07, 0x63, 0xd4, 0x38, 0xd4, 0xdb, 0xe2, 0x7b, 0xdd, 0xc9, 0x3d, 0xf3, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x24, 0xa1, 0xc4, 0x2b, 0xdf, 0x02, 0x00, 0x00,
}
