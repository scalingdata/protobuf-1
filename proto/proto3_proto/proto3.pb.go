// Code generated by protoc-gen-gogo.
// source: proto3_proto/proto3.proto
// DO NOT EDIT!

/*
Package proto3_proto is a generated protocol buffer package.

It is generated from these files:
	proto3_proto/proto3.proto

It has these top-level messages:
	Message
	Nested
	MessageWithMap
*/
package proto3_proto

import proto "github.com/scalingdata/gogo-protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/scalingdata/gogo-protobuf/types"
import testdata "github.com/scalingdata/gogo-protobuf/proto/testdata"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Message_Humour int32

const (
	Message_UNKNOWN     Message_Humour = 0
	Message_PUNS        Message_Humour = 1
	Message_SLAPSTICK   Message_Humour = 2
	Message_BILL_BAILEY Message_Humour = 3
)

var Message_Humour_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUNS",
	2: "SLAPSTICK",
	3: "BILL_BAILEY",
}
var Message_Humour_value = map[string]int32{
	"UNKNOWN":     0,
	"PUNS":        1,
	"SLAPSTICK":   2,
	"BILL_BAILEY": 3,
}

func (x Message_Humour) String() string {
	return proto.EnumName(Message_Humour_name, int32(x))
}
func (Message_Humour) EnumDescriptor() ([]byte, []int) { return fileDescriptorProto3, []int{0, 0} }

type Message struct {
	Name         string                           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hilarity     Message_Humour                   `protobuf:"varint,2,opt,name=hilarity,proto3,enum=proto3_proto.Message_Humour" json:"hilarity,omitempty"`
	HeightInCm   uint32                           `protobuf:"varint,3,opt,name=height_in_cm,json=heightInCm,proto3" json:"height_in_cm,omitempty"`
	Data         []byte                           `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ResultCount  int64                            `protobuf:"varint,7,opt,name=result_count,json=resultCount,proto3" json:"result_count,omitempty"`
	TrueScotsman bool                             `protobuf:"varint,8,opt,name=true_scotsman,json=trueScotsman,proto3" json:"true_scotsman,omitempty"`
	Score        float32                          `protobuf:"fixed32,9,opt,name=score,proto3" json:"score,omitempty"`
	Key          []uint64                         `protobuf:"varint,5,rep,packed,name=key" json:"key,omitempty"`
	Nested       *Nested                          `protobuf:"bytes,6,opt,name=nested" json:"nested,omitempty"`
	RFunny       []Message_Humour                 `protobuf:"varint,16,rep,packed,name=r_funny,json=rFunny,enum=proto3_proto.Message_Humour" json:"r_funny,omitempty"`
	Terrain      map[string]*Nested               `protobuf:"bytes,10,rep,name=terrain" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Proto2Field  *testdata.SubDefaults            `protobuf:"bytes,11,opt,name=proto2_field,json=proto2Field" json:"proto2_field,omitempty"`
	Proto2Value  map[string]*testdata.SubDefaults `protobuf:"bytes,13,rep,name=proto2_value,json=proto2Value" json:"proto2_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Anything     *google_protobuf.Any             `protobuf:"bytes,14,opt,name=anything" json:"anything,omitempty"`
	ManyThings   []*google_protobuf.Any           `protobuf:"bytes,15,rep,name=many_things,json=manyThings" json:"many_things,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{0} }

func (m *Message) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Message) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

func (m *Message) GetProto2Field() *testdata.SubDefaults {
	if m != nil {
		return m.Proto2Field
	}
	return nil
}

func (m *Message) GetProto2Value() map[string]*testdata.SubDefaults {
	if m != nil {
		return m.Proto2Value
	}
	return nil
}

func (m *Message) GetAnything() *google_protobuf.Any {
	if m != nil {
		return m.Anything
	}
	return nil
}

func (m *Message) GetManyThings() []*google_protobuf.Any {
	if m != nil {
		return m.ManyThings
	}
	return nil
}

type Nested struct {
	Bunny string `protobuf:"bytes,1,opt,name=bunny,proto3" json:"bunny,omitempty"`
	Cute  bool   `protobuf:"varint,2,opt,name=cute,proto3" json:"cute,omitempty"`
}

func (m *Nested) Reset()                    { *m = Nested{} }
func (m *Nested) String() string            { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()               {}
func (*Nested) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{1} }

