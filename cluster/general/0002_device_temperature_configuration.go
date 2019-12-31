package general

import "hemtjan.st/zcl"

// DeviceTemperatureConfiguration
const DeviceTemperatureConfigurationID zcl.ClusterID = 2

var DeviceTemperatureConfigurationCluster = zcl.Cluster{
	Name:      "Device Temperature Configuration",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		CurrentTemperatureAttr:     func() zcl.Attr { return new(CurrentTemperature) },
		MinTempExperiencedAttr:     func() zcl.Attr { return new(MinTempExperienced) },
		MaxTempExperiencedAttr:     func() zcl.Attr { return new(MaxTempExperienced) },
		OverTempTotalDwellAttr:     func() zcl.Attr { return new(OverTempTotalDwell) },
		DeviceTempAlarmMaskAttr:    func() zcl.Attr { return new(DeviceTempAlarmMask) },
		LowTempThresholdAttr:       func() zcl.Attr { return new(LowTempThreshold) },
		HighTempThresholdAttr:      func() zcl.Attr { return new(HighTempThreshold) },
		LowTempDwellTripPointAttr:  func() zcl.Attr { return new(LowTempDwellTripPoint) },
		HighTempDwellTripPointAttr: func() zcl.Attr { return new(HighTempDwellTripPoint) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
