package closures

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID

const ConfigStatusAttr zcl.AttrID = 7

func (ConfigStatus) ID() zcl.AttrID   { return ConfigStatusAttr }
func (ConfigStatus) Readable() bool   { return true }
func (ConfigStatus) Writable() bool   { return false }
func (ConfigStatus) Reportable() bool { return false }
func (ConfigStatus) SceneIndex() int  { return -1 }

func (ConfigStatus) Name() string { return "Config / Status" }

type ConfigStatus zcl.Zbmp8

func (a *ConfigStatus) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (a *ConfigStatus) Value() zcl.Val     { return a }

func (a ConfigStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *ConfigStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = ConfigStatus(*nt)
	return br, err
}

func (a ConfigStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(a))
}

func (a *ConfigStatus) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = ConfigStatus(*v)
	return nil
}

func (a *ConfigStatus) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = ConfigStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ConfigStatus) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Operational")
		case 1:
			bstr = append(bstr, "Online")
		case 2:
			bstr = append(bstr, "Commands Reversed")
		case 3:
			bstr = append(bstr, "Lift control is Closed Loop")
		case 4:
			bstr = append(bstr, "Tilt control is Closed Loop")
		case 5:
			bstr = append(bstr, "Lift: Encoder Controlled")
		case 6:
			bstr = append(bstr, "Tilt: Encoder Controlled")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a ConfigStatus) IsOperational() bool             { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a ConfigStatus) IsOnline() bool                  { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a ConfigStatus) IsCommandsReversed() bool        { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a ConfigStatus) IsLiftControlIsClosedLoop() bool { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a ConfigStatus) IsTiltControlIsClosedLoop() bool { return zcl.BitmapTest([]byte(a[:]), 4) }
func (a ConfigStatus) IsLiftEncoderControlled() bool   { return zcl.BitmapTest([]byte(a[:]), 5) }
func (a ConfigStatus) IsTiltEncoderControlled() bool   { return zcl.BitmapTest([]byte(a[:]), 6) }
func (a *ConfigStatus) SetOperational(b bool)          { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *ConfigStatus) SetOnline(b bool)               { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b)) }
func (a *ConfigStatus) SetCommandsReversed(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *ConfigStatus) SetLiftControlIsClosedLoop(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}
func (a *ConfigStatus) SetTiltControlIsClosedLoop(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 4, b))
}
func (a *ConfigStatus) SetLiftEncoderControlled(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 5, b))
}
func (a *ConfigStatus) SetTiltEncoderControlled(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 6, b))
}

func (ConfigStatus) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Operational"},
		{Value: 1, Name: "Online"},
		{Value: 2, Name: "Commands Reversed"},
		{Value: 3, Name: "Lift control is Closed Loop"},
		{Value: 4, Name: "Tilt control is Closed Loop"},
		{Value: 5, Name: "Lift: Encoder Controlled"},
		{Value: 6, Name: "Tilt: Encoder Controlled"},
	}
}

const LiftAccelerationTimeAttr zcl.AttrID = 21

func (LiftAccelerationTime) ID() zcl.AttrID   { return LiftAccelerationTimeAttr }
func (LiftAccelerationTime) Readable() bool   { return true }
func (LiftAccelerationTime) Writable() bool   { return true }
func (LiftAccelerationTime) Reportable() bool { return false }
func (LiftAccelerationTime) SceneIndex() int  { return -1 }

func (LiftAccelerationTime) Name() string { return "Lift - Acceleration Time" }

type LiftAccelerationTime zcl.Zu16

func (a *LiftAccelerationTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *LiftAccelerationTime) Value() zcl.Val     { return a }

func (a LiftAccelerationTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LiftAccelerationTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LiftAccelerationTime(*nt)
	return br, err
}

func (a LiftAccelerationTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *LiftAccelerationTime) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LiftAccelerationTime(*v)
	return nil
}

