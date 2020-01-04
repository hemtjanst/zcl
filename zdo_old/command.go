package zdo_old

import "hemtjan.st/zcl"

func (NWKAddrRequest) Cluster() uint16      { return 0x0000 }
func (n *NWKAddrRequest) Values() []zcl.Val { return []zcl.Val{&n.IEEE, &n.RequestType, &n.StartIndex} }

type NWKAddrRequest struct {
	IEEE zcl.Zuid
	// RequestType
	// 0x00 = single device response
	// 0x01 = extended response
	RequestType zcl.Zu8
	StartIndex  zcl.Zu8
}

func (IEEEAddrRequest) Cluster() uint16      { return 0x0001 }
func (n *IEEEAddrRequest) Values() []zcl.Val { return []zcl.Val{&n.NWKI, &n.RequestType, &n.StartIndex} }

type IEEEAddrRequest struct {
	NWKI zcl.Zu16
	// RequestType
	// 0x00 = single device response
	// 0x01 = extended response
	RequestType zcl.Zu8
	StartIndex  zcl.Zu8
}

func (NodeDescRequest) Cluster() uint16      { return 0x0002 }
func (n *NodeDescRequest) Values() []zcl.Val { return []zcl.Val{&n.NWKI} }

type NodeDescRequest struct {
	NWKI zcl.Zu16
}

func (PowerDescRequest) Cluster() uint16      { return 0x0003 }
func (n *PowerDescRequest) Values() []zcl.Val { return []zcl.Val{&n.NWKI} }

type PowerDescRequest struct {
	NWKI zcl.Zu16
}

func (SimpleDescRequest) Cluster() uint16      { return 0x0004 }
func (n *SimpleDescRequest) Values() []zcl.Val { return []zcl.Val{&n.NWKI, &n.Endpoint} }

type SimpleDescRequest struct {
	NWKI     zcl.Zu16
	Endpoint zcl.Zu8
}

func (ActiveEndpointRequest) Cluster() uint16      { return 0x0005 }
func (n *ActiveEndpointRequest) Values() []zcl.Val { return []zcl.Val{&n.NWKI} }

type ActiveEndpointRequest struct {
	NWKI zcl.Zu16
}

func (MatchDescRequest) Cluster() uint16 { return 0x0006 }
func (n *MatchDescRequest) Values() []zcl.Val {
	return []zcl.Val{&n.NWKI, &n.ProfileID, &n.InClusterList, &n.OutClusterList}
}

type MatchDescRequest struct {
	NWKI           zcl.Zu16
	ProfileID      zcl.Zu16
	InClusterList  LVuint16
	OutClusterList LVuint16
}

func (ComplexDescRequest) Cluster() uint16      { return 0x0010 }
func (n *ComplexDescRequest) Values() []zcl.Val { return []zcl.Val{&n.NWKI} }

type ComplexDescRequest struct {
	NWKI zcl.Zu16
}

func (UserDescRequest) Cluster() uint16      { return 0x0011 }
func (n *UserDescRequest) Values() []zcl.Val { return []zcl.Val{&n.NWKI} }

type UserDescRequest struct {
	NWKI zcl.Zu16
}

func (DiscoveryCacheRequest) Cluster() uint16      { return 0x0012 }
func (n *DiscoveryCacheRequest) Values() []zcl.Val { return []zcl.Val{&n.NWK, &n.IEEE} }

type DiscoveryCacheRequest struct {
	NWK  zcl.Zu16
	IEEE zcl.Zuid
}

func (DeviceAnnounce) Cluster() uint16      { return 0x0013 }
func (n *DeviceAnnounce) Values() []zcl.Val { return []zcl.Val{&n.NWK, &n.IEEE, &n.Capability} }

type DeviceCapability zcl.Zbmp8

type DeviceAnnounce struct {
	NWK        zcl.Zu16
	IEEE       zcl.Zuid
	Capability DeviceCapability
}

func (n DeviceAnnounce) String() string {
	return zcl.Sprintf(
		"DeviceAnnounce{NWK:0x%04X IEEE:%s, Cap:(%s)}",
		uint16(n.NWK),
		n.IEEE.String(),
		n.Capability.String(),
	)
}

