package general

import "hemtjan.st/zcl"

// OnOffSwitchConfiguration
const OnOffSwitchConfigurationID zcl.ClusterID = 7

var OnOffSwitchConfigurationCluster = zcl.Cluster{
	Name:      "On/Off Switch Configuration",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		SwitchTypeAttr:    func() zcl.Attr { return new(SwitchType) },
		SwitchActionsAttr: func() zcl.Attr { return new(SwitchActions) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
