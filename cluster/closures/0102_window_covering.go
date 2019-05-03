// The window covering cluster provides an interface for controlling and adjusting automatic window coverings such as drapery motors, automatic shades, and blinds.
package closures

import (
	"neotor.se/zcl"
)

// WindowCovering
const WindowCoveringID zcl.ClusterID = 258

var WindowCoveringCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		UpOpenCommand:             func() zcl.Command { return new(UpOpen) },
		DownCloseCommand:          func() zcl.Command { return new(DownClose) },
		StopCommand:               func() zcl.Command { return new(Stop) },
		GoToLiftValueCommand:      func() zcl.Command { return new(GoToLiftValue) },
		GoToLiftPercentageCommand: func() zcl.Command { return new(GoToLiftPercentage) },
		GoToTiltValueCommand:      func() zcl.Command { return new(GoToTiltValue) },
		GoToTiltPercentageCommand: func() zcl.Command { return new(GoToTiltPercentage) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		WindowCoveringTypeAttr:            func() zcl.Attr { return new(WindowCoveringType) },
		PhysicalClosedLimitLiftAttr:       func() zcl.Attr { return new(PhysicalClosedLimitLift) },
		PhysicalClosedLimitTiltAttr:       func() zcl.Attr { return new(PhysicalClosedLimitTilt) },
		CurrentPositionLiftAttr:           func() zcl.Attr { return new(CurrentPositionLift) },
		CurrentPositionTiltAttr:           func() zcl.Attr { return new(CurrentPositionTilt) },
		NumberOfActuationsLiftAttr:        func() zcl.Attr { return new(NumberOfActuationsLift) },
		NumberOfActuationsTiltAttr:        func() zcl.Attr { return new(NumberOfActuationsTilt) },
		ConfigStatusAttr:                  func() zcl.Attr { return new(ConfigStatus) },
		CurrentPositionLiftPercentageAttr: func() zcl.Attr { return new(CurrentPositionLiftPercentage) },
		CurrentPositionTiltPercentageAttr: func() zcl.Attr { return new(CurrentPositionTiltPercentage) },
		InstalledOpenLimitLiftAttr:        func() zcl.Attr { return new(InstalledOpenLimitLift) },
		InstalledClosedLimitLiftAttr:      func() zcl.Attr { return new(InstalledClosedLimitLift) },
		InstalledOpenLimitTiltAAttr:       func() zcl.Attr { return new(InstalledOpenLimitTiltA) },
		InstalledOpenLimitTiltBAttr:       func() zcl.Attr { return new(InstalledOpenLimitTiltB) },
		VelocityLiftAttr:                  func() zcl.Attr { return new(VelocityLift) },
		AccelerationTimeLiftAttr:          func() zcl.Attr { return new(AccelerationTimeLift) },
		DecelerationTimeLiftAttr:          func() zcl.Attr { return new(DecelerationTimeLift) },
		WindowCoveringModeAttr:            func() zcl.Attr { return new(WindowCoveringMode) },
		IntermediateSetpointsLiftAttr:     func() zcl.Attr { return new(IntermediateSetpointsLift) },
		IntermediateSetpointsTiltAttr:     func() zcl.Attr { return new(IntermediateSetpointsTilt) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type UpOpen struct {
}

const UpOpenCommand zcl.CommandID = 0

func (v *UpOpen) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v UpOpen) ID() zcl.CommandID {
	return UpOpenCommand
}

func (v UpOpen) Cluster() zcl.ClusterID {
	return WindowCoveringID
}

func (v UpOpen) MnfCode() []byte {
	return []byte{}
}

func (v UpOpen) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *UpOpen) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

func (v UpOpen) String() string {
	var str []string
	return "UpOpen{" + zcl.StrJoin(str, " ") + "}"
}

func (UpOpen) Name() string { return "Up / Open" }

type DownClose struct {
}

const DownCloseCommand zcl.CommandID = 1

func (v *DownClose) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v DownClose) ID() zcl.CommandID {
	return DownCloseCommand
}

func (v DownClose) Cluster() zcl.ClusterID {
	return WindowCoveringID
}

func (v DownClose) MnfCode() []byte {
	return []byte{}
}

func (v DownClose) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *DownClose) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

func (v DownClose) String() string {
	var str []string
	return "DownClose{" + zcl.StrJoin(str, " ") + "}"
}

func (DownClose) Name() string { return "Down / Close" }

type Stop struct {
}

const StopCommand zcl.CommandID = 2

func (v *Stop) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v Stop) ID() zcl.CommandID {
	return StopCommand
}

func (v Stop) Cluster() zcl.ClusterID {
	return WindowCoveringID
}

func (v Stop) MnfCode() []byte {
	return []byte{}
}

func (v Stop) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *Stop) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

