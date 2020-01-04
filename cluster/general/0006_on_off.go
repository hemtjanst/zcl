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

// OffCommand is the Command ID of Off
const OffCommand CommandID = 0x0000

// Values returns all values of Off
func (v *Off) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of Off
func (v *Off) Arguments() []zcl.Argument {
	return []zcl.Argument{}
}

// Name of the command
func (Off) Name() string { return "Off" }

// ID of the command
func (Off) ID() CommandID { return OffCommand }

// Required
func (Off) Required() bool { return true }

// Cluster ID of the command
func (Off) Cluster() zcl.ClusterID { return OnOffID }

// MnfCode returns the manufacturer code (if any) of the command
func (Off) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (Off) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

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

// OnCommand is the Command ID of On
const OnCommand CommandID = 0x0001

// Values returns all values of On
func (v *On) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of On
func (v *On) Arguments() []zcl.Argument {
	return []zcl.Argument{}
}

// Name of the command
func (On) Name() string { return "On" }

// ID of the command
func (On) ID() CommandID { return OnCommand }

// Required
func (On) Required() bool { return true }

// Cluster ID of the command
func (On) Cluster() zcl.ClusterID { return OnOffID }

// MnfCode returns the manufacturer code (if any) of the command
func (On) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (On) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

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

// ToggleCommand is the Command ID of Toggle
const ToggleCommand CommandID = 0x0002

// Values returns all values of Toggle
func (v *Toggle) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of Toggle
func (v *Toggle) Arguments() []zcl.Argument {
	return []zcl.Argument{}
}

// Name of the command
func (Toggle) Name() string { return "Toggle" }

// ID of the command
func (Toggle) ID() CommandID { return ToggleCommand }

// Required
func (Toggle) Required() bool { return true }

// Cluster ID of the command
func (Toggle) Cluster() zcl.ClusterID { return OnOffID }

// MnfCode returns the manufacturer code (if any) of the command
func (Toggle) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (Toggle) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

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
func (v *OffWithEffect) Arguments() []zcl.Argument {
	return []zcl.Argument{
		&v.EffectIdentifier,
		&v.EffectVariant,
	}
}

// Name of the command
func (OffWithEffect) Name() string { return "Off with effect" }

// ID of the command
func (OffWithEffect) ID() CommandID { return OffWithEffectCommand }

// Required
func (OffWithEffect) Required() bool { return true }

// Cluster ID of the command
func (OffWithEffect) Cluster() zcl.ClusterID { return OnOffID }

// MnfCode returns the manufacturer code (if any) of the command
func (OffWithEffect) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (OffWithEffect) MarshalJSON() ([]byte, error) { return []byte("64"), nil }

// MarshalZcl returns the wire format representation of OffWithEffect
func (v OffWithEffect) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
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

// OnWithRecallGlobalSceneCommand is the Command ID of OnWithRecallGlobalScene
const OnWithRecallGlobalSceneCommand CommandID = 0x0041

// Values returns all values of OnWithRecallGlobalScene
func (v *OnWithRecallGlobalScene) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of OnWithRecallGlobalScene
func (v *OnWithRecallGlobalScene) Arguments() []zcl.Argument {
	return []zcl.Argument{}
}

// Name of the command
func (OnWithRecallGlobalScene) Name() string { return "On with recall global scene" }

// ID of the command
func (OnWithRecallGlobalScene) ID() CommandID { return OnWithRecallGlobalSceneCommand }

// Required
func (OnWithRecallGlobalScene) Required() bool { return true }

// Cluster ID of the command
func (OnWithRecallGlobalScene) Cluster() zcl.ClusterID { return OnOffID }

// MnfCode returns the manufacturer code (if any) of the command
func (OnWithRecallGlobalScene) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (OnWithRecallGlobalScene) MarshalJSON() ([]byte, error) { return []byte("65"), nil }

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
func (v *OnWithTimedOff) Arguments() []zcl.Argument {
	return []zcl.Argument{
		&v.OnOffControl,
		&v.OnTime,
		&v.OffWaitTime,
	}
}

// Name of the command
func (OnWithTimedOff) Name() string { return "On with timed off" }

// ID of the command
func (OnWithTimedOff) ID() CommandID { return OnWithTimedOffCommand }

// Required
func (OnWithTimedOff) Required() bool { return true }

// Cluster ID of the command
func (OnWithTimedOff) Cluster() zcl.ClusterID { return OnOffID }

// MnfCode returns the manufacturer code (if any) of the command
func (OnWithTimedOff) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (OnWithTimedOff) MarshalJSON() ([]byte, error) { return []byte("66"), nil }

// MarshalZcl returns the wire format representation of OnWithTimedOff
func (v OnWithTimedOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
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
