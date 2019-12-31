// provides general/common attributes and operations for most zigbee devices
package general

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID

// AlarmCount Number of alarms currently defined
type AlarmCount zcl.Zu16

const AlarmCountAttr zcl.AttrID = 0

func (AlarmCount) ID() zcl.AttrID                { return AlarmCountAttr }
func (AlarmCount) Name() string                  { return "Alarm Count" }
func (AlarmCount) Readable() bool                { return true }
func (AlarmCount) Writable() bool                { return true }
func (AlarmCount) Reportable() bool              { return false }
func (AlarmCount) SceneIndex() int               { return -1 }
func (a *AlarmCount) Value() *AlarmCount         { return a }
func (a AlarmCount) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AlarmCount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmCount(*nt)
	return br, err
}

func (a *AlarmCount) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = AlarmCount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AlarmCount) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type AlarmMask zcl.Zbmp8

const AlarmMaskAttr zcl.AttrID = 19

func (AlarmMask) ID() zcl.AttrID                { return AlarmMaskAttr }
func (AlarmMask) Name() string                  { return "Alarm Mask" }
func (AlarmMask) Readable() bool                { return true }
func (AlarmMask) Writable() bool                { return true }
func (AlarmMask) Reportable() bool              { return false }
func (AlarmMask) SceneIndex() int               { return -1 }
func (a *AlarmMask) Value() *AlarmMask          { return a }
func (a AlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *AlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmMask(*nt)
	return br, err
}

func (a *AlarmMask) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = AlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "General Hardware Fault")
		case 1:
			bstr = append(bstr, "General Software Fault")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a AlarmMask) IsGeneralHardwareFault() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a AlarmMask) IsGeneralSoftwareFault() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a *AlarmMask) SetGeneralHardwareFault(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *AlarmMask) SetGeneralSoftwareFault(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}

type AlarmCode zcl.Zenum8

func (a *AlarmCode) Value() *AlarmCode          { return a }
func (a AlarmCode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *AlarmCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmCode(*nt)
	return br, err
}

func (a *AlarmCode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = AlarmCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a AlarmCode) String() string {
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

type ApplicationVersion zcl.Zu8

const ApplicationVersionAttr zcl.AttrID = 1

func (ApplicationVersion) ID() zcl.AttrID                { return ApplicationVersionAttr }
func (ApplicationVersion) Name() string                  { return "Application Version" }
func (ApplicationVersion) Readable() bool                { return true }
func (ApplicationVersion) Writable() bool                { return false }
func (ApplicationVersion) Reportable() bool              { return false }
func (ApplicationVersion) SceneIndex() int               { return -1 }
func (a *ApplicationVersion) Value() *ApplicationVersion { return a }
func (a ApplicationVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ApplicationVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ApplicationVersion(*nt)
	return br, err
}

func (a *ApplicationVersion) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ApplicationVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ApplicationVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type BatteryAlarmMask zcl.Zbmp8

const BatteryAlarmMaskAttr zcl.AttrID = 53

func (BatteryAlarmMask) ID() zcl.AttrID                { return BatteryAlarmMaskAttr }
func (BatteryAlarmMask) Name() string                  { return "Battery Alarm Mask" }
func (BatteryAlarmMask) Readable() bool                { return true }
func (BatteryAlarmMask) Writable() bool                { return true }
func (BatteryAlarmMask) Reportable() bool              { return false }
func (BatteryAlarmMask) SceneIndex() int               { return -1 }
func (a *BatteryAlarmMask) Value() *BatteryAlarmMask   { return a }
func (a BatteryAlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *BatteryAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryAlarmMask(*nt)
	return br, err
}

func (a *BatteryAlarmMask) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = BatteryAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Battery Voltage too low")
		case 1:
			bstr = append(bstr, "Battery Alarm 1")
		case 2:
			bstr = append(bstr, "Battery Alarm 2")
		case 3:
			bstr = append(bstr, "Battery Alarm 3")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a BatteryAlarmMask) IsBatteryVoltageTooLow() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a BatteryAlarmMask) IsBatteryAlarm1() bool        { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a BatteryAlarmMask) IsBatteryAlarm2() bool        { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a BatteryAlarmMask) IsBatteryAlarm3() bool        { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *BatteryAlarmMask) SetBatteryVoltageTooLow(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *BatteryAlarmMask) SetBatteryAlarm1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *BatteryAlarmMask) SetBatteryAlarm2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *BatteryAlarmMask) SetBatteryAlarm3(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}

type BatteryAlarmState zcl.Zbmp32

const BatteryAlarmStateAttr zcl.AttrID = 62

func (BatteryAlarmState) ID() zcl.AttrID                { return BatteryAlarmStateAttr }
func (BatteryAlarmState) Name() string                  { return "Battery Alarm State" }
func (BatteryAlarmState) Readable() bool                { return true }
func (BatteryAlarmState) Writable() bool                { return true }
func (BatteryAlarmState) Reportable() bool              { return false }
func (BatteryAlarmState) SceneIndex() int               { return -1 }
func (a *BatteryAlarmState) Value() *BatteryAlarmState  { return a }
func (a BatteryAlarmState) MarshalZcl() ([]byte, error) { return zcl.Zbmp32(a).MarshalZcl() }

func (a *BatteryAlarmState) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryAlarmState(*nt)
	return br, err
}

func (a *BatteryAlarmState) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp32); ok {
		*a = BatteryAlarmState(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryAlarmState) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 1")
		case 1:
			bstr = append(bstr, "Threshold 1 for Voltage or Percentage reached for Battery Source 1")
		case 10:
			bstr = append(bstr, "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 2")
		case 11:
			bstr = append(bstr, "Threshold 1 for Voltage or Percentage reached for Battery Source 2")
		case 12:
			bstr = append(bstr, "Threshold 2 for Voltage or Percentage reached for Battery Source 2")
		case 13:
			bstr = append(bstr, "Threshold 3 for Voltage or Percentage reached for Battery Source 2")
		case 2:
			bstr = append(bstr, "Threshold 2 for Voltage or Percentage reached for Battery Source 1")
		case 20:
			bstr = append(bstr, "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 3")
		case 21:
			bstr = append(bstr, "Threshold 1 for Voltage or Percentage reached for Battery Source 3")
		case 22:
			bstr = append(bstr, "Threshold 2 for Voltage or Percentage reached for Battery Source 3")
		case 23:
			bstr = append(bstr, "Threshold 3 for Voltage or Percentage reached for Battery Source 3")
		case 3:
			bstr = append(bstr, "Threshold 3 for Voltage or Percentage reached for Battery Source 1")
		case 30:
			bstr = append(bstr, "Mains power lost/unavailable")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a[:]), 0)
}
func (a BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a[:]), 1)
}
func (a BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a[:]), 10)
}
func (a BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a[:]), 11)
}
func (a BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a[:]), 12)
}
func (a BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a[:]), 13)
}
func (a BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a[:]), 2)
}
func (a BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a[:]), 20)
}
func (a BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a[:]), 21)
}
func (a BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a[:]), 22)
}
func (a BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a[:]), 23)
}
func (a BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a[:]), 3)
}
func (a BatteryAlarmState) IsMainsPowerLostUnavailable() bool { return zcl.BitmapTest([]byte(a[:]), 30) }
func (a *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 10, b))
}
func (a *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 11, b))
}
func (a *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 12, b))
}
func (a *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 13, b))
}
func (a *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 20, b))
}
func (a *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 21, b))
}
func (a *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 22, b))
}
func (a *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 23, b))
}
func (a *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}
func (a *BatteryAlarmState) SetMainsPowerLostUnavailable(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 30, b))
}

type BatteryManufacturer zcl.Zcstring

const BatteryManufacturerAttr zcl.AttrID = 48

func (BatteryManufacturer) ID() zcl.AttrID                 { return BatteryManufacturerAttr }
func (BatteryManufacturer) Name() string                   { return "Battery Manufacturer" }
func (BatteryManufacturer) Readable() bool                 { return true }
func (BatteryManufacturer) Writable() bool                 { return true }
func (BatteryManufacturer) Reportable() bool               { return false }
func (BatteryManufacturer) SceneIndex() int                { return -1 }
func (a *BatteryManufacturer) Value() *BatteryManufacturer { return a }
func (a BatteryManufacturer) MarshalZcl() ([]byte, error)  { return zcl.Zcstring(a).MarshalZcl() }

func (a *BatteryManufacturer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryManufacturer(*nt)
	return br, err
}

func (a *BatteryManufacturer) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = BatteryManufacturer(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryManufacturer) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type BatteryPercentageMinThreshold zcl.Zu8

const BatteryPercentageMinThresholdAttr zcl.AttrID = 58

func (BatteryPercentageMinThreshold) ID() zcl.AttrID                           { return BatteryPercentageMinThresholdAttr }
func (BatteryPercentageMinThreshold) Name() string                             { return "Battery Percentage Min Threshold" }
func (BatteryPercentageMinThreshold) Readable() bool                           { return true }
func (BatteryPercentageMinThreshold) Writable() bool                           { return true }
func (BatteryPercentageMinThreshold) Reportable() bool                         { return false }
func (BatteryPercentageMinThreshold) SceneIndex() int                          { return -1 }
func (a *BatteryPercentageMinThreshold) Value() *BatteryPercentageMinThreshold { return a }
func (a BatteryPercentageMinThreshold) MarshalZcl() ([]byte, error)            { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryPercentageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageMinThreshold(*nt)
	return br, err
}

func (a *BatteryPercentageMinThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryPercentageMinThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryPercentageMinThreshold) String() string {
	return zcl.Percent.Format(float64(a) / 2)
}

type BatteryPercentageThreshold1 zcl.Zu8

const BatteryPercentageThreshold1Attr zcl.AttrID = 59

func (BatteryPercentageThreshold1) ID() zcl.AttrID                         { return BatteryPercentageThreshold1Attr }
func (BatteryPercentageThreshold1) Name() string                           { return "Battery Percentage Threshold 1" }
func (BatteryPercentageThreshold1) Readable() bool                         { return true }
func (BatteryPercentageThreshold1) Writable() bool                         { return true }
func (BatteryPercentageThreshold1) Reportable() bool                       { return false }
func (BatteryPercentageThreshold1) SceneIndex() int                        { return -1 }
func (a *BatteryPercentageThreshold1) Value() *BatteryPercentageThreshold1 { return a }
func (a BatteryPercentageThreshold1) MarshalZcl() ([]byte, error)          { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryPercentageThreshold1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageThreshold1(*nt)
	return br, err
}

func (a *BatteryPercentageThreshold1) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryPercentageThreshold1(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryPercentageThreshold1) String() string {
	return zcl.Percent.Format(float64(a) / 2)
}

type BatteryPercentageThreshold2 zcl.Zu8

const BatteryPercentageThreshold2Attr zcl.AttrID = 60

func (BatteryPercentageThreshold2) ID() zcl.AttrID                         { return BatteryPercentageThreshold2Attr }
func (BatteryPercentageThreshold2) Name() string                           { return "Battery Percentage Threshold 2" }
func (BatteryPercentageThreshold2) Readable() bool                         { return true }
func (BatteryPercentageThreshold2) Writable() bool                         { return true }
func (BatteryPercentageThreshold2) Reportable() bool                       { return false }
func (BatteryPercentageThreshold2) SceneIndex() int                        { return -1 }
func (a *BatteryPercentageThreshold2) Value() *BatteryPercentageThreshold2 { return a }
func (a BatteryPercentageThreshold2) MarshalZcl() ([]byte, error)          { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryPercentageThreshold2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageThreshold2(*nt)
	return br, err
}

func (a *BatteryPercentageThreshold2) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryPercentageThreshold2(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryPercentageThreshold2) String() string {
	return zcl.Percent.Format(float64(a) / 2)
}

type BatteryPercentageThreshold3 zcl.Zu8

const BatteryPercentageThreshold3Attr zcl.AttrID = 61

func (BatteryPercentageThreshold3) ID() zcl.AttrID                         { return BatteryPercentageThreshold3Attr }
func (BatteryPercentageThreshold3) Name() string                           { return "Battery Percentage Threshold 3" }
func (BatteryPercentageThreshold3) Readable() bool                         { return true }
func (BatteryPercentageThreshold3) Writable() bool                         { return true }
func (BatteryPercentageThreshold3) Reportable() bool                       { return false }
func (BatteryPercentageThreshold3) SceneIndex() int                        { return -1 }
func (a *BatteryPercentageThreshold3) Value() *BatteryPercentageThreshold3 { return a }
func (a BatteryPercentageThreshold3) MarshalZcl() ([]byte, error)          { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryPercentageThreshold3) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageThreshold3(*nt)
	return br, err
}

func (a *BatteryPercentageThreshold3) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryPercentageThreshold3(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryPercentageThreshold3) String() string {
	return zcl.Percent.Format(float64(a) / 2)
}

type BatteryQuantity zcl.Zu8

const BatteryQuantityAttr zcl.AttrID = 51

func (BatteryQuantity) ID() zcl.AttrID                { return BatteryQuantityAttr }
func (BatteryQuantity) Name() string                  { return "Battery Quantity" }
func (BatteryQuantity) Readable() bool                { return true }
func (BatteryQuantity) Writable() bool                { return true }
func (BatteryQuantity) Reportable() bool              { return false }
func (BatteryQuantity) SceneIndex() int               { return -1 }
func (a *BatteryQuantity) Value() *BatteryQuantity    { return a }
func (a BatteryQuantity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryQuantity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryQuantity(*nt)
	return br, err
}

func (a *BatteryQuantity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryQuantity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryQuantity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type BatteryRatedVoltage zcl.Zu8

const BatteryRatedVoltageAttr zcl.AttrID = 52

func (BatteryRatedVoltage) ID() zcl.AttrID                 { return BatteryRatedVoltageAttr }
func (BatteryRatedVoltage) Name() string                   { return "Battery Rated Voltage" }
func (BatteryRatedVoltage) Readable() bool                 { return true }
func (BatteryRatedVoltage) Writable() bool                 { return true }
func (BatteryRatedVoltage) Reportable() bool               { return false }
func (BatteryRatedVoltage) SceneIndex() int                { return -1 }
func (a *BatteryRatedVoltage) Value() *BatteryRatedVoltage { return a }
func (a BatteryRatedVoltage) MarshalZcl() ([]byte, error)  { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryRatedVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryRatedVoltage(*nt)
	return br, err
}

func (a *BatteryRatedVoltage) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryRatedVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryRatedVoltage) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryRemaining zcl.Zu8

const BatteryRemainingAttr zcl.AttrID = 33

func (BatteryRemaining) ID() zcl.AttrID                { return BatteryRemainingAttr }
func (BatteryRemaining) Name() string                  { return "Battery Remaining" }
func (BatteryRemaining) Readable() bool                { return true }
func (BatteryRemaining) Writable() bool                { return false }
func (BatteryRemaining) Reportable() bool              { return true }
func (BatteryRemaining) SceneIndex() int               { return -1 }
func (a *BatteryRemaining) Value() *BatteryRemaining   { return a }
func (a BatteryRemaining) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryRemaining) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryRemaining(*nt)
	return br, err
}

func (a *BatteryRemaining) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryRemaining(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryRemaining) String() string {
	return zcl.Percent.Format(float64(a) / 2)
}

type BatterySize zcl.Zenum8

const BatterySizeAttr zcl.AttrID = 49

func (BatterySize) ID() zcl.AttrID                { return BatterySizeAttr }
func (BatterySize) Name() string                  { return "Battery Size" }
func (BatterySize) Readable() bool                { return true }
func (BatterySize) Writable() bool                { return true }
func (BatterySize) Reportable() bool              { return false }
func (BatterySize) SceneIndex() int               { return -1 }
func (a *BatterySize) Value() *BatterySize        { return a }
func (a BatterySize) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *BatterySize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatterySize(*nt)
	return br, err
}

func (a *BatterySize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = BatterySize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatterySize) String() string {
	switch a {
	case 0x00:
		return "No Battery"
	case 0x01:
		return "Built in"
	case 0x02:
		return "Other"
	case 0x03:
		return "AA"
	case 0x04:
		return "AAA"
	case 0x05:
		return "C"
	case 0x06:
		return "D"
	case 0x07:
		return "CR2"
	case 0x08:
		return "CR123A"
	case 0xFF:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a BatterySize) IsNoBattery() bool { return a == 0x00 }
func (a BatterySize) IsBuiltIn() bool   { return a == 0x01 }
func (a BatterySize) IsOther() bool     { return a == 0x02 }
func (a BatterySize) IsAa() bool        { return a == 0x03 }
func (a BatterySize) IsAaa() bool       { return a == 0x04 }
func (a BatterySize) IsC() bool         { return a == 0x05 }
func (a BatterySize) IsD() bool         { return a == 0x06 }
func (a BatterySize) IsCr2() bool       { return a == 0x07 }
func (a BatterySize) IsCr123a() bool    { return a == 0x08 }
func (a BatterySize) IsUnknown() bool   { return a == 0xFF }
func (a *BatterySize) SetNoBattery()    { *a = 0x00 }
func (a *BatterySize) SetBuiltIn()      { *a = 0x01 }
func (a *BatterySize) SetOther()        { *a = 0x02 }
func (a *BatterySize) SetAa()           { *a = 0x03 }
func (a *BatterySize) SetAaa()          { *a = 0x04 }
func (a *BatterySize) SetC()            { *a = 0x05 }
func (a *BatterySize) SetD()            { *a = 0x06 }
func (a *BatterySize) SetCr2()          { *a = 0x07 }
func (a *BatterySize) SetCr123a()       { *a = 0x08 }
func (a *BatterySize) SetUnknown()      { *a = 0xFF }

type BatteryVoltage zcl.Zu8

const BatteryVoltageAttr zcl.AttrID = 32

func (BatteryVoltage) ID() zcl.AttrID                { return BatteryVoltageAttr }
func (BatteryVoltage) Name() string                  { return "Battery Voltage" }
func (BatteryVoltage) Readable() bool                { return true }
func (BatteryVoltage) Writable() bool                { return false }
func (BatteryVoltage) Reportable() bool              { return false }
func (BatteryVoltage) SceneIndex() int               { return -1 }
func (a *BatteryVoltage) Value() *BatteryVoltage     { return a }
func (a BatteryVoltage) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltage(*nt)
	return br, err
}

func (a *BatteryVoltage) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryVoltage) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryVoltageMinThreshold zcl.Zu8

const BatteryVoltageMinThresholdAttr zcl.AttrID = 54

func (BatteryVoltageMinThreshold) ID() zcl.AttrID                        { return BatteryVoltageMinThresholdAttr }
func (BatteryVoltageMinThreshold) Name() string                          { return "Battery Voltage Min Threshold" }
func (BatteryVoltageMinThreshold) Readable() bool                        { return true }
func (BatteryVoltageMinThreshold) Writable() bool                        { return true }
func (BatteryVoltageMinThreshold) Reportable() bool                      { return false }
func (BatteryVoltageMinThreshold) SceneIndex() int                       { return -1 }
func (a *BatteryVoltageMinThreshold) Value() *BatteryVoltageMinThreshold { return a }
func (a BatteryVoltageMinThreshold) MarshalZcl() ([]byte, error)         { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryVoltageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageMinThreshold(*nt)
	return br, err
}

func (a *BatteryVoltageMinThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryVoltageMinThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryVoltageMinThreshold) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryVoltageThreshold1 zcl.Zu8

const BatteryVoltageThreshold1Attr zcl.AttrID = 55

func (BatteryVoltageThreshold1) ID() zcl.AttrID                      { return BatteryVoltageThreshold1Attr }
func (BatteryVoltageThreshold1) Name() string                        { return "Battery Voltage Threshold 1" }
func (BatteryVoltageThreshold1) Readable() bool                      { return true }
func (BatteryVoltageThreshold1) Writable() bool                      { return true }
func (BatteryVoltageThreshold1) Reportable() bool                    { return false }
func (BatteryVoltageThreshold1) SceneIndex() int                     { return -1 }
func (a *BatteryVoltageThreshold1) Value() *BatteryVoltageThreshold1 { return a }
func (a BatteryVoltageThreshold1) MarshalZcl() ([]byte, error)       { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryVoltageThreshold1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageThreshold1(*nt)
	return br, err
}

func (a *BatteryVoltageThreshold1) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryVoltageThreshold1(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryVoltageThreshold1) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryVoltageThreshold2 zcl.Zu8

const BatteryVoltageThreshold2Attr zcl.AttrID = 56

func (BatteryVoltageThreshold2) ID() zcl.AttrID                      { return BatteryVoltageThreshold2Attr }
func (BatteryVoltageThreshold2) Name() string                        { return "Battery Voltage Threshold 2" }
func (BatteryVoltageThreshold2) Readable() bool                      { return true }
func (BatteryVoltageThreshold2) Writable() bool                      { return true }
func (BatteryVoltageThreshold2) Reportable() bool                    { return false }
func (BatteryVoltageThreshold2) SceneIndex() int                     { return -1 }
func (a *BatteryVoltageThreshold2) Value() *BatteryVoltageThreshold2 { return a }
func (a BatteryVoltageThreshold2) MarshalZcl() ([]byte, error)       { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryVoltageThreshold2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageThreshold2(*nt)
	return br, err
}

func (a *BatteryVoltageThreshold2) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryVoltageThreshold2(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryVoltageThreshold2) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryVoltageThreshold3 zcl.Zu8

const BatteryVoltageThreshold3Attr zcl.AttrID = 57

func (BatteryVoltageThreshold3) ID() zcl.AttrID                      { return BatteryVoltageThreshold3Attr }
func (BatteryVoltageThreshold3) Name() string                        { return "Battery Voltage Threshold 3" }
func (BatteryVoltageThreshold3) Readable() bool                      { return true }
func (BatteryVoltageThreshold3) Writable() bool                      { return true }
func (BatteryVoltageThreshold3) Reportable() bool                    { return false }
func (BatteryVoltageThreshold3) SceneIndex() int                     { return -1 }
func (a *BatteryVoltageThreshold3) Value() *BatteryVoltageThreshold3 { return a }
func (a BatteryVoltageThreshold3) MarshalZcl() ([]byte, error)       { return zcl.Zu8(a).MarshalZcl() }

func (a *BatteryVoltageThreshold3) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageThreshold3(*nt)
	return br, err
}

func (a *BatteryVoltageThreshold3) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BatteryVoltageThreshold3(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryVoltageThreshold3) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type BatteryCapacity zcl.Zu16

const BatteryCapacityAttr zcl.AttrID = 50

func (BatteryCapacity) ID() zcl.AttrID                { return BatteryCapacityAttr }
func (BatteryCapacity) Name() string                  { return "Battery capacity" }
func (BatteryCapacity) Readable() bool                { return true }
func (BatteryCapacity) Writable() bool                { return true }
func (BatteryCapacity) Reportable() bool              { return false }
func (BatteryCapacity) SceneIndex() int               { return -1 }
func (a *BatteryCapacity) Value() *BatteryCapacity    { return a }
func (a BatteryCapacity) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *BatteryCapacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryCapacity(*nt)
	return br, err
}

func (a *BatteryCapacity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = BatteryCapacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BatteryCapacity) String() string {
	return zcl.MilliAmpereHours.Format(float64(a) / 0.1)
}

type CalculationPeriod zcl.Zu16

const CalculationPeriodAttr zcl.AttrID = 22

func (CalculationPeriod) ID() zcl.AttrID                { return CalculationPeriodAttr }
func (CalculationPeriod) Name() string                  { return "Calculation Period" }
func (CalculationPeriod) Readable() bool                { return true }
func (CalculationPeriod) Writable() bool                { return true }
func (CalculationPeriod) Reportable() bool              { return false }
func (CalculationPeriod) SceneIndex() int               { return -1 }
func (a *CalculationPeriod) Value() *CalculationPeriod  { return a }
func (a CalculationPeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *CalculationPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CalculationPeriod(*nt)
	return br, err
}

func (a *CalculationPeriod) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = CalculationPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CalculationPeriod) String() string {
	return zcl.Seconds.Format(float64(a))
}

type ClusterId zcl.Zu16

func (a *ClusterId) Value() *ClusterId          { return a }
func (a ClusterId) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ClusterId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ClusterId(*nt)
	return br, err
}

func (a *ClusterId) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ClusterId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ClusterId) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ClusterRevision zcl.Zu16

const ClusterRevisionAttr zcl.AttrID = 65533

func (ClusterRevision) ID() zcl.AttrID                { return ClusterRevisionAttr }
func (ClusterRevision) Name() string                  { return "Cluster Revision" }
func (ClusterRevision) Readable() bool                { return true }
func (ClusterRevision) Writable() bool                { return true }
func (ClusterRevision) Reportable() bool              { return false }
func (ClusterRevision) SceneIndex() int               { return -1 }
func (a *ClusterRevision) Value() *ClusterRevision    { return a }
func (a ClusterRevision) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ClusterRevision) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ClusterRevision(*nt)
	return br, err
}

func (a *ClusterRevision) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ClusterRevision(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ClusterRevision) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type Configuration zcl.Zbmp16

const ConfigurationAttr zcl.AttrID = 49

func (Configuration) ID() zcl.AttrID                { return ConfigurationAttr }
func (Configuration) Name() string                  { return "Configuration" }
func (Configuration) Readable() bool                { return true }
func (Configuration) Writable() bool                { return true }
func (Configuration) Reportable() bool              { return false }
func (Configuration) SceneIndex() int               { return -1 }
func (a *Configuration) Value() *Configuration      { return a }
func (a Configuration) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(a).MarshalZcl() }

func (a *Configuration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = Configuration(*nt)
	return br, err
}

func (a *Configuration) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp16); ok {
		*a = Configuration(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Configuration) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Touchlink enabled 0")
		case 1:
			bstr = append(bstr, "Touchlink enabled 1")
		case 3:
			bstr = append(bstr, "Touchlink enabled 2")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a Configuration) IsTouchlinkEnabled0() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a Configuration) IsTouchlinkEnabled1() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a Configuration) IsTouchlinkEnabled2() bool { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *Configuration) SetTouchlinkEnabled0(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *Configuration) SetTouchlinkEnabled1(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *Configuration) SetTouchlinkEnabled2(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}

type CurrentGroup zcl.Zu16

const CurrentGroupAttr zcl.AttrID = 2

func (CurrentGroup) ID() zcl.AttrID                { return CurrentGroupAttr }
func (CurrentGroup) Name() string                  { return "Current Group" }
func (CurrentGroup) Readable() bool                { return true }
func (CurrentGroup) Writable() bool                { return false }
func (CurrentGroup) Reportable() bool              { return false }
func (CurrentGroup) SceneIndex() int               { return -1 }
func (a *CurrentGroup) Value() *CurrentGroup       { return a }
func (a CurrentGroup) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *CurrentGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentGroup(*nt)
	return br, err
}

func (a *CurrentGroup) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = CurrentGroup(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentGroup) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// CurrentLevel The CurrentLevel attribute represents the current level of this device. meaning of 'level' is device dependent.
type CurrentLevel zcl.Zu8

const CurrentLevelAttr zcl.AttrID = 0

func (CurrentLevel) ID() zcl.AttrID                { return CurrentLevelAttr }
func (CurrentLevel) Name() string                  { return "Current Level" }
func (CurrentLevel) Readable() bool                { return true }
func (CurrentLevel) Writable() bool                { return false }
func (CurrentLevel) Reportable() bool              { return true }
func (CurrentLevel) SceneIndex() int               { return 1 }
func (a *CurrentLevel) Value() *CurrentLevel       { return a }
func (a CurrentLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *CurrentLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentLevel(*nt)
	return br, err
}

func (a *CurrentLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = CurrentLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentLevel) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

type CurrentScene zcl.Zu8

const CurrentSceneAttr zcl.AttrID = 1

func (CurrentScene) ID() zcl.AttrID                { return CurrentSceneAttr }
func (CurrentScene) Name() string                  { return "Current Scene" }
func (CurrentScene) Readable() bool                { return true }
func (CurrentScene) Writable() bool                { return false }
func (CurrentScene) Reportable() bool              { return false }
func (CurrentScene) SceneIndex() int               { return -1 }
func (a *CurrentScene) Value() *CurrentScene       { return a }
func (a CurrentScene) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *CurrentScene) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentScene(*nt)
	return br, err
}

func (a *CurrentScene) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = CurrentScene(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentScene) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type CurrentTemperature zcl.Zs16

const CurrentTemperatureAttr zcl.AttrID = 0

func (CurrentTemperature) ID() zcl.AttrID                { return CurrentTemperatureAttr }
func (CurrentTemperature) Name() string                  { return "Current Temperature" }
func (CurrentTemperature) Readable() bool                { return true }
func (CurrentTemperature) Writable() bool                { return false }
func (CurrentTemperature) Reportable() bool              { return false }
func (CurrentTemperature) SceneIndex() int               { return -1 }
func (a *CurrentTemperature) Value() *CurrentTemperature { return a }
func (a CurrentTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *CurrentTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentTemperature(*nt)
	return br, err
}

func (a *CurrentTemperature) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = CurrentTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

type DateCode zcl.Zcstring

const DateCodeAttr zcl.AttrID = 6

func (DateCode) ID() zcl.AttrID                { return DateCodeAttr }
func (DateCode) Name() string                  { return "Date Code" }
func (DateCode) Readable() bool                { return true }
func (DateCode) Writable() bool                { return false }
func (DateCode) Reportable() bool              { return false }
func (DateCode) SceneIndex() int               { return -1 }
func (a *DateCode) Value() *DateCode           { return a }
func (a DateCode) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *DateCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = DateCode(*nt)
	return br, err
}

func (a *DateCode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = DateCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DateCode) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type DefaultMoveRate zcl.Zu8

const DefaultMoveRateAttr zcl.AttrID = 20

func (DefaultMoveRate) ID() zcl.AttrID                { return DefaultMoveRateAttr }
func (DefaultMoveRate) Name() string                  { return "Default Move Rate" }
func (DefaultMoveRate) Readable() bool                { return true }
func (DefaultMoveRate) Writable() bool                { return true }
func (DefaultMoveRate) Reportable() bool              { return false }
func (DefaultMoveRate) SceneIndex() int               { return -1 }
func (a *DefaultMoveRate) Value() *DefaultMoveRate    { return a }
func (a DefaultMoveRate) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *DefaultMoveRate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = DefaultMoveRate(*nt)
	return br, err
}

func (a *DefaultMoveRate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = DefaultMoveRate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DefaultMoveRate) String() string {
	return zcl.PercentPerSecond.Format(float64(a) / 2.54)
}

type Device zcl.Zuid

func (a *Device) Value() *Device             { return a }
func (a Device) MarshalZcl() ([]byte, error) { return zcl.Zuid(a).MarshalZcl() }

func (a *Device) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = Device(*nt)
	return br, err
}

func (a *Device) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zuid); ok {
		*a = Device(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Device) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(a))
}

type DeviceEnabled zcl.Zbool

const DeviceEnabledAttr zcl.AttrID = 18

func (DeviceEnabled) ID() zcl.AttrID                { return DeviceEnabledAttr }
func (DeviceEnabled) Name() string                  { return "Device Enabled" }
func (DeviceEnabled) Readable() bool                { return true }
func (DeviceEnabled) Writable() bool                { return true }
func (DeviceEnabled) Reportable() bool              { return false }
func (DeviceEnabled) SceneIndex() int               { return -1 }
func (a *DeviceEnabled) Value() *DeviceEnabled      { return a }
func (a DeviceEnabled) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *DeviceEnabled) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceEnabled(*nt)
	return br, err
}

func (a *DeviceEnabled) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = DeviceEnabled(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DeviceEnabled) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type DeviceTempAlarmMask zcl.Zbmp8

const DeviceTempAlarmMaskAttr zcl.AttrID = 16

func (DeviceTempAlarmMask) ID() zcl.AttrID                 { return DeviceTempAlarmMaskAttr }
func (DeviceTempAlarmMask) Name() string                   { return "Device Temp Alarm Mask" }
func (DeviceTempAlarmMask) Readable() bool                 { return true }
func (DeviceTempAlarmMask) Writable() bool                 { return true }
func (DeviceTempAlarmMask) Reportable() bool               { return false }
func (DeviceTempAlarmMask) SceneIndex() int                { return -1 }
func (a *DeviceTempAlarmMask) Value() *DeviceTempAlarmMask { return a }
func (a DeviceTempAlarmMask) MarshalZcl() ([]byte, error)  { return zcl.Zbmp8(a).MarshalZcl() }

func (a *DeviceTempAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceTempAlarmMask(*nt)
	return br, err
}

func (a *DeviceTempAlarmMask) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = DeviceTempAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DeviceTempAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Temperature too low")
		case 1:
			bstr = append(bstr, "Temperature too high")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a DeviceTempAlarmMask) IsTemperatureTooLow() bool  { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a DeviceTempAlarmMask) IsTemperatureTooHigh() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a *DeviceTempAlarmMask) SetTemperatureTooLow(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *DeviceTempAlarmMask) SetTemperatureTooHigh(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}

type DisableLocalConfig zcl.Zbmp8

const DisableLocalConfigAttr zcl.AttrID = 20

func (DisableLocalConfig) ID() zcl.AttrID                { return DisableLocalConfigAttr }
func (DisableLocalConfig) Name() string                  { return "Disable Local Config" }
func (DisableLocalConfig) Readable() bool                { return true }
func (DisableLocalConfig) Writable() bool                { return true }
func (DisableLocalConfig) Reportable() bool              { return false }
func (DisableLocalConfig) SceneIndex() int               { return -1 }
func (a *DisableLocalConfig) Value() *DisableLocalConfig { return a }
func (a DisableLocalConfig) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *DisableLocalConfig) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DisableLocalConfig(*nt)
	return br, err
}

func (a *DisableLocalConfig) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = DisableLocalConfig(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DisableLocalConfig) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Disable Factory Reset")
		case 1:
			bstr = append(bstr, "Disable Device Configuration")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a DisableLocalConfig) IsDisableFactoryReset() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a DisableLocalConfig) IsDisableDeviceConfiguration() bool {
	return zcl.BitmapTest([]byte(a[:]), 1)
}
func (a *DisableLocalConfig) SetDisableFactoryReset(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *DisableLocalConfig) SetDisableDeviceConfiguration(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}

