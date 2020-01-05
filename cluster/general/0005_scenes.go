package general

import "hemtjan.st/zcl"

// Scenes
const ScenesID zcl.ClusterID = 5

var ScenesCluster = zcl.Cluster{
	Name: "Scenes",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		AddSceneCommand:                 func() zcl.Command { return new(AddScene) },
		ViewSceneCommand:                func() zcl.Command { return new(ViewScene) },
		RemoveSceneCommand:              func() zcl.Command { return new(RemoveScene) },
		RemoveAllScenesCommand:          func() zcl.Command { return new(RemoveAllScenes) },
		StoreSceneCommand:               func() zcl.Command { return new(StoreScene) },
		RecallSceneCommand:              func() zcl.Command { return new(RecallScene) },
		GetSceneMembershipCommand:       func() zcl.Command { return new(GetSceneMembership) },
		EnhancedAddSceneCommand:         func() zcl.Command { return new(EnhancedAddScene) },
		EnhancedViewSceneCommand:        func() zcl.Command { return new(EnhancedViewScene) },
		CopySceneCommand:                func() zcl.Command { return new(CopyScene) },
		IkeaRemotePressCommand:          func() zcl.Command { return new(IkeaRemotePress) },
		IkeaRemoteLongpressStartCommand: func() zcl.Command { return new(IkeaRemoteLongpressStart) },
		IkeaRemoteLongpressStopCommand:  func() zcl.Command { return new(IkeaRemoteLongpressStop) },
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
		SceneCountAttr:            func() zcl.Attr { return new(SceneCount) },
		CurrentSceneAttr:          func() zcl.Attr { return new(CurrentScene) },
		CurrentGroupAttr:          func() zcl.Attr { return new(CurrentGroup) },
		SceneValidAttr:            func() zcl.Attr { return new(SceneValid) },
		SceneNameSupportAttr:      func() zcl.Attr { return new(SceneNameSupport) },
		SceneLastConfiguredByAttr: func() zcl.Attr { return new(SceneLastConfiguredBy) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// AddScene Add a scenes to the group.
type AddScene struct {
	GroupId           GroupId
	SceneId           SceneId
	TransitionTimeSec TransitionTimeSec
	SceneName         SceneName
	// SceneExtensions The format of each extension field set is a 16 bit field carrying the cluster ID,
	// followed by an 8 bit length field and the set of scene extension fields specified
	// in  the  relevant  cluster. The length field holds the length in octets of that
	// extension field set. Extension field set format:
	// {{clusterId1, length 1, {extension field set 1}}, {clusterId2, length 2, {extension field set 2}} ...}
	// I.e. the field would be a repeating struct with [ClusterID uint16] [OctetLength uint8] [AttrID ...uint16]
	SceneExtensions SceneExtensions
}

// AddSceneCommand is the Command ID of AddScene
const AddSceneCommand CommandID = 0x0000

// Values returns all values of AddScene
func (v *AddScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
		&v.TransitionTimeSec,
		&v.SceneName,
		&v.SceneExtensions,
	}
}

// Arguments returns all values of AddScene
func (v *AddScene) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
		{Name: "TransitionTimeSec", Argument: &v.TransitionTimeSec},
		{Name: "SceneName", Argument: &v.SceneName},
		{Name: "SceneExtensions", Argument: &v.SceneExtensions},
	}
}

// Name of the command
func (AddScene) Name() string { return "Add scene" }

// ID of the command
func (AddScene) ID() CommandID { return AddSceneCommand }

// Required
func (AddScene) Required() bool { return true }

// Cluster ID of the command
func (AddScene) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (AddScene) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (AddScene) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

// MarshalZcl returns the wire format representation of AddScene
func (v AddScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TransitionTimeSec.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneName.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneExtensions.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the AddScene struct
func (v *AddScene) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTimeSec).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneExtensions).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v AddScene) String() string {
	return zcl.Sprintf(
		"AddScene{"+zcl.StrJoin([]string{
			"GroupId(%v)",
			"SceneId(%v)",
			"TransitionTimeSec(%v)",
			"SceneName(%v)",
			"SceneExtensions(%v)",
		}, " ")+"}",
		v.GroupId,
		v.SceneId,
		v.TransitionTimeSec,
		v.SceneName,
		v.SceneExtensions,
	)
}

// ViewScene Views the scenes of a group.
type ViewScene struct {
	GroupId GroupId
	SceneId SceneId
}

// ViewSceneCommand is the Command ID of ViewScene
const ViewSceneCommand CommandID = 0x0001

// Values returns all values of ViewScene
func (v *ViewScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
	}
}

// Arguments returns all values of ViewScene
func (v *ViewScene) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
	}
}

// Name of the command
func (ViewScene) Name() string { return "View scene" }

// ID of the command
func (ViewScene) ID() CommandID { return ViewSceneCommand }

// Required
func (ViewScene) Required() bool { return true }

// Cluster ID of the command
func (ViewScene) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (ViewScene) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (ViewScene) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

// MarshalZcl returns the wire format representation of ViewScene
func (v ViewScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ViewScene struct
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

// String returns a log-friendly string representation of the struct
func (v ViewScene) String() string {
	return zcl.Sprintf(
		"ViewScene{"+zcl.StrJoin([]string{
			"GroupId(%v)",
			"SceneId(%v)",
		}, " ")+"}",
		v.GroupId,
		v.SceneId,
	)
}

// RemoveScene Removes a scenes of a group.
type RemoveScene struct {
	GroupId GroupId
	SceneId SceneId
}

// RemoveSceneCommand is the Command ID of RemoveScene
const RemoveSceneCommand CommandID = 0x0002

// Values returns all values of RemoveScene
func (v *RemoveScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
	}
}

// Arguments returns all values of RemoveScene
func (v *RemoveScene) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
	}
}

// Name of the command
func (RemoveScene) Name() string { return "Remove scene" }

// ID of the command
func (RemoveScene) ID() CommandID { return RemoveSceneCommand }

// Required
func (RemoveScene) Required() bool { return true }

// Cluster ID of the command
func (RemoveScene) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (RemoveScene) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (RemoveScene) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

// MarshalZcl returns the wire format representation of RemoveScene
func (v RemoveScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RemoveScene struct
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

// String returns a log-friendly string representation of the struct
func (v RemoveScene) String() string {
	return zcl.Sprintf(
		"RemoveScene{"+zcl.StrJoin([]string{
			"GroupId(%v)",
			"SceneId(%v)",
		}, " ")+"}",
		v.GroupId,
		v.SceneId,
	)
}

// RemoveAllScenes Removes all scenes of a group.
type RemoveAllScenes struct {
	GroupId GroupId
}

// RemoveAllScenesCommand is the Command ID of RemoveAllScenes
const RemoveAllScenesCommand CommandID = 0x0003

// Values returns all values of RemoveAllScenes
func (v *RemoveAllScenes) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
	}
}

// Arguments returns all values of RemoveAllScenes
func (v *RemoveAllScenes) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
	}
}

// Name of the command
func (RemoveAllScenes) Name() string { return "Remove all scenes" }

// ID of the command
func (RemoveAllScenes) ID() CommandID { return RemoveAllScenesCommand }

// Required
func (RemoveAllScenes) Required() bool { return true }

// Cluster ID of the command
func (RemoveAllScenes) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (RemoveAllScenes) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (RemoveAllScenes) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

// MarshalZcl returns the wire format representation of RemoveAllScenes
func (v RemoveAllScenes) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RemoveAllScenes struct
func (v *RemoveAllScenes) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v RemoveAllScenes) String() string {
	return zcl.Sprintf(
		"RemoveAllScenes{"+zcl.StrJoin([]string{
			"GroupId(%v)",
		}, " ")+"}",
		v.GroupId,
	)
}

// StoreScene Stores a scene of a group for a device.
type StoreScene struct {
	GroupId GroupId
	SceneId SceneId
}

// StoreSceneCommand is the Command ID of StoreScene
const StoreSceneCommand CommandID = 0x0004

// Values returns all values of StoreScene
func (v *StoreScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
	}
}

// Arguments returns all values of StoreScene
func (v *StoreScene) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
	}
}

// Name of the command
func (StoreScene) Name() string { return "Store scene" }

// ID of the command
func (StoreScene) ID() CommandID { return StoreSceneCommand }

// Required
func (StoreScene) Required() bool { return true }

// Cluster ID of the command
func (StoreScene) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (StoreScene) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (StoreScene) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

// MarshalZcl returns the wire format representation of StoreScene
func (v StoreScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the StoreScene struct
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

// String returns a log-friendly string representation of the struct
func (v StoreScene) String() string {
	return zcl.Sprintf(
		"StoreScene{"+zcl.StrJoin([]string{
			"GroupId(%v)",
			"SceneId(%v)",
		}, " ")+"}",
		v.GroupId,
		v.SceneId,
	)
}

// RecallScene Recalls a scene of a group for a device.
type RecallScene struct {
	GroupId GroupId
	SceneId SceneId
}

// RecallSceneCommand is the Command ID of RecallScene
const RecallSceneCommand CommandID = 0x0005

// Values returns all values of RecallScene
func (v *RecallScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
	}
}

// Arguments returns all values of RecallScene
func (v *RecallScene) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
	}
}

// Name of the command
func (RecallScene) Name() string { return "Recall scene" }

// ID of the command
func (RecallScene) ID() CommandID { return RecallSceneCommand }

// Required
func (RecallScene) Required() bool { return true }

// Cluster ID of the command
func (RecallScene) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (RecallScene) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (RecallScene) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

// MarshalZcl returns the wire format representation of RecallScene
func (v RecallScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RecallScene struct
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

// String returns a log-friendly string representation of the struct
func (v RecallScene) String() string {
	return zcl.Sprintf(
		"RecallScene{"+zcl.StrJoin([]string{
			"GroupId(%v)",
			"SceneId(%v)",
		}, " ")+"}",
		v.GroupId,
		v.SceneId,
	)
}

// GetSceneMembership Get the scenes of a group for a device.
type GetSceneMembership struct {
	GroupId GroupId
}

// GetSceneMembershipCommand is the Command ID of GetSceneMembership
const GetSceneMembershipCommand CommandID = 0x0006

// Values returns all values of GetSceneMembership
func (v *GetSceneMembership) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
	}
}

// Arguments returns all values of GetSceneMembership
func (v *GetSceneMembership) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
	}
}

// Name of the command
func (GetSceneMembership) Name() string { return "Get scene membership" }

// ID of the command
func (GetSceneMembership) ID() CommandID { return GetSceneMembershipCommand }

// Required
func (GetSceneMembership) Required() bool { return true }

// Cluster ID of the command
func (GetSceneMembership) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (GetSceneMembership) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (GetSceneMembership) MarshalJSON() ([]byte, error) { return []byte("6"), nil }

// MarshalZcl returns the wire format representation of GetSceneMembership
func (v GetSceneMembership) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetSceneMembership struct
func (v *GetSceneMembership) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.GroupId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetSceneMembership) String() string {
	return zcl.Sprintf(
		"GetSceneMembership{"+zcl.StrJoin([]string{
			"GroupId(%v)",
		}, " ")+"}",
		v.GroupId,
	)
}

// EnhancedAddScene Same as Add Scene, except that transition time is specified in 1/10 s
type EnhancedAddScene struct {
	GroupId        GroupId
	SceneId        SceneId
	TransitionTime TransitionTime
	SceneName      SceneName
	// SceneExtensions The format of each extension field set is a 16 bit field carrying the cluster ID,
	// followed by an 8 bit length field and the set of scene extension fields specified
	// in  the  relevant  cluster. The length field holds the length in octets of that
	// extension field set. Extension field set format:
	// {{clusterId1, length 1, {extension field set 1}}, {clusterId2, length 2, {extension field set 2}} ...}
	// I.e. the field would be a repeating struct with [ClusterID uint16] [OctetLength uint8] [AttrID ...uint16]
	SceneExtensions SceneExtensions
}

// EnhancedAddSceneCommand is the Command ID of EnhancedAddScene
const EnhancedAddSceneCommand CommandID = 0x0040

// Values returns all values of EnhancedAddScene
func (v *EnhancedAddScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
		&v.TransitionTime,
		&v.SceneName,
		&v.SceneExtensions,
	}
}

// Arguments returns all values of EnhancedAddScene
func (v *EnhancedAddScene) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
		{Name: "SceneName", Argument: &v.SceneName},
		{Name: "SceneExtensions", Argument: &v.SceneExtensions},
	}
}

// Name of the command
func (EnhancedAddScene) Name() string { return "Enhanced Add Scene" }

// ID of the command
func (EnhancedAddScene) ID() CommandID { return EnhancedAddSceneCommand }

// Required
func (EnhancedAddScene) Required() bool { return true }

// Cluster ID of the command
func (EnhancedAddScene) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (EnhancedAddScene) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (EnhancedAddScene) MarshalJSON() ([]byte, error) { return []byte("64"), nil }

// MarshalZcl returns the wire format representation of EnhancedAddScene
func (v EnhancedAddScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneName.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneExtensions.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the EnhancedAddScene struct
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

	if b, err = (&v.SceneExtensions).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v EnhancedAddScene) String() string {
	return zcl.Sprintf(
		"EnhancedAddScene{"+zcl.StrJoin([]string{
			"GroupId(%v)",
			"SceneId(%v)",
			"TransitionTime(%v)",
			"SceneName(%v)",
			"SceneExtensions(%v)",
		}, " ")+"}",
		v.GroupId,
		v.SceneId,
		v.TransitionTime,
		v.SceneName,
		v.SceneExtensions,
	)
}

// EnhancedViewScene Views the scenes of a group (returning transition time with 1/10s precision).
type EnhancedViewScene struct {
	GroupId GroupId
	SceneId SceneId
}

// EnhancedViewSceneCommand is the Command ID of EnhancedViewScene
const EnhancedViewSceneCommand CommandID = 0x0041

// Values returns all values of EnhancedViewScene
func (v *EnhancedViewScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.GroupId,
		&v.SceneId,
	}
}

// Arguments returns all values of EnhancedViewScene
func (v *EnhancedViewScene) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
	}
}

// Name of the command
func (EnhancedViewScene) Name() string { return "Enhanced View Scene" }

// ID of the command
func (EnhancedViewScene) ID() CommandID { return EnhancedViewSceneCommand }

// Required
func (EnhancedViewScene) Required() bool { return false }

// Cluster ID of the command
func (EnhancedViewScene) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (EnhancedViewScene) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (EnhancedViewScene) MarshalJSON() ([]byte, error) { return []byte("65"), nil }

// MarshalZcl returns the wire format representation of EnhancedViewScene
func (v EnhancedViewScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the EnhancedViewScene struct
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

// String returns a log-friendly string representation of the struct
func (v EnhancedViewScene) String() string {
	return zcl.Sprintf(
		"EnhancedViewScene{"+zcl.StrJoin([]string{
			"GroupId(%v)",
			"SceneId(%v)",
		}, " ")+"}",
		v.GroupId,
		v.SceneId,
	)
}

type CopyScene struct {
	SceneCopyMode SceneCopyMode
	FromGroupId   GroupId
	FromSceneId   SceneId
	ToGroupId     GroupId
	ToSceneId     SceneId
}

// CopySceneCommand is the Command ID of CopyScene
const CopySceneCommand CommandID = 0x0042

// Values returns all values of CopyScene
func (v *CopyScene) Values() []zcl.Val {
	return []zcl.Val{
		&v.SceneCopyMode,
		&v.FromGroupId,
		&v.FromSceneId,
		&v.ToGroupId,
		&v.ToSceneId,
	}
}

// Arguments returns all values of CopyScene
func (v *CopyScene) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "SceneCopyMode", Argument: &v.SceneCopyMode},
		{Name: "FromGroupId", Argument: &v.FromGroupId},
		{Name: "FromSceneId", Argument: &v.FromSceneId},
		{Name: "ToGroupId", Argument: &v.ToGroupId},
		{Name: "ToSceneId", Argument: &v.ToSceneId},
	}
}

// Name of the command
func (CopyScene) Name() string { return "Copy Scene" }

// ID of the command
func (CopyScene) ID() CommandID { return CopySceneCommand }

// Required
func (CopyScene) Required() bool { return false }

// Cluster ID of the command
func (CopyScene) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (CopyScene) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (CopyScene) MarshalJSON() ([]byte, error) { return []byte("66"), nil }

// MarshalZcl returns the wire format representation of CopyScene
func (v CopyScene) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.SceneCopyMode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FromGroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FromSceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ToGroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ToSceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the CopyScene struct
func (v *CopyScene) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.SceneCopyMode).UnmarshalZcl(b); err != nil {
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

// String returns a log-friendly string representation of the struct
func (v CopyScene) String() string {
	return zcl.Sprintf(
		"CopyScene{"+zcl.StrJoin([]string{
			"SceneCopyMode(%v)",
			"FromGroupId(%v)",
			"FromSceneId(%v)",
			"ToGroupId(%v)",
			"ToSceneId(%v)",
		}, " ")+"}",
		v.SceneCopyMode,
		v.FromGroupId,
		v.FromSceneId,
		v.ToGroupId,
		v.ToSceneId,
	)
}

// IkeaRemotePress Custom left/right action of Ikea remote
type IkeaRemotePress struct {
	IkeaRemoteDirection IkeaRemoteDirection
}

// IkeaRemotePressCommand is the Command ID of IkeaRemotePress
const IkeaRemotePressCommand CommandID = 0x0007

// Values returns all values of IkeaRemotePress
func (v *IkeaRemotePress) Values() []zcl.Val {
	return []zcl.Val{
		&v.IkeaRemoteDirection,
	}
}

// Arguments returns all values of IkeaRemotePress
func (v *IkeaRemotePress) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "IkeaRemoteDirection", Argument: &v.IkeaRemoteDirection},
	}
}

// Name of the command
func (IkeaRemotePress) Name() string { return "Ikea Remote Press" }

// ID of the command
func (IkeaRemotePress) ID() CommandID { return IkeaRemotePressCommand }

// Required
func (IkeaRemotePress) Required() bool { return false }

// Cluster ID of the command
func (IkeaRemotePress) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (IkeaRemotePress) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (IkeaRemotePress) MarshalJSON() ([]byte, error) { return []byte("7"), nil }

// MarshalZcl returns the wire format representation of IkeaRemotePress
func (v IkeaRemotePress) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.IkeaRemoteDirection.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the IkeaRemotePress struct
func (v *IkeaRemotePress) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.IkeaRemoteDirection).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v IkeaRemotePress) String() string {
	return zcl.Sprintf(
		"IkeaRemotePress{"+zcl.StrJoin([]string{
			"IkeaRemoteDirection(%v)",
		}, " ")+"}",
		v.IkeaRemoteDirection,
	)
}

type IkeaRemoteLongpressStart struct {
	IkeaRemoteDirection IkeaRemoteDirection
}

// IkeaRemoteLongpressStartCommand is the Command ID of IkeaRemoteLongpressStart
const IkeaRemoteLongpressStartCommand CommandID = 0x0008

// Values returns all values of IkeaRemoteLongpressStart
func (v *IkeaRemoteLongpressStart) Values() []zcl.Val {
	return []zcl.Val{
		&v.IkeaRemoteDirection,
	}
}

// Arguments returns all values of IkeaRemoteLongpressStart
func (v *IkeaRemoteLongpressStart) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "IkeaRemoteDirection", Argument: &v.IkeaRemoteDirection},
	}
}

// Name of the command
func (IkeaRemoteLongpressStart) Name() string { return "Ikea Remote Longpress Start" }

// ID of the command
func (IkeaRemoteLongpressStart) ID() CommandID { return IkeaRemoteLongpressStartCommand }

// Required
func (IkeaRemoteLongpressStart) Required() bool { return false }

// Cluster ID of the command
func (IkeaRemoteLongpressStart) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (IkeaRemoteLongpressStart) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (IkeaRemoteLongpressStart) MarshalJSON() ([]byte, error) { return []byte("8"), nil }

// MarshalZcl returns the wire format representation of IkeaRemoteLongpressStart
func (v IkeaRemoteLongpressStart) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.IkeaRemoteDirection.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the IkeaRemoteLongpressStart struct
func (v *IkeaRemoteLongpressStart) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.IkeaRemoteDirection).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v IkeaRemoteLongpressStart) String() string {
	return zcl.Sprintf(
		"IkeaRemoteLongpressStart{"+zcl.StrJoin([]string{
			"IkeaRemoteDirection(%v)",
		}, " ")+"}",
		v.IkeaRemoteDirection,
	)
}

type IkeaRemoteLongpressStop struct {
}

// IkeaRemoteLongpressStopCommand is the Command ID of IkeaRemoteLongpressStop
const IkeaRemoteLongpressStopCommand CommandID = 0x0009

// Values returns all values of IkeaRemoteLongpressStop
func (v *IkeaRemoteLongpressStop) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of IkeaRemoteLongpressStop
func (v *IkeaRemoteLongpressStop) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (IkeaRemoteLongpressStop) Name() string { return "Ikea Remote Longpress Stop" }

// ID of the command
func (IkeaRemoteLongpressStop) ID() CommandID { return IkeaRemoteLongpressStopCommand }

// Required
func (IkeaRemoteLongpressStop) Required() bool { return false }

// Cluster ID of the command
func (IkeaRemoteLongpressStop) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (IkeaRemoteLongpressStop) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (IkeaRemoteLongpressStop) MarshalJSON() ([]byte, error) { return []byte("9"), nil }

// MarshalZcl returns the wire format representation of IkeaRemoteLongpressStop
func (v IkeaRemoteLongpressStop) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the IkeaRemoteLongpressStop struct
func (v *IkeaRemoteLongpressStop) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v IkeaRemoteLongpressStop) String() string {
	return zcl.Sprintf(
		"IkeaRemoteLongpressStop{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

// AddSceneResponse Response to the add scene command.
type AddSceneResponse struct {
	Status  Status
	GroupId GroupId
	SceneId SceneId
}

// AddSceneResponseCommand is the Command ID of AddSceneResponse
const AddSceneResponseCommand CommandID = 0x0000

// Values returns all values of AddSceneResponse
func (v *AddSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
	}
}

// Arguments returns all values of AddSceneResponse
func (v *AddSceneResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
	}
}

// Name of the command
func (AddSceneResponse) Name() string { return "Add scene response" }

// ID of the command
func (AddSceneResponse) ID() CommandID { return AddSceneResponseCommand }

// Required
func (AddSceneResponse) Required() bool { return true }

// Cluster ID of the command
func (AddSceneResponse) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (AddSceneResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (AddSceneResponse) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

// MarshalZcl returns the wire format representation of AddSceneResponse
func (v AddSceneResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the AddSceneResponse struct
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

// String returns a log-friendly string representation of the struct
func (v AddSceneResponse) String() string {
	return zcl.Sprintf(
		"AddSceneResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"GroupId(%v)",
			"SceneId(%v)",
		}, " ")+"}",
		v.Status,
		v.GroupId,
		v.SceneId,
	)
}

// ViewSceneResponse Response to the view scene command.
type ViewSceneResponse struct {
	Status            Status
	GroupId           GroupId
	SceneId           SceneId
	TransitionTimeSec TransitionTimeSec
	SceneName         SceneName
	// SceneExtensions The format of each extension field set is a 16 bit field carrying the cluster ID,
	// followed by an 8 bit length field and the set of scene extension fields specified
	// in  the  relevant  cluster. The length field holds the length in octets of that
	// extension field set. Extension field set format:
	// {{clusterId1, length 1, {extension field set 1}}, {clusterId2, length 2, {extension field set 2}} ...}
	// I.e. the field would be a repeating struct with [ClusterID uint16] [OctetLength uint8] [AttrID ...uint16]
	SceneExtensions SceneExtensions
}

// ViewSceneResponseCommand is the Command ID of ViewSceneResponse
const ViewSceneResponseCommand CommandID = 0x0001

// Values returns all values of ViewSceneResponse
func (v *ViewSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
		&v.TransitionTimeSec,
		&v.SceneName,
		&v.SceneExtensions,
	}
}

// Arguments returns all values of ViewSceneResponse
func (v *ViewSceneResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
		{Name: "TransitionTimeSec", Argument: &v.TransitionTimeSec},
		{Name: "SceneName", Argument: &v.SceneName},
		{Name: "SceneExtensions", Argument: &v.SceneExtensions},
	}
}

// Name of the command
func (ViewSceneResponse) Name() string { return "View scene response" }

// ID of the command
func (ViewSceneResponse) ID() CommandID { return ViewSceneResponseCommand }

// Required
func (ViewSceneResponse) Required() bool { return true }

// Cluster ID of the command
func (ViewSceneResponse) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (ViewSceneResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (ViewSceneResponse) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

// MarshalZcl returns the wire format representation of ViewSceneResponse
func (v ViewSceneResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TransitionTimeSec.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneName.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneExtensions.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ViewSceneResponse struct
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

	if b, err = (&v.TransitionTimeSec).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneExtensions).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ViewSceneResponse) String() string {
	return zcl.Sprintf(
		"ViewSceneResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"GroupId(%v)",
			"SceneId(%v)",
			"TransitionTimeSec(%v)",
			"SceneName(%v)",
			"SceneExtensions(%v)",
		}, " ")+"}",
		v.Status,
		v.GroupId,
		v.SceneId,
		v.TransitionTimeSec,
		v.SceneName,
		v.SceneExtensions,
	)
}

// RemoveSceneResponse Response to the remove scene command.
type RemoveSceneResponse struct {
	Status  Status
	GroupId GroupId
	SceneId SceneId
}

// RemoveSceneResponseCommand is the Command ID of RemoveSceneResponse
const RemoveSceneResponseCommand CommandID = 0x0002

// Values returns all values of RemoveSceneResponse
func (v *RemoveSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
	}
}

// Arguments returns all values of RemoveSceneResponse
func (v *RemoveSceneResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
	}
}

// Name of the command
func (RemoveSceneResponse) Name() string { return "Remove scene response" }

// ID of the command
func (RemoveSceneResponse) ID() CommandID { return RemoveSceneResponseCommand }

// Required
func (RemoveSceneResponse) Required() bool { return true }

// Cluster ID of the command
func (RemoveSceneResponse) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (RemoveSceneResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (RemoveSceneResponse) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

// MarshalZcl returns the wire format representation of RemoveSceneResponse
func (v RemoveSceneResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RemoveSceneResponse struct
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

// String returns a log-friendly string representation of the struct
func (v RemoveSceneResponse) String() string {
	return zcl.Sprintf(
		"RemoveSceneResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"GroupId(%v)",
			"SceneId(%v)",
		}, " ")+"}",
		v.Status,
		v.GroupId,
		v.SceneId,
	)
}

// RemoveAllScenesResponse Response to the remove all scenes command.
type RemoveAllScenesResponse struct {
	Status  Status
	GroupId GroupId
}

// RemoveAllScenesResponseCommand is the Command ID of RemoveAllScenesResponse
const RemoveAllScenesResponseCommand CommandID = 0x0003

// Values returns all values of RemoveAllScenesResponse
func (v *RemoveAllScenesResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
	}
}

// Arguments returns all values of RemoveAllScenesResponse
func (v *RemoveAllScenesResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "GroupId", Argument: &v.GroupId},
	}
}

// Name of the command
func (RemoveAllScenesResponse) Name() string { return "Remove all scenes response" }

// ID of the command
func (RemoveAllScenesResponse) ID() CommandID { return RemoveAllScenesResponseCommand }

// Required
func (RemoveAllScenesResponse) Required() bool { return true }

// Cluster ID of the command
func (RemoveAllScenesResponse) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (RemoveAllScenesResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (RemoveAllScenesResponse) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

// MarshalZcl returns the wire format representation of RemoveAllScenesResponse
func (v RemoveAllScenesResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RemoveAllScenesResponse struct
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

// String returns a log-friendly string representation of the struct
func (v RemoveAllScenesResponse) String() string {
	return zcl.Sprintf(
		"RemoveAllScenesResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"GroupId(%v)",
		}, " ")+"}",
		v.Status,
		v.GroupId,
	)
}

// StoreSceneResponse Response to the store scene command.
type StoreSceneResponse struct {
	Status  Status
	GroupId GroupId
	SceneId SceneId
}

// StoreSceneResponseCommand is the Command ID of StoreSceneResponse
const StoreSceneResponseCommand CommandID = 0x0004

// Values returns all values of StoreSceneResponse
func (v *StoreSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
	}
}

// Arguments returns all values of StoreSceneResponse
func (v *StoreSceneResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
	}
}

// Name of the command
func (StoreSceneResponse) Name() string { return "Store scene response" }

// ID of the command
func (StoreSceneResponse) ID() CommandID { return StoreSceneResponseCommand }

// Required
func (StoreSceneResponse) Required() bool { return true }

// Cluster ID of the command
func (StoreSceneResponse) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (StoreSceneResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (StoreSceneResponse) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

// MarshalZcl returns the wire format representation of StoreSceneResponse
func (v StoreSceneResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the StoreSceneResponse struct
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

// String returns a log-friendly string representation of the struct
func (v StoreSceneResponse) String() string {
	return zcl.Sprintf(
		"StoreSceneResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"GroupId(%v)",
			"SceneId(%v)",
		}, " ")+"}",
		v.Status,
		v.GroupId,
		v.SceneId,
	)
}

// GetSceneMembershipResponse Shows details about scene membership.
type GetSceneMembershipResponse struct {
	Status Status
	// SceneCapacity specifies remaining number of scenes that can be added
	SceneCapacity SceneCapacity
	GroupId       GroupId
	SceneList     SceneList
}

// GetSceneMembershipResponseCommand is the Command ID of GetSceneMembershipResponse
const GetSceneMembershipResponseCommand CommandID = 0x0006

// Values returns all values of GetSceneMembershipResponse
func (v *GetSceneMembershipResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.SceneCapacity,
		&v.GroupId,
		&v.SceneList,
	}
}

// Arguments returns all values of GetSceneMembershipResponse
func (v *GetSceneMembershipResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "SceneCapacity", Argument: &v.SceneCapacity},
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneList", Argument: &v.SceneList},
	}
}

// Name of the command
func (GetSceneMembershipResponse) Name() string { return "Get scene membership response" }

// ID of the command
func (GetSceneMembershipResponse) ID() CommandID { return GetSceneMembershipResponseCommand }

// Required
func (GetSceneMembershipResponse) Required() bool { return true }

// Cluster ID of the command
func (GetSceneMembershipResponse) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (GetSceneMembershipResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (GetSceneMembershipResponse) MarshalJSON() ([]byte, error) { return []byte("6"), nil }

// MarshalZcl returns the wire format representation of GetSceneMembershipResponse
func (v GetSceneMembershipResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneCapacity.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetSceneMembershipResponse struct
func (v *GetSceneMembershipResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneCapacity).UnmarshalZcl(b); err != nil {
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

// String returns a log-friendly string representation of the struct
func (v GetSceneMembershipResponse) String() string {
	return zcl.Sprintf(
		"GetSceneMembershipResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"SceneCapacity(%v)",
			"GroupId(%v)",
			"SceneList(%v)",
		}, " ")+"}",
		v.Status,
		v.SceneCapacity,
		v.GroupId,
		v.SceneList,
	)
}

type EnhancedAddSceneResponse struct {
	Status  Status
	GroupId GroupId
	SceneId SceneId
}

// EnhancedAddSceneResponseCommand is the Command ID of EnhancedAddSceneResponse
const EnhancedAddSceneResponseCommand CommandID = 0x0040

// Values returns all values of EnhancedAddSceneResponse
func (v *EnhancedAddSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
	}
}

// Arguments returns all values of EnhancedAddSceneResponse
func (v *EnhancedAddSceneResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
	}
}

// Name of the command
func (EnhancedAddSceneResponse) Name() string { return "Enhanced Add Scene response" }

// ID of the command
func (EnhancedAddSceneResponse) ID() CommandID { return EnhancedAddSceneResponseCommand }

// Required
func (EnhancedAddSceneResponse) Required() bool { return false }

// Cluster ID of the command
func (EnhancedAddSceneResponse) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (EnhancedAddSceneResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (EnhancedAddSceneResponse) MarshalJSON() ([]byte, error) { return []byte("64"), nil }

// MarshalZcl returns the wire format representation of EnhancedAddSceneResponse
func (v EnhancedAddSceneResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the EnhancedAddSceneResponse struct
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

// String returns a log-friendly string representation of the struct
func (v EnhancedAddSceneResponse) String() string {
	return zcl.Sprintf(
		"EnhancedAddSceneResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"GroupId(%v)",
			"SceneId(%v)",
		}, " ")+"}",
		v.Status,
		v.GroupId,
		v.SceneId,
	)
}

// EnhancedViewSceneResponse A scene description.
type EnhancedViewSceneResponse struct {
	Status         Status
	GroupId        GroupId
	SceneId        SceneId
	TransitionTime TransitionTime
	SceneName      SceneName
	// SceneExtensions The format of each extension field set is a 16 bit field carrying the cluster ID,
	// followed by an 8 bit length field and the set of scene extension fields specified
	// in  the  relevant  cluster. The length field holds the length in octets of that
	// extension field set. Extension field set format:
	// {{clusterId1, length 1, {extension field set 1}}, {clusterId2, length 2, {extension field set 2}} ...}
	// I.e. the field would be a repeating struct with [ClusterID uint16] [OctetLength uint8] [AttrID ...uint16]
	SceneExtensions SceneExtensions
}

// EnhancedViewSceneResponseCommand is the Command ID of EnhancedViewSceneResponse
const EnhancedViewSceneResponseCommand CommandID = 0x0041

// Values returns all values of EnhancedViewSceneResponse
func (v *EnhancedViewSceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.GroupId,
		&v.SceneId,
		&v.TransitionTime,
		&v.SceneName,
		&v.SceneExtensions,
	}
}

// Arguments returns all values of EnhancedViewSceneResponse
func (v *EnhancedViewSceneResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "GroupId", Argument: &v.GroupId},
		{Name: "SceneId", Argument: &v.SceneId},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
		{Name: "SceneName", Argument: &v.SceneName},
		{Name: "SceneExtensions", Argument: &v.SceneExtensions},
	}
}

// Name of the command
func (EnhancedViewSceneResponse) Name() string { return "Enhanced View Scene response" }

// ID of the command
func (EnhancedViewSceneResponse) ID() CommandID { return EnhancedViewSceneResponseCommand }

// Required
func (EnhancedViewSceneResponse) Required() bool { return false }

// Cluster ID of the command
func (EnhancedViewSceneResponse) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (EnhancedViewSceneResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (EnhancedViewSceneResponse) MarshalJSON() ([]byte, error) { return []byte("65"), nil }

// MarshalZcl returns the wire format representation of EnhancedViewSceneResponse
func (v EnhancedViewSceneResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneName.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SceneExtensions.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the EnhancedViewSceneResponse struct
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

	if b, err = (&v.SceneName).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SceneExtensions).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v EnhancedViewSceneResponse) String() string {
	return zcl.Sprintf(
		"EnhancedViewSceneResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"GroupId(%v)",
			"SceneId(%v)",
			"TransitionTime(%v)",
			"SceneName(%v)",
			"SceneExtensions(%v)",
		}, " ")+"}",
		v.Status,
		v.GroupId,
		v.SceneId,
		v.TransitionTime,
		v.SceneName,
		v.SceneExtensions,
	)
}

type CopySceneResponse struct {
	Status      Status
	FromGroupId GroupId
	FromSceneId SceneId
}

// CopySceneResponseCommand is the Command ID of CopySceneResponse
const CopySceneResponseCommand CommandID = 0x0042

// Values returns all values of CopySceneResponse
func (v *CopySceneResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.FromGroupId,
		&v.FromSceneId,
	}
}

// Arguments returns all values of CopySceneResponse
func (v *CopySceneResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "FromGroupId", Argument: &v.FromGroupId},
		{Name: "FromSceneId", Argument: &v.FromSceneId},
	}
}

// Name of the command
func (CopySceneResponse) Name() string { return "Copy Scene Response" }

// ID of the command
func (CopySceneResponse) ID() CommandID { return CopySceneResponseCommand }

// Required
func (CopySceneResponse) Required() bool { return false }

// Cluster ID of the command
func (CopySceneResponse) Cluster() zcl.ClusterID { return ScenesID }

// MnfCode returns the manufacturer code (if any) of the command
func (CopySceneResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (CopySceneResponse) MarshalJSON() ([]byte, error) { return []byte("66"), nil }

// MarshalZcl returns the wire format representation of CopySceneResponse
func (v CopySceneResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FromGroupId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.FromSceneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the CopySceneResponse struct
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

// String returns a log-friendly string representation of the struct
func (v CopySceneResponse) String() string {
	return zcl.Sprintf(
		"CopySceneResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"FromGroupId(%v)",
			"FromSceneId(%v)",
		}, " ")+"}",
		v.Status,
		v.FromGroupId,
		v.FromSceneId,
	)
}
