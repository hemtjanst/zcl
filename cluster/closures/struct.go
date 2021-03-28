package closures

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

const ConfigStatusAttr zcl.AttrID = 7

func (ConfigStatus) ID() zcl.AttrID   { return ConfigStatusAttr }
func (ConfigStatus) Readable() bool   { return true }
func (ConfigStatus) Writable() bool   { return false }
func (ConfigStatus) Reportable() bool { return false }
func (ConfigStatus) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ConfigStatus) AttrID() zcl.AttrID   { return v.ID() }
func (v ConfigStatus) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ConfigStatus) AttrValue() zcl.Val  { return v.Value() }

func (ConfigStatus) Name() string        { return `Config / Status` }
func (ConfigStatus) Description() string { return `` }

type ConfigStatus zcl.Zbmp8

func (v *ConfigStatus) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *ConfigStatus) Value() zcl.Val     { return v }

func (v ConfigStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *ConfigStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = ConfigStatus(*nt)
	return br, err
}

func (v ConfigStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *ConfigStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ConfigStatus(*a)
	return nil
}

func (v *ConfigStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = ConfigStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ConfigStatus) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v ConfigStatus) IsOperational() bool             { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v ConfigStatus) IsOnline() bool                  { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v ConfigStatus) IsCommandsReversed() bool        { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v ConfigStatus) IsLiftControlIsClosedLoop() bool { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v ConfigStatus) IsTiltControlIsClosedLoop() bool { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v ConfigStatus) IsLiftEncoderControlled() bool   { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v ConfigStatus) IsTiltEncoderControlled() bool   { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v *ConfigStatus) SetOperational(b bool)          { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *ConfigStatus) SetOnline(b bool)               { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }
func (v *ConfigStatus) SetCommandsReversed(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *ConfigStatus) SetLiftControlIsClosedLoop(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *ConfigStatus) SetTiltControlIsClosedLoop(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
}
func (v *ConfigStatus) SetLiftEncoderControlled(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b))
}
func (v *ConfigStatus) SetTiltEncoderControlled(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b))
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

// Implements AttrDef/AttrValue interfaces
func (v LiftAccelerationTime) AttrID() zcl.AttrID   { return v.ID() }
func (v LiftAccelerationTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LiftAccelerationTime) AttrValue() zcl.Val  { return v.Value() }

func (LiftAccelerationTime) Name() string        { return `Lift - Acceleration Time` }
func (LiftAccelerationTime) Description() string { return `` }

type LiftAccelerationTime zcl.Zu16

func (v *LiftAccelerationTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LiftAccelerationTime) Value() zcl.Val     { return v }

func (v LiftAccelerationTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LiftAccelerationTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LiftAccelerationTime(*nt)
	return br, err
}

func (v LiftAccelerationTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LiftAccelerationTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LiftAccelerationTime(*a)
	return nil
}

func (v *LiftAccelerationTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LiftAccelerationTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LiftAccelerationTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const LiftCurrentPositionAttr zcl.AttrID = 3

func (LiftCurrentPosition) ID() zcl.AttrID   { return LiftCurrentPositionAttr }
func (LiftCurrentPosition) Readable() bool   { return true }
func (LiftCurrentPosition) Writable() bool   { return false }
func (LiftCurrentPosition) Reportable() bool { return true }
func (LiftCurrentPosition) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LiftCurrentPosition) AttrID() zcl.AttrID   { return v.ID() }
func (v LiftCurrentPosition) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LiftCurrentPosition) AttrValue() zcl.Val  { return v.Value() }

func (LiftCurrentPosition) Name() string        { return `Lift - Current Position` }
func (LiftCurrentPosition) Description() string { return `` }

type LiftCurrentPosition zcl.Zu16

func (v *LiftCurrentPosition) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LiftCurrentPosition) Value() zcl.Val     { return v }

func (v LiftCurrentPosition) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LiftCurrentPosition) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LiftCurrentPosition(*nt)
	return br, err
}

func (v LiftCurrentPosition) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LiftCurrentPosition) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LiftCurrentPosition(*a)
	return nil
}

func (v *LiftCurrentPosition) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LiftCurrentPosition(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LiftCurrentPosition) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const LiftDecelerationTimeAttr zcl.AttrID = 22

