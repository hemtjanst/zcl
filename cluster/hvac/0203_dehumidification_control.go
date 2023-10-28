package hvac

import "hemtjan.st/zcl"

// DehumidificationControl
const DehumidificationControlID zcl.ClusterID = 515

var DehumidificationControlCluster = zcl.Cluster{
	Name:      "Dehumidification Control",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		RelativeHumidityAttr:           func() zcl.Attr { return new(RelativeHumidity) },
		DehumidificationCoolingAttr:    func() zcl.Attr { return new(DehumidificationCooling) },
		RhDehumidificationSetpointAttr: func() zcl.Attr { return new(RhDehumidificationSetpoint) },
		RelativeHumidityModeAttr:       func() zcl.Attr { return new(RelativeHumidityMode) },
		DehumidificationLockoutAttr:    func() zcl.Attr { return new(DehumidificationLockout) },
		DehumidificationHysteresisAttr: func() zcl.Attr { return new(DehumidificationHysteresis) },
		DehumidificationMaxCoolAttr:    func() zcl.Attr { return new(DehumidificationMaxCool) },
		RelativeHumidityDisplayAttr:    func() zcl.Attr { return new(RelativeHumidityDisplay) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
