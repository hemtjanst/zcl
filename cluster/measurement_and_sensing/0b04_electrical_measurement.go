// Provides a mechanism for querying data about the electrical properties as measured by the device.
package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// ElectricalMeasurement
const ElectricalMeasurementID zcl.ClusterID = 2820

var ElectricalMeasurementCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		GetProfileInfoResponseCommand:        func() zcl.Command { return new(GetProfileInfoResponse) },
		GetMeasurementProfileResponseCommand: func() zcl.Command { return new(GetMeasurementProfileResponse) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		GetProfileInfoCommand:        func() zcl.Command { return new(GetProfileInfo) },
		GetMeasurementProfileCommand: func() zcl.Command { return new(GetMeasurementProfile) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MeasurementTypeAttr:                          func() zcl.Attr { return new(MeasurementType) },
		DcVoltageAttr:                                func() zcl.Attr { return new(DcVoltage) },
		DcVoltageMinAttr:                             func() zcl.Attr { return new(DcVoltageMin) },
		DcVoltageMaxAttr:                             func() zcl.Attr { return new(DcVoltageMax) },
		DcCurrentAttr:                                func() zcl.Attr { return new(DcCurrent) },
		DcCurrentMinAttr:                             func() zcl.Attr { return new(DcCurrentMin) },
		DcCurrentMaxAttr:                             func() zcl.Attr { return new(DcCurrentMax) },
		DcPowerAttr:                                  func() zcl.Attr { return new(DcPower) },
		DcPowerMinAttr:                               func() zcl.Attr { return new(DcPowerMin) },
		DcPowerMaxAttr:                               func() zcl.Attr { return new(DcPowerMax) },
		DcVoltageMultiplierAttr:                      func() zcl.Attr { return new(DcVoltageMultiplier) },
		DcVoltageDivisorAttr:                         func() zcl.Attr { return new(DcVoltageDivisor) },
		DcCurrentMultiplierAttr:                      func() zcl.Attr { return new(DcCurrentMultiplier) },
		DcCurrentDivisorAttr:                         func() zcl.Attr { return new(DcCurrentDivisor) },
		DcPowerMultiplierAttr:                        func() zcl.Attr { return new(DcPowerMultiplier) },
		DcPowerDivisorAttr:                           func() zcl.Attr { return new(DcPowerDivisor) },
		AcFrequencyAttr:                              func() zcl.Attr { return new(AcFrequency) },
		AcFrequencyMinAttr:                           func() zcl.Attr { return new(AcFrequencyMin) },
		AcFrequencyMaxAttr:                           func() zcl.Attr { return new(AcFrequencyMax) },
		NeutralCurrentAttr:                           func() zcl.Attr { return new(NeutralCurrent) },
		TotalActivePowerAttr:                         func() zcl.Attr { return new(TotalActivePower) },
		TotalReactivePowerAttr:                       func() zcl.Attr { return new(TotalReactivePower) },
		TotalApparentPowerAttr:                       func() zcl.Attr { return new(TotalApparentPower) },
		Measured1StHarmonicCurrentAttr:               func() zcl.Attr { return new(Measured1StHarmonicCurrent) },
		Measured3RdHarmonicCurrentAttr:               func() zcl.Attr { return new(Measured3RdHarmonicCurrent) },
		Measured5ThHarmonicCurrentAttr:               func() zcl.Attr { return new(Measured5ThHarmonicCurrent) },
		Measured7ThHarmonicCurrentAttr:               func() zcl.Attr { return new(Measured7ThHarmonicCurrent) },
		Measured9ThHarmonicCurrentAttr:               func() zcl.Attr { return new(Measured9ThHarmonicCurrent) },
		Measured11ThHarmonicCurrentAttr:              func() zcl.Attr { return new(Measured11ThHarmonicCurrent) },
		MeasuredPhase1StHarmonicCurrentAttr:          func() zcl.Attr { return new(MeasuredPhase1StHarmonicCurrent) },
		MeasuredPhase3RdHarmonicCurrentAttr:          func() zcl.Attr { return new(MeasuredPhase3RdHarmonicCurrent) },
		MeasuredPhase5ThHarmonicCurrentAttr:          func() zcl.Attr { return new(MeasuredPhase5ThHarmonicCurrent) },
		MeasuredPhase7ThHarmonicCurrentAttr:          func() zcl.Attr { return new(MeasuredPhase7ThHarmonicCurrent) },
		MeasuredPhase9ThHarmonicCurrentAttr:          func() zcl.Attr { return new(MeasuredPhase9ThHarmonicCurrent) },
		MeasuredPhase11ThHarmonicCurrentAttr:         func() zcl.Attr { return new(MeasuredPhase11ThHarmonicCurrent) },
		AcFrequencyMultiplierAttr:                    func() zcl.Attr { return new(AcFrequencyMultiplier) },
		AcFrequencyDivisorAttr:                       func() zcl.Attr { return new(AcFrequencyDivisor) },
		PowerMultiplierAttr:                          func() zcl.Attr { return new(PowerMultiplier) },
		PowerDivisorAttr:                             func() zcl.Attr { return new(PowerDivisor) },
		HarmonicCurrentMultiplierAttr:                func() zcl.Attr { return new(HarmonicCurrentMultiplier) },
		PhaseHarmonicCurrentMultiplierAttr:           func() zcl.Attr { return new(PhaseHarmonicCurrentMultiplier) },
		LineCurrentAttr:                              func() zcl.Attr { return new(LineCurrent) },
		ActiveCurrentAttr:                            func() zcl.Attr { return new(ActiveCurrent) },
		ReactiveCurrentAttr:                          func() zcl.Attr { return new(ReactiveCurrent) },
		RmsVoltageAttr:                               func() zcl.Attr { return new(RmsVoltage) },
		RmsVoltageMinAttr:                            func() zcl.Attr { return new(RmsVoltageMin) },
		RmsVoltageMaxAttr:                            func() zcl.Attr { return new(RmsVoltageMax) },
		RmsCurrentAttr:                               func() zcl.Attr { return new(RmsCurrent) },
		RmsCurrentMinAttr:                            func() zcl.Attr { return new(RmsCurrentMin) },
		RmsCurrentMaxAttr:                            func() zcl.Attr { return new(RmsCurrentMax) },
		ActivePowerAttr:                              func() zcl.Attr { return new(ActivePower) },
		ActivePowerMinAttr:                           func() zcl.Attr { return new(ActivePowerMin) },
		ActivePowerMaxAttr:                           func() zcl.Attr { return new(ActivePowerMax) },
		ReactivePowerAttr:                            func() zcl.Attr { return new(ReactivePower) },
		ApparentPowerAttr:                            func() zcl.Attr { return new(ApparentPower) },
		PowerFactorAttr:                              func() zcl.Attr { return new(PowerFactor) },
		AverageRmsVoltageMeasurementPeriodAttr:       func() zcl.Attr { return new(AverageRmsVoltageMeasurementPeriod) },
		AverageRmsOvervoltageCounterAttr:             func() zcl.Attr { return new(AverageRmsOvervoltageCounter) },
		AverageRmsUndervoltageCounterAttr:            func() zcl.Attr { return new(AverageRmsUndervoltageCounter) },
		RmsExtremeOvervoltagePeriodAttr:              func() zcl.Attr { return new(RmsExtremeOvervoltagePeriod) },
		RmsExtremeUndervoltagePeriodAttr:             func() zcl.Attr { return new(RmsExtremeUndervoltagePeriod) },
		RmsVoltageSagPeriodAttr:                      func() zcl.Attr { return new(RmsVoltageSagPeriod) },
		RmsVoltageSwellPeriodAttr:                    func() zcl.Attr { return new(RmsVoltageSwellPeriod) },
		AcVoltageMultiplierAttr:                      func() zcl.Attr { return new(AcVoltageMultiplier) },
		AcVoltageDivisorAttr:                         func() zcl.Attr { return new(AcVoltageDivisor) },
		AcCurrentMultiplierAttr:                      func() zcl.Attr { return new(AcCurrentMultiplier) },
		AcCurrentDivisorAttr:                         func() zcl.Attr { return new(AcCurrentDivisor) },
		AcPowerMultiplierAttr:                        func() zcl.Attr { return new(AcPowerMultiplier) },
		AcPowerDivisorAttr:                           func() zcl.Attr { return new(AcPowerDivisor) },
		DcOverloadAlarmsMaskAttr:                     func() zcl.Attr { return new(DcOverloadAlarmsMask) },
		DcVoltageOverloadAttr:                        func() zcl.Attr { return new(DcVoltageOverload) },
		DcCurrentOverloadAttr:                        func() zcl.Attr { return new(DcCurrentOverload) },
		AcOverloadAlarmsMaskAttr:                     func() zcl.Attr { return new(AcOverloadAlarmsMask) },
		AcVoltageOverloadAttr:                        func() zcl.Attr { return new(AcVoltageOverload) },
		AcCurrentOverloadAttr:                        func() zcl.Attr { return new(AcCurrentOverload) },
		AcActivePowerOverloadAttr:                    func() zcl.Attr { return new(AcActivePowerOverload) },
		AcReactivePowerOverloadAttr:                  func() zcl.Attr { return new(AcReactivePowerOverload) },
		AverageRmsOvervoltageAttr:                    func() zcl.Attr { return new(AverageRmsOvervoltage) },
		AverageRmsUndervoltageAttr:                   func() zcl.Attr { return new(AverageRmsUndervoltage) },
		RmsExtremeOvervoltageAttr:                    func() zcl.Attr { return new(RmsExtremeOvervoltage) },
		RmsExtremeUndervoltageAttr:                   func() zcl.Attr { return new(RmsExtremeUndervoltage) },
		RmsVoltageSagAttr:                            func() zcl.Attr { return new(RmsVoltageSag) },
		RmsVoltageSwellAttr:                          func() zcl.Attr { return new(RmsVoltageSwell) },
		LineCurrentPhaseBAttr:                        func() zcl.Attr { return new(LineCurrentPhaseB) },
		ActiveCurrentPhaseBAttr:                      func() zcl.Attr { return new(ActiveCurrentPhaseB) },
		ReactiveCurrentPhaseBAttr:                    func() zcl.Attr { return new(ReactiveCurrentPhaseB) },
		RmsVoltagePhaseBAttr:                         func() zcl.Attr { return new(RmsVoltagePhaseB) },
		RmsVoltageMinPhaseBAttr:                      func() zcl.Attr { return new(RmsVoltageMinPhaseB) },
		RmsVoltageMaxPhaseBAttr:                      func() zcl.Attr { return new(RmsVoltageMaxPhaseB) },
		RmsCurrentPhaseBAttr:                         func() zcl.Attr { return new(RmsCurrentPhaseB) },
		RmsCurrentMinPhaseBAttr:                      func() zcl.Attr { return new(RmsCurrentMinPhaseB) },
		RmsCurrentMaxPhaseBAttr:                      func() zcl.Attr { return new(RmsCurrentMaxPhaseB) },
		ActivePowerPhaseBAttr:                        func() zcl.Attr { return new(ActivePowerPhaseB) },
		ActivePowerMinPhaseBAttr:                     func() zcl.Attr { return new(ActivePowerMinPhaseB) },
		ActivePowerMaxPhaseBAttr:                     func() zcl.Attr { return new(ActivePowerMaxPhaseB) },
		ReactivePowerPhaseBAttr:                      func() zcl.Attr { return new(ReactivePowerPhaseB) },
		ApparentPowerPhaseBAttr:                      func() zcl.Attr { return new(ApparentPowerPhaseB) },
		PowerFactorPhaseBAttr:                        func() zcl.Attr { return new(PowerFactorPhaseB) },
		AverageRmsVoltageMeasurementPeriodPhaseBAttr: func() zcl.Attr { return new(AverageRmsVoltageMeasurementPeriodPhaseB) },
		AverageRmsOvervoltageCounterPhaseBAttr:       func() zcl.Attr { return new(AverageRmsOvervoltageCounterPhaseB) },
		AverageRmsUndervoltageCounterPhaseBAttr:      func() zcl.Attr { return new(AverageRmsUndervoltageCounterPhaseB) },
		RmsExtremeOvervoltagePeriodPhaseBAttr:        func() zcl.Attr { return new(RmsExtremeOvervoltagePeriodPhaseB) },
		RmsExtremeUndervoltagePeriodPhaseBAttr:       func() zcl.Attr { return new(RmsExtremeUndervoltagePeriodPhaseB) },
		RmsVoltageSagPeriodPhaseBAttr:                func() zcl.Attr { return new(RmsVoltageSagPeriodPhaseB) },
		RmsVoltageSwellPeriodPhaseBAttr:              func() zcl.Attr { return new(RmsVoltageSwellPeriodPhaseB) },
		LineCurrentPhaseCAttr:                        func() zcl.Attr { return new(LineCurrentPhaseC) },
		ActiveCurrentPhaseCAttr:                      func() zcl.Attr { return new(ActiveCurrentPhaseC) },
		ReactiveCurrentPhaseCAttr:                    func() zcl.Attr { return new(ReactiveCurrentPhaseC) },
		RmsVoltagePhaseCAttr:                         func() zcl.Attr { return new(RmsVoltagePhaseC) },
		RmsVoltageMinPhaseCAttr:                      func() zcl.Attr { return new(RmsVoltageMinPhaseC) },
		RmsVoltageMaxPhaseCAttr:                      func() zcl.Attr { return new(RmsVoltageMaxPhaseC) },
		RmsCurrentPhaseCAttr:                         func() zcl.Attr { return new(RmsCurrentPhaseC) },
		RmsCurrentMinPhaseCAttr:                      func() zcl.Attr { return new(RmsCurrentMinPhaseC) },
		RmsCurrentMaxPhaseCAttr:                      func() zcl.Attr { return new(RmsCurrentMaxPhaseC) },
		ActivePowerPhaseCAttr:                        func() zcl.Attr { return new(ActivePowerPhaseC) },
		ActivePowerMinPhaseCAttr:                     func() zcl.Attr { return new(ActivePowerMinPhaseC) },
		ActivePowerMaxPhaseCAttr:                     func() zcl.Attr { return new(ActivePowerMaxPhaseC) },
		ReactivePowerPhaseCAttr:                      func() zcl.Attr { return new(ReactivePowerPhaseC) },
		ApparentPowerPhaseCAttr:                      func() zcl.Attr { return new(ApparentPowerPhaseC) },
		PowerFactorPhaseCAttr:                        func() zcl.Attr { return new(PowerFactorPhaseC) },
		AverageRmsVoltageMeasurementPeriodPhaseCAttr: func() zcl.Attr { return new(AverageRmsVoltageMeasurementPeriodPhaseC) },
		AverageRmsOvervoltageCounterPhaseCAttr:       func() zcl.Attr { return new(AverageRmsOvervoltageCounterPhaseC) },
		AverageRmsUndervoltageCounterPhaseCAttr:      func() zcl.Attr { return new(AverageRmsUndervoltageCounterPhaseC) },
		RmsExtremeOvervoltagePeriodPhaseCAttr:        func() zcl.Attr { return new(RmsExtremeOvervoltagePeriodPhaseC) },
		RmsExtremeUndervoltagePeriodPhaseCAttr:       func() zcl.Attr { return new(RmsExtremeUndervoltagePeriodPhaseC) },
		RmsVoltageSagPeriodPhaseCAttr:                func() zcl.Attr { return new(RmsVoltageSagPeriodPhaseC) },
		RmsVoltageSwellPeriodPhaseCAttr:              func() zcl.Attr { return new(RmsVoltageSwellPeriodPhaseC) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type GetProfileInfoResponse struct {
	ProfileCount          zcl.Zu8
	ProfileIntervalPeriod zcl.Zenum8
	MaxNumberOfIntervals  zcl.Zu8
	ListOfAttributes      zcl.Zarray
}

const GetProfileInfoResponseCommand zcl.CommandID = 0

func (v *GetProfileInfoResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.ProfileCount,
		&v.ProfileIntervalPeriod,
		&v.MaxNumberOfIntervals,
		&v.ListOfAttributes,
	}
}

func (v GetProfileInfoResponse) ID() zcl.CommandID {
	return GetProfileInfoResponseCommand
}

func (v GetProfileInfoResponse) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}

func (v GetProfileInfoResponse) MnfCode() []byte {
	return []byte{}
}

func (v GetProfileInfoResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.ProfileCount.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ProfileIntervalPeriod.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.MaxNumberOfIntervals.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ListOfAttributes.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetProfileInfoResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.ProfileCount).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ProfileIntervalPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.MaxNumberOfIntervals).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ListOfAttributes).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v GetProfileInfoResponse) ProfileCountString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.ProfileCount))
}
func (v GetProfileInfoResponse) ProfileIntervalPeriodString() string {
	switch v.ProfileIntervalPeriod {
	case 0x00:
		return "Daily"
	case 0x01:
		return "60 minutes"
	case 0x02:
		return "30 minutes"
	case 0x03:
		return "15 minutes"
	case 0x04:
		return "10 minutes"
	case 0x05:
		return "7.5 minutes"
	case 0x06:
		return "5 minutes"
	case 0x07:
		return "2.5 minutes"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v.ProfileIntervalPeriod))
}
func (v GetProfileInfoResponse) MaxNumberOfIntervalsString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.MaxNumberOfIntervals))
}
func (v GetProfileInfoResponse) ListOfAttributesString() string {
	return zcl.Sprintf("%v", zcl.Zarray(v.ListOfAttributes))
}

func (v GetProfileInfoResponse) String() string {
	var str []string
	str = append(str, "ProfileCount["+v.ProfileCountString()+"]")
	str = append(str, "ProfileIntervalPeriod["+v.ProfileIntervalPeriodString()+"]")
	str = append(str, "MaxNumberOfIntervals["+v.MaxNumberOfIntervalsString()+"]")
	str = append(str, "ListOfAttributes["+v.ListOfAttributesString()+"]")
	return "GetProfileInfoResponse{" + zcl.StrJoin(str, " ") + "}"
}

func (GetProfileInfoResponse) Name() string { return "Get Profile Info Response" }

type GetMeasurementProfileResponse struct {
	StartTime                  zcl.Zutc
	Status                     zcl.Zenum8
	ProfileIntervalPeriod      zcl.Zenum8
	NumberOfIntervalsDelivered zcl.Zu8
	AttributeId                zcl.Zu16
	Intervals                  zcl.Zarray
}

const GetMeasurementProfileResponseCommand zcl.CommandID = 1

func (v *GetMeasurementProfileResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.StartTime,
		&v.Status,
		&v.ProfileIntervalPeriod,
		&v.NumberOfIntervalsDelivered,
		&v.AttributeId,
		&v.Intervals,
	}
}

func (v GetMeasurementProfileResponse) ID() zcl.CommandID {
	return GetMeasurementProfileResponseCommand
}

func (v GetMeasurementProfileResponse) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}

func (v GetMeasurementProfileResponse) MnfCode() []byte {
	return []byte{}
}

func (v GetMeasurementProfileResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.StartTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ProfileIntervalPeriod.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NumberOfIntervalsDelivered.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.AttributeId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Intervals.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetMeasurementProfileResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.StartTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ProfileIntervalPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberOfIntervalsDelivered).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.AttributeId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Intervals).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v GetMeasurementProfileResponse) StartTimeString() string {
	return zcl.Sprintf("%v", zcl.Zutc(v.StartTime))
}
func (v GetMeasurementProfileResponse) StatusString() string {
	switch v.Status {
	case 0x00:
		return "Success"
	case 0x01:
		return "Attribute Profile not supported"
	case 0x02:
		return "Invalid Start Time"
	case 0x03:
		return "More intervals requested than can be returned"
	case 0x04:
		return "No intervals available for the requested time"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v.Status))
}
func (v GetMeasurementProfileResponse) ProfileIntervalPeriodString() string {
	switch v.ProfileIntervalPeriod {
	case 0x00:
		return "Daily"
	case 0x01:
		return "60 minutes"
	case 0x02:
		return "30 minutes"
	case 0x03:
		return "15 minutes"
	case 0x04:
		return "10 minutes"
	case 0x05:
		return "7.5 minutes"
	case 0x06:
		return "5 minutes"
	case 0x07:
		return "2.5 minutes"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v.ProfileIntervalPeriod))
}
func (v GetMeasurementProfileResponse) NumberOfIntervalsDeliveredString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.NumberOfIntervalsDelivered))
}
func (v GetMeasurementProfileResponse) AttributeIdString() string {
	return zcl.Sprintf("%v", zcl.Zu16(v.AttributeId))
}
func (v GetMeasurementProfileResponse) IntervalsString() string {
	return zcl.Sprintf("%v", zcl.Zarray(v.Intervals))
}