func (LiftDecelerationTime) ID() zcl.AttrID   { return LiftDecelerationTimeAttr }
func (LiftDecelerationTime) Readable() bool   { return true }
func (LiftDecelerationTime) Writable() bool   { return true }
func (LiftDecelerationTime) Reportable() bool { return false }
func (LiftDecelerationTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LiftDecelerationTime) AttrID() zcl.AttrID   { return v.ID() }
func (v LiftDecelerationTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LiftDecelerationTime) AttrValue() zcl.Val  { return v.Value() }

func (LiftDecelerationTime) Name() string        { return `Lift - Deceleration Time` }
func (LiftDecelerationTime) Description() string { return `` }

type LiftDecelerationTime zcl.Zu16

func (v *LiftDecelerationTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LiftDecelerationTime) Value() zcl.Val     { return v }

func (v LiftDecelerationTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LiftDecelerationTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LiftDecelerationTime(*nt)
	return br, err
}

func (v LiftDecelerationTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LiftDecelerationTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LiftDecelerationTime(*a)
	return nil
}

func (v *LiftDecelerationTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LiftDecelerationTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LiftDecelerationTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const LiftInstalledClosedLimitAttr zcl.AttrID = 17

func (LiftInstalledClosedLimit) ID() zcl.AttrID   { return LiftInstalledClosedLimitAttr }
func (LiftInstalledClosedLimit) Readable() bool   { return true }
func (LiftInstalledClosedLimit) Writable() bool   { return false }
func (LiftInstalledClosedLimit) Reportable() bool { return false }
func (LiftInstalledClosedLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LiftInstalledClosedLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v LiftInstalledClosedLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LiftInstalledClosedLimit) AttrValue() zcl.Val  { return v.Value() }

func (LiftInstalledClosedLimit) Name() string        { return `Lift - Installed Closed Limit` }
func (LiftInstalledClosedLimit) Description() string { return `` }

type LiftInstalledClosedLimit zcl.Zu16

func (v *LiftInstalledClosedLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LiftInstalledClosedLimit) Value() zcl.Val     { return v }

func (v LiftInstalledClosedLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LiftInstalledClosedLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LiftInstalledClosedLimit(*nt)
	return br, err
}

func (v LiftInstalledClosedLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LiftInstalledClosedLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LiftInstalledClosedLimit(*a)
	return nil
}

func (v *LiftInstalledClosedLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LiftInstalledClosedLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LiftInstalledClosedLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const LiftInstalledOpenLimitAttr zcl.AttrID = 16

func (LiftInstalledOpenLimit) ID() zcl.AttrID   { return LiftInstalledOpenLimitAttr }
func (LiftInstalledOpenLimit) Readable() bool   { return true }
func (LiftInstalledOpenLimit) Writable() bool   { return false }
func (LiftInstalledOpenLimit) Reportable() bool { return false }
func (LiftInstalledOpenLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LiftInstalledOpenLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v LiftInstalledOpenLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LiftInstalledOpenLimit) AttrValue() zcl.Val  { return v.Value() }

func (LiftInstalledOpenLimit) Name() string        { return `Lift - Installed Open Limit` }
func (LiftInstalledOpenLimit) Description() string { return `` }

type LiftInstalledOpenLimit zcl.Zu16

func (v *LiftInstalledOpenLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LiftInstalledOpenLimit) Value() zcl.Val     { return v }

func (v LiftInstalledOpenLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LiftInstalledOpenLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LiftInstalledOpenLimit(*nt)
	return br, err
}

func (v LiftInstalledOpenLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LiftInstalledOpenLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LiftInstalledOpenLimit(*a)
	return nil
}

func (v *LiftInstalledOpenLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LiftInstalledOpenLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LiftInstalledOpenLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const LiftIntermediateSetpointsAttr zcl.AttrID = 24

func (LiftIntermediateSetpoints) ID() zcl.AttrID   { return LiftIntermediateSetpointsAttr }
func (LiftIntermediateSetpoints) Readable() bool   { return true }
func (LiftIntermediateSetpoints) Writable() bool   { return true }
func (LiftIntermediateSetpoints) Reportable() bool { return false }
func (LiftIntermediateSetpoints) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LiftIntermediateSetpoints) AttrID() zcl.AttrID   { return v.ID() }
func (v LiftIntermediateSetpoints) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LiftIntermediateSetpoints) AttrValue() zcl.Val  { return v.Value() }

func (LiftIntermediateSetpoints) Name() string        { return `Lift - Intermediate Setpoints` }
func (LiftIntermediateSetpoints) Description() string { return `` }

type LiftIntermediateSetpoints zcl.Zostring

func (v *LiftIntermediateSetpoints) TypeID() zcl.TypeID { return new(zcl.Zostring).TypeID() }
func (v *LiftIntermediateSetpoints) Value() zcl.Val     { return v }

func (v LiftIntermediateSetpoints) MarshalZcl() ([]byte, error) { return zcl.Zostring(v).MarshalZcl() }

func (v *LiftIntermediateSetpoints) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*v = LiftIntermediateSetpoints(*nt)
	return br, err
}