func (n DeviceCapability) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(n).MarshalZcl()
}
func (n *DeviceCapability) UnmarshalZcl(b []byte) ([]byte, error) {
	v := new(zcl.Zbmp8)
	b, err := v.UnmarshalZcl(b)
	*n = DeviceCapability(*v)
	return b, err
}

func (n DeviceCapability) String() string {
	var str []string
	if zcl.BitmapTest(n[:], 0) {
		str = append(str, "PAN Coordinator")
	}
	if zcl.BitmapTest(n[:], 1) {
		str = append(str, "Full Function")
	} else {
		str = append(str, "Reduced Function")
	}
	if zcl.BitmapTest(n[:], 2) {
		str = append(str, "Mains Power")
	} else {
		str = append(str, "Other Power")
	}
	if zcl.BitmapTest(n[:], 3) {
		str = append(str, "Recv on Idle")
	}
	if zcl.BitmapTest(n[:], 4) {
		str = append(str, "Reserved(4)")
	}
	if zcl.BitmapTest(n[:], 5) {
		str = append(str, "Reserved(5)")
	}
	if zcl.BitmapTest(n[:], 6) {
		str = append(str, "Security")
	}
	if zcl.BitmapTest(n[:], 7) {
		str = append(str, "Allocate Address")
	}
	return zcl.StrJoin(str, ", ")
}

func (SetUserDescRequest) Cluster() uint16      { return 0x0014 }
func (n *SetUserDescRequest) Values() []zcl.Val { return []zcl.Val{&n.NWKI, &n.UserDescriptor} }

type SetUserDescRequest struct {
	NWKI           zcl.Zu16
	UserDescriptor zcl.Zbytes
}

func (SystemServerDiscoveryRequest) Cluster() uint16      { return 0x0015 }
func (n *SystemServerDiscoveryRequest) Values() []zcl.Val { return []zcl.Val{&n.ServerMask} }

type SystemServerDiscoveryRequest struct {
	ServerMask zcl.Zbmp16
}

func (n SystemServerDiscoveryRequest) String() string {
	return zcl.Sprintf(
		"SystemServerDiscoveryRequest{Mask: %s}",
		n.ServerMaskString(),
	)
}

func (n SystemServerDiscoveryRequest) ServerMaskString() string {
	var str []string
	for _, bit := range zcl.BitmapList(n.ServerMask[:]) {
		switch bit {
		case 0:
			str = append(str, "Primary Trust Center")
		case 1:
			str = append(str, "Backup Trust Center")
		case 2:
			str = append(str, "Primary Binding Table Cache")
		case 3:
			str = append(str, "Backup Binding Table Cache")
		case 4:
			str = append(str, "Primary Discovery Cache")
		case 5:
			str = append(str, "Backup Discovery Cache")
		case 6:
			str = append(str, "Network Manager")
		default:
			str = append(str, zcl.Sprintf("Reserved(%d)", bit))
		}
	}
	return zcl.StrJoin(str, ", ")
}

func (DiscoveryStoreRequest) Cluster() uint16 { return 0x0016 }
func (n *DiscoveryStoreRequest) Values() []zcl.Val {
	return []zcl.Val{&n.NWK, &n.IEEE, &n.NodeDescSize, &n.PowerDescSize, &n.ActiveEPSize}
}

type DiscoveryStoreRequest struct {
	NWK                zcl.Zu16
	IEEE               zcl.Zuid
	NodeDescSize       zcl.Zu8
	PowerDescSize      zcl.Zu8
	ActiveEPSize       zcl.Zu8
	SimpleDescSizeList LVuint8
}

func (NodeDescStoreRequest) Cluster() uint16 { return 0x0017 }
func (n *NodeDescStoreRequest) Values() []zcl.Val {
	return []zcl.Val{
		&n.NWK, &n.IEEE, &n.Byte1, &n.Byte2, &n.MacCapabilityFlags, &n.ManufacturerCode, &n.MaximumBufferSize,
		&n.MaximumIncomingTransferSize, &n.ServerMask, &n.MaximumOutgoingTransferSize, &n.DescriptorCapabilityField,
	}
}

