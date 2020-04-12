package measurement

import "hemtjan.st/zcl"

// TemperatureMeasurement
const TemperatureMeasurementID zcl.ClusterID = 1026

var TemperatureMeasurementCluster = zcl.Cluster{
	Name:      "Temperature measurement",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredTemperatureAttr:    func() zcl.Attr { return new(MeasuredTemperature) },
		MinMeasuredTemperatureAttr: func() zcl.Attr { return new(MinMeasuredTemperature) },
		MaxMeasuredTemperatureAttr: func() zcl.Attr { return new(MaxMeasuredTemperature) },
		ToleranceAttr:              func() zcl.Attr { return new(Tolerance) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
