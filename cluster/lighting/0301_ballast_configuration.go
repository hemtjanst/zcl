package lighting

import "hemtjan.st/zcl"

// BallastConfiguration
const BallastConfigurationID zcl.ClusterID = 769

var BallastConfigurationCluster = zcl.Cluster{
	Name:      "Ballast Configuration",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		PhysicalMinLevelAttr:        func() zcl.Attr { return new(PhysicalMinLevel) },
		PhysicalMaxLevelAttr:        func() zcl.Attr { return new(PhysicalMaxLevel) },
		BallastStatusAttr:           func() zcl.Attr { return new(BallastStatus) },
		MinLevelAttr:                func() zcl.Attr { return new(MinLevel) },
		MaxLevelAttr:                func() zcl.Attr { return new(MaxLevel) },
		PowerOnLevelAttr:            func() zcl.Attr { return new(PowerOnLevel) },
		PowerOnTimeAttr:             func() zcl.Attr { return new(PowerOnTime) },
		IntrinsicBallastFactorAttr:  func() zcl.Attr { return new(IntrinsicBallastFactor) },
		BallastFactorAdjustmentAttr: func() zcl.Attr { return new(BallastFactorAdjustment) },
		LampQuantityAttr:            func() zcl.Attr { return new(LampQuantity) },
		LampTypeAttr:                func() zcl.Attr { return new(LampType) },
		LampManufacturerAttr:        func() zcl.Attr { return new(LampManufacturer) },
		LampRatedHoursAttr:          func() zcl.Attr { return new(LampRatedHours) },
		LampBurnHoursAttr:           func() zcl.Attr { return new(LampBurnHours) },
		LampAlarmModeAttr:           func() zcl.Attr { return new(LampAlarmMode) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