func (v Stop) String() string {
	var str []string
	return "Stop{" + zcl.StrJoin(str, " ") + "}"
}

func (Stop) Name() string { return "Stop" }

type GoToLiftValue struct {
	LiftValue zcl.Zu16
}

const GoToLiftValueCommand zcl.CommandID = 4

func (v *GoToLiftValue) Values() []zcl.Val {
	return []zcl.Val{
		&v.LiftValue,
	}
}

func (v GoToLiftValue) ID() zcl.CommandID {
	return GoToLiftValueCommand
}

func (v GoToLiftValue) Cluster() zcl.ClusterID {
	return WindowCoveringID
}

func (v GoToLiftValue) MnfCode() []byte {
	return []byte{}
}

func (v GoToLiftValue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.LiftValue.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GoToLiftValue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.LiftValue).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v GoToLiftValue) LiftValueString() string {
	return zcl.Sprintf("%v", zcl.Zu16(v.LiftValue))
}

func (v GoToLiftValue) String() string {
	var str []string
	str = append(str, "LiftValue["+v.LiftValueString()+"]")
	return "GoToLiftValue{" + zcl.StrJoin(str, " ") + "}"
}

func (GoToLiftValue) Name() string { return "Go To Lift Value" }

type GoToLiftPercentage struct {
	LiftPercentage zcl.Zu8
}

const GoToLiftPercentageCommand zcl.CommandID = 5

func (v *GoToLiftPercentage) Values() []zcl.Val {
	return []zcl.Val{
		&v.LiftPercentage,
	}
}

func (v GoToLiftPercentage) ID() zcl.CommandID {
	return GoToLiftPercentageCommand
}

func (v GoToLiftPercentage) Cluster() zcl.ClusterID {
	return WindowCoveringID
}

func (v GoToLiftPercentage) MnfCode() []byte {
	return []byte{}
}

func (v GoToLiftPercentage) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.LiftPercentage.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GoToLiftPercentage) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.LiftPercentage).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v GoToLiftPercentage) LiftPercentageString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.LiftPercentage))
}

func (v GoToLiftPercentage) String() string {
	var str []string
	str = append(str, "LiftPercentage["+v.LiftPercentageString()+"]")
	return "GoToLiftPercentage{" + zcl.StrJoin(str, " ") + "}"
}

func (GoToLiftPercentage) Name() string { return "Go to Lift Percentage" }

type GoToTiltValue struct {
	TiltValue zcl.Zu16
}

const GoToTiltValueCommand zcl.CommandID = 7

func (v *GoToTiltValue) Values() []zcl.Val {
	return []zcl.Val{
		&v.TiltValue,
	}
}

func (v GoToTiltValue) ID() zcl.CommandID {
	return GoToTiltValueCommand
}

func (v GoToTiltValue) Cluster() zcl.ClusterID {
	return WindowCoveringID
}

func (v GoToTiltValue) MnfCode() []byte {
	return []byte{}
}

func (v GoToTiltValue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.TiltValue.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GoToTiltValue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.TiltValue).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v GoToTiltValue) TiltValueString() string {
	return zcl.Sprintf("%v", zcl.Zu16(v.TiltValue))
}

func (v GoToTiltValue) String() string {
	var str []string
	str = append(str, "TiltValue["+v.TiltValueString()+"]")
	return "GoToTiltValue{" + zcl.StrJoin(str, " ") + "}"
}

func (GoToTiltValue) Name() string { return "Go to Tilt Value" }

type GoToTiltPercentage struct {
	TiltPercentage zcl.Zu8
}

const GoToTiltPercentageCommand zcl.CommandID = 8

func (v *GoToTiltPercentage) Values() []zcl.Val {
	return []zcl.Val{
		&v.TiltPercentage,
	}
}

func (v GoToTiltPercentage) ID() zcl.CommandID {
	return GoToTiltPercentageCommand
}

func (v GoToTiltPercentage) Cluster() zcl.ClusterID {
	return WindowCoveringID
}

func (v GoToTiltPercentage) MnfCode() []byte {
	return []byte{}
}

func (v GoToTiltPercentage) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.TiltPercentage.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GoToTiltPercentage) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.TiltPercentage).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v GoToTiltPercentage) TiltPercentageString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.TiltPercentage))
}

func (v GoToTiltPercentage) String() string {
	var str []string
	str = append(str, "TiltPercentage["+v.TiltPercentageString()+"]")
	return "GoToTiltPercentage{" + zcl.StrJoin(str, " ") + "}"
}

func (GoToTiltPercentage) Name() string { return "Go to Tilt Percentage" }

// WindowCoveringType is an autogenerated attribute in the WindowCovering cluster
type WindowCoveringType zcl.Zenum8

const WindowCoveringTypeAttr zcl.AttrID = 0

