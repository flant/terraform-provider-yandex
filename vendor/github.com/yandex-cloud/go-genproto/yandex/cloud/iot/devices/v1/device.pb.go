// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/iot/devices/v1/device.proto

package devices

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Device struct {
	Id          string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RegistryId  string               `protobuf:"bytes,2,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	CreatedAt   *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name        string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// map from alias to canonical topic name prefix, e.g. my/custom/alias -> $device/abcdef/events
	TopicAliases         map[string]string `protobuf:"bytes,6,rep,name=topic_aliases,json=topicAliases,proto3" json:"topic_aliases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5d5b38a6ba5ae6, []int{0}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Device) GetRegistryId() string {
	if m != nil {
		return m.RegistryId
	}
	return ""
}

func (m *Device) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Device) GetTopicAliases() map[string]string {
	if m != nil {
		return m.TopicAliases
	}
	return nil
}

type DeviceCertificate struct {
	DeviceId             string               `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Fingerprint          string               `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	CertificateData      string               `protobuf:"bytes,3,opt,name=certificate_data,json=certificateData,proto3" json:"certificate_data,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeviceCertificate) Reset()         { *m = DeviceCertificate{} }
func (m *DeviceCertificate) String() string { return proto.CompactTextString(m) }
func (*DeviceCertificate) ProtoMessage()    {}
func (*DeviceCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5d5b38a6ba5ae6, []int{1}
}

func (m *DeviceCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceCertificate.Unmarshal(m, b)
}
func (m *DeviceCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceCertificate.Marshal(b, m, deterministic)
}
func (m *DeviceCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceCertificate.Merge(m, src)
}
func (m *DeviceCertificate) XXX_Size() int {
	return xxx_messageInfo_DeviceCertificate.Size(m)
}
func (m *DeviceCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceCertificate proto.InternalMessageInfo

func (m *DeviceCertificate) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DeviceCertificate) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *DeviceCertificate) GetCertificateData() string {
	if m != nil {
		return m.CertificateData
	}
	return ""
}

func (m *DeviceCertificate) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type DevicePassword struct {
	DeviceId             string               `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DevicePassword) Reset()         { *m = DevicePassword{} }
func (m *DevicePassword) String() string { return proto.CompactTextString(m) }
func (*DevicePassword) ProtoMessage()    {}
func (*DevicePassword) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc5d5b38a6ba5ae6, []int{2}
}

func (m *DevicePassword) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevicePassword.Unmarshal(m, b)
}
func (m *DevicePassword) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevicePassword.Marshal(b, m, deterministic)
}
func (m *DevicePassword) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevicePassword.Merge(m, src)
}
func (m *DevicePassword) XXX_Size() int {
	return xxx_messageInfo_DevicePassword.Size(m)
}
func (m *DevicePassword) XXX_DiscardUnknown() {
	xxx_messageInfo_DevicePassword.DiscardUnknown(m)
}

var xxx_messageInfo_DevicePassword proto.InternalMessageInfo

func (m *DevicePassword) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DevicePassword) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DevicePassword) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Device)(nil), "yandex.cloud.iot.devices.v1.Device")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.iot.devices.v1.Device.TopicAliasesEntry")
	proto.RegisterType((*DeviceCertificate)(nil), "yandex.cloud.iot.devices.v1.DeviceCertificate")
	proto.RegisterType((*DevicePassword)(nil), "yandex.cloud.iot.devices.v1.DevicePassword")
}

func init() {
	proto.RegisterFile("yandex/cloud/iot/devices/v1/device.proto", fileDescriptor_dc5d5b38a6ba5ae6)
}

