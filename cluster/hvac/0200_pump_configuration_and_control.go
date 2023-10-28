package hvac

import "hemtjan.st/zcl"

// PumpConfigurationAndControl
const PumpConfigurationAndControlID zcl.ClusterID = 512

var PumpConfigurationAndControlCluster = zcl.Cluster{
	Name:      "Pump Configuration and Control",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MaxPressureAttr:            func() zcl.Attr { return new(MaxPressure) },
		MaxSpeedAttr:               func() zcl.Attr { return new(MaxSpeed) },
		MaxFlowAttr:                func() zcl.Attr { return new(MaxFlow) },
		MinConstPressureAttr:       func() zcl.Attr { return new(MinConstPressure) },
		MaxConstPressureAttr:       func() zcl.Attr { return new(MaxConstPressure) },
		MinCompPressureAttr:        func() zcl.Attr { return new(MinCompPressure) },
		MaxCompPressureAttr:        func() zcl.Attr { return new(MaxCompPressure) },
		MinConstSpeedAttr:          func() zcl.Attr { return new(MinConstSpeed) },
		MaxConstSpeedAttr:          func() zcl.Attr { return new(MaxConstSpeed) },
		MinConstFlowAttr:           func() zcl.Attr { return new(MinConstFlow) },
		MaxConstFlowAttr:           func() zcl.Attr { return new(MaxConstFlow) },
		MinConstTempAttr:           func() zcl.Attr { return new(MinConstTemp) },
		MaxConstTempAttr:           func() zcl.Attr { return new(MaxConstTemp) },
		PumpStatusAttr:             func() zcl.Attr { return new(PumpStatus) },
		EffectiveOperationModeAttr: func() zcl.Attr { return new(EffectiveOperationMode) },
		EffectiveControlModeAttr:   func() zcl.Attr { return new(EffectiveControlMode) },
		CapacityAttr:               func() zcl.Attr { return new(Capacity) },
		SpeedAttr:                  func() zcl.Attr { return new(Speed) },
		LifetimeRunningHoursAttr:   func() zcl.Attr { return new(LifetimeRunningHours) },
		PowerAttr:                  func() zcl.Attr { return new(Power) },
		LifetimeEnergyConsumedAttr: func() zcl.Attr { return new(LifetimeEnergyConsumed) },
		OperationModeAttr:          func() zcl.Attr { return new(OperationMode) },
		ControlModeAttr:            func() zcl.Attr { return new(ControlMode) },
		PumpAlarmMaskAttr:          func() zcl.Attr { return new(PumpAlarmMask) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
