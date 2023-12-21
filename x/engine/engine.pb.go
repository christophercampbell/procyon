// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: polygon/engine/v1/engine.proto

package engine

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// State is the current external engine state
type State struct {
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	NumTxs int64 `protobuf:"varint,2,opt,name=num_txs,json=numTxs,proto3" json:"num_txs,omitempty"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5e50b763ec0bcba, []int{0}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_State.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return m.Size()
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *State) GetNumTxs() int64 {
	if m != nil {
		return m.NumTxs
	}
	return 0
}

// EngineBlock is an array of external engine transaction byte arrays
type EngineBlock struct {
	Txs [][]byte `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
}

func (m *EngineBlock) Reset()         { *m = EngineBlock{} }
func (m *EngineBlock) String() string { return proto.CompactTextString(m) }
func (*EngineBlock) ProtoMessage()    {}
func (*EngineBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5e50b763ec0bcba, []int{1}
}
func (m *EngineBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EngineBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EngineBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EngineBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EngineBlock.Merge(m, src)
}
func (m *EngineBlock) XXX_Size() int {
	return m.Size()
}
func (m *EngineBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_EngineBlock.DiscardUnknown(m)
}

var xxx_messageInfo_EngineBlock proto.InternalMessageInfo

func (m *EngineBlock) GetTxs() [][]byte {
	if m != nil {
		return m.Txs
	}
	return nil
}

func init() {
	proto.RegisterType((*State)(nil), "polygon.engine.v1.State")
	proto.RegisterType((*EngineBlock)(nil), "polygon.engine.v1.EngineBlock")
}

func init() { proto.RegisterFile("polygon/engine/v1/engine.proto", fileDescriptor_a5e50b763ec0bcba) }

var fileDescriptor_a5e50b763ec0bcba = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xc8, 0xcf, 0xa9,
	0x4c, 0xcf, 0xcf, 0xd3, 0x4f, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0x2f, 0x33, 0x84, 0xb2, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0xa1, 0xf2, 0x7a, 0x50, 0xd1, 0x32, 0x43, 0x25, 0x0b,
	0x2e, 0xd6, 0xe0, 0x92, 0xc4, 0x92, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0x8c, 0xd4, 0xcc, 0xf4, 0x8c,
	0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x28, 0x4f, 0x48, 0x9c, 0x8b, 0x3d, 0xaf, 0x34,
	0x37, 0xbe, 0xa4, 0xa2, 0x58, 0x82, 0x09, 0x22, 0x91, 0x57, 0x9a, 0x1b, 0x52, 0x51, 0xac, 0x24,
	0xcf, 0xc5, 0xed, 0x0a, 0x36, 0xc6, 0x29, 0x27, 0x3f, 0x39, 0x5b, 0x48, 0x80, 0x8b, 0x19, 0xa4,
	0x86, 0x51, 0x81, 0x59, 0x83, 0x27, 0x08, 0xc4, 0x74, 0x32, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1,
	0xc6, 0x63, 0x39, 0x86, 0x28, 0xd9, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c,
	0x7d, 0x83, 0x8a, 0x00, 0xa8, 0xa3, 0x2b, 0xa0, 0x8e, 0x4d, 0x62, 0x03, 0xbb, 0xd6, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xef, 0x31, 0x63, 0x48, 0xcf, 0x00, 0x00, 0x00,
}

func (m *State) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *State) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *State) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumTxs != 0 {
		i = encodeVarintEngine(dAtA, i, uint64(m.NumTxs))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintEngine(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EngineBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EngineBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EngineBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Txs[iNdEx])
			copy(dAtA[i:], m.Txs[iNdEx])
			i = encodeVarintEngine(dAtA, i, uint64(len(m.Txs[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintEngine(dAtA []byte, offset int, v uint64) int {
	offset -= sovEngine(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *State) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovEngine(uint64(m.Height))
	}
	if m.NumTxs != 0 {
		n += 1 + sovEngine(uint64(m.NumTxs))
	}
	return n
}

func (m *EngineBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, b := range m.Txs {
			l = len(b)
			n += 1 + l + sovEngine(uint64(l))
		}
	}
	return n
}

func sovEngine(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEngine(x uint64) (n int) {
	return sovEngine(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *State) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEngine
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
			return fmt.Errorf("proto: State: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: State: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEngine
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumTxs", wireType)
			}
			m.NumTxs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEngine
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumTxs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEngine(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEngine
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
func (m *EngineBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEngine
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
			return fmt.Errorf("proto: EngineBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EngineBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEngine
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEngine
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEngine
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, make([]byte, postIndex-iNdEx))
			copy(m.Txs[len(m.Txs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEngine(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEngine
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
func skipEngine(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEngine
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
					return 0, ErrIntOverflowEngine
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
					return 0, ErrIntOverflowEngine
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
				return 0, ErrInvalidLengthEngine
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEngine
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEngine
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEngine        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEngine          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEngine = fmt.Errorf("proto: unexpected end of group")
)