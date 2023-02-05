// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: htlc/htlc.proto

package htlc

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// HTLCStatus defines the state of an OldHTLC
type HTLCStatus int32

const (
	// HTLC_STATE_OPEN defines an open state.
	Open HTLCStatus = 0
	// HTLC_STATE_COMPLETED defines a completed state.
	Completed HTLCStatus = 1
	// HTLC_STATE_EXPIRED defines an expired state.
	Expired HTLCStatus = 2
	// HTLC_STATE_REFUNDED defines a refunded state.
	Refunded HTLCStatus = 3
)

var HTLCStatus_name = map[int32]string{
	0: "HTLC_STATE_OPEN",
	1: "HTLC_STATE_COMPLETED",
	2: "HTLC_STATE_EXPIRED",
	3: "HTLC_STATE_REFUNDED",
}

var HTLCStatus_value = map[string]int32{
	"HTLC_STATE_OPEN":      0,
	"HTLC_STATE_COMPLETED": 1,
	"HTLC_STATE_EXPIRED":   2,
	"HTLC_STATE_REFUNDED":  3,
}

func (x HTLCStatus) String() string {
	return proto.EnumName(HTLCStatus_name, int32(x))
}

// OldHTLC defines the struct of an OldHTLC
type OldHTLC struct {
	Sender               string                                   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	To                   string                                   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	ReceiverOnOtherChain string                                   `protobuf:"bytes,3,opt,name=receiver_on_other_chain,json=receiverOnOtherChain,proto3" json:"receiver_on_other_chain,omitempty" yaml:"receiver_on_other_chain"`
	Amount               github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	Secret               string                                   `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
	Timestamp            uint64                                   `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ExpirationHeight     uint64                                   `protobuf:"varint,7,opt,name=expiration_height,json=expirationHeight,proto3" json:"expiration_height,omitempty" yaml:"expiration_height"`
	State                HTLCStatus                               `protobuf:"varint,8,opt,name=state,proto3,enum=gridmod.htlc.HTLCStatus" json:"state,omitempty"`
}

func (m *OldHTLC) Reset()         { *m = OldHTLC{} }
func (m *OldHTLC) String() string { return proto.CompactTextString(m) }
func (*OldHTLC) ProtoMessage()    {}

func (m *OldHTLC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OldHTLC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HTLC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OldHTLC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTLC.Merge(m, src)
}
func (m *OldHTLC) XXX_Size() int {
	return m.Size()
}
func (m *OldHTLC) XXX_DiscardUnknown() {
	xxx_messageInfo_HTLC.DiscardUnknown(m)
}

var xxx_messageInfo_HTLC proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("gridmod.htlc.HTLCStatus", HTLCStatus_name, HTLCStatus_value)
	proto.RegisterType((*OldHTLC)(nil), "gridmod.htlc.OldHTLC")
}

func (this *OldHTLC) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*OldHTLC)
	if !ok {
		that2, ok := that.(OldHTLC)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Sender != that1.Sender {
		return false
	}
	if this.To != that1.To {
		return false
	}
	if this.ReceiverOnOtherChain != that1.ReceiverOnOtherChain {
		return false
	}
	if len(this.Amount) != len(that1.Amount) {
		return false
	}
	for i := range this.Amount {
		if !this.Amount[i].Equal(&that1.Amount[i]) {
			return false
		}
	}
	if this.Secret != that1.Secret {
		return false
	}
	if this.Timestamp != that1.Timestamp {
		return false
	}
	if this.ExpirationHeight != that1.ExpirationHeight {
		return false
	}
	if this.State != that1.State {
		return false
	}
	return true
}
func (m *OldHTLC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OldHTLC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OldHTLC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		i = encodeVarintHtlc(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x40
	}
	if m.ExpirationHeight != 0 {
		i = encodeVarintHtlc(dAtA, i, uint64(m.ExpirationHeight))
		i--
		dAtA[i] = 0x38
	}
	if m.Timestamp != 0 {
		i = encodeVarintHtlc(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintHtlc(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHtlc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ReceiverOnOtherChain) > 0 {
		i -= len(m.ReceiverOnOtherChain)
		copy(dAtA[i:], m.ReceiverOnOtherChain)
		i = encodeVarintHtlc(dAtA, i, uint64(len(m.ReceiverOnOtherChain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintHtlc(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintHtlc(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHtlc(dAtA []byte, offset int, v uint64) int {
	offset -= sovHtlc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OldHTLC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovHtlc(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovHtlc(uint64(l))
	}
	l = len(m.ReceiverOnOtherChain)
	if l > 0 {
		n += 1 + l + sovHtlc(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovHtlc(uint64(l))
		}
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovHtlc(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovHtlc(uint64(m.Timestamp))
	}
	if m.ExpirationHeight != 0 {
		n += 1 + sovHtlc(uint64(m.ExpirationHeight))
	}
	if m.State != 0 {
		n += 1 + sovHtlc(uint64(m.State))
	}
	return n
}

func sovHtlc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHtlc(x uint64) (n int) {
	return sovHtlc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OldHTLC) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHtlc
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
			return fmt.Errorf("proto: OldHTLC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OldHTLC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
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
				return ErrInvalidLengthHtlc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHtlc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
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
				return ErrInvalidLengthHtlc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHtlc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.To = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceiverOnOtherChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
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
				return ErrInvalidLengthHtlc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHtlc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceiverOnOtherChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHtlc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHtlc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
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
				return ErrInvalidLengthHtlc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHtlc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationHeight", wireType)
			}
			m.ExpirationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHtlc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= HTLCStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHtlc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHtlc
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
func skipHtlc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHtlc
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
					return 0, ErrIntOverflowHtlc
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
					return 0, ErrIntOverflowHtlc
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
				return 0, ErrInvalidLengthHtlc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHtlc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHtlc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHtlc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHtlc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHtlc = fmt.Errorf("proto: unexpected end of group")
)
