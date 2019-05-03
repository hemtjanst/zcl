// Attributes and commands for group configuration and manipulation.
package general

import (
	"neotor.se/zcl"
)

// Groups
const GroupsID zcl.ClusterID = 4

var GroupsCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		AddGroupCommand:              func() zcl.Command { return new(AddGroup) },
		ViewGroupCommand:             func() zcl.Command { return new(ViewGroup) },
		GetGroupMembershipCommand:    func() zcl.Command { return new(GetGroupMembership) },
		RemoveGroupCommand:           func() zcl.Command { return new(RemoveGroup) },
		RemoveAllGroupsCommand:       func() zcl.Command { return new(RemoveAllGroups) },
		AddGroupIfIdentifyingCommand: func() zcl.Command { return new(AddGroupIfIdentifying) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		AddGroupResponseCommand:           func() zcl.Command { return new(AddGroupResponse) },
		ViewGroupResponseCommand:          func() zcl.Command { return new(ViewGroupResponse) },
		GetGroupMembershipResponseCommand: func() zcl.Command { return new(GetGroupMembershipResponse) },
		RemoveGroupResponseCommand:        func() zcl.Command { return new(RemoveGroupResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		GroupNameSupportAttr: func() zcl.Attr { return new(GroupNameSupport) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// Add a group to the device.
type AddGroup struct {
	GroupId   zcl.Zu16
	GroupName zcl.Zcstring
}

const AddGroupCommand zcl.CommandID = 0

func (v *AddGroup) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.GroupName,
	}
}

func (v AddGroup) ID() zcl.CommandID {
	return AddGroupCommand
}

func (v AddGroup) Cluster() zcl.ClusterID {
	return GroupsID
}

func (v AddGroup) MnfCode() []byte {
	return []byte{}
}

func (v AddGroup) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupName.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *AddGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v AddGroup) GroupIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.GroupId))
}
func (v AddGroup) GroupNameString() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v.GroupName))
}

func (v AddGroup) String() string {
	var str []string
	str = append(str, "GroupId["+v.GroupIdString()+"]")
	str = append(str, "GroupName["+v.GroupNameString()+"]")
	return "AddGroup{" + zcl.StrJoin(str, " ") + "}"
}

func (AddGroup) Name() string { return "Add group" }

// Get the name of a group.
type ViewGroup struct {
	GroupId zcl.Zu16
}

const ViewGroupCommand zcl.CommandID = 1

func (v *ViewGroup) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
	}
}

func (v ViewGroup) ID() zcl.CommandID {
	return ViewGroupCommand
}

func (v ViewGroup) Cluster() zcl.ClusterID {
	return GroupsID
}

func (v ViewGroup) MnfCode() []byte {
	return []byte{}
}

func (v ViewGroup) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ViewGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v ViewGroup) GroupIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.GroupId))
}

func (v ViewGroup) String() string {
	var str []string
	str = append(str, "GroupId["+v.GroupIdString()+"]")
	return "ViewGroup{" + zcl.StrJoin(str, " ") + "}"
}

func (ViewGroup) Name() string { return "View group" }

// Get the group membership of the device. Send an empty group list to request all group memberships
type GetGroupMembership struct {
	GroupList zcl.Zset
}

const GetGroupMembershipCommand zcl.CommandID = 2

func (v *GetGroupMembership) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupList,
	}
}

func (v GetGroupMembership) ID() zcl.CommandID {
	return GetGroupMembershipCommand
}

func (v GetGroupMembership) Cluster() zcl.ClusterID {
	return GroupsID
}

func (v GetGroupMembership) MnfCode() []byte {
	return []byte{}
}

func (v GetGroupMembership) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupList.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetGroupMembership) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v GetGroupMembership) GroupListString() string {
	return zcl.Sprintf("0x%X", zcl.Zset(v.GroupList))
}

func (v GetGroupMembership) String() string {
	var str []string
	str = append(str, "GroupList["+v.GroupListString()+"]")
	return "GetGroupMembership{" + zcl.StrJoin(str, " ") + "}"
}

func (GetGroupMembership) Name() string { return "Get group membership" }

// Remove a group from the device.
type RemoveGroup struct {
	GroupId zcl.Zu16
}

const RemoveGroupCommand zcl.CommandID = 3

func (v *RemoveGroup) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
	}
}

func (v RemoveGroup) ID() zcl.CommandID {
	return RemoveGroupCommand
}

