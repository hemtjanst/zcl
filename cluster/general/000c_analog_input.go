package general

import "hemtjan.st/zcl"

// AnalogInput
const AnalogInputID zcl.ClusterID = 12

var AnalogInputCluster = zcl.Cluster{
	Name:      "Analog Input",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		IODescriptionAttr:         func() zcl.Attr { return new(IODescription) },
		AnalogMaxPresentValueAttr: func() zcl.Attr { return new(AnalogMaxPresentValue) },
		AnalogMinPresentValueAttr: func() zcl.Attr { return new(AnalogMinPresentValue) },
		IOOutOfServiceAttr:        func() zcl.Attr { return new(IOOutOfService) },
		AnalogPresentValueAttr:    func() zcl.Attr { return new(AnalogPresentValue) },
		IOReliabilityAttr:         func() zcl.Attr { return new(IOReliability) },
		AnalogResolutionAttr:      func() zcl.Attr { return new(AnalogResolution) },
		IOStatusFlagsAttr:         func() zcl.Attr { return new(IOStatusFlags) },
		IOUnitTypeAttr:            func() zcl.Attr { return new(IOUnitType) },
		IOApplicationTypeAttr:     func() zcl.Attr { return new(IOApplicationType) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
