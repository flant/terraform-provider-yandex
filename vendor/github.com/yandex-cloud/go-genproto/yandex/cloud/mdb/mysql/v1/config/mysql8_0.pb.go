// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1/config/mysql8_0.proto

package mysql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type MysqlConfig8_0_SQLMode int32

const (
	MysqlConfig8_0_SQLMODE_UNSPECIFIED        MysqlConfig8_0_SQLMode = 0
	MysqlConfig8_0_ALLOW_INVALID_DATES        MysqlConfig8_0_SQLMode = 1
	MysqlConfig8_0_ANSI_QUOTES                MysqlConfig8_0_SQLMode = 2
	MysqlConfig8_0_ERROR_FOR_DIVISION_BY_ZERO MysqlConfig8_0_SQLMode = 3
	MysqlConfig8_0_HIGH_NOT_PRECEDENCE        MysqlConfig8_0_SQLMode = 4
	MysqlConfig8_0_IGNORE_SPACE               MysqlConfig8_0_SQLMode = 5
	MysqlConfig8_0_NO_AUTO_VALUE_ON_ZERO      MysqlConfig8_0_SQLMode = 6
	MysqlConfig8_0_NO_BACKSLASH_ESCAPES       MysqlConfig8_0_SQLMode = 7
	MysqlConfig8_0_NO_ENGINE_SUBSTITUTION     MysqlConfig8_0_SQLMode = 8
	MysqlConfig8_0_NO_UNSIGNED_SUBTRACTION    MysqlConfig8_0_SQLMode = 9
	MysqlConfig8_0_NO_ZERO_DATE               MysqlConfig8_0_SQLMode = 10
	MysqlConfig8_0_NO_ZERO_IN_DATE            MysqlConfig8_0_SQLMode = 11
	MysqlConfig8_0_ONLY_FULL_GROUP_BY         MysqlConfig8_0_SQLMode = 15
	MysqlConfig8_0_PAD_CHAR_TO_FULL_LENGTH    MysqlConfig8_0_SQLMode = 16
	MysqlConfig8_0_PIPES_AS_CONCAT            MysqlConfig8_0_SQLMode = 17
	MysqlConfig8_0_REAL_AS_FLOAT              MysqlConfig8_0_SQLMode = 18
	MysqlConfig8_0_STRICT_ALL_TABLES          MysqlConfig8_0_SQLMode = 19
	MysqlConfig8_0_STRICT_TRANS_TABLES        MysqlConfig8_0_SQLMode = 20
	MysqlConfig8_0_TIME_TRUNCATE_FRACTIONAL   MysqlConfig8_0_SQLMode = 21
	MysqlConfig8_0_ANSI                       MysqlConfig8_0_SQLMode = 22
	MysqlConfig8_0_TRADITIONAL                MysqlConfig8_0_SQLMode = 23
)

var MysqlConfig8_0_SQLMode_name = map[int32]string{
	0:  "SQLMODE_UNSPECIFIED",
	1:  "ALLOW_INVALID_DATES",
	2:  "ANSI_QUOTES",
	3:  "ERROR_FOR_DIVISION_BY_ZERO",
	4:  "HIGH_NOT_PRECEDENCE",
	5:  "IGNORE_SPACE",
	6:  "NO_AUTO_VALUE_ON_ZERO",
	7:  "NO_BACKSLASH_ESCAPES",
	8:  "NO_ENGINE_SUBSTITUTION",
	9:  "NO_UNSIGNED_SUBTRACTION",
	10: "NO_ZERO_DATE",
	11: "NO_ZERO_IN_DATE",
	15: "ONLY_FULL_GROUP_BY",
	16: "PAD_CHAR_TO_FULL_LENGTH",
	17: "PIPES_AS_CONCAT",
	18: "REAL_AS_FLOAT",
	19: "STRICT_ALL_TABLES",
	20: "STRICT_TRANS_TABLES",
	21: "TIME_TRUNCATE_FRACTIONAL",
	22: "ANSI",
	23: "TRADITIONAL",
}

