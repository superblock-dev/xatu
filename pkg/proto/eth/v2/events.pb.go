// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: pkg/proto/eth/v2/events.proto

package v2

import (
	v1 "github.com/ethpandaops/xatu/pkg/proto/eth/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockVersion int32

const (
	BlockVersion_UNKNOWN   BlockVersion = 0
	BlockVersion_PHASE0    BlockVersion = 1
	BlockVersion_ALTAIR    BlockVersion = 2
	BlockVersion_BELLATRIX BlockVersion = 3
	BlockVersion_CAPELLA   BlockVersion = 4
	BlockVersion_DENEB     BlockVersion = 5
	BlockVersion_ELECTRA   BlockVersion = 6
)

// Enum value maps for BlockVersion.
var (
	BlockVersion_name = map[int32]string{
		0: "UNKNOWN",
		1: "PHASE0",
		2: "ALTAIR",
		3: "BELLATRIX",
		4: "CAPELLA",
		5: "DENEB",
		6: "ELECTRA",
	}
	BlockVersion_value = map[string]int32{
		"UNKNOWN":   0,
		"PHASE0":    1,
		"ALTAIR":    2,
		"BELLATRIX": 3,
		"CAPELLA":   4,
		"DENEB":     5,
		"ELECTRA":   6,
	}
)

func (x BlockVersion) Enum() *BlockVersion {
	p := new(BlockVersion)
	*p = x
	return p
}

func (x BlockVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlockVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_proto_eth_v2_events_proto_enumTypes[0].Descriptor()
}

func (BlockVersion) Type() protoreflect.EnumType {
	return &file_pkg_proto_eth_v2_events_proto_enumTypes[0]
}

func (x BlockVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlockVersion.Descriptor instead.
func (BlockVersion) EnumDescriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v2_events_proto_rawDescGZIP(), []int{0}
}

type EventBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*EventBlock_Phase0Block
	//	*EventBlock_AltairBlock
	//	*EventBlock_BellatrixBlock
	//	*EventBlock_CapellaBlock
	//	*EventBlock_DenebBlock
	//	*EventBlock_ElectraBlock
	Message   isEventBlock_Message `protobuf_oneof:"message"`
	Signature string               `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Version   BlockVersion         `protobuf:"varint,6,opt,name=version,proto3,enum=xatu.eth.v2.BlockVersion" json:"version,omitempty"`
}

func (x *EventBlock) Reset() {
	*x = EventBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v2_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventBlock) ProtoMessage() {}

func (x *EventBlock) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v2_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventBlock.ProtoReflect.Descriptor instead.
func (*EventBlock) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v2_events_proto_rawDescGZIP(), []int{0}
}

func (m *EventBlock) GetMessage() isEventBlock_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *EventBlock) GetPhase0Block() *v1.BeaconBlock {
	if x, ok := x.GetMessage().(*EventBlock_Phase0Block); ok {
		return x.Phase0Block
	}
	return nil
}

func (x *EventBlock) GetAltairBlock() *BeaconBlockAltair {
	if x, ok := x.GetMessage().(*EventBlock_AltairBlock); ok {
		return x.AltairBlock
	}
	return nil
}

func (x *EventBlock) GetBellatrixBlock() *BeaconBlockBellatrix {
	if x, ok := x.GetMessage().(*EventBlock_BellatrixBlock); ok {
		return x.BellatrixBlock
	}
	return nil
}

func (x *EventBlock) GetCapellaBlock() *BeaconBlockCapella {
	if x, ok := x.GetMessage().(*EventBlock_CapellaBlock); ok {
		return x.CapellaBlock
	}
	return nil
}

func (x *EventBlock) GetDenebBlock() *BeaconBlockDeneb {
	if x, ok := x.GetMessage().(*EventBlock_DenebBlock); ok {
		return x.DenebBlock
	}
	return nil
}

func (x *EventBlock) GetElectraBlock() *BeaconBlockElectra {
	if x, ok := x.GetMessage().(*EventBlock_ElectraBlock); ok {
		return x.ElectraBlock
	}
	return nil
}

func (x *EventBlock) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *EventBlock) GetVersion() BlockVersion {
	if x != nil {
		return x.Version
	}
	return BlockVersion_UNKNOWN
}

type isEventBlock_Message interface {
	isEventBlock_Message()
}

type EventBlock_Phase0Block struct {
	Phase0Block *v1.BeaconBlock `protobuf:"bytes,1,opt,name=phase0_block,json=PHASE0,proto3,oneof"`
}

type EventBlock_AltairBlock struct {
	AltairBlock *BeaconBlockAltair `protobuf:"bytes,2,opt,name=altair_block,json=ALTAIR,proto3,oneof"`
}

type EventBlock_BellatrixBlock struct {
	BellatrixBlock *BeaconBlockBellatrix `protobuf:"bytes,3,opt,name=bellatrix_block,json=BELLATRIX,proto3,oneof"`
}

type EventBlock_CapellaBlock struct {
	CapellaBlock *BeaconBlockCapella `protobuf:"bytes,4,opt,name=capella_block,json=CAPELLA,proto3,oneof"`
}

type EventBlock_DenebBlock struct {
	DenebBlock *BeaconBlockDeneb `protobuf:"bytes,7,opt,name=deneb_block,json=DENEB,proto3,oneof"`
}

type EventBlock_ElectraBlock struct {
	ElectraBlock *BeaconBlockElectra `protobuf:"bytes,8,opt,name=electra_block,json=ELECTRA,proto3,oneof"`
}

func (*EventBlock_Phase0Block) isEventBlock_Message() {}

func (*EventBlock_AltairBlock) isEventBlock_Message() {}

func (*EventBlock_BellatrixBlock) isEventBlock_Message() {}

func (*EventBlock_CapellaBlock) isEventBlock_Message() {}

func (*EventBlock_DenebBlock) isEventBlock_Message() {}

func (*EventBlock_ElectraBlock) isEventBlock_Message() {}

type EventBlockV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*EventBlockV2_Phase0Block
	//	*EventBlockV2_AltairBlock
	//	*EventBlockV2_BellatrixBlock
	//	*EventBlockV2_CapellaBlock
	//	*EventBlockV2_DenebBlock
	//	*EventBlockV2_ElectraBlock
	Message   isEventBlockV2_Message `protobuf_oneof:"message"`
	Signature string                 `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Version   BlockVersion           `protobuf:"varint,6,opt,name=version,proto3,enum=xatu.eth.v2.BlockVersion" json:"version,omitempty"`
}

