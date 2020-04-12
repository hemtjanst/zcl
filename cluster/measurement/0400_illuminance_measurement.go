package measurement

import "hemtjan.st/zcl"

// IlluminanceMeasurement
const IlluminanceMeasurementID zcl.ClusterID = 1024

var IlluminanceMeasurementCluster = zcl.Cluster{
	Name:      "Illuminance measurement",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredIlluminanceAttr:    func() zcl.Attr { return new(MeasuredIlluminance) },
		MinMeasuredIlluminanceAttr: func() zcl.Attr { return new(MinMeasuredIlluminance) },
		MaxMeasuredIlluminanceAttr: func() zcl.Attr { return new(MaxMeasuredIlluminance) },
		ToleranceAttr:              func() zcl.Attr { return new(Tolerance) },
		LightSensorTypeAttr:        func() zcl.Attr { return new(LightSensorType) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
