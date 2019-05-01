// The server cluster provides an interface to temperature measurement functionality, including configuration and provision of notifications of temperature measurements.
package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// TemperatureMeasurement
const TemperatureMeasurementID zcl.ClusterID = 1026

var TemperatureMeasurementCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredTemperatureAttr:    func() zcl.Attr { return new(MeasuredTemperature) },
		MinMeasuredTemperatureAttr: func() zcl.Attr { return new(MinMeasuredTemperature) },
		MaxMeasuredTemperatureAttr: func() zcl.Attr { return new(MaxMeasuredTemperature) },
		TemperatureToleranceAttr:   func() zcl.Attr { return new(TemperatureTolerance) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const MeasuredTemperatureAttr zcl.AttrID = 0

type MeasuredTemperature zcl.Zs16

func (a MeasuredTemperature) ID() zcl.AttrID               { return MeasuredTemperatureAttr }
func (a MeasuredTemperature) Cluster() zcl.ClusterID       { return TemperatureMeasurementID }
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

const MinMeasuredTemperatureAttr zcl.AttrID = 1

type MinMeasuredTemperature zcl.Zs16

func (a MinMeasuredTemperature) ID() zcl.AttrID                  { return MinMeasuredTemperatureAttr }
func (a MinMeasuredTemperature) Cluster() zcl.ClusterID          { return TemperatureMeasurementID }
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

const MaxMeasuredTemperatureAttr zcl.AttrID = 2

type MaxMeasuredTemperature zcl.Zs16

func (a MaxMeasuredTemperature) ID() zcl.AttrID                  { return MaxMeasuredTemperatureAttr }
func (a MaxMeasuredTemperature) Cluster() zcl.ClusterID          { return TemperatureMeasurementID }
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

const TemperatureToleranceAttr zcl.AttrID = 3

type TemperatureTolerance zcl.Zu16

func (a TemperatureTolerance) ID() zcl.AttrID                { return TemperatureToleranceAttr }
func (a TemperatureTolerance) Cluster() zcl.ClusterID        { return TemperatureMeasurementID }
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