func (a WindowCoveringType) ID() zcl.AttrID              { return WindowCoveringTypeAttr }
func (a WindowCoveringType) Cluster() zcl.ClusterID      { return WindowCoveringID }
func (a *WindowCoveringType) Value() *WindowCoveringType { return a }
func (a WindowCoveringType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *WindowCoveringType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = WindowCoveringType(*nt)
	return br, err
}
func (WindowCoveringType) Name() string     { return "Window Covering Type" }
func (WindowCoveringType) Readable() bool   { return true }
func (WindowCoveringType) Writable() bool   { return false }
func (WindowCoveringType) Reportable() bool { return false }
func (WindowCoveringType) SceneIndex() int  { return -1 }

func (a WindowCoveringType) String() string {
	switch a {
	case 0x00:
		return "Rollershade"
	case 0x01:
		return "Rollershade - 2 Motor"
	case 0x02:
		return "Rollershade - Exterior"
	case 0x03:
		return "Rollershade - Exterior - 2 Motor"
	case 0x04:
		return "Drapery"
	case 0x05:
		return "Awning"
	case 0x06:
		return "Shutter"
	case 0x07:
		return "Tilt Blind - Tilt Only"
	case 0x08:
		return "Tilt Blind - Lift and Tilt"
	case 0x09:
		return "Projector Screen"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

// IsRollershade checks if WindowCoveringType equals the value for Rollershade (0x00)
func (a WindowCoveringType) IsRollershade() bool { return a == 0x00 }

// SetRollershade sets WindowCoveringType to Rollershade (0x00)
func (a *WindowCoveringType) SetRollershade() { *a = 0x00 }

// IsRollershade2Motor checks if WindowCoveringType equals the value for Rollershade - 2 Motor (0x01)
func (a WindowCoveringType) IsRollershade2Motor() bool { return a == 0x01 }

// SetRollershade2Motor sets WindowCoveringType to Rollershade - 2 Motor (0x01)
func (a *WindowCoveringType) SetRollershade2Motor() { *a = 0x01 }

// IsRollershadeExterior checks if WindowCoveringType equals the value for Rollershade - Exterior (0x02)
func (a WindowCoveringType) IsRollershadeExterior() bool { return a == 0x02 }

// SetRollershadeExterior sets WindowCoveringType to Rollershade - Exterior (0x02)
func (a *WindowCoveringType) SetRollershadeExterior() { *a = 0x02 }

// IsRollershadeExterior2Motor checks if WindowCoveringType equals the value for Rollershade - Exterior - 2 Motor (0x03)
func (a WindowCoveringType) IsRollershadeExterior2Motor() bool { return a == 0x03 }

// SetRollershadeExterior2Motor sets WindowCoveringType to Rollershade - Exterior - 2 Motor (0x03)
func (a *WindowCoveringType) SetRollershadeExterior2Motor() { *a = 0x03 }

// IsDrapery checks if WindowCoveringType equals the value for Drapery (0x04)
func (a WindowCoveringType) IsDrapery() bool { return a == 0x04 }

// SetDrapery sets WindowCoveringType to Drapery (0x04)
func (a *WindowCoveringType) SetDrapery() { *a = 0x04 }

// IsAwning checks if WindowCoveringType equals the value for Awning (0x05)
func (a WindowCoveringType) IsAwning() bool { return a == 0x05 }

// SetAwning sets WindowCoveringType to Awning (0x05)
func (a *WindowCoveringType) SetAwning() { *a = 0x05 }

// IsShutter checks if WindowCoveringType equals the value for Shutter (0x06)
func (a WindowCoveringType) IsShutter() bool { return a == 0x06 }

// SetShutter sets WindowCoveringType to Shutter (0x06)
func (a *WindowCoveringType) SetShutter() { *a = 0x06 }

// IsTiltBlindTiltOnly checks if WindowCoveringType equals the value for Tilt Blind - Tilt Only (0x07)
func (a WindowCoveringType) IsTiltBlindTiltOnly() bool { return a == 0x07 }

// SetTiltBlindTiltOnly sets WindowCoveringType to Tilt Blind - Tilt Only (0x07)
func (a *WindowCoveringType) SetTiltBlindTiltOnly() { *a = 0x07 }

// IsTiltBlindLiftAndTilt checks if WindowCoveringType equals the value for Tilt Blind - Lift and Tilt (0x08)
func (a WindowCoveringType) IsTiltBlindLiftAndTilt() bool { return a == 0x08 }

// SetTiltBlindLiftAndTilt sets WindowCoveringType to Tilt Blind - Lift and Tilt (0x08)
func (a *WindowCoveringType) SetTiltBlindLiftAndTilt() { *a = 0x08 }

// IsProjectorScreen checks if WindowCoveringType equals the value for Projector Screen (0x09)
func (a WindowCoveringType) IsProjectorScreen() bool { return a == 0x09 }

// SetProjectorScreen sets WindowCoveringType to Projector Screen (0x09)
func (a *WindowCoveringType) SetProjectorScreen() { *a = 0x09 }

// PhysicalClosedLimitLift is an autogenerated attribute in the WindowCovering cluster
type PhysicalClosedLimitLift zcl.Zu16

const PhysicalClosedLimitLiftAttr zcl.AttrID = 1

func (a PhysicalClosedLimitLift) ID() zcl.AttrID                   { return PhysicalClosedLimitLiftAttr }
func (a PhysicalClosedLimitLift) Cluster() zcl.ClusterID           { return WindowCoveringID }
func (a *PhysicalClosedLimitLift) Value() *PhysicalClosedLimitLift { return a }
func (a PhysicalClosedLimitLift) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PhysicalClosedLimitLift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PhysicalClosedLimitLift(*nt)
	return br, err
}
func (PhysicalClosedLimitLift) Name() string     { return "Physical Closed Limit - Lift" }
func (PhysicalClosedLimitLift) Readable() bool   { return true }
func (PhysicalClosedLimitLift) Writable() bool   { return false }
func (PhysicalClosedLimitLift) Reportable() bool { return false }
func (PhysicalClosedLimitLift) SceneIndex() int  { return -1 }

func (a PhysicalClosedLimitLift) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// PhysicalClosedLimitTilt is an autogenerated attribute in the WindowCovering cluster
type PhysicalClosedLimitTilt zcl.Zu16

const PhysicalClosedLimitTiltAttr zcl.AttrID = 2

func (a PhysicalClosedLimitTilt) ID() zcl.AttrID                   { return PhysicalClosedLimitTiltAttr }
func (a PhysicalClosedLimitTilt) Cluster() zcl.ClusterID           { return WindowCoveringID }
func (a *PhysicalClosedLimitTilt) Value() *PhysicalClosedLimitTilt { return a }
func (a PhysicalClosedLimitTilt) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PhysicalClosedLimitTilt) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PhysicalClosedLimitTilt(*nt)
	return br, err
}
func (PhysicalClosedLimitTilt) Name() string     { return "Physical Closed Limit - Tilt" }
func (PhysicalClosedLimitTilt) Readable() bool   { return true }
func (PhysicalClosedLimitTilt) Writable() bool   { return false }
func (PhysicalClosedLimitTilt) Reportable() bool { return false }
func (PhysicalClosedLimitTilt) SceneIndex() int  { return -1 }

