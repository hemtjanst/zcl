// provides generic measurement and sensing interfaces
package measurement

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

const AcActivePowerOverloadAttr zcl.AttrID = 2051

func (AcActivePowerOverload) ID() zcl.AttrID   { return AcActivePowerOverloadAttr }
func (AcActivePowerOverload) Readable() bool   { return false }
func (AcActivePowerOverload) Writable() bool   { return false }
func (AcActivePowerOverload) Reportable() bool { return false }
func (AcActivePowerOverload) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcActivePowerOverload) AttrID() zcl.AttrID   { return v.ID() }
func (v AcActivePowerOverload) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcActivePowerOverload) AttrValue() zcl.Val  { return v.Value() }

func (AcActivePowerOverload) Name() string { return `AC Active Power Overload` }
func (AcActivePowerOverload) Description() string {
	return `specifies the alarm threshold, set by the manufacturer, for the maximum output active power supported by
device. The value is multiplied and divided by the ACPowerMultiplier and ACPowerDivisor, respectively.`
}

// AcActivePowerOverload specifies the alarm threshold, set by the manufacturer, for the maximum output active power supported by
// device. The value is multiplied and divided by the ACPowerMultiplier and ACPowerDivisor, respectively.
type AcActivePowerOverload zcl.Zs16

func (v *AcActivePowerOverload) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AcActivePowerOverload) Value() zcl.Val     { return v }

func (v AcActivePowerOverload) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AcActivePowerOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcActivePowerOverload(*nt)
	return br, err
}

func (v AcActivePowerOverload) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AcActivePowerOverload) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcActivePowerOverload(*a)
	return nil
}

func (v *AcActivePowerOverload) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AcActivePowerOverload(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcActivePowerOverload) String() string {
	return zcl.Watts.Format(float64(v))
}

const AcAlarmsMaskAttr zcl.AttrID = 2048

func (AcAlarmsMask) ID() zcl.AttrID   { return AcAlarmsMaskAttr }
func (AcAlarmsMask) Readable() bool   { return true }
func (AcAlarmsMask) Writable() bool   { return true }
func (AcAlarmsMask) Reportable() bool { return false }
func (AcAlarmsMask) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcAlarmsMask) AttrID() zcl.AttrID   { return v.ID() }
func (v AcAlarmsMask) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcAlarmsMask) AttrValue() zcl.Val  { return v.Value() }

func (AcAlarmsMask) Name() string { return `AC Alarms Mask` }
func (AcAlarmsMask) Description() string {
	return `specifies which configurable alarms may be generated`
}

// AcAlarmsMask specifies which configurable alarms may be generated
type AcAlarmsMask zcl.Zbmp16

func (v *AcAlarmsMask) TypeID() zcl.TypeID { return new(zcl.Zbmp16).TypeID() }
func (v *AcAlarmsMask) Value() zcl.Val     { return v }

func (v AcAlarmsMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(v).MarshalZcl() }

func (v *AcAlarmsMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcAlarmsMask(*nt)
	return br, err
}

func (v AcAlarmsMask) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp16(v))
}

func (v *AcAlarmsMask) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcAlarmsMask(*a)
	return nil
}

func (v *AcAlarmsMask) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp16); ok {
		*v = AcAlarmsMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcAlarmsMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Voltage Overload")
		case 1:
			bstr = append(bstr, "Current Overload")
		case 2:
			bstr = append(bstr, "Active Power Overload")
		case 3:
			bstr = append(bstr, "Reactive Power Overload")
		case 4:
			bstr = append(bstr, "Average RMS Over Voltage")
		case 5:
			bstr = append(bstr, "Average RMS Under Voltage")
		case 6:
			bstr = append(bstr, "RMS Extreme Over Voltage")
		case 7:
			bstr = append(bstr, "RMS Extreme Under Voltage")
		case 8:
			bstr = append(bstr, "RMS Voltage Sag")
		case 9:
			bstr = append(bstr, "RMS Voltage Swell")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v AcAlarmsMask) IsVoltageOverload() bool        { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v AcAlarmsMask) IsCurrentOverload() bool        { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v AcAlarmsMask) IsActivePowerOverload() bool    { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v AcAlarmsMask) IsReactivePowerOverload() bool  { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v AcAlarmsMask) IsAverageRmsOverVoltage() bool  { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v AcAlarmsMask) IsAverageRmsUnderVoltage() bool { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v AcAlarmsMask) IsRmsExtremeOverVoltage() bool  { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v AcAlarmsMask) IsRmsExtremeUnderVoltage() bool { return zcl.BitmapTest([]byte(v[:]), 7) }
func (v AcAlarmsMask) IsRmsVoltageSag() bool          { return zcl.BitmapTest([]byte(v[:]), 8) }
func (v AcAlarmsMask) IsRmsVoltageSwell() bool        { return zcl.BitmapTest([]byte(v[:]), 9) }
func (v *AcAlarmsMask) SetVoltageOverload(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *AcAlarmsMask) SetCurrentOverload(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *AcAlarmsMask) SetActivePowerOverload(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *AcAlarmsMask) SetReactivePowerOverload(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *AcAlarmsMask) SetAverageRmsOverVoltage(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
}
func (v *AcAlarmsMask) SetAverageRmsUnderVoltage(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b))
}
func (v *AcAlarmsMask) SetRmsExtremeOverVoltage(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b))
}
func (v *AcAlarmsMask) SetRmsExtremeUnderVoltage(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 7, b))
}
func (v *AcAlarmsMask) SetRmsVoltageSag(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 8, b)) }
func (v *AcAlarmsMask) SetRmsVoltageSwell(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 9, b))
}

func (AcAlarmsMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Voltage Overload"},
		{Value: 1, Name: "Current Overload"},
		{Value: 2, Name: "Active Power Overload"},
		{Value: 3, Name: "Reactive Power Overload"},
		{Value: 4, Name: "Average RMS Over Voltage"},
		{Value: 5, Name: "Average RMS Under Voltage"},
		{Value: 6, Name: "RMS Extreme Over Voltage"},
		{Value: 7, Name: "RMS Extreme Under Voltage"},
		{Value: 8, Name: "RMS Voltage Sag"},
		{Value: 9, Name: "RMS Voltage Swell"},
	}
}

const AcCurrentDivisorAttr zcl.AttrID = 1539

func (AcCurrentDivisor) ID() zcl.AttrID   { return AcCurrentDivisorAttr }
func (AcCurrentDivisor) Readable() bool   { return false }
func (AcCurrentDivisor) Writable() bool   { return false }
func (AcCurrentDivisor) Reportable() bool { return false }
func (AcCurrentDivisor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcCurrentDivisor) AttrID() zcl.AttrID   { return v.ID() }
func (v AcCurrentDivisor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcCurrentDivisor) AttrValue() zcl.Val  { return v.Value() }

func (AcCurrentDivisor) Name() string { return `AC Current Divisor` }
func (AcCurrentDivisor) Description() string {
	return `Provides a value to be divided against the ACCurrent, InstantaneousCurrent and RMSCurrent attributes.
This attribute must be used in conjunction with the ACCurrentMultiplier attribute. 0x0000 is an invalid value
for this attribute`
}

// AcCurrentDivisor Provides a value to be divided against the ACCurrent, InstantaneousCurrent and RMSCurrent attributes.
// This attribute must be used in conjunction with the ACCurrentMultiplier attribute. 0x0000 is an invalid value
// for this attribute
type AcCurrentDivisor zcl.Zu16

func (v *AcCurrentDivisor) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcCurrentDivisor) Value() zcl.Val     { return v }

func (v AcCurrentDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcCurrentDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcCurrentDivisor(*nt)
	return br, err
}

func (v AcCurrentDivisor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcCurrentDivisor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcCurrentDivisor(*a)
	return nil
}

func (v *AcCurrentDivisor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcCurrentDivisor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcCurrentDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AcCurrentMultiplierAttr zcl.AttrID = 1538

func (AcCurrentMultiplier) ID() zcl.AttrID   { return AcCurrentMultiplierAttr }
func (AcCurrentMultiplier) Readable() bool   { return false }
func (AcCurrentMultiplier) Writable() bool   { return false }
func (AcCurrentMultiplier) Reportable() bool { return false }
func (AcCurrentMultiplier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcCurrentMultiplier) AttrID() zcl.AttrID   { return v.ID() }
func (v AcCurrentMultiplier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcCurrentMultiplier) AttrValue() zcl.Val  { return v.Value() }

func (AcCurrentMultiplier) Name() string { return `AC Current Multiplier` }
func (AcCurrentMultiplier) Description() string {
	return `Provides a value to be multiplied against the InstantaneousCurrent and RMSCurrent attributes. This attribute
must be used in conjunction with the ACCurrentDivisor attribute. 0x0000 is an invalid value for this attribute.`
}

// AcCurrentMultiplier Provides a value to be multiplied against the InstantaneousCurrent and RMSCurrent attributes. This attribute
// must be used in conjunction with the ACCurrentDivisor attribute. 0x0000 is an invalid value for this attribute.
type AcCurrentMultiplier zcl.Zu16

func (v *AcCurrentMultiplier) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcCurrentMultiplier) Value() zcl.Val     { return v }

func (v AcCurrentMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcCurrentMultiplier(*nt)
	return br, err
}

func (v AcCurrentMultiplier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcCurrentMultiplier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcCurrentMultiplier(*a)
	return nil
}

func (v *AcCurrentMultiplier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcCurrentMultiplier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcCurrentMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AcCurrentOverloadAttr zcl.AttrID = 2050

func (AcCurrentOverload) ID() zcl.AttrID   { return AcCurrentOverloadAttr }
func (AcCurrentOverload) Readable() bool   { return false }
func (AcCurrentOverload) Writable() bool   { return false }
func (AcCurrentOverload) Reportable() bool { return false }
func (AcCurrentOverload) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcCurrentOverload) AttrID() zcl.AttrID   { return v.ID() }
func (v AcCurrentOverload) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcCurrentOverload) AttrValue() zcl.Val  { return v.Value() }

func (AcCurrentOverload) Name() string { return `AC Current Overload` }
func (AcCurrentOverload) Description() string {
	return `specifies the alarm threshold, set by the manufacturer, for the maximum output current supported by device.
The value is multiplied and divided by the ACCurrentMultiplier and ACCurrentDivider, respectively. The
value is current RMS.`
}

// AcCurrentOverload specifies the alarm threshold, set by the manufacturer, for the maximum output current supported by device.
// The value is multiplied and divided by the ACCurrentMultiplier and ACCurrentDivider, respectively. The
// value is current RMS.
type AcCurrentOverload zcl.Zs16

func (v *AcCurrentOverload) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AcCurrentOverload) Value() zcl.Val     { return v }

func (v AcCurrentOverload) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AcCurrentOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcCurrentOverload(*nt)
	return br, err
}

func (v AcCurrentOverload) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AcCurrentOverload) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcCurrentOverload(*a)
	return nil
}

func (v *AcCurrentOverload) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AcCurrentOverload(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcCurrentOverload) String() string {
	return zcl.Amperes.Format(float64(v))
}

const AcFrequencyAttr zcl.AttrID = 768

func (AcFrequency) ID() zcl.AttrID   { return AcFrequencyAttr }
func (AcFrequency) Readable() bool   { return false }
func (AcFrequency) Writable() bool   { return false }
func (AcFrequency) Reportable() bool { return false }
func (AcFrequency) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcFrequency) AttrID() zcl.AttrID   { return v.ID() }
func (v AcFrequency) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcFrequency) AttrValue() zcl.Val  { return v.Value() }

func (AcFrequency) Name() string { return `AC Frequency` }
func (AcFrequency) Description() string {
	return `represents the most recent AC Frequency reading in Hertz (Hz). If the frequency
cannot be measured, a value of 0xFFFF is returned.`
}

// AcFrequency represents the most recent AC Frequency reading in Hertz (Hz). If the frequency
// cannot be measured, a value of 0xFFFF is returned.
type AcFrequency zcl.Zu16

func (v *AcFrequency) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcFrequency) Value() zcl.Val     { return v }

func (v AcFrequency) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcFrequency) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcFrequency(*nt)
	return br, err
}

func (v AcFrequency) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcFrequency) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcFrequency(*a)
	return nil
}

func (v *AcFrequency) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcFrequency(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcFrequency) String() string {
	return zcl.Hertz.Format(float64(v))
}

const AcFrequencyDivisorAttr zcl.AttrID = 1025

func (AcFrequencyDivisor) ID() zcl.AttrID   { return AcFrequencyDivisorAttr }
func (AcFrequencyDivisor) Readable() bool   { return false }
func (AcFrequencyDivisor) Writable() bool   { return false }
func (AcFrequencyDivisor) Reportable() bool { return false }
func (AcFrequencyDivisor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcFrequencyDivisor) AttrID() zcl.AttrID   { return v.ID() }
func (v AcFrequencyDivisor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcFrequencyDivisor) AttrValue() zcl.Val  { return v.Value() }

func (AcFrequencyDivisor) Name() string { return `AC Frequency Divisor` }
func (AcFrequencyDivisor) Description() string {
	return `Provides a value to be divided against the ACFrequency attribute. This attribute must be used in conjunction
with the ACFrequencyMultiplier attribute. 0x0000 is an invalid value for this attribute.`
}

// AcFrequencyDivisor Provides a value to be divided against the ACFrequency attribute. This attribute must be used in conjunction
// with the ACFrequencyMultiplier attribute. 0x0000 is an invalid value for this attribute.
type AcFrequencyDivisor zcl.Zu16

func (v *AcFrequencyDivisor) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcFrequencyDivisor) Value() zcl.Val     { return v }

func (v AcFrequencyDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcFrequencyDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcFrequencyDivisor(*nt)
	return br, err
}

func (v AcFrequencyDivisor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcFrequencyDivisor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcFrequencyDivisor(*a)
	return nil
}

func (v *AcFrequencyDivisor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcFrequencyDivisor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcFrequencyDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AcFrequencyMaxAttr zcl.AttrID = 770

func (AcFrequencyMax) ID() zcl.AttrID   { return AcFrequencyMaxAttr }
func (AcFrequencyMax) Readable() bool   { return false }
func (AcFrequencyMax) Writable() bool   { return false }
func (AcFrequencyMax) Reportable() bool { return false }
func (AcFrequencyMax) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcFrequencyMax) AttrID() zcl.AttrID   { return v.ID() }
func (v AcFrequencyMax) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcFrequencyMax) AttrValue() zcl.Val  { return v.Value() }

func (AcFrequencyMax) Name() string { return `AC Frequency Max` }
func (AcFrequencyMax) Description() string {
	return `attribute represents the highest AC Frequency value measured in Hertz (Hz). After
resetting, this attribute will return a value of 0xFFFF until a measurement is made.`
}

// AcFrequencyMax attribute represents the highest AC Frequency value measured in Hertz (Hz). After
// resetting, this attribute will return a value of 0xFFFF until a measurement is made.
type AcFrequencyMax zcl.Zu16

func (v *AcFrequencyMax) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcFrequencyMax) Value() zcl.Val     { return v }

func (v AcFrequencyMax) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcFrequencyMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcFrequencyMax(*nt)
	return br, err
}

func (v AcFrequencyMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcFrequencyMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcFrequencyMax(*a)
	return nil
}

func (v *AcFrequencyMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcFrequencyMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcFrequencyMax) String() string {
	return zcl.Hertz.Format(float64(v))
}

const AcFrequencyMinAttr zcl.AttrID = 769

func (AcFrequencyMin) ID() zcl.AttrID   { return AcFrequencyMinAttr }
func (AcFrequencyMin) Readable() bool   { return false }
func (AcFrequencyMin) Writable() bool   { return false }
func (AcFrequencyMin) Reportable() bool { return false }
func (AcFrequencyMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcFrequencyMin) AttrID() zcl.AttrID   { return v.ID() }
func (v AcFrequencyMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcFrequencyMin) AttrValue() zcl.Val  { return v.Value() }

func (AcFrequencyMin) Name() string { return `AC Frequency Min` }
func (AcFrequencyMin) Description() string {
	return `attribute represents the lowest AC Frequency value measured in Hertz (Hz). After
resetting, this attribute will return a value of 0xFFFF until a measurement is made.`
}

// AcFrequencyMin attribute represents the lowest AC Frequency value measured in Hertz (Hz). After
// resetting, this attribute will return a value of 0xFFFF until a measurement is made.
type AcFrequencyMin zcl.Zu16

func (v *AcFrequencyMin) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcFrequencyMin) Value() zcl.Val     { return v }

func (v AcFrequencyMin) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcFrequencyMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcFrequencyMin(*nt)
	return br, err
}

func (v AcFrequencyMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcFrequencyMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcFrequencyMin(*a)
	return nil
}

func (v *AcFrequencyMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcFrequencyMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcFrequencyMin) String() string {
	return zcl.Hertz.Format(float64(v))
}

const AcFrequencyMultiplierAttr zcl.AttrID = 1024

func (AcFrequencyMultiplier) ID() zcl.AttrID   { return AcFrequencyMultiplierAttr }
func (AcFrequencyMultiplier) Readable() bool   { return false }
func (AcFrequencyMultiplier) Writable() bool   { return false }
func (AcFrequencyMultiplier) Reportable() bool { return false }
func (AcFrequencyMultiplier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcFrequencyMultiplier) AttrID() zcl.AttrID   { return v.ID() }
func (v AcFrequencyMultiplier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcFrequencyMultiplier) AttrValue() zcl.Val  { return v.Value() }

func (AcFrequencyMultiplier) Name() string { return `AC Frequency Multiplier` }
func (AcFrequencyMultiplier) Description() string {
	return `Provides a value to be multiplied against the ACFrequency attribute.
This attribute must be used in conjunction with the ACFrequencyDivisor attribute.
0x0000 is an invalid value for this attribute.`
}

// AcFrequencyMultiplier Provides a value to be multiplied against the ACFrequency attribute.
// This attribute must be used in conjunction with the ACFrequencyDivisor attribute.
// 0x0000 is an invalid value for this attribute.
type AcFrequencyMultiplier zcl.Zu16

func (v *AcFrequencyMultiplier) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcFrequencyMultiplier) Value() zcl.Val     { return v }

func (v AcFrequencyMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcFrequencyMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcFrequencyMultiplier(*nt)
	return br, err
}

func (v AcFrequencyMultiplier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcFrequencyMultiplier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcFrequencyMultiplier(*a)
	return nil
}

func (v *AcFrequencyMultiplier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcFrequencyMultiplier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcFrequencyMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AcPowerDivisorAttr zcl.AttrID = 1541

func (AcPowerDivisor) ID() zcl.AttrID   { return AcPowerDivisorAttr }
func (AcPowerDivisor) Readable() bool   { return false }
func (AcPowerDivisor) Writable() bool   { return false }
func (AcPowerDivisor) Reportable() bool { return false }
func (AcPowerDivisor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcPowerDivisor) AttrID() zcl.AttrID   { return v.ID() }
func (v AcPowerDivisor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcPowerDivisor) AttrValue() zcl.Val  { return v.Value() }

func (AcPowerDivisor) Name() string { return `AC Power Divisor` }
func (AcPowerDivisor) Description() string {
	return `Provides a value to be divided against the InstantaneousPower and ActivePower attributes. This attribute
must be used in conjunction with the ACPowerMultiplier attribute. 0x0000 is an invalid value for this attribute.`
}

// AcPowerDivisor Provides a value to be divided against the InstantaneousPower and ActivePower attributes. This attribute
// must be used in conjunction with the ACPowerMultiplier attribute. 0x0000 is an invalid value for this attribute.
type AcPowerDivisor zcl.Zu16

func (v *AcPowerDivisor) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcPowerDivisor) Value() zcl.Val     { return v }

func (v AcPowerDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcPowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcPowerDivisor(*nt)
	return br, err
}

func (v AcPowerDivisor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcPowerDivisor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcPowerDivisor(*a)
	return nil
}

func (v *AcPowerDivisor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcPowerDivisor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcPowerDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AcPowerMultiplierAttr zcl.AttrID = 1540

func (AcPowerMultiplier) ID() zcl.AttrID   { return AcPowerMultiplierAttr }
func (AcPowerMultiplier) Readable() bool   { return false }
func (AcPowerMultiplier) Writable() bool   { return false }
func (AcPowerMultiplier) Reportable() bool { return false }
func (AcPowerMultiplier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcPowerMultiplier) AttrID() zcl.AttrID   { return v.ID() }
func (v AcPowerMultiplier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcPowerMultiplier) AttrValue() zcl.Val  { return v.Value() }

func (AcPowerMultiplier) Name() string { return `AC Power Multiplier` }
func (AcPowerMultiplier) Description() string {
	return `Provides a value to be multiplied against the InstantaneousPower and ActivePower attributes. This attribute
must be used in conjunction with the ACPowerDivisor attribute. 0x0000 is an invalid value for this attribute.`
}

// AcPowerMultiplier Provides a value to be multiplied against the InstantaneousPower and ActivePower attributes. This attribute
// must be used in conjunction with the ACPowerDivisor attribute. 0x0000 is an invalid value for this attribute.
type AcPowerMultiplier zcl.Zu16

func (v *AcPowerMultiplier) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcPowerMultiplier) Value() zcl.Val     { return v }

func (v AcPowerMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcPowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcPowerMultiplier(*nt)
	return br, err
}

func (v AcPowerMultiplier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcPowerMultiplier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcPowerMultiplier(*a)
	return nil
}

func (v *AcPowerMultiplier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcPowerMultiplier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcPowerMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AcReactivePowerOverloadAttr zcl.AttrID = 2052

func (AcReactivePowerOverload) ID() zcl.AttrID   { return AcReactivePowerOverloadAttr }
func (AcReactivePowerOverload) Readable() bool   { return false }
func (AcReactivePowerOverload) Writable() bool   { return false }
func (AcReactivePowerOverload) Reportable() bool { return false }
func (AcReactivePowerOverload) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcReactivePowerOverload) AttrID() zcl.AttrID   { return v.ID() }
func (v AcReactivePowerOverload) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcReactivePowerOverload) AttrValue() zcl.Val  { return v.Value() }

func (AcReactivePowerOverload) Name() string { return `AC Reactive Power Overload` }
func (AcReactivePowerOverload) Description() string {
	return `specifies the alarm threshold, set by the manufacturer, for the maximum output reactive power supported by
device. The value is multiplied and divided by the ACPowerMultiplier and ACPowerDivisor, respectively.`
}

// AcReactivePowerOverload specifies the alarm threshold, set by the manufacturer, for the maximum output reactive power supported by
// device. The value is multiplied and divided by the ACPowerMultiplier and ACPowerDivisor, respectively.
type AcReactivePowerOverload zcl.Zs16

func (v *AcReactivePowerOverload) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AcReactivePowerOverload) Value() zcl.Val     { return v }

func (v AcReactivePowerOverload) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AcReactivePowerOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcReactivePowerOverload(*nt)
	return br, err
}

func (v AcReactivePowerOverload) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AcReactivePowerOverload) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcReactivePowerOverload(*a)
	return nil
}

func (v *AcReactivePowerOverload) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AcReactivePowerOverload(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcReactivePowerOverload) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const AcVoltageDivisorAttr zcl.AttrID = 1537

func (AcVoltageDivisor) ID() zcl.AttrID   { return AcVoltageDivisorAttr }
func (AcVoltageDivisor) Readable() bool   { return false }
func (AcVoltageDivisor) Writable() bool   { return false }
func (AcVoltageDivisor) Reportable() bool { return false }
func (AcVoltageDivisor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcVoltageDivisor) AttrID() zcl.AttrID   { return v.ID() }
func (v AcVoltageDivisor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcVoltageDivisor) AttrValue() zcl.Val  { return v.Value() }

func (AcVoltageDivisor) Name() string { return `AC Voltage Divisor` }
func (AcVoltageDivisor) Description() string {
	return `Provides a value to be divided against the InstantaneousVoltage and RMSVoltage attributes. This attribute
must be used in conjunction with the ACVoltageMultiplier attribute. 0x0000 is an invalid value for this
attribute.`
}

// AcVoltageDivisor Provides a value to be divided against the InstantaneousVoltage and RMSVoltage attributes. This attribute
// must be used in conjunction with the ACVoltageMultiplier attribute. 0x0000 is an invalid value for this
// attribute.
type AcVoltageDivisor zcl.Zu16

func (v *AcVoltageDivisor) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcVoltageDivisor) Value() zcl.Val     { return v }

func (v AcVoltageDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcVoltageDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcVoltageDivisor(*nt)
	return br, err
}

func (v AcVoltageDivisor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcVoltageDivisor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcVoltageDivisor(*a)
	return nil
}

func (v *AcVoltageDivisor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcVoltageDivisor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcVoltageDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AcVoltageMultiplierAttr zcl.AttrID = 1536

func (AcVoltageMultiplier) ID() zcl.AttrID   { return AcVoltageMultiplierAttr }
func (AcVoltageMultiplier) Readable() bool   { return false }
func (AcVoltageMultiplier) Writable() bool   { return false }
func (AcVoltageMultiplier) Reportable() bool { return false }
func (AcVoltageMultiplier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcVoltageMultiplier) AttrID() zcl.AttrID   { return v.ID() }
func (v AcVoltageMultiplier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcVoltageMultiplier) AttrValue() zcl.Val  { return v.Value() }

func (AcVoltageMultiplier) Name() string { return `AC Voltage Multiplier` }
func (AcVoltageMultiplier) Description() string {
	return `Provides a value to be multiplied against the InstantaneousVoltage and RMSVoltage attributes. This attribute
must be used in conjunction with the ACVoltageDivisor attribute. 0x0000 is an invalid value for this attribute.`
}

// AcVoltageMultiplier Provides a value to be multiplied against the InstantaneousVoltage and RMSVoltage attributes. This attribute
// must be used in conjunction with the ACVoltageDivisor attribute. 0x0000 is an invalid value for this attribute.
type AcVoltageMultiplier zcl.Zu16

func (v *AcVoltageMultiplier) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcVoltageMultiplier) Value() zcl.Val     { return v }

func (v AcVoltageMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcVoltageMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcVoltageMultiplier(*nt)
	return br, err
}

func (v AcVoltageMultiplier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcVoltageMultiplier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcVoltageMultiplier(*a)
	return nil
}

func (v *AcVoltageMultiplier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcVoltageMultiplier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcVoltageMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AcVoltageOverloadAttr zcl.AttrID = 2049

func (AcVoltageOverload) ID() zcl.AttrID   { return AcVoltageOverloadAttr }
func (AcVoltageOverload) Readable() bool   { return false }
func (AcVoltageOverload) Writable() bool   { return false }
func (AcVoltageOverload) Reportable() bool { return false }
func (AcVoltageOverload) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcVoltageOverload) AttrID() zcl.AttrID   { return v.ID() }
func (v AcVoltageOverload) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcVoltageOverload) AttrValue() zcl.Val  { return v.Value() }

func (AcVoltageOverload) Name() string { return `AC Voltage Overload` }
func (AcVoltageOverload) Description() string {
	return `specifies the alarm threshold, set by the manufacturer, for the maximum output voltage supported by device.
The value is multiplied and divided by the ACVoltageMultiplier the ACVoltageDivisor, respectively. The
value is voltage RMS.`
}

// AcVoltageOverload specifies the alarm threshold, set by the manufacturer, for the maximum output voltage supported by device.
// The value is multiplied and divided by the ACVoltageMultiplier the ACVoltageDivisor, respectively. The
// value is voltage RMS.
type AcVoltageOverload zcl.Zs16

func (v *AcVoltageOverload) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AcVoltageOverload) Value() zcl.Val     { return v }

func (v AcVoltageOverload) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AcVoltageOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcVoltageOverload(*nt)
	return br, err
}

func (v AcVoltageOverload) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AcVoltageOverload) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcVoltageOverload(*a)
	return nil
}

func (v *AcVoltageOverload) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AcVoltageOverload(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcVoltageOverload) String() string {
	return zcl.Volts.Format(float64(v))
}

const ActiveCurrentAttr zcl.AttrID = 1282

func (ActiveCurrent) ID() zcl.AttrID   { return ActiveCurrentAttr }
func (ActiveCurrent) Readable() bool   { return false }
func (ActiveCurrent) Writable() bool   { return false }
func (ActiveCurrent) Reportable() bool { return false }
func (ActiveCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActiveCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v ActiveCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActiveCurrent) AttrValue() zcl.Val  { return v.Value() }

func (ActiveCurrent) Name() string { return `Active Current` }
func (ActiveCurrent) Description() string {
	return `Represents the single phase or Phase A, AC active/resistive current value at the moment in time the attribute
is read, in Amps (A). Positive values indicate power delivered to the premises where negative values indicate
power received from the premises.`
}

// ActiveCurrent Represents the single phase or Phase A, AC active/resistive current value at the moment in time the attribute
// is read, in Amps (A). Positive values indicate power delivered to the premises where negative values indicate
// power received from the premises.
type ActiveCurrent zcl.Zs16

func (v *ActiveCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActiveCurrent) Value() zcl.Val     { return v }

func (v ActiveCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActiveCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActiveCurrent(*nt)
	return br, err
}

func (v ActiveCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActiveCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActiveCurrent(*a)
	return nil
}

func (v *ActiveCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActiveCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActiveCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const ActiveCurrentPhBAttr zcl.AttrID = 2306

func (ActiveCurrentPhB) ID() zcl.AttrID   { return ActiveCurrentPhBAttr }
func (ActiveCurrentPhB) Readable() bool   { return false }
func (ActiveCurrentPhB) Writable() bool   { return false }
func (ActiveCurrentPhB) Reportable() bool { return false }
func (ActiveCurrentPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActiveCurrentPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v ActiveCurrentPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActiveCurrentPhB) AttrValue() zcl.Val  { return v.Value() }

func (ActiveCurrentPhB) Name() string { return `Active Current Ph B` }
func (ActiveCurrentPhB) Description() string {
	return `represents the Phase B, AC active/resistive current value at the moment in time the attribute is read.
Positive values indicate power delivered to the premises where negative values indicate power received
from the premises.`
}

// ActiveCurrentPhB represents the Phase B, AC active/resistive current value at the moment in time the attribute is read.
// Positive values indicate power delivered to the premises where negative values indicate power received
// from the premises.
type ActiveCurrentPhB zcl.Zs16

func (v *ActiveCurrentPhB) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActiveCurrentPhB) Value() zcl.Val     { return v }

func (v ActiveCurrentPhB) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActiveCurrentPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActiveCurrentPhB(*nt)
	return br, err
}

func (v ActiveCurrentPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActiveCurrentPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActiveCurrentPhB(*a)
	return nil
}

func (v *ActiveCurrentPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActiveCurrentPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActiveCurrentPhB) String() string {
	return zcl.Amperes.Format(float64(v))
}

const ActiveCurrentPhCAttr zcl.AttrID = 2562

func (ActiveCurrentPhC) ID() zcl.AttrID   { return ActiveCurrentPhCAttr }
func (ActiveCurrentPhC) Readable() bool   { return false }
func (ActiveCurrentPhC) Writable() bool   { return false }
func (ActiveCurrentPhC) Reportable() bool { return false }
func (ActiveCurrentPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActiveCurrentPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v ActiveCurrentPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActiveCurrentPhC) AttrValue() zcl.Val  { return v.Value() }

