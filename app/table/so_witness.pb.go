// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_witness.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoWitness struct {
	Owner                 *prototype.AccountName         `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	WitnessScheduleType   *prototype.WitnessScheduleType `protobuf:"bytes,2,opt,name=witness_schedule_type,json=witnessScheduleType,proto3" json:"witness_schedule_type,omitempty"`
	CreatedTime           *prototype.TimePointSec        `protobuf:"bytes,3,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	Url                   string                         `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	VoteCount             uint64                         `protobuf:"varint,5,opt,name=vote_count,json=voteCount,proto3" json:"vote_count,omitempty"`
	LastConfirmedBlockNum uint32                         `protobuf:"varint,6,opt,name=last_confirmed_block_num,json=lastConfirmedBlockNum,proto3" json:"last_confirmed_block_num,omitempty"`
	TotalMissed           uint32                         `protobuf:"varint,7,opt,name=total_missed,json=totalMissed,proto3" json:"total_missed,omitempty"`
	PowWorker             uint32                         `protobuf:"varint,8,opt,name=pow_worker,json=powWorker,proto3" json:"pow_worker,omitempty"`
	SigningKey            *prototype.PublicKeyType       `protobuf:"bytes,9,opt,name=signing_key,json=signingKey,proto3" json:"signing_key,omitempty"`
	LastWork              *prototype.Sha256              `protobuf:"bytes,10,opt,name=last_work,json=lastWork,proto3" json:"last_work,omitempty"`
	RunningVersion        uint32                         `protobuf:"varint,11,opt,name=running_version,json=runningVersion,proto3" json:"running_version,omitempty"`
	LastAslot             uint32                         `protobuf:"varint,12,opt,name=last_aslot,json=lastAslot,proto3" json:"last_aslot,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                       `json:"-"`
	XXX_unrecognized      []byte                         `json:"-"`
	XXX_sizecache         int32                          `json:"-"`
}

func (m *SoWitness) Reset()         { *m = SoWitness{} }
func (m *SoWitness) String() string { return proto.CompactTextString(m) }
func (*SoWitness) ProtoMessage()    {}
func (*SoWitness) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{0}
}

func (m *SoWitness) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoWitness.Unmarshal(m, b)
}
func (m *SoWitness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoWitness.Marshal(b, m, deterministic)
}
func (m *SoWitness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoWitness.Merge(m, src)
}
func (m *SoWitness) XXX_Size() int {
	return xxx_messageInfo_SoWitness.Size(m)
}
func (m *SoWitness) XXX_DiscardUnknown() {
	xxx_messageInfo_SoWitness.DiscardUnknown(m)
}

var xxx_messageInfo_SoWitness proto.InternalMessageInfo

func (m *SoWitness) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *SoWitness) GetWitnessScheduleType() *prototype.WitnessScheduleType {
	if m != nil {
		return m.WitnessScheduleType
	}
	return nil
}

func (m *SoWitness) GetCreatedTime() *prototype.TimePointSec {
	if m != nil {
		return m.CreatedTime
	}
	return nil
}

func (m *SoWitness) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SoWitness) GetVoteCount() uint64 {
	if m != nil {
		return m.VoteCount
	}
	return 0
}

func (m *SoWitness) GetLastConfirmedBlockNum() uint32 {
	if m != nil {
		return m.LastConfirmedBlockNum
	}
	return 0
}

func (m *SoWitness) GetTotalMissed() uint32 {
	if m != nil {
		return m.TotalMissed
	}
	return 0
}

func (m *SoWitness) GetPowWorker() uint32 {
	if m != nil {
		return m.PowWorker
	}
	return 0
}

func (m *SoWitness) GetSigningKey() *prototype.PublicKeyType {
	if m != nil {
		return m.SigningKey
	}
	return nil
}

func (m *SoWitness) GetLastWork() *prototype.Sha256 {
	if m != nil {
		return m.LastWork
	}
	return nil
}

func (m *SoWitness) GetRunningVersion() uint32 {
	if m != nil {
		return m.RunningVersion
	}
	return 0
}

func (m *SoWitness) GetLastAslot() uint32 {
	if m != nil {
		return m.LastAslot
	}
	return 0
}

type SoListWitnessByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoListWitnessByOwner) Reset()         { *m = SoListWitnessByOwner{} }
func (m *SoListWitnessByOwner) String() string { return proto.CompactTextString(m) }
func (*SoListWitnessByOwner) ProtoMessage()    {}
func (*SoListWitnessByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{1}
}

func (m *SoListWitnessByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoListWitnessByOwner.Unmarshal(m, b)
}
func (m *SoListWitnessByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoListWitnessByOwner.Marshal(b, m, deterministic)
}
func (m *SoListWitnessByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoListWitnessByOwner.Merge(m, src)
}
func (m *SoListWitnessByOwner) XXX_Size() int {
	return xxx_messageInfo_SoListWitnessByOwner.Size(m)
}
func (m *SoListWitnessByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoListWitnessByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoListWitnessByOwner proto.InternalMessageInfo

func (m *SoListWitnessByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

type SoUniqueWitnessByOwner struct {
	Owner                *prototype.AccountName `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SoUniqueWitnessByOwner) Reset()         { *m = SoUniqueWitnessByOwner{} }
