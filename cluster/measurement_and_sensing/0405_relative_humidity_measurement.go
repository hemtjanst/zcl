package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// RelativeHumidityMeasurement

func NewRelativeHumidityMeasurementServer(profile zcl.ProfileID) *RelativeHumidityMeasurementServer {
	return &RelativeHumidityMeasurementServer{p: profile}
}
func NewRelativeHumidityMeasurementClient(profile zcl.ProfileID) *RelativeHumidityMeasurementClient {
	return &RelativeHumidityMeasurementClient{p: profile}
}

const RelativeHumidityMeasurementCluster zcl.ClusterID = 1029

type RelativeHumidityMeasurementServer struct {
	p zcl.ProfileID

	MeasuredRelativeHumidity    *MeasuredRelativeHumidity
	MinMeasuredRelativeHumidity *MinMeasuredRelativeHumidity
	MaxMeasuredRelativeHumidity *MaxMeasuredRelativeHumidity
	RelativeHumidityTolerance   *RelativeHumidityTolerance
}

type RelativeHumidityMeasurementClient struct {
	p zcl.ProfileID
}

/*
var RelativeHumidityMeasurementServer = map[zcl.CommandID]func() zcl.Command{
}

var RelativeHumidityMeasurementClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type MeasuredRelativeHumidity zcl.Zu16

func (a MeasuredRelativeHumidity) ID() zcl.AttrID                    { return 0 }
func (a MeasuredRelativeHumidity) Cluster() zcl.ClusterID            { return RelativeHumidityMeasurementCluster }
func (a *MeasuredRelativeHumidity) Value() *MeasuredRelativeHumidity { return a }
func (a MeasuredRelativeHumidity) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MeasuredRelativeHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredRelativeHumidity(*nt)
	return br, err
}

func (a MeasuredRelativeHumidity) Readable() bool   { return true }
func (a MeasuredRelativeHumidity) Writable() bool   { return false }
func (a MeasuredRelativeHumidity) Reportable() bool { return true }
func (a MeasuredRelativeHumidity) SceneIndex() int  { return -1 }

func (a MeasuredRelativeHumidity) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type MinMeasuredRelativeHumidity zcl.Zu16

func (a MinMeasuredRelativeHumidity) ID() zcl.AttrID { return 1 }
func (a MinMeasuredRelativeHumidity) Cluster() zcl.ClusterID {
	return RelativeHumidityMeasurementCluster
}
func (a *MinMeasuredRelativeHumidity) Value() *MinMeasuredRelativeHumidity { return a }
func (a MinMeasuredRelativeHumidity) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MinMeasuredRelativeHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinMeasuredRelativeHumidity(*nt)
	return br, err
}

func (a MinMeasuredRelativeHumidity) Readable() bool   { return true }
func (a MinMeasuredRelativeHumidity) Writable() bool   { return false }
func (a MinMeasuredRelativeHumidity) Reportable() bool { return false }
func (a MinMeasuredRelativeHumidity) SceneIndex() int  { return -1 }

func (a MinMeasuredRelativeHumidity) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type MaxMeasuredRelativeHumidity zcl.Zu16

func (a MaxMeasuredRelativeHumidity) ID() zcl.AttrID { return 2 }
func (a MaxMeasuredRelativeHumidity) Cluster() zcl.ClusterID {
	return RelativeHumidityMeasurementCluster
}
func (a *MaxMeasuredRelativeHumidity) Value() *MaxMeasuredRelativeHumidity { return a }
func (a MaxMeasuredRelativeHumidity) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MaxMeasuredRelativeHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxMeasuredRelativeHumidity(*nt)
	return br, err
}

func (a MaxMeasuredRelativeHumidity) Readable() bool   { return true }
func (a MaxMeasuredRelativeHumidity) Writable() bool   { return false }
func (a MaxMeasuredRelativeHumidity) Reportable() bool { return false }
func (a MaxMeasuredRelativeHumidity) SceneIndex() int  { return -1 }

func (a MaxMeasuredRelativeHumidity) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type RelativeHumidityTolerance zcl.Zu16

func (a RelativeHumidityTolerance) ID() zcl.AttrID                     { return 3 }
func (a RelativeHumidityTolerance) Cluster() zcl.ClusterID             { return RelativeHumidityMeasurementCluster }
func (a *RelativeHumidityTolerance) Value() *RelativeHumidityTolerance { return a }
func (a RelativeHumidityTolerance) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RelativeHumidityTolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RelativeHumidityTolerance(*nt)
	return br, err
}

func (a RelativeHumidityTolerance) Readable() bool   { return true }
func (a RelativeHumidityTolerance) Writable() bool   { return false }
func (a RelativeHumidityTolerance) Reportable() bool { return true }
func (a RelativeHumidityTolerance) SceneIndex() int  { return -1 }

func (a RelativeHumidityTolerance) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}