func (a PhysicalClosedLimitTilt) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// CurrentPositionLift is an autogenerated attribute in the WindowCovering cluster
type CurrentPositionLift zcl.Zu16

const CurrentPositionLiftAttr zcl.AttrID = 3

func (a CurrentPositionLift) ID() zcl.AttrID               { return CurrentPositionLiftAttr }
func (a CurrentPositionLift) Cluster() zcl.ClusterID       { return WindowCoveringID }
func (a *CurrentPositionLift) Value() *CurrentPositionLift { return a }
func (a CurrentPositionLift) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *CurrentPositionLift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentPositionLift(*nt)
	return br, err
}
func (CurrentPositionLift) Name() string     { return "Current Position - Lift" }
func (CurrentPositionLift) Readable() bool   { return true }
func (CurrentPositionLift) Writable() bool   { return false }
func (CurrentPositionLift) Reportable() bool { return false }
func (CurrentPositionLift) SceneIndex() int  { return -1 }

func (a CurrentPositionLift) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// CurrentPositionTilt is an autogenerated attribute in the WindowCovering cluster
type CurrentPositionTilt zcl.Zu16

const CurrentPositionTiltAttr zcl.AttrID = 4

func (a CurrentPositionTilt) ID() zcl.AttrID               { return CurrentPositionTiltAttr }
func (a CurrentPositionTilt) Cluster() zcl.ClusterID       { return WindowCoveringID }
func (a *CurrentPositionTilt) Value() *CurrentPositionTilt { return a }
func (a CurrentPositionTilt) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *CurrentPositionTilt) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentPositionTilt(*nt)
	return br, err
}
func (CurrentPositionTilt) Name() string     { return "Current Position - Tilt" }
func (CurrentPositionTilt) Readable() bool   { return true }
func (CurrentPositionTilt) Writable() bool   { return false }
func (CurrentPositionTilt) Reportable() bool { return false }
func (CurrentPositionTilt) SceneIndex() int  { return -1 }

func (a CurrentPositionTilt) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// NumberOfActuationsLift is an autogenerated attribute in the WindowCovering cluster
type NumberOfActuationsLift zcl.Zu16

const NumberOfActuationsLiftAttr zcl.AttrID = 5

