// Attributes and commands for scene configuration and manipulation.
package general

import (
	"neotor.se/zcl"
)

// Scenes
const ScenesID zcl.ClusterID = 5

var ScenesCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		AddSceneCommand:           func() zcl.Command { return new(AddScene) },
		ViewSceneCommand:          func() zcl.Command { return new(ViewScene) },
		RemoveSceneCommand:        func() zcl.Command { return new(RemoveScene) },
		RemoveAllScenesCommand:    func() zcl.Command { return new(RemoveAllScenes) },
		StoreSceneCommand:         func() zcl.Command { return new(StoreScene) },
		RecallSceneCommand:        func() zcl.Command { return new(RecallScene) },
		GetSceneMembershipCommand: func() zcl.Command { return new(GetSceneMembership) },
		EnhancedAddSceneCommand:   func() zcl.Command { return new(EnhancedAddScene) },
		EnhancedViewSceneCommand:  func() zcl.Command { return new(EnhancedViewScene) },
		CopySceneCommand:          func() zcl.Command { return new(CopyScene) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		AddSceneResponseCommand:           func() zcl.Command { return new(AddSceneResponse) },
		ViewSceneResponseCommand:          func() zcl.Command { return new(ViewSceneResponse) },
		RemoveSceneResponseCommand:        func() zcl.Command { return new(RemoveSceneResponse) },
		RemoveAllScenesResponseCommand:    func() zcl.Command { return new(RemoveAllScenesResponse) },
		StoreSceneResponseCommand:         func() zcl.Command { return new(StoreSceneResponse) },
		GetSceneMembershipResponseCommand: func() zcl.Command { return new(GetSceneMembershipResponse) },
		EnhancedAddSceneResponseCommand:   func() zcl.Command { return new(EnhancedAddSceneResponse) },
		EnhancedViewSceneResponseCommand:  func() zcl.Command { return new(EnhancedViewSceneResponse) },
		CopySceneResponseCommand:          func() zcl.Command { return new(CopySceneResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		SceneCountAttr:       func() zcl.Attr { return new(SceneCount) },
		CurrentSceneAttr:     func() zcl.Attr { return new(CurrentScene) },
		CurrentGroupAttr:     func() zcl.Attr { return new(CurrentGroup) },
		SceneValidAttr:       func() zcl.Attr { return new(SceneValid) },
		SceneNameSupportAttr: func() zcl.Attr { return new(SceneNameSupport) },
		LastConfiguredbyAttr: func() zcl.Attr { return new(LastConfiguredby) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// Add a scenes to the group.
type AddScene struct {
	GroupId        zcl.Zu16
	SceneId        zcl.Zu8
	TransitionTime zcl.Zu16
	SceneName      zcl.Zcstring
	// The format of each extension field set is a 16 bit field carrying the cluster ID,
	// followed by an 8 bit length field and the set of scene extension fields specified
	// in  the  relevant  cluster. The length field holds the length in octets of that
	// extension field set. Extension field set format:
	// {{clusterId1, length 1, {extension field set 1}}, {clusterId2, length 2, {extension field set 2}} ...}
	// I.e. the field would be a repeating struct with [ClusterID uint16] [OctetLength uint8] [AttrID ...uint16]
	//
	ExtensionFieldSets zcl.SceneExtensionSet
}

const AddSceneCommand zcl.CommandID = 0

func (v *AddScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
		&v.TransitionTime,
		&v.SceneName,
		&v.ExtensionFieldSets,
	}
}

func (v AddScene) ID() zcl.CommandID {
	return AddSceneCommand
}

func (v AddScene) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v AddScene) MnfCode() []byte {
	return []byte{}
}

func (v AddScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneName.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ExtensionFieldSets.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *AddScene) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ExtensionFieldSets).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Views the scenes of a group.
type ViewScene struct {
	GroupId zcl.Zu16
	SceneId zcl.Zu8
}

const ViewSceneCommand zcl.CommandID = 1

func (v *ViewScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
	}
}

func (v ViewScene) ID() zcl.CommandID {
	return ViewSceneCommand
}

func (v ViewScene) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v ViewScene) MnfCode() []byte {
	return []byte{}
}

func (v ViewScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ViewScene) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Removes a scenes of a group.
type RemoveScene struct {
	GroupId zcl.Zu16
	SceneId zcl.Zu8
}

const RemoveSceneCommand zcl.CommandID = 2

func (v *RemoveScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
	}
}

func (v RemoveScene) ID() zcl.CommandID {
	return RemoveSceneCommand
}

func (v RemoveScene) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v RemoveScene) MnfCode() []byte {
	return []byte{}
}

func (v RemoveScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RemoveScene) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Removes all scenes of a group.
type RemoveAllScenes struct {
	GroupId zcl.Zu16
}

const RemoveAllScenesCommand zcl.CommandID = 3

func (v *RemoveAllScenes) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
	}
}

func (v RemoveAllScenes) ID() zcl.CommandID {
	return RemoveAllScenesCommand
}

func (v RemoveAllScenes) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v RemoveAllScenes) MnfCode() []byte {
	return []byte{}
}

func (v RemoveAllScenes) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RemoveAllScenes) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Stores a scene of a group for a device.
type StoreScene struct {
	GroupId zcl.Zu16
	SceneId zcl.Zu8
}

const StoreSceneCommand zcl.CommandID = 4

func (v *StoreScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
	}
}

func (v StoreScene) ID() zcl.CommandID {
	return StoreSceneCommand
}

func (v StoreScene) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v StoreScene) MnfCode() []byte {
	return []byte{}
}

func (v StoreScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *StoreScene) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Recalls a scene of a group for a device.
type RecallScene struct {
	GroupId zcl.Zu16
	SceneId zcl.Zu8
}

const RecallSceneCommand zcl.CommandID = 5

func (v *RecallScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
	}
}

func (v RecallScene) ID() zcl.CommandID {
	return RecallSceneCommand
}

func (v RecallScene) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v RecallScene) MnfCode() []byte {
	return []byte{}
}

func (v RecallScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RecallScene) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Get the scenes of a group for a device.
type GetSceneMembership struct {
	GroupId zcl.Zu16
}

const GetSceneMembershipCommand zcl.CommandID = 6

func (v *GetSceneMembership) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
	}
}

func (v GetSceneMembership) ID() zcl.CommandID {
	return GetSceneMembershipCommand
}

func (v GetSceneMembership) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v GetSceneMembership) MnfCode() []byte {
	return []byte{}
}

func (v GetSceneMembership) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetSceneMembership) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Same as Add Scene, except that transition time is specified in 1/10 s
type EnhancedAddScene struct {
	GroupId            zcl.Zu16
	SceneId            zcl.Zu8
	TransitionTime     zcl.Zu16
	SceneName          zcl.Zcstring
	ExtensionFieldSets zcl.SceneExtensionSet
}

const EnhancedAddSceneCommand zcl.CommandID = 64

func (v *EnhancedAddScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
		&v.TransitionTime,
		&v.SceneName,
		&v.ExtensionFieldSets,
	}
}

func (v EnhancedAddScene) ID() zcl.CommandID {
	return EnhancedAddSceneCommand
}

func (v EnhancedAddScene) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v EnhancedAddScene) MnfCode() []byte {
	return []byte{}
}

func (v EnhancedAddScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneName.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ExtensionFieldSets.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *EnhancedAddScene) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ExtensionFieldSets).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Views the scenes of a group (returning transition time with 1/10s precision).
type EnhancedViewScene struct {
	GroupId zcl.Zu16
	SceneId zcl.Zu8
}

const EnhancedViewSceneCommand zcl.CommandID = 65

func (v *EnhancedViewScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
	}
}

func (v EnhancedViewScene) ID() zcl.CommandID {
	return EnhancedViewSceneCommand
}

func (v EnhancedViewScene) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v EnhancedViewScene) MnfCode() []byte {
	return []byte{}
}

func (v EnhancedViewScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *EnhancedViewScene) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type CopyScene struct {
	Mode        zcl.Zbmp8
	FromGroupId zcl.Zu16
	FromSceneId zcl.Zu8
	ToGroupId   zcl.Zu16
	ToSceneId   zcl.Zu8
}

const CopySceneCommand zcl.CommandID = 66

func (v *CopyScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.Mode,
		&v.FromGroupId,
		&v.FromSceneId,
		&v.ToGroupId,
		&v.ToSceneId,
	}
}

func (v CopyScene) ID() zcl.CommandID {
	return CopySceneCommand
}

func (v CopyScene) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v CopyScene) MnfCode() []byte {
	return []byte{}
}

func (v CopyScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Mode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.FromGroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.FromSceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ToGroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ToSceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *CopyScene) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Mode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FromGroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FromSceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ToGroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ToSceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Response to the add scene command.
type AddSceneResponse struct {
	Status  zcl.Status
	GroupId zcl.Zu16
	SceneId zcl.Zu8
}

const AddSceneResponseCommand zcl.CommandID = 0

func (v *AddSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
	}
}

func (v AddSceneResponse) ID() zcl.CommandID {
	return AddSceneResponseCommand
}

func (v AddSceneResponse) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v AddSceneResponse) MnfCode() []byte {
	return []byte{}
}

func (v AddSceneResponse) MarshalZcl() ([]byte, error) {
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

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *AddSceneResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Response to the view scene command.
type ViewSceneResponse struct {
	Status             zcl.Status
	GroupId            zcl.Zu16
	SceneId            zcl.Zu8
	TransitionTime     zcl.Zu16
	SceneName          zcl.Zcstring
	ExtensionFieldSets zcl.SceneExtensionSet
}

const ViewSceneResponseCommand zcl.CommandID = 1

func (v *ViewSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
		&v.TransitionTime,
		&v.SceneName,
		&v.ExtensionFieldSets,
	}
}

func (v ViewSceneResponse) ID() zcl.CommandID {
	return ViewSceneResponseCommand
}

func (v ViewSceneResponse) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v ViewSceneResponse) MnfCode() []byte {
	return []byte{}
}

func (v ViewSceneResponse) MarshalZcl() ([]byte, error) {
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

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneName.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ExtensionFieldSets.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ViewSceneResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ExtensionFieldSets).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Response to the remove scene command.
type RemoveSceneResponse struct {
	Status  zcl.Status
	GroupId zcl.Zu16
	SceneId zcl.Zu8
}

const RemoveSceneResponseCommand zcl.CommandID = 2

func (v *RemoveSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
	}
}

func (v RemoveSceneResponse) ID() zcl.CommandID {
	return RemoveSceneResponseCommand
}

func (v RemoveSceneResponse) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v RemoveSceneResponse) MnfCode() []byte {
	return []byte{}
}

func (v RemoveSceneResponse) MarshalZcl() ([]byte, error) {
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

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *RemoveSceneResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Response to the remove all scenes command.
type RemoveAllScenesResponse struct {
	Status  zcl.Status
	GroupId zcl.Zu16
}

const RemoveAllScenesResponseCommand zcl.CommandID = 3

func (v *RemoveAllScenesResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
	}
}

func (v RemoveAllScenesResponse) ID() zcl.CommandID {
	return RemoveAllScenesResponseCommand
}

func (v RemoveAllScenesResponse) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v RemoveAllScenesResponse) MnfCode() []byte {
	return []byte{}
}

func (v RemoveAllScenesResponse) MarshalZcl() ([]byte, error) {
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

func (v *RemoveAllScenesResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Response to the store scene command.
type StoreSceneResponse struct {
	Status  zcl.Status
	GroupId zcl.Zu16
	SceneId zcl.Zu8
}

const StoreSceneResponseCommand zcl.CommandID = 4

func (v *StoreSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
	}
}

func (v StoreSceneResponse) ID() zcl.CommandID {
	return StoreSceneResponseCommand
}

func (v StoreSceneResponse) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v StoreSceneResponse) MnfCode() []byte {
	return []byte{}
}

func (v StoreSceneResponse) MarshalZcl() ([]byte, error) {
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

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *StoreSceneResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Shows details about scene membership.
type GetSceneMembershipResponse struct {
	Status    zcl.Status
	Capacity  zcl.Zu8
	GroupId   zcl.Zu16
	SceneList zcl.Zset
}

const GetSceneMembershipResponseCommand zcl.CommandID = 6

func (v *GetSceneMembershipResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.Capacity,
		&v.GroupId,
		&v.SceneList,
	}
}

func (v GetSceneMembershipResponse) ID() zcl.CommandID {
	return GetSceneMembershipResponseCommand
}

func (v GetSceneMembershipResponse) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v GetSceneMembershipResponse) MnfCode() []byte {
	return []byte{}
}

func (v GetSceneMembershipResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Capacity.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.GroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.SceneList.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetSceneMembershipResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Capacity).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type EnhancedAddSceneResponse struct {
	Status  zcl.Status
	GroupId zcl.Zu16
	SceneId zcl.Zu8
}

const EnhancedAddSceneResponseCommand zcl.CommandID = 64

func (v *EnhancedAddSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
	}
}

func (v EnhancedAddSceneResponse) ID() zcl.CommandID {
	return EnhancedAddSceneResponseCommand
}

func (v EnhancedAddSceneResponse) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v EnhancedAddSceneResponse) MnfCode() []byte {
	return []byte{}
}

func (v EnhancedAddSceneResponse) MarshalZcl() ([]byte, error) {
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

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *EnhancedAddSceneResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// A scene description.
type EnhancedViewSceneResponse struct {
	Status             zcl.Status
	GroupId            zcl.Zu16
	SceneId            zcl.Zu8
	TransitionTime     zcl.Zu16
	Name               zcl.Zcstring
	ExtensionFieldSets zcl.SceneExtensionSet
}

const EnhancedViewSceneResponseCommand zcl.CommandID = 65

func (v *EnhancedViewSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
		&v.TransitionTime,
		&v.Name,
		&v.ExtensionFieldSets,
	}
}

func (v EnhancedViewSceneResponse) ID() zcl.CommandID {
	return EnhancedViewSceneResponseCommand
}

func (v EnhancedViewSceneResponse) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v EnhancedViewSceneResponse) MnfCode() []byte {
	return []byte{}
}

func (v EnhancedViewSceneResponse) MarshalZcl() ([]byte, error) {
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

	if tmp, err = v.SceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Name.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ExtensionFieldSets.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *EnhancedViewSceneResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Name).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ExtensionFieldSets).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type CopySceneResponse struct {
	Status      zcl.Status
	FromGroupId zcl.Zu16
	FromSceneId zcl.Zu8
}

const CopySceneResponseCommand zcl.CommandID = 66

func (v *CopySceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.FromGroupId,
		&v.FromSceneId,
	}
}

func (v CopySceneResponse) ID() zcl.CommandID {
	return CopySceneResponseCommand
}

func (v CopySceneResponse) Cluster() zcl.ClusterID {
	return ScenesID
}

func (v CopySceneResponse) MnfCode() []byte {
	return []byte{}
}

func (v CopySceneResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.FromGroupId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.FromSceneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *CopySceneResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FromGroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.FromSceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

const SceneCountAttr zcl.AttrID = 0

type SceneCount zcl.Zu8

func (a SceneCount) ID() zcl.AttrID         { return SceneCountAttr }
func (a SceneCount) Cluster() zcl.ClusterID { return ScenesID }
func (a *SceneCount) Value() *SceneCount    { return a }
func (a SceneCount) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *SceneCount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneCount(*nt)
	return br, err
}

func (a SceneCount) Readable() bool   { return true }
func (a SceneCount) Writable() bool   { return false }
func (a SceneCount) Reportable() bool { return false }
func (a SceneCount) SceneIndex() int  { return -1 }

func (a SceneCount) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu8(a))
}

const CurrentSceneAttr zcl.AttrID = 1

type CurrentScene zcl.Zu8

