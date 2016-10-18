// Code generated by protoc-gen-gogo.
// source: nopackage.proto
// DO NOT EDIT!

/*
Package nopackage is a generated protocol buffer package.

It is generated from these files:
	nopackage.proto

It has these top-level messages:
	M
*/
package nopackage

import proto "github.com/scalingdata/gogo-protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type M struct {
	F map[string]float64 `protobuf:"bytes,1,rep,name=f" json:"f,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (m *M) Reset()                    { *m = M{} }
func (m *M) String() string            { return proto.CompactTextString(m) }
func (*M) ProtoMessage()               {}
func (*M) Descriptor() ([]byte, []int) { return fileDescriptorNopackage, []int{0} }

func (m *M) GetF() map[string]float64 {
	if m != nil {
		return m.F
	}
	return nil
}

func init() {
	proto.RegisterType((*M)(nil), "M")
}
func (m *M) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *M) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.F) > 0 {
		for k := range m.F {
			data[i] = 0xa
			i++
			v := m.F[k]
			mapSize := 1 + len(k) + sovNopackage(uint64(len(k))) + 1 + 8
			i = encodeVarintNopackage(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintNopackage(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x11
			i++
			i = encodeFixed64Nopackage(data, i, uint64(math.Float64bits(float64(v))))
		}
	}
	return i, nil
}

func encodeFixed64Nopackage(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Nopackage(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintNopackage(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *M) Size() (n int) {
	var l int
	_ = l
	if len(m.F) > 0 {
		for k, v := range m.F {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovNopackage(uint64(len(k))) + 1 + 8
			n += mapEntrySize + 1 + sovNopackage(uint64(mapEntrySize))
		}
	}
	return n
}

func sovNopackage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNopackage(x uint64) (n int) {
	return sovNopackage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *M) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNopackage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: M: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: M: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field F", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNopackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNopackage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNopackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNopackage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthNopackage
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.F == nil {
				m.F = make(map[string]float64)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNopackage
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapvaluetemp uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				iNdEx += 8
				mapvaluetemp = uint64(data[iNdEx-8])
				mapvaluetemp |= uint64(data[iNdEx-7]) << 8
				mapvaluetemp |= uint64(data[iNdEx-6]) << 16
				mapvaluetemp |= uint64(data[iNdEx-5]) << 24
				mapvaluetemp |= uint64(data[iNdEx-4]) << 32
				mapvaluetemp |= uint64(data[iNdEx-3]) << 40
				mapvaluetemp |= uint64(data[iNdEx-2]) << 48
				mapvaluetemp |= uint64(data[iNdEx-1]) << 56
				mapvalue := math.Float64frombits(mapvaluetemp)
				m.F[mapkey] = mapvalue
			} else {
				var mapvalue float64
				m.F[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNopackage(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNopackage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNopackage(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNopackage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNopackage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNopackage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthNopackage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNopackage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNopackage(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNopackage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNopackage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("nopackage.proto", fileDescriptorNopackage) }

var fileDescriptorNopackage = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xcb, 0x2f, 0x48,
	0x4c, 0xce, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x0a, 0xe2, 0x62, 0xf4,
	0x15, 0x12, 0xe7, 0x62, 0x4c, 0x93, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd4, 0xf3, 0xd5,
	0x73, 0x73, 0xcd, 0x2b, 0x29, 0xaa, 0x0c, 0x62, 0x4c, 0x93, 0x32, 0xe1, 0x62, 0x83, 0x70, 0x84,
	0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21,
	0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xc6, 0x20, 0x08,
	0xc7, 0x8a, 0xc9, 0x82, 0xd1, 0x89, 0xe7, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x4c, 0x62, 0x03, 0x5b, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x62, 0x62,
	0xb2, 0xed, 0x7b, 0x00, 0x00, 0x00,
}