func (m *SoUniqueWitnessByOwner) String() string { return proto.CompactTextString(m) }
func (*SoUniqueWitnessByOwner) ProtoMessage()    {}
func (*SoUniqueWitnessByOwner) Descriptor() ([]byte, []int) {
	return fileDescriptor_00097e516fc05425, []int{2}
}

func (m *SoUniqueWitnessByOwner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Unmarshal(m, b)
}
func (m *SoUniqueWitnessByOwner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Marshal(b, m, deterministic)
}
func (m *SoUniqueWitnessByOwner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueWitnessByOwner.Merge(m, src)
}
func (m *SoUniqueWitnessByOwner) XXX_Size() int {
	return xxx_messageInfo_SoUniqueWitnessByOwner.Size(m)
}
func (m *SoUniqueWitnessByOwner) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueWitnessByOwner.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueWitnessByOwner proto.InternalMessageInfo

func (m *SoUniqueWitnessByOwner) GetOwner() *prototype.AccountName {
	if m != nil {
		return m.Owner
	}
	return nil
}

func init() {
	proto.RegisterType((*SoWitness)(nil), "table.so_witness")
	proto.RegisterType((*SoListWitnessByOwner)(nil), "table.so_list_witness_by_owner")
	proto.RegisterType((*SoUniqueWitnessByOwner)(nil), "table.so_unique_witness_by_owner")
}

func init() { proto.RegisterFile("app/table/so_witness.proto", fileDescriptor_00097e516fc05425) }

