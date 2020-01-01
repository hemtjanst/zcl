package general

import "hemtjan.st/zcl"

// BinaryInput
const BinaryInputID zcl.ClusterID = 15

var BinaryInputCluster = zcl.Cluster{
	Name:      "Binary Input",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		BinaryActiveTextAttr:   func() zcl.Attr { return new(BinaryActiveText) },
		IODescriptionAttr:      func() zcl.Attr { return new(IODescription) },
		BinaryInactiveTextAttr: func() zcl.Attr { return new(BinaryInactiveText) },
		IOOutOfServiceAttr:     func() zcl.Attr { return new(IOOutOfService) },
		BinaryPolarityAttr:     func() zcl.Attr { return new(BinaryPolarity) },
		BinaryPresentValueAttr: func() zcl.Attr { return new(BinaryPresentValue) },
		IOReliabilityAttr:      func() zcl.Attr { return new(IOReliability) },
		IOStatusFlagsAttr:      func() zcl.Attr { return new(IOStatusFlags) },
		IOApplicationTypeAttr:  func() zcl.Attr { return new(IOApplicationType) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
