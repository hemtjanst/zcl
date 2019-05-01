// The server cluster provides an interface to air pressure measurement functionality, including configuration and provision of notifications of air pressure measurements.
package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// PressureMeasurement
// The server cluster provides an interface to air pressure measurement functionality, including configuration and provision of notifications of air pressure measurements.

func NewPressureMeasurementServer(profile zcl.ProfileID) *PressureMeasurementServer {
	return &PressureMeasurementServer{p: profile}
}
func NewPressureMeasurementClient(profile zcl.ProfileID) *PressureMeasurementClient {
	return &PressureMeasurementClient{p: profile}
}

const PressureMeasurementCluster zcl.ClusterID = 1027

type PressureMeasurementServer struct {
	p zcl.ProfileID

	MeasuredPressure        *MeasuredPressure
	MinMeasuredPressure     *MinMeasuredPressure
	MaxMeasuredPressure     *MaxMeasuredPressure
	PressureTolerance       *PressureTolerance
	ScaledPressure          *ScaledPressure
	MinScaledPressure       *MinScaledPressure
	MaxScaledPressure       *MaxScaledPressure
	ScaledPressureTolerance *ScaledPressureTolerance
	Scale                   *Scale
}

type PressureMeasurementClient struct {
	p zcl.ProfileID
}

/*
var PressureMeasurementServer = map[zcl.CommandID]func() zcl.Command{
}

var PressureMeasurementClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type MeasuredPressure zcl.Zs16

func (a MeasuredPressure) ID() zcl.AttrID            { return 0 }
func (a MeasuredPressure) Cluster() zcl.ClusterID    { return PressureMeasurementCluster }
func (a *MeasuredPressure) Value() *MeasuredPressure { return a }
func (a MeasuredPressure) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPressure(*nt)
	return br, err
}

func (a MeasuredPressure) Readable() bool   { return true }
func (a MeasuredPressure) Writable() bool   { return false }
func (a MeasuredPressure) Reportable() bool { return true }
func (a MeasuredPressure) SceneIndex() int  { return -1 }

func (a MeasuredPressure) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MinMeasuredPressure zcl.Zs16

func (a MinMeasuredPressure) ID() zcl.AttrID               { return 1 }
func (a MinMeasuredPressure) Cluster() zcl.ClusterID       { return PressureMeasurementCluster }
func (a *MinMeasuredPressure) Value() *MinMeasuredPressure { return a }
func (a MinMeasuredPressure) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MinMeasuredPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinMeasuredPressure(*nt)
	return br, err
}

func (a MinMeasuredPressure) Readable() bool   { return true }
func (a MinMeasuredPressure) Writable() bool   { return false }
func (a MinMeasuredPressure) Reportable() bool { return false }
func (a MinMeasuredPressure) SceneIndex() int  { return -1 }

func (a MinMeasuredPressure) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MaxMeasuredPressure zcl.Zs16

func (a MaxMeasuredPressure) ID() zcl.AttrID               { return 2 }
func (a MaxMeasuredPressure) Cluster() zcl.ClusterID       { return PressureMeasurementCluster }
func (a *MaxMeasuredPressure) Value() *MaxMeasuredPressure { return a }
func (a MaxMeasuredPressure) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MaxMeasuredPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxMeasuredPressure(*nt)
	return br, err
}

func (a MaxMeasuredPressure) Readable() bool   { return true }
func (a MaxMeasuredPressure) Writable() bool   { return false }
func (a MaxMeasuredPressure) Reportable() bool { return false }
func (a MaxMeasuredPressure) SceneIndex() int  { return -1 }

func (a MaxMeasuredPressure) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type PressureTolerance zcl.Zu16

func (a PressureTolerance) ID() zcl.AttrID             { return 3 }
func (a PressureTolerance) Cluster() zcl.ClusterID     { return PressureMeasurementCluster }
func (a *PressureTolerance) Value() *PressureTolerance { return a }
func (a PressureTolerance) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PressureTolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PressureTolerance(*nt)
	return br, err
}

func (a PressureTolerance) Readable() bool   { return true }
func (a PressureTolerance) Writable() bool   { return false }
func (a PressureTolerance) Reportable() bool { return true }
func (a PressureTolerance) SceneIndex() int  { return -1 }

func (a PressureTolerance) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type ScaledPressure zcl.Zs16

func (a ScaledPressure) ID() zcl.AttrID          { return 16 }
func (a ScaledPressure) Cluster() zcl.ClusterID  { return PressureMeasurementCluster }
func (a *ScaledPressure) Value() *ScaledPressure { return a }
func (a ScaledPressure) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ScaledPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ScaledPressure(*nt)
	return br, err
}

func (a ScaledPressure) Readable() bool   { return true }
func (a ScaledPressure) Writable() bool   { return false }
func (a ScaledPressure) Reportable() bool { return false }
func (a ScaledPressure) SceneIndex() int  { return -1 }

func (a ScaledPressure) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MinScaledPressure zcl.Zs16

func (a MinScaledPressure) ID() zcl.AttrID             { return 17 }
func (a MinScaledPressure) Cluster() zcl.ClusterID     { return PressureMeasurementCluster }
func (a *MinScaledPressure) Value() *MinScaledPressure { return a }
func (a MinScaledPressure) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MinScaledPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinScaledPressure(*nt)
	return br, err
}

func (a MinScaledPressure) Readable() bool   { return true }
func (a MinScaledPressure) Writable() bool   { return false }
func (a MinScaledPressure) Reportable() bool { return false }
func (a MinScaledPressure) SceneIndex() int  { return -1 }

func (a MinScaledPressure) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MaxScaledPressure zcl.Zs16

func (a MaxScaledPressure) ID() zcl.AttrID             { return 18 }
func (a MaxScaledPressure) Cluster() zcl.ClusterID     { return PressureMeasurementCluster }
func (a *MaxScaledPressure) Value() *MaxScaledPressure { return a }
func (a MaxScaledPressure) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MaxScaledPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxScaledPressure(*nt)
	return br, err
}

func (a MaxScaledPressure) Readable() bool   { return true }
func (a MaxScaledPressure) Writable() bool   { return false }
func (a MaxScaledPressure) Reportable() bool { return false }
func (a MaxScaledPressure) SceneIndex() int  { return -1 }

func (a MaxScaledPressure) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type ScaledPressureTolerance zcl.Zu16

func (a ScaledPressureTolerance) ID() zcl.AttrID                   { return 19 }
func (a ScaledPressureTolerance) Cluster() zcl.ClusterID           { return PressureMeasurementCluster }
func (a *ScaledPressureTolerance) Value() *ScaledPressureTolerance { return a }
func (a ScaledPressureTolerance) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ScaledPressureTolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ScaledPressureTolerance(*nt)
	return br, err
}

func (a ScaledPressureTolerance) Readable() bool   { return true }
func (a ScaledPressureTolerance) Writable() bool   { return false }
func (a ScaledPressureTolerance) Reportable() bool { return false }
func (a ScaledPressureTolerance) SceneIndex() int  { return -1 }

func (a ScaledPressureTolerance) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type Scale zcl.Zs8

func (a Scale) ID() zcl.AttrID         { return 20 }
func (a Scale) Cluster() zcl.ClusterID { return PressureMeasurementCluster }
func (a *Scale) Value() *Scale         { return a }
func (a Scale) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *Scale) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = Scale(*nt)
	return br, err
}

func (a Scale) Readable() bool   { return true }
func (a Scale) Writable() bool   { return false }
func (a Scale) Reportable() bool { return false }
func (a Scale) SceneIndex() int  { return -1 }

func (a Scale) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}