func (v LiftIntermediateSetpoints) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zostring(v))
}

func (v *LiftIntermediateSetpoints) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zostring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LiftIntermediateSetpoints(*a)
	return nil
}

func (v *LiftIntermediateSetpoints) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zostring); ok {
		*v = LiftIntermediateSetpoints(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LiftIntermediateSetpoints) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(v))
}

const LiftNumberOfActuationsAttr zcl.AttrID = 5

func (LiftNumberOfActuations) ID() zcl.AttrID   { return LiftNumberOfActuationsAttr }
func (LiftNumberOfActuations) Readable() bool   { return true }
func (LiftNumberOfActuations) Writable() bool   { return false }
func (LiftNumberOfActuations) Reportable() bool { return false }
func (LiftNumberOfActuations) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LiftNumberOfActuations) AttrID() zcl.AttrID   { return v.ID() }
func (v LiftNumberOfActuations) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LiftNumberOfActuations) AttrValue() zcl.Val  { return v.Value() }

func (LiftNumberOfActuations) Name() string        { return `Lift - Number of Actuations` }
func (LiftNumberOfActuations) Description() string { return `` }

type LiftNumberOfActuations zcl.Zu16

func (v *LiftNumberOfActuations) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LiftNumberOfActuations) Value() zcl.Val     { return v }

func (v LiftNumberOfActuations) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LiftNumberOfActuations) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LiftNumberOfActuations(*nt)
	return br, err
}

func (v LiftNumberOfActuations) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LiftNumberOfActuations) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LiftNumberOfActuations(*a)
	return nil
}

func (v *LiftNumberOfActuations) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LiftNumberOfActuations(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LiftNumberOfActuations) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const LiftPhysicalClosedLimitAttr zcl.AttrID = 1

func (LiftPhysicalClosedLimit) ID() zcl.AttrID   { return LiftPhysicalClosedLimitAttr }
func (LiftPhysicalClosedLimit) Readable() bool   { return true }
func (LiftPhysicalClosedLimit) Writable() bool   { return false }
func (LiftPhysicalClosedLimit) Reportable() bool { return false }
func (LiftPhysicalClosedLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LiftPhysicalClosedLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v LiftPhysicalClosedLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LiftPhysicalClosedLimit) AttrValue() zcl.Val  { return v.Value() }

func (LiftPhysicalClosedLimit) Name() string        { return `Lift - Physical Closed Limit` }
func (LiftPhysicalClosedLimit) Description() string { return `` }

type LiftPhysicalClosedLimit zcl.Zu16

func (v *LiftPhysicalClosedLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LiftPhysicalClosedLimit) Value() zcl.Val     { return v }

func (v LiftPhysicalClosedLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LiftPhysicalClosedLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LiftPhysicalClosedLimit(*nt)
	return br, err
}

func (v LiftPhysicalClosedLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LiftPhysicalClosedLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LiftPhysicalClosedLimit(*a)
	return nil
}

func (v *LiftPhysicalClosedLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LiftPhysicalClosedLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LiftPhysicalClosedLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const LiftVelocityAttr zcl.AttrID = 20

func (LiftVelocity) ID() zcl.AttrID   { return LiftVelocityAttr }
func (LiftVelocity) Readable() bool   { return true }
func (LiftVelocity) Writable() bool   { return true }
func (LiftVelocity) Reportable() bool { return false }
func (LiftVelocity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LiftVelocity) AttrID() zcl.AttrID   { return v.ID() }
func (v LiftVelocity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LiftVelocity) AttrValue() zcl.Val  { return v.Value() }

func (LiftVelocity) Name() string        { return `Lift - Velocity` }
func (LiftVelocity) Description() string { return `` }

type LiftVelocity zcl.Zu16

func (v *LiftVelocity) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LiftVelocity) Value() zcl.Val     { return v }

