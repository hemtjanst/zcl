// This cluster provides an interface to the functionality of Smart Energy Demand Response and Load Control. Devices targeted by this cluster include thermostats and devices that support load control.
package smart_energy

import (
	"neotor.se/zcl"
)

// DemandResponseAndLoadControl
const DemandResponseAndLoadControlID zcl.ClusterID = 1793

var DemandResponseAndLoadControlCluster = zcl.Cluster{
	ServerCmd:  map[zcl.CommandID]func() zcl.Command{},
	ClientCmd:  map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{
		UtilityEnrolmentGroupAttr: func() zcl.Attr { return new(UtilityEnrolmentGroup) },
		StartRandomizeMinutesAttr: func() zcl.Attr { return new(StartRandomizeMinutes) },
		StopRandomizeMinutesAttr:  func() zcl.Attr { return new(StopRandomizeMinutes) },
		DeviceClassValueAttr:      func() zcl.Attr { return new(DeviceClassValue) },
	},
	SceneAttr: []zcl.AttrID{},
}

// UtilityEnrolmentGroup is an autogenerated attribute in the DemandResponseAndLoadControl cluster
type UtilityEnrolmentGroup zcl.Zu8

const UtilityEnrolmentGroupAttr zcl.AttrID = 0

func (UtilityEnrolmentGroup) ID() zcl.AttrID                   { return UtilityEnrolmentGroupAttr }
func (UtilityEnrolmentGroup) Cluster() zcl.ClusterID           { return DemandResponseAndLoadControlID }
func (UtilityEnrolmentGroup) Name() string                     { return "Utility Enrolment Group" }
func (UtilityEnrolmentGroup) Readable() bool                   { return true }
func (UtilityEnrolmentGroup) Writable() bool                   { return true }
func (UtilityEnrolmentGroup) Reportable() bool                 { return false }
func (UtilityEnrolmentGroup) SceneIndex() int                  { return -1 }
func (a *UtilityEnrolmentGroup) Value() *UtilityEnrolmentGroup { return a }
func (a UtilityEnrolmentGroup) MarshalZcl() ([]byte, error)    { return zcl.Zu8(a).MarshalZcl() }

func (a *UtilityEnrolmentGroup) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = UtilityEnrolmentGroup(*nt)
	return br, err
}

func (a UtilityEnrolmentGroup) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// StartRandomizeMinutes is an autogenerated attribute in the DemandResponseAndLoadControl cluster
type StartRandomizeMinutes zcl.Zu8

const StartRandomizeMinutesAttr zcl.AttrID = 1

func (StartRandomizeMinutes) ID() zcl.AttrID                   { return StartRandomizeMinutesAttr }
func (StartRandomizeMinutes) Cluster() zcl.ClusterID           { return DemandResponseAndLoadControlID }
func (StartRandomizeMinutes) Name() string                     { return "Start Randomize Minutes" }
func (StartRandomizeMinutes) Readable() bool                   { return true }
func (StartRandomizeMinutes) Writable() bool                   { return true }
func (StartRandomizeMinutes) Reportable() bool                 { return false }
func (StartRandomizeMinutes) SceneIndex() int                  { return -1 }
func (a *StartRandomizeMinutes) Value() *StartRandomizeMinutes { return a }
func (a StartRandomizeMinutes) MarshalZcl() ([]byte, error)    { return zcl.Zu8(a).MarshalZcl() }

func (a *StartRandomizeMinutes) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StartRandomizeMinutes(*nt)
	return br, err
}

func (a StartRandomizeMinutes) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// StopRandomizeMinutes is an autogenerated attribute in the DemandResponseAndLoadControl cluster
type StopRandomizeMinutes zcl.Zu8

const StopRandomizeMinutesAttr zcl.AttrID = 2

func (StopRandomizeMinutes) ID() zcl.AttrID                  { return StopRandomizeMinutesAttr }
func (StopRandomizeMinutes) Cluster() zcl.ClusterID          { return DemandResponseAndLoadControlID }
func (StopRandomizeMinutes) Name() string                    { return "Stop Randomize Minutes" }
func (StopRandomizeMinutes) Readable() bool                  { return true }
func (StopRandomizeMinutes) Writable() bool                  { return true }
func (StopRandomizeMinutes) Reportable() bool                { return false }
func (StopRandomizeMinutes) SceneIndex() int                 { return -1 }
func (a *StopRandomizeMinutes) Value() *StopRandomizeMinutes { return a }
func (a StopRandomizeMinutes) MarshalZcl() ([]byte, error)   { return zcl.Zu8(a).MarshalZcl() }

func (a *StopRandomizeMinutes) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StopRandomizeMinutes(*nt)
	return br, err
}

func (a StopRandomizeMinutes) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// DeviceClassValue is an autogenerated attribute in the DemandResponseAndLoadControl cluster
type DeviceClassValue zcl.Zu16

const DeviceClassValueAttr zcl.AttrID = 3

func (DeviceClassValue) ID() zcl.AttrID                { return DeviceClassValueAttr }
func (DeviceClassValue) Cluster() zcl.ClusterID        { return DemandResponseAndLoadControlID }
func (DeviceClassValue) Name() string                  { return "Device Class Value" }
func (DeviceClassValue) Readable() bool                { return true }
func (DeviceClassValue) Writable() bool                { return false }
func (DeviceClassValue) Reportable() bool              { return false }
func (DeviceClassValue) SceneIndex() int               { return -1 }
func (a *DeviceClassValue) Value() *DeviceClassValue   { return a }
func (a DeviceClassValue) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *DeviceClassValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DeviceClassValue(*nt)
	return br, err
}

func (a DeviceClassValue) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}
