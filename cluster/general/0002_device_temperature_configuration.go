// Attributes for determining information about a device’s internal temperature, and for configuring under/over temperature alarms.
package general

import (
	"neotor.se/zcl/cluster/zcl"
)

// DeviceTemperatureConfiguration
// Attributes for determining information about a device’s internal temperature, and for configuring under/over temperature alarms.

func NewDeviceTemperatureConfigurationServer(profile zcl.ProfileID) *DeviceTemperatureConfigurationServer {
	return &DeviceTemperatureConfigurationServer{p: profile}
}
func NewDeviceTemperatureConfigurationClient(profile zcl.ProfileID) *DeviceTemperatureConfigurationClient {
	return &DeviceTemperatureConfigurationClient{p: profile}
}

const DeviceTemperatureConfigurationCluster zcl.ClusterID = 2

type DeviceTemperatureConfigurationServer struct {
	p zcl.ProfileID

	CurrentTemperature     *CurrentTemperature
	MinTempExperienced     *MinTempExperienced
	MaxTempExperienced     *MaxTempExperienced
	OverTempTotalDwell     *OverTempTotalDwell
	DeviceTempAlarmMask    *DeviceTempAlarmMask
	LowTempThreshold       *LowTempThreshold
	HighTempThreshold      *HighTempThreshold
	LowTempDwellTripPoint  *LowTempDwellTripPoint
	HighTempDwellTripPoint *HighTempDwellTripPoint
}

type DeviceTemperatureConfigurationClient struct {
	p zcl.ProfileID
}

