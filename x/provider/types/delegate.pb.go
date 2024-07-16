// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/provider/v1beta1/delegate.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// BondStatus is the status of a validator.
type BondStatus int32

const (
	// UNSPECIFIED defines an invalid validator status.
	Unspecified BondStatus = 0
	// UNBONDED defines a validator that is not bonded.
	Unbonded BondStatus = 1
	// UNBONDING defines a validator that is unbonding.
	Unbonding BondStatus = 2
	// BONDED defines a validator that is bonded.
	Bonded BondStatus = 3
)

var BondStatus_name = map[int32]string{
	0: "BOND_STATUS_UNSPECIFIED",
	1: "BOND_STATUS_UNBONDED",
	2: "BOND_STATUS_UNBONDING",
	3: "BOND_STATUS_BONDED",
}

var BondStatus_value = map[string]int32{
	"BOND_STATUS_UNSPECIFIED": 0,
	"BOND_STATUS_UNBONDED":    1,
	"BOND_STATUS_UNBONDING":   2,
	"BOND_STATUS_BONDED":      3,
}

func (x BondStatus) String() string {
	return proto.EnumName(BondStatus_name, int32(x))
}

func (BondStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cd75e4a5b29b30f7, []int{0}
}

// Params defines the parameters for the x/meshsecurity module.
type Depositors struct {
	Address string                                   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Tokens  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=tokens,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"tokens"`
}

func (m *Depositors) Reset()         { *m = Depositors{} }
func (m *Depositors) String() string { return proto.CompactTextString(m) }
func (*Depositors) ProtoMessage()    {}
func (*Depositors) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd75e4a5b29b30f7, []int{0}
}
func (m *Depositors) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Depositors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Depositors.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Depositors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Depositors.Merge(m, src)
}
func (m *Depositors) XXX_Size() int {
	return m.Size()
}
func (m *Depositors) XXX_DiscardUnknown() {
	xxx_messageInfo_Depositors.DiscardUnknown(m)
}

var xxx_messageInfo_Depositors proto.InternalMessageInfo

// vault-staker
type Intermediary struct {
	ConsumerValidator string      `protobuf:"bytes,1,opt,name=consumer_validator,json=consumerValidator,proto3" json:"consumer_validator,omitempty"`
	ChainId           string      `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	ContractAddress   string      `protobuf:"bytes,3,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Jailed            bool        `protobuf:"varint,4,opt,name=jailed,proto3" json:"jailed,omitempty"`
	Tombstoned        bool        `protobuf:"varint,5,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	Status            BondStatus  `protobuf:"varint,6,opt,name=status,proto3,enum=osmosis.provider.v1beta1.BondStatus" json:"status,omitempty"`
	Token             *types.Coin `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *Intermediary) Reset()         { *m = Intermediary{} }
func (m *Intermediary) String() string { return proto.CompactTextString(m) }
func (*Intermediary) ProtoMessage()    {}
func (*Intermediary) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd75e4a5b29b30f7, []int{1}
}
func (m *Intermediary) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Intermediary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Intermediary.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Intermediary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Intermediary.Merge(m, src)
}
func (m *Intermediary) XXX_Size() int {
	return m.Size()
}
func (m *Intermediary) XXX_DiscardUnknown() {
	xxx_messageInfo_Intermediary.DiscardUnknown(m)
}

var xxx_messageInfo_Intermediary proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("osmosis.provider.v1beta1.BondStatus", BondStatus_name, BondStatus_value)
	proto.RegisterType((*Depositors)(nil), "osmosis.provider.v1beta1.Depositors")
	proto.RegisterType((*Intermediary)(nil), "osmosis.provider.v1beta1.Intermediary")
}

func init() {
	proto.RegisterFile("osmosis/provider/v1beta1/delegate.proto", fileDescriptor_cd75e4a5b29b30f7)
}

var fileDescriptor_cd75e4a5b29b30f7 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xbd, 0x69, 0xeb, 0xb6, 0xdb, 0x42, 0xd3, 0x55, 0x01, 0xd7, 0x07, 0xd7, 0xaa, 0x10,
	0x98, 0x8a, 0xd8, 0x6a, 0x11, 0x17, 0xc4, 0xa5, 0x69, 0x0a, 0xca, 0xa5, 0xa0, 0xa4, 0xe1, 0x80,
	0x90, 0xa2, 0xb5, 0x77, 0x49, 0x96, 0xc6, 0xbb, 0x91, 0x77, 0x13, 0x91, 0x37, 0x40, 0x39, 0xf5,
	0x05, 0x72, 0xe2, 0x52, 0xf5, 0xc4, 0x89, 0x67, 0xc8, 0xb1, 0x47, 0x4e, 0x7c, 0x24, 0x07, 0x5e,
	0x03, 0x65, 0x6d, 0xd3, 0x20, 0x04, 0x17, 0x7b, 0x3e, 0x7e, 0x33, 0xfa, 0xcf, 0x78, 0x0c, 0xef,
	0x0b, 0x19, 0x0b, 0xc9, 0x64, 0xd0, 0x4d, 0x44, 0x9f, 0x11, 0x9a, 0x04, 0xfd, 0xfd, 0x90, 0x2a,
	0xbc, 0x1f, 0x10, 0xda, 0xa1, 0x2d, 0xac, 0xa8, 0xdf, 0x4d, 0x84, 0x12, 0xc8, 0xca, 0x40, 0x3f,
	0x07, 0xfd, 0x0c, 0xb4, 0x9d, 0x48, 0xa7, 0x82, 0x10, 0x4b, 0xfa, 0xbb, 0x3a, 0x12, 0x8c, 0xa7,
	0x95, 0xf6, 0x56, 0x4b, 0xb4, 0x84, 0x36, 0x83, 0x99, 0x95, 0x45, 0x37, 0x71, 0xcc, 0xb8, 0x08,
	0xf4, 0x33, 0x0d, 0xed, 0x9e, 0x03, 0x08, 0x2b, 0xb4, 0x2b, 0x24, 0x53, 0x22, 0x91, 0xc8, 0x82,
	0xcb, 0x98, 0x90, 0x84, 0x4a, 0x69, 0x01, 0x17, 0x78, 0xab, 0xb5, 0xdc, 0x45, 0x6d, 0x68, 0x2a,
	0x71, 0x46, 0xb9, 0xb4, 0x0a, 0xee, 0x82, 0xb7, 0x76, 0xb0, 0xed, 0xa7, 0x12, 0xfc, 0x99, 0x84,
	0x5c, 0x97, 0x7f, 0x24, 0x18, 0x2f, 0x3f, 0x1e, 0x7f, 0xdd, 0x31, 0x2e, 0xbf, 0xed, 0x78, 0x2d,
	0xa6, 0xda, 0xbd, 0xd0, 0x8f, 0x44, 0x1c, 0x64, 0x7a, 0xd3, 0x57, 0x49, 0x92, 0xb3, 0x40, 0x0d,
	0xba, 0x54, 0xea, 0x02, 0x79, 0xf1, 0xf3, 0xd3, 0x1e, 0xa8, 0x65, 0xfd, 0x77, 0x2f, 0x0b, 0x70,
	0xbd, 0xca, 0x15, 0x4d, 0x62, 0x4a, 0x18, 0x4e, 0x06, 0xa8, 0x04, 0x51, 0x24, 0xb8, 0xec, 0xc5,
	0x34, 0x69, 0xf6, 0x71, 0x87, 0x11, 0xac, 0x44, 0x92, 0xe9, 0xdb, 0xcc, 0x33, 0xaf, 0xf2, 0x04,
	0xda, 0x86, 0x2b, 0x51, 0x1b, 0x33, 0xde, 0x64, 0xc4, 0x2a, 0xa4, 0x43, 0x68, 0xbf, 0x4a, 0xd0,
	0x03, 0x58, 0x8c, 0x04, 0x57, 0x09, 0x8e, 0x54, 0x33, 0x9f, 0x73, 0x41, 0x23, 0x1b, 0x79, 0xfc,
	0x30, 0x9b, 0xf7, 0x36, 0x34, 0xdf, 0x61, 0xd6, 0xa1, 0xc4, 0x5a, 0x74, 0x81, 0xb7, 0x52, 0xcb,
	0x3c, 0xe4, 0x40, 0xa8, 0x44, 0x1c, 0x4a, 0x25, 0x38, 0x25, 0xd6, 0x92, 0xce, 0xcd, 0x45, 0xd0,
	0x53, 0x68, 0x4a, 0x85, 0x55, 0x4f, 0x5a, 0xa6, 0x0b, 0xbc, 0x9b, 0x07, 0x77, 0xfd, 0x7f, 0x7d,
	0x44, 0xbf, 0x2c, 0x38, 0xa9, 0x6b, 0xb6, 0x96, 0xd5, 0xa0, 0x00, 0x2e, 0xe9, 0x2d, 0x58, 0xcb,
	0x2e, 0xf8, 0xef, 0x92, 0x6b, 0x29, 0xb7, 0xf7, 0x19, 0x40, 0x78, 0xdd, 0x07, 0x3d, 0x84, 0x77,
	0xca, 0x2f, 0x4e, 0x2a, 0xcd, 0xfa, 0xe9, 0xe1, 0x69, 0xa3, 0xde, 0x6c, 0x9c, 0xd4, 0x5f, 0x1e,
	0x1f, 0x55, 0x9f, 0x55, 0x8f, 0x2b, 0x45, 0xc3, 0xde, 0x18, 0x8e, 0xdc, 0xb5, 0x06, 0x97, 0x5d,
	0x1a, 0xb1, 0xb7, 0x8c, 0x12, 0x74, 0x0f, 0x6e, 0xfd, 0x49, 0xcf, 0xbc, 0xe3, 0x4a, 0x11, 0xd8,
	0xeb, 0xc3, 0x91, 0xbb, 0xd2, 0xe0, 0xa1, 0xe0, 0x84, 0x12, 0xe4, 0xc1, 0x5b, 0x7f, 0x73, 0xd5,
	0x93, 0xe7, 0xc5, 0x82, 0x7d, 0x63, 0x38, 0x72, 0x57, 0x53, 0x90, 0xf1, 0x16, 0xda, 0x85, 0x68,
	0x9e, 0xcc, 0xfa, 0x2d, 0xd8, 0x70, 0x38, 0x72, 0xcd, 0xb2, 0xee, 0x66, 0x2f, 0x7e, 0xf8, 0xe8,
	0x18, 0xe5, 0x37, 0xe3, 0x1f, 0x8e, 0x71, 0x31, 0x71, 0x8c, 0xf1, 0xc4, 0x01, 0x57, 0x13, 0x07,
	0x7c, 0x9f, 0x38, 0xe0, 0x7c, 0xea, 0x18, 0x57, 0x53, 0xc7, 0xf8, 0x32, 0x75, 0x8c, 0xd7, 0x4f,
	0xe6, 0xce, 0x27, 0xdb, 0x61, 0xa9, 0x83, 0x43, 0x19, 0xc4, 0x54, 0xb6, 0x4b, 0x92, 0x46, 0xbd,
	0x84, 0xa9, 0x81, 0xbe, 0xa5, 0xf7, 0xd7, 0xbf, 0x92, 0x3e, 0xab, 0xd0, 0xd4, 0xd7, 0xfd, 0xe8,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x83, 0xdd, 0x14, 0x3d, 0x6b, 0x03, 0x00, 0x00,
}

func (m *Depositors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Depositors) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Depositors) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tokens) > 0 {
		for iNdEx := len(m.Tokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDelegate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintDelegate(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Intermediary) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Intermediary) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Intermediary) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Token != nil {
		{
			size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDelegate(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Status != 0 {
		i = encodeVarintDelegate(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if m.Tombstoned {
		i--
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintDelegate(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintDelegate(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConsumerValidator) > 0 {
		i -= len(m.ConsumerValidator)
		copy(dAtA[i:], m.ConsumerValidator)
		i = encodeVarintDelegate(dAtA, i, uint64(len(m.ConsumerValidator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelegate(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Depositors) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovDelegate(uint64(l))
	}
	if len(m.Tokens) > 0 {
		for _, e := range m.Tokens {
			l = e.Size()
			n += 1 + l + sovDelegate(uint64(l))
		}
	}
	return n
}

func (m *Intermediary) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConsumerValidator)
	if l > 0 {
		n += 1 + l + sovDelegate(uint64(l))
	}
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovDelegate(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovDelegate(uint64(l))
	}
	if m.Jailed {
		n += 2
	}
	if m.Tombstoned {
		n += 2
	}
	if m.Status != 0 {
		n += 1 + sovDelegate(uint64(m.Status))
	}
	if m.Token != nil {
		l = m.Token.Size()
		n += 1 + l + sovDelegate(uint64(l))
	}
	return n
}

func sovDelegate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegate(x uint64) (n int) {
	return sovDelegate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Depositors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegate
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
			return fmt.Errorf("proto: Depositors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Depositors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
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
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
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
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tokens = append(m.Tokens, types.Coin{})
			if err := m.Tokens[len(m.Tokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegate
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
func (m *Intermediary) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegate
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
			return fmt.Errorf("proto: Intermediary: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Intermediary: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsumerValidator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
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
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsumerValidator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
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
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
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
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstoned = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= BondStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegate
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
				return ErrInvalidLengthDelegate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelegate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Token == nil {
				m.Token = &types.Coin{}
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegate
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
func skipDelegate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegate
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
					return 0, ErrIntOverflowDelegate
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
					return 0, ErrIntOverflowDelegate
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
				return 0, ErrInvalidLengthDelegate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegate = fmt.Errorf("proto: unexpected end of group")
)