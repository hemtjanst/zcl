package ias

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

const CurrentZoneSensitivityLevelAttr zcl.AttrID = 19

func (CurrentZoneSensitivityLevel) ID() zcl.AttrID   { return CurrentZoneSensitivityLevelAttr }
func (CurrentZoneSensitivityLevel) Readable() bool   { return true }
func (CurrentZoneSensitivityLevel) Writable() bool   { return true }
func (CurrentZoneSensitivityLevel) Reportable() bool { return false }
func (CurrentZoneSensitivityLevel) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentZoneSensitivityLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentZoneSensitivityLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentZoneSensitivityLevel) AttrValue() zcl.Val  { return v.Value() }

func (CurrentZoneSensitivityLevel) Name() string        { return `Current Zone Sensitivity Level` }
func (CurrentZoneSensitivityLevel) Description() string { return `` }

type CurrentZoneSensitivityLevel zcl.Zu8

func (v *CurrentZoneSensitivityLevel) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *CurrentZoneSensitivityLevel) Value() zcl.Val     { return v }

func (v CurrentZoneSensitivityLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *CurrentZoneSensitivityLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentZoneSensitivityLevel(*nt)
	return br, err
}

func (v CurrentZoneSensitivityLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *CurrentZoneSensitivityLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentZoneSensitivityLevel(*a)
	return nil
}

func (v *CurrentZoneSensitivityLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = CurrentZoneSensitivityLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentZoneSensitivityLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (Delay) Name() string        { return `Delay` }
func (Delay) Description() string { return `` }

type Delay zcl.Zu16

func (v *Delay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Delay) Value() zcl.Val     { return v }

func (v Delay) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Delay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Delay(*nt)
	return br, err
}

func (v Delay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Delay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Delay(*a)
	return nil
}

func (v *Delay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Delay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Delay) String() string {
	return zcl.Seconds.Format(float64(v) / 4)
}

func (v Delay) Scaled() float64 {
	return float64(v) / 4
}

func (EnrollResponseCode) Name() string        { return `Enroll response code` }
func (EnrollResponseCode) Description() string { return `` }

type EnrollResponseCode zcl.Zenum8

func (v *EnrollResponseCode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *EnrollResponseCode) Value() zcl.Val     { return v }

func (v EnrollResponseCode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *EnrollResponseCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = EnrollResponseCode(*nt)
	return br, err
}

func (v EnrollResponseCode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *EnrollResponseCode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = EnrollResponseCode(*a)
	return nil
}

func (v *EnrollResponseCode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = EnrollResponseCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EnrollResponseCode) String() string {
	switch v {
	case 0x00:
		return "Success"
	case 0x01:
		return "Not Supported"
	case 0x02:
		return "Not permitted"
	case 0x03:
		return "Too many zones"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v EnrollResponseCode) IsSuccess() bool      { return v == 0x00 }
func (v EnrollResponseCode) IsNotSupported() bool { return v == 0x01 }
func (v EnrollResponseCode) IsNotPermitted() bool { return v == 0x02 }
func (v EnrollResponseCode) IsTooManyZones() bool { return v == 0x03 }
func (v *EnrollResponseCode) SetSuccess()         { *v = 0x00 }
func (v *EnrollResponseCode) SetNotSupported()    { *v = 0x01 }
func (v *EnrollResponseCode) SetNotPermitted()    { *v = 0x02 }
func (v *EnrollResponseCode) SetTooManyZones()    { *v = 0x03 }

func (EnrollResponseCode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Success"},
		{Value: 0x01, Name: "Not Supported"},
		{Value: 0x02, Name: "Not permitted"},
		{Value: 0x03, Name: "Too many zones"},
	}
}

func (ExtendedStatus) Name() string        { return `Extended Status` }
func (ExtendedStatus) Description() string { return `` }

type ExtendedStatus zcl.Zbmp8

func (v *ExtendedStatus) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *ExtendedStatus) Value() zcl.Val     { return v }

func (v ExtendedStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *ExtendedStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = ExtendedStatus(*nt)
	return br, err
}

func (v ExtendedStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *ExtendedStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ExtendedStatus(*a)
	return nil
}

func (v *ExtendedStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = ExtendedStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ExtendedStatus) String() string {
	return zcl.Sprintf("%v", zcl.Zbmp8(v))
}

const IasCieAddressMacAttr zcl.AttrID = 16

func (IasCieAddressMac) ID() zcl.AttrID   { return IasCieAddressMacAttr }
func (IasCieAddressMac) Readable() bool   { return true }
func (IasCieAddressMac) Writable() bool   { return true }
func (IasCieAddressMac) Reportable() bool { return false }
func (IasCieAddressMac) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IasCieAddressMac) AttrID() zcl.AttrID   { return v.ID() }
func (v IasCieAddressMac) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IasCieAddressMac) AttrValue() zcl.Val  { return v.Value() }

func (IasCieAddressMac) Name() string        { return `IAS CIE Address (MAC)` }
func (IasCieAddressMac) Description() string { return `` }

type IasCieAddressMac zcl.Zuid

func (v *IasCieAddressMac) TypeID() zcl.TypeID { return new(zcl.Zuid).TypeID() }
func (v *IasCieAddressMac) Value() zcl.Val     { return v }

func (v IasCieAddressMac) MarshalZcl() ([]byte, error) { return zcl.Zuid(v).MarshalZcl() }

func (v *IasCieAddressMac) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*v = IasCieAddressMac(*nt)
	return br, err
}

func (v IasCieAddressMac) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zuid(v))
}

func (v *IasCieAddressMac) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zuid)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IasCieAddressMac(*a)
	return nil
}

func (v *IasCieAddressMac) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zuid); ok {
		*v = IasCieAddressMac(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IasCieAddressMac) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(v))
}

func (ManufacturerCode) Name() string        { return `Manufacturer Code` }
func (ManufacturerCode) Description() string { return `` }

type ManufacturerCode zcl.Zu16

func (v *ManufacturerCode) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ManufacturerCode) Value() zcl.Val     { return v }

func (v ManufacturerCode) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ManufacturerCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ManufacturerCode(*nt)
	return br, err
}

func (v ManufacturerCode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ManufacturerCode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ManufacturerCode(*a)
	return nil
}

func (v *ManufacturerCode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ManufacturerCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ManufacturerCode) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const NumberOfZoneSensitivityLevelsSupportedAttr zcl.AttrID = 18

func (NumberOfZoneSensitivityLevelsSupported) ID() zcl.AttrID {
	return NumberOfZoneSensitivityLevelsSupportedAttr
}
func (NumberOfZoneSensitivityLevelsSupported) Readable() bool   { return false }
func (NumberOfZoneSensitivityLevelsSupported) Writable() bool   { return false }
func (NumberOfZoneSensitivityLevelsSupported) Reportable() bool { return false }
func (NumberOfZoneSensitivityLevelsSupported) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NumberOfZoneSensitivityLevelsSupported) AttrID() zcl.AttrID   { return v.ID() }
func (v NumberOfZoneSensitivityLevelsSupported) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NumberOfZoneSensitivityLevelsSupported) AttrValue() zcl.Val  { return v.Value() }

func (NumberOfZoneSensitivityLevelsSupported) Name() string {
	return `Number Of Zone Sensitivity Levels Supported`
}
func (NumberOfZoneSensitivityLevelsSupported) Description() string { return `` }

type NumberOfZoneSensitivityLevelsSupported zcl.Zu8

func (v *NumberOfZoneSensitivityLevelsSupported) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *NumberOfZoneSensitivityLevelsSupported) Value() zcl.Val     { return v }

func (v NumberOfZoneSensitivityLevelsSupported) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(v).MarshalZcl()
}

func (v *NumberOfZoneSensitivityLevelsSupported) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = NumberOfZoneSensitivityLevelsSupported(*nt)
	return br, err
}

func (v NumberOfZoneSensitivityLevelsSupported) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *NumberOfZoneSensitivityLevelsSupported) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NumberOfZoneSensitivityLevelsSupported(*a)
	return nil
}

func (v *NumberOfZoneSensitivityLevelsSupported) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = NumberOfZoneSensitivityLevelsSupported(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NumberOfZoneSensitivityLevelsSupported) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (TestModeDuration) Name() string        { return `Test mode duration` }
func (TestModeDuration) Description() string { return `` }

type TestModeDuration zcl.Zu8

func (v *TestModeDuration) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *TestModeDuration) Value() zcl.Val     { return v }