func (x *EventBlockV2) Reset() {
	*x = EventBlockV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_eth_v2_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventBlockV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventBlockV2) ProtoMessage() {}

func (x *EventBlockV2) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_eth_v2_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventBlockV2.ProtoReflect.Descriptor instead.
func (*EventBlockV2) Descriptor() ([]byte, []int) {
	return file_pkg_proto_eth_v2_events_proto_rawDescGZIP(), []int{1}
}

func (m *EventBlockV2) GetMessage() isEventBlockV2_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *EventBlockV2) GetPhase0Block() *v1.BeaconBlockV2 {
	if x, ok := x.GetMessage().(*EventBlockV2_Phase0Block); ok {
		return x.Phase0Block
	}
	return nil
}

func (x *EventBlockV2) GetAltairBlock() *BeaconBlockAltairV2 {
	if x, ok := x.GetMessage().(*EventBlockV2_AltairBlock); ok {
		return x.AltairBlock
	}
	return nil
}

func (x *EventBlockV2) GetBellatrixBlock() *BeaconBlockBellatrixV2 {
	if x, ok := x.GetMessage().(*EventBlockV2_BellatrixBlock); ok {
		return x.BellatrixBlock
	}
	return nil
}

func (x *EventBlockV2) GetCapellaBlock() *BeaconBlockCapellaV2 {
	if x, ok := x.GetMessage().(*EventBlockV2_CapellaBlock); ok {
		return x.CapellaBlock
	}
	return nil
}

func (x *EventBlockV2) GetDenebBlock() *BeaconBlockDeneb {
	if x, ok := x.GetMessage().(*EventBlockV2_DenebBlock); ok {
		return x.DenebBlock
	}
	return nil
}

func (x *EventBlockV2) GetElectraBlock() *BeaconBlockElectra {
	if x, ok := x.GetMessage().(*EventBlockV2_ElectraBlock); ok {
		return x.ElectraBlock
	}
	return nil
}