func (a *LiftAccelerationTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = LiftAccelerationTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LiftAccelerationTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const LiftCurrentPositionAttr zcl.AttrID = 3

func (LiftCurrentPosition) ID() zcl.AttrID   { return LiftCurrentPositionAttr }
func (LiftCurrentPosition) Readable() bool   { return true }
func (LiftCurrentPosition) Writable() bool   { return false }
func (LiftCurrentPosition) Reportable() bool { return false }
func (LiftCurrentPosition) SceneIndex() int  { return -1 }

func (LiftCurrentPosition) Name() string { return "Lift - Current Position" }

type LiftCurrentPosition zcl.Zu16

func (a *LiftCurrentPosition) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *LiftCurrentPosition) Value() zcl.Val     { return a }

func (a LiftCurrentPosition) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LiftCurrentPosition) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LiftCurrentPosition(*nt)
	return br, err
}

func (a LiftCurrentPosition) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *LiftCurrentPosition) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LiftCurrentPosition(*v)
	return nil
}

func (a *LiftCurrentPosition) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = LiftCurrentPosition(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LiftCurrentPosition) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const LiftDecelerationTimeAttr zcl.AttrID = 22

func (LiftDecelerationTime) ID() zcl.AttrID   { return LiftDecelerationTimeAttr }
func (LiftDecelerationTime) Readable() bool   { return true }
func (LiftDecelerationTime) Writable() bool   { return true }
func (LiftDecelerationTime) Reportable() bool { return false }
func (LiftDecelerationTime) SceneIndex() int  { return -1 }

func (LiftDecelerationTime) Name() string { return "Lift - Deceleration Time" }

type LiftDecelerationTime zcl.Zu16

func (a *LiftDecelerationTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *LiftDecelerationTime) Value() zcl.Val     { return a }

func (a LiftDecelerationTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LiftDecelerationTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LiftDecelerationTime(*nt)
	return br, err
}

func (a LiftDecelerationTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *LiftDecelerationTime) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LiftDecelerationTime(*v)
	return nil
}

func (a *LiftDecelerationTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = LiftDecelerationTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LiftDecelerationTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const LiftInstalledClosedLimitAttr zcl.AttrID = 17

func (LiftInstalledClosedLimit) ID() zcl.AttrID   { return LiftInstalledClosedLimitAttr }
func (LiftInstalledClosedLimit) Readable() bool   { return true }
func (LiftInstalledClosedLimit) Writable() bool   { return false }
func (LiftInstalledClosedLimit) Reportable() bool { return false }
func (LiftInstalledClosedLimit) SceneIndex() int  { return -1 }

func (LiftInstalledClosedLimit) Name() string { return "Lift - Installed Closed Limit" }

type LiftInstalledClosedLimit zcl.Zu16

func (a *LiftInstalledClosedLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *LiftInstalledClosedLimit) Value() zcl.Val     { return a }

func (a LiftInstalledClosedLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LiftInstalledClosedLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LiftInstalledClosedLimit(*nt)
	return br, err
}

func (a LiftInstalledClosedLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *LiftInstalledClosedLimit) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LiftInstalledClosedLimit(*v)
	return nil
}

func (a *LiftInstalledClosedLimit) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = LiftInstalledClosedLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LiftInstalledClosedLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const LiftInstalledOpenLimitAttr zcl.AttrID = 16

func (LiftInstalledOpenLimit) ID() zcl.AttrID   { return LiftInstalledOpenLimitAttr }
func (LiftInstalledOpenLimit) Readable() bool   { return true }
func (LiftInstalledOpenLimit) Writable() bool   { return false }
func (LiftInstalledOpenLimit) Reportable() bool { return false }
func (LiftInstalledOpenLimit) SceneIndex() int  { return -1 }

func (LiftInstalledOpenLimit) Name() string { return "Lift - Installed Open Limit" }

type LiftInstalledOpenLimit zcl.Zu16

func (a *LiftInstalledOpenLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *LiftInstalledOpenLimit) Value() zcl.Val     { return a }

func (a LiftInstalledOpenLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LiftInstalledOpenLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LiftInstalledOpenLimit(*nt)
	return br, err
}

func (a LiftInstalledOpenLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *LiftInstalledOpenLimit) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LiftInstalledOpenLimit(*v)
	return nil
}

func (a *LiftInstalledOpenLimit) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = LiftInstalledOpenLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LiftInstalledOpenLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const LiftIntermediateSetpointsAttr zcl.AttrID = 24