var MysqlConfig8_0_SQLMode_value = map[string]int32{
	"SQLMODE_UNSPECIFIED":        0,
	"ALLOW_INVALID_DATES":        1,
	"ANSI_QUOTES":                2,
	"ERROR_FOR_DIVISION_BY_ZERO": 3,
	"HIGH_NOT_PRECEDENCE":        4,
	"IGNORE_SPACE":               5,
	"NO_AUTO_VALUE_ON_ZERO":      6,
	"NO_BACKSLASH_ESCAPES":       7,
	"NO_ENGINE_SUBSTITUTION":     8,
	"NO_UNSIGNED_SUBTRACTION":    9,
	"NO_ZERO_DATE":               10,
	"NO_ZERO_IN_DATE":            11,
	"ONLY_FULL_GROUP_BY":         15,
	"PAD_CHAR_TO_FULL_LENGTH":    16,
	"PIPES_AS_CONCAT":            17,
	"REAL_AS_FLOAT":              18,
	"STRICT_ALL_TABLES":          19,
	"STRICT_TRANS_TABLES":        20,
	"TIME_TRUNCATE_FRACTIONAL":   21,
	"ANSI":                       22,
	"TRADITIONAL":                23,
}

func (x MysqlConfig8_0_SQLMode) String() string {
	return proto.EnumName(MysqlConfig8_0_SQLMode_name, int32(x))
}

func (MysqlConfig8_0_SQLMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d7b57328ffba4745, []int{0, 0}
}

type MysqlConfig8_0_AuthPlugin int32

const (
	MysqlConfig8_0_AUTH_PLUGIN_UNSPECIFIED MysqlConfig8_0_AuthPlugin = 0
	MysqlConfig8_0_MYSQL_NATIVE_PASSWORD   MysqlConfig8_0_AuthPlugin = 1
	MysqlConfig8_0_CACHING_SHA2_PASSWORD   MysqlConfig8_0_AuthPlugin = 2
	MysqlConfig8_0_SHA256_PASSWORD         MysqlConfig8_0_AuthPlugin = 3
)

var MysqlConfig8_0_AuthPlugin_name = map[int32]string{
	0: "AUTH_PLUGIN_UNSPECIFIED",
	1: "MYSQL_NATIVE_PASSWORD",
	2: "CACHING_SHA2_PASSWORD",
	3: "SHA256_PASSWORD",
}

var MysqlConfig8_0_AuthPlugin_value = map[string]int32{
	"AUTH_PLUGIN_UNSPECIFIED": 0,
	"MYSQL_NATIVE_PASSWORD":   1,
	"CACHING_SHA2_PASSWORD":   2,
	"SHA256_PASSWORD":         3,
}

func (x MysqlConfig8_0_AuthPlugin) String() string {
	return proto.EnumName(MysqlConfig8_0_AuthPlugin_name, int32(x))
}

func (MysqlConfig8_0_AuthPlugin) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d7b57328ffba4745, []int{0, 1}
}