func (x *EventBlockV2) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *EventBlockV2) GetVersion() BlockVersion {
	if x != nil {
		return x.Version
	}
	return BlockVersion_UNKNOWN
}

type isEventBlockV2_Message interface {
	isEventBlockV2_Message()
}

type EventBlockV2_Phase0Block struct {
	Phase0Block *v1.BeaconBlockV2 `protobuf:"bytes,1,opt,name=phase0_block,json=PHASE0,proto3,oneof"`
}

type EventBlockV2_AltairBlock struct {
	AltairBlock *BeaconBlockAltairV2 `protobuf:"bytes,2,opt,name=altair_block,json=ALTAIR,proto3,oneof"`
}

type EventBlockV2_BellatrixBlock struct {
	BellatrixBlock *BeaconBlockBellatrixV2 `protobuf:"bytes,3,opt,name=bellatrix_block,json=BELLATRIX,proto3,oneof"`
}

type EventBlockV2_CapellaBlock struct {
	CapellaBlock *BeaconBlockCapellaV2 `protobuf:"bytes,4,opt,name=capella_block,json=CAPELLA,proto3,oneof"`
}

type EventBlockV2_DenebBlock struct {
	DenebBlock *BeaconBlockDeneb `protobuf:"bytes,7,opt,name=deneb_block,json=DENEB,proto3,oneof"`
}

type EventBlockV2_ElectraBlock struct {
	ElectraBlock *BeaconBlockElectra `protobuf:"bytes,8,opt,name=electra_block,json=ELECTRA,proto3,oneof"`
}

func (*EventBlockV2_Phase0Block) isEventBlockV2_Message() {}

func (*EventBlockV2_AltairBlock) isEventBlockV2_Message() {}

func (*EventBlockV2_BellatrixBlock) isEventBlockV2_Message() {}

func (*EventBlockV2_CapellaBlock) isEventBlockV2_Message() {}

func (*EventBlockV2_DenebBlock) isEventBlockV2_Message() {}

func (*EventBlockV2_ElectraBlock) isEventBlockV2_Message() {}

var File_pkg_proto_eth_v2_events_proto protoreflect.FileDescriptor