func (a CurrentScene) ID() zcl.AttrID         { return CurrentSceneAttr }
func (a CurrentScene) Cluster() zcl.ClusterID { return ScenesID }
func (a *CurrentScene) Value() *CurrentScene  { return a }
func (a CurrentScene) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *CurrentScene) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentScene(*nt)
	return br, err
}

func (a CurrentScene) Readable() bool   { return true }
func (a CurrentScene) Writable() bool   { return false }
func (a CurrentScene) Reportable() bool { return false }
func (a CurrentScene) SceneIndex() int  { return -1 }

func (a CurrentScene) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu8(a))
}

const CurrentGroupAttr zcl.AttrID = 2

type CurrentGroup zcl.Zu16

func (a CurrentGroup) ID() zcl.AttrID         { return CurrentGroupAttr }
func (a CurrentGroup) Cluster() zcl.ClusterID { return ScenesID }
func (a *CurrentGroup) Value() *CurrentGroup  { return a }
func (a CurrentGroup) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *CurrentGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentGroup(*nt)
	return br, err
}

func (a CurrentGroup) Readable() bool   { return true }
func (a CurrentGroup) Writable() bool   { return false }
func (a CurrentGroup) Reportable() bool { return false }
func (a CurrentGroup) SceneIndex() int  { return -1 }

func (a CurrentGroup) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu16(a))
}

const SceneValidAttr zcl.AttrID = 3

type SceneValid zcl.Zbool

func (a SceneValid) ID() zcl.AttrID         { return SceneValidAttr }
func (a SceneValid) Cluster() zcl.ClusterID { return ScenesID }
func (a *SceneValid) Value() *SceneValid    { return a }
func (a SceneValid) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *SceneValid) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneValid(*nt)
	return br, err
}

func (a SceneValid) Readable() bool   { return true }
func (a SceneValid) Writable() bool   { return false }
func (a SceneValid) Reportable() bool { return false }
func (a SceneValid) SceneIndex() int  { return -1 }

func (a SceneValid) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

const SceneNameSupportAttr zcl.AttrID = 4

type SceneNameSupport zcl.Zbmp8

func (a SceneNameSupport) ID() zcl.AttrID            { return SceneNameSupportAttr }
func (a SceneNameSupport) Cluster() zcl.ClusterID    { return ScenesID }
func (a *SceneNameSupport) Value() *SceneNameSupport { return a }
func (a SceneNameSupport) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *SceneNameSupport) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneNameSupport(*nt)
	return br, err
}

func (a SceneNameSupport) Readable() bool   { return true }
func (a SceneNameSupport) Writable() bool   { return false }
func (a SceneNameSupport) Reportable() bool { return false }
func (a SceneNameSupport) SceneIndex() int  { return -1 }

func (a SceneNameSupport) String() string {

	var bstr []string
	if a.IsNamesSupported() {
		bstr = append(bstr, "Names Supported")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a SceneNameSupport) IsNamesSupported() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *SceneNameSupport) SetNamesSupported(b bool) {
	*a = SceneNameSupport(zcl.BitmapSet([]byte(*a), 7, b))
}

const LastConfiguredbyAttr zcl.AttrID = 5

type LastConfiguredby zcl.Zuid

func (a LastConfiguredby) ID() zcl.AttrID            { return LastConfiguredbyAttr }
func (a LastConfiguredby) Cluster() zcl.ClusterID    { return ScenesID }
func (a *LastConfiguredby) Value() *LastConfiguredby { return a }
func (a LastConfiguredby) MarshalZcl() ([]byte, error) {
	return zcl.Zuid(a).MarshalZcl()
}
func (a *LastConfiguredby) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = LastConfiguredby(*nt)
	return br, err
}

func (a LastConfiguredby) Readable() bool   { return true }
func (a LastConfiguredby) Writable() bool   { return false }
func (a LastConfiguredby) Reportable() bool { return false }
func (a LastConfiguredby) SceneIndex() int  { return -1 }

func (a LastConfiguredby) String() string {

	return zcl.Sprintf("%s", zcl.Zuid(a))
}