/*
var DeviceTemperatureConfigurationServer = map[zcl.CommandID]func() zcl.Command{
}

var DeviceTemperatureConfigurationClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type CurrentTemperature zcl.Zs16

func (a CurrentTemperature) ID() zcl.AttrID              { return 0 }
func (a CurrentTemperature) Cluster() zcl.ClusterID      { return DeviceTemperatureConfigurationCluster }
func (a *CurrentTemperature) Value() *CurrentTemperature { return a }
func (a CurrentTemperature) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *CurrentTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentTemperature(*nt)
	return br, err
}

func (a CurrentTemperature) Readable() bool   { return true }
func (a CurrentTemperature) Writable() bool   { return false }
func (a CurrentTemperature) Reportable() bool { return false }
func (a CurrentTemperature) SceneIndex() int  { return -1 }

func (a CurrentTemperature) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MinTempExperienced zcl.Zs16

func (a MinTempExperienced) ID() zcl.AttrID              { return 1 }
func (a MinTempExperienced) Cluster() zcl.ClusterID      { return DeviceTemperatureConfigurationCluster }
func (a *MinTempExperienced) Value() *MinTempExperienced { return a }
func (a MinTempExperienced) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MinTempExperienced) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinTempExperienced(*nt)
	return br, err
}

func (a MinTempExperienced) Readable() bool   { return true }
func (a MinTempExperienced) Writable() bool   { return false }
func (a MinTempExperienced) Reportable() bool { return false }
func (a MinTempExperienced) SceneIndex() int  { return -1 }

func (a MinTempExperienced) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MaxTempExperienced zcl.Zs16

func (a MaxTempExperienced) ID() zcl.AttrID              { return 2 }
func (a MaxTempExperienced) Cluster() zcl.ClusterID      { return DeviceTemperatureConfigurationCluster }
func (a *MaxTempExperienced) Value() *MaxTempExperienced { return a }
func (a MaxTempExperienced) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MaxTempExperienced) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxTempExperienced(*nt)
	return br, err
}

func (a MaxTempExperienced) Readable() bool   { return true }
func (a MaxTempExperienced) Writable() bool   { return false }
func (a MaxTempExperienced) Reportable() bool { return false }
func (a MaxTempExperienced) SceneIndex() int  { return -1 }

func (a MaxTempExperienced) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type OverTempTotalDwell zcl.Zu16

func (a OverTempTotalDwell) ID() zcl.AttrID              { return 3 }
func (a OverTempTotalDwell) Cluster() zcl.ClusterID      { return DeviceTemperatureConfigurationCluster }
func (a *OverTempTotalDwell) Value() *OverTempTotalDwell { return a }
func (a OverTempTotalDwell) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *OverTempTotalDwell) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OverTempTotalDwell(*nt)
	return br, err
}

func (a OverTempTotalDwell) Readable() bool   { return false }
func (a OverTempTotalDwell) Writable() bool   { return false }
func (a OverTempTotalDwell) Reportable() bool { return false }
func (a OverTempTotalDwell) SceneIndex() int  { return -1 }

func (a OverTempTotalDwell) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type DeviceTempAlarmMask zcl.Zbmp8

func (a DeviceTempAlarmMask) ID() zcl.AttrID               { return 16 }
func (a DeviceTempAlarmMask) Cluster() zcl.ClusterID       { return DeviceTemperatureConfigurationCluster }
func (a *DeviceTempAlarmMask) Value() *DeviceTempAlarmMask { return a }
func (a DeviceTempAlarmMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *DeviceTempAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceTempAlarmMask(*nt)
	return br, err
}

func (a DeviceTempAlarmMask) Readable() bool   { return true }
func (a DeviceTempAlarmMask) Writable() bool   { return true }
func (a DeviceTempAlarmMask) Reportable() bool { return false }
func (a DeviceTempAlarmMask) SceneIndex() int  { return -1 }

func (a DeviceTempAlarmMask) String() string {

	var bstr []string
	if a.IsTemperatureTooLow() {
		bstr = append(bstr, "Temperature too low")
	}
	if a.IsTemperatureTooHigh() {
		bstr = append(bstr, "Temperature too high")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a DeviceTempAlarmMask) IsTemperatureTooLow() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *DeviceTempAlarmMask) SetTemperatureTooLow(b bool) {
	*a = DeviceTempAlarmMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a DeviceTempAlarmMask) IsTemperatureTooHigh() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *DeviceTempAlarmMask) SetTemperatureTooHigh(b bool) {
	*a = DeviceTempAlarmMask(zcl.BitmapSet([]byte(*a), 1, b))
}

type LowTempThreshold zcl.Zs16

func (a LowTempThreshold) ID() zcl.AttrID            { return 17 }
func (a LowTempThreshold) Cluster() zcl.ClusterID    { return DeviceTemperatureConfigurationCluster }
func (a *LowTempThreshold) Value() *LowTempThreshold { return a }
func (a LowTempThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *LowTempThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = LowTempThreshold(*nt)
	return br, err
}

func (a LowTempThreshold) Readable() bool   { return true }
func (a LowTempThreshold) Writable() bool   { return true }
func (a LowTempThreshold) Reportable() bool { return false }
func (a LowTempThreshold) SceneIndex() int  { return -1 }

func (a LowTempThreshold) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type HighTempThreshold zcl.Zs16

func (a HighTempThreshold) ID() zcl.AttrID             { return 18 }
func (a HighTempThreshold) Cluster() zcl.ClusterID     { return DeviceTemperatureConfigurationCluster }
func (a *HighTempThreshold) Value() *HighTempThreshold { return a }
func (a HighTempThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *HighTempThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = HighTempThreshold(*nt)
	return br, err
}

func (a HighTempThreshold) Readable() bool   { return true }
func (a HighTempThreshold) Writable() bool   { return true }
func (a HighTempThreshold) Reportable() bool { return false }
func (a HighTempThreshold) SceneIndex() int  { return -1 }

func (a HighTempThreshold) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type LowTempDwellTripPoint zcl.Zu24

func (a LowTempDwellTripPoint) ID() zcl.AttrID                 { return 19 }
func (a LowTempDwellTripPoint) Cluster() zcl.ClusterID         { return DeviceTemperatureConfigurationCluster }
func (a *LowTempDwellTripPoint) Value() *LowTempDwellTripPoint { return a }
func (a LowTempDwellTripPoint) MarshalZcl() ([]byte, error) {
	return zcl.Zu24(a).MarshalZcl()
}
func (a *LowTempDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = LowTempDwellTripPoint(*nt)
	return br, err
}

func (a LowTempDwellTripPoint) Readable() bool   { return false }
func (a LowTempDwellTripPoint) Writable() bool   { return false }
func (a LowTempDwellTripPoint) Reportable() bool { return false }
func (a LowTempDwellTripPoint) SceneIndex() int  { return -1 }

func (a LowTempDwellTripPoint) String() string {

	return zcl.Sprintf("%s", zcl.Zu24(a))
}

type HighTempDwellTripPoint zcl.Zu24

func (a HighTempDwellTripPoint) ID() zcl.AttrID                  { return 20 }
func (a HighTempDwellTripPoint) Cluster() zcl.ClusterID          { return DeviceTemperatureConfigurationCluster }
func (a *HighTempDwellTripPoint) Value() *HighTempDwellTripPoint { return a }
func (a HighTempDwellTripPoint) MarshalZcl() ([]byte, error) {
	return zcl.Zu24(a).MarshalZcl()
}
func (a *HighTempDwellTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = HighTempDwellTripPoint(*nt)
	return br, err
}

func (a HighTempDwellTripPoint) Readable() bool   { return false }
func (a HighTempDwellTripPoint) Writable() bool   { return false }
func (a HighTempDwellTripPoint) Reportable() bool { return false }
func (a HighTempDwellTripPoint) SceneIndex() int  { return -1 }

func (a HighTempDwellTripPoint) String() string {

	return zcl.Sprintf("%s", zcl.Zu24(a))
}