func (v GetMeasurementProfileResponse) String() string {
	var str []string
	str = append(str, "StartTime["+v.StartTimeString()+"]")
	str = append(str, "Status["+v.StatusString()+"]")
	str = append(str, "ProfileIntervalPeriod["+v.ProfileIntervalPeriodString()+"]")
	str = append(str, "NumberOfIntervalsDelivered["+v.NumberOfIntervalsDeliveredString()+"]")
	str = append(str, "AttributeId["+v.AttributeIdString()+"]")
	str = append(str, "Intervals["+v.IntervalsString()+"]")
	return "GetMeasurementProfileResponse{" + zcl.StrJoin(str, " ") + "}"
}

func (GetMeasurementProfileResponse) Name() string { return "Get Measurement Profile Response" }

type GetProfileInfo struct {
}

const GetProfileInfoCommand zcl.CommandID = 0

func (v *GetProfileInfo) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v GetProfileInfo) ID() zcl.CommandID {
	return GetProfileInfoCommand
}

func (v GetProfileInfo) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}

func (v GetProfileInfo) MnfCode() []byte {
	return []byte{}
}

func (v GetProfileInfo) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *GetProfileInfo) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

func (v GetProfileInfo) String() string {
	var str []string
	return "GetProfileInfo{" + zcl.StrJoin(str, " ") + "}"
}

func (GetProfileInfo) Name() string { return "Get Profile Info" }

type GetMeasurementProfile struct {
	AttributeId       zcl.Zu16
	StartTime         zcl.Zutc
	NumberOfIntervals zcl.Zu8
}

const GetMeasurementProfileCommand zcl.CommandID = 1

func (v *GetMeasurementProfile) Values() []zcl.Val {
	return []zcl.Val{
		&v.AttributeId,
		&v.StartTime,
		&v.NumberOfIntervals,
	}
}

func (v GetMeasurementProfile) ID() zcl.CommandID {
	return GetMeasurementProfileCommand
}

func (v GetMeasurementProfile) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}

func (v GetMeasurementProfile) MnfCode() []byte {
	return []byte{}
}

func (v GetMeasurementProfile) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.AttributeId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StartTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.NumberOfIntervals.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetMeasurementProfile) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.AttributeId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StartTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberOfIntervals).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

func (v GetMeasurementProfile) AttributeIdString() string {
	return zcl.Sprintf("%v", zcl.Zu16(v.AttributeId))
}
func (v GetMeasurementProfile) StartTimeString() string {
	return zcl.Sprintf("%v", zcl.Zutc(v.StartTime))
}
func (v GetMeasurementProfile) NumberOfIntervalsString() string {
	return zcl.Sprintf("%v", zcl.Zu8(v.NumberOfIntervals))
}

func (v GetMeasurementProfile) String() string {
	var str []string
	str = append(str, "AttributeId["+v.AttributeIdString()+"]")
	str = append(str, "StartTime["+v.StartTimeString()+"]")
	str = append(str, "NumberOfIntervals["+v.NumberOfIntervalsString()+"]")
	return "GetMeasurementProfile{" + zcl.StrJoin(str, " ") + "}"
}

func (GetMeasurementProfile) Name() string { return "Get Measurement Profile" }

// MeasurementType is an autogenerated attribute in the ElectricalMeasurement cluster
type MeasurementType zcl.Zbmp32

const MeasurementTypeAttr zcl.AttrID = 0

func (MeasurementType) ID() zcl.AttrID                { return MeasurementTypeAttr }
func (MeasurementType) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (MeasurementType) Name() string                  { return "Measurement Type" }
func (MeasurementType) Readable() bool                { return true }
func (MeasurementType) Writable() bool                { return false }
func (MeasurementType) Reportable() bool              { return false }
func (MeasurementType) SceneIndex() int               { return -1 }
func (a *MeasurementType) Value() *MeasurementType    { return a }
func (a MeasurementType) MarshalZcl() ([]byte, error) { return zcl.Zbmp32(a).MarshalZcl() }

func (a *MeasurementType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasurementType(*nt)
	return br, err
}

func (a MeasurementType) String() string {
	var bstr []string
	if a.IsActiveMeasurementAc() {
		bstr = append(bstr, "Active measurement (AC)")
	}
	if a.IsReactiveMeasurementAc() {
		bstr = append(bstr, "Reactive measurement (AC)")
	}
	if a.IsApparentMeasurementAc() {
		bstr = append(bstr, "Apparent measurement (AC)")
	}
	if a.IsPhaseAMeasurement() {
		bstr = append(bstr, "Phase A measurement")
	}
	if a.IsPhaseBMeasurement() {
		bstr = append(bstr, "Phase B measurement")
	}
	if a.IsPhaseCMeasurement() {
		bstr = append(bstr, "Phase C measurement")
	}
	if a.IsDcMeasurement() {
		bstr = append(bstr, "DC measurement")
	}
	if a.IsHarmonicsMeasurement() {
		bstr = append(bstr, "Harmonics measurement")
	}
	if a.IsPowerQualityMeasurement() {
		bstr = append(bstr, "Power quality measurement")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a MeasurementType) IsActiveMeasurementAc() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *MeasurementType) SetActiveMeasurementAc(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a MeasurementType) IsReactiveMeasurementAc() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *MeasurementType) SetReactiveMeasurementAc(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a MeasurementType) IsApparentMeasurementAc() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *MeasurementType) SetApparentMeasurementAc(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a MeasurementType) IsPhaseAMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *MeasurementType) SetPhaseAMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a MeasurementType) IsPhaseBMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *MeasurementType) SetPhaseBMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a MeasurementType) IsPhaseCMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *MeasurementType) SetPhaseCMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a MeasurementType) IsDcMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *MeasurementType) SetDcMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 6, b))
}

func (a MeasurementType) IsHarmonicsMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *MeasurementType) SetHarmonicsMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 7, b))
}

func (a MeasurementType) IsPowerQualityMeasurement() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *MeasurementType) SetPowerQualityMeasurement(b bool) {
	*a = MeasurementType(zcl.BitmapSet([]byte(*a), 8, b))
}

// DcVoltage is an autogenerated attribute in the ElectricalMeasurement cluster
type DcVoltage zcl.Zs16

const DcVoltageAttr zcl.AttrID = 256

func (DcVoltage) ID() zcl.AttrID                { return DcVoltageAttr }
func (DcVoltage) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcVoltage) Name() string                  { return "DC Voltage" }
func (DcVoltage) Readable() bool                { return true }
func (DcVoltage) Writable() bool                { return false }
func (DcVoltage) Reportable() bool              { return false }
func (DcVoltage) SceneIndex() int               { return -1 }
func (a *DcVoltage) Value() *DcVoltage          { return a }
func (a DcVoltage) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltage(*nt)
	return br, err
}

func (a DcVoltage) String() string {
	return zcl.Volts.Format(float64(a))
}

// DcVoltageMin is an autogenerated attribute in the ElectricalMeasurement cluster
type DcVoltageMin zcl.Zs16

const DcVoltageMinAttr zcl.AttrID = 257

func (DcVoltageMin) ID() zcl.AttrID                { return DcVoltageMinAttr }
func (DcVoltageMin) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcVoltageMin) Name() string                  { return "DC Voltage Min" }
func (DcVoltageMin) Readable() bool                { return true }
func (DcVoltageMin) Writable() bool                { return false }
func (DcVoltageMin) Reportable() bool              { return false }
func (DcVoltageMin) SceneIndex() int               { return -1 }
func (a *DcVoltageMin) Value() *DcVoltageMin       { return a }
func (a DcVoltageMin) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcVoltageMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageMin(*nt)
	return br, err
}

func (a DcVoltageMin) String() string {
	return zcl.Volts.Format(float64(a))
}

// DcVoltageMax is an autogenerated attribute in the ElectricalMeasurement cluster
type DcVoltageMax zcl.Zs16

const DcVoltageMaxAttr zcl.AttrID = 258

func (DcVoltageMax) ID() zcl.AttrID                { return DcVoltageMaxAttr }
func (DcVoltageMax) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcVoltageMax) Name() string                  { return "DC Voltage Max" }
func (DcVoltageMax) Readable() bool                { return true }
func (DcVoltageMax) Writable() bool                { return false }
func (DcVoltageMax) Reportable() bool              { return false }
func (DcVoltageMax) SceneIndex() int               { return -1 }
func (a *DcVoltageMax) Value() *DcVoltageMax       { return a }
func (a DcVoltageMax) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcVoltageMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageMax(*nt)
	return br, err
}

func (a DcVoltageMax) String() string {
	return zcl.Volts.Format(float64(a))
}

// DcCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type DcCurrent zcl.Zs16

const DcCurrentAttr zcl.AttrID = 259

func (DcCurrent) ID() zcl.AttrID                { return DcCurrentAttr }
func (DcCurrent) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcCurrent) Name() string                  { return "DC Current" }
func (DcCurrent) Readable() bool                { return true }
func (DcCurrent) Writable() bool                { return false }
func (DcCurrent) Reportable() bool              { return false }
func (DcCurrent) SceneIndex() int               { return -1 }
func (a *DcCurrent) Value() *DcCurrent          { return a }
func (a DcCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrent(*nt)
	return br, err
}

func (a DcCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// DcCurrentMin is an autogenerated attribute in the ElectricalMeasurement cluster
type DcCurrentMin zcl.Zs16

const DcCurrentMinAttr zcl.AttrID = 260

func (DcCurrentMin) ID() zcl.AttrID                { return DcCurrentMinAttr }
func (DcCurrentMin) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcCurrentMin) Name() string                  { return "DC Current Min" }
func (DcCurrentMin) Readable() bool                { return true }
func (DcCurrentMin) Writable() bool                { return false }
func (DcCurrentMin) Reportable() bool              { return false }
func (DcCurrentMin) SceneIndex() int               { return -1 }
func (a *DcCurrentMin) Value() *DcCurrentMin       { return a }
func (a DcCurrentMin) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcCurrentMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentMin(*nt)
	return br, err
}

func (a DcCurrentMin) String() string {
	return zcl.Amperes.Format(float64(a))
}

// DcCurrentMax is an autogenerated attribute in the ElectricalMeasurement cluster
type DcCurrentMax zcl.Zs16

const DcCurrentMaxAttr zcl.AttrID = 261

func (DcCurrentMax) ID() zcl.AttrID                { return DcCurrentMaxAttr }
func (DcCurrentMax) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcCurrentMax) Name() string                  { return "DC Current Max" }
func (DcCurrentMax) Readable() bool                { return true }
func (DcCurrentMax) Writable() bool                { return false }
func (DcCurrentMax) Reportable() bool              { return false }
func (DcCurrentMax) SceneIndex() int               { return -1 }
func (a *DcCurrentMax) Value() *DcCurrentMax       { return a }
func (a DcCurrentMax) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcCurrentMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentMax(*nt)
	return br, err
}

func (a DcCurrentMax) String() string {
	return zcl.Amperes.Format(float64(a))
}

// DcPower is an autogenerated attribute in the ElectricalMeasurement cluster
type DcPower zcl.Zs16

const DcPowerAttr zcl.AttrID = 262

func (DcPower) ID() zcl.AttrID                { return DcPowerAttr }
func (DcPower) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcPower) Name() string                  { return "DC Power" }
func (DcPower) Readable() bool                { return true }
func (DcPower) Writable() bool                { return false }
func (DcPower) Reportable() bool              { return false }
func (DcPower) SceneIndex() int               { return -1 }
func (a *DcPower) Value() *DcPower            { return a }
func (a DcPower) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPower(*nt)
	return br, err
}

func (a DcPower) String() string {
	return zcl.Watts.Format(float64(a))
}

// DcPowerMin is an autogenerated attribute in the ElectricalMeasurement cluster
type DcPowerMin zcl.Zs16

const DcPowerMinAttr zcl.AttrID = 263

func (DcPowerMin) ID() zcl.AttrID                { return DcPowerMinAttr }
func (DcPowerMin) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcPowerMin) Name() string                  { return "DC Power Min" }
func (DcPowerMin) Readable() bool                { return true }
func (DcPowerMin) Writable() bool                { return false }
func (DcPowerMin) Reportable() bool              { return false }
func (DcPowerMin) SceneIndex() int               { return -1 }
func (a *DcPowerMin) Value() *DcPowerMin         { return a }
func (a DcPowerMin) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcPowerMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerMin(*nt)
	return br, err
}

func (a DcPowerMin) String() string {
	return zcl.Watts.Format(float64(a))
}

// DcPowerMax is an autogenerated attribute in the ElectricalMeasurement cluster
type DcPowerMax zcl.Zs16

const DcPowerMaxAttr zcl.AttrID = 264

func (DcPowerMax) ID() zcl.AttrID                { return DcPowerMaxAttr }
func (DcPowerMax) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcPowerMax) Name() string                  { return "DC Power Max" }
func (DcPowerMax) Readable() bool                { return true }
func (DcPowerMax) Writable() bool                { return false }
func (DcPowerMax) Reportable() bool              { return false }
func (DcPowerMax) SceneIndex() int               { return -1 }
func (a *DcPowerMax) Value() *DcPowerMax         { return a }
func (a DcPowerMax) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcPowerMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerMax(*nt)
	return br, err
}

func (a DcPowerMax) String() string {
	return zcl.Watts.Format(float64(a))
}

// DcVoltageMultiplier is an autogenerated attribute in the ElectricalMeasurement cluster
type DcVoltageMultiplier zcl.Zu16

const DcVoltageMultiplierAttr zcl.AttrID = 512

func (DcVoltageMultiplier) ID() zcl.AttrID                 { return DcVoltageMultiplierAttr }
func (DcVoltageMultiplier) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (DcVoltageMultiplier) Name() string                   { return "DC Voltage Multiplier" }
func (DcVoltageMultiplier) Readable() bool                 { return true }
func (DcVoltageMultiplier) Writable() bool                 { return false }
func (DcVoltageMultiplier) Reportable() bool               { return false }
func (DcVoltageMultiplier) SceneIndex() int                { return -1 }
func (a *DcVoltageMultiplier) Value() *DcVoltageMultiplier { return a }
func (a DcVoltageMultiplier) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *DcVoltageMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageMultiplier(*nt)
	return br, err
}

func (a DcVoltageMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// DcVoltageDivisor is an autogenerated attribute in the ElectricalMeasurement cluster
type DcVoltageDivisor zcl.Zu16

const DcVoltageDivisorAttr zcl.AttrID = 513

func (DcVoltageDivisor) ID() zcl.AttrID                { return DcVoltageDivisorAttr }
func (DcVoltageDivisor) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcVoltageDivisor) Name() string                  { return "DC Voltage Divisor" }
func (DcVoltageDivisor) Readable() bool                { return true }
func (DcVoltageDivisor) Writable() bool                { return false }
func (DcVoltageDivisor) Reportable() bool              { return false }
func (DcVoltageDivisor) SceneIndex() int               { return -1 }
func (a *DcVoltageDivisor) Value() *DcVoltageDivisor   { return a }
func (a DcVoltageDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *DcVoltageDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageDivisor(*nt)
	return br, err
}

func (a DcVoltageDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// DcCurrentMultiplier is an autogenerated attribute in the ElectricalMeasurement cluster
type DcCurrentMultiplier zcl.Zu16

const DcCurrentMultiplierAttr zcl.AttrID = 514

func (DcCurrentMultiplier) ID() zcl.AttrID                 { return DcCurrentMultiplierAttr }
func (DcCurrentMultiplier) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (DcCurrentMultiplier) Name() string                   { return "DC Current Multiplier" }
func (DcCurrentMultiplier) Readable() bool                 { return true }
func (DcCurrentMultiplier) Writable() bool                 { return false }
func (DcCurrentMultiplier) Reportable() bool               { return false }
func (DcCurrentMultiplier) SceneIndex() int                { return -1 }
func (a *DcCurrentMultiplier) Value() *DcCurrentMultiplier { return a }
func (a DcCurrentMultiplier) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *DcCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentMultiplier(*nt)
	return br, err
}

func (a DcCurrentMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// DcCurrentDivisor is an autogenerated attribute in the ElectricalMeasurement cluster
type DcCurrentDivisor zcl.Zu16

const DcCurrentDivisorAttr zcl.AttrID = 515

func (DcCurrentDivisor) ID() zcl.AttrID                { return DcCurrentDivisorAttr }
func (DcCurrentDivisor) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcCurrentDivisor) Name() string                  { return "DC Current Divisor" }
func (DcCurrentDivisor) Readable() bool                { return true }
func (DcCurrentDivisor) Writable() bool                { return false }
func (DcCurrentDivisor) Reportable() bool              { return false }
func (DcCurrentDivisor) SceneIndex() int               { return -1 }
func (a *DcCurrentDivisor) Value() *DcCurrentDivisor   { return a }
func (a DcCurrentDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *DcCurrentDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentDivisor(*nt)
	return br, err
}

func (a DcCurrentDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// DcPowerMultiplier is an autogenerated attribute in the ElectricalMeasurement cluster
type DcPowerMultiplier zcl.Zu16

const DcPowerMultiplierAttr zcl.AttrID = 516

func (DcPowerMultiplier) ID() zcl.AttrID                { return DcPowerMultiplierAttr }
func (DcPowerMultiplier) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcPowerMultiplier) Name() string                  { return "DC Power Multiplier" }
func (DcPowerMultiplier) Readable() bool                { return true }
func (DcPowerMultiplier) Writable() bool                { return false }
func (DcPowerMultiplier) Reportable() bool              { return false }
func (DcPowerMultiplier) SceneIndex() int               { return -1 }
func (a *DcPowerMultiplier) Value() *DcPowerMultiplier  { return a }
func (a DcPowerMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *DcPowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerMultiplier(*nt)
	return br, err
}

func (a DcPowerMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// DcPowerDivisor is an autogenerated attribute in the ElectricalMeasurement cluster
type DcPowerDivisor zcl.Zu16

const DcPowerDivisorAttr zcl.AttrID = 517

func (DcPowerDivisor) ID() zcl.AttrID                { return DcPowerDivisorAttr }
func (DcPowerDivisor) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcPowerDivisor) Name() string                  { return "DC Power Divisor" }
func (DcPowerDivisor) Readable() bool                { return true }
func (DcPowerDivisor) Writable() bool                { return false }
func (DcPowerDivisor) Reportable() bool              { return false }
func (DcPowerDivisor) SceneIndex() int               { return -1 }
func (a *DcPowerDivisor) Value() *DcPowerDivisor     { return a }
func (a DcPowerDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *DcPowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerDivisor(*nt)
	return br, err
}

func (a DcPowerDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AcFrequency is an autogenerated attribute in the ElectricalMeasurement cluster
type AcFrequency zcl.Zu16

const AcFrequencyAttr zcl.AttrID = 768

func (AcFrequency) ID() zcl.AttrID                { return AcFrequencyAttr }
func (AcFrequency) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (AcFrequency) Name() string                  { return "AC Frequency" }
func (AcFrequency) Readable() bool                { return true }
func (AcFrequency) Writable() bool                { return false }
func (AcFrequency) Reportable() bool              { return false }
func (AcFrequency) SceneIndex() int               { return -1 }
func (a *AcFrequency) Value() *AcFrequency        { return a }
func (a AcFrequency) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AcFrequency) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequency(*nt)
	return br, err
}

func (a AcFrequency) String() string {
	return zcl.Hertz.Format(float64(a))
}

// AcFrequencyMin is an autogenerated attribute in the ElectricalMeasurement cluster
type AcFrequencyMin zcl.Zu16

const AcFrequencyMinAttr zcl.AttrID = 769

func (AcFrequencyMin) ID() zcl.AttrID                { return AcFrequencyMinAttr }
func (AcFrequencyMin) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (AcFrequencyMin) Name() string                  { return "AC Frequency Min" }
func (AcFrequencyMin) Readable() bool                { return true }
func (AcFrequencyMin) Writable() bool                { return false }
func (AcFrequencyMin) Reportable() bool              { return false }
func (AcFrequencyMin) SceneIndex() int               { return -1 }
func (a *AcFrequencyMin) Value() *AcFrequencyMin     { return a }
func (a AcFrequencyMin) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AcFrequencyMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyMin(*nt)
	return br, err
}

