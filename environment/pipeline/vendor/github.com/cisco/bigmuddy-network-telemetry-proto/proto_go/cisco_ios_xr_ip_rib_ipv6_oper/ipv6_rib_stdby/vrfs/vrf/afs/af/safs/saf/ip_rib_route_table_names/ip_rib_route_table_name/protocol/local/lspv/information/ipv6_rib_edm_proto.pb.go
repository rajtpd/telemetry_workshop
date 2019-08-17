// Code generated by protoc-gen-go.
// source: ipv6_rib_edm_proto.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_local_lspv_information is a generated protocol buffer package.

It is generated from these files:
	ipv6_rib_edm_proto.proto

It has these top-level messages:
	Ipv6RibEdmProto_KEYS
	Ipv6RibEdmProto
*/
package cisco_ios_xr_ip_rib_ipv6_oper_ipv6_rib_stdby_vrfs_vrf_afs_af_safs_saf_ip_rib_route_table_names_ip_rib_route_table_name_protocol_local_lspv_information

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

// Information of a rib protocol
type Ipv6RibEdmProto_KEYS struct {
	VrfName        string `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AfName         string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName        string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	RouteTableName string `protobuf:"bytes,4,opt,name=route_table_name,json=routeTableName" json:"route_table_name,omitempty"`
}

func (m *Ipv6RibEdmProto_KEYS) Reset()                    { *m = Ipv6RibEdmProto_KEYS{} }
func (m *Ipv6RibEdmProto_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmProto_KEYS) ProtoMessage()               {}
func (*Ipv6RibEdmProto_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ipv6RibEdmProto_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ipv6RibEdmProto_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *Ipv6RibEdmProto_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *Ipv6RibEdmProto_KEYS) GetRouteTableName() string {
	if m != nil {
		return m.RouteTableName
	}
	return ""
}

type Ipv6RibEdmProto struct {
	// Name
	ProtocolNames string `protobuf:"bytes,50,opt,name=protocol_names,json=protocolNames" json:"protocol_names,omitempty"`
	// Instance
	Instance string `protobuf:"bytes,51,opt,name=instance" json:"instance,omitempty"`
	// Proto version
	Version uint32 `protobuf:"varint,52,opt,name=version" json:"version,omitempty"`
	// Number of redist clients
	RedistributionClientCount uint32 `protobuf:"varint,53,opt,name=redistribution_client_count,json=redistributionClientCount" json:"redistribution_client_count,omitempty"`
	// Number of proto clients
	ProtocolClientsCount uint32 `protobuf:"varint,54,opt,name=protocol_clients_count,json=protocolClientsCount" json:"protocol_clients_count,omitempty"`
	// Number of routes (including active, backup and deleted), where, number of backup routes = RoutesCounts - ActiveRoutesCount - DeletedRoutesCount
	RoutesCounts uint32 `protobuf:"varint,55,opt,name=routes_counts,json=routesCounts" json:"routes_counts,omitempty"`
	// Number of active routes (not deleted)
	ActiveRoutesCount uint32 `protobuf:"varint,56,opt,name=active_routes_count,json=activeRoutesCount" json:"active_routes_count,omitempty"`
	// Number of deleted routes
	DeletedRoutesCount uint32 `protobuf:"varint,57,opt,name=deleted_routes_count,json=deletedRoutesCount" json:"deleted_routes_count,omitempty"`
	// Number of paths for all routes
	PathsCount uint32 `protobuf:"varint,58,opt,name=paths_count,json=pathsCount" json:"paths_count,omitempty"`
	// Memory for proto's routes and paths in bytes
	ProtocolRouteMemory uint32 `protobuf:"varint,59,opt,name=protocol_route_memory,json=protocolRouteMemory" json:"protocol_route_memory,omitempty"`
	// Number of backup routes
	BackupRoutesCount uint32 `protobuf:"varint,60,opt,name=backup_routes_count,json=backupRoutesCount" json:"backup_routes_count,omitempty"`
}

func (m *Ipv6RibEdmProto) Reset()                    { *m = Ipv6RibEdmProto{} }
func (m *Ipv6RibEdmProto) String() string            { return proto.CompactTextString(m) }
func (*Ipv6RibEdmProto) ProtoMessage()               {}
func (*Ipv6RibEdmProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv6RibEdmProto) GetProtocolNames() string {
	if m != nil {
		return m.ProtocolNames
	}
	return ""
}

func (m *Ipv6RibEdmProto) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Ipv6RibEdmProto) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetRedistributionClientCount() uint32 {
	if m != nil {
		return m.RedistributionClientCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetProtocolClientsCount() uint32 {
	if m != nil {
		return m.ProtocolClientsCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetRoutesCounts() uint32 {
	if m != nil {
		return m.RoutesCounts
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetActiveRoutesCount() uint32 {
	if m != nil {
		return m.ActiveRoutesCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetDeletedRoutesCount() uint32 {
	if m != nil {
		return m.DeletedRoutesCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetPathsCount() uint32 {
	if m != nil {
		return m.PathsCount
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetProtocolRouteMemory() uint32 {
	if m != nil {
		return m.ProtocolRouteMemory
	}
	return 0
}

func (m *Ipv6RibEdmProto) GetBackupRoutesCount() uint32 {
	if m != nil {
		return m.BackupRoutesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6RibEdmProto_KEYS)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.information.ipv6_rib_edm_proto_KEYS")
	proto.RegisterType((*Ipv6RibEdmProto)(nil), "cisco_ios_xr_ip_rib_ipv6_oper.ipv6_rib_stdby.vrfs.vrf.afs.af.safs.saf.ip_rib_route_table_names.ip_rib_route_table_name.protocol.local.lspv.information.ipv6_rib_edm_proto")
}

func init() { proto.RegisterFile("ipv6_rib_edm_proto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x14, 0x35, 0x61, 0x20, 0x15, 0xb8, 0x85, 0x3a, 0x70, 0xa0, 0x2a, 0x42, 0xca,
	0xc9, 0x42, 0x6d, 0x29, 0x7f, 0xc5, 0xa5, 0xe2, 0x84, 0xe0, 0x10, 0xb8, 0x70, 0xb2, 0xbc, 0x5e,
	0xaf, 0xb0, 0xd8, 0xb5, 0x57, 0x1e, 0x67, 0x45, 0x5f, 0x83, 0x03, 0x2f, 0xca, 0x0b, 0x20, 0x8f,
	0x77, 0x43, 0x42, 0xc4, 0xc5, 0x91, 0xe7, 0xfb, 0xfd, 0x9c, 0x6f, 0x12, 0xe0, 0xb6, 0xeb, 0x2f,
	0x65, 0xb0, 0xa5, 0x34, 0x55, 0x2b, 0xbb, 0xe0, 0xa3, 0x17, 0x74, 0xb2, 0x5f, 0x85, 0xb6, 0xa8,
	0xbd, 0xb4, 0x1e, 0xe5, 0x8f, 0x20, 0x6d, 0x47, 0x14, 0xe1, 0xbe, 0x33, 0x41, 0xac, 0x45, 0x8c,
	0x55, 0x79, 0x2d, 0xfa, 0x50, 0x63, 0x3a, 0x84, 0xaa, 0x51, 0xa8, 0x5a, 0x60, 0xfa, 0x44, 0x55,
	0x8b, 0x41, 0x0c, 0x7e, 0x15, 0x8d, 0x8c, 0xaa, 0x6c, 0x8c, 0x74, 0xaa, 0x35, 0xf8, 0xbf, 0x20,
	0x7f, 0xbd, 0xf6, 0x8d, 0x68, 0xbc, 0x56, 0x8d, 0x68, 0xb0, 0xeb, 0x85, 0x75, 0xb5, 0x0f, 0xad,
	0x8a, 0xd6, 0xbb, 0xd3, 0x9f, 0x05, 0x1c, 0xef, 0xb6, 0x96, 0x1f, 0xde, 0x7f, 0xfd, 0xcc, 0xe6,
	0x30, 0xed, 0x43, 0x4d, 0xef, 0xf0, 0xe2, 0xa4, 0x58, 0xdc, 0x5a, 0x4e, 0xfa, 0x50, 0x7f, 0x52,
	0xad, 0x61, 0xc7, 0x30, 0x51, 0x43, 0x72, 0x83, 0x92, 0x7d, 0x95, 0x83, 0x39, 0x4c, 0x71, 0x4c,
	0xf6, 0xb2, 0x83, 0x43, 0xb4, 0x80, 0xbb, 0xff, 0xd6, 0xe3, 0x37, 0x09, 0x39, 0xa0, 0xf9, 0x97,
	0x34, 0x4e, 0xe4, 0xe9, 0xef, 0x3d, 0x60, 0xbb, 0xa5, 0xd8, 0x53, 0x38, 0x18, 0xd7, 0xc9, 0x5b,
	0xf3, 0x33, 0xd2, 0x67, 0xe3, 0x34, 0xc9, 0xc8, 0x1e, 0xc2, 0xd4, 0x3a, 0x8c, 0xca, 0x69, 0xc3,
	0xcf, 0x09, 0x58, 0xdf, 0x19, 0x87, 0x49, 0x6f, 0x02, 0x5a, 0xef, 0xf8, 0xc5, 0x49, 0xb1, 0x98,
	0x2d, 0xc7, 0x2b, 0x7b, 0x07, 0x8f, 0x82, 0xa9, 0x2c, 0xc6, 0x60, 0xcb, 0x55, 0xfa, 0x69, 0xa4,
	0x6e, 0xac, 0x71, 0x51, 0x6a, 0xbf, 0x72, 0x91, 0x3f, 0x27, 0x7a, 0xbe, 0x8d, 0x5c, 0x11, 0x71,
	0x95, 0x00, 0x76, 0x01, 0x0f, 0xd6, 0xe5, 0xb2, 0x89, 0x83, 0x7a, 0x49, 0xea, 0xd1, 0x98, 0x66,
	0x09, 0xb3, 0xf5, 0x04, 0x66, 0xb4, 0xfb, 0xc0, 0x22, 0x7f, 0x41, 0xf0, 0x9d, 0x3c, 0x24, 0x06,
	0x99, 0x80, 0x43, 0xa5, 0xa3, 0xed, 0x8d, 0xdc, 0x64, 0xf9, 0x4b, 0x42, 0xef, 0xe5, 0x68, 0xf9,
	0x57, 0x60, 0xcf, 0xe0, 0xa8, 0x32, 0x8d, 0x89, 0xa6, 0xda, 0x16, 0x5e, 0x91, 0xc0, 0x86, 0x6c,
	0xd3, 0x78, 0x0c, 0xb7, 0x3b, 0x15, 0xbf, 0x8d, 0xe0, 0x6b, 0x02, 0x81, 0x46, 0x19, 0x38, 0x83,
	0xfb, 0xeb, 0xed, 0xf2, 0x9f, 0xd8, 0x9a, 0xd6, 0x87, 0x6b, 0xfe, 0x86, 0xd0, 0xc3, 0x31, 0xa4,
	0x47, 0x3f, 0x52, 0x94, 0x6a, 0x97, 0x4a, 0x7f, 0x5f, 0x75, 0xdb, 0x2d, 0xde, 0xe6, 0xda, 0x39,
	0xda, 0x28, 0x51, 0xee, 0xd3, 0x23, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x74, 0x9d, 0xc6,
	0xfc, 0x46, 0x03, 0x00, 0x00,
}