type NodeDescStoreRequest struct {
	NWK                         zcl.Zu16
	IEEE                        zcl.Zuid
	Byte1                       zcl.Zu8
	Byte2                       zcl.Zu8
	MacCapabilityFlags          zcl.Zu8
	ManufacturerCode            zcl.Zu16
	MaximumBufferSize           zcl.Zu8
	MaximumIncomingTransferSize zcl.Zu16
	ServerMask                  zcl.Zu16
	MaximumOutgoingTransferSize zcl.Zu16
	DescriptorCapabilityField   zcl.Zu8
}

func (ActiveEndpointStoreRequest) Cluster() uint16 { return 0x0019 }
func (n *ActiveEndpointStoreRequest) Values() []zcl.Val {
	return []zcl.Val{&n.NWK, &n.IEEE, &n.ActiveEPList}
}

type ActiveEndpointStoreRequest struct {
	NWK          zcl.Zu16
	IEEE         zcl.Zuid
	ActiveEPList zcl.Zbytes
}

func (SimpleDescStoreRequest) Cluster() uint16 { return 0x001a }
func (n *SimpleDescStoreRequest) Values() []zcl.Val {
	return []zcl.Val{&n.NWK, &n.IEEE, &n.LVSimpleDescriptor}
}

type SimpleDescStoreRequest struct {
	NWK  zcl.Zu16
	IEEE zcl.Zuid
	LVSimpleDescriptor
}

func (RemoveNodeCacheRequest) Cluster() uint16      { return 0x001b }
func (n *RemoveNodeCacheRequest) Values() []zcl.Val { return []zcl.Val{&n.NWK, &n.IEEE} }

type RemoveNodeCacheRequest struct {
	NWK  zcl.Zu16
	IEEE zcl.Zuid
}

func (FindNodeCacheRequest) Cluster() uint16      { return 0x001c }
func (n *FindNodeCacheRequest) Values() []zcl.Val { return []zcl.Val{&n.NWK, &n.IEEE} }

type FindNodeCacheRequest struct {
	NWK  zcl.Zu16
	IEEE zcl.Zuid
}

func (ExtendedSimpleDescRequest) Cluster() uint16 { return 0x001d }
func (n *ExtendedSimpleDescRequest) Values() []zcl.Val {
	return []zcl.Val{&n.NWKI, &n.Endpoint, &n.StartIndex}
}

type ExtendedSimpleDescRequest struct {
	NWKI       zcl.Zu16
	Endpoint   zcl.Zu8
	StartIndex zcl.Zu8
}

func (ExtendedActiveEndpointRequest) Cluster() uint16      { return 0x001e }
func (n *ExtendedActiveEndpointRequest) Values() []zcl.Val { return []zcl.Val{&n.NWKI, &n.StartIndex} }

type ExtendedActiveEndpointRequest struct {
	NWKI       zcl.Zu16
	StartIndex zcl.Zu8
}

func (EndDeviceBindRequest) Cluster() uint16 { return 0x0020 }
func (n *EndDeviceBindRequest) Values() []zcl.Val {
	return []zcl.Val{
		&n.BindingTarget, &n.SrcAddress, &n.SrcEndpoint, &n.ProfileID, &n.InClusterList, &n.OutClusterList,
	}
}

type EndDeviceBindRequest struct {
	BindingTarget  zcl.Zu16
	SrcAddress     zcl.Zuid
	SrcEndpoint    zcl.Zu8
	ProfileID      zcl.Zu8
	InClusterList  LVuint16
	OutClusterList LVuint16
}

func (BindRequest) Cluster() uint16 { return 0x0021 }
func (n *BindRequest) Values() []zcl.Val {
	return []zcl.Val{&n.SrcAddress, &n.SrcEndpoint, &n.ClusterStr, &n.DstAddress}
}

