// Attributes and commands for Touchlink commissioning.
package commissioning

import (
	"neotor.se/zcl"
)

// Touchlink
const TouchlinkID zcl.ClusterID = 4096

var TouchlinkCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		GetGroupIdentifiersCommand: func() zcl.Command { return new(GetGroupIdentifiers) },
		GetEndpointListCommand:     func() zcl.Command { return new(GetEndpointList) },
		WriteMacAddressCommand:     func() zcl.Command { return new(WriteMacAddress) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		GetGroupIdentifiersResponseCommand: func() zcl.Command { return new(GetGroupIdentifiersResponse) },
		GetEndpointListResponseCommand:     func() zcl.Command { return new(GetEndpointListResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// The get group identifiers request command is used to retrieve the actual group identifiers that the endpoint is using in its multicast communication in controlling different (remote) devices.
type GetGroupIdentifiers struct {
	StartIndex zcl.Zu8
}

const GetGroupIdentifiersCommand zcl.CommandID = 65

func (v *GetGroupIdentifiers) Values() []zcl.Val {
	return []zcl.Val{
		&v.StartIndex,
	}
}

func (v GetGroupIdentifiers) ID() zcl.CommandID {
	return GetGroupIdentifiersCommand
}

func (v GetGroupIdentifiers) Cluster() zcl.ClusterID {
	return TouchlinkID
}

func (v GetGroupIdentifiers) MnfCode() []byte {
	return []byte{}
}

func (v GetGroupIdentifiers) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetGroupIdentifiers) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *GetGroupIdentifiers) StartIndexString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.StartIndex))
}

func (v *GetGroupIdentifiers) String() string {
	var str []string
	str = append(str, "StartIndex["+v.StartIndexString()+"]")
	return "GetGroupIdentifiers{" + zcl.StrJoin(str, " ") + "}"
}

// The get endpoint list request command is used to retrieve addressing information for each endpoint the device is using in its unicast communication in controlling different (remote) devices.
type GetEndpointList struct {
	StartIndex zcl.Zu8
}

const GetEndpointListCommand zcl.CommandID = 66

func (v *GetEndpointList) Values() []zcl.Val {
	return []zcl.Val{
		&v.StartIndex,
	}
}

func (v GetEndpointList) ID() zcl.CommandID {
	return GetEndpointListCommand
}

func (v GetEndpointList) Cluster() zcl.ClusterID {
	return TouchlinkID
}

func (v GetEndpointList) MnfCode() []byte {
	return []byte{}
}

func (v GetEndpointList) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetEndpointList) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *GetEndpointList) StartIndexString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.StartIndex))
}

func (v *GetEndpointList) String() string {
	var str []string
	str = append(str, "StartIndex["+v.StartIndexString()+"]")
	return "GetEndpointList{" + zcl.StrJoin(str, " ") + "}"
}

// Non standard write MAC address (DDEL).
type WriteMacAddress struct {
	MacAddress zcl.Zu64
}

const WriteMacAddressCommand zcl.CommandID = 208

func (v *WriteMacAddress) Values() []zcl.Val {
	return []zcl.Val{
		&v.MacAddress,
	}
}

func (v WriteMacAddress) ID() zcl.CommandID {
	return WriteMacAddressCommand
}

func (v WriteMacAddress) Cluster() zcl.ClusterID {
	return TouchlinkID
}

func (v WriteMacAddress) MnfCode() []byte {
	return []byte{}
}

func (v WriteMacAddress) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.MacAddress.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *WriteMacAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.MacAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *WriteMacAddress) MacAddressString() string {
	return zcl.Sprintf("0x%X", zcl.Zu64(v.MacAddress))
}

func (v *WriteMacAddress) String() string {
	var str []string
	str = append(str, "MacAddress["+v.MacAddressString()+"]")
	return "WriteMacAddress{" + zcl.StrJoin(str, " ") + "}"
}

// The get group identifiers response command allows a remote device to respond to the get group identifiers request command.
type GetGroupIdentifiersResponse struct {
	Total          zcl.Zu8
	StartIndex     zcl.Zu8
	Count          zcl.Zu8
	FirstGroupId   zcl.Zu16
	FirstGroupType zcl.Zu8
}

const GetGroupIdentifiersResponseCommand zcl.CommandID = 65

func (v *GetGroupIdentifiersResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Total,
		&v.StartIndex,
		&v.Count,
		&v.FirstGroupId,
		&v.FirstGroupType,
	}
}