func (v TestModeDuration) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *TestModeDuration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = TestModeDuration(*nt)
	return br, err
}

func (v TestModeDuration) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *TestModeDuration) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TestModeDuration(*a)
	return nil
}

func (v *TestModeDuration) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = TestModeDuration(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TestModeDuration) String() string {
	return zcl.Seconds.Format(float64(v))
}

const ZoneIdAttr zcl.AttrID = 17

func (ZoneId) ID() zcl.AttrID   { return ZoneIdAttr }
func (ZoneId) Readable() bool   { return false }
func (ZoneId) Writable() bool   { return false }
func (ZoneId) Reportable() bool { return false }
func (ZoneId) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ZoneId) AttrID() zcl.AttrID   { return v.ID() }
func (v ZoneId) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ZoneId) AttrValue() zcl.Val  { return v.Value() }

func (ZoneId) Name() string        { return `Zone ID` }
func (ZoneId) Description() string { return `` }

type ZoneId zcl.Zu8

func (v *ZoneId) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ZoneId) Value() zcl.Val     { return v }

func (v ZoneId) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ZoneId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ZoneId(*nt)
	return br, err
}

func (v ZoneId) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ZoneId) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ZoneId(*a)
	return nil
}

func (v *ZoneId) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ZoneId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ZoneId) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const ZoneStateAttr zcl.AttrID = 0

func (ZoneState) ID() zcl.AttrID   { return ZoneStateAttr }
func (ZoneState) Readable() bool   { return false }
func (ZoneState) Writable() bool   { return false }
func (ZoneState) Reportable() bool { return true }
func (ZoneState) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ZoneState) AttrID() zcl.AttrID   { return v.ID() }
func (v ZoneState) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ZoneState) AttrValue() zcl.Val  { return v.Value() }

func (ZoneState) Name() string        { return `Zone State` }
func (ZoneState) Description() string { return `` }

type ZoneState zcl.Zenum8

func (v *ZoneState) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *ZoneState) Value() zcl.Val     { return v }

func (v ZoneState) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *ZoneState) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = ZoneState(*nt)
	return br, err
}

func (v ZoneState) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *ZoneState) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ZoneState(*a)
	return nil
}

func (v *ZoneState) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = ZoneState(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ZoneState) String() string {
	switch v {
	case 0x00:
		return "Not Enrolled"
	case 0x01:
		return "Enrolled"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v ZoneState) IsNotEnrolled() bool { return v == 0x00 }
func (v ZoneState) IsEnrolled() bool    { return v == 0x01 }
func (v *ZoneState) SetNotEnrolled()    { *v = 0x00 }
func (v *ZoneState) SetEnrolled()       { *v = 0x01 }

func (ZoneState) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Not Enrolled"},
		{Value: 0x01, Name: "Enrolled"},
	}
}

const ZoneStatusAttr zcl.AttrID = 2

func (ZoneStatus) ID() zcl.AttrID   { return ZoneStatusAttr }
func (ZoneStatus) Readable() bool   { return false }
func (ZoneStatus) Writable() bool   { return false }
func (ZoneStatus) Reportable() bool { return true }
func (ZoneStatus) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ZoneStatus) AttrID() zcl.AttrID   { return v.ID() }
func (v ZoneStatus) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ZoneStatus) AttrValue() zcl.Val  { return v.Value() }

func (ZoneStatus) Name() string        { return `Zone Status` }
func (ZoneStatus) Description() string { return `` }

type ZoneStatus zcl.Zbmp16

func (v *ZoneStatus) TypeID() zcl.TypeID { return new(zcl.Zbmp16).TypeID() }
func (v *ZoneStatus) Value() zcl.Val     { return v }

func (v ZoneStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(v).MarshalZcl() }

func (v *ZoneStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*v = ZoneStatus(*nt)
	return br, err
}

func (v ZoneStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp16(v))
}

func (v *ZoneStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ZoneStatus(*a)
	return nil
}