// Options and structure of `MysqlConfig8_0` reflects MySQL 8.0 configuration file
type MysqlConfig8_0 struct {
	// Size of the InnoDB buffer pool used for caching table and index data.
	//
	// For details, see [MySQL documentation for the parameter](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_buffer_pool_size).
	InnodbBufferPoolSize *wrappers.Int64Value `protobuf:"bytes,1,opt,name=innodb_buffer_pool_size,json=innodbBufferPoolSize,proto3" json:"innodb_buffer_pool_size,omitempty"`
	// The maximum permitted number of simultaneous client connections.
	//
	// For details, see [MySQL documentation for the variable](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_connections).
	MaxConnections *wrappers.Int64Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// Time that it takes to process a query before it is considered slow.
	//
	// For details, see [MySQL documentation for the variable](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_long_query_time).
	LongQueryTime *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=long_query_time,json=longQueryTime,proto3" json:"long_query_time,omitempty"`
	// Enable writing of general query log of MySQL.
	//
	// For details, see [MySQL documentation for the variable](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_general_log).
	GeneralLog *wrappers.BoolValue `protobuf:"bytes,4,opt,name=general_log,json=generalLog,proto3" json:"general_log,omitempty"`
	// Enable writing of audit log of MySQL.
	//
	// For details, see [MySQL documentation for the variable](https://dev.mysql.com/doc/refman/8.0/en/audit-log-reference.html#audit-log-options-variables).
	AuditLog *wrappers.BoolValue `protobuf:"bytes,5,opt,name=audit_log,json=auditLog,proto3" json:"audit_log,omitempty"`
	// Server SQL mode of MySQL.
	//
	// For details, see [MySQL documentation for the variable](https://dev.mysql.com/doc/refman/8.0/en/sql-mode.html#sql-mode-setting).
	SqlMode []MysqlConfig8_0_SQLMode `protobuf:"varint,6,rep,packed,name=sql_mode,json=sqlMode,proto3,enum=yandex.cloud.mdb.mysql.v1.config.MysqlConfig8_0_SQLMode" json:"sql_mode,omitempty"`
	// The maximum size in bytes of one packet.
	//
	// For details, see [MySQL documentation for the variable](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_allowed_packet).
	MaxAllowedPacket            *wrappers.Int64Value      `protobuf:"bytes,7,opt,name=max_allowed_packet,json=maxAllowedPacket,proto3" json:"max_allowed_packet,omitempty"`
	DefaultAuthenticationPlugin MysqlConfig8_0_AuthPlugin `protobuf:"varint,8,opt,name=default_authentication_plugin,json=defaultAuthenticationPlugin,proto3,enum=yandex.cloud.mdb.mysql.v1.config.MysqlConfig8_0_AuthPlugin" json:"default_authentication_plugin,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                  `json:"-"`
	XXX_unrecognized            []byte                    `json:"-"`
	XXX_sizecache               int32                     `json:"-"`
}

func (m *MysqlConfig8_0) Reset()         { *m = MysqlConfig8_0{} }
func (m *MysqlConfig8_0) String() string { return proto.CompactTextString(m) }
func (*MysqlConfig8_0) ProtoMessage()    {}
func (*MysqlConfig8_0) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7b57328ffba4745, []int{0}
}

func (m *MysqlConfig8_0) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MysqlConfig8_0.Unmarshal(m, b)
}
func (m *MysqlConfig8_0) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MysqlConfig8_0.Marshal(b, m, deterministic)
}
func (m *MysqlConfig8_0) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MysqlConfig8_0.Merge(m, src)
}
func (m *MysqlConfig8_0) XXX_Size() int {
	return xxx_messageInfo_MysqlConfig8_0.Size(m)
}
func (m *MysqlConfig8_0) XXX_DiscardUnknown() {
	xxx_messageInfo_MysqlConfig8_0.DiscardUnknown(m)
}

var xxx_messageInfo_MysqlConfig8_0 proto.InternalMessageInfo

func (m *MysqlConfig8_0) GetInnodbBufferPoolSize() *wrappers.Int64Value {
	if m != nil {
		return m.InnodbBufferPoolSize
	}
	return nil
}

func (m *MysqlConfig8_0) GetMaxConnections() *wrappers.Int64Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *MysqlConfig8_0) GetLongQueryTime() *wrappers.DoubleValue {
	if m != nil {
		return m.LongQueryTime
	}
	return nil
}

func (m *MysqlConfig8_0) GetGeneralLog() *wrappers.BoolValue {
	if m != nil {
		return m.GeneralLog
	}
	return nil
}

func (m *MysqlConfig8_0) GetAuditLog() *wrappers.BoolValue {
	if m != nil {
		return m.AuditLog
	}
	return nil
}

func (m *MysqlConfig8_0) GetSqlMode() []MysqlConfig8_0_SQLMode {
	if m != nil {
		return m.SqlMode
	}
	return nil
}

func (m *MysqlConfig8_0) GetMaxAllowedPacket() *wrappers.Int64Value {
	if m != nil {
		return m.MaxAllowedPacket
	}
	return nil
}

func (m *MysqlConfig8_0) GetDefaultAuthenticationPlugin() MysqlConfig8_0_AuthPlugin {
	if m != nil {
		return m.DefaultAuthenticationPlugin
	}
	return MysqlConfig8_0_AUTH_PLUGIN_UNSPECIFIED
}

type MysqlConfigSet8_0 struct {
	// Effective settings for a MySQL 8.0 cluster (a combination of settings defined
	// in [user_config] and [default_config]).
	EffectiveConfig *MysqlConfig8_0 `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	// User-defined settings for a MySQL 8.0 cluster.
	UserConfig *MysqlConfig8_0 `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	// Default configuration for a MySQL 8.0 cluster.
	DefaultConfig        *MysqlConfig8_0 `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MysqlConfigSet8_0) Reset()         { *m = MysqlConfigSet8_0{} }
