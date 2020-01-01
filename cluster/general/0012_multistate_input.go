package general

import "hemtjan.st/zcl"

// MultistateInput
const MultistateInputID zcl.ClusterID = 18

var MultistateInputCluster = zcl.Cluster{
	Name:      "Multistate Input",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MultistateTextAttr:           func() zcl.Attr { return new(MultistateText) },
		IODescriptionAttr:            func() zcl.Attr { return new(IODescription) },
		MultistateNumberOfStatesAttr: func() zcl.Attr { return new(MultistateNumberOfStates) },
		IOOutOfServiceAttr:           func() zcl.Attr { return new(IOOutOfService) },
		MultistatePresentValueAttr:   func() zcl.Attr { return new(MultistatePresentValue) },
		IOReliabilityAttr:            func() zcl.Attr { return new(IOReliability) },
		IOStatusFlagsAttr:            func() zcl.Attr { return new(IOStatusFlags) },
		IOApplicationTypeAttr:        func() zcl.Attr { return new(IOApplicationType) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