type Distance zcl.Zu16

func (a *Distance) Value() *Distance           { return a }
func (a Distance) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Distance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Distance(*nt)
	return br, err
}

func (a *Distance) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Distance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Distance) String() string {
	return zcl.Meters.Format(float64(a) / 10)
}

type DstEnd zcl.Zutc

const DstEndAttr zcl.AttrID = 4

func (DstEnd) ID() zcl.AttrID                { return DstEndAttr }
func (DstEnd) Name() string                  { return "Dst End" }
func (DstEnd) Readable() bool                { return true }
func (DstEnd) Writable() bool                { return true }
func (DstEnd) Reportable() bool              { return false }
func (DstEnd) SceneIndex() int               { return -1 }
func (a *DstEnd) Value() *DstEnd             { return a }
func (a DstEnd) MarshalZcl() ([]byte, error) { return zcl.Zutc(a).MarshalZcl() }

func (a *DstEnd) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = DstEnd(*nt)
	return br, err
}

func (a *DstEnd) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zutc); ok {
		*a = DstEnd(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DstEnd) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(a))
}

type DstShift zcl.Zs32

const DstShiftAttr zcl.AttrID = 5

func (DstShift) ID() zcl.AttrID                { return DstShiftAttr }
func (DstShift) Name() string                  { return "Dst Shift" }
func (DstShift) Readable() bool                { return true }
func (DstShift) Writable() bool                { return true }
func (DstShift) Reportable() bool              { return false }
func (DstShift) SceneIndex() int               { return -1 }
func (a *DstShift) Value() *DstShift           { return a }
func (a DstShift) MarshalZcl() ([]byte, error) { return zcl.Zs32(a).MarshalZcl() }

func (a *DstShift) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = DstShift(*nt)
	return br, err
}

func (a *DstShift) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs32); ok {
		*a = DstShift(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DstShift) String() string {
	return zcl.Seconds.Format(float64(a))
}

type DstStart zcl.Zutc

const DstStartAttr zcl.AttrID = 3

func (DstStart) ID() zcl.AttrID                { return DstStartAttr }
func (DstStart) Name() string                  { return "Dst Start" }
func (DstStart) Readable() bool                { return true }
func (DstStart) Writable() bool                { return true }
func (DstStart) Reportable() bool              { return false }
func (DstStart) SceneIndex() int               { return -1 }
func (a *DstStart) Value() *DstStart           { return a }
func (a DstStart) MarshalZcl() ([]byte, error) { return zcl.Zutc(a).MarshalZcl() }

func (a *DstStart) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = DstStart(*nt)
	return br, err
}

func (a *DstStart) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zutc); ok {
		*a = DstStart(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a DstStart) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(a))
}

// EffectIdentifier when turning lights off
type EffectIdentifier zcl.Zenum8

func (a *EffectIdentifier) Value() *EffectIdentifier   { return a }
func (a EffectIdentifier) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *EffectIdentifier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = EffectIdentifier(*nt)
	return br, err
}

func (a *EffectIdentifier) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = EffectIdentifier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a EffectIdentifier) String() string {
	switch a {
	case 0x00:
		return "Delayed all off"
	case 0x01:
		return "Dying light"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a EffectIdentifier) IsDelayedAllOff() bool { return a == 0x00 }
func (a EffectIdentifier) IsDyingLight() bool    { return a == 0x01 }
func (a *EffectIdentifier) SetDelayedAllOff()    { *a = 0x00 }
func (a *EffectIdentifier) SetDyingLight()       { *a = 0x01 }

type EffectVariant zcl.Zenum8

func (a *EffectVariant) Value() *EffectVariant      { return a }
func (a EffectVariant) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *EffectVariant) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = EffectVariant(*nt)
	return br, err
}

func (a *EffectVariant) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = EffectVariant(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a EffectVariant) String() string {
	switch a {
	case 0x00:
		return "Fade to off in 0.8 seconds / 20% dim up in 0.5 then fade to off in 1 second"
	case 0x01:
		return "No Fade"
	case 0x02:
		return "50% dim down in 0.8s then fade off in 12s"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a EffectVariant) IsFadeToOffIn08Seconds20DimUpIn05ThenFadeToOffIn1Second() bool {
	return a == 0x00
}
func (a EffectVariant) IsNoFade() bool                                             { return a == 0x01 }
func (a EffectVariant) Is50DimDownIn08SThenFadeOffIn12S() bool                     { return a == 0x02 }
func (a *EffectVariant) SetFadeToOffIn08Seconds20DimUpIn05ThenFadeToOffIn1Second() { *a = 0x00 }
func (a *EffectVariant) SetNoFade()                                                { *a = 0x01 }
func (a *EffectVariant) Set50DimDownIn08SThenFadeOffIn12S()                        { *a = 0x02 }

type GlobalSceneControl zcl.Zbool

const GlobalSceneControlAttr zcl.AttrID = 16384

func (GlobalSceneControl) ID() zcl.AttrID                { return GlobalSceneControlAttr }
func (GlobalSceneControl) Name() string                  { return "Global Scene Control" }
func (GlobalSceneControl) Readable() bool                { return true }
func (GlobalSceneControl) Writable() bool                { return false }
func (GlobalSceneControl) Reportable() bool              { return false }
func (GlobalSceneControl) SceneIndex() int               { return -1 }
func (a *GlobalSceneControl) Value() *GlobalSceneControl { return a }
func (a GlobalSceneControl) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *GlobalSceneControl) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = GlobalSceneControl(*nt)
	return br, err
}

func (a *GlobalSceneControl) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = GlobalSceneControl(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GlobalSceneControl) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type GroupId zcl.Zu16

func (a *GroupId) Value() *GroupId            { return a }
func (a GroupId) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *GroupId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupId(*nt)
	return br, err
}

func (a *GroupId) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = GroupId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GroupId) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type GroupNameSupport zcl.Zbmp8

const GroupNameSupportAttr zcl.AttrID = 0

func (GroupNameSupport) ID() zcl.AttrID                { return GroupNameSupportAttr }
func (GroupNameSupport) Name() string                  { return "Group Name Support" }
func (GroupNameSupport) Readable() bool                { return true }
func (GroupNameSupport) Writable() bool                { return false }
func (GroupNameSupport) Reportable() bool              { return false }
func (GroupNameSupport) SceneIndex() int               { return -1 }
func (a *GroupNameSupport) Value() *GroupNameSupport   { return a }
func (a GroupNameSupport) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *GroupNameSupport) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupNameSupport(*nt)
	return br, err
}

func (a *GroupNameSupport) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = GroupNameSupport(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GroupNameSupport) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 7:
			bstr = append(bstr, "Names Supported")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a GroupNameSupport) IsNamesSupported() bool { return zcl.BitmapTest([]byte(a[:]), 7) }
func (a *GroupNameSupport) SetNamesSupported(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 7, b))
}

// GroupCapacity specifies remaining number of groups that can be added.
// If set to 0xFE, at least one more group can be added (exact number unknown)
// If set to 0xFF, it's unknown if any more groups can be added
type GroupCapacity zcl.Zu8

func (a *GroupCapacity) Value() *GroupCapacity      { return a }
func (a GroupCapacity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *GroupCapacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupCapacity(*nt)
	return br, err
}

func (a *GroupCapacity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = GroupCapacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GroupCapacity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type GroupList zcl.Zset

func (a *GroupList) Value() *GroupList { return a }
func (a GroupList) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *GroupList) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zset)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupList(*nt)
	return br, err
}

func (a *GroupList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zset); ok {
		*a = GroupList(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GroupList) String() string {
	return zcl.Sprintf("%v", zcl.Zset(a))
}

type GroupName zcl.Zcstring

func (a *GroupName) Value() *GroupName          { return a }
func (a GroupName) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *GroupName) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = GroupName(*nt)
	return br, err
}

func (a *GroupName) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = GroupName(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a GroupName) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type HwVersion zcl.Zu8

const HwVersionAttr zcl.AttrID = 3

func (HwVersion) ID() zcl.AttrID                { return HwVersionAttr }
func (HwVersion) Name() string                  { return "HW Version" }
func (HwVersion) Readable() bool                { return true }
func (HwVersion) Writable() bool                { return false }
func (HwVersion) Reportable() bool              { return false }
func (HwVersion) SceneIndex() int               { return -1 }
func (a *HwVersion) Value() *HwVersion          { return a }
func (a HwVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *HwVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = HwVersion(*nt)
	return br, err
}

func (a *HwVersion) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = HwVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a HwVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type HighTempDwellTripPoint zcl.Zu24

const HighTempDwellTripPointAttr zcl.AttrID = 20

func (HighTempDwellTripPoint) ID() zcl.AttrID                    { return HighTempDwellTripPointAttr }
func (HighTempDwellTripPoint) Name() string                      { return "High Temp Dwell Trip Point" }
func (HighTempDwellTripPoint) Readable() bool                    { return false }
func (HighTempDwellTripPoint) Writable() bool                    { return false }
func (HighTempDwellTripPoint) Reportable() bool                  { return false }
func (HighTempDwellTripPoint) SceneIndex() int                   { return -1 }
func (a *HighTempDwellTripPoint) Value() *HighTempDwellTripPoint { return a }
func (a HighTempDwellTripPoint) MarshalZcl() ([]byte, error)     { return zcl.Zu24(a).MarshalZcl() }

func (a *HighTempDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = HighTempDwellTripPoint(*nt)
	return br, err
}

func (a *HighTempDwellTripPoint) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu24); ok {
		*a = HighTempDwellTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a HighTempDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(a))
}

// HighTempThreshold If the current temperature goes above the threshold for longer
// than the time specified by high temp dwell, an alarm will be triggered
type HighTempThreshold zcl.Zs16

const HighTempThresholdAttr zcl.AttrID = 18

func (HighTempThreshold) ID() zcl.AttrID                { return HighTempThresholdAttr }
func (HighTempThreshold) Name() string                  { return "High Temp Threshold" }
func (HighTempThreshold) Readable() bool                { return true }
func (HighTempThreshold) Writable() bool                { return true }
func (HighTempThreshold) Reportable() bool              { return false }
func (HighTempThreshold) SceneIndex() int               { return -1 }
func (a *HighTempThreshold) Value() *HighTempThreshold  { return a }
func (a HighTempThreshold) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *HighTempThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = HighTempThreshold(*nt)
	return br, err
}

func (a *HighTempThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = HighTempThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a HighTempThreshold) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

// IdentifyEffect The effect identifier field specifies the identify effect to use.
type IdentifyEffect zcl.Zenum8

func (a *IdentifyEffect) Value() *IdentifyEffect     { return a }
func (a IdentifyEffect) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *IdentifyEffect) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = IdentifyEffect(*nt)
	return br, err
}

func (a *IdentifyEffect) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = IdentifyEffect(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IdentifyEffect) String() string {
	switch a {
	case 0x00:
		return "Blink"
	case 0x01:
		return "Breathe"
	case 0x02:
		return "Okay"
	case 0x0B:
		return "Channel change"
	case 0xFE:
		return "Finish"
	case 0xFF:
		return "Stop"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a IdentifyEffect) IsBlink() bool         { return a == 0x00 }
func (a IdentifyEffect) IsBreathe() bool       { return a == 0x01 }
func (a IdentifyEffect) IsOkay() bool          { return a == 0x02 }
func (a IdentifyEffect) IsChannelChange() bool { return a == 0x0B }
func (a IdentifyEffect) IsFinish() bool        { return a == 0xFE }
func (a IdentifyEffect) IsStop() bool          { return a == 0xFF }
func (a *IdentifyEffect) SetBlink()            { *a = 0x00 }
func (a *IdentifyEffect) SetBreathe()          { *a = 0x01 }
func (a *IdentifyEffect) SetOkay()             { *a = 0x02 }
func (a *IdentifyEffect) SetChannelChange()    { *a = 0x0B }
func (a *IdentifyEffect) SetFinish()           { *a = 0xFE }
func (a *IdentifyEffect) SetStop()             { *a = 0xFF }

// IdentifyEffectVariant The effect identifier field specifies the identify effect to use.
type IdentifyEffectVariant zcl.Zenum8

func (a *IdentifyEffectVariant) Value() *IdentifyEffectVariant { return a }
func (a IdentifyEffectVariant) MarshalZcl() ([]byte, error)    { return zcl.Zenum8(a).MarshalZcl() }

func (a *IdentifyEffectVariant) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = IdentifyEffectVariant(*nt)
	return br, err
}

func (a *IdentifyEffectVariant) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = IdentifyEffectVariant(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IdentifyEffectVariant) String() string {
	switch a {
	case 0x00:
		return "Default"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a IdentifyEffectVariant) IsDefault() bool { return a == 0x00 }
func (a *IdentifyEffectVariant) SetDefault()    { *a = 0x00 }

// IdentifyTime The time in seconds for which a device will stay in identify mode.
type IdentifyTime zcl.Zu16

const IdentifyTimeAttr zcl.AttrID = 0

func (IdentifyTime) ID() zcl.AttrID                { return IdentifyTimeAttr }
func (IdentifyTime) Name() string                  { return "Identify Time" }
func (IdentifyTime) Readable() bool                { return true }
func (IdentifyTime) Writable() bool                { return true }
func (IdentifyTime) Reportable() bool              { return false }
func (IdentifyTime) SceneIndex() int               { return -1 }
func (a *IdentifyTime) Value() *IdentifyTime       { return a }
func (a IdentifyTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *IdentifyTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = IdentifyTime(*nt)
	return br, err
}

func (a *IdentifyTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = IdentifyTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IdentifyTime) String() string {
	return zcl.Seconds.Format(float64(a))
}

// IdentifyTimeout The time in seconds for which a device will stay in identify mode.
type IdentifyTimeout zcl.Zu16

func (a *IdentifyTimeout) Value() *IdentifyTimeout    { return a }
func (a IdentifyTimeout) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *IdentifyTimeout) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = IdentifyTimeout(*nt)
	return br, err
}

func (a *IdentifyTimeout) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = IdentifyTimeout(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IdentifyTimeout) String() string {
	return zcl.Seconds.Format(float64(a))
}

type LedIndication zcl.Zbool

const LedIndicationAttr zcl.AttrID = 51

func (LedIndication) ID() zcl.AttrID                { return LedIndicationAttr }
func (LedIndication) Name() string                  { return "LED Indication" }
func (LedIndication) Readable() bool                { return true }
func (LedIndication) Writable() bool                { return true }
func (LedIndication) Reportable() bool              { return false }
func (LedIndication) SceneIndex() int               { return -1 }
func (a *LedIndication) Value() *LedIndication      { return a }
func (a LedIndication) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *LedIndication) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = LedIndication(*nt)
	return br, err
}

func (a *LedIndication) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = LedIndication(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LedIndication) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type LastSetTime zcl.Zutc

const LastSetTimeAttr zcl.AttrID = 8

func (LastSetTime) ID() zcl.AttrID                { return LastSetTimeAttr }
func (LastSetTime) Name() string                  { return "Last Set Time" }
func (LastSetTime) Readable() bool                { return true }
func (LastSetTime) Writable() bool                { return false }
func (LastSetTime) Reportable() bool              { return false }
func (LastSetTime) SceneIndex() int               { return -1 }
func (a *LastSetTime) Value() *LastSetTime        { return a }
func (a LastSetTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(a).MarshalZcl() }

func (a *LastSetTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = LastSetTime(*nt)
	return br, err
}

func (a *LastSetTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zutc); ok {
		*a = LastSetTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LastSetTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(a))
}

type Level zcl.Zu8

func (a *Level) Value() *Level              { return a }
func (a Level) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Level) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Level(*nt)
	return br, err
}

func (a *Level) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Level(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Level) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

type LevelDirection zcl.Zenum8

func (a *LevelDirection) Value() *LevelDirection     { return a }
func (a LevelDirection) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *LevelDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LevelDirection(*nt)
	return br, err
}

func (a *LevelDirection) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = LevelDirection(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LevelDirection) String() string {
	switch a {
	case 0x00:
		return "Up"
	case 0x01:
		return "Down"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a LevelDirection) IsUp() bool   { return a == 0x00 }
func (a LevelDirection) IsDown() bool { return a == 0x01 }
func (a *LevelDirection) SetUp()      { *a = 0x00 }
func (a *LevelDirection) SetDown()    { *a = 0x01 }

// LocalTime Local time
type LocalTime zcl.Zu32

const LocalTimeAttr zcl.AttrID = 7

func (LocalTime) ID() zcl.AttrID                { return LocalTimeAttr }
func (LocalTime) Name() string                  { return "Local Time" }
func (LocalTime) Readable() bool                { return true }
func (LocalTime) Writable() bool                { return false }
func (LocalTime) Reportable() bool              { return false }
func (LocalTime) SceneIndex() int               { return -1 }
func (a *LocalTime) Value() *LocalTime          { return a }
func (a LocalTime) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *LocalTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = LocalTime(*nt)
	return br, err
}

func (a *LocalTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = LocalTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocalTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type LocationAge zcl.Zu16

const LocationAgeAttr zcl.AttrID = 2

func (LocationAge) ID() zcl.AttrID                { return LocationAgeAttr }
func (LocationAge) Name() string                  { return "Location Age" }
func (LocationAge) Readable() bool                { return true }
func (LocationAge) Writable() bool                { return false }
func (LocationAge) Reportable() bool              { return false }
func (LocationAge) SceneIndex() int               { return -1 }
func (a *LocationAge) Value() *LocationAge        { return a }
func (a LocationAge) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LocationAge) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationAge(*nt)
	return br, err
}

func (a *LocationAge) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = LocationAge(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocationAge) String() string {
	return zcl.Seconds.Format(float64(a))
}

type LocationDescription zcl.Zcstring

const LocationDescriptionAttr zcl.AttrID = 16

func (LocationDescription) ID() zcl.AttrID                 { return LocationDescriptionAttr }
func (LocationDescription) Name() string                   { return "Location Description" }
func (LocationDescription) Readable() bool                 { return true }
func (LocationDescription) Writable() bool                 { return true }
func (LocationDescription) Reportable() bool               { return false }
func (LocationDescription) SceneIndex() int                { return -1 }
func (a *LocationDescription) Value() *LocationDescription { return a }
func (a LocationDescription) MarshalZcl() ([]byte, error)  { return zcl.Zcstring(a).MarshalZcl() }

func (a *LocationDescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationDescription(*nt)
	return br, err
}

func (a *LocationDescription) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = LocationDescription(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocationDescription) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type LocationMethod zcl.Zenum8

const LocationMethodAttr zcl.AttrID = 1

func (LocationMethod) ID() zcl.AttrID                { return LocationMethodAttr }
func (LocationMethod) Name() string                  { return "Location Method" }
func (LocationMethod) Readable() bool                { return true }
func (LocationMethod) Writable() bool                { return true }
func (LocationMethod) Reportable() bool              { return false }
func (LocationMethod) SceneIndex() int               { return -1 }
func (a *LocationMethod) Value() *LocationMethod     { return a }
func (a LocationMethod) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *LocationMethod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationMethod(*nt)
	return br, err
}

func (a *LocationMethod) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = LocationMethod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocationMethod) String() string {
	switch a {
	case 0x00:
		return "Lateration"
	case 0x01:
		return "Signposting"
	case 0x02:
		return "RF fingerprinting"
	case 0x03:
		return "Out of band"
	case 0x04:
		return "Centralized"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a LocationMethod) IsLateration() bool       { return a == 0x00 }
func (a LocationMethod) IsSignposting() bool      { return a == 0x01 }
func (a LocationMethod) IsRfFingerprinting() bool { return a == 0x02 }
func (a LocationMethod) IsOutOfBand() bool        { return a == 0x03 }
func (a LocationMethod) IsCentralized() bool      { return a == 0x04 }
func (a *LocationMethod) SetLateration()          { *a = 0x00 }
func (a *LocationMethod) SetSignposting()         { *a = 0x01 }
func (a *LocationMethod) SetRfFingerprinting()    { *a = 0x02 }
func (a *LocationMethod) SetOutOfBand()           { *a = 0x03 }
func (a *LocationMethod) SetCentralized()         { *a = 0x04 }

type LocationType zcl.Zenum8

const LocationTypeAttr zcl.AttrID = 0

func (LocationType) ID() zcl.AttrID                { return LocationTypeAttr }
func (LocationType) Name() string                  { return "Location Type" }
func (LocationType) Readable() bool                { return true }
func (LocationType) Writable() bool                { return true }
func (LocationType) Reportable() bool              { return false }
func (LocationType) SceneIndex() int               { return -1 }
func (a *LocationType) Value() *LocationType       { return a }
func (a LocationType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *LocationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationType(*nt)
	return br, err
}

func (a *LocationType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = LocationType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocationType) String() string {
	switch a {
	case 0x00:
		return "3D Location"
	case 0x01:
		return "Absolute 3D Location"
	case 0x02:
		return "2D Location"
	case 0x03:
		return "Absolute 2D Location"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a LocationType) Is3DLocation() bool         { return a == 0x00 }
func (a LocationType) IsAbsolute3DLocation() bool { return a == 0x01 }
func (a LocationType) Is2DLocation() bool         { return a == 0x02 }
func (a LocationType) IsAbsolute2DLocation() bool { return a == 0x03 }
func (a *LocationType) Set3DLocation()            { *a = 0x00 }
func (a *LocationType) SetAbsolute3DLocation()    { *a = 0x01 }
func (a *LocationType) Set2DLocation()            { *a = 0x02 }
func (a *LocationType) SetAbsolute2DLocation()    { *a = 0x03 }

type LocationFlags zcl.Zbmp8

func (a *LocationFlags) Value() *LocationFlags      { return a }
func (a LocationFlags) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *LocationFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = LocationFlags(*nt)
	return br, err
}

func (a *LocationFlags) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = LocationFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LocationFlags) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Absolute Only")
		case 1:
			bstr = append(bstr, "Recalculate")
		case 2:
			bstr = append(bstr, "Broadcast Indicator")
		case 3:
			bstr = append(bstr, "Broadcast Response")
		case 4:
			bstr = append(bstr, "Compact Response")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a LocationFlags) IsAbsoluteOnly() bool       { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a LocationFlags) IsRecalculate() bool        { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a LocationFlags) IsBroadcastIndicator() bool { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a LocationFlags) IsBroadcastResponse() bool  { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a LocationFlags) IsCompactResponse() bool    { return zcl.BitmapTest([]byte(a[:]), 4) }
func (a *LocationFlags) SetAbsoluteOnly(b bool)    { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *LocationFlags) SetRecalculate(b bool)     { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b)) }
func (a *LocationFlags) SetBroadcastIndicator(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *LocationFlags) SetBroadcastResponse(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b))
}
func (a *LocationFlags) SetCompactResponse(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 4, b))
}

type LowTempDwellTripPoint zcl.Zu24

const LowTempDwellTripPointAttr zcl.AttrID = 19

func (LowTempDwellTripPoint) ID() zcl.AttrID                   { return LowTempDwellTripPointAttr }
func (LowTempDwellTripPoint) Name() string                     { return "Low Temp Dwell Trip Point" }
func (LowTempDwellTripPoint) Readable() bool                   { return false }
func (LowTempDwellTripPoint) Writable() bool                   { return false }
func (LowTempDwellTripPoint) Reportable() bool                 { return false }
func (LowTempDwellTripPoint) SceneIndex() int                  { return -1 }
func (a *LowTempDwellTripPoint) Value() *LowTempDwellTripPoint { return a }
func (a LowTempDwellTripPoint) MarshalZcl() ([]byte, error)    { return zcl.Zu24(a).MarshalZcl() }

func (a *LowTempDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = LowTempDwellTripPoint(*nt)
	return br, err
}

func (a *LowTempDwellTripPoint) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu24); ok {
		*a = LowTempDwellTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LowTempDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(a))
}

// LowTempThreshold If the current temperature drops below the threshold for longer
// than the time specified by low temp dwell, an alarm will be triggered
type LowTempThreshold zcl.Zs16

const LowTempThresholdAttr zcl.AttrID = 17

func (LowTempThreshold) ID() zcl.AttrID                { return LowTempThresholdAttr }
func (LowTempThreshold) Name() string                  { return "Low Temp Threshold" }
func (LowTempThreshold) Readable() bool                { return true }
func (LowTempThreshold) Writable() bool                { return true }
func (LowTempThreshold) Reportable() bool              { return false }
func (LowTempThreshold) SceneIndex() int               { return -1 }
func (a *LowTempThreshold) Value() *LowTempThreshold   { return a }
func (a LowTempThreshold) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *LowTempThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = LowTempThreshold(*nt)
	return br, err
}

func (a *LowTempThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = LowTempThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LowTempThreshold) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

type MainsAlarmMask zcl.Zbmp8

const MainsAlarmMaskAttr zcl.AttrID = 16

func (MainsAlarmMask) ID() zcl.AttrID                { return MainsAlarmMaskAttr }
func (MainsAlarmMask) Name() string                  { return "Mains Alarm Mask" }
func (MainsAlarmMask) Readable() bool                { return true }
func (MainsAlarmMask) Writable() bool                { return true }
func (MainsAlarmMask) Reportable() bool              { return false }
func (MainsAlarmMask) SceneIndex() int               { return -1 }
func (a *MainsAlarmMask) Value() *MainsAlarmMask     { return a }
func (a MainsAlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *MainsAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsAlarmMask(*nt)
	return br, err
}

func (a *MainsAlarmMask) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = MainsAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Mains Voltage too low")
		case 1:
			bstr = append(bstr, "Mains Voltage too high")
		case 2:
			bstr = append(bstr, "Mains power supply lost/unavailable")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a MainsAlarmMask) IsMainsVoltageTooLow() bool  { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a MainsAlarmMask) IsMainsVoltageTooHigh() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a MainsAlarmMask) IsMainsPowerSupplyLostUnavailable() bool {
	return zcl.BitmapTest([]byte(a[:]), 2)
}
func (a *MainsAlarmMask) SetMainsVoltageTooLow(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *MainsAlarmMask) SetMainsVoltageTooHigh(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *MainsAlarmMask) SetMainsPowerSupplyLostUnavailable(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}

// MainsFrequency Resolution of 2Hz
// Special values:
// * 0x00 frequency too low to measure
// * 0xfe frequency too high to measure
// * 0xff unable to measure
type MainsFrequency zcl.Zu8

const MainsFrequencyAttr zcl.AttrID = 1

func (MainsFrequency) ID() zcl.AttrID                { return MainsFrequencyAttr }
func (MainsFrequency) Name() string                  { return "Mains Frequency" }
func (MainsFrequency) Readable() bool                { return true }
func (MainsFrequency) Writable() bool                { return false }
func (MainsFrequency) Reportable() bool              { return false }
func (MainsFrequency) SceneIndex() int               { return -1 }
func (a *MainsFrequency) Value() *MainsFrequency     { return a }
func (a MainsFrequency) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *MainsFrequency) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsFrequency(*nt)
	return br, err
}

func (a *MainsFrequency) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = MainsFrequency(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsFrequency) String() string {
	return zcl.Hertz.Format(float64(a) / 2)
}

type MainsVoltage zcl.Zu16

const MainsVoltageAttr zcl.AttrID = 0

func (MainsVoltage) ID() zcl.AttrID                { return MainsVoltageAttr }
func (MainsVoltage) Name() string                  { return "Mains Voltage" }
func (MainsVoltage) Readable() bool                { return true }
func (MainsVoltage) Writable() bool                { return false }
func (MainsVoltage) Reportable() bool              { return false }
func (MainsVoltage) SceneIndex() int               { return -1 }
func (a *MainsVoltage) Value() *MainsVoltage       { return a }
func (a MainsVoltage) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *MainsVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltage(*nt)
	return br, err
}

func (a *MainsVoltage) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MainsVoltage(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsVoltage) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

// MainsVoltageDwellTripPoint Length of time that the value of MainsVoltage MAY exist beyond either
// of its thresholds before an alarm is generated
type MainsVoltageDwellTripPoint zcl.Zu16

const MainsVoltageDwellTripPointAttr zcl.AttrID = 19

func (MainsVoltageDwellTripPoint) ID() zcl.AttrID                        { return MainsVoltageDwellTripPointAttr }
func (MainsVoltageDwellTripPoint) Name() string                          { return "Mains Voltage Dwell Trip Point" }
func (MainsVoltageDwellTripPoint) Readable() bool                        { return true }
func (MainsVoltageDwellTripPoint) Writable() bool                        { return true }
func (MainsVoltageDwellTripPoint) Reportable() bool                      { return false }
func (MainsVoltageDwellTripPoint) SceneIndex() int                       { return -1 }
func (a *MainsVoltageDwellTripPoint) Value() *MainsVoltageDwellTripPoint { return a }
func (a MainsVoltageDwellTripPoint) MarshalZcl() ([]byte, error)         { return zcl.Zu16(a).MarshalZcl() }

func (a *MainsVoltageDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltageDwellTripPoint(*nt)
	return br, err
}

func (a *MainsVoltageDwellTripPoint) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MainsVoltageDwellTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsVoltageDwellTripPoint) String() string {
	return zcl.Seconds.Format(float64(a))
}

type MainsVoltageMaxThreshold zcl.Zu16

const MainsVoltageMaxThresholdAttr zcl.AttrID = 18

func (MainsVoltageMaxThreshold) ID() zcl.AttrID                      { return MainsVoltageMaxThresholdAttr }
func (MainsVoltageMaxThreshold) Name() string                        { return "Mains Voltage Max Threshold" }
func (MainsVoltageMaxThreshold) Readable() bool                      { return true }
func (MainsVoltageMaxThreshold) Writable() bool                      { return true }
func (MainsVoltageMaxThreshold) Reportable() bool                    { return false }
func (MainsVoltageMaxThreshold) SceneIndex() int                     { return -1 }
func (a *MainsVoltageMaxThreshold) Value() *MainsVoltageMaxThreshold { return a }
func (a MainsVoltageMaxThreshold) MarshalZcl() ([]byte, error)       { return zcl.Zu16(a).MarshalZcl() }

func (a *MainsVoltageMaxThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltageMaxThreshold(*nt)
	return br, err
}

func (a *MainsVoltageMaxThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MainsVoltageMaxThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsVoltageMaxThreshold) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type MainsVoltageMinThreshold zcl.Zu16

const MainsVoltageMinThresholdAttr zcl.AttrID = 17

func (MainsVoltageMinThreshold) ID() zcl.AttrID                      { return MainsVoltageMinThresholdAttr }
func (MainsVoltageMinThreshold) Name() string                        { return "Mains Voltage Min Threshold" }
func (MainsVoltageMinThreshold) Readable() bool                      { return true }
func (MainsVoltageMinThreshold) Writable() bool                      { return true }
func (MainsVoltageMinThreshold) Reportable() bool                    { return false }
func (MainsVoltageMinThreshold) SceneIndex() int                     { return -1 }
func (a *MainsVoltageMinThreshold) Value() *MainsVoltageMinThreshold { return a }
func (a MainsVoltageMinThreshold) MarshalZcl() ([]byte, error)       { return zcl.Zu16(a).MarshalZcl() }

func (a *MainsVoltageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltageMinThreshold(*nt)
	return br, err
}

func (a *MainsVoltageMinThreshold) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = MainsVoltageMinThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MainsVoltageMinThreshold) String() string {
	return zcl.Volts.Format(float64(a) / 10)
}

type ManufacturerName zcl.Zcstring

const ManufacturerNameAttr zcl.AttrID = 4

func (ManufacturerName) ID() zcl.AttrID                { return ManufacturerNameAttr }
func (ManufacturerName) Name() string                  { return "Manufacturer Name" }
func (ManufacturerName) Readable() bool                { return true }
func (ManufacturerName) Writable() bool                { return false }
func (ManufacturerName) Reportable() bool              { return false }
func (ManufacturerName) SceneIndex() int               { return -1 }
func (a *ManufacturerName) Value() *ManufacturerName   { return a }
func (a ManufacturerName) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *ManufacturerName) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = ManufacturerName(*nt)
	return br, err
}

func (a *ManufacturerName) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = ManufacturerName(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ManufacturerName) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type MaxTempExperienced zcl.Zs16

const MaxTempExperiencedAttr zcl.AttrID = 2

func (MaxTempExperienced) ID() zcl.AttrID                { return MaxTempExperiencedAttr }
func (MaxTempExperienced) Name() string                  { return "Max Temp Experienced" }
func (MaxTempExperienced) Readable() bool                { return true }
func (MaxTempExperienced) Writable() bool                { return false }
func (MaxTempExperienced) Reportable() bool              { return false }
func (MaxTempExperienced) SceneIndex() int               { return -1 }
func (a *MaxTempExperienced) Value() *MaxTempExperienced { return a }
func (a MaxTempExperienced) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *MaxTempExperienced) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxTempExperienced(*nt)
	return br, err
}

func (a *MaxTempExperienced) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = MaxTempExperienced(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MaxTempExperienced) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

type MinTempExperienced zcl.Zs16

const MinTempExperiencedAttr zcl.AttrID = 1

func (MinTempExperienced) ID() zcl.AttrID                { return MinTempExperiencedAttr }
func (MinTempExperienced) Name() string                  { return "Min Temp Experienced" }
func (MinTempExperienced) Readable() bool                { return true }
func (MinTempExperienced) Writable() bool                { return false }
func (MinTempExperienced) Reportable() bool              { return false }
func (MinTempExperienced) SceneIndex() int               { return -1 }
func (a *MinTempExperienced) Value() *MinTempExperienced { return a }
func (a MinTempExperienced) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *MinTempExperienced) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinTempExperienced(*nt)
	return br, err
}

func (a *MinTempExperienced) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = MinTempExperienced(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MinTempExperienced) String() string {
	return zcl.DegreesCelsius.Format(float64(a))
}

type ModelIdentifier zcl.Zcstring

const ModelIdentifierAttr zcl.AttrID = 5

func (ModelIdentifier) ID() zcl.AttrID                { return ModelIdentifierAttr }
func (ModelIdentifier) Name() string                  { return "Model Identifier" }
func (ModelIdentifier) Readable() bool                { return true }
func (ModelIdentifier) Writable() bool                { return false }
func (ModelIdentifier) Reportable() bool              { return false }
func (ModelIdentifier) SceneIndex() int               { return -1 }
func (a *ModelIdentifier) Value() *ModelIdentifier    { return a }
func (a ModelIdentifier) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *ModelIdentifier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = ModelIdentifier(*nt)
	return br, err
}

func (a *ModelIdentifier) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = ModelIdentifier(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ModelIdentifier) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type NeighborsInfo zcl.Zset

func (a *NeighborsInfo) Value() *NeighborsInfo { return a }
func (a NeighborsInfo) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *NeighborsInfo) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zset)
	br, err := nt.UnmarshalZcl(b)
	*a = NeighborsInfo(*nt)
	return br, err
}

func (a *NeighborsInfo) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zset); ok {
		*a = NeighborsInfo(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NeighborsInfo) String() string {
	return zcl.Sprintf("%v", zcl.Zset(a))
}

// NumberRssiMeasurements is the number of measurements to use to generate one location estimate
type NumberRssiMeasurements zcl.Zu8

const NumberRssiMeasurementsAttr zcl.AttrID = 23

func (NumberRssiMeasurements) ID() zcl.AttrID                    { return NumberRssiMeasurementsAttr }
func (NumberRssiMeasurements) Name() string                      { return "Number RSSI Measurements" }
func (NumberRssiMeasurements) Readable() bool                    { return true }
func (NumberRssiMeasurements) Writable() bool                    { return true }
func (NumberRssiMeasurements) Reportable() bool                  { return false }
func (NumberRssiMeasurements) SceneIndex() int                   { return -1 }
func (a *NumberRssiMeasurements) Value() *NumberRssiMeasurements { return a }
func (a NumberRssiMeasurements) MarshalZcl() ([]byte, error)     { return zcl.Zu8(a).MarshalZcl() }

func (a *NumberRssiMeasurements) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberRssiMeasurements(*nt)
	return br, err
}

func (a *NumberRssiMeasurements) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = NumberRssiMeasurements(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NumberRssiMeasurements) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type NumberResponses zcl.Zu8

func (a *NumberResponses) Value() *NumberResponses    { return a }
func (a NumberResponses) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *NumberResponses) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberResponses(*nt)
	return br, err
}

func (a *NumberResponses) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = NumberResponses(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NumberResponses) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type NumberOfDevices zcl.Zu8

const NumberOfDevicesAttr zcl.AttrID = 4

func (NumberOfDevices) ID() zcl.AttrID                { return NumberOfDevicesAttr }
func (NumberOfDevices) Name() string                  { return "Number of Devices" }
func (NumberOfDevices) Readable() bool                { return true }
func (NumberOfDevices) Writable() bool                { return false }
func (NumberOfDevices) Reportable() bool              { return false }
func (NumberOfDevices) SceneIndex() int               { return -1 }
func (a *NumberOfDevices) Value() *NumberOfDevices    { return a }
func (a NumberOfDevices) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *NumberOfDevices) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfDevices(*nt)
	return br, err
}

func (a *NumberOfDevices) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = NumberOfDevices(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NumberOfDevices) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type OffTransitionTime zcl.Zu16

const OffTransitionTimeAttr zcl.AttrID = 19

func (OffTransitionTime) ID() zcl.AttrID                { return OffTransitionTimeAttr }
func (OffTransitionTime) Name() string                  { return "Off Transition Time" }
func (OffTransitionTime) Readable() bool                { return true }
func (OffTransitionTime) Writable() bool                { return true }
func (OffTransitionTime) Reportable() bool              { return false }
func (OffTransitionTime) SceneIndex() int               { return -1 }
func (a *OffTransitionTime) Value() *OffTransitionTime  { return a }
func (a OffTransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OffTransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OffTransitionTime(*nt)
	return br, err
}

func (a *OffTransitionTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OffTransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OffTransitionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type OffWaitTime zcl.Zu16

const OffWaitTimeAttr zcl.AttrID = 16386

func (OffWaitTime) ID() zcl.AttrID                { return OffWaitTimeAttr }
func (OffWaitTime) Name() string                  { return "Off Wait Time" }
func (OffWaitTime) Readable() bool                { return true }
func (OffWaitTime) Writable() bool                { return false }
func (OffWaitTime) Reportable() bool              { return false }
func (OffWaitTime) SceneIndex() int               { return -1 }
func (a *OffWaitTime) Value() *OffWaitTime        { return a }
func (a OffWaitTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OffWaitTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OffWaitTime(*nt)
	return br, err
}

func (a *OffWaitTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OffWaitTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OffWaitTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

// OnLevel determines the value that the CurrentLevel attribute is set to when the OnOff attribute of an On/Off cluster on the
// same endpoint is set to On. If the OnLevel attribute is not implemented, or is set to 0xff, it has no effect.
type OnLevel zcl.Zu8

const OnLevelAttr zcl.AttrID = 17

func (OnLevel) ID() zcl.AttrID                { return OnLevelAttr }
func (OnLevel) Name() string                  { return "On Level" }
func (OnLevel) Readable() bool                { return true }
func (OnLevel) Writable() bool                { return true }
func (OnLevel) Reportable() bool              { return false }
func (OnLevel) SceneIndex() int               { return -1 }
func (a *OnLevel) Value() *OnLevel            { return a }
func (a OnLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *OnLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = OnLevel(*nt)
	return br, err
}

func (a *OnLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = OnLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnLevel) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

type OnOff zcl.Zbool

const OnOffAttr zcl.AttrID = 0

func (OnOff) ID() zcl.AttrID                { return OnOffAttr }
func (OnOff) Name() string                  { return "On Off" }
func (OnOff) Readable() bool                { return true }
func (OnOff) Writable() bool                { return false }
func (OnOff) Reportable() bool              { return true }
func (OnOff) SceneIndex() int               { return 1 }
func (a *OnOff) Value() *OnOff              { return a }
func (a OnOff) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *OnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = OnOff(*nt)
	return br, err
}

func (a *OnOff) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = OnOff(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnOff) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x01:
		return "On"
	}
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

func (a OnOff) IsOff() bool { return a == 0x00 }
func (a OnOff) IsOn() bool  { return a == 0x01 }
func (a *OnOff) SetOff()    { *a = 0x00 }
func (a *OnOff) SetOn()     { *a = 0x01 }

type OnTime zcl.Zu16

const OnTimeAttr zcl.AttrID = 16385

func (OnTime) ID() zcl.AttrID                { return OnTimeAttr }
func (OnTime) Name() string                  { return "On Time" }
func (OnTime) Readable() bool                { return true }
func (OnTime) Writable() bool                { return false }
func (OnTime) Reportable() bool              { return false }
func (OnTime) SceneIndex() int               { return -1 }
func (a *OnTime) Value() *OnTime             { return a }
func (a OnTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OnTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OnTime(*nt)
	return br, err
}

func (a *OnTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OnTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type OnTransitionTime zcl.Zu16

const OnTransitionTimeAttr zcl.AttrID = 18

func (OnTransitionTime) ID() zcl.AttrID                { return OnTransitionTimeAttr }
func (OnTransitionTime) Name() string                  { return "On Transition Time" }
func (OnTransitionTime) Readable() bool                { return true }
func (OnTransitionTime) Writable() bool                { return true }
func (OnTransitionTime) Reportable() bool              { return false }
func (OnTransitionTime) SceneIndex() int               { return -1 }
func (a *OnTransitionTime) Value() *OnTransitionTime   { return a }
func (a OnTransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OnTransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OnTransitionTime(*nt)
	return br, err
}

func (a *OnTransitionTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OnTransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnTransitionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

// OnOffTransistionTime represents the time taken to move to or from the target level when On of Off commands are received
// by an On/Off cluster on the same endpoint.
// The actual time taken should be as close to OnOffTransitionTime as the device is able.
type OnOffTransistionTime zcl.Zu16

const OnOffTransistionTimeAttr zcl.AttrID = 16

func (OnOffTransistionTime) ID() zcl.AttrID                  { return OnOffTransistionTimeAttr }
func (OnOffTransistionTime) Name() string                    { return "On/Off Transistion Time" }
func (OnOffTransistionTime) Readable() bool                  { return true }
func (OnOffTransistionTime) Writable() bool                  { return true }
func (OnOffTransistionTime) Reportable() bool                { return false }
func (OnOffTransistionTime) SceneIndex() int                 { return -1 }
func (a *OnOffTransistionTime) Value() *OnOffTransistionTime { return a }
func (a OnOffTransistionTime) MarshalZcl() ([]byte, error)   { return zcl.Zu16(a).MarshalZcl() }

func (a *OnOffTransistionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OnOffTransistionTime(*nt)
	return br, err
}

func (a *OnOffTransistionTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OnOffTransistionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnOffTransistionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type OnOffControl zcl.Zbmp8

func (a *OnOffControl) Value() *OnOffControl       { return a }
func (a OnOffControl) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *OnOffControl) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = OnOffControl(*nt)
	return br, err
}

func (a *OnOffControl) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = OnOffControl(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OnOffControl) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Accept only when on")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a OnOffControl) IsAcceptOnlyWhenOn() bool { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a *OnOffControl) SetAcceptOnlyWhenOn(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}

// OverTempTotalDwell Total time the device has spent above the tmperature specified by High Temp Threshold
type OverTempTotalDwell zcl.Zu16

const OverTempTotalDwellAttr zcl.AttrID = 3

func (OverTempTotalDwell) ID() zcl.AttrID                { return OverTempTotalDwellAttr }
func (OverTempTotalDwell) Name() string                  { return "Over Temp Total Dwell" }
func (OverTempTotalDwell) Readable() bool                { return false }
func (OverTempTotalDwell) Writable() bool                { return false }
func (OverTempTotalDwell) Reportable() bool              { return false }
func (OverTempTotalDwell) SceneIndex() int               { return -1 }
func (a *OverTempTotalDwell) Value() *OverTempTotalDwell { return a }
func (a OverTempTotalDwell) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *OverTempTotalDwell) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OverTempTotalDwell(*nt)
	return br, err
}

func (a *OverTempTotalDwell) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = OverTempTotalDwell(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a OverTempTotalDwell) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// PathLossExponent is the rate at which the signal power decays with increasing distance
type PathLossExponent zcl.Zu16

const PathLossExponentAttr zcl.AttrID = 20

func (PathLossExponent) ID() zcl.AttrID                { return PathLossExponentAttr }
func (PathLossExponent) Name() string                  { return "Path loss Exponent" }
func (PathLossExponent) Readable() bool                { return true }
func (PathLossExponent) Writable() bool                { return true }
func (PathLossExponent) Reportable() bool              { return false }
func (PathLossExponent) SceneIndex() int               { return -1 }
func (a *PathLossExponent) Value() *PathLossExponent   { return a }
func (a PathLossExponent) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *PathLossExponent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PathLossExponent(*nt)
	return br, err
}

func (a *PathLossExponent) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = PathLossExponent(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PathLossExponent) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type PhysicalEnvironment zcl.Zenum8

const PhysicalEnvironmentAttr zcl.AttrID = 17

func (PhysicalEnvironment) ID() zcl.AttrID                 { return PhysicalEnvironmentAttr }
func (PhysicalEnvironment) Name() string                   { return "Physical Environment" }
func (PhysicalEnvironment) Readable() bool                 { return true }
func (PhysicalEnvironment) Writable() bool                 { return true }
func (PhysicalEnvironment) Reportable() bool               { return false }
func (PhysicalEnvironment) SceneIndex() int                { return -1 }
func (a *PhysicalEnvironment) Value() *PhysicalEnvironment { return a }
func (a PhysicalEnvironment) MarshalZcl() ([]byte, error)  { return zcl.Zenum8(a).MarshalZcl() }

func (a *PhysicalEnvironment) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PhysicalEnvironment(*nt)
	return br, err
}

func (a *PhysicalEnvironment) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = PhysicalEnvironment(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PhysicalEnvironment) String() string {
	switch a {
	case 0x00:
		return "Unspecified Environment"
	case 0x01:
		return "Atrium"
	case 0x02:
		return "Bar"
	case 0x03:
		return "Courtyard"
	case 0x04:
		return "Bathroom"
	case 0x05:
		return "Bedroom"
	case 0x06:
		return "Billiard Room"
	case 0x07:
		return "Utility Room"
	case 0x08:
		return "Cellar"
	case 0x09:
		return "Storage Closet"
	case 0x0a:
		return "Theater"
	case 0x0b:
		return "Office"
	case 0x0c:
		return "Deck"
	case 0x0d:
		return "Den"
	case 0x0e:
		return "Dining Room"
	case 0x0f:
		return "Electrical Room"
	case 0x10:
		return "Elevator"
	case 0x11:
		return "Entry"
	case 0x12:
		return "Family Room"
	case 0x13:
		return "Main Floor"
	case 0x14:
		return "Upstairs"
	case 0x15:
		return "Downstairs"
	case 0x16:
		return "Basement/Lower Level"
	case 0x17:
		return "Gallery"
	case 0x18:
		return "Game Room"
	case 0x19:
		return "Garage"
	case 0x1a:
		return "Gym"
	case 0x1b:
		return "Hallway"
	case 0x1c:
		return "House"
	case 0x1d:
		return "Kitchen"
	case 0x1e:
		return "Laundry Room"
	case 0x1f:
		return "Library"
	case 0x20:
		return "Master Bedroom"
	case 0x21:
		return "Mud Room (small room for coats and boots)"
	case 0x22:
		return "Nursery"
	case 0x23:
		return "Pantry"
	case 0x24:
		return "Office 2"
	case 0x25:
		return "Outside"
	case 0x26:
		return "Pool"
	case 0x27:
		return "Porch"
	case 0x28:
		return "Sewing Room"
	case 0x29:
		return "Sitting Room"
	case 0x2a:
		return "Stairway"
	case 0x2b:
		return "Yard"
	case 0x2c:
		return "Attic"
	case 0x2d:
		return "Hot Tub"
	case 0x2e:
		return "Living Room"
	case 0x2f:
		return "Sauna"
	case 0x30:
		return "Shop/Workshop"
	case 0x31:
		return "Guest Bedroom"
	case 0x32:
		return "Guest Bath"
	case 0x33:
		return "Powder Room (1/2 bath)"
	case 0x34:
		return "Back Yard"
	case 0x35:
		return "Front Yard"
	case 0x36:
		return "Patio"
	case 0x37:
		return "Driveway"
	case 0x38:
		return "Sun Room"
	case 0x39:
		return "Living Room 2"
	case 0x3a:
		return "Spa"
	case 0x3b:
		return "Whirlpool"
	case 0x3c:
		return "Shed"
	case 0x3d:
		return "Equipment Storage"
	case 0x3e:
		return "Hobby/Craft Room"
	case 0x3f:
		return "Fountain"
	case 0x40:
		return "Pond"
	case 0x41:
		return "Reception Room"
	case 0x42:
		return "Breakfast Room"
	case 0x43:
		return "Nook"
	case 0x44:
		return "Garden"
	case 0x45:
		return "Balcony"
	case 0x46:
		return "Panic Room"
	case 0x47:
		return "Terrace"
	case 0x48:
		return "Roof"
	case 0x49:
		return "Toilet"
	case 0x4a:
		return "Toilet Main"
	case 0x4b:
		return "Outside Toilet"
	case 0x4c:
		return "Shower room"
	case 0x4d:
		return "Study"
	case 0x4e:
		return "Front Garden"
	case 0x4f:
		return "Back Garden"
	case 0x50:
		return "Kettle"
	case 0x51:
		return "Television"
	case 0x52:
		return "Stove"
	case 0x53:
		return "Microwave"
	case 0x54:
		return "Toaster"
	case 0x55:
		return "Vacuum"
	case 0x56:
		return "Appliance"
	case 0x57:
		return "Front Door"
	case 0x58:
		return "Back Door"
	case 0x59:
		return "Fridge Door"
	case 0x60:
		return "Medication Cabinet Door"
	case 0x61:
		return "Wardrobe Door"
	case 0x62:
		return "Front Cupboard Door"
	case 0x63:
		return "Other Door"
	case 0x64:
		return "Waiting Room"
	case 0x65:
		return "Triage Room"
	case 0x66:
		return "Doctor’s Office"
	case 0x67:
		return "Patient’s Private Room"
	case 0x68:
		return "Consultation Room"
	case 0x69:
		return "Nurse Station"
	case 0x6a:
		return "Ward"
	case 0x6b:
		return "Corridor"
	case 0x6c:
		return "Operating Theatre"
	case 0x6d:
		return "Dental Surgery Room"
	case 0x6e:
		return "Medical Imaging Room"
	case 0x6f:
		return "Decontamination Room"
	case 0xFF:
		return "Unknown Environment"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a PhysicalEnvironment) IsUnspecifiedEnvironment() bool           { return a == 0x00 }
func (a PhysicalEnvironment) IsAtrium() bool                           { return a == 0x01 }
func (a PhysicalEnvironment) IsBar() bool                              { return a == 0x02 }
func (a PhysicalEnvironment) IsCourtyard() bool                        { return a == 0x03 }
func (a PhysicalEnvironment) IsBathroom() bool                         { return a == 0x04 }
func (a PhysicalEnvironment) IsBedroom() bool                          { return a == 0x05 }
func (a PhysicalEnvironment) IsBilliardRoom() bool                     { return a == 0x06 }
func (a PhysicalEnvironment) IsUtilityRoom() bool                      { return a == 0x07 }
func (a PhysicalEnvironment) IsCellar() bool                           { return a == 0x08 }
func (a PhysicalEnvironment) IsStorageCloset() bool                    { return a == 0x09 }
func (a PhysicalEnvironment) IsTheater() bool                          { return a == 0x0a }
func (a PhysicalEnvironment) IsOffice() bool                           { return a == 0x0b }
func (a PhysicalEnvironment) IsDeck() bool                             { return a == 0x0c }
func (a PhysicalEnvironment) IsDen() bool                              { return a == 0x0d }
func (a PhysicalEnvironment) IsDiningRoom() bool                       { return a == 0x0e }
func (a PhysicalEnvironment) IsElectricalRoom() bool                   { return a == 0x0f }
func (a PhysicalEnvironment) IsElevator() bool                         { return a == 0x10 }
func (a PhysicalEnvironment) IsEntry() bool                            { return a == 0x11 }
func (a PhysicalEnvironment) IsFamilyRoom() bool                       { return a == 0x12 }
func (a PhysicalEnvironment) IsMainFloor() bool                        { return a == 0x13 }
func (a PhysicalEnvironment) IsUpstairs() bool                         { return a == 0x14 }
func (a PhysicalEnvironment) IsDownstairs() bool                       { return a == 0x15 }
func (a PhysicalEnvironment) IsBasementLowerLevel() bool               { return a == 0x16 }
func (a PhysicalEnvironment) IsGallery() bool                          { return a == 0x17 }
func (a PhysicalEnvironment) IsGameRoom() bool                         { return a == 0x18 }
func (a PhysicalEnvironment) IsGarage() bool                           { return a == 0x19 }
func (a PhysicalEnvironment) IsGym() bool                              { return a == 0x1a }
func (a PhysicalEnvironment) IsHallway() bool                          { return a == 0x1b }
func (a PhysicalEnvironment) IsHouse() bool                            { return a == 0x1c }
func (a PhysicalEnvironment) IsKitchen() bool                          { return a == 0x1d }
func (a PhysicalEnvironment) IsLaundryRoom() bool                      { return a == 0x1e }
func (a PhysicalEnvironment) IsLibrary() bool                          { return a == 0x1f }
func (a PhysicalEnvironment) IsMasterBedroom() bool                    { return a == 0x20 }
func (a PhysicalEnvironment) IsMudRoomSmallRoomForCoatsAndBoots() bool { return a == 0x21 }
func (a PhysicalEnvironment) IsNursery() bool                          { return a == 0x22 }
func (a PhysicalEnvironment) IsPantry() bool                           { return a == 0x23 }
func (a PhysicalEnvironment) IsOffice2() bool                          { return a == 0x24 }
func (a PhysicalEnvironment) IsOutside() bool                          { return a == 0x25 }
func (a PhysicalEnvironment) IsPool() bool                             { return a == 0x26 }
func (a PhysicalEnvironment) IsPorch() bool                            { return a == 0x27 }
func (a PhysicalEnvironment) IsSewingRoom() bool                       { return a == 0x28 }
func (a PhysicalEnvironment) IsSittingRoom() bool                      { return a == 0x29 }
func (a PhysicalEnvironment) IsStairway() bool                         { return a == 0x2a }
func (a PhysicalEnvironment) IsYard() bool                             { return a == 0x2b }
func (a PhysicalEnvironment) IsAttic() bool                            { return a == 0x2c }
func (a PhysicalEnvironment) IsHotTub() bool                           { return a == 0x2d }
func (a PhysicalEnvironment) IsLivingRoom() bool                       { return a == 0x2e }
func (a PhysicalEnvironment) IsSauna() bool                            { return a == 0x2f }
func (a PhysicalEnvironment) IsShopWorkshop() bool                     { return a == 0x30 }
func (a PhysicalEnvironment) IsGuestBedroom() bool                     { return a == 0x31 }
func (a PhysicalEnvironment) IsGuestBath() bool                        { return a == 0x32 }
func (a PhysicalEnvironment) IsPowderRoom12Bath() bool                 { return a == 0x33 }
func (a PhysicalEnvironment) IsBackYard() bool                         { return a == 0x34 }
func (a PhysicalEnvironment) IsFrontYard() bool                        { return a == 0x35 }
func (a PhysicalEnvironment) IsPatio() bool                            { return a == 0x36 }
func (a PhysicalEnvironment) IsDriveway() bool                         { return a == 0x37 }
func (a PhysicalEnvironment) IsSunRoom() bool                          { return a == 0x38 }
func (a PhysicalEnvironment) IsLivingRoom2() bool                      { return a == 0x39 }
func (a PhysicalEnvironment) IsSpa() bool                              { return a == 0x3a }
func (a PhysicalEnvironment) IsWhirlpool() bool                        { return a == 0x3b }
func (a PhysicalEnvironment) IsShed() bool                             { return a == 0x3c }
func (a PhysicalEnvironment) IsEquipmentStorage() bool                 { return a == 0x3d }
func (a PhysicalEnvironment) IsHobbyCraftRoom() bool                   { return a == 0x3e }
func (a PhysicalEnvironment) IsFountain() bool                         { return a == 0x3f }
func (a PhysicalEnvironment) IsPond() bool                             { return a == 0x40 }
func (a PhysicalEnvironment) IsReceptionRoom() bool                    { return a == 0x41 }
func (a PhysicalEnvironment) IsBreakfastRoom() bool                    { return a == 0x42 }
func (a PhysicalEnvironment) IsNook() bool                             { return a == 0x43 }
func (a PhysicalEnvironment) IsGarden() bool                           { return a == 0x44 }
func (a PhysicalEnvironment) IsBalcony() bool                          { return a == 0x45 }
func (a PhysicalEnvironment) IsPanicRoom() bool                        { return a == 0x46 }
func (a PhysicalEnvironment) IsTerrace() bool                          { return a == 0x47 }
func (a PhysicalEnvironment) IsRoof() bool                             { return a == 0x48 }
func (a PhysicalEnvironment) IsToilet() bool                           { return a == 0x49 }
func (a PhysicalEnvironment) IsToiletMain() bool                       { return a == 0x4a }
func (a PhysicalEnvironment) IsOutsideToilet() bool                    { return a == 0x4b }
func (a PhysicalEnvironment) IsShowerRoom() bool                       { return a == 0x4c }
func (a PhysicalEnvironment) IsStudy() bool                            { return a == 0x4d }
func (a PhysicalEnvironment) IsFrontGarden() bool                      { return a == 0x4e }
func (a PhysicalEnvironment) IsBackGarden() bool                       { return a == 0x4f }
func (a PhysicalEnvironment) IsKettle() bool                           { return a == 0x50 }
func (a PhysicalEnvironment) IsTelevision() bool                       { return a == 0x51 }
func (a PhysicalEnvironment) IsStove() bool                            { return a == 0x52 }
func (a PhysicalEnvironment) IsMicrowave() bool                        { return a == 0x53 }
func (a PhysicalEnvironment) IsToaster() bool                          { return a == 0x54 }
func (a PhysicalEnvironment) IsVacuum() bool                           { return a == 0x55 }
func (a PhysicalEnvironment) IsAppliance() bool                        { return a == 0x56 }
func (a PhysicalEnvironment) IsFrontDoor() bool                        { return a == 0x57 }
func (a PhysicalEnvironment) IsBackDoor() bool                         { return a == 0x58 }
func (a PhysicalEnvironment) IsFridgeDoor() bool                       { return a == 0x59 }
func (a PhysicalEnvironment) IsMedicationCabinetDoor() bool            { return a == 0x60 }
func (a PhysicalEnvironment) IsWardrobeDoor() bool                     { return a == 0x61 }
func (a PhysicalEnvironment) IsFrontCupboardDoor() bool                { return a == 0x62 }
func (a PhysicalEnvironment) IsOtherDoor() bool                        { return a == 0x63 }
func (a PhysicalEnvironment) IsWaitingRoom() bool                      { return a == 0x64 }
func (a PhysicalEnvironment) IsTriageRoom() bool                       { return a == 0x65 }
func (a PhysicalEnvironment) IsDoctorSOffice() bool                    { return a == 0x66 }
func (a PhysicalEnvironment) IsPatientSPrivateRoom() bool              { return a == 0x67 }
func (a PhysicalEnvironment) IsConsultationRoom() bool                 { return a == 0x68 }
func (a PhysicalEnvironment) IsNurseStation() bool                     { return a == 0x69 }
func (a PhysicalEnvironment) IsWard() bool                             { return a == 0x6a }
func (a PhysicalEnvironment) IsCorridor() bool                         { return a == 0x6b }
func (a PhysicalEnvironment) IsOperatingTheatre() bool                 { return a == 0x6c }
func (a PhysicalEnvironment) IsDentalSurgeryRoom() bool                { return a == 0x6d }
func (a PhysicalEnvironment) IsMedicalImagingRoom() bool               { return a == 0x6e }
func (a PhysicalEnvironment) IsDecontaminationRoom() bool              { return a == 0x6f }
func (a PhysicalEnvironment) IsUnknownEnvironment() bool               { return a == 0xFF }
func (a *PhysicalEnvironment) SetUnspecifiedEnvironment()              { *a = 0x00 }
func (a *PhysicalEnvironment) SetAtrium()                              { *a = 0x01 }
func (a *PhysicalEnvironment) SetBar()                                 { *a = 0x02 }
func (a *PhysicalEnvironment) SetCourtyard()                           { *a = 0x03 }
func (a *PhysicalEnvironment) SetBathroom()                            { *a = 0x04 }
func (a *PhysicalEnvironment) SetBedroom()                             { *a = 0x05 }
func (a *PhysicalEnvironment) SetBilliardRoom()                        { *a = 0x06 }
func (a *PhysicalEnvironment) SetUtilityRoom()                         { *a = 0x07 }
func (a *PhysicalEnvironment) SetCellar()                              { *a = 0x08 }
func (a *PhysicalEnvironment) SetStorageCloset()                       { *a = 0x09 }
func (a *PhysicalEnvironment) SetTheater()                             { *a = 0x0a }
func (a *PhysicalEnvironment) SetOffice()                              { *a = 0x0b }
func (a *PhysicalEnvironment) SetDeck()                                { *a = 0x0c }
func (a *PhysicalEnvironment) SetDen()                                 { *a = 0x0d }
func (a *PhysicalEnvironment) SetDiningRoom()                          { *a = 0x0e }
func (a *PhysicalEnvironment) SetElectricalRoom()                      { *a = 0x0f }
func (a *PhysicalEnvironment) SetElevator()                            { *a = 0x10 }
func (a *PhysicalEnvironment) SetEntry()                               { *a = 0x11 }
func (a *PhysicalEnvironment) SetFamilyRoom()                          { *a = 0x12 }
func (a *PhysicalEnvironment) SetMainFloor()                           { *a = 0x13 }
func (a *PhysicalEnvironment) SetUpstairs()                            { *a = 0x14 }
func (a *PhysicalEnvironment) SetDownstairs()                          { *a = 0x15 }
func (a *PhysicalEnvironment) SetBasementLowerLevel()                  { *a = 0x16 }
func (a *PhysicalEnvironment) SetGallery()                             { *a = 0x17 }
func (a *PhysicalEnvironment) SetGameRoom()                            { *a = 0x18 }
func (a *PhysicalEnvironment) SetGarage()                              { *a = 0x19 }
func (a *PhysicalEnvironment) SetGym()                                 { *a = 0x1a }
func (a *PhysicalEnvironment) SetHallway()                             { *a = 0x1b }
func (a *PhysicalEnvironment) SetHouse()                               { *a = 0x1c }
func (a *PhysicalEnvironment) SetKitchen()                             { *a = 0x1d }
func (a *PhysicalEnvironment) SetLaundryRoom()                         { *a = 0x1e }
func (a *PhysicalEnvironment) SetLibrary()                             { *a = 0x1f }
func (a *PhysicalEnvironment) SetMasterBedroom()                       { *a = 0x20 }
func (a *PhysicalEnvironment) SetMudRoomSmallRoomForCoatsAndBoots()    { *a = 0x21 }
func (a *PhysicalEnvironment) SetNursery()                             { *a = 0x22 }
func (a *PhysicalEnvironment) SetPantry()                              { *a = 0x23 }
func (a *PhysicalEnvironment) SetOffice2()                             { *a = 0x24 }
func (a *PhysicalEnvironment) SetOutside()                             { *a = 0x25 }
func (a *PhysicalEnvironment) SetPool()                                { *a = 0x26 }
func (a *PhysicalEnvironment) SetPorch()                               { *a = 0x27 }
func (a *PhysicalEnvironment) SetSewingRoom()                          { *a = 0x28 }
func (a *PhysicalEnvironment) SetSittingRoom()                         { *a = 0x29 }
func (a *PhysicalEnvironment) SetStairway()                            { *a = 0x2a }
func (a *PhysicalEnvironment) SetYard()                                { *a = 0x2b }
func (a *PhysicalEnvironment) SetAttic()                               { *a = 0x2c }
func (a *PhysicalEnvironment) SetHotTub()                              { *a = 0x2d }
func (a *PhysicalEnvironment) SetLivingRoom()                          { *a = 0x2e }
func (a *PhysicalEnvironment) SetSauna()                               { *a = 0x2f }
func (a *PhysicalEnvironment) SetShopWorkshop()                        { *a = 0x30 }
func (a *PhysicalEnvironment) SetGuestBedroom()                        { *a = 0x31 }
func (a *PhysicalEnvironment) SetGuestBath()                           { *a = 0x32 }
func (a *PhysicalEnvironment) SetPowderRoom12Bath()                    { *a = 0x33 }
func (a *PhysicalEnvironment) SetBackYard()                            { *a = 0x34 }
func (a *PhysicalEnvironment) SetFrontYard()                           { *a = 0x35 }
func (a *PhysicalEnvironment) SetPatio()                               { *a = 0x36 }
func (a *PhysicalEnvironment) SetDriveway()                            { *a = 0x37 }
func (a *PhysicalEnvironment) SetSunRoom()                             { *a = 0x38 }
func (a *PhysicalEnvironment) SetLivingRoom2()                         { *a = 0x39 }
func (a *PhysicalEnvironment) SetSpa()                                 { *a = 0x3a }
func (a *PhysicalEnvironment) SetWhirlpool()                           { *a = 0x3b }
func (a *PhysicalEnvironment) SetShed()                                { *a = 0x3c }
func (a *PhysicalEnvironment) SetEquipmentStorage()                    { *a = 0x3d }
func (a *PhysicalEnvironment) SetHobbyCraftRoom()                      { *a = 0x3e }
func (a *PhysicalEnvironment) SetFountain()                            { *a = 0x3f }
func (a *PhysicalEnvironment) SetPond()                                { *a = 0x40 }
func (a *PhysicalEnvironment) SetReceptionRoom()                       { *a = 0x41 }
func (a *PhysicalEnvironment) SetBreakfastRoom()                       { *a = 0x42 }
func (a *PhysicalEnvironment) SetNook()                                { *a = 0x43 }
func (a *PhysicalEnvironment) SetGarden()                              { *a = 0x44 }
func (a *PhysicalEnvironment) SetBalcony()                             { *a = 0x45 }
func (a *PhysicalEnvironment) SetPanicRoom()                           { *a = 0x46 }
func (a *PhysicalEnvironment) SetTerrace()                             { *a = 0x47 }
func (a *PhysicalEnvironment) SetRoof()                                { *a = 0x48 }
func (a *PhysicalEnvironment) SetToilet()                              { *a = 0x49 }
func (a *PhysicalEnvironment) SetToiletMain()                          { *a = 0x4a }
func (a *PhysicalEnvironment) SetOutsideToilet()                       { *a = 0x4b }
func (a *PhysicalEnvironment) SetShowerRoom()                          { *a = 0x4c }
func (a *PhysicalEnvironment) SetStudy()                               { *a = 0x4d }
func (a *PhysicalEnvironment) SetFrontGarden()                         { *a = 0x4e }
func (a *PhysicalEnvironment) SetBackGarden()                          { *a = 0x4f }
func (a *PhysicalEnvironment) SetKettle()                              { *a = 0x50 }
func (a *PhysicalEnvironment) SetTelevision()                          { *a = 0x51 }
func (a *PhysicalEnvironment) SetStove()                               { *a = 0x52 }
func (a *PhysicalEnvironment) SetMicrowave()                           { *a = 0x53 }
func (a *PhysicalEnvironment) SetToaster()                             { *a = 0x54 }
func (a *PhysicalEnvironment) SetVacuum()                              { *a = 0x55 }
func (a *PhysicalEnvironment) SetAppliance()                           { *a = 0x56 }
func (a *PhysicalEnvironment) SetFrontDoor()                           { *a = 0x57 }
func (a *PhysicalEnvironment) SetBackDoor()                            { *a = 0x58 }
func (a *PhysicalEnvironment) SetFridgeDoor()                          { *a = 0x59 }
func (a *PhysicalEnvironment) SetMedicationCabinetDoor()               { *a = 0x60 }
func (a *PhysicalEnvironment) SetWardrobeDoor()                        { *a = 0x61 }
func (a *PhysicalEnvironment) SetFrontCupboardDoor()                   { *a = 0x62 }
func (a *PhysicalEnvironment) SetOtherDoor()                           { *a = 0x63 }
func (a *PhysicalEnvironment) SetWaitingRoom()                         { *a = 0x64 }
func (a *PhysicalEnvironment) SetTriageRoom()                          { *a = 0x65 }
func (a *PhysicalEnvironment) SetDoctorSOffice()                       { *a = 0x66 }
func (a *PhysicalEnvironment) SetPatientSPrivateRoom()                 { *a = 0x67 }
func (a *PhysicalEnvironment) SetConsultationRoom()                    { *a = 0x68 }
func (a *PhysicalEnvironment) SetNurseStation()                        { *a = 0x69 }
func (a *PhysicalEnvironment) SetWard()                                { *a = 0x6a }
func (a *PhysicalEnvironment) SetCorridor()                            { *a = 0x6b }
func (a *PhysicalEnvironment) SetOperatingTheatre()                    { *a = 0x6c }
func (a *PhysicalEnvironment) SetDentalSurgeryRoom()                   { *a = 0x6d }
func (a *PhysicalEnvironment) SetMedicalImagingRoom()                  { *a = 0x6e }
func (a *PhysicalEnvironment) SetDecontaminationRoom()                 { *a = 0x6f }
func (a *PhysicalEnvironment) SetUnknownEnvironment()                  { *a = 0xFF }

type Power zcl.Zs16

const PowerAttr zcl.AttrID = 19

func (Power) ID() zcl.AttrID                { return PowerAttr }
func (Power) Name() string                  { return "Power" }
func (Power) Readable() bool                { return true }
func (Power) Writable() bool                { return true }
func (Power) Reportable() bool              { return false }
func (Power) SceneIndex() int               { return -1 }
func (a *Power) Value() *Power              { return a }
func (a Power) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *Power) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Power(*nt)
	return br, err
}

func (a *Power) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = Power(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Power) String() string {
	return zcl.DecibelMilliWatts.Format(float64(a) / 100)
}

type PowerOnLevel zcl.Zu8

const PowerOnLevelAttr zcl.AttrID = 16384

func (PowerOnLevel) ID() zcl.AttrID                { return PowerOnLevelAttr }
func (PowerOnLevel) Name() string                  { return "Power On level" }
func (PowerOnLevel) Readable() bool                { return true }
func (PowerOnLevel) Writable() bool                { return true }
func (PowerOnLevel) Reportable() bool              { return false }
func (PowerOnLevel) SceneIndex() int               { return -1 }
func (a *PowerOnLevel) Value() *PowerOnLevel       { return a }
func (a PowerOnLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *PowerOnLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerOnLevel(*nt)
	return br, err
}

func (a *PowerOnLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = PowerOnLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerOnLevel) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

type PowerSource zcl.Zenum8

const PowerSourceAttr zcl.AttrID = 7

func (PowerSource) ID() zcl.AttrID                { return PowerSourceAttr }
func (PowerSource) Name() string                  { return "Power Source" }
func (PowerSource) Readable() bool                { return true }
func (PowerSource) Writable() bool                { return false }
func (PowerSource) Reportable() bool              { return false }
func (PowerSource) SceneIndex() int               { return -1 }
func (a *PowerSource) Value() *PowerSource        { return a }
func (a PowerSource) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *PowerSource) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerSource(*nt)
	return br, err
}

func (a *PowerSource) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = PowerSource(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerSource) String() string {
	switch a {
	case 0x00:
		return "Unknown"
	case 0x01:
		return "Mains (single phase)"
	case 0x02:
		return "Mains (3 phase)"
	case 0x03:
		return "Battery"
	case 0x04:
		return "DC Source"
	case 0x05:
		return "Emergency mains constantly powered"
	case 0x06:
		return "Emergency mains and transfer switch"
	case 0x80:
		return "Unknown with battery backup"
	case 0x81:
		return "Mains (single phase) with battery backup"
	case 0x82:
		return "Mains (3 phase) with battery backup"
	case 0x83:
		return "Battery with battery backup"
	case 0x84:
		return "DC Source with battery backup"
	case 0x85:
		return "Emergency mains constantly powered with battery backup"
	case 0x86:
		return "Emergency mains and transfer switch with battery backup"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a PowerSource) IsUnknown() bool                                          { return a == 0x00 }
func (a PowerSource) IsMainsSinglePhase() bool                                 { return a == 0x01 }
func (a PowerSource) IsMains3Phase() bool                                      { return a == 0x02 }
func (a PowerSource) IsBattery() bool                                          { return a == 0x03 }
func (a PowerSource) IsDcSource() bool                                         { return a == 0x04 }
func (a PowerSource) IsEmergencyMainsConstantlyPowered() bool                  { return a == 0x05 }
func (a PowerSource) IsEmergencyMainsAndTransferSwitch() bool                  { return a == 0x06 }
func (a PowerSource) IsUnknownWithBatteryBackup() bool                         { return a == 0x80 }
func (a PowerSource) IsMainsSinglePhaseWithBatteryBackup() bool                { return a == 0x81 }
func (a PowerSource) IsMains3PhaseWithBatteryBackup() bool                     { return a == 0x82 }
func (a PowerSource) IsBatteryWithBatteryBackup() bool                         { return a == 0x83 }
func (a PowerSource) IsDcSourceWithBatteryBackup() bool                        { return a == 0x84 }
func (a PowerSource) IsEmergencyMainsConstantlyPoweredWithBatteryBackup() bool { return a == 0x85 }
func (a PowerSource) IsEmergencyMainsAndTransferSwitchWithBatteryBackup() bool { return a == 0x86 }
func (a *PowerSource) SetUnknown()                                             { *a = 0x00 }
func (a *PowerSource) SetMainsSinglePhase()                                    { *a = 0x01 }
func (a *PowerSource) SetMains3Phase()                                         { *a = 0x02 }
func (a *PowerSource) SetBattery()                                             { *a = 0x03 }
func (a *PowerSource) SetDcSource()                                            { *a = 0x04 }
func (a *PowerSource) SetEmergencyMainsConstantlyPowered()                     { *a = 0x05 }
func (a *PowerSource) SetEmergencyMainsAndTransferSwitch()                     { *a = 0x06 }
func (a *PowerSource) SetUnknownWithBatteryBackup()                            { *a = 0x80 }
func (a *PowerSource) SetMainsSinglePhaseWithBatteryBackup()                   { *a = 0x81 }
func (a *PowerSource) SetMains3PhaseWithBatteryBackup()                        { *a = 0x82 }
func (a *PowerSource) SetBatteryWithBatteryBackup()                            { *a = 0x83 }
func (a *PowerSource) SetDcSourceWithBatteryBackup()                           { *a = 0x84 }
func (a *PowerSource) SetEmergencyMainsConstantlyPoweredWithBatteryBackup()    { *a = 0x85 }
func (a *PowerSource) SetEmergencyMainsAndTransferSwitchWithBatteryBackup()    { *a = 0x86 }

type PoweronOnOff zcl.Zenum8

const PoweronOnOffAttr zcl.AttrID = 16387

func (PoweronOnOff) ID() zcl.AttrID                { return PoweronOnOffAttr }
func (PoweronOnOff) Name() string                  { return "PowerOn On/Off" }
func (PoweronOnOff) Readable() bool                { return true }
func (PoweronOnOff) Writable() bool                { return true }
func (PoweronOnOff) Reportable() bool              { return false }
func (PoweronOnOff) SceneIndex() int               { return -1 }
func (a *PoweronOnOff) Value() *PoweronOnOff       { return a }
func (a PoweronOnOff) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *PoweronOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = PoweronOnOff(*nt)
	return br, err
}

func (a *PoweronOnOff) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = PoweronOnOff(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PoweronOnOff) String() string {
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

func (a PoweronOnOff) IsOff() bool      { return a == 0x00 }
func (a PoweronOnOff) IsOn() bool       { return a == 0x01 }
func (a PoweronOnOff) IsPrevious() bool { return a == 0xFF }
func (a *PoweronOnOff) SetOff()         { *a = 0x00 }
func (a *PoweronOnOff) SetOn()          { *a = 0x01 }
func (a *PoweronOnOff) SetPrevious()    { *a = 0xFF }

// ProductCode As printed on the product.
type ProductCode zcl.Zostring

const ProductCodeAttr zcl.AttrID = 10

func (ProductCode) ID() zcl.AttrID                { return ProductCodeAttr }
func (ProductCode) Name() string                  { return "Product code" }
func (ProductCode) Readable() bool                { return true }
func (ProductCode) Writable() bool                { return false }
func (ProductCode) Reportable() bool              { return false }
func (ProductCode) SceneIndex() int               { return -1 }
func (a *ProductCode) Value() *ProductCode        { return a }
func (a ProductCode) MarshalZcl() ([]byte, error) { return zcl.Zostring(a).MarshalZcl() }

func (a *ProductCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zostring)
	br, err := nt.UnmarshalZcl(b)
	*a = ProductCode(*nt)
	return br, err
}

func (a *ProductCode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zostring); ok {
		*a = ProductCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ProductCode) String() string {
	return zcl.Sprintf("%v", zcl.Zostring(a))
}

type QualityMeasure zcl.Zu8

const QualityMeasureAttr zcl.AttrID = 3

func (QualityMeasure) ID() zcl.AttrID                { return QualityMeasureAttr }
func (QualityMeasure) Name() string                  { return "Quality Measure" }
func (QualityMeasure) Readable() bool                { return true }
func (QualityMeasure) Writable() bool                { return false }
func (QualityMeasure) Reportable() bool              { return false }
func (QualityMeasure) SceneIndex() int               { return -1 }
func (a *QualityMeasure) Value() *QualityMeasure     { return a }
func (a QualityMeasure) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *QualityMeasure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = QualityMeasure(*nt)
	return br, err
}

func (a *QualityMeasure) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = QualityMeasure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a QualityMeasure) String() string {
	return zcl.Percent.Format(float64(a))
}

type QualityIndex zcl.Zu16

func (a *QualityIndex) Value() *QualityIndex       { return a }
func (a QualityIndex) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *QualityIndex) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = QualityIndex(*nt)
	return br, err
}

func (a *QualityIndex) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = QualityIndex(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a QualityIndex) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type Rssi zcl.Zs8

func (a *Rssi) Value() *Rssi               { return a }
func (a Rssi) MarshalZcl() ([]byte, error) { return zcl.Zs8(a).MarshalZcl() }

func (a *Rssi) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = Rssi(*nt)
	return br, err
}

func (a *Rssi) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs8); ok {
		*a = Rssi(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Rssi) String() string {
	return zcl.DecibelMilliWatts.Format(float64(a))
}

type Rate zcl.Zu8

func (a *Rate) Value() *Rate               { return a }
func (a Rate) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Rate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Rate(*nt)
	return br, err
}

func (a *Rate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Rate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Rate) String() string {
	return zcl.PercentPerSecond.Format(float64(a) / 2.54)
}

// RemainingTime represents the time remaining until the current command is complete - it is specified in 1/10ths of a second.
type RemainingTime zcl.Zu16

const RemainingTimeAttr zcl.AttrID = 1

func (RemainingTime) ID() zcl.AttrID                { return RemainingTimeAttr }
func (RemainingTime) Name() string                  { return "Remaining Time" }
func (RemainingTime) Readable() bool                { return true }
func (RemainingTime) Writable() bool                { return false }
func (RemainingTime) Reportable() bool              { return false }
func (RemainingTime) SceneIndex() int               { return -1 }
func (a *RemainingTime) Value() *RemainingTime      { return a }
func (a RemainingTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RemainingTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RemainingTime(*nt)
	return br, err
}

func (a *RemainingTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = RemainingTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a RemainingTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type ReportingPeriod zcl.Zu16

const ReportingPeriodAttr zcl.AttrID = 21

func (ReportingPeriod) ID() zcl.AttrID                { return ReportingPeriodAttr }
func (ReportingPeriod) Name() string                  { return "Reporting Period" }
func (ReportingPeriod) Readable() bool                { return true }
func (ReportingPeriod) Writable() bool                { return true }
func (ReportingPeriod) Reportable() bool              { return false }
func (ReportingPeriod) SceneIndex() int               { return -1 }
func (a *ReportingPeriod) Value() *ReportingPeriod    { return a }
func (a ReportingPeriod) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ReportingPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReportingPeriod(*nt)
	return br, err
}

func (a *ReportingPeriod) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ReportingPeriod(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ReportingPeriod) String() string {
	return zcl.Seconds.Format(float64(a))
}

type Resolution zcl.Zenum8

func (a *Resolution) Value() *Resolution         { return a }
func (a Resolution) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *Resolution) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Resolution(*nt)
	return br, err
}

func (a *Resolution) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = Resolution(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Resolution) String() string {
	switch a {
	case 0x00:
		return "High"
	case 0x01:
		return "Mid"
	case 0x02:
		return "Low"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a Resolution) IsHigh() bool { return a == 0x00 }
func (a Resolution) IsMid() bool  { return a == 0x01 }
func (a Resolution) IsLow() bool  { return a == 0x02 }
func (a *Resolution) SetHigh()    { *a = 0x00 }
func (a *Resolution) SetMid()     { *a = 0x01 }
func (a *Resolution) SetLow()     { *a = 0x02 }

type SwBuildId zcl.Zcstring

const SwBuildIdAttr zcl.AttrID = 16384

func (SwBuildId) ID() zcl.AttrID                { return SwBuildIdAttr }
func (SwBuildId) Name() string                  { return "SW Build ID" }
func (SwBuildId) Readable() bool                { return true }
func (SwBuildId) Writable() bool                { return false }
func (SwBuildId) Reportable() bool              { return false }
func (SwBuildId) SceneIndex() int               { return -1 }
func (a *SwBuildId) Value() *SwBuildId          { return a }
func (a SwBuildId) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *SwBuildId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = SwBuildId(*nt)
	return br, err
}

func (a *SwBuildId) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = SwBuildId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SwBuildId) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

// SceneCapacity specifies remaining number of scenes that can be added
type SceneCapacity zcl.Zu8

func (a *SceneCapacity) Value() *SceneCapacity      { return a }
func (a SceneCapacity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *SceneCapacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneCapacity(*nt)
	return br, err
}

func (a *SceneCapacity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = SceneCapacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneCapacity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type SceneCopyMode zcl.Zbmp8

func (a *SceneCopyMode) Value() *SceneCopyMode      { return a }
func (a SceneCopyMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *SceneCopyMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneCopyMode(*nt)
	return br, err
}

func (a *SceneCopyMode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = SceneCopyMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneCopyMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Copy All Scenes")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a SceneCopyMode) IsCopyAllScenes() bool    { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a *SceneCopyMode) SetCopyAllScenes(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }

type SceneCount zcl.Zu8

const SceneCountAttr zcl.AttrID = 0

func (SceneCount) ID() zcl.AttrID                { return SceneCountAttr }
func (SceneCount) Name() string                  { return "Scene Count" }
func (SceneCount) Readable() bool                { return true }
func (SceneCount) Writable() bool                { return false }
func (SceneCount) Reportable() bool              { return false }
func (SceneCount) SceneIndex() int               { return -1 }
func (a *SceneCount) Value() *SceneCount         { return a }
func (a SceneCount) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *SceneCount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneCount(*nt)
	return br, err
}

func (a *SceneCount) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = SceneCount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneCount) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// SceneExtensions The format of each extension field set is a 16 bit field carrying the cluster ID,
// followed by an 8 bit length field and the set of scene extension fields specified
// in  the  relevant  cluster. The length field holds the length in octets of that
// extension field set. Extension field set format:
// {{clusterId1, length 1, {extension field set 1}}, {clusterId2, length 2, {extension field set 2}} ...}
// I.e. the field would be a repeating struct with [ClusterID uint16] [OctetLength uint8] [AttrID ...uint16]
type SceneExtensions zcl.SceneExtensionSet

func (a *SceneExtensions) Value() *SceneExtensions    { return a }
func (a SceneExtensions) MarshalZcl() ([]byte, error) { return zcl.SceneExtensionSet(a).MarshalZcl() }

func (a *SceneExtensions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.SceneExtensionSet)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneExtensions(*nt)
	return br, err
}

func (a *SceneExtensions) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.SceneExtensionSet); ok {
		*a = SceneExtensions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneExtensions) String() string {
	return zcl.Sprintf("%v", zcl.SceneExtensionSet(a))
}

type SceneId zcl.Zu8

func (a *SceneId) Value() *SceneId            { return a }
func (a SceneId) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *SceneId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneId(*nt)
	return br, err
}

func (a *SceneId) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = SceneId(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneId) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type SceneLastConfiguredBy zcl.Zuid

const SceneLastConfiguredByAttr zcl.AttrID = 5

