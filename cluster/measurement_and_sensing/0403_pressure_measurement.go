// The server cluster provides an interface to air pressure measurement functionality, including configuration and provision of notifications of air pressure measurements.
package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// PressureMeasurement
const PressureMeasurementID zcl.ClusterID = 1027

var PressureMeasurementCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredPressureAttr:        func() zcl.Attr { return new(MeasuredPressure) },
		MinMeasuredPressureAttr:     func() zcl.Attr { return new(MinMeasuredPressure) },
		MaxMeasuredPressureAttr:     func() zcl.Attr { return new(MaxMeasuredPressure) },
		PressureToleranceAttr:       func() zcl.Attr { return new(PressureTolerance) },
		ScaledPressureAttr:          func() zcl.Attr { return new(ScaledPressure) },
		MinScaledPressureAttr:       func() zcl.Attr { return new(MinScaledPressure) },
		MaxScaledPressureAttr:       func() zcl.Attr { return new(MaxScaledPressure) },
		ScaledPressureToleranceAttr: func() zcl.Attr { return new(ScaledPressureTolerance) },
		ScaleAttr:                   func() zcl.Attr { return new(Scale) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// MeasuredPressure is an autogenerated attribute in the PressureMeasurement cluster
type MeasuredPressure zcl.Zs16

const MeasuredPressureAttr zcl.AttrID = 0

func (a MeasuredPressure) ID() zcl.AttrID            { return MeasuredPressureAttr }
func (a MeasuredPressure) Cluster() zcl.ClusterID    { return PressureMeasurementID }
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
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// MinMeasuredPressure is an autogenerated attribute in the PressureMeasurement cluster
type MinMeasuredPressure zcl.Zs16

const MinMeasuredPressureAttr zcl.AttrID = 1

func (a MinMeasuredPressure) ID() zcl.AttrID               { return MinMeasuredPressureAttr }
func (a MinMeasuredPressure) Cluster() zcl.ClusterID       { return PressureMeasurementID }
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
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// MaxMeasuredPressure is an autogenerated attribute in the PressureMeasurement cluster
type MaxMeasuredPressure zcl.Zs16

const MaxMeasuredPressureAttr zcl.AttrID = 2

func (a MaxMeasuredPressure) ID() zcl.AttrID               { return MaxMeasuredPressureAttr }
func (a MaxMeasuredPressure) Cluster() zcl.ClusterID       { return PressureMeasurementID }
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
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// PressureTolerance is an autogenerated attribute in the PressureMeasurement cluster
type PressureTolerance zcl.Zu16

const PressureToleranceAttr zcl.AttrID = 3

func (a PressureTolerance) ID() zcl.AttrID             { return PressureToleranceAttr }
func (a PressureTolerance) Cluster() zcl.ClusterID     { return PressureMeasurementID }
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
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ScaledPressure is an autogenerated attribute in the PressureMeasurement cluster
type ScaledPressure zcl.Zs16

const ScaledPressureAttr zcl.AttrID = 16

func (a ScaledPressure) ID() zcl.AttrID          { return ScaledPressureAttr }
func (a ScaledPressure) Cluster() zcl.ClusterID  { return PressureMeasurementID }
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
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// MinScaledPressure is an autogenerated attribute in the PressureMeasurement cluster
type MinScaledPressure zcl.Zs16

const MinScaledPressureAttr zcl.AttrID = 17

func (a MinScaledPressure) ID() zcl.AttrID             { return MinScaledPressureAttr }
func (a MinScaledPressure) Cluster() zcl.ClusterID     { return PressureMeasurementID }
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
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// MaxScaledPressure is an autogenerated attribute in the PressureMeasurement cluster
type MaxScaledPressure zcl.Zs16

const MaxScaledPressureAttr zcl.AttrID = 18

func (a MaxScaledPressure) ID() zcl.AttrID             { return MaxScaledPressureAttr }
func (a MaxScaledPressure) Cluster() zcl.ClusterID     { return PressureMeasurementID }
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
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// ScaledPressureTolerance is an autogenerated attribute in the PressureMeasurement cluster
type ScaledPressureTolerance zcl.Zu16

const ScaledPressureToleranceAttr zcl.AttrID = 19

func (a ScaledPressureTolerance) ID() zcl.AttrID                   { return ScaledPressureToleranceAttr }
func (a ScaledPressureTolerance) Cluster() zcl.ClusterID           { return PressureMeasurementID }
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
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Scale is an autogenerated attribute in the PressureMeasurement cluster
type Scale zcl.Zs8

const ScaleAttr zcl.AttrID = 20

func (a Scale) ID() zcl.AttrID         { return ScaleAttr }
func (a Scale) Cluster() zcl.ClusterID { return PressureMeasurementID }
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
	return zcl.Sprintf("%v", zcl.Zs8(a))
}