func (v RemoveGroup) Cluster() zcl.ClusterID {
	return GroupsID
}

func (v RemoveGroup) MnfCode() []byte {
	return []byte{}
}

func (v RemoveGroup) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RemoveGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v RemoveGroup) GroupIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.GroupId))
}

func (v RemoveGroup) String() string {
	var str []string
	str = append(str, "GroupId["+v.GroupIdString()+"]")
	return "RemoveGroup{" + zcl.StrJoin(str, " ") + "}"
}

func (RemoveGroup) Name() string { return "Remove group" }

// Remove all group from the device.
type RemoveAllGroups struct {
}

const RemoveAllGroupsCommand zcl.CommandID = 4

func (v *RemoveAllGroups) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v RemoveAllGroups) ID() zcl.CommandID {
	return RemoveAllGroupsCommand
}

func (v RemoveAllGroups) Cluster() zcl.ClusterID {
	return GroupsID
}

func (v RemoveAllGroups) MnfCode() []byte {
	return []byte{}
}

func (v RemoveAllGroups) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *RemoveAllGroups) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

func (v RemoveAllGroups) String() string {
	var str []string
	return "RemoveAllGroups{" + zcl.StrJoin(str, " ") + "}"
}

func (RemoveAllGroups) Name() string { return "Remove all groups" }

// Add a group to the device if the device is currently identifying itself (using the identify cluster)
type AddGroupIfIdentifying struct {
	GroupId   zcl.Zu16
	GroupName zcl.Zcstring
}

const AddGroupIfIdentifyingCommand zcl.CommandID = 5

func (v *AddGroupIfIdentifying) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.GroupName,
	}
}

func (v AddGroupIfIdentifying) ID() zcl.CommandID {
	return AddGroupIfIdentifyingCommand
}

func (v AddGroupIfIdentifying) Cluster() zcl.ClusterID {
	return GroupsID
}

func (v AddGroupIfIdentifying) MnfCode() []byte {
	return []byte{}
}

func (v AddGroupIfIdentifying) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupName.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *AddGroupIfIdentifying) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v AddGroupIfIdentifying) GroupIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.GroupId))
}
func (v AddGroupIfIdentifying) GroupNameString() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v.GroupName))
}

func (v AddGroupIfIdentifying) String() string {
	var str []string
	str = append(str, "GroupId["+v.GroupIdString()+"]")
	str = append(str, "GroupName["+v.GroupNameString()+"]")
	return "AddGroupIfIdentifying{" + zcl.StrJoin(str, " ") + "}"
}

func (AddGroupIfIdentifying) Name() string { return "Add group if identifying" }

// The Response to the add group request.
type AddGroupResponse struct {
	Status  zcl.Status
	GroupId zcl.Zu16
}

const AddGroupResponseCommand zcl.CommandID = 0

func (v *AddGroupResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
	}
}

func (v AddGroupResponse) ID() zcl.CommandID {
	return AddGroupResponseCommand
}

func (v AddGroupResponse) Cluster() zcl.ClusterID {
	return GroupsID
}

func (v AddGroupResponse) MnfCode() []byte {
	return []byte{}
}

func (v AddGroupResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *AddGroupResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v AddGroupResponse) StatusString() string {
	return zcl.Sprintf("%v", zcl.Status(v.Status))
}
func (v AddGroupResponse) GroupIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.GroupId))
}

func (v AddGroupResponse) String() string {
	var str []string
	str = append(str, "Status["+v.StatusString()+"]")
	str = append(str, "GroupId["+v.GroupIdString()+"]")
	return "AddGroupResponse{" + zcl.StrJoin(str, " ") + "}"
}

func (AddGroupResponse) Name() string { return "Add group response" }

// The Response to the view group request.
type ViewGroupResponse struct {
	Status    zcl.Status
	GroupId   zcl.Zu16
	GroupName zcl.Zcstring
}

const ViewGroupResponseCommand zcl.CommandID = 1

func (v *ViewGroupResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.GroupName,
	}
}

func (v ViewGroupResponse) ID() zcl.CommandID {
	return ViewGroupResponseCommand
}

func (v ViewGroupResponse) Cluster() zcl.ClusterID {
	return GroupsID
}

func (v ViewGroupResponse) MnfCode() []byte {
	return []byte{}
}

func (v ViewGroupResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupName.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ViewGroupResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v ViewGroupResponse) StatusString() string {
	return zcl.Sprintf("%v", zcl.Status(v.Status))
}
func (v ViewGroupResponse) GroupIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.GroupId))
}
func (v ViewGroupResponse) GroupNameString() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v.GroupName))
}

