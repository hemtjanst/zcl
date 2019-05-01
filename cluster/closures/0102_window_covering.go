// The window covering cluster provides an interface for controlling and adjusting automatic window coverings such as drapery motors, automatic shades, and blinds.
package closures

import (
	"neotor.se/zcl"
)

// WindowCovering
// The window covering cluster provides an interface for controlling and adjusting automatic window coverings such as drapery motors, automatic shades, and blinds.

func NewWindowCoveringServer(profile zcl.ProfileID) *WindowCoveringServer {
	return &WindowCoveringServer{p: profile}
}
func NewWindowCoveringClient(profile zcl.ProfileID) *WindowCoveringClient {
	return &WindowCoveringClient{p: profile}
}

const WindowCoveringCluster zcl.ClusterID = 258

type WindowCoveringServer struct {
	p zcl.ProfileID

	WindowCoveringType            *WindowCoveringType
	PhysicalClosedLimitLift       *PhysicalClosedLimitLift
	PhysicalClosedLimitTilt       *PhysicalClosedLimitTilt
	CurrentPositionLift           *CurrentPositionLift
	CurrentPositionTilt           *CurrentPositionTilt
	NumberOfActuationsLift        *NumberOfActuationsLift
	NumberOfActuationsTilt        *NumberOfActuationsTilt
	ConfigStatus                  *ConfigStatus
	CurrentPositionLiftPercentage *CurrentPositionLiftPercentage
	CurrentPositionTiltPercentage *CurrentPositionTiltPercentage
	InstalledOpenLimitLift        *InstalledOpenLimitLift
	InstalledClosedLimitLift      *InstalledClosedLimitLift
	InstalledOpenLimitTiltA       *InstalledOpenLimitTiltA
	InstalledOpenLimitTiltB       *InstalledOpenLimitTiltB
	VelocityLift                  *VelocityLift
	AccelerationTimeLift          *AccelerationTimeLift
	DecelerationTimeLift          *DecelerationTimeLift
	WindowCoveringMode            *WindowCoveringMode
	IntermediateSetpointsLift     *IntermediateSetpointsLift
	IntermediateSetpointsTilt     *IntermediateSetpointsTilt
}

func (s *WindowCoveringServer) UpOpen() *UpOpen               { return new(UpOpen) }
func (s *WindowCoveringServer) DownClose() *DownClose         { return new(DownClose) }
func (s *WindowCoveringServer) Stop() *Stop                   { return new(Stop) }
func (s *WindowCoveringServer) GoToLiftValue() *GoToLiftValue { return new(GoToLiftValue) }
func (s *WindowCoveringServer) GoToLiftPercentage() *GoToLiftPercentage {
	return new(GoToLiftPercentage)
}
func (s *WindowCoveringServer) GoToTiltValue() *GoToTiltValue { return new(GoToTiltValue) }
func (s *WindowCoveringServer) GoToTiltPercentage() *GoToTiltPercentage {
	return new(GoToTiltPercentage)
}

type WindowCoveringClient struct {
	p zcl.ProfileID
}

/*
var WindowCoveringServer = map[zcl.CommandID]func() zcl.Command{
    UpOpenID: func() zcl.Command { return new(UpOpen) },
    DownCloseID: func() zcl.Command { return new(DownClose) },
    StopID: func() zcl.Command { return new(Stop) },
    GoToLiftValueID: func() zcl.Command { return new(GoToLiftValue) },
    GoToLiftPercentageID: func() zcl.Command { return new(GoToLiftPercentage) },
    GoToTiltValueID: func() zcl.Command { return new(GoToTiltValue) },
    GoToTiltPercentageID: func() zcl.Command { return new(GoToTiltPercentage) },
}

var WindowCoveringClient = map[zcl.CommandID]func() zcl.Command{
}
*/

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
	return WindowCoveringCluster
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
	return WindowCoveringCluster
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
	return WindowCoveringCluster
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
	return WindowCoveringCluster
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
	return WindowCoveringCluster
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
	return WindowCoveringCluster
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
	return WindowCoveringCluster
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