func (a NumberOfActuationsLift) ID() zcl.AttrID                  { return NumberOfActuationsLiftAttr }
func (a NumberOfActuationsLift) Cluster() zcl.ClusterID          { return WindowCoveringID }
func (a *NumberOfActuationsLift) Value() *NumberOfActuationsLift { return a }
func (a NumberOfActuationsLift) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NumberOfActuationsLift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfActuationsLift(*nt)
	return br, err
}
func (NumberOfActuationsLift) Name() string     { return "Number of Actuations - Lift" }
func (NumberOfActuationsLift) Readable() bool   { return true }
func (NumberOfActuationsLift) Writable() bool   { return false }
func (NumberOfActuationsLift) Reportable() bool { return false }
func (NumberOfActuationsLift) SceneIndex() int  { return -1 }

func (a NumberOfActuationsLift) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// NumberOfActuationsTilt is an autogenerated attribute in the WindowCovering cluster
type NumberOfActuationsTilt zcl.Zu16

const NumberOfActuationsTiltAttr zcl.AttrID = 6

func (a NumberOfActuationsTilt) ID() zcl.AttrID                  { return NumberOfActuationsTiltAttr }
func (a NumberOfActuationsTilt) Cluster() zcl.ClusterID          { return WindowCoveringID }
func (a *NumberOfActuationsTilt) Value() *NumberOfActuationsTilt { return a }
func (a NumberOfActuationsTilt) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NumberOfActuationsTilt) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfActuationsTilt(*nt)
	return br, err
}
func (NumberOfActuationsTilt) Name() string     { return "Number of Actuations - Tilt" }
func (NumberOfActuationsTilt) Readable() bool   { return true }
func (NumberOfActuationsTilt) Writable() bool   { return false }
func (NumberOfActuationsTilt) Reportable() bool { return false }
func (NumberOfActuationsTilt) SceneIndex() int  { return -1 }

func (a NumberOfActuationsTilt) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ConfigStatus is an autogenerated attribute in the WindowCovering cluster
type ConfigStatus zcl.Zbmp8

const ConfigStatusAttr zcl.AttrID = 7

func (a ConfigStatus) ID() zcl.AttrID         { return ConfigStatusAttr }
func (a ConfigStatus) Cluster() zcl.ClusterID { return WindowCoveringID }
func (a *ConfigStatus) Value() *ConfigStatus  { return a }
func (a ConfigStatus) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *ConfigStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = ConfigStatus(*nt)
	return br, err
}
func (ConfigStatus) Name() string     { return "Config / Status" }
func (ConfigStatus) Readable() bool   { return true }
func (ConfigStatus) Writable() bool   { return false }
func (ConfigStatus) Reportable() bool { return false }
func (ConfigStatus) SceneIndex() int  { return -1 }

func (a ConfigStatus) String() string {
	var bstr []string
	if a.IsOperational() {
		bstr = append(bstr, "Operational")
	}
	if a.IsOnline() {
		bstr = append(bstr, "Online")
	}
	if a.IsCommandsReversed() {
		bstr = append(bstr, "Commands Reversed")
	}
	if a.IsLiftControlIsClosedLoop() {
		bstr = append(bstr, "Lift control is Closed Loop")
	}
	if a.IsTiltControlIsClosedLoop() {
		bstr = append(bstr, "Tilt control is Closed Loop")
	}
	if a.IsLiftEncoderControlled() {
		bstr = append(bstr, "Lift: Encoder Controlled")
	}
	if a.IsTiltEncoderControlled() {
		bstr = append(bstr, "Tilt: Encoder Controlled")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a ConfigStatus) IsOperational() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *ConfigStatus) SetOperational(b bool) {
	*a = ConfigStatus(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a ConfigStatus) IsOnline() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *ConfigStatus) SetOnline(b bool) {
	*a = ConfigStatus(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a ConfigStatus) IsCommandsReversed() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *ConfigStatus) SetCommandsReversed(b bool) {
	*a = ConfigStatus(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a ConfigStatus) IsLiftControlIsClosedLoop() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *ConfigStatus) SetLiftControlIsClosedLoop(b bool) {
	*a = ConfigStatus(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a ConfigStatus) IsTiltControlIsClosedLoop() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *ConfigStatus) SetTiltControlIsClosedLoop(b bool) {
	*a = ConfigStatus(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a ConfigStatus) IsLiftEncoderControlled() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *ConfigStatus) SetLiftEncoderControlled(b bool) {
	*a = ConfigStatus(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a ConfigStatus) IsTiltEncoderControlled() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *ConfigStatus) SetTiltEncoderControlled(b bool) {
	*a = ConfigStatus(zcl.BitmapSet([]byte(*a), 6, b))
}

// CurrentPositionLiftPercentage is an autogenerated attribute in the WindowCovering cluster
type CurrentPositionLiftPercentage zcl.Zu8

