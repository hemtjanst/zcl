package cluster

import (
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/cluster/general"
	"hemtjan.st/zcl/cluster/lighting"
)

var Clusters = map[zcl.ClusterID]zcl.Cluster{
	general.AlarmsID:                         general.AlarmsCluster,
	general.BasicID:                          general.BasicCluster,
	general.DeviceTemperatureConfigurationID: general.DeviceTemperatureConfigurationCluster,
	general.GroupsID:                         general.GroupsCluster,
	general.IdentifyID:                       general.IdentifyCluster,
	general.LevelControlID:                   general.LevelControlCluster,
	general.LocationID:                       general.LocationCluster,
	general.OnOffID:                          general.OnOffCluster,
	general.OnOffSwitchConfigurationID:       general.OnOffSwitchConfigurationCluster,
	general.PowerConfigurationID:             general.PowerConfigurationCluster,
	general.ScenesID:                         general.ScenesCluster,
	general.TimeID:                           general.TimeCluster,
	lighting.BallastConfigurationID:          lighting.BallastConfigurationCluster,
	lighting.ColorControlID:                  lighting.ColorControlCluster,
}