func (LiftIntermediateSetpoints) ID() zcl.AttrID   { return LiftIntermediateSetpointsAttr }
func (LiftIntermediateSetpoints) Readable() bool   { return true }
func (LiftIntermediateSetpoints) Writable() bool   { return true }
func (LiftIntermediateSetpoints) Reportable() bool { return false }
func (LiftIntermediateSetpoints) SceneIndex() int  { return -1 }

func (LiftIntermediateSetpoints) Name() string { return "Lift - Intermediate Setpoints" }

type LiftIntermediateSetpoints zcl.Zostring

func (a *LiftIntermediateSetpoints) TypeID() zcl.TypeID { return new(zcl.Zostring).TypeID() }
func (a *LiftIntermediateSetpoints) Value() zcl.Val     { return a }

func (a LiftIntermediateSetpoints) MarshalZcl() ([]byte, error) { return zcl.Zostring(a).MarshalZcl() }

func (a *LiftIntermediateSetpoints) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = LiftIntermediateSetpoints(*nt)
	return br, err
}

func (a LiftIntermediateSetpoints) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zostring(a))
}

func (a *LiftIntermediateSetpoints) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zostring)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LiftIntermediateSetpoints(*v)
	return nil
}

func (a *LiftIntermediateSetpoints) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zostring); ok {
		*a = LiftIntermediateSetpoints(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LiftIntermediateSetpoints) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(a))
}

const LiftNumberOfActuationsAttr zcl.AttrID = 5

func (LiftNumberOfActuations) ID() zcl.AttrID   { return LiftNumberOfActuationsAttr }
func (LiftNumberOfActuations) Readable() bool   { return true }
func (LiftNumberOfActuations) Writable() bool   { return false }
func (LiftNumberOfActuations) Reportable() bool { return false }
func (LiftNumberOfActuations) SceneIndex() int  { return -1 }

func (LiftNumberOfActuations) Name() string { return "Lift - Number of Actuations" }

type LiftNumberOfActuations zcl.Zu16

func (a *LiftNumberOfActuations) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *LiftNumberOfActuations) Value() zcl.Val     { return a }

func (a LiftNumberOfActuations) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LiftNumberOfActuations) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LiftNumberOfActuations(*nt)
	return br, err
}

func (a LiftNumberOfActuations) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *LiftNumberOfActuations) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LiftNumberOfActuations(*v)
	return nil
}

func (a *LiftNumberOfActuations) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = LiftNumberOfActuations(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LiftNumberOfActuations) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const LiftPhysicalClosedLimitAttr zcl.AttrID = 1

func (LiftPhysicalClosedLimit) ID() zcl.AttrID   { return LiftPhysicalClosedLimitAttr }
func (LiftPhysicalClosedLimit) Readable() bool   { return true }
func (LiftPhysicalClosedLimit) Writable() bool   { return false }
func (LiftPhysicalClosedLimit) Reportable() bool { return false }
func (LiftPhysicalClosedLimit) SceneIndex() int  { return -1 }

func (LiftPhysicalClosedLimit) Name() string { return "Lift - Physical Closed Limit" }

type LiftPhysicalClosedLimit zcl.Zu16

func (a *LiftPhysicalClosedLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *LiftPhysicalClosedLimit) Value() zcl.Val     { return a }

func (a LiftPhysicalClosedLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LiftPhysicalClosedLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LiftPhysicalClosedLimit(*nt)
	return br, err
}

func (a LiftPhysicalClosedLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *LiftPhysicalClosedLimit) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LiftPhysicalClosedLimit(*v)
	return nil
}

func (a *LiftPhysicalClosedLimit) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = LiftPhysicalClosedLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LiftPhysicalClosedLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const LiftVelocityAttr zcl.AttrID = 20

func (LiftVelocity) ID() zcl.AttrID   { return LiftVelocityAttr }
func (LiftVelocity) Readable() bool   { return true }
func (LiftVelocity) Writable() bool   { return true }
func (LiftVelocity) Reportable() bool { return false }
func (LiftVelocity) SceneIndex() int  { return -1 }

