// This cluster specifies an interface to control the speed of a fan as part of a heating / cooling system.
package hvac

import (
	"neotor.se/zcl"
)

// FanControl
const FanControlID zcl.ClusterID = 514

var FanControlCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		FanModeAttr:         func() zcl.Attr { return new(FanMode) },
		FanModeSequenceAttr: func() zcl.Attr { return new(FanModeSequence) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const FanModeAttr zcl.AttrID = 0

type FanMode zcl.Zenum8

func (a FanMode) ID() zcl.AttrID         { return FanModeAttr }
func (a FanMode) Cluster() zcl.ClusterID { return FanControlID }
func (a *FanMode) Value() *FanMode       { return a }
func (a FanMode) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *FanMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = FanMode(*nt)
	return br, err
}

func (a FanMode) Readable() bool   { return true }
func (a FanMode) Writable() bool   { return true }
func (a FanMode) Reportable() bool { return false }
func (a FanMode) SceneIndex() int  { return -1 }

func (a FanMode) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x01:
		return "Low"
	case 0x02:
		return "Medium"
	case 0x03:
		return "High"
	case 0x04:
		return "On"
	case 0x05:
		return "Auto"
	case 0x06:
		return "Smart"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsOff checks if FanMode equals the value for Off (0x00)
func (a FanMode) IsOff() bool { return a == 0x00 }

// SetOff sets FanMode to Off (0x00)
func (a *FanMode) SetOff() { *a = 0x00 }

// IsLow checks if FanMode equals the value for Low (0x01)
func (a FanMode) IsLow() bool { return a == 0x01 }

// SetLow sets FanMode to Low (0x01)
func (a *FanMode) SetLow() { *a = 0x01 }

// IsMedium checks if FanMode equals the value for Medium (0x02)
func (a FanMode) IsMedium() bool { return a == 0x02 }

// SetMedium sets FanMode to Medium (0x02)
func (a *FanMode) SetMedium() { *a = 0x02 }

// IsHigh checks if FanMode equals the value for High (0x03)
func (a FanMode) IsHigh() bool { return a == 0x03 }

// SetHigh sets FanMode to High (0x03)
func (a *FanMode) SetHigh() { *a = 0x03 }

// IsOn checks if FanMode equals the value for On (0x04)
func (a FanMode) IsOn() bool { return a == 0x04 }

// SetOn sets FanMode to On (0x04)
func (a *FanMode) SetOn() { *a = 0x04 }

// IsAuto checks if FanMode equals the value for Auto (0x05)
func (a FanMode) IsAuto() bool { return a == 0x05 }

// SetAuto sets FanMode to Auto (0x05)
func (a *FanMode) SetAuto() { *a = 0x05 }

// IsSmart checks if FanMode equals the value for Smart (0x06)
func (a FanMode) IsSmart() bool { return a == 0x06 }

// SetSmart sets FanMode to Smart (0x06)
func (a *FanMode) SetSmart() { *a = 0x06 }

const FanModeSequenceAttr zcl.AttrID = 1

type FanModeSequence zcl.Zenum8

func (a FanModeSequence) ID() zcl.AttrID           { return FanModeSequenceAttr }
func (a FanModeSequence) Cluster() zcl.ClusterID   { return FanControlID }
func (a *FanModeSequence) Value() *FanModeSequence { return a }
func (a FanModeSequence) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *FanModeSequence) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = FanModeSequence(*nt)
	return br, err
}

func (a FanModeSequence) Readable() bool   { return true }
func (a FanModeSequence) Writable() bool   { return true }
func (a FanModeSequence) Reportable() bool { return false }
func (a FanModeSequence) SceneIndex() int  { return -1 }

func (a FanModeSequence) String() string {
	switch a {
	case 0x00:
		return "Low/Med/High"
	case 0x01:
		return "Low/High"
	case 0x02:
		return "Low/Med/High/Auto"
	case 0x03:
		return "Low/High/Auto"
	case 0x04:
		return "On/Auto"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsLowMedHigh checks if FanModeSequence equals the value for Low/Med/High (0x00)
func (a FanModeSequence) IsLowMedHigh() bool { return a == 0x00 }

// SetLowMedHigh sets FanModeSequence to Low/Med/High (0x00)
func (a *FanModeSequence) SetLowMedHigh() { *a = 0x00 }

// IsLowHigh checks if FanModeSequence equals the value for Low/High (0x01)
func (a FanModeSequence) IsLowHigh() bool { return a == 0x01 }

// SetLowHigh sets FanModeSequence to Low/High (0x01)
func (a *FanModeSequence) SetLowHigh() { *a = 0x01 }

// IsLowMedHighAuto checks if FanModeSequence equals the value for Low/Med/High/Auto (0x02)
func (a FanModeSequence) IsLowMedHighAuto() bool { return a == 0x02 }

// SetLowMedHighAuto sets FanModeSequence to Low/Med/High/Auto (0x02)
func (a *FanModeSequence) SetLowMedHighAuto() { *a = 0x02 }

// IsLowHighAuto checks if FanModeSequence equals the value for Low/High/Auto (0x03)
func (a FanModeSequence) IsLowHighAuto() bool { return a == 0x03 }

// SetLowHighAuto sets FanModeSequence to Low/High/Auto (0x03)
func (a *FanModeSequence) SetLowHighAuto() { *a = 0x03 }

// IsOnAuto checks if FanModeSequence equals the value for On/Auto (0x04)
func (a FanModeSequence) IsOnAuto() bool { return a == 0x04 }

// SetOnAuto sets FanModeSequence to On/Auto (0x04)
func (a *FanModeSequence) SetOnAuto() { *a = 0x04 }