func (m *MysqlConfigSet8_0) String() string { return proto.CompactTextString(m) }
func (*MysqlConfigSet8_0) ProtoMessage()    {}
func (*MysqlConfigSet8_0) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7b57328ffba4745, []int{1}
}

func (m *MysqlConfigSet8_0) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MysqlConfigSet8_0.Unmarshal(m, b)
}
func (m *MysqlConfigSet8_0) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MysqlConfigSet8_0.Marshal(b, m, deterministic)
}
func (m *MysqlConfigSet8_0) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MysqlConfigSet8_0.Merge(m, src)
}
func (m *MysqlConfigSet8_0) XXX_Size() int {
	return xxx_messageInfo_MysqlConfigSet8_0.Size(m)
}
func (m *MysqlConfigSet8_0) XXX_DiscardUnknown() {
	xxx_messageInfo_MysqlConfigSet8_0.DiscardUnknown(m)
}

var xxx_messageInfo_MysqlConfigSet8_0 proto.InternalMessageInfo

func (m *MysqlConfigSet8_0) GetEffectiveConfig() *MysqlConfig8_0 {
	if m != nil {
		return m.EffectiveConfig
	}
	return nil
}

func (m *MysqlConfigSet8_0) GetUserConfig() *MysqlConfig8_0 {
	if m != nil {
		return m.UserConfig
	}
	return nil
}

func (m *MysqlConfigSet8_0) GetDefaultConfig() *MysqlConfig8_0 {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.mdb.mysql.v1.config.MysqlConfig8_0_SQLMode", MysqlConfig8_0_SQLMode_name, MysqlConfig8_0_SQLMode_value)
	proto.RegisterEnum("yandex.cloud.mdb.mysql.v1.config.MysqlConfig8_0_AuthPlugin", MysqlConfig8_0_AuthPlugin_name, MysqlConfig8_0_AuthPlugin_value)
	proto.RegisterType((*MysqlConfig8_0)(nil), "yandex.cloud.mdb.mysql.v1.config.MysqlConfig8_0")
	proto.RegisterType((*MysqlConfigSet8_0)(nil), "yandex.cloud.mdb.mysql.v1.config.MysqlConfigSet8_0")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1/config/mysql8_0.proto", fileDescriptor_d7b57328ffba4745)
}