type WindowCoveringType zcl.Zenum8

func (a WindowCoveringType) ID() zcl.AttrID              { return 0 }
func (a WindowCoveringType) Cluster() zcl.ClusterID      { return WindowCoveringCluster }
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

func (a WindowCoveringType) Readable() bool   { return true }
func (a WindowCoveringType) Writable() bool   { return false }
func (a WindowCoveringType) Reportable() bool { return false }
func (a WindowCoveringType) SceneIndex() int  { return -1 }

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

	return zcl.Sprintf("%s", zcl.Zenum8(a))
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

type PhysicalClosedLimitLift zcl.Zu16

func (a PhysicalClosedLimitLift) ID() zcl.AttrID                   { return 1 }
func (a PhysicalClosedLimitLift) Cluster() zcl.ClusterID           { return WindowCoveringCluster }
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

func (a PhysicalClosedLimitLift) Readable() bool   { return true }
func (a PhysicalClosedLimitLift) Writable() bool   { return false }
func (a PhysicalClosedLimitLift) Reportable() bool { return false }
func (a PhysicalClosedLimitLift) SceneIndex() int  { return -1 }

func (a PhysicalClosedLimitLift) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type PhysicalClosedLimitTilt zcl.Zu16

func (a PhysicalClosedLimitTilt) ID() zcl.AttrID                   { return 2 }
func (a PhysicalClosedLimitTilt) Cluster() zcl.ClusterID           { return WindowCoveringCluster }
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

func (a PhysicalClosedLimitTilt) Readable() bool   { return true }
func (a PhysicalClosedLimitTilt) Writable() bool   { return false }
func (a PhysicalClosedLimitTilt) Reportable() bool { return false }
func (a PhysicalClosedLimitTilt) SceneIndex() int  { return -1 }

func (a PhysicalClosedLimitTilt) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type CurrentPositionLift zcl.Zu16

func (a CurrentPositionLift) ID() zcl.AttrID               { return 3 }
func (a CurrentPositionLift) Cluster() zcl.ClusterID       { return WindowCoveringCluster }
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

func (a CurrentPositionLift) Readable() bool   { return true }
func (a CurrentPositionLift) Writable() bool   { return false }
func (a CurrentPositionLift) Reportable() bool { return false }
func (a CurrentPositionLift) SceneIndex() int  { return -1 }

func (a CurrentPositionLift) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type CurrentPositionTilt zcl.Zu16

func (a CurrentPositionTilt) ID() zcl.AttrID               { return 4 }
func (a CurrentPositionTilt) Cluster() zcl.ClusterID       { return WindowCoveringCluster }
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

func (a CurrentPositionTilt) Readable() bool   { return true }
func (a CurrentPositionTilt) Writable() bool   { return false }
func (a CurrentPositionTilt) Reportable() bool { return false }
func (a CurrentPositionTilt) SceneIndex() int  { return -1 }

func (a CurrentPositionTilt) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type NumberOfActuationsLift zcl.Zu16

func (a NumberOfActuationsLift) ID() zcl.AttrID                  { return 5 }
func (a NumberOfActuationsLift) Cluster() zcl.ClusterID          { return WindowCoveringCluster }
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

func (a NumberOfActuationsLift) Readable() bool   { return true }
func (a NumberOfActuationsLift) Writable() bool   { return false }
func (a NumberOfActuationsLift) Reportable() bool { return false }
func (a NumberOfActuationsLift) SceneIndex() int  { return -1 }

func (a NumberOfActuationsLift) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type NumberOfActuationsTilt zcl.Zu16

func (a NumberOfActuationsTilt) ID() zcl.AttrID                  { return 6 }
func (a NumberOfActuationsTilt) Cluster() zcl.ClusterID          { return WindowCoveringCluster }
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

