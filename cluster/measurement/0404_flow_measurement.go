package measurement

import "hemtjan.st/zcl"

// FlowMeasurement
const FlowMeasurementID zcl.ClusterID = 1028

var FlowMeasurementCluster = zcl.Cluster{
	Name:      "Flow measurement",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredFlowAttr:    func() zcl.Attr { return new(MeasuredFlow) },
		MinMeasuredFlowAttr: func() zcl.Attr { return new(MinMeasuredFlow) },
		MaxMeasuredFlowAttr: func() zcl.Attr { return new(MaxMeasuredFlow) },
		ToleranceAttr:       func() zcl.Attr { return new(Tolerance) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
