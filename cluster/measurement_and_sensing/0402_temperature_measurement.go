// The server cluster provides an interface to temperature measurement functionality, including configuration and provision of notifications of temperature measurements.
package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// TemperatureMeasurement
// The server cluster provides an interface to temperature measurement functionality, including configuration and provision of notifications of temperature measurements.

func NewTemperatureMeasurementServer(profile zcl.ProfileID) *TemperatureMeasurementServer {
	return &TemperatureMeasurementServer{p: profile}
}
func NewTemperatureMeasurementClient(profile zcl.ProfileID) *TemperatureMeasurementClient {
	return &TemperatureMeasurementClient{p: profile}
}

const TemperatureMeasurementCluster zcl.ClusterID = 1026

type TemperatureMeasurementServer struct {
	p zcl.ProfileID

	MeasuredTemperature    *MeasuredTemperature
	MinMeasuredTemperature *MinMeasuredTemperature
	MaxMeasuredTemperature *MaxMeasuredTemperature
	TemperatureTolerance   *TemperatureTolerance
}

type TemperatureMeasurementClient struct {
	p zcl.ProfileID
}

/*
var TemperatureMeasurementServer = map[zcl.CommandID]func() zcl.Command{
}

var TemperatureMeasurementClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type MeasuredTemperature zcl.Zs16

func (a MeasuredTemperature) ID() zcl.AttrID               { return 0 }
func (a MeasuredTemperature) Cluster() zcl.ClusterID       { return TemperatureMeasurementCluster }
func (a *MeasuredTemperature) Value() *MeasuredTemperature { return a }
func (a MeasuredTemperature) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredTemperature(*nt)
	return br, err
}

func (a MeasuredTemperature) Readable() bool   { return true }
func (a MeasuredTemperature) Writable() bool   { return false }
func (a MeasuredTemperature) Reportable() bool { return true }
func (a MeasuredTemperature) SceneIndex() int  { return -1 }

func (a MeasuredTemperature) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MinMeasuredTemperature zcl.Zs16

func (a MinMeasuredTemperature) ID() zcl.AttrID                  { return 1 }
func (a MinMeasuredTemperature) Cluster() zcl.ClusterID          { return TemperatureMeasurementCluster }
func (a *MinMeasuredTemperature) Value() *MinMeasuredTemperature { return a }
func (a MinMeasuredTemperature) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MinMeasuredTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinMeasuredTemperature(*nt)
	return br, err
}

func (a MinMeasuredTemperature) Readable() bool   { return true }
func (a MinMeasuredTemperature) Writable() bool   { return false }
func (a MinMeasuredTemperature) Reportable() bool { return false }
func (a MinMeasuredTemperature) SceneIndex() int  { return -1 }

func (a MinMeasuredTemperature) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type MaxMeasuredTemperature zcl.Zs16

func (a MaxMeasuredTemperature) ID() zcl.AttrID                  { return 2 }
func (a MaxMeasuredTemperature) Cluster() zcl.ClusterID          { return TemperatureMeasurementCluster }
func (a *MaxMeasuredTemperature) Value() *MaxMeasuredTemperature { return a }
func (a MaxMeasuredTemperature) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MaxMeasuredTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxMeasuredTemperature(*nt)
	return br, err
}

func (a MaxMeasuredTemperature) Readable() bool   { return true }
func (a MaxMeasuredTemperature) Writable() bool   { return false }
func (a MaxMeasuredTemperature) Reportable() bool { return false }
func (a MaxMeasuredTemperature) SceneIndex() int  { return -1 }

func (a MaxMeasuredTemperature) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

type TemperatureTolerance zcl.Zu16

func (a TemperatureTolerance) ID() zcl.AttrID                { return 3 }
func (a TemperatureTolerance) Cluster() zcl.ClusterID        { return TemperatureMeasurementCluster }
func (a *TemperatureTolerance) Value() *TemperatureTolerance { return a }
func (a TemperatureTolerance) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *TemperatureTolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TemperatureTolerance(*nt)
	return br, err
}

func (a TemperatureTolerance) Readable() bool   { return true }
func (a TemperatureTolerance) Writable() bool   { return false }
func (a TemperatureTolerance) Reportable() bool { return true }
func (a TemperatureTolerance) SceneIndex() int  { return -1 }

func (a TemperatureTolerance) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}