func (a NumberOfActuationsTilt) Readable() bool   { return true }
func (a NumberOfActuationsTilt) Writable() bool   { return false }
func (a NumberOfActuationsTilt) Reportable() bool { return false }
func (a NumberOfActuationsTilt) SceneIndex() int  { return -1 }

func (a NumberOfActuationsTilt) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type ConfigStatus zcl.Zbmp8

func (a ConfigStatus) ID() zcl.AttrID         { return 7 }
func (a ConfigStatus) Cluster() zcl.ClusterID { return WindowCoveringCluster }
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

func (a ConfigStatus) Readable() bool   { return true }
func (a ConfigStatus) Writable() bool   { return false }
func (a ConfigStatus) Reportable() bool { return false }
func (a ConfigStatus) SceneIndex() int  { return -1 }

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

type CurrentPositionLiftPercentage zcl.Zu8

func (a CurrentPositionLiftPercentage) ID() zcl.AttrID                         { return 8 }
func (a CurrentPositionLiftPercentage) Cluster() zcl.ClusterID                 { return WindowCoveringCluster }
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

func (a CurrentPositionLiftPercentage) Readable() bool   { return true }
func (a CurrentPositionLiftPercentage) Writable() bool   { return false }
func (a CurrentPositionLiftPercentage) Reportable() bool { return false }
func (a CurrentPositionLiftPercentage) SceneIndex() int  { return -1 }

func (a CurrentPositionLiftPercentage) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type CurrentPositionTiltPercentage zcl.Zu8

func (a CurrentPositionTiltPercentage) ID() zcl.AttrID                         { return 9 }
func (a CurrentPositionTiltPercentage) Cluster() zcl.ClusterID                 { return WindowCoveringCluster }
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

func (a CurrentPositionTiltPercentage) Readable() bool   { return true }
func (a CurrentPositionTiltPercentage) Writable() bool   { return false }
func (a CurrentPositionTiltPercentage) Reportable() bool { return false }
func (a CurrentPositionTiltPercentage) SceneIndex() int  { return -1 }

func (a CurrentPositionTiltPercentage) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type InstalledOpenLimitLift zcl.Zu16

func (a InstalledOpenLimitLift) ID() zcl.AttrID                  { return 16 }
func (a InstalledOpenLimitLift) Cluster() zcl.ClusterID          { return WindowCoveringCluster }
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

func (a InstalledOpenLimitLift) Readable() bool   { return true }
func (a InstalledOpenLimitLift) Writable() bool   { return false }
func (a InstalledOpenLimitLift) Reportable() bool { return false }
func (a InstalledOpenLimitLift) SceneIndex() int  { return -1 }

func (a InstalledOpenLimitLift) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type InstalledClosedLimitLift zcl.Zu16

func (a InstalledClosedLimitLift) ID() zcl.AttrID                    { return 17 }
func (a InstalledClosedLimitLift) Cluster() zcl.ClusterID            { return WindowCoveringCluster }
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

func (a InstalledClosedLimitLift) Readable() bool   { return true }
func (a InstalledClosedLimitLift) Writable() bool   { return false }
func (a InstalledClosedLimitLift) Reportable() bool { return false }
func (a InstalledClosedLimitLift) SceneIndex() int  { return -1 }

func (a InstalledClosedLimitLift) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type InstalledOpenLimitTiltA zcl.Zu16

func (a InstalledOpenLimitTiltA) ID() zcl.AttrID                   { return 18 }
func (a InstalledOpenLimitTiltA) Cluster() zcl.ClusterID           { return WindowCoveringCluster }
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

func (a InstalledOpenLimitTiltA) Readable() bool   { return true }
func (a InstalledOpenLimitTiltA) Writable() bool   { return false }
func (a InstalledOpenLimitTiltA) Reportable() bool { return false }
func (a InstalledOpenLimitTiltA) SceneIndex() int  { return -1 }

func (a InstalledOpenLimitTiltA) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type InstalledOpenLimitTiltB zcl.Zu16

