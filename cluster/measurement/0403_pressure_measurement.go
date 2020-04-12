package measurement

import "hemtjan.st/zcl"

// PressureMeasurement
const PressureMeasurementID zcl.ClusterID = 1027

var PressureMeasurementCluster = zcl.Cluster{
	Name:      "Pressure measurement",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredPressureAttr:    func() zcl.Attr { return new(MeasuredPressure) },
		MinMeasuredPressureAttr: func() zcl.Attr { return new(MinMeasuredPressure) },
		MaxMeasuredPressureAttr: func() zcl.Attr { return new(MaxMeasuredPressure) },
		ToleranceAttr:           func() zcl.Attr { return new(Tolerance) },
		ScaledPressureAttr:      func() zcl.Attr { return new(ScaledPressure) },
		ScaledMinPressureAttr:   func() zcl.Attr { return new(ScaledMinPressure) },
		ScaledMaxPressureAttr:   func() zcl.Attr { return new(ScaledMaxPressure) },
		ScaledToleranceAttr:     func() zcl.Attr { return new(ScaledTolerance) },
		ScaleAttr:               func() zcl.Attr { return new(Scale) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