func (LiftVelocity) Name() string { return "Lift - Velocity" }

type LiftVelocity zcl.Zu16

func (a *LiftVelocity) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *LiftVelocity) Value() zcl.Val     { return a }

func (a LiftVelocity) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LiftVelocity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LiftVelocity(*nt)
	return br, err
}

func (a LiftVelocity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *LiftVelocity) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LiftVelocity(*v)
	return nil
}

func (a *LiftVelocity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = LiftVelocity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LiftVelocity) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const LiftCurrentPositionPercentageAttr zcl.AttrID = 8

func (LiftCurrentPositionPercentage) ID() zcl.AttrID   { return LiftCurrentPositionPercentageAttr }
func (LiftCurrentPositionPercentage) Readable() bool   { return true }
func (LiftCurrentPositionPercentage) Writable() bool   { return false }
func (LiftCurrentPositionPercentage) Reportable() bool { return false }
func (LiftCurrentPositionPercentage) SceneIndex() int  { return -1 }

func (LiftCurrentPositionPercentage) Name() string { return "Lift Current Position Percentage" }

type LiftCurrentPositionPercentage zcl.Zu8

func (a *LiftCurrentPositionPercentage) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *LiftCurrentPositionPercentage) Value() zcl.Val     { return a }

func (a LiftCurrentPositionPercentage) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *LiftCurrentPositionPercentage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = LiftCurrentPositionPercentage(*nt)
	return br, err
}

func (a LiftCurrentPositionPercentage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *LiftCurrentPositionPercentage) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = LiftCurrentPositionPercentage(*v)
	return nil
}

func (a *LiftCurrentPositionPercentage) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = LiftCurrentPositionPercentage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LiftCurrentPositionPercentage) String() string {
	return zcl.Percent.Format(float64(a))
}

func (Percentage) Name() string { return "Percentage" }

type Percentage zcl.Zu8

func (a *Percentage) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *Percentage) Value() zcl.Val     { return a }

func (a Percentage) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Percentage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Percentage(*nt)
	return br, err
}

func (a Percentage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *Percentage) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = Percentage(*v)
	return nil
}

func (a *Percentage) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Percentage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Percentage) String() string {
	return zcl.Percent.Format(float64(a))
}

func (Position) Name() string { return "Position" }

type Position zcl.Zu16

func (a *Position) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *Position) Value() zcl.Val     { return a }

func (a Position) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Position) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Position(*nt)
	return br, err
}

func (a Position) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *Position) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = Position(*v)
	return nil
}

func (a *Position) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Position(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Position) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const TiltCurrentPositionAttr zcl.AttrID = 4

func (TiltCurrentPosition) ID() zcl.AttrID   { return TiltCurrentPositionAttr }
func (TiltCurrentPosition) Readable() bool   { return true }
func (TiltCurrentPosition) Writable() bool   { return false }
func (TiltCurrentPosition) Reportable() bool { return false }
func (TiltCurrentPosition) SceneIndex() int  { return -1 }

func (TiltCurrentPosition) Name() string { return "Tilt - Current Position" }

type TiltCurrentPosition zcl.Zu16

func (a *TiltCurrentPosition) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *TiltCurrentPosition) Value() zcl.Val     { return a }

func (a TiltCurrentPosition) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *TiltCurrentPosition) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TiltCurrentPosition(*nt)
	return br, err
}

func (a TiltCurrentPosition) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *TiltCurrentPosition) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = TiltCurrentPosition(*v)
	return nil
}

func (a *TiltCurrentPosition) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = TiltCurrentPosition(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TiltCurrentPosition) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const TiltIntermediateSetpointsAttr zcl.AttrID = 25

func (TiltIntermediateSetpoints) ID() zcl.AttrID   { return TiltIntermediateSetpointsAttr }
func (TiltIntermediateSetpoints) Readable() bool   { return true }
func (TiltIntermediateSetpoints) Writable() bool   { return true }
func (TiltIntermediateSetpoints) Reportable() bool { return false }
func (TiltIntermediateSetpoints) SceneIndex() int  { return -1 }

func (TiltIntermediateSetpoints) Name() string { return "Tilt - Intermediate Setpoints" }