func (a InstalledOpenLimitTiltB) ID() zcl.AttrID                   { return 19 }
func (a InstalledOpenLimitTiltB) Cluster() zcl.ClusterID           { return WindowCoveringCluster }
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

func (a InstalledOpenLimitTiltB) Readable() bool   { return true }
func (a InstalledOpenLimitTiltB) Writable() bool   { return false }
func (a InstalledOpenLimitTiltB) Reportable() bool { return false }
func (a InstalledOpenLimitTiltB) SceneIndex() int  { return -1 }

func (a InstalledOpenLimitTiltB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type VelocityLift zcl.Zu16

func (a VelocityLift) ID() zcl.AttrID         { return 20 }
func (a VelocityLift) Cluster() zcl.ClusterID { return WindowCoveringCluster }
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

func (a VelocityLift) Readable() bool   { return true }
func (a VelocityLift) Writable() bool   { return true }
func (a VelocityLift) Reportable() bool { return false }
func (a VelocityLift) SceneIndex() int  { return -1 }

func (a VelocityLift) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type AccelerationTimeLift zcl.Zu16

func (a AccelerationTimeLift) ID() zcl.AttrID                { return 21 }
func (a AccelerationTimeLift) Cluster() zcl.ClusterID        { return WindowCoveringCluster }
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

func (a AccelerationTimeLift) Readable() bool   { return true }
func (a AccelerationTimeLift) Writable() bool   { return true }
func (a AccelerationTimeLift) Reportable() bool { return false }
func (a AccelerationTimeLift) SceneIndex() int  { return -1 }

func (a AccelerationTimeLift) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type DecelerationTimeLift zcl.Zu16

func (a DecelerationTimeLift) ID() zcl.AttrID                { return 22 }
func (a DecelerationTimeLift) Cluster() zcl.ClusterID        { return WindowCoveringCluster }
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

func (a DecelerationTimeLift) Readable() bool   { return true }
func (a DecelerationTimeLift) Writable() bool   { return true }
func (a DecelerationTimeLift) Reportable() bool { return false }
func (a DecelerationTimeLift) SceneIndex() int  { return -1 }

func (a DecelerationTimeLift) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type WindowCoveringMode zcl.Zbmp8

func (a WindowCoveringMode) ID() zcl.AttrID              { return 23 }
func (a WindowCoveringMode) Cluster() zcl.ClusterID      { return WindowCoveringCluster }
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

func (a WindowCoveringMode) Readable() bool   { return true }
func (a WindowCoveringMode) Writable() bool   { return true }
func (a WindowCoveringMode) Reportable() bool { return false }
func (a WindowCoveringMode) SceneIndex() int  { return -1 }

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

type IntermediateSetpointsLift zcl.Zostring

func (a IntermediateSetpointsLift) ID() zcl.AttrID                     { return 24 }
func (a IntermediateSetpointsLift) Cluster() zcl.ClusterID             { return WindowCoveringCluster }
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

func (a IntermediateSetpointsLift) Readable() bool   { return true }
func (a IntermediateSetpointsLift) Writable() bool   { return true }
func (a IntermediateSetpointsLift) Reportable() bool { return false }
func (a IntermediateSetpointsLift) SceneIndex() int  { return -1 }

func (a IntermediateSetpointsLift) String() string {

	return zcl.Sprintf("%s", zcl.Zostring(a))
}

type IntermediateSetpointsTilt zcl.Zostring

func (a IntermediateSetpointsTilt) ID() zcl.AttrID                     { return 25 }
func (a IntermediateSetpointsTilt) Cluster() zcl.ClusterID             { return WindowCoveringCluster }
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

func (a IntermediateSetpointsTilt) Readable() bool   { return true }
func (a IntermediateSetpointsTilt) Writable() bool   { return true }
func (a IntermediateSetpointsTilt) Reportable() bool { return false }
func (a IntermediateSetpointsTilt) SceneIndex() int  { return -1 }

func (a IntermediateSetpointsTilt) String() string {

	return zcl.Sprintf("%s", zcl.Zostring(a))
}
