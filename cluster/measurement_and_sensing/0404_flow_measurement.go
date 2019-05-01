package measurement_and_sensing

import (
	"neotor.se/zcl/cluster/zcl"
)

// FlowMeasurement

func NewFlowMeasurementServer(profile zcl.ProfileID) *FlowMeasurementServer {
	return &FlowMeasurementServer{p: profile}
}
func NewFlowMeasurementClient(profile zcl.ProfileID) *FlowMeasurementClient {
	return &FlowMeasurementClient{p: profile}
}

const FlowMeasurementCluster zcl.ClusterID = 1028

type FlowMeasurementServer struct {
	p zcl.ProfileID

	MeasuredFlow    *MeasuredFlow
	MinMeasuredFlow *MinMeasuredFlow
	MaxMeasuredFlow *MaxMeasuredFlow
	FlowTolerance   *FlowTolerance
}

type FlowMeasurementClient struct {
	p zcl.ProfileID
}

/*
var FlowMeasurementServer = map[zcl.CommandID]func() zcl.Command{
}

var FlowMeasurementClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type MeasuredFlow zcl.Zu16

func (a MeasuredFlow) ID() zcl.AttrID         { return 0 }
func (a MeasuredFlow) Cluster() zcl.ClusterID { return FlowMeasurementCluster }
func (a *MeasuredFlow) Value() *MeasuredFlow  { return a }
func (a MeasuredFlow) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredFlow(*nt)
	return br, err
}

func (a MeasuredFlow) Readable() bool   { return true }
func (a MeasuredFlow) Writable() bool   { return false }
func (a MeasuredFlow) Reportable() bool { return true }
func (a MeasuredFlow) SceneIndex() int  { return -1 }

func (a MeasuredFlow) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type MinMeasuredFlow zcl.Zu16

func (a MinMeasuredFlow) ID() zcl.AttrID           { return 1 }
func (a MinMeasuredFlow) Cluster() zcl.ClusterID   { return FlowMeasurementCluster }
func (a *MinMeasuredFlow) Value() *MinMeasuredFlow { return a }
func (a MinMeasuredFlow) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MinMeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinMeasuredFlow(*nt)
	return br, err
}

func (a MinMeasuredFlow) Readable() bool   { return true }
func (a MinMeasuredFlow) Writable() bool   { return false }
func (a MinMeasuredFlow) Reportable() bool { return false }
func (a MinMeasuredFlow) SceneIndex() int  { return -1 }

func (a MinMeasuredFlow) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type MaxMeasuredFlow zcl.Zu16

func (a MaxMeasuredFlow) ID() zcl.AttrID           { return 2 }
func (a MaxMeasuredFlow) Cluster() zcl.ClusterID   { return FlowMeasurementCluster }
func (a *MaxMeasuredFlow) Value() *MaxMeasuredFlow { return a }
func (a MaxMeasuredFlow) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MaxMeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxMeasuredFlow(*nt)
	return br, err
}

func (a MaxMeasuredFlow) Readable() bool   { return true }
func (a MaxMeasuredFlow) Writable() bool   { return false }
func (a MaxMeasuredFlow) Reportable() bool { return false }
func (a MaxMeasuredFlow) SceneIndex() int  { return -1 }

func (a MaxMeasuredFlow) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type FlowTolerance zcl.Zu16

func (a FlowTolerance) ID() zcl.AttrID         { return 3 }
func (a FlowTolerance) Cluster() zcl.ClusterID { return FlowMeasurementCluster }
func (a *FlowTolerance) Value() *FlowTolerance { return a }
func (a FlowTolerance) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *FlowTolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = FlowTolerance(*nt)
	return br, err
}

func (a FlowTolerance) Readable() bool   { return true }
func (a FlowTolerance) Writable() bool   { return false }
func (a FlowTolerance) Reportable() bool { return true }
func (a FlowTolerance) SceneIndex() int  { return -1 }

func (a FlowTolerance) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}
