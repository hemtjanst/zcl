package general

import "hemtjan.st/zcl"

// PowerConfiguration
const PowerConfigurationID zcl.ClusterID = 1

var PowerConfigurationCluster = zcl.Cluster{
	Name:      "Power Configuration",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MainsVoltageAttr:                  func() zcl.Attr { return new(MainsVoltage) },
		MainsFrequencyAttr:                func() zcl.Attr { return new(MainsFrequency) },
		MainsAlarmMaskAttr:                func() zcl.Attr { return new(MainsAlarmMask) },
		MainsVoltageMinThresholdAttr:      func() zcl.Attr { return new(MainsVoltageMinThreshold) },
		MainsVoltageMaxThresholdAttr:      func() zcl.Attr { return new(MainsVoltageMaxThreshold) },
		MainsVoltageDwellTripPointAttr:    func() zcl.Attr { return new(MainsVoltageDwellTripPoint) },
		BatteryVoltageAttr:                func() zcl.Attr { return new(BatteryVoltage) },
		BatteryRemainingAttr:              func() zcl.Attr { return new(BatteryRemaining) },
		BatteryManufacturerAttr:           func() zcl.Attr { return new(BatteryManufacturer) },
		BatterySizeAttr:                   func() zcl.Attr { return new(BatterySize) },
		BatteryCapacityAttr:               func() zcl.Attr { return new(BatteryCapacity) },
		BatteryQuantityAttr:               func() zcl.Attr { return new(BatteryQuantity) },
		BatteryRatedVoltageAttr:           func() zcl.Attr { return new(BatteryRatedVoltage) },
		BatteryAlarmMaskAttr:              func() zcl.Attr { return new(BatteryAlarmMask) },
		BatteryVoltageMinThresholdAttr:    func() zcl.Attr { return new(BatteryVoltageMinThreshold) },
		BatteryVoltageThreshold1Attr:      func() zcl.Attr { return new(BatteryVoltageThreshold1) },
		BatteryVoltageThreshold2Attr:      func() zcl.Attr { return new(BatteryVoltageThreshold2) },
		BatteryVoltageThreshold3Attr:      func() zcl.Attr { return new(BatteryVoltageThreshold3) },
		BatteryPercentageMinThresholdAttr: func() zcl.Attr { return new(BatteryPercentageMinThreshold) },
		BatteryPercentageThreshold1Attr:   func() zcl.Attr { return new(BatteryPercentageThreshold1) },
		BatteryPercentageThreshold2Attr:   func() zcl.Attr { return new(BatteryPercentageThreshold2) },
		BatteryPercentageThreshold3Attr:   func() zcl.Attr { return new(BatteryPercentageThreshold3) },
		BatteryAlarmStateAttr:             func() zcl.Attr { return new(BatteryAlarmState) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
