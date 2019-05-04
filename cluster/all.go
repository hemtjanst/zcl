package cluster

import (
	"hemtjan.st/zcl"
	"hemtjan.st/zcl/cluster/closures"
	"hemtjan.st/zcl/cluster/commissioning"
	"hemtjan.st/zcl/cluster/general"
	"hemtjan.st/zcl/cluster/green_power"
	"hemtjan.st/zcl/cluster/hvac"
	"hemtjan.st/zcl/cluster/lighting"
	"hemtjan.st/zcl/cluster/measurement_and_sensing"
	"hemtjan.st/zcl/cluster/ota_upgrade"
	"hemtjan.st/zcl/cluster/other"
	"hemtjan.st/zcl/cluster/smart_energy"
)

var Clusters = map[zcl.ClusterID]zcl.Cluster{
	closures.DoorLockID:                                   closures.DoorLockCluster,
	closures.ShadeConfigurationID:                         closures.ShadeConfigurationCluster,
	closures.WindowCoveringID:                             closures.WindowCoveringCluster,
	commissioning.CommissioningID:                         commissioning.CommissioningCluster,
	commissioning.TouchlinkID:                             commissioning.TouchlinkCluster,
	general.AlarmsID:                                      general.AlarmsCluster,
	general.AnalogInputBasicID:                            general.AnalogInputBasicCluster,
	general.AnalogOutputBasicID:                           general.AnalogOutputBasicCluster,
	general.AnalogValueBasicID:                            general.AnalogValueBasicCluster,
	general.BasicID:                                       general.BasicCluster,
	general.BinaryInputBasicID:                            general.BinaryInputBasicCluster,
	general.BinaryOutputBasicID:                           general.BinaryOutputBasicCluster,
	general.BinaryValueBasicID:                            general.BinaryValueBasicCluster,
	general.DeviceTemperatureConfigurationID:              general.DeviceTemperatureConfigurationCluster,
	general.DiagnosticsID:                                 general.DiagnosticsCluster,
	general.GroupsID:                                      general.GroupsCluster,
	general.IdentifyID:                                    general.IdentifyCluster,
	general.LevelControlID:                                general.LevelControlCluster,
	general.LocationID:                                    general.LocationCluster,
	general.MeterIdentificationID:                         general.MeterIdentificationCluster,
	general.MultistateInputBasicID:                        general.MultistateInputBasicCluster,
	general.MultistateOutputBasicID:                       general.MultistateOutputBasicCluster,
	general.MultistateValueBasicID:                        general.MultistateValueBasicCluster,
	general.OnOffID:                                       general.OnOffCluster,
	general.OnOffSwitchConfigurationID:                    general.OnOffSwitchConfigurationCluster,
	general.PollControlID:                                 general.PollControlCluster,
	general.PowerConfigurationID:                          general.PowerConfigurationCluster,
	general.PowerProfileID:                                general.PowerProfileCluster,
	general.ScenesID:                                      general.ScenesCluster,
	general.TimeID:                                        general.TimeCluster,
	green_power.GreenPowerID:                              green_power.GreenPowerCluster,
	hvac.DeDebugID:                                        hvac.DeDebugCluster,
	hvac.DehumidificationControlID:                        hvac.DehumidificationControlCluster,
	hvac.FanControlID:                                     hvac.FanControlCluster,
	hvac.PumpConfigurationAndControlID:                    hvac.PumpConfigurationAndControlCluster,
	hvac.ThermostatID:                                     hvac.ThermostatCluster,
	hvac.ThermostatUserInterfaceConfigurationID:           hvac.ThermostatUserInterfaceConfigurationCluster,
	lighting.BallastConfigurationID:                       lighting.BallastConfigurationCluster,
	lighting.ColorControlID:                               lighting.ColorControlCluster,
	measurement_and_sensing.ElectricalMeasurementID:       measurement_and_sensing.ElectricalMeasurementCluster,
	measurement_and_sensing.FlowMeasurementID:             measurement_and_sensing.FlowMeasurementCluster,
	measurement_and_sensing.IasWdID:                       measurement_and_sensing.IasWdCluster,
	measurement_and_sensing.IasZoneID:                     measurement_and_sensing.IasZoneCluster,
	measurement_and_sensing.IlluminanceLevelSensingID:     measurement_and_sensing.IlluminanceLevelSensingCluster,
	measurement_and_sensing.IlluminanceMeasurementID:      measurement_and_sensing.IlluminanceMeasurementCluster,
	measurement_and_sensing.OccupancySensingID:            measurement_and_sensing.OccupancySensingCluster,
	measurement_and_sensing.PressureMeasurementID:         measurement_and_sensing.PressureMeasurementCluster,
	measurement_and_sensing.RelativeHumidityMeasurementID: measurement_and_sensing.RelativeHumidityMeasurementCluster,
	measurement_and_sensing.TemperatureMeasurementID:      measurement_and_sensing.TemperatureMeasurementCluster,
	ota_upgrade.OtauID:                                    ota_upgrade.OtauCluster,
	other.KeyEstablishmentID:                              other.KeyEstablishmentCluster,
	other.RgbColorID:                                      other.RgbColorCluster,
	smart_energy.DemandResponseAndLoadControlID:           smart_energy.DemandResponseAndLoadControlCluster,
	smart_energy.MessageID:                                smart_energy.MessageCluster,
	smart_energy.PrepaymentID:                             smart_energy.PrepaymentCluster,
	smart_energy.PriceID:                                  smart_energy.PriceCluster,
	smart_energy.SimpleMeteringID:                         smart_energy.SimpleMeteringCluster,
	smart_energy.SmartEnergyTunnelingComplexMeteringID:    smart_energy.SmartEnergyTunnelingComplexMeteringCluster,
}