func (v GetGroupIdentifiersResponse) ID() zcl.CommandID {
	return GetGroupIdentifiersResponseCommand
}

func (v GetGroupIdentifiersResponse) Cluster() zcl.ClusterID {
	return TouchlinkID
}

func (v GetGroupIdentifiersResponse) MnfCode() []byte {
	return []byte{}
}

func (v GetGroupIdentifiersResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Total.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Count.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.FirstGroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.FirstGroupType.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetGroupIdentifiersResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Total).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Count).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FirstGroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FirstGroupType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *GetGroupIdentifiersResponse) TotalString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.Total))
}
func (v *GetGroupIdentifiersResponse) StartIndexString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.StartIndex))
}
func (v *GetGroupIdentifiersResponse) CountString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.Count))
}
func (v *GetGroupIdentifiersResponse) FirstGroupIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.FirstGroupId))
}
func (v *GetGroupIdentifiersResponse) FirstGroupTypeString() string {
	return zcl.Sprintf("0x%X", zcl.Zu8(v.FirstGroupType))
}

func (v *GetGroupIdentifiersResponse) String() string {
	var str []string
	str = append(str, "Total["+v.TotalString()+"]")
	str = append(str, "StartIndex["+v.StartIndexString()+"]")
	str = append(str, "Count["+v.CountString()+"]")
	str = append(str, "FirstGroupId["+v.FirstGroupIdString()+"]")
	str = append(str, "FirstGroupType["+v.FirstGroupTypeString()+"]")
	return "GetGroupIdentifiersResponse{" + zcl.StrJoin(str, " ") + "}"
}

// The get group identifiers response command allows a remote device to respond to the get group identifiers request command.
type GetEndpointListResponse struct {
	Total      zcl.Zu8
	StartIndex zcl.Zu8
	Count      zcl.Zu8
	NwkAddress zcl.Zu16
	Endpoint   zcl.Zu8
	ProfileId  zcl.Zu16
	DeviceId   zcl.Zu16
	Version    zcl.Zu8
}

const GetEndpointListResponseCommand zcl.CommandID = 66

func (v *GetEndpointListResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Total,
		&v.StartIndex,
		&v.Count,
		&v.NwkAddress,
		&v.Endpoint,
		&v.ProfileId,
		&v.DeviceId,
		&v.Version,
	}
}

func (v GetEndpointListResponse) ID() zcl.CommandID {
	return GetEndpointListResponseCommand
}

func (v GetEndpointListResponse) Cluster() zcl.ClusterID {
	return TouchlinkID
}

func (v GetEndpointListResponse) MnfCode() []byte {
	return []byte{}
}

func (v GetEndpointListResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Total.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StartIndex.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Count.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NwkAddress.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Endpoint.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ProfileId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.DeviceId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Version.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetEndpointListResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Total).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StartIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Count).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NwkAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Endpoint).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ProfileId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.DeviceId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Version).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v *GetEndpointListResponse) TotalString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.Total))
}
func (v *GetEndpointListResponse) StartIndexString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.StartIndex))
}
func (v *GetEndpointListResponse) CountString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.Count))
}
func (v *GetEndpointListResponse) NwkAddressString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.NwkAddress))
}
func (v *GetEndpointListResponse) EndpointString() string {
	return zcl.Sprintf("0x%X", zcl.Zu8(v.Endpoint))
}
func (v *GetEndpointListResponse) ProfileIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.ProfileId))
}
func (v *GetEndpointListResponse) DeviceIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.DeviceId))
}
func (v *GetEndpointListResponse) VersionString() string {
	return zcl.Sprintf("0x%X", zcl.Zu8(v.Version))
}

func (v *GetEndpointListResponse) String() string {
	var str []string
	str = append(str, "Total["+v.TotalString()+"]")
	str = append(str, "StartIndex["+v.StartIndexString()+"]")
	str = append(str, "Count["+v.CountString()+"]")
	str = append(str, "NwkAddress["+v.NwkAddressString()+"]")
	str = append(str, "Endpoint["+v.EndpointString()+"]")
	str = append(str, "ProfileId["+v.ProfileIdString()+"]")
	str = append(str, "DeviceId["+v.DeviceIdString()+"]")
	str = append(str, "Version["+v.VersionString()+"]")
	return "GetEndpointListResponse{" + zcl.StrJoin(str, " ") + "}"
}