var fileDescriptor_dc5d5b38a6ba5ae6 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9d, 0x34, 0x22, 0x13, 0x28, 0xed, 0x8a, 0x83, 0x95, 0x0a, 0x25, 0xca, 0x29, 0x1c,
	0xba, 0x56, 0x8b, 0x90, 0xf8, 0x38, 0xa0, 0x42, 0x11, 0xea, 0x0d, 0x45, 0x3d, 0xf5, 0x62, 0x6d,
	0xbc, 0x13, 0x33, 0xe0, 0x78, 0xad, 0xdd, 0x89, 0x69, 0xfe, 0x19, 0xff, 0x80, 0xbf, 0x85, 0xb2,
	0xeb, 0x94, 0xb4, 0x48, 0x91, 0xe0, 0x36, 0x7e, 0xfb, 0xe6, 0xed, 0xdb, 0x37, 0x63, 0x98, 0xae,
	0x55, 0xa5, 0xf1, 0x36, 0xcd, 0x4b, 0xb3, 0xd2, 0x29, 0x19, 0x4e, 0x35, 0x36, 0x94, 0xa3, 0x4b,
	0x9b, 0xb3, 0xb6, 0x94, 0xb5, 0x35, 0x6c, 0xc4, 0x49, 0x60, 0x4a, 0xcf, 0x94, 0x64, 0x58, 0xb6,
	0x4c, 0xd9, 0x9c, 0x0d, 0x47, 0x85, 0x31, 0x45, 0x89, 0xa9, 0xa7, 0xce, 0x57, 0x8b, 0x94, 0x69,
	0x89, 0x8e, 0xd5, 0xb2, 0x0e, 0xdd, 0xc3, 0xe7, 0xf7, 0xee, 0x69, 0x54, 0x49, 0x5a, 0x31, 0x99,
	0x2a, 0x1c, 0x4f, 0x7e, 0xc5, 0xd0, 0xbb, 0xf4, 0x72, 0xe2, 0x10, 0x62, 0xd2, 0x49, 0x34, 0x8e,
	0xa6, 0xfd, 0x59, 0x4c, 0x5a, 0x8c, 0x60, 0x60, 0xb1, 0x20, 0xc7, 0x76, 0x9d, 0x91, 0x4e, 0x62,
	0x7f, 0x00, 0x5b, 0xe8, 0x4a, 0x8b, 0x37, 0x00, 0xb9, 0x45, 0xc5, 0xa8, 0x33, 0xc5, 0x49, 0x67,
	0x1c, 0x4d, 0x07, 0xe7, 0x43, 0x19, 0x0c, 0xc9, 0xad, 0x21, 0x79, 0xbd, 0x35, 0x34, 0xeb, 0xb7,
	0xec, 0x0b, 0x16, 0x02, 0xba, 0x95, 0x5a, 0x62, 0xd2, 0xf5, 0xa2, 0xbe, 0x16, 0x63, 0x18, 0x68,
	0x74, 0xb9, 0xa5, 0x7a, 0xe3, 0x2f, 0x39, 0xf0, 0x47, 0xbb, 0x90, 0xb8, 0x81, 0x27, 0x6c, 0x6a,
	0xca, 0x33, 0x55, 0x92, 0x72, 0xe8, 0x92, 0xde, 0xb8, 0x33, 0x1d, 0x9c, 0xbf, 0x92, 0x7b, 0x12,
	0x92, 0xe1, 0x75, 0xf2, 0x7a, 0xd3, 0x78, 0x11, 0xfa, 0x3e, 0x55, 0x6c, 0xd7, 0xb3, 0xc7, 0xbc,
	0x03, 0x0d, 0xdf, 0xc3, 0xf1, 0x5f, 0x14, 0x71, 0x04, 0x9d, 0xef, 0xb8, 0x6e, 0x33, 0xd9, 0x94,
	0xe2, 0x19, 0x1c, 0x34, 0xaa, 0x5c, 0x61, 0x1b, 0x47, 0xf8, 0x78, 0x1b, 0xbf, 0x8e, 0x26, 0x3f,
	0x23, 0x38, 0x0e, 0x77, 0x7d, 0x44, 0xcb, 0xb4, 0xa0, 0x5c, 0x31, 0x8a, 0x13, 0xe8, 0x07, 0x2f,
	0xd9, 0x5d, 0xb6, 0x8f, 0x02, 0x70, 0xa5, 0x37, 0x2f, 0x5e, 0x50, 0x55, 0xa0, 0xad, 0x2d, 0x55,
	0xdc, 0x4a, 0xee, 0x42, 0xe2, 0x05, 0x1c, 0xe5, 0x7f, 0xd4, 0x32, 0xad, 0x58, 0xf9, 0xa0, 0xfb,
	0xb3, 0xa7, 0x3b, 0xf8, 0xa5, 0x62, 0xf5, 0x60, 0x1a, 0xdd, 0x7f, 0x98, 0xc6, 0xe4, 0x16, 0x0e,
	0x83, 0xf3, 0x2f, 0xca, 0xb9, 0x1f, 0xc6, 0xea, 0xfd, 0xb6, 0xc3, 0xa2, 0xc4, 0x77, 0x8b, 0xf2,
	0xff, 0x7b, 0xf0, 0xe1, 0x1b, 0x8c, 0xee, 0xcd, 0x4e, 0xd5, 0xf4, 0x60, 0x7e, 0x37, 0x9f, 0x0b,
	0xe2, 0xaf, 0xab, 0xb9, 0xcc, 0xcd, 0x32, 0x0d, 0xdc, 0xd3, 0xb0, 0xcb, 0x85, 0x39, 0x2d, 0xb0,
	0xf2, 0xfa, 0xe9, 0x9e, 0x9f, 0xe9, 0x5d, 0x5b, 0xce, 0x7b, 0x9e, 0xfa, 0xf2, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6f, 0x2e, 0xe3, 0xc9, 0x7a, 0x03, 0x00, 0x00,
}