var file_pkg_proto_eth_v2_events_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f,
	0x76, 0x32, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65,
	0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x03, 0x0a, 0x0a, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x38, 0x0a, 0x0c, 0x70, 0x68, 0x61, 0x73, 0x65,
	0x30, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x61, 0x63,
	0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x06, 0x50, 0x48, 0x41, 0x53, 0x45,
	0x30, 0x12, 0x3e, 0x0a, 0x0c, 0x61, 0x6c, 0x74, 0x61, 0x69, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x41, 0x6c, 0x74, 0x61, 0x69, 0x72, 0x48, 0x00, 0x52, 0x06, 0x41, 0x4c, 0x54, 0x41, 0x49,
	0x52, 0x12, 0x47, 0x0a, 0x0f, 0x62, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x78, 0x61, 0x74,
	0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x72, 0x69, 0x78, 0x48, 0x00, 0x52,
	0x09, 0x42, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x52, 0x49, 0x58, 0x12, 0x41, 0x0a, 0x0d, 0x63, 0x61,
	0x70, 0x65, 0x6c, 0x6c, 0x61, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e,
	0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x70, 0x65, 0x6c,
	0x6c, 0x61, 0x48, 0x00, 0x52, 0x07, 0x43, 0x41, 0x50, 0x45, 0x4c, 0x4c, 0x41, 0x12, 0x3b, 0x0a,
	0x0b, 0x64, 0x65, 0x6e, 0x65, 0x62, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32,
	0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x6e, 0x65,
	0x62, 0x48, 0x00, 0x52, 0x05, 0x44, 0x45, 0x4e, 0x45, 0x42, 0x12, 0x41, 0x0a, 0x0d, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x72, 0x61, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e,
	0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6c, 0x65, 0x63, 0x74,
	0x72, 0x61, 0x48, 0x00, 0x52, 0x07, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x52, 0x41, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x78,
	0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xfa, 0x03, 0x0a, 0x0c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x32, 0x12, 0x3a, 0x0a, 0x0c,
	0x70, 0x68, 0x61, 0x73, 0x65, 0x30, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x32, 0x48, 0x00,
	0x52, 0x06, 0x50, 0x48, 0x41, 0x53, 0x45, 0x30, 0x12, 0x40, 0x0a, 0x0c, 0x61, 0x6c, 0x74, 0x61,
	0x69, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x6c, 0x74, 0x61, 0x69, 0x72, 0x56, 0x32,
	0x48, 0x00, 0x52, 0x06, 0x41, 0x4c, 0x54, 0x41, 0x49, 0x52, 0x12, 0x49, 0x0a, 0x0f, 0x62, 0x65,
	0x6c, 0x6c, 0x61, 0x74, 0x72, 0x69, 0x78, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x32, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x65, 0x6c,
	0x6c, 0x61, 0x74, 0x72, 0x69, 0x78, 0x56, 0x32, 0x48, 0x00, 0x52, 0x09, 0x42, 0x45, 0x4c, 0x4c,
	0x41, 0x54, 0x52, 0x49, 0x58, 0x12, 0x43, 0x0a, 0x0d, 0x63, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x61,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x78,
	0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x70, 0x65, 0x6c, 0x6c, 0x61, 0x56, 0x32, 0x48,
	0x00, 0x52, 0x07, 0x43, 0x41, 0x50, 0x45, 0x4c, 0x4c, 0x41, 0x12, 0x3b, 0x0a, 0x0b, 0x64, 0x65,
	0x6e, 0x65, 0x62, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x65,
	0x61, 0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x6e, 0x65, 0x62, 0x48, 0x00,
	0x52, 0x05, 0x44, 0x45, 0x4e, 0x45, 0x42, 0x12, 0x41, 0x0a, 0x0d, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x72, 0x61, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x78, 0x61, 0x74, 0x75, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x61, 0x48,
	0x00, 0x52, 0x07, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x52, 0x41, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x78, 0x61, 0x74, 0x75,
	0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x67, 0x0a, 0x0c, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x48, 0x41, 0x53, 0x45, 0x30, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4c, 0x54, 0x41, 0x49, 0x52, 0x10, 0x02, 0x12, 0x0d, 0x0a,
	0x09, 0x42, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x52, 0x49, 0x58, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07,
	0x43, 0x41, 0x50, 0x45, 0x4c, 0x4c, 0x41, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x4e,
	0x45, 0x42, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x52, 0x41, 0x10,
	0x06, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x74, 0x68, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x70, 0x73, 0x2f, 0x78, 0x61, 0x74, 0x75,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76,
	0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_eth_v2_events_proto_rawDescOnce sync.Once
	file_pkg_proto_eth_v2_events_proto_rawDescData = file_pkg_proto_eth_v2_events_proto_rawDesc
)

func file_pkg_proto_eth_v2_events_proto_rawDescGZIP() []byte {
	file_pkg_proto_eth_v2_events_proto_rawDescOnce.Do(func() {
		file_pkg_proto_eth_v2_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_eth_v2_events_proto_rawDescData)
	})
	return file_pkg_proto_eth_v2_events_proto_rawDescData
}

