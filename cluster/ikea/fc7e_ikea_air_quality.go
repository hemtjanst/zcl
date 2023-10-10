package ikea

import "hemtjan.st/zcl"

// IkeaAirQuality
const IkeaAirQualityID zcl.ClusterID = 64638

var IkeaAirQualityCluster = zcl.Cluster{
	Name:      "IKEA Air Quality",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		VocIndexAttr: func() zcl.Attr { return new(VocIndex) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
