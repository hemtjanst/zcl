package general

import "hemtjan.st/zcl"

// MultistateOutput
const MultistateOutputID zcl.ClusterID = 19

var MultistateOutputCluster = zcl.Cluster{
	Name:      "Multistate Output",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MultistateTextAttr:              func() zcl.Attr { return new(MultistateText) },
		IODescriptionAttr:               func() zcl.Attr { return new(IODescription) },
		MultistateNumberOfStatesAttr:    func() zcl.Attr { return new(MultistateNumberOfStates) },
		IOOutOfServiceAttr:              func() zcl.Attr { return new(IOOutOfService) },
		MultistatePresentValueAttr:      func() zcl.Attr { return new(MultistatePresentValue) },
		MultistatePriorityArrayAttr:     func() zcl.Attr { return new(MultistatePriorityArray) },
		IOReliabilityAttr:               func() zcl.Attr { return new(IOReliability) },
		MultistateRelinquishDefaultAttr: func() zcl.Attr { return new(MultistateRelinquishDefault) },
		IOStatusFlagsAttr:               func() zcl.Attr { return new(IOStatusFlags) },
		IOApplicationTypeAttr:           func() zcl.Attr { return new(IOApplicationType) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
