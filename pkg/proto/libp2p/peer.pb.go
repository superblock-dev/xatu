// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: pkg/proto/libp2p/peer.proto

package libp2p

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Peer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the ID of the peer.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Port is the port of the peer.
	Port *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty"`
	// IP is the IP of the peer.
	Ip string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	// UserAgent is the user agent of the peer.
	UserAgent string `protobuf:"bytes,4,opt,name=user_agent,proto3" json:"user_agent,omitempty"`
	// ProtocolVersion is the protocol version of the peer.
	ProtocolVersion string `protobuf:"bytes,5,opt,name=protocol_version,proto3" json:"protocol_version,omitempty"`
	// Protocols is the protocols of the peer.
	Protocols []string `protobuf:"bytes,6,rep,name=protocols,proto3" json:"protocols,omitempty"`
	// Latency is the latency of the peer (in milliseconds).
	Latency *wrapperspb.UInt64Value `protobuf:"bytes,7,opt,name=latency,proto3" json:"latency,omitempty"`
}

func (x *Peer) Reset() {
	*x = Peer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_libp2p_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Peer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Peer) ProtoMessage() {}

func (x *Peer) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_libp2p_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Peer.ProtoReflect.Descriptor instead.
func (*Peer) Descriptor() ([]byte, []int) {
	return file_pkg_proto_libp2p_peer_proto_rawDescGZIP(), []int{0}
}

func (x *Peer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Peer) GetPort() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *Peer) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Peer) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *Peer) GetProtocolVersion() string {
	if x != nil {
		return x.ProtocolVersion
	}
	return ""
}

func (x *Peer) GetProtocols() []string {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *Peer) GetLatency() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Latency
	}
	return nil
}

var File_pkg_proto_libp2p_peer_proto protoreflect.FileDescriptor

var file_pkg_proto_libp2p_peer_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x70,
	0x32, 0x70, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x78,
	0x61, 0x74, 0x75, 0x2e, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a, 0x04, 0x50,
	0x65, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12,
	0x36, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x74, 0x68, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x70,
	0x73, 0x2f, 0x78, 0x61, 0x74, 0x75, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6c, 0x69, 0x62, 0x70, 0x32, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_libp2p_peer_proto_rawDescOnce sync.Once
	file_pkg_proto_libp2p_peer_proto_rawDescData = file_pkg_proto_libp2p_peer_proto_rawDesc
)

func file_pkg_proto_libp2p_peer_proto_rawDescGZIP() []byte {
	file_pkg_proto_libp2p_peer_proto_rawDescOnce.Do(func() {
		file_pkg_proto_libp2p_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_libp2p_peer_proto_rawDescData)
	})
	return file_pkg_proto_libp2p_peer_proto_rawDescData
}

var file_pkg_proto_libp2p_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_libp2p_peer_proto_goTypes = []any{
	(*Peer)(nil),                   // 0: xatu.libp2p.Peer
	(*wrapperspb.UInt32Value)(nil), // 1: google.protobuf.UInt32Value
	(*wrapperspb.UInt64Value)(nil), // 2: google.protobuf.UInt64Value
}
var file_pkg_proto_libp2p_peer_proto_depIdxs = []int32{
	1, // 0: xatu.libp2p.Peer.port:type_name -> google.protobuf.UInt32Value
	2, // 1: xatu.libp2p.Peer.latency:type_name -> google.protobuf.UInt64Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_proto_libp2p_peer_proto_init() }
func file_pkg_proto_libp2p_peer_proto_init() {
	if File_pkg_proto_libp2p_peer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_libp2p_peer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Peer); i {
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
			RawDescriptor: file_pkg_proto_libp2p_peer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_libp2p_peer_proto_goTypes,
		DependencyIndexes: file_pkg_proto_libp2p_peer_proto_depIdxs,
		MessageInfos:      file_pkg_proto_libp2p_peer_proto_msgTypes,
	}.Build()
	File_pkg_proto_libp2p_peer_proto = out.File
	file_pkg_proto_libp2p_peer_proto_rawDesc = nil
	file_pkg_proto_libp2p_peer_proto_goTypes = nil
	file_pkg_proto_libp2p_peer_proto_depIdxs = nil
}