func (v ViewGroupResponse) String() string {
	var str []string
	str = append(str, "Status["+v.StatusString()+"]")
	str = append(str, "GroupId["+v.GroupIdString()+"]")
	str = append(str, "GroupName["+v.GroupNameString()+"]")
	return "ViewGroupResponse{" + zcl.StrJoin(str, " ") + "}"
}

func (ViewGroupResponse) Name() string { return "View group response" }

// The Response to the get group membership request.
type GetGroupMembershipResponse struct {
	// The remaining number of groups that can be added.
	// If set to 0xFE, at least one more group can be added (exact number unknown)
	// If set to 0xFF, it's unknown if any more groups can be added
	Capacity  zcl.Zu8
	GroupList zcl.Zset
}

const GetGroupMembershipResponseCommand zcl.CommandID = 2

func (v *GetGroupMembershipResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Capacity,
		&v.GroupList,
	}
}

func (v GetGroupMembershipResponse) ID() zcl.CommandID {
	return GetGroupMembershipResponseCommand
}

func (v GetGroupMembershipResponse) Cluster() zcl.ClusterID {
	return GroupsID
}

func (v GetGroupMembershipResponse) MnfCode() []byte {
	return []byte{}
}

func (v GetGroupMembershipResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Capacity.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupList.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetGroupMembershipResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Capacity).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v GetGroupMembershipResponse) CapacityString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.Capacity))
}
func (v GetGroupMembershipResponse) GroupListString() string {
	return zcl.Sprintf("0x%X", zcl.Zset(v.GroupList))
}

func (v GetGroupMembershipResponse) String() string {
	var str []string
	str = append(str, "Capacity["+v.CapacityString()+"]")
	str = append(str, "GroupList["+v.GroupListString()+"]")
	return "GetGroupMembershipResponse{" + zcl.StrJoin(str, " ") + "}"
}

func (GetGroupMembershipResponse) Name() string { return "Get group membership response" }

// The Response to the remove group request.
type RemoveGroupResponse struct {
	Status  zcl.Status
	GroupId zcl.Zu16
}

const RemoveGroupResponseCommand zcl.CommandID = 3

func (v *RemoveGroupResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
	}
}

func (v RemoveGroupResponse) ID() zcl.CommandID {
	return RemoveGroupResponseCommand
}

func (v RemoveGroupResponse) Cluster() zcl.ClusterID {
	return GroupsID
}

func (v RemoveGroupResponse) MnfCode() []byte {
	return []byte{}
}

func (v RemoveGroupResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RemoveGroupResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v RemoveGroupResponse) StatusString() string {
	return zcl.Sprintf("%v", zcl.Status(v.Status))
}
func (v RemoveGroupResponse) GroupIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.GroupId))
}

func (v RemoveGroupResponse) String() string {
	var str []string
	str = append(str, "Status["+v.StatusString()+"]")
	str = append(str, "GroupId["+v.GroupIdString()+"]")
	return "RemoveGroupResponse{" + zcl.StrJoin(str, " ") + "}"
}

func (RemoveGroupResponse) Name() string { return "Remove group response" }

// GroupNameSupport is an autogenerated attribute in the Groups cluster
type GroupNameSupport zcl.Zbmp8

const GroupNameSupportAttr zcl.AttrID = 0

func (a GroupNameSupport) ID() zcl.AttrID            { return GroupNameSupportAttr }
func (a GroupNameSupport) Cluster() zcl.ClusterID    { return GroupsID }
func (a *GroupNameSupport) Value() *GroupNameSupport { return a }
func (a GroupNameSupport) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *GroupNameSupport) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupNameSupport(*nt)
	return br, err
}
func (GroupNameSupport) Name() string     { return "Group Name Support" }
func (GroupNameSupport) Readable() bool   { return true }
func (GroupNameSupport) Writable() bool   { return false }
func (GroupNameSupport) Reportable() bool { return false }
func (GroupNameSupport) SceneIndex() int  { return -1 }

func (a GroupNameSupport) String() string {
	var bstr []string
	if a.IsNamesSupported() {
		bstr = append(bstr, "Names Supported")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a GroupNameSupport) IsNamesSupported() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *GroupNameSupport) SetNamesSupported(b bool) {
	*a = GroupNameSupport(zcl.BitmapSet([]byte(*a), 7, b))
}
