// This cluster provides an interface to the functionality of Smart Energy Demand Response and Load Control. Devices targeted by this cluster include thermostats and devices that support load control.
package smart_energy

import (
	"neotor.se/zcl/cluster/zcl"
)

// DemandResponseAndLoadControl
// This cluster provides an interface to the functionality of Smart Energy Demand Response and Load Control. Devices targeted by this cluster include thermostats and devices that support load control.

func NewDemandResponseAndLoadControlServer(profile zcl.ProfileID) *DemandResponseAndLoadControlServer {
	return &DemandResponseAndLoadControlServer{p: profile}
}
func NewDemandResponseAndLoadControlClient(profile zcl.ProfileID) *DemandResponseAndLoadControlClient {
	return &DemandResponseAndLoadControlClient{p: profile}
}

const DemandResponseAndLoadControlCluster zcl.ClusterID = 1793

type DemandResponseAndLoadControlServer struct {
	p zcl.ProfileID
}

type DemandResponseAndLoadControlClient struct {
	p zcl.ProfileID

	UtilityEnrolmentGroup *UtilityEnrolmentGroup
	StartRandomizeMinutes *StartRandomizeMinutes
	StopRandomizeMinutes  *StopRandomizeMinutes
	DeviceClassValue      *DeviceClassValue
}

/*
var DemandResponseAndLoadControlServer = map[zcl.CommandID]func() zcl.Command{
}

var DemandResponseAndLoadControlClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type UtilityEnrolmentGroup zcl.Zu8

func (a UtilityEnrolmentGroup) ID() zcl.AttrID                 { return 0 }
func (a UtilityEnrolmentGroup) Cluster() zcl.ClusterID         { return DemandResponseAndLoadControlCluster }
func (a *UtilityEnrolmentGroup) Value() *UtilityEnrolmentGroup { return a }
func (a UtilityEnrolmentGroup) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *UtilityEnrolmentGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = UtilityEnrolmentGroup(*nt)
	return br, err
}

func (a UtilityEnrolmentGroup) Readable() bool   { return true }
func (a UtilityEnrolmentGroup) Writable() bool   { return true }
func (a UtilityEnrolmentGroup) Reportable() bool { return false }
func (a UtilityEnrolmentGroup) SceneIndex() int  { return -1 }

func (a UtilityEnrolmentGroup) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type StartRandomizeMinutes zcl.Zu8

func (a StartRandomizeMinutes) ID() zcl.AttrID                 { return 1 }
func (a StartRandomizeMinutes) Cluster() zcl.ClusterID         { return DemandResponseAndLoadControlCluster }
func (a *StartRandomizeMinutes) Value() *StartRandomizeMinutes { return a }
func (a StartRandomizeMinutes) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *StartRandomizeMinutes) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StartRandomizeMinutes(*nt)
	return br, err
}

func (a StartRandomizeMinutes) Readable() bool   { return true }
func (a StartRandomizeMinutes) Writable() bool   { return true }
func (a StartRandomizeMinutes) Reportable() bool { return false }
func (a StartRandomizeMinutes) SceneIndex() int  { return -1 }

func (a StartRandomizeMinutes) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type StopRandomizeMinutes zcl.Zu8

func (a StopRandomizeMinutes) ID() zcl.AttrID                { return 2 }
func (a StopRandomizeMinutes) Cluster() zcl.ClusterID        { return DemandResponseAndLoadControlCluster }
func (a *StopRandomizeMinutes) Value() *StopRandomizeMinutes { return a }
func (a StopRandomizeMinutes) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *StopRandomizeMinutes) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StopRandomizeMinutes(*nt)
	return br, err
}

func (a StopRandomizeMinutes) Readable() bool   { return true }
func (a StopRandomizeMinutes) Writable() bool   { return true }
func (a StopRandomizeMinutes) Reportable() bool { return false }
func (a StopRandomizeMinutes) SceneIndex() int  { return -1 }

func (a StopRandomizeMinutes) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type DeviceClassValue zcl.Zu16

func (a DeviceClassValue) ID() zcl.AttrID            { return 3 }
func (a DeviceClassValue) Cluster() zcl.ClusterID    { return DemandResponseAndLoadControlCluster }
func (a *DeviceClassValue) Value() *DeviceClassValue { return a }
func (a DeviceClassValue) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DeviceClassValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceClassValue(*nt)
	return br, err
}

func (a DeviceClassValue) Readable() bool   { return true }
func (a DeviceClassValue) Writable() bool   { return false }
func (a DeviceClassValue) Reportable() bool { return false }
func (a DeviceClassValue) SceneIndex() int  { return -1 }

func (a DeviceClassValue) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}