func (v LiftVelocity) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LiftVelocity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LiftVelocity(*nt)
	return br, err
}

func (v LiftVelocity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LiftVelocity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LiftVelocity(*a)
	return nil
}

func (v *LiftVelocity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LiftVelocity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LiftVelocity) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const LiftCurrentPositionPercentageAttr zcl.AttrID = 8

func (LiftCurrentPositionPercentage) ID() zcl.AttrID   { return LiftCurrentPositionPercentageAttr }
func (LiftCurrentPositionPercentage) Readable() bool   { return true }
func (LiftCurrentPositionPercentage) Writable() bool   { return false }
func (LiftCurrentPositionPercentage) Reportable() bool { return true }
func (LiftCurrentPositionPercentage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LiftCurrentPositionPercentage) AttrID() zcl.AttrID   { return v.ID() }
func (v LiftCurrentPositionPercentage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LiftCurrentPositionPercentage) AttrValue() zcl.Val  { return v.Value() }

func (LiftCurrentPositionPercentage) Name() string        { return `Lift Current Position Percentage` }
func (LiftCurrentPositionPercentage) Description() string { return `` }

type LiftCurrentPositionPercentage zcl.Zu8

func (v *LiftCurrentPositionPercentage) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *LiftCurrentPositionPercentage) Value() zcl.Val     { return v }

func (v LiftCurrentPositionPercentage) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *LiftCurrentPositionPercentage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = LiftCurrentPositionPercentage(*nt)
	return br, err
}

func (v LiftCurrentPositionPercentage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *LiftCurrentPositionPercentage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LiftCurrentPositionPercentage(*a)
	return nil
}

func (v *LiftCurrentPositionPercentage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = LiftCurrentPositionPercentage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LiftCurrentPositionPercentage) String() string {
	return zcl.Percent.Format(float64(v))
}

func (Percentage) Name() string        { return `Percentage` }
func (Percentage) Description() string { return `` }

type Percentage zcl.Zu8

func (v *Percentage) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Percentage) Value() zcl.Val     { return v }

func (v Percentage) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Percentage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Percentage(*nt)
	return br, err
}

func (v Percentage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Percentage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Percentage(*a)
	return nil
}

func (v *Percentage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Percentage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Percentage) String() string {
	return zcl.Percent.Format(float64(v))
}

func (Position) Name() string        { return `Position` }
func (Position) Description() string { return `` }

type Position zcl.Zu16

func (v *Position) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Position) Value() zcl.Val     { return v }

func (v Position) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Position) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Position(*nt)
	return br, err
}

func (v Position) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Position) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Position(*a)
	return nil
}

func (v *Position) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Position(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Position) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const TiltCurrentPositionAttr zcl.AttrID = 4

func (TiltCurrentPosition) ID() zcl.AttrID   { return TiltCurrentPositionAttr }
func (TiltCurrentPosition) Readable() bool   { return true }
func (TiltCurrentPosition) Writable() bool   { return false }
func (TiltCurrentPosition) Reportable() bool { return true }
func (TiltCurrentPosition) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TiltCurrentPosition) AttrID() zcl.AttrID   { return v.ID() }
func (v TiltCurrentPosition) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TiltCurrentPosition) AttrValue() zcl.Val  { return v.Value() }

func (TiltCurrentPosition) Name() string        { return `Tilt - Current Position` }
func (TiltCurrentPosition) Description() string { return `` }

type TiltCurrentPosition zcl.Zu16

func (v *TiltCurrentPosition) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TiltCurrentPosition) Value() zcl.Val     { return v }

func (v TiltCurrentPosition) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *TiltCurrentPosition) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TiltCurrentPosition(*nt)
	return br, err
}

func (v TiltCurrentPosition) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TiltCurrentPosition) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TiltCurrentPosition(*a)
	return nil
}

func (v *TiltCurrentPosition) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TiltCurrentPosition(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TiltCurrentPosition) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const TiltIntermediateSetpointsAttr zcl.AttrID = 25

