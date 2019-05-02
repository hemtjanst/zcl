// The IAS Zone cluster defines an interface to the functionality of an IAS security zone device. IAS Zone supports up totwo alarm types per zone, low battery reports and supervision of the IAS network.
package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// IasZone
const IasZoneID zcl.ClusterID = 1280

var IasZoneCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		ZoneEnrollResponseCommand:           func() zcl.Command { return new(ZoneEnrollResponse) },
		InitiateNormalOperationModeCommand:  func() zcl.Command { return new(InitiateNormalOperationMode) },
		ZoneStatusChangeNotificationCommand: func() zcl.Command { return new(ZoneStatusChangeNotification) },
		ZoneEnrollRequestCommand:            func() zcl.Command { return new(ZoneEnrollRequest) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		ZoneStateAttr:     func() zcl.Attr { return new(ZoneState) },
		ZoneTypeAttr:      func() zcl.Attr { return new(ZoneType) },
		ZoneStatusAttr:    func() zcl.Attr { return new(ZoneStatus) },
		IasCieAddressAttr: func() zcl.Attr { return new(IasCieAddress) },
		ZoneIdAttr:        func() zcl.Attr { return new(ZoneId) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type ZoneEnrollResponse struct {
	EnrollResponseCode zcl.Zenum16
	ZoneId             zcl.Zu8
}

const ZoneEnrollResponseCommand zcl.CommandID = 0

func (v *ZoneEnrollResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.EnrollResponseCode,
		&v.ZoneId,
	}
}

func (v ZoneEnrollResponse) ID() zcl.CommandID {
	return ZoneEnrollResponseCommand
}

func (v ZoneEnrollResponse) Cluster() zcl.ClusterID {
	return IasZoneID
}

func (v ZoneEnrollResponse) MnfCode() []byte {
	return []byte{}
}

func (v ZoneEnrollResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.EnrollResponseCode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ZoneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ZoneEnrollResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.EnrollResponseCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZoneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type InitiateNormalOperationMode struct {
}

const InitiateNormalOperationModeCommand zcl.CommandID = 1

func (v *InitiateNormalOperationMode) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v InitiateNormalOperationMode) ID() zcl.CommandID {
	return InitiateNormalOperationModeCommand
}

func (v InitiateNormalOperationMode) Cluster() zcl.ClusterID {
	return IasZoneID
}

func (v InitiateNormalOperationMode) MnfCode() []byte {
	return []byte{}
}

func (v InitiateNormalOperationMode) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *InitiateNormalOperationMode) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

type ZoneStatusChangeNotification struct {
	ZoneStatus     zcl.Zbmp16
	ExtendedStatus zcl.Zu8
	ZoneId         zcl.Zu8
	Delay          zcl.Zu16
}

const ZoneStatusChangeNotificationCommand zcl.CommandID = 0

func (v *ZoneStatusChangeNotification) Values() []zcl.Val {
	return []zcl.Val{
		&v.ZoneStatus,
		&v.ExtendedStatus,
		&v.ZoneId,
		&v.Delay,
	}
}

func (v ZoneStatusChangeNotification) ID() zcl.CommandID {
	return ZoneStatusChangeNotificationCommand
}

func (v ZoneStatusChangeNotification) Cluster() zcl.ClusterID {
	return IasZoneID
}

func (v ZoneStatusChangeNotification) MnfCode() []byte {
	return []byte{}
}

func (v ZoneStatusChangeNotification) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.ZoneStatus.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ExtendedStatus.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ZoneId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Delay.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ZoneStatusChangeNotification) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.ZoneStatus).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ExtendedStatus).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZoneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Delay).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type ZoneEnrollRequest struct {
	ZoneStatus       zcl.Zenum16
	ManufacturerCode zcl.Zu16
}

const ZoneEnrollRequestCommand zcl.CommandID = 1

func (v *ZoneEnrollRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.ZoneStatus,
		&v.ManufacturerCode,
	}
}

func (v ZoneEnrollRequest) ID() zcl.CommandID {
	return ZoneEnrollRequestCommand
}

func (v ZoneEnrollRequest) Cluster() zcl.ClusterID {
	return IasZoneID
}

func (v ZoneEnrollRequest) MnfCode() []byte {
	return []byte{}
}

func (v ZoneEnrollRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.ZoneStatus.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ZoneEnrollRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.ZoneStatus).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// ZoneState is an autogenerated attribute in the IasZone cluster
type ZoneState zcl.Zenum8

const ZoneStateAttr zcl.AttrID = 0

func (a ZoneState) ID() zcl.AttrID         { return ZoneStateAttr }
func (a ZoneState) Cluster() zcl.ClusterID { return IasZoneID }
func (a *ZoneState) Value() *ZoneState     { return a }
func (a ZoneState) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *ZoneState) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = ZoneState(*nt)
	return br, err
}

func (a ZoneState) Readable() bool   { return true }
func (a ZoneState) Writable() bool   { return false }
func (a ZoneState) Reportable() bool { return false }
func (a ZoneState) SceneIndex() int  { return -1 }

func (a ZoneState) String() string {
	switch a {
	case 0x00:
		return "Not enrolled"
	case 0x01:
		return "Enrolled"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsNotEnrolled checks if ZoneState equals the value for Not enrolled (0x00)
func (a ZoneState) IsNotEnrolled() bool { return a == 0x00 }

// SetNotEnrolled sets ZoneState to Not enrolled (0x00)
func (a *ZoneState) SetNotEnrolled() { *a = 0x00 }

// IsEnrolled checks if ZoneState equals the value for Enrolled (0x01)
func (a ZoneState) IsEnrolled() bool { return a == 0x01 }

// SetEnrolled sets ZoneState to Enrolled (0x01)
func (a *ZoneState) SetEnrolled() { *a = 0x01 }

// ZoneType is an autogenerated attribute in the IasZone cluster
type ZoneType zcl.Zenum16

const ZoneTypeAttr zcl.AttrID = 1

func (a ZoneType) ID() zcl.AttrID         { return ZoneTypeAttr }
func (a ZoneType) Cluster() zcl.ClusterID { return IasZoneID }
func (a *ZoneType) Value() *ZoneType      { return a }
func (a ZoneType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum16(a).MarshalZcl()
}
func (a *ZoneType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*a = ZoneType(*nt)
	return br, err
}

func (a ZoneType) Readable() bool   { return true }
func (a ZoneType) Writable() bool   { return false }
func (a ZoneType) Reportable() bool { return false }
func (a ZoneType) SceneIndex() int  { return -1 }

func (a ZoneType) String() string {
	switch a {
	case 0x0000:
		return "Standard CIE"
	case 0x000D:
		return "Motion sensor"
	case 0x0015:
		return "Contact switch"
	case 0x0028:
		return "Fire sensor"
	case 0x002A:
		return "Water sensor"
	case 0x002B:
		return "Carbon Monoxide (CO) sensor"
	case 0x002C:
		return "Personal emergency device"
	case 0x002D:
		return "Vibration / Movement sensor"
	case 0x010F:
		return "Remote Control"
	case 0x0115:
		return "Key fob"
	case 0x021D:
		return "Keypad"
	case 0x0225:
		return "Standard Warning Device"
	case 0x0226:
		return "Glass break sensor"
	case 0x0229:
		return "Security repeater"
	case 0xFFFF:
		return "Invalid Zone Type"
	}
	return zcl.Sprintf("%s", zcl.Zenum16(a))
}

// IsStandardCie checks if ZoneType equals the value for Standard CIE (0x0000)
func (a ZoneType) IsStandardCie() bool { return a == 0x0000 }

// SetStandardCie sets ZoneType to Standard CIE (0x0000)
func (a *ZoneType) SetStandardCie() { *a = 0x0000 }

// IsMotionSensor checks if ZoneType equals the value for Motion sensor (0x000D)
func (a ZoneType) IsMotionSensor() bool { return a == 0x000D }

// SetMotionSensor sets ZoneType to Motion sensor (0x000D)
func (a *ZoneType) SetMotionSensor() { *a = 0x000D }

// IsContactSwitch checks if ZoneType equals the value for Contact switch (0x0015)
func (a ZoneType) IsContactSwitch() bool { return a == 0x0015 }

// SetContactSwitch sets ZoneType to Contact switch (0x0015)
func (a *ZoneType) SetContactSwitch() { *a = 0x0015 }

// IsFireSensor checks if ZoneType equals the value for Fire sensor (0x0028)
func (a ZoneType) IsFireSensor() bool { return a == 0x0028 }

// SetFireSensor sets ZoneType to Fire sensor (0x0028)
func (a *ZoneType) SetFireSensor() { *a = 0x0028 }

// IsWaterSensor checks if ZoneType equals the value for Water sensor (0x002A)
func (a ZoneType) IsWaterSensor() bool { return a == 0x002A }

// SetWaterSensor sets ZoneType to Water sensor (0x002A)
func (a *ZoneType) SetWaterSensor() { *a = 0x002A }

// IsCarbonMonoxideCoSensor checks if ZoneType equals the value for Carbon Monoxide (CO) sensor (0x002B)
func (a ZoneType) IsCarbonMonoxideCoSensor() bool { return a == 0x002B }

// SetCarbonMonoxideCoSensor sets ZoneType to Carbon Monoxide (CO) sensor (0x002B)
func (a *ZoneType) SetCarbonMonoxideCoSensor() { *a = 0x002B }

// IsPersonalEmergencyDevice checks if ZoneType equals the value for Personal emergency device (0x002C)
func (a ZoneType) IsPersonalEmergencyDevice() bool { return a == 0x002C }

// SetPersonalEmergencyDevice sets ZoneType to Personal emergency device (0x002C)
func (a *ZoneType) SetPersonalEmergencyDevice() { *a = 0x002C }

// IsVibrationMovementSensor checks if ZoneType equals the value for Vibration / Movement sensor (0x002D)
func (a ZoneType) IsVibrationMovementSensor() bool { return a == 0x002D }

// SetVibrationMovementSensor sets ZoneType to Vibration / Movement sensor (0x002D)
func (a *ZoneType) SetVibrationMovementSensor() { *a = 0x002D }

// IsRemoteControl checks if ZoneType equals the value for Remote Control (0x010F)
func (a ZoneType) IsRemoteControl() bool { return a == 0x010F }

// SetRemoteControl sets ZoneType to Remote Control (0x010F)
func (a *ZoneType) SetRemoteControl() { *a = 0x010F }

// IsKeyFob checks if ZoneType equals the value for Key fob (0x0115)
func (a ZoneType) IsKeyFob() bool { return a == 0x0115 }

// SetKeyFob sets ZoneType to Key fob (0x0115)
func (a *ZoneType) SetKeyFob() { *a = 0x0115 }

// IsKeypad checks if ZoneType equals the value for Keypad (0x021D)
func (a ZoneType) IsKeypad() bool { return a == 0x021D }

// SetKeypad sets ZoneType to Keypad (0x021D)
func (a *ZoneType) SetKeypad() { *a = 0x021D }

// IsStandardWarningDevice checks if ZoneType equals the value for Standard Warning Device (0x0225)
func (a ZoneType) IsStandardWarningDevice() bool { return a == 0x0225 }

// SetStandardWarningDevice sets ZoneType to Standard Warning Device (0x0225)
func (a *ZoneType) SetStandardWarningDevice() { *a = 0x0225 }

// IsGlassBreakSensor checks if ZoneType equals the value for Glass break sensor (0x0226)
func (a ZoneType) IsGlassBreakSensor() bool { return a == 0x0226 }

// SetGlassBreakSensor sets ZoneType to Glass break sensor (0x0226)
func (a *ZoneType) SetGlassBreakSensor() { *a = 0x0226 }

// IsSecurityRepeater checks if ZoneType equals the value for Security repeater (0x0229)
func (a ZoneType) IsSecurityRepeater() bool { return a == 0x0229 }

// SetSecurityRepeater sets ZoneType to Security repeater (0x0229)
func (a *ZoneType) SetSecurityRepeater() { *a = 0x0229 }

// IsInvalidZoneType checks if ZoneType equals the value for Invalid Zone Type (0xFFFF)
func (a ZoneType) IsInvalidZoneType() bool { return a == 0xFFFF }

// SetInvalidZoneType sets ZoneType to Invalid Zone Type (0xFFFF)
func (a *ZoneType) SetInvalidZoneType() { *a = 0xFFFF }

// ZoneStatus is an autogenerated attribute in the IasZone cluster
type ZoneStatus zcl.Zbmp16

const ZoneStatusAttr zcl.AttrID = 2

func (a ZoneStatus) ID() zcl.AttrID         { return ZoneStatusAttr }
func (a ZoneStatus) Cluster() zcl.ClusterID { return IasZoneID }
func (a *ZoneStatus) Value() *ZoneStatus    { return a }
func (a ZoneStatus) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp16(a).MarshalZcl()
}
func (a *ZoneStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = ZoneStatus(*nt)
	return br, err
}

func (a ZoneStatus) Readable() bool   { return true }
func (a ZoneStatus) Writable() bool   { return false }
func (a ZoneStatus) Reportable() bool { return false }
func (a ZoneStatus) SceneIndex() int  { return -1 }

func (a ZoneStatus) String() string {
	var bstr []string
	if a.IsAlarm1() {
		bstr = append(bstr, "Alarm 1")
	}
	if a.IsAlarm2() {
		bstr = append(bstr, "Alarm 2")
	}
	if a.IsTamper() {
		bstr = append(bstr, "Tamper")
	}
	if a.IsBattery() {
		bstr = append(bstr, "Battery")
	}
	if a.IsSupervisionReports() {
		bstr = append(bstr, "Supervision reports")
	}
	if a.IsRestoreReports() {
		bstr = append(bstr, "Restore reports")
	}
	if a.IsTrouble() {
		bstr = append(bstr, "Trouble")
	}
	if a.IsAcMains() {
		bstr = append(bstr, "AC (mains)")
	}
	if a.IsTest() {
		bstr = append(bstr, "Test")
	}
	if a.IsBatteryDefect() {
		bstr = append(bstr, "Battery defect")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a ZoneStatus) IsAlarm1() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *ZoneStatus) SetAlarm1(b bool) {
	*a = ZoneStatus(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a ZoneStatus) IsAlarm2() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *ZoneStatus) SetAlarm2(b bool) {
	*a = ZoneStatus(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a ZoneStatus) IsTamper() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *ZoneStatus) SetTamper(b bool) {
	*a = ZoneStatus(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a ZoneStatus) IsBattery() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *ZoneStatus) SetBattery(b bool) {
	*a = ZoneStatus(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a ZoneStatus) IsSupervisionReports() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *ZoneStatus) SetSupervisionReports(b bool) {
	*a = ZoneStatus(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a ZoneStatus) IsRestoreReports() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *ZoneStatus) SetRestoreReports(b bool) {
	*a = ZoneStatus(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a ZoneStatus) IsTrouble() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *ZoneStatus) SetTrouble(b bool) {
	*a = ZoneStatus(zcl.BitmapSet([]byte(*a), 6, b))
}

func (a ZoneStatus) IsAcMains() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *ZoneStatus) SetAcMains(b bool) {
	*a = ZoneStatus(zcl.BitmapSet([]byte(*a), 7, b))
}

func (a ZoneStatus) IsTest() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *ZoneStatus) SetTest(b bool) {
	*a = ZoneStatus(zcl.BitmapSet([]byte(*a), 8, b))
}

func (a ZoneStatus) IsBatteryDefect() bool {
	return zcl.BitmapTest([]byte(a), 9)
}
func (a *ZoneStatus) SetBatteryDefect(b bool) {
	*a = ZoneStatus(zcl.BitmapSet([]byte(*a), 9, b))
}

// IasCieAddress is an autogenerated attribute in the IasZone cluster
type IasCieAddress zcl.Zuid

const IasCieAddressAttr zcl.AttrID = 16

func (a IasCieAddress) ID() zcl.AttrID         { return IasCieAddressAttr }
func (a IasCieAddress) Cluster() zcl.ClusterID { return IasZoneID }
func (a *IasCieAddress) Value() *IasCieAddress { return a }
func (a IasCieAddress) MarshalZcl() ([]byte, error) {
	return zcl.Zuid(a).MarshalZcl()
}
func (a *IasCieAddress) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zuid)
	br, err := nt.UnmarshalZcl(b)
	*a = IasCieAddress(*nt)
	return br, err
}

func (a IasCieAddress) Readable() bool   { return true }
func (a IasCieAddress) Writable() bool   { return true }
func (a IasCieAddress) Reportable() bool { return false }
func (a IasCieAddress) SceneIndex() int  { return -1 }

func (a IasCieAddress) String() string {
	return zcl.Sprintf("%s", zcl.Zuid(a))
}

// ZoneId is an autogenerated attribute in the IasZone cluster
type ZoneId zcl.Zu8

const ZoneIdAttr zcl.AttrID = 17

func (a ZoneId) ID() zcl.AttrID         { return ZoneIdAttr }
func (a ZoneId) Cluster() zcl.ClusterID { return IasZoneID }
func (a *ZoneId) Value() *ZoneId        { return a }
func (a ZoneId) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ZoneId) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ZoneId(*nt)
	return br, err
}

func (a ZoneId) Readable() bool   { return true }
func (a ZoneId) Writable() bool   { return false }
func (a ZoneId) Reportable() bool { return false }
func (a ZoneId) SceneIndex() int  { return -1 }

func (a ZoneId) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}