func (v *ZoneStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp16); ok {
		*v = ZoneStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ZoneStatus) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Test")
		case 1:
			bstr = append(bstr, "Battery defect")
		case 10:
			bstr = append(bstr, "Tamper")
		case 11:
			bstr = append(bstr, "Battery")
		case 12:
			bstr = append(bstr, "Supervision notify")
		case 13:
			bstr = append(bstr, "Restore notify")
		case 14:
			bstr = append(bstr, "Trouble")
		case 15:
			bstr = append(bstr, "AC (mains)")
		case 8:
			bstr = append(bstr, "Alarm1")
		case 9:
			bstr = append(bstr, "Alarm2")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v ZoneStatus) IsTest() bool              { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v ZoneStatus) IsBatteryDefect() bool     { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v ZoneStatus) IsTamper() bool            { return zcl.BitmapTest([]byte(v[:]), 10) }
func (v ZoneStatus) IsBattery() bool           { return zcl.BitmapTest([]byte(v[:]), 11) }
func (v ZoneStatus) IsSupervisionNotify() bool { return zcl.BitmapTest([]byte(v[:]), 12) }
func (v ZoneStatus) IsRestoreNotify() bool     { return zcl.BitmapTest([]byte(v[:]), 13) }
func (v ZoneStatus) IsTrouble() bool           { return zcl.BitmapTest([]byte(v[:]), 14) }
func (v ZoneStatus) IsAcMains() bool           { return zcl.BitmapTest([]byte(v[:]), 15) }
func (v ZoneStatus) IsAlarm1() bool            { return zcl.BitmapTest([]byte(v[:]), 8) }
func (v ZoneStatus) IsAlarm2() bool            { return zcl.BitmapTest([]byte(v[:]), 9) }
func (v *ZoneStatus) SetTest(b bool)           { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *ZoneStatus) SetBatteryDefect(b bool)  { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }
func (v *ZoneStatus) SetTamper(b bool)         { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 10, b)) }
func (v *ZoneStatus) SetBattery(b bool)        { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 11, b)) }
func (v *ZoneStatus) SetSupervisionNotify(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 12, b))
}
func (v *ZoneStatus) SetRestoreNotify(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 13, b)) }
func (v *ZoneStatus) SetTrouble(b bool)       { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 14, b)) }
func (v *ZoneStatus) SetAcMains(b bool)       { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 15, b)) }
func (v *ZoneStatus) SetAlarm1(b bool)        { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 8, b)) }
func (v *ZoneStatus) SetAlarm2(b bool)        { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 9, b)) }

func (ZoneStatus) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Test"},
		{Value: 1, Name: "Battery defect"},
		{Value: 10, Name: "Tamper"},
		{Value: 11, Name: "Battery"},
		{Value: 12, Name: "Supervision notify"},
		{Value: 13, Name: "Restore notify"},
		{Value: 14, Name: "Trouble"},
		{Value: 15, Name: "AC (mains)"},
		{Value: 8, Name: "Alarm1"},
		{Value: 9, Name: "Alarm2"},
	}
}

const ZoneTypeAttr zcl.AttrID = 1

func (ZoneType) ID() zcl.AttrID   { return ZoneTypeAttr }
func (ZoneType) Readable() bool   { return false }
func (ZoneType) Writable() bool   { return false }
func (ZoneType) Reportable() bool { return false }
func (ZoneType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ZoneType) AttrID() zcl.AttrID   { return v.ID() }
func (v ZoneType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ZoneType) AttrValue() zcl.Val  { return v.Value() }

func (ZoneType) Name() string        { return `Zone Type` }
func (ZoneType) Description() string { return `` }

type ZoneType zcl.Zenum16

func (v *ZoneType) TypeID() zcl.TypeID { return new(zcl.Zenum16).TypeID() }
func (v *ZoneType) Value() zcl.Val     { return v }

func (v ZoneType) MarshalZcl() ([]byte, error) { return zcl.Zenum16(v).MarshalZcl() }

func (v *ZoneType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*v = ZoneType(*nt)
	return br, err
}

func (v ZoneType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum16(v))
}

func (v *ZoneType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ZoneType(*a)
	return nil
}