type BindRequest struct {
	SrcAddress  zcl.Zuid
	SrcEndpoint zcl.Zu8
	ClusterStr  zcl.Zu16
	DstAddress  MultiAddress
}

func (UnbindRequest) Cluster() uint16 { return 0x0022 }
func (n *UnbindRequest) Values() []zcl.Val {
	return []zcl.Val{&n.SrcAddress, &n.SrcEndpoint, &n.ClusterStr, &n.DstAddress}
}

type UnbindRequest struct {
	SrcAddress  zcl.Zuid
	SrcEndpoint zcl.Zu8
	ClusterStr  zcl.Zu16
	DstAddress  MultiAddress
}

func (MgmtLqiRequest) Cluster() uint16      { return 0x0031 }
func (n *MgmtLqiRequest) Values() []zcl.Val { return []zcl.Val{&n.StartIndex} }

type MgmtLqiRequest struct {
	StartIndex zcl.Zu8
}

func (MgmtRtgRequest) Cluster() uint16      { return 0x0032 }
func (n *MgmtRtgRequest) Values() []zcl.Val { return []zcl.Val{&n.StartIndex} }

type MgmtRtgRequest struct {
	StartIndex zcl.Zu8
}

func (MgmtLeaveRequest) Cluster() uint16      { return 0x0034 }
func (n *MgmtLeaveRequest) Values() []zcl.Val { return []zcl.Val{&n.DeviceAddress, &n.Options} }

type MgmtLeaveRequest struct {
	DeviceAddress zcl.Zuid
	Options       zcl.Zu8
}

func (MgmtPermitJoiningRequest) Cluster() uint16 { return 0x0036 }
func (n *MgmtPermitJoiningRequest) Values() []zcl.Val {
	return []zcl.Val{&n.PermitDuration, &n.TCSignificant}
}

type MgmtPermitJoiningRequest struct {
	PermitDuration zcl.Zu8
	TCSignificant  zcl.Zu8
}

func (NWKAddrResponse) Cluster() uint16 { return 0x8000 }
func (n *NWKAddrResponse) Values() []zcl.Val {
	return []zcl.Val{&n.Status, &n.IEEE, &n.NWK, &n.NumAssocDev, &n.StartIndex, &n.NWKAddrAssocDevList}
}

type NWKAddrResponse struct {
	Status              ResponseStatus
	IEEE                zcl.Zuid
	NWK                 zcl.Zu16
	NumAssocDev         zcl.Zu8
	StartIndex          zcl.Zu8
	NWKAddrAssocDevList Uint16List
}

func (IEEEAddrResponse) Cluster() uint16 { return 0x8001 }

func (n *IEEEAddrResponse) Values() []zcl.Val {
	return []zcl.Val{&n.Status, &n.IEEE, &n.NWK, &n.NumAssocDev, &n.StartIndex, &n.NWKAddrAssocDevList}
}

type IEEEAddrResponse struct {
	Status              ResponseStatus
	IEEE                zcl.Zuid
	NWK                 zcl.Zu16
	NumAssocDev         zcl.Zu8
	StartIndex          zcl.Zu8
	NWKAddrAssocDevList Uint16List
}

func (NodeDescResponse) Cluster() uint16 { return 0x8002 }

func (n *NodeDescResponse) Values() []zcl.Val {
	return []zcl.Val{
		&n.Status, &n.NWKI, &n.Byte1, &n.Byte2, &n.MacCapabilityFlags, &n.ManufacturerCode, &n.MaximumBufferSize,
		&n.MaximumIncomingTransferSize, &n.ServerMask, &n.MaximumOutgoingTransferSize, &n.DescriptorCapabilityField,
	}
}

type NodeDescResponse struct {
	Status                      ResponseStatus
	NWKI                        zcl.Zu16
	Byte1                       zcl.Zu8
	Byte2                       zcl.Zu8
	MacCapabilityFlags          zcl.Zu8
	ManufacturerCode            zcl.Zu16
	MaximumBufferSize           zcl.Zu8
	MaximumIncomingTransferSize zcl.Zu16
	ServerMask                  zcl.Zu16
	MaximumOutgoingTransferSize zcl.Zu16
	DescriptorCapabilityField   zcl.Zu8
}

