package measurement

import "hemtjan.st/zcl"

// Pm25Measurement
const Pm25MeasurementID zcl.ClusterID = 1066

var Pm25MeasurementCluster = zcl.Cluster{
	Name:      "PM2.5 Measurement",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredPm25Attr:    func() zcl.Attr { return new(MeasuredPm25) },
		MinMeasuredPm25Attr: func() zcl.Attr { return new(MinMeasuredPm25) },
		MaxMeasuredPm25Attr: func() zcl.Attr { return new(MaxMeasuredPm25) },
		Pm25ToleranceAttr:   func() zcl.Attr { return new(Pm25Tolerance) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