type MessageWithMap struct {
	ByteMapping map[bool][]byte `protobuf:"bytes,1,rep,name=byte_mapping,json=byteMapping" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *MessageWithMap) Reset()                    { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string            { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()               {}
func (*MessageWithMap) Descriptor() ([]byte, []int) { return fileDescriptorProto3, []int{2} }

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "proto3_proto.Message")
	proto.RegisterType((*Nested)(nil), "proto3_proto.Nested")
	proto.RegisterType((*MessageWithMap)(nil), "proto3_proto.MessageWithMap")
	proto.RegisterEnum("proto3_proto.Message_Humour", Message_Humour_name, Message_Humour_value)
}

func init() { proto.RegisterFile("proto3_proto/proto3.proto", fileDescriptorProto3) }

var fileDescriptorProto3 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x25, 0x4d, 0x97, 0x66, 0x37, 0xe9, 0x16, 0x99, 0x21, 0x79, 0x15, 0x0f, 0xa1, 0x48, 0x28,
	0xe2, 0x23, 0x43, 0x45, 0x93, 0x26, 0x84, 0x40, 0xdb, 0xd8, 0x44, 0xb5, 0xae, 0x54, 0xee, 0xc6,
	0xc4, 0x53, 0xe4, 0x76, 0x6e, 0x1b, 0xd1, 0x38, 0x55, 0xe2, 0x20, 0xe5, 0xef, 0xf0, 0x2b, 0x79,
	0x44, 0xb6, 0xd3, 0xae, 0x9b, 0x0a, 0x3c, 0xe5, 0xfa, 0xf8, 0xdc, 0x0f, 0x9f, 0x73, 0x03, 0xfb,
	0x8b, 0x2c, 0x15, 0xe9, 0xbb, 0x48, 0x7d, 0x0e, 0xf4, 0x21, 0x54, 0x1f, 0xe4, 0xae, 0x5f, 0xb5,
	0xf6, 0xa7, 0x69, 0x3a, 0x9d, 0x33, 0x4d, 0x19, 0x15, 0x93, 0x03, 0xca, 0x4b, 0x4d, 0x6c, 0x3d,
	0x16, 0x2c, 0x17, 0xb7, 0x54, 0xd0, 0x03, 0x19, 0x68, 0xb0, 0xfd, 0xdb, 0x82, 0xc6, 0x25, 0xcb,
	0x73, 0x3a, 0x65, 0x08, 0x41, 0x9d, 0xd3, 0x84, 0x61, 0xc3, 0x37, 0x82, 0x6d, 0xa2, 0x62, 0x74,
	0x04, 0xf6, 0x2c, 0x9e, 0xd3, 0x2c, 0x16, 0x25, 0xae, 0xf9, 0x46, 0xb0, 0xd3, 0x79, 0x1a, 0xae,
	0x37, 0x0c, 0xab, 0xe4, 0xf0, 0x4b, 0x91, 0xa4, 0x45, 0x46, 0x56, 0x6c, 0xe4, 0x83, 0x3b, 0x63,
	0xf1, 0x74, 0x26, 0xa2, 0x98, 0x47, 0xe3, 0x04, 0x9b, 0xbe, 0x11, 0x34, 0x09, 0x68, 0xac, 0xcb,
	0x4f, 0x13, 0xd9, 0x4f, 0x8e, 0x83, 0xeb, 0xbe, 0x11, 0xb8, 0x44, 0xc5, 0xe8, 0x19, 0xb8, 0x19,
	0xcb, 0x8b, 0xb9, 0x88, 0xc6, 0x69, 0xc1, 0x05, 0x6e, 0xf8, 0x46, 0x60, 0x12, 0x47, 0x63, 0xa7,
	0x12, 0x42, 0xcf, 0xa1, 0x29, 0xb2, 0x82, 0x45, 0xf9, 0x38, 0x15, 0x79, 0x42, 0x39, 0xb6, 0x7d,
	0x23, 0xb0, 0x89, 0x2b, 0xc1, 0x61, 0x85, 0xa1, 0x3d, 0xd8, 0xca, 0xc7, 0x69, 0xc6, 0xf0, 0xb6,
	0x6f, 0x04, 0x35, 0xa2, 0x0f, 0xc8, 0x03, 0xf3, 0x07, 0x2b, 0xf1, 0x96, 0x6f, 0x06, 0x75, 0x22,
	0x43, 0xf4, 0x1a, 0x2c, 0xce, 0x72, 0xc1, 0x6e, 0xb1, 0xe5, 0x1b, 0x81, 0xd3, 0xd9, 0xbb, 0xff,
	0xba, 0xbe, 0xba, 0x23, 0x15, 0x07, 0x1d, 0x42, 0x23, 0x8b, 0x26, 0x05, 0xe7, 0x25, 0xf6, 0x7c,
	0xf3, 0xbf, 0x62, 0x58, 0xd9, 0xb9, 0xe4, 0xa2, 0x0f, 0xd0, 0x10, 0x2c, 0xcb, 0x68, 0xcc, 0x31,
	0xf8, 0x66, 0xe0, 0x74, 0xda, 0x9b, 0xd3, 0xae, 0x34, 0xe9, 0x8c, 0x8b, 0xac, 0x24, 0xcb, 0x14,
	0x74, 0x04, 0xda, 0xe2, 0x4e, 0x34, 0x89, 0xd9, 0xfc, 0x16, 0x3b, 0x6a, 0xd0, 0x27, 0xe1, 0xd2,
	0xce, 0x70, 0x58, 0x8c, 0x3e, 0xb3, 0x09, 0x2d, 0xe6, 0x22, 0x27, 0x8e, 0xa6, 0x9e, 0x4b, 0x26,
	0xea, 0xae, 0x32, 0x7f, 0xd2, 0x79, 0xc1, 0x70, 0x53, 0x35, 0x7f, 0xb1, 0xb9, 0xf9, 0x40, 0x31,
	0xbf, 0x49, 0xa2, 0x1e, 0xa0, 0x2a, 0xa5, 0x10, 0xf4, 0x16, 0x6c, 0xca, 0x4b, 0x31, 0x8b, 0xf9,
	0x14, 0xef, 0x54, 0x4a, 0xe9, 0x55, 0x0b, 0x97, 0xab, 0x16, 0x1e, 0xf3, 0x92, 0xac, 0x58, 0xe8,
	0x10, 0x9c, 0x84, 0xf2, 0x32, 0x52, 0xa7, 0x1c, 0xef, 0xaa, 0xde, 0x9b, 0x93, 0x40, 0x12, 0xaf,
	0x14, 0xaf, 0x35, 0x00, 0x77, 0x5d, 0x86, 0xa5, 0x65, 0x7a, 0x27, 0x95, 0x65, 0x2f, 0x61, 0x4b,
	0x3f, 0xa7, 0xf6, 0x0f, 0xc7, 0x34, 0xe5, 0x7d, 0xed, 0xc8, 0x68, 0x5d, 0x83, 0xf7, 0xf0, 0x6d,
	0x1b, 0xaa, 0xbe, 0xba, 0x5f, 0xf5, 0x2f, 0xf2, 0xde, 0x95, 0x6d, 0x7f, 0x02, 0x4b, 0xdb, 0x8c,
	0x1c, 0x68, 0x5c, 0xf7, 0x2f, 0xfa, 0x5f, 0x6f, 0xfa, 0xde, 0x23, 0x64, 0x43, 0x7d, 0x70, 0xdd,
	0x1f, 0x7a, 0x06, 0x6a, 0xc2, 0xf6, 0xb0, 0x77, 0x3c, 0x18, 0x5e, 0x75, 0x4f, 0x2f, 0xbc, 0x1a,
	0xda, 0x05, 0xe7, 0xa4, 0xdb, 0xeb, 0x45, 0x27, 0xc7, 0xdd, 0xde, 0xd9, 0x77, 0xcf, 0x6c, 0x77,
	0xc0, 0xd2, 0xc3, 0xca, 0x65, 0x1d, 0xa9, 0xa5, 0xd2, 0xf3, 0xe8, 0x83, 0xfc, 0x3d, 0xc6, 0x85,
	0xd0, 0x03, 0xd9, 0x44, 0xc5, 0xed, 0x5f, 0x06, 0xec, 0x54, 0x86, 0xdd, 0xc4, 0x62, 0x76, 0x49,
	0x17, 0x68, 0x00, 0xee, 0xa8, 0x14, 0x2c, 0x4a, 0xe8, 0x62, 0x21, 0xdd, 0x31, 0x94, 0xd0, 0x6f,
	0x36, 0x9a, 0x5c, 0xe5, 0x84, 0x27, 0xa5, 0x60, 0x97, 0x9a, 0x5f, 0x79, 0x3d, 0xba, 0x43, 0x5a,
	0x1f, 0xc1, 0x7b, 0x48, 0x58, 0x17, 0xcc, 0xd6, 0x82, 0xed, 0xad, 0x0b, 0xe6, 0xae, 0x29, 0x33,
	0xb2, 0x74, 0xeb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0xc9, 0x75, 0x36, 0xb5, 0x04, 0x00,
	0x00,
}
