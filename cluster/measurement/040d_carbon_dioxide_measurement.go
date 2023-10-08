package measurement

import "hemtjan.st/zcl"

// CarbonDioxideMeasurement
const CarbonDioxideMeasurementID zcl.ClusterID = 1037

var CarbonDioxideMeasurementCluster = zcl.Cluster{
	Name:      "Carbon Dioxide Measurement",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasuredConcentrationAttr:    func() zcl.Attr { return new(MeasuredConcentration) },
		MinMeasuredConcentrationAttr: func() zcl.Attr { return new(MinMeasuredConcentration) },
		MaxMeasuredConcentrationAttr: func() zcl.Attr { return new(MaxMeasuredConcentration) },
		ConcentrationToleranceAttr:   func() zcl.Attr { return new(ConcentrationTolerance) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