const CurrentPositionLiftPercentageAttr zcl.AttrID = 8

func (a CurrentPositionLiftPercentage) ID() zcl.AttrID                         { return CurrentPositionLiftPercentageAttr }
func (a CurrentPositionLiftPercentage) Cluster() zcl.ClusterID                 { return WindowCoveringID }
func (a *CurrentPositionLiftPercentage) Value() *CurrentPositionLiftPercentage { return a }
func (a CurrentPositionLiftPercentage) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *CurrentPositionLiftPercentage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentPositionLiftPercentage(*nt)
	return br, err
}
func (CurrentPositionLiftPercentage) Name() string     { return "Current Position Lift Percentage" }
func (CurrentPositionLiftPercentage) Readable() bool   { return true }
func (CurrentPositionLiftPercentage) Writable() bool   { return false }
func (CurrentPositionLiftPercentage) Reportable() bool { return false }
func (CurrentPositionLiftPercentage) SceneIndex() int  { return -1 }

func (a CurrentPositionLiftPercentage) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// CurrentPositionTiltPercentage is an autogenerated attribute in the WindowCovering cluster
type CurrentPositionTiltPercentage zcl.Zu8

const CurrentPositionTiltPercentageAttr zcl.AttrID = 9

func (a CurrentPositionTiltPercentage) ID() zcl.AttrID                         { return CurrentPositionTiltPercentageAttr }
func (a CurrentPositionTiltPercentage) Cluster() zcl.ClusterID                 { return WindowCoveringID }
func (a *CurrentPositionTiltPercentage) Value() *CurrentPositionTiltPercentage { return a }
func (a CurrentPositionTiltPercentage) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *CurrentPositionTiltPercentage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentPositionTiltPercentage(*nt)
	return br, err
}
func (CurrentPositionTiltPercentage) Name() string     { return "Current Position Tilt Percentage" }
func (CurrentPositionTiltPercentage) Readable() bool   { return true }
func (CurrentPositionTiltPercentage) Writable() bool   { return false }
func (CurrentPositionTiltPercentage) Reportable() bool { return false }
func (CurrentPositionTiltPercentage) SceneIndex() int  { return -1 }

func (a CurrentPositionTiltPercentage) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// InstalledOpenLimitLift is an autogenerated attribute in the WindowCovering cluster
type InstalledOpenLimitLift zcl.Zu16

const InstalledOpenLimitLiftAttr zcl.AttrID = 16

func (a InstalledOpenLimitLift) ID() zcl.AttrID                  { return InstalledOpenLimitLiftAttr }
func (a InstalledOpenLimitLift) Cluster() zcl.ClusterID          { return WindowCoveringID }
func (a *InstalledOpenLimitLift) Value() *InstalledOpenLimitLift { return a }
func (a InstalledOpenLimitLift) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *InstalledOpenLimitLift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = InstalledOpenLimitLift(*nt)
	return br, err
}
func (InstalledOpenLimitLift) Name() string     { return "Installed Open Limit - Lift" }
func (InstalledOpenLimitLift) Readable() bool   { return true }
func (InstalledOpenLimitLift) Writable() bool   { return false }
func (InstalledOpenLimitLift) Reportable() bool { return false }
func (InstalledOpenLimitLift) SceneIndex() int  { return -1 }

func (a InstalledOpenLimitLift) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// InstalledClosedLimitLift is an autogenerated attribute in the WindowCovering cluster
type InstalledClosedLimitLift zcl.Zu16

const InstalledClosedLimitLiftAttr zcl.AttrID = 17

func (a InstalledClosedLimitLift) ID() zcl.AttrID                    { return InstalledClosedLimitLiftAttr }
func (a InstalledClosedLimitLift) Cluster() zcl.ClusterID            { return WindowCoveringID }
func (a *InstalledClosedLimitLift) Value() *InstalledClosedLimitLift { return a }
func (a InstalledClosedLimitLift) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *InstalledClosedLimitLift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = InstalledClosedLimitLift(*nt)
	return br, err
}
func (InstalledClosedLimitLift) Name() string     { return "Installed Closed Limit - Lift" }
func (InstalledClosedLimitLift) Readable() bool   { return true }
func (InstalledClosedLimitLift) Writable() bool   { return false }
func (InstalledClosedLimitLift) Reportable() bool { return false }
func (InstalledClosedLimitLift) SceneIndex() int  { return -1 }

func (a InstalledClosedLimitLift) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// InstalledOpenLimitTiltA is an autogenerated attribute in the WindowCovering cluster
type InstalledOpenLimitTiltA zcl.Zu16

const InstalledOpenLimitTiltAAttr zcl.AttrID = 18

