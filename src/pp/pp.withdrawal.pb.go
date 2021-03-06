// Code generated by protoc-gen-go.
// source: pp.withdrawal.proto
// DO NOT EDIT!

package pp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WithdrawalReq struct {
	Pubkey           *string `protobuf:"bytes,10,opt,name=pubkey" json:"pubkey,omitempty"`
	CoinType         *string `protobuf:"bytes,11,opt,name=coin_type" json:"coin_type,omitempty"`
	Coins            *uint64 `protobuf:"varint,12,opt,name=coins" json:"coins,omitempty"`
	OutputAddress    *string `protobuf:"bytes,13,opt,name=output_address" json:"output_address,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WithdrawalReq) Reset()                    { *m = WithdrawalReq{} }
func (m *WithdrawalReq) String() string            { return proto.CompactTextString(m) }
func (*WithdrawalReq) ProtoMessage()               {}
func (*WithdrawalReq) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *WithdrawalReq) GetPubkey() string {
	if m != nil && m.Pubkey != nil {
		return *m.Pubkey
	}
	return ""
}

func (m *WithdrawalReq) GetCoinType() string {
	if m != nil && m.CoinType != nil {
		return *m.CoinType
	}
	return ""
}

func (m *WithdrawalReq) GetCoins() uint64 {
	if m != nil && m.Coins != nil {
		return *m.Coins
	}
	return 0
}

func (m *WithdrawalReq) GetOutputAddress() string {
	if m != nil && m.OutputAddress != nil {
		return *m.OutputAddress
	}
	return ""
}

type WithdrawalRes struct {
	Result           *Result `protobuf:"bytes,1,req,name=result" json:"result,omitempty"`
	NewTxid          *string `protobuf:"bytes,20,opt,name=new_txid" json:"new_txid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WithdrawalRes) Reset()                    { *m = WithdrawalRes{} }
func (m *WithdrawalRes) String() string            { return proto.CompactTextString(m) }
func (*WithdrawalRes) ProtoMessage()               {}
func (*WithdrawalRes) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *WithdrawalRes) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *WithdrawalRes) GetNewTxid() string {
	if m != nil && m.NewTxid != nil {
		return *m.NewTxid
	}
	return ""
}

func init() {
	proto.RegisterType((*WithdrawalReq)(nil), "pp.WithdrawalReq")
	proto.RegisterType((*WithdrawalRes)(nil), "pp.WithdrawalRes")
}

func init() { proto.RegisterFile("pp.withdrawal.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0xcc, 0xc1, 0xaa, 0x82, 0x40,
	0x14, 0xc6, 0x71, 0x94, 0x7b, 0x25, 0x8f, 0x69, 0x35, 0x45, 0x0c, 0xae, 0xc4, 0x95, 0xab, 0x59,
	0xb4, 0xef, 0x25, 0xdc, 0x44, 0x2b, 0x31, 0x67, 0x20, 0x49, 0x9d, 0xd3, 0xcc, 0x19, 0xcc, 0xb7,
	0x0f, 0x0d, 0x82, 0xb6, 0x3f, 0xfe, 0xdf, 0x07, 0x7b, 0x44, 0x31, 0xb6, 0x74, 0x97, 0xa6, 0x1e,
	0xeb, 0x4e, 0xa0, 0xd1, 0xa4, 0x99, 0x8f, 0x98, 0x6e, 0x10, 0x45, 0xa3, 0xfb, 0x5e, 0x0f, 0x1f,
	0xcc, 0xaf, 0x10, 0x5f, 0xbe, 0x61, 0xa9, 0x9e, 0x2c, 0x81, 0x00, 0xdd, 0xed, 0xa1, 0x26, 0x0e,
	0x99, 0x57, 0x84, 0x6c, 0x07, 0x61, 0xa3, 0xdb, 0xa1, 0xa2, 0x09, 0x15, 0x8f, 0x16, 0x8a, 0xe1,
	0x7f, 0x26, 0xcb, 0xd7, 0x99, 0x57, 0xfc, 0xb1, 0x23, 0x24, 0xda, 0x11, 0x3a, 0xaa, 0x6a, 0x29,
	0x8d, 0xb2, 0x96, 0xc7, 0x73, 0x96, 0x9f, 0x7f, 0xaf, 0x2d, 0x4b, 0x21, 0x30, 0xca, 0xba, 0x8e,
	0xb8, 0x97, 0xf9, 0x45, 0x74, 0x02, 0x81, 0x28, 0xca, 0x45, 0xd8, 0x16, 0x56, 0x83, 0x1a, 0x2b,
	0x7a, 0xb5, 0x92, 0x1f, 0xe6, 0xf9, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x04, 0xf5, 0x53, 0xc7, 0xc4,
	0x00, 0x00, 0x00,
}
