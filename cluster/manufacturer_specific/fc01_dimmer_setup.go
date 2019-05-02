// Attributes and commands.
package manufacturer_specific

import (
	"neotor.se/zcl"
)

// UbisysDimmerSetup
// Manufacturer code 0x10F2
const UbisysDimmerSetupID zcl.ClusterID = 64513

var UbisysDimmerSetupCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		CapabilitiesAttr: func() zcl.Attr { return new(Capabilities) },
		StatusAttr:       func() zcl.Attr { return new(Status) },
		ModeAttr:         func() zcl.Attr { return new(Mode) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const CapabilitiesAttr zcl.AttrID = 0

type Capabilities zcl.Zbmp8

func (a Capabilities) ID() zcl.AttrID         { return CapabilitiesAttr }
func (a Capabilities) Cluster() zcl.ClusterID { return UbisysDimmerSetupID }
func (a *Capabilities) Value() *Capabilities  { return a }
func (a Capabilities) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *Capabilities) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Capabilities(*nt)
	return br, err
}

func (a Capabilities) Readable() bool   { return true }
func (a Capabilities) Writable() bool   { return false }
func (a Capabilities) Reportable() bool { return false }
func (a Capabilities) SceneIndex() int  { return -1 }

func (a Capabilities) String() string {

	var bstr []string
	if a.IsForwardPhaseControl() {
		bstr = append(bstr, "Forward Phase Control")
	}
	if a.IsReversePhaseControl() {
		bstr = append(bstr, "Reverse Phase Control")
	}
	if a.IsReactanceDiscriminator() {
		bstr = append(bstr, "Reactance Discriminator")
	}
	if a.IsConfigurableCurve() {
		bstr = append(bstr, "Configurable Curve")
	}
	if a.IsOverloadDetection() {
		bstr = append(bstr, "Overload detection")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a Capabilities) IsForwardPhaseControl() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *Capabilities) SetForwardPhaseControl(b bool) {
	*a = Capabilities(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a Capabilities) IsReversePhaseControl() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *Capabilities) SetReversePhaseControl(b bool) {
	*a = Capabilities(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a Capabilities) IsReactanceDiscriminator() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *Capabilities) SetReactanceDiscriminator(b bool) {
	*a = Capabilities(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a Capabilities) IsConfigurableCurve() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *Capabilities) SetConfigurableCurve(b bool) {
	*a = Capabilities(zcl.BitmapSet([]byte(*a), 6, b))
}

func (a Capabilities) IsOverloadDetection() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *Capabilities) SetOverloadDetection(b bool) {
	*a = Capabilities(zcl.BitmapSet([]byte(*a), 7, b))
}

const StatusAttr zcl.AttrID = 1

type Status zcl.Zbmp8

func (a Status) ID() zcl.AttrID         { return StatusAttr }
func (a Status) Cluster() zcl.ClusterID { return UbisysDimmerSetupID }
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
func (a Status) Writable() bool   { return false }
func (a Status) Reportable() bool { return false }
func (a Status) SceneIndex() int  { return -1 }

func (a Status) String() string {

	var bstr []string
	if a.IsForwardPhaseControl() {
		bstr = append(bstr, "Forward Phase Control")
	}
	if a.IsReversePhaseControl() {
		bstr = append(bstr, "Reverse Phase Control")
	}
	if a.IsOperational() {
		bstr = append(bstr, "Operational")
	}
	if a.IsOverload() {
		bstr = append(bstr, "Overload")
	}
	if a.IsCapacitiveLoad() {
		bstr = append(bstr, "Capacitive Load")
	}
	if a.IsInductiveLoad() {
		bstr = append(bstr, "Inductive Load")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a Status) IsForwardPhaseControl() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *Status) SetForwardPhaseControl(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a Status) IsReversePhaseControl() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *Status) SetReversePhaseControl(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a Status) IsOperational() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *Status) SetOperational(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a Status) IsOverload() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *Status) SetOverload(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a Status) IsCapacitiveLoad() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *Status) SetCapacitiveLoad(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 6, b))
}

func (a Status) IsInductiveLoad() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *Status) SetInductiveLoad(b bool) {
	*a = Status(zcl.BitmapSet([]byte(*a), 7, b))
}

const ModeAttr zcl.AttrID = 2

type Mode zcl.Zbmp8

func (a Mode) ID() zcl.AttrID         { return ModeAttr }
func (a Mode) Cluster() zcl.ClusterID { return UbisysDimmerSetupID }
func (a *Mode) Value() *Mode          { return a }
func (a Mode) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *Mode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Mode(*nt)
	return br, err
}

func (a Mode) Readable() bool   { return true }
func (a Mode) Writable() bool   { return false }
func (a Mode) Reportable() bool { return false }
func (a Mode) SceneIndex() int  { return -1 }

func (a Mode) String() string {

	var bstr []string
	if a.IsAutomaticPhaseControl() {
		bstr = append(bstr, "Automatic Phase Control")
	}
	if a.IsForwardPhaseControl() {
		bstr = append(bstr, "Forward Phase Control")
	}
	if a.IsReversePhaseControl() {
		bstr = append(bstr, "Reverse Phase Control")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a Mode) IsAutomaticPhaseControl() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *Mode) SetAutomaticPhaseControl(b bool) {
	*a = Mode(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a Mode) IsForwardPhaseControl() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *Mode) SetForwardPhaseControl(b bool) {
	*a = Mode(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a Mode) IsReversePhaseControl() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *Mode) SetReversePhaseControl(b bool) {
	*a = Mode(zcl.BitmapSet([]byte(*a), 2, b))
}