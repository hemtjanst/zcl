package hvac

import "hemtjan.st/zcl"

// FanControl
const FanControlID zcl.ClusterID = 514

var FanControlCluster = zcl.Cluster{
	Name:      "Fan Control",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		FanModeAttr:         func() zcl.Attr { return new(FanMode) },
		FanModeSequenceAttr: func() zcl.Attr { return new(FanModeSequence) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
