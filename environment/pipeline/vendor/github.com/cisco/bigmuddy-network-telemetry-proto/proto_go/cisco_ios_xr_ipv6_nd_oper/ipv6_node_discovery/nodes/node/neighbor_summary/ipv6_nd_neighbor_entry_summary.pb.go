// Code generated by protoc-gen-go.
// source: ipv6_nd_neighbor_entry_summary.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_neighbor_summary is a generated protocol buffer package.

It is generated from these files:
	ipv6_nd_neighbor_entry_summary.proto

It has these top-level messages:
	Ipv6NdNeighborEntrySummary_KEYS
	Ipv6NdNeighborEntrySummary
	BagNbrEntrySum
*/
package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_neighbor_summary

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

// IPv6 ND neighbor entry summary
type Ipv6NdNeighborEntrySummary_KEYS struct {
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
}

func (m *Ipv6NdNeighborEntrySummary_KEYS) Reset()                    { *m = Ipv6NdNeighborEntrySummary_KEYS{} }
func (m *Ipv6NdNeighborEntrySummary_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6NdNeighborEntrySummary_KEYS) ProtoMessage()               {}
func (*Ipv6NdNeighborEntrySummary_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6NdNeighborEntrySummary_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type Ipv6NdNeighborEntrySummary struct {
	// Multicast neighbor summary
	Multicast *BagNbrEntrySum `protobuf:"bytes,50,opt,name=multicast" json:"multicast,omitempty"`
	// Static neighbor summary
	Static *BagNbrEntrySum `protobuf:"bytes,51,opt,name=static" json:"static,omitempty"`
	// Dynamic neighbor summary
	Dynamic *BagNbrEntrySum `protobuf:"bytes,52,opt,name=dynamic" json:"dynamic,omitempty"`
	// Total number of entries
	TotalNeighborEntries uint32 `protobuf:"varint,53,opt,name=total_neighbor_entries,json=totalNeighborEntries" json:"total_neighbor_entries,omitempty"`
}

func (m *Ipv6NdNeighborEntrySummary) Reset()                    { *m = Ipv6NdNeighborEntrySummary{} }
func (m *Ipv6NdNeighborEntrySummary) String() string            { return proto.CompactTextString(m) }
func (*Ipv6NdNeighborEntrySummary) ProtoMessage()               {}
func (*Ipv6NdNeighborEntrySummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6NdNeighborEntrySummary) GetMulticast() *BagNbrEntrySum {
	if m != nil {
		return m.Multicast
	}
	return nil
}

func (m *Ipv6NdNeighborEntrySummary) GetStatic() *BagNbrEntrySum {
	if m != nil {
		return m.Static
	}
	return nil
}

func (m *Ipv6NdNeighborEntrySummary) GetDynamic() *BagNbrEntrySum {
	if m != nil {
		return m.Dynamic
	}
	return nil
}

func (m *Ipv6NdNeighborEntrySummary) GetTotalNeighborEntries() uint32 {
	if m != nil {
		return m.TotalNeighborEntries
	}
	return 0
}

// IPv6 ND summary information
type BagNbrEntrySum struct {
	// Total incomplete entries
	IncompleteEntries uint32 `protobuf:"varint,1,opt,name=incomplete_entries,json=incompleteEntries" json:"incomplete_entries,omitempty"`
	// Total reachable entries
	ReachableEntries uint32 `protobuf:"varint,2,opt,name=reachable_entries,json=reachableEntries" json:"reachable_entries,omitempty"`
	// Total stale entries
	StaleEntries uint32 `protobuf:"varint,3,opt,name=stale_entries,json=staleEntries" json:"stale_entries,omitempty"`
	// Total delayed entries
	DelayedEntries uint32 `protobuf:"varint,4,opt,name=delayed_entries,json=delayedEntries" json:"delayed_entries,omitempty"`
	// Total probe entries
	ProbeEntries uint32 `protobuf:"varint,5,opt,name=probe_entries,json=probeEntries" json:"probe_entries,omitempty"`
	// Total deleted entries
	DeletedEntries uint32 `protobuf:"varint,6,opt,name=deleted_entries,json=deletedEntries" json:"deleted_entries,omitempty"`
	// Total number of entries
	SubtotalNeighborEntries uint32 `protobuf:"varint,7,opt,name=subtotal_neighbor_entries,json=subtotalNeighborEntries" json:"subtotal_neighbor_entries,omitempty"`
}

func (m *BagNbrEntrySum) Reset()                    { *m = BagNbrEntrySum{} }
func (m *BagNbrEntrySum) String() string            { return proto.CompactTextString(m) }
func (*BagNbrEntrySum) ProtoMessage()               {}
func (*BagNbrEntrySum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BagNbrEntrySum) GetIncompleteEntries() uint32 {
	if m != nil {
		return m.IncompleteEntries
	}
	return 0
}

func (m *BagNbrEntrySum) GetReachableEntries() uint32 {
	if m != nil {
		return m.ReachableEntries
	}
	return 0
}

func (m *BagNbrEntrySum) GetStaleEntries() uint32 {
	if m != nil {
		return m.StaleEntries
	}
	return 0
}

func (m *BagNbrEntrySum) GetDelayedEntries() uint32 {
	if m != nil {
		return m.DelayedEntries
	}
	return 0
}

func (m *BagNbrEntrySum) GetProbeEntries() uint32 {
	if m != nil {
		return m.ProbeEntries
	}
	return 0
}

func (m *BagNbrEntrySum) GetDeletedEntries() uint32 {
	if m != nil {
		return m.DeletedEntries
	}
	return 0
}

func (m *BagNbrEntrySum) GetSubtotalNeighborEntries() uint32 {
	if m != nil {
		return m.SubtotalNeighborEntries
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6NdNeighborEntrySummary_KEYS)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.neighbor_summary.ipv6_nd_neighbor_entry_summary_KEYS")
	proto.RegisterType((*Ipv6NdNeighborEntrySummary)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.neighbor_summary.ipv6_nd_neighbor_entry_summary")
	proto.RegisterType((*BagNbrEntrySum)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.neighbor_summary.bag_nbr_entry_sum")
}

func init() { proto.RegisterFile("ipv6_nd_neighbor_entry_summary.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x86, 0x49, 0x7b, 0x6f, 0x7b, 0x7b, 0xae, 0x55, 0x3b, 0x88, 0x46, 0x04, 0x29, 0xad, 0x60,
	0x41, 0xcc, 0xa2, 0xad, 0x2e, 0x5c, 0x0a, 0x5d, 0x88, 0xd0, 0x45, 0x5d, 0x09, 0xc2, 0x30, 0x49,
	0x0e, 0xed, 0x40, 0x92, 0x09, 0x33, 0xd3, 0x62, 0xfc, 0x75, 0xee, 0xfc, 0x5b, 0x92, 0x69, 0xbe,
	0xaa, 0xd8, 0x95, 0x74, 0x13, 0x98, 0xf3, 0x3e, 0xf3, 0xbc, 0xe4, 0x84, 0xc0, 0x05, 0x8f, 0x57,
	0xb7, 0x34, 0xf2, 0x69, 0x84, 0x7c, 0xbe, 0x70, 0x85, 0xa4, 0x18, 0x69, 0x99, 0x50, 0xb5, 0x0c,
	0x43, 0x26, 0x13, 0x27, 0x96, 0x42, 0x0b, 0xf2, 0xe0, 0x71, 0xe5, 0x09, 0xca, 0x85, 0xa2, 0xaf,
	0x92, 0xe6, 0x57, 0x44, 0x8c, 0xd2, 0x59, 0x1f, 0x84, 0x8f, 0xd4, 0x4f, 0x99, 0x15, 0xca, 0xc4,
	0x49, 0x8f, 0xca, 0x3c, 0x9d, 0x42, 0x9b, 0x09, 0x7b, 0xf7, 0xd0, 0xdf, 0x5e, 0x49, 0x1f, 0x27,
	0xcf, 0x4f, 0xe4, 0x0c, 0x5a, 0x46, 0x1a, 0xb1, 0x10, 0x6d, 0xab, 0x6b, 0x0d, 0x5a, 0xb3, 0x7f,
	0xe9, 0x60, 0xca, 0x42, 0xec, 0xbd, 0xd7, 0xe1, 0x7c, 0xbb, 0x84, 0xbc, 0x41, 0x2b, 0x5c, 0x06,
	0x9a, 0x7b, 0x4c, 0x69, 0x7b, 0xd8, 0xb5, 0x06, 0xff, 0x87, 0x2f, 0xce, 0xaf, 0xbd, 0x85, 0xe3,
	0xb2, 0x39, 0x8d, 0xdc, 0x4a, 0xe9, 0xac, 0xac, 0x23, 0x1a, 0x1a, 0x4a, 0x33, 0xcd, 0x3d, 0x7b,
	0xb4, 0x83, 0xe2, 0xac, 0x8b, 0xac, 0xa0, 0xe9, 0x27, 0x11, 0x0b, 0xb9, 0x67, 0x8f, 0x77, 0x50,
	0x9b, 0x97, 0x91, 0x31, 0x1c, 0x6b, 0xa1, 0x59, 0xb0, 0xf9, 0x25, 0x38, 0x2a, 0xfb, 0xa6, 0x6b,
	0x0d, 0xda, 0xb3, 0x23, 0x93, 0x4e, 0xb3, 0x70, 0xb2, 0xce, 0x7a, 0x1f, 0x35, 0xe8, 0x7c, 0x93,
	0x92, 0x6b, 0x20, 0x3c, 0xf2, 0x44, 0x18, 0x07, 0xa8, 0xb1, 0xf0, 0x58, 0xc6, 0xd3, 0x29, 0x93,
	0x4c, 0x42, 0xae, 0xa0, 0x23, 0x91, 0x79, 0x0b, 0xe6, 0x06, 0x25, 0x5d, 0x33, 0xf4, 0x61, 0x11,
	0xe4, 0x70, 0x1f, 0xda, 0x4a, 0xb3, 0x0a, 0x58, 0x37, 0xe0, 0x9e, 0x19, 0xe6, 0xd0, 0x25, 0x1c,
	0xf8, 0x18, 0xb0, 0x04, 0xfd, 0x02, 0xfb, 0x63, 0xb0, 0xfd, 0x6c, 0x5c, 0xb1, 0xc5, 0x52, 0xb8,
	0xa5, 0xed, 0xef, 0xda, 0x66, 0x86, 0x9b, 0x36, 0xd4, 0x15, 0x5b, 0xa3, 0xb0, 0xa5, 0xe3, 0x1c,
	0xbc, 0x83, 0x53, 0xb5, 0x74, 0x7f, 0x58, 0x63, 0xd3, 0x5c, 0x39, 0xc9, 0x81, 0x2f, 0x9b, 0x74,
	0x1b, 0xe6, 0x17, 0x1d, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0x19, 0xb9, 0x58, 0x67, 0xca, 0x03,
	0x00, 0x00,
}
