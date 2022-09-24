// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc_api/bilibili/metadata/locale/locale.proto

package bilibili_metadata_locale

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// 区域标识
// gRPC头部:x-bili-locale-bin
type Locale struct {
	// App设置的locale
	CLocale *LocaleIds `protobuf:"bytes,1,opt,name=c_locale,json=cLocale,proto3" json:"c_locale,omitempty"`
	// 系统默认的locale
	SLocale *LocaleIds `protobuf:"bytes,2,opt,name=s_locale,json=sLocale,proto3" json:"s_locale,omitempty"`
	// sim卡的国家码+运营商码
	SimCode string `protobuf:"bytes,3,opt,name=sim_code,json=simCode,proto3" json:"sim_code,omitempty"`
	// 时区
	Timezone             string   `protobuf:"bytes,4,opt,name=timezone,proto3" json:"timezone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Locale) Reset()         { *m = Locale{} }
func (m *Locale) String() string { return proto.CompactTextString(m) }
func (*Locale) ProtoMessage()    {}
func (*Locale) Descriptor() ([]byte, []int) {
	return fileDescriptor_352efcc66b584675, []int{0}
}

func (m *Locale) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Locale.Unmarshal(m, b)
}
func (m *Locale) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Locale.Marshal(b, m, deterministic)
}
func (m *Locale) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Locale.Merge(m, src)
}
func (m *Locale) XXX_Size() int {
	return xxx_messageInfo_Locale.Size(m)
}
func (m *Locale) XXX_DiscardUnknown() {
	xxx_messageInfo_Locale.DiscardUnknown(m)
}

var xxx_messageInfo_Locale proto.InternalMessageInfo

func (m *Locale) GetCLocale() *LocaleIds {
	if m != nil {
		return m.CLocale
	}
	return nil
}

func (m *Locale) GetSLocale() *LocaleIds {
	if m != nil {
		return m.SLocale
	}
	return nil
}

func (m *Locale) GetSimCode() string {
	if m != nil {
		return m.SimCode
	}
	return ""
}

func (m *Locale) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

// Defined by https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/LanguageandLocaleIDs/LanguageandLocaleIDs.html
type LocaleIds struct {
	// A language designator is a code that represents a language.
	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	// Writing systems.
	Script string `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	// A region designator is a code that represents a country or an area.
	Region               string   `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocaleIds) Reset()         { *m = LocaleIds{} }
func (m *LocaleIds) String() string { return proto.CompactTextString(m) }
func (*LocaleIds) ProtoMessage()    {}
func (*LocaleIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_352efcc66b584675, []int{1}
}

func (m *LocaleIds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocaleIds.Unmarshal(m, b)
}
func (m *LocaleIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocaleIds.Marshal(b, m, deterministic)
}
func (m *LocaleIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocaleIds.Merge(m, src)
}
func (m *LocaleIds) XXX_Size() int {
	return xxx_messageInfo_LocaleIds.Size(m)
}
func (m *LocaleIds) XXX_DiscardUnknown() {
	xxx_messageInfo_LocaleIds.DiscardUnknown(m)
}

var xxx_messageInfo_LocaleIds proto.InternalMessageInfo

func (m *LocaleIds) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *LocaleIds) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *LocaleIds) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func init() {
	proto.RegisterType((*Locale)(nil), "bilibili.metadata.locale.Locale")
	proto.RegisterType((*LocaleIds)(nil), "bilibili.metadata.locale.LocaleIds")
}

func init() {
	proto.RegisterFile("grpc_api/bilibili/metadata/locale/locale.proto", fileDescriptor_352efcc66b584675)
}

var fileDescriptor_352efcc66b584675 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbd, 0x4e, 0x04, 0x21,
	0x14, 0x46, 0x83, 0x9a, 0xd9, 0x01, 0x3b, 0x0a, 0x83, 0x56, 0x9b, 0xb5, 0xd9, 0x8a, 0x49, 0xb4,
	0xb7, 0xb1, 0x32, 0xb1, 0xa2, 0xb1, 0x9c, 0xb0, 0x40, 0xc8, 0x4d, 0x60, 0x20, 0x80, 0x8d, 0x2f,
	0xe7, 0xab, 0x99, 0x01, 0x66, 0xac, 0x2c, 0x2c, 0x08, 0x9c, 0x7b, 0xbf, 0xc3, 0x1f, 0xe1, 0x36,
	0x45, 0x35, 0xcb, 0x08, 0xd3, 0x05, 0x1c, 0xac, 0x63, 0xf2, 0xa6, 0x48, 0x2d, 0x8b, 0x9c, 0x5c,
	0x50, 0xd2, 0x99, 0x3e, 0xf1, 0x98, 0x42, 0x09, 0x94, 0x6d, 0x31, 0xbe, 0xc5, 0x78, 0xeb, 0x9f,
	0xbe, 0x11, 0x19, 0xde, 0xeb, 0x92, 0xbe, 0x90, 0x51, 0xcd, 0xad, 0xcc, 0xd0, 0x11, 0x9d, 0x6f,
	0x9f, 0x1e, 0xf9, 0x5f, 0x1e, 0x6f, 0xce, 0x9b, 0xce, 0xe2, 0xa0, 0x7e, 0xfd, 0xbc, 0xf9, 0x57,
	0xff, 0xf0, 0x73, 0xf7, 0xef, 0xc9, 0x98, 0xc1, 0xcf, 0x2a, 0x68, 0xc3, 0xae, 0x8f, 0xe8, 0x8c,
	0xc5, 0x21, 0x83, 0x7f, 0x0d, 0xda, 0xd0, 0x07, 0x32, 0x16, 0xf0, 0xe6, 0x2b, 0x2c, 0x86, 0xdd,
	0xd4, 0xd6, 0xce, 0xa7, 0x0f, 0x82, 0xf7, 0xcd, 0xd6, 0xa0, 0x93, 0x8b, 0xfd, 0x94, 0xb6, 0xbd,
	0x01, 0x8b, 0x9d, 0xe9, 0x1d, 0x19, 0xb2, 0x4a, 0x10, 0x4b, 0xbd, 0x1d, 0x16, 0x9d, 0xd6, 0x7a,
	0x32, 0x16, 0xc2, 0xd2, 0x4f, 0xed, 0x74, 0x19, 0xea, 0xdf, 0x3d, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x52, 0x54, 0x3d, 0x98, 0x6d, 0x01, 0x00, 0x00,
}