func (ActiveCurrentPhC) Name() string        { return `Active Current Ph C` }
func (ActiveCurrentPhC) Description() string { return `` }

type ActiveCurrentPhC zcl.Zs16

func (v *ActiveCurrentPhC) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActiveCurrentPhC) Value() zcl.Val     { return v }

func (v ActiveCurrentPhC) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActiveCurrentPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActiveCurrentPhC(*nt)
	return br, err
}

func (v ActiveCurrentPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActiveCurrentPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActiveCurrentPhC(*a)
	return nil
}

func (v *ActiveCurrentPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActiveCurrentPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActiveCurrentPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const ActivePowerAttr zcl.AttrID = 1291

func (ActivePower) ID() zcl.AttrID   { return ActivePowerAttr }
func (ActivePower) Readable() bool   { return false }
func (ActivePower) Writable() bool   { return false }
func (ActivePower) Reportable() bool { return true }
func (ActivePower) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActivePower) AttrID() zcl.AttrID   { return v.ID() }
func (v ActivePower) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActivePower) AttrValue() zcl.Val  { return v.Value() }

func (ActivePower) Name() string { return `Active Power` }
func (ActivePower) Description() string {
	return `Represents the single phase or Phase A, current demand of active power delivered or received at the premises.
Positive values indicate power delivered to the premises where negative values indicate power received from the
premises.`
}

// ActivePower Represents the single phase or Phase A, current demand of active power delivered or received at the premises.
// Positive values indicate power delivered to the premises where negative values indicate power received from the
// premises.
type ActivePower zcl.Zs16

func (v *ActivePower) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActivePower) Value() zcl.Val     { return v }

func (v ActivePower) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActivePower(*nt)
	return br, err
}

func (v ActivePower) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActivePower) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActivePower(*a)
	return nil
}

func (v *ActivePower) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActivePower(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActivePower) String() string {
	return zcl.Watts.Format(float64(v))
}

const ActivePowerMaxAttr zcl.AttrID = 1293

func (ActivePowerMax) ID() zcl.AttrID   { return ActivePowerMaxAttr }
func (ActivePowerMax) Readable() bool   { return false }
func (ActivePowerMax) Writable() bool   { return false }
func (ActivePowerMax) Reportable() bool { return false }
func (ActivePowerMax) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActivePowerMax) AttrID() zcl.AttrID   { return v.ID() }
func (v ActivePowerMax) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActivePowerMax) AttrValue() zcl.Val  { return v.Value() }

func (ActivePowerMax) Name() string        { return `Active Power Max` }
func (ActivePowerMax) Description() string { return `Represents the highest AC power value` }

// ActivePowerMax Represents the highest AC power value
type ActivePowerMax zcl.Zs16

func (v *ActivePowerMax) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActivePowerMax) Value() zcl.Val     { return v }

func (v ActivePowerMax) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActivePowerMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActivePowerMax(*nt)
	return br, err
}

func (v ActivePowerMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActivePowerMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActivePowerMax(*a)
	return nil
}

func (v *ActivePowerMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActivePowerMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActivePowerMax) String() string {
	return zcl.Watts.Format(float64(v))
}

const ActivePowerMaxPhBAttr zcl.AttrID = 2317

func (ActivePowerMaxPhB) ID() zcl.AttrID   { return ActivePowerMaxPhBAttr }
func (ActivePowerMaxPhB) Readable() bool   { return false }
func (ActivePowerMaxPhB) Writable() bool   { return false }
func (ActivePowerMaxPhB) Reportable() bool { return false }
func (ActivePowerMaxPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActivePowerMaxPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v ActivePowerMaxPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActivePowerMaxPhB) AttrValue() zcl.Val  { return v.Value() }

func (ActivePowerMaxPhB) Name() string        { return `Active Power Max Ph B` }
func (ActivePowerMaxPhB) Description() string { return `` }

type ActivePowerMaxPhB zcl.Zs16

func (v *ActivePowerMaxPhB) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActivePowerMaxPhB) Value() zcl.Val     { return v }

func (v ActivePowerMaxPhB) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActivePowerMaxPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActivePowerMaxPhB(*nt)
	return br, err
}

func (v ActivePowerMaxPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActivePowerMaxPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActivePowerMaxPhB(*a)
	return nil
}

func (v *ActivePowerMaxPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActivePowerMaxPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActivePowerMaxPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const ActivePowerMaxPhCAttr zcl.AttrID = 2573

func (ActivePowerMaxPhC) ID() zcl.AttrID   { return ActivePowerMaxPhCAttr }
func (ActivePowerMaxPhC) Readable() bool   { return false }
func (ActivePowerMaxPhC) Writable() bool   { return false }
func (ActivePowerMaxPhC) Reportable() bool { return false }
func (ActivePowerMaxPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActivePowerMaxPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v ActivePowerMaxPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActivePowerMaxPhC) AttrValue() zcl.Val  { return v.Value() }

func (ActivePowerMaxPhC) Name() string        { return `Active Power Max Ph C` }
func (ActivePowerMaxPhC) Description() string { return `` }

type ActivePowerMaxPhC zcl.Zs16

func (v *ActivePowerMaxPhC) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActivePowerMaxPhC) Value() zcl.Val     { return v }

func (v ActivePowerMaxPhC) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActivePowerMaxPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActivePowerMaxPhC(*nt)
	return br, err
}

func (v ActivePowerMaxPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActivePowerMaxPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActivePowerMaxPhC(*a)
	return nil
}

func (v *ActivePowerMaxPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActivePowerMaxPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActivePowerMaxPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const ActivePowerMinAttr zcl.AttrID = 1292

func (ActivePowerMin) ID() zcl.AttrID   { return ActivePowerMinAttr }
func (ActivePowerMin) Readable() bool   { return false }
func (ActivePowerMin) Writable() bool   { return false }
func (ActivePowerMin) Reportable() bool { return false }
func (ActivePowerMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActivePowerMin) AttrID() zcl.AttrID   { return v.ID() }
func (v ActivePowerMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActivePowerMin) AttrValue() zcl.Val  { return v.Value() }

func (ActivePowerMin) Name() string        { return `Active Power Min` }
func (ActivePowerMin) Description() string { return `Represents the lowest AC power value` }

// ActivePowerMin Represents the lowest AC power value
type ActivePowerMin zcl.Zs16

func (v *ActivePowerMin) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActivePowerMin) Value() zcl.Val     { return v }

func (v ActivePowerMin) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActivePowerMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActivePowerMin(*nt)
	return br, err
}

func (v ActivePowerMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActivePowerMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActivePowerMin(*a)
	return nil
}

func (v *ActivePowerMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActivePowerMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActivePowerMin) String() string {
	return zcl.Watts.Format(float64(v))
}

const ActivePowerMinPhBAttr zcl.AttrID = 2316

func (ActivePowerMinPhB) ID() zcl.AttrID   { return ActivePowerMinPhBAttr }
func (ActivePowerMinPhB) Readable() bool   { return false }
func (ActivePowerMinPhB) Writable() bool   { return false }
func (ActivePowerMinPhB) Reportable() bool { return false }
func (ActivePowerMinPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActivePowerMinPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v ActivePowerMinPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActivePowerMinPhB) AttrValue() zcl.Val  { return v.Value() }

func (ActivePowerMinPhB) Name() string        { return `Active Power Min Ph B` }
func (ActivePowerMinPhB) Description() string { return `` }

type ActivePowerMinPhB zcl.Zs16

func (v *ActivePowerMinPhB) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActivePowerMinPhB) Value() zcl.Val     { return v }

func (v ActivePowerMinPhB) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActivePowerMinPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActivePowerMinPhB(*nt)
	return br, err
}

func (v ActivePowerMinPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActivePowerMinPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActivePowerMinPhB(*a)
	return nil
}

func (v *ActivePowerMinPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActivePowerMinPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActivePowerMinPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const ActivePowerMinPhCAttr zcl.AttrID = 2572

func (ActivePowerMinPhC) ID() zcl.AttrID   { return ActivePowerMinPhCAttr }
func (ActivePowerMinPhC) Readable() bool   { return false }
func (ActivePowerMinPhC) Writable() bool   { return false }
func (ActivePowerMinPhC) Reportable() bool { return false }
func (ActivePowerMinPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActivePowerMinPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v ActivePowerMinPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActivePowerMinPhC) AttrValue() zcl.Val  { return v.Value() }

func (ActivePowerMinPhC) Name() string        { return `Active Power Min Ph C` }
func (ActivePowerMinPhC) Description() string { return `` }

type ActivePowerMinPhC zcl.Zs16

func (v *ActivePowerMinPhC) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActivePowerMinPhC) Value() zcl.Val     { return v }

func (v ActivePowerMinPhC) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActivePowerMinPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActivePowerMinPhC(*nt)
	return br, err
}

func (v ActivePowerMinPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActivePowerMinPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActivePowerMinPhC(*a)
	return nil
}

func (v *ActivePowerMinPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActivePowerMinPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActivePowerMinPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const ActivePowerPhBAttr zcl.AttrID = 2315

func (ActivePowerPhB) ID() zcl.AttrID   { return ActivePowerPhBAttr }
func (ActivePowerPhB) Readable() bool   { return false }
func (ActivePowerPhB) Writable() bool   { return false }
func (ActivePowerPhB) Reportable() bool { return false }
func (ActivePowerPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActivePowerPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v ActivePowerPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActivePowerPhB) AttrValue() zcl.Val  { return v.Value() }

func (ActivePowerPhB) Name() string        { return `Active Power Ph B` }
func (ActivePowerPhB) Description() string { return `` }

type ActivePowerPhB zcl.Zs16

func (v *ActivePowerPhB) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActivePowerPhB) Value() zcl.Val     { return v }

func (v ActivePowerPhB) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActivePowerPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActivePowerPhB(*nt)
	return br, err
}

func (v ActivePowerPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActivePowerPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActivePowerPhB(*a)
	return nil
}

func (v *ActivePowerPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActivePowerPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActivePowerPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const ActivePowerPhCAttr zcl.AttrID = 2571

func (ActivePowerPhC) ID() zcl.AttrID   { return ActivePowerPhCAttr }
func (ActivePowerPhC) Readable() bool   { return false }
func (ActivePowerPhC) Writable() bool   { return false }
func (ActivePowerPhC) Reportable() bool { return false }
func (ActivePowerPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ActivePowerPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v ActivePowerPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ActivePowerPhC) AttrValue() zcl.Val  { return v.Value() }

func (ActivePowerPhC) Name() string        { return `Active Power Ph C` }
func (ActivePowerPhC) Description() string { return `` }

type ActivePowerPhC zcl.Zs16

func (v *ActivePowerPhC) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ActivePowerPhC) Value() zcl.Val     { return v }

func (v ActivePowerPhC) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ActivePowerPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ActivePowerPhC(*nt)
	return br, err
}

func (v ActivePowerPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ActivePowerPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ActivePowerPhC(*a)
	return nil
}

func (v *ActivePowerPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ActivePowerPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ActivePowerPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const ApparentPowerAttr zcl.AttrID = 1295

func (ApparentPower) ID() zcl.AttrID   { return ApparentPowerAttr }
func (ApparentPower) Readable() bool   { return false }
func (ApparentPower) Writable() bool   { return false }
func (ApparentPower) Reportable() bool { return false }
func (ApparentPower) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApparentPower) AttrID() zcl.AttrID   { return v.ID() }
func (v ApparentPower) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApparentPower) AttrValue() zcl.Val  { return v.Value() }

func (ApparentPower) Name() string { return `Apparent Power` }
func (ApparentPower) Description() string {
	return `Represents the single phase or Phase A, current demand of apparent (Square root of active and reactive
power) power, in VA.`
}

// ApparentPower Represents the single phase or Phase A, current demand of apparent (Square root of active and reactive
// power) power, in VA.
type ApparentPower zcl.Zu16

func (v *ApparentPower) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApparentPower) Value() zcl.Val     { return v }

func (v ApparentPower) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApparentPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApparentPower(*nt)
	return br, err
}

func (v ApparentPower) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApparentPower) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApparentPower(*a)
	return nil
}

func (v *ApparentPower) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApparentPower(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApparentPower) String() string {
	return zcl.VoltAmperes.Format(float64(v))
}

const ApparentPowerPhBAttr zcl.AttrID = 2319

func (ApparentPowerPhB) ID() zcl.AttrID   { return ApparentPowerPhBAttr }
func (ApparentPowerPhB) Readable() bool   { return false }
func (ApparentPowerPhB) Writable() bool   { return false }
func (ApparentPowerPhB) Reportable() bool { return false }
func (ApparentPowerPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApparentPowerPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v ApparentPowerPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApparentPowerPhB) AttrValue() zcl.Val  { return v.Value() }

func (ApparentPowerPhB) Name() string        { return `Apparent Power Ph B` }
func (ApparentPowerPhB) Description() string { return `` }

type ApparentPowerPhB zcl.Zu16

func (v *ApparentPowerPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApparentPowerPhB) Value() zcl.Val     { return v }

func (v ApparentPowerPhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApparentPowerPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApparentPowerPhB(*nt)
	return br, err
}

func (v ApparentPowerPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApparentPowerPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApparentPowerPhB(*a)
	return nil
}

func (v *ApparentPowerPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApparentPowerPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApparentPowerPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ApparentPowerPhCAttr zcl.AttrID = 2575

func (ApparentPowerPhC) ID() zcl.AttrID   { return ApparentPowerPhCAttr }
func (ApparentPowerPhC) Readable() bool   { return false }
func (ApparentPowerPhC) Writable() bool   { return false }
func (ApparentPowerPhC) Reportable() bool { return false }
func (ApparentPowerPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ApparentPowerPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v ApparentPowerPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ApparentPowerPhC) AttrValue() zcl.Val  { return v.Value() }

func (ApparentPowerPhC) Name() string        { return `Apparent Power Ph C` }
func (ApparentPowerPhC) Description() string { return `` }

type ApparentPowerPhC zcl.Zu16

func (v *ApparentPowerPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ApparentPowerPhC) Value() zcl.Val     { return v }

func (v ApparentPowerPhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ApparentPowerPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ApparentPowerPhC(*nt)
	return br, err
}

func (v ApparentPowerPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ApparentPowerPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ApparentPowerPhC(*a)
	return nil
}

func (v *ApparentPowerPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ApparentPowerPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ApparentPowerPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (AttributeId) Name() string        { return `Attribute ID` }
func (AttributeId) Description() string { return `` }

type AttributeId zcl.Zu16

func (v *AttributeId) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AttributeId) Value() zcl.Val     { return v }

func (v AttributeId) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AttributeId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AttributeId(*nt)
	return br, err
}

func (v AttributeId) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AttributeId) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AttributeId(*a)
	return nil
}

func (v *AttributeId) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AttributeId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AttributeId) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AverageRmsOverVoltageCounterAttr zcl.AttrID = 1298

func (AverageRmsOverVoltageCounter) ID() zcl.AttrID   { return AverageRmsOverVoltageCounterAttr }
func (AverageRmsOverVoltageCounter) Readable() bool   { return true }
func (AverageRmsOverVoltageCounter) Writable() bool   { return true }
func (AverageRmsOverVoltageCounter) Reportable() bool { return false }
func (AverageRmsOverVoltageCounter) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsOverVoltageCounter) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsOverVoltageCounter) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsOverVoltageCounter) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsOverVoltageCounter) Name() string { return `Average RMS Over Voltage Counter` }
func (AverageRmsOverVoltageCounter) Description() string {
	return `is the number of times the average RMS voltage, has been above the AverageRMS OverVoltage threshold
since last reset. This counter may be reset by writing zero to the attribute.`
}

// AverageRmsOverVoltageCounter is the number of times the average RMS voltage, has been above the AverageRMS OverVoltage threshold
// since last reset. This counter may be reset by writing zero to the attribute.
type AverageRmsOverVoltageCounter zcl.Zu16

func (v *AverageRmsOverVoltageCounter) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AverageRmsOverVoltageCounter) Value() zcl.Val     { return v }

func (v AverageRmsOverVoltageCounter) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AverageRmsOverVoltageCounter) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsOverVoltageCounter(*nt)
	return br, err
}

func (v AverageRmsOverVoltageCounter) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AverageRmsOverVoltageCounter) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsOverVoltageCounter(*a)
	return nil
}

func (v *AverageRmsOverVoltageCounter) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AverageRmsOverVoltageCounter(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsOverVoltageCounter) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AverageRmsOverVoltageCounterPhBAttr zcl.AttrID = 2322

func (AverageRmsOverVoltageCounterPhB) ID() zcl.AttrID   { return AverageRmsOverVoltageCounterPhBAttr }
func (AverageRmsOverVoltageCounterPhB) Readable() bool   { return true }
func (AverageRmsOverVoltageCounterPhB) Writable() bool   { return true }
func (AverageRmsOverVoltageCounterPhB) Reportable() bool { return false }
func (AverageRmsOverVoltageCounterPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsOverVoltageCounterPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsOverVoltageCounterPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsOverVoltageCounterPhB) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsOverVoltageCounterPhB) Name() string        { return `Average RMS Over Voltage Counter Ph B` }
func (AverageRmsOverVoltageCounterPhB) Description() string { return `` }

type AverageRmsOverVoltageCounterPhB zcl.Zu16

func (v *AverageRmsOverVoltageCounterPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AverageRmsOverVoltageCounterPhB) Value() zcl.Val     { return v }

func (v AverageRmsOverVoltageCounterPhB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *AverageRmsOverVoltageCounterPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsOverVoltageCounterPhB(*nt)
	return br, err
}

func (v AverageRmsOverVoltageCounterPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AverageRmsOverVoltageCounterPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsOverVoltageCounterPhB(*a)
	return nil
}

func (v *AverageRmsOverVoltageCounterPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AverageRmsOverVoltageCounterPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsOverVoltageCounterPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AverageRmsOverVoltageCounterPhCAttr zcl.AttrID = 2578

func (AverageRmsOverVoltageCounterPhC) ID() zcl.AttrID   { return AverageRmsOverVoltageCounterPhCAttr }
func (AverageRmsOverVoltageCounterPhC) Readable() bool   { return true }
func (AverageRmsOverVoltageCounterPhC) Writable() bool   { return true }
func (AverageRmsOverVoltageCounterPhC) Reportable() bool { return false }
func (AverageRmsOverVoltageCounterPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsOverVoltageCounterPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsOverVoltageCounterPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsOverVoltageCounterPhC) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsOverVoltageCounterPhC) Name() string        { return `Average RMS Over Voltage Counter Ph C` }
func (AverageRmsOverVoltageCounterPhC) Description() string { return `` }

type AverageRmsOverVoltageCounterPhC zcl.Zu16

func (v *AverageRmsOverVoltageCounterPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AverageRmsOverVoltageCounterPhC) Value() zcl.Val     { return v }

func (v AverageRmsOverVoltageCounterPhC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *AverageRmsOverVoltageCounterPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsOverVoltageCounterPhC(*nt)
	return br, err
}

func (v AverageRmsOverVoltageCounterPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AverageRmsOverVoltageCounterPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsOverVoltageCounterPhC(*a)
	return nil
}

func (v *AverageRmsOverVoltageCounterPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AverageRmsOverVoltageCounterPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsOverVoltageCounterPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AverageRmsUnderVoltageCounterAttr zcl.AttrID = 1299

func (AverageRmsUnderVoltageCounter) ID() zcl.AttrID   { return AverageRmsUnderVoltageCounterAttr }
func (AverageRmsUnderVoltageCounter) Readable() bool   { return true }
func (AverageRmsUnderVoltageCounter) Writable() bool   { return true }
func (AverageRmsUnderVoltageCounter) Reportable() bool { return false }
func (AverageRmsUnderVoltageCounter) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsUnderVoltageCounter) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsUnderVoltageCounter) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsUnderVoltageCounter) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsUnderVoltageCounter) Name() string { return `Average RMS Under Voltage Counter` }
func (AverageRmsUnderVoltageCounter) Description() string {
	return `is the number of times the average RMS voltage, has been below the AverageRMS underVoltage threshold
since last reset. This counter may be reset by writing zero to the attribute.`
}

// AverageRmsUnderVoltageCounter is the number of times the average RMS voltage, has been below the AverageRMS underVoltage threshold
// since last reset. This counter may be reset by writing zero to the attribute.
type AverageRmsUnderVoltageCounter zcl.Zu16

func (v *AverageRmsUnderVoltageCounter) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AverageRmsUnderVoltageCounter) Value() zcl.Val     { return v }

func (v AverageRmsUnderVoltageCounter) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AverageRmsUnderVoltageCounter) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsUnderVoltageCounter(*nt)
	return br, err
}

func (v AverageRmsUnderVoltageCounter) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AverageRmsUnderVoltageCounter) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsUnderVoltageCounter(*a)
	return nil
}

func (v *AverageRmsUnderVoltageCounter) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AverageRmsUnderVoltageCounter(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsUnderVoltageCounter) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AverageRmsUnderVoltageCounterPhBAttr zcl.AttrID = 2323

func (AverageRmsUnderVoltageCounterPhB) ID() zcl.AttrID   { return AverageRmsUnderVoltageCounterPhBAttr }
func (AverageRmsUnderVoltageCounterPhB) Readable() bool   { return true }
func (AverageRmsUnderVoltageCounterPhB) Writable() bool   { return true }
func (AverageRmsUnderVoltageCounterPhB) Reportable() bool { return false }
func (AverageRmsUnderVoltageCounterPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsUnderVoltageCounterPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsUnderVoltageCounterPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsUnderVoltageCounterPhB) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsUnderVoltageCounterPhB) Name() string {
	return `Average RMS Under Voltage Counter Ph B`
}
func (AverageRmsUnderVoltageCounterPhB) Description() string { return `` }

type AverageRmsUnderVoltageCounterPhB zcl.Zu16

func (v *AverageRmsUnderVoltageCounterPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AverageRmsUnderVoltageCounterPhB) Value() zcl.Val     { return v }

func (v AverageRmsUnderVoltageCounterPhB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *AverageRmsUnderVoltageCounterPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsUnderVoltageCounterPhB(*nt)
	return br, err
}

func (v AverageRmsUnderVoltageCounterPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AverageRmsUnderVoltageCounterPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsUnderVoltageCounterPhB(*a)
	return nil
}

func (v *AverageRmsUnderVoltageCounterPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AverageRmsUnderVoltageCounterPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsUnderVoltageCounterPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AverageRmsUnderVoltageCounterPhCAttr zcl.AttrID = 2579

func (AverageRmsUnderVoltageCounterPhC) ID() zcl.AttrID   { return AverageRmsUnderVoltageCounterPhCAttr }
func (AverageRmsUnderVoltageCounterPhC) Readable() bool   { return true }
func (AverageRmsUnderVoltageCounterPhC) Writable() bool   { return true }
func (AverageRmsUnderVoltageCounterPhC) Reportable() bool { return false }
func (AverageRmsUnderVoltageCounterPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsUnderVoltageCounterPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsUnderVoltageCounterPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsUnderVoltageCounterPhC) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsUnderVoltageCounterPhC) Name() string {
	return `Average RMS Under Voltage Counter Ph C`
}
func (AverageRmsUnderVoltageCounterPhC) Description() string { return `` }

type AverageRmsUnderVoltageCounterPhC zcl.Zu16

func (v *AverageRmsUnderVoltageCounterPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AverageRmsUnderVoltageCounterPhC) Value() zcl.Val     { return v }

func (v AverageRmsUnderVoltageCounterPhC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *AverageRmsUnderVoltageCounterPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsUnderVoltageCounterPhC(*nt)
	return br, err
}

func (v AverageRmsUnderVoltageCounterPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AverageRmsUnderVoltageCounterPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsUnderVoltageCounterPhC(*a)
	return nil
}

func (v *AverageRmsUnderVoltageCounterPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AverageRmsUnderVoltageCounterPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsUnderVoltageCounterPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AverageRmsVoltageMeasurementPeriodAttr zcl.AttrID = 1297

func (AverageRmsVoltageMeasurementPeriod) ID() zcl.AttrID {
	return AverageRmsVoltageMeasurementPeriodAttr
}
func (AverageRmsVoltageMeasurementPeriod) Readable() bool   { return true }
func (AverageRmsVoltageMeasurementPeriod) Writable() bool   { return true }
func (AverageRmsVoltageMeasurementPeriod) Reportable() bool { return false }
func (AverageRmsVoltageMeasurementPeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsVoltageMeasurementPeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsVoltageMeasurementPeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsVoltageMeasurementPeriod) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsVoltageMeasurementPeriod) Name() string {
	return `Average RMS Voltage Measurement Period`
}
func (AverageRmsVoltageMeasurementPeriod) Description() string {
	return `is the period that the RMS voltage is averaged over.`
}

// AverageRmsVoltageMeasurementPeriod is the period that the RMS voltage is averaged over.
type AverageRmsVoltageMeasurementPeriod zcl.Zu16

func (v *AverageRmsVoltageMeasurementPeriod) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AverageRmsVoltageMeasurementPeriod) Value() zcl.Val     { return v }

func (v AverageRmsVoltageMeasurementPeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *AverageRmsVoltageMeasurementPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsVoltageMeasurementPeriod(*nt)
	return br, err
}

func (v AverageRmsVoltageMeasurementPeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AverageRmsVoltageMeasurementPeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsVoltageMeasurementPeriod(*a)
	return nil
}

func (v *AverageRmsVoltageMeasurementPeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AverageRmsVoltageMeasurementPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsVoltageMeasurementPeriod) String() string {
	return zcl.Seconds.Format(float64(v))
}

const AverageRmsVoltageMeasurementPeriodPhBAttr zcl.AttrID = 2321

func (AverageRmsVoltageMeasurementPeriodPhB) ID() zcl.AttrID {
	return AverageRmsVoltageMeasurementPeriodPhBAttr
}
func (AverageRmsVoltageMeasurementPeriodPhB) Readable() bool   { return true }
func (AverageRmsVoltageMeasurementPeriodPhB) Writable() bool   { return true }
func (AverageRmsVoltageMeasurementPeriodPhB) Reportable() bool { return false }
func (AverageRmsVoltageMeasurementPeriodPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsVoltageMeasurementPeriodPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsVoltageMeasurementPeriodPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsVoltageMeasurementPeriodPhB) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsVoltageMeasurementPeriodPhB) Name() string {
	return `Average RMS Voltage Measurement Period Ph B`
}
func (AverageRmsVoltageMeasurementPeriodPhB) Description() string { return `` }

type AverageRmsVoltageMeasurementPeriodPhB zcl.Zu16

func (v *AverageRmsVoltageMeasurementPeriodPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AverageRmsVoltageMeasurementPeriodPhB) Value() zcl.Val     { return v }

func (v AverageRmsVoltageMeasurementPeriodPhB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *AverageRmsVoltageMeasurementPeriodPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsVoltageMeasurementPeriodPhB(*nt)
	return br, err
}

func (v AverageRmsVoltageMeasurementPeriodPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AverageRmsVoltageMeasurementPeriodPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsVoltageMeasurementPeriodPhB(*a)
	return nil
}

func (v *AverageRmsVoltageMeasurementPeriodPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AverageRmsVoltageMeasurementPeriodPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsVoltageMeasurementPeriodPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AverageRmsVoltageMeasurementPeriodPhCAttr zcl.AttrID = 2577

func (AverageRmsVoltageMeasurementPeriodPhC) ID() zcl.AttrID {
	return AverageRmsVoltageMeasurementPeriodPhCAttr
}
func (AverageRmsVoltageMeasurementPeriodPhC) Readable() bool   { return true }
func (AverageRmsVoltageMeasurementPeriodPhC) Writable() bool   { return true }
func (AverageRmsVoltageMeasurementPeriodPhC) Reportable() bool { return false }
func (AverageRmsVoltageMeasurementPeriodPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsVoltageMeasurementPeriodPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsVoltageMeasurementPeriodPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsVoltageMeasurementPeriodPhC) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsVoltageMeasurementPeriodPhC) Name() string {
	return `Average RMS Voltage Measurement Period Ph C`
}
func (AverageRmsVoltageMeasurementPeriodPhC) Description() string { return `` }

type AverageRmsVoltageMeasurementPeriodPhC zcl.Zu16

func (v *AverageRmsVoltageMeasurementPeriodPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AverageRmsVoltageMeasurementPeriodPhC) Value() zcl.Val     { return v }

func (v AverageRmsVoltageMeasurementPeriodPhC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *AverageRmsVoltageMeasurementPeriodPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsVoltageMeasurementPeriodPhC(*nt)
	return br, err
}

func (v AverageRmsVoltageMeasurementPeriodPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AverageRmsVoltageMeasurementPeriodPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsVoltageMeasurementPeriodPhC(*a)
	return nil
}

func (v *AverageRmsVoltageMeasurementPeriodPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AverageRmsVoltageMeasurementPeriodPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsVoltageMeasurementPeriodPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const AverageRmsOverVoltageAttr zcl.AttrID = 2053

func (AverageRmsOverVoltage) ID() zcl.AttrID   { return AverageRmsOverVoltageAttr }
func (AverageRmsOverVoltage) Readable() bool   { return false }
func (AverageRmsOverVoltage) Writable() bool   { return false }
func (AverageRmsOverVoltage) Reportable() bool { return false }
func (AverageRmsOverVoltage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsOverVoltage) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsOverVoltage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsOverVoltage) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsOverVoltage) Name() string { return `Average RMS over-voltage` }
func (AverageRmsOverVoltage) Description() string {
	return `is the average RMS voltage above which an over voltage condition is reported.
The threshold shall be configurable within the specified operating range of the electricity meter.
The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively.`
}

// AverageRmsOverVoltage is the average RMS voltage above which an over voltage condition is reported.
// The threshold shall be configurable within the specified operating range of the electricity meter.
// The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively.
type AverageRmsOverVoltage zcl.Zs16

func (v *AverageRmsOverVoltage) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AverageRmsOverVoltage) Value() zcl.Val     { return v }

func (v AverageRmsOverVoltage) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AverageRmsOverVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsOverVoltage(*nt)
	return br, err
}

func (v AverageRmsOverVoltage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AverageRmsOverVoltage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsOverVoltage(*a)
	return nil
}

func (v *AverageRmsOverVoltage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AverageRmsOverVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsOverVoltage) String() string {
	return zcl.Volts.Format(float64(v))
}

const AverageRmsUnderVoltageAttr zcl.AttrID = 2054

func (AverageRmsUnderVoltage) ID() zcl.AttrID   { return AverageRmsUnderVoltageAttr }
func (AverageRmsUnderVoltage) Readable() bool   { return false }
func (AverageRmsUnderVoltage) Writable() bool   { return false }
func (AverageRmsUnderVoltage) Reportable() bool { return false }
func (AverageRmsUnderVoltage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AverageRmsUnderVoltage) AttrID() zcl.AttrID   { return v.ID() }
func (v AverageRmsUnderVoltage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AverageRmsUnderVoltage) AttrValue() zcl.Val  { return v.Value() }