func (a InstalledOpenLimitTiltA) ID() zcl.AttrID                   { return InstalledOpenLimitTiltAAttr }
func (a InstalledOpenLimitTiltA) Cluster() zcl.ClusterID           { return WindowCoveringID }
func (a *InstalledOpenLimitTiltA) Value() *InstalledOpenLimitTiltA { return a }
func (a InstalledOpenLimitTiltA) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *InstalledOpenLimitTiltA) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = InstalledOpenLimitTiltA(*nt)
	return br, err
}
func (InstalledOpenLimitTiltA) Name() string     { return "Installed Open Limit - Tilt A" }
func (InstalledOpenLimitTiltA) Readable() bool   { return true }
func (InstalledOpenLimitTiltA) Writable() bool   { return false }
func (InstalledOpenLimitTiltA) Reportable() bool { return false }
func (InstalledOpenLimitTiltA) SceneIndex() int  { return -1 }

func (a InstalledOpenLimitTiltA) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// InstalledOpenLimitTiltB is an autogenerated attribute in the WindowCovering cluster
type InstalledOpenLimitTiltB zcl.Zu16

const InstalledOpenLimitTiltBAttr zcl.AttrID = 19

func (a InstalledOpenLimitTiltB) ID() zcl.AttrID                   { return InstalledOpenLimitTiltBAttr }
func (a InstalledOpenLimitTiltB) Cluster() zcl.ClusterID           { return WindowCoveringID }
func (a *InstalledOpenLimitTiltB) Value() *InstalledOpenLimitTiltB { return a }
func (a InstalledOpenLimitTiltB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *InstalledOpenLimitTiltB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = InstalledOpenLimitTiltB(*nt)
	return br, err
}
func (InstalledOpenLimitTiltB) Name() string     { return "Installed Open Limit - Tilt B" }
func (InstalledOpenLimitTiltB) Readable() bool   { return true }
func (InstalledOpenLimitTiltB) Writable() bool   { return false }
func (InstalledOpenLimitTiltB) Reportable() bool { return false }
func (InstalledOpenLimitTiltB) SceneIndex() int  { return -1 }

func (a InstalledOpenLimitTiltB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// VelocityLift is an autogenerated attribute in the WindowCovering cluster
type VelocityLift zcl.Zu16

const VelocityLiftAttr zcl.AttrID = 20

func (a VelocityLift) ID() zcl.AttrID         { return VelocityLiftAttr }
func (a VelocityLift) Cluster() zcl.ClusterID { return WindowCoveringID }
func (a *VelocityLift) Value() *VelocityLift  { return a }
func (a VelocityLift) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *VelocityLift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = VelocityLift(*nt)
	return br, err
}
func (VelocityLift) Name() string     { return "Velocity - Lift" }
func (VelocityLift) Readable() bool   { return true }
func (VelocityLift) Writable() bool   { return true }
func (VelocityLift) Reportable() bool { return false }
func (VelocityLift) SceneIndex() int  { return -1 }

func (a VelocityLift) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AccelerationTimeLift is an autogenerated attribute in the WindowCovering cluster
type AccelerationTimeLift zcl.Zu16

const AccelerationTimeLiftAttr zcl.AttrID = 21

func (a AccelerationTimeLift) ID() zcl.AttrID                { return AccelerationTimeLiftAttr }
func (a AccelerationTimeLift) Cluster() zcl.ClusterID        { return WindowCoveringID }
func (a *AccelerationTimeLift) Value() *AccelerationTimeLift { return a }
func (a AccelerationTimeLift) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AccelerationTimeLift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AccelerationTimeLift(*nt)
	return br, err
}
func (AccelerationTimeLift) Name() string     { return "Acceleration Time - Lift" }
func (AccelerationTimeLift) Readable() bool   { return true }
func (AccelerationTimeLift) Writable() bool   { return true }
func (AccelerationTimeLift) Reportable() bool { return false }
func (AccelerationTimeLift) SceneIndex() int  { return -1 }

func (a AccelerationTimeLift) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// DecelerationTimeLift is an autogenerated attribute in the WindowCovering cluster
type DecelerationTimeLift zcl.Zu16

const DecelerationTimeLiftAttr zcl.AttrID = 22

func (a DecelerationTimeLift) ID() zcl.AttrID                { return DecelerationTimeLiftAttr }
func (a DecelerationTimeLift) Cluster() zcl.ClusterID        { return WindowCoveringID }
func (a *DecelerationTimeLift) Value() *DecelerationTimeLift { return a }
func (a DecelerationTimeLift) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DecelerationTimeLift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DecelerationTimeLift(*nt)
	return br, err
}
func (DecelerationTimeLift) Name() string     { return "Deceleration Time - Lift" }
func (DecelerationTimeLift) Readable() bool   { return true }
func (DecelerationTimeLift) Writable() bool   { return true }
func (DecelerationTimeLift) Reportable() bool { return false }
func (DecelerationTimeLift) SceneIndex() int  { return -1 }

