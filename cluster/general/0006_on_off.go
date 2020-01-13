package general

import "hemtjan.st/zcl"

// OnOff
const OnOffID zcl.ClusterID = 6

var OnOffCluster = zcl.Cluster{
	Name: "On/Off",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		OffCommand:                     func() zcl.Command { return new(Off) },
		OnCommand:                      func() zcl.Command { return new(On) },
		ToggleCommand:                  func() zcl.Command { return new(Toggle) },
		OffWithEffectCommand:           func() zcl.Command { return new(OffWithEffect) },
		OnWithRecallGlobalSceneCommand: func() zcl.Command { return new(OnWithRecallGlobalScene) },
		OnWithTimedOffCommand:          func() zcl.Command { return new(OnWithTimedOff) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		OnOffAttr:              func() zcl.Attr { return new(OnOff) },
		GlobalSceneControlAttr: func() zcl.Attr { return new(GlobalSceneControl) },
		OnTimeAttr:             func() zcl.Attr { return new(OnTime) },
		OffWaitTimeAttr:        func() zcl.Attr { return new(OffWaitTime) },
		PoweronOnOffAttr:       func() zcl.Attr { return new(PoweronOnOff) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr: []zcl.AttrID{
		OnOffAttr,
	},
}

// Off On receipt of this command, a device shall enter its 'Off' state. This state is device dependent, but it is recommended that it is used for power off or similar functions.
type Off struct {
}

type OffHandler interface {
	HandleOff(frame Frame, cmd *Off) error
}

// OffCommand is the Command ID of Off
const OffCommand CommandID = 0x0000

// Values returns all values of Off
func (v *Off) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of Off
func (v *Off) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (Off) Name() string { return `Off` }

// Description of the command
func (Off) Description() string {
	return `On receipt of this command, a device shall enter its 'Off' state. This state is device dependent, but it is recommended that it is used for power off or similar functions.`
}

// ID of the command
func (Off) ID() CommandID { return OffCommand }

// Required
func (Off) Required() bool { return true }

// Cluster ID of the command
func (Off) Cluster() zcl.ClusterID { return OnOffID }

// Direction of the command
func (Off) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (Off) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (Off) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *Off) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h OffHandler
	if h, found = handler.(OffHandler); found {
		err = h.HandleOff(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of Off
func (v Off) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the Off struct
func (v *Off) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v Off) String() string {
	return zcl.Sprintf(
		"Off{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

// On On receipt of this command, a device shall enter its 'On' state. This state is device dependent, but it is recommended that it is used for power on or similar functions.
type On struct {
}

type OnHandler interface {
	HandleOn(frame Frame, cmd *On) error
}

// OnCommand is the Command ID of On
const OnCommand CommandID = 0x0001

// Values returns all values of On
func (v *On) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of On
func (v *On) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (On) Name() string { return `On` }

// Description of the command
func (On) Description() string {
	return `On receipt of this command, a device shall enter its 'On' state. This state is device dependent, but it is recommended that it is used for power on or similar functions.`
}

// ID of the command
func (On) ID() CommandID { return OnCommand }

// Required
func (On) Required() bool { return true }

// Cluster ID of the command
func (On) Cluster() zcl.ClusterID { return OnOffID }

// Direction of the command
func (On) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (On) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (On) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *On) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h OnHandler
	if h, found = handler.(OnHandler); found {
		err = h.HandleOn(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of On
func (v On) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the On struct
func (v *On) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v On) String() string {
	return zcl.Sprintf(
		"On{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

// Toggle On receipt of this command, if a device is in its ‘Off’ state it shall enter its 'On' state. Otherwise, if it is in its ‘On’ state it shall enter its 'Off' state.
type Toggle struct {
}

type ToggleHandler interface {
	HandleToggle(frame Frame, cmd *Toggle) error
}

// ToggleCommand is the Command ID of Toggle
const ToggleCommand CommandID = 0x0002

// Values returns all values of Toggle
func (v *Toggle) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of Toggle
func (v *Toggle) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (Toggle) Name() string { return `Toggle` }

// Description of the command
func (Toggle) Description() string {
	return `On receipt of this command, if a device is in its ‘Off’ state it shall enter its 'On' state. Otherwise, if it is in its ‘On’ state it shall enter its 'Off' state.`
}

// ID of the command
func (Toggle) ID() CommandID { return ToggleCommand }

// Required
func (Toggle) Required() bool { return true }

// Cluster ID of the command
func (Toggle) Cluster() zcl.ClusterID { return OnOffID }

// Direction of the command
func (Toggle) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (Toggle) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (Toggle) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

func (v *Toggle) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ToggleHandler
	if h, found = handler.(ToggleHandler); found {
		err = h.HandleToggle(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of Toggle
func (v Toggle) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the Toggle struct
func (v *Toggle) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v Toggle) String() string {
	return zcl.Sprintf(
		"Toggle{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type OffWithEffect struct {
	// EffectIdentifier when turning lights off
	EffectIdentifier EffectIdentifier
	EffectVariant    EffectVariant
}

type OffWithEffectHandler interface {
	HandleOffWithEffect(frame Frame, cmd *OffWithEffect) error
}

// OffWithEffectCommand is the Command ID of OffWithEffect
const OffWithEffectCommand CommandID = 0x0040

// Values returns all values of OffWithEffect
func (v *OffWithEffect) Values() []zcl.Val {
	return []zcl.Val{
		&v.EffectIdentifier,
		&v.EffectVariant,
	}
}

// Arguments returns all values of OffWithEffect
func (v *OffWithEffect) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "EffectIdentifier", Argument: &v.EffectIdentifier},
		{Name: "EffectVariant", Argument: &v.EffectVariant},
	}
}

// Name of the command
func (OffWithEffect) Name() string { return `Off with effect` }

// Description of the command
func (OffWithEffect) Description() string { return `` }

// ID of the command
func (OffWithEffect) ID() CommandID { return OffWithEffectCommand }

// Required
func (OffWithEffect) Required() bool { return true }

// Cluster ID of the command
func (OffWithEffect) Cluster() zcl.ClusterID { return OnOffID }

// Direction of the command
func (OffWithEffect) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (OffWithEffect) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (OffWithEffect) MarshalJSON() ([]byte, error) { return []byte("64"), nil }

func (v *OffWithEffect) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h OffWithEffectHandler
	if h, found = handler.(OffWithEffectHandler); found {
		err = h.HandleOffWithEffect(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of OffWithEffect
func (v OffWithEffect) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.EffectIdentifier.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.EffectVariant.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the OffWithEffect struct
func (v *OffWithEffect) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.EffectIdentifier).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.EffectVariant).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v OffWithEffect) String() string {
	return zcl.Sprintf(
		"OffWithEffect{"+zcl.StrJoin([]string{
			"EffectIdentifier(%v)",
			"EffectVariant(%v)",
		}, " ")+"}",
		v.EffectIdentifier,
		v.EffectVariant,
	)
}

// OnWithRecallGlobalScene The on with recall global scene command allows the recall of the light settings when the light was turned off.
type OnWithRecallGlobalScene struct {
}

type OnWithRecallGlobalSceneHandler interface {
	HandleOnWithRecallGlobalScene(frame Frame, cmd *OnWithRecallGlobalScene) error
}

// OnWithRecallGlobalSceneCommand is the Command ID of OnWithRecallGlobalScene
const OnWithRecallGlobalSceneCommand CommandID = 0x0041

// Values returns all values of OnWithRecallGlobalScene
func (v *OnWithRecallGlobalScene) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of OnWithRecallGlobalScene
func (v *OnWithRecallGlobalScene) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (OnWithRecallGlobalScene) Name() string { return `On with recall global scene` }

// Description of the command
func (OnWithRecallGlobalScene) Description() string {
	return `The on with recall global scene command allows the recall of the light settings when the light was turned off.`
}

// ID of the command
func (OnWithRecallGlobalScene) ID() CommandID { return OnWithRecallGlobalSceneCommand }

// Required
func (OnWithRecallGlobalScene) Required() bool { return true }

// Cluster ID of the command
func (OnWithRecallGlobalScene) Cluster() zcl.ClusterID { return OnOffID }

// Direction of the command
func (OnWithRecallGlobalScene) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (OnWithRecallGlobalScene) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (OnWithRecallGlobalScene) MarshalJSON() ([]byte, error) { return []byte("65"), nil }

func (v *OnWithRecallGlobalScene) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h OnWithRecallGlobalSceneHandler
	if h, found = handler.(OnWithRecallGlobalSceneHandler); found {
		err = h.HandleOnWithRecallGlobalScene(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of OnWithRecallGlobalScene
func (v OnWithRecallGlobalScene) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the OnWithRecallGlobalScene struct
func (v *OnWithRecallGlobalScene) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v OnWithRecallGlobalScene) String() string {
	return zcl.Sprintf(
		"OnWithRecallGlobalScene{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

// OnWithTimedOff Allows lamps to be turned on for a specific duration with a guarded off duration so that should the lamp be subsequently switched off, further on with timed off commands, received during this time, are prevented from turning the lamps back on.
type OnWithTimedOff struct {
	OnOffControl OnOffControl
	OnTime       OnTime
	OffWaitTime  OffWaitTime
}

type OnWithTimedOffHandler interface {
	HandleOnWithTimedOff(frame Frame, cmd *OnWithTimedOff) error
}

// OnWithTimedOffCommand is the Command ID of OnWithTimedOff
const OnWithTimedOffCommand CommandID = 0x0042

// Values returns all values of OnWithTimedOff
func (v *OnWithTimedOff) Values() []zcl.Val {
	return []zcl.Val{
		&v.OnOffControl,
		&v.OnTime,
		&v.OffWaitTime,
	}
}

// Arguments returns all values of OnWithTimedOff
func (v *OnWithTimedOff) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "OnOffControl", Argument: &v.OnOffControl},
		{Name: "OnTime", Argument: &v.OnTime},
		{Name: "OffWaitTime", Argument: &v.OffWaitTime},
	}
}

// Name of the command
func (OnWithTimedOff) Name() string { return `On with timed off` }

// Description of the command
func (OnWithTimedOff) Description() string {
	return `Allows lamps to be turned on for a specific duration with a guarded off duration so that should the lamp be subsequently switched off, further on with timed off commands, received during this time, are prevented from turning the lamps back on.`
}

// ID of the command
func (OnWithTimedOff) ID() CommandID { return OnWithTimedOffCommand }

// Required
func (OnWithTimedOff) Required() bool { return true }

// Cluster ID of the command
func (OnWithTimedOff) Cluster() zcl.ClusterID { return OnOffID }

// Direction of the command
func (OnWithTimedOff) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (OnWithTimedOff) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (OnWithTimedOff) MarshalJSON() ([]byte, error) { return []byte("66"), nil }

func (v *OnWithTimedOff) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h OnWithTimedOffHandler
	if h, found = handler.(OnWithTimedOffHandler); found {
		err = h.HandleOnWithTimedOff(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of OnWithTimedOff
func (v OnWithTimedOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.OnOffControl.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.OnTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.OffWaitTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the OnWithTimedOff struct
func (v *OnWithTimedOff) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.OnOffControl).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.OnTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.OffWaitTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v OnWithTimedOff) String() string {
	return zcl.Sprintf(
		"OnWithTimedOff{"+zcl.StrJoin([]string{
			"OnOffControl(%v)",
			"OnTime(%v)",
			"OffWaitTime(%v)",
		}, " ")+"}",
		v.OnOffControl,
		v.OnTime,
		v.OffWaitTime,
	)
}