func (AverageRmsUnderVoltage) Name() string { return `Average RMS under-voltage` }
func (AverageRmsUnderVoltage) Description() string {
	return `is the average RMS voltage below which an under voltage condition is reported.
The threshold shall be configurable within the specified operating range of the electricity meter.
The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively.`
}

// AverageRmsUnderVoltage is the average RMS voltage below which an under voltage condition is reported.
// The threshold shall be configurable within the specified operating range of the electricity meter.
// The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively.
type AverageRmsUnderVoltage zcl.Zs16

func (v *AverageRmsUnderVoltage) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AverageRmsUnderVoltage) Value() zcl.Val     { return v }

func (v AverageRmsUnderVoltage) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AverageRmsUnderVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AverageRmsUnderVoltage(*nt)
	return br, err
}

func (v AverageRmsUnderVoltage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AverageRmsUnderVoltage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AverageRmsUnderVoltage(*a)
	return nil
}

func (v *AverageRmsUnderVoltage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AverageRmsUnderVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AverageRmsUnderVoltage) String() string {
	return zcl.Volts.Format(float64(v))
}

const DcCurrentAttr zcl.AttrID = 259

func (DcCurrent) ID() zcl.AttrID   { return DcCurrentAttr }
func (DcCurrent) Readable() bool   { return false }
func (DcCurrent) Writable() bool   { return false }
func (DcCurrent) Reportable() bool { return false }
func (DcCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v DcCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcCurrent) AttrValue() zcl.Val  { return v.Value() }

func (DcCurrent) Name() string { return `DC Current` }
func (DcCurrent) Description() string {
	return `represents the most recent DC current reading in Amps (A). If the current cannot be
measured, a value of 0x8000 is returned.`
}

// DcCurrent represents the most recent DC current reading in Amps (A). If the current cannot be
// measured, a value of 0x8000 is returned.
type DcCurrent zcl.Zs16

func (v *DcCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcCurrent) Value() zcl.Val     { return v }

func (v DcCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcCurrent(*nt)
	return br, err
}

func (v DcCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcCurrent(*a)
	return nil
}

func (v *DcCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const DcCurrentDivisorAttr zcl.AttrID = 515

func (DcCurrentDivisor) ID() zcl.AttrID   { return DcCurrentDivisorAttr }
func (DcCurrentDivisor) Readable() bool   { return false }
func (DcCurrentDivisor) Writable() bool   { return false }
func (DcCurrentDivisor) Reportable() bool { return false }
func (DcCurrentDivisor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcCurrentDivisor) AttrID() zcl.AttrID   { return v.ID() }
func (v DcCurrentDivisor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcCurrentDivisor) AttrValue() zcl.Val  { return v.Value() }

func (DcCurrentDivisor) Name() string { return `DC Current Divisor` }
func (DcCurrentDivisor) Description() string {
	return `provides a value to be divided against the DCCurrent, DCCurrentMin, and DCCurrentMax attributes.
This attribute must be used in conjunction with the DCCurrentMultiplier attribute.
0x0000 is an invalid value for this attribute.`
}

// DcCurrentDivisor provides a value to be divided against the DCCurrent, DCCurrentMin, and DCCurrentMax attributes.
// This attribute must be used in conjunction with the DCCurrentMultiplier attribute.
// 0x0000 is an invalid value for this attribute.
type DcCurrentDivisor zcl.Zu16

func (v *DcCurrentDivisor) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *DcCurrentDivisor) Value() zcl.Val     { return v }

func (v DcCurrentDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *DcCurrentDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcCurrentDivisor(*nt)
	return br, err
}

func (v DcCurrentDivisor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *DcCurrentDivisor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcCurrentDivisor(*a)
	return nil
}

func (v *DcCurrentDivisor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = DcCurrentDivisor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcCurrentDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const DcCurrentMaxAttr zcl.AttrID = 261

func (DcCurrentMax) ID() zcl.AttrID   { return DcCurrentMaxAttr }
func (DcCurrentMax) Readable() bool   { return false }
func (DcCurrentMax) Writable() bool   { return false }
func (DcCurrentMax) Reportable() bool { return false }
func (DcCurrentMax) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcCurrentMax) AttrID() zcl.AttrID   { return v.ID() }
func (v DcCurrentMax) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcCurrentMax) AttrValue() zcl.Val  { return v.Value() }

func (DcCurrentMax) Name() string { return `DC Current Max` }
func (DcCurrentMax) Description() string {
	return `represents the highest DC current value measured in Amps (A). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`
}

// DcCurrentMax represents the highest DC current value measured in Amps (A). After resetting,
// this attribute will return a value of 0x8000 until a measurement is made.
type DcCurrentMax zcl.Zs16

func (v *DcCurrentMax) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcCurrentMax) Value() zcl.Val     { return v }

func (v DcCurrentMax) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcCurrentMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcCurrentMax(*nt)
	return br, err
}

func (v DcCurrentMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcCurrentMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcCurrentMax(*a)
	return nil
}

func (v *DcCurrentMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcCurrentMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcCurrentMax) String() string {
	return zcl.Amperes.Format(float64(v))
}

const DcCurrentMinAttr zcl.AttrID = 260

func (DcCurrentMin) ID() zcl.AttrID   { return DcCurrentMinAttr }
func (DcCurrentMin) Readable() bool   { return false }
func (DcCurrentMin) Writable() bool   { return false }
func (DcCurrentMin) Reportable() bool { return false }
func (DcCurrentMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcCurrentMin) AttrID() zcl.AttrID   { return v.ID() }
func (v DcCurrentMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcCurrentMin) AttrValue() zcl.Val  { return v.Value() }

func (DcCurrentMin) Name() string { return `DC Current Min` }
func (DcCurrentMin) Description() string {
	return `represents the lowest DC current value measured in Amps (A). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`
}

// DcCurrentMin represents the lowest DC current value measured in Amps (A). After resetting,
// this attribute will return a value of 0x8000 until a measurement is made.
type DcCurrentMin zcl.Zs16

func (v *DcCurrentMin) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcCurrentMin) Value() zcl.Val     { return v }

func (v DcCurrentMin) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcCurrentMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcCurrentMin(*nt)
	return br, err
}

func (v DcCurrentMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcCurrentMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcCurrentMin(*a)
	return nil
}

func (v *DcCurrentMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcCurrentMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcCurrentMin) String() string {
	return zcl.Amperes.Format(float64(v))
}

const DcCurrentMultiplierAttr zcl.AttrID = 514

func (DcCurrentMultiplier) ID() zcl.AttrID   { return DcCurrentMultiplierAttr }
func (DcCurrentMultiplier) Readable() bool   { return false }
func (DcCurrentMultiplier) Writable() bool   { return false }
func (DcCurrentMultiplier) Reportable() bool { return false }
func (DcCurrentMultiplier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcCurrentMultiplier) AttrID() zcl.AttrID   { return v.ID() }
func (v DcCurrentMultiplier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcCurrentMultiplier) AttrValue() zcl.Val  { return v.Value() }

func (DcCurrentMultiplier) Name() string { return `DC Current Multiplier` }
func (DcCurrentMultiplier) Description() string {
	return `provides a value to be multiplied against the DCCurrent, DCCurrentMin, and DCCurrentMax attributes.
This attribute must be used in conjunction with the DCCurrentDivisor attribute.
0x0000 is an invalid value for this attribute.`
}

// DcCurrentMultiplier provides a value to be multiplied against the DCCurrent, DCCurrentMin, and DCCurrentMax attributes.
// This attribute must be used in conjunction with the DCCurrentDivisor attribute.
// 0x0000 is an invalid value for this attribute.
type DcCurrentMultiplier zcl.Zu16

func (v *DcCurrentMultiplier) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *DcCurrentMultiplier) Value() zcl.Val     { return v }

func (v DcCurrentMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *DcCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcCurrentMultiplier(*nt)
	return br, err
}

func (v DcCurrentMultiplier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *DcCurrentMultiplier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcCurrentMultiplier(*a)
	return nil
}

func (v *DcCurrentMultiplier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = DcCurrentMultiplier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcCurrentMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const DcCurrentOverloadAttr zcl.AttrID = 1794

func (DcCurrentOverload) ID() zcl.AttrID   { return DcCurrentOverloadAttr }
func (DcCurrentOverload) Readable() bool   { return false }
func (DcCurrentOverload) Writable() bool   { return false }
func (DcCurrentOverload) Reportable() bool { return false }
func (DcCurrentOverload) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcCurrentOverload) AttrID() zcl.AttrID   { return v.ID() }
func (v DcCurrentOverload) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcCurrentOverload) AttrValue() zcl.Val  { return v.Value() }

func (DcCurrentOverload) Name() string { return `DC Current Overload` }
func (DcCurrentOverload) Description() string {
	return `Specifies the alarm threshold, set by the manufacturer, for the maximum output current supported by device.
The value is multiplied and divided by the DCCurrentMultiplier and DCCurrentDivider respectively.`
}

// DcCurrentOverload Specifies the alarm threshold, set by the manufacturer, for the maximum output current supported by device.
// The value is multiplied and divided by the DCCurrentMultiplier and DCCurrentDivider respectively.
type DcCurrentOverload zcl.Zs16

func (v *DcCurrentOverload) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcCurrentOverload) Value() zcl.Val     { return v }

func (v DcCurrentOverload) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcCurrentOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcCurrentOverload(*nt)
	return br, err
}

func (v DcCurrentOverload) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcCurrentOverload) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcCurrentOverload(*a)
	return nil
}

func (v *DcCurrentOverload) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcCurrentOverload(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcCurrentOverload) String() string {
	return zcl.Amperes.Format(float64(v))
}

const DcOverloadAlarmsMaskAttr zcl.AttrID = 1792

func (DcOverloadAlarmsMask) ID() zcl.AttrID   { return DcOverloadAlarmsMaskAttr }
func (DcOverloadAlarmsMask) Readable() bool   { return true }
func (DcOverloadAlarmsMask) Writable() bool   { return true }
func (DcOverloadAlarmsMask) Reportable() bool { return false }
func (DcOverloadAlarmsMask) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcOverloadAlarmsMask) AttrID() zcl.AttrID   { return v.ID() }
func (v DcOverloadAlarmsMask) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcOverloadAlarmsMask) AttrValue() zcl.Val  { return v.Value() }

func (DcOverloadAlarmsMask) Name() string { return `DC Overload Alarms Mask` }
func (DcOverloadAlarmsMask) Description() string {
	return `Specifies which configurable alarms may be generated`
}

// DcOverloadAlarmsMask Specifies which configurable alarms may be generated
type DcOverloadAlarmsMask zcl.Zbmp8

func (v *DcOverloadAlarmsMask) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *DcOverloadAlarmsMask) Value() zcl.Val     { return v }

func (v DcOverloadAlarmsMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *DcOverloadAlarmsMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = DcOverloadAlarmsMask(*nt)
	return br, err
}

func (v DcOverloadAlarmsMask) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *DcOverloadAlarmsMask) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcOverloadAlarmsMask(*a)
	return nil
}

func (v *DcOverloadAlarmsMask) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = DcOverloadAlarmsMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcOverloadAlarmsMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Voltage Overload")
		case 1:
			bstr = append(bstr, "Current Overload")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v DcOverloadAlarmsMask) IsVoltageOverload() bool { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v DcOverloadAlarmsMask) IsCurrentOverload() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v *DcOverloadAlarmsMask) SetVoltageOverload(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *DcOverloadAlarmsMask) SetCurrentOverload(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}

func (DcOverloadAlarmsMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Voltage Overload"},
		{Value: 1, Name: "Current Overload"},
	}
}

const DcPowerAttr zcl.AttrID = 262

func (DcPower) ID() zcl.AttrID   { return DcPowerAttr }
func (DcPower) Readable() bool   { return false }
func (DcPower) Writable() bool   { return false }
func (DcPower) Reportable() bool { return false }
func (DcPower) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcPower) AttrID() zcl.AttrID   { return v.ID() }
func (v DcPower) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcPower) AttrValue() zcl.Val  { return v.Value() }

func (DcPower) Name() string { return `DC Power` }
func (DcPower) Description() string {
	return `represents the most recent DC power reading in Watts (W). If the power cannot be
measured, a value of 0x8000 is returned.`
}

// DcPower represents the most recent DC power reading in Watts (W). If the power cannot be
// measured, a value of 0x8000 is returned.
type DcPower zcl.Zs16

func (v *DcPower) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcPower) Value() zcl.Val     { return v }

func (v DcPower) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcPower(*nt)
	return br, err
}

func (v DcPower) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcPower) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcPower(*a)
	return nil
}

func (v *DcPower) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcPower(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcPower) String() string {
	return zcl.Watts.Format(float64(v))
}

const DcPowerDivisorAttr zcl.AttrID = 517

func (DcPowerDivisor) ID() zcl.AttrID   { return DcPowerDivisorAttr }
func (DcPowerDivisor) Readable() bool   { return false }
func (DcPowerDivisor) Writable() bool   { return false }
func (DcPowerDivisor) Reportable() bool { return false }
func (DcPowerDivisor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcPowerDivisor) AttrID() zcl.AttrID   { return v.ID() }
func (v DcPowerDivisor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcPowerDivisor) AttrValue() zcl.Val  { return v.Value() }

func (DcPowerDivisor) Name() string { return `DC Power Divisor` }
func (DcPowerDivisor) Description() string {
	return `provides a value to be divided against the DCPower, DCPowerMin, and DCPowerMax attributes.
This attribute must be used in conjunction with the DCPowerMultiplier attribute.
0x0000 is an invalid value for this attribute`
}

// DcPowerDivisor provides a value to be divided against the DCPower, DCPowerMin, and DCPowerMax attributes.
// This attribute must be used in conjunction with the DCPowerMultiplier attribute.
// 0x0000 is an invalid value for this attribute
type DcPowerDivisor zcl.Zu16

func (v *DcPowerDivisor) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *DcPowerDivisor) Value() zcl.Val     { return v }

func (v DcPowerDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *DcPowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcPowerDivisor(*nt)
	return br, err
}

func (v DcPowerDivisor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *DcPowerDivisor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcPowerDivisor(*a)
	return nil
}

func (v *DcPowerDivisor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = DcPowerDivisor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcPowerDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const DcPowerMaxAttr zcl.AttrID = 264

func (DcPowerMax) ID() zcl.AttrID   { return DcPowerMaxAttr }
func (DcPowerMax) Readable() bool   { return false }
func (DcPowerMax) Writable() bool   { return false }
func (DcPowerMax) Reportable() bool { return false }
func (DcPowerMax) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcPowerMax) AttrID() zcl.AttrID   { return v.ID() }
func (v DcPowerMax) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcPowerMax) AttrValue() zcl.Val  { return v.Value() }

func (DcPowerMax) Name() string { return `DC Power Max` }
func (DcPowerMax) Description() string {
	return `represents the highest DC power value measured in Watts (W). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`
}

// DcPowerMax represents the highest DC power value measured in Watts (W). After resetting,
// this attribute will return a value of 0x8000 until a measurement is made.
type DcPowerMax zcl.Zs16

func (v *DcPowerMax) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcPowerMax) Value() zcl.Val     { return v }

func (v DcPowerMax) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcPowerMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcPowerMax(*nt)
	return br, err
}

func (v DcPowerMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcPowerMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcPowerMax(*a)
	return nil
}

func (v *DcPowerMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcPowerMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcPowerMax) String() string {
	return zcl.Watts.Format(float64(v))
}

const DcPowerMinAttr zcl.AttrID = 263

func (DcPowerMin) ID() zcl.AttrID   { return DcPowerMinAttr }
func (DcPowerMin) Readable() bool   { return false }
func (DcPowerMin) Writable() bool   { return false }
func (DcPowerMin) Reportable() bool { return false }
func (DcPowerMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcPowerMin) AttrID() zcl.AttrID   { return v.ID() }
func (v DcPowerMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcPowerMin) AttrValue() zcl.Val  { return v.Value() }

func (DcPowerMin) Name() string { return `DC Power Min` }
func (DcPowerMin) Description() string {
	return `represents the lowest DC power value measured in Watts (W). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`
}

// DcPowerMin represents the lowest DC power value measured in Watts (W). After resetting,
// this attribute will return a value of 0x8000 until a measurement is made.
type DcPowerMin zcl.Zs16

func (v *DcPowerMin) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcPowerMin) Value() zcl.Val     { return v }

func (v DcPowerMin) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcPowerMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcPowerMin(*nt)
	return br, err
}

func (v DcPowerMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcPowerMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcPowerMin(*a)
	return nil
}

func (v *DcPowerMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcPowerMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcPowerMin) String() string {
	return zcl.Watts.Format(float64(v))
}

const DcPowerMultiplierAttr zcl.AttrID = 516

func (DcPowerMultiplier) ID() zcl.AttrID   { return DcPowerMultiplierAttr }
func (DcPowerMultiplier) Readable() bool   { return false }
func (DcPowerMultiplier) Writable() bool   { return false }
func (DcPowerMultiplier) Reportable() bool { return false }
func (DcPowerMultiplier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcPowerMultiplier) AttrID() zcl.AttrID   { return v.ID() }
func (v DcPowerMultiplier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcPowerMultiplier) AttrValue() zcl.Val  { return v.Value() }

func (DcPowerMultiplier) Name() string { return `DC Power Multiplier` }
func (DcPowerMultiplier) Description() string {
	return `provides a value to be multiplied against the DCPower, DCPowerMin, and DCPowerMax attributes.
This attribute must be used in conjunction with the DCPowerDivisor attribute.
0x0000 is an invalid value for this attribute.`
}

// DcPowerMultiplier provides a value to be multiplied against the DCPower, DCPowerMin, and DCPowerMax attributes.
// This attribute must be used in conjunction with the DCPowerDivisor attribute.
// 0x0000 is an invalid value for this attribute.
type DcPowerMultiplier zcl.Zu16

func (v *DcPowerMultiplier) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *DcPowerMultiplier) Value() zcl.Val     { return v }

func (v DcPowerMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *DcPowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcPowerMultiplier(*nt)
	return br, err
}

func (v DcPowerMultiplier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *DcPowerMultiplier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcPowerMultiplier(*a)
	return nil
}

func (v *DcPowerMultiplier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = DcPowerMultiplier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcPowerMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const DcVoltageAttr zcl.AttrID = 256

func (DcVoltage) ID() zcl.AttrID   { return DcVoltageAttr }
func (DcVoltage) Readable() bool   { return false }
func (DcVoltage) Writable() bool   { return false }
func (DcVoltage) Reportable() bool { return false }
func (DcVoltage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcVoltage) AttrID() zcl.AttrID   { return v.ID() }
func (v DcVoltage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcVoltage) AttrValue() zcl.Val  { return v.Value() }

func (DcVoltage) Name() string { return `DC Voltage` }
func (DcVoltage) Description() string {
	return `represents the most recent DC voltage reading in Volts (V). If the voltage cannot be
measured, a value of 0x8000 is returned.`
}

// DcVoltage represents the most recent DC voltage reading in Volts (V). If the voltage cannot be
// measured, a value of 0x8000 is returned.
type DcVoltage zcl.Zs16

func (v *DcVoltage) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcVoltage) Value() zcl.Val     { return v }

func (v DcVoltage) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcVoltage(*nt)
	return br, err
}

func (v DcVoltage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcVoltage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcVoltage(*a)
	return nil
}

func (v *DcVoltage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcVoltage) String() string {
	return zcl.Volts.Format(float64(v))
}

const DcVoltageDivisorAttr zcl.AttrID = 513

func (DcVoltageDivisor) ID() zcl.AttrID   { return DcVoltageDivisorAttr }
func (DcVoltageDivisor) Readable() bool   { return false }
func (DcVoltageDivisor) Writable() bool   { return false }
func (DcVoltageDivisor) Reportable() bool { return false }
func (DcVoltageDivisor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcVoltageDivisor) AttrID() zcl.AttrID   { return v.ID() }
func (v DcVoltageDivisor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcVoltageDivisor) AttrValue() zcl.Val  { return v.Value() }

func (DcVoltageDivisor) Name() string { return `DC Voltage Divisor` }
func (DcVoltageDivisor) Description() string {
	return `provides a value to be divided against the DCVoltage, DCVoltageMin, and DCVoltageMax attributes.
This attribute must be used in conjunction with the DCVoltageMultiplier attribute.
0x0000 is an invalid value for this attribute.`
}

// DcVoltageDivisor provides a value to be divided against the DCVoltage, DCVoltageMin, and DCVoltageMax attributes.
// This attribute must be used in conjunction with the DCVoltageMultiplier attribute.
// 0x0000 is an invalid value for this attribute.
type DcVoltageDivisor zcl.Zu16

func (v *DcVoltageDivisor) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *DcVoltageDivisor) Value() zcl.Val     { return v }

func (v DcVoltageDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *DcVoltageDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcVoltageDivisor(*nt)
	return br, err
}

func (v DcVoltageDivisor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *DcVoltageDivisor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcVoltageDivisor(*a)
	return nil
}

func (v *DcVoltageDivisor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = DcVoltageDivisor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcVoltageDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const DcVoltageMaxAttr zcl.AttrID = 258

func (DcVoltageMax) ID() zcl.AttrID   { return DcVoltageMaxAttr }
func (DcVoltageMax) Readable() bool   { return false }
func (DcVoltageMax) Writable() bool   { return false }
func (DcVoltageMax) Reportable() bool { return false }
func (DcVoltageMax) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcVoltageMax) AttrID() zcl.AttrID   { return v.ID() }
func (v DcVoltageMax) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcVoltageMax) AttrValue() zcl.Val  { return v.Value() }

func (DcVoltageMax) Name() string { return `DC Voltage Max` }
func (DcVoltageMax) Description() string {
	return `represents the highest DC voltage value measured in Volts (V). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`
}

// DcVoltageMax represents the highest DC voltage value measured in Volts (V). After resetting,
// this attribute will return a value of 0x8000 until a measurement is made.
type DcVoltageMax zcl.Zs16

func (v *DcVoltageMax) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcVoltageMax) Value() zcl.Val     { return v }

func (v DcVoltageMax) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcVoltageMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcVoltageMax(*nt)
	return br, err
}

func (v DcVoltageMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcVoltageMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcVoltageMax(*a)
	return nil
}

func (v *DcVoltageMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcVoltageMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcVoltageMax) String() string {
	return zcl.Volts.Format(float64(v))
}

const DcVoltageMinAttr zcl.AttrID = 257

func (DcVoltageMin) ID() zcl.AttrID   { return DcVoltageMinAttr }
func (DcVoltageMin) Readable() bool   { return false }
func (DcVoltageMin) Writable() bool   { return false }
func (DcVoltageMin) Reportable() bool { return false }
func (DcVoltageMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcVoltageMin) AttrID() zcl.AttrID   { return v.ID() }
func (v DcVoltageMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcVoltageMin) AttrValue() zcl.Val  { return v.Value() }

func (DcVoltageMin) Name() string { return `DC Voltage Min` }
func (DcVoltageMin) Description() string {
	return `represents the lowest DC voltage value measured in Volts (V). After resetting,
this attribute will return a value of 0x8000 until a measurement is made.`
}

// DcVoltageMin represents the lowest DC voltage value measured in Volts (V). After resetting,
// this attribute will return a value of 0x8000 until a measurement is made.
type DcVoltageMin zcl.Zs16

func (v *DcVoltageMin) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcVoltageMin) Value() zcl.Val     { return v }

func (v DcVoltageMin) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcVoltageMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcVoltageMin(*nt)
	return br, err
}

func (v DcVoltageMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcVoltageMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcVoltageMin(*a)
	return nil
}

func (v *DcVoltageMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcVoltageMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcVoltageMin) String() string {
	return zcl.Volts.Format(float64(v))
}

const DcVoltageMultiplierAttr zcl.AttrID = 512

func (DcVoltageMultiplier) ID() zcl.AttrID   { return DcVoltageMultiplierAttr }
func (DcVoltageMultiplier) Readable() bool   { return false }
func (DcVoltageMultiplier) Writable() bool   { return false }
func (DcVoltageMultiplier) Reportable() bool { return false }
func (DcVoltageMultiplier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcVoltageMultiplier) AttrID() zcl.AttrID   { return v.ID() }
func (v DcVoltageMultiplier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcVoltageMultiplier) AttrValue() zcl.Val  { return v.Value() }

func (DcVoltageMultiplier) Name() string { return `DC Voltage Multiplier` }
func (DcVoltageMultiplier) Description() string {
	return `provides a value to be multiplied against the DCVoltage, DCVoltageMin, and DCVoltageMax attributes.
This attribute must be used in conjunction with the DCVoltageDivisor attribute.
0x0000 is an invalid value for this attribute`
}

// DcVoltageMultiplier provides a value to be multiplied against the DCVoltage, DCVoltageMin, and DCVoltageMax attributes.
// This attribute must be used in conjunction with the DCVoltageDivisor attribute.
// 0x0000 is an invalid value for this attribute
type DcVoltageMultiplier zcl.Zu16

func (v *DcVoltageMultiplier) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *DcVoltageMultiplier) Value() zcl.Val     { return v }

func (v DcVoltageMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *DcVoltageMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcVoltageMultiplier(*nt)
	return br, err
}

func (v DcVoltageMultiplier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *DcVoltageMultiplier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcVoltageMultiplier(*a)
	return nil
}

func (v *DcVoltageMultiplier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = DcVoltageMultiplier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcVoltageMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const DcVoltageOverloadAttr zcl.AttrID = 1793

func (DcVoltageOverload) ID() zcl.AttrID   { return DcVoltageOverloadAttr }
func (DcVoltageOverload) Readable() bool   { return false }
func (DcVoltageOverload) Writable() bool   { return false }
func (DcVoltageOverload) Reportable() bool { return false }
func (DcVoltageOverload) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DcVoltageOverload) AttrID() zcl.AttrID   { return v.ID() }
func (v DcVoltageOverload) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DcVoltageOverload) AttrValue() zcl.Val  { return v.Value() }

func (DcVoltageOverload) Name() string { return `DC Voltage Overload` }
func (DcVoltageOverload) Description() string {
	return `Specifies the alarm threshold, set by the manufacturer, for the maximum output voltage supported by device.
The value is multiplied and divided by the DCVoltageMultiplier the DCVoltageDivisor respectively.`
}

// DcVoltageOverload Specifies the alarm threshold, set by the manufacturer, for the maximum output voltage supported by device.
// The value is multiplied and divided by the DCVoltageMultiplier the DCVoltageDivisor respectively.
type DcVoltageOverload zcl.Zs16

func (v *DcVoltageOverload) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *DcVoltageOverload) Value() zcl.Val     { return v }

func (v DcVoltageOverload) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *DcVoltageOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = DcVoltageOverload(*nt)
	return br, err
}

func (v DcVoltageOverload) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *DcVoltageOverload) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DcVoltageOverload(*a)
	return nil
}

func (v *DcVoltageOverload) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = DcVoltageOverload(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DcVoltageOverload) String() string {
	return zcl.Volts.Format(float64(v))
}

const ElectricalMeasurementTypeAttr zcl.AttrID = 0

func (ElectricalMeasurementType) ID() zcl.AttrID   { return ElectricalMeasurementTypeAttr }
func (ElectricalMeasurementType) Readable() bool   { return false }
func (ElectricalMeasurementType) Writable() bool   { return false }
func (ElectricalMeasurementType) Reportable() bool { return false }
func (ElectricalMeasurementType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ElectricalMeasurementType) AttrID() zcl.AttrID   { return v.ID() }
func (v ElectricalMeasurementType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ElectricalMeasurementType) AttrValue() zcl.Val  { return v.Value() }

func (ElectricalMeasurementType) Name() string        { return `Electrical Measurement Type` }
func (ElectricalMeasurementType) Description() string { return `` }

type ElectricalMeasurementType zcl.Zbmp32

func (v *ElectricalMeasurementType) TypeID() zcl.TypeID { return new(zcl.Zbmp32).TypeID() }
func (v *ElectricalMeasurementType) Value() zcl.Val     { return v }

func (v ElectricalMeasurementType) MarshalZcl() ([]byte, error) { return zcl.Zbmp32(v).MarshalZcl() }

func (v *ElectricalMeasurementType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*v = ElectricalMeasurementType(*nt)
	return br, err
}

func (v ElectricalMeasurementType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp32(v))
}

func (v *ElectricalMeasurementType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ElectricalMeasurementType(*a)
	return nil
}

func (v *ElectricalMeasurementType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp32); ok {
		*v = ElectricalMeasurementType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ElectricalMeasurementType) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Active measurement (AC)")
		case 1:
			bstr = append(bstr, "Reactive measurement (AC)")
		case 2:
			bstr = append(bstr, "Apparent measurement (AC)")
		case 3:
			bstr = append(bstr, "Phase A measurement")
		case 4:
			bstr = append(bstr, "Phase B measurement")
		case 5:
			bstr = append(bstr, "Phase C measurement")
		case 6:
			bstr = append(bstr, "DC measurement")
		case 7:
			bstr = append(bstr, "Harmonics measurement")
		case 8:
			bstr = append(bstr, "Power quality measurement")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v ElectricalMeasurementType) IsActiveMeasurementAc() bool {
	return zcl.BitmapTest([]byte(v[:]), 0)
}
func (v ElectricalMeasurementType) IsReactiveMeasurementAc() bool {
	return zcl.BitmapTest([]byte(v[:]), 1)
}
func (v ElectricalMeasurementType) IsApparentMeasurementAc() bool {
	return zcl.BitmapTest([]byte(v[:]), 2)
}
func (v ElectricalMeasurementType) IsPhaseAMeasurement() bool { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v ElectricalMeasurementType) IsPhaseBMeasurement() bool { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v ElectricalMeasurementType) IsPhaseCMeasurement() bool { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v ElectricalMeasurementType) IsDcMeasurement() bool     { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v ElectricalMeasurementType) IsHarmonicsMeasurement() bool {
	return zcl.BitmapTest([]byte(v[:]), 7)
}
func (v ElectricalMeasurementType) IsPowerQualityMeasurement() bool {
	return zcl.BitmapTest([]byte(v[:]), 8)
}
func (v *ElectricalMeasurementType) SetActiveMeasurementAc(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *ElectricalMeasurementType) SetReactiveMeasurementAc(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *ElectricalMeasurementType) SetApparentMeasurementAc(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *ElectricalMeasurementType) SetPhaseAMeasurement(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *ElectricalMeasurementType) SetPhaseBMeasurement(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
}
func (v *ElectricalMeasurementType) SetPhaseCMeasurement(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b))
}
func (v *ElectricalMeasurementType) SetDcMeasurement(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b))
}
func (v *ElectricalMeasurementType) SetHarmonicsMeasurement(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 7, b))
}
func (v *ElectricalMeasurementType) SetPowerQualityMeasurement(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 8, b))
}

func (ElectricalMeasurementType) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Active measurement (AC)"},
		{Value: 1, Name: "Reactive measurement (AC)"},
		{Value: 2, Name: "Apparent measurement (AC)"},
		{Value: 3, Name: "Phase A measurement"},
		{Value: 4, Name: "Phase B measurement"},
		{Value: 5, Name: "Phase C measurement"},
		{Value: 6, Name: "DC measurement"},
		{Value: 7, Name: "Harmonics measurement"},
		{Value: 8, Name: "Power quality measurement"},
	}
}

const HarmonicCurrentMultiplierAttr zcl.AttrID = 1028