func (TiltIntermediateSetpoints) ID() zcl.AttrID   { return TiltIntermediateSetpointsAttr }
func (TiltIntermediateSetpoints) Readable() bool   { return true }
func (TiltIntermediateSetpoints) Writable() bool   { return true }
func (TiltIntermediateSetpoints) Reportable() bool { return false }
func (TiltIntermediateSetpoints) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TiltIntermediateSetpoints) AttrID() zcl.AttrID   { return v.ID() }
func (v TiltIntermediateSetpoints) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TiltIntermediateSetpoints) AttrValue() zcl.Val  { return v.Value() }

func (TiltIntermediateSetpoints) Name() string        { return `Tilt - Intermediate Setpoints` }
func (TiltIntermediateSetpoints) Description() string { return `` }

type TiltIntermediateSetpoints zcl.Zostring

func (v *TiltIntermediateSetpoints) TypeID() zcl.TypeID { return new(zcl.Zostring).TypeID() }
func (v *TiltIntermediateSetpoints) Value() zcl.Val     { return v }

func (v TiltIntermediateSetpoints) MarshalZcl() ([]byte, error) { return zcl.Zostring(v).MarshalZcl() }

func (v *TiltIntermediateSetpoints) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*v = TiltIntermediateSetpoints(*nt)
	return br, err
}

func (v TiltIntermediateSetpoints) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zostring(v))
}

func (v *TiltIntermediateSetpoints) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zostring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TiltIntermediateSetpoints(*a)
	return nil
}

func (v *TiltIntermediateSetpoints) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zostring); ok {
		*v = TiltIntermediateSetpoints(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TiltIntermediateSetpoints) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(v))
}

const TiltNumberOfActuationsAttr zcl.AttrID = 6

func (TiltNumberOfActuations) ID() zcl.AttrID   { return TiltNumberOfActuationsAttr }
func (TiltNumberOfActuations) Readable() bool   { return true }
func (TiltNumberOfActuations) Writable() bool   { return false }
func (TiltNumberOfActuations) Reportable() bool { return false }
func (TiltNumberOfActuations) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TiltNumberOfActuations) AttrID() zcl.AttrID   { return v.ID() }
func (v TiltNumberOfActuations) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TiltNumberOfActuations) AttrValue() zcl.Val  { return v.Value() }

func (TiltNumberOfActuations) Name() string        { return `Tilt - Number of Actuations` }
func (TiltNumberOfActuations) Description() string { return `` }

type TiltNumberOfActuations zcl.Zu16

func (v *TiltNumberOfActuations) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TiltNumberOfActuations) Value() zcl.Val     { return v }

func (v TiltNumberOfActuations) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *TiltNumberOfActuations) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TiltNumberOfActuations(*nt)
	return br, err
}

func (v TiltNumberOfActuations) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TiltNumberOfActuations) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TiltNumberOfActuations(*a)
	return nil
}

func (v *TiltNumberOfActuations) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TiltNumberOfActuations(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TiltNumberOfActuations) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const TiltPhysicalClosedLimitAttr zcl.AttrID = 2

func (TiltPhysicalClosedLimit) ID() zcl.AttrID   { return TiltPhysicalClosedLimitAttr }
func (TiltPhysicalClosedLimit) Readable() bool   { return true }
func (TiltPhysicalClosedLimit) Writable() bool   { return false }
func (TiltPhysicalClosedLimit) Reportable() bool { return false }
func (TiltPhysicalClosedLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TiltPhysicalClosedLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v TiltPhysicalClosedLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TiltPhysicalClosedLimit) AttrValue() zcl.Val  { return v.Value() }

func (TiltPhysicalClosedLimit) Name() string        { return `Tilt - Physical Closed Limit` }
func (TiltPhysicalClosedLimit) Description() string { return `` }

type TiltPhysicalClosedLimit zcl.Zu16

func (v *TiltPhysicalClosedLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TiltPhysicalClosedLimit) Value() zcl.Val     { return v }

func (v TiltPhysicalClosedLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *TiltPhysicalClosedLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TiltPhysicalClosedLimit(*nt)
	return br, err
}

func (v TiltPhysicalClosedLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TiltPhysicalClosedLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TiltPhysicalClosedLimit(*a)
	return nil
}

