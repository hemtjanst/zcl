// Attributes for determining more detailed information about a device’s power source(s), and for configuring under/over voltage alarms.
package general

import (
	"neotor.se/zcl"
)

// PowerConfiguration
// Attributes for determining more detailed information about a device’s power source(s), and for configuring under/over voltage alarms.

func NewPowerConfigurationServer(profile zcl.ProfileID) *PowerConfigurationServer {
	return &PowerConfigurationServer{p: profile}
}
func NewPowerConfigurationClient(profile zcl.ProfileID) *PowerConfigurationClient {
	return &PowerConfigurationClient{p: profile}
}

const PowerConfigurationCluster zcl.ClusterID = 1

type PowerConfigurationServer struct {
	p zcl.ProfileID

	MainsVoltage                  *MainsVoltage
	MainsFrequency                *MainsFrequency
	MainsAlarmMask                *MainsAlarmMask
	MainsVoltageMinThreshold      *MainsVoltageMinThreshold
	MainsVoltageMaxThreshold      *MainsVoltageMaxThreshold
	MainsVoltageDwellTripPoint    *MainsVoltageDwellTripPoint
	BatteryVoltage                *BatteryVoltage
	BatteryPercentageRemaining    *BatteryPercentageRemaining
	BatteryManufacturer           *BatteryManufacturer
	BatterySize                   *BatterySize
	BatteryAhrRating              *BatteryAhrRating
	BatteryQuantity               *BatteryQuantity
	BatteryRatedVoltage           *BatteryRatedVoltage
	BatteryAlarmMask              *BatteryAlarmMask
	BatteryVoltageMinThreshold    *BatteryVoltageMinThreshold
	BatteryVoltageThreshold1      *BatteryVoltageThreshold1
	BatteryVoltageThreshold2      *BatteryVoltageThreshold2
	BatteryVoltageThreshold3      *BatteryVoltageThreshold3
	BatteryPercentageMinThreshold *BatteryPercentageMinThreshold
	BatteryPercentageThreshold1   *BatteryPercentageThreshold1
	BatteryPercentageThreshold2   *BatteryPercentageThreshold2
	BatteryPercentageThreshold3   *BatteryPercentageThreshold3
	BatteryAlarmState             *BatteryAlarmState
}

type PowerConfigurationClient struct {
	p zcl.ProfileID
}

/*
var PowerConfigurationServer = map[zcl.CommandID]func() zcl.Command{
}

var PowerConfigurationClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type MainsVoltage zcl.Zu16

func (a MainsVoltage) ID() zcl.AttrID         { return 0 }
func (a MainsVoltage) Cluster() zcl.ClusterID { return PowerConfigurationCluster }
func (a *MainsVoltage) Value() *MainsVoltage  { return a }
func (a MainsVoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MainsVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltage(*nt)
	return br, err
}

func (a MainsVoltage) Readable() bool   { return true }
func (a MainsVoltage) Writable() bool   { return false }
func (a MainsVoltage) Reportable() bool { return false }
func (a MainsVoltage) SceneIndex() int  { return -1 }

func (a MainsVoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type MainsFrequency zcl.Zu8

func (a MainsFrequency) ID() zcl.AttrID          { return 1 }
func (a MainsFrequency) Cluster() zcl.ClusterID  { return PowerConfigurationCluster }
func (a *MainsFrequency) Value() *MainsFrequency { return a }
func (a MainsFrequency) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *MainsFrequency) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsFrequency(*nt)
	return br, err
}

func (a MainsFrequency) Readable() bool   { return true }
func (a MainsFrequency) Writable() bool   { return false }
func (a MainsFrequency) Reportable() bool { return false }
func (a MainsFrequency) SceneIndex() int  { return -1 }

func (a MainsFrequency) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type MainsAlarmMask zcl.Zbmp8

func (a MainsAlarmMask) ID() zcl.AttrID          { return 16 }
func (a MainsAlarmMask) Cluster() zcl.ClusterID  { return PowerConfigurationCluster }
func (a *MainsAlarmMask) Value() *MainsAlarmMask { return a }
func (a MainsAlarmMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *MainsAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsAlarmMask(*nt)
	return br, err
}

func (a MainsAlarmMask) Readable() bool   { return true }
func (a MainsAlarmMask) Writable() bool   { return true }
func (a MainsAlarmMask) Reportable() bool { return false }
func (a MainsAlarmMask) SceneIndex() int  { return -1 }

func (a MainsAlarmMask) String() string {

	var bstr []string
	if a.IsMainsVoltageTooLow() {
		bstr = append(bstr, "Mains Voltage too low")
	}
	if a.IsMainsVoltageTooHigh() {
		bstr = append(bstr, "Mains Voltage too high")
	}
	if a.IsMainsPowerSupplyLostUnavailable() {
		bstr = append(bstr, "Mains power supply lost/unavailable")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a MainsAlarmMask) IsMainsVoltageTooLow() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *MainsAlarmMask) SetMainsVoltageTooLow(b bool) {
	*a = MainsAlarmMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a MainsAlarmMask) IsMainsVoltageTooHigh() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *MainsAlarmMask) SetMainsVoltageTooHigh(b bool) {
	*a = MainsAlarmMask(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a MainsAlarmMask) IsMainsPowerSupplyLostUnavailable() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *MainsAlarmMask) SetMainsPowerSupplyLostUnavailable(b bool) {
	*a = MainsAlarmMask(zcl.BitmapSet([]byte(*a), 2, b))
}

type MainsVoltageMinThreshold zcl.Zu16

func (a MainsVoltageMinThreshold) ID() zcl.AttrID                    { return 17 }
func (a MainsVoltageMinThreshold) Cluster() zcl.ClusterID            { return PowerConfigurationCluster }
func (a *MainsVoltageMinThreshold) Value() *MainsVoltageMinThreshold { return a }
func (a MainsVoltageMinThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MainsVoltageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltageMinThreshold(*nt)
	return br, err
}

func (a MainsVoltageMinThreshold) Readable() bool   { return true }
func (a MainsVoltageMinThreshold) Writable() bool   { return true }
func (a MainsVoltageMinThreshold) Reportable() bool { return false }
func (a MainsVoltageMinThreshold) SceneIndex() int  { return -1 }

func (a MainsVoltageMinThreshold) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type MainsVoltageMaxThreshold zcl.Zu16

func (a MainsVoltageMaxThreshold) ID() zcl.AttrID                    { return 18 }
func (a MainsVoltageMaxThreshold) Cluster() zcl.ClusterID            { return PowerConfigurationCluster }
func (a *MainsVoltageMaxThreshold) Value() *MainsVoltageMaxThreshold { return a }
func (a MainsVoltageMaxThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MainsVoltageMaxThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltageMaxThreshold(*nt)
	return br, err
}

func (a MainsVoltageMaxThreshold) Readable() bool   { return true }
func (a MainsVoltageMaxThreshold) Writable() bool   { return true }
func (a MainsVoltageMaxThreshold) Reportable() bool { return false }
func (a MainsVoltageMaxThreshold) SceneIndex() int  { return -1 }

func (a MainsVoltageMaxThreshold) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type MainsVoltageDwellTripPoint zcl.Zu16

func (a MainsVoltageDwellTripPoint) ID() zcl.AttrID                      { return 19 }
func (a MainsVoltageDwellTripPoint) Cluster() zcl.ClusterID              { return PowerConfigurationCluster }
func (a *MainsVoltageDwellTripPoint) Value() *MainsVoltageDwellTripPoint { return a }
func (a MainsVoltageDwellTripPoint) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MainsVoltageDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MainsVoltageDwellTripPoint(*nt)
	return br, err
}

func (a MainsVoltageDwellTripPoint) Readable() bool   { return true }
func (a MainsVoltageDwellTripPoint) Writable() bool   { return true }
func (a MainsVoltageDwellTripPoint) Reportable() bool { return false }
func (a MainsVoltageDwellTripPoint) SceneIndex() int  { return -1 }

func (a MainsVoltageDwellTripPoint) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type BatteryVoltage zcl.Zu8

func (a BatteryVoltage) ID() zcl.AttrID          { return 32 }
func (a BatteryVoltage) Cluster() zcl.ClusterID  { return PowerConfigurationCluster }
func (a *BatteryVoltage) Value() *BatteryVoltage { return a }
func (a BatteryVoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltage(*nt)
	return br, err
}

func (a BatteryVoltage) Readable() bool   { return true }
func (a BatteryVoltage) Writable() bool   { return false }
func (a BatteryVoltage) Reportable() bool { return false }
func (a BatteryVoltage) SceneIndex() int  { return -1 }

func (a BatteryVoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryPercentageRemaining zcl.Zu8

func (a BatteryPercentageRemaining) ID() zcl.AttrID                      { return 33 }
func (a BatteryPercentageRemaining) Cluster() zcl.ClusterID              { return PowerConfigurationCluster }
func (a *BatteryPercentageRemaining) Value() *BatteryPercentageRemaining { return a }
func (a BatteryPercentageRemaining) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryPercentageRemaining) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageRemaining(*nt)
	return br, err
}

func (a BatteryPercentageRemaining) Readable() bool   { return true }
func (a BatteryPercentageRemaining) Writable() bool   { return false }
func (a BatteryPercentageRemaining) Reportable() bool { return true }
func (a BatteryPercentageRemaining) SceneIndex() int  { return -1 }

func (a BatteryPercentageRemaining) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryManufacturer zcl.Zcstring

func (a BatteryManufacturer) ID() zcl.AttrID               { return 48 }
func (a BatteryManufacturer) Cluster() zcl.ClusterID       { return PowerConfigurationCluster }
func (a *BatteryManufacturer) Value() *BatteryManufacturer { return a }
func (a BatteryManufacturer) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *BatteryManufacturer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryManufacturer(*nt)
	return br, err
}

func (a BatteryManufacturer) Readable() bool   { return true }
func (a BatteryManufacturer) Writable() bool   { return true }
func (a BatteryManufacturer) Reportable() bool { return false }
func (a BatteryManufacturer) SceneIndex() int  { return -1 }

func (a BatteryManufacturer) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

type BatterySize zcl.Zenum8

func (a BatterySize) ID() zcl.AttrID         { return 49 }
func (a BatterySize) Cluster() zcl.ClusterID { return PowerConfigurationCluster }
func (a *BatterySize) Value() *BatterySize   { return a }
func (a BatterySize) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *BatterySize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatterySize(*nt)
	return br, err
}

func (a BatterySize) Readable() bool   { return true }
func (a BatterySize) Writable() bool   { return true }
func (a BatterySize) Reportable() bool { return false }
func (a BatterySize) SceneIndex() int  { return -1 }

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

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsNoBattery checks if BatterySize equals the value for No Battery (0x00)
func (a BatterySize) IsNoBattery() bool { return a == 0x00 }

// SetNoBattery sets BatterySize to No Battery (0x00)
func (a *BatterySize) SetNoBattery() { *a = 0x00 }

// IsBuiltIn checks if BatterySize equals the value for Built in (0x01)
func (a BatterySize) IsBuiltIn() bool { return a == 0x01 }

// SetBuiltIn sets BatterySize to Built in (0x01)
func (a *BatterySize) SetBuiltIn() { *a = 0x01 }

// IsOther checks if BatterySize equals the value for Other (0x02)
func (a BatterySize) IsOther() bool { return a == 0x02 }

// SetOther sets BatterySize to Other (0x02)
func (a *BatterySize) SetOther() { *a = 0x02 }

// IsAa checks if BatterySize equals the value for AA (0x03)
func (a BatterySize) IsAa() bool { return a == 0x03 }

// SetAa sets BatterySize to AA (0x03)
func (a *BatterySize) SetAa() { *a = 0x03 }

// IsAaa checks if BatterySize equals the value for AAA (0x04)
func (a BatterySize) IsAaa() bool { return a == 0x04 }

// SetAaa sets BatterySize to AAA (0x04)
func (a *BatterySize) SetAaa() { *a = 0x04 }

// IsC checks if BatterySize equals the value for C (0x05)
func (a BatterySize) IsC() bool { return a == 0x05 }

// SetC sets BatterySize to C (0x05)
func (a *BatterySize) SetC() { *a = 0x05 }

// IsD checks if BatterySize equals the value for D (0x06)
func (a BatterySize) IsD() bool { return a == 0x06 }

// SetD sets BatterySize to D (0x06)
func (a *BatterySize) SetD() { *a = 0x06 }

// IsCr2 checks if BatterySize equals the value for CR2 (0x07)
func (a BatterySize) IsCr2() bool { return a == 0x07 }

// SetCr2 sets BatterySize to CR2 (0x07)
func (a *BatterySize) SetCr2() { *a = 0x07 }

// IsCr123a checks if BatterySize equals the value for CR123A (0x08)
func (a BatterySize) IsCr123a() bool { return a == 0x08 }

// SetCr123a sets BatterySize to CR123A (0x08)
func (a *BatterySize) SetCr123a() { *a = 0x08 }

// IsUnknown checks if BatterySize equals the value for Unknown (0xFF)
func (a BatterySize) IsUnknown() bool { return a == 0xFF }

// SetUnknown sets BatterySize to Unknown (0xFF)
func (a *BatterySize) SetUnknown() { *a = 0xFF }

type BatteryAhrRating zcl.Zu16

func (a BatteryAhrRating) ID() zcl.AttrID            { return 50 }
func (a BatteryAhrRating) Cluster() zcl.ClusterID    { return PowerConfigurationCluster }
func (a *BatteryAhrRating) Value() *BatteryAhrRating { return a }
func (a BatteryAhrRating) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *BatteryAhrRating) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryAhrRating(*nt)
	return br, err
}

func (a BatteryAhrRating) Readable() bool   { return true }
func (a BatteryAhrRating) Writable() bool   { return true }
func (a BatteryAhrRating) Reportable() bool { return false }
func (a BatteryAhrRating) SceneIndex() int  { return -1 }

func (a BatteryAhrRating) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type BatteryQuantity zcl.Zu8

func (a BatteryQuantity) ID() zcl.AttrID           { return 51 }
func (a BatteryQuantity) Cluster() zcl.ClusterID   { return PowerConfigurationCluster }
func (a *BatteryQuantity) Value() *BatteryQuantity { return a }
func (a BatteryQuantity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryQuantity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryQuantity(*nt)
	return br, err
}

func (a BatteryQuantity) Readable() bool   { return true }
func (a BatteryQuantity) Writable() bool   { return true }
func (a BatteryQuantity) Reportable() bool { return false }
func (a BatteryQuantity) SceneIndex() int  { return -1 }

func (a BatteryQuantity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryRatedVoltage zcl.Zu8

func (a BatteryRatedVoltage) ID() zcl.AttrID               { return 52 }
func (a BatteryRatedVoltage) Cluster() zcl.ClusterID       { return PowerConfigurationCluster }
func (a *BatteryRatedVoltage) Value() *BatteryRatedVoltage { return a }
func (a BatteryRatedVoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryRatedVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryRatedVoltage(*nt)
	return br, err
}

func (a BatteryRatedVoltage) Readable() bool   { return true }
func (a BatteryRatedVoltage) Writable() bool   { return true }
func (a BatteryRatedVoltage) Reportable() bool { return false }
func (a BatteryRatedVoltage) SceneIndex() int  { return -1 }

func (a BatteryRatedVoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryAlarmMask zcl.Zbmp8

func (a BatteryAlarmMask) ID() zcl.AttrID            { return 53 }
func (a BatteryAlarmMask) Cluster() zcl.ClusterID    { return PowerConfigurationCluster }
func (a *BatteryAlarmMask) Value() *BatteryAlarmMask { return a }
func (a BatteryAlarmMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *BatteryAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryAlarmMask(*nt)
	return br, err
}

func (a BatteryAlarmMask) Readable() bool   { return true }
func (a BatteryAlarmMask) Writable() bool   { return true }
func (a BatteryAlarmMask) Reportable() bool { return false }
func (a BatteryAlarmMask) SceneIndex() int  { return -1 }

func (a BatteryAlarmMask) String() string {

	var bstr []string
	if a.IsBatteryVoltageTooLow() {
		bstr = append(bstr, "Battery Voltage too low")
	}
	if a.IsBatteryAlarm1() {
		bstr = append(bstr, "Battery Alarm 1")
	}
	if a.IsBatteryAlarm2() {
		bstr = append(bstr, "Battery Alarm 2")
	}
	if a.IsBatteryAlarm3() {
		bstr = append(bstr, "Battery Alarm 3")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a BatteryAlarmMask) IsBatteryVoltageTooLow() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *BatteryAlarmMask) SetBatteryVoltageTooLow(b bool) {
	*a = BatteryAlarmMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a BatteryAlarmMask) IsBatteryAlarm1() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *BatteryAlarmMask) SetBatteryAlarm1(b bool) {
	*a = BatteryAlarmMask(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a BatteryAlarmMask) IsBatteryAlarm2() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *BatteryAlarmMask) SetBatteryAlarm2(b bool) {
	*a = BatteryAlarmMask(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a BatteryAlarmMask) IsBatteryAlarm3() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *BatteryAlarmMask) SetBatteryAlarm3(b bool) {
	*a = BatteryAlarmMask(zcl.BitmapSet([]byte(*a), 3, b))
}

type BatteryVoltageMinThreshold zcl.Zu8

func (a BatteryVoltageMinThreshold) ID() zcl.AttrID                      { return 54 }
func (a BatteryVoltageMinThreshold) Cluster() zcl.ClusterID              { return PowerConfigurationCluster }
func (a *BatteryVoltageMinThreshold) Value() *BatteryVoltageMinThreshold { return a }
func (a BatteryVoltageMinThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryVoltageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageMinThreshold(*nt)
	return br, err
}

func (a BatteryVoltageMinThreshold) Readable() bool   { return true }
func (a BatteryVoltageMinThreshold) Writable() bool   { return true }
func (a BatteryVoltageMinThreshold) Reportable() bool { return false }
func (a BatteryVoltageMinThreshold) SceneIndex() int  { return -1 }

func (a BatteryVoltageMinThreshold) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryVoltageThreshold1 zcl.Zu8

func (a BatteryVoltageThreshold1) ID() zcl.AttrID                    { return 55 }
func (a BatteryVoltageThreshold1) Cluster() zcl.ClusterID            { return PowerConfigurationCluster }
func (a *BatteryVoltageThreshold1) Value() *BatteryVoltageThreshold1 { return a }
func (a BatteryVoltageThreshold1) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryVoltageThreshold1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageThreshold1(*nt)
	return br, err
}

func (a BatteryVoltageThreshold1) Readable() bool   { return true }
func (a BatteryVoltageThreshold1) Writable() bool   { return true }
func (a BatteryVoltageThreshold1) Reportable() bool { return false }
func (a BatteryVoltageThreshold1) SceneIndex() int  { return -1 }

func (a BatteryVoltageThreshold1) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryVoltageThreshold2 zcl.Zu8

func (a BatteryVoltageThreshold2) ID() zcl.AttrID                    { return 56 }
func (a BatteryVoltageThreshold2) Cluster() zcl.ClusterID            { return PowerConfigurationCluster }
func (a *BatteryVoltageThreshold2) Value() *BatteryVoltageThreshold2 { return a }
func (a BatteryVoltageThreshold2) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryVoltageThreshold2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageThreshold2(*nt)
	return br, err
}

func (a BatteryVoltageThreshold2) Readable() bool   { return true }
func (a BatteryVoltageThreshold2) Writable() bool   { return true }
func (a BatteryVoltageThreshold2) Reportable() bool { return false }
func (a BatteryVoltageThreshold2) SceneIndex() int  { return -1 }

func (a BatteryVoltageThreshold2) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryVoltageThreshold3 zcl.Zu8

func (a BatteryVoltageThreshold3) ID() zcl.AttrID                    { return 57 }
func (a BatteryVoltageThreshold3) Cluster() zcl.ClusterID            { return PowerConfigurationCluster }
func (a *BatteryVoltageThreshold3) Value() *BatteryVoltageThreshold3 { return a }
func (a BatteryVoltageThreshold3) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryVoltageThreshold3) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryVoltageThreshold3(*nt)
	return br, err
}

func (a BatteryVoltageThreshold3) Readable() bool   { return true }
func (a BatteryVoltageThreshold3) Writable() bool   { return true }
func (a BatteryVoltageThreshold3) Reportable() bool { return false }
func (a BatteryVoltageThreshold3) SceneIndex() int  { return -1 }

func (a BatteryVoltageThreshold3) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryPercentageMinThreshold zcl.Zu8

func (a BatteryPercentageMinThreshold) ID() zcl.AttrID                         { return 58 }
func (a BatteryPercentageMinThreshold) Cluster() zcl.ClusterID                 { return PowerConfigurationCluster }
func (a *BatteryPercentageMinThreshold) Value() *BatteryPercentageMinThreshold { return a }
func (a BatteryPercentageMinThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryPercentageMinThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageMinThreshold(*nt)
	return br, err
}

func (a BatteryPercentageMinThreshold) Readable() bool   { return true }
func (a BatteryPercentageMinThreshold) Writable() bool   { return true }
func (a BatteryPercentageMinThreshold) Reportable() bool { return false }
func (a BatteryPercentageMinThreshold) SceneIndex() int  { return -1 }

func (a BatteryPercentageMinThreshold) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryPercentageThreshold1 zcl.Zu8

func (a BatteryPercentageThreshold1) ID() zcl.AttrID                       { return 59 }
func (a BatteryPercentageThreshold1) Cluster() zcl.ClusterID               { return PowerConfigurationCluster }
func (a *BatteryPercentageThreshold1) Value() *BatteryPercentageThreshold1 { return a }
func (a BatteryPercentageThreshold1) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryPercentageThreshold1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageThreshold1(*nt)
	return br, err
}

func (a BatteryPercentageThreshold1) Readable() bool   { return true }
func (a BatteryPercentageThreshold1) Writable() bool   { return true }
func (a BatteryPercentageThreshold1) Reportable() bool { return false }
func (a BatteryPercentageThreshold1) SceneIndex() int  { return -1 }

func (a BatteryPercentageThreshold1) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryPercentageThreshold2 zcl.Zu8

func (a BatteryPercentageThreshold2) ID() zcl.AttrID                       { return 60 }
func (a BatteryPercentageThreshold2) Cluster() zcl.ClusterID               { return PowerConfigurationCluster }
func (a *BatteryPercentageThreshold2) Value() *BatteryPercentageThreshold2 { return a }
func (a BatteryPercentageThreshold2) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryPercentageThreshold2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageThreshold2(*nt)
	return br, err
}

func (a BatteryPercentageThreshold2) Readable() bool   { return true }
func (a BatteryPercentageThreshold2) Writable() bool   { return true }
func (a BatteryPercentageThreshold2) Reportable() bool { return false }
func (a BatteryPercentageThreshold2) SceneIndex() int  { return -1 }

func (a BatteryPercentageThreshold2) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryPercentageThreshold3 zcl.Zu8

func (a BatteryPercentageThreshold3) ID() zcl.AttrID                       { return 61 }
func (a BatteryPercentageThreshold3) Cluster() zcl.ClusterID               { return PowerConfigurationCluster }
func (a *BatteryPercentageThreshold3) Value() *BatteryPercentageThreshold3 { return a }
func (a BatteryPercentageThreshold3) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *BatteryPercentageThreshold3) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryPercentageThreshold3(*nt)
	return br, err
}

func (a BatteryPercentageThreshold3) Readable() bool   { return true }
func (a BatteryPercentageThreshold3) Writable() bool   { return true }
func (a BatteryPercentageThreshold3) Reportable() bool { return false }
func (a BatteryPercentageThreshold3) SceneIndex() int  { return -1 }

func (a BatteryPercentageThreshold3) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type BatteryAlarmState zcl.Zbmp32

func (a BatteryAlarmState) ID() zcl.AttrID             { return 62 }
func (a BatteryAlarmState) Cluster() zcl.ClusterID     { return PowerConfigurationCluster }
func (a *BatteryAlarmState) Value() *BatteryAlarmState { return a }
func (a BatteryAlarmState) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp32(a).MarshalZcl()
}
func (a *BatteryAlarmState) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*a = BatteryAlarmState(*nt)
	return br, err
}

func (a BatteryAlarmState) Readable() bool   { return true }
func (a BatteryAlarmState) Writable() bool   { return true }
func (a BatteryAlarmState) Reportable() bool { return false }
func (a BatteryAlarmState) SceneIndex() int  { return -1 }

func (a BatteryAlarmState) String() string {

	var bstr []string
	if a.IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource1() {
		bstr = append(bstr, "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 1")
	}
	if a.IsThreshold1ForVoltageOrPercentageReachedForBatterySource1() {
		bstr = append(bstr, "Threshold 1 for Voltage or Percentage reached for Battery Source 1")
	}
	if a.IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource2() {
		bstr = append(bstr, "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 2")
	}
	if a.IsThreshold1ForVoltageOrPercentageReachedForBatterySource2() {
		bstr = append(bstr, "Threshold 1 for Voltage or Percentage reached for Battery Source 2")
	}
	if a.IsThreshold2ForVoltageOrPercentageReachedForBatterySource2() {
		bstr = append(bstr, "Threshold 2 for Voltage or Percentage reached for Battery Source 2")
	}
	if a.IsThreshold3ForVoltageOrPercentageReachedForBatterySource2() {
		bstr = append(bstr, "Threshold 3 for Voltage or Percentage reached for Battery Source 2")
	}
	if a.IsThreshold2ForVoltageOrPercentageReachedForBatterySource1() {
		bstr = append(bstr, "Threshold 2 for Voltage or Percentage reached for Battery Source 1")
	}
	if a.IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource3() {
		bstr = append(bstr, "Minimum Threshold 1 for Voltage or Percentage reached for Battery Source 3")
	}
	if a.IsThreshold1ForVoltageOrPercentageReachedForBatterySource3() {
		bstr = append(bstr, "Threshold 1 for Voltage or Percentage reached for Battery Source 3")
	}
	if a.IsThreshold2ForVoltageOrPercentageReachedForBatterySource3() {
		bstr = append(bstr, "Threshold 2 for Voltage or Percentage reached for Battery Source 3")
	}
	if a.IsThreshold3ForVoltageOrPercentageReachedForBatterySource3() {
		bstr = append(bstr, "Threshold 3 for Voltage or Percentage reached for Battery Source 3")
	}
	if a.IsThreshold3ForVoltageOrPercentageReachedForBatterySource1() {
		bstr = append(bstr, "Threshold 3 for Voltage or Percentage reached for Battery Source 1")
	}
	if a.IsMainsPowerLostUnavailable() {
		bstr = append(bstr, "Mains power lost/unavailable")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a), 10)
}
func (a *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 10, b))
}

func (a BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a), 11)
}
func (a *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 11, b))
}

func (a BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a), 12)
}
func (a *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 12, b))
}

func (a BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource2() bool {
	return zcl.BitmapTest([]byte(a), 13)
}
func (a *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource2(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 13, b))
}

func (a BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a BatteryAlarmState) IsMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a), 20)
}
func (a *BatteryAlarmState) SetMinimumThreshold1ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 20, b))
}

func (a BatteryAlarmState) IsThreshold1ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a), 21)
}
func (a *BatteryAlarmState) SetThreshold1ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 21, b))
}

func (a BatteryAlarmState) IsThreshold2ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a), 22)
}
func (a *BatteryAlarmState) SetThreshold2ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 22, b))
}

func (a BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource3() bool {
	return zcl.BitmapTest([]byte(a), 23)
}
func (a *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource3(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 23, b))
}

func (a BatteryAlarmState) IsThreshold3ForVoltageOrPercentageReachedForBatterySource1() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *BatteryAlarmState) SetThreshold3ForVoltageOrPercentageReachedForBatterySource1(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a BatteryAlarmState) IsMainsPowerLostUnavailable() bool {
	return zcl.BitmapTest([]byte(a), 30)
}
func (a *BatteryAlarmState) SetMainsPowerLostUnavailable(b bool) {
	*a = BatteryAlarmState(zcl.BitmapSet([]byte(*a), 30, b))
}