func (HarmonicCurrentMultiplier) ID() zcl.AttrID   { return HarmonicCurrentMultiplierAttr }
func (HarmonicCurrentMultiplier) Readable() bool   { return false }
func (HarmonicCurrentMultiplier) Writable() bool   { return false }
func (HarmonicCurrentMultiplier) Reportable() bool { return false }
func (HarmonicCurrentMultiplier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v HarmonicCurrentMultiplier) AttrID() zcl.AttrID   { return v.ID() }
func (v HarmonicCurrentMultiplier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *HarmonicCurrentMultiplier) AttrValue() zcl.Val  { return v.Value() }

func (HarmonicCurrentMultiplier) Name() string { return `Harmonic Current Multiplier` }
func (HarmonicCurrentMultiplier) Description() string {
	return `Represents the unit value for the MeasuredNthHarmonicCurrent attribute in the format
MeasuredNthHarmonicCurrent * 10 ^ HarmonicCurrentMultiplier amperes.`
}

// HarmonicCurrentMultiplier Represents the unit value for the MeasuredNthHarmonicCurrent attribute in the format
// MeasuredNthHarmonicCurrent * 10 ^ HarmonicCurrentMultiplier amperes.
type HarmonicCurrentMultiplier zcl.Zs8

func (v *HarmonicCurrentMultiplier) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *HarmonicCurrentMultiplier) Value() zcl.Val     { return v }

func (v HarmonicCurrentMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *HarmonicCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = HarmonicCurrentMultiplier(*nt)
	return br, err
}

func (v HarmonicCurrentMultiplier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *HarmonicCurrentMultiplier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = HarmonicCurrentMultiplier(*a)
	return nil
}

func (v *HarmonicCurrentMultiplier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = HarmonicCurrentMultiplier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v HarmonicCurrentMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const IlluminanceLightSensorTypeAttr zcl.AttrID = 1

func (IlluminanceLightSensorType) ID() zcl.AttrID   { return IlluminanceLightSensorTypeAttr }
func (IlluminanceLightSensorType) Readable() bool   { return true }
func (IlluminanceLightSensorType) Writable() bool   { return false }
func (IlluminanceLightSensorType) Reportable() bool { return false }
func (IlluminanceLightSensorType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IlluminanceLightSensorType) AttrID() zcl.AttrID   { return v.ID() }
func (v IlluminanceLightSensorType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IlluminanceLightSensorType) AttrValue() zcl.Val  { return v.Value() }

func (IlluminanceLightSensorType) Name() string { return `Illuminance Light Sensor Type` }
func (IlluminanceLightSensorType) Description() string {
	return `specifies the electronic type of the light sensor. The range 0x40-0xfe is reserved for
manufacturer specific light sensor types`
}

// IlluminanceLightSensorType specifies the electronic type of the light sensor. The range 0x40-0xfe is reserved for
// manufacturer specific light sensor types
type IlluminanceLightSensorType zcl.Zenum8

func (v *IlluminanceLightSensorType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *IlluminanceLightSensorType) Value() zcl.Val     { return v }

func (v IlluminanceLightSensorType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *IlluminanceLightSensorType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = IlluminanceLightSensorType(*nt)
	return br, err
}

func (v IlluminanceLightSensorType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *IlluminanceLightSensorType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IlluminanceLightSensorType(*a)
	return nil
}

func (v *IlluminanceLightSensorType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = IlluminanceLightSensorType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IlluminanceLightSensorType) String() string {
	switch v {
	case 0x00:
		return "Photodiode"
	case 0x01:
		return "CMOS"
	case 0xff:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v IlluminanceLightSensorType) IsPhotodiode() bool { return v == 0x00 }
func (v IlluminanceLightSensorType) IsCmos() bool       { return v == 0x01 }
func (v IlluminanceLightSensorType) IsUnknown() bool    { return v == 0xff }
func (v *IlluminanceLightSensorType) SetPhotodiode()    { *v = 0x00 }
func (v *IlluminanceLightSensorType) SetCmos()          { *v = 0x01 }
func (v *IlluminanceLightSensorType) SetUnknown()       { *v = 0xff }

func (IlluminanceLightSensorType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Photodiode"},
		{Value: 0x01, Name: "CMOS"},
		{Value: 0xff, Name: "Unknown"},
	}
}

const IlluminanceTargetLevelAttr zcl.AttrID = 16

func (IlluminanceTargetLevel) ID() zcl.AttrID   { return IlluminanceTargetLevelAttr }
func (IlluminanceTargetLevel) Readable() bool   { return true }
func (IlluminanceTargetLevel) Writable() bool   { return true }
func (IlluminanceTargetLevel) Reportable() bool { return false }
func (IlluminanceTargetLevel) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IlluminanceTargetLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v IlluminanceTargetLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IlluminanceTargetLevel) AttrValue() zcl.Val  { return v.Value() }

func (IlluminanceTargetLevel) Name() string { return `Illuminance Target Level` }
func (IlluminanceTargetLevel) Description() string {
	return `specifies the target Illuminance level. This target level is taken as the
centre of a 'dead band', which must be sufficient in width, with hysteresis
bands at both top and bottom, to provide reliable notifications without
'chatter'. IlluminanceTargetLevel represents Illuminance in Lux (symbol lx) as
follows: IlluminanceTargetLevel = 10,000 x log 10 Illuminance`
}

// IlluminanceTargetLevel specifies the target Illuminance level. This target level is taken as the
// centre of a 'dead band', which must be sufficient in width, with hysteresis
// bands at both top and bottom, to provide reliable notifications without
// 'chatter'. IlluminanceTargetLevel represents Illuminance in Lux (symbol lx) as
// follows: IlluminanceTargetLevel = 10,000 x log 10 Illuminance
type IlluminanceTargetLevel zcl.Zu16

func (v *IlluminanceTargetLevel) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *IlluminanceTargetLevel) Value() zcl.Val     { return v }

func (v IlluminanceTargetLevel) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *IlluminanceTargetLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = IlluminanceTargetLevel(*nt)
	return br, err
}

func (v IlluminanceTargetLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *IlluminanceTargetLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IlluminanceTargetLevel(*a)
	return nil
}

func (v *IlluminanceTargetLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = IlluminanceTargetLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IlluminanceTargetLevel) String() string {
	return zcl.Luxes.Format(float64(v))
}

func (Intervals) Name() string { return `Intervals` }
func (Intervals) Description() string {
	return `is a series of interval data captured using the period specified by the ProfileIntervalPeriod field. The
content of the interval data depend of the type of information requested using the AttributeID field in the Get
Measurement Profile Command. Data is organized in a reverse chronological order, the oldest intervals are
transmitted first and the newest interval is transmitted last. Invalid intervals should be marked as 0xFFFF.
For scaling and data type use the respective attribute set as defined above in attribute sets.`
}

// Intervals is a series of interval data captured using the period specified by the ProfileIntervalPeriod field. The
// content of the interval data depend of the type of information requested using the AttributeID field in the Get
// Measurement Profile Command. Data is organized in a reverse chronological order, the oldest intervals are
// transmitted first and the newest interval is transmitted last. Invalid intervals should be marked as 0xFFFF.
// For scaling and data type use the respective attribute set as defined above in attribute sets.
type Intervals []*zcl.Zu16

func (v *Intervals) TypeID() zcl.TypeID { return new(zcl.Zarray).TypeID() }
func (v *Intervals) Value() zcl.Val     { return v }

func (Intervals) ArrayTypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }

func (v *Intervals) ArrayValues() (o []uint16) {
	for _, a := range *v {
		o = append(o, uint16(*a))
	}
	return o
}

func (v *Intervals) SetValues(val []uint16) error {
	*v = []*zcl.Zu16{}
	return v.AddValues(val...)
}

func (v *Intervals) AddValues(val ...uint16) error {
	for _, a := range val {
		nv := zcl.Zu16(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v Intervals) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *Intervals) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu16{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu16)
		*v = append(*v, nv)
		return nv
	})
}

func (v *Intervals) SetValue(a zcl.Val) error {
	if nv, ok := a.(*Intervals); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Intervals) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

const LevelStatusAttr zcl.AttrID = 0

func (LevelStatus) ID() zcl.AttrID   { return LevelStatusAttr }
func (LevelStatus) Readable() bool   { return true }
func (LevelStatus) Writable() bool   { return false }
func (LevelStatus) Reportable() bool { return true }
func (LevelStatus) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LevelStatus) AttrID() zcl.AttrID   { return v.ID() }
func (v LevelStatus) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LevelStatus) AttrValue() zcl.Val  { return v.Value() }

func (LevelStatus) Name() string { return `Level Status` }
func (LevelStatus) Description() string {
	return `indicates whether the measured Flow is above, below, or within a band
around FlowTargetLevel`
}

// LevelStatus indicates whether the measured Flow is above, below, or within a band
// around FlowTargetLevel
type LevelStatus zcl.Zenum8

func (v *LevelStatus) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *LevelStatus) Value() zcl.Val     { return v }

func (v LevelStatus) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *LevelStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = LevelStatus(*nt)
	return br, err
}

func (v LevelStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *LevelStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LevelStatus(*a)
	return nil
}

func (v *LevelStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = LevelStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LevelStatus) String() string {
	switch v {
	case 0x00:
		return "Flow on target"
	case 0x01:
		return "Flow below target"
	case 0x03:
		return "Flow above target"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v LevelStatus) IsFlowOnTarget() bool    { return v == 0x00 }
func (v LevelStatus) IsFlowBelowTarget() bool { return v == 0x01 }
func (v LevelStatus) IsFlowAboveTarget() bool { return v == 0x03 }
func (v *LevelStatus) SetFlowOnTarget()       { *v = 0x00 }
func (v *LevelStatus) SetFlowBelowTarget()    { *v = 0x01 }
func (v *LevelStatus) SetFlowAboveTarget()    { *v = 0x03 }

func (LevelStatus) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Flow on target"},
		{Value: 0x01, Name: "Flow below target"},
		{Value: 0x03, Name: "Flow above target"},
	}
}

const LightSensorTypeAttr zcl.AttrID = 4

func (LightSensorType) ID() zcl.AttrID   { return LightSensorTypeAttr }
func (LightSensorType) Readable() bool   { return true }
func (LightSensorType) Writable() bool   { return false }
func (LightSensorType) Reportable() bool { return false }
func (LightSensorType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LightSensorType) AttrID() zcl.AttrID   { return v.ID() }
func (v LightSensorType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LightSensorType) AttrValue() zcl.Val  { return v.Value() }

func (LightSensorType) Name() string { return `Light Sensor Type` }
func (LightSensorType) Description() string {
	return `specifies the electronic type of the light sensor. The range 0x40-0xfe is reserved for
manufacturer specific light sensor types`
}

// LightSensorType specifies the electronic type of the light sensor. The range 0x40-0xfe is reserved for
// manufacturer specific light sensor types
type LightSensorType zcl.Zenum8

func (v *LightSensorType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *LightSensorType) Value() zcl.Val     { return v }

func (v LightSensorType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *LightSensorType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = LightSensorType(*nt)
	return br, err
}

func (v LightSensorType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *LightSensorType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LightSensorType(*a)
	return nil
}

func (v *LightSensorType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = LightSensorType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LightSensorType) String() string {
	switch v {
	case 0x00:
		return "Photodiode"
	case 0x01:
		return "CMOS"
	case 0xff:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v LightSensorType) IsPhotodiode() bool { return v == 0x00 }
func (v LightSensorType) IsCmos() bool       { return v == 0x01 }
func (v LightSensorType) IsUnknown() bool    { return v == 0xff }
func (v *LightSensorType) SetPhotodiode()    { *v = 0x00 }
func (v *LightSensorType) SetCmos()          { *v = 0x01 }
func (v *LightSensorType) SetUnknown()       { *v = 0xff }

func (LightSensorType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Photodiode"},
		{Value: 0x01, Name: "CMOS"},
		{Value: 0xff, Name: "Unknown"},
	}
}

const LineCurrentAttr zcl.AttrID = 1281

func (LineCurrent) ID() zcl.AttrID   { return LineCurrentAttr }
func (LineCurrent) Readable() bool   { return false }
func (LineCurrent) Writable() bool   { return false }
func (LineCurrent) Reportable() bool { return false }
func (LineCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LineCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v LineCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LineCurrent) AttrValue() zcl.Val  { return v.Value() }

func (LineCurrent) Name() string { return `Line Current` }
func (LineCurrent) Description() string {
	return `Represents the single phase or Phase A, AC line current (Square root of active and reactive current) value at
the moment in time the attribute is read. If it cannot be measured, a value of 0x8000 is returned.`
}

// LineCurrent Represents the single phase or Phase A, AC line current (Square root of active and reactive current) value at
// the moment in time the attribute is read. If it cannot be measured, a value of 0x8000 is returned.
type LineCurrent zcl.Zu16

func (v *LineCurrent) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LineCurrent) Value() zcl.Val     { return v }

func (v LineCurrent) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LineCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LineCurrent(*nt)
	return br, err
}

func (v LineCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LineCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LineCurrent(*a)
	return nil
}

func (v *LineCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LineCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LineCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const LineCurrentPhBAttr zcl.AttrID = 2305

func (LineCurrentPhB) ID() zcl.AttrID   { return LineCurrentPhBAttr }
func (LineCurrentPhB) Readable() bool   { return false }
func (LineCurrentPhB) Writable() bool   { return false }
func (LineCurrentPhB) Reportable() bool { return false }
func (LineCurrentPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LineCurrentPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v LineCurrentPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LineCurrentPhB) AttrValue() zcl.Val  { return v.Value() }

func (LineCurrentPhB) Name() string { return `Line Current Ph B` }
func (LineCurrentPhB) Description() string {
	return `represents the Phase B, AC line current (Square root sum of active and reactive currents) value at the moment
in time the attribute is read.
If the instantaneous current cannot be measured, a value of 0x8000 is returned.`
}

// LineCurrentPhB represents the Phase B, AC line current (Square root sum of active and reactive currents) value at the moment
// in time the attribute is read.
// If the instantaneous current cannot be measured, a value of 0x8000 is returned.
type LineCurrentPhB zcl.Zu16

func (v *LineCurrentPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LineCurrentPhB) Value() zcl.Val     { return v }

func (v LineCurrentPhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LineCurrentPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LineCurrentPhB(*nt)
	return br, err
}

func (v LineCurrentPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LineCurrentPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LineCurrentPhB(*a)
	return nil
}

func (v *LineCurrentPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LineCurrentPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LineCurrentPhB) String() string {
	return zcl.Amperes.Format(float64(v))
}

const LineCurrentPhCAttr zcl.AttrID = 2561

func (LineCurrentPhC) ID() zcl.AttrID   { return LineCurrentPhCAttr }
func (LineCurrentPhC) Readable() bool   { return false }
func (LineCurrentPhC) Writable() bool   { return false }
func (LineCurrentPhC) Reportable() bool { return false }
func (LineCurrentPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LineCurrentPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v LineCurrentPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LineCurrentPhC) AttrValue() zcl.Val  { return v.Value() }

func (LineCurrentPhC) Name() string        { return `Line Current Ph C` }
func (LineCurrentPhC) Description() string { return `` }

type LineCurrentPhC zcl.Zu16

func (v *LineCurrentPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *LineCurrentPhC) Value() zcl.Val     { return v }

func (v LineCurrentPhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *LineCurrentPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = LineCurrentPhC(*nt)
	return br, err
}

func (v LineCurrentPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *LineCurrentPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LineCurrentPhC(*a)
	return nil
}

func (v *LineCurrentPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = LineCurrentPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LineCurrentPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (ListOfAttributes) Name() string        { return `List Of Attributes` }
func (ListOfAttributes) Description() string { return `is the list of attributes being profiled` }

// ListOfAttributes is the list of attributes being profiled
type ListOfAttributes []*zcl.Zu16

func (v *ListOfAttributes) TypeID() zcl.TypeID { return new(zcl.Zset).TypeID() }
func (v *ListOfAttributes) Value() zcl.Val     { return v }

func (ListOfAttributes) ArrayTypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }

func (v *ListOfAttributes) ArrayValues() (o []uint16) {
	for _, a := range *v {
		o = append(o, uint16(*a))
	}
	return o
}

func (v *ListOfAttributes) SetValues(val []uint16) error {
	*v = []*zcl.Zu16{}
	return v.AddValues(val...)
}

func (v *ListOfAttributes) AddValues(val ...uint16) error {
	for _, a := range val {
		nv := zcl.Zu16(a)
		*v = append(*v, &nv)
	}
	return nil
}

func (v ListOfAttributes) MarshalZcl() ([]byte, error) {
	var vars []zcl.Val
	for _, a := range v {
		vars = append(vars, a)
	}
	return zcl.ArrayNoTypeMarshalZcl("sloc", vars)
}

func (v *ListOfAttributes) UnmarshalZcl(b []byte) ([]byte, error) {
	*v = []*zcl.Zu16{}
	return zcl.UnmarshalList("sloc", b, func() zcl.Val {
		nv := new(zcl.Zu16)
		*v = append(*v, nv)
		return nv
	})
}

func (v *ListOfAttributes) SetValue(a zcl.Val) error {
	if nv, ok := a.(*ListOfAttributes); ok {
		*v = *nv
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ListOfAttributes) String() string {
	var s []string
	for _, a := range v {
		s = append(s, zcl.Sprintf("%v", a))
	}
	return "[" + zcl.StrJoin(s, ",") + "]"
}

const MaxMeasuredFlowAttr zcl.AttrID = 2

func (MaxMeasuredFlow) ID() zcl.AttrID   { return MaxMeasuredFlowAttr }
func (MaxMeasuredFlow) Readable() bool   { return true }
func (MaxMeasuredFlow) Writable() bool   { return false }
func (MaxMeasuredFlow) Reportable() bool { return false }
func (MaxMeasuredFlow) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxMeasuredFlow) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxMeasuredFlow) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxMeasuredFlow) AttrValue() zcl.Val  { return v.Value() }

func (MaxMeasuredFlow) Name() string { return `Max Measured Flow` }
func (MaxMeasuredFlow) Description() string {
	return `indicates the maximum value of MeasuredFlow that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MaxMeasuredFlow indicates the maximum value of MeasuredFlow that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MaxMeasuredFlow zcl.Zu16

func (v *MaxMeasuredFlow) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxMeasuredFlow) Value() zcl.Val     { return v }

func (v MaxMeasuredFlow) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxMeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxMeasuredFlow(*nt)
	return br, err
}

func (v MaxMeasuredFlow) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxMeasuredFlow) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxMeasuredFlow(*a)
	return nil
}

func (v *MaxMeasuredFlow) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxMeasuredFlow(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxMeasuredFlow) String() string {
	return zcl.CubicMetersPerHour.Format(float64(v) / 10)
}

func (v MaxMeasuredFlow) Scaled() float64 {
	return float64(v) / 10
}

const MaxMeasuredIlluminanceAttr zcl.AttrID = 2

func (MaxMeasuredIlluminance) ID() zcl.AttrID   { return MaxMeasuredIlluminanceAttr }
func (MaxMeasuredIlluminance) Readable() bool   { return true }
func (MaxMeasuredIlluminance) Writable() bool   { return false }
func (MaxMeasuredIlluminance) Reportable() bool { return false }
func (MaxMeasuredIlluminance) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxMeasuredIlluminance) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxMeasuredIlluminance) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxMeasuredIlluminance) AttrValue() zcl.Val  { return v.Value() }

func (MaxMeasuredIlluminance) Name() string { return `Max Measured Illuminance` }
func (MaxMeasuredIlluminance) Description() string {
	return `indicates the maximum value of MeasuredIlluminance that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MaxMeasuredIlluminance indicates the maximum value of MeasuredIlluminance that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MaxMeasuredIlluminance zcl.Zu16

func (v *MaxMeasuredIlluminance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxMeasuredIlluminance) Value() zcl.Val     { return v }

func (v MaxMeasuredIlluminance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxMeasuredIlluminance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxMeasuredIlluminance(*nt)
	return br, err
}

func (v MaxMeasuredIlluminance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxMeasuredIlluminance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxMeasuredIlluminance(*a)
	return nil
}

func (v *MaxMeasuredIlluminance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxMeasuredIlluminance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxMeasuredIlluminance) String() string {
	return zcl.Luxes.Format(float64(v))
}

const MaxMeasuredPressureAttr zcl.AttrID = 2

func (MaxMeasuredPressure) ID() zcl.AttrID   { return MaxMeasuredPressureAttr }
func (MaxMeasuredPressure) Readable() bool   { return true }
func (MaxMeasuredPressure) Writable() bool   { return false }
func (MaxMeasuredPressure) Reportable() bool { return false }
func (MaxMeasuredPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxMeasuredPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxMeasuredPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxMeasuredPressure) AttrValue() zcl.Val  { return v.Value() }

func (MaxMeasuredPressure) Name() string { return `Max Measured Pressure` }
func (MaxMeasuredPressure) Description() string {
	return `indicates the maximum value of MeasuredPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// MaxMeasuredPressure indicates the maximum value of MeasuredPressure that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type MaxMeasuredPressure zcl.Zs16

func (v *MaxMeasuredPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxMeasuredPressure) Value() zcl.Val     { return v }

func (v MaxMeasuredPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxMeasuredPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxMeasuredPressure(*nt)
	return br, err
}

func (v MaxMeasuredPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxMeasuredPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxMeasuredPressure(*a)
	return nil
}

func (v *MaxMeasuredPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxMeasuredPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxMeasuredPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MaxMeasuredPressure) Scaled() float64 {
	return float64(v) / 10
}

const MaxMeasuredRelativeHumidityAttr zcl.AttrID = 2

func (MaxMeasuredRelativeHumidity) ID() zcl.AttrID   { return MaxMeasuredRelativeHumidityAttr }
func (MaxMeasuredRelativeHumidity) Readable() bool   { return true }
func (MaxMeasuredRelativeHumidity) Writable() bool   { return false }
func (MaxMeasuredRelativeHumidity) Reportable() bool { return false }
func (MaxMeasuredRelativeHumidity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxMeasuredRelativeHumidity) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxMeasuredRelativeHumidity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxMeasuredRelativeHumidity) AttrValue() zcl.Val  { return v.Value() }

func (MaxMeasuredRelativeHumidity) Name() string { return `Max Measured Relative Humidity` }
func (MaxMeasuredRelativeHumidity) Description() string {
	return `indicates the maximum value of MeasuredRH that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MaxMeasuredRelativeHumidity indicates the maximum value of MeasuredRH that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MaxMeasuredRelativeHumidity zcl.Zu16

func (v *MaxMeasuredRelativeHumidity) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxMeasuredRelativeHumidity) Value() zcl.Val     { return v }

func (v MaxMeasuredRelativeHumidity) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxMeasuredRelativeHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxMeasuredRelativeHumidity(*nt)
	return br, err
}

func (v MaxMeasuredRelativeHumidity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxMeasuredRelativeHumidity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxMeasuredRelativeHumidity(*a)
	return nil
}

func (v *MaxMeasuredRelativeHumidity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxMeasuredRelativeHumidity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxMeasuredRelativeHumidity) String() string {
	return zcl.PercentRelativeHumidity.Format(float64(v) / 100)
}

func (v MaxMeasuredRelativeHumidity) Scaled() float64 {
	return float64(v) / 100
}

const MaxMeasuredTemperatureAttr zcl.AttrID = 2

func (MaxMeasuredTemperature) ID() zcl.AttrID   { return MaxMeasuredTemperatureAttr }
func (MaxMeasuredTemperature) Readable() bool   { return true }
func (MaxMeasuredTemperature) Writable() bool   { return false }
func (MaxMeasuredTemperature) Reportable() bool { return false }
func (MaxMeasuredTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxMeasuredTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxMeasuredTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxMeasuredTemperature) AttrValue() zcl.Val  { return v.Value() }

