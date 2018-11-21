// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iam/im/config.proto

package pbim // import "openpitrix.io/iam/pkg/pb/im"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Config struct {
	LogLevel             *string           `protobuf:"bytes,1,opt,name=log_level,json=logLevel" json:"log_level,omitempty"`
	Mysql                *MysqlConfig      `protobuf:"bytes,2,opt,name=mysql" json:"mysql,omitempty"`
	Extra                map[string]string `protobuf:"bytes,100,rep,name=extra" json:"extra,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_97a6ca372799db2e, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetLogLevel() string {
	if m != nil && m.LogLevel != nil {
		return *m.LogLevel
	}
	return ""
}

func (m *Config) GetMysql() *MysqlConfig {
	if m != nil {
		return m.Mysql
	}
	return nil
}

func (m *Config) GetExtra() map[string]string {
	if m != nil {
		return m.Extra
	}
	return nil
}

type MysqlConfig struct {
	Host                 *string  `protobuf:"bytes,1,opt,name=host,def=openpitrix-db" json:"host,omitempty"`
	Port                 *int32   `protobuf:"varint,2,opt,name=port,def=3306" json:"port,omitempty"`
	User                 *string  `protobuf:"bytes,3,opt,name=user,def=root" json:"user,omitempty"`
	Password             *string  `protobuf:"bytes,4,opt,name=password,def=password" json:"password,omitempty"`
	Database             *string  `protobuf:"bytes,5,opt,name=database,def=openpitrix" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MysqlConfig) Reset()         { *m = MysqlConfig{} }
func (m *MysqlConfig) String() string { return proto.CompactTextString(m) }
func (*MysqlConfig) ProtoMessage()    {}
func (*MysqlConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_97a6ca372799db2e, []int{1}
}
func (m *MysqlConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MysqlConfig.Unmarshal(m, b)
}
func (m *MysqlConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MysqlConfig.Marshal(b, m, deterministic)
}
func (dst *MysqlConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MysqlConfig.Merge(dst, src)
}
func (m *MysqlConfig) XXX_Size() int {
	return xxx_messageInfo_MysqlConfig.Size(m)
}
func (m *MysqlConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MysqlConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MysqlConfig proto.InternalMessageInfo

const Default_MysqlConfig_Host string = "openpitrix-db"
const Default_MysqlConfig_Port int32 = 3306
const Default_MysqlConfig_User string = "root"
const Default_MysqlConfig_Password string = "password"
const Default_MysqlConfig_Database string = "openpitrix"

func (m *MysqlConfig) GetHost() string {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return Default_MysqlConfig_Host
}

func (m *MysqlConfig) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return Default_MysqlConfig_Port
}

func (m *MysqlConfig) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return Default_MysqlConfig_User
}

func (m *MysqlConfig) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return Default_MysqlConfig_Password
}

func (m *MysqlConfig) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return Default_MysqlConfig_Database
}

func init() {
	proto.RegisterType((*Config)(nil), "iam.im.Config")
	proto.RegisterMapType((map[string]string)(nil), "iam.im.Config.ExtraEntry")
	proto.RegisterType((*MysqlConfig)(nil), "iam.im.MysqlConfig")
}

func init() { proto.RegisterFile("iam/im/config.proto", fileDescriptor_config_97a6ca372799db2e) }

var fileDescriptor_config_97a6ca372799db2e = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x4e, 0x2a, 0x31,
	0x14, 0x87, 0x33, 0x30, 0x43, 0xe0, 0x90, 0x9b, 0xdc, 0x94, 0xbb, 0x98, 0x8b, 0x0b, 0x47, 0x62,
	0x0c, 0x2e, 0x9c, 0x1a, 0x48, 0x0c, 0xc1, 0x9d, 0x86, 0x9d, 0x6e, 0xba, 0x74, 0x63, 0x3a, 0x4e,
	0x19, 0x1b, 0xa6, 0x9c, 0xda, 0x16, 0x84, 0xd7, 0x72, 0xeb, 0xcb, 0x99, 0x69, 0x15, 0xdc, 0xf5,
	0x7c, 0xbf, 0xef, 0xfc, 0x49, 0x0a, 0x03, 0xc9, 0x15, 0x95, 0x8a, 0xbe, 0xe0, 0x7a, 0x29, 0xab,
	0x5c, 0x1b, 0x74, 0x48, 0x3a, 0x92, 0xab, 0x5c, 0xaa, 0xe1, 0x69, 0x85, 0x58, 0xd5, 0x82, 0x7a,
	0x5a, 0x6c, 0x96, 0xd4, 0x49, 0x25, 0xac, 0xe3, 0x4a, 0x07, 0x71, 0xf4, 0x19, 0x41, 0xe7, 0xde,
	0x77, 0x92, 0x13, 0xe8, 0xd5, 0x58, 0x3d, 0xd7, 0x62, 0x2b, 0xea, 0x34, 0xca, 0xa2, 0x71, 0x8f,
	0x75, 0x6b, 0xac, 0x1e, 0x9a, 0x9a, 0x5c, 0x42, 0xa2, 0xf6, 0xf6, 0xad, 0x4e, 0x5b, 0x59, 0x34,
	0xee, 0x4f, 0x06, 0x79, 0x58, 0x90, 0x3f, 0x36, 0x30, 0x0c, 0x60, 0xc1, 0x20, 0x14, 0x12, 0xb1,
	0x73, 0x86, 0xa7, 0x65, 0xd6, 0x1e, 0xf7, 0x27, 0xff, 0x7f, 0xd4, 0x60, 0xe5, 0x8b, 0x26, 0x5b,
	0xac, 0x9d, 0xd9, 0xb3, 0xe0, 0x0d, 0x67, 0x00, 0x47, 0x48, 0xfe, 0x42, 0x7b, 0x25, 0xf6, 0xdf,
	0x07, 0x34, 0x4f, 0xf2, 0x0f, 0x92, 0x2d, 0xaf, 0x37, 0xc2, 0xef, 0xee, 0xb1, 0x50, 0xcc, 0x5b,
	0xb3, 0x68, 0xf4, 0x11, 0x41, 0xff, 0xd7, 0x05, 0xe4, 0x0c, 0xe2, 0x57, 0xb4, 0x2e, 0x34, 0xcf,
	0xff, 0xa0, 0x16, 0x6b, 0x2d, 0x9d, 0x91, 0xbb, 0xab, 0xb2, 0x60, 0x3e, 0x22, 0x29, 0xc4, 0x1a,
	0x8d, 0xf3, 0xb3, 0x92, 0x79, 0x3c, 0x9d, 0x5e, 0xdf, 0x30, 0x4f, 0x9a, 0x64, 0x63, 0x85, 0x49,
	0xdb, 0xbe, 0x39, 0x36, 0x88, 0x8e, 0x79, 0x42, 0xce, 0xa1, 0xab, 0xb9, 0xb5, 0xef, 0x68, 0xca,
	0x34, 0xf6, 0xe9, 0xa1, 0x66, 0x87, 0x17, 0xb9, 0x80, 0x6e, 0xc9, 0x1d, 0x2f, 0xb8, 0x15, 0x69,
	0xe2, 0x2d, 0x38, 0x1e, 0xc0, 0x0e, 0xd9, 0xdd, 0xe8, 0x29, 0x3b, 0xf2, 0x5c, 0x22, 0x6d, 0x3e,
	0x50, 0xaf, 0x2a, 0xaa, 0x0b, 0x2a, 0xd5, 0xad, 0x2e, 0xa4, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x17, 0xff, 0x35, 0xa2, 0xd5, 0x01, 0x00, 0x00,
}