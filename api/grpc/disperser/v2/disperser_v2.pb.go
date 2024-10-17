// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: disperser/v2/disperser_v2.proto

package v2

import (
	common "github.com/Layr-Labs/eigenda/api/grpc/common"
	v2 "github.com/Layr-Labs/eigenda/api/grpc/common/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BlobStatus represents the status of a blob.
// The status of a blob is updated as the blob is processed by the disperser.
// The status of a blob can be queried by the client using the GetBlobStatus API.
// Intermediate states are states that the blob can be in while being processed, and it can be updated to a differet state:
// - QUEUED
// - ENCODED
// Terminal states are states that will not be updated to a different state:
// - CERTIFIED
// - FAILED
// - INSUFFICIENT_SIGNATURES
type BlobStatus int32

const (
	BlobStatus_UNKNOWN BlobStatus = 0
	// QUEUED means that the blob has been queued by the disperser for processing
	BlobStatus_QUEUED BlobStatus = 1
	// ENCODED means that the blob has been encoded and is ready to be dispersed to DA Nodes
	BlobStatus_ENCODED BlobStatus = 2
	// CERTIFIED means the blob has been dispersed and attested by the DA nodes
	BlobStatus_CERTIFIED BlobStatus = 3
	// FAILED means that the blob has failed permanently (for reasons other than insufficient
	// signatures, which is a separate state)
	BlobStatus_FAILED BlobStatus = 4
	// INSUFFICIENT_SIGNATURES means that the confirmation threshold for the blob was not met
	// for at least one quorum.
	BlobStatus_INSUFFICIENT_SIGNATURES BlobStatus = 5
)

// Enum value maps for BlobStatus.
var (
	BlobStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "QUEUED",
		2: "ENCODED",
		3: "CERTIFIED",
		4: "FAILED",
		5: "INSUFFICIENT_SIGNATURES",
	}
	BlobStatus_value = map[string]int32{
		"UNKNOWN":                 0,
		"QUEUED":                  1,
		"ENCODED":                 2,
		"CERTIFIED":               3,
		"FAILED":                  4,
		"INSUFFICIENT_SIGNATURES": 5,
	}
)

func (x BlobStatus) Enum() *BlobStatus {
	p := new(BlobStatus)
	*p = x
	return p
}

func (x BlobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_disperser_v2_disperser_v2_proto_enumTypes[0].Descriptor()
}

func (BlobStatus) Type() protoreflect.EnumType {
	return &file_disperser_v2_disperser_v2_proto_enumTypes[0]
}

func (x BlobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlobStatus.Descriptor instead.
func (BlobStatus) EnumDescriptor() ([]byte, []int) {
	return file_disperser_v2_disperser_v2_proto_rawDescGZIP(), []int{0}
}

type DisperseBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The data to be dispersed.
	// The size of data must be <= 16MiB. Every 32 bytes of data is interpreted as an integer in big endian format
	// where the lower address has more significant bits. The integer must stay in the valid range to be interpreted
	// as a field element on the bn254 curve. The valid range is
	// 0 <= x < 21888242871839275222246405745257275088548364400416034343698204186575808495617
	// If any one of the 32 bytes elements is outside the range, the whole request is deemed as invalid, and rejected.
	Data       []byte         `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	BlobHeader *v2.BlobHeader `protobuf:"bytes,2,opt,name=blob_header,json=blobHeader,proto3" json:"blob_header,omitempty"`
}

func (x *DisperseBlobRequest) Reset() {
	*x = DisperseBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disperser_v2_disperser_v2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisperseBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisperseBlobRequest) ProtoMessage() {}

func (x *DisperseBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_disperser_v2_disperser_v2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisperseBlobRequest.ProtoReflect.Descriptor instead.
func (*DisperseBlobRequest) Descriptor() ([]byte, []int) {
	return file_disperser_v2_disperser_v2_proto_rawDescGZIP(), []int{0}
}

func (x *DisperseBlobRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DisperseBlobRequest) GetBlobHeader() *v2.BlobHeader {
	if x != nil {
		return x.BlobHeader
	}
	return nil
}

type DisperseBlobReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the blob associated with the blob key.
	Result  BlobStatus `protobuf:"varint,1,opt,name=result,proto3,enum=disperser.v2.BlobStatus" json:"result,omitempty"`
	BlobKey []byte     `protobuf:"bytes,2,opt,name=blob_key,json=blobKey,proto3" json:"blob_key,omitempty"`
}

func (x *DisperseBlobReply) Reset() {
	*x = DisperseBlobReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disperser_v2_disperser_v2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisperseBlobReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisperseBlobReply) ProtoMessage() {}

func (x *DisperseBlobReply) ProtoReflect() protoreflect.Message {
	mi := &file_disperser_v2_disperser_v2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisperseBlobReply.ProtoReflect.Descriptor instead.
func (*DisperseBlobReply) Descriptor() ([]byte, []int) {
	return file_disperser_v2_disperser_v2_proto_rawDescGZIP(), []int{1}
}

func (x *DisperseBlobReply) GetResult() BlobStatus {
	if x != nil {
		return x.Result
	}
	return BlobStatus_UNKNOWN
}

func (x *DisperseBlobReply) GetBlobKey() []byte {
	if x != nil {
		return x.BlobKey
	}
	return nil
}

// BlobStatusRequest is used to query the status of a blob.
type BlobStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlobKey []byte `protobuf:"bytes,1,opt,name=blob_key,json=blobKey,proto3" json:"blob_key,omitempty"`
}

func (x *BlobStatusRequest) Reset() {
	*x = BlobStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disperser_v2_disperser_v2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobStatusRequest) ProtoMessage() {}

func (x *BlobStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_disperser_v2_disperser_v2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobStatusRequest.ProtoReflect.Descriptor instead.
func (*BlobStatusRequest) Descriptor() ([]byte, []int) {
	return file_disperser_v2_disperser_v2_proto_rawDescGZIP(), []int{2}
}

func (x *BlobStatusRequest) GetBlobKey() []byte {
	if x != nil {
		return x.BlobKey
	}
	return nil
}

type BlobStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the blob.
	Status BlobStatus `protobuf:"varint,1,opt,name=status,proto3,enum=disperser.v2.BlobStatus" json:"status,omitempty"`
	// The signed blob certificate
	SignedCertificate *SignedCertificate `protobuf:"bytes,2,opt,name=signed_certificate,json=signedCertificate,proto3" json:"signed_certificate,omitempty"`
}

func (x *BlobStatusReply) Reset() {
	*x = BlobStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disperser_v2_disperser_v2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobStatusReply) ProtoMessage() {}

func (x *BlobStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_disperser_v2_disperser_v2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobStatusReply.ProtoReflect.Descriptor instead.
func (*BlobStatusReply) Descriptor() ([]byte, []int) {
	return file_disperser_v2_disperser_v2_proto_rawDescGZIP(), []int{3}
}

func (x *BlobStatusReply) GetStatus() BlobStatus {
	if x != nil {
		return x.Status
	}
	return BlobStatus_UNKNOWN
}

func (x *BlobStatusReply) GetSignedCertificate() *SignedCertificate {
	if x != nil {
		return x.SignedCertificate
	}
	return nil
}

// Utility method used to generate the commitment of blob given its data.
// This can be used to construct BlobHeader.commitment
type BlobCommitmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BlobCommitmentRequest) Reset() {
	*x = BlobCommitmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disperser_v2_disperser_v2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobCommitmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobCommitmentRequest) ProtoMessage() {}

func (x *BlobCommitmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_disperser_v2_disperser_v2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobCommitmentRequest.ProtoReflect.Descriptor instead.
func (*BlobCommitmentRequest) Descriptor() ([]byte, []int) {
	return file_disperser_v2_disperser_v2_proto_rawDescGZIP(), []int{4}
}

func (x *BlobCommitmentRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BlobCommitmentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlobCommitment *common.BlobCommitment `protobuf:"bytes,1,opt,name=blob_commitment,json=blobCommitment,proto3" json:"blob_commitment,omitempty"`
}

func (x *BlobCommitmentReply) Reset() {
	*x = BlobCommitmentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disperser_v2_disperser_v2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlobCommitmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlobCommitmentReply) ProtoMessage() {}

func (x *BlobCommitmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_disperser_v2_disperser_v2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlobCommitmentReply.ProtoReflect.Descriptor instead.
func (*BlobCommitmentReply) Descriptor() ([]byte, []int) {
	return file_disperser_v2_disperser_v2_proto_rawDescGZIP(), []int{5}
}

func (x *BlobCommitmentReply) GetBlobCommitment() *common.BlobCommitment {
	if x != nil {
		return x.BlobCommitment
	}
	return nil
}

type SignedCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlobCertificate             *v2.BlobCertificate `protobuf:"bytes,1,opt,name=blob_certificate,json=blobCertificate,proto3" json:"blob_certificate,omitempty"`
	NonSignerStakesAndSignature *Attestation        `protobuf:"bytes,2,opt,name=non_signer_stakes_and_signature,json=nonSignerStakesAndSignature,proto3" json:"non_signer_stakes_and_signature,omitempty"`
}

func (x *SignedCertificate) Reset() {
	*x = SignedCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disperser_v2_disperser_v2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedCertificate) ProtoMessage() {}

func (x *SignedCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_disperser_v2_disperser_v2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedCertificate.ProtoReflect.Descriptor instead.
func (*SignedCertificate) Descriptor() ([]byte, []int) {
	return file_disperser_v2_disperser_v2_proto_rawDescGZIP(), []int{6}
}

func (x *SignedCertificate) GetBlobCertificate() *v2.BlobCertificate {
	if x != nil {
		return x.BlobCertificate
	}
	return nil
}

func (x *SignedCertificate) GetNonSignerStakesAndSignature() *Attestation {
	if x != nil {
		return x.NonSignerStakesAndSignature
	}
	return nil
}

type Attestation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NonSignerQuorumBitmapIndices []uint32                          `protobuf:"varint,1,rep,packed,name=nonSignerQuorumBitmapIndices,proto3" json:"nonSignerQuorumBitmapIndices,omitempty"`
	NonSignerPubkeys             []*common.G1Commitment            `protobuf:"bytes,2,rep,name=nonSignerPubkeys,proto3" json:"nonSignerPubkeys,omitempty"`
	QuorumApks                   []*common.G1Commitment            `protobuf:"bytes,3,rep,name=quorumApks,proto3" json:"quorumApks,omitempty"`
	ApkG2                        *common.G2Commitment              `protobuf:"bytes,4,opt,name=apkG2,proto3" json:"apkG2,omitempty"`
	Sigma                        *common.G1Commitment              `protobuf:"bytes,5,opt,name=sigma,proto3" json:"sigma,omitempty"`
	QuorumApkIndices             []uint32                          `protobuf:"varint,6,rep,packed,name=quorumApkIndices,proto3" json:"quorumApkIndices,omitempty"`
	TotalStakeIndices            []uint32                          `protobuf:"varint,7,rep,packed,name=totalStakeIndices,proto3" json:"totalStakeIndices,omitempty"`
	NonSignerStakeIndices        []*NonSignerStakeIndicesForQuorum `protobuf:"bytes,8,rep,name=nonSignerStakeIndices,proto3" json:"nonSignerStakeIndices,omitempty"`
}

func (x *Attestation) Reset() {
	*x = Attestation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disperser_v2_disperser_v2_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attestation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attestation) ProtoMessage() {}

func (x *Attestation) ProtoReflect() protoreflect.Message {
	mi := &file_disperser_v2_disperser_v2_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attestation.ProtoReflect.Descriptor instead.
func (*Attestation) Descriptor() ([]byte, []int) {
	return file_disperser_v2_disperser_v2_proto_rawDescGZIP(), []int{7}
}

func (x *Attestation) GetNonSignerQuorumBitmapIndices() []uint32 {
	if x != nil {
		return x.NonSignerQuorumBitmapIndices
	}
	return nil
}

func (x *Attestation) GetNonSignerPubkeys() []*common.G1Commitment {
	if x != nil {
		return x.NonSignerPubkeys
	}
	return nil
}

func (x *Attestation) GetQuorumApks() []*common.G1Commitment {
	if x != nil {
		return x.QuorumApks
	}
	return nil
}

func (x *Attestation) GetApkG2() *common.G2Commitment {
	if x != nil {
		return x.ApkG2
	}
	return nil
}

func (x *Attestation) GetSigma() *common.G1Commitment {
	if x != nil {
		return x.Sigma
	}
	return nil
}

func (x *Attestation) GetQuorumApkIndices() []uint32 {
	if x != nil {
		return x.QuorumApkIndices
	}
	return nil
}

func (x *Attestation) GetTotalStakeIndices() []uint32 {
	if x != nil {
		return x.TotalStakeIndices
	}
	return nil
}

func (x *Attestation) GetNonSignerStakeIndices() []*NonSignerStakeIndicesForQuorum {
	if x != nil {
		return x.NonSignerStakeIndices
	}
	return nil
}

type NonSignerStakeIndicesForQuorum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indices []uint32 `protobuf:"varint,1,rep,packed,name=indices,proto3" json:"indices,omitempty"`
}

func (x *NonSignerStakeIndicesForQuorum) Reset() {
	*x = NonSignerStakeIndicesForQuorum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disperser_v2_disperser_v2_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonSignerStakeIndicesForQuorum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonSignerStakeIndicesForQuorum) ProtoMessage() {}

func (x *NonSignerStakeIndicesForQuorum) ProtoReflect() protoreflect.Message {
	mi := &file_disperser_v2_disperser_v2_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonSignerStakeIndicesForQuorum.ProtoReflect.Descriptor instead.
func (*NonSignerStakeIndicesForQuorum) Descriptor() ([]byte, []int) {
	return file_disperser_v2_disperser_v2_proto_rawDescGZIP(), []int{8}
}

func (x *NonSignerStakeIndicesForQuorum) GetIndices() []uint32 {
	if x != nil {
		return x.Indices
	}
	return nil
}

var File_disperser_v2_disperser_v2_proto protoreflect.FileDescriptor

var file_disperser_v2_disperser_v2_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x64,
	0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x5f, 0x76, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x1a,
	0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x13,
	0x44, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x62, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x62, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22,
	0x60, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x62, 0x4b, 0x65,
	0x79, 0x22, 0x2e, 0x0a, 0x11, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x62, 0x4b, 0x65,
	0x79, 0x22, 0x93, 0x01, 0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4e, 0x0a, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x11, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x15, 0x42, 0x6c, 0x6f, 0x62, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x13, 0x42, 0x6c, 0x6f, 0x62, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3f, 0x0a, 0x0f, 0x62,
	0x6c, 0x6f, 0x62, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x6c,
	0x6f, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0e, 0x62, 0x6c,
	0x6f, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xbb, 0x01, 0x0a,
	0x11, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x62, 0x6c, 0x6f, 0x62, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x5f, 0x0a, 0x1f, 0x6e, 0x6f, 0x6e,
	0x5f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x73, 0x5f, 0x61,
	0x6e, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1b, 0x6e,
	0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x73, 0x41, 0x6e,
	0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xdf, 0x03, 0x0a, 0x0b, 0x41,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x1c, 0x6e, 0x6f,
	0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x42, 0x69, 0x74,
	0x6d, 0x61, 0x70, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x1c, 0x6e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x51, 0x75, 0x6f, 0x72, 0x75,
	0x6d, 0x42, 0x69, 0x74, 0x6d, 0x61, 0x70, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x40,
	0x0a, 0x10, 0x6e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65,
	0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x47, 0x31, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10,
	0x6e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x73,
	0x12, 0x34, 0x0a, 0x0a, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x41, 0x70, 0x6b, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x31,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x71, 0x75, 0x6f, 0x72,
	0x75, 0x6d, 0x41, 0x70, 0x6b, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x61, 0x70, 0x6b, 0x47, 0x32, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47,
	0x32, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x70, 0x6b,
	0x47, 0x32, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x31, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x12, 0x2a,
	0x0a, 0x10, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x41, 0x70, 0x6b, 0x49, 0x6e, 0x64, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x71, 0x75, 0x6f, 0x72, 0x75, 0x6d,
	0x41, 0x70, 0x6b, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x6b,
	0x65, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x15, 0x6e, 0x6f, 0x6e, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x6b, 0x65, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x51,
	0x75, 0x6f, 0x72, 0x75, 0x6d, 0x52, 0x15, 0x6e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x6b, 0x65, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x1e,
	0x4e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x49, 0x6e,
	0x64, 0x69, 0x63, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x07, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x2a, 0x6a, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x62,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x53, 0x55, 0x46,
	0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x41, 0x54, 0x55, 0x52,
	0x45, 0x53, 0x10, 0x05, 0x32, 0x93, 0x02, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73,
	0x65, 0x72, 0x12, 0x54, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x42, 0x6c,
	0x6f, 0x62, 0x12, 0x21, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x42, 0x6c, 0x6f,
	0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x2e, 0x64, 0x69, 0x73, 0x70,
	0x65, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x69, 0x73,
	0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x23, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e,
	0x42, 0x6c, 0x6f, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x79, 0x72, 0x2d, 0x4c, 0x61,
	0x62, 0x73, 0x2f, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x65, 0x72, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_disperser_v2_disperser_v2_proto_rawDescOnce sync.Once
	file_disperser_v2_disperser_v2_proto_rawDescData = file_disperser_v2_disperser_v2_proto_rawDesc
)

func file_disperser_v2_disperser_v2_proto_rawDescGZIP() []byte {
	file_disperser_v2_disperser_v2_proto_rawDescOnce.Do(func() {
		file_disperser_v2_disperser_v2_proto_rawDescData = protoimpl.X.CompressGZIP(file_disperser_v2_disperser_v2_proto_rawDescData)
	})
	return file_disperser_v2_disperser_v2_proto_rawDescData
}

var file_disperser_v2_disperser_v2_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_disperser_v2_disperser_v2_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_disperser_v2_disperser_v2_proto_goTypes = []interface{}{
	(BlobStatus)(0),                        // 0: disperser.v2.BlobStatus
	(*DisperseBlobRequest)(nil),            // 1: disperser.v2.DisperseBlobRequest
	(*DisperseBlobReply)(nil),              // 2: disperser.v2.DisperseBlobReply
	(*BlobStatusRequest)(nil),              // 3: disperser.v2.BlobStatusRequest
	(*BlobStatusReply)(nil),                // 4: disperser.v2.BlobStatusReply
	(*BlobCommitmentRequest)(nil),          // 5: disperser.v2.BlobCommitmentRequest
	(*BlobCommitmentReply)(nil),            // 6: disperser.v2.BlobCommitmentReply
	(*SignedCertificate)(nil),              // 7: disperser.v2.SignedCertificate
	(*Attestation)(nil),                    // 8: disperser.v2.Attestation
	(*NonSignerStakeIndicesForQuorum)(nil), // 9: disperser.v2.NonSignerStakeIndicesForQuorum
	(*v2.BlobHeader)(nil),                  // 10: common.v2.BlobHeader
	(*common.BlobCommitment)(nil),          // 11: common.BlobCommitment
	(*v2.BlobCertificate)(nil),             // 12: common.v2.BlobCertificate
	(*common.G1Commitment)(nil),            // 13: common.G1Commitment
	(*common.G2Commitment)(nil),            // 14: common.G2Commitment
}
var file_disperser_v2_disperser_v2_proto_depIdxs = []int32{
	10, // 0: disperser.v2.DisperseBlobRequest.blob_header:type_name -> common.v2.BlobHeader
	0,  // 1: disperser.v2.DisperseBlobReply.result:type_name -> disperser.v2.BlobStatus
	0,  // 2: disperser.v2.BlobStatusReply.status:type_name -> disperser.v2.BlobStatus
	7,  // 3: disperser.v2.BlobStatusReply.signed_certificate:type_name -> disperser.v2.SignedCertificate
	11, // 4: disperser.v2.BlobCommitmentReply.blob_commitment:type_name -> common.BlobCommitment
	12, // 5: disperser.v2.SignedCertificate.blob_certificate:type_name -> common.v2.BlobCertificate
	8,  // 6: disperser.v2.SignedCertificate.non_signer_stakes_and_signature:type_name -> disperser.v2.Attestation
	13, // 7: disperser.v2.Attestation.nonSignerPubkeys:type_name -> common.G1Commitment
	13, // 8: disperser.v2.Attestation.quorumApks:type_name -> common.G1Commitment
	14, // 9: disperser.v2.Attestation.apkG2:type_name -> common.G2Commitment
	13, // 10: disperser.v2.Attestation.sigma:type_name -> common.G1Commitment
	9,  // 11: disperser.v2.Attestation.nonSignerStakeIndices:type_name -> disperser.v2.NonSignerStakeIndicesForQuorum
	1,  // 12: disperser.v2.Disperser.DisperseBlob:input_type -> disperser.v2.DisperseBlobRequest
	3,  // 13: disperser.v2.Disperser.GetBlobStatus:input_type -> disperser.v2.BlobStatusRequest
	5,  // 14: disperser.v2.Disperser.GetBlobCommitment:input_type -> disperser.v2.BlobCommitmentRequest
	2,  // 15: disperser.v2.Disperser.DisperseBlob:output_type -> disperser.v2.DisperseBlobReply
	4,  // 16: disperser.v2.Disperser.GetBlobStatus:output_type -> disperser.v2.BlobStatusReply
	6,  // 17: disperser.v2.Disperser.GetBlobCommitment:output_type -> disperser.v2.BlobCommitmentReply
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_disperser_v2_disperser_v2_proto_init() }
func file_disperser_v2_disperser_v2_proto_init() {
	if File_disperser_v2_disperser_v2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_disperser_v2_disperser_v2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisperseBlobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_disperser_v2_disperser_v2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisperseBlobReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_disperser_v2_disperser_v2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobStatusRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_disperser_v2_disperser_v2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobStatusReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_disperser_v2_disperser_v2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobCommitmentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_disperser_v2_disperser_v2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlobCommitmentReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_disperser_v2_disperser_v2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedCertificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_disperser_v2_disperser_v2_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attestation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_disperser_v2_disperser_v2_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonSignerStakeIndicesForQuorum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_disperser_v2_disperser_v2_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_disperser_v2_disperser_v2_proto_goTypes,
		DependencyIndexes: file_disperser_v2_disperser_v2_proto_depIdxs,
		EnumInfos:         file_disperser_v2_disperser_v2_proto_enumTypes,
		MessageInfos:      file_disperser_v2_disperser_v2_proto_msgTypes,
	}.Build()
	File_disperser_v2_disperser_v2_proto = out.File
	file_disperser_v2_disperser_v2_proto_rawDesc = nil
	file_disperser_v2_disperser_v2_proto_goTypes = nil
	file_disperser_v2_disperser_v2_proto_depIdxs = nil
}
