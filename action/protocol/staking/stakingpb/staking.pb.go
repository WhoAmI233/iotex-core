// Code generated by protoc-gen-go. DO NOT EDIT.
// source: staking.proto

package stakingpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Bucket struct {
	Index                uint64               `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	CandidateAddress     string               `protobuf:"bytes,2,opt,name=candidateAddress,proto3" json:"candidateAddress,omitempty"`
	StakedAmount         string               `protobuf:"bytes,3,opt,name=stakedAmount,proto3" json:"stakedAmount,omitempty"`
	StakedDuration       uint32               `protobuf:"varint,4,opt,name=stakedDuration,proto3" json:"stakedDuration,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	StakeStartTime       *timestamp.Timestamp `protobuf:"bytes,6,opt,name=stakeStartTime,proto3" json:"stakeStartTime,omitempty"`
	UnstakeStartTime     *timestamp.Timestamp `protobuf:"bytes,7,opt,name=unstakeStartTime,proto3" json:"unstakeStartTime,omitempty"`
	AutoStake            bool                 `protobuf:"varint,8,opt,name=autoStake,proto3" json:"autoStake,omitempty"`
	Owner                string               `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{0}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Bucket) GetCandidateAddress() string {
	if m != nil {
		return m.CandidateAddress
	}
	return ""
}

func (m *Bucket) GetStakedAmount() string {
	if m != nil {
		return m.StakedAmount
	}
	return ""
}

func (m *Bucket) GetStakedDuration() uint32 {
	if m != nil {
		return m.StakedDuration
	}
	return 0
}

func (m *Bucket) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Bucket) GetStakeStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StakeStartTime
	}
	return nil
}

func (m *Bucket) GetUnstakeStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.UnstakeStartTime
	}
	return nil
}

func (m *Bucket) GetAutoStake() bool {
	if m != nil {
		return m.AutoStake
	}
	return false
}

func (m *Bucket) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type BucketIndices struct {
	Indices              []uint64 `protobuf:"varint,1,rep,packed,name=indices,proto3" json:"indices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BucketIndices) Reset()         { *m = BucketIndices{} }
func (m *BucketIndices) String() string { return proto.CompactTextString(m) }
func (*BucketIndices) ProtoMessage()    {}
func (*BucketIndices) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{1}
}

func (m *BucketIndices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BucketIndices.Unmarshal(m, b)
}
func (m *BucketIndices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BucketIndices.Marshal(b, m, deterministic)
}
func (m *BucketIndices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketIndices.Merge(m, src)
}
func (m *BucketIndices) XXX_Size() int {
	return xxx_messageInfo_BucketIndices.Size(m)
}
func (m *BucketIndices) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketIndices.DiscardUnknown(m)
}

var xxx_messageInfo_BucketIndices proto.InternalMessageInfo

func (m *BucketIndices) GetIndices() []uint64 {
	if m != nil {
		return m.Indices
	}
	return nil
}

type Candidate struct {
	OwnerAddress         string   `protobuf:"bytes,1,opt,name=ownerAddress,proto3" json:"ownerAddress,omitempty"`
	OperatorAddress      string   `protobuf:"bytes,2,opt,name=operatorAddress,proto3" json:"operatorAddress,omitempty"`
	RewardAddress        string   `protobuf:"bytes,3,opt,name=rewardAddress,proto3" json:"rewardAddress,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Votes                string   `protobuf:"bytes,5,opt,name=votes,proto3" json:"votes,omitempty"`
	SelfStakeBucketIdx   uint64   `protobuf:"varint,6,opt,name=selfStakeBucketIdx,proto3" json:"selfStakeBucketIdx,omitempty"`
	SelfStake            string   `protobuf:"bytes,7,opt,name=selfStake,proto3" json:"selfStake,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Candidate) Reset()         { *m = Candidate{} }
func (m *Candidate) String() string { return proto.CompactTextString(m) }
func (*Candidate) ProtoMessage()    {}
func (*Candidate) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{2}
}