var file_pkg_proto_eth_v2_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_proto_eth_v2_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_proto_eth_v2_events_proto_goTypes = []interface{}{
	(BlockVersion)(0),              // 0: xatu.eth.v2.BlockVersion
	(*EventBlock)(nil),             // 1: xatu.eth.v2.EventBlock
	(*EventBlockV2)(nil),           // 2: xatu.eth.v2.EventBlockV2
	(*v1.BeaconBlock)(nil),         // 3: xatu.eth.v1.BeaconBlock
	(*BeaconBlockAltair)(nil),      // 4: xatu.eth.v2.BeaconBlockAltair
	(*BeaconBlockBellatrix)(nil),   // 5: xatu.eth.v2.BeaconBlockBellatrix
	(*BeaconBlockCapella)(nil),     // 6: xatu.eth.v2.BeaconBlockCapella
	(*BeaconBlockDeneb)(nil),       // 7: xatu.eth.v2.BeaconBlockDeneb
	(*BeaconBlockElectra)(nil),     // 8: xatu.eth.v2.BeaconBlockElectra
	(*v1.BeaconBlockV2)(nil),       // 9: xatu.eth.v1.BeaconBlockV2
	(*BeaconBlockAltairV2)(nil),    // 10: xatu.eth.v2.BeaconBlockAltairV2
	(*BeaconBlockBellatrixV2)(nil), // 11: xatu.eth.v2.BeaconBlockBellatrixV2
	(*BeaconBlockCapellaV2)(nil),   // 12: xatu.eth.v2.BeaconBlockCapellaV2
}
var file_pkg_proto_eth_v2_events_proto_depIdxs = []int32{
	3,  // 0: xatu.eth.v2.EventBlock.phase0_block:type_name -> xatu.eth.v1.BeaconBlock
	4,  // 1: xatu.eth.v2.EventBlock.altair_block:type_name -> xatu.eth.v2.BeaconBlockAltair
	5,  // 2: xatu.eth.v2.EventBlock.bellatrix_block:type_name -> xatu.eth.v2.BeaconBlockBellatrix
	6,  // 3: xatu.eth.v2.EventBlock.capella_block:type_name -> xatu.eth.v2.BeaconBlockCapella
	7,  // 4: xatu.eth.v2.EventBlock.deneb_block:type_name -> xatu.eth.v2.BeaconBlockDeneb
	8,  // 5: xatu.eth.v2.EventBlock.electra_block:type_name -> xatu.eth.v2.BeaconBlockElectra
	0,  // 6: xatu.eth.v2.EventBlock.version:type_name -> xatu.eth.v2.BlockVersion
	9,  // 7: xatu.eth.v2.EventBlockV2.phase0_block:type_name -> xatu.eth.v1.BeaconBlockV2
	10, // 8: xatu.eth.v2.EventBlockV2.altair_block:type_name -> xatu.eth.v2.BeaconBlockAltairV2
	11, // 9: xatu.eth.v2.EventBlockV2.bellatrix_block:type_name -> xatu.eth.v2.BeaconBlockBellatrixV2
	12, // 10: xatu.eth.v2.EventBlockV2.capella_block:type_name -> xatu.eth.v2.BeaconBlockCapellaV2
	7,  // 11: xatu.eth.v2.EventBlockV2.deneb_block:type_name -> xatu.eth.v2.BeaconBlockDeneb
	8,  // 12: xatu.eth.v2.EventBlockV2.electra_block:type_name -> xatu.eth.v2.BeaconBlockElectra
	0,  // 13: xatu.eth.v2.EventBlockV2.version:type_name -> xatu.eth.v2.BlockVersion
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_pkg_proto_eth_v2_events_proto_init() }
func file_pkg_proto_eth_v2_events_proto_init() {
	if File_pkg_proto_eth_v2_events_proto != nil {
		return
	}
	file_pkg_proto_eth_v2_beacon_block_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_eth_v2_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventBlock); i {
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
		file_pkg_proto_eth_v2_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventBlockV2); i {
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
	file_pkg_proto_eth_v2_events_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EventBlock_Phase0Block)(nil),
		(*EventBlock_AltairBlock)(nil),
		(*EventBlock_BellatrixBlock)(nil),
		(*EventBlock_CapellaBlock)(nil),
		(*EventBlock_DenebBlock)(nil),
		(*EventBlock_ElectraBlock)(nil),
	}
	file_pkg_proto_eth_v2_events_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*EventBlockV2_Phase0Block)(nil),
		(*EventBlockV2_AltairBlock)(nil),
		(*EventBlockV2_BellatrixBlock)(nil),
		(*EventBlockV2_CapellaBlock)(nil),
		(*EventBlockV2_DenebBlock)(nil),
		(*EventBlockV2_ElectraBlock)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_eth_v2_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_eth_v2_events_proto_goTypes,
		DependencyIndexes: file_pkg_proto_eth_v2_events_proto_depIdxs,
		EnumInfos:         file_pkg_proto_eth_v2_events_proto_enumTypes,
		MessageInfos:      file_pkg_proto_eth_v2_events_proto_msgTypes,
	}.Build()
	File_pkg_proto_eth_v2_events_proto = out.File
	file_pkg_proto_eth_v2_events_proto_rawDesc = nil
	file_pkg_proto_eth_v2_events_proto_goTypes = nil
	file_pkg_proto_eth_v2_events_proto_depIdxs = nil
}
