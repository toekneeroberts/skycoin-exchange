// Code generated by protoc-gen-go.
// source: pp.admin.proto
// DO NOT EDIT!

package pp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UpdateCreditReq struct {
	Pubkey           *string `protobuf:"bytes,10,opt,name=pubkey" json:"pubkey,omitempty"`
	CoinType         *string `protobuf:"bytes,20,opt,name=coin_type" json:"coin_type,omitempty"`
	Amount           *uint64 `protobuf:"varint,30,opt,name=amount" json:"amount,omitempty"`
	Dst              *string `protobuf:"bytes,40,opt,name=dst" json:"dst,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateCreditReq) Reset()                    { *m = UpdateCreditReq{} }
func (m *UpdateCreditReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateCreditReq) ProtoMessage()               {}
func (*UpdateCreditReq) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *UpdateCreditReq) GetPubkey() string {
	if m != nil && m.Pubkey != nil {
		return *m.Pubkey
	}
	return ""
}

func (m *UpdateCreditReq) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *UpdateCreditReq) GetAmount() uint64 {
	if m != nil && m.Amount != nil {
		return *m.Amount
	}
	return 0
}

func (m *UpdateCreditReq) GetDst() string {
	if m != nil && m.Dst != nil {
		return *m.Dst
	}
	return ""
}

type UpdateCreditRes struct {
	Result           *Result `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateCreditRes) Reset()                    { *m = UpdateCreditRes{} }
func (m *UpdateCreditRes) String() string            { return proto.CompactTextString(m) }
func (*UpdateCreditRes) ProtoMessage()               {}
func (*UpdateCreditRes) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *UpdateCreditRes) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateCreditReq)(nil), "pp.UpdateCreditReq")
	proto.RegisterType((*UpdateCreditRes)(nil), "pp.UpdateCreditRes")
}

func init() { proto.RegisterFile("pp.admin.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0xcc, 0xb1, 0xaa, 0x83, 0x30,
	0x14, 0x87, 0x71, 0xf4, 0x5e, 0x04, 0x8f, 0xa0, 0x34, 0x74, 0x08, 0x0e, 0x45, 0x9c, 0xb2, 0x34,
	0x43, 0x5f, 0xa1, 0x4f, 0x50, 0xa1, 0x73, 0xb1, 0x26, 0x83, 0xb4, 0xf1, 0xfc, 0xab, 0x27, 0x83,
	0x6f, 0x5f, 0x2a, 0x9d, 0xba, 0xfe, 0x3e, 0xf8, 0xa8, 0x04, 0x6c, 0xef, 0xc2, 0x38, 0x59, 0xcc,
	0x2c, 0xac, 0x52, 0xa0, 0xae, 0x00, 0x3b, 0x70, 0x08, 0xfc, 0xc5, 0xf6, 0x42, 0xd5, 0x15, 0xae,
	0x17, 0x7f, 0x9e, 0xbd, 0x1b, 0xa5, 0xf3, 0x2f, 0x55, 0x52, 0x86, 0x78, 0x7f, 0xf8, 0x55, 0x53,
	0x93, 0x98, 0x5c, 0xed, 0x28, 0x1f, 0x78, 0x9c, 0x6e, 0xb2, 0xc2, 0xeb, 0xfd, 0x46, 0x25, 0x65,
	0x7d, 0xe0, 0x38, 0x89, 0x3e, 0x34, 0x89, 0xf9, 0x57, 0x05, 0xfd, 0xb9, 0x45, 0xb4, 0xf9, 0xc4,
	0xf6, 0xf8, 0xbb, 0x5c, 0x54, 0x4d, 0xd9, 0xec, 0x97, 0xf8, 0x14, 0x9d, 0x34, 0xa9, 0x29, 0x4e,
	0x64, 0x01, 0xdb, 0x6d, 0xf2, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb5, 0x78, 0x04, 0x24, 0xa7, 0x00,
	0x00, 0x00,
}
