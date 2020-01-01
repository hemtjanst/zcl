package general

import "hemtjan.st/zcl"

// AnalogValue
const AnalogValueID zcl.ClusterID = 14

var AnalogValueCluster = zcl.Cluster{
	Name:      "Analog Value",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		IODescriptionAttr:           func() zcl.Attr { return new(IODescription) },
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
