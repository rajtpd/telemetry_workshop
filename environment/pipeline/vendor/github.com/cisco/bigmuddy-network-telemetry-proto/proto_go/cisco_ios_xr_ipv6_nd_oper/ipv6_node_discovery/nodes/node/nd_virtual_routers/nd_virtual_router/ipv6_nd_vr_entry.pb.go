// Code generated by protoc-gen-go.
// source: ipv6_nd_vr_entry.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_nd_virtual_routers_nd_virtual_router is a generated protocol buffer package.

It is generated from these files:
	ipv6_nd_vr_entry.proto

It has these top-level messages:
	Ipv6NdVrEntry_KEYS
	Ipv6NdVrEntry
	Ipv6NdAddr
*/
package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_nd_virtual_routers_nd_virtual_router

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Detailed Info of ND IPv6 Virtual Router entry
type Ipv6NdVrEntry_KEYS struct {
	NodeName      string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	InterfaceName string `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *Ipv6NdVrEntry_KEYS) Reset()                    { *m = Ipv6NdVrEntry_KEYS{} }
func (m *Ipv6NdVrEntry_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6NdVrEntry_KEYS) ProtoMessage()               {}
func (*Ipv6NdVrEntry_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6NdVrEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6NdVrEntry_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ipv6NdVrEntry struct {
	// Link-Layer Address
	LinkLayerAddress string `protobuf:"bytes,50,opt,name=link_layer_address,json=linkLayerAddress" json:"link_layer_address,omitempty"`
	// Link local address
	LocalAddress *Ipv6NdAddr `protobuf:"bytes,51,opt,name=local_address,json=localAddress" json:"local_address,omitempty"`
	// Virtual Router ID
	Context uint32 `protobuf:"varint,52,opt,name=context" json:"context,omitempty"`
	// VR state
	State string `protobuf:"bytes,53,opt,name=state" json:"state,omitempty"`
	// VR Flags
	Flags string `protobuf:"bytes,54,opt,name=flags" json:"flags,omitempty"`
	// Virtual Global Address Count
	VrGlAddrCt uint32 `protobuf:"varint,55,opt,name=vr_gl_addr_ct,json=vrGlAddrCt" json:"vr_gl_addr_ct,omitempty"`
	// List of ND global addresses
	VrGlobalAddressList []*Ipv6NdAddr `protobuf:"bytes,56,rep,name=vr_global_address_list,json=vrGlobalAddressList" json:"vr_global_address_list,omitempty"`
}

func (m *Ipv6NdVrEntry) Reset()                    { *m = Ipv6NdVrEntry{} }
func (m *Ipv6NdVrEntry) String() string            { return proto.CompactTextString(m) }
func (*Ipv6NdVrEntry) ProtoMessage()               {}
func (*Ipv6NdVrEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6NdVrEntry) GetLinkLayerAddress() string {
	if m != nil {
		return m.LinkLayerAddress
	}
	return ""
}

func (m *Ipv6NdVrEntry) GetLocalAddress() *Ipv6NdAddr {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *Ipv6NdVrEntry) GetContext() uint32 {
	if m != nil {
		return m.Context
	}
	return 0
}

func (m *Ipv6NdVrEntry) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Ipv6NdVrEntry) GetFlags() string {
	if m != nil {
		return m.Flags
	}
	return ""
}

func (m *Ipv6NdVrEntry) GetVrGlAddrCt() uint32 {
	if m != nil {
		return m.VrGlAddrCt
	}
	return 0
}

func (m *Ipv6NdVrEntry) GetVrGlobalAddressList() []*Ipv6NdAddr {
	if m != nil {
		return m.VrGlobalAddressList
	}
	return nil
}

// List of IPv6 ND Addresses
type Ipv6NdAddr struct {
	// IPv6 address
	Ipv6Address string `protobuf:"bytes,1,opt,name=ipv6_address,json=ipv6Address" json:"ipv6_address,omitempty"`
}

func (m *Ipv6NdAddr) Reset()                    { *m = Ipv6NdAddr{} }
func (m *Ipv6NdAddr) String() string            { return proto.CompactTextString(m) }
func (*Ipv6NdAddr) ProtoMessage()               {}
func (*Ipv6NdAddr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ipv6NdAddr) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv6NdVrEntry_KEYS)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.nd_virtual_routers.nd_virtual_router.ipv6_nd_vr_entry_KEYS")
	proto.RegisterType((*Ipv6NdVrEntry)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.nd_virtual_routers.nd_virtual_router.ipv6_nd_vr_entry")
	proto.RegisterType((*Ipv6NdAddr)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.nd_virtual_routers.nd_virtual_router.ipv6_nd_addr")
}

func init() { proto.RegisterFile("ipv6_nd_vr_entry.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x4a, 0x2b, 0x31,
	0x14, 0xc6, 0x99, 0x5b, 0xee, 0xbd, 0x36, 0xed, 0x48, 0x89, 0x5a, 0x02, 0x6e, 0xa6, 0x05, 0xa1,
	0x0b, 0x19, 0xb0, 0xd5, 0xea, 0x56, 0x44, 0x5c, 0x58, 0x5c, 0xd4, 0x95, 0x88, 0x1c, 0xd2, 0x99,
	0xb4, 0x84, 0xa6, 0x93, 0x72, 0x92, 0x0e, 0xed, 0x5b, 0xf8, 0x08, 0xbe, 0x99, 0xaf, 0x22, 0xc9,
	0xcc, 0x54, 0xa9, 0x6b, 0xdd, 0x0c, 0x7c, 0x5f, 0x7e, 0xf9, 0xce, 0x9f, 0x09, 0x69, 0xcb, 0x65,
	0x3e, 0x84, 0x2c, 0x85, 0x1c, 0x41, 0x64, 0x16, 0x37, 0xf1, 0x12, 0xb5, 0xd5, 0xf4, 0x25, 0x91,
	0x26, 0xd1, 0x20, 0xb5, 0x81, 0x35, 0x42, 0x05, 0xe9, 0xa5, 0xc0, 0xb8, 0x10, 0x3a, 0x15, 0x90,
	0x3a, 0x26, 0x17, 0xb8, 0x89, 0x9d, 0x34, 0xfe, 0x1b, 0xbb, 0x2c, 0x89, 0x76, 0xc5, 0x15, 0xa0,
	0x5e, 0x59, 0x81, 0xe6, 0xbb, 0xd5, 0x7d, 0x26, 0x47, 0xbb, 0x85, 0xe1, 0xfe, 0xf6, 0xe9, 0x91,
	0x1e, 0x93, 0xba, 0x8f, 0xce, 0xf8, 0x42, 0xb0, 0x20, 0x0a, 0x7a, 0xf5, 0xf1, 0x9e, 0x33, 0x1e,
	0xf8, 0x42, 0xd0, 0x13, 0xb2, 0x2f, 0x33, 0x2b, 0x70, 0xca, 0x93, 0x92, 0xf8, 0xe3, 0x89, 0x70,
	0xeb, 0x3a, 0xac, 0xfb, 0x5e, 0x23, 0xad, 0xdd, 0x74, 0x7a, 0x4a, 0xa8, 0x92, 0xd9, 0x1c, 0x14,
	0xdf, 0x08, 0x04, 0x9e, 0xa6, 0x28, 0x8c, 0x61, 0x7d, 0x7f, 0xbf, 0xe5, 0x4e, 0x46, 0xee, 0xe0,
	0xba, 0xf0, 0xe9, 0x6b, 0x40, 0x42, 0xa5, 0x13, 0xae, 0xb6, 0xe4, 0x20, 0x0a, 0x7a, 0x8d, 0xfe,
	0x3c, 0xfe, 0xd1, 0xbd, 0xc4, 0x55, 0xa0, 0xab, 0x3a, 0x6e, 0xfa, 0x0e, 0xaa, 0x96, 0x18, 0xf9,
	0x9f, 0xe8, 0xcc, 0x8a, 0xb5, 0x65, 0xe7, 0x51, 0xd0, 0x0b, 0xc7, 0x95, 0xa4, 0x87, 0xe4, 0xaf,
	0xb1, 0xdc, 0x0a, 0x76, 0xe1, 0xa7, 0x29, 0x84, 0x73, 0xa7, 0x8a, 0xcf, 0x0c, 0x1b, 0x16, 0xae,
	0x17, 0xb4, 0x43, 0xc2, 0x1c, 0x61, 0x56, 0xcc, 0x05, 0x89, 0x65, 0x97, 0x3e, 0x8b, 0xe4, 0x78,
	0xe7, 0x2b, 0xdd, 0x58, 0xfa, 0x16, 0x90, 0xb6, 0x67, 0xf4, 0xe4, 0x73, 0x7e, 0x50, 0xd2, 0x58,
	0x76, 0x15, 0xd5, 0x7e, 0x7b, 0x09, 0x07, 0xae, 0x33, 0xd7, 0x49, 0xb9, 0x87, 0x91, 0x34, 0xb6,
	0x7b, 0x46, 0x9a, 0x5f, 0x21, 0xda, 0x29, 0x75, 0xf5, 0xb3, 0x8a, 0x87, 0xd3, 0x70, 0x5e, 0x79,
	0x6d, 0xf2, 0xcf, 0xbf, 0xeb, 0xc1, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xd1, 0x33, 0xa6,
	0xf1, 0x02, 0x00, 0x00,
}
