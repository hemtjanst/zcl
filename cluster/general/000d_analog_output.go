package general

import "hemtjan.st/zcl"

// AnalogOutput
const AnalogOutputID zcl.ClusterID = 13

var AnalogOutputCluster = zcl.Cluster{
	Name:      "Analog Output",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		IODescriptionAttr:           func() zcl.Attr { return new(IODescription) },
		AnalogMaxPresentValueAttr:   func() zcl.Attr { return new(AnalogMaxPresentValue) },
		AnalogMinPresentValueAttr:   func() zcl.Attr { return new(AnalogMinPresentValue) },
		IOOutOfServiceAttr:          func() zcl.Attr { return new(IOOutOfService) },
		AnalogPresentValueAttr:      func() zcl.Attr { return new(AnalogPresentValue) },
		AnalogPriorityArrayAttr:     func() zcl.Attr { return new(AnalogPriorityArray) },
		IOReliabilityAttr:           func() zcl.Attr { return new(IOReliability) },
		AnalogRelinquishDefaultAttr: func() zcl.Attr { return new(AnalogRelinquishDefault) },
		AnalogResolutionAttr:        func() zcl.Attr { return new(AnalogResolution) },
		IOStatusFlagsAttr:           func() zcl.Attr { return new(IOStatusFlags) },
		IOUnitTypeAttr:              func() zcl.Attr { return new(IOUnitType) },
		IOApplicationTypeAttr:       func() zcl.Attr { return new(IOApplicationType) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
