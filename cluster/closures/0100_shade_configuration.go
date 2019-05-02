// The shade configuration cluster provides an interface for reading information about a shade, and configuring its open and closed limits.
package closures

import (
	"neotor.se/zcl"
)

// ShadeConfiguration
const ShadeConfigurationID zcl.ClusterID = 256

var ShadeConfigurationCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		PhysicalClosedLimitAttr: func() zcl.Attr { return new(PhysicalClosedLimit) },
		MotorstepsizeAttr:       func() zcl.Attr { return new(Motorstepsize) },
		StatusAttr:              func() zcl.Attr { return new(Status) },
		ClosedLimitAttr:         func() zcl.Attr { return new(ClosedLimit) },
		ModeAttr:                func() zcl.Attr { return new(Mode) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const PhysicalClosedLimitAttr zcl.AttrID = 0

type PhysicalClosedLimit zcl.Zu16

func (a PhysicalClosedLimit) ID() zcl.AttrID               { return PhysicalClosedLimitAttr }
func (a PhysicalClosedLimit) Cluster() zcl.ClusterID       { return ShadeConfigurationID }
func (a *PhysicalClosedLimit) Value() *PhysicalClosedLimit { return a }
func (a PhysicalClosedLimit) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PhysicalClosedLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PhysicalClosedLimit(*nt)
	return br, err
}

func (a PhysicalClosedLimit) Readable() bool   { return true }
func (a PhysicalClosedLimit) Writable() bool   { return false }
func (a PhysicalClosedLimit) Reportable() bool { return false }
func (a PhysicalClosedLimit) SceneIndex() int  { return -1 }

func (a PhysicalClosedLimit) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MotorstepsizeAttr zcl.AttrID = 1

type Motorstepsize zcl.Zu8

func (a Motorstepsize) ID() zcl.AttrID         { return MotorstepsizeAttr }
func (a Motorstepsize) Cluster() zcl.ClusterID { return ShadeConfigurationID }
func (a *Motorstepsize) Value() *Motorstepsize { return a }
func (a Motorstepsize) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Motorstepsize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Motorstepsize(*nt)
	return br, err
}

func (a Motorstepsize) Readable() bool   { return true }
func (a Motorstepsize) Writable() bool   { return false }
func (a Motorstepsize) Reportable() bool { return false }
func (a Motorstepsize) SceneIndex() int  { return -1 }

func (a Motorstepsize) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const StatusAttr zcl.AttrID = 2

type Status zcl.Zbmp8

func (a Status) ID() zcl.AttrID         { return StatusAttr }
func (a Status) Cluster() zcl.ClusterID { return ShadeConfigurationID }
func (a *Status) Value() *Status        { return a }
func (a Status) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *Status) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Status(*nt)
	return br, err
}

func (a Status) Readable() bool   { return true }
func (a Status) Writable() bool   { return true }
func (a Status) Reportable() bool { return false }
func (a Status) SceneIndex() int  { return -1 }

func (a Status) String() string {

	var bstr []string
	if a.IsShadeOperational() {
		bstr = append(bstr, "Shade operational")
	}
	if a.IsShadeAdjusting() {
		bstr = append(bstr, "Shade adjusting")
	}
	if a.IsShadeDirection() {
		bstr = append(bstr, "Shade direction")
	}
	if a.IsDirectionReversed() {
		bstr = append(bstr, "Direction Reversed")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a Status) IsShadeOperational() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *Status) SetShadeOperational(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a Status) IsShadeAdjusting() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *Status) SetShadeAdjusting(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a Status) IsShadeDirection() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *Status) SetShadeDirection(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a Status) IsDirectionReversed() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *Status) SetDirectionReversed(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 3, b))
}

const ClosedLimitAttr zcl.AttrID = 16

type ClosedLimit zcl.Zu16

func (a ClosedLimit) ID() zcl.AttrID         { return ClosedLimitAttr }
func (a ClosedLimit) Cluster() zcl.ClusterID { return ShadeConfigurationID }
func (a *ClosedLimit) Value() *ClosedLimit   { return a }
func (a ClosedLimit) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ClosedLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ClosedLimit(*nt)
	return br, err
}

func (a ClosedLimit) Readable() bool   { return true }
func (a ClosedLimit) Writable() bool   { return true }
func (a ClosedLimit) Reportable() bool { return false }
func (a ClosedLimit) SceneIndex() int  { return -1 }

func (a ClosedLimit) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ModeAttr zcl.AttrID = 17

type Mode zcl.Zenum8

func (a Mode) ID() zcl.AttrID         { return ModeAttr }
func (a Mode) Cluster() zcl.ClusterID { return ShadeConfigurationID }
func (a *Mode) Value() *Mode          { return a }
func (a Mode) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *Mode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Mode(*nt)
	return br, err
}

func (a Mode) Readable() bool   { return true }
func (a Mode) Writable() bool   { return true }
func (a Mode) Reportable() bool { return false }
func (a Mode) SceneIndex() int  { return -1 }

func (a Mode) String() string {
	switch a {
	case 0x00:
		return "Normal"
	case 0x01:
		return "Configure"
	case 0xFF:
		return "Invalid"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsNormal checks if Mode equals the value for Normal (0x00)
func (a Mode) IsNormal() bool { return a == 0x00 }

// SetNormal sets Mode to Normal (0x00)
func (a *Mode) SetNormal() { *a = 0x00 }

// IsConfigure checks if Mode equals the value for Configure (0x01)
func (a Mode) IsConfigure() bool { return a == 0x01 }

// SetConfigure sets Mode to Configure (0x01)
func (a *Mode) SetConfigure() { *a = 0x01 }

// IsInvalid checks if Mode equals the value for Invalid (0xFF)
func (a Mode) IsInvalid() bool { return a == 0xFF }

// SetInvalid sets Mode to Invalid (0xFF)
func (a *Mode) SetInvalid() { *a = 0xFF }