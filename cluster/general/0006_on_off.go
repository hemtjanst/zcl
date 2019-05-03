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

func (v Off) String() string {
	var str []string
	return "Off{" + zcl.StrJoin(str, " ") + "}"
}

func (Off) Name() string { return "Off" }

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

func (v On) String() string {
	var str []string
	return "On{" + zcl.StrJoin(str, " ") + "}"
}

func (On) Name() string { return "On" }

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

func (v Toggle) String() string {
	var str []string
	return "Toggle{" + zcl.StrJoin(str, " ") + "}"
}

func (Toggle) Name() string { return "Toggle" }

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

func (v OffWithEffect) EffectIdentifierString() string {
	switch v.EffectIdentifier {
	case 0x00:
		return "Delayed all off"
	case 0x01:
		return "Dying light"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v.EffectIdentifier))
}
func (v OffWithEffect) EffectVariantString() string {
	switch v.EffectVariant {
	case 0x00:
		return "Fade to off in 0.8 seconds / 20% dim up in 0.5 then fade to off in 1 second"
	case 0x01:
		return "No Fade"
	case 0x02:
		return "50% dim down in 0.8s then fade off in 12s"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v.EffectVariant))
}

func (v OffWithEffect) String() string {
	var str []string
	str = append(str, "EffectIdentifier["+v.EffectIdentifierString()+"]")
	str = append(str, "EffectVariant["+v.EffectVariantString()+"]")
	return "OffWithEffect{" + zcl.StrJoin(str, " ") + "}"
}

func (OffWithEffect) Name() string { return "Off with effect" }

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

func (v OnWithRecallGlobalScene) String() string {
	var str []string
	return "OnWithRecallGlobalScene{" + zcl.StrJoin(str, " ") + "}"
}

func (OnWithRecallGlobalScene) Name() string { return "On with recall global scene" }

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

func (v OnWithTimedOff) OnOffControlString() string {
	var bstr []string
	if zcl.BitmapTest([]byte(v.OnOffControl), 0) {
		bstr = append(bstr, "Accept only when on")
	}
	return zcl.StrJoin(bstr, ", ")
}
func (v OnWithTimedOff) OnTimeString() string {
	return zcl.Seconds.Format(float64(v.OnTime) / 10)
}
func (v OnWithTimedOff) OffWaitTimeString() string {
	return zcl.Seconds.Format(float64(v.OffWaitTime) / 10)
}

func (v OnWithTimedOff) String() string {
	var str []string
	str = append(str, "OnOffControl["+v.OnOffControlString()+"]")
	str = append(str, "OnTime["+v.OnTimeString()+"]")
	str = append(str, "OffWaitTime["+v.OffWaitTimeString()+"]")
	return "OnWithTimedOff{" + zcl.StrJoin(str, " ") + "}"
}

func (OnWithTimedOff) Name() string { return "On with timed off" }

// Onoff is an autogenerated attribute in the OnOff cluster
type Onoff zcl.Zbool

const OnoffAttr zcl.AttrID = 0

func (Onoff) ID() zcl.AttrID                { return OnoffAttr }
func (Onoff) Cluster() zcl.ClusterID        { return OnOffID }
func (Onoff) Name() string                  { return "OnOff" }
func (Onoff) Readable() bool                { return true }
func (Onoff) Writable() bool                { return false }
func (Onoff) Reportable() bool              { return true }
func (Onoff) SceneIndex() int               { return 1 }
func (a *Onoff) Value() *Onoff              { return a }
func (a Onoff) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *Onoff) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = Onoff(*nt)
	return br, err
}

func (a Onoff) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x01:
		return "On"
	}
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

// IsOff checks if Onoff equals the value for Off (0x00)
func (a Onoff) IsOff() bool { return a == 0x00 }

// SetOff sets Onoff to Off (0x00)
func (a *Onoff) SetOff() { *a = 0x00 }

// IsOn checks if Onoff equals the value for On (0x01)
func (a Onoff) IsOn() bool { return a == 0x01 }

// SetOn sets Onoff to On (0x01)
func (a *Onoff) SetOn() { *a = 0x01 }

// Globalscenecontrol is an autogenerated attribute in the OnOff cluster
type Globalscenecontrol zcl.Zbool

const GlobalscenecontrolAttr zcl.AttrID = 16384

func (Globalscenecontrol) ID() zcl.AttrID                { return GlobalscenecontrolAttr }
func (Globalscenecontrol) Cluster() zcl.ClusterID        { return OnOffID }
func (Globalscenecontrol) Name() string                  { return "GlobalSceneControl" }
func (Globalscenecontrol) Readable() bool                { return true }
func (Globalscenecontrol) Writable() bool                { return false }
func (Globalscenecontrol) Reportable() bool              { return false }
func (Globalscenecontrol) SceneIndex() int               { return -1 }
func (a *Globalscenecontrol) Value() *Globalscenecontrol { return a }
func (a Globalscenecontrol) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *Globalscenecontrol) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = Globalscenecontrol(*nt)
	return br, err
}

func (a Globalscenecontrol) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

// Ontime is an autogenerated attribute in the OnOff cluster
type Ontime zcl.Zu16

const OntimeAttr zcl.AttrID = 16385

func (Ontime) ID() zcl.AttrID                { return OntimeAttr }
func (Ontime) Cluster() zcl.ClusterID        { return OnOffID }
func (Ontime) Name() string                  { return "OnTime" }
func (Ontime) Readable() bool                { return true }
func (Ontime) Writable() bool                { return false }
func (Ontime) Reportable() bool              { return false }
func (Ontime) SceneIndex() int               { return -1 }
func (a *Ontime) Value() *Ontime             { return a }
func (a Ontime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Ontime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Ontime(*nt)
	return br, err
}

func (a Ontime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

// Offwaittime is an autogenerated attribute in the OnOff cluster
type Offwaittime zcl.Zu16

const OffwaittimeAttr zcl.AttrID = 16386

func (Offwaittime) ID() zcl.AttrID                { return OffwaittimeAttr }
func (Offwaittime) Cluster() zcl.ClusterID        { return OnOffID }
func (Offwaittime) Name() string                  { return "OffWaitTime" }
func (Offwaittime) Readable() bool                { return true }
func (Offwaittime) Writable() bool                { return false }
func (Offwaittime) Reportable() bool              { return false }
func (Offwaittime) SceneIndex() int               { return -1 }
func (a *Offwaittime) Value() *Offwaittime        { return a }
func (a Offwaittime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Offwaittime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Offwaittime(*nt)
	return br, err
}

func (a Offwaittime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

// PoweronOnoff is an autogenerated attribute in the OnOff cluster
type PoweronOnoff zcl.Zenum8

const PoweronOnoffAttr zcl.AttrID = 16387

func (PoweronOnoff) ID() zcl.AttrID                { return PoweronOnoffAttr }
func (PoweronOnoff) Cluster() zcl.ClusterID        { return OnOffID }
func (PoweronOnoff) Name() string                  { return "PowerOn OnOff" }
func (PoweronOnoff) Readable() bool                { return true }
func (PoweronOnoff) Writable() bool                { return true }
func (PoweronOnoff) Reportable() bool              { return false }
func (PoweronOnoff) SceneIndex() int               { return -1 }
func (a *PoweronOnoff) Value() *PoweronOnoff       { return a }
func (a PoweronOnoff) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *PoweronOnoff) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PoweronOnoff(*nt)
	return br, err
}

func (a PoweronOnoff) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x01:
		return "On"
	case 0xFF:
		return "Previous"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
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
