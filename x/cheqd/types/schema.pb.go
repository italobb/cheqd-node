// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cheqd/schema.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Schema struct {
	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Authors   []string `protobuf:"bytes,2,rep,name=authors,proto3" json:"authors,omitempty"`
	Name      string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Version   string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	AttrNames []string `protobuf:"bytes,5,rep,name=attrNames,json=attrNames,proto3" json:"attrNames,omitempty"`
}

func (m *Schema) Reset()         { *m = Schema{} }
func (m *Schema) String() string { return proto.CompactTextString(m) }
func (*Schema) ProtoMessage()    {}
func (*Schema) Descriptor() ([]byte, []int) {
	return fileDescriptor_96d8c360e747a40a, []int{0}
}
func (m *Schema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Schema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Schema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Schema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schema.Merge(m, src)
}
func (m *Schema) XXX_Size() int {
	return m.Size()
}
func (m *Schema) XXX_DiscardUnknown() {
	xxx_messageInfo_Schema.DiscardUnknown(m)
}

var xxx_messageInfo_Schema proto.InternalMessageInfo

func (m *Schema) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Schema) GetAuthors() []string {
	if m != nil {
		return m.Authors
	}
	return nil
}

func (m *Schema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Schema) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Schema) GetAttrNames() []string {
	if m != nil {
		return m.AttrNames
	}
	return nil
}

func init() {
	proto.RegisterType((*Schema)(nil), "cheqdid.cheqdnode.cheqd.Schema")
}

func init() { proto.RegisterFile("cheqd/schema.proto", fileDescriptor_96d8c360e747a40a) }

var fileDescriptor_96d8c360e747a40a = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0x48, 0x2d,
	0x4c, 0xd1, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x07, 0x8b, 0x65, 0xa6, 0xe8, 0x81, 0xe9, 0xbc, 0xfc, 0x94, 0x54, 0x08, 0x4b, 0x4a, 0x24, 0x3d,
	0x3f, 0x3d, 0x1f, 0xac, 0x46, 0x1f, 0xc4, 0x82, 0x28, 0x57, 0xaa, 0xe7, 0x62, 0x0b, 0x06, 0x6b,
	0x17, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c,
	0x11, 0x92, 0xe0, 0x62, 0x4f, 0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0x2a, 0x96, 0x60, 0x52, 0x60, 0xd6,
	0xe0, 0x0c, 0x82, 0x71, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0xc1, 0x6a,
	0xc1, 0x6c, 0x90, 0xea, 0xb2, 0xd4, 0xa2, 0xe2, 0xcc, 0xfc, 0x3c, 0x09, 0x16, 0xb0, 0x30, 0x8c,
	0x2b, 0x24, 0xcb, 0xc5, 0x95, 0x58, 0x52, 0x52, 0x14, 0x0f, 0x52, 0x56, 0x2c, 0xc1, 0x0a, 0x36,
	0x8a, 0x13, 0x24, 0xe2, 0x07, 0x12, 0x70, 0x72, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0xcd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x88,
	0x47, 0xc1, 0xa4, 0x2e, 0xc8, 0x4f, 0xfa, 0x15, 0x50, 0xa1, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24,
	0x36, 0xb0, 0x67, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x3f, 0x5e, 0x48, 0x11, 0x01,
	0x00, 0x00,
}

func (m *Schema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Schema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Schema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AttrNames) > 0 {
		for iNdEx := len(m.AttrNames) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AttrNames[iNdEx])
			copy(dAtA[i:], m.AttrNames[iNdEx])
			i = encodeVarintSchema(dAtA, i, uint64(len(m.AttrNames[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Authors) > 0 {
		for iNdEx := len(m.Authors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Authors[iNdEx])
			copy(dAtA[i:], m.Authors[iNdEx])
			i = encodeVarintSchema(dAtA, i, uint64(len(m.Authors[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSchema(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Schema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if len(m.Authors) > 0 {
		for _, s := range m.Authors {
			l = len(s)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovSchema(uint64(l))
	}
	if len(m.AttrNames) > 0 {
		for _, s := range m.AttrNames {
			l = len(s)
			n += 1 + l + sovSchema(uint64(l))
		}
	}
	return n
}

func sovSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSchema(x uint64) (n int) {
	return sovSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Schema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Schema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Schema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authors = append(m.Authors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttrNames", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AttrNames = append(m.AttrNames, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSchema
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
func skipSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSchema
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
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
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSchema = fmt.Errorf("proto: unexpected end of group")
)
