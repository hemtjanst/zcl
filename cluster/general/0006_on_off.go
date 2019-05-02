// Attributes and commands for switching devices between 'On' and 'Off' states.
package general

import (
	"neotor.se/zcl"
)

// OnOff
const OnOffID zcl.ClusterID = 6

var OnOffCluster = zcl.Cluster{
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
		OnoffAttr:              func() zcl.Attr { return new(Onoff) },
		GlobalscenecontrolAttr: func() zcl.Attr { return new(Globalscenecontrol) },
		OntimeAttr:             func() zcl.Attr { return new(Ontime) },
		OffwaittimeAttr:        func() zcl.Attr { return new(Offwaittime) },
		PoweronOnoffAttr:       func() zcl.Attr { return new(PoweronOnoff) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr: []zcl.AttrID{
		OnoffAttr,
	},
}

// On receipt of this command, a device shall enter its 'Off' state. This state is device dependent, but it is recommended that it is used for power off or similar functions.
type Off struct {
}

const OffCommand zcl.CommandID = 0

func (v *Off) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v Off) ID() zcl.CommandID {
	return OffCommand
}

func (v Off) Cluster() zcl.ClusterID {
	return OnOffID
}

func (v Off) MnfCode() []byte {
	return []byte{}
}

func (v Off) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *Off) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// On receipt of this command, a device shall enter its 'On' state. This state is device dependent, but it is recommended that it is used for power on or similar functions.
type On struct {
}

const OnCommand zcl.CommandID = 1

func (v *On) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v On) ID() zcl.CommandID {
	return OnCommand
}

func (v On) Cluster() zcl.ClusterID {
	return OnOffID
}

func (v On) MnfCode() []byte {
	return []byte{}
}

func (v On) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *On) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// On receipt of this command, if a device is in its ‘Off’ state it shall enter its 'On' state. Otherwise, if it is in its ‘On’ state it shall enter its 'Off' state.
type Toggle struct {
}

const ToggleCommand zcl.CommandID = 2

func (v *Toggle) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v Toggle) ID() zcl.CommandID {
	return ToggleCommand
}

func (v Toggle) Cluster() zcl.ClusterID {
	return OnOffID
}

func (v Toggle) MnfCode() []byte {
	return []byte{}
}

func (v Toggle) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *Toggle) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

type OffWithEffect struct {
	// The effect identifier field specifies the fading effect to use.
	EffectIdentifier zcl.Zenum8
	// The effect identifier field specifies the effect variant to use.
	EffectVariant zcl.Zenum8
}

const OffWithEffectCommand zcl.CommandID = 64

func (v *OffWithEffect) Values() []zcl.Val {
	return []zcl.Val{
		&v.EffectIdentifier,
		&v.EffectVariant,
	}
}

func (v OffWithEffect) ID() zcl.CommandID {
	return OffWithEffectCommand
}

func (v OffWithEffect) Cluster() zcl.ClusterID {
	return OnOffID
}

func (v OffWithEffect) MnfCode() []byte {
	return []byte{}
}

func (v OffWithEffect) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.EffectIdentifier.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.EffectVariant.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

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

// The on with recall global scene command allows the recall of the light settings when the light was turned off.
type OnWithRecallGlobalScene struct {
}

const OnWithRecallGlobalSceneCommand zcl.CommandID = 65

func (v *OnWithRecallGlobalScene) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v OnWithRecallGlobalScene) ID() zcl.CommandID {
	return OnWithRecallGlobalSceneCommand
}

func (v OnWithRecallGlobalScene) Cluster() zcl.ClusterID {
	return OnOffID
}

func (v OnWithRecallGlobalScene) MnfCode() []byte {
	return []byte{}
}

func (v OnWithRecallGlobalScene) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *OnWithRecallGlobalScene) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// Allows lamps to be turned on for a specific duration with a guarded off duration so that should the lamp be subsequently switched off, further on with timed off commands, received during this time, are prevented from turning the lamps back on.
type OnWithTimedOff struct {
	// The effect identifier field specifies the fading effect to use.
	OnOffControl zcl.Zbmp8
	// Length of time (in 1/10ths second) that the lamp is to remain on.
	OnTime zcl.Zu16
	// Length of time (in 1/10ths second) that the lamp shall remain off.
	OffWaitTime zcl.Zu16
}

const OnWithTimedOffCommand zcl.CommandID = 66

func (v *OnWithTimedOff) Values() []zcl.Val {
	return []zcl.Val{
		&v.OnOffControl,
		&v.OnTime,
		&v.OffWaitTime,
	}
}

func (v OnWithTimedOff) ID() zcl.CommandID {
	return OnWithTimedOffCommand
}

func (v OnWithTimedOff) Cluster() zcl.ClusterID {
	return OnOffID
}

func (v OnWithTimedOff) MnfCode() []byte {
	return []byte{}
}

func (v OnWithTimedOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.OnOffControl.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.OnTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.OffWaitTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

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

const OnoffAttr zcl.AttrID = 0

type Onoff zcl.Zbool

func (a Onoff) ID() zcl.AttrID         { return OnoffAttr }
func (a Onoff) Cluster() zcl.ClusterID { return OnOffID }
func (a *Onoff) Value() *Onoff         { return a }
func (a Onoff) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *Onoff) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = Onoff(*nt)
	return br, err
}

