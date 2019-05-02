package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// RelativeHumidityMeasurement
const RelativeHumidityMeasurementID zcl.ClusterID = 1029

var RelativeHumidityMeasurementCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredRelativeHumidityAttr:    func() zcl.Attr { return new(MeasuredRelativeHumidity) },
		MinMeasuredRelativeHumidityAttr: func() zcl.Attr { return new(MinMeasuredRelativeHumidity) },
		MaxMeasuredRelativeHumidityAttr: func() zcl.Attr { return new(MaxMeasuredRelativeHumidity) },
		RelativeHumidityToleranceAttr:   func() zcl.Attr { return new(RelativeHumidityTolerance) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const MeasuredRelativeHumidityAttr zcl.AttrID = 0

type MeasuredRelativeHumidity zcl.Zu16

func (a MeasuredRelativeHumidity) ID() zcl.AttrID                    { return MeasuredRelativeHumidityAttr }
func (a MeasuredRelativeHumidity) Cluster() zcl.ClusterID            { return RelativeHumidityMeasurementID }
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

const MinMeasuredRelativeHumidityAttr zcl.AttrID = 1

type MinMeasuredRelativeHumidity zcl.Zu16

func (a MinMeasuredRelativeHumidity) ID() zcl.AttrID                       { return MinMeasuredRelativeHumidityAttr }
func (a MinMeasuredRelativeHumidity) Cluster() zcl.ClusterID               { return RelativeHumidityMeasurementID }
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

const MaxMeasuredRelativeHumidityAttr zcl.AttrID = 2

type MaxMeasuredRelativeHumidity zcl.Zu16

func (a MaxMeasuredRelativeHumidity) ID() zcl.AttrID                       { return MaxMeasuredRelativeHumidityAttr }
func (a MaxMeasuredRelativeHumidity) Cluster() zcl.ClusterID               { return RelativeHumidityMeasurementID }
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

const RelativeHumidityToleranceAttr zcl.AttrID = 3

type RelativeHumidityTolerance zcl.Zu16

func (a RelativeHumidityTolerance) ID() zcl.AttrID                     { return RelativeHumidityToleranceAttr }
func (a RelativeHumidityTolerance) Cluster() zcl.ClusterID             { return RelativeHumidityMeasurementID }
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