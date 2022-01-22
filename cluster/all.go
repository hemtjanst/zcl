package cluster

import (
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/cluster/closures"
	"hemtjan.st/zcl/cluster/general"
	"hemtjan.st/zcl/cluster/ias"
	"hemtjan.st/zcl/cluster/lighting"
	"hemtjan.st/zcl/cluster/measurement"
	"hemtjan.st/zcl/cluster/otau"
)

var Clusters = map[zcl.ClusterID]zcl.Cluster{
	closures.WindowCoveringID:                 closures.WindowCoveringCluster,
	general.AlarmsID:                          general.AlarmsCluster,
	general.AnalogInputID:                     general.AnalogInputCluster,
	general.AnalogOutputID:                    general.AnalogOutputCluster,
	general.AnalogValueID:                     general.AnalogValueCluster,
	general.BasicID:                           general.BasicCluster,
	general.BinaryInputID:                     general.BinaryInputCluster,
	general.BinaryOutputID:                    general.BinaryOutputCluster,
	general.BinaryValueID:                     general.BinaryValueCluster,
	general.DeviceTemperatureConfigurationID:  general.DeviceTemperatureConfigurationCluster,
	general.DiagnosticsID:                     general.DiagnosticsCluster,
	general.GroupsID:                          general.GroupsCluster,
	general.IdentifyID:                        general.IdentifyCluster,
	general.LevelControlID:                    general.LevelControlCluster,
	general.LocationID:                        general.LocationCluster,
	general.MeterIdentificationID:             general.MeterIdentificationCluster,
	general.MultistateInputID:                 general.MultistateInputCluster,
	general.MultistateOutputID:                general.MultistateOutputCluster,
	general.MultistateValueID:                 general.MultistateValueCluster,
	general.OnOffID:                           general.OnOffCluster,
	general.OnOffSwitchConfigurationID:        general.OnOffSwitchConfigurationCluster,
	general.PollControlID:                     general.PollControlCluster,
	general.PowerConfigurationID:              general.PowerConfigurationCluster,
	general.ScenesID:                          general.ScenesCluster,
	general.TimeID:                            general.TimeCluster,
	ias.IasZoneID:                             ias.IasZoneCluster,
	lighting.BallastConfigurationID:           lighting.BallastConfigurationCluster,
	lighting.ColorControlID:                   lighting.ColorControlCluster,
	measurement.FlowMeasurementID:             measurement.FlowMeasurementCluster,
	measurement.IlluminanceLevelSensingID:     measurement.IlluminanceLevelSensingCluster,
	measurement.IlluminanceMeasurementID:      measurement.IlluminanceMeasurementCluster,
	measurement.OccupancySensingID:            measurement.OccupancySensingCluster,
	measurement.PressureMeasurementID:         measurement.PressureMeasurementCluster,
	measurement.RelativeHumidityMeasurementID: measurement.RelativeHumidityMeasurementCluster,
	measurement.TemperatureMeasurementID:      measurement.TemperatureMeasurementCluster,
	otau.OtauID:                               otau.OtauCluster,
}