var fileDescriptor_d7b57328ffba4745 = []byte{
	// 1010 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xdd, 0x4e, 0xe3, 0x46,
	0x14, 0xc7, 0x0b, 0x61, 0x21, 0x0c, 0x4b, 0x62, 0x86, 0xaf, 0x2c, 0xec, 0xae, 0x10, 0xea, 0x05,
	0x37, 0x38, 0x1f, 0xcb, 0x2e, 0x91, 0x50, 0x2b, 0x4d, 0xec, 0x21, 0x71, 0x6b, 0xc6, 0x66, 0xc6,
	0x01, 0xb1, 0x55, 0x35, 0x72, 0x92, 0x49, 0xb0, 0xea, 0x78, 0x42, 0x62, 0xb3, 0xb0, 0x37, 0x7d,
	0x91, 0xde, 0xf5, 0x61, 0x78, 0x9f, 0x5e, 0x72, 0x55, 0x8d, 0x1d, 0x4a, 0xe9, 0x56, 0x42, 0xbb,
	0x77, 0xf1, 0xf9, 0x9f, 0xff, 0xef, 0xf8, 0xf8, 0x9c, 0xcc, 0x80, 0xf2, 0xad, 0x1f, 0xf5, 0xc4,
	0x4d, 0xb9, 0x1b, 0xca, 0xa4, 0x57, 0x1e, 0xf6, 0x3a, 0xe5, 0xe1, 0xed, 0xe4, 0x2a, 0x2c, 0x5f,
	0x57, 0xcb, 0x5d, 0x19, 0xf5, 0x83, 0x41, 0xf6, 0x5c, 0xe7, 0x15, 0x7d, 0x34, 0x96, 0xb1, 0x84,
	0x3b, 0x99, 0x41, 0x4f, 0x0d, 0xfa, 0xb0, 0xd7, 0xd1, 0xd3, 0x04, 0xfd, 0xba, 0xaa, 0x67, 0x86,
	0xad, 0xb7, 0x03, 0x29, 0x07, 0xa1, 0x28, 0xa7, 0xf9, 0x9d, 0xa4, 0x5f, 0xfe, 0x34, 0xf6, 0x47,
	0x23, 0x31, 0x9e, 0x64, 0x84, 0xad, 0x37, 0x4f, 0x4a, 0x5e, 0xfb, 0x61, 0xd0, 0xf3, 0xe3, 0x40,
	0x46, 0x99, 0xbc, 0xfb, 0x27, 0x00, 0x85, 0x13, 0x85, 0x34, 0x52, 0x5c, 0x9d, 0x57, 0xa0, 0x0f,
	0x36, 0x83, 0x28, 0x92, 0xbd, 0x0e, 0xef, 0x24, 0xfd, 0xbe, 0x18, 0xf3, 0x91, 0x94, 0x21, 0x9f,
	0x04, 0x9f, 0x45, 0x69, 0x66, 0x67, 0x66, 0x6f, 0xa9, 0xb6, 0xad, 0x67, 0x35, 0xf5, 0x87, 0x9a,
	0xba, 0x15, 0xc5, 0x1f, 0x0e, 0xce, 0xfc, 0x30, 0x11, 0x8d, 0xe5, 0xfb, 0xbb, 0xea, 0xe2, 0x8f,
	0x3f, 0xbc, 0xaf, 0x1d, 0xd4, 0xea, 0xf5, 0x0a, 0x5d, 0xcb, 0x50, 0x8d, 0x94, 0xe4, 0x4a, 0x19,
	0xb2, 0xe0, 0xb3, 0x80, 0x14, 0x14, 0x87, 0xfe, 0x0d, 0xef, 0xca, 0x28, 0x12, 0x5d, 0xf5, 0x36,
	0x93, 0xd2, 0xec, 0xf3, 0xe8, 0x97, 0xf7, 0x77, 0xd5, 0x7c, 0xb5, 0xb2, 0x5f, 0xad, 0x54, 0x2a,
	0x15, 0x5a, 0x18, 0xfa, 0x37, 0xc6, 0x23, 0x00, 0x9a, 0xa0, 0x18, 0xca, 0x68, 0xc0, 0xaf, 0x12,
	0x31, 0xbe, 0xe5, 0x71, 0x30, 0x14, 0xa5, 0x5c, 0xca, 0x7c, 0xfd, 0x05, 0xd3, 0x94, 0x49, 0x27,
	0x14, 0x29, 0x94, 0x2e, 0x2b, 0xd3, 0xa9, 0xf2, 0x78, 0xc1, 0x50, 0xc0, 0x23, 0xb0, 0x34, 0x10,
	0x91, 0x18, 0xfb, 0x21, 0x0f, 0xe5, 0xa0, 0x34, 0x97, 0x12, 0xb6, 0xbe, 0x20, 0x34, 0xa4, 0x0c,
	0x33, 0x3f, 0x98, 0xa6, 0xdb, 0x72, 0x00, 0x0f, 0xc1, 0xa2, 0x9f, 0xf4, 0x82, 0x38, 0xb5, 0xbe,
	0x78, 0xd6, 0x9a, 0x4f, 0x93, 0x95, 0x91, 0x81, 0xfc, 0xe4, 0x2a, 0xe4, 0x43, 0xd9, 0x13, 0xa5,
	0xf9, 0x9d, 0xdc, 0x5e, 0xa1, 0x56, 0xd7, 0x9f, 0x9b, 0xbc, 0xfe, 0x74, 0x6c, 0x3a, 0x3b, 0xb5,
	0x4f, 0x64, 0x4f, 0xd0, 0x85, 0xc9, 0x55, 0xa8, 0x7e, 0xc0, 0x5f, 0x01, 0x54, 0x1f, 0xd9, 0x0f,
	0x43, 0xf9, 0x49, 0xf4, 0xf8, 0xc8, 0xef, 0xfe, 0x26, 0xe2, 0xd2, 0xc2, 0xf3, 0xdf, 0x19, 0xde,
	0xdf, 0x55, 0x0b, 0xd5, 0x4a, 0xed, 0x60, 0xbf, 0xfa, 0xee, 0xa0, 0x56, 0x3d, 0x3c, 0xac, 0xd5,
	0xa9, 0x36, 0xf4, 0x6f, 0x50, 0x46, 0x72, 0x53, 0x10, 0xfc, 0x1d, 0xbc, 0xe9, 0x89, 0xbe, 0x9f,
	0x84, 0x31, 0xf7, 0x93, 0xf8, 0x52, 0x44, 0x71, 0xd0, 0x4d, 0x37, 0x8b, 0x8f, 0xc2, 0x64, 0x10,
	0x44, 0xa5, 0xfc, 0xce, 0xcc, 0x5e, 0xa1, 0x76, 0xf4, 0xd5, 0x8d, 0xa0, 0x24, 0xbe, 0x74, 0x53,
	0x04, 0xdd, 0x9e, 0x56, 0x40, 0x4f, 0x0a, 0x64, 0xe2, 0xee, 0x5f, 0x39, 0xb0, 0x30, 0x6d, 0x1a,
	0x6e, 0x82, 0x55, 0xf5, 0xd3, 0x31, 0x31, 0x6f, 0x13, 0xe6, 0x62, 0xc3, 0x3a, 0xb6, 0xb0, 0xa9,
	0x7d, 0xa7, 0x04, 0x64, 0xdb, 0xce, 0x39, 0xb7, 0xc8, 0x19, 0xb2, 0x2d, 0x93, 0x9b, 0xc8, 0xc3,
	0x4c, 0x9b, 0x81, 0x45, 0xb0, 0x84, 0x08, 0xb3, 0xf8, 0x69, 0xdb, 0x51, 0x81, 0x59, 0xf8, 0x16,
	0x6c, 0x61, 0x4a, 0x1d, 0xca, 0x8f, 0x1d, 0xca, 0x4d, 0xeb, 0xcc, 0x62, 0x96, 0x43, 0x78, 0xe3,
	0x82, 0x7f, 0xc4, 0xd4, 0xd1, 0x72, 0x8a, 0xd4, 0xb2, 0x9a, 0x2d, 0x4e, 0x1c, 0x8f, 0xbb, 0x14,
	0x1b, 0xd8, 0xc4, 0xc4, 0xc0, 0xda, 0x1c, 0xd4, 0xc0, 0x4b, 0xab, 0x49, 0x1c, 0x8a, 0x39, 0x73,
	0x91, 0x81, 0xb5, 0x17, 0xf0, 0x15, 0x58, 0x27, 0x0e, 0x47, 0x6d, 0xcf, 0xe1, 0x67, 0xc8, 0x6e,
	0x63, 0xee, 0x90, 0x8c, 0x32, 0x0f, 0x4b, 0x60, 0x8d, 0x38, 0xbc, 0x81, 0x8c, 0x9f, 0x99, 0x8d,
	0x58, 0x8b, 0x63, 0x66, 0x20, 0x17, 0x33, 0x6d, 0x01, 0x6e, 0x81, 0x0d, 0xe2, 0x70, 0x4c, 0x9a,
	0x16, 0xc1, 0x9c, 0xb5, 0x1b, 0xcc, 0xb3, 0xbc, 0xb6, 0x67, 0x39, 0x44, 0xcb, 0xc3, 0x6d, 0xb0,
	0x49, 0x1c, 0xd5, 0x99, 0xd5, 0x24, 0xd8, 0x54, 0xaa, 0x47, 0x91, 0x91, 0x8a, 0x8b, 0xaa, 0x3e,
	0x71, 0x52, 0x7e, 0xda, 0x9c, 0x06, 0xe0, 0x2a, 0x28, 0x3e, 0x44, 0x2c, 0x92, 0x05, 0x97, 0xe0,
	0x06, 0x80, 0x0e, 0xb1, 0x2f, 0xf8, 0x71, 0xdb, 0xb6, 0x79, 0x93, 0x3a, 0x6d, 0x97, 0x37, 0x2e,
	0xb4, 0xa2, 0x62, 0xbb, 0xc8, 0xe4, 0x46, 0x0b, 0x51, 0xee, 0x39, 0x99, 0x6c, 0x63, 0xd2, 0xf4,
	0x5a, 0x9a, 0xa6, 0x48, 0xae, 0xe5, 0x62, 0xc6, 0x11, 0xe3, 0x86, 0x43, 0x0c, 0xe4, 0x69, 0x2b,
	0x70, 0x05, 0x2c, 0x53, 0x8c, 0x6c, 0x15, 0x3b, 0xb6, 0x1d, 0xe4, 0x69, 0x10, 0xae, 0x83, 0x15,
	0xe6, 0x51, 0xcb, 0xf0, 0x38, 0xb2, 0x6d, 0xee, 0xa1, 0x86, 0x8d, 0x99, 0xb6, 0x9a, 0x8e, 0x25,
	0x0b, 0x7b, 0x14, 0x11, 0xf6, 0x20, 0xac, 0xc1, 0xd7, 0xa0, 0xe4, 0x59, 0x27, 0x98, 0x7b, 0xb4,
	0xad, 0xa0, 0x98, 0x1f, 0x4f, 0xfb, 0x41, 0xb6, 0xb6, 0x0e, 0xf3, 0x60, 0x4e, 0xcd, 0x46, 0xdb,
	0x50, 0x53, 0xf2, 0x28, 0x32, 0xad, 0xa9, 0xb4, 0xb9, 0x1b, 0x03, 0xf0, 0xb8, 0x1f, 0xea, 0xdd,
	0x51, 0xdb, 0x6b, 0x71, 0xd7, 0x6e, 0x37, 0x2d, 0xf2, 0x9f, 0xd1, 0xbf, 0x02, 0xeb, 0x27, 0x17,
	0xec, 0xd4, 0xe6, 0x04, 0x79, 0xd6, 0x19, 0xe6, 0x2e, 0x62, 0xec, 0xdc, 0xa1, 0xa6, 0x36, 0xa3,
	0x24, 0x03, 0x19, 0x2d, 0x8b, 0x34, 0x39, 0x6b, 0xa1, 0xda, 0xa3, 0x34, 0xab, 0x3a, 0x56, 0xa1,
	0xf7, 0x1f, 0x1e, 0x83, 0xb9, 0xdd, 0x3f, 0x66, 0xc1, 0xca, 0xbf, 0xb6, 0x94, 0x89, 0x58, 0x1d,
	0x94, 0xbf, 0x00, 0x4d, 0xf4, 0xfb, 0xea, 0xfc, 0xb9, 0x16, 0x3c, 0xdb, 0xe5, 0xe9, 0x09, 0x59,
	0xf9, 0xda, 0xa5, 0xa7, 0xc5, 0x7f, 0x48, 0x59, 0x0c, 0x9e, 0x82, 0xa5, 0x64, 0x22, 0xc6, 0x0f,
	0xdc, 0xd9, 0x6f, 0xe4, 0x02, 0x05, 0x99, 0x22, 0xcf, 0x41, 0xe1, 0xe1, 0x1f, 0x3b, 0xa5, 0xe6,
	0xbe, 0x91, 0xba, 0x3c, 0xe5, 0x64, 0x91, 0xc6, 0x18, 0x7c, 0xff, 0x84, 0xe0, 0x8f, 0x82, 0xff,
	0xa3, 0x7c, 0xfc, 0x69, 0x10, 0xc4, 0x97, 0x49, 0x47, 0xef, 0xca, 0xe1, 0xf4, 0x26, 0xdc, 0xcf,
	0xae, 0xa5, 0x81, 0xdc, 0x1f, 0x88, 0x28, 0x3d, 0x8b, 0x9e, 0xbd, 0x22, 0x8f, 0xd2, 0xe7, 0xce,
	0x7c, 0x9a, 0xfd, 0xee, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0xa8, 0x01, 0x0f, 0x53, 0x07,
	0x00, 0x00,
}