func (MaxMeasuredTemperature) Name() string { return `Max Measured Temperature` }
func (MaxMeasuredTemperature) Description() string {
	return `indicates the maximum value of MeasuredTemperature that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// MaxMeasuredTemperature indicates the maximum value of MeasuredTemperature that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type MaxMeasuredTemperature zcl.Zs16

func (v *MaxMeasuredTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxMeasuredTemperature) Value() zcl.Val     { return v }

func (v MaxMeasuredTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxMeasuredTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxMeasuredTemperature(*nt)
	return br, err
}

func (v MaxMeasuredTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxMeasuredTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxMeasuredTemperature(*a)
	return nil
}

func (v *MaxMeasuredTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxMeasuredTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxMeasuredTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v MaxMeasuredTemperature) Scaled() float64 {
	return float64(v) / 100
}

func (MaxNumberOfIntervals) Name() string { return `Max Number Of Intervals` }
func (MaxNumberOfIntervals) Description() string {
	return `represents the maximum number of intervals the device is capable of returning in one command.
It is required MaxNumberofIntervals fit within the default Fragmentation ASDU size of 128 bytes,
or an optionally agreed upon larger Fragmentation ASDU size supported by both devices.`
}

// MaxNumberOfIntervals represents the maximum number of intervals the device is capable of returning in one command.
// It is required MaxNumberofIntervals fit within the default Fragmentation ASDU size of 128 bytes,
// or an optionally agreed upon larger Fragmentation ASDU size supported by both devices.
type MaxNumberOfIntervals zcl.Zu8

func (v *MaxNumberOfIntervals) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *MaxNumberOfIntervals) Value() zcl.Val     { return v }

func (v MaxNumberOfIntervals) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *MaxNumberOfIntervals) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxNumberOfIntervals(*nt)
	return br, err
}

func (v MaxNumberOfIntervals) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *MaxNumberOfIntervals) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxNumberOfIntervals(*a)
	return nil
}

func (v *MaxNumberOfIntervals) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = MaxNumberOfIntervals(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxNumberOfIntervals) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const Measured11ThHarmonicCurrentAttr zcl.AttrID = 780

func (Measured11ThHarmonicCurrent) ID() zcl.AttrID   { return Measured11ThHarmonicCurrentAttr }
func (Measured11ThHarmonicCurrent) Readable() bool   { return false }
func (Measured11ThHarmonicCurrent) Writable() bool   { return false }
func (Measured11ThHarmonicCurrent) Reportable() bool { return false }
func (Measured11ThHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Measured11ThHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v Measured11ThHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Measured11ThHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (Measured11ThHarmonicCurrent) Name() string { return `Measured 11th Harmonic Current` }
func (Measured11ThHarmonicCurrent) Description() string {
	return `represent the most recent 11th harmonic current reading in an AC frequency.`
}

// Measured11ThHarmonicCurrent represent the most recent 11th harmonic current reading in an AC frequency.
type Measured11ThHarmonicCurrent zcl.Zs16

func (v *Measured11ThHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *Measured11ThHarmonicCurrent) Value() zcl.Val     { return v }

func (v Measured11ThHarmonicCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *Measured11ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = Measured11ThHarmonicCurrent(*nt)
	return br, err
}

func (v Measured11ThHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *Measured11ThHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Measured11ThHarmonicCurrent(*a)
	return nil
}

func (v *Measured11ThHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = Measured11ThHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Measured11ThHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const Measured1StHarmonicCurrentAttr zcl.AttrID = 775

func (Measured1StHarmonicCurrent) ID() zcl.AttrID   { return Measured1StHarmonicCurrentAttr }
func (Measured1StHarmonicCurrent) Readable() bool   { return false }
func (Measured1StHarmonicCurrent) Writable() bool   { return false }
func (Measured1StHarmonicCurrent) Reportable() bool { return false }
func (Measured1StHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Measured1StHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v Measured1StHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Measured1StHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (Measured1StHarmonicCurrent) Name() string { return `Measured 1st Harmonic Current` }
func (Measured1StHarmonicCurrent) Description() string {
	return `represent the most recent 1st harmonic current reading in an AC frequency.
The unit for this measurement is 10 ^ 1stHarmonicCurrentMultiplier amperes.
If 1stHarmonicCurrentMultiplier is not implemented the unit is in amperes.
If the 1st harmonic current cannot be measured a value of 0x8000 is returned.
A positive value indicates the measured 1st harmonic current is positive, and a
negative value indicates that the measured 1st harmonic current is negative.`
}

// Measured1StHarmonicCurrent represent the most recent 1st harmonic current reading in an AC frequency.
// The unit for this measurement is 10 ^ 1stHarmonicCurrentMultiplier amperes.
// If 1stHarmonicCurrentMultiplier is not implemented the unit is in amperes.
// If the 1st harmonic current cannot be measured a value of 0x8000 is returned.
// A positive value indicates the measured 1st harmonic current is positive, and a
// negative value indicates that the measured 1st harmonic current is negative.
type Measured1StHarmonicCurrent zcl.Zs16

func (v *Measured1StHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *Measured1StHarmonicCurrent) Value() zcl.Val     { return v }

func (v Measured1StHarmonicCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *Measured1StHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = Measured1StHarmonicCurrent(*nt)
	return br, err
}

func (v Measured1StHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *Measured1StHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Measured1StHarmonicCurrent(*a)
	return nil
}

func (v *Measured1StHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = Measured1StHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Measured1StHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const Measured3RdHarmonicCurrentAttr zcl.AttrID = 776

func (Measured3RdHarmonicCurrent) ID() zcl.AttrID   { return Measured3RdHarmonicCurrentAttr }
func (Measured3RdHarmonicCurrent) Readable() bool   { return false }
func (Measured3RdHarmonicCurrent) Writable() bool   { return false }
func (Measured3RdHarmonicCurrent) Reportable() bool { return false }
func (Measured3RdHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Measured3RdHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v Measured3RdHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Measured3RdHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (Measured3RdHarmonicCurrent) Name() string { return `Measured 3rd Harmonic Current` }
func (Measured3RdHarmonicCurrent) Description() string {
	return `represent the most recent 3rd harmonic current reading in an AC frequency.`
}

// Measured3RdHarmonicCurrent represent the most recent 3rd harmonic current reading in an AC frequency.
type Measured3RdHarmonicCurrent zcl.Zs16

func (v *Measured3RdHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *Measured3RdHarmonicCurrent) Value() zcl.Val     { return v }

func (v Measured3RdHarmonicCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *Measured3RdHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = Measured3RdHarmonicCurrent(*nt)
	return br, err
}

func (v Measured3RdHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *Measured3RdHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Measured3RdHarmonicCurrent(*a)
	return nil
}

func (v *Measured3RdHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = Measured3RdHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Measured3RdHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const Measured5ThHarmonicCurrentAttr zcl.AttrID = 777

func (Measured5ThHarmonicCurrent) ID() zcl.AttrID   { return Measured5ThHarmonicCurrentAttr }
func (Measured5ThHarmonicCurrent) Readable() bool   { return false }
func (Measured5ThHarmonicCurrent) Writable() bool   { return false }
func (Measured5ThHarmonicCurrent) Reportable() bool { return false }
func (Measured5ThHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Measured5ThHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v Measured5ThHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Measured5ThHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (Measured5ThHarmonicCurrent) Name() string { return `Measured 5th Harmonic Current` }
func (Measured5ThHarmonicCurrent) Description() string {
	return `represent the most recent 5th harmonic current reading in an AC frequency.`
}

// Measured5ThHarmonicCurrent represent the most recent 5th harmonic current reading in an AC frequency.
type Measured5ThHarmonicCurrent zcl.Zs16

func (v *Measured5ThHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *Measured5ThHarmonicCurrent) Value() zcl.Val     { return v }

func (v Measured5ThHarmonicCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *Measured5ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = Measured5ThHarmonicCurrent(*nt)
	return br, err
}

func (v Measured5ThHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *Measured5ThHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Measured5ThHarmonicCurrent(*a)
	return nil
}

func (v *Measured5ThHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = Measured5ThHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Measured5ThHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const Measured7ThHarmonicCurrentAttr zcl.AttrID = 778

func (Measured7ThHarmonicCurrent) ID() zcl.AttrID   { return Measured7ThHarmonicCurrentAttr }
func (Measured7ThHarmonicCurrent) Readable() bool   { return false }
func (Measured7ThHarmonicCurrent) Writable() bool   { return false }
func (Measured7ThHarmonicCurrent) Reportable() bool { return false }
func (Measured7ThHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Measured7ThHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v Measured7ThHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Measured7ThHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (Measured7ThHarmonicCurrent) Name() string { return `Measured 7th Harmonic Current` }
func (Measured7ThHarmonicCurrent) Description() string {
	return `represent the most recent 7th harmonic current reading in an AC frequency.`
}

// Measured7ThHarmonicCurrent represent the most recent 7th harmonic current reading in an AC frequency.
type Measured7ThHarmonicCurrent zcl.Zs16

func (v *Measured7ThHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *Measured7ThHarmonicCurrent) Value() zcl.Val     { return v }

func (v Measured7ThHarmonicCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *Measured7ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = Measured7ThHarmonicCurrent(*nt)
	return br, err
}

func (v Measured7ThHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *Measured7ThHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Measured7ThHarmonicCurrent(*a)
	return nil
}

func (v *Measured7ThHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = Measured7ThHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Measured7ThHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const Measured9ThHarmonicCurrentAttr zcl.AttrID = 779

func (Measured9ThHarmonicCurrent) ID() zcl.AttrID   { return Measured9ThHarmonicCurrentAttr }
func (Measured9ThHarmonicCurrent) Readable() bool   { return false }
func (Measured9ThHarmonicCurrent) Writable() bool   { return false }
func (Measured9ThHarmonicCurrent) Reportable() bool { return false }
func (Measured9ThHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Measured9ThHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v Measured9ThHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Measured9ThHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (Measured9ThHarmonicCurrent) Name() string { return `Measured 9th Harmonic Current` }
func (Measured9ThHarmonicCurrent) Description() string {
	return `represent the most recent 9th harmonic current reading in an AC frequency.`
}

// Measured9ThHarmonicCurrent represent the most recent 9th harmonic current reading in an AC frequency.
type Measured9ThHarmonicCurrent zcl.Zs16

func (v *Measured9ThHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *Measured9ThHarmonicCurrent) Value() zcl.Val     { return v }

func (v Measured9ThHarmonicCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *Measured9ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = Measured9ThHarmonicCurrent(*nt)
	return br, err
}

func (v Measured9ThHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *Measured9ThHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Measured9ThHarmonicCurrent(*a)
	return nil
}

func (v *Measured9ThHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = Measured9ThHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Measured9ThHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const MeasuredFlowAttr zcl.AttrID = 0

func (MeasuredFlow) ID() zcl.AttrID   { return MeasuredFlowAttr }
func (MeasuredFlow) Readable() bool   { return true }
func (MeasuredFlow) Writable() bool   { return false }
func (MeasuredFlow) Reportable() bool { return true }
func (MeasuredFlow) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredFlow) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredFlow) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredFlow) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredFlow) Name() string        { return `Measured Flow` }
func (MeasuredFlow) Description() string { return `represents the flow in m^3/h` }

// MeasuredFlow represents the flow in m^3/h
type MeasuredFlow zcl.Zu16

func (v *MeasuredFlow) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MeasuredFlow) Value() zcl.Val     { return v }

func (v MeasuredFlow) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredFlow(*nt)
	return br, err
}

func (v MeasuredFlow) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MeasuredFlow) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredFlow(*a)
	return nil
}

func (v *MeasuredFlow) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MeasuredFlow(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredFlow) String() string {
	return zcl.CubicMetersPerHour.Format(float64(v) / 10)
}

func (v MeasuredFlow) Scaled() float64 {
	return float64(v) / 10
}

const MeasuredIlluminanceAttr zcl.AttrID = 0

func (MeasuredIlluminance) ID() zcl.AttrID   { return MeasuredIlluminanceAttr }
func (MeasuredIlluminance) Readable() bool   { return true }
func (MeasuredIlluminance) Writable() bool   { return false }
func (MeasuredIlluminance) Reportable() bool { return true }
func (MeasuredIlluminance) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredIlluminance) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredIlluminance) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredIlluminance) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredIlluminance) Name() string { return `Measured Illuminance` }
func (MeasuredIlluminance) Description() string {
	return `represents the Illuminance in Lux (symbol lx) as follows
MeasuredValue = 10,000 x log10 Flow + 1 where 1 lx
<= Flow <=3.576 Mlx, corresponding to a MeasuredValue
in the range 1 to 0xfffe`
}

// MeasuredIlluminance represents the Illuminance in Lux (symbol lx) as follows
// MeasuredValue = 10,000 x log10 Flow + 1 where 1 lx
// <= Flow <=3.576 Mlx, corresponding to a MeasuredValue
// in the range 1 to 0xfffe
type MeasuredIlluminance zcl.Zu16

func (v *MeasuredIlluminance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MeasuredIlluminance) Value() zcl.Val     { return v }

func (v MeasuredIlluminance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MeasuredIlluminance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredIlluminance(*nt)
	return br, err
}

func (v MeasuredIlluminance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MeasuredIlluminance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredIlluminance(*a)
	return nil
}

func (v *MeasuredIlluminance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MeasuredIlluminance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredIlluminance) String() string {
	return zcl.Luxes.Format(float64(v))
}

const MeasuredPhase11ThHarmonicCurrentAttr zcl.AttrID = 786

func (MeasuredPhase11ThHarmonicCurrent) ID() zcl.AttrID   { return MeasuredPhase11ThHarmonicCurrentAttr }
func (MeasuredPhase11ThHarmonicCurrent) Readable() bool   { return false }
func (MeasuredPhase11ThHarmonicCurrent) Writable() bool   { return false }
func (MeasuredPhase11ThHarmonicCurrent) Reportable() bool { return false }
func (MeasuredPhase11ThHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredPhase11ThHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredPhase11ThHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredPhase11ThHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredPhase11ThHarmonicCurrent) Name() string { return `Measured Phase 11th Harmonic Current` }
func (MeasuredPhase11ThHarmonicCurrent) Description() string {
	return `represent the most recent phase of the 11th harmonic current reading in an AC frequency.`
}

// MeasuredPhase11ThHarmonicCurrent represent the most recent phase of the 11th harmonic current reading in an AC frequency.
type MeasuredPhase11ThHarmonicCurrent zcl.Zs16

func (v *MeasuredPhase11ThHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MeasuredPhase11ThHarmonicCurrent) Value() zcl.Val     { return v }

func (v MeasuredPhase11ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(v).MarshalZcl()
}

func (v *MeasuredPhase11ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredPhase11ThHarmonicCurrent(*nt)
	return br, err
}

func (v MeasuredPhase11ThHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MeasuredPhase11ThHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredPhase11ThHarmonicCurrent(*a)
	return nil
}

func (v *MeasuredPhase11ThHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MeasuredPhase11ThHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredPhase11ThHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(v))
}

const MeasuredPhase1StHarmonicCurrentAttr zcl.AttrID = 781

func (MeasuredPhase1StHarmonicCurrent) ID() zcl.AttrID   { return MeasuredPhase1StHarmonicCurrentAttr }
func (MeasuredPhase1StHarmonicCurrent) Readable() bool   { return false }
func (MeasuredPhase1StHarmonicCurrent) Writable() bool   { return false }
func (MeasuredPhase1StHarmonicCurrent) Reportable() bool { return false }
func (MeasuredPhase1StHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredPhase1StHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredPhase1StHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredPhase1StHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredPhase1StHarmonicCurrent) Name() string { return `Measured Phase 1st Harmonic Current` }
func (MeasuredPhase1StHarmonicCurrent) Description() string {
	return `represent the most recent phase of the 1st harmonic current reading in an AC frequency.
The unit for this measurement is 10 ^ Phase1StHarmonicCurrentMultiplier degree.
If Phase1StHarmonicCurrentMultiplier is not implemented the unit is in degree.
If the phase of cannot be measured a value of 0x8000 is returned.
A positive value indicates the measured phase is prehurry,
and a negative value indicates that the measured phase is lagging.`
}

// MeasuredPhase1StHarmonicCurrent represent the most recent phase of the 1st harmonic current reading in an AC frequency.
// The unit for this measurement is 10 ^ Phase1StHarmonicCurrentMultiplier degree.
// If Phase1StHarmonicCurrentMultiplier is not implemented the unit is in degree.
// If the phase of cannot be measured a value of 0x8000 is returned.
// A positive value indicates the measured phase is prehurry,
// and a negative value indicates that the measured phase is lagging.
type MeasuredPhase1StHarmonicCurrent zcl.Zs16

func (v *MeasuredPhase1StHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MeasuredPhase1StHarmonicCurrent) Value() zcl.Val     { return v }

func (v MeasuredPhase1StHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(v).MarshalZcl()
}

func (v *MeasuredPhase1StHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredPhase1StHarmonicCurrent(*nt)
	return br, err
}

func (v MeasuredPhase1StHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MeasuredPhase1StHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredPhase1StHarmonicCurrent(*a)
	return nil
}

func (v *MeasuredPhase1StHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MeasuredPhase1StHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredPhase1StHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(v))
}

const MeasuredPhase3RdHarmonicCurrentAttr zcl.AttrID = 782

func (MeasuredPhase3RdHarmonicCurrent) ID() zcl.AttrID   { return MeasuredPhase3RdHarmonicCurrentAttr }
func (MeasuredPhase3RdHarmonicCurrent) Readable() bool   { return false }
func (MeasuredPhase3RdHarmonicCurrent) Writable() bool   { return false }
func (MeasuredPhase3RdHarmonicCurrent) Reportable() bool { return false }
func (MeasuredPhase3RdHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredPhase3RdHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredPhase3RdHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredPhase3RdHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredPhase3RdHarmonicCurrent) Name() string { return `Measured Phase 3rd Harmonic Current` }
func (MeasuredPhase3RdHarmonicCurrent) Description() string {
	return `represent the most recent phase of the 3rd harmonic current reading in an AC frequency.`
}

// MeasuredPhase3RdHarmonicCurrent represent the most recent phase of the 3rd harmonic current reading in an AC frequency.
type MeasuredPhase3RdHarmonicCurrent zcl.Zs16

func (v *MeasuredPhase3RdHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MeasuredPhase3RdHarmonicCurrent) Value() zcl.Val     { return v }

func (v MeasuredPhase3RdHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(v).MarshalZcl()
}

func (v *MeasuredPhase3RdHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredPhase3RdHarmonicCurrent(*nt)
	return br, err
}

func (v MeasuredPhase3RdHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MeasuredPhase3RdHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredPhase3RdHarmonicCurrent(*a)
	return nil
}

func (v *MeasuredPhase3RdHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MeasuredPhase3RdHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredPhase3RdHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(v))
}

const MeasuredPhase5ThHarmonicCurrentAttr zcl.AttrID = 783

func (MeasuredPhase5ThHarmonicCurrent) ID() zcl.AttrID   { return MeasuredPhase5ThHarmonicCurrentAttr }
func (MeasuredPhase5ThHarmonicCurrent) Readable() bool   { return false }
func (MeasuredPhase5ThHarmonicCurrent) Writable() bool   { return false }
func (MeasuredPhase5ThHarmonicCurrent) Reportable() bool { return false }
func (MeasuredPhase5ThHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredPhase5ThHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredPhase5ThHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredPhase5ThHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredPhase5ThHarmonicCurrent) Name() string { return `Measured Phase 5th Harmonic Current` }
func (MeasuredPhase5ThHarmonicCurrent) Description() string {
	return `represent the most recent phase of the 5th harmonic current reading in an AC frequency.`
}

// MeasuredPhase5ThHarmonicCurrent represent the most recent phase of the 5th harmonic current reading in an AC frequency.
type MeasuredPhase5ThHarmonicCurrent zcl.Zs16

func (v *MeasuredPhase5ThHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MeasuredPhase5ThHarmonicCurrent) Value() zcl.Val     { return v }

func (v MeasuredPhase5ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(v).MarshalZcl()
}

func (v *MeasuredPhase5ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredPhase5ThHarmonicCurrent(*nt)
	return br, err
}

func (v MeasuredPhase5ThHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MeasuredPhase5ThHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredPhase5ThHarmonicCurrent(*a)
	return nil
}

func (v *MeasuredPhase5ThHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MeasuredPhase5ThHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredPhase5ThHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(v))
}

const MeasuredPhase7ThHarmonicCurrentAttr zcl.AttrID = 784

func (MeasuredPhase7ThHarmonicCurrent) ID() zcl.AttrID   { return MeasuredPhase7ThHarmonicCurrentAttr }
func (MeasuredPhase7ThHarmonicCurrent) Readable() bool   { return false }
func (MeasuredPhase7ThHarmonicCurrent) Writable() bool   { return false }
func (MeasuredPhase7ThHarmonicCurrent) Reportable() bool { return false }
func (MeasuredPhase7ThHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredPhase7ThHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredPhase7ThHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredPhase7ThHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredPhase7ThHarmonicCurrent) Name() string { return `Measured Phase 7th Harmonic Current` }
func (MeasuredPhase7ThHarmonicCurrent) Description() string {
	return `represent the most recent phase of the 7th harmonic current reading in an AC frequency.`
}

// MeasuredPhase7ThHarmonicCurrent represent the most recent phase of the 7th harmonic current reading in an AC frequency.
type MeasuredPhase7ThHarmonicCurrent zcl.Zs16

func (v *MeasuredPhase7ThHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MeasuredPhase7ThHarmonicCurrent) Value() zcl.Val     { return v }

func (v MeasuredPhase7ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(v).MarshalZcl()
}

func (v *MeasuredPhase7ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredPhase7ThHarmonicCurrent(*nt)
	return br, err
}

func (v MeasuredPhase7ThHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MeasuredPhase7ThHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredPhase7ThHarmonicCurrent(*a)
	return nil
}

func (v *MeasuredPhase7ThHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MeasuredPhase7ThHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredPhase7ThHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(v))
}

const MeasuredPhase9ThHarmonicCurrentAttr zcl.AttrID = 785

func (MeasuredPhase9ThHarmonicCurrent) ID() zcl.AttrID   { return MeasuredPhase9ThHarmonicCurrentAttr }
func (MeasuredPhase9ThHarmonicCurrent) Readable() bool   { return false }
func (MeasuredPhase9ThHarmonicCurrent) Writable() bool   { return false }
func (MeasuredPhase9ThHarmonicCurrent) Reportable() bool { return false }
func (MeasuredPhase9ThHarmonicCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredPhase9ThHarmonicCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredPhase9ThHarmonicCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredPhase9ThHarmonicCurrent) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredPhase9ThHarmonicCurrent) Name() string { return `Measured Phase 9th Harmonic Current` }
func (MeasuredPhase9ThHarmonicCurrent) Description() string {
	return `represent the most recent phase of the 9th harmonic current reading in an AC frequency.`
}

// MeasuredPhase9ThHarmonicCurrent represent the most recent phase of the 9th harmonic current reading in an AC frequency.
type MeasuredPhase9ThHarmonicCurrent zcl.Zs16

func (v *MeasuredPhase9ThHarmonicCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MeasuredPhase9ThHarmonicCurrent) Value() zcl.Val     { return v }

func (v MeasuredPhase9ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(v).MarshalZcl()
}

func (v *MeasuredPhase9ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredPhase9ThHarmonicCurrent(*nt)
	return br, err
}

func (v MeasuredPhase9ThHarmonicCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MeasuredPhase9ThHarmonicCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredPhase9ThHarmonicCurrent(*a)
	return nil
}

func (v *MeasuredPhase9ThHarmonicCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MeasuredPhase9ThHarmonicCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredPhase9ThHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(v))
}

const MeasuredPressureAttr zcl.AttrID = 0

func (MeasuredPressure) ID() zcl.AttrID   { return MeasuredPressureAttr }
func (MeasuredPressure) Readable() bool   { return true }
func (MeasuredPressure) Writable() bool   { return false }
func (MeasuredPressure) Reportable() bool { return true }
func (MeasuredPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredPressure) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredPressure) Name() string { return `Measured Pressure` }
func (MeasuredPressure) Description() string {
	return `represents the temperature in degrees DegreesCelsius`
}

// MeasuredPressure represents the temperature in degrees DegreesCelsius
type MeasuredPressure zcl.Zs16

func (v *MeasuredPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MeasuredPressure) Value() zcl.Val     { return v }

func (v MeasuredPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MeasuredPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredPressure(*nt)
	return br, err
}

func (v MeasuredPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MeasuredPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredPressure(*a)
	return nil
}

func (v *MeasuredPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MeasuredPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MeasuredPressure) Scaled() float64 {
	return float64(v) / 10
}

const MeasuredRelativeHumidityAttr zcl.AttrID = 0

func (MeasuredRelativeHumidity) ID() zcl.AttrID   { return MeasuredRelativeHumidityAttr }
func (MeasuredRelativeHumidity) Readable() bool   { return true }
func (MeasuredRelativeHumidity) Writable() bool   { return false }
func (MeasuredRelativeHumidity) Reportable() bool { return true }
func (MeasuredRelativeHumidity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredRelativeHumidity) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredRelativeHumidity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredRelativeHumidity) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredRelativeHumidity) Name() string { return `Measured Relative Humidity` }
func (MeasuredRelativeHumidity) Description() string {
	return `represents the relative humidity in percent`
}

// MeasuredRelativeHumidity represents the relative humidity in percent
type MeasuredRelativeHumidity zcl.Zu16

func (v *MeasuredRelativeHumidity) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MeasuredRelativeHumidity) Value() zcl.Val     { return v }

func (v MeasuredRelativeHumidity) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MeasuredRelativeHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredRelativeHumidity(*nt)
	return br, err
}

func (v MeasuredRelativeHumidity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MeasuredRelativeHumidity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredRelativeHumidity(*a)
	return nil
}

func (v *MeasuredRelativeHumidity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MeasuredRelativeHumidity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredRelativeHumidity) String() string {
	return zcl.PercentRelativeHumidity.Format(float64(v) / 100)
}

func (v MeasuredRelativeHumidity) Scaled() float64 {
	return float64(v) / 100
}

const MeasuredTemperatureAttr zcl.AttrID = 0

func (MeasuredTemperature) ID() zcl.AttrID   { return MeasuredTemperatureAttr }
func (MeasuredTemperature) Readable() bool   { return true }
func (MeasuredTemperature) Writable() bool   { return false }
func (MeasuredTemperature) Reportable() bool { return true }
func (MeasuredTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredTemperature) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredTemperature) Name() string { return `Measured Temperature` }
func (MeasuredTemperature) Description() string {
	return `represents the temperature in degrees DegreesCelsius`
}

// MeasuredTemperature represents the temperature in degrees DegreesCelsius
type MeasuredTemperature zcl.Zs16

func (v *MeasuredTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MeasuredTemperature) Value() zcl.Val     { return v }

func (v MeasuredTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MeasuredTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredTemperature(*nt)
	return br, err
}

func (v MeasuredTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MeasuredTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredTemperature(*a)
	return nil
}

func (v *MeasuredTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MeasuredTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v MeasuredTemperature) Scaled() float64 {
	return float64(v) / 100
}

func (MeasurementResponseStatus) Name() string        { return `Measurement Response Status` }
func (MeasurementResponseStatus) Description() string { return `` }

type MeasurementResponseStatus zcl.Zenum8

func (v *MeasurementResponseStatus) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *MeasurementResponseStatus) Value() zcl.Val     { return v }

func (v MeasurementResponseStatus) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *MeasurementResponseStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasurementResponseStatus(*nt)
	return br, err
}

func (v MeasurementResponseStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *MeasurementResponseStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasurementResponseStatus(*a)
	return nil
}

func (v *MeasurementResponseStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = MeasurementResponseStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasurementResponseStatus) String() string {
	switch v {
	case 0x00:
		return "Success"
	case 0x01:
		return "Attribute Profile not supported"
	case 0x02:
		return "Invalid Start Time"
	case 0x03:
		return "More intervals requested than can be returned"
	case 0x04:
		return "No intervals available for the requested time"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v MeasurementResponseStatus) IsSuccess() bool                                 { return v == 0x00 }
func (v MeasurementResponseStatus) IsAttributeProfileNotSupported() bool            { return v == 0x01 }
func (v MeasurementResponseStatus) IsInvalidStartTime() bool                        { return v == 0x02 }
func (v MeasurementResponseStatus) IsMoreIntervalsRequestedThanCanBeReturned() bool { return v == 0x03 }
func (v MeasurementResponseStatus) IsNoIntervalsAvailableForTheRequestedTime() bool { return v == 0x04 }
func (v *MeasurementResponseStatus) SetSuccess()                                    { *v = 0x00 }
func (v *MeasurementResponseStatus) SetAttributeProfileNotSupported()               { *v = 0x01 }
func (v *MeasurementResponseStatus) SetInvalidStartTime()                           { *v = 0x02 }
func (v *MeasurementResponseStatus) SetMoreIntervalsRequestedThanCanBeReturned()    { *v = 0x03 }
func (v *MeasurementResponseStatus) SetNoIntervalsAvailableForTheRequestedTime()    { *v = 0x04 }

func (MeasurementResponseStatus) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Success"},
		{Value: 0x01, Name: "Attribute Profile not supported"},
		{Value: 0x02, Name: "Invalid Start Time"},
		{Value: 0x03, Name: "More intervals requested than can be returned"},
		{Value: 0x04, Name: "No intervals available for the requested time"},
	}
}

const MinMeasuredFlowAttr zcl.AttrID = 1

func (MinMeasuredFlow) ID() zcl.AttrID   { return MinMeasuredFlowAttr }
func (MinMeasuredFlow) Readable() bool   { return true }
func (MinMeasuredFlow) Writable() bool   { return false }
func (MinMeasuredFlow) Reportable() bool { return false }
func (MinMeasuredFlow) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinMeasuredFlow) AttrID() zcl.AttrID   { return v.ID() }
func (v MinMeasuredFlow) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinMeasuredFlow) AttrValue() zcl.Val  { return v.Value() }

func (MinMeasuredFlow) Name() string { return `Min Measured Flow` }
func (MinMeasuredFlow) Description() string {
	return `indicates the minimum value of MeasuredFlow that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MinMeasuredFlow indicates the minimum value of MeasuredFlow that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MinMeasuredFlow zcl.Zu16

func (v *MinMeasuredFlow) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MinMeasuredFlow) Value() zcl.Val     { return v }

func (v MinMeasuredFlow) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MinMeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinMeasuredFlow(*nt)
	return br, err
}

func (v MinMeasuredFlow) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MinMeasuredFlow) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinMeasuredFlow(*a)
	return nil
}

func (v *MinMeasuredFlow) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MinMeasuredFlow(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinMeasuredFlow) String() string {
	return zcl.CubicMetersPerHour.Format(float64(v) / 10)
}

func (v MinMeasuredFlow) Scaled() float64 {
	return float64(v) / 10
}

const MinMeasuredIlluminanceAttr zcl.AttrID = 1

func (MinMeasuredIlluminance) ID() zcl.AttrID   { return MinMeasuredIlluminanceAttr }
func (MinMeasuredIlluminance) Readable() bool   { return true }
func (MinMeasuredIlluminance) Writable() bool   { return false }
func (MinMeasuredIlluminance) Reportable() bool { return false }
func (MinMeasuredIlluminance) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinMeasuredIlluminance) AttrID() zcl.AttrID   { return v.ID() }
func (v MinMeasuredIlluminance) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinMeasuredIlluminance) AttrValue() zcl.Val  { return v.Value() }

func (MinMeasuredIlluminance) Name() string { return `Min Measured Illuminance` }
func (MinMeasuredIlluminance) Description() string {
	return `indicates the minimum value of MeasuredIlluminance that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MinMeasuredIlluminance indicates the minimum value of MeasuredIlluminance that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MinMeasuredIlluminance zcl.Zu16

func (v *MinMeasuredIlluminance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MinMeasuredIlluminance) Value() zcl.Val     { return v }

func (v MinMeasuredIlluminance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MinMeasuredIlluminance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinMeasuredIlluminance(*nt)
	return br, err
}

func (v MinMeasuredIlluminance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MinMeasuredIlluminance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinMeasuredIlluminance(*a)
	return nil
}

func (v *MinMeasuredIlluminance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MinMeasuredIlluminance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinMeasuredIlluminance) String() string {
	return zcl.Luxes.Format(float64(v))
}

const MinMeasuredPressureAttr zcl.AttrID = 1

func (MinMeasuredPressure) ID() zcl.AttrID   { return MinMeasuredPressureAttr }
func (MinMeasuredPressure) Readable() bool   { return true }
func (MinMeasuredPressure) Writable() bool   { return false }
func (MinMeasuredPressure) Reportable() bool { return false }
func (MinMeasuredPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinMeasuredPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MinMeasuredPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinMeasuredPressure) AttrValue() zcl.Val  { return v.Value() }

func (MinMeasuredPressure) Name() string { return `Min Measured Pressure` }
func (MinMeasuredPressure) Description() string {
	return `indicates the minimum value of MeasuredPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// MinMeasuredPressure indicates the minimum value of MeasuredPressure that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type MinMeasuredPressure zcl.Zs16

func (v *MinMeasuredPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MinMeasuredPressure) Value() zcl.Val     { return v }

func (v MinMeasuredPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MinMeasuredPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinMeasuredPressure(*nt)
	return br, err
}

func (v MinMeasuredPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MinMeasuredPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinMeasuredPressure(*a)
	return nil
}

func (v *MinMeasuredPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MinMeasuredPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinMeasuredPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MinMeasuredPressure) Scaled() float64 {
	return float64(v) / 10
}

const MinMeasuredRelativeHumidityAttr zcl.AttrID = 1

func (MinMeasuredRelativeHumidity) ID() zcl.AttrID   { return MinMeasuredRelativeHumidityAttr }
func (MinMeasuredRelativeHumidity) Readable() bool   { return true }
func (MinMeasuredRelativeHumidity) Writable() bool   { return false }
func (MinMeasuredRelativeHumidity) Reportable() bool { return false }
func (MinMeasuredRelativeHumidity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinMeasuredRelativeHumidity) AttrID() zcl.AttrID   { return v.ID() }
func (v MinMeasuredRelativeHumidity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinMeasuredRelativeHumidity) AttrValue() zcl.Val  { return v.Value() }

func (MinMeasuredRelativeHumidity) Name() string { return `Min Measured Relative Humidity` }
func (MinMeasuredRelativeHumidity) Description() string {
	return `indicates the minimum value of MeasuredRH that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MinMeasuredRelativeHumidity indicates the minimum value of MeasuredRH that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MinMeasuredRelativeHumidity zcl.Zu16

func (v *MinMeasuredRelativeHumidity) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MinMeasuredRelativeHumidity) Value() zcl.Val     { return v }

func (v MinMeasuredRelativeHumidity) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MinMeasuredRelativeHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinMeasuredRelativeHumidity(*nt)
	return br, err
}

func (v MinMeasuredRelativeHumidity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MinMeasuredRelativeHumidity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinMeasuredRelativeHumidity(*a)
	return nil
}

func (v *MinMeasuredRelativeHumidity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MinMeasuredRelativeHumidity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinMeasuredRelativeHumidity) String() string {
	return zcl.PercentRelativeHumidity.Format(float64(v) / 100)
}

func (v MinMeasuredRelativeHumidity) Scaled() float64 {
	return float64(v) / 100
}

const MinMeasuredTemperatureAttr zcl.AttrID = 1

func (MinMeasuredTemperature) ID() zcl.AttrID   { return MinMeasuredTemperatureAttr }
func (MinMeasuredTemperature) Readable() bool   { return true }
func (MinMeasuredTemperature) Writable() bool   { return false }
func (MinMeasuredTemperature) Reportable() bool { return false }
func (MinMeasuredTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinMeasuredTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v MinMeasuredTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinMeasuredTemperature) AttrValue() zcl.Val  { return v.Value() }

func (MinMeasuredTemperature) Name() string { return `Min Measured Temperature` }
func (MinMeasuredTemperature) Description() string {
	return `indicates the minimum value of MeasuredTemperature that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// MinMeasuredTemperature indicates the minimum value of MeasuredTemperature that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type MinMeasuredTemperature zcl.Zs16

func (v *MinMeasuredTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MinMeasuredTemperature) Value() zcl.Val     { return v }

func (v MinMeasuredTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MinMeasuredTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinMeasuredTemperature(*nt)
	return br, err
}

func (v MinMeasuredTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MinMeasuredTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinMeasuredTemperature(*a)
	return nil
}

func (v *MinMeasuredTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MinMeasuredTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinMeasuredTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v MinMeasuredTemperature) Scaled() float64 {
	return float64(v) / 100
}

const NeutralCurrentAttr zcl.AttrID = 771

func (NeutralCurrent) ID() zcl.AttrID   { return NeutralCurrentAttr }
func (NeutralCurrent) Readable() bool   { return false }
func (NeutralCurrent) Writable() bool   { return false }
func (NeutralCurrent) Reportable() bool { return false }
func (NeutralCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NeutralCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v NeutralCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NeutralCurrent) AttrValue() zcl.Val  { return v.Value() }

func (NeutralCurrent) Name() string { return `Neutral Current` }
func (NeutralCurrent) Description() string {
	return `represents the AC neutral (Line-Out) current value at the moment in time the
attribute is read. If the instantaneous current cannot be measured, a value of 0xFFFF is returned`
}

// NeutralCurrent represents the AC neutral (Line-Out) current value at the moment in time the
// attribute is read. If the instantaneous current cannot be measured, a value of 0xFFFF is returned
type NeutralCurrent zcl.Zu16

func (v *NeutralCurrent) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *NeutralCurrent) Value() zcl.Val     { return v }

func (v NeutralCurrent) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *NeutralCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = NeutralCurrent(*nt)
	return br, err
}

func (v NeutralCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *NeutralCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NeutralCurrent(*a)
	return nil
}

func (v *NeutralCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = NeutralCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NeutralCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

func (NumberOfIntervals) Name() string { return `Number Of Intervals` }
func (NumberOfIntervals) Description() string {
	return `is the number of intervals requested or returned`
}

// NumberOfIntervals is the number of intervals requested or returned
type NumberOfIntervals zcl.Zu8

func (v *NumberOfIntervals) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *NumberOfIntervals) Value() zcl.Val     { return v }

func (v NumberOfIntervals) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *NumberOfIntervals) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = NumberOfIntervals(*nt)
	return br, err
}

func (v NumberOfIntervals) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *NumberOfIntervals) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NumberOfIntervals(*a)
	return nil
}

func (v *NumberOfIntervals) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = NumberOfIntervals(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NumberOfIntervals) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const OccupancyAttr zcl.AttrID = 0