func (v *ZoneType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum16); ok {
		*v = ZoneType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ZoneType) String() string {
	switch v {
	case 0x0000:
		return "Standard CIE"
	case 0x000d:
		return "Motion sensor"
	case 0x0015:
		return "Contact Switch"
	case 0x0016:
		return "Door/Window handle"
	case 0x0028:
		return "Fire sensor"
	case 0x002a:
		return "Water sensor"
	case 0x002b:
		return "Carbon Monoxide sensor"
	case 0x002c:
		return "Personal emergency device"
	case 0x002d:
		return "Viburation/Movement sensor"
	case 0x010f:
		return "Remote control"
	case 0x0115:
		return "Key fob"
	case 0x021d:
		return "Key pad"
	case 0x0225:
		return "Standard warning device"
	case 0x0226:
		return "Glass break sensor"
	case 0x0229:
		return "Security repeater"
	case 0xffff:
		return "Invalid zone type"
	}
	return zcl.Sprintf("%v", zcl.Zenum16(v))
}

func (v ZoneType) IsStandardCie() bool              { return v == 0x0000 }
func (v ZoneType) IsMotionSensor() bool             { return v == 0x000d }
func (v ZoneType) IsContactSwitch() bool            { return v == 0x0015 }
func (v ZoneType) IsDoorWindowHandle() bool         { return v == 0x0016 }
func (v ZoneType) IsFireSensor() bool               { return v == 0x0028 }
func (v ZoneType) IsWaterSensor() bool              { return v == 0x002a }
func (v ZoneType) IsCarbonMonoxideSensor() bool     { return v == 0x002b }
func (v ZoneType) IsPersonalEmergencyDevice() bool  { return v == 0x002c }
func (v ZoneType) IsViburationMovementSensor() bool { return v == 0x002d }
func (v ZoneType) IsRemoteControl() bool            { return v == 0x010f }
func (v ZoneType) IsKeyFob() bool                   { return v == 0x0115 }
func (v ZoneType) IsKeyPad() bool                   { return v == 0x021d }
func (v ZoneType) IsStandardWarningDevice() bool    { return v == 0x0225 }
func (v ZoneType) IsGlassBreakSensor() bool         { return v == 0x0226 }
func (v ZoneType) IsSecurityRepeater() bool         { return v == 0x0229 }
func (v ZoneType) IsInvalidZoneType() bool          { return v == 0xffff }
func (v *ZoneType) SetStandardCie()                 { *v = 0x0000 }
func (v *ZoneType) SetMotionSensor()                { *v = 0x000d }
func (v *ZoneType) SetContactSwitch()               { *v = 0x0015 }
func (v *ZoneType) SetDoorWindowHandle()            { *v = 0x0016 }
func (v *ZoneType) SetFireSensor()                  { *v = 0x0028 }
func (v *ZoneType) SetWaterSensor()                 { *v = 0x002a }
func (v *ZoneType) SetCarbonMonoxideSensor()        { *v = 0x002b }
func (v *ZoneType) SetPersonalEmergencyDevice()     { *v = 0x002c }
func (v *ZoneType) SetViburationMovementSensor()    { *v = 0x002d }
func (v *ZoneType) SetRemoteControl()               { *v = 0x010f }
func (v *ZoneType) SetKeyFob()                      { *v = 0x0115 }
func (v *ZoneType) SetKeyPad()                      { *v = 0x021d }
func (v *ZoneType) SetStandardWarningDevice()       { *v = 0x0225 }
func (v *ZoneType) SetGlassBreakSensor()            { *v = 0x0226 }
func (v *ZoneType) SetSecurityRepeater()            { *v = 0x0229 }
func (v *ZoneType) SetInvalidZoneType()             { *v = 0xffff }

func (ZoneType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x0000, Name: "Standard CIE"},
		{Value: 0x000d, Name: "Motion sensor"},
		{Value: 0x0015, Name: "Contact Switch"},
		{Value: 0x0016, Name: "Door/Window handle"},
		{Value: 0x0028, Name: "Fire sensor"},
		{Value: 0x002a, Name: "Water sensor"},
		{Value: 0x002b, Name: "Carbon Monoxide sensor"},
		{Value: 0x002c, Name: "Personal emergency device"},
		{Value: 0x002d, Name: "Viburation/Movement sensor"},
		{Value: 0x010f, Name: "Remote control"},
		{Value: 0x0115, Name: "Key fob"},
		{Value: 0x021d, Name: "Key pad"},
		{Value: 0x0225, Name: "Standard warning device"},
		{Value: 0x0226, Name: "Glass break sensor"},
		{Value: 0x0229, Name: "Security repeater"},
		{Value: 0xffff, Name: "Invalid zone type"},
	}
}