func (a Onoff) Readable() bool   { return true }
func (a Onoff) Writable() bool   { return false }
func (a Onoff) Reportable() bool { return true }
func (a Onoff) SceneIndex() int  { return 1 }

func (a Onoff) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x01:
		return "On"
	}
	return zcl.Sprintf("%s", zcl.Zbool(a))
}

// IsOff checks if Onoff equals the value for Off (0x00)
func (a Onoff) IsOff() bool { return a == 0x00 }

// SetOff sets Onoff to Off (0x00)
func (a *Onoff) SetOff() { *a = 0x00 }

// IsOn checks if Onoff equals the value for On (0x01)
func (a Onoff) IsOn() bool { return a == 0x01 }

// SetOn sets Onoff to On (0x01)
func (a *Onoff) SetOn() { *a = 0x01 }

const GlobalscenecontrolAttr zcl.AttrID = 16384

type Globalscenecontrol zcl.Zbool

func (a Globalscenecontrol) ID() zcl.AttrID              { return GlobalscenecontrolAttr }
func (a Globalscenecontrol) Cluster() zcl.ClusterID      { return OnOffID }
func (a *Globalscenecontrol) Value() *Globalscenecontrol { return a }
func (a Globalscenecontrol) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *Globalscenecontrol) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = Globalscenecontrol(*nt)
	return br, err
}

func (a Globalscenecontrol) Readable() bool   { return true }
func (a Globalscenecontrol) Writable() bool   { return false }
func (a Globalscenecontrol) Reportable() bool { return false }
func (a Globalscenecontrol) SceneIndex() int  { return -1 }

func (a Globalscenecontrol) String() string {
	return zcl.Sprintf("%s", zcl.Zbool(a))
}

const OntimeAttr zcl.AttrID = 16385

type Ontime zcl.Zu16

func (a Ontime) ID() zcl.AttrID         { return OntimeAttr }
func (a Ontime) Cluster() zcl.ClusterID { return OnOffID }
func (a *Ontime) Value() *Ontime        { return a }
func (a Ontime) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Ontime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Ontime(*nt)
	return br, err
}

func (a Ontime) Readable() bool   { return true }
func (a Ontime) Writable() bool   { return false }
func (a Ontime) Reportable() bool { return false }
func (a Ontime) SceneIndex() int  { return -1 }

func (a Ontime) String() string {
	return zcl.Duration(int(a), 10).String()
}

const OffwaittimeAttr zcl.AttrID = 16386

type Offwaittime zcl.Zu16

func (a Offwaittime) ID() zcl.AttrID         { return OffwaittimeAttr }
func (a Offwaittime) Cluster() zcl.ClusterID { return OnOffID }
func (a *Offwaittime) Value() *Offwaittime   { return a }
func (a Offwaittime) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Offwaittime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Offwaittime(*nt)
	return br, err
}

func (a Offwaittime) Readable() bool   { return true }
func (a Offwaittime) Writable() bool   { return false }
func (a Offwaittime) Reportable() bool { return false }
func (a Offwaittime) SceneIndex() int  { return -1 }

func (a Offwaittime) String() string {
	return zcl.Duration(int(a), 10).String()
}

const PoweronOnoffAttr zcl.AttrID = 16387

type PoweronOnoff zcl.Zenum8

func (a PoweronOnoff) ID() zcl.AttrID         { return PoweronOnoffAttr }
func (a PoweronOnoff) Cluster() zcl.ClusterID { return OnOffID }
func (a *PoweronOnoff) Value() *PoweronOnoff  { return a }
func (a PoweronOnoff) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *PoweronOnoff) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PoweronOnoff(*nt)
	return br, err
}

func (a PoweronOnoff) Readable() bool   { return true }
func (a PoweronOnoff) Writable() bool   { return true }
func (a PoweronOnoff) Reportable() bool { return false }
func (a PoweronOnoff) SceneIndex() int  { return -1 }

func (a PoweronOnoff) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x01:
		return "On"
	case 0xFF:
		return "Previous"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsOff checks if PoweronOnoff equals the value for Off (0x00)
func (a PoweronOnoff) IsOff() bool { return a == 0x00 }

// SetOff sets PoweronOnoff to Off (0x00)
func (a *PoweronOnoff) SetOff() { *a = 0x00 }

// IsOn checks if PoweronOnoff equals the value for On (0x01)
func (a PoweronOnoff) IsOn() bool { return a == 0x01 }

// SetOn sets PoweronOnoff to On (0x01)
func (a *PoweronOnoff) SetOn() { *a = 0x01 }

// IsPrevious checks if PoweronOnoff equals the value for Previous (0xFF)
func (a PoweronOnoff) IsPrevious() bool { return a == 0xFF }

// SetPrevious sets PoweronOnoff to Previous (0xFF)
func (a *PoweronOnoff) SetPrevious() { *a = 0xFF }