func (v *TiltPhysicalClosedLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TiltPhysicalClosedLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TiltPhysicalClosedLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const TiltAInstalledOpenLimitAttr zcl.AttrID = 18

func (TiltAInstalledOpenLimit) ID() zcl.AttrID   { return TiltAInstalledOpenLimitAttr }
func (TiltAInstalledOpenLimit) Readable() bool   { return true }
func (TiltAInstalledOpenLimit) Writable() bool   { return false }
func (TiltAInstalledOpenLimit) Reportable() bool { return false }
func (TiltAInstalledOpenLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TiltAInstalledOpenLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v TiltAInstalledOpenLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TiltAInstalledOpenLimit) AttrValue() zcl.Val  { return v.Value() }

func (TiltAInstalledOpenLimit) Name() string        { return `Tilt A - Installed Open Limit` }
func (TiltAInstalledOpenLimit) Description() string { return `` }

type TiltAInstalledOpenLimit zcl.Zu16

func (v *TiltAInstalledOpenLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TiltAInstalledOpenLimit) Value() zcl.Val     { return v }

func (v TiltAInstalledOpenLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *TiltAInstalledOpenLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TiltAInstalledOpenLimit(*nt)
	return br, err
}

func (v TiltAInstalledOpenLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TiltAInstalledOpenLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TiltAInstalledOpenLimit(*a)
	return nil
}

func (v *TiltAInstalledOpenLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TiltAInstalledOpenLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TiltAInstalledOpenLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const TiltBInstalledOpenLimitAttr zcl.AttrID = 19

func (TiltBInstalledOpenLimit) ID() zcl.AttrID   { return TiltBInstalledOpenLimitAttr }
func (TiltBInstalledOpenLimit) Readable() bool   { return true }
func (TiltBInstalledOpenLimit) Writable() bool   { return false }
func (TiltBInstalledOpenLimit) Reportable() bool { return false }
func (TiltBInstalledOpenLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TiltBInstalledOpenLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v TiltBInstalledOpenLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TiltBInstalledOpenLimit) AttrValue() zcl.Val  { return v.Value() }

func (TiltBInstalledOpenLimit) Name() string        { return `Tilt B - Installed Open Limit` }
func (TiltBInstalledOpenLimit) Description() string { return `` }

type TiltBInstalledOpenLimit zcl.Zu16

func (v *TiltBInstalledOpenLimit) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TiltBInstalledOpenLimit) Value() zcl.Val     { return v }

func (v TiltBInstalledOpenLimit) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *TiltBInstalledOpenLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TiltBInstalledOpenLimit(*nt)
	return br, err
}

func (v TiltBInstalledOpenLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TiltBInstalledOpenLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TiltBInstalledOpenLimit(*a)
	return nil
}

func (v *TiltBInstalledOpenLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TiltBInstalledOpenLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TiltBInstalledOpenLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const TiltCurrentPositionPercentageAttr zcl.AttrID = 9

func (TiltCurrentPositionPercentage) ID() zcl.AttrID   { return TiltCurrentPositionPercentageAttr }
func (TiltCurrentPositionPercentage) Readable() bool   { return true }
func (TiltCurrentPositionPercentage) Writable() bool   { return false }
func (TiltCurrentPositionPercentage) Reportable() bool { return true }
func (TiltCurrentPositionPercentage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TiltCurrentPositionPercentage) AttrID() zcl.AttrID   { return v.ID() }
func (v TiltCurrentPositionPercentage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TiltCurrentPositionPercentage) AttrValue() zcl.Val  { return v.Value() }

func (TiltCurrentPositionPercentage) Name() string        { return `Tilt Current Position Percentage` }
func (TiltCurrentPositionPercentage) Description() string { return `` }

type TiltCurrentPositionPercentage zcl.Zu8

func (v *TiltCurrentPositionPercentage) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *TiltCurrentPositionPercentage) Value() zcl.Val     { return v }

func (v TiltCurrentPositionPercentage) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *TiltCurrentPositionPercentage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = TiltCurrentPositionPercentage(*nt)
	return br, err
}

func (v TiltCurrentPositionPercentage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *TiltCurrentPositionPercentage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TiltCurrentPositionPercentage(*a)
	return nil
}

func (v *TiltCurrentPositionPercentage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = TiltCurrentPositionPercentage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TiltCurrentPositionPercentage) String() string {
	return zcl.Percent.Format(float64(v))
}