func (SceneLastConfiguredBy) ID() zcl.AttrID                   { return SceneLastConfiguredByAttr }
func (SceneLastConfiguredBy) Name() string                     { return "Scene Last Configured By" }
func (SceneLastConfiguredBy) Readable() bool                   { return true }
func (SceneLastConfiguredBy) Writable() bool                   { return false }
func (SceneLastConfiguredBy) Reportable() bool                 { return false }
func (SceneLastConfiguredBy) SceneIndex() int                  { return -1 }
func (a *SceneLastConfiguredBy) Value() *SceneLastConfiguredBy { return a }
func (a SceneLastConfiguredBy) MarshalZcl() ([]byte, error)    { return zcl.Zuid(a).MarshalZcl() }

func (a *SceneLastConfiguredBy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneLastConfiguredBy(*nt)
	return br, err
}

func (a *SceneLastConfiguredBy) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zuid); ok {
		*a = SceneLastConfiguredBy(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneLastConfiguredBy) String() string {
	return zcl.Sprintf("%v", zcl.Zuid(a))
}

type SceneName zcl.Zcstring

func (a *SceneName) Value() *SceneName          { return a }
func (a SceneName) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *SceneName) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneName(*nt)
	return br, err
}

func (a *SceneName) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = SceneName(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneName) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

type SceneNameSupport zcl.Zbmp8

const SceneNameSupportAttr zcl.AttrID = 4

func (SceneNameSupport) ID() zcl.AttrID                { return SceneNameSupportAttr }
func (SceneNameSupport) Name() string                  { return "Scene Name Support" }
func (SceneNameSupport) Readable() bool                { return true }
func (SceneNameSupport) Writable() bool                { return false }
func (SceneNameSupport) Reportable() bool              { return false }
func (SceneNameSupport) SceneIndex() int               { return -1 }
func (a *SceneNameSupport) Value() *SceneNameSupport   { return a }
func (a SceneNameSupport) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *SceneNameSupport) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneNameSupport(*nt)
	return br, err
}

func (a *SceneNameSupport) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = SceneNameSupport(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneNameSupport) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 7:
			bstr = append(bstr, "Names Supported")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a SceneNameSupport) IsNamesSupported() bool { return zcl.BitmapTest([]byte(a[:]), 7) }
func (a *SceneNameSupport) SetNamesSupported(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 7, b))
}

type SceneValid zcl.Zbool

const SceneValidAttr zcl.AttrID = 3

func (SceneValid) ID() zcl.AttrID                { return SceneValidAttr }
func (SceneValid) Name() string                  { return "Scene Valid" }
func (SceneValid) Readable() bool                { return true }
func (SceneValid) Writable() bool                { return false }
func (SceneValid) Reportable() bool              { return false }
func (SceneValid) SceneIndex() int               { return -1 }
func (a *SceneValid) Value() *SceneValid         { return a }
func (a SceneValid) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *SceneValid) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneValid(*nt)
	return br, err
}

func (a *SceneValid) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = SceneValid(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneValid) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type SceneList zcl.Zset

func (a *SceneList) Value() *SceneList { return a }
func (a SceneList) MarshalZcl() ([]byte, error) {
	return nil, zcl.Errorf("not implemented")
}

func (a *SceneList) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zset)
	br, err := nt.UnmarshalZcl(b)
	*a = SceneList(*nt)
	return br, err
}

func (a *SceneList) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zset); ok {
		*a = SceneList(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SceneList) String() string {
	return zcl.Sprintf("%v", zcl.Zset(a))
}

type Sensitivity zcl.Zenum8

const SensitivityAttr zcl.AttrID = 48

func (Sensitivity) ID() zcl.AttrID                { return SensitivityAttr }
func (Sensitivity) Name() string                  { return "Sensitivity" }
func (Sensitivity) Readable() bool                { return true }
func (Sensitivity) Writable() bool                { return true }
func (Sensitivity) Reportable() bool              { return false }
func (Sensitivity) SceneIndex() int               { return -1 }
func (a *Sensitivity) Value() *Sensitivity        { return a }
func (a Sensitivity) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *Sensitivity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Sensitivity(*nt)
	return br, err
}

func (a *Sensitivity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = Sensitivity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Sensitivity) String() string {
	switch a {
	case 0x00:
		return "default"
	case 0x01:
		return "high"
	case 0x02:
		return "max"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a Sensitivity) IsDefault() bool { return a == 0x00 }
func (a Sensitivity) IsHigh() bool    { return a == 0x01 }
func (a Sensitivity) IsMax() bool     { return a == 0x02 }
func (a *Sensitivity) SetDefault()    { *a = 0x00 }
func (a *Sensitivity) SetHigh()       { *a = 0x01 }
func (a *Sensitivity) SetMax()        { *a = 0x02 }

type StackVersion zcl.Zu8

const StackVersionAttr zcl.AttrID = 2

func (StackVersion) ID() zcl.AttrID                { return StackVersionAttr }
func (StackVersion) Name() string                  { return "Stack Version" }
func (StackVersion) Readable() bool                { return true }
func (StackVersion) Writable() bool                { return false }
func (StackVersion) Reportable() bool              { return false }
func (StackVersion) SceneIndex() int               { return -1 }
func (a *StackVersion) Value() *StackVersion       { return a }
func (a StackVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *StackVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StackVersion(*nt)
	return br, err
}

func (a *StackVersion) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = StackVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StackVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// StandardTime Local time (without DST offset)
type StandardTime zcl.Zu32

const StandardTimeAttr zcl.AttrID = 6

func (StandardTime) ID() zcl.AttrID                { return StandardTimeAttr }
func (StandardTime) Name() string                  { return "Standard Time" }
func (StandardTime) Readable() bool                { return true }
func (StandardTime) Writable() bool                { return false }
func (StandardTime) Reportable() bool              { return false }
func (StandardTime) SceneIndex() int               { return -1 }
func (a *StandardTime) Value() *StandardTime       { return a }
func (a StandardTime) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *StandardTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = StandardTime(*nt)
	return br, err
}

func (a *StandardTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu32); ok {
		*a = StandardTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StandardTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

type Status zcl.Status

func (a *Status) Value() *Status             { return a }
func (a Status) MarshalZcl() ([]byte, error) { return zcl.Status(a).MarshalZcl() }

func (a *Status) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Status)
	br, err := nt.UnmarshalZcl(b)
	*a = Status(*nt)
	return br, err
}

func (a *Status) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Status); ok {
		*a = Status(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Status) String() string {
	return zcl.Sprintf("%v", zcl.Status(a))
}

type StepSize zcl.Zu8

func (a *StepSize) Value() *StepSize           { return a }
func (a StepSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *StepSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StepSize(*nt)
	return br, err
}

func (a *StepSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = StepSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StepSize) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

// SwitchActions specifies the commands of the On/Off cluster to be generated when the switch moves between its two states.
type SwitchActions zcl.Zenum8

const SwitchActionsAttr zcl.AttrID = 16

func (SwitchActions) ID() zcl.AttrID                { return SwitchActionsAttr }
func (SwitchActions) Name() string                  { return "Switch actions" }
func (SwitchActions) Readable() bool                { return true }
func (SwitchActions) Writable() bool                { return true }
func (SwitchActions) Reportable() bool              { return false }
func (SwitchActions) SceneIndex() int               { return -1 }
func (a *SwitchActions) Value() *SwitchActions      { return a }
func (a SwitchActions) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *SwitchActions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = SwitchActions(*nt)
	return br, err
}

func (a *SwitchActions) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = SwitchActions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SwitchActions) String() string {
	switch a {
	case 0x00:
		return "On-Off"
	case 0x01:
		return "Off-On"
	case 0x02:
		return "Toggle"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a SwitchActions) IsOnOff() bool  { return a == 0x00 }
func (a SwitchActions) IsOffOn() bool  { return a == 0x01 }
func (a SwitchActions) IsToggle() bool { return a == 0x02 }
func (a *SwitchActions) SetOnOff()     { *a = 0x00 }
func (a *SwitchActions) SetOffOn()     { *a = 0x01 }
func (a *SwitchActions) SetToggle()    { *a = 0x02 }

// SwitchType specifies the basic functionality of the On/Off switching device.
type SwitchType zcl.Zenum8

const SwitchTypeAttr zcl.AttrID = 0

func (SwitchType) ID() zcl.AttrID                { return SwitchTypeAttr }
func (SwitchType) Name() string                  { return "Switch type" }
func (SwitchType) Readable() bool                { return true }
func (SwitchType) Writable() bool                { return false }
func (SwitchType) Reportable() bool              { return false }
func (SwitchType) SceneIndex() int               { return -1 }
func (a *SwitchType) Value() *SwitchType         { return a }
func (a SwitchType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *SwitchType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = SwitchType(*nt)
	return br, err
}

func (a *SwitchType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = SwitchType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a SwitchType) String() string {
	switch a {
	case 0x00:
		return "Toggle"
	case 0x01:
		return "Momentary"
	case 0x02:
		return "Multifunction"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a SwitchType) IsToggle() bool        { return a == 0x00 }
func (a SwitchType) IsMomentary() bool     { return a == 0x01 }
func (a SwitchType) IsMultifunction() bool { return a == 0x02 }
func (a *SwitchType) SetToggle()           { *a = 0x00 }
func (a *SwitchType) SetMomentary()        { *a = 0x01 }
func (a *SwitchType) SetMultifunction()    { *a = 0x02 }

type Time zcl.Zutc

const TimeAttr zcl.AttrID = 0

func (Time) ID() zcl.AttrID                { return TimeAttr }
func (Time) Name() string                  { return "Time" }
func (Time) Readable() bool                { return true }
func (Time) Writable() bool                { return true }
func (Time) Reportable() bool              { return false }
func (Time) SceneIndex() int               { return -1 }
func (a *Time) Value() *Time               { return a }
func (a Time) MarshalZcl() ([]byte, error) { return zcl.Zutc(a).MarshalZcl() }

func (a *Time) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = Time(*nt)
	return br, err
}

func (a *Time) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zutc); ok {
		*a = Time(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Time) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(a))
}

type TimeStatus zcl.Zbmp8

const TimeStatusAttr zcl.AttrID = 1

func (TimeStatus) ID() zcl.AttrID                { return TimeStatusAttr }
func (TimeStatus) Name() string                  { return "Time Status" }
func (TimeStatus) Readable() bool                { return true }
func (TimeStatus) Writable() bool                { return true }
func (TimeStatus) Reportable() bool              { return false }
func (TimeStatus) SceneIndex() int               { return -1 }
func (a *TimeStatus) Value() *TimeStatus         { return a }
func (a TimeStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *TimeStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = TimeStatus(*nt)
	return br, err
}

func (a *TimeStatus) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = TimeStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TimeStatus) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Master Clock")
		case 1:
			bstr = append(bstr, "Synchronized")
		case 2:
			bstr = append(bstr, "Master for Timezone and Dst")
		case 3:
			bstr = append(bstr, "Superseding")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a TimeStatus) IsMasterClock() bool             { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a TimeStatus) IsSynchronized() bool            { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a TimeStatus) IsMasterForTimezoneAndDst() bool { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a TimeStatus) IsSuperseding() bool             { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *TimeStatus) SetMasterClock(b bool)          { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *TimeStatus) SetSynchronized(b bool)         { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b)) }
func (a *TimeStatus) SetMasterForTimezoneAndDst(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b))
}
func (a *TimeStatus) SetSuperseding(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b)) }

// TimeZone Offset during normal time from UTC in seconds
type TimeZone zcl.Zs32

const TimeZoneAttr zcl.AttrID = 2

func (TimeZone) ID() zcl.AttrID                { return TimeZoneAttr }
func (TimeZone) Name() string                  { return "Time Zone" }
func (TimeZone) Readable() bool                { return true }
func (TimeZone) Writable() bool                { return true }
func (TimeZone) Reportable() bool              { return false }
func (TimeZone) SceneIndex() int               { return -1 }
func (a *TimeZone) Value() *TimeZone           { return a }
func (a TimeZone) MarshalZcl() ([]byte, error) { return zcl.Zs32(a).MarshalZcl() }

func (a *TimeZone) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = TimeZone(*nt)
	return br, err
}

func (a *TimeZone) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs32); ok {
		*a = TimeZone(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TimeZone) String() string {
	return zcl.Seconds.Format(float64(a))
}

type TransitionTime zcl.Zu16

func (a *TransitionTime) Value() *TransitionTime     { return a }
func (a TransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *TransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TransitionTime(*nt)
	return br, err
}

func (a *TransitionTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = TransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TransitionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type TransitionTimeSec zcl.Zu16

func (a *TransitionTimeSec) Value() *TransitionTimeSec  { return a }
func (a TransitionTimeSec) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *TransitionTimeSec) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TransitionTimeSec(*nt)
	return br, err
}

func (a *TransitionTimeSec) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = TransitionTimeSec(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TransitionTimeSec) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type UserTest zcl.Zbool

const UserTestAttr zcl.AttrID = 50

func (UserTest) ID() zcl.AttrID                { return UserTestAttr }
func (UserTest) Name() string                  { return "User test" }
func (UserTest) Readable() bool                { return true }
func (UserTest) Writable() bool                { return true }
func (UserTest) Reportable() bool              { return false }
func (UserTest) SceneIndex() int               { return -1 }
func (a *UserTest) Value() *UserTest           { return a }
func (a UserTest) MarshalZcl() ([]byte, error) { return zcl.Zbool(a).MarshalZcl() }

func (a *UserTest) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = UserTest(*nt)
	return br, err
}

func (a *UserTest) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbool); ok {
		*a = UserTest(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a UserTest) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

type ValidUntilTime zcl.Zutc

const ValidUntilTimeAttr zcl.AttrID = 9

func (ValidUntilTime) ID() zcl.AttrID                { return ValidUntilTimeAttr }
func (ValidUntilTime) Name() string                  { return "Valid Until Time" }
func (ValidUntilTime) Readable() bool                { return true }
func (ValidUntilTime) Writable() bool                { return true }
func (ValidUntilTime) Reportable() bool              { return false }
func (ValidUntilTime) SceneIndex() int               { return -1 }
func (a *ValidUntilTime) Value() *ValidUntilTime     { return a }
func (a ValidUntilTime) MarshalZcl() ([]byte, error) { return zcl.Zutc(a).MarshalZcl() }

func (a *ValidUntilTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*a = ValidUntilTime(*nt)
	return br, err
}

func (a *ValidUntilTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zutc); ok {
		*a = ValidUntilTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ValidUntilTime) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(a))
}

type XCoordinate zcl.Zs16

const XCoordinateAttr zcl.AttrID = 16

func (XCoordinate) ID() zcl.AttrID                { return XCoordinateAttr }
func (XCoordinate) Name() string                  { return "X Coordinate" }
func (XCoordinate) Readable() bool                { return true }
func (XCoordinate) Writable() bool                { return true }
func (XCoordinate) Reportable() bool              { return false }
func (XCoordinate) SceneIndex() int               { return -1 }
func (a *XCoordinate) Value() *XCoordinate        { return a }
func (a XCoordinate) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *XCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = XCoordinate(*nt)
	return br, err
}

func (a *XCoordinate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = XCoordinate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a XCoordinate) String() string {
	return zcl.Meters.Format(float64(a) / 10)
}

type YCoordinate zcl.Zs16

const YCoordinateAttr zcl.AttrID = 17

func (YCoordinate) ID() zcl.AttrID                { return YCoordinateAttr }
func (YCoordinate) Name() string                  { return "Y Coordinate" }
func (YCoordinate) Readable() bool                { return true }
func (YCoordinate) Writable() bool                { return true }
func (YCoordinate) Reportable() bool              { return false }
func (YCoordinate) SceneIndex() int               { return -1 }
func (a *YCoordinate) Value() *YCoordinate        { return a }
func (a YCoordinate) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *YCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = YCoordinate(*nt)
	return br, err
}

func (a *YCoordinate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = YCoordinate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a YCoordinate) String() string {
	return zcl.Meters.Format(float64(a) / 10)
}

type ZCoordinate zcl.Zs16

const ZCoordinateAttr zcl.AttrID = 18

func (ZCoordinate) ID() zcl.AttrID                { return ZCoordinateAttr }
func (ZCoordinate) Name() string                  { return "Z Coordinate" }
func (ZCoordinate) Readable() bool                { return true }
func (ZCoordinate) Writable() bool                { return true }
func (ZCoordinate) Reportable() bool              { return false }
func (ZCoordinate) SceneIndex() int               { return -1 }
func (a *ZCoordinate) Value() *ZCoordinate        { return a }
func (a ZCoordinate) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *ZCoordinate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ZCoordinate(*nt)
	return br, err
}

func (a *ZCoordinate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = ZCoordinate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ZCoordinate) String() string {
	return zcl.Meters.Format(float64(a) / 10)
}

type ZclVersion zcl.Zu8

const ZclVersionAttr zcl.AttrID = 0

func (ZclVersion) ID() zcl.AttrID                { return ZclVersionAttr }
func (ZclVersion) Name() string                  { return "ZCL Version" }
func (ZclVersion) Readable() bool                { return true }
func (ZclVersion) Writable() bool                { return false }
func (ZclVersion) Reportable() bool              { return false }
func (ZclVersion) SceneIndex() int               { return -1 }
func (a *ZclVersion) Value() *ZclVersion         { return a }
func (a ZclVersion) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ZclVersion) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ZclVersion(*nt)
	return br, err
}

func (a *ZclVersion) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ZclVersion(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ZclVersion) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}
