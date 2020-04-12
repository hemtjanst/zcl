package measurement

import "hemtjan.st/zcl"

// IlluminanceLevelSensing
const IlluminanceLevelSensingID zcl.ClusterID = 1025

var IlluminanceLevelSensingCluster = zcl.Cluster{
	Name:      "Illuminance level sensing",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		LevelStatusAttr:                func() zcl.Attr { return new(LevelStatus) },
		IlluminanceLightSensorTypeAttr: func() zcl.Attr { return new(IlluminanceLightSensorType) },
		IlluminanceTargetLevelAttr:     func() zcl.Attr { return new(IlluminanceTargetLevel) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
