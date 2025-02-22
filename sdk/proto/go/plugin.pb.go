// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin.proto

package pulumirpc

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

// PluginInfo is meta-information about a plugin that is used by the system.
type PluginInfo struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginInfo) Reset()         { *m = PluginInfo{} }
func (m *PluginInfo) String() string { return proto.CompactTextString(m) }
func (*PluginInfo) ProtoMessage()    {}
func (*PluginInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{0}
}

func (m *PluginInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginInfo.Unmarshal(m, b)
}
func (m *PluginInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginInfo.Marshal(b, m, deterministic)
}
func (m *PluginInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginInfo.Merge(m, src)
}
func (m *PluginInfo) XXX_Size() int {
	return xxx_messageInfo_PluginInfo.Size(m)
}
func (m *PluginInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PluginInfo proto.InternalMessageInfo

func (m *PluginInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// PluginDependency is information about a plugin that a program may depend upon.
type PluginDependency struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Server               string   `protobuf:"bytes,4,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginDependency) Reset()         { *m = PluginDependency{} }
func (m *PluginDependency) String() string { return proto.CompactTextString(m) }
func (*PluginDependency) ProtoMessage()    {}
func (*PluginDependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{1}
}

func (m *PluginDependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginDependency.Unmarshal(m, b)
}
func (m *PluginDependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginDependency.Marshal(b, m, deterministic)
}
func (m *PluginDependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginDependency.Merge(m, src)
}
func (m *PluginDependency) XXX_Size() int {
	return xxx_messageInfo_PluginDependency.Size(m)
}
func (m *PluginDependency) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginDependency.DiscardUnknown(m)
}

var xxx_messageInfo_PluginDependency proto.InternalMessageInfo

func (m *PluginDependency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PluginDependency) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *PluginDependency) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PluginDependency) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

// PluginAttach is used to attach an already running plugin to the engine.
//
// Normally the engine starts the plugin process itself and passes the engine address as the first argumnent.
// But when debugging it can be useful to have an already running provider that the engine instead attaches
// to, this message is used so the provider can still be passed the engine address to communicate with.
type PluginAttach struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PluginAttach) Reset()         { *m = PluginAttach{} }
func (m *PluginAttach) String() string { return proto.CompactTextString(m) }
func (*PluginAttach) ProtoMessage()    {}
func (*PluginAttach) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{2}
}

func (m *PluginAttach) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PluginAttach.Unmarshal(m, b)
}
func (m *PluginAttach) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PluginAttach.Marshal(b, m, deterministic)
}
func (m *PluginAttach) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PluginAttach.Merge(m, src)
}
func (m *PluginAttach) XXX_Size() int {
	return xxx_messageInfo_PluginAttach.Size(m)
}
func (m *PluginAttach) XXX_DiscardUnknown() {
	xxx_messageInfo_PluginAttach.DiscardUnknown(m)
}

var xxx_messageInfo_PluginAttach proto.InternalMessageInfo

func (m *PluginAttach) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*PluginInfo)(nil), "pulumirpc.PluginInfo")
	proto.RegisterType((*PluginDependency)(nil), "pulumirpc.PluginDependency")
	proto.RegisterType((*PluginAttach)(nil), "pulumirpc.PluginAttach")
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptor_22a625af4bc1cc87) }

var fileDescriptor_22a625af4bc1cc87 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0x29, 0x4d,
	0xcf, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c, 0x28, 0xcd, 0x29, 0xcd, 0xcd,
	0x2c, 0x2a, 0x48, 0x56, 0x52, 0xe3, 0xe2, 0x0a, 0x00, 0x4b, 0x79, 0xe6, 0xa5, 0xe5, 0x0b, 0x49,
	0x70, 0xb1, 0x97, 0xa5, 0x16, 0x15, 0x67, 0xe6, 0xe7, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xc1, 0xb8, 0x4a, 0x39, 0x5c, 0x02, 0x10, 0x75, 0x2e, 0xa9, 0x05, 0xa9, 0x79, 0x29, 0xa9, 0x79,
	0xc9, 0x95, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x50, 0xa5, 0x60, 0x36, 0x48, 0x2c,
	0x3b, 0x33, 0x2f, 0x45, 0x82, 0x09, 0x22, 0x06, 0x62, 0x23, 0x9b, 0xca, 0x8c, 0x62, 0xaa, 0x90,
	0x18, 0x17, 0x5b, 0x71, 0x6a, 0x51, 0x59, 0x6a, 0x91, 0x04, 0x0b, 0x58, 0x02, 0xca, 0x53, 0xd2,
	0xe0, 0xe2, 0x81, 0xd8, 0xe6, 0x58, 0x52, 0x92, 0x98, 0x9c, 0x01, 0x32, 0x21, 0x31, 0x25, 0xa5,
	0x28, 0xb5, 0xb8, 0x18, 0xe6, 0x2e, 0x28, 0x37, 0x89, 0x0d, 0xec, 0x23, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8d, 0x82, 0x2c, 0xaa, 0xe1, 0x00, 0x00, 0x00,
}