func (m *Candidate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candidate.Unmarshal(m, b)
}
func (m *Candidate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candidate.Marshal(b, m, deterministic)
}
func (m *Candidate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candidate.Merge(m, src)
}
func (m *Candidate) XXX_Size() int {
	return xxx_messageInfo_Candidate.Size(m)
}
func (m *Candidate) XXX_DiscardUnknown() {
	xxx_messageInfo_Candidate.DiscardUnknown(m)
}

var xxx_messageInfo_Candidate proto.InternalMessageInfo

func (m *Candidate) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *Candidate) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *Candidate) GetRewardAddress() string {
	if m != nil {
		return m.RewardAddress
	}
	return ""
}

func (m *Candidate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Candidate) GetVotes() string {
	if m != nil {
		return m.Votes
	}
	return ""
}

func (m *Candidate) GetSelfStakeBucketIdx() uint64 {
	if m != nil {
		return m.SelfStakeBucketIdx
	}
	return 0
}

func (m *Candidate) GetSelfStake() string {
	if m != nil {
		return m.SelfStake
	}
	return ""
}

type Candidates struct {
	Candidates           []*Candidate `protobuf:"bytes,1,rep,name=candidates,proto3" json:"candidates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Candidates) Reset()         { *m = Candidates{} }
func (m *Candidates) String() string { return proto.CompactTextString(m) }
func (*Candidates) ProtoMessage()    {}
func (*Candidates) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{3}
}

func (m *Candidates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candidates.Unmarshal(m, b)
}
func (m *Candidates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candidates.Marshal(b, m, deterministic)
}
func (m *Candidates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candidates.Merge(m, src)
}
func (m *Candidates) XXX_Size() int {
	return xxx_messageInfo_Candidates.Size(m)
}
func (m *Candidates) XXX_DiscardUnknown() {
	xxx_messageInfo_Candidates.DiscardUnknown(m)
}

var xxx_messageInfo_Candidates proto.InternalMessageInfo

func (m *Candidates) GetCandidates() []*Candidate {
	if m != nil {
		return m.Candidates
	}
	return nil
}

func init() {
	proto.RegisterType((*Bucket)(nil), "stakingpb.Bucket")
	proto.RegisterType((*BucketIndices)(nil), "stakingpb.BucketIndices")
	proto.RegisterType((*Candidate)(nil), "stakingpb.Candidate")
	proto.RegisterType((*Candidates)(nil), "stakingpb.Candidates")
}

func init() { proto.RegisterFile("staking.proto", fileDescriptor_289e7c8aea278311) }

var fileDescriptor_289e7c8aea278311 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6a, 0x1b, 0x31,
	0x10, 0x86, 0x51, 0xbd, 0x71, 0xa2, 0x49, 0xdd, 0x86, 0x21, 0x07, 0x11, 0x0a, 0x5d, 0x96, 0x52,
	0xb6, 0x3d, 0x6c, 0x20, 0xed, 0xa9, 0xb7, 0xb8, 0xa5, 0xd0, 0xab, 0x92, 0x17, 0x90, 0xad, 0x89,
	0x11, 0xc9, 0x4a, 0x46, 0xd2, 0x36, 0x79, 0x8e, 0xbe, 0x6c, 0xaf, 0x65, 0x25, 0xef, 0xc6, 0x76,
	0x0a, 0xbe, 0x69, 0x7e, 0xfd, 0x33, 0xe2, 0x9f, 0x6f, 0x17, 0x66, 0x21, 0xaa, 0x7b, 0x63, 0x57,
	0xcd, 0xda, 0xbb, 0xe8, 0x90, 0x6f, 0xca, 0xf5, 0xe2, 0xe2, 0xfd, 0xca, 0xb9, 0xd5, 0x03, 0x5d,
	0xa6, 0x8b, 0x45, 0x77, 0x77, 0x19, 0x4d, 0x4b, 0x21, 0xaa, 0x76, 0x9d, 0xbd, 0xd5, 0x9f, 0x09,
	0x4c, 0xe7, 0xdd, 0xf2, 0x9e, 0x22, 0x9e, 0xc3, 0x91, 0xb1, 0x9a, 0x9e, 0x04, 0x2b, 0x59, 0x5d,
	0xc8, 0x5c, 0xe0, 0x67, 0x38, 0x5b, 0x2a, 0xab, 0x8d, 0x56, 0x91, 0xae, 0xb5, 0xf6, 0x14, 0x82,
	0x78, 0x55, 0xb2, 0x9a, 0xcb, 0x17, 0x3a, 0x56, 0xf0, 0xba, 0x7f, 0x9a, 0xf4, 0x75, 0xeb, 0x3a,
	0x1b, 0xc5, 0x24, 0xf9, 0x76, 0x34, 0xfc, 0x08, 0x6f, 0x72, 0xfd, 0xa3, 0xf3, 0x2a, 0x1a, 0x67,
	0x45, 0x51, 0xb2, 0x7a, 0x26, 0xf7, 0x54, 0xfc, 0x06, 0xb0, 0xf4, 0xa4, 0x22, 0xdd, 0x9a, 0x96,
	0xc4, 0x51, 0xc9, 0xea, 0xd3, 0xab, 0x8b, 0x26, 0xc7, 0x69, 0x86, 0x38, 0xcd, 0xed, 0x10, 0x47,
	0x6e, 0xb9, 0x71, 0xbe, 0x79, 0xe3, 0x26, 0x2a, 0x1f, 0x53, 0xff, 0xf4, 0x60, 0xff, 0x5e, 0x07,
	0xfe, 0x84, 0xb3, 0xce, 0xee, 0x4d, 0x39, 0x3e, 0x38, 0xe5, 0x45, 0x0f, 0xbe, 0x03, 0xae, 0xba,
	0xe8, 0x6e, 0x7a, 0x55, 0x9c, 0x94, 0xac, 0x3e, 0x91, 0xcf, 0x42, 0xbf, 0x73, 0xf7, 0x68, 0xc9,
	0x0b, 0x9e, 0x56, 0x95, 0x8b, 0xea, 0x13, 0xcc, 0x32, 0x93, 0x5f, 0x56, 0x9b, 0x25, 0x05, 0x14,
	0x70, 0x6c, 0xf2, 0x51, 0xb0, 0x72, 0x52, 0x17, 0x72, 0x28, 0xab, 0xbf, 0x0c, 0xf8, 0xf7, 0x81,
	0x43, 0x0f, 0x20, 0x4d, 0x18, 0x40, 0xb1, 0x0c, 0x60, 0x5b, 0xc3, 0x1a, 0xde, 0xba, 0x35, 0x79,
	0x15, 0x9d, 0xdf, 0xe5, 0xb9, 0x2f, 0xe3, 0x07, 0x98, 0x79, 0x7a, 0x54, 0x5e, 0x0f, 0xbe, 0xcc,
	0x73, 0x57, 0x44, 0x84, 0xc2, 0xaa, 0x96, 0x12, 0x46, 0x2e, 0xd3, 0xb9, 0x8f, 0xf5, 0xdb, 0x45,
	0x0a, 0x89, 0x1b, 0x97, 0xb9, 0xc0, 0x06, 0x30, 0xd0, 0xc3, 0x5d, 0x4a, 0xbe, 0xc9, 0xa7, 0x9f,
	0x12, 0x9a, 0x42, 0xfe, 0xe7, 0xa6, 0x5f, 0xdd, 0xa8, 0xa6, 0xdd, 0x73, 0xf9, 0x2c, 0x54, 0x73,
	0x80, 0x31, 0x78, 0xc0, 0xaf, 0x00, 0xe3, 0xe7, 0x98, 0x97, 0x74, 0x7a, 0x75, 0xde, 0x8c, 0x3f,
	0x42, 0x33, 0x5a, 0xe5, 0x96, 0x6f, 0x31, 0x4d, 0x08, 0xbf, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0xc4, 0x39, 0x2e, 0xe7, 0x41, 0x03, 0x00, 0x00,
}
