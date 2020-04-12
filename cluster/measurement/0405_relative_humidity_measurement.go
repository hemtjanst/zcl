package measurement

import "hemtjan.st/zcl"

// RelativeHumidityMeasurement
const RelativeHumidityMeasurementID zcl.ClusterID = 1029

var RelativeHumidityMeasurementCluster = zcl.Cluster{
	Name:      "Relative Humidity measurement",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredRelativeHumidityAttr:    func() zcl.Attr { return new(MeasuredRelativeHumidity) },
		MinMeasuredRelativeHumidityAttr: func() zcl.Attr { return new(MinMeasuredRelativeHumidity) },
		MaxMeasuredRelativeHumidityAttr: func() zcl.Attr { return new(MaxMeasuredRelativeHumidity) },
		ToleranceAttr:                   func() zcl.Attr { return new(Tolerance) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