func (Occupancy) ID() zcl.AttrID   { return OccupancyAttr }
func (Occupancy) Readable() bool   { return true }
func (Occupancy) Writable() bool   { return false }
func (Occupancy) Reportable() bool { return true }
func (Occupancy) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Occupancy) AttrID() zcl.AttrID   { return v.ID() }
func (v Occupancy) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Occupancy) AttrValue() zcl.Val  { return v.Value() }

func (Occupancy) Name() string        { return `Occupancy` }
func (Occupancy) Description() string { return `` }

type Occupancy zcl.Zbmp8

func (v *Occupancy) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *Occupancy) Value() zcl.Val     { return v }

func (v Occupancy) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *Occupancy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = Occupancy(*nt)
	return br, err
}

func (v Occupancy) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *Occupancy) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Occupancy(*a)
	return nil
}

func (v *Occupancy) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = Occupancy(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Occupancy) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Occupied")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v Occupancy) IsOccupied() bool    { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v *Occupancy) SetOccupied(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }

func (Occupancy) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Occupied"},
	}
}

const OccupancyTypeAttr zcl.AttrID = 1

func (OccupancyType) ID() zcl.AttrID   { return OccupancyTypeAttr }
func (OccupancyType) Readable() bool   { return true }
func (OccupancyType) Writable() bool   { return false }
func (OccupancyType) Reportable() bool { return false }
func (OccupancyType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OccupancyType) AttrID() zcl.AttrID   { return v.ID() }
func (v OccupancyType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OccupancyType) AttrValue() zcl.Val  { return v.Value() }

func (OccupancyType) Name() string        { return `Occupancy Type` }
func (OccupancyType) Description() string { return `` }

type OccupancyType zcl.Zenum8

func (v *OccupancyType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *OccupancyType) Value() zcl.Val     { return v }

func (v OccupancyType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *OccupancyType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = OccupancyType(*nt)
	return br, err
}

func (v OccupancyType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *OccupancyType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OccupancyType(*a)
	return nil
}

func (v *OccupancyType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = OccupancyType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OccupancyType) String() string {
	switch v {
	case 0x00:
		return "PIR"
	case 0x01:
		return "Ultrasonic"
	case 0x03:
		return "PIR and ultrasonic"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v OccupancyType) IsPir() bool              { return v == 0x00 }
func (v OccupancyType) IsUltrasonic() bool       { return v == 0x01 }
func (v OccupancyType) IsPirAndUltrasonic() bool { return v == 0x03 }
func (v *OccupancyType) SetPir()                 { *v = 0x00 }
func (v *OccupancyType) SetUltrasonic()          { *v = 0x01 }
func (v *OccupancyType) SetPirAndUltrasonic()    { *v = 0x03 }

func (OccupancyType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "PIR"},
		{Value: 0x01, Name: "Ultrasonic"},
		{Value: 0x03, Name: "PIR and ultrasonic"},
	}
}

const PirOccupiedToUnoccupiedDelayAttr zcl.AttrID = 16

func (PirOccupiedToUnoccupiedDelay) ID() zcl.AttrID   { return PirOccupiedToUnoccupiedDelayAttr }
func (PirOccupiedToUnoccupiedDelay) Readable() bool   { return true }
func (PirOccupiedToUnoccupiedDelay) Writable() bool   { return true }
func (PirOccupiedToUnoccupiedDelay) Reportable() bool { return false }
func (PirOccupiedToUnoccupiedDelay) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PirOccupiedToUnoccupiedDelay) AttrID() zcl.AttrID   { return v.ID() }
func (v PirOccupiedToUnoccupiedDelay) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PirOccupiedToUnoccupiedDelay) AttrValue() zcl.Val  { return v.Value() }

func (PirOccupiedToUnoccupiedDelay) Name() string { return `PIR Occupied To Unoccupied Delay` }
func (PirOccupiedToUnoccupiedDelay) Description() string {
	return `is 16 bits in length and specifies the time delay, in seconds, before the PIR sensor
changes to its unoccupied state after the last detection of movement in the sensed area`
}

// PirOccupiedToUnoccupiedDelay is 16 bits in length and specifies the time delay, in seconds, before the PIR sensor
// changes to its unoccupied state after the last detection of movement in the sensed area
type PirOccupiedToUnoccupiedDelay zcl.Zu16

func (v *PirOccupiedToUnoccupiedDelay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PirOccupiedToUnoccupiedDelay) Value() zcl.Val     { return v }

func (v PirOccupiedToUnoccupiedDelay) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PirOccupiedToUnoccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PirOccupiedToUnoccupiedDelay(*nt)
	return br, err
}

func (v PirOccupiedToUnoccupiedDelay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PirOccupiedToUnoccupiedDelay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PirOccupiedToUnoccupiedDelay(*a)
	return nil
}

func (v *PirOccupiedToUnoccupiedDelay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PirOccupiedToUnoccupiedDelay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PirOccupiedToUnoccupiedDelay) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PirUnoccupiedToOccupiedDelayAttr zcl.AttrID = 17

func (PirUnoccupiedToOccupiedDelay) ID() zcl.AttrID   { return PirUnoccupiedToOccupiedDelayAttr }
func (PirUnoccupiedToOccupiedDelay) Readable() bool   { return true }
func (PirUnoccupiedToOccupiedDelay) Writable() bool   { return true }
func (PirUnoccupiedToOccupiedDelay) Reportable() bool { return false }
func (PirUnoccupiedToOccupiedDelay) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PirUnoccupiedToOccupiedDelay) AttrID() zcl.AttrID   { return v.ID() }
func (v PirUnoccupiedToOccupiedDelay) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PirUnoccupiedToOccupiedDelay) AttrValue() zcl.Val  { return v.Value() }

func (PirUnoccupiedToOccupiedDelay) Name() string { return `PIR Unoccupied To Occupied Delay` }
func (PirUnoccupiedToOccupiedDelay) Description() string {
	return `is 16 bits in length and specifies the time delay, in seconds, before the PIR sensor changes
to its occupied state after the detection of movement in the sensed area. This attribute is
mandatory if the PIRUnoccupiedToOccupiedThreshold attribute is implemented`
}

// PirUnoccupiedToOccupiedDelay is 16 bits in length and specifies the time delay, in seconds, before the PIR sensor changes
// to its occupied state after the detection of movement in the sensed area. This attribute is
// mandatory if the PIRUnoccupiedToOccupiedThreshold attribute is implemented
type PirUnoccupiedToOccupiedDelay zcl.Zu16

func (v *PirUnoccupiedToOccupiedDelay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PirUnoccupiedToOccupiedDelay) Value() zcl.Val     { return v }

func (v PirUnoccupiedToOccupiedDelay) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PirUnoccupiedToOccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PirUnoccupiedToOccupiedDelay(*nt)
	return br, err
}

func (v PirUnoccupiedToOccupiedDelay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PirUnoccupiedToOccupiedDelay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PirUnoccupiedToOccupiedDelay(*a)
	return nil
}

func (v *PirUnoccupiedToOccupiedDelay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PirUnoccupiedToOccupiedDelay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PirUnoccupiedToOccupiedDelay) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PirUnoccupiedToOccupiedThresholdAttr zcl.AttrID = 18

func (PirUnoccupiedToOccupiedThreshold) ID() zcl.AttrID   { return PirUnoccupiedToOccupiedThresholdAttr }
func (PirUnoccupiedToOccupiedThreshold) Readable() bool   { return true }
func (PirUnoccupiedToOccupiedThreshold) Writable() bool   { return true }
func (PirUnoccupiedToOccupiedThreshold) Reportable() bool { return false }
func (PirUnoccupiedToOccupiedThreshold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PirUnoccupiedToOccupiedThreshold) AttrID() zcl.AttrID   { return v.ID() }
func (v PirUnoccupiedToOccupiedThreshold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PirUnoccupiedToOccupiedThreshold) AttrValue() zcl.Val  { return v.Value() }

func (PirUnoccupiedToOccupiedThreshold) Name() string { return `PIR Unoccupied To Occupied Threshold` }
func (PirUnoccupiedToOccupiedThreshold) Description() string {
	return `is 8 bits in length and specifies the number of movement detection events that must occur in the period PIRUnoccupiedToOccupiedDelay, before the PIR sensor changes to its occupied state. This attribute is mandatory if the PIRUnoccupiedToOccupiedDelay attribute is implemented`
}

// PirUnoccupiedToOccupiedThreshold is 8 bits in length and specifies the number of movement detection events that must occur in the period PIRUnoccupiedToOccupiedDelay, before the PIR sensor changes to its occupied state. This attribute is mandatory if the PIRUnoccupiedToOccupiedDelay attribute is implemented
type PirUnoccupiedToOccupiedThreshold zcl.Zu8

func (v *PirUnoccupiedToOccupiedThreshold) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *PirUnoccupiedToOccupiedThreshold) Value() zcl.Val     { return v }

func (v PirUnoccupiedToOccupiedThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(v).MarshalZcl()
}

func (v *PirUnoccupiedToOccupiedThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = PirUnoccupiedToOccupiedThreshold(*nt)
	return br, err
}

func (v PirUnoccupiedToOccupiedThreshold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *PirUnoccupiedToOccupiedThreshold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PirUnoccupiedToOccupiedThreshold(*a)
	return nil
}

func (v *PirUnoccupiedToOccupiedThreshold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = PirUnoccupiedToOccupiedThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PirUnoccupiedToOccupiedThreshold) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const PhaseHarmonicCurrentMultiplierAttr zcl.AttrID = 1029

func (PhaseHarmonicCurrentMultiplier) ID() zcl.AttrID   { return PhaseHarmonicCurrentMultiplierAttr }
func (PhaseHarmonicCurrentMultiplier) Readable() bool   { return false }
func (PhaseHarmonicCurrentMultiplier) Writable() bool   { return false }
func (PhaseHarmonicCurrentMultiplier) Reportable() bool { return false }
func (PhaseHarmonicCurrentMultiplier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PhaseHarmonicCurrentMultiplier) AttrID() zcl.AttrID   { return v.ID() }
func (v PhaseHarmonicCurrentMultiplier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PhaseHarmonicCurrentMultiplier) AttrValue() zcl.Val  { return v.Value() }

func (PhaseHarmonicCurrentMultiplier) Name() string { return `Phase Harmonic Current Multiplier` }
func (PhaseHarmonicCurrentMultiplier) Description() string {
	return `Represents the unit value for the MeasuredPhaseNthHarmonicCurrent attribute in the format
MeasuredPhaseNthHarmonicCurrent * 10 ^ PhaseHarmonicCurrentMultiplier degrees.`
}

// PhaseHarmonicCurrentMultiplier Represents the unit value for the MeasuredPhaseNthHarmonicCurrent attribute in the format
// MeasuredPhaseNthHarmonicCurrent * 10 ^ PhaseHarmonicCurrentMultiplier degrees.
type PhaseHarmonicCurrentMultiplier zcl.Zs8

func (v *PhaseHarmonicCurrentMultiplier) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *PhaseHarmonicCurrentMultiplier) Value() zcl.Val     { return v }

func (v PhaseHarmonicCurrentMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *PhaseHarmonicCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = PhaseHarmonicCurrentMultiplier(*nt)
	return br, err
}

func (v PhaseHarmonicCurrentMultiplier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *PhaseHarmonicCurrentMultiplier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PhaseHarmonicCurrentMultiplier(*a)
	return nil
}

func (v *PhaseHarmonicCurrentMultiplier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = PhaseHarmonicCurrentMultiplier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PhaseHarmonicCurrentMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const PowerDivisorAttr zcl.AttrID = 1027

func (PowerDivisor) ID() zcl.AttrID   { return PowerDivisorAttr }
func (PowerDivisor) Readable() bool   { return false }
func (PowerDivisor) Writable() bool   { return false }
func (PowerDivisor) Reportable() bool { return false }
func (PowerDivisor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerDivisor) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerDivisor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerDivisor) AttrValue() zcl.Val  { return v.Value() }

func (PowerDivisor) Name() string { return `Power Divisor` }
func (PowerDivisor) Description() string {
	return `Provides a value to divide against the results of applying the Multiplier attribute against a raw or
uncompensated sensor count of power being measured by the metering device. If present, this attribute must
be applied against all demand/power values to derive the delivered and received values expressed in the
specified units. This attribute must be used in conjunction with the PowerMultiplier attribute.`
}

// PowerDivisor Provides a value to divide against the results of applying the Multiplier attribute against a raw or
// uncompensated sensor count of power being measured by the metering device. If present, this attribute must
// be applied against all demand/power values to derive the delivered and received values expressed in the
// specified units. This attribute must be used in conjunction with the PowerMultiplier attribute.
type PowerDivisor zcl.Zu32

func (v *PowerDivisor) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *PowerDivisor) Value() zcl.Val     { return v }

func (v PowerDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *PowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerDivisor(*nt)
	return br, err
}

func (v PowerDivisor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *PowerDivisor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerDivisor(*a)
	return nil
}

func (v *PowerDivisor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = PowerDivisor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

const PowerFactorAttr zcl.AttrID = 1296

func (PowerFactor) ID() zcl.AttrID   { return PowerFactorAttr }
func (PowerFactor) Readable() bool   { return false }
func (PowerFactor) Writable() bool   { return false }
func (PowerFactor) Reportable() bool { return false }
func (PowerFactor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerFactor) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerFactor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerFactor) AttrValue() zcl.Val  { return v.Value() }

func (PowerFactor) Name() string { return `Power Factor` }
func (PowerFactor) Description() string {
	return `contains the single phase or PhaseA, Power Factor ratio in 1/100ths`
}

// PowerFactor contains the single phase or PhaseA, Power Factor ratio in 1/100ths
type PowerFactor zcl.Zs8

func (v *PowerFactor) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *PowerFactor) Value() zcl.Val     { return v }

func (v PowerFactor) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *PowerFactor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerFactor(*nt)
	return br, err
}

func (v PowerFactor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *PowerFactor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerFactor(*a)
	return nil
}

func (v *PowerFactor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = PowerFactor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerFactor) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const PowerFactorPhBAttr zcl.AttrID = 2320

func (PowerFactorPhB) ID() zcl.AttrID   { return PowerFactorPhBAttr }
func (PowerFactorPhB) Readable() bool   { return false }
func (PowerFactorPhB) Writable() bool   { return false }
func (PowerFactorPhB) Reportable() bool { return false }
func (PowerFactorPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerFactorPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerFactorPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerFactorPhB) AttrValue() zcl.Val  { return v.Value() }

func (PowerFactorPhB) Name() string        { return `Power Factor Ph B` }
func (PowerFactorPhB) Description() string { return `` }

type PowerFactorPhB zcl.Zs8

func (v *PowerFactorPhB) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *PowerFactorPhB) Value() zcl.Val     { return v }

func (v PowerFactorPhB) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *PowerFactorPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerFactorPhB(*nt)
	return br, err
}

func (v PowerFactorPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *PowerFactorPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerFactorPhB(*a)
	return nil
}

func (v *PowerFactorPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = PowerFactorPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerFactorPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const PowerFactorPhCAttr zcl.AttrID = 2576

func (PowerFactorPhC) ID() zcl.AttrID   { return PowerFactorPhCAttr }
func (PowerFactorPhC) Readable() bool   { return false }
func (PowerFactorPhC) Writable() bool   { return false }
func (PowerFactorPhC) Reportable() bool { return false }
func (PowerFactorPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerFactorPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerFactorPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerFactorPhC) AttrValue() zcl.Val  { return v.Value() }

func (PowerFactorPhC) Name() string        { return `Power Factor Ph C` }
func (PowerFactorPhC) Description() string { return `` }

type PowerFactorPhC zcl.Zs8

func (v *PowerFactorPhC) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *PowerFactorPhC) Value() zcl.Val     { return v }

func (v PowerFactorPhC) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *PowerFactorPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerFactorPhC(*nt)
	return br, err
}

func (v PowerFactorPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *PowerFactorPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerFactorPhC(*a)
	return nil
}

func (v *PowerFactorPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = PowerFactorPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerFactorPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const PowerMultiplierAttr zcl.AttrID = 1026

func (PowerMultiplier) ID() zcl.AttrID   { return PowerMultiplierAttr }
func (PowerMultiplier) Readable() bool   { return false }
func (PowerMultiplier) Writable() bool   { return false }
func (PowerMultiplier) Reportable() bool { return false }
func (PowerMultiplier) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerMultiplier) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerMultiplier) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerMultiplier) AttrValue() zcl.Val  { return v.Value() }

func (PowerMultiplier) Name() string { return `Power Multiplier` }
func (PowerMultiplier) Description() string {
	return `Provides a value to be multiplied against a raw or uncompensated sensor count of power being measured by
the metering device. If present, this attribute must be applied against all power/demand values to derive the
delivered and received values expressed in the specified units. This attribute must be used in conjunction with
the PowerDivisor attribute.`
}

// PowerMultiplier Provides a value to be multiplied against a raw or uncompensated sensor count of power being measured by
// the metering device. If present, this attribute must be applied against all power/demand values to derive the
// delivered and received values expressed in the specified units. This attribute must be used in conjunction with
// the PowerDivisor attribute.
type PowerMultiplier zcl.Zu32

func (v *PowerMultiplier) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *PowerMultiplier) Value() zcl.Val     { return v }

func (v PowerMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *PowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerMultiplier(*nt)
	return br, err
}

func (v PowerMultiplier) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *PowerMultiplier) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerMultiplier(*a)
	return nil
}

func (v *PowerMultiplier) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = PowerMultiplier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(v))
}

func (ProfileCount) Name() string        { return `Profile Count` }
func (ProfileCount) Description() string { return `is the total number of supported profiles` }

// ProfileCount is the total number of supported profiles
type ProfileCount zcl.Zu8

func (v *ProfileCount) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ProfileCount) Value() zcl.Val     { return v }

func (v ProfileCount) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ProfileCount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ProfileCount(*nt)
	return br, err
}

func (v ProfileCount) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ProfileCount) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ProfileCount(*a)
	return nil
}

func (v *ProfileCount) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ProfileCount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ProfileCount) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (ProfileIntervalPeriod) Name() string { return `Profile Interval Period` }
func (ProfileIntervalPeriod) Description() string {
	return `represents the interval or time frame used to capture parameter for profiling purposes`
}

// ProfileIntervalPeriod represents the interval or time frame used to capture parameter for profiling purposes
type ProfileIntervalPeriod zcl.Zenum8

func (v *ProfileIntervalPeriod) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *ProfileIntervalPeriod) Value() zcl.Val     { return v }

func (v ProfileIntervalPeriod) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *ProfileIntervalPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = ProfileIntervalPeriod(*nt)
	return br, err
}

func (v ProfileIntervalPeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *ProfileIntervalPeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ProfileIntervalPeriod(*a)
	return nil
}

func (v *ProfileIntervalPeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = ProfileIntervalPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ProfileIntervalPeriod) String() string {
	switch v {
	case 0x00:
		return "Daily"
	case 0x01:
		return "60 minutes"
	case 0x02:
		return "30 minutes"
	case 0x03:
		return "15 minutes"
	case 0x04:
		return "10 minutes"
	case 0x05:
		return "7.5 minutes"
	case 0x06:
		return "5 minutes"
	case 0x07:
		return "2.5 minutes"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v ProfileIntervalPeriod) IsDaily() bool     { return v == 0x00 }
func (v ProfileIntervalPeriod) Is60Minutes() bool { return v == 0x01 }
func (v ProfileIntervalPeriod) Is30Minutes() bool { return v == 0x02 }
func (v ProfileIntervalPeriod) Is15Minutes() bool { return v == 0x03 }
func (v ProfileIntervalPeriod) Is10Minutes() bool { return v == 0x04 }
func (v ProfileIntervalPeriod) Is75Minutes() bool { return v == 0x05 }
func (v ProfileIntervalPeriod) Is5Minutes() bool  { return v == 0x06 }
func (v ProfileIntervalPeriod) Is25Minutes() bool { return v == 0x07 }
func (v *ProfileIntervalPeriod) SetDaily()        { *v = 0x00 }
func (v *ProfileIntervalPeriod) Set60Minutes()    { *v = 0x01 }
func (v *ProfileIntervalPeriod) Set30Minutes()    { *v = 0x02 }
func (v *ProfileIntervalPeriod) Set15Minutes()    { *v = 0x03 }
func (v *ProfileIntervalPeriod) Set10Minutes()    { *v = 0x04 }
func (v *ProfileIntervalPeriod) Set75Minutes()    { *v = 0x05 }
func (v *ProfileIntervalPeriod) Set5Minutes()     { *v = 0x06 }
func (v *ProfileIntervalPeriod) Set25Minutes()    { *v = 0x07 }

func (ProfileIntervalPeriod) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Daily"},
		{Value: 0x01, Name: "60 minutes"},
		{Value: 0x02, Name: "30 minutes"},
		{Value: 0x03, Name: "15 minutes"},
		{Value: 0x04, Name: "10 minutes"},
		{Value: 0x05, Name: "7.5 minutes"},
		{Value: 0x06, Name: "5 minutes"},
		{Value: 0x07, Name: "2.5 minutes"},
	}
}

const RmsCurrentAttr zcl.AttrID = 1288

func (RmsCurrent) ID() zcl.AttrID   { return RmsCurrentAttr }
func (RmsCurrent) Readable() bool   { return false }
func (RmsCurrent) Writable() bool   { return false }
func (RmsCurrent) Reportable() bool { return true }
func (RmsCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsCurrent) AttrValue() zcl.Val  { return v.Value() }

func (RmsCurrent) Name() string        { return `RMS Current` }
func (RmsCurrent) Description() string { return `Represents the most recent RMS current reading.` }

// RmsCurrent Represents the most recent RMS current reading.
type RmsCurrent zcl.Zu16

func (v *RmsCurrent) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsCurrent) Value() zcl.Val     { return v }

func (v RmsCurrent) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsCurrent(*nt)
	return br, err
}

func (v RmsCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsCurrent(*a)
	return nil
}

func (v *RmsCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const RmsCurrentMaxAttr zcl.AttrID = 1290

func (RmsCurrentMax) ID() zcl.AttrID   { return RmsCurrentMaxAttr }
func (RmsCurrentMax) Readable() bool   { return false }
func (RmsCurrentMax) Writable() bool   { return false }
func (RmsCurrentMax) Reportable() bool { return false }
func (RmsCurrentMax) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsCurrentMax) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsCurrentMax) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsCurrentMax) AttrValue() zcl.Val  { return v.Value() }

func (RmsCurrentMax) Name() string        { return `RMS Current Max` }
func (RmsCurrentMax) Description() string { return `Represents the highest RMS current value.` }

// RmsCurrentMax Represents the highest RMS current value.
type RmsCurrentMax zcl.Zu16

func (v *RmsCurrentMax) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsCurrentMax) Value() zcl.Val     { return v }

func (v RmsCurrentMax) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsCurrentMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsCurrentMax(*nt)
	return br, err
}

func (v RmsCurrentMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsCurrentMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsCurrentMax(*a)
	return nil
}

func (v *RmsCurrentMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsCurrentMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsCurrentMax) String() string {
	return zcl.Amperes.Format(float64(v))
}

const RmsCurrentMaxPhBAttr zcl.AttrID = 2314

func (RmsCurrentMaxPhB) ID() zcl.AttrID   { return RmsCurrentMaxPhBAttr }
func (RmsCurrentMaxPhB) Readable() bool   { return false }
func (RmsCurrentMaxPhB) Writable() bool   { return false }
func (RmsCurrentMaxPhB) Reportable() bool { return false }
func (RmsCurrentMaxPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsCurrentMaxPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsCurrentMaxPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsCurrentMaxPhB) AttrValue() zcl.Val  { return v.Value() }

func (RmsCurrentMaxPhB) Name() string        { return `RMS Current Max Ph B` }
func (RmsCurrentMaxPhB) Description() string { return `` }

type RmsCurrentMaxPhB zcl.Zu16

func (v *RmsCurrentMaxPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsCurrentMaxPhB) Value() zcl.Val     { return v }

func (v RmsCurrentMaxPhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsCurrentMaxPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsCurrentMaxPhB(*nt)
	return br, err
}

func (v RmsCurrentMaxPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsCurrentMaxPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsCurrentMaxPhB(*a)
	return nil
}

func (v *RmsCurrentMaxPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsCurrentMaxPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsCurrentMaxPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsCurrentMaxPhCAttr zcl.AttrID = 2570

func (RmsCurrentMaxPhC) ID() zcl.AttrID   { return RmsCurrentMaxPhCAttr }
func (RmsCurrentMaxPhC) Readable() bool   { return false }
func (RmsCurrentMaxPhC) Writable() bool   { return false }
func (RmsCurrentMaxPhC) Reportable() bool { return false }
func (RmsCurrentMaxPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsCurrentMaxPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsCurrentMaxPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsCurrentMaxPhC) AttrValue() zcl.Val  { return v.Value() }

func (RmsCurrentMaxPhC) Name() string        { return `RMS Current Max Ph C` }
func (RmsCurrentMaxPhC) Description() string { return `` }

type RmsCurrentMaxPhC zcl.Zu16

func (v *RmsCurrentMaxPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsCurrentMaxPhC) Value() zcl.Val     { return v }

func (v RmsCurrentMaxPhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsCurrentMaxPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsCurrentMaxPhC(*nt)
	return br, err
}

func (v RmsCurrentMaxPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsCurrentMaxPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsCurrentMaxPhC(*a)
	return nil
}

func (v *RmsCurrentMaxPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsCurrentMaxPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsCurrentMaxPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsCurrentMinAttr zcl.AttrID = 1289

func (RmsCurrentMin) ID() zcl.AttrID   { return RmsCurrentMinAttr }
func (RmsCurrentMin) Readable() bool   { return false }
func (RmsCurrentMin) Writable() bool   { return false }
func (RmsCurrentMin) Reportable() bool { return false }
func (RmsCurrentMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsCurrentMin) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsCurrentMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsCurrentMin) AttrValue() zcl.Val  { return v.Value() }

func (RmsCurrentMin) Name() string        { return `RMS Current Min` }
func (RmsCurrentMin) Description() string { return `Represents the lowest RMS current value.` }

// RmsCurrentMin Represents the lowest RMS current value.
type RmsCurrentMin zcl.Zu16

func (v *RmsCurrentMin) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsCurrentMin) Value() zcl.Val     { return v }

func (v RmsCurrentMin) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsCurrentMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsCurrentMin(*nt)
	return br, err
}

func (v RmsCurrentMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsCurrentMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsCurrentMin(*a)
	return nil
}

func (v *RmsCurrentMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsCurrentMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsCurrentMin) String() string {
	return zcl.Amperes.Format(float64(v))
}

const RmsCurrentMinPhBAttr zcl.AttrID = 2313

func (RmsCurrentMinPhB) ID() zcl.AttrID   { return RmsCurrentMinPhBAttr }
func (RmsCurrentMinPhB) Readable() bool   { return false }
func (RmsCurrentMinPhB) Writable() bool   { return false }
func (RmsCurrentMinPhB) Reportable() bool { return false }
func (RmsCurrentMinPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsCurrentMinPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsCurrentMinPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsCurrentMinPhB) AttrValue() zcl.Val  { return v.Value() }

func (RmsCurrentMinPhB) Name() string        { return `RMS Current Min Ph B` }
func (RmsCurrentMinPhB) Description() string { return `` }

type RmsCurrentMinPhB zcl.Zu16

func (v *RmsCurrentMinPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsCurrentMinPhB) Value() zcl.Val     { return v }

func (v RmsCurrentMinPhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsCurrentMinPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsCurrentMinPhB(*nt)
	return br, err
}

func (v RmsCurrentMinPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsCurrentMinPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsCurrentMinPhB(*a)
	return nil
}