func (a AcFrequencyMin) String() string {
	return zcl.Hertz.Format(float64(a))
}

// AcFrequencyMax is an autogenerated attribute in the ElectricalMeasurement cluster
type AcFrequencyMax zcl.Zu16

const AcFrequencyMaxAttr zcl.AttrID = 770

func (AcFrequencyMax) ID() zcl.AttrID                { return AcFrequencyMaxAttr }
func (AcFrequencyMax) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (AcFrequencyMax) Name() string                  { return "AC Frequency Max" }
func (AcFrequencyMax) Readable() bool                { return true }
func (AcFrequencyMax) Writable() bool                { return false }
func (AcFrequencyMax) Reportable() bool              { return false }
func (AcFrequencyMax) SceneIndex() int               { return -1 }
func (a *AcFrequencyMax) Value() *AcFrequencyMax     { return a }
func (a AcFrequencyMax) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AcFrequencyMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyMax(*nt)
	return br, err
}

func (a AcFrequencyMax) String() string {
	return zcl.Hertz.Format(float64(a))
}

// NeutralCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type NeutralCurrent zcl.Zu16

const NeutralCurrentAttr zcl.AttrID = 771

func (NeutralCurrent) ID() zcl.AttrID                { return NeutralCurrentAttr }
func (NeutralCurrent) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (NeutralCurrent) Name() string                  { return "Neutral Current" }
func (NeutralCurrent) Readable() bool                { return true }
func (NeutralCurrent) Writable() bool                { return false }
func (NeutralCurrent) Reportable() bool              { return false }
func (NeutralCurrent) SceneIndex() int               { return -1 }
func (a *NeutralCurrent) Value() *NeutralCurrent     { return a }
func (a NeutralCurrent) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *NeutralCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeutralCurrent(*nt)
	return br, err
}

func (a NeutralCurrent) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// TotalActivePower is an autogenerated attribute in the ElectricalMeasurement cluster
type TotalActivePower zcl.Zs32

const TotalActivePowerAttr zcl.AttrID = 772

func (TotalActivePower) ID() zcl.AttrID                { return TotalActivePowerAttr }
func (TotalActivePower) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (TotalActivePower) Name() string                  { return "Total Active Power" }
func (TotalActivePower) Readable() bool                { return true }
func (TotalActivePower) Writable() bool                { return false }
func (TotalActivePower) Reportable() bool              { return false }
func (TotalActivePower) SceneIndex() int               { return -1 }
func (a *TotalActivePower) Value() *TotalActivePower   { return a }
func (a TotalActivePower) MarshalZcl() ([]byte, error) { return zcl.Zs32(a).MarshalZcl() }

func (a *TotalActivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = TotalActivePower(*nt)
	return br, err
}

func (a TotalActivePower) String() string {
	return zcl.Kilowatts.Format(float64(a))
}

// TotalReactivePower is an autogenerated attribute in the ElectricalMeasurement cluster
type TotalReactivePower zcl.Zs32

const TotalReactivePowerAttr zcl.AttrID = 773

func (TotalReactivePower) ID() zcl.AttrID                { return TotalReactivePowerAttr }
func (TotalReactivePower) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (TotalReactivePower) Name() string                  { return "Total Reactive Power" }
func (TotalReactivePower) Readable() bool                { return true }
func (TotalReactivePower) Writable() bool                { return false }
func (TotalReactivePower) Reportable() bool              { return false }
func (TotalReactivePower) SceneIndex() int               { return -1 }
func (a *TotalReactivePower) Value() *TotalReactivePower { return a }
func (a TotalReactivePower) MarshalZcl() ([]byte, error) { return zcl.Zs32(a).MarshalZcl() }

func (a *TotalReactivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = TotalReactivePower(*nt)
	return br, err
}

func (a TotalReactivePower) String() string {
	return zcl.KiloVoltAmperes.Format(float64(a))
}

// TotalApparentPower is an autogenerated attribute in the ElectricalMeasurement cluster
type TotalApparentPower zcl.Zu32

const TotalApparentPowerAttr zcl.AttrID = 774

func (TotalApparentPower) ID() zcl.AttrID                { return TotalApparentPowerAttr }
func (TotalApparentPower) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (TotalApparentPower) Name() string                  { return "Total Apparent Power" }
func (TotalApparentPower) Readable() bool                { return true }
func (TotalApparentPower) Writable() bool                { return false }
func (TotalApparentPower) Reportable() bool              { return false }
func (TotalApparentPower) SceneIndex() int               { return -1 }
func (a *TotalApparentPower) Value() *TotalApparentPower { return a }
func (a TotalApparentPower) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *TotalApparentPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = TotalApparentPower(*nt)
	return br, err
}

func (a TotalApparentPower) String() string {
	return zcl.KiloVoltAmperes.Format(float64(a))
}

// Measured1StHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type Measured1StHarmonicCurrent zcl.Zs16

const Measured1StHarmonicCurrentAttr zcl.AttrID = 775

func (Measured1StHarmonicCurrent) ID() zcl.AttrID                        { return Measured1StHarmonicCurrentAttr }
func (Measured1StHarmonicCurrent) Cluster() zcl.ClusterID                { return ElectricalMeasurementID }
func (Measured1StHarmonicCurrent) Name() string                          { return "Measured 1st Harmonic Current" }
func (Measured1StHarmonicCurrent) Readable() bool                        { return true }
func (Measured1StHarmonicCurrent) Writable() bool                        { return false }
func (Measured1StHarmonicCurrent) Reportable() bool                      { return false }
func (Measured1StHarmonicCurrent) SceneIndex() int                       { return -1 }
func (a *Measured1StHarmonicCurrent) Value() *Measured1StHarmonicCurrent { return a }
func (a Measured1StHarmonicCurrent) MarshalZcl() ([]byte, error)         { return zcl.Zs16(a).MarshalZcl() }

func (a *Measured1StHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured1StHarmonicCurrent(*nt)
	return br, err
}

func (a Measured1StHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// Measured3RdHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type Measured3RdHarmonicCurrent zcl.Zs16

const Measured3RdHarmonicCurrentAttr zcl.AttrID = 776

func (Measured3RdHarmonicCurrent) ID() zcl.AttrID                        { return Measured3RdHarmonicCurrentAttr }
func (Measured3RdHarmonicCurrent) Cluster() zcl.ClusterID                { return ElectricalMeasurementID }
func (Measured3RdHarmonicCurrent) Name() string                          { return "Measured 3rd Harmonic Current" }
func (Measured3RdHarmonicCurrent) Readable() bool                        { return true }
func (Measured3RdHarmonicCurrent) Writable() bool                        { return false }
func (Measured3RdHarmonicCurrent) Reportable() bool                      { return false }
func (Measured3RdHarmonicCurrent) SceneIndex() int                       { return -1 }
func (a *Measured3RdHarmonicCurrent) Value() *Measured3RdHarmonicCurrent { return a }
func (a Measured3RdHarmonicCurrent) MarshalZcl() ([]byte, error)         { return zcl.Zs16(a).MarshalZcl() }

func (a *Measured3RdHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured3RdHarmonicCurrent(*nt)
	return br, err
}

func (a Measured3RdHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// Measured5ThHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type Measured5ThHarmonicCurrent zcl.Zs16

const Measured5ThHarmonicCurrentAttr zcl.AttrID = 777

func (Measured5ThHarmonicCurrent) ID() zcl.AttrID                        { return Measured5ThHarmonicCurrentAttr }
func (Measured5ThHarmonicCurrent) Cluster() zcl.ClusterID                { return ElectricalMeasurementID }
func (Measured5ThHarmonicCurrent) Name() string                          { return "Measured 5th Harmonic Current" }
func (Measured5ThHarmonicCurrent) Readable() bool                        { return true }
func (Measured5ThHarmonicCurrent) Writable() bool                        { return false }
func (Measured5ThHarmonicCurrent) Reportable() bool                      { return false }
func (Measured5ThHarmonicCurrent) SceneIndex() int                       { return -1 }
func (a *Measured5ThHarmonicCurrent) Value() *Measured5ThHarmonicCurrent { return a }
func (a Measured5ThHarmonicCurrent) MarshalZcl() ([]byte, error)         { return zcl.Zs16(a).MarshalZcl() }

func (a *Measured5ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured5ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured5ThHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// Measured7ThHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type Measured7ThHarmonicCurrent zcl.Zs16

const Measured7ThHarmonicCurrentAttr zcl.AttrID = 778

func (Measured7ThHarmonicCurrent) ID() zcl.AttrID                        { return Measured7ThHarmonicCurrentAttr }
func (Measured7ThHarmonicCurrent) Cluster() zcl.ClusterID                { return ElectricalMeasurementID }
func (Measured7ThHarmonicCurrent) Name() string                          { return "Measured 7th Harmonic Current" }
func (Measured7ThHarmonicCurrent) Readable() bool                        { return true }
func (Measured7ThHarmonicCurrent) Writable() bool                        { return false }
func (Measured7ThHarmonicCurrent) Reportable() bool                      { return false }
func (Measured7ThHarmonicCurrent) SceneIndex() int                       { return -1 }
func (a *Measured7ThHarmonicCurrent) Value() *Measured7ThHarmonicCurrent { return a }
func (a Measured7ThHarmonicCurrent) MarshalZcl() ([]byte, error)         { return zcl.Zs16(a).MarshalZcl() }

func (a *Measured7ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured7ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured7ThHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// Measured9ThHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type Measured9ThHarmonicCurrent zcl.Zs16

const Measured9ThHarmonicCurrentAttr zcl.AttrID = 779

func (Measured9ThHarmonicCurrent) ID() zcl.AttrID                        { return Measured9ThHarmonicCurrentAttr }
func (Measured9ThHarmonicCurrent) Cluster() zcl.ClusterID                { return ElectricalMeasurementID }
func (Measured9ThHarmonicCurrent) Name() string                          { return "Measured 9th Harmonic Current" }
func (Measured9ThHarmonicCurrent) Readable() bool                        { return true }
func (Measured9ThHarmonicCurrent) Writable() bool                        { return false }
func (Measured9ThHarmonicCurrent) Reportable() bool                      { return false }
func (Measured9ThHarmonicCurrent) SceneIndex() int                       { return -1 }
func (a *Measured9ThHarmonicCurrent) Value() *Measured9ThHarmonicCurrent { return a }
func (a Measured9ThHarmonicCurrent) MarshalZcl() ([]byte, error)         { return zcl.Zs16(a).MarshalZcl() }

func (a *Measured9ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured9ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured9ThHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// Measured11ThHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type Measured11ThHarmonicCurrent zcl.Zs16

const Measured11ThHarmonicCurrentAttr zcl.AttrID = 780

func (Measured11ThHarmonicCurrent) ID() zcl.AttrID                         { return Measured11ThHarmonicCurrentAttr }
func (Measured11ThHarmonicCurrent) Cluster() zcl.ClusterID                 { return ElectricalMeasurementID }
func (Measured11ThHarmonicCurrent) Name() string                           { return "Measured 11th Harmonic Current" }
func (Measured11ThHarmonicCurrent) Readable() bool                         { return true }
func (Measured11ThHarmonicCurrent) Writable() bool                         { return false }
func (Measured11ThHarmonicCurrent) Reportable() bool                       { return false }
func (Measured11ThHarmonicCurrent) SceneIndex() int                        { return -1 }
func (a *Measured11ThHarmonicCurrent) Value() *Measured11ThHarmonicCurrent { return a }
func (a Measured11ThHarmonicCurrent) MarshalZcl() ([]byte, error)          { return zcl.Zs16(a).MarshalZcl() }

func (a *Measured11ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured11ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured11ThHarmonicCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// MeasuredPhase1StHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type MeasuredPhase1StHarmonicCurrent zcl.Zs16

const MeasuredPhase1StHarmonicCurrentAttr zcl.AttrID = 781

func (MeasuredPhase1StHarmonicCurrent) ID() zcl.AttrID                             { return MeasuredPhase1StHarmonicCurrentAttr }
func (MeasuredPhase1StHarmonicCurrent) Cluster() zcl.ClusterID                     { return ElectricalMeasurementID }
func (MeasuredPhase1StHarmonicCurrent) Name() string                               { return "Measured Phase 1st Harmonic Current" }
func (MeasuredPhase1StHarmonicCurrent) Readable() bool                             { return true }
func (MeasuredPhase1StHarmonicCurrent) Writable() bool                             { return false }
func (MeasuredPhase1StHarmonicCurrent) Reportable() bool                           { return false }
func (MeasuredPhase1StHarmonicCurrent) SceneIndex() int                            { return -1 }
func (a *MeasuredPhase1StHarmonicCurrent) Value() *MeasuredPhase1StHarmonicCurrent { return a }
func (a MeasuredPhase1StHarmonicCurrent) MarshalZcl() ([]byte, error)              { return zcl.Zs16(a).MarshalZcl() }

func (a *MeasuredPhase1StHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase1StHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase1StHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(a))
}

// MeasuredPhase3RdHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type MeasuredPhase3RdHarmonicCurrent zcl.Zs16

const MeasuredPhase3RdHarmonicCurrentAttr zcl.AttrID = 782

func (MeasuredPhase3RdHarmonicCurrent) ID() zcl.AttrID                             { return MeasuredPhase3RdHarmonicCurrentAttr }
func (MeasuredPhase3RdHarmonicCurrent) Cluster() zcl.ClusterID                     { return ElectricalMeasurementID }
func (MeasuredPhase3RdHarmonicCurrent) Name() string                               { return "Measured Phase 3rd Harmonic Current" }
func (MeasuredPhase3RdHarmonicCurrent) Readable() bool                             { return true }
func (MeasuredPhase3RdHarmonicCurrent) Writable() bool                             { return false }
func (MeasuredPhase3RdHarmonicCurrent) Reportable() bool                           { return false }
func (MeasuredPhase3RdHarmonicCurrent) SceneIndex() int                            { return -1 }
func (a *MeasuredPhase3RdHarmonicCurrent) Value() *MeasuredPhase3RdHarmonicCurrent { return a }
func (a MeasuredPhase3RdHarmonicCurrent) MarshalZcl() ([]byte, error)              { return zcl.Zs16(a).MarshalZcl() }

func (a *MeasuredPhase3RdHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase3RdHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase3RdHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(a))
}

// MeasuredPhase5ThHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type MeasuredPhase5ThHarmonicCurrent zcl.Zs16

const MeasuredPhase5ThHarmonicCurrentAttr zcl.AttrID = 783

func (MeasuredPhase5ThHarmonicCurrent) ID() zcl.AttrID                             { return MeasuredPhase5ThHarmonicCurrentAttr }
func (MeasuredPhase5ThHarmonicCurrent) Cluster() zcl.ClusterID                     { return ElectricalMeasurementID }
func (MeasuredPhase5ThHarmonicCurrent) Name() string                               { return "Measured Phase 5th Harmonic Current" }
func (MeasuredPhase5ThHarmonicCurrent) Readable() bool                             { return true }
func (MeasuredPhase5ThHarmonicCurrent) Writable() bool                             { return false }
func (MeasuredPhase5ThHarmonicCurrent) Reportable() bool                           { return false }
func (MeasuredPhase5ThHarmonicCurrent) SceneIndex() int                            { return -1 }
func (a *MeasuredPhase5ThHarmonicCurrent) Value() *MeasuredPhase5ThHarmonicCurrent { return a }
func (a MeasuredPhase5ThHarmonicCurrent) MarshalZcl() ([]byte, error)              { return zcl.Zs16(a).MarshalZcl() }

func (a *MeasuredPhase5ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase5ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase5ThHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(a))
}

// MeasuredPhase7ThHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type MeasuredPhase7ThHarmonicCurrent zcl.Zs16

const MeasuredPhase7ThHarmonicCurrentAttr zcl.AttrID = 784

func (MeasuredPhase7ThHarmonicCurrent) ID() zcl.AttrID                             { return MeasuredPhase7ThHarmonicCurrentAttr }
func (MeasuredPhase7ThHarmonicCurrent) Cluster() zcl.ClusterID                     { return ElectricalMeasurementID }
func (MeasuredPhase7ThHarmonicCurrent) Name() string                               { return "Measured Phase 7th Harmonic Current" }
func (MeasuredPhase7ThHarmonicCurrent) Readable() bool                             { return true }
func (MeasuredPhase7ThHarmonicCurrent) Writable() bool                             { return false }
func (MeasuredPhase7ThHarmonicCurrent) Reportable() bool                           { return false }
func (MeasuredPhase7ThHarmonicCurrent) SceneIndex() int                            { return -1 }
func (a *MeasuredPhase7ThHarmonicCurrent) Value() *MeasuredPhase7ThHarmonicCurrent { return a }
func (a MeasuredPhase7ThHarmonicCurrent) MarshalZcl() ([]byte, error)              { return zcl.Zs16(a).MarshalZcl() }

func (a *MeasuredPhase7ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase7ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase7ThHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(a))
}

// MeasuredPhase9ThHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type MeasuredPhase9ThHarmonicCurrent zcl.Zs16

const MeasuredPhase9ThHarmonicCurrentAttr zcl.AttrID = 785

func (MeasuredPhase9ThHarmonicCurrent) ID() zcl.AttrID                             { return MeasuredPhase9ThHarmonicCurrentAttr }
func (MeasuredPhase9ThHarmonicCurrent) Cluster() zcl.ClusterID                     { return ElectricalMeasurementID }
func (MeasuredPhase9ThHarmonicCurrent) Name() string                               { return "Measured Phase 9th Harmonic Current" }
func (MeasuredPhase9ThHarmonicCurrent) Readable() bool                             { return true }
func (MeasuredPhase9ThHarmonicCurrent) Writable() bool                             { return false }
func (MeasuredPhase9ThHarmonicCurrent) Reportable() bool                           { return false }
func (MeasuredPhase9ThHarmonicCurrent) SceneIndex() int                            { return -1 }
func (a *MeasuredPhase9ThHarmonicCurrent) Value() *MeasuredPhase9ThHarmonicCurrent { return a }
func (a MeasuredPhase9ThHarmonicCurrent) MarshalZcl() ([]byte, error)              { return zcl.Zs16(a).MarshalZcl() }

func (a *MeasuredPhase9ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase9ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase9ThHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(a))
}

// MeasuredPhase11ThHarmonicCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type MeasuredPhase11ThHarmonicCurrent zcl.Zs16

const MeasuredPhase11ThHarmonicCurrentAttr zcl.AttrID = 786

func (MeasuredPhase11ThHarmonicCurrent) ID() zcl.AttrID                              { return MeasuredPhase11ThHarmonicCurrentAttr }
func (MeasuredPhase11ThHarmonicCurrent) Cluster() zcl.ClusterID                      { return ElectricalMeasurementID }
func (MeasuredPhase11ThHarmonicCurrent) Name() string                                { return "Measured Phase 11th Harmonic Current" }
func (MeasuredPhase11ThHarmonicCurrent) Readable() bool                              { return true }
func (MeasuredPhase11ThHarmonicCurrent) Writable() bool                              { return false }
func (MeasuredPhase11ThHarmonicCurrent) Reportable() bool                            { return false }
func (MeasuredPhase11ThHarmonicCurrent) SceneIndex() int                             { return -1 }
func (a *MeasuredPhase11ThHarmonicCurrent) Value() *MeasuredPhase11ThHarmonicCurrent { return a }
func (a MeasuredPhase11ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}

func (a *MeasuredPhase11ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase11ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase11ThHarmonicCurrent) String() string {
	return zcl.DegreesPhase.Format(float64(a))
}

// AcFrequencyMultiplier is an autogenerated attribute in the ElectricalMeasurement cluster
type AcFrequencyMultiplier zcl.Zu16

const AcFrequencyMultiplierAttr zcl.AttrID = 1024

func (AcFrequencyMultiplier) ID() zcl.AttrID                   { return AcFrequencyMultiplierAttr }
func (AcFrequencyMultiplier) Cluster() zcl.ClusterID           { return ElectricalMeasurementID }
func (AcFrequencyMultiplier) Name() string                     { return "AC Frequency Multiplier" }
func (AcFrequencyMultiplier) Readable() bool                   { return true }
func (AcFrequencyMultiplier) Writable() bool                   { return false }
func (AcFrequencyMultiplier) Reportable() bool                 { return false }
func (AcFrequencyMultiplier) SceneIndex() int                  { return -1 }
func (a *AcFrequencyMultiplier) Value() *AcFrequencyMultiplier { return a }
func (a AcFrequencyMultiplier) MarshalZcl() ([]byte, error)    { return zcl.Zu16(a).MarshalZcl() }

func (a *AcFrequencyMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyMultiplier(*nt)
	return br, err
}

func (a AcFrequencyMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AcFrequencyDivisor is an autogenerated attribute in the ElectricalMeasurement cluster
type AcFrequencyDivisor zcl.Zu16

const AcFrequencyDivisorAttr zcl.AttrID = 1025

func (AcFrequencyDivisor) ID() zcl.AttrID                { return AcFrequencyDivisorAttr }
func (AcFrequencyDivisor) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (AcFrequencyDivisor) Name() string                  { return "AC Frequency Divisor" }
func (AcFrequencyDivisor) Readable() bool                { return true }
func (AcFrequencyDivisor) Writable() bool                { return false }
func (AcFrequencyDivisor) Reportable() bool              { return false }
func (AcFrequencyDivisor) SceneIndex() int               { return -1 }
func (a *AcFrequencyDivisor) Value() *AcFrequencyDivisor { return a }
func (a AcFrequencyDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AcFrequencyDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyDivisor(*nt)
	return br, err
}

func (a AcFrequencyDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// PowerMultiplier is an autogenerated attribute in the ElectricalMeasurement cluster
type PowerMultiplier zcl.Zu32

const PowerMultiplierAttr zcl.AttrID = 1026

func (PowerMultiplier) ID() zcl.AttrID                { return PowerMultiplierAttr }
func (PowerMultiplier) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (PowerMultiplier) Name() string                  { return "Power Multiplier" }
func (PowerMultiplier) Readable() bool                { return true }
func (PowerMultiplier) Writable() bool                { return false }
func (PowerMultiplier) Reportable() bool              { return false }
func (PowerMultiplier) SceneIndex() int               { return -1 }
func (a *PowerMultiplier) Value() *PowerMultiplier    { return a }
func (a PowerMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *PowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerMultiplier(*nt)
	return br, err
}

func (a PowerMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

// PowerDivisor is an autogenerated attribute in the ElectricalMeasurement cluster
type PowerDivisor zcl.Zu32

const PowerDivisorAttr zcl.AttrID = 1027

func (PowerDivisor) ID() zcl.AttrID                { return PowerDivisorAttr }
func (PowerDivisor) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (PowerDivisor) Name() string                  { return "Power Divisor" }
func (PowerDivisor) Readable() bool                { return true }
func (PowerDivisor) Writable() bool                { return false }
func (PowerDivisor) Reportable() bool              { return false }
func (PowerDivisor) SceneIndex() int               { return -1 }
func (a *PowerDivisor) Value() *PowerDivisor       { return a }
func (a PowerDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu32(a).MarshalZcl() }

func (a *PowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerDivisor(*nt)
	return br, err
}

func (a PowerDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

// HarmonicCurrentMultiplier is an autogenerated attribute in the ElectricalMeasurement cluster
type HarmonicCurrentMultiplier zcl.Zs8

const HarmonicCurrentMultiplierAttr zcl.AttrID = 1028

func (HarmonicCurrentMultiplier) ID() zcl.AttrID                       { return HarmonicCurrentMultiplierAttr }
func (HarmonicCurrentMultiplier) Cluster() zcl.ClusterID               { return ElectricalMeasurementID }
func (HarmonicCurrentMultiplier) Name() string                         { return "Harmonic Current Multiplier" }
func (HarmonicCurrentMultiplier) Readable() bool                       { return true }
func (HarmonicCurrentMultiplier) Writable() bool                       { return false }
func (HarmonicCurrentMultiplier) Reportable() bool                     { return false }
func (HarmonicCurrentMultiplier) SceneIndex() int                      { return -1 }
func (a *HarmonicCurrentMultiplier) Value() *HarmonicCurrentMultiplier { return a }
func (a HarmonicCurrentMultiplier) MarshalZcl() ([]byte, error)        { return zcl.Zs8(a).MarshalZcl() }

func (a *HarmonicCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = HarmonicCurrentMultiplier(*nt)
	return br, err
}

func (a HarmonicCurrentMultiplier) String() string {
	return zcl.Amperes.Format(float64(a))
}

// PhaseHarmonicCurrentMultiplier is an autogenerated attribute in the ElectricalMeasurement cluster
type PhaseHarmonicCurrentMultiplier zcl.Zs8

const PhaseHarmonicCurrentMultiplierAttr zcl.AttrID = 1029

func (PhaseHarmonicCurrentMultiplier) ID() zcl.AttrID                            { return PhaseHarmonicCurrentMultiplierAttr }
func (PhaseHarmonicCurrentMultiplier) Cluster() zcl.ClusterID                    { return ElectricalMeasurementID }
func (PhaseHarmonicCurrentMultiplier) Name() string                              { return "Phase Harmonic Current Multiplier" }
func (PhaseHarmonicCurrentMultiplier) Readable() bool                            { return true }
func (PhaseHarmonicCurrentMultiplier) Writable() bool                            { return false }
func (PhaseHarmonicCurrentMultiplier) Reportable() bool                          { return false }
func (PhaseHarmonicCurrentMultiplier) SceneIndex() int                           { return -1 }
func (a *PhaseHarmonicCurrentMultiplier) Value() *PhaseHarmonicCurrentMultiplier { return a }
func (a PhaseHarmonicCurrentMultiplier) MarshalZcl() ([]byte, error)             { return zcl.Zs8(a).MarshalZcl() }

func (a *PhaseHarmonicCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PhaseHarmonicCurrentMultiplier(*nt)
	return br, err
}

func (a PhaseHarmonicCurrentMultiplier) String() string {
	return zcl.DegreesPhase.Format(float64(a))
}

// LineCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type LineCurrent zcl.Zu16

const LineCurrentAttr zcl.AttrID = 1281

func (LineCurrent) ID() zcl.AttrID                { return LineCurrentAttr }
func (LineCurrent) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (LineCurrent) Name() string                  { return "Line Current" }
func (LineCurrent) Readable() bool                { return true }
func (LineCurrent) Writable() bool                { return false }
func (LineCurrent) Reportable() bool              { return false }
func (LineCurrent) SceneIndex() int               { return -1 }
func (a *LineCurrent) Value() *LineCurrent        { return a }
func (a LineCurrent) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LineCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LineCurrent(*nt)
	return br, err
}

func (a LineCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// ActiveCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type ActiveCurrent zcl.Zs16

const ActiveCurrentAttr zcl.AttrID = 1282

func (ActiveCurrent) ID() zcl.AttrID                { return ActiveCurrentAttr }
func (ActiveCurrent) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (ActiveCurrent) Name() string                  { return "Active Current" }
func (ActiveCurrent) Readable() bool                { return true }
func (ActiveCurrent) Writable() bool                { return false }
func (ActiveCurrent) Reportable() bool              { return false }
func (ActiveCurrent) SceneIndex() int               { return -1 }
func (a *ActiveCurrent) Value() *ActiveCurrent      { return a }
func (a ActiveCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *ActiveCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveCurrent(*nt)
	return br, err
}

func (a ActiveCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// ReactiveCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type ReactiveCurrent zcl.Zs16

const ReactiveCurrentAttr zcl.AttrID = 1283

func (ReactiveCurrent) ID() zcl.AttrID                { return ReactiveCurrentAttr }
func (ReactiveCurrent) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (ReactiveCurrent) Name() string                  { return "Reactive Current" }
func (ReactiveCurrent) Readable() bool                { return true }
func (ReactiveCurrent) Writable() bool                { return false }
func (ReactiveCurrent) Reportable() bool              { return false }
func (ReactiveCurrent) SceneIndex() int               { return -1 }
func (a *ReactiveCurrent) Value() *ReactiveCurrent    { return a }
func (a ReactiveCurrent) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *ReactiveCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactiveCurrent(*nt)
	return br, err
}

func (a ReactiveCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// RmsVoltage is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltage zcl.Zu16

const RmsVoltageAttr zcl.AttrID = 1285

func (RmsVoltage) ID() zcl.AttrID                { return RmsVoltageAttr }
func (RmsVoltage) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsVoltage) Name() string                  { return "RMS Voltage" }
func (RmsVoltage) Readable() bool                { return true }
func (RmsVoltage) Writable() bool                { return false }
func (RmsVoltage) Reportable() bool              { return false }
func (RmsVoltage) SceneIndex() int               { return -1 }
func (a *RmsVoltage) Value() *RmsVoltage         { return a }
func (a RmsVoltage) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltage(*nt)
	return br, err
}

func (a RmsVoltage) String() string {
	return zcl.Volts.Format(float64(a))
}

// RmsVoltageMin is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageMin zcl.Zu16

const RmsVoltageMinAttr zcl.AttrID = 1286

func (RmsVoltageMin) ID() zcl.AttrID                { return RmsVoltageMinAttr }
func (RmsVoltageMin) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsVoltageMin) Name() string                  { return "RMS Voltage Min" }
func (RmsVoltageMin) Readable() bool                { return true }
func (RmsVoltageMin) Writable() bool                { return false }
func (RmsVoltageMin) Reportable() bool              { return false }
func (RmsVoltageMin) SceneIndex() int               { return -1 }
func (a *RmsVoltageMin) Value() *RmsVoltageMin      { return a }
func (a RmsVoltageMin) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMin(*nt)
	return br, err
}

func (a RmsVoltageMin) String() string {
	return zcl.Volts.Format(float64(a))
}

// RmsVoltageMax is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageMax zcl.Zu16

const RmsVoltageMaxAttr zcl.AttrID = 1287

func (RmsVoltageMax) ID() zcl.AttrID                { return RmsVoltageMaxAttr }
func (RmsVoltageMax) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsVoltageMax) Name() string                  { return "RMS Voltage Max" }
func (RmsVoltageMax) Readable() bool                { return true }
func (RmsVoltageMax) Writable() bool                { return false }
func (RmsVoltageMax) Reportable() bool              { return false }
func (RmsVoltageMax) SceneIndex() int               { return -1 }
func (a *RmsVoltageMax) Value() *RmsVoltageMax      { return a }
func (a RmsVoltageMax) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMax(*nt)
	return br, err
}

func (a RmsVoltageMax) String() string {
	return zcl.Volts.Format(float64(a))
}

// RmsCurrent is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsCurrent zcl.Zu16

const RmsCurrentAttr zcl.AttrID = 1288

func (RmsCurrent) ID() zcl.AttrID                { return RmsCurrentAttr }
func (RmsCurrent) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsCurrent) Name() string                  { return "RMS Current" }
func (RmsCurrent) Readable() bool                { return true }
func (RmsCurrent) Writable() bool                { return false }
func (RmsCurrent) Reportable() bool              { return false }
func (RmsCurrent) SceneIndex() int               { return -1 }
func (a *RmsCurrent) Value() *RmsCurrent         { return a }
func (a RmsCurrent) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrent(*nt)
	return br, err
}

func (a RmsCurrent) String() string {
	return zcl.Amperes.Format(float64(a))
}

// RmsCurrentMin is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsCurrentMin zcl.Zu16

const RmsCurrentMinAttr zcl.AttrID = 1289

func (RmsCurrentMin) ID() zcl.AttrID                { return RmsCurrentMinAttr }
func (RmsCurrentMin) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsCurrentMin) Name() string                  { return "RMS Current Min" }
func (RmsCurrentMin) Readable() bool                { return true }
func (RmsCurrentMin) Writable() bool                { return false }
func (RmsCurrentMin) Reportable() bool              { return false }
func (RmsCurrentMin) SceneIndex() int               { return -1 }
func (a *RmsCurrentMin) Value() *RmsCurrentMin      { return a }
func (a RmsCurrentMin) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsCurrentMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMin(*nt)
	return br, err
}

func (a RmsCurrentMin) String() string {
	return zcl.Amperes.Format(float64(a))
}

// RmsCurrentMax is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsCurrentMax zcl.Zu16

const RmsCurrentMaxAttr zcl.AttrID = 1290

func (RmsCurrentMax) ID() zcl.AttrID                { return RmsCurrentMaxAttr }
func (RmsCurrentMax) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsCurrentMax) Name() string                  { return "RMS Current Max" }
func (RmsCurrentMax) Readable() bool                { return true }
func (RmsCurrentMax) Writable() bool                { return false }
func (RmsCurrentMax) Reportable() bool              { return false }
func (RmsCurrentMax) SceneIndex() int               { return -1 }
func (a *RmsCurrentMax) Value() *RmsCurrentMax      { return a }
func (a RmsCurrentMax) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsCurrentMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMax(*nt)
	return br, err
}

func (a RmsCurrentMax) String() string {
	return zcl.Amperes.Format(float64(a))
}

// ActivePower is an autogenerated attribute in the ElectricalMeasurement cluster
// Represents the single phase or Phase A, current demand of active power delivered or received at the premises, in Watts (W). Positive values indicate power delivered to the premises where negative values indicate power received from the premises.
type ActivePower zcl.Zs16

const ActivePowerAttr zcl.AttrID = 1291

func (ActivePower) ID() zcl.AttrID                { return ActivePowerAttr }
func (ActivePower) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (ActivePower) Name() string                  { return "Active Power" }
func (ActivePower) Readable() bool                { return true }
func (ActivePower) Writable() bool                { return false }
func (ActivePower) Reportable() bool              { return false }
func (ActivePower) SceneIndex() int               { return -1 }
func (a *ActivePower) Value() *ActivePower        { return a }
func (a ActivePower) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *ActivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePower(*nt)
	return br, err
}

func (a ActivePower) String() string {
	return zcl.Watts.Format(float64(a))
}

// ActivePowerMin is an autogenerated attribute in the ElectricalMeasurement cluster
type ActivePowerMin zcl.Zs16

const ActivePowerMinAttr zcl.AttrID = 1292

func (ActivePowerMin) ID() zcl.AttrID                { return ActivePowerMinAttr }
func (ActivePowerMin) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (ActivePowerMin) Name() string                  { return "Active Power Min" }
func (ActivePowerMin) Readable() bool                { return true }
func (ActivePowerMin) Writable() bool                { return false }
func (ActivePowerMin) Reportable() bool              { return false }
func (ActivePowerMin) SceneIndex() int               { return -1 }
func (a *ActivePowerMin) Value() *ActivePowerMin     { return a }
func (a ActivePowerMin) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *ActivePowerMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMin(*nt)
	return br, err
}

func (a ActivePowerMin) String() string {
	return zcl.Watts.Format(float64(a))
}

// ActivePowerMax is an autogenerated attribute in the ElectricalMeasurement cluster
type ActivePowerMax zcl.Zs16

const ActivePowerMaxAttr zcl.AttrID = 1293

func (ActivePowerMax) ID() zcl.AttrID                { return ActivePowerMaxAttr }
func (ActivePowerMax) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (ActivePowerMax) Name() string                  { return "Active Power Max" }
func (ActivePowerMax) Readable() bool                { return true }
func (ActivePowerMax) Writable() bool                { return false }
func (ActivePowerMax) Reportable() bool              { return false }
func (ActivePowerMax) SceneIndex() int               { return -1 }
func (a *ActivePowerMax) Value() *ActivePowerMax     { return a }
func (a ActivePowerMax) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *ActivePowerMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMax(*nt)
	return br, err
}

func (a ActivePowerMax) String() string {
	return zcl.Watts.Format(float64(a))
}

// ReactivePower is an autogenerated attribute in the ElectricalMeasurement cluster
type ReactivePower zcl.Zs16

const ReactivePowerAttr zcl.AttrID = 1294

func (ReactivePower) ID() zcl.AttrID                { return ReactivePowerAttr }
func (ReactivePower) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (ReactivePower) Name() string                  { return "Reactive Power" }
func (ReactivePower) Readable() bool                { return true }
func (ReactivePower) Writable() bool                { return false }
func (ReactivePower) Reportable() bool              { return false }
func (ReactivePower) SceneIndex() int               { return -1 }
func (a *ReactivePower) Value() *ReactivePower      { return a }
func (a ReactivePower) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *ReactivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactivePower(*nt)
	return br, err
}

func (a ReactivePower) String() string {
	return zcl.VoltAmperesReactive.Format(float64(a))
}

// ApparentPower is an autogenerated attribute in the ElectricalMeasurement cluster
type ApparentPower zcl.Zu16

const ApparentPowerAttr zcl.AttrID = 1295

func (ApparentPower) ID() zcl.AttrID                { return ApparentPowerAttr }
func (ApparentPower) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (ApparentPower) Name() string                  { return "Apparent Power" }
func (ApparentPower) Readable() bool                { return true }
func (ApparentPower) Writable() bool                { return false }
func (ApparentPower) Reportable() bool              { return false }
func (ApparentPower) SceneIndex() int               { return -1 }
func (a *ApparentPower) Value() *ApparentPower      { return a }
func (a ApparentPower) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ApparentPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApparentPower(*nt)
	return br, err
}

func (a ApparentPower) String() string {
	return zcl.VoltAmperes.Format(float64(a))
}

// PowerFactor is an autogenerated attribute in the ElectricalMeasurement cluster
type PowerFactor zcl.Zs8

const PowerFactorAttr zcl.AttrID = 1296

func (PowerFactor) ID() zcl.AttrID                { return PowerFactorAttr }
func (PowerFactor) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (PowerFactor) Name() string                  { return "Power Factor" }
func (PowerFactor) Readable() bool                { return true }
func (PowerFactor) Writable() bool                { return false }
func (PowerFactor) Reportable() bool              { return false }
func (PowerFactor) SceneIndex() int               { return -1 }
func (a *PowerFactor) Value() *PowerFactor        { return a }
func (a PowerFactor) MarshalZcl() ([]byte, error) { return zcl.Zs8(a).MarshalZcl() }

func (a *PowerFactor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerFactor(*nt)
	return br, err
}

func (a PowerFactor) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(a))
}

// AverageRmsVoltageMeasurementPeriod is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsVoltageMeasurementPeriod zcl.Zu16

const AverageRmsVoltageMeasurementPeriodAttr zcl.AttrID = 1297

func (AverageRmsVoltageMeasurementPeriod) ID() zcl.AttrID {
	return AverageRmsVoltageMeasurementPeriodAttr
}
func (AverageRmsVoltageMeasurementPeriod) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (AverageRmsVoltageMeasurementPeriod) Name() string {
	return "Average RMS Voltage Measurement Period"
}
func (AverageRmsVoltageMeasurementPeriod) Readable() bool                                { return true }
func (AverageRmsVoltageMeasurementPeriod) Writable() bool                                { return true }
func (AverageRmsVoltageMeasurementPeriod) Reportable() bool                              { return false }
func (AverageRmsVoltageMeasurementPeriod) SceneIndex() int                               { return -1 }
func (a *AverageRmsVoltageMeasurementPeriod) Value() *AverageRmsVoltageMeasurementPeriod { return a }
func (a AverageRmsVoltageMeasurementPeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *AverageRmsVoltageMeasurementPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsVoltageMeasurementPeriod(*nt)
	return br, err
}

func (a AverageRmsVoltageMeasurementPeriod) String() string {
	return zcl.Seconds.Format(float64(a))
}

// AverageRmsOvervoltageCounter is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsOvervoltageCounter zcl.Zu16

const AverageRmsOvervoltageCounterAttr zcl.AttrID = 1298

func (AverageRmsOvervoltageCounter) ID() zcl.AttrID                          { return AverageRmsOvervoltageCounterAttr }
func (AverageRmsOvervoltageCounter) Cluster() zcl.ClusterID                  { return ElectricalMeasurementID }
func (AverageRmsOvervoltageCounter) Name() string                            { return "Average RMS Overvoltage Counter" }
func (AverageRmsOvervoltageCounter) Readable() bool                          { return true }
func (AverageRmsOvervoltageCounter) Writable() bool                          { return true }
func (AverageRmsOvervoltageCounter) Reportable() bool                        { return false }
func (AverageRmsOvervoltageCounter) SceneIndex() int                         { return -1 }
func (a *AverageRmsOvervoltageCounter) Value() *AverageRmsOvervoltageCounter { return a }
func (a AverageRmsOvervoltageCounter) MarshalZcl() ([]byte, error)           { return zcl.Zu16(a).MarshalZcl() }

func (a *AverageRmsOvervoltageCounter) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsOvervoltageCounter(*nt)
	return br, err
}

func (a AverageRmsOvervoltageCounter) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AverageRmsUndervoltageCounter is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsUndervoltageCounter zcl.Zu16

const AverageRmsUndervoltageCounterAttr zcl.AttrID = 1299

func (AverageRmsUndervoltageCounter) ID() zcl.AttrID                           { return AverageRmsUndervoltageCounterAttr }
func (AverageRmsUndervoltageCounter) Cluster() zcl.ClusterID                   { return ElectricalMeasurementID }
func (AverageRmsUndervoltageCounter) Name() string                             { return "Average RMS Undervoltage Counter" }
func (AverageRmsUndervoltageCounter) Readable() bool                           { return true }
func (AverageRmsUndervoltageCounter) Writable() bool                           { return true }
func (AverageRmsUndervoltageCounter) Reportable() bool                         { return false }
func (AverageRmsUndervoltageCounter) SceneIndex() int                          { return -1 }
func (a *AverageRmsUndervoltageCounter) Value() *AverageRmsUndervoltageCounter { return a }
func (a AverageRmsUndervoltageCounter) MarshalZcl() ([]byte, error)            { return zcl.Zu16(a).MarshalZcl() }

func (a *AverageRmsUndervoltageCounter) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsUndervoltageCounter(*nt)
	return br, err
}

func (a AverageRmsUndervoltageCounter) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// RmsExtremeOvervoltagePeriod is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsExtremeOvervoltagePeriod zcl.Zu16

const RmsExtremeOvervoltagePeriodAttr zcl.AttrID = 1300

func (RmsExtremeOvervoltagePeriod) ID() zcl.AttrID                         { return RmsExtremeOvervoltagePeriodAttr }
func (RmsExtremeOvervoltagePeriod) Cluster() zcl.ClusterID                 { return ElectricalMeasurementID }
func (RmsExtremeOvervoltagePeriod) Name() string                           { return "RMS Extreme Overvoltage Period" }
func (RmsExtremeOvervoltagePeriod) Readable() bool                         { return true }
func (RmsExtremeOvervoltagePeriod) Writable() bool                         { return true }
func (RmsExtremeOvervoltagePeriod) Reportable() bool                       { return false }
func (RmsExtremeOvervoltagePeriod) SceneIndex() int                        { return -1 }
func (a *RmsExtremeOvervoltagePeriod) Value() *RmsExtremeOvervoltagePeriod { return a }
func (a RmsExtremeOvervoltagePeriod) MarshalZcl() ([]byte, error)          { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsExtremeOvervoltagePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeOvervoltagePeriod(*nt)
	return br, err
}

func (a RmsExtremeOvervoltagePeriod) String() string {
	return zcl.Seconds.Format(float64(a))
}

// RmsExtremeUndervoltagePeriod is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsExtremeUndervoltagePeriod zcl.Zu16

const RmsExtremeUndervoltagePeriodAttr zcl.AttrID = 1301

func (RmsExtremeUndervoltagePeriod) ID() zcl.AttrID                          { return RmsExtremeUndervoltagePeriodAttr }
func (RmsExtremeUndervoltagePeriod) Cluster() zcl.ClusterID                  { return ElectricalMeasurementID }
func (RmsExtremeUndervoltagePeriod) Name() string                            { return "RMS Extreme Undervoltage Period" }
func (RmsExtremeUndervoltagePeriod) Readable() bool                          { return true }
func (RmsExtremeUndervoltagePeriod) Writable() bool                          { return true }
func (RmsExtremeUndervoltagePeriod) Reportable() bool                        { return false }
func (RmsExtremeUndervoltagePeriod) SceneIndex() int                         { return -1 }
func (a *RmsExtremeUndervoltagePeriod) Value() *RmsExtremeUndervoltagePeriod { return a }
func (a RmsExtremeUndervoltagePeriod) MarshalZcl() ([]byte, error)           { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsExtremeUndervoltagePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeUndervoltagePeriod(*nt)
	return br, err
}

func (a RmsExtremeUndervoltagePeriod) String() string {
	return zcl.Seconds.Format(float64(a))
}

// RmsVoltageSagPeriod is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageSagPeriod zcl.Zu16

const RmsVoltageSagPeriodAttr zcl.AttrID = 1302

func (RmsVoltageSagPeriod) ID() zcl.AttrID                 { return RmsVoltageSagPeriodAttr }
func (RmsVoltageSagPeriod) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (RmsVoltageSagPeriod) Name() string                   { return "RMS Voltage Sag Period" }
func (RmsVoltageSagPeriod) Readable() bool                 { return true }
func (RmsVoltageSagPeriod) Writable() bool                 { return true }
func (RmsVoltageSagPeriod) Reportable() bool               { return false }
func (RmsVoltageSagPeriod) SceneIndex() int                { return -1 }
func (a *RmsVoltageSagPeriod) Value() *RmsVoltageSagPeriod { return a }
func (a RmsVoltageSagPeriod) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageSagPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSagPeriod(*nt)
	return br, err
}

func (a RmsVoltageSagPeriod) String() string {
	return zcl.Seconds.Format(float64(a))
}

// RmsVoltageSwellPeriod is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageSwellPeriod zcl.Zu16

const RmsVoltageSwellPeriodAttr zcl.AttrID = 1303

func (RmsVoltageSwellPeriod) ID() zcl.AttrID                   { return RmsVoltageSwellPeriodAttr }
func (RmsVoltageSwellPeriod) Cluster() zcl.ClusterID           { return ElectricalMeasurementID }
func (RmsVoltageSwellPeriod) Name() string                     { return "RMS Voltage Swell Period" }
func (RmsVoltageSwellPeriod) Readable() bool                   { return true }
func (RmsVoltageSwellPeriod) Writable() bool                   { return true }
func (RmsVoltageSwellPeriod) Reportable() bool                 { return false }
func (RmsVoltageSwellPeriod) SceneIndex() int                  { return -1 }
func (a *RmsVoltageSwellPeriod) Value() *RmsVoltageSwellPeriod { return a }
func (a RmsVoltageSwellPeriod) MarshalZcl() ([]byte, error)    { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageSwellPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwellPeriod(*nt)
	return br, err
}

func (a RmsVoltageSwellPeriod) String() string {
	return zcl.Seconds.Format(float64(a))
}

// AcVoltageMultiplier is an autogenerated attribute in the ElectricalMeasurement cluster
type AcVoltageMultiplier zcl.Zu16

const AcVoltageMultiplierAttr zcl.AttrID = 1536

func (AcVoltageMultiplier) ID() zcl.AttrID                 { return AcVoltageMultiplierAttr }
func (AcVoltageMultiplier) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (AcVoltageMultiplier) Name() string                   { return "AC Voltage Multiplier" }
func (AcVoltageMultiplier) Readable() bool                 { return true }
func (AcVoltageMultiplier) Writable() bool                 { return false }
func (AcVoltageMultiplier) Reportable() bool               { return false }
func (AcVoltageMultiplier) SceneIndex() int                { return -1 }
func (a *AcVoltageMultiplier) Value() *AcVoltageMultiplier { return a }
func (a AcVoltageMultiplier) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *AcVoltageMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcVoltageMultiplier(*nt)
	return br, err
}

func (a AcVoltageMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AcVoltageDivisor is an autogenerated attribute in the ElectricalMeasurement cluster
type AcVoltageDivisor zcl.Zu16

const AcVoltageDivisorAttr zcl.AttrID = 1537

func (AcVoltageDivisor) ID() zcl.AttrID                { return AcVoltageDivisorAttr }
func (AcVoltageDivisor) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (AcVoltageDivisor) Name() string                  { return "AC Voltage Divisor" }
func (AcVoltageDivisor) Readable() bool                { return true }
func (AcVoltageDivisor) Writable() bool                { return false }
func (AcVoltageDivisor) Reportable() bool              { return false }
func (AcVoltageDivisor) SceneIndex() int               { return -1 }
func (a *AcVoltageDivisor) Value() *AcVoltageDivisor   { return a }
func (a AcVoltageDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AcVoltageDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcVoltageDivisor(*nt)
	return br, err
}

func (a AcVoltageDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AcCurrentMultiplier is an autogenerated attribute in the ElectricalMeasurement cluster
type AcCurrentMultiplier zcl.Zu16

const AcCurrentMultiplierAttr zcl.AttrID = 1538

func (AcCurrentMultiplier) ID() zcl.AttrID                 { return AcCurrentMultiplierAttr }
func (AcCurrentMultiplier) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (AcCurrentMultiplier) Name() string                   { return "AC Current Multiplier" }
func (AcCurrentMultiplier) Readable() bool                 { return true }
func (AcCurrentMultiplier) Writable() bool                 { return false }
func (AcCurrentMultiplier) Reportable() bool               { return false }
func (AcCurrentMultiplier) SceneIndex() int                { return -1 }
func (a *AcCurrentMultiplier) Value() *AcCurrentMultiplier { return a }
func (a AcCurrentMultiplier) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *AcCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcCurrentMultiplier(*nt)
	return br, err
}

func (a AcCurrentMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AcCurrentDivisor is an autogenerated attribute in the ElectricalMeasurement cluster
type AcCurrentDivisor zcl.Zu16

const AcCurrentDivisorAttr zcl.AttrID = 1539

func (AcCurrentDivisor) ID() zcl.AttrID                { return AcCurrentDivisorAttr }
func (AcCurrentDivisor) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (AcCurrentDivisor) Name() string                  { return "AC Current Divisor" }
func (AcCurrentDivisor) Readable() bool                { return true }
func (AcCurrentDivisor) Writable() bool                { return false }
func (AcCurrentDivisor) Reportable() bool              { return false }
func (AcCurrentDivisor) SceneIndex() int               { return -1 }
func (a *AcCurrentDivisor) Value() *AcCurrentDivisor   { return a }
func (a AcCurrentDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AcCurrentDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcCurrentDivisor(*nt)
	return br, err
}

func (a AcCurrentDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AcPowerMultiplier is an autogenerated attribute in the ElectricalMeasurement cluster
type AcPowerMultiplier zcl.Zu16

const AcPowerMultiplierAttr zcl.AttrID = 1540

func (AcPowerMultiplier) ID() zcl.AttrID                { return AcPowerMultiplierAttr }
func (AcPowerMultiplier) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (AcPowerMultiplier) Name() string                  { return "AC Power Multiplier" }
func (AcPowerMultiplier) Readable() bool                { return true }
func (AcPowerMultiplier) Writable() bool                { return false }
func (AcPowerMultiplier) Reportable() bool              { return false }
func (AcPowerMultiplier) SceneIndex() int               { return -1 }
func (a *AcPowerMultiplier) Value() *AcPowerMultiplier  { return a }
func (a AcPowerMultiplier) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AcPowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcPowerMultiplier(*nt)
	return br, err
}

func (a AcPowerMultiplier) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AcPowerDivisor is an autogenerated attribute in the ElectricalMeasurement cluster
type AcPowerDivisor zcl.Zu16

const AcPowerDivisorAttr zcl.AttrID = 1541

func (AcPowerDivisor) ID() zcl.AttrID                { return AcPowerDivisorAttr }
func (AcPowerDivisor) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (AcPowerDivisor) Name() string                  { return "AC Power Divisor" }
func (AcPowerDivisor) Readable() bool                { return true }
func (AcPowerDivisor) Writable() bool                { return false }
func (AcPowerDivisor) Reportable() bool              { return false }
func (AcPowerDivisor) SceneIndex() int               { return -1 }
func (a *AcPowerDivisor) Value() *AcPowerDivisor     { return a }
func (a AcPowerDivisor) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AcPowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcPowerDivisor(*nt)
	return br, err
}

func (a AcPowerDivisor) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// DcOverloadAlarmsMask is an autogenerated attribute in the ElectricalMeasurement cluster
type DcOverloadAlarmsMask zcl.Zbmp8

const DcOverloadAlarmsMaskAttr zcl.AttrID = 1792

func (DcOverloadAlarmsMask) ID() zcl.AttrID                  { return DcOverloadAlarmsMaskAttr }
func (DcOverloadAlarmsMask) Cluster() zcl.ClusterID          { return ElectricalMeasurementID }
func (DcOverloadAlarmsMask) Name() string                    { return "DC Overload Alarms Mask" }
func (DcOverloadAlarmsMask) Readable() bool                  { return true }
func (DcOverloadAlarmsMask) Writable() bool                  { return true }
func (DcOverloadAlarmsMask) Reportable() bool                { return false }
func (DcOverloadAlarmsMask) SceneIndex() int                 { return -1 }
func (a *DcOverloadAlarmsMask) Value() *DcOverloadAlarmsMask { return a }
func (a DcOverloadAlarmsMask) MarshalZcl() ([]byte, error)   { return zcl.Zbmp8(a).MarshalZcl() }

func (a *DcOverloadAlarmsMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DcOverloadAlarmsMask(*nt)
	return br, err
}

func (a DcOverloadAlarmsMask) String() string {
	var bstr []string
	if a.IsVoltageOverload() {
		bstr = append(bstr, "Voltage Overload")
	}
	if a.IsCurrentOverload() {
		bstr = append(bstr, "Current Overload")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a DcOverloadAlarmsMask) IsVoltageOverload() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *DcOverloadAlarmsMask) SetVoltageOverload(b bool) {
	*a = DcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a DcOverloadAlarmsMask) IsCurrentOverload() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *DcOverloadAlarmsMask) SetCurrentOverload(b bool) {
	*a = DcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 1, b))
}

// DcVoltageOverload is an autogenerated attribute in the ElectricalMeasurement cluster
type DcVoltageOverload zcl.Zs16

const DcVoltageOverloadAttr zcl.AttrID = 1793

func (DcVoltageOverload) ID() zcl.AttrID                { return DcVoltageOverloadAttr }
func (DcVoltageOverload) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcVoltageOverload) Name() string                  { return "DC Voltage Overload" }
func (DcVoltageOverload) Readable() bool                { return true }
func (DcVoltageOverload) Writable() bool                { return false }
func (DcVoltageOverload) Reportable() bool              { return false }
func (DcVoltageOverload) SceneIndex() int               { return -1 }
func (a *DcVoltageOverload) Value() *DcVoltageOverload  { return a }
func (a DcVoltageOverload) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcVoltageOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageOverload(*nt)
	return br, err
}

func (a DcVoltageOverload) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// DcCurrentOverload is an autogenerated attribute in the ElectricalMeasurement cluster
type DcCurrentOverload zcl.Zs16

const DcCurrentOverloadAttr zcl.AttrID = 1794

func (DcCurrentOverload) ID() zcl.AttrID                { return DcCurrentOverloadAttr }
func (DcCurrentOverload) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (DcCurrentOverload) Name() string                  { return "DC Current Overload" }
func (DcCurrentOverload) Readable() bool                { return true }
func (DcCurrentOverload) Writable() bool                { return false }
func (DcCurrentOverload) Reportable() bool              { return false }
func (DcCurrentOverload) SceneIndex() int               { return -1 }
func (a *DcCurrentOverload) Value() *DcCurrentOverload  { return a }
func (a DcCurrentOverload) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *DcCurrentOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentOverload(*nt)
	return br, err
}

func (a DcCurrentOverload) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// AcOverloadAlarmsMask is an autogenerated attribute in the ElectricalMeasurement cluster
type AcOverloadAlarmsMask zcl.Zbmp16

const AcOverloadAlarmsMaskAttr zcl.AttrID = 2048

func (AcOverloadAlarmsMask) ID() zcl.AttrID                  { return AcOverloadAlarmsMaskAttr }
func (AcOverloadAlarmsMask) Cluster() zcl.ClusterID          { return ElectricalMeasurementID }
func (AcOverloadAlarmsMask) Name() string                    { return "AC Overload Alarms Mask" }
func (AcOverloadAlarmsMask) Readable() bool                  { return true }
func (AcOverloadAlarmsMask) Writable() bool                  { return true }
func (AcOverloadAlarmsMask) Reportable() bool                { return false }
func (AcOverloadAlarmsMask) SceneIndex() int                 { return -1 }
func (a *AcOverloadAlarmsMask) Value() *AcOverloadAlarmsMask { return a }
func (a AcOverloadAlarmsMask) MarshalZcl() ([]byte, error)   { return zcl.Zbmp16(a).MarshalZcl() }

func (a *AcOverloadAlarmsMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcOverloadAlarmsMask(*nt)
	return br, err
}

func (a AcOverloadAlarmsMask) String() string {
	var bstr []string
	if a.IsVoltageOverload() {
		bstr = append(bstr, "Voltage Overload")
	}
	if a.IsCurrentOverload() {
		bstr = append(bstr, "Current Overload")
	}
	if a.IsActivePowerOverload() {
		bstr = append(bstr, "Active Power Overload")
	}
	if a.IsReactivePowerOverload() {
		bstr = append(bstr, "Reactive Power Overload")
	}
	if a.IsAverageRmsOvervoltage() {
		bstr = append(bstr, "Average RMS Overvoltage")
	}
	if a.IsAverageRmsUndervoltage() {
		bstr = append(bstr, "Average RMS Undervoltage")
	}
	if a.IsRmsExtremeOvervoltage() {
		bstr = append(bstr, "RMS Extreme Overvoltage")
	}
	if a.IsRmsExtremeUndervoltage() {
		bstr = append(bstr, "RMS Extreme Undervoltage")
	}
	if a.IsRmsVoltageSag() {
		bstr = append(bstr, "RMS Voltage Sag")
	}
	if a.IsRmsVoltageSwell() {
		bstr = append(bstr, "RMS Voltage Swell")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a AcOverloadAlarmsMask) IsVoltageOverload() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *AcOverloadAlarmsMask) SetVoltageOverload(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a AcOverloadAlarmsMask) IsCurrentOverload() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *AcOverloadAlarmsMask) SetCurrentOverload(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a AcOverloadAlarmsMask) IsActivePowerOverload() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *AcOverloadAlarmsMask) SetActivePowerOverload(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a AcOverloadAlarmsMask) IsReactivePowerOverload() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *AcOverloadAlarmsMask) SetReactivePowerOverload(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a AcOverloadAlarmsMask) IsAverageRmsOvervoltage() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *AcOverloadAlarmsMask) SetAverageRmsOvervoltage(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a AcOverloadAlarmsMask) IsAverageRmsUndervoltage() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *AcOverloadAlarmsMask) SetAverageRmsUndervoltage(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a AcOverloadAlarmsMask) IsRmsExtremeOvervoltage() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *AcOverloadAlarmsMask) SetRmsExtremeOvervoltage(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 6, b))
}

func (a AcOverloadAlarmsMask) IsRmsExtremeUndervoltage() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *AcOverloadAlarmsMask) SetRmsExtremeUndervoltage(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 7, b))
}

func (a AcOverloadAlarmsMask) IsRmsVoltageSag() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *AcOverloadAlarmsMask) SetRmsVoltageSag(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 8, b))
}

func (a AcOverloadAlarmsMask) IsRmsVoltageSwell() bool {
	return zcl.BitmapTest([]byte(a), 9)
}
func (a *AcOverloadAlarmsMask) SetRmsVoltageSwell(b bool) {
	*a = AcOverloadAlarmsMask(zcl.BitmapSet([]byte(*a), 9, b))
}

// AcVoltageOverload is an autogenerated attribute in the ElectricalMeasurement cluster
type AcVoltageOverload zcl.Zs16

const AcVoltageOverloadAttr zcl.AttrID = 2049

func (AcVoltageOverload) ID() zcl.AttrID                { return AcVoltageOverloadAttr }
func (AcVoltageOverload) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (AcVoltageOverload) Name() string                  { return "AC Voltage Overload" }
func (AcVoltageOverload) Readable() bool                { return true }
func (AcVoltageOverload) Writable() bool                { return false }
func (AcVoltageOverload) Reportable() bool              { return false }
func (AcVoltageOverload) SceneIndex() int               { return -1 }
func (a *AcVoltageOverload) Value() *AcVoltageOverload  { return a }
func (a AcVoltageOverload) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *AcVoltageOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcVoltageOverload(*nt)
	return br, err
}

func (a AcVoltageOverload) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// AcCurrentOverload is an autogenerated attribute in the ElectricalMeasurement cluster
type AcCurrentOverload zcl.Zs16

const AcCurrentOverloadAttr zcl.AttrID = 2050

func (AcCurrentOverload) ID() zcl.AttrID                { return AcCurrentOverloadAttr }
func (AcCurrentOverload) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (AcCurrentOverload) Name() string                  { return "AC Current Overload" }
func (AcCurrentOverload) Readable() bool                { return true }
func (AcCurrentOverload) Writable() bool                { return false }
func (AcCurrentOverload) Reportable() bool              { return false }
func (AcCurrentOverload) SceneIndex() int               { return -1 }
func (a *AcCurrentOverload) Value() *AcCurrentOverload  { return a }
func (a AcCurrentOverload) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *AcCurrentOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcCurrentOverload(*nt)
	return br, err
}

func (a AcCurrentOverload) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// AcActivePowerOverload is an autogenerated attribute in the ElectricalMeasurement cluster
type AcActivePowerOverload zcl.Zs16

const AcActivePowerOverloadAttr zcl.AttrID = 2051

func (AcActivePowerOverload) ID() zcl.AttrID                   { return AcActivePowerOverloadAttr }
func (AcActivePowerOverload) Cluster() zcl.ClusterID           { return ElectricalMeasurementID }
func (AcActivePowerOverload) Name() string                     { return "AC Active Power Overload" }
func (AcActivePowerOverload) Readable() bool                   { return true }
func (AcActivePowerOverload) Writable() bool                   { return false }
func (AcActivePowerOverload) Reportable() bool                 { return false }
func (AcActivePowerOverload) SceneIndex() int                  { return -1 }
func (a *AcActivePowerOverload) Value() *AcActivePowerOverload { return a }
func (a AcActivePowerOverload) MarshalZcl() ([]byte, error)    { return zcl.Zs16(a).MarshalZcl() }

func (a *AcActivePowerOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcActivePowerOverload(*nt)
	return br, err
}

func (a AcActivePowerOverload) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// AcReactivePowerOverload is an autogenerated attribute in the ElectricalMeasurement cluster
type AcReactivePowerOverload zcl.Zs16

const AcReactivePowerOverloadAttr zcl.AttrID = 2052

func (AcReactivePowerOverload) ID() zcl.AttrID                     { return AcReactivePowerOverloadAttr }
func (AcReactivePowerOverload) Cluster() zcl.ClusterID             { return ElectricalMeasurementID }
func (AcReactivePowerOverload) Name() string                       { return "AC Reactive Power Overload" }
func (AcReactivePowerOverload) Readable() bool                     { return true }
func (AcReactivePowerOverload) Writable() bool                     { return false }
func (AcReactivePowerOverload) Reportable() bool                   { return false }
func (AcReactivePowerOverload) SceneIndex() int                    { return -1 }
func (a *AcReactivePowerOverload) Value() *AcReactivePowerOverload { return a }
func (a AcReactivePowerOverload) MarshalZcl() ([]byte, error)      { return zcl.Zs16(a).MarshalZcl() }

func (a *AcReactivePowerOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcReactivePowerOverload(*nt)
	return br, err
}

func (a AcReactivePowerOverload) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// AverageRmsOvervoltage is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsOvervoltage zcl.Zs16

const AverageRmsOvervoltageAttr zcl.AttrID = 2053

func (AverageRmsOvervoltage) ID() zcl.AttrID                   { return AverageRmsOvervoltageAttr }
func (AverageRmsOvervoltage) Cluster() zcl.ClusterID           { return ElectricalMeasurementID }
func (AverageRmsOvervoltage) Name() string                     { return "Average RMS Overvoltage" }
func (AverageRmsOvervoltage) Readable() bool                   { return true }
func (AverageRmsOvervoltage) Writable() bool                   { return false }
func (AverageRmsOvervoltage) Reportable() bool                 { return false }
func (AverageRmsOvervoltage) SceneIndex() int                  { return -1 }
func (a *AverageRmsOvervoltage) Value() *AverageRmsOvervoltage { return a }
func (a AverageRmsOvervoltage) MarshalZcl() ([]byte, error)    { return zcl.Zs16(a).MarshalZcl() }

func (a *AverageRmsOvervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsOvervoltage(*nt)
	return br, err
}

func (a AverageRmsOvervoltage) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// AverageRmsUndervoltage is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsUndervoltage zcl.Zs16

const AverageRmsUndervoltageAttr zcl.AttrID = 2054

func (AverageRmsUndervoltage) ID() zcl.AttrID                    { return AverageRmsUndervoltageAttr }
func (AverageRmsUndervoltage) Cluster() zcl.ClusterID            { return ElectricalMeasurementID }
func (AverageRmsUndervoltage) Name() string                      { return "Average RMS Undervoltage" }
func (AverageRmsUndervoltage) Readable() bool                    { return true }
func (AverageRmsUndervoltage) Writable() bool                    { return false }
func (AverageRmsUndervoltage) Reportable() bool                  { return false }
func (AverageRmsUndervoltage) SceneIndex() int                   { return -1 }
func (a *AverageRmsUndervoltage) Value() *AverageRmsUndervoltage { return a }
func (a AverageRmsUndervoltage) MarshalZcl() ([]byte, error)     { return zcl.Zs16(a).MarshalZcl() }

func (a *AverageRmsUndervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsUndervoltage(*nt)
	return br, err
}

func (a AverageRmsUndervoltage) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// RmsExtremeOvervoltage is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsExtremeOvervoltage zcl.Zs16

const RmsExtremeOvervoltageAttr zcl.AttrID = 2055

func (RmsExtremeOvervoltage) ID() zcl.AttrID                   { return RmsExtremeOvervoltageAttr }
func (RmsExtremeOvervoltage) Cluster() zcl.ClusterID           { return ElectricalMeasurementID }
func (RmsExtremeOvervoltage) Name() string                     { return "RMS Extreme Overvoltage" }
func (RmsExtremeOvervoltage) Readable() bool                   { return true }
func (RmsExtremeOvervoltage) Writable() bool                   { return false }
func (RmsExtremeOvervoltage) Reportable() bool                 { return false }
func (RmsExtremeOvervoltage) SceneIndex() int                  { return -1 }
func (a *RmsExtremeOvervoltage) Value() *RmsExtremeOvervoltage { return a }
func (a RmsExtremeOvervoltage) MarshalZcl() ([]byte, error)    { return zcl.Zs16(a).MarshalZcl() }

func (a *RmsExtremeOvervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeOvervoltage(*nt)
	return br, err
}

func (a RmsExtremeOvervoltage) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// RmsExtremeUndervoltage is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsExtremeUndervoltage zcl.Zs16

const RmsExtremeUndervoltageAttr zcl.AttrID = 2056

func (RmsExtremeUndervoltage) ID() zcl.AttrID                    { return RmsExtremeUndervoltageAttr }
func (RmsExtremeUndervoltage) Cluster() zcl.ClusterID            { return ElectricalMeasurementID }
func (RmsExtremeUndervoltage) Name() string                      { return "RMS Extreme Undervoltage" }
func (RmsExtremeUndervoltage) Readable() bool                    { return true }
func (RmsExtremeUndervoltage) Writable() bool                    { return false }
func (RmsExtremeUndervoltage) Reportable() bool                  { return false }
func (RmsExtremeUndervoltage) SceneIndex() int                   { return -1 }
func (a *RmsExtremeUndervoltage) Value() *RmsExtremeUndervoltage { return a }
func (a RmsExtremeUndervoltage) MarshalZcl() ([]byte, error)     { return zcl.Zs16(a).MarshalZcl() }

func (a *RmsExtremeUndervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeUndervoltage(*nt)
	return br, err
}

func (a RmsExtremeUndervoltage) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// RmsVoltageSag is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageSag zcl.Zs16

const RmsVoltageSagAttr zcl.AttrID = 2057

func (RmsVoltageSag) ID() zcl.AttrID                { return RmsVoltageSagAttr }
func (RmsVoltageSag) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsVoltageSag) Name() string                  { return "RMS Voltage Sag" }
func (RmsVoltageSag) Readable() bool                { return true }
func (RmsVoltageSag) Writable() bool                { return false }
func (RmsVoltageSag) Reportable() bool              { return false }
func (RmsVoltageSag) SceneIndex() int               { return -1 }
func (a *RmsVoltageSag) Value() *RmsVoltageSag      { return a }
func (a RmsVoltageSag) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *RmsVoltageSag) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSag(*nt)
	return br, err
}

func (a RmsVoltageSag) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// RmsVoltageSwell is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageSwell zcl.Zs16

const RmsVoltageSwellAttr zcl.AttrID = 2058

func (RmsVoltageSwell) ID() zcl.AttrID                { return RmsVoltageSwellAttr }
func (RmsVoltageSwell) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsVoltageSwell) Name() string                  { return "RMS Voltage Swell" }
func (RmsVoltageSwell) Readable() bool                { return true }
func (RmsVoltageSwell) Writable() bool                { return false }
func (RmsVoltageSwell) Reportable() bool              { return false }
func (RmsVoltageSwell) SceneIndex() int               { return -1 }
func (a *RmsVoltageSwell) Value() *RmsVoltageSwell    { return a }
func (a RmsVoltageSwell) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *RmsVoltageSwell) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwell(*nt)
	return br, err
}

func (a RmsVoltageSwell) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// LineCurrentPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type LineCurrentPhaseB zcl.Zu16

const LineCurrentPhaseBAttr zcl.AttrID = 2305

func (LineCurrentPhaseB) ID() zcl.AttrID                { return LineCurrentPhaseBAttr }
func (LineCurrentPhaseB) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (LineCurrentPhaseB) Name() string                  { return "Line Current Phase B" }
func (LineCurrentPhaseB) Readable() bool                { return true }
func (LineCurrentPhaseB) Writable() bool                { return false }
func (LineCurrentPhaseB) Reportable() bool              { return false }
func (LineCurrentPhaseB) SceneIndex() int               { return -1 }
func (a *LineCurrentPhaseB) Value() *LineCurrentPhaseB  { return a }
func (a LineCurrentPhaseB) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LineCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LineCurrentPhaseB(*nt)
	return br, err
}

func (a LineCurrentPhaseB) String() string {
	return zcl.Amperes.Format(float64(a))
}

// ActiveCurrentPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type ActiveCurrentPhaseB zcl.Zs16

const ActiveCurrentPhaseBAttr zcl.AttrID = 2306

func (ActiveCurrentPhaseB) ID() zcl.AttrID                 { return ActiveCurrentPhaseBAttr }
func (ActiveCurrentPhaseB) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (ActiveCurrentPhaseB) Name() string                   { return "Active Current Phase B" }
func (ActiveCurrentPhaseB) Readable() bool                 { return true }
func (ActiveCurrentPhaseB) Writable() bool                 { return false }
func (ActiveCurrentPhaseB) Reportable() bool               { return false }
func (ActiveCurrentPhaseB) SceneIndex() int                { return -1 }
func (a *ActiveCurrentPhaseB) Value() *ActiveCurrentPhaseB { return a }
func (a ActiveCurrentPhaseB) MarshalZcl() ([]byte, error)  { return zcl.Zs16(a).MarshalZcl() }

func (a *ActiveCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveCurrentPhaseB(*nt)
	return br, err
}

func (a ActiveCurrentPhaseB) String() string {
	return zcl.Amperes.Format(float64(a))
}

// ReactiveCurrentPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type ReactiveCurrentPhaseB zcl.Zs16

const ReactiveCurrentPhaseBAttr zcl.AttrID = 2307

func (ReactiveCurrentPhaseB) ID() zcl.AttrID                   { return ReactiveCurrentPhaseBAttr }
func (ReactiveCurrentPhaseB) Cluster() zcl.ClusterID           { return ElectricalMeasurementID }
func (ReactiveCurrentPhaseB) Name() string                     { return "Reactive Current Phase B" }
func (ReactiveCurrentPhaseB) Readable() bool                   { return true }
func (ReactiveCurrentPhaseB) Writable() bool                   { return false }
func (ReactiveCurrentPhaseB) Reportable() bool                 { return false }
func (ReactiveCurrentPhaseB) SceneIndex() int                  { return -1 }
func (a *ReactiveCurrentPhaseB) Value() *ReactiveCurrentPhaseB { return a }
func (a ReactiveCurrentPhaseB) MarshalZcl() ([]byte, error)    { return zcl.Zs16(a).MarshalZcl() }

func (a *ReactiveCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactiveCurrentPhaseB(*nt)
	return br, err
}

func (a ReactiveCurrentPhaseB) String() string {
	return zcl.Amperes.Format(float64(a))
}

// RmsVoltagePhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltagePhaseB zcl.Zu16

const RmsVoltagePhaseBAttr zcl.AttrID = 2309

func (RmsVoltagePhaseB) ID() zcl.AttrID                { return RmsVoltagePhaseBAttr }
func (RmsVoltagePhaseB) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsVoltagePhaseB) Name() string                  { return "RMS Voltage Phase B" }
func (RmsVoltagePhaseB) Readable() bool                { return true }
func (RmsVoltagePhaseB) Writable() bool                { return false }
func (RmsVoltagePhaseB) Reportable() bool              { return false }
func (RmsVoltagePhaseB) SceneIndex() int               { return -1 }
func (a *RmsVoltagePhaseB) Value() *RmsVoltagePhaseB   { return a }
func (a RmsVoltagePhaseB) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltagePhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltagePhaseB(*nt)
	return br, err
}

func (a RmsVoltagePhaseB) String() string {
	return zcl.Volts.Format(float64(a))
}

// RmsVoltageMinPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageMinPhaseB zcl.Zu16

const RmsVoltageMinPhaseBAttr zcl.AttrID = 2310

func (RmsVoltageMinPhaseB) ID() zcl.AttrID                 { return RmsVoltageMinPhaseBAttr }
func (RmsVoltageMinPhaseB) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (RmsVoltageMinPhaseB) Name() string                   { return "RMS Voltage Min Phase B" }
func (RmsVoltageMinPhaseB) Readable() bool                 { return true }
func (RmsVoltageMinPhaseB) Writable() bool                 { return false }
func (RmsVoltageMinPhaseB) Reportable() bool               { return false }
func (RmsVoltageMinPhaseB) SceneIndex() int                { return -1 }
func (a *RmsVoltageMinPhaseB) Value() *RmsVoltageMinPhaseB { return a }
func (a RmsVoltageMinPhaseB) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageMinPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMinPhaseB(*nt)
	return br, err
}

func (a RmsVoltageMinPhaseB) String() string {
	return zcl.Volts.Format(float64(a))
}

// RmsVoltageMaxPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageMaxPhaseB zcl.Zu16

const RmsVoltageMaxPhaseBAttr zcl.AttrID = 2311

func (RmsVoltageMaxPhaseB) ID() zcl.AttrID                 { return RmsVoltageMaxPhaseBAttr }
func (RmsVoltageMaxPhaseB) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (RmsVoltageMaxPhaseB) Name() string                   { return "RMS Voltage Max Phase B" }
func (RmsVoltageMaxPhaseB) Readable() bool                 { return true }
func (RmsVoltageMaxPhaseB) Writable() bool                 { return false }
func (RmsVoltageMaxPhaseB) Reportable() bool               { return false }
func (RmsVoltageMaxPhaseB) SceneIndex() int                { return -1 }
func (a *RmsVoltageMaxPhaseB) Value() *RmsVoltageMaxPhaseB { return a }
func (a RmsVoltageMaxPhaseB) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageMaxPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMaxPhaseB(*nt)
	return br, err
}

func (a RmsVoltageMaxPhaseB) String() string {
	return zcl.Volts.Format(float64(a))
}

// RmsCurrentPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsCurrentPhaseB zcl.Zu16

const RmsCurrentPhaseBAttr zcl.AttrID = 2312

func (RmsCurrentPhaseB) ID() zcl.AttrID                { return RmsCurrentPhaseBAttr }
func (RmsCurrentPhaseB) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsCurrentPhaseB) Name() string                  { return "RMS Current Phase B" }
func (RmsCurrentPhaseB) Readable() bool                { return true }
func (RmsCurrentPhaseB) Writable() bool                { return false }
func (RmsCurrentPhaseB) Reportable() bool              { return false }
func (RmsCurrentPhaseB) SceneIndex() int               { return -1 }
func (a *RmsCurrentPhaseB) Value() *RmsCurrentPhaseB   { return a }
func (a RmsCurrentPhaseB) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentPhaseB(*nt)
	return br, err
}

func (a RmsCurrentPhaseB) String() string {
	return zcl.Amperes.Format(float64(a))
}

// RmsCurrentMinPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsCurrentMinPhaseB zcl.Zu16

const RmsCurrentMinPhaseBAttr zcl.AttrID = 2313

func (RmsCurrentMinPhaseB) ID() zcl.AttrID                 { return RmsCurrentMinPhaseBAttr }
func (RmsCurrentMinPhaseB) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (RmsCurrentMinPhaseB) Name() string                   { return "RMS Current Min Phase B" }
func (RmsCurrentMinPhaseB) Readable() bool                 { return true }
func (RmsCurrentMinPhaseB) Writable() bool                 { return false }
func (RmsCurrentMinPhaseB) Reportable() bool               { return false }
func (RmsCurrentMinPhaseB) SceneIndex() int                { return -1 }
func (a *RmsCurrentMinPhaseB) Value() *RmsCurrentMinPhaseB { return a }
func (a RmsCurrentMinPhaseB) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsCurrentMinPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMinPhaseB(*nt)
	return br, err
}

func (a RmsCurrentMinPhaseB) String() string {
	return zcl.Amperes.Format(float64(a))
}

// RmsCurrentMaxPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsCurrentMaxPhaseB zcl.Zu16

const RmsCurrentMaxPhaseBAttr zcl.AttrID = 2314

func (RmsCurrentMaxPhaseB) ID() zcl.AttrID                 { return RmsCurrentMaxPhaseBAttr }
func (RmsCurrentMaxPhaseB) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (RmsCurrentMaxPhaseB) Name() string                   { return "RMS Current Max Phase B" }
func (RmsCurrentMaxPhaseB) Readable() bool                 { return true }
func (RmsCurrentMaxPhaseB) Writable() bool                 { return false }
func (RmsCurrentMaxPhaseB) Reportable() bool               { return false }
func (RmsCurrentMaxPhaseB) SceneIndex() int                { return -1 }
func (a *RmsCurrentMaxPhaseB) Value() *RmsCurrentMaxPhaseB { return a }
func (a RmsCurrentMaxPhaseB) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsCurrentMaxPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMaxPhaseB(*nt)
	return br, err
}

func (a RmsCurrentMaxPhaseB) String() string {
	return zcl.Amperes.Format(float64(a))
}

// ActivePowerPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
// Represents the Phase B, current demand of active power delivered or received at the premises, in Watts (W). Positive values indicate power delivered to the premises where negative values indicate power received from the premises.
type ActivePowerPhaseB zcl.Zs16

const ActivePowerPhaseBAttr zcl.AttrID = 2315

func (ActivePowerPhaseB) ID() zcl.AttrID                { return ActivePowerPhaseBAttr }
func (ActivePowerPhaseB) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (ActivePowerPhaseB) Name() string                  { return "Active Power Phase B" }
func (ActivePowerPhaseB) Readable() bool                { return true }
func (ActivePowerPhaseB) Writable() bool                { return false }
func (ActivePowerPhaseB) Reportable() bool              { return false }
func (ActivePowerPhaseB) SceneIndex() int               { return -1 }
func (a *ActivePowerPhaseB) Value() *ActivePowerPhaseB  { return a }
func (a ActivePowerPhaseB) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *ActivePowerPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerPhaseB(*nt)
	return br, err
}

func (a ActivePowerPhaseB) String() string {
	return zcl.Watts.Format(float64(a))
}

// ActivePowerMinPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type ActivePowerMinPhaseB zcl.Zs16

const ActivePowerMinPhaseBAttr zcl.AttrID = 2316

func (ActivePowerMinPhaseB) ID() zcl.AttrID                  { return ActivePowerMinPhaseBAttr }
func (ActivePowerMinPhaseB) Cluster() zcl.ClusterID          { return ElectricalMeasurementID }
func (ActivePowerMinPhaseB) Name() string                    { return "Active Power Min Phase B" }
func (ActivePowerMinPhaseB) Readable() bool                  { return true }
func (ActivePowerMinPhaseB) Writable() bool                  { return false }
func (ActivePowerMinPhaseB) Reportable() bool                { return false }
func (ActivePowerMinPhaseB) SceneIndex() int                 { return -1 }
func (a *ActivePowerMinPhaseB) Value() *ActivePowerMinPhaseB { return a }
func (a ActivePowerMinPhaseB) MarshalZcl() ([]byte, error)   { return zcl.Zs16(a).MarshalZcl() }

func (a *ActivePowerMinPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMinPhaseB(*nt)
	return br, err
}

func (a ActivePowerMinPhaseB) String() string {
	return zcl.Watts.Format(float64(a))
}

// ActivePowerMaxPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type ActivePowerMaxPhaseB zcl.Zs16

const ActivePowerMaxPhaseBAttr zcl.AttrID = 2317

func (ActivePowerMaxPhaseB) ID() zcl.AttrID                  { return ActivePowerMaxPhaseBAttr }
func (ActivePowerMaxPhaseB) Cluster() zcl.ClusterID          { return ElectricalMeasurementID }
func (ActivePowerMaxPhaseB) Name() string                    { return "Active Power Max Phase B" }
func (ActivePowerMaxPhaseB) Readable() bool                  { return true }
func (ActivePowerMaxPhaseB) Writable() bool                  { return false }
func (ActivePowerMaxPhaseB) Reportable() bool                { return false }
func (ActivePowerMaxPhaseB) SceneIndex() int                 { return -1 }
func (a *ActivePowerMaxPhaseB) Value() *ActivePowerMaxPhaseB { return a }
func (a ActivePowerMaxPhaseB) MarshalZcl() ([]byte, error)   { return zcl.Zs16(a).MarshalZcl() }

func (a *ActivePowerMaxPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMaxPhaseB(*nt)
	return br, err
}

func (a ActivePowerMaxPhaseB) String() string {
	return zcl.Watts.Format(float64(a))
}

// ReactivePowerPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type ReactivePowerPhaseB zcl.Zs16

const ReactivePowerPhaseBAttr zcl.AttrID = 2318

func (ReactivePowerPhaseB) ID() zcl.AttrID                 { return ReactivePowerPhaseBAttr }
func (ReactivePowerPhaseB) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (ReactivePowerPhaseB) Name() string                   { return "Reactive Power Phase B" }
func (ReactivePowerPhaseB) Readable() bool                 { return true }
func (ReactivePowerPhaseB) Writable() bool                 { return false }
func (ReactivePowerPhaseB) Reportable() bool               { return false }
func (ReactivePowerPhaseB) SceneIndex() int                { return -1 }
func (a *ReactivePowerPhaseB) Value() *ReactivePowerPhaseB { return a }
func (a ReactivePowerPhaseB) MarshalZcl() ([]byte, error)  { return zcl.Zs16(a).MarshalZcl() }

func (a *ReactivePowerPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactivePowerPhaseB(*nt)
	return br, err
}

func (a ReactivePowerPhaseB) String() string {
	return zcl.VoltAmperesReactive.Format(float64(a))
}

// ApparentPowerPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type ApparentPowerPhaseB zcl.Zu16

const ApparentPowerPhaseBAttr zcl.AttrID = 2319

func (ApparentPowerPhaseB) ID() zcl.AttrID                 { return ApparentPowerPhaseBAttr }
func (ApparentPowerPhaseB) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (ApparentPowerPhaseB) Name() string                   { return "Apparent Power Phase B" }
func (ApparentPowerPhaseB) Readable() bool                 { return true }
func (ApparentPowerPhaseB) Writable() bool                 { return false }
func (ApparentPowerPhaseB) Reportable() bool               { return false }
func (ApparentPowerPhaseB) SceneIndex() int                { return -1 }
func (a *ApparentPowerPhaseB) Value() *ApparentPowerPhaseB { return a }
func (a ApparentPowerPhaseB) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *ApparentPowerPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApparentPowerPhaseB(*nt)
	return br, err
}

func (a ApparentPowerPhaseB) String() string {
	return zcl.VoltAmperes.Format(float64(a))
}

// PowerFactorPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type PowerFactorPhaseB zcl.Zs8

const PowerFactorPhaseBAttr zcl.AttrID = 2320

func (PowerFactorPhaseB) ID() zcl.AttrID                { return PowerFactorPhaseBAttr }
func (PowerFactorPhaseB) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (PowerFactorPhaseB) Name() string                  { return "Power Factor Phase B" }
func (PowerFactorPhaseB) Readable() bool                { return true }
func (PowerFactorPhaseB) Writable() bool                { return false }
func (PowerFactorPhaseB) Reportable() bool              { return false }
func (PowerFactorPhaseB) SceneIndex() int               { return -1 }
func (a *PowerFactorPhaseB) Value() *PowerFactorPhaseB  { return a }
func (a PowerFactorPhaseB) MarshalZcl() ([]byte, error) { return zcl.Zs8(a).MarshalZcl() }

func (a *PowerFactorPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerFactorPhaseB(*nt)
	return br, err
}

func (a PowerFactorPhaseB) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(a))
}

// AverageRmsVoltageMeasurementPeriodPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsVoltageMeasurementPeriodPhaseB zcl.Zu16

const AverageRmsVoltageMeasurementPeriodPhaseBAttr zcl.AttrID = 2321

func (AverageRmsVoltageMeasurementPeriodPhaseB) ID() zcl.AttrID {
	return AverageRmsVoltageMeasurementPeriodPhaseBAttr
}
func (AverageRmsVoltageMeasurementPeriodPhaseB) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}
func (AverageRmsVoltageMeasurementPeriodPhaseB) Name() string {
	return "Average RMS Voltage Measurement Period Phase B"
}
func (AverageRmsVoltageMeasurementPeriodPhaseB) Readable() bool   { return true }
func (AverageRmsVoltageMeasurementPeriodPhaseB) Writable() bool   { return true }
func (AverageRmsVoltageMeasurementPeriodPhaseB) Reportable() bool { return false }
func (AverageRmsVoltageMeasurementPeriodPhaseB) SceneIndex() int  { return -1 }
func (a *AverageRmsVoltageMeasurementPeriodPhaseB) Value() *AverageRmsVoltageMeasurementPeriodPhaseB {
	return a
}
func (a AverageRmsVoltageMeasurementPeriodPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *AverageRmsVoltageMeasurementPeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsVoltageMeasurementPeriodPhaseB(*nt)
	return br, err
}

func (a AverageRmsVoltageMeasurementPeriodPhaseB) String() string {
	return zcl.Seconds.Format(float64(a))
}

// AverageRmsOvervoltageCounterPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsOvervoltageCounterPhaseB zcl.Zu16

const AverageRmsOvervoltageCounterPhaseBAttr zcl.AttrID = 2322

func (AverageRmsOvervoltageCounterPhaseB) ID() zcl.AttrID {
	return AverageRmsOvervoltageCounterPhaseBAttr
}
func (AverageRmsOvervoltageCounterPhaseB) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (AverageRmsOvervoltageCounterPhaseB) Name() string {
	return "Average RMS Overvoltage Counter Phase B"
}
func (AverageRmsOvervoltageCounterPhaseB) Readable() bool                                { return true }
func (AverageRmsOvervoltageCounterPhaseB) Writable() bool                                { return true }
func (AverageRmsOvervoltageCounterPhaseB) Reportable() bool                              { return false }
func (AverageRmsOvervoltageCounterPhaseB) SceneIndex() int                               { return -1 }
func (a *AverageRmsOvervoltageCounterPhaseB) Value() *AverageRmsOvervoltageCounterPhaseB { return a }
func (a AverageRmsOvervoltageCounterPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *AverageRmsOvervoltageCounterPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsOvervoltageCounterPhaseB(*nt)
	return br, err
}

func (a AverageRmsOvervoltageCounterPhaseB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AverageRmsUndervoltageCounterPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsUndervoltageCounterPhaseB zcl.Zu16

const AverageRmsUndervoltageCounterPhaseBAttr zcl.AttrID = 2323

func (AverageRmsUndervoltageCounterPhaseB) ID() zcl.AttrID {
	return AverageRmsUndervoltageCounterPhaseBAttr
}
func (AverageRmsUndervoltageCounterPhaseB) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (AverageRmsUndervoltageCounterPhaseB) Name() string {
	return "Average RMS Undervoltage Counter Phase B"
}
func (AverageRmsUndervoltageCounterPhaseB) Readable() bool                                 { return true }
func (AverageRmsUndervoltageCounterPhaseB) Writable() bool                                 { return true }
func (AverageRmsUndervoltageCounterPhaseB) Reportable() bool                               { return false }
func (AverageRmsUndervoltageCounterPhaseB) SceneIndex() int                                { return -1 }
func (a *AverageRmsUndervoltageCounterPhaseB) Value() *AverageRmsUndervoltageCounterPhaseB { return a }
func (a AverageRmsUndervoltageCounterPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *AverageRmsUndervoltageCounterPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsUndervoltageCounterPhaseB(*nt)
	return br, err
}

func (a AverageRmsUndervoltageCounterPhaseB) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// RmsExtremeOvervoltagePeriodPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsExtremeOvervoltagePeriodPhaseB zcl.Zu16

const RmsExtremeOvervoltagePeriodPhaseBAttr zcl.AttrID = 2324

func (RmsExtremeOvervoltagePeriodPhaseB) ID() zcl.AttrID         { return RmsExtremeOvervoltagePeriodPhaseBAttr }
func (RmsExtremeOvervoltagePeriodPhaseB) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (RmsExtremeOvervoltagePeriodPhaseB) Name() string {
	return "RMS Extreme Overvoltage Period Phase B"
}
func (RmsExtremeOvervoltagePeriodPhaseB) Readable() bool                               { return true }
func (RmsExtremeOvervoltagePeriodPhaseB) Writable() bool                               { return true }
func (RmsExtremeOvervoltagePeriodPhaseB) Reportable() bool                             { return false }
func (RmsExtremeOvervoltagePeriodPhaseB) SceneIndex() int                              { return -1 }
func (a *RmsExtremeOvervoltagePeriodPhaseB) Value() *RmsExtremeOvervoltagePeriodPhaseB { return a }
func (a RmsExtremeOvervoltagePeriodPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *RmsExtremeOvervoltagePeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeOvervoltagePeriodPhaseB(*nt)
	return br, err
}

func (a RmsExtremeOvervoltagePeriodPhaseB) String() string {
	return zcl.Seconds.Format(float64(a))
}

// RmsExtremeUndervoltagePeriodPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsExtremeUndervoltagePeriodPhaseB zcl.Zu16

const RmsExtremeUndervoltagePeriodPhaseBAttr zcl.AttrID = 2325

func (RmsExtremeUndervoltagePeriodPhaseB) ID() zcl.AttrID {
	return RmsExtremeUndervoltagePeriodPhaseBAttr
}
func (RmsExtremeUndervoltagePeriodPhaseB) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (RmsExtremeUndervoltagePeriodPhaseB) Name() string {
	return "RMS Extreme Undervoltage Period Phase B"
}
func (RmsExtremeUndervoltagePeriodPhaseB) Readable() bool                                { return true }
func (RmsExtremeUndervoltagePeriodPhaseB) Writable() bool                                { return true }
func (RmsExtremeUndervoltagePeriodPhaseB) Reportable() bool                              { return false }
func (RmsExtremeUndervoltagePeriodPhaseB) SceneIndex() int                               { return -1 }
func (a *RmsExtremeUndervoltagePeriodPhaseB) Value() *RmsExtremeUndervoltagePeriodPhaseB { return a }
func (a RmsExtremeUndervoltagePeriodPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *RmsExtremeUndervoltagePeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeUndervoltagePeriodPhaseB(*nt)
	return br, err
}

func (a RmsExtremeUndervoltagePeriodPhaseB) String() string {
	return zcl.Seconds.Format(float64(a))
}

// RmsVoltageSagPeriodPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageSagPeriodPhaseB zcl.Zu16

const RmsVoltageSagPeriodPhaseBAttr zcl.AttrID = 2326

func (RmsVoltageSagPeriodPhaseB) ID() zcl.AttrID                       { return RmsVoltageSagPeriodPhaseBAttr }
func (RmsVoltageSagPeriodPhaseB) Cluster() zcl.ClusterID               { return ElectricalMeasurementID }
func (RmsVoltageSagPeriodPhaseB) Name() string                         { return "RMS Voltage Sag Period Phase B" }
func (RmsVoltageSagPeriodPhaseB) Readable() bool                       { return true }
func (RmsVoltageSagPeriodPhaseB) Writable() bool                       { return true }
func (RmsVoltageSagPeriodPhaseB) Reportable() bool                     { return false }
func (RmsVoltageSagPeriodPhaseB) SceneIndex() int                      { return -1 }
func (a *RmsVoltageSagPeriodPhaseB) Value() *RmsVoltageSagPeriodPhaseB { return a }
func (a RmsVoltageSagPeriodPhaseB) MarshalZcl() ([]byte, error)        { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageSagPeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSagPeriodPhaseB(*nt)
	return br, err
}

func (a RmsVoltageSagPeriodPhaseB) String() string {
	return zcl.Seconds.Format(float64(a))
}

// RmsVoltageSwellPeriodPhaseB is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageSwellPeriodPhaseB zcl.Zu16

const RmsVoltageSwellPeriodPhaseBAttr zcl.AttrID = 2327

func (RmsVoltageSwellPeriodPhaseB) ID() zcl.AttrID                         { return RmsVoltageSwellPeriodPhaseBAttr }
func (RmsVoltageSwellPeriodPhaseB) Cluster() zcl.ClusterID                 { return ElectricalMeasurementID }
func (RmsVoltageSwellPeriodPhaseB) Name() string                           { return "RMS Voltage Swell Period Phase B" }
func (RmsVoltageSwellPeriodPhaseB) Readable() bool                         { return true }
func (RmsVoltageSwellPeriodPhaseB) Writable() bool                         { return true }
func (RmsVoltageSwellPeriodPhaseB) Reportable() bool                       { return false }
func (RmsVoltageSwellPeriodPhaseB) SceneIndex() int                        { return -1 }
func (a *RmsVoltageSwellPeriodPhaseB) Value() *RmsVoltageSwellPeriodPhaseB { return a }
func (a RmsVoltageSwellPeriodPhaseB) MarshalZcl() ([]byte, error)          { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageSwellPeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwellPeriodPhaseB(*nt)
	return br, err
}

func (a RmsVoltageSwellPeriodPhaseB) String() string {
	return zcl.Seconds.Format(float64(a))
}

// LineCurrentPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type LineCurrentPhaseC zcl.Zu16

const LineCurrentPhaseCAttr zcl.AttrID = 2561

func (LineCurrentPhaseC) ID() zcl.AttrID                { return LineCurrentPhaseCAttr }
func (LineCurrentPhaseC) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (LineCurrentPhaseC) Name() string                  { return "Line Current Phase C" }
func (LineCurrentPhaseC) Readable() bool                { return true }
func (LineCurrentPhaseC) Writable() bool                { return false }
func (LineCurrentPhaseC) Reportable() bool              { return false }
func (LineCurrentPhaseC) SceneIndex() int               { return -1 }
func (a *LineCurrentPhaseC) Value() *LineCurrentPhaseC  { return a }
func (a LineCurrentPhaseC) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *LineCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LineCurrentPhaseC(*nt)
	return br, err
}

func (a LineCurrentPhaseC) String() string {
	return zcl.Amperes.Format(float64(a))
}

// ActiveCurrentPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type ActiveCurrentPhaseC zcl.Zs16

const ActiveCurrentPhaseCAttr zcl.AttrID = 2562

func (ActiveCurrentPhaseC) ID() zcl.AttrID                 { return ActiveCurrentPhaseCAttr }
func (ActiveCurrentPhaseC) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (ActiveCurrentPhaseC) Name() string                   { return "Active Current Phase C" }
func (ActiveCurrentPhaseC) Readable() bool                 { return true }
func (ActiveCurrentPhaseC) Writable() bool                 { return false }
func (ActiveCurrentPhaseC) Reportable() bool               { return false }
func (ActiveCurrentPhaseC) SceneIndex() int                { return -1 }
func (a *ActiveCurrentPhaseC) Value() *ActiveCurrentPhaseC { return a }
func (a ActiveCurrentPhaseC) MarshalZcl() ([]byte, error)  { return zcl.Zs16(a).MarshalZcl() }

func (a *ActiveCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveCurrentPhaseC(*nt)
	return br, err
}

func (a ActiveCurrentPhaseC) String() string {
	return zcl.Amperes.Format(float64(a))
}

// ReactiveCurrentPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type ReactiveCurrentPhaseC zcl.Zs16

const ReactiveCurrentPhaseCAttr zcl.AttrID = 2563

func (ReactiveCurrentPhaseC) ID() zcl.AttrID                   { return ReactiveCurrentPhaseCAttr }
func (ReactiveCurrentPhaseC) Cluster() zcl.ClusterID           { return ElectricalMeasurementID }
func (ReactiveCurrentPhaseC) Name() string                     { return "Reactive Current Phase C" }
func (ReactiveCurrentPhaseC) Readable() bool                   { return true }
func (ReactiveCurrentPhaseC) Writable() bool                   { return false }
func (ReactiveCurrentPhaseC) Reportable() bool                 { return false }
func (ReactiveCurrentPhaseC) SceneIndex() int                  { return -1 }
func (a *ReactiveCurrentPhaseC) Value() *ReactiveCurrentPhaseC { return a }
func (a ReactiveCurrentPhaseC) MarshalZcl() ([]byte, error)    { return zcl.Zs16(a).MarshalZcl() }

func (a *ReactiveCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactiveCurrentPhaseC(*nt)
	return br, err
}

func (a ReactiveCurrentPhaseC) String() string {
	return zcl.Amperes.Format(float64(a))
}

// RmsVoltagePhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltagePhaseC zcl.Zu16

const RmsVoltagePhaseCAttr zcl.AttrID = 2565

func (RmsVoltagePhaseC) ID() zcl.AttrID                { return RmsVoltagePhaseCAttr }
func (RmsVoltagePhaseC) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsVoltagePhaseC) Name() string                  { return "RMS Voltage Phase C" }
func (RmsVoltagePhaseC) Readable() bool                { return true }
func (RmsVoltagePhaseC) Writable() bool                { return false }
func (RmsVoltagePhaseC) Reportable() bool              { return false }
func (RmsVoltagePhaseC) SceneIndex() int               { return -1 }
func (a *RmsVoltagePhaseC) Value() *RmsVoltagePhaseC   { return a }
func (a RmsVoltagePhaseC) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltagePhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltagePhaseC(*nt)
	return br, err
}

func (a RmsVoltagePhaseC) String() string {
	return zcl.Volts.Format(float64(a))
}

// RmsVoltageMinPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageMinPhaseC zcl.Zu16

const RmsVoltageMinPhaseCAttr zcl.AttrID = 2566

func (RmsVoltageMinPhaseC) ID() zcl.AttrID                 { return RmsVoltageMinPhaseCAttr }
func (RmsVoltageMinPhaseC) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (RmsVoltageMinPhaseC) Name() string                   { return "RMS Voltage Min Phase C" }
func (RmsVoltageMinPhaseC) Readable() bool                 { return true }
func (RmsVoltageMinPhaseC) Writable() bool                 { return false }
func (RmsVoltageMinPhaseC) Reportable() bool               { return false }
func (RmsVoltageMinPhaseC) SceneIndex() int                { return -1 }
func (a *RmsVoltageMinPhaseC) Value() *RmsVoltageMinPhaseC { return a }
func (a RmsVoltageMinPhaseC) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageMinPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMinPhaseC(*nt)
	return br, err
}

func (a RmsVoltageMinPhaseC) String() string {
	return zcl.Volts.Format(float64(a))
}

// RmsVoltageMaxPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageMaxPhaseC zcl.Zu16

const RmsVoltageMaxPhaseCAttr zcl.AttrID = 2567

func (RmsVoltageMaxPhaseC) ID() zcl.AttrID                 { return RmsVoltageMaxPhaseCAttr }
func (RmsVoltageMaxPhaseC) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (RmsVoltageMaxPhaseC) Name() string                   { return "RMS Voltage Max Phase C" }
func (RmsVoltageMaxPhaseC) Readable() bool                 { return true }
func (RmsVoltageMaxPhaseC) Writable() bool                 { return false }
func (RmsVoltageMaxPhaseC) Reportable() bool               { return false }
func (RmsVoltageMaxPhaseC) SceneIndex() int                { return -1 }
func (a *RmsVoltageMaxPhaseC) Value() *RmsVoltageMaxPhaseC { return a }
func (a RmsVoltageMaxPhaseC) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageMaxPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMaxPhaseC(*nt)
	return br, err
}

func (a RmsVoltageMaxPhaseC) String() string {
	return zcl.Volts.Format(float64(a))
}

// RmsCurrentPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsCurrentPhaseC zcl.Zu16

const RmsCurrentPhaseCAttr zcl.AttrID = 2568

func (RmsCurrentPhaseC) ID() zcl.AttrID                { return RmsCurrentPhaseCAttr }
func (RmsCurrentPhaseC) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (RmsCurrentPhaseC) Name() string                  { return "RMS Current Phase C" }
func (RmsCurrentPhaseC) Readable() bool                { return true }
func (RmsCurrentPhaseC) Writable() bool                { return false }
func (RmsCurrentPhaseC) Reportable() bool              { return false }
func (RmsCurrentPhaseC) SceneIndex() int               { return -1 }
func (a *RmsCurrentPhaseC) Value() *RmsCurrentPhaseC   { return a }
func (a RmsCurrentPhaseC) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentPhaseC(*nt)
	return br, err
}

func (a RmsCurrentPhaseC) String() string {
	return zcl.Amperes.Format(float64(a))
}

// RmsCurrentMinPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsCurrentMinPhaseC zcl.Zu16

const RmsCurrentMinPhaseCAttr zcl.AttrID = 2569

func (RmsCurrentMinPhaseC) ID() zcl.AttrID                 { return RmsCurrentMinPhaseCAttr }
func (RmsCurrentMinPhaseC) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (RmsCurrentMinPhaseC) Name() string                   { return "RMS Current Min Phase C" }
func (RmsCurrentMinPhaseC) Readable() bool                 { return true }
func (RmsCurrentMinPhaseC) Writable() bool                 { return false }
func (RmsCurrentMinPhaseC) Reportable() bool               { return false }
func (RmsCurrentMinPhaseC) SceneIndex() int                { return -1 }
func (a *RmsCurrentMinPhaseC) Value() *RmsCurrentMinPhaseC { return a }
func (a RmsCurrentMinPhaseC) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsCurrentMinPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMinPhaseC(*nt)
	return br, err
}

func (a RmsCurrentMinPhaseC) String() string {
	return zcl.Amperes.Format(float64(a))
}

// RmsCurrentMaxPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsCurrentMaxPhaseC zcl.Zu16

const RmsCurrentMaxPhaseCAttr zcl.AttrID = 2570

func (RmsCurrentMaxPhaseC) ID() zcl.AttrID                 { return RmsCurrentMaxPhaseCAttr }
func (RmsCurrentMaxPhaseC) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (RmsCurrentMaxPhaseC) Name() string                   { return "RMS Current Max Phase C" }
func (RmsCurrentMaxPhaseC) Readable() bool                 { return true }
func (RmsCurrentMaxPhaseC) Writable() bool                 { return false }
func (RmsCurrentMaxPhaseC) Reportable() bool               { return false }
func (RmsCurrentMaxPhaseC) SceneIndex() int                { return -1 }
func (a *RmsCurrentMaxPhaseC) Value() *RmsCurrentMaxPhaseC { return a }
func (a RmsCurrentMaxPhaseC) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsCurrentMaxPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMaxPhaseC(*nt)
	return br, err
}

func (a RmsCurrentMaxPhaseC) String() string {
	return zcl.Amperes.Format(float64(a))
}

// ActivePowerPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
// Represents the Phase C, current demand of active power delivered or received at the premises, in Watts (W). Positive values indicate power delivered to the premises where negative values indicate power received from the premises.
type ActivePowerPhaseC zcl.Zs16

const ActivePowerPhaseCAttr zcl.AttrID = 2571

func (ActivePowerPhaseC) ID() zcl.AttrID                { return ActivePowerPhaseCAttr }
func (ActivePowerPhaseC) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (ActivePowerPhaseC) Name() string                  { return "Active Power Phase C" }
func (ActivePowerPhaseC) Readable() bool                { return true }
func (ActivePowerPhaseC) Writable() bool                { return false }
func (ActivePowerPhaseC) Reportable() bool              { return false }
func (ActivePowerPhaseC) SceneIndex() int               { return -1 }
func (a *ActivePowerPhaseC) Value() *ActivePowerPhaseC  { return a }
func (a ActivePowerPhaseC) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *ActivePowerPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerPhaseC(*nt)
	return br, err
}

func (a ActivePowerPhaseC) String() string {
	return zcl.Watts.Format(float64(a))
}

// ActivePowerMinPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type ActivePowerMinPhaseC zcl.Zs16

const ActivePowerMinPhaseCAttr zcl.AttrID = 2572

func (ActivePowerMinPhaseC) ID() zcl.AttrID                  { return ActivePowerMinPhaseCAttr }
func (ActivePowerMinPhaseC) Cluster() zcl.ClusterID          { return ElectricalMeasurementID }
func (ActivePowerMinPhaseC) Name() string                    { return "Active Power Min Phase C" }
func (ActivePowerMinPhaseC) Readable() bool                  { return true }
func (ActivePowerMinPhaseC) Writable() bool                  { return false }
func (ActivePowerMinPhaseC) Reportable() bool                { return false }
func (ActivePowerMinPhaseC) SceneIndex() int                 { return -1 }
func (a *ActivePowerMinPhaseC) Value() *ActivePowerMinPhaseC { return a }
func (a ActivePowerMinPhaseC) MarshalZcl() ([]byte, error)   { return zcl.Zs16(a).MarshalZcl() }

func (a *ActivePowerMinPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMinPhaseC(*nt)
	return br, err
}

func (a ActivePowerMinPhaseC) String() string {
	return zcl.Watts.Format(float64(a))
}

// ActivePowerMaxPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type ActivePowerMaxPhaseC zcl.Zs16

const ActivePowerMaxPhaseCAttr zcl.AttrID = 2573

func (ActivePowerMaxPhaseC) ID() zcl.AttrID                  { return ActivePowerMaxPhaseCAttr }
func (ActivePowerMaxPhaseC) Cluster() zcl.ClusterID          { return ElectricalMeasurementID }
func (ActivePowerMaxPhaseC) Name() string                    { return "Active Power Max Phase C" }
func (ActivePowerMaxPhaseC) Readable() bool                  { return true }
func (ActivePowerMaxPhaseC) Writable() bool                  { return false }
func (ActivePowerMaxPhaseC) Reportable() bool                { return false }
func (ActivePowerMaxPhaseC) SceneIndex() int                 { return -1 }
func (a *ActivePowerMaxPhaseC) Value() *ActivePowerMaxPhaseC { return a }
func (a ActivePowerMaxPhaseC) MarshalZcl() ([]byte, error)   { return zcl.Zs16(a).MarshalZcl() }

func (a *ActivePowerMaxPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMaxPhaseC(*nt)
	return br, err
}

func (a ActivePowerMaxPhaseC) String() string {
	return zcl.Watts.Format(float64(a))
}

// ReactivePowerPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type ReactivePowerPhaseC zcl.Zs16

const ReactivePowerPhaseCAttr zcl.AttrID = 2574

func (ReactivePowerPhaseC) ID() zcl.AttrID                 { return ReactivePowerPhaseCAttr }
func (ReactivePowerPhaseC) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (ReactivePowerPhaseC) Name() string                   { return "Reactive Power Phase C" }
func (ReactivePowerPhaseC) Readable() bool                 { return true }
func (ReactivePowerPhaseC) Writable() bool                 { return false }
func (ReactivePowerPhaseC) Reportable() bool               { return false }
func (ReactivePowerPhaseC) SceneIndex() int                { return -1 }
func (a *ReactivePowerPhaseC) Value() *ReactivePowerPhaseC { return a }
func (a ReactivePowerPhaseC) MarshalZcl() ([]byte, error)  { return zcl.Zs16(a).MarshalZcl() }

func (a *ReactivePowerPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactivePowerPhaseC(*nt)
	return br, err
}

func (a ReactivePowerPhaseC) String() string {
	return zcl.VoltAmperesReactive.Format(float64(a))
}

// ApparentPowerPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type ApparentPowerPhaseC zcl.Zu16

const ApparentPowerPhaseCAttr zcl.AttrID = 2575

func (ApparentPowerPhaseC) ID() zcl.AttrID                 { return ApparentPowerPhaseCAttr }
func (ApparentPowerPhaseC) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (ApparentPowerPhaseC) Name() string                   { return "Apparent Power Phase C" }
func (ApparentPowerPhaseC) Readable() bool                 { return true }
func (ApparentPowerPhaseC) Writable() bool                 { return false }
func (ApparentPowerPhaseC) Reportable() bool               { return false }
func (ApparentPowerPhaseC) SceneIndex() int                { return -1 }
func (a *ApparentPowerPhaseC) Value() *ApparentPowerPhaseC { return a }
func (a ApparentPowerPhaseC) MarshalZcl() ([]byte, error)  { return zcl.Zu16(a).MarshalZcl() }

func (a *ApparentPowerPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApparentPowerPhaseC(*nt)
	return br, err
}

func (a ApparentPowerPhaseC) String() string {
	return zcl.VoltAmperes.Format(float64(a))
}

// PowerFactorPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type PowerFactorPhaseC zcl.Zs8

const PowerFactorPhaseCAttr zcl.AttrID = 2576

func (PowerFactorPhaseC) ID() zcl.AttrID                { return PowerFactorPhaseCAttr }
func (PowerFactorPhaseC) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (PowerFactorPhaseC) Name() string                  { return "Power Factor Phase C" }
func (PowerFactorPhaseC) Readable() bool                { return true }
func (PowerFactorPhaseC) Writable() bool                { return false }
func (PowerFactorPhaseC) Reportable() bool              { return false }
func (PowerFactorPhaseC) SceneIndex() int               { return -1 }
func (a *PowerFactorPhaseC) Value() *PowerFactorPhaseC  { return a }
func (a PowerFactorPhaseC) MarshalZcl() ([]byte, error) { return zcl.Zs8(a).MarshalZcl() }

func (a *PowerFactorPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerFactorPhaseC(*nt)
	return br, err
}

func (a PowerFactorPhaseC) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(a))
}

// AverageRmsVoltageMeasurementPeriodPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsVoltageMeasurementPeriodPhaseC zcl.Zu16

const AverageRmsVoltageMeasurementPeriodPhaseCAttr zcl.AttrID = 2577

func (AverageRmsVoltageMeasurementPeriodPhaseC) ID() zcl.AttrID {
	return AverageRmsVoltageMeasurementPeriodPhaseCAttr
}
func (AverageRmsVoltageMeasurementPeriodPhaseC) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}
func (AverageRmsVoltageMeasurementPeriodPhaseC) Name() string {
	return "Average RMS Voltage Measurement Period Phase C"
}
func (AverageRmsVoltageMeasurementPeriodPhaseC) Readable() bool   { return true }
func (AverageRmsVoltageMeasurementPeriodPhaseC) Writable() bool   { return true }
func (AverageRmsVoltageMeasurementPeriodPhaseC) Reportable() bool { return false }
func (AverageRmsVoltageMeasurementPeriodPhaseC) SceneIndex() int  { return -1 }
func (a *AverageRmsVoltageMeasurementPeriodPhaseC) Value() *AverageRmsVoltageMeasurementPeriodPhaseC {
	return a
}
func (a AverageRmsVoltageMeasurementPeriodPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *AverageRmsVoltageMeasurementPeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsVoltageMeasurementPeriodPhaseC(*nt)
	return br, err
}

func (a AverageRmsVoltageMeasurementPeriodPhaseC) String() string {
	return zcl.Seconds.Format(float64(a))
}

// AverageRmsOvervoltageCounterPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsOvervoltageCounterPhaseC zcl.Zu16

const AverageRmsOvervoltageCounterPhaseCAttr zcl.AttrID = 2578

func (AverageRmsOvervoltageCounterPhaseC) ID() zcl.AttrID {
	return AverageRmsOvervoltageCounterPhaseCAttr
}
func (AverageRmsOvervoltageCounterPhaseC) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (AverageRmsOvervoltageCounterPhaseC) Name() string {
	return "Average RMS Overvoltage Counter Phase C"
}
func (AverageRmsOvervoltageCounterPhaseC) Readable() bool                                { return true }
func (AverageRmsOvervoltageCounterPhaseC) Writable() bool                                { return true }
func (AverageRmsOvervoltageCounterPhaseC) Reportable() bool                              { return false }
func (AverageRmsOvervoltageCounterPhaseC) SceneIndex() int                               { return -1 }
func (a *AverageRmsOvervoltageCounterPhaseC) Value() *AverageRmsOvervoltageCounterPhaseC { return a }
func (a AverageRmsOvervoltageCounterPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *AverageRmsOvervoltageCounterPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsOvervoltageCounterPhaseC(*nt)
	return br, err
}

func (a AverageRmsOvervoltageCounterPhaseC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// AverageRmsUndervoltageCounterPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type AverageRmsUndervoltageCounterPhaseC zcl.Zu16

const AverageRmsUndervoltageCounterPhaseCAttr zcl.AttrID = 2579

func (AverageRmsUndervoltageCounterPhaseC) ID() zcl.AttrID {
	return AverageRmsUndervoltageCounterPhaseCAttr
}
func (AverageRmsUndervoltageCounterPhaseC) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (AverageRmsUndervoltageCounterPhaseC) Name() string {
	return "Average RMS Undervoltage Counter Phase C"
}
func (AverageRmsUndervoltageCounterPhaseC) Readable() bool                                 { return true }
func (AverageRmsUndervoltageCounterPhaseC) Writable() bool                                 { return true }
func (AverageRmsUndervoltageCounterPhaseC) Reportable() bool                               { return false }
func (AverageRmsUndervoltageCounterPhaseC) SceneIndex() int                                { return -1 }
func (a *AverageRmsUndervoltageCounterPhaseC) Value() *AverageRmsUndervoltageCounterPhaseC { return a }
func (a AverageRmsUndervoltageCounterPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *AverageRmsUndervoltageCounterPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsUndervoltageCounterPhaseC(*nt)
	return br, err
}

func (a AverageRmsUndervoltageCounterPhaseC) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// RmsExtremeOvervoltagePeriodPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsExtremeOvervoltagePeriodPhaseC zcl.Zu16

const RmsExtremeOvervoltagePeriodPhaseCAttr zcl.AttrID = 2580

func (RmsExtremeOvervoltagePeriodPhaseC) ID() zcl.AttrID         { return RmsExtremeOvervoltagePeriodPhaseCAttr }
func (RmsExtremeOvervoltagePeriodPhaseC) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (RmsExtremeOvervoltagePeriodPhaseC) Name() string {
	return "RMS Extreme Overvoltage Period Phase C"
}
func (RmsExtremeOvervoltagePeriodPhaseC) Readable() bool                               { return true }
func (RmsExtremeOvervoltagePeriodPhaseC) Writable() bool                               { return true }
func (RmsExtremeOvervoltagePeriodPhaseC) Reportable() bool                             { return false }
func (RmsExtremeOvervoltagePeriodPhaseC) SceneIndex() int                              { return -1 }
func (a *RmsExtremeOvervoltagePeriodPhaseC) Value() *RmsExtremeOvervoltagePeriodPhaseC { return a }
func (a RmsExtremeOvervoltagePeriodPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *RmsExtremeOvervoltagePeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeOvervoltagePeriodPhaseC(*nt)
	return br, err
}

func (a RmsExtremeOvervoltagePeriodPhaseC) String() string {
	return zcl.Seconds.Format(float64(a))
}

// RmsExtremeUndervoltagePeriodPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsExtremeUndervoltagePeriodPhaseC zcl.Zu16

const RmsExtremeUndervoltagePeriodPhaseCAttr zcl.AttrID = 2581

func (RmsExtremeUndervoltagePeriodPhaseC) ID() zcl.AttrID {
	return RmsExtremeUndervoltagePeriodPhaseCAttr
}
func (RmsExtremeUndervoltagePeriodPhaseC) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (RmsExtremeUndervoltagePeriodPhaseC) Name() string {
	return "RMS Extreme Undervoltage Period Phase C"
}
func (RmsExtremeUndervoltagePeriodPhaseC) Readable() bool                                { return true }
func (RmsExtremeUndervoltagePeriodPhaseC) Writable() bool                                { return true }
func (RmsExtremeUndervoltagePeriodPhaseC) Reportable() bool                              { return false }
func (RmsExtremeUndervoltagePeriodPhaseC) SceneIndex() int                               { return -1 }
func (a *RmsExtremeUndervoltagePeriodPhaseC) Value() *RmsExtremeUndervoltagePeriodPhaseC { return a }
func (a RmsExtremeUndervoltagePeriodPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *RmsExtremeUndervoltagePeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeUndervoltagePeriodPhaseC(*nt)
	return br, err
}

func (a RmsExtremeUndervoltagePeriodPhaseC) String() string {
	return zcl.Seconds.Format(float64(a))
}

// RmsVoltageSagPeriodPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageSagPeriodPhaseC zcl.Zu16

const RmsVoltageSagPeriodPhaseCAttr zcl.AttrID = 2582

func (RmsVoltageSagPeriodPhaseC) ID() zcl.AttrID                       { return RmsVoltageSagPeriodPhaseCAttr }
func (RmsVoltageSagPeriodPhaseC) Cluster() zcl.ClusterID               { return ElectricalMeasurementID }
func (RmsVoltageSagPeriodPhaseC) Name() string                         { return "RMS Voltage Sag Period Phase C" }
func (RmsVoltageSagPeriodPhaseC) Readable() bool                       { return true }
func (RmsVoltageSagPeriodPhaseC) Writable() bool                       { return true }
func (RmsVoltageSagPeriodPhaseC) Reportable() bool                     { return false }
func (RmsVoltageSagPeriodPhaseC) SceneIndex() int                      { return -1 }
func (a *RmsVoltageSagPeriodPhaseC) Value() *RmsVoltageSagPeriodPhaseC { return a }
func (a RmsVoltageSagPeriodPhaseC) MarshalZcl() ([]byte, error)        { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageSagPeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSagPeriodPhaseC(*nt)
	return br, err
}

func (a RmsVoltageSagPeriodPhaseC) String() string {
	return zcl.Seconds.Format(float64(a))
}

// RmsVoltageSwellPeriodPhaseC is an autogenerated attribute in the ElectricalMeasurement cluster
type RmsVoltageSwellPeriodPhaseC zcl.Zu16

const RmsVoltageSwellPeriodPhaseCAttr zcl.AttrID = 2583

func (RmsVoltageSwellPeriodPhaseC) ID() zcl.AttrID                         { return RmsVoltageSwellPeriodPhaseCAttr }
func (RmsVoltageSwellPeriodPhaseC) Cluster() zcl.ClusterID                 { return ElectricalMeasurementID }
func (RmsVoltageSwellPeriodPhaseC) Name() string                           { return "RMS Voltage Swell Period Phase C" }
func (RmsVoltageSwellPeriodPhaseC) Readable() bool                         { return true }
func (RmsVoltageSwellPeriodPhaseC) Writable() bool                         { return true }
func (RmsVoltageSwellPeriodPhaseC) Reportable() bool                       { return false }
func (RmsVoltageSwellPeriodPhaseC) SceneIndex() int                        { return -1 }
func (a *RmsVoltageSwellPeriodPhaseC) Value() *RmsVoltageSwellPeriodPhaseC { return a }
func (a RmsVoltageSwellPeriodPhaseC) MarshalZcl() ([]byte, error)          { return zcl.Zu16(a).MarshalZcl() }

func (a *RmsVoltageSwellPeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwellPeriodPhaseC(*nt)
	return br, err
}

func (a RmsVoltageSwellPeriodPhaseC) String() string {
	return zcl.Seconds.Format(float64(a))
}