type TiltIntermediateSetpoints zcl.Zostring

func (a *TiltIntermediateSetpoints) TypeID() zcl.TypeID { return new(zcl.Zostring).TypeID() }
func (a *TiltIntermediateSetpoints) Value() zcl.Val     { return a }

func (a TiltIntermediateSetpoints) MarshalZcl() ([]byte, error) { return zcl.Zostring(a).MarshalZcl() }

func (a *TiltIntermediateSetpoints) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = TiltIntermediateSetpoints(*nt)
	return br, err
}

func (a TiltIntermediateSetpoints) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zostring(a))
}

func (a *TiltIntermediateSetpoints) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zostring)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = TiltIntermediateSetpoints(*v)
	return nil
}

func (a *TiltIntermediateSetpoints) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zostring); ok {
		*a = TiltIntermediateSetpoints(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TiltIntermediateSetpoints) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(a))
}

const TiltNumberOfActuationsAttr zcl.AttrID = 6

func (TiltNumberOfActuations) ID() zcl.AttrID   { return TiltNumberOfActuationsAttr }
func (TiltNumberOfActuations) Readable() bool   { return true }
func (TiltNumberOfActuations) Writable() bool   { return false }
func (TiltNumberOfActuations) Reportable() bool { return false }
func (TiltNumberOfActuations) SceneIndex() int  { return -1 }

func (TiltNumberOfActuations) Name() string { return "Tilt - Number of Actuations" }

type TiltNumberOfActuations zcl.Zu16

func (a *TiltNumberOfActuations) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *TiltNumberOfActuations) Value() zcl.Val     { return a }

func (a TiltNumberOfActuations) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *TiltNumberOfActuations) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TiltNumberOfActuations(*nt)
	return br, err
}

func (a TiltNumberOfActuations) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *TiltNumberOfActuations) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = TiltNumberOfActuations(*v)
	return nil
}

func (a *TiltNumberOfActuations) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = TiltNumberOfActuations(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TiltNumberOfActuations) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const TiltPhysicalClosedLimitAttr zcl.AttrID = 2

func (TiltPhysicalClosedLimit) ID() zcl.AttrID   { return TiltPhysicalClosedLimitAttr }
func (TiltPhysicalClosedLimit) Readable() bool   { return true }
func (TiltPhysicalClosedLimit) Writable() bool   { return false }
func (TiltPhysicalClosedLimit) Reportable() bool { return false }
func (TiltPhysicalClosedLimit) SceneIndex() int  { return -1 }

func (TiltPhysicalClosedLimit) Name() string { return "Tilt - Physical Closed Limit" }

type TiltPhysicalClosedLimit zcl.Zu16

func (a *TiltPhysicalClosedLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *TiltPhysicalClosedLimit) Value() zcl.Val     { return a }

func (a TiltPhysicalClosedLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *TiltPhysicalClosedLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TiltPhysicalClosedLimit(*nt)
	return br, err
}

func (a TiltPhysicalClosedLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *TiltPhysicalClosedLimit) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = TiltPhysicalClosedLimit(*v)
	return nil
}

func (a *TiltPhysicalClosedLimit) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = TiltPhysicalClosedLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TiltPhysicalClosedLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const TiltAInstalledOpenLimitAttr zcl.AttrID = 18

func (TiltAInstalledOpenLimit) ID() zcl.AttrID   { return TiltAInstalledOpenLimitAttr }
func (TiltAInstalledOpenLimit) Readable() bool   { return true }
func (TiltAInstalledOpenLimit) Writable() bool   { return false }
func (TiltAInstalledOpenLimit) Reportable() bool { return false }
func (TiltAInstalledOpenLimit) SceneIndex() int  { return -1 }

func (TiltAInstalledOpenLimit) Name() string { return "Tilt A - Installed Open Limit" }

type TiltAInstalledOpenLimit zcl.Zu16

func (a *TiltAInstalledOpenLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *TiltAInstalledOpenLimit) Value() zcl.Val     { return a }

func (a TiltAInstalledOpenLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *TiltAInstalledOpenLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TiltAInstalledOpenLimit(*nt)
	return br, err
}