var fileDescriptor_00097e516fc05425 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4f, 0x6f, 0xd4, 0x30,
	0x10, 0xc5, 0x15, 0xba, 0x5b, 0x1a, 0x67, 0xf9, 0x67, 0xa8, 0x30, 0x2b, 0x21, 0x85, 0x3d, 0x40,
	0x2e, 0x4d, 0xa4, 0x22, 0xe0, 0x00, 0x17, 0xda, 0x13, 0xaa, 0xe0, 0x10, 0x56, 0x20, 0x71, 0xb1,
	0x1c, 0xef, 0xb0, 0x6b, 0x6d, 0xe2, 0x09, 0xb1, 0xdd, 0x28, 0xdf, 0x8a, 0x8f, 0x88, 0xec, 0x04,
	0xd8, 0x03, 0x17, 0xd4, 0x4b, 0x94, 0xfc, 0xe6, 0xbd, 0x37, 0xe3, 0x64, 0x42, 0x96, 0xa2, 0x6d,
	0x0b, 0x2b, 0xaa, 0x1a, 0x0a, 0x83, 0xbc, 0x57, 0x56, 0x83, 0x31, 0x79, 0xdb, 0xa1, 0x45, 0x3a,
	0x0f, 0x7c, 0xf9, 0x28, 0x3c, 0xd9, 0xa1, 0x85, 0xc2, 0x5f, 0xc6, 0xe2, 0xea, 0xe7, 0x8c, 0x90,
	0xbf, 0x0e, 0x7a, 0x46, 0xe6, 0xd8, 0x6b, 0xe8, 0x58, 0x94, 0x46, 0x59, 0x72, 0xfe, 0x38, 0xff,
	0x63, 0xca, 0x85, 0x94, 0xe8, 0xb4, 0xe5, 0x5a, 0x34, 0x50, 0x8e, 0x2a, 0xba, 0x26, 0xa7, 0x93,
	0x93, 0x1b, 0xb9, 0x83, 0x8d, 0xab, 0x81, 0x7b, 0x31, 0xbb, 0x15, 0xec, 0xe9, 0x81, 0xfd, 0x9f,
	0xba, 0xf2, 0xe1, 0x84, 0x3f, 0x4f, 0x74, 0x3d, 0xb4, 0x40, 0xdf, 0x91, 0x85, 0xec, 0x40, 0x58,
	0xd8, 0x70, 0xab, 0x1a, 0x60, 0x47, 0x21, 0xec, 0xc9, 0x41, 0x98, 0xc7, 0xbc, 0x45, 0xa5, 0x2d,
	0x37, 0x20, 0xcb, 0x64, 0x92, 0xaf, 0x55, 0x03, 0xf4, 0x3e, 0x39, 0x72, 0x5d, 0xcd, 0x66, 0x69,
	0x94, 0xc5, 0xa5, 0xbf, 0xa5, 0x4f, 0x09, 0xb9, 0x46, 0x0b, 0x3c, 0xcc, 0xcf, 0xe6, 0x69, 0x94,
	0xcd, 0xca, 0xd8, 0x93, 0x4b, 0x0f, 0xe8, 0x1b, 0xc2, 0x6a, 0x61, 0x2c, 0x97, 0xa8, 0xbf, 0xab,
	0xae, 0x81, 0x0d, 0xaf, 0x6a, 0x94, 0x7b, 0xae, 0x5d, 0xc3, 0x8e, 0xd3, 0x28, 0xbb, 0x53, 0x9e,
	0xfa, 0xfa, 0xe5, 0xef, 0xf2, 0x85, 0xaf, 0x7e, 0x72, 0x0d, 0x7d, 0x46, 0x16, 0x16, 0xad, 0xa8,
	0x79, 0xa3, 0x8c, 0x81, 0x0d, 0xbb, 0x1d, 0xc4, 0x49, 0x60, 0x1f, 0x03, 0xf2, 0xad, 0x5b, 0xec,
	0x79, 0x8f, 0xdd, 0x1e, 0x3a, 0x76, 0x12, 0x04, 0x71, 0x8b, 0xfd, 0xd7, 0x00, 0xe8, 0x5b, 0x92,
	0x18, 0xb5, 0xd5, 0x4a, 0x6f, 0xf9, 0x1e, 0x06, 0x16, 0x87, 0x83, 0x2e, 0x0f, 0x0e, 0xda, 0xba,
	0xaa, 0x56, 0xd2, 0x17, 0xc7, 0xf7, 0x45, 0x26, 0xf9, 0x15, 0x0c, 0x34, 0x27, 0x71, 0x98, 0xdb,
	0x87, 0x33, 0x12, 0xac, 0x0f, 0x0e, 0xac, 0x66, 0x27, 0xce, 0x5f, 0xbd, 0x2e, 0x4f, 0xbc, 0xc6,
	0xb7, 0xa3, 0x2f, 0xc8, 0xbd, 0xce, 0xe9, 0xd0, 0xec, 0x1a, 0x3a, 0xa3, 0x50, 0xb3, 0x24, 0x0c,
	0x74, 0x77, 0xc2, 0x5f, 0x46, 0xea, 0x87, 0x0e, 0xc1, 0xc2, 0xd4, 0x68, 0xd9, 0x62, 0x1c, 0xda,
	0x93, 0xf7, 0x1e, 0xac, 0x3e, 0x10, 0x66, 0x90, 0xd7, 0xca, 0xb7, 0x9e, 0x3e, 0x6a, 0x35, 0xf0,
	0x71, 0x21, 0xfe, 0x6f, 0x7f, 0x56, 0x57, 0x64, 0x69, 0x90, 0x3b, 0xad, 0x7e, 0x38, 0xb8, 0x69,
	0xd8, 0x45, 0xf6, 0xed, 0xf9, 0x56, 0xd9, 0x9d, 0xab, 0x72, 0x89, 0x4d, 0x21, 0xd1, 0xc8, 0x9d,
	0x50, 0xba, 0x90, 0xa8, 0x2d, 0x68, 0x8b, 0xe6, 0x6c, 0x8b, 0xe3, 0x2f, 0x52, 0x1d, 0x87, 0xa0,
	0x97, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc9, 0xc8, 0xde, 0x89, 0x36, 0x03, 0x00, 0x00,
}
