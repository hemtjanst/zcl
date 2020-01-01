package general

import "hemtjan.st/zcl"

// BinaryOutput
const BinaryOutputID zcl.ClusterID = 16

var BinaryOutputCluster = zcl.Cluster{
	Name:      "Binary Output",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		BinaryActiveTextAttr:        func() zcl.Attr { return new(BinaryActiveText) },
		IODescriptionAttr:           func() zcl.Attr { return new(IODescription) },
		BinaryInactiveTextAttr:      func() zcl.Attr { return new(BinaryInactiveText) },
		BinaryMinOffTimeAttr:        func() zcl.Attr { return new(BinaryMinOffTime) },
		BinaryMaxOffTimeAttr:        func() zcl.Attr { return new(BinaryMaxOffTime) },
		IOOutOfServiceAttr:          func() zcl.Attr { return new(IOOutOfService) },
		BinaryPolarityAttr:          func() zcl.Attr { return new(BinaryPolarity) },
		BinaryPresentValueAttr:      func() zcl.Attr { return new(BinaryPresentValue) },
		BinaryPriorityArrayAttr:     func() zcl.Attr { return new(BinaryPriorityArray) },
		IOReliabilityAttr:           func() zcl.Attr { return new(IOReliability) },
		BinaryRelinquishDefaultAttr: func() zcl.Attr { return new(BinaryRelinquishDefault) },
		IOStatusFlagsAttr:           func() zcl.Attr { return new(IOStatusFlags) },
		IOApplicationTypeAttr:       func() zcl.Attr { return new(IOApplicationType) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