func (v *RmsCurrentMinPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsCurrentMinPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsCurrentMinPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsCurrentMinPhCAttr zcl.AttrID = 2569

func (RmsCurrentMinPhC) ID() zcl.AttrID   { return RmsCurrentMinPhCAttr }
func (RmsCurrentMinPhC) Readable() bool   { return false }
func (RmsCurrentMinPhC) Writable() bool   { return false }
func (RmsCurrentMinPhC) Reportable() bool { return false }
func (RmsCurrentMinPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsCurrentMinPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsCurrentMinPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsCurrentMinPhC) AttrValue() zcl.Val  { return v.Value() }

func (RmsCurrentMinPhC) Name() string        { return `RMS Current Min Ph C` }
func (RmsCurrentMinPhC) Description() string { return `` }

type RmsCurrentMinPhC zcl.Zu16

func (v *RmsCurrentMinPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsCurrentMinPhC) Value() zcl.Val     { return v }

func (v RmsCurrentMinPhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsCurrentMinPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsCurrentMinPhC(*nt)
	return br, err
}

func (v RmsCurrentMinPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsCurrentMinPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsCurrentMinPhC(*a)
	return nil
}

func (v *RmsCurrentMinPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsCurrentMinPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsCurrentMinPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsCurrentPhBAttr zcl.AttrID = 2312

func (RmsCurrentPhB) ID() zcl.AttrID   { return RmsCurrentPhBAttr }
func (RmsCurrentPhB) Readable() bool   { return false }
func (RmsCurrentPhB) Writable() bool   { return false }
func (RmsCurrentPhB) Reportable() bool { return false }
func (RmsCurrentPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsCurrentPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsCurrentPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsCurrentPhB) AttrValue() zcl.Val  { return v.Value() }

func (RmsCurrentPhB) Name() string        { return `RMS Current Ph B` }
func (RmsCurrentPhB) Description() string { return `` }

type RmsCurrentPhB zcl.Zu16

func (v *RmsCurrentPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsCurrentPhB) Value() zcl.Val     { return v }

func (v RmsCurrentPhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsCurrentPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsCurrentPhB(*nt)
	return br, err
}

func (v RmsCurrentPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsCurrentPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsCurrentPhB(*a)
	return nil
}

func (v *RmsCurrentPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsCurrentPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsCurrentPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsCurrentPhCAttr zcl.AttrID = 2568

func (RmsCurrentPhC) ID() zcl.AttrID   { return RmsCurrentPhCAttr }
func (RmsCurrentPhC) Readable() bool   { return false }
func (RmsCurrentPhC) Writable() bool   { return false }
func (RmsCurrentPhC) Reportable() bool { return false }
func (RmsCurrentPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsCurrentPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsCurrentPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsCurrentPhC) AttrValue() zcl.Val  { return v.Value() }

func (RmsCurrentPhC) Name() string        { return `RMS Current Ph C` }
func (RmsCurrentPhC) Description() string { return `` }

type RmsCurrentPhC zcl.Zu16

func (v *RmsCurrentPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsCurrentPhC) Value() zcl.Val     { return v }

func (v RmsCurrentPhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsCurrentPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsCurrentPhC(*nt)
	return br, err
}

func (v RmsCurrentPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsCurrentPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsCurrentPhC(*a)
	return nil
}

func (v *RmsCurrentPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsCurrentPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsCurrentPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsExtremeOverVoltagePeriodAttr zcl.AttrID = 1300

func (RmsExtremeOverVoltagePeriod) ID() zcl.AttrID   { return RmsExtremeOverVoltagePeriodAttr }
func (RmsExtremeOverVoltagePeriod) Readable() bool   { return true }
func (RmsExtremeOverVoltagePeriod) Writable() bool   { return true }
func (RmsExtremeOverVoltagePeriod) Reportable() bool { return false }
func (RmsExtremeOverVoltagePeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsExtremeOverVoltagePeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsExtremeOverVoltagePeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsExtremeOverVoltagePeriod) AttrValue() zcl.Val  { return v.Value() }

func (RmsExtremeOverVoltagePeriod) Name() string { return `RMS Extreme Over Voltage Period` }
func (RmsExtremeOverVoltagePeriod) Description() string {
	return `is the durations used to measure an extreme over voltage condition.`
}

// RmsExtremeOverVoltagePeriod is the durations used to measure an extreme over voltage condition.
type RmsExtremeOverVoltagePeriod zcl.Zu16

func (v *RmsExtremeOverVoltagePeriod) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsExtremeOverVoltagePeriod) Value() zcl.Val     { return v }

func (v RmsExtremeOverVoltagePeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsExtremeOverVoltagePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsExtremeOverVoltagePeriod(*nt)
	return br, err
}

func (v RmsExtremeOverVoltagePeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsExtremeOverVoltagePeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsExtremeOverVoltagePeriod(*a)
	return nil
}

func (v *RmsExtremeOverVoltagePeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsExtremeOverVoltagePeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsExtremeOverVoltagePeriod) String() string {
	return zcl.Seconds.Format(float64(v))
}

const RmsExtremeOverVoltagePeriodPhBAttr zcl.AttrID = 2324

func (RmsExtremeOverVoltagePeriodPhB) ID() zcl.AttrID   { return RmsExtremeOverVoltagePeriodPhBAttr }
func (RmsExtremeOverVoltagePeriodPhB) Readable() bool   { return true }
func (RmsExtremeOverVoltagePeriodPhB) Writable() bool   { return true }
func (RmsExtremeOverVoltagePeriodPhB) Reportable() bool { return false }
func (RmsExtremeOverVoltagePeriodPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsExtremeOverVoltagePeriodPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsExtremeOverVoltagePeriodPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsExtremeOverVoltagePeriodPhB) AttrValue() zcl.Val  { return v.Value() }

func (RmsExtremeOverVoltagePeriodPhB) Name() string        { return `RMS Extreme Over Voltage Period Ph B` }
func (RmsExtremeOverVoltagePeriodPhB) Description() string { return `` }

type RmsExtremeOverVoltagePeriodPhB zcl.Zu16

func (v *RmsExtremeOverVoltagePeriodPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsExtremeOverVoltagePeriodPhB) Value() zcl.Val     { return v }

func (v RmsExtremeOverVoltagePeriodPhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsExtremeOverVoltagePeriodPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsExtremeOverVoltagePeriodPhB(*nt)
	return br, err
}

func (v RmsExtremeOverVoltagePeriodPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsExtremeOverVoltagePeriodPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsExtremeOverVoltagePeriodPhB(*a)
	return nil
}

func (v *RmsExtremeOverVoltagePeriodPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsExtremeOverVoltagePeriodPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsExtremeOverVoltagePeriodPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsExtremeOverVoltagePeriodPhCAttr zcl.AttrID = 2580

func (RmsExtremeOverVoltagePeriodPhC) ID() zcl.AttrID   { return RmsExtremeOverVoltagePeriodPhCAttr }
func (RmsExtremeOverVoltagePeriodPhC) Readable() bool   { return true }
func (RmsExtremeOverVoltagePeriodPhC) Writable() bool   { return true }
func (RmsExtremeOverVoltagePeriodPhC) Reportable() bool { return false }
func (RmsExtremeOverVoltagePeriodPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsExtremeOverVoltagePeriodPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsExtremeOverVoltagePeriodPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsExtremeOverVoltagePeriodPhC) AttrValue() zcl.Val  { return v.Value() }

func (RmsExtremeOverVoltagePeriodPhC) Name() string        { return `RMS Extreme Over-voltage Period Ph C` }
func (RmsExtremeOverVoltagePeriodPhC) Description() string { return `` }

type RmsExtremeOverVoltagePeriodPhC zcl.Zu16

func (v *RmsExtremeOverVoltagePeriodPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsExtremeOverVoltagePeriodPhC) Value() zcl.Val     { return v }

func (v RmsExtremeOverVoltagePeriodPhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsExtremeOverVoltagePeriodPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsExtremeOverVoltagePeriodPhC(*nt)
	return br, err
}

func (v RmsExtremeOverVoltagePeriodPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsExtremeOverVoltagePeriodPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsExtremeOverVoltagePeriodPhC(*a)
	return nil
}

func (v *RmsExtremeOverVoltagePeriodPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsExtremeOverVoltagePeriodPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsExtremeOverVoltagePeriodPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsExtremeUnderVoltagePeriodAttr zcl.AttrID = 1301

func (RmsExtremeUnderVoltagePeriod) ID() zcl.AttrID   { return RmsExtremeUnderVoltagePeriodAttr }
func (RmsExtremeUnderVoltagePeriod) Readable() bool   { return true }
func (RmsExtremeUnderVoltagePeriod) Writable() bool   { return true }
func (RmsExtremeUnderVoltagePeriod) Reportable() bool { return false }
func (RmsExtremeUnderVoltagePeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsExtremeUnderVoltagePeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsExtremeUnderVoltagePeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsExtremeUnderVoltagePeriod) AttrValue() zcl.Val  { return v.Value() }

func (RmsExtremeUnderVoltagePeriod) Name() string { return `RMS Extreme Under Voltage Period` }
func (RmsExtremeUnderVoltagePeriod) Description() string {
	return `is the duration used to measure an extreme under voltage condition.`
}

// RmsExtremeUnderVoltagePeriod is the duration used to measure an extreme under voltage condition.
type RmsExtremeUnderVoltagePeriod zcl.Zu16

func (v *RmsExtremeUnderVoltagePeriod) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsExtremeUnderVoltagePeriod) Value() zcl.Val     { return v }

func (v RmsExtremeUnderVoltagePeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsExtremeUnderVoltagePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsExtremeUnderVoltagePeriod(*nt)
	return br, err
}

func (v RmsExtremeUnderVoltagePeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsExtremeUnderVoltagePeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsExtremeUnderVoltagePeriod(*a)
	return nil
}

func (v *RmsExtremeUnderVoltagePeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsExtremeUnderVoltagePeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsExtremeUnderVoltagePeriod) String() string {
	return zcl.Seconds.Format(float64(v))
}

const RmsExtremeUnderVoltagePeriodPhBAttr zcl.AttrID = 2325

func (RmsExtremeUnderVoltagePeriodPhB) ID() zcl.AttrID   { return RmsExtremeUnderVoltagePeriodPhBAttr }
func (RmsExtremeUnderVoltagePeriodPhB) Readable() bool   { return true }
func (RmsExtremeUnderVoltagePeriodPhB) Writable() bool   { return true }
func (RmsExtremeUnderVoltagePeriodPhB) Reportable() bool { return false }
func (RmsExtremeUnderVoltagePeriodPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsExtremeUnderVoltagePeriodPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsExtremeUnderVoltagePeriodPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsExtremeUnderVoltagePeriodPhB) AttrValue() zcl.Val  { return v.Value() }

func (RmsExtremeUnderVoltagePeriodPhB) Name() string        { return `RMS Extreme Under Voltage Period Ph B` }
func (RmsExtremeUnderVoltagePeriodPhB) Description() string { return `` }

type RmsExtremeUnderVoltagePeriodPhB zcl.Zu16

func (v *RmsExtremeUnderVoltagePeriodPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsExtremeUnderVoltagePeriodPhB) Value() zcl.Val     { return v }

func (v RmsExtremeUnderVoltagePeriodPhB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *RmsExtremeUnderVoltagePeriodPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsExtremeUnderVoltagePeriodPhB(*nt)
	return br, err
}

func (v RmsExtremeUnderVoltagePeriodPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsExtremeUnderVoltagePeriodPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsExtremeUnderVoltagePeriodPhB(*a)
	return nil
}

func (v *RmsExtremeUnderVoltagePeriodPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsExtremeUnderVoltagePeriodPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsExtremeUnderVoltagePeriodPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsExtremeUnderVoltagePeriodPhCAttr zcl.AttrID = 2581

func (RmsExtremeUnderVoltagePeriodPhC) ID() zcl.AttrID   { return RmsExtremeUnderVoltagePeriodPhCAttr }
func (RmsExtremeUnderVoltagePeriodPhC) Readable() bool   { return true }
func (RmsExtremeUnderVoltagePeriodPhC) Writable() bool   { return true }
func (RmsExtremeUnderVoltagePeriodPhC) Reportable() bool { return false }
func (RmsExtremeUnderVoltagePeriodPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsExtremeUnderVoltagePeriodPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsExtremeUnderVoltagePeriodPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsExtremeUnderVoltagePeriodPhC) AttrValue() zcl.Val  { return v.Value() }

func (RmsExtremeUnderVoltagePeriodPhC) Name() string        { return `RMS Extreme Under-voltage Period Ph C` }
func (RmsExtremeUnderVoltagePeriodPhC) Description() string { return `` }

type RmsExtremeUnderVoltagePeriodPhC zcl.Zu16

func (v *RmsExtremeUnderVoltagePeriodPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsExtremeUnderVoltagePeriodPhC) Value() zcl.Val     { return v }

func (v RmsExtremeUnderVoltagePeriodPhC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *RmsExtremeUnderVoltagePeriodPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsExtremeUnderVoltagePeriodPhC(*nt)
	return br, err
}

func (v RmsExtremeUnderVoltagePeriodPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsExtremeUnderVoltagePeriodPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsExtremeUnderVoltagePeriodPhC(*a)
	return nil
}

func (v *RmsExtremeUnderVoltagePeriodPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsExtremeUnderVoltagePeriodPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsExtremeUnderVoltagePeriodPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsVoltageAttr zcl.AttrID = 1285

func (RmsVoltage) ID() zcl.AttrID   { return RmsVoltageAttr }
func (RmsVoltage) Readable() bool   { return false }
func (RmsVoltage) Writable() bool   { return false }
func (RmsVoltage) Reportable() bool { return true }
func (RmsVoltage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltage) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltage) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltage) Name() string { return `RMS Voltage` }
func (RmsVoltage) Description() string {
	return `represents the most recent RMS voltage reading.
If the RMS voltage cannot be measured, a value of 0xFFFF is returned`
}

// RmsVoltage represents the most recent RMS voltage reading.
// If the RMS voltage cannot be measured, a value of 0xFFFF is returned
type RmsVoltage zcl.Zu16

func (v *RmsVoltage) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltage) Value() zcl.Val     { return v }

func (v RmsVoltage) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltage(*nt)
	return br, err
}

func (v RmsVoltage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltage(*a)
	return nil
}

func (v *RmsVoltage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltage) String() string {
	return zcl.Volts.Format(float64(v))
}

const RmsVoltageMaxAttr zcl.AttrID = 1287

func (RmsVoltageMax) ID() zcl.AttrID   { return RmsVoltageMaxAttr }
func (RmsVoltageMax) Readable() bool   { return false }
func (RmsVoltageMax) Writable() bool   { return false }
func (RmsVoltageMax) Reportable() bool { return false }
func (RmsVoltageMax) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageMax) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageMax) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageMax) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageMax) Name() string { return `RMS Voltage Max` }
func (RmsVoltageMax) Description() string {
	return `Represents the highest RMS voltage value.
After resetting, this attribute will return a value of 0xFFFF until a measurement is made.`
}

// RmsVoltageMax Represents the highest RMS voltage value.
// After resetting, this attribute will return a value of 0xFFFF until a measurement is made.
type RmsVoltageMax zcl.Zu16

func (v *RmsVoltageMax) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageMax) Value() zcl.Val     { return v }

func (v RmsVoltageMax) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageMax(*nt)
	return br, err
}

func (v RmsVoltageMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageMax(*a)
	return nil
}

func (v *RmsVoltageMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageMax) String() string {
	return zcl.Volts.Format(float64(v))
}

const RmsVoltageMaxPhBAttr zcl.AttrID = 2311

func (RmsVoltageMaxPhB) ID() zcl.AttrID   { return RmsVoltageMaxPhBAttr }
func (RmsVoltageMaxPhB) Readable() bool   { return false }
func (RmsVoltageMaxPhB) Writable() bool   { return false }
func (RmsVoltageMaxPhB) Reportable() bool { return false }
func (RmsVoltageMaxPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageMaxPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageMaxPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageMaxPhB) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageMaxPhB) Name() string { return `RMS Voltage Max Ph B` }
func (RmsVoltageMaxPhB) Description() string {
	return `represents the highest RMS voltage value measured. After resetting, this attribute will return a
value of 0xFFFF until a measurement is made.`
}

// RmsVoltageMaxPhB represents the highest RMS voltage value measured. After resetting, this attribute will return a
// value of 0xFFFF until a measurement is made.
type RmsVoltageMaxPhB zcl.Zu16

func (v *RmsVoltageMaxPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageMaxPhB) Value() zcl.Val     { return v }

func (v RmsVoltageMaxPhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageMaxPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageMaxPhB(*nt)
	return br, err
}

func (v RmsVoltageMaxPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageMaxPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageMaxPhB(*a)
	return nil
}

func (v *RmsVoltageMaxPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageMaxPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageMaxPhB) String() string {
	return zcl.Volts.Format(float64(v))
}

const RmsVoltageMaxPhCAttr zcl.AttrID = 2567

func (RmsVoltageMaxPhC) ID() zcl.AttrID   { return RmsVoltageMaxPhCAttr }
func (RmsVoltageMaxPhC) Readable() bool   { return false }
func (RmsVoltageMaxPhC) Writable() bool   { return false }
func (RmsVoltageMaxPhC) Reportable() bool { return false }
func (RmsVoltageMaxPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageMaxPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageMaxPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageMaxPhC) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageMaxPhC) Name() string        { return `RMS Voltage Max Ph C` }
func (RmsVoltageMaxPhC) Description() string { return `` }

type RmsVoltageMaxPhC zcl.Zu16

func (v *RmsVoltageMaxPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageMaxPhC) Value() zcl.Val     { return v }

func (v RmsVoltageMaxPhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageMaxPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageMaxPhC(*nt)
	return br, err
}

func (v RmsVoltageMaxPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageMaxPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageMaxPhC(*a)
	return nil
}

func (v *RmsVoltageMaxPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageMaxPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageMaxPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsVoltageMinAttr zcl.AttrID = 1286

func (RmsVoltageMin) ID() zcl.AttrID   { return RmsVoltageMinAttr }
func (RmsVoltageMin) Readable() bool   { return false }
func (RmsVoltageMin) Writable() bool   { return false }
func (RmsVoltageMin) Reportable() bool { return false }
func (RmsVoltageMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageMin) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageMin) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageMin) Name() string { return `RMS Voltage Min` }
func (RmsVoltageMin) Description() string {
	return `Represents the lowest RMS voltage value.
After resetting, this attribute will return a value of 0xFFFF until a measurement is made.`
}

// RmsVoltageMin Represents the lowest RMS voltage value.
// After resetting, this attribute will return a value of 0xFFFF until a measurement is made.
type RmsVoltageMin zcl.Zu16

func (v *RmsVoltageMin) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageMin) Value() zcl.Val     { return v }

func (v RmsVoltageMin) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageMin(*nt)
	return br, err
}

func (v RmsVoltageMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageMin(*a)
	return nil
}

func (v *RmsVoltageMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageMin) String() string {
	return zcl.Volts.Format(float64(v))
}

const RmsVoltageMinPhBAttr zcl.AttrID = 2310

func (RmsVoltageMinPhB) ID() zcl.AttrID   { return RmsVoltageMinPhBAttr }
func (RmsVoltageMinPhB) Readable() bool   { return false }
func (RmsVoltageMinPhB) Writable() bool   { return false }
func (RmsVoltageMinPhB) Reportable() bool { return false }
func (RmsVoltageMinPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageMinPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageMinPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageMinPhB) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageMinPhB) Name() string { return `RMS Voltage Min Ph B` }
func (RmsVoltageMinPhB) Description() string {
	return `represents the lowest RMS voltage value measured. After resetting, this attribute will return a
value of 0xFFFF until a measurement is made.`
}

// RmsVoltageMinPhB represents the lowest RMS voltage value measured. After resetting, this attribute will return a
// value of 0xFFFF until a measurement is made.
type RmsVoltageMinPhB zcl.Zu16

func (v *RmsVoltageMinPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageMinPhB) Value() zcl.Val     { return v }

func (v RmsVoltageMinPhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageMinPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageMinPhB(*nt)
	return br, err
}

func (v RmsVoltageMinPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageMinPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageMinPhB(*a)
	return nil
}

func (v *RmsVoltageMinPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageMinPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageMinPhB) String() string {
	return zcl.Volts.Format(float64(v))
}

const RmsVoltageMinPhCAttr zcl.AttrID = 2566

func (RmsVoltageMinPhC) ID() zcl.AttrID   { return RmsVoltageMinPhCAttr }
func (RmsVoltageMinPhC) Readable() bool   { return false }
func (RmsVoltageMinPhC) Writable() bool   { return false }
func (RmsVoltageMinPhC) Reportable() bool { return false }
func (RmsVoltageMinPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageMinPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageMinPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageMinPhC) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageMinPhC) Name() string        { return `RMS Voltage Min Ph C` }
func (RmsVoltageMinPhC) Description() string { return `` }

type RmsVoltageMinPhC zcl.Zu16

func (v *RmsVoltageMinPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageMinPhC) Value() zcl.Val     { return v }

func (v RmsVoltageMinPhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageMinPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageMinPhC(*nt)
	return br, err
}

func (v RmsVoltageMinPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageMinPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageMinPhC(*a)
	return nil
}

func (v *RmsVoltageMinPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageMinPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageMinPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsVoltagePhBAttr zcl.AttrID = 2309

func (RmsVoltagePhB) ID() zcl.AttrID   { return RmsVoltagePhBAttr }
func (RmsVoltagePhB) Readable() bool   { return false }
func (RmsVoltagePhB) Writable() bool   { return false }
func (RmsVoltagePhB) Reportable() bool { return false }
func (RmsVoltagePhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltagePhB) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltagePhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltagePhB) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltagePhB) Name() string { return `RMS Voltage Ph B` }
func (RmsVoltagePhB) Description() string {
	return `represents the most recent RMS voltage reading. If the RMS voltage cannot be measured, a value of 0xFFFF
is returned.`
}

// RmsVoltagePhB represents the most recent RMS voltage reading. If the RMS voltage cannot be measured, a value of 0xFFFF
// is returned.
type RmsVoltagePhB zcl.Zu16

func (v *RmsVoltagePhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltagePhB) Value() zcl.Val     { return v }

func (v RmsVoltagePhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltagePhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltagePhB(*nt)
	return br, err
}

func (v RmsVoltagePhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltagePhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltagePhB(*a)
	return nil
}

func (v *RmsVoltagePhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltagePhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltagePhB) String() string {
	return zcl.Volts.Format(float64(v))
}

const RmsVoltagePhCAttr zcl.AttrID = 2565

func (RmsVoltagePhC) ID() zcl.AttrID   { return RmsVoltagePhCAttr }
func (RmsVoltagePhC) Readable() bool   { return false }
func (RmsVoltagePhC) Writable() bool   { return false }
func (RmsVoltagePhC) Reportable() bool { return false }
func (RmsVoltagePhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltagePhC) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltagePhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltagePhC) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltagePhC) Name() string        { return `RMS Voltage Ph C` }
func (RmsVoltagePhC) Description() string { return `` }

type RmsVoltagePhC zcl.Zu16

func (v *RmsVoltagePhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltagePhC) Value() zcl.Val     { return v }

func (v RmsVoltagePhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltagePhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltagePhC(*nt)
	return br, err
}

func (v RmsVoltagePhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltagePhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltagePhC(*a)
	return nil
}

func (v *RmsVoltagePhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltagePhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltagePhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsVoltageSagAttr zcl.AttrID = 2057

func (RmsVoltageSag) ID() zcl.AttrID   { return RmsVoltageSagAttr }
func (RmsVoltageSag) Readable() bool   { return true }
func (RmsVoltageSag) Writable() bool   { return true }
func (RmsVoltageSag) Reportable() bool { return false }
func (RmsVoltageSag) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageSag) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageSag) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageSag) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageSag) Name() string { return `RMS Voltage Sag` }
func (RmsVoltageSag) Description() string {
	return `is the RMS voltage below which a sag condition is reported.
The threshold shall be configurable within the specified operating range of the electricity meter.
The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively`
}

// RmsVoltageSag is the RMS voltage below which a sag condition is reported.
// The threshold shall be configurable within the specified operating range of the electricity meter.
// The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively
type RmsVoltageSag zcl.Zs16

func (v *RmsVoltageSag) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *RmsVoltageSag) Value() zcl.Val     { return v }

func (v RmsVoltageSag) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *RmsVoltageSag) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageSag(*nt)
	return br, err
}

func (v RmsVoltageSag) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *RmsVoltageSag) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageSag(*a)
	return nil
}

func (v *RmsVoltageSag) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = RmsVoltageSag(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageSag) String() string {
	return zcl.Volts.Format(float64(v))
}

const RmsVoltageSagPeriodAttr zcl.AttrID = 1302

func (RmsVoltageSagPeriod) ID() zcl.AttrID   { return RmsVoltageSagPeriodAttr }
func (RmsVoltageSagPeriod) Readable() bool   { return true }
func (RmsVoltageSagPeriod) Writable() bool   { return true }
func (RmsVoltageSagPeriod) Reportable() bool { return false }
func (RmsVoltageSagPeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageSagPeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageSagPeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageSagPeriod) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageSagPeriod) Name() string { return `RMS Voltage Sag Period` }
func (RmsVoltageSagPeriod) Description() string {
	return `is the duration used to measure a voltage sag condition.`
}

// RmsVoltageSagPeriod is the duration used to measure a voltage sag condition.
type RmsVoltageSagPeriod zcl.Zu16

func (v *RmsVoltageSagPeriod) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageSagPeriod) Value() zcl.Val     { return v }

func (v RmsVoltageSagPeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageSagPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageSagPeriod(*nt)
	return br, err
}

func (v RmsVoltageSagPeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageSagPeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageSagPeriod(*a)
	return nil
}

func (v *RmsVoltageSagPeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageSagPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageSagPeriod) String() string {
	return zcl.Seconds.Format(float64(v))
}

const RmsVoltageSagPeriodPhBAttr zcl.AttrID = 2326

func (RmsVoltageSagPeriodPhB) ID() zcl.AttrID   { return RmsVoltageSagPeriodPhBAttr }
func (RmsVoltageSagPeriodPhB) Readable() bool   { return true }
func (RmsVoltageSagPeriodPhB) Writable() bool   { return true }
func (RmsVoltageSagPeriodPhB) Reportable() bool { return false }
func (RmsVoltageSagPeriodPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageSagPeriodPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageSagPeriodPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageSagPeriodPhB) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageSagPeriodPhB) Name() string        { return `RMS Voltage Sag Period Ph B` }
func (RmsVoltageSagPeriodPhB) Description() string { return `` }

type RmsVoltageSagPeriodPhB zcl.Zu16

func (v *RmsVoltageSagPeriodPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageSagPeriodPhB) Value() zcl.Val     { return v }

func (v RmsVoltageSagPeriodPhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageSagPeriodPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageSagPeriodPhB(*nt)
	return br, err
}

func (v RmsVoltageSagPeriodPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageSagPeriodPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageSagPeriodPhB(*a)
	return nil
}

func (v *RmsVoltageSagPeriodPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageSagPeriodPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageSagPeriodPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsVoltageSagPeriodPhCAttr zcl.AttrID = 2582

func (RmsVoltageSagPeriodPhC) ID() zcl.AttrID   { return RmsVoltageSagPeriodPhCAttr }
func (RmsVoltageSagPeriodPhC) Readable() bool   { return true }
func (RmsVoltageSagPeriodPhC) Writable() bool   { return true }
func (RmsVoltageSagPeriodPhC) Reportable() bool { return false }
func (RmsVoltageSagPeriodPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageSagPeriodPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageSagPeriodPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageSagPeriodPhC) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageSagPeriodPhC) Name() string        { return `RMS Voltage Sag Period Ph C` }
func (RmsVoltageSagPeriodPhC) Description() string { return `` }

type RmsVoltageSagPeriodPhC zcl.Zu16

func (v *RmsVoltageSagPeriodPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageSagPeriodPhC) Value() zcl.Val     { return v }

func (v RmsVoltageSagPeriodPhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageSagPeriodPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageSagPeriodPhC(*nt)
	return br, err
}

func (v RmsVoltageSagPeriodPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageSagPeriodPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageSagPeriodPhC(*a)
	return nil
}

func (v *RmsVoltageSagPeriodPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageSagPeriodPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageSagPeriodPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsVoltageSwellAttr zcl.AttrID = 2058

func (RmsVoltageSwell) ID() zcl.AttrID   { return RmsVoltageSwellAttr }
func (RmsVoltageSwell) Readable() bool   { return true }
func (RmsVoltageSwell) Writable() bool   { return true }
func (RmsVoltageSwell) Reportable() bool { return false }
func (RmsVoltageSwell) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageSwell) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageSwell) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageSwell) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageSwell) Name() string { return `RMS Voltage Swell` }
func (RmsVoltageSwell) Description() string {
	return `is the RMS voltage above which a swell condition is reported.
The threshold shall be configurable within the specified operating range of the electricity meter.
The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively`
}

// RmsVoltageSwell is the RMS voltage above which a swell condition is reported.
// The threshold shall be configurable within the specified operating range of the electricity meter.
// The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively
type RmsVoltageSwell zcl.Zs16

func (v *RmsVoltageSwell) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *RmsVoltageSwell) Value() zcl.Val     { return v }

func (v RmsVoltageSwell) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *RmsVoltageSwell) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageSwell(*nt)
	return br, err
}

func (v RmsVoltageSwell) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *RmsVoltageSwell) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageSwell(*a)
	return nil
}

func (v *RmsVoltageSwell) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = RmsVoltageSwell(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageSwell) String() string {
	return zcl.Volts.Format(float64(v))
}

const RmsVoltageSwellPeriodAttr zcl.AttrID = 1303

func (RmsVoltageSwellPeriod) ID() zcl.AttrID   { return RmsVoltageSwellPeriodAttr }
func (RmsVoltageSwellPeriod) Readable() bool   { return true }
func (RmsVoltageSwellPeriod) Writable() bool   { return true }
func (RmsVoltageSwellPeriod) Reportable() bool { return false }
func (RmsVoltageSwellPeriod) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageSwellPeriod) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageSwellPeriod) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageSwellPeriod) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageSwellPeriod) Name() string { return `RMS Voltage Swell Period` }
func (RmsVoltageSwellPeriod) Description() string {
	return `is the duration used to measure a voltage swell condition.`
}

// RmsVoltageSwellPeriod is the duration used to measure a voltage swell condition.
type RmsVoltageSwellPeriod zcl.Zu16

func (v *RmsVoltageSwellPeriod) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageSwellPeriod) Value() zcl.Val     { return v }

func (v RmsVoltageSwellPeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageSwellPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageSwellPeriod(*nt)
	return br, err
}

func (v RmsVoltageSwellPeriod) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageSwellPeriod) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageSwellPeriod(*a)
	return nil
}

func (v *RmsVoltageSwellPeriod) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageSwellPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageSwellPeriod) String() string {
	return zcl.Seconds.Format(float64(v))
}

const RmsVoltageSwellPeriodPhBAttr zcl.AttrID = 2327

func (RmsVoltageSwellPeriodPhB) ID() zcl.AttrID   { return RmsVoltageSwellPeriodPhBAttr }
func (RmsVoltageSwellPeriodPhB) Readable() bool   { return true }
func (RmsVoltageSwellPeriodPhB) Writable() bool   { return true }
func (RmsVoltageSwellPeriodPhB) Reportable() bool { return false }
func (RmsVoltageSwellPeriodPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageSwellPeriodPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageSwellPeriodPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageSwellPeriodPhB) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageSwellPeriodPhB) Name() string        { return `RMS Voltage Swell Period Ph B` }
func (RmsVoltageSwellPeriodPhB) Description() string { return `` }

type RmsVoltageSwellPeriodPhB zcl.Zu16

func (v *RmsVoltageSwellPeriodPhB) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageSwellPeriodPhB) Value() zcl.Val     { return v }

func (v RmsVoltageSwellPeriodPhB) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageSwellPeriodPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageSwellPeriodPhB(*nt)
	return br, err
}

func (v RmsVoltageSwellPeriodPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageSwellPeriodPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageSwellPeriodPhB(*a)
	return nil
}

func (v *RmsVoltageSwellPeriodPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageSwellPeriodPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageSwellPeriodPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsVoltageSwellPeriodPhCAttr zcl.AttrID = 2583

func (RmsVoltageSwellPeriodPhC) ID() zcl.AttrID   { return RmsVoltageSwellPeriodPhCAttr }
func (RmsVoltageSwellPeriodPhC) Readable() bool   { return true }
func (RmsVoltageSwellPeriodPhC) Writable() bool   { return true }
func (RmsVoltageSwellPeriodPhC) Reportable() bool { return false }
func (RmsVoltageSwellPeriodPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsVoltageSwellPeriodPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsVoltageSwellPeriodPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsVoltageSwellPeriodPhC) AttrValue() zcl.Val  { return v.Value() }

func (RmsVoltageSwellPeriodPhC) Name() string        { return `RMS Voltage Swell Period Ph C` }
func (RmsVoltageSwellPeriodPhC) Description() string { return `` }

type RmsVoltageSwellPeriodPhC zcl.Zu16

func (v *RmsVoltageSwellPeriodPhC) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RmsVoltageSwellPeriodPhC) Value() zcl.Val     { return v }

func (v RmsVoltageSwellPeriodPhC) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RmsVoltageSwellPeriodPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsVoltageSwellPeriodPhC(*nt)
	return br, err
}

func (v RmsVoltageSwellPeriodPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RmsVoltageSwellPeriodPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsVoltageSwellPeriodPhC(*a)
	return nil
}

func (v *RmsVoltageSwellPeriodPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RmsVoltageSwellPeriodPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsVoltageSwellPeriodPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RmsExtremeOverVoltageAttr zcl.AttrID = 2055

func (RmsExtremeOverVoltage) ID() zcl.AttrID   { return RmsExtremeOverVoltageAttr }
func (RmsExtremeOverVoltage) Readable() bool   { return true }
func (RmsExtremeOverVoltage) Writable() bool   { return true }
func (RmsExtremeOverVoltage) Reportable() bool { return false }
func (RmsExtremeOverVoltage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsExtremeOverVoltage) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsExtremeOverVoltage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsExtremeOverVoltage) AttrValue() zcl.Val  { return v.Value() }

func (RmsExtremeOverVoltage) Name() string { return `RMS extreme over-voltage` }
func (RmsExtremeOverVoltage) Description() string {
	return `is the RMS voltage above which an extreme under voltage condition is reported.
The threshold shall be configurable within the specified operating range of the electricity meter.
The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively.`
}

// RmsExtremeOverVoltage is the RMS voltage above which an extreme under voltage condition is reported.
// The threshold shall be configurable within the specified operating range of the electricity meter.
// The value is multiplied and divided by the ACVoltageMultiplier and ACVoltageDivisor, respectively.
type RmsExtremeOverVoltage zcl.Zs16

func (v *RmsExtremeOverVoltage) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *RmsExtremeOverVoltage) Value() zcl.Val     { return v }

func (v RmsExtremeOverVoltage) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *RmsExtremeOverVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsExtremeOverVoltage(*nt)
	return br, err
}

func (v RmsExtremeOverVoltage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *RmsExtremeOverVoltage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsExtremeOverVoltage(*a)
	return nil
}

func (v *RmsExtremeOverVoltage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = RmsExtremeOverVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsExtremeOverVoltage) String() string {
	return zcl.Volts.Format(float64(v))
}

const RmsExtremeUnderVoltageAttr zcl.AttrID = 2056

func (RmsExtremeUnderVoltage) ID() zcl.AttrID   { return RmsExtremeUnderVoltageAttr }
func (RmsExtremeUnderVoltage) Readable() bool   { return true }
func (RmsExtremeUnderVoltage) Writable() bool   { return true }
func (RmsExtremeUnderVoltage) Reportable() bool { return false }
func (RmsExtremeUnderVoltage) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RmsExtremeUnderVoltage) AttrID() zcl.AttrID   { return v.ID() }
func (v RmsExtremeUnderVoltage) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RmsExtremeUnderVoltage) AttrValue() zcl.Val  { return v.Value() }

func (RmsExtremeUnderVoltage) Name() string { return `RMS extreme under-voltage` }
func (RmsExtremeUnderVoltage) Description() string {
	return `is the RMS voltage below which an extreme under voltage condition is reported. The threshold shall be
configurable within the specified operating range of the electricity meter. The value is multiplied and divided
by the ACVoltageMultiplier and ACVoltageDivisor, respectively.`
}

// RmsExtremeUnderVoltage is the RMS voltage below which an extreme under voltage condition is reported. The threshold shall be
// configurable within the specified operating range of the electricity meter. The value is multiplied and divided
// by the ACVoltageMultiplier and ACVoltageDivisor, respectively.
type RmsExtremeUnderVoltage zcl.Zs16

func (v *RmsExtremeUnderVoltage) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *RmsExtremeUnderVoltage) Value() zcl.Val     { return v }

func (v RmsExtremeUnderVoltage) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *RmsExtremeUnderVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = RmsExtremeUnderVoltage(*nt)
	return br, err
}

func (v RmsExtremeUnderVoltage) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *RmsExtremeUnderVoltage) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RmsExtremeUnderVoltage(*a)
	return nil
}

func (v *RmsExtremeUnderVoltage) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = RmsExtremeUnderVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RmsExtremeUnderVoltage) String() string {
	return zcl.Volts.Format(float64(v))
}

const ReactiveCurrentAttr zcl.AttrID = 1283

func (ReactiveCurrent) ID() zcl.AttrID   { return ReactiveCurrentAttr }
func (ReactiveCurrent) Readable() bool   { return false }
func (ReactiveCurrent) Writable() bool   { return false }
func (ReactiveCurrent) Reportable() bool { return false }
func (ReactiveCurrent) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ReactiveCurrent) AttrID() zcl.AttrID   { return v.ID() }
func (v ReactiveCurrent) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ReactiveCurrent) AttrValue() zcl.Val  { return v.Value() }

func (ReactiveCurrent) Name() string { return `Reactive Current` }
func (ReactiveCurrent) Description() string {
	return `Represents the single phase or Phase A, AC reactive current value at the moment in time the attribute is read,
in Amps (A). Positive values indicate power delivered to the premises where negative values indicate power
received from the premises.`
}

// ReactiveCurrent Represents the single phase or Phase A, AC reactive current value at the moment in time the attribute is read,
// in Amps (A). Positive values indicate power delivered to the premises where negative values indicate power
// received from the premises.
type ReactiveCurrent zcl.Zs16

func (v *ReactiveCurrent) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ReactiveCurrent) Value() zcl.Val     { return v }

func (v ReactiveCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ReactiveCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ReactiveCurrent(*nt)
	return br, err
}

func (v ReactiveCurrent) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ReactiveCurrent) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ReactiveCurrent(*a)
	return nil
}

func (v *ReactiveCurrent) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ReactiveCurrent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ReactiveCurrent) String() string {
	return zcl.Amperes.Format(float64(v))
}

const ReactiveCurrentPhBAttr zcl.AttrID = 2307

func (ReactiveCurrentPhB) ID() zcl.AttrID   { return ReactiveCurrentPhBAttr }
func (ReactiveCurrentPhB) Readable() bool   { return false }
func (ReactiveCurrentPhB) Writable() bool   { return false }
func (ReactiveCurrentPhB) Reportable() bool { return false }
func (ReactiveCurrentPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ReactiveCurrentPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v ReactiveCurrentPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ReactiveCurrentPhB) AttrValue() zcl.Val  { return v.Value() }

func (ReactiveCurrentPhB) Name() string { return `Reactive Current Ph B` }
func (ReactiveCurrentPhB) Description() string {
	return `represents the Phase B, AC reactive current value at the moment in time the attribute is read.
Positive values indicate power delivered to the premises where negative values indicate power received from
the premises.`
}

// ReactiveCurrentPhB represents the Phase B, AC reactive current value at the moment in time the attribute is read.
// Positive values indicate power delivered to the premises where negative values indicate power received from
// the premises.
type ReactiveCurrentPhB zcl.Zs16

func (v *ReactiveCurrentPhB) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ReactiveCurrentPhB) Value() zcl.Val     { return v }

func (v ReactiveCurrentPhB) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ReactiveCurrentPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ReactiveCurrentPhB(*nt)
	return br, err
}

func (v ReactiveCurrentPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ReactiveCurrentPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ReactiveCurrentPhB(*a)
	return nil
}

func (v *ReactiveCurrentPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ReactiveCurrentPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ReactiveCurrentPhB) String() string {
	return zcl.Amperes.Format(float64(v))
}

const ReactiveCurrentPhCAttr zcl.AttrID = 2563

func (ReactiveCurrentPhC) ID() zcl.AttrID   { return ReactiveCurrentPhCAttr }
func (ReactiveCurrentPhC) Readable() bool   { return false }
func (ReactiveCurrentPhC) Writable() bool   { return false }
func (ReactiveCurrentPhC) Reportable() bool { return false }
func (ReactiveCurrentPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ReactiveCurrentPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v ReactiveCurrentPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ReactiveCurrentPhC) AttrValue() zcl.Val  { return v.Value() }

func (ReactiveCurrentPhC) Name() string        { return `Reactive Current Ph C` }
func (ReactiveCurrentPhC) Description() string { return `` }

type ReactiveCurrentPhC zcl.Zs16

func (v *ReactiveCurrentPhC) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ReactiveCurrentPhC) Value() zcl.Val     { return v }

func (v ReactiveCurrentPhC) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ReactiveCurrentPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ReactiveCurrentPhC(*nt)
	return br, err
}

func (v ReactiveCurrentPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ReactiveCurrentPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ReactiveCurrentPhC(*a)
	return nil
}

func (v *ReactiveCurrentPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ReactiveCurrentPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ReactiveCurrentPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const ReactivePowerAttr zcl.AttrID = 1294

func (ReactivePower) ID() zcl.AttrID   { return ReactivePowerAttr }
func (ReactivePower) Readable() bool   { return false }
func (ReactivePower) Writable() bool   { return false }
func (ReactivePower) Reportable() bool { return false }
func (ReactivePower) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ReactivePower) AttrID() zcl.AttrID   { return v.ID() }
func (v ReactivePower) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ReactivePower) AttrValue() zcl.Val  { return v.Value() }

func (ReactivePower) Name() string { return `Reactive Power` }
func (ReactivePower) Description() string {
	return `Represents the single phase or Phase A, current demand of reactive power delivered or received at the
premises, in VAr. Positive values indicate power delivered to the premises where negative values indicate
power received from the premises.`
}

// ReactivePower Represents the single phase or Phase A, current demand of reactive power delivered or received at the
// premises, in VAr. Positive values indicate power delivered to the premises where negative values indicate
// power received from the premises.
type ReactivePower zcl.Zs16

func (v *ReactivePower) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ReactivePower) Value() zcl.Val     { return v }

func (v ReactivePower) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ReactivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ReactivePower(*nt)
	return br, err
}

func (v ReactivePower) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ReactivePower) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ReactivePower(*a)
	return nil
}

func (v *ReactivePower) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ReactivePower(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ReactivePower) String() string {
	return zcl.VoltAmperesReactive.Format(float64(v))
}

const ReactivePowerPhBAttr zcl.AttrID = 2318

func (ReactivePowerPhB) ID() zcl.AttrID   { return ReactivePowerPhBAttr }
func (ReactivePowerPhB) Readable() bool   { return false }
func (ReactivePowerPhB) Writable() bool   { return false }
func (ReactivePowerPhB) Reportable() bool { return false }
func (ReactivePowerPhB) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ReactivePowerPhB) AttrID() zcl.AttrID   { return v.ID() }
func (v ReactivePowerPhB) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ReactivePowerPhB) AttrValue() zcl.Val  { return v.Value() }

func (ReactivePowerPhB) Name() string        { return `Reactive Power Ph B` }
func (ReactivePowerPhB) Description() string { return `` }

type ReactivePowerPhB zcl.Zs16

func (v *ReactivePowerPhB) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ReactivePowerPhB) Value() zcl.Val     { return v }

func (v ReactivePowerPhB) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ReactivePowerPhB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ReactivePowerPhB(*nt)
	return br, err
}

func (v ReactivePowerPhB) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ReactivePowerPhB) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ReactivePowerPhB(*a)
	return nil
}

func (v *ReactivePowerPhB) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ReactivePowerPhB(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ReactivePowerPhB) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const ReactivePowerPhCAttr zcl.AttrID = 2574

func (ReactivePowerPhC) ID() zcl.AttrID   { return ReactivePowerPhCAttr }
func (ReactivePowerPhC) Readable() bool   { return false }
func (ReactivePowerPhC) Writable() bool   { return false }
func (ReactivePowerPhC) Reportable() bool { return false }
func (ReactivePowerPhC) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ReactivePowerPhC) AttrID() zcl.AttrID   { return v.ID() }
func (v ReactivePowerPhC) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ReactivePowerPhC) AttrValue() zcl.Val  { return v.Value() }

func (ReactivePowerPhC) Name() string        { return `Reactive Power Ph C` }
func (ReactivePowerPhC) Description() string { return `` }

type ReactivePowerPhC zcl.Zs16

func (v *ReactivePowerPhC) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ReactivePowerPhC) Value() zcl.Val     { return v }

func (v ReactivePowerPhC) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ReactivePowerPhC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ReactivePowerPhC(*nt)
	return br, err
}

func (v ReactivePowerPhC) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ReactivePowerPhC) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ReactivePowerPhC(*a)
	return nil
}

func (v *ReactivePowerPhC) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ReactivePowerPhC(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ReactivePowerPhC) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const ScaleAttr zcl.AttrID = 20

func (Scale) ID() zcl.AttrID   { return ScaleAttr }
func (Scale) Readable() bool   { return true }
func (Scale) Writable() bool   { return false }
func (Scale) Reportable() bool { return false }
func (Scale) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Scale) AttrID() zcl.AttrID   { return v.ID() }
func (v Scale) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Scale) AttrValue() zcl.Val  { return v.Value() }

func (Scale) Name() string { return `Scale` }
func (Scale) Description() string {
	return `indicates the base 10 exponent used to obtain ScaledPressure`
}

// Scale indicates the base 10 exponent used to obtain ScaledPressure
type Scale zcl.Zs8

func (v *Scale) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *Scale) Value() zcl.Val     { return v }

func (v Scale) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *Scale) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = Scale(*nt)
	return br, err
}

func (v Scale) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *Scale) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Scale(*a)
	return nil
}

func (v *Scale) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = Scale(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Scale) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const ScaledMaxPressureAttr zcl.AttrID = 18

func (ScaledMaxPressure) ID() zcl.AttrID   { return ScaledMaxPressureAttr }
func (ScaledMaxPressure) Readable() bool   { return true }
func (ScaledMaxPressure) Writable() bool   { return false }
func (ScaledMaxPressure) Reportable() bool { return false }
func (ScaledMaxPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ScaledMaxPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v ScaledMaxPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ScaledMaxPressure) AttrValue() zcl.Val  { return v.Value() }

func (ScaledMaxPressure) Name() string { return `Scaled Max Pressure` }
func (ScaledMaxPressure) Description() string {
	return `indicates the maximum value of ScaledPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// ScaledMaxPressure indicates the maximum value of ScaledPressure that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type ScaledMaxPressure zcl.Zs16

func (v *ScaledMaxPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ScaledMaxPressure) Value() zcl.Val     { return v }

func (v ScaledMaxPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ScaledMaxPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ScaledMaxPressure(*nt)
	return br, err
}

func (v ScaledMaxPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ScaledMaxPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ScaledMaxPressure(*a)
	return nil
}

func (v *ScaledMaxPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ScaledMaxPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ScaledMaxPressure) String() string {
	return zcl.Pascals.Format(float64(v))
}

const ScaledMinPressureAttr zcl.AttrID = 17

func (ScaledMinPressure) ID() zcl.AttrID   { return ScaledMinPressureAttr }
func (ScaledMinPressure) Readable() bool   { return true }
func (ScaledMinPressure) Writable() bool   { return false }
func (ScaledMinPressure) Reportable() bool { return false }
func (ScaledMinPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ScaledMinPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v ScaledMinPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ScaledMinPressure) AttrValue() zcl.Val  { return v.Value() }

func (ScaledMinPressure) Name() string { return `Scaled Min Pressure` }
func (ScaledMinPressure) Description() string {
	return `indicates the minimum value of ScaledPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// ScaledMinPressure indicates the minimum value of ScaledPressure that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type ScaledMinPressure zcl.Zs16

func (v *ScaledMinPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ScaledMinPressure) Value() zcl.Val     { return v }

func (v ScaledMinPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ScaledMinPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ScaledMinPressure(*nt)
	return br, err
}

func (v ScaledMinPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ScaledMinPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ScaledMinPressure(*a)
	return nil
}

func (v *ScaledMinPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ScaledMinPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ScaledMinPressure) String() string {
	return zcl.Pascals.Format(float64(v))
}

const ScaledPressureAttr zcl.AttrID = 16

func (ScaledPressure) ID() zcl.AttrID   { return ScaledPressureAttr }
func (ScaledPressure) Readable() bool   { return true }
func (ScaledPressure) Writable() bool   { return false }
func (ScaledPressure) Reportable() bool { return true }
func (ScaledPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ScaledPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v ScaledPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ScaledPressure) AttrValue() zcl.Val  { return v.Value() }

func (ScaledPressure) Name() string { return `Scaled Pressure` }
func (ScaledPressure) Description() string {
	return `represents the pressure in Pascals as follows:
ScaledPressure = 10^Scale x Pressure in Pa`
}

// ScaledPressure represents the pressure in Pascals as follows:
// ScaledPressure = 10^Scale x Pressure in Pa
type ScaledPressure zcl.Zs16

func (v *ScaledPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ScaledPressure) Value() zcl.Val     { return v }

func (v ScaledPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ScaledPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ScaledPressure(*nt)
	return br, err
}

func (v ScaledPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ScaledPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ScaledPressure(*a)
	return nil
}

func (v *ScaledPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ScaledPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ScaledPressure) String() string {
	return zcl.Pascals.Format(float64(v))
}

const ScaledToleranceAttr zcl.AttrID = 19

func (ScaledTolerance) ID() zcl.AttrID   { return ScaledToleranceAttr }
func (ScaledTolerance) Readable() bool   { return true }
func (ScaledTolerance) Writable() bool   { return false }
func (ScaledTolerance) Reportable() bool { return true }
func (ScaledTolerance) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ScaledTolerance) AttrID() zcl.AttrID   { return v.ID() }
func (v ScaledTolerance) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ScaledTolerance) AttrValue() zcl.Val  { return v.Value() }

func (ScaledTolerance) Name() string { return `Scaled Tolerance` }
func (ScaledTolerance) Description() string {
	return `indicates the magnitude of the possible error that is associated with ScaledPressure.
The true value is located in the range (ScaledPressure – ScaledTolerance) to
(ScaledPressure + ScaledTolerance)`
}

// ScaledTolerance indicates the magnitude of the possible error that is associated with ScaledPressure.
// The true value is located in the range (ScaledPressure – ScaledTolerance) to
// (ScaledPressure + ScaledTolerance)
type ScaledTolerance zcl.Zu16

func (v *ScaledTolerance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ScaledTolerance) Value() zcl.Val     { return v }

func (v ScaledTolerance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ScaledTolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ScaledTolerance(*nt)
	return br, err
}

func (v ScaledTolerance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ScaledTolerance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ScaledTolerance(*a)
	return nil
}

func (v *ScaledTolerance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ScaledTolerance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ScaledTolerance) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (StartTime) Name() string { return `Start Time` }
func (StartTime) Description() string {
	return `is the end time of the most chronologically recent interval being requested.
Example: Data collected from 2:00 PM to 3:00 PM would be specified as a 3:00 PM interval (end time).`
}

// StartTime is the end time of the most chronologically recent interval being requested.
// Example: Data collected from 2:00 PM to 3:00 PM would be specified as a 3:00 PM interval (end time).
type StartTime zcl.Zutc

func (v *StartTime) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *StartTime) Value() zcl.Val     { return v }

func (v StartTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *StartTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = StartTime(*nt)
	return br, err
}

func (v StartTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *StartTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StartTime(*a)
	return nil
}

func (v *StartTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = StartTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StartTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

const ToleranceAttr zcl.AttrID = 3

func (Tolerance) ID() zcl.AttrID   { return ToleranceAttr }
func (Tolerance) Readable() bool   { return true }
func (Tolerance) Writable() bool   { return false }
func (Tolerance) Reportable() bool { return true }
func (Tolerance) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Tolerance) AttrID() zcl.AttrID   { return v.ID() }
func (v Tolerance) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Tolerance) AttrValue() zcl.Val  { return v.Value() }

func (Tolerance) Name() string { return `Tolerance` }
func (Tolerance) Description() string {
	return `indicates the magnitude of the possible error that is associated with MeasuredValue.
The true value is located in the range (MeasuredValue – Tolerance) to
(MeasuredValue + Tolerance)`
}

// Tolerance indicates the magnitude of the possible error that is associated with MeasuredValue.
// The true value is located in the range (MeasuredValue – Tolerance) to
// (MeasuredValue + Tolerance)
type Tolerance zcl.Zu16

func (v *Tolerance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Tolerance) Value() zcl.Val     { return v }

func (v Tolerance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Tolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Tolerance(*nt)
	return br, err
}

func (v Tolerance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Tolerance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Tolerance(*a)
	return nil
}

func (v *Tolerance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Tolerance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Tolerance) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const TotalActivePowerAttr zcl.AttrID = 772

func (TotalActivePower) ID() zcl.AttrID   { return TotalActivePowerAttr }
func (TotalActivePower) Readable() bool   { return false }
func (TotalActivePower) Writable() bool   { return false }
func (TotalActivePower) Reportable() bool { return false }
func (TotalActivePower) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TotalActivePower) AttrID() zcl.AttrID   { return v.ID() }
func (v TotalActivePower) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TotalActivePower) AttrValue() zcl.Val  { return v.Value() }

func (TotalActivePower) Name() string { return `Total Active Power` }
func (TotalActivePower) Description() string {
	return `represents the current demand of active power delivered or received at the premises, in kW.
Positive values indicate power delivered to the premises where negative values indicate power received from
the premises. In case if device is capable of measuring multi elements or phases then this will be net active
power value.`
}

// TotalActivePower represents the current demand of active power delivered or received at the premises, in kW.
// Positive values indicate power delivered to the premises where negative values indicate power received from
// the premises. In case if device is capable of measuring multi elements or phases then this will be net active
// power value.
type TotalActivePower zcl.Zs32

func (v *TotalActivePower) TypeID() zcl.TypeID { return new(zcl.Zs32).TypeID() }
func (v *TotalActivePower) Value() zcl.Val     { return v }

func (v TotalActivePower) MarshalZcl() ([]byte, error) { return zcl.Zs32(v).MarshalZcl() }

func (v *TotalActivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*v = TotalActivePower(*nt)
	return br, err
}

func (v TotalActivePower) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs32(v))
}

func (v *TotalActivePower) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TotalActivePower(*a)
	return nil
}

func (v *TotalActivePower) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs32); ok {
		*v = TotalActivePower(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TotalActivePower) String() string {
	return zcl.Kilowatts.Format(float64(v))
}

const TotalApparentPowerAttr zcl.AttrID = 774

func (TotalApparentPower) ID() zcl.AttrID   { return TotalApparentPowerAttr }
func (TotalApparentPower) Readable() bool   { return false }
func (TotalApparentPower) Writable() bool   { return false }
func (TotalApparentPower) Reportable() bool { return false }
func (TotalApparentPower) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TotalApparentPower) AttrID() zcl.AttrID   { return v.ID() }
func (v TotalApparentPower) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TotalApparentPower) AttrValue() zcl.Val  { return v.Value() }

func (TotalApparentPower) Name() string { return `Total Apparent Power` }
func (TotalApparentPower) Description() string {
	return `represents the current demand of apparent power.`
}

// TotalApparentPower represents the current demand of apparent power.
type TotalApparentPower zcl.Zu32

func (v *TotalApparentPower) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *TotalApparentPower) Value() zcl.Val     { return v }

func (v TotalApparentPower) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *TotalApparentPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = TotalApparentPower(*nt)
	return br, err
}

func (v TotalApparentPower) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *TotalApparentPower) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TotalApparentPower(*a)
	return nil
}

func (v *TotalApparentPower) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = TotalApparentPower(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TotalApparentPower) String() string {
	return zcl.KiloVoltAmperes.Format(float64(v))
}

const TotalReactivePowerAttr zcl.AttrID = 773

func (TotalReactivePower) ID() zcl.AttrID   { return TotalReactivePowerAttr }
func (TotalReactivePower) Readable() bool   { return false }
func (TotalReactivePower) Writable() bool   { return false }
func (TotalReactivePower) Reportable() bool { return false }
func (TotalReactivePower) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TotalReactivePower) AttrID() zcl.AttrID   { return v.ID() }
func (v TotalReactivePower) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TotalReactivePower) AttrValue() zcl.Val  { return v.Value() }

func (TotalReactivePower) Name() string { return `Total Reactive Power` }
func (TotalReactivePower) Description() string {
	return `represents the current demand of reactive power delivered or received at the premises, in kVAr. Positive values
indicate power delivered to the premises where negative values indicate power received from the premises.
In case if device is capable of measuring multi elements or phases then this will be net reactive
power value.`
}

// TotalReactivePower represents the current demand of reactive power delivered or received at the premises, in kVAr. Positive values
// indicate power delivered to the premises where negative values indicate power received from the premises.
// In case if device is capable of measuring multi elements or phases then this will be net reactive
// power value.
type TotalReactivePower zcl.Zs32

func (v *TotalReactivePower) TypeID() zcl.TypeID { return new(zcl.Zs32).TypeID() }
func (v *TotalReactivePower) Value() zcl.Val     { return v }

func (v TotalReactivePower) MarshalZcl() ([]byte, error) { return zcl.Zs32(v).MarshalZcl() }

func (v *TotalReactivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*v = TotalReactivePower(*nt)
	return br, err
}

func (v TotalReactivePower) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs32(v))
}

func (v *TotalReactivePower) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TotalReactivePower(*a)
	return nil
}

func (v *TotalReactivePower) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs32); ok {
		*v = TotalReactivePower(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TotalReactivePower) String() string {
	return zcl.KiloVoltAmperesReactive.Format(float64(v))
}

const UltrasonicOccupiedToUnoccupiedDelayAttr zcl.AttrID = 32

func (UltrasonicOccupiedToUnoccupiedDelay) ID() zcl.AttrID {
	return UltrasonicOccupiedToUnoccupiedDelayAttr
}
func (UltrasonicOccupiedToUnoccupiedDelay) Readable() bool   { return true }
func (UltrasonicOccupiedToUnoccupiedDelay) Writable() bool   { return true }
func (UltrasonicOccupiedToUnoccupiedDelay) Reportable() bool { return false }
func (UltrasonicOccupiedToUnoccupiedDelay) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UltrasonicOccupiedToUnoccupiedDelay) AttrID() zcl.AttrID   { return v.ID() }
func (v UltrasonicOccupiedToUnoccupiedDelay) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UltrasonicOccupiedToUnoccupiedDelay) AttrValue() zcl.Val  { return v.Value() }

func (UltrasonicOccupiedToUnoccupiedDelay) Name() string {
	return `Ultrasonic Occupied To Unoccupied Delay`
}
func (UltrasonicOccupiedToUnoccupiedDelay) Description() string {
	return `is 16 bits in length and specifies the time delay, in seconds, before the Ultrasonic sensor
changes to its unoccupied state after the last detection of movement in the sensed area`
}

// UltrasonicOccupiedToUnoccupiedDelay is 16 bits in length and specifies the time delay, in seconds, before the Ultrasonic sensor
// changes to its unoccupied state after the last detection of movement in the sensed area
type UltrasonicOccupiedToUnoccupiedDelay zcl.Zu16

func (v *UltrasonicOccupiedToUnoccupiedDelay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *UltrasonicOccupiedToUnoccupiedDelay) Value() zcl.Val     { return v }

func (v UltrasonicOccupiedToUnoccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *UltrasonicOccupiedToUnoccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = UltrasonicOccupiedToUnoccupiedDelay(*nt)
	return br, err
}

func (v UltrasonicOccupiedToUnoccupiedDelay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *UltrasonicOccupiedToUnoccupiedDelay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UltrasonicOccupiedToUnoccupiedDelay(*a)
	return nil
}

func (v *UltrasonicOccupiedToUnoccupiedDelay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = UltrasonicOccupiedToUnoccupiedDelay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UltrasonicOccupiedToUnoccupiedDelay) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const UltrasonicUnoccupiedToOccupiedDelayAttr zcl.AttrID = 33

func (UltrasonicUnoccupiedToOccupiedDelay) ID() zcl.AttrID {
	return UltrasonicUnoccupiedToOccupiedDelayAttr
}
func (UltrasonicUnoccupiedToOccupiedDelay) Readable() bool   { return true }
func (UltrasonicUnoccupiedToOccupiedDelay) Writable() bool   { return true }
func (UltrasonicUnoccupiedToOccupiedDelay) Reportable() bool { return false }
func (UltrasonicUnoccupiedToOccupiedDelay) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UltrasonicUnoccupiedToOccupiedDelay) AttrID() zcl.AttrID   { return v.ID() }
func (v UltrasonicUnoccupiedToOccupiedDelay) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UltrasonicUnoccupiedToOccupiedDelay) AttrValue() zcl.Val  { return v.Value() }

func (UltrasonicUnoccupiedToOccupiedDelay) Name() string {
	return `Ultrasonic Unoccupied To Occupied Delay`
}
func (UltrasonicUnoccupiedToOccupiedDelay) Description() string {
	return `is 16 bits in length and specifies the time delay, in seconds, before the Ultrasonic sensor changes
to its occupied state after the detection of movement in the sensed area. This attribute is
mandatory if the UltrasonicUnoccupiedToOccupiedThreshold attribute is implemented`
}

// UltrasonicUnoccupiedToOccupiedDelay is 16 bits in length and specifies the time delay, in seconds, before the Ultrasonic sensor changes
// to its occupied state after the detection of movement in the sensed area. This attribute is
// mandatory if the UltrasonicUnoccupiedToOccupiedThreshold attribute is implemented
type UltrasonicUnoccupiedToOccupiedDelay zcl.Zu16

func (v *UltrasonicUnoccupiedToOccupiedDelay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *UltrasonicUnoccupiedToOccupiedDelay) Value() zcl.Val     { return v }

func (v UltrasonicUnoccupiedToOccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *UltrasonicUnoccupiedToOccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = UltrasonicUnoccupiedToOccupiedDelay(*nt)
	return br, err
}

func (v UltrasonicUnoccupiedToOccupiedDelay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *UltrasonicUnoccupiedToOccupiedDelay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UltrasonicUnoccupiedToOccupiedDelay(*a)
	return nil
}

func (v *UltrasonicUnoccupiedToOccupiedDelay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = UltrasonicUnoccupiedToOccupiedDelay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UltrasonicUnoccupiedToOccupiedDelay) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const UltrasonicUnoccupiedToOccupiedThresholdAttr zcl.AttrID = 34

func (UltrasonicUnoccupiedToOccupiedThreshold) ID() zcl.AttrID {
	return UltrasonicUnoccupiedToOccupiedThresholdAttr
}
func (UltrasonicUnoccupiedToOccupiedThreshold) Readable() bool   { return true }
func (UltrasonicUnoccupiedToOccupiedThreshold) Writable() bool   { return true }
func (UltrasonicUnoccupiedToOccupiedThreshold) Reportable() bool { return false }
func (UltrasonicUnoccupiedToOccupiedThreshold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UltrasonicUnoccupiedToOccupiedThreshold) AttrID() zcl.AttrID   { return v.ID() }
func (v UltrasonicUnoccupiedToOccupiedThreshold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UltrasonicUnoccupiedToOccupiedThreshold) AttrValue() zcl.Val  { return v.Value() }

func (UltrasonicUnoccupiedToOccupiedThreshold) Name() string {
	return `Ultrasonic Unoccupied To Occupied Threshold`
}
func (UltrasonicUnoccupiedToOccupiedThreshold) Description() string {
	return `is 8 bits in length and specifies the number of movement detection events that must occur in the period UltrasonicUnoccupiedToOccupiedDelay, before the Ultrasonic sensor changes to its occupied state. This attribute is mandatory if the UltrasonicUnoccupiedToOccupiedDelay attribute is implemented`
}

// UltrasonicUnoccupiedToOccupiedThreshold is 8 bits in length and specifies the number of movement detection events that must occur in the period UltrasonicUnoccupiedToOccupiedDelay, before the Ultrasonic sensor changes to its occupied state. This attribute is mandatory if the UltrasonicUnoccupiedToOccupiedDelay attribute is implemented
type UltrasonicUnoccupiedToOccupiedThreshold zcl.Zu8

func (v *UltrasonicUnoccupiedToOccupiedThreshold) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *UltrasonicUnoccupiedToOccupiedThreshold) Value() zcl.Val     { return v }

func (v UltrasonicUnoccupiedToOccupiedThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(v).MarshalZcl()
}

func (v *UltrasonicUnoccupiedToOccupiedThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = UltrasonicUnoccupiedToOccupiedThreshold(*nt)
	return br, err
}

func (v UltrasonicUnoccupiedToOccupiedThreshold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *UltrasonicUnoccupiedToOccupiedThreshold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UltrasonicUnoccupiedToOccupiedThreshold(*a)
	return nil
}

func (v *UltrasonicUnoccupiedToOccupiedThreshold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = UltrasonicUnoccupiedToOccupiedThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UltrasonicUnoccupiedToOccupiedThreshold) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}