func (PowerDescResponse) Cluster() uint16 { return 0x8003 }
func (n *PowerDescResponse) Values() []zcl.Val {
	return []zcl.Val{&n.Status, &n.NWKI, &n.Byte1, &n.Byte2}
}

type PowerDescResponse struct {
	Status ResponseStatus
	NWKI   zcl.Zu16
	Byte1  zcl.Zu8
	Byte2  zcl.Zu8
}

func (SimpleDescResponse) Cluster() uint16      { return 0x8004 }
func (n *SimpleDescResponse) Values() []zcl.Val { return []zcl.Val{&n.Status, &n.NWKI, &n.Desc} }

type SimpleDescResponse struct {
	Status ResponseStatus
	NWKI   zcl.Zu16
	Desc   LVSimpleDescriptor
}

func (ActiveEndpointResponse) Cluster() uint16 { return 0x8005 }
func (n *ActiveEndpointResponse) Values() []zcl.Val {
	return []zcl.Val{&n.Status, &n.NWKI, &n.ActiveEPList}
}

type ActiveEndpointResponse struct {
	Status       ResponseStatus
	NWKI         zcl.Zu16
	ActiveEPList LVuint8
}

func (MatchDescResponse) Cluster() uint16      { return 0x8006 }
func (n *MatchDescResponse) Values() []zcl.Val { return []zcl.Val{&n.Status, &n.MatchList} }

type MatchDescResponse struct {
	Status    ResponseStatus
	MatchList LVuint8
}

func (ComplexDescResponse) Cluster() uint16      { return 0x8010 }
func (n *ComplexDescResponse) Values() []zcl.Val { return []zcl.Val{&n.Status, &n.NWKI, &n.Desc} }

type ComplexDescResponse struct {
	Status ResponseStatus
	NWKI   zcl.Zu16
	Desc   zcl.Zostring
}

func (UserDescResponse) Cluster() uint16      { return 0x8011 }
func (n *UserDescResponse) Values() []zcl.Val { return []zcl.Val{&n.Status, &n.UserDescriptor} }

type UserDescResponse struct {
	Status         ResponseStatus
	UserDescriptor zcl.Zbytes
}