func (a TiltAInstalledOpenLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *TiltAInstalledOpenLimit) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = TiltAInstalledOpenLimit(*v)
	return nil
}

func (a *TiltAInstalledOpenLimit) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = TiltAInstalledOpenLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TiltAInstalledOpenLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const TiltBInstalledOpenLimitAttr zcl.AttrID = 19

func (TiltBInstalledOpenLimit) ID() zcl.AttrID   { return TiltBInstalledOpenLimitAttr }
func (TiltBInstalledOpenLimit) Readable() bool   { return true }
func (TiltBInstalledOpenLimit) Writable() bool   { return false }
func (TiltBInstalledOpenLimit) Reportable() bool { return false }
func (TiltBInstalledOpenLimit) SceneIndex() int  { return -1 }

func (TiltBInstalledOpenLimit) Name() string { return "Tilt B - Installed Open Limit" }

type TiltBInstalledOpenLimit zcl.Zu16

func (a *TiltBInstalledOpenLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (a *TiltBInstalledOpenLimit) Value() zcl.Val     { return a }

func (a TiltBInstalledOpenLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *TiltBInstalledOpenLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TiltBInstalledOpenLimit(*nt)
	return br, err
}

func (a TiltBInstalledOpenLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(a))
}

func (a *TiltBInstalledOpenLimit) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu16)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = TiltBInstalledOpenLimit(*v)
	return nil
}

func (a *TiltBInstalledOpenLimit) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = TiltBInstalledOpenLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TiltBInstalledOpenLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

const TiltCurrentPositionPercentageAttr zcl.AttrID = 9

func (TiltCurrentPositionPercentage) ID() zcl.AttrID   { return TiltCurrentPositionPercentageAttr }
func (TiltCurrentPositionPercentage) Readable() bool   { return true }
func (TiltCurrentPositionPercentage) Writable() bool   { return false }
func (TiltCurrentPositionPercentage) Reportable() bool { return false }
func (TiltCurrentPositionPercentage) SceneIndex() int  { return -1 }

func (TiltCurrentPositionPercentage) Name() string { return "Tilt Current Position Percentage" }

type TiltCurrentPositionPercentage zcl.Zu8

func (a *TiltCurrentPositionPercentage) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (a *TiltCurrentPositionPercentage) Value() zcl.Val     { return a }

func (a TiltCurrentPositionPercentage) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *TiltCurrentPositionPercentage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = TiltCurrentPositionPercentage(*nt)
	return br, err
}

func (a TiltCurrentPositionPercentage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(a))
}

func (a *TiltCurrentPositionPercentage) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zu8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = TiltCurrentPositionPercentage(*v)
	return nil
}

func (a *TiltCurrentPositionPercentage) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = TiltCurrentPositionPercentage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TiltCurrentPositionPercentage) String() string {
	return zcl.Percent.Format(float64(a))
}

const WindowCoveringModeAttr zcl.AttrID = 23

func (WindowCoveringMode) ID() zcl.AttrID   { return WindowCoveringModeAttr }
func (WindowCoveringMode) Readable() bool   { return true }
func (WindowCoveringMode) Writable() bool   { return true }
func (WindowCoveringMode) Reportable() bool { return false }
func (WindowCoveringMode) SceneIndex() int  { return -1 }

func (WindowCoveringMode) Name() string { return "Window Covering Mode" }

type WindowCoveringMode zcl.Zbmp8

func (a *WindowCoveringMode) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (a *WindowCoveringMode) Value() zcl.Val     { return a }

func (a WindowCoveringMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *WindowCoveringMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = WindowCoveringMode(*nt)
	return br, err
}

func (a WindowCoveringMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(a))
}

func (a *WindowCoveringMode) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = WindowCoveringMode(*v)
	return nil
}

func (a *WindowCoveringMode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = WindowCoveringMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a WindowCoveringMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Reversed")
		case 1:
			bstr = append(bstr, "Calibration Mode")
		case 2:
			bstr = append(bstr, "Maintenance Mode")
		case 3:
			bstr = append(bstr, "LED feedback")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a WindowCoveringMode) IsReversed() bool        { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a WindowCoveringMode) IsCalibrationMode() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a WindowCoveringMode) IsMaintenanceMode() bool { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a WindowCoveringMode) IsLedFeedback() bool     { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *WindowCoveringMode) SetReversed(b bool)     { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *WindowCoveringMode) SetCalibrationMode(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *WindowCoveringMode) SetMaintenanceMode(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *WindowCoveringMode) SetLedFeedback(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}