const WindowCoveringModeAttr zcl.AttrID = 23

func (WindowCoveringMode) ID() zcl.AttrID   { return WindowCoveringModeAttr }
func (WindowCoveringMode) Readable() bool   { return true }
func (WindowCoveringMode) Writable() bool   { return true }
func (WindowCoveringMode) Reportable() bool { return false }
func (WindowCoveringMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v WindowCoveringMode) AttrID() zcl.AttrID   { return v.ID() }
func (v WindowCoveringMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *WindowCoveringMode) AttrValue() zcl.Val  { return v.Value() }

func (WindowCoveringMode) Name() string        { return `Window Covering Mode` }
func (WindowCoveringMode) Description() string { return `` }

type WindowCoveringMode zcl.Zbmp8

func (v *WindowCoveringMode) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *WindowCoveringMode) Value() zcl.Val     { return v }

func (v WindowCoveringMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *WindowCoveringMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = WindowCoveringMode(*nt)
	return br, err
}

func (v WindowCoveringMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *WindowCoveringMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = WindowCoveringMode(*a)
	return nil
}

func (v *WindowCoveringMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = WindowCoveringMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v WindowCoveringMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v WindowCoveringMode) IsReversed() bool        { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v WindowCoveringMode) IsCalibrationMode() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v WindowCoveringMode) IsMaintenanceMode() bool { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v WindowCoveringMode) IsLedFeedback() bool     { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v *WindowCoveringMode) SetReversed(b bool)     { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *WindowCoveringMode) SetCalibrationMode(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *WindowCoveringMode) SetMaintenanceMode(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *WindowCoveringMode) SetLedFeedback(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
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

// Implements AttrDef/AttrValue interfaces
func (v WindowCoveringType) AttrID() zcl.AttrID   { return v.ID() }
func (v WindowCoveringType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *WindowCoveringType) AttrValue() zcl.Val  { return v.Value() }

func (WindowCoveringType) Name() string        { return `Window Covering Type` }
func (WindowCoveringType) Description() string { return `` }

type WindowCoveringType zcl.Zenum8

func (v *WindowCoveringType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *WindowCoveringType) Value() zcl.Val     { return v }

func (v WindowCoveringType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *WindowCoveringType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = WindowCoveringType(*nt)
	return br, err
}

func (v WindowCoveringType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *WindowCoveringType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = WindowCoveringType(*a)
	return nil
}

func (v *WindowCoveringType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = WindowCoveringType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v WindowCoveringType) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v WindowCoveringType) IsRollershade() bool               { return v == 0x00 }
func (v WindowCoveringType) IsRollershade2Motor() bool         { return v == 0x01 }
func (v WindowCoveringType) IsRollershadeExterior() bool       { return v == 0x02 }
func (v WindowCoveringType) IsRollershadeExterior2Motor() bool { return v == 0x03 }
func (v WindowCoveringType) IsDrapery() bool                   { return v == 0x04 }
func (v WindowCoveringType) IsAwning() bool                    { return v == 0x05 }
func (v WindowCoveringType) IsShutter() bool                   { return v == 0x06 }
func (v WindowCoveringType) IsTiltBlindTiltOnly() bool         { return v == 0x07 }
func (v WindowCoveringType) IsTiltBlindLiftAndTilt() bool      { return v == 0x08 }
func (v WindowCoveringType) IsProjectorScreen() bool           { return v == 0x09 }
func (v *WindowCoveringType) SetRollershade()                  { *v = 0x00 }
func (v *WindowCoveringType) SetRollershade2Motor()            { *v = 0x01 }
func (v *WindowCoveringType) SetRollershadeExterior()          { *v = 0x02 }
func (v *WindowCoveringType) SetRollershadeExterior2Motor()    { *v = 0x03 }
func (v *WindowCoveringType) SetDrapery()                      { *v = 0x04 }
func (v *WindowCoveringType) SetAwning()                       { *v = 0x05 }
func (v *WindowCoveringType) SetShutter()                      { *v = 0x06 }
func (v *WindowCoveringType) SetTiltBlindTiltOnly()            { *v = 0x07 }
func (v *WindowCoveringType) SetTiltBlindLiftAndTilt()         { *v = 0x08 }
func (v *WindowCoveringType) SetProjectorScreen()              { *v = 0x09 }

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