func (DiscoveryCacheResponse) Cluster() uint16      { return 0x8012 }
func (n *DiscoveryCacheResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type DiscoveryCacheResponse struct {
	Status ResponseStatus
}

func (SetUserDescResponse) Cluster() uint16      { return 0x8014 }
func (n *SetUserDescResponse) Values() []zcl.Val { return []zcl.Val{&n.Status, &n.NWKI} }

type SetUserDescResponse struct {
	Status ResponseStatus
	NWKI   zcl.Zu16
}

func (SystemServerDiscoveryResponse) Cluster() uint16      { return 0x8015 }
func (n *SystemServerDiscoveryResponse) Values() []zcl.Val { return []zcl.Val{&n.Status, &n.ServerMask} }

type SystemServerDiscoveryResponse struct {
	Status     ResponseStatus
	ServerMask zcl.Zu16
}

func (DiscoveryStoreResponse) Cluster() uint16      { return 0x8016 }
func (n *DiscoveryStoreResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type DiscoveryStoreResponse struct {
	Status ResponseStatus
}

func (NodeDescStoreResponse) Cluster() uint16      { return 0x8017 }
func (n *NodeDescStoreResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type NodeDescStoreResponse struct {
	Status ResponseStatus
}

func (PowerDescStoreResponse) Cluster() uint16 { return 0x8018 }
func (n *PowerDescStoreResponse) Values() []zcl.Val {
	return []zcl.Val{&n.Status, &n.IEEE, &n.Byte1, &n.Byte2}
}

type PowerDescStoreResponse struct {
	Status ResponseStatus
	IEEE   zcl.Zuid
	Byte1  zcl.Zu8
	Byte2  zcl.Zu8
}

func (ActiveEndpointStoreResponse) Cluster() uint16      { return 0x8019 }
func (n *ActiveEndpointStoreResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type ActiveEndpointStoreResponse struct {
	Status ResponseStatus
}

func (SimpleDescStoreResponse) Cluster() uint16      { return 0x801a }
func (n *SimpleDescStoreResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type SimpleDescStoreResponse struct {
	Status ResponseStatus
}

func (RemoveNodeCacheResponse) Cluster() uint16      { return 0x801b }
func (n *RemoveNodeCacheResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type RemoveNodeCacheResponse struct {
	Status ResponseStatus
}

func (FindNodeCacheResponse) Cluster() uint16      { return 0x801c }
func (n *FindNodeCacheResponse) Values() []zcl.Val { return []zcl.Val{&n.CacheNWKAddr, &n.NWK, &n.IEEE} }

type FindNodeCacheResponse struct {
	CacheNWKAddr zcl.Zuid
	NWK          zcl.Zu16
	IEEE         zcl.Zuid
}

func (ExtendedSimpleDescResponse) Cluster() uint16 { return 0x801d }
func (n *ExtendedSimpleDescResponse) Values() []zcl.Val {
	return []zcl.Val{
		&n.Status, &n.NWK, &n.Endpoint, &n.AppInputClusterCount,
		&n.AppOutputClusterCount, &n.StartIndex, &n.AppClusterList,
	}
}

type ExtendedSimpleDescResponse struct {
	Status                ResponseStatus
	NWK                   zcl.Zu16
	Endpoint              zcl.Zu8
	AppInputClusterCount  zcl.Zu8
	AppOutputClusterCount zcl.Zu8
	StartIndex            zcl.Zu8
	AppClusterList        Uint16List
}

func (ExtendedActiveEndpointResponse) Cluster() uint16 { return 0x801e }
func (n *ExtendedActiveEndpointResponse) Values() []zcl.Val {
	return []zcl.Val{&n.Status, &n.NWKI, &n.ActiveEPCount, &n.StartIndex, &n.ActiveEPList}
}

type ExtendedActiveEndpointResponse struct {
	Status        ResponseStatus
	NWKI          zcl.Zu16
	ActiveEPCount zcl.Zu8
	StartIndex    zcl.Zu8
	ActiveEPList  Uint8List
}

func (EndDeviceBindResponse) Cluster() uint16      { return 0x8020 }
func (n *EndDeviceBindResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type EndDeviceBindResponse struct {
	Status ResponseStatus
}

func (BindResponse) Cluster() uint16      { return 0x8021 }
func (n *BindResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type BindResponse struct {
	Status ResponseStatus
}

func (UnbindResponse) Cluster() uint16      { return 0x8022 }
func (n *UnbindResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type UnbindResponse struct {
	Status ResponseStatus
}

func (MgmtLqiResponse) Cluster() uint16 { return 0x8031 }
func (n *MgmtLqiResponse) Values() []zcl.Val {
	return []zcl.Val{&n.Status, &n.Entries, &n.StartIndex, &n.NeighborTableList}
}

type MgmtLqiResponse struct {
	Status            ResponseStatus
	Entries           zcl.Zu8
	StartIndex        zcl.Zu8
	NeighborTableList LVNeighbor
}

func (MgmtRtgResponse) Cluster() uint16 { return 0x8032 }
func (n *MgmtRtgResponse) Values() []zcl.Val {
	return []zcl.Val{&n.Status, &n.Entries, &n.StartIndex, &n.Routes}
}

type MgmtRtgResponse struct {
	Status     ResponseStatus
	Entries    zcl.Zu8
	StartIndex zcl.Zu8
	Routes     LVRoute
}

func (MgmtLeaveResponse) Cluster() uint16      { return 0x8034 }
func (n *MgmtLeaveResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type MgmtLeaveResponse struct {
	Status ResponseStatus
}

func (MgmtPermitJoiningResponse) Cluster() uint16      { return 0x8036 }
func (n *MgmtPermitJoiningResponse) Values() []zcl.Val { return []zcl.Val{&n.Status} }

type MgmtPermitJoiningResponse struct {
	Status ResponseStatus
}