func (WindowCoveringMode) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Reversed"},
		{Value: 1, Name: "Calibration Mode"},
		{Value: 2, Name: "Maintenance Mode"},
		{Value: 3, Name: "LED feedback"},
	}
}

const WindowCoveringTypeAttr zcl.AttrID = 0

func (WindowCoveringType) ID() zcl.AttrID   { return WindowCoveringTypeAttr }
func (WindowCoveringType) Readable() bool   { return true }
func (WindowCoveringType) Writable() bool   { return false }
func (WindowCoveringType) Reportable() bool { return false }
func (WindowCoveringType) SceneIndex() int  { return -1 }

func (WindowCoveringType) Name() string { return "Window Covering Type" }

type WindowCoveringType zcl.Zenum8

func (a *WindowCoveringType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (a *WindowCoveringType) Value() zcl.Val     { return a }

func (a WindowCoveringType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *WindowCoveringType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = WindowCoveringType(*nt)
	return br, err
}

func (a WindowCoveringType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(a))
}

func (a *WindowCoveringType) UnmarshalJSON(b []byte) error {
	v := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, v); err != nil {
		return err
	}
	*a = WindowCoveringType(*v)
	return nil
}

func (a *WindowCoveringType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = WindowCoveringType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

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

func (a WindowCoveringType) IsRollershade() bool               { return a == 0x00 }
func (a WindowCoveringType) IsRollershade2Motor() bool         { return a == 0x01 }
func (a WindowCoveringType) IsRollershadeExterior() bool       { return a == 0x02 }
func (a WindowCoveringType) IsRollershadeExterior2Motor() bool { return a == 0x03 }
func (a WindowCoveringType) IsDrapery() bool                   { return a == 0x04 }
func (a WindowCoveringType) IsAwning() bool                    { return a == 0x05 }
func (a WindowCoveringType) IsShutter() bool                   { return a == 0x06 }
func (a WindowCoveringType) IsTiltBlindTiltOnly() bool         { return a == 0x07 }
func (a WindowCoveringType) IsTiltBlindLiftAndTilt() bool      { return a == 0x08 }
func (a WindowCoveringType) IsProjectorScreen() bool           { return a == 0x09 }
func (a *WindowCoveringType) SetRollershade()                  { *a = 0x00 }
func (a *WindowCoveringType) SetRollershade2Motor()            { *a = 0x01 }
func (a *WindowCoveringType) SetRollershadeExterior()          { *a = 0x02 }
func (a *WindowCoveringType) SetRollershadeExterior2Motor()    { *a = 0x03 }
func (a *WindowCoveringType) SetDrapery()                      { *a = 0x04 }
func (a *WindowCoveringType) SetAwning()                       { *a = 0x05 }
func (a *WindowCoveringType) SetShutter()                      { *a = 0x06 }
func (a *WindowCoveringType) SetTiltBlindTiltOnly()            { *a = 0x07 }
func (a *WindowCoveringType) SetTiltBlindLiftAndTilt()         { *a = 0x08 }
func (a *WindowCoveringType) SetProjectorScreen()              { *a = 0x09 }

func (WindowCoveringType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Rollershade"},
		{Value: 0x01, Name: "Rollershade - 2 Motor"},
		{Value: 0x02, Name: "Rollershade - Exterior"},
		{Value: 0x03, Name: "Rollershade - Exterior - 2 Motor"},
		{Value: 0x04, Name: "Drapery"},
		{Value: 0x05, Name: "Awning"},
		{Value: 0x06, Name: "Shutter"},
		{Value: 0x07, Name: "Tilt Blind - Tilt Only"},
		{Value: 0x08, Name: "Tilt Blind - Lift and Tilt"},
		{Value: 0x09, Name: "Projector Screen"},
	}
}
