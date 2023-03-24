package smartenergy

import "hemtjan.st/zcl"

// Metering
const MeteringID zcl.ClusterID = 1794

var MeteringCluster = zcl.Cluster{
	Name:      "Metering",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		CurrentSummationDeliveredAttr:               func() zcl.Attr { return new(CurrentSummationDelivered) },
		CurrentSummationReceivedAttr:                func() zcl.Attr { return new(CurrentSummationReceived) },
		CurrentMaxDemandDeliveredAttr:               func() zcl.Attr { return new(CurrentMaxDemandDelivered) },
		CurrentMaxDemandReceivedAttr:                func() zcl.Attr { return new(CurrentMaxDemandReceived) },
		DftSummationAttr:                            func() zcl.Attr { return new(DftSummation) },
		DailyFreezeTimeAttr:                         func() zcl.Attr { return new(DailyFreezeTime) },
		PowerFactorAttr:                             func() zcl.Attr { return new(PowerFactor) },
		ReadingSnapShotTimeAttr:                     func() zcl.Attr { return new(ReadingSnapShotTime) },
		CurrentMaxDemandDeliveredTimeAttr:           func() zcl.Attr { return new(CurrentMaxDemandDeliveredTime) },
		CurrentMaxDemandReceivedTimeAttr:            func() zcl.Attr { return new(CurrentMaxDemandReceivedTime) },
		DefaultUpdatePeriodAttr:                     func() zcl.Attr { return new(DefaultUpdatePeriod) },
		FastPollUpdatePeriodAttr:                    func() zcl.Attr { return new(FastPollUpdatePeriod) },
		CurrentBlockPeriodConsumptionDeliveredAttr:  func() zcl.Attr { return new(CurrentBlockPeriodConsumptionDelivered) },
		DailyConsumptionTargetAttr:                  func() zcl.Attr { return new(DailyConsumptionTarget) },
		CurrentBlockAttr:                            func() zcl.Attr { return new(CurrentBlock) },
		ProfileIntervalPeriodAttr:                   func() zcl.Attr { return new(ProfileIntervalPeriod) },
		IntervalReadReportingPeriodAttr:             func() zcl.Attr { return new(IntervalReadReportingPeriod) },
		PresetReadingTimeAttr:                       func() zcl.Attr { return new(PresetReadingTime) },
		VolumePerReportAttr:                         func() zcl.Attr { return new(VolumePerReport) },
		FlowRestrictionAttr:                         func() zcl.Attr { return new(FlowRestriction) },
		SupplyStatusAttr:                            func() zcl.Attr { return new(SupplyStatus) },
		CurrentInletEnergyCarrierSummationAttr:      func() zcl.Attr { return new(CurrentInletEnergyCarrierSummation) },
		CurrentOutletEnergyCarrierSummationAttr:     func() zcl.Attr { return new(CurrentOutletEnergyCarrierSummation) },
		InletTemperatureAttr:                        func() zcl.Attr { return new(InletTemperature) },
		OutletTemperatureAttr:                       func() zcl.Attr { return new(OutletTemperature) },
		ControlTemperatureAttr:                      func() zcl.Attr { return new(ControlTemperature) },
		CurrentInletEnergyCarrierDemandAttr:         func() zcl.Attr { return new(CurrentInletEnergyCarrierDemand) },
		CurrentOutletEnergyCarrierDemandAttr:        func() zcl.Attr { return new(CurrentOutletEnergyCarrierDemand) },
		PreviousBlockPeriodConsumptionDeliveredAttr: func() zcl.Attr { return new(PreviousBlockPeriodConsumptionDelivered) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}
