package hvac

import "hemtjan.st/zcl"

// ThermostatUiConfiguration
const ThermostatUiConfigurationID zcl.ClusterID = 516

var ThermostatUiConfigurationCluster = zcl.Cluster{
	Name:      "Thermostat UI configuration",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		TemperatureDisplayModeAttr:        func() zcl.Attr { return new(TemperatureDisplayMode) },
		KeypadLockoutAttr:                 func() zcl.Attr { return new(KeypadLockout) },
		ScheduleProgrammingVisibilityAttr: func() zcl.Attr { return new(ScheduleProgrammingVisibility) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
