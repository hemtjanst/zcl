package measurement

import "hemtjan.st/zcl"

// OccupancySensing
const OccupancySensingID zcl.ClusterID = 1030

var OccupancySensingCluster = zcl.Cluster{
	Name:      "Occupancy sensing",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		OccupancyAttr:                               func() zcl.Attr { return new(Occupancy) },
		OccupancyTypeAttr:                           func() zcl.Attr { return new(OccupancyType) },
		PirOccupiedToUnoccupiedDelayAttr:            func() zcl.Attr { return new(PirOccupiedToUnoccupiedDelay) },
		PirUnoccupiedToOccupiedDelayAttr:            func() zcl.Attr { return new(PirUnoccupiedToOccupiedDelay) },
		PirUnoccupiedToOccupiedThresholdAttr:        func() zcl.Attr { return new(PirUnoccupiedToOccupiedThreshold) },
		UltrasonicOccupiedToUnoccupiedDelayAttr:     func() zcl.Attr { return new(UltrasonicOccupiedToUnoccupiedDelay) },
		UltrasonicUnoccupiedToOccupiedDelayAttr:     func() zcl.Attr { return new(UltrasonicUnoccupiedToOccupiedDelay) },
		UltrasonicUnoccupiedToOccupiedThresholdAttr: func() zcl.Attr { return new(UltrasonicUnoccupiedToOccupiedThreshold) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