func (a DecelerationTimeLift) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// WindowCoveringMode is an autogenerated attribute in the WindowCovering cluster
type WindowCoveringMode zcl.Zbmp8

const WindowCoveringModeAttr zcl.AttrID = 23

func (a WindowCoveringMode) ID() zcl.AttrID              { return WindowCoveringModeAttr }
func (a WindowCoveringMode) Cluster() zcl.ClusterID      { return WindowCoveringID }
func (a *WindowCoveringMode) Value() *WindowCoveringMode { return a }
func (a WindowCoveringMode) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *WindowCoveringMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = WindowCoveringMode(*nt)
	return br, err
}
func (WindowCoveringMode) Name() string     { return "Window Covering Mode" }
func (WindowCoveringMode) Readable() bool   { return true }
func (WindowCoveringMode) Writable() bool   { return true }
func (WindowCoveringMode) Reportable() bool { return false }
func (WindowCoveringMode) SceneIndex() int  { return -1 }

func (a WindowCoveringMode) String() string {
	var bstr []string
	if a.IsReversed() {
		bstr = append(bstr, "Reversed")
	}
	if a.IsCalibrationMode() {
		bstr = append(bstr, "Calibration Mode")
	}
	if a.IsMaintenanceMode() {
		bstr = append(bstr, "Maintenance Mode")
	}
	if a.IsLedFeedback() {
		bstr = append(bstr, "LED feedback")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a WindowCoveringMode) IsReversed() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *WindowCoveringMode) SetReversed(b bool) {
	*a = WindowCoveringMode(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a WindowCoveringMode) IsCalibrationMode() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *WindowCoveringMode) SetCalibrationMode(b bool) {
	*a = WindowCoveringMode(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a WindowCoveringMode) IsMaintenanceMode() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *WindowCoveringMode) SetMaintenanceMode(b bool) {
	*a = WindowCoveringMode(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a WindowCoveringMode) IsLedFeedback() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *WindowCoveringMode) SetLedFeedback(b bool) {
	*a = WindowCoveringMode(zcl.BitmapSet([]byte(*a), 3, b))
}

// IntermediateSetpointsLift is an autogenerated attribute in the WindowCovering cluster
type IntermediateSetpointsLift zcl.Zostring

const IntermediateSetpointsLiftAttr zcl.AttrID = 24

func (a IntermediateSetpointsLift) ID() zcl.AttrID                     { return IntermediateSetpointsLiftAttr }
func (a IntermediateSetpointsLift) Cluster() zcl.ClusterID             { return WindowCoveringID }
func (a *IntermediateSetpointsLift) Value() *IntermediateSetpointsLift { return a }
func (a IntermediateSetpointsLift) MarshalZcl() ([]byte, error) {
	return zcl.Zostring(a).MarshalZcl()
}
func (a *IntermediateSetpointsLift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = IntermediateSetpointsLift(*nt)
	return br, err
}
func (IntermediateSetpointsLift) Name() string     { return "Intermediate Setpoints - Lift" }
func (IntermediateSetpointsLift) Readable() bool   { return true }
func (IntermediateSetpointsLift) Writable() bool   { return true }
func (IntermediateSetpointsLift) Reportable() bool { return false }
func (IntermediateSetpointsLift) SceneIndex() int  { return -1 }

func (a IntermediateSetpointsLift) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(a))
}

// IntermediateSetpointsTilt is an autogenerated attribute in the WindowCovering cluster
type IntermediateSetpointsTilt zcl.Zostring

const IntermediateSetpointsTiltAttr zcl.AttrID = 25

func (a IntermediateSetpointsTilt) ID() zcl.AttrID                     { return IntermediateSetpointsTiltAttr }
func (a IntermediateSetpointsTilt) Cluster() zcl.ClusterID             { return WindowCoveringID }
func (a *IntermediateSetpointsTilt) Value() *IntermediateSetpointsTilt { return a }
func (a IntermediateSetpointsTilt) MarshalZcl() ([]byte, error) {
	return zcl.Zostring(a).MarshalZcl()
}
func (a *IntermediateSetpointsTilt) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = IntermediateSetpointsTilt(*nt)
	return br, err
}
func (IntermediateSetpointsTilt) Name() string     { return "Intermediate Setpoints - Tilt" }
func (IntermediateSetpointsTilt) Readable() bool   { return true }
func (IntermediateSetpointsTilt) Writable() bool   { return true }
func (IntermediateSetpointsTilt) Reportable() bool { return false }
func (IntermediateSetpointsTilt) SceneIndex() int  { return -1 }

func (a IntermediateSetpointsTilt) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(a))
}
