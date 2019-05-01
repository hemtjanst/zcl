// Provides a mechanism for querying data about the electrical properties as measured by the device.
package measurement_and_sensing

import (
	"neotor.se/zcl"
)

// ElectricalMeasurement
const ElectricalMeasurementID zcl.ClusterID = 2820

var ElectricalMeasurementCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		GetProfileInfoResponseCommandCommand:        func() zcl.Command { return new(GetProfileInfoResponseCommand) },
		GetMeasurementProfileResponseCommandCommand: func() zcl.Command { return new(GetMeasurementProfileResponseCommand) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		GetProfileInfoCommandCommand:        func() zcl.Command { return new(GetProfileInfoCommand) },
		GetMeasurementProfileCommandCommand: func() zcl.Command { return new(GetMeasurementProfileCommand) },
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

type GetProfileInfoResponseCommand struct {
	ProfileCount          zcl.Zu8
	ProfileIntervalPeriod zcl.Zenum8
	MaxNumberOfIntervals  zcl.Zu8
	ListOfAttributes      zcl.Zarray
}

const GetProfileInfoResponseCommandCommand zcl.CommandID = 0

func (v *GetProfileInfoResponseCommand) Values() []zcl.Val {
	return []zcl.Val{
		&v.ProfileCount,
		&v.ProfileIntervalPeriod,
		&v.MaxNumberOfIntervals,
		&v.ListOfAttributes,
	}
}

func (v GetProfileInfoResponseCommand) ID() zcl.CommandID {
	return GetProfileInfoResponseCommandCommand
}

func (v GetProfileInfoResponseCommand) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}

func (v GetProfileInfoResponseCommand) MnfCode() []byte {
	return []byte{}
}

func (v GetProfileInfoResponseCommand) MarshalZcl() ([]byte, error) {
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

func (v *GetProfileInfoResponseCommand) UnmarshalZcl(b []byte) ([]byte, error) {
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

type GetMeasurementProfileResponseCommand struct {
	StartTime                  zcl.Zutc
	Status                     zcl.Zenum8
	ProfileIntervalPeriod      zcl.Zenum8
	NumberOfIntervalsDelivered zcl.Zu8
	AttributeId                zcl.Zu16
	Intervals                  zcl.Zarray
}

const GetMeasurementProfileResponseCommandCommand zcl.CommandID = 1

func (v *GetMeasurementProfileResponseCommand) Values() []zcl.Val {
	return []zcl.Val{
		&v.StartTime,
		&v.Status,
		&v.ProfileIntervalPeriod,
		&v.NumberOfIntervalsDelivered,
		&v.AttributeId,
		&v.Intervals,
	}
}

func (v GetMeasurementProfileResponseCommand) ID() zcl.CommandID {
	return GetMeasurementProfileResponseCommandCommand
}

func (v GetMeasurementProfileResponseCommand) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}

func (v GetMeasurementProfileResponseCommand) MnfCode() []byte {
	return []byte{}
}

func (v GetMeasurementProfileResponseCommand) MarshalZcl() ([]byte, error) {
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

func (v *GetMeasurementProfileResponseCommand) UnmarshalZcl(b []byte) ([]byte, error) {
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

type GetProfileInfoCommand struct {
}

const GetProfileInfoCommandCommand zcl.CommandID = 0

func (v *GetProfileInfoCommand) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v GetProfileInfoCommand) ID() zcl.CommandID {
	return GetProfileInfoCommandCommand
}

func (v GetProfileInfoCommand) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}

func (v GetProfileInfoCommand) MnfCode() []byte {
	return []byte{}
}

func (v GetProfileInfoCommand) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *GetProfileInfoCommand) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

type GetMeasurementProfileCommand struct {
	AttributeId       zcl.Zu16
	StartTime         zcl.Zutc
	NumberOfIntervals zcl.Zu8
}

const GetMeasurementProfileCommandCommand zcl.CommandID = 1

func (v *GetMeasurementProfileCommand) Values() []zcl.Val {
	return []zcl.Val{
		&v.AttributeId,
		&v.StartTime,
		&v.NumberOfIntervals,
	}
}

func (v GetMeasurementProfileCommand) ID() zcl.CommandID {
	return GetMeasurementProfileCommandCommand
}

func (v GetMeasurementProfileCommand) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}

func (v GetMeasurementProfileCommand) MnfCode() []byte {
	return []byte{}
}

func (v GetMeasurementProfileCommand) MarshalZcl() ([]byte, error) {
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

func (v *GetMeasurementProfileCommand) UnmarshalZcl(b []byte) ([]byte, error) {
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

const MeasurementTypeAttr zcl.AttrID = 0

type MeasurementType zcl.Zbmp32

func (a MeasurementType) ID() zcl.AttrID           { return MeasurementTypeAttr }
func (a MeasurementType) Cluster() zcl.ClusterID   { return ElectricalMeasurementID }
func (a *MeasurementType) Value() *MeasurementType { return a }
func (a MeasurementType) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp32(a).MarshalZcl()
}
func (a *MeasurementType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasurementType(*nt)
	return br, err
}

func (a MeasurementType) Readable() bool   { return true }
func (a MeasurementType) Writable() bool   { return false }
func (a MeasurementType) Reportable() bool { return false }
func (a MeasurementType) SceneIndex() int  { return -1 }

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

const DcVoltageAttr zcl.AttrID = 256

type DcVoltage zcl.Zs16

func (a DcVoltage) ID() zcl.AttrID         { return DcVoltageAttr }
func (a DcVoltage) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *DcVoltage) Value() *DcVoltage     { return a }
func (a DcVoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltage(*nt)
	return br, err
}

func (a DcVoltage) Readable() bool   { return true }
func (a DcVoltage) Writable() bool   { return false }
func (a DcVoltage) Reportable() bool { return false }
func (a DcVoltage) SceneIndex() int  { return -1 }

func (a DcVoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const DcVoltageMinAttr zcl.AttrID = 257

type DcVoltageMin zcl.Zs16

func (a DcVoltageMin) ID() zcl.AttrID         { return DcVoltageMinAttr }
func (a DcVoltageMin) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *DcVoltageMin) Value() *DcVoltageMin  { return a }
func (a DcVoltageMin) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcVoltageMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageMin(*nt)
	return br, err
}

func (a DcVoltageMin) Readable() bool   { return true }
func (a DcVoltageMin) Writable() bool   { return false }
func (a DcVoltageMin) Reportable() bool { return false }
func (a DcVoltageMin) SceneIndex() int  { return -1 }

func (a DcVoltageMin) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const DcVoltageMaxAttr zcl.AttrID = 258

type DcVoltageMax zcl.Zs16

func (a DcVoltageMax) ID() zcl.AttrID         { return DcVoltageMaxAttr }
func (a DcVoltageMax) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *DcVoltageMax) Value() *DcVoltageMax  { return a }
func (a DcVoltageMax) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcVoltageMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageMax(*nt)
	return br, err
}

func (a DcVoltageMax) Readable() bool   { return true }
func (a DcVoltageMax) Writable() bool   { return false }
func (a DcVoltageMax) Reportable() bool { return false }
func (a DcVoltageMax) SceneIndex() int  { return -1 }

func (a DcVoltageMax) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const DcCurrentAttr zcl.AttrID = 259

type DcCurrent zcl.Zs16

func (a DcCurrent) ID() zcl.AttrID         { return DcCurrentAttr }
func (a DcCurrent) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *DcCurrent) Value() *DcCurrent     { return a }
func (a DcCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrent(*nt)
	return br, err
}

func (a DcCurrent) Readable() bool   { return true }
func (a DcCurrent) Writable() bool   { return false }
func (a DcCurrent) Reportable() bool { return false }
func (a DcCurrent) SceneIndex() int  { return -1 }

func (a DcCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const DcCurrentMinAttr zcl.AttrID = 260

type DcCurrentMin zcl.Zs16

func (a DcCurrentMin) ID() zcl.AttrID         { return DcCurrentMinAttr }
func (a DcCurrentMin) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *DcCurrentMin) Value() *DcCurrentMin  { return a }
func (a DcCurrentMin) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcCurrentMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentMin(*nt)
	return br, err
}

func (a DcCurrentMin) Readable() bool   { return true }
func (a DcCurrentMin) Writable() bool   { return false }
func (a DcCurrentMin) Reportable() bool { return false }
func (a DcCurrentMin) SceneIndex() int  { return -1 }

func (a DcCurrentMin) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const DcCurrentMaxAttr zcl.AttrID = 261

type DcCurrentMax zcl.Zs16

func (a DcCurrentMax) ID() zcl.AttrID         { return DcCurrentMaxAttr }
func (a DcCurrentMax) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *DcCurrentMax) Value() *DcCurrentMax  { return a }
func (a DcCurrentMax) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcCurrentMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentMax(*nt)
	return br, err
}

func (a DcCurrentMax) Readable() bool   { return true }
func (a DcCurrentMax) Writable() bool   { return false }
func (a DcCurrentMax) Reportable() bool { return false }
func (a DcCurrentMax) SceneIndex() int  { return -1 }

func (a DcCurrentMax) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const DcPowerAttr zcl.AttrID = 262

type DcPower zcl.Zs16

func (a DcPower) ID() zcl.AttrID         { return DcPowerAttr }
func (a DcPower) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *DcPower) Value() *DcPower       { return a }
func (a DcPower) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPower(*nt)
	return br, err
}

func (a DcPower) Readable() bool   { return true }
func (a DcPower) Writable() bool   { return false }
func (a DcPower) Reportable() bool { return false }
func (a DcPower) SceneIndex() int  { return -1 }

func (a DcPower) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const DcPowerMinAttr zcl.AttrID = 263

type DcPowerMin zcl.Zs16

func (a DcPowerMin) ID() zcl.AttrID         { return DcPowerMinAttr }
func (a DcPowerMin) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *DcPowerMin) Value() *DcPowerMin    { return a }
func (a DcPowerMin) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcPowerMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerMin(*nt)
	return br, err
}

func (a DcPowerMin) Readable() bool   { return true }
func (a DcPowerMin) Writable() bool   { return false }
func (a DcPowerMin) Reportable() bool { return false }
func (a DcPowerMin) SceneIndex() int  { return -1 }

func (a DcPowerMin) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const DcPowerMaxAttr zcl.AttrID = 264

type DcPowerMax zcl.Zs16

func (a DcPowerMax) ID() zcl.AttrID         { return DcPowerMaxAttr }
func (a DcPowerMax) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *DcPowerMax) Value() *DcPowerMax    { return a }
func (a DcPowerMax) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcPowerMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerMax(*nt)
	return br, err
}

func (a DcPowerMax) Readable() bool   { return true }
func (a DcPowerMax) Writable() bool   { return false }
func (a DcPowerMax) Reportable() bool { return false }
func (a DcPowerMax) SceneIndex() int  { return -1 }

func (a DcPowerMax) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const DcVoltageMultiplierAttr zcl.AttrID = 512

type DcVoltageMultiplier zcl.Zu16

func (a DcVoltageMultiplier) ID() zcl.AttrID               { return DcVoltageMultiplierAttr }
func (a DcVoltageMultiplier) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *DcVoltageMultiplier) Value() *DcVoltageMultiplier { return a }
func (a DcVoltageMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcVoltageMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageMultiplier(*nt)
	return br, err
}

func (a DcVoltageMultiplier) Readable() bool   { return true }
func (a DcVoltageMultiplier) Writable() bool   { return false }
func (a DcVoltageMultiplier) Reportable() bool { return false }
func (a DcVoltageMultiplier) SceneIndex() int  { return -1 }

func (a DcVoltageMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const DcVoltageDivisorAttr zcl.AttrID = 513

type DcVoltageDivisor zcl.Zu16

func (a DcVoltageDivisor) ID() zcl.AttrID            { return DcVoltageDivisorAttr }
func (a DcVoltageDivisor) Cluster() zcl.ClusterID    { return ElectricalMeasurementID }
func (a *DcVoltageDivisor) Value() *DcVoltageDivisor { return a }
func (a DcVoltageDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcVoltageDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageDivisor(*nt)
	return br, err
}

func (a DcVoltageDivisor) Readable() bool   { return true }
func (a DcVoltageDivisor) Writable() bool   { return false }
func (a DcVoltageDivisor) Reportable() bool { return false }
func (a DcVoltageDivisor) SceneIndex() int  { return -1 }

func (a DcVoltageDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const DcCurrentMultiplierAttr zcl.AttrID = 514

type DcCurrentMultiplier zcl.Zu16

func (a DcCurrentMultiplier) ID() zcl.AttrID               { return DcCurrentMultiplierAttr }
func (a DcCurrentMultiplier) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *DcCurrentMultiplier) Value() *DcCurrentMultiplier { return a }
func (a DcCurrentMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentMultiplier(*nt)
	return br, err
}

func (a DcCurrentMultiplier) Readable() bool   { return true }
func (a DcCurrentMultiplier) Writable() bool   { return false }
func (a DcCurrentMultiplier) Reportable() bool { return false }
func (a DcCurrentMultiplier) SceneIndex() int  { return -1 }

func (a DcCurrentMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const DcCurrentDivisorAttr zcl.AttrID = 515

type DcCurrentDivisor zcl.Zu16

func (a DcCurrentDivisor) ID() zcl.AttrID            { return DcCurrentDivisorAttr }
func (a DcCurrentDivisor) Cluster() zcl.ClusterID    { return ElectricalMeasurementID }
func (a *DcCurrentDivisor) Value() *DcCurrentDivisor { return a }
func (a DcCurrentDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcCurrentDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentDivisor(*nt)
	return br, err
}

func (a DcCurrentDivisor) Readable() bool   { return true }
func (a DcCurrentDivisor) Writable() bool   { return false }
func (a DcCurrentDivisor) Reportable() bool { return false }
func (a DcCurrentDivisor) SceneIndex() int  { return -1 }

func (a DcCurrentDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const DcPowerMultiplierAttr zcl.AttrID = 516

type DcPowerMultiplier zcl.Zu16

func (a DcPowerMultiplier) ID() zcl.AttrID             { return DcPowerMultiplierAttr }
func (a DcPowerMultiplier) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *DcPowerMultiplier) Value() *DcPowerMultiplier { return a }
func (a DcPowerMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcPowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerMultiplier(*nt)
	return br, err
}

func (a DcPowerMultiplier) Readable() bool   { return true }
func (a DcPowerMultiplier) Writable() bool   { return false }
func (a DcPowerMultiplier) Reportable() bool { return false }
func (a DcPowerMultiplier) SceneIndex() int  { return -1 }

func (a DcPowerMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const DcPowerDivisorAttr zcl.AttrID = 517

type DcPowerDivisor zcl.Zu16

func (a DcPowerDivisor) ID() zcl.AttrID          { return DcPowerDivisorAttr }
func (a DcPowerDivisor) Cluster() zcl.ClusterID  { return ElectricalMeasurementID }
func (a *DcPowerDivisor) Value() *DcPowerDivisor { return a }
func (a DcPowerDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DcPowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcPowerDivisor(*nt)
	return br, err
}

func (a DcPowerDivisor) Readable() bool   { return true }
func (a DcPowerDivisor) Writable() bool   { return false }
func (a DcPowerDivisor) Reportable() bool { return false }
func (a DcPowerDivisor) SceneIndex() int  { return -1 }

func (a DcPowerDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AcFrequencyAttr zcl.AttrID = 768

type AcFrequency zcl.Zu16

func (a AcFrequency) ID() zcl.AttrID         { return AcFrequencyAttr }
func (a AcFrequency) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *AcFrequency) Value() *AcFrequency   { return a }
func (a AcFrequency) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcFrequency) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequency(*nt)
	return br, err
}

func (a AcFrequency) Readable() bool   { return true }
func (a AcFrequency) Writable() bool   { return false }
func (a AcFrequency) Reportable() bool { return false }
func (a AcFrequency) SceneIndex() int  { return -1 }

func (a AcFrequency) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AcFrequencyMinAttr zcl.AttrID = 769

type AcFrequencyMin zcl.Zu16

func (a AcFrequencyMin) ID() zcl.AttrID          { return AcFrequencyMinAttr }
func (a AcFrequencyMin) Cluster() zcl.ClusterID  { return ElectricalMeasurementID }
func (a *AcFrequencyMin) Value() *AcFrequencyMin { return a }
func (a AcFrequencyMin) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcFrequencyMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyMin(*nt)
	return br, err
}

func (a AcFrequencyMin) Readable() bool   { return true }
func (a AcFrequencyMin) Writable() bool   { return false }
func (a AcFrequencyMin) Reportable() bool { return false }
func (a AcFrequencyMin) SceneIndex() int  { return -1 }

func (a AcFrequencyMin) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AcFrequencyMaxAttr zcl.AttrID = 770

type AcFrequencyMax zcl.Zu16

func (a AcFrequencyMax) ID() zcl.AttrID          { return AcFrequencyMaxAttr }
func (a AcFrequencyMax) Cluster() zcl.ClusterID  { return ElectricalMeasurementID }
func (a *AcFrequencyMax) Value() *AcFrequencyMax { return a }
func (a AcFrequencyMax) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcFrequencyMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyMax(*nt)
	return br, err
}

func (a AcFrequencyMax) Readable() bool   { return true }
func (a AcFrequencyMax) Writable() bool   { return false }
func (a AcFrequencyMax) Reportable() bool { return false }
func (a AcFrequencyMax) SceneIndex() int  { return -1 }

func (a AcFrequencyMax) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const NeutralCurrentAttr zcl.AttrID = 771

type NeutralCurrent zcl.Zu16

func (a NeutralCurrent) ID() zcl.AttrID          { return NeutralCurrentAttr }
func (a NeutralCurrent) Cluster() zcl.ClusterID  { return ElectricalMeasurementID }
func (a *NeutralCurrent) Value() *NeutralCurrent { return a }
func (a NeutralCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NeutralCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NeutralCurrent(*nt)
	return br, err
}

func (a NeutralCurrent) Readable() bool   { return true }
func (a NeutralCurrent) Writable() bool   { return false }
func (a NeutralCurrent) Reportable() bool { return false }
func (a NeutralCurrent) SceneIndex() int  { return -1 }

func (a NeutralCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const TotalActivePowerAttr zcl.AttrID = 772

type TotalActivePower zcl.Zs32

func (a TotalActivePower) ID() zcl.AttrID            { return TotalActivePowerAttr }
func (a TotalActivePower) Cluster() zcl.ClusterID    { return ElectricalMeasurementID }
func (a *TotalActivePower) Value() *TotalActivePower { return a }
func (a TotalActivePower) MarshalZcl() ([]byte, error) {
	return zcl.Zs32(a).MarshalZcl()
}
func (a *TotalActivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = TotalActivePower(*nt)
	return br, err
}

func (a TotalActivePower) Readable() bool   { return true }
func (a TotalActivePower) Writable() bool   { return false }
func (a TotalActivePower) Reportable() bool { return false }
func (a TotalActivePower) SceneIndex() int  { return -1 }

func (a TotalActivePower) String() string {

	return zcl.Sprintf("%s", zcl.Zs32(a))
}

const TotalReactivePowerAttr zcl.AttrID = 773

type TotalReactivePower zcl.Zs32

func (a TotalReactivePower) ID() zcl.AttrID              { return TotalReactivePowerAttr }
func (a TotalReactivePower) Cluster() zcl.ClusterID      { return ElectricalMeasurementID }
func (a *TotalReactivePower) Value() *TotalReactivePower { return a }
func (a TotalReactivePower) MarshalZcl() ([]byte, error) {
	return zcl.Zs32(a).MarshalZcl()
}
func (a *TotalReactivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs32)
	br, err := nt.UnmarshalZcl(b)
	*a = TotalReactivePower(*nt)
	return br, err
}

func (a TotalReactivePower) Readable() bool   { return true }
func (a TotalReactivePower) Writable() bool   { return false }
func (a TotalReactivePower) Reportable() bool { return false }
func (a TotalReactivePower) SceneIndex() int  { return -1 }

func (a TotalReactivePower) String() string {

	return zcl.Sprintf("%s", zcl.Zs32(a))
}

const TotalApparentPowerAttr zcl.AttrID = 774

type TotalApparentPower zcl.Zu32

func (a TotalApparentPower) ID() zcl.AttrID              { return TotalApparentPowerAttr }
func (a TotalApparentPower) Cluster() zcl.ClusterID      { return ElectricalMeasurementID }
func (a *TotalApparentPower) Value() *TotalApparentPower { return a }
func (a TotalApparentPower) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *TotalApparentPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = TotalApparentPower(*nt)
	return br, err
}

func (a TotalApparentPower) Readable() bool   { return true }
func (a TotalApparentPower) Writable() bool   { return false }
func (a TotalApparentPower) Reportable() bool { return false }
func (a TotalApparentPower) SceneIndex() int  { return -1 }

func (a TotalApparentPower) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const Measured1StHarmonicCurrentAttr zcl.AttrID = 775

type Measured1StHarmonicCurrent zcl.Zs16

func (a Measured1StHarmonicCurrent) ID() zcl.AttrID                      { return Measured1StHarmonicCurrentAttr }
func (a Measured1StHarmonicCurrent) Cluster() zcl.ClusterID              { return ElectricalMeasurementID }
func (a *Measured1StHarmonicCurrent) Value() *Measured1StHarmonicCurrent { return a }
func (a Measured1StHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured1StHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured1StHarmonicCurrent(*nt)
	return br, err
}

func (a Measured1StHarmonicCurrent) Readable() bool   { return true }
func (a Measured1StHarmonicCurrent) Writable() bool   { return false }
func (a Measured1StHarmonicCurrent) Reportable() bool { return false }
func (a Measured1StHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured1StHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const Measured3RdHarmonicCurrentAttr zcl.AttrID = 776

type Measured3RdHarmonicCurrent zcl.Zs16

func (a Measured3RdHarmonicCurrent) ID() zcl.AttrID                      { return Measured3RdHarmonicCurrentAttr }
func (a Measured3RdHarmonicCurrent) Cluster() zcl.ClusterID              { return ElectricalMeasurementID }
func (a *Measured3RdHarmonicCurrent) Value() *Measured3RdHarmonicCurrent { return a }
func (a Measured3RdHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured3RdHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured3RdHarmonicCurrent(*nt)
	return br, err
}

func (a Measured3RdHarmonicCurrent) Readable() bool   { return true }
func (a Measured3RdHarmonicCurrent) Writable() bool   { return false }
func (a Measured3RdHarmonicCurrent) Reportable() bool { return false }
func (a Measured3RdHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured3RdHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const Measured5ThHarmonicCurrentAttr zcl.AttrID = 777

type Measured5ThHarmonicCurrent zcl.Zs16

func (a Measured5ThHarmonicCurrent) ID() zcl.AttrID                      { return Measured5ThHarmonicCurrentAttr }
func (a Measured5ThHarmonicCurrent) Cluster() zcl.ClusterID              { return ElectricalMeasurementID }
func (a *Measured5ThHarmonicCurrent) Value() *Measured5ThHarmonicCurrent { return a }
func (a Measured5ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured5ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured5ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured5ThHarmonicCurrent) Readable() bool   { return true }
func (a Measured5ThHarmonicCurrent) Writable() bool   { return false }
func (a Measured5ThHarmonicCurrent) Reportable() bool { return false }
func (a Measured5ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured5ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const Measured7ThHarmonicCurrentAttr zcl.AttrID = 778

type Measured7ThHarmonicCurrent zcl.Zs16

func (a Measured7ThHarmonicCurrent) ID() zcl.AttrID                      { return Measured7ThHarmonicCurrentAttr }
func (a Measured7ThHarmonicCurrent) Cluster() zcl.ClusterID              { return ElectricalMeasurementID }
func (a *Measured7ThHarmonicCurrent) Value() *Measured7ThHarmonicCurrent { return a }
func (a Measured7ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured7ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured7ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured7ThHarmonicCurrent) Readable() bool   { return true }
func (a Measured7ThHarmonicCurrent) Writable() bool   { return false }
func (a Measured7ThHarmonicCurrent) Reportable() bool { return false }
func (a Measured7ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured7ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const Measured9ThHarmonicCurrentAttr zcl.AttrID = 779

type Measured9ThHarmonicCurrent zcl.Zs16

func (a Measured9ThHarmonicCurrent) ID() zcl.AttrID                      { return Measured9ThHarmonicCurrentAttr }
func (a Measured9ThHarmonicCurrent) Cluster() zcl.ClusterID              { return ElectricalMeasurementID }
func (a *Measured9ThHarmonicCurrent) Value() *Measured9ThHarmonicCurrent { return a }
func (a Measured9ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured9ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured9ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured9ThHarmonicCurrent) Readable() bool   { return true }
func (a Measured9ThHarmonicCurrent) Writable() bool   { return false }
func (a Measured9ThHarmonicCurrent) Reportable() bool { return false }
func (a Measured9ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured9ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const Measured11ThHarmonicCurrentAttr zcl.AttrID = 780

type Measured11ThHarmonicCurrent zcl.Zs16

func (a Measured11ThHarmonicCurrent) ID() zcl.AttrID                       { return Measured11ThHarmonicCurrentAttr }
func (a Measured11ThHarmonicCurrent) Cluster() zcl.ClusterID               { return ElectricalMeasurementID }
func (a *Measured11ThHarmonicCurrent) Value() *Measured11ThHarmonicCurrent { return a }
func (a Measured11ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *Measured11ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = Measured11ThHarmonicCurrent(*nt)
	return br, err
}

func (a Measured11ThHarmonicCurrent) Readable() bool   { return true }
func (a Measured11ThHarmonicCurrent) Writable() bool   { return false }
func (a Measured11ThHarmonicCurrent) Reportable() bool { return false }
func (a Measured11ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a Measured11ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MeasuredPhase1StHarmonicCurrentAttr zcl.AttrID = 781

type MeasuredPhase1StHarmonicCurrent zcl.Zs16

func (a MeasuredPhase1StHarmonicCurrent) ID() zcl.AttrID                           { return MeasuredPhase1StHarmonicCurrentAttr }
func (a MeasuredPhase1StHarmonicCurrent) Cluster() zcl.ClusterID                   { return ElectricalMeasurementID }
func (a *MeasuredPhase1StHarmonicCurrent) Value() *MeasuredPhase1StHarmonicCurrent { return a }
func (a MeasuredPhase1StHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase1StHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase1StHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase1StHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase1StHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase1StHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase1StHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase1StHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MeasuredPhase3RdHarmonicCurrentAttr zcl.AttrID = 782

type MeasuredPhase3RdHarmonicCurrent zcl.Zs16

func (a MeasuredPhase3RdHarmonicCurrent) ID() zcl.AttrID                           { return MeasuredPhase3RdHarmonicCurrentAttr }
func (a MeasuredPhase3RdHarmonicCurrent) Cluster() zcl.ClusterID                   { return ElectricalMeasurementID }
func (a *MeasuredPhase3RdHarmonicCurrent) Value() *MeasuredPhase3RdHarmonicCurrent { return a }
func (a MeasuredPhase3RdHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase3RdHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase3RdHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase3RdHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase3RdHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase3RdHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase3RdHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase3RdHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MeasuredPhase5ThHarmonicCurrentAttr zcl.AttrID = 783

type MeasuredPhase5ThHarmonicCurrent zcl.Zs16

func (a MeasuredPhase5ThHarmonicCurrent) ID() zcl.AttrID                           { return MeasuredPhase5ThHarmonicCurrentAttr }
func (a MeasuredPhase5ThHarmonicCurrent) Cluster() zcl.ClusterID                   { return ElectricalMeasurementID }
func (a *MeasuredPhase5ThHarmonicCurrent) Value() *MeasuredPhase5ThHarmonicCurrent { return a }
func (a MeasuredPhase5ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase5ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase5ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase5ThHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase5ThHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase5ThHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase5ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase5ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MeasuredPhase7ThHarmonicCurrentAttr zcl.AttrID = 784

type MeasuredPhase7ThHarmonicCurrent zcl.Zs16

func (a MeasuredPhase7ThHarmonicCurrent) ID() zcl.AttrID                           { return MeasuredPhase7ThHarmonicCurrentAttr }
func (a MeasuredPhase7ThHarmonicCurrent) Cluster() zcl.ClusterID                   { return ElectricalMeasurementID }
func (a *MeasuredPhase7ThHarmonicCurrent) Value() *MeasuredPhase7ThHarmonicCurrent { return a }
func (a MeasuredPhase7ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase7ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase7ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase7ThHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase7ThHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase7ThHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase7ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase7ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MeasuredPhase9ThHarmonicCurrentAttr zcl.AttrID = 785

type MeasuredPhase9ThHarmonicCurrent zcl.Zs16

func (a MeasuredPhase9ThHarmonicCurrent) ID() zcl.AttrID                           { return MeasuredPhase9ThHarmonicCurrentAttr }
func (a MeasuredPhase9ThHarmonicCurrent) Cluster() zcl.ClusterID                   { return ElectricalMeasurementID }
func (a *MeasuredPhase9ThHarmonicCurrent) Value() *MeasuredPhase9ThHarmonicCurrent { return a }
func (a MeasuredPhase9ThHarmonicCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MeasuredPhase9ThHarmonicCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MeasuredPhase9ThHarmonicCurrent(*nt)
	return br, err
}

func (a MeasuredPhase9ThHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase9ThHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase9ThHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase9ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase9ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MeasuredPhase11ThHarmonicCurrentAttr zcl.AttrID = 786

type MeasuredPhase11ThHarmonicCurrent zcl.Zs16

func (a MeasuredPhase11ThHarmonicCurrent) ID() zcl.AttrID                            { return MeasuredPhase11ThHarmonicCurrentAttr }
func (a MeasuredPhase11ThHarmonicCurrent) Cluster() zcl.ClusterID                    { return ElectricalMeasurementID }
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

func (a MeasuredPhase11ThHarmonicCurrent) Readable() bool   { return true }
func (a MeasuredPhase11ThHarmonicCurrent) Writable() bool   { return false }
func (a MeasuredPhase11ThHarmonicCurrent) Reportable() bool { return false }
func (a MeasuredPhase11ThHarmonicCurrent) SceneIndex() int  { return -1 }

func (a MeasuredPhase11ThHarmonicCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const AcFrequencyMultiplierAttr zcl.AttrID = 1024

type AcFrequencyMultiplier zcl.Zu16

func (a AcFrequencyMultiplier) ID() zcl.AttrID                 { return AcFrequencyMultiplierAttr }
func (a AcFrequencyMultiplier) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (a *AcFrequencyMultiplier) Value() *AcFrequencyMultiplier { return a }
func (a AcFrequencyMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcFrequencyMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyMultiplier(*nt)
	return br, err
}

func (a AcFrequencyMultiplier) Readable() bool   { return true }
func (a AcFrequencyMultiplier) Writable() bool   { return false }
func (a AcFrequencyMultiplier) Reportable() bool { return false }
func (a AcFrequencyMultiplier) SceneIndex() int  { return -1 }

func (a AcFrequencyMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AcFrequencyDivisorAttr zcl.AttrID = 1025

type AcFrequencyDivisor zcl.Zu16

func (a AcFrequencyDivisor) ID() zcl.AttrID              { return AcFrequencyDivisorAttr }
func (a AcFrequencyDivisor) Cluster() zcl.ClusterID      { return ElectricalMeasurementID }
func (a *AcFrequencyDivisor) Value() *AcFrequencyDivisor { return a }
func (a AcFrequencyDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcFrequencyDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcFrequencyDivisor(*nt)
	return br, err
}

func (a AcFrequencyDivisor) Readable() bool   { return true }
func (a AcFrequencyDivisor) Writable() bool   { return false }
func (a AcFrequencyDivisor) Reportable() bool { return false }
func (a AcFrequencyDivisor) SceneIndex() int  { return -1 }

func (a AcFrequencyDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PowerMultiplierAttr zcl.AttrID = 1026

type PowerMultiplier zcl.Zu32

func (a PowerMultiplier) ID() zcl.AttrID           { return PowerMultiplierAttr }
func (a PowerMultiplier) Cluster() zcl.ClusterID   { return ElectricalMeasurementID }
func (a *PowerMultiplier) Value() *PowerMultiplier { return a }
func (a PowerMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *PowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerMultiplier(*nt)
	return br, err
}

func (a PowerMultiplier) Readable() bool   { return true }
func (a PowerMultiplier) Writable() bool   { return false }
func (a PowerMultiplier) Reportable() bool { return false }
func (a PowerMultiplier) SceneIndex() int  { return -1 }

func (a PowerMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const PowerDivisorAttr zcl.AttrID = 1027

type PowerDivisor zcl.Zu32

func (a PowerDivisor) ID() zcl.AttrID         { return PowerDivisorAttr }
func (a PowerDivisor) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *PowerDivisor) Value() *PowerDivisor  { return a }
func (a PowerDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *PowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerDivisor(*nt)
	return br, err
}

func (a PowerDivisor) Readable() bool   { return true }
func (a PowerDivisor) Writable() bool   { return false }
func (a PowerDivisor) Reportable() bool { return false }
func (a PowerDivisor) SceneIndex() int  { return -1 }

func (a PowerDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

const HarmonicCurrentMultiplierAttr zcl.AttrID = 1028

type HarmonicCurrentMultiplier zcl.Zs8

func (a HarmonicCurrentMultiplier) ID() zcl.AttrID                     { return HarmonicCurrentMultiplierAttr }
func (a HarmonicCurrentMultiplier) Cluster() zcl.ClusterID             { return ElectricalMeasurementID }
func (a *HarmonicCurrentMultiplier) Value() *HarmonicCurrentMultiplier { return a }
func (a HarmonicCurrentMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *HarmonicCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = HarmonicCurrentMultiplier(*nt)
	return br, err
}

func (a HarmonicCurrentMultiplier) Readable() bool   { return true }
func (a HarmonicCurrentMultiplier) Writable() bool   { return false }
func (a HarmonicCurrentMultiplier) Reportable() bool { return false }
func (a HarmonicCurrentMultiplier) SceneIndex() int  { return -1 }

func (a HarmonicCurrentMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}

const PhaseHarmonicCurrentMultiplierAttr zcl.AttrID = 1029

type PhaseHarmonicCurrentMultiplier zcl.Zs8

func (a PhaseHarmonicCurrentMultiplier) ID() zcl.AttrID                          { return PhaseHarmonicCurrentMultiplierAttr }
func (a PhaseHarmonicCurrentMultiplier) Cluster() zcl.ClusterID                  { return ElectricalMeasurementID }
func (a *PhaseHarmonicCurrentMultiplier) Value() *PhaseHarmonicCurrentMultiplier { return a }
func (a PhaseHarmonicCurrentMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *PhaseHarmonicCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PhaseHarmonicCurrentMultiplier(*nt)
	return br, err
}

func (a PhaseHarmonicCurrentMultiplier) Readable() bool   { return true }
func (a PhaseHarmonicCurrentMultiplier) Writable() bool   { return false }
func (a PhaseHarmonicCurrentMultiplier) Reportable() bool { return false }
func (a PhaseHarmonicCurrentMultiplier) SceneIndex() int  { return -1 }

func (a PhaseHarmonicCurrentMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}

const LineCurrentAttr zcl.AttrID = 1281

type LineCurrent zcl.Zu16

func (a LineCurrent) ID() zcl.AttrID         { return LineCurrentAttr }
func (a LineCurrent) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *LineCurrent) Value() *LineCurrent   { return a }
func (a LineCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *LineCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LineCurrent(*nt)
	return br, err
}

func (a LineCurrent) Readable() bool   { return true }
func (a LineCurrent) Writable() bool   { return false }
func (a LineCurrent) Reportable() bool { return false }
func (a LineCurrent) SceneIndex() int  { return -1 }

func (a LineCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ActiveCurrentAttr zcl.AttrID = 1282

type ActiveCurrent zcl.Zs16

func (a ActiveCurrent) ID() zcl.AttrID         { return ActiveCurrentAttr }
func (a ActiveCurrent) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *ActiveCurrent) Value() *ActiveCurrent { return a }
func (a ActiveCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActiveCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveCurrent(*nt)
	return br, err
}

func (a ActiveCurrent) Readable() bool   { return true }
func (a ActiveCurrent) Writable() bool   { return false }
func (a ActiveCurrent) Reportable() bool { return false }
func (a ActiveCurrent) SceneIndex() int  { return -1 }

func (a ActiveCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ReactiveCurrentAttr zcl.AttrID = 1283

type ReactiveCurrent zcl.Zs16

func (a ReactiveCurrent) ID() zcl.AttrID           { return ReactiveCurrentAttr }
func (a ReactiveCurrent) Cluster() zcl.ClusterID   { return ElectricalMeasurementID }
func (a *ReactiveCurrent) Value() *ReactiveCurrent { return a }
func (a ReactiveCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactiveCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactiveCurrent(*nt)
	return br, err
}

func (a ReactiveCurrent) Readable() bool   { return true }
func (a ReactiveCurrent) Writable() bool   { return false }
func (a ReactiveCurrent) Reportable() bool { return false }
func (a ReactiveCurrent) SceneIndex() int  { return -1 }

func (a ReactiveCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const RmsVoltageAttr zcl.AttrID = 1285

type RmsVoltage zcl.Zu16

func (a RmsVoltage) ID() zcl.AttrID         { return RmsVoltageAttr }
func (a RmsVoltage) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *RmsVoltage) Value() *RmsVoltage    { return a }
func (a RmsVoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltage(*nt)
	return br, err
}

func (a RmsVoltage) Readable() bool   { return true }
func (a RmsVoltage) Writable() bool   { return false }
func (a RmsVoltage) Reportable() bool { return false }
func (a RmsVoltage) SceneIndex() int  { return -1 }

func (a RmsVoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageMinAttr zcl.AttrID = 1286

type RmsVoltageMin zcl.Zu16

func (a RmsVoltageMin) ID() zcl.AttrID         { return RmsVoltageMinAttr }
func (a RmsVoltageMin) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *RmsVoltageMin) Value() *RmsVoltageMin { return a }
func (a RmsVoltageMin) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMin(*nt)
	return br, err
}

func (a RmsVoltageMin) Readable() bool   { return true }
func (a RmsVoltageMin) Writable() bool   { return false }
func (a RmsVoltageMin) Reportable() bool { return false }
func (a RmsVoltageMin) SceneIndex() int  { return -1 }

func (a RmsVoltageMin) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageMaxAttr zcl.AttrID = 1287

type RmsVoltageMax zcl.Zu16

func (a RmsVoltageMax) ID() zcl.AttrID         { return RmsVoltageMaxAttr }
func (a RmsVoltageMax) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *RmsVoltageMax) Value() *RmsVoltageMax { return a }
func (a RmsVoltageMax) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMax(*nt)
	return br, err
}

func (a RmsVoltageMax) Readable() bool   { return true }
func (a RmsVoltageMax) Writable() bool   { return false }
func (a RmsVoltageMax) Reportable() bool { return false }
func (a RmsVoltageMax) SceneIndex() int  { return -1 }

func (a RmsVoltageMax) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsCurrentAttr zcl.AttrID = 1288

type RmsCurrent zcl.Zu16

func (a RmsCurrent) ID() zcl.AttrID         { return RmsCurrentAttr }
func (a RmsCurrent) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *RmsCurrent) Value() *RmsCurrent    { return a }
func (a RmsCurrent) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrent) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrent(*nt)
	return br, err
}

func (a RmsCurrent) Readable() bool   { return true }
func (a RmsCurrent) Writable() bool   { return false }
func (a RmsCurrent) Reportable() bool { return false }
func (a RmsCurrent) SceneIndex() int  { return -1 }

func (a RmsCurrent) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsCurrentMinAttr zcl.AttrID = 1289

type RmsCurrentMin zcl.Zu16

func (a RmsCurrentMin) ID() zcl.AttrID         { return RmsCurrentMinAttr }
func (a RmsCurrentMin) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *RmsCurrentMin) Value() *RmsCurrentMin { return a }
func (a RmsCurrentMin) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMin(*nt)
	return br, err
}

func (a RmsCurrentMin) Readable() bool   { return true }
func (a RmsCurrentMin) Writable() bool   { return false }
func (a RmsCurrentMin) Reportable() bool { return false }
func (a RmsCurrentMin) SceneIndex() int  { return -1 }

func (a RmsCurrentMin) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsCurrentMaxAttr zcl.AttrID = 1290

type RmsCurrentMax zcl.Zu16

func (a RmsCurrentMax) ID() zcl.AttrID         { return RmsCurrentMaxAttr }
func (a RmsCurrentMax) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *RmsCurrentMax) Value() *RmsCurrentMax { return a }
func (a RmsCurrentMax) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMax(*nt)
	return br, err
}

func (a RmsCurrentMax) Readable() bool   { return true }
func (a RmsCurrentMax) Writable() bool   { return false }
func (a RmsCurrentMax) Reportable() bool { return false }
func (a RmsCurrentMax) SceneIndex() int  { return -1 }

func (a RmsCurrentMax) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ActivePowerAttr zcl.AttrID = 1291

type ActivePower zcl.Zs16

func (a ActivePower) ID() zcl.AttrID         { return ActivePowerAttr }
func (a ActivePower) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *ActivePower) Value() *ActivePower   { return a }
func (a ActivePower) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePower(*nt)
	return br, err
}

func (a ActivePower) Readable() bool   { return true }
func (a ActivePower) Writable() bool   { return false }
func (a ActivePower) Reportable() bool { return false }
func (a ActivePower) SceneIndex() int  { return -1 }

func (a ActivePower) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ActivePowerMinAttr zcl.AttrID = 1292

type ActivePowerMin zcl.Zs16

func (a ActivePowerMin) ID() zcl.AttrID          { return ActivePowerMinAttr }
func (a ActivePowerMin) Cluster() zcl.ClusterID  { return ElectricalMeasurementID }
func (a *ActivePowerMin) Value() *ActivePowerMin { return a }
func (a ActivePowerMin) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMin(*nt)
	return br, err
}

func (a ActivePowerMin) Readable() bool   { return true }
func (a ActivePowerMin) Writable() bool   { return false }
func (a ActivePowerMin) Reportable() bool { return false }
func (a ActivePowerMin) SceneIndex() int  { return -1 }

func (a ActivePowerMin) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ActivePowerMaxAttr zcl.AttrID = 1293

type ActivePowerMax zcl.Zs16

func (a ActivePowerMax) ID() zcl.AttrID          { return ActivePowerMaxAttr }
func (a ActivePowerMax) Cluster() zcl.ClusterID  { return ElectricalMeasurementID }
func (a *ActivePowerMax) Value() *ActivePowerMax { return a }
func (a ActivePowerMax) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMax(*nt)
	return br, err
}

func (a ActivePowerMax) Readable() bool   { return true }
func (a ActivePowerMax) Writable() bool   { return false }
func (a ActivePowerMax) Reportable() bool { return false }
func (a ActivePowerMax) SceneIndex() int  { return -1 }

func (a ActivePowerMax) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ReactivePowerAttr zcl.AttrID = 1294

type ReactivePower zcl.Zs16

func (a ReactivePower) ID() zcl.AttrID         { return ReactivePowerAttr }
func (a ReactivePower) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *ReactivePower) Value() *ReactivePower { return a }
func (a ReactivePower) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactivePower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactivePower(*nt)
	return br, err
}

func (a ReactivePower) Readable() bool   { return true }
func (a ReactivePower) Writable() bool   { return false }
func (a ReactivePower) Reportable() bool { return false }
func (a ReactivePower) SceneIndex() int  { return -1 }

func (a ReactivePower) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ApparentPowerAttr zcl.AttrID = 1295

type ApparentPower zcl.Zu16

func (a ApparentPower) ID() zcl.AttrID         { return ApparentPowerAttr }
func (a ApparentPower) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *ApparentPower) Value() *ApparentPower { return a }
func (a ApparentPower) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApparentPower) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApparentPower(*nt)
	return br, err
}

func (a ApparentPower) Readable() bool   { return true }
func (a ApparentPower) Writable() bool   { return false }
func (a ApparentPower) Reportable() bool { return false }
func (a ApparentPower) SceneIndex() int  { return -1 }

func (a ApparentPower) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PowerFactorAttr zcl.AttrID = 1296

type PowerFactor zcl.Zs8

func (a PowerFactor) ID() zcl.AttrID         { return PowerFactorAttr }
func (a PowerFactor) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *PowerFactor) Value() *PowerFactor   { return a }
func (a PowerFactor) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *PowerFactor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerFactor(*nt)
	return br, err
}

func (a PowerFactor) Readable() bool   { return true }
func (a PowerFactor) Writable() bool   { return false }
func (a PowerFactor) Reportable() bool { return false }
func (a PowerFactor) SceneIndex() int  { return -1 }

func (a PowerFactor) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}

const AverageRmsVoltageMeasurementPeriodAttr zcl.AttrID = 1297

type AverageRmsVoltageMeasurementPeriod zcl.Zu16

func (a AverageRmsVoltageMeasurementPeriod) ID() zcl.AttrID {
	return AverageRmsVoltageMeasurementPeriodAttr
}
func (a AverageRmsVoltageMeasurementPeriod) Cluster() zcl.ClusterID                      { return ElectricalMeasurementID }
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

func (a AverageRmsVoltageMeasurementPeriod) Readable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriod) Writable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriod) Reportable() bool { return false }
func (a AverageRmsVoltageMeasurementPeriod) SceneIndex() int  { return -1 }

func (a AverageRmsVoltageMeasurementPeriod) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AverageRmsOvervoltageCounterAttr zcl.AttrID = 1298

type AverageRmsOvervoltageCounter zcl.Zu16

func (a AverageRmsOvervoltageCounter) ID() zcl.AttrID                        { return AverageRmsOvervoltageCounterAttr }
func (a AverageRmsOvervoltageCounter) Cluster() zcl.ClusterID                { return ElectricalMeasurementID }
func (a *AverageRmsOvervoltageCounter) Value() *AverageRmsOvervoltageCounter { return a }
func (a AverageRmsOvervoltageCounter) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsOvervoltageCounter) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsOvervoltageCounter(*nt)
	return br, err
}

func (a AverageRmsOvervoltageCounter) Readable() bool   { return true }
func (a AverageRmsOvervoltageCounter) Writable() bool   { return true }
func (a AverageRmsOvervoltageCounter) Reportable() bool { return false }
func (a AverageRmsOvervoltageCounter) SceneIndex() int  { return -1 }

func (a AverageRmsOvervoltageCounter) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AverageRmsUndervoltageCounterAttr zcl.AttrID = 1299

type AverageRmsUndervoltageCounter zcl.Zu16

func (a AverageRmsUndervoltageCounter) ID() zcl.AttrID                         { return AverageRmsUndervoltageCounterAttr }
func (a AverageRmsUndervoltageCounter) Cluster() zcl.ClusterID                 { return ElectricalMeasurementID }
func (a *AverageRmsUndervoltageCounter) Value() *AverageRmsUndervoltageCounter { return a }
func (a AverageRmsUndervoltageCounter) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AverageRmsUndervoltageCounter) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsUndervoltageCounter(*nt)
	return br, err
}

func (a AverageRmsUndervoltageCounter) Readable() bool   { return true }
func (a AverageRmsUndervoltageCounter) Writable() bool   { return true }
func (a AverageRmsUndervoltageCounter) Reportable() bool { return false }
func (a AverageRmsUndervoltageCounter) SceneIndex() int  { return -1 }

func (a AverageRmsUndervoltageCounter) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsExtremeOvervoltagePeriodAttr zcl.AttrID = 1300

type RmsExtremeOvervoltagePeriod zcl.Zu16

func (a RmsExtremeOvervoltagePeriod) ID() zcl.AttrID                       { return RmsExtremeOvervoltagePeriodAttr }
func (a RmsExtremeOvervoltagePeriod) Cluster() zcl.ClusterID               { return ElectricalMeasurementID }
func (a *RmsExtremeOvervoltagePeriod) Value() *RmsExtremeOvervoltagePeriod { return a }
func (a RmsExtremeOvervoltagePeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsExtremeOvervoltagePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeOvervoltagePeriod(*nt)
	return br, err
}

func (a RmsExtremeOvervoltagePeriod) Readable() bool   { return true }
func (a RmsExtremeOvervoltagePeriod) Writable() bool   { return true }
func (a RmsExtremeOvervoltagePeriod) Reportable() bool { return false }
func (a RmsExtremeOvervoltagePeriod) SceneIndex() int  { return -1 }

func (a RmsExtremeOvervoltagePeriod) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsExtremeUndervoltagePeriodAttr zcl.AttrID = 1301

type RmsExtremeUndervoltagePeriod zcl.Zu16

func (a RmsExtremeUndervoltagePeriod) ID() zcl.AttrID                        { return RmsExtremeUndervoltagePeriodAttr }
func (a RmsExtremeUndervoltagePeriod) Cluster() zcl.ClusterID                { return ElectricalMeasurementID }
func (a *RmsExtremeUndervoltagePeriod) Value() *RmsExtremeUndervoltagePeriod { return a }
func (a RmsExtremeUndervoltagePeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsExtremeUndervoltagePeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeUndervoltagePeriod(*nt)
	return br, err
}

func (a RmsExtremeUndervoltagePeriod) Readable() bool   { return true }
func (a RmsExtremeUndervoltagePeriod) Writable() bool   { return true }
func (a RmsExtremeUndervoltagePeriod) Reportable() bool { return false }
func (a RmsExtremeUndervoltagePeriod) SceneIndex() int  { return -1 }

func (a RmsExtremeUndervoltagePeriod) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageSagPeriodAttr zcl.AttrID = 1302

type RmsVoltageSagPeriod zcl.Zu16

func (a RmsVoltageSagPeriod) ID() zcl.AttrID               { return RmsVoltageSagPeriodAttr }
func (a RmsVoltageSagPeriod) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *RmsVoltageSagPeriod) Value() *RmsVoltageSagPeriod { return a }
func (a RmsVoltageSagPeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSagPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSagPeriod(*nt)
	return br, err
}

func (a RmsVoltageSagPeriod) Readable() bool   { return true }
func (a RmsVoltageSagPeriod) Writable() bool   { return true }
func (a RmsVoltageSagPeriod) Reportable() bool { return false }
func (a RmsVoltageSagPeriod) SceneIndex() int  { return -1 }

func (a RmsVoltageSagPeriod) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageSwellPeriodAttr zcl.AttrID = 1303

type RmsVoltageSwellPeriod zcl.Zu16

func (a RmsVoltageSwellPeriod) ID() zcl.AttrID                 { return RmsVoltageSwellPeriodAttr }
func (a RmsVoltageSwellPeriod) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (a *RmsVoltageSwellPeriod) Value() *RmsVoltageSwellPeriod { return a }
func (a RmsVoltageSwellPeriod) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSwellPeriod) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwellPeriod(*nt)
	return br, err
}

func (a RmsVoltageSwellPeriod) Readable() bool   { return true }
func (a RmsVoltageSwellPeriod) Writable() bool   { return true }
func (a RmsVoltageSwellPeriod) Reportable() bool { return false }
func (a RmsVoltageSwellPeriod) SceneIndex() int  { return -1 }

func (a RmsVoltageSwellPeriod) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AcVoltageMultiplierAttr zcl.AttrID = 1536

type AcVoltageMultiplier zcl.Zu16

func (a AcVoltageMultiplier) ID() zcl.AttrID               { return AcVoltageMultiplierAttr }
func (a AcVoltageMultiplier) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *AcVoltageMultiplier) Value() *AcVoltageMultiplier { return a }
func (a AcVoltageMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcVoltageMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcVoltageMultiplier(*nt)
	return br, err
}

func (a AcVoltageMultiplier) Readable() bool   { return true }
func (a AcVoltageMultiplier) Writable() bool   { return false }
func (a AcVoltageMultiplier) Reportable() bool { return false }
func (a AcVoltageMultiplier) SceneIndex() int  { return -1 }

func (a AcVoltageMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AcVoltageDivisorAttr zcl.AttrID = 1537

type AcVoltageDivisor zcl.Zu16

func (a AcVoltageDivisor) ID() zcl.AttrID            { return AcVoltageDivisorAttr }
func (a AcVoltageDivisor) Cluster() zcl.ClusterID    { return ElectricalMeasurementID }
func (a *AcVoltageDivisor) Value() *AcVoltageDivisor { return a }
func (a AcVoltageDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcVoltageDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcVoltageDivisor(*nt)
	return br, err
}

func (a AcVoltageDivisor) Readable() bool   { return true }
func (a AcVoltageDivisor) Writable() bool   { return false }
func (a AcVoltageDivisor) Reportable() bool { return false }
func (a AcVoltageDivisor) SceneIndex() int  { return -1 }

func (a AcVoltageDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AcCurrentMultiplierAttr zcl.AttrID = 1538

type AcCurrentMultiplier zcl.Zu16

func (a AcCurrentMultiplier) ID() zcl.AttrID               { return AcCurrentMultiplierAttr }
func (a AcCurrentMultiplier) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *AcCurrentMultiplier) Value() *AcCurrentMultiplier { return a }
func (a AcCurrentMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcCurrentMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcCurrentMultiplier(*nt)
	return br, err
}

func (a AcCurrentMultiplier) Readable() bool   { return true }
func (a AcCurrentMultiplier) Writable() bool   { return false }
func (a AcCurrentMultiplier) Reportable() bool { return false }
func (a AcCurrentMultiplier) SceneIndex() int  { return -1 }

func (a AcCurrentMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AcCurrentDivisorAttr zcl.AttrID = 1539

type AcCurrentDivisor zcl.Zu16

func (a AcCurrentDivisor) ID() zcl.AttrID            { return AcCurrentDivisorAttr }
func (a AcCurrentDivisor) Cluster() zcl.ClusterID    { return ElectricalMeasurementID }
func (a *AcCurrentDivisor) Value() *AcCurrentDivisor { return a }
func (a AcCurrentDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcCurrentDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcCurrentDivisor(*nt)
	return br, err
}

func (a AcCurrentDivisor) Readable() bool   { return true }
func (a AcCurrentDivisor) Writable() bool   { return false }
func (a AcCurrentDivisor) Reportable() bool { return false }
func (a AcCurrentDivisor) SceneIndex() int  { return -1 }

func (a AcCurrentDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AcPowerMultiplierAttr zcl.AttrID = 1540

type AcPowerMultiplier zcl.Zu16

func (a AcPowerMultiplier) ID() zcl.AttrID             { return AcPowerMultiplierAttr }
func (a AcPowerMultiplier) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *AcPowerMultiplier) Value() *AcPowerMultiplier { return a }
func (a AcPowerMultiplier) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcPowerMultiplier) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcPowerMultiplier(*nt)
	return br, err
}

func (a AcPowerMultiplier) Readable() bool   { return true }
func (a AcPowerMultiplier) Writable() bool   { return false }
func (a AcPowerMultiplier) Reportable() bool { return false }
func (a AcPowerMultiplier) SceneIndex() int  { return -1 }

func (a AcPowerMultiplier) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AcPowerDivisorAttr zcl.AttrID = 1541

type AcPowerDivisor zcl.Zu16

func (a AcPowerDivisor) ID() zcl.AttrID          { return AcPowerDivisorAttr }
func (a AcPowerDivisor) Cluster() zcl.ClusterID  { return ElectricalMeasurementID }
func (a *AcPowerDivisor) Value() *AcPowerDivisor { return a }
func (a AcPowerDivisor) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AcPowerDivisor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcPowerDivisor(*nt)
	return br, err
}

func (a AcPowerDivisor) Readable() bool   { return true }
func (a AcPowerDivisor) Writable() bool   { return false }
func (a AcPowerDivisor) Reportable() bool { return false }
func (a AcPowerDivisor) SceneIndex() int  { return -1 }

func (a AcPowerDivisor) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const DcOverloadAlarmsMaskAttr zcl.AttrID = 1792

type DcOverloadAlarmsMask zcl.Zbmp8

func (a DcOverloadAlarmsMask) ID() zcl.AttrID                { return DcOverloadAlarmsMaskAttr }
func (a DcOverloadAlarmsMask) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (a *DcOverloadAlarmsMask) Value() *DcOverloadAlarmsMask { return a }
func (a DcOverloadAlarmsMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *DcOverloadAlarmsMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = DcOverloadAlarmsMask(*nt)
	return br, err
}

func (a DcOverloadAlarmsMask) Readable() bool   { return true }
func (a DcOverloadAlarmsMask) Writable() bool   { return true }
func (a DcOverloadAlarmsMask) Reportable() bool { return false }
func (a DcOverloadAlarmsMask) SceneIndex() int  { return -1 }

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

const DcVoltageOverloadAttr zcl.AttrID = 1793

type DcVoltageOverload zcl.Zs16

func (a DcVoltageOverload) ID() zcl.AttrID             { return DcVoltageOverloadAttr }
func (a DcVoltageOverload) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *DcVoltageOverload) Value() *DcVoltageOverload { return a }
func (a DcVoltageOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcVoltageOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcVoltageOverload(*nt)
	return br, err
}

func (a DcVoltageOverload) Readable() bool   { return true }
func (a DcVoltageOverload) Writable() bool   { return false }
func (a DcVoltageOverload) Reportable() bool { return false }
func (a DcVoltageOverload) SceneIndex() int  { return -1 }

func (a DcVoltageOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const DcCurrentOverloadAttr zcl.AttrID = 1794

type DcCurrentOverload zcl.Zs16

func (a DcCurrentOverload) ID() zcl.AttrID             { return DcCurrentOverloadAttr }
func (a DcCurrentOverload) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *DcCurrentOverload) Value() *DcCurrentOverload { return a }
func (a DcCurrentOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *DcCurrentOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = DcCurrentOverload(*nt)
	return br, err
}

func (a DcCurrentOverload) Readable() bool   { return true }
func (a DcCurrentOverload) Writable() bool   { return false }
func (a DcCurrentOverload) Reportable() bool { return false }
func (a DcCurrentOverload) SceneIndex() int  { return -1 }

func (a DcCurrentOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const AcOverloadAlarmsMaskAttr zcl.AttrID = 2048

type AcOverloadAlarmsMask zcl.Zbmp16

func (a AcOverloadAlarmsMask) ID() zcl.AttrID                { return AcOverloadAlarmsMaskAttr }
func (a AcOverloadAlarmsMask) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (a *AcOverloadAlarmsMask) Value() *AcOverloadAlarmsMask { return a }
func (a AcOverloadAlarmsMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp16(a).MarshalZcl()
}
func (a *AcOverloadAlarmsMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcOverloadAlarmsMask(*nt)
	return br, err
}

func (a AcOverloadAlarmsMask) Readable() bool   { return true }
func (a AcOverloadAlarmsMask) Writable() bool   { return true }
func (a AcOverloadAlarmsMask) Reportable() bool { return false }
func (a AcOverloadAlarmsMask) SceneIndex() int  { return -1 }

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

const AcVoltageOverloadAttr zcl.AttrID = 2049

type AcVoltageOverload zcl.Zs16

func (a AcVoltageOverload) ID() zcl.AttrID             { return AcVoltageOverloadAttr }
func (a AcVoltageOverload) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *AcVoltageOverload) Value() *AcVoltageOverload { return a }
func (a AcVoltageOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AcVoltageOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcVoltageOverload(*nt)
	return br, err
}

func (a AcVoltageOverload) Readable() bool   { return true }
func (a AcVoltageOverload) Writable() bool   { return false }
func (a AcVoltageOverload) Reportable() bool { return false }
func (a AcVoltageOverload) SceneIndex() int  { return -1 }

func (a AcVoltageOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const AcCurrentOverloadAttr zcl.AttrID = 2050

type AcCurrentOverload zcl.Zs16

func (a AcCurrentOverload) ID() zcl.AttrID             { return AcCurrentOverloadAttr }
func (a AcCurrentOverload) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *AcCurrentOverload) Value() *AcCurrentOverload { return a }
func (a AcCurrentOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AcCurrentOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcCurrentOverload(*nt)
	return br, err
}

func (a AcCurrentOverload) Readable() bool   { return true }
func (a AcCurrentOverload) Writable() bool   { return false }
func (a AcCurrentOverload) Reportable() bool { return false }
func (a AcCurrentOverload) SceneIndex() int  { return -1 }

func (a AcCurrentOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const AcActivePowerOverloadAttr zcl.AttrID = 2051

type AcActivePowerOverload zcl.Zs16

func (a AcActivePowerOverload) ID() zcl.AttrID                 { return AcActivePowerOverloadAttr }
func (a AcActivePowerOverload) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (a *AcActivePowerOverload) Value() *AcActivePowerOverload { return a }
func (a AcActivePowerOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AcActivePowerOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcActivePowerOverload(*nt)
	return br, err
}

func (a AcActivePowerOverload) Readable() bool   { return true }
func (a AcActivePowerOverload) Writable() bool   { return false }
func (a AcActivePowerOverload) Reportable() bool { return false }
func (a AcActivePowerOverload) SceneIndex() int  { return -1 }

func (a AcActivePowerOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const AcReactivePowerOverloadAttr zcl.AttrID = 2052

type AcReactivePowerOverload zcl.Zs16

func (a AcReactivePowerOverload) ID() zcl.AttrID                   { return AcReactivePowerOverloadAttr }
func (a AcReactivePowerOverload) Cluster() zcl.ClusterID           { return ElectricalMeasurementID }
func (a *AcReactivePowerOverload) Value() *AcReactivePowerOverload { return a }
func (a AcReactivePowerOverload) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AcReactivePowerOverload) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AcReactivePowerOverload(*nt)
	return br, err
}

func (a AcReactivePowerOverload) Readable() bool   { return true }
func (a AcReactivePowerOverload) Writable() bool   { return false }
func (a AcReactivePowerOverload) Reportable() bool { return false }
func (a AcReactivePowerOverload) SceneIndex() int  { return -1 }

func (a AcReactivePowerOverload) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const AverageRmsOvervoltageAttr zcl.AttrID = 2053

type AverageRmsOvervoltage zcl.Zs16

func (a AverageRmsOvervoltage) ID() zcl.AttrID                 { return AverageRmsOvervoltageAttr }
func (a AverageRmsOvervoltage) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (a *AverageRmsOvervoltage) Value() *AverageRmsOvervoltage { return a }
func (a AverageRmsOvervoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AverageRmsOvervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsOvervoltage(*nt)
	return br, err
}

func (a AverageRmsOvervoltage) Readable() bool   { return true }
func (a AverageRmsOvervoltage) Writable() bool   { return false }
func (a AverageRmsOvervoltage) Reportable() bool { return false }
func (a AverageRmsOvervoltage) SceneIndex() int  { return -1 }

func (a AverageRmsOvervoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const AverageRmsUndervoltageAttr zcl.AttrID = 2054

type AverageRmsUndervoltage zcl.Zs16

func (a AverageRmsUndervoltage) ID() zcl.AttrID                  { return AverageRmsUndervoltageAttr }
func (a AverageRmsUndervoltage) Cluster() zcl.ClusterID          { return ElectricalMeasurementID }
func (a *AverageRmsUndervoltage) Value() *AverageRmsUndervoltage { return a }
func (a AverageRmsUndervoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AverageRmsUndervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AverageRmsUndervoltage(*nt)
	return br, err
}

func (a AverageRmsUndervoltage) Readable() bool   { return true }
func (a AverageRmsUndervoltage) Writable() bool   { return false }
func (a AverageRmsUndervoltage) Reportable() bool { return false }
func (a AverageRmsUndervoltage) SceneIndex() int  { return -1 }

func (a AverageRmsUndervoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const RmsExtremeOvervoltageAttr zcl.AttrID = 2055

type RmsExtremeOvervoltage zcl.Zs16

func (a RmsExtremeOvervoltage) ID() zcl.AttrID                 { return RmsExtremeOvervoltageAttr }
func (a RmsExtremeOvervoltage) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (a *RmsExtremeOvervoltage) Value() *RmsExtremeOvervoltage { return a }
func (a RmsExtremeOvervoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *RmsExtremeOvervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeOvervoltage(*nt)
	return br, err
}

func (a RmsExtremeOvervoltage) Readable() bool   { return true }
func (a RmsExtremeOvervoltage) Writable() bool   { return false }
func (a RmsExtremeOvervoltage) Reportable() bool { return false }
func (a RmsExtremeOvervoltage) SceneIndex() int  { return -1 }

func (a RmsExtremeOvervoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const RmsExtremeUndervoltageAttr zcl.AttrID = 2056

type RmsExtremeUndervoltage zcl.Zs16

func (a RmsExtremeUndervoltage) ID() zcl.AttrID                  { return RmsExtremeUndervoltageAttr }
func (a RmsExtremeUndervoltage) Cluster() zcl.ClusterID          { return ElectricalMeasurementID }
func (a *RmsExtremeUndervoltage) Value() *RmsExtremeUndervoltage { return a }
func (a RmsExtremeUndervoltage) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *RmsExtremeUndervoltage) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsExtremeUndervoltage(*nt)
	return br, err
}

func (a RmsExtremeUndervoltage) Readable() bool   { return true }
func (a RmsExtremeUndervoltage) Writable() bool   { return false }
func (a RmsExtremeUndervoltage) Reportable() bool { return false }
func (a RmsExtremeUndervoltage) SceneIndex() int  { return -1 }

func (a RmsExtremeUndervoltage) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const RmsVoltageSagAttr zcl.AttrID = 2057

type RmsVoltageSag zcl.Zs16

func (a RmsVoltageSag) ID() zcl.AttrID         { return RmsVoltageSagAttr }
func (a RmsVoltageSag) Cluster() zcl.ClusterID { return ElectricalMeasurementID }
func (a *RmsVoltageSag) Value() *RmsVoltageSag { return a }
func (a RmsVoltageSag) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *RmsVoltageSag) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSag(*nt)
	return br, err
}

func (a RmsVoltageSag) Readable() bool   { return true }
func (a RmsVoltageSag) Writable() bool   { return false }
func (a RmsVoltageSag) Reportable() bool { return false }
func (a RmsVoltageSag) SceneIndex() int  { return -1 }

func (a RmsVoltageSag) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const RmsVoltageSwellAttr zcl.AttrID = 2058

type RmsVoltageSwell zcl.Zs16

func (a RmsVoltageSwell) ID() zcl.AttrID           { return RmsVoltageSwellAttr }
func (a RmsVoltageSwell) Cluster() zcl.ClusterID   { return ElectricalMeasurementID }
func (a *RmsVoltageSwell) Value() *RmsVoltageSwell { return a }
func (a RmsVoltageSwell) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *RmsVoltageSwell) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwell(*nt)
	return br, err
}

func (a RmsVoltageSwell) Readable() bool   { return true }
func (a RmsVoltageSwell) Writable() bool   { return false }
func (a RmsVoltageSwell) Reportable() bool { return false }
func (a RmsVoltageSwell) SceneIndex() int  { return -1 }

func (a RmsVoltageSwell) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const LineCurrentPhaseBAttr zcl.AttrID = 2305

type LineCurrentPhaseB zcl.Zu16

func (a LineCurrentPhaseB) ID() zcl.AttrID             { return LineCurrentPhaseBAttr }
func (a LineCurrentPhaseB) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *LineCurrentPhaseB) Value() *LineCurrentPhaseB { return a }
func (a LineCurrentPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *LineCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LineCurrentPhaseB(*nt)
	return br, err
}

func (a LineCurrentPhaseB) Readable() bool   { return true }
func (a LineCurrentPhaseB) Writable() bool   { return false }
func (a LineCurrentPhaseB) Reportable() bool { return false }
func (a LineCurrentPhaseB) SceneIndex() int  { return -1 }

func (a LineCurrentPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ActiveCurrentPhaseBAttr zcl.AttrID = 2306

type ActiveCurrentPhaseB zcl.Zs16

func (a ActiveCurrentPhaseB) ID() zcl.AttrID               { return ActiveCurrentPhaseBAttr }
func (a ActiveCurrentPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *ActiveCurrentPhaseB) Value() *ActiveCurrentPhaseB { return a }
func (a ActiveCurrentPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActiveCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveCurrentPhaseB(*nt)
	return br, err
}

func (a ActiveCurrentPhaseB) Readable() bool   { return true }
func (a ActiveCurrentPhaseB) Writable() bool   { return false }
func (a ActiveCurrentPhaseB) Reportable() bool { return false }
func (a ActiveCurrentPhaseB) SceneIndex() int  { return -1 }

func (a ActiveCurrentPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ReactiveCurrentPhaseBAttr zcl.AttrID = 2307

type ReactiveCurrentPhaseB zcl.Zs16

func (a ReactiveCurrentPhaseB) ID() zcl.AttrID                 { return ReactiveCurrentPhaseBAttr }
func (a ReactiveCurrentPhaseB) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (a *ReactiveCurrentPhaseB) Value() *ReactiveCurrentPhaseB { return a }
func (a ReactiveCurrentPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactiveCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactiveCurrentPhaseB(*nt)
	return br, err
}

func (a ReactiveCurrentPhaseB) Readable() bool   { return true }
func (a ReactiveCurrentPhaseB) Writable() bool   { return false }
func (a ReactiveCurrentPhaseB) Reportable() bool { return false }
func (a ReactiveCurrentPhaseB) SceneIndex() int  { return -1 }

func (a ReactiveCurrentPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const RmsVoltagePhaseBAttr zcl.AttrID = 2309

type RmsVoltagePhaseB zcl.Zu16

func (a RmsVoltagePhaseB) ID() zcl.AttrID            { return RmsVoltagePhaseBAttr }
func (a RmsVoltagePhaseB) Cluster() zcl.ClusterID    { return ElectricalMeasurementID }
func (a *RmsVoltagePhaseB) Value() *RmsVoltagePhaseB { return a }
func (a RmsVoltagePhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltagePhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltagePhaseB(*nt)
	return br, err
}

func (a RmsVoltagePhaseB) Readable() bool   { return true }
func (a RmsVoltagePhaseB) Writable() bool   { return false }
func (a RmsVoltagePhaseB) Reportable() bool { return false }
func (a RmsVoltagePhaseB) SceneIndex() int  { return -1 }

func (a RmsVoltagePhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageMinPhaseBAttr zcl.AttrID = 2310

type RmsVoltageMinPhaseB zcl.Zu16

func (a RmsVoltageMinPhaseB) ID() zcl.AttrID               { return RmsVoltageMinPhaseBAttr }
func (a RmsVoltageMinPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *RmsVoltageMinPhaseB) Value() *RmsVoltageMinPhaseB { return a }
func (a RmsVoltageMinPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMinPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMinPhaseB(*nt)
	return br, err
}

func (a RmsVoltageMinPhaseB) Readable() bool   { return true }
func (a RmsVoltageMinPhaseB) Writable() bool   { return false }
func (a RmsVoltageMinPhaseB) Reportable() bool { return false }
func (a RmsVoltageMinPhaseB) SceneIndex() int  { return -1 }

func (a RmsVoltageMinPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageMaxPhaseBAttr zcl.AttrID = 2311

type RmsVoltageMaxPhaseB zcl.Zu16

func (a RmsVoltageMaxPhaseB) ID() zcl.AttrID               { return RmsVoltageMaxPhaseBAttr }
func (a RmsVoltageMaxPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *RmsVoltageMaxPhaseB) Value() *RmsVoltageMaxPhaseB { return a }
func (a RmsVoltageMaxPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMaxPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMaxPhaseB(*nt)
	return br, err
}

func (a RmsVoltageMaxPhaseB) Readable() bool   { return true }
func (a RmsVoltageMaxPhaseB) Writable() bool   { return false }
func (a RmsVoltageMaxPhaseB) Reportable() bool { return false }
func (a RmsVoltageMaxPhaseB) SceneIndex() int  { return -1 }

func (a RmsVoltageMaxPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsCurrentPhaseBAttr zcl.AttrID = 2312

type RmsCurrentPhaseB zcl.Zu16

func (a RmsCurrentPhaseB) ID() zcl.AttrID            { return RmsCurrentPhaseBAttr }
func (a RmsCurrentPhaseB) Cluster() zcl.ClusterID    { return ElectricalMeasurementID }
func (a *RmsCurrentPhaseB) Value() *RmsCurrentPhaseB { return a }
func (a RmsCurrentPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentPhaseB(*nt)
	return br, err
}

func (a RmsCurrentPhaseB) Readable() bool   { return true }
func (a RmsCurrentPhaseB) Writable() bool   { return false }
func (a RmsCurrentPhaseB) Reportable() bool { return false }
func (a RmsCurrentPhaseB) SceneIndex() int  { return -1 }

func (a RmsCurrentPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsCurrentMinPhaseBAttr zcl.AttrID = 2313

type RmsCurrentMinPhaseB zcl.Zu16

func (a RmsCurrentMinPhaseB) ID() zcl.AttrID               { return RmsCurrentMinPhaseBAttr }
func (a RmsCurrentMinPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *RmsCurrentMinPhaseB) Value() *RmsCurrentMinPhaseB { return a }
func (a RmsCurrentMinPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMinPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMinPhaseB(*nt)
	return br, err
}

func (a RmsCurrentMinPhaseB) Readable() bool   { return true }
func (a RmsCurrentMinPhaseB) Writable() bool   { return false }
func (a RmsCurrentMinPhaseB) Reportable() bool { return false }
func (a RmsCurrentMinPhaseB) SceneIndex() int  { return -1 }

func (a RmsCurrentMinPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsCurrentMaxPhaseBAttr zcl.AttrID = 2314

type RmsCurrentMaxPhaseB zcl.Zu16

func (a RmsCurrentMaxPhaseB) ID() zcl.AttrID               { return RmsCurrentMaxPhaseBAttr }
func (a RmsCurrentMaxPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *RmsCurrentMaxPhaseB) Value() *RmsCurrentMaxPhaseB { return a }
func (a RmsCurrentMaxPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMaxPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMaxPhaseB(*nt)
	return br, err
}

func (a RmsCurrentMaxPhaseB) Readable() bool   { return true }
func (a RmsCurrentMaxPhaseB) Writable() bool   { return false }
func (a RmsCurrentMaxPhaseB) Reportable() bool { return false }
func (a RmsCurrentMaxPhaseB) SceneIndex() int  { return -1 }

func (a RmsCurrentMaxPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ActivePowerPhaseBAttr zcl.AttrID = 2315

type ActivePowerPhaseB zcl.Zs16

func (a ActivePowerPhaseB) ID() zcl.AttrID             { return ActivePowerPhaseBAttr }
func (a ActivePowerPhaseB) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *ActivePowerPhaseB) Value() *ActivePowerPhaseB { return a }
func (a ActivePowerPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerPhaseB(*nt)
	return br, err
}

func (a ActivePowerPhaseB) Readable() bool   { return true }
func (a ActivePowerPhaseB) Writable() bool   { return false }
func (a ActivePowerPhaseB) Reportable() bool { return false }
func (a ActivePowerPhaseB) SceneIndex() int  { return -1 }

func (a ActivePowerPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ActivePowerMinPhaseBAttr zcl.AttrID = 2316

type ActivePowerMinPhaseB zcl.Zs16

func (a ActivePowerMinPhaseB) ID() zcl.AttrID                { return ActivePowerMinPhaseBAttr }
func (a ActivePowerMinPhaseB) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (a *ActivePowerMinPhaseB) Value() *ActivePowerMinPhaseB { return a }
func (a ActivePowerMinPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMinPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMinPhaseB(*nt)
	return br, err
}

func (a ActivePowerMinPhaseB) Readable() bool   { return true }
func (a ActivePowerMinPhaseB) Writable() bool   { return false }
func (a ActivePowerMinPhaseB) Reportable() bool { return false }
func (a ActivePowerMinPhaseB) SceneIndex() int  { return -1 }

func (a ActivePowerMinPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ActivePowerMaxPhaseBAttr zcl.AttrID = 2317

type ActivePowerMaxPhaseB zcl.Zs16

func (a ActivePowerMaxPhaseB) ID() zcl.AttrID                { return ActivePowerMaxPhaseBAttr }
func (a ActivePowerMaxPhaseB) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (a *ActivePowerMaxPhaseB) Value() *ActivePowerMaxPhaseB { return a }
func (a ActivePowerMaxPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMaxPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMaxPhaseB(*nt)
	return br, err
}

func (a ActivePowerMaxPhaseB) Readable() bool   { return true }
func (a ActivePowerMaxPhaseB) Writable() bool   { return false }
func (a ActivePowerMaxPhaseB) Reportable() bool { return false }
func (a ActivePowerMaxPhaseB) SceneIndex() int  { return -1 }

func (a ActivePowerMaxPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ReactivePowerPhaseBAttr zcl.AttrID = 2318

type ReactivePowerPhaseB zcl.Zs16

func (a ReactivePowerPhaseB) ID() zcl.AttrID               { return ReactivePowerPhaseBAttr }
func (a ReactivePowerPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *ReactivePowerPhaseB) Value() *ReactivePowerPhaseB { return a }
func (a ReactivePowerPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactivePowerPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactivePowerPhaseB(*nt)
	return br, err
}

func (a ReactivePowerPhaseB) Readable() bool   { return true }
func (a ReactivePowerPhaseB) Writable() bool   { return false }
func (a ReactivePowerPhaseB) Reportable() bool { return false }
func (a ReactivePowerPhaseB) SceneIndex() int  { return -1 }

func (a ReactivePowerPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ApparentPowerPhaseBAttr zcl.AttrID = 2319

type ApparentPowerPhaseB zcl.Zu16

func (a ApparentPowerPhaseB) ID() zcl.AttrID               { return ApparentPowerPhaseBAttr }
func (a ApparentPowerPhaseB) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *ApparentPowerPhaseB) Value() *ApparentPowerPhaseB { return a }
func (a ApparentPowerPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApparentPowerPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApparentPowerPhaseB(*nt)
	return br, err
}

func (a ApparentPowerPhaseB) Readable() bool   { return true }
func (a ApparentPowerPhaseB) Writable() bool   { return false }
func (a ApparentPowerPhaseB) Reportable() bool { return false }
func (a ApparentPowerPhaseB) SceneIndex() int  { return -1 }

func (a ApparentPowerPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PowerFactorPhaseBAttr zcl.AttrID = 2320

type PowerFactorPhaseB zcl.Zs8

func (a PowerFactorPhaseB) ID() zcl.AttrID             { return PowerFactorPhaseBAttr }
func (a PowerFactorPhaseB) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *PowerFactorPhaseB) Value() *PowerFactorPhaseB { return a }
func (a PowerFactorPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *PowerFactorPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerFactorPhaseB(*nt)
	return br, err
}

func (a PowerFactorPhaseB) Readable() bool   { return true }
func (a PowerFactorPhaseB) Writable() bool   { return false }
func (a PowerFactorPhaseB) Reportable() bool { return false }
func (a PowerFactorPhaseB) SceneIndex() int  { return -1 }

func (a PowerFactorPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}

const AverageRmsVoltageMeasurementPeriodPhaseBAttr zcl.AttrID = 2321

type AverageRmsVoltageMeasurementPeriodPhaseB zcl.Zu16

func (a AverageRmsVoltageMeasurementPeriodPhaseB) ID() zcl.AttrID {
	return AverageRmsVoltageMeasurementPeriodPhaseBAttr
}
func (a AverageRmsVoltageMeasurementPeriodPhaseB) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}
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

func (a AverageRmsVoltageMeasurementPeriodPhaseB) Readable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriodPhaseB) Writable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriodPhaseB) Reportable() bool { return false }
func (a AverageRmsVoltageMeasurementPeriodPhaseB) SceneIndex() int  { return -1 }

func (a AverageRmsVoltageMeasurementPeriodPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AverageRmsOvervoltageCounterPhaseBAttr zcl.AttrID = 2322

type AverageRmsOvervoltageCounterPhaseB zcl.Zu16

func (a AverageRmsOvervoltageCounterPhaseB) ID() zcl.AttrID {
	return AverageRmsOvervoltageCounterPhaseBAttr
}
func (a AverageRmsOvervoltageCounterPhaseB) Cluster() zcl.ClusterID                      { return ElectricalMeasurementID }
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

func (a AverageRmsOvervoltageCounterPhaseB) Readable() bool   { return true }
func (a AverageRmsOvervoltageCounterPhaseB) Writable() bool   { return true }
func (a AverageRmsOvervoltageCounterPhaseB) Reportable() bool { return false }
func (a AverageRmsOvervoltageCounterPhaseB) SceneIndex() int  { return -1 }

func (a AverageRmsOvervoltageCounterPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AverageRmsUndervoltageCounterPhaseBAttr zcl.AttrID = 2323

type AverageRmsUndervoltageCounterPhaseB zcl.Zu16

func (a AverageRmsUndervoltageCounterPhaseB) ID() zcl.AttrID {
	return AverageRmsUndervoltageCounterPhaseBAttr
}
func (a AverageRmsUndervoltageCounterPhaseB) Cluster() zcl.ClusterID                       { return ElectricalMeasurementID }
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

func (a AverageRmsUndervoltageCounterPhaseB) Readable() bool   { return true }
func (a AverageRmsUndervoltageCounterPhaseB) Writable() bool   { return true }
func (a AverageRmsUndervoltageCounterPhaseB) Reportable() bool { return false }
func (a AverageRmsUndervoltageCounterPhaseB) SceneIndex() int  { return -1 }

func (a AverageRmsUndervoltageCounterPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsExtremeOvervoltagePeriodPhaseBAttr zcl.AttrID = 2324

type RmsExtremeOvervoltagePeriodPhaseB zcl.Zu16

func (a RmsExtremeOvervoltagePeriodPhaseB) ID() zcl.AttrID {
	return RmsExtremeOvervoltagePeriodPhaseBAttr
}
func (a RmsExtremeOvervoltagePeriodPhaseB) Cluster() zcl.ClusterID                     { return ElectricalMeasurementID }
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

func (a RmsExtremeOvervoltagePeriodPhaseB) Readable() bool   { return true }
func (a RmsExtremeOvervoltagePeriodPhaseB) Writable() bool   { return true }
func (a RmsExtremeOvervoltagePeriodPhaseB) Reportable() bool { return false }
func (a RmsExtremeOvervoltagePeriodPhaseB) SceneIndex() int  { return -1 }

func (a RmsExtremeOvervoltagePeriodPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsExtremeUndervoltagePeriodPhaseBAttr zcl.AttrID = 2325

type RmsExtremeUndervoltagePeriodPhaseB zcl.Zu16

func (a RmsExtremeUndervoltagePeriodPhaseB) ID() zcl.AttrID {
	return RmsExtremeUndervoltagePeriodPhaseBAttr
}
func (a RmsExtremeUndervoltagePeriodPhaseB) Cluster() zcl.ClusterID                      { return ElectricalMeasurementID }
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

func (a RmsExtremeUndervoltagePeriodPhaseB) Readable() bool   { return true }
func (a RmsExtremeUndervoltagePeriodPhaseB) Writable() bool   { return true }
func (a RmsExtremeUndervoltagePeriodPhaseB) Reportable() bool { return false }
func (a RmsExtremeUndervoltagePeriodPhaseB) SceneIndex() int  { return -1 }

func (a RmsExtremeUndervoltagePeriodPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageSagPeriodPhaseBAttr zcl.AttrID = 2326

type RmsVoltageSagPeriodPhaseB zcl.Zu16

func (a RmsVoltageSagPeriodPhaseB) ID() zcl.AttrID                     { return RmsVoltageSagPeriodPhaseBAttr }
func (a RmsVoltageSagPeriodPhaseB) Cluster() zcl.ClusterID             { return ElectricalMeasurementID }
func (a *RmsVoltageSagPeriodPhaseB) Value() *RmsVoltageSagPeriodPhaseB { return a }
func (a RmsVoltageSagPeriodPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSagPeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSagPeriodPhaseB(*nt)
	return br, err
}

func (a RmsVoltageSagPeriodPhaseB) Readable() bool   { return true }
func (a RmsVoltageSagPeriodPhaseB) Writable() bool   { return true }
func (a RmsVoltageSagPeriodPhaseB) Reportable() bool { return false }
func (a RmsVoltageSagPeriodPhaseB) SceneIndex() int  { return -1 }

func (a RmsVoltageSagPeriodPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageSwellPeriodPhaseBAttr zcl.AttrID = 2327

type RmsVoltageSwellPeriodPhaseB zcl.Zu16

func (a RmsVoltageSwellPeriodPhaseB) ID() zcl.AttrID                       { return RmsVoltageSwellPeriodPhaseBAttr }
func (a RmsVoltageSwellPeriodPhaseB) Cluster() zcl.ClusterID               { return ElectricalMeasurementID }
func (a *RmsVoltageSwellPeriodPhaseB) Value() *RmsVoltageSwellPeriodPhaseB { return a }
func (a RmsVoltageSwellPeriodPhaseB) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSwellPeriodPhaseB) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwellPeriodPhaseB(*nt)
	return br, err
}

func (a RmsVoltageSwellPeriodPhaseB) Readable() bool   { return true }
func (a RmsVoltageSwellPeriodPhaseB) Writable() bool   { return true }
func (a RmsVoltageSwellPeriodPhaseB) Reportable() bool { return false }
func (a RmsVoltageSwellPeriodPhaseB) SceneIndex() int  { return -1 }

func (a RmsVoltageSwellPeriodPhaseB) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const LineCurrentPhaseCAttr zcl.AttrID = 2561

type LineCurrentPhaseC zcl.Zu16

func (a LineCurrentPhaseC) ID() zcl.AttrID             { return LineCurrentPhaseCAttr }
func (a LineCurrentPhaseC) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *LineCurrentPhaseC) Value() *LineCurrentPhaseC { return a }
func (a LineCurrentPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *LineCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = LineCurrentPhaseC(*nt)
	return br, err
}

func (a LineCurrentPhaseC) Readable() bool   { return true }
func (a LineCurrentPhaseC) Writable() bool   { return false }
func (a LineCurrentPhaseC) Reportable() bool { return false }
func (a LineCurrentPhaseC) SceneIndex() int  { return -1 }

func (a LineCurrentPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ActiveCurrentPhaseCAttr zcl.AttrID = 2562

type ActiveCurrentPhaseC zcl.Zs16

func (a ActiveCurrentPhaseC) ID() zcl.AttrID               { return ActiveCurrentPhaseCAttr }
func (a ActiveCurrentPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *ActiveCurrentPhaseC) Value() *ActiveCurrentPhaseC { return a }
func (a ActiveCurrentPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActiveCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActiveCurrentPhaseC(*nt)
	return br, err
}

func (a ActiveCurrentPhaseC) Readable() bool   { return true }
func (a ActiveCurrentPhaseC) Writable() bool   { return false }
func (a ActiveCurrentPhaseC) Reportable() bool { return false }
func (a ActiveCurrentPhaseC) SceneIndex() int  { return -1 }

func (a ActiveCurrentPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ReactiveCurrentPhaseCAttr zcl.AttrID = 2563

type ReactiveCurrentPhaseC zcl.Zs16

func (a ReactiveCurrentPhaseC) ID() zcl.AttrID                 { return ReactiveCurrentPhaseCAttr }
func (a ReactiveCurrentPhaseC) Cluster() zcl.ClusterID         { return ElectricalMeasurementID }
func (a *ReactiveCurrentPhaseC) Value() *ReactiveCurrentPhaseC { return a }
func (a ReactiveCurrentPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactiveCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactiveCurrentPhaseC(*nt)
	return br, err
}

func (a ReactiveCurrentPhaseC) Readable() bool   { return true }
func (a ReactiveCurrentPhaseC) Writable() bool   { return false }
func (a ReactiveCurrentPhaseC) Reportable() bool { return false }
func (a ReactiveCurrentPhaseC) SceneIndex() int  { return -1 }

func (a ReactiveCurrentPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const RmsVoltagePhaseCAttr zcl.AttrID = 2565

type RmsVoltagePhaseC zcl.Zu16

func (a RmsVoltagePhaseC) ID() zcl.AttrID            { return RmsVoltagePhaseCAttr }
func (a RmsVoltagePhaseC) Cluster() zcl.ClusterID    { return ElectricalMeasurementID }
func (a *RmsVoltagePhaseC) Value() *RmsVoltagePhaseC { return a }
func (a RmsVoltagePhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltagePhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltagePhaseC(*nt)
	return br, err
}

func (a RmsVoltagePhaseC) Readable() bool   { return true }
func (a RmsVoltagePhaseC) Writable() bool   { return false }
func (a RmsVoltagePhaseC) Reportable() bool { return false }
func (a RmsVoltagePhaseC) SceneIndex() int  { return -1 }

func (a RmsVoltagePhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageMinPhaseCAttr zcl.AttrID = 2566

type RmsVoltageMinPhaseC zcl.Zu16

func (a RmsVoltageMinPhaseC) ID() zcl.AttrID               { return RmsVoltageMinPhaseCAttr }
func (a RmsVoltageMinPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *RmsVoltageMinPhaseC) Value() *RmsVoltageMinPhaseC { return a }
func (a RmsVoltageMinPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMinPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMinPhaseC(*nt)
	return br, err
}

func (a RmsVoltageMinPhaseC) Readable() bool   { return true }
func (a RmsVoltageMinPhaseC) Writable() bool   { return false }
func (a RmsVoltageMinPhaseC) Reportable() bool { return false }
func (a RmsVoltageMinPhaseC) SceneIndex() int  { return -1 }

func (a RmsVoltageMinPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageMaxPhaseCAttr zcl.AttrID = 2567

type RmsVoltageMaxPhaseC zcl.Zu16

func (a RmsVoltageMaxPhaseC) ID() zcl.AttrID               { return RmsVoltageMaxPhaseCAttr }
func (a RmsVoltageMaxPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *RmsVoltageMaxPhaseC) Value() *RmsVoltageMaxPhaseC { return a }
func (a RmsVoltageMaxPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageMaxPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageMaxPhaseC(*nt)
	return br, err
}

func (a RmsVoltageMaxPhaseC) Readable() bool   { return true }
func (a RmsVoltageMaxPhaseC) Writable() bool   { return false }
func (a RmsVoltageMaxPhaseC) Reportable() bool { return false }
func (a RmsVoltageMaxPhaseC) SceneIndex() int  { return -1 }

func (a RmsVoltageMaxPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsCurrentPhaseCAttr zcl.AttrID = 2568

type RmsCurrentPhaseC zcl.Zu16

func (a RmsCurrentPhaseC) ID() zcl.AttrID            { return RmsCurrentPhaseCAttr }
func (a RmsCurrentPhaseC) Cluster() zcl.ClusterID    { return ElectricalMeasurementID }
func (a *RmsCurrentPhaseC) Value() *RmsCurrentPhaseC { return a }
func (a RmsCurrentPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentPhaseC(*nt)
	return br, err
}

func (a RmsCurrentPhaseC) Readable() bool   { return true }
func (a RmsCurrentPhaseC) Writable() bool   { return false }
func (a RmsCurrentPhaseC) Reportable() bool { return false }
func (a RmsCurrentPhaseC) SceneIndex() int  { return -1 }

func (a RmsCurrentPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsCurrentMinPhaseCAttr zcl.AttrID = 2569

type RmsCurrentMinPhaseC zcl.Zu16

func (a RmsCurrentMinPhaseC) ID() zcl.AttrID               { return RmsCurrentMinPhaseCAttr }
func (a RmsCurrentMinPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *RmsCurrentMinPhaseC) Value() *RmsCurrentMinPhaseC { return a }
func (a RmsCurrentMinPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMinPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMinPhaseC(*nt)
	return br, err
}

func (a RmsCurrentMinPhaseC) Readable() bool   { return true }
func (a RmsCurrentMinPhaseC) Writable() bool   { return false }
func (a RmsCurrentMinPhaseC) Reportable() bool { return false }
func (a RmsCurrentMinPhaseC) SceneIndex() int  { return -1 }

func (a RmsCurrentMinPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsCurrentMaxPhaseCAttr zcl.AttrID = 2570

type RmsCurrentMaxPhaseC zcl.Zu16

func (a RmsCurrentMaxPhaseC) ID() zcl.AttrID               { return RmsCurrentMaxPhaseCAttr }
func (a RmsCurrentMaxPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *RmsCurrentMaxPhaseC) Value() *RmsCurrentMaxPhaseC { return a }
func (a RmsCurrentMaxPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsCurrentMaxPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsCurrentMaxPhaseC(*nt)
	return br, err
}

func (a RmsCurrentMaxPhaseC) Readable() bool   { return true }
func (a RmsCurrentMaxPhaseC) Writable() bool   { return false }
func (a RmsCurrentMaxPhaseC) Reportable() bool { return false }
func (a RmsCurrentMaxPhaseC) SceneIndex() int  { return -1 }

func (a RmsCurrentMaxPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ActivePowerPhaseCAttr zcl.AttrID = 2571

type ActivePowerPhaseC zcl.Zs16

func (a ActivePowerPhaseC) ID() zcl.AttrID             { return ActivePowerPhaseCAttr }
func (a ActivePowerPhaseC) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *ActivePowerPhaseC) Value() *ActivePowerPhaseC { return a }
func (a ActivePowerPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerPhaseC(*nt)
	return br, err
}

func (a ActivePowerPhaseC) Readable() bool   { return true }
func (a ActivePowerPhaseC) Writable() bool   { return false }
func (a ActivePowerPhaseC) Reportable() bool { return false }
func (a ActivePowerPhaseC) SceneIndex() int  { return -1 }

func (a ActivePowerPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ActivePowerMinPhaseCAttr zcl.AttrID = 2572

type ActivePowerMinPhaseC zcl.Zs16

func (a ActivePowerMinPhaseC) ID() zcl.AttrID                { return ActivePowerMinPhaseCAttr }
func (a ActivePowerMinPhaseC) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (a *ActivePowerMinPhaseC) Value() *ActivePowerMinPhaseC { return a }
func (a ActivePowerMinPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMinPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMinPhaseC(*nt)
	return br, err
}

func (a ActivePowerMinPhaseC) Readable() bool   { return true }
func (a ActivePowerMinPhaseC) Writable() bool   { return false }
func (a ActivePowerMinPhaseC) Reportable() bool { return false }
func (a ActivePowerMinPhaseC) SceneIndex() int  { return -1 }

func (a ActivePowerMinPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ActivePowerMaxPhaseCAttr zcl.AttrID = 2573

type ActivePowerMaxPhaseC zcl.Zs16

func (a ActivePowerMaxPhaseC) ID() zcl.AttrID                { return ActivePowerMaxPhaseCAttr }
func (a ActivePowerMaxPhaseC) Cluster() zcl.ClusterID        { return ElectricalMeasurementID }
func (a *ActivePowerMaxPhaseC) Value() *ActivePowerMaxPhaseC { return a }
func (a ActivePowerMaxPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ActivePowerMaxPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ActivePowerMaxPhaseC(*nt)
	return br, err
}

func (a ActivePowerMaxPhaseC) Readable() bool   { return true }
func (a ActivePowerMaxPhaseC) Writable() bool   { return false }
func (a ActivePowerMaxPhaseC) Reportable() bool { return false }
func (a ActivePowerMaxPhaseC) SceneIndex() int  { return -1 }

func (a ActivePowerMaxPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ReactivePowerPhaseCAttr zcl.AttrID = 2574

type ReactivePowerPhaseC zcl.Zs16

func (a ReactivePowerPhaseC) ID() zcl.AttrID               { return ReactivePowerPhaseCAttr }
func (a ReactivePowerPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *ReactivePowerPhaseC) Value() *ReactivePowerPhaseC { return a }
func (a ReactivePowerPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *ReactivePowerPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = ReactivePowerPhaseC(*nt)
	return br, err
}

func (a ReactivePowerPhaseC) Readable() bool   { return true }
func (a ReactivePowerPhaseC) Writable() bool   { return false }
func (a ReactivePowerPhaseC) Reportable() bool { return false }
func (a ReactivePowerPhaseC) SceneIndex() int  { return -1 }

func (a ReactivePowerPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const ApparentPowerPhaseCAttr zcl.AttrID = 2575

type ApparentPowerPhaseC zcl.Zu16

func (a ApparentPowerPhaseC) ID() zcl.AttrID               { return ApparentPowerPhaseCAttr }
func (a ApparentPowerPhaseC) Cluster() zcl.ClusterID       { return ElectricalMeasurementID }
func (a *ApparentPowerPhaseC) Value() *ApparentPowerPhaseC { return a }
func (a ApparentPowerPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ApparentPowerPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ApparentPowerPhaseC(*nt)
	return br, err
}

func (a ApparentPowerPhaseC) Readable() bool   { return true }
func (a ApparentPowerPhaseC) Writable() bool   { return false }
func (a ApparentPowerPhaseC) Reportable() bool { return false }
func (a ApparentPowerPhaseC) SceneIndex() int  { return -1 }

func (a ApparentPowerPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PowerFactorPhaseCAttr zcl.AttrID = 2576

type PowerFactorPhaseC zcl.Zs8

func (a PowerFactorPhaseC) ID() zcl.AttrID             { return PowerFactorPhaseCAttr }
func (a PowerFactorPhaseC) Cluster() zcl.ClusterID     { return ElectricalMeasurementID }
func (a *PowerFactorPhaseC) Value() *PowerFactorPhaseC { return a }
func (a PowerFactorPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *PowerFactorPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerFactorPhaseC(*nt)
	return br, err
}

func (a PowerFactorPhaseC) Readable() bool   { return true }
func (a PowerFactorPhaseC) Writable() bool   { return false }
func (a PowerFactorPhaseC) Reportable() bool { return false }
func (a PowerFactorPhaseC) SceneIndex() int  { return -1 }

func (a PowerFactorPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zs8(a))
}

const AverageRmsVoltageMeasurementPeriodPhaseCAttr zcl.AttrID = 2577

type AverageRmsVoltageMeasurementPeriodPhaseC zcl.Zu16

func (a AverageRmsVoltageMeasurementPeriodPhaseC) ID() zcl.AttrID {
	return AverageRmsVoltageMeasurementPeriodPhaseCAttr
}
func (a AverageRmsVoltageMeasurementPeriodPhaseC) Cluster() zcl.ClusterID {
	return ElectricalMeasurementID
}
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

func (a AverageRmsVoltageMeasurementPeriodPhaseC) Readable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriodPhaseC) Writable() bool   { return true }
func (a AverageRmsVoltageMeasurementPeriodPhaseC) Reportable() bool { return false }
func (a AverageRmsVoltageMeasurementPeriodPhaseC) SceneIndex() int  { return -1 }

func (a AverageRmsVoltageMeasurementPeriodPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AverageRmsOvervoltageCounterPhaseCAttr zcl.AttrID = 2578

type AverageRmsOvervoltageCounterPhaseC zcl.Zu16

func (a AverageRmsOvervoltageCounterPhaseC) ID() zcl.AttrID {
	return AverageRmsOvervoltageCounterPhaseCAttr
}
func (a AverageRmsOvervoltageCounterPhaseC) Cluster() zcl.ClusterID                      { return ElectricalMeasurementID }
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

func (a AverageRmsOvervoltageCounterPhaseC) Readable() bool   { return true }
func (a AverageRmsOvervoltageCounterPhaseC) Writable() bool   { return true }
func (a AverageRmsOvervoltageCounterPhaseC) Reportable() bool { return false }
func (a AverageRmsOvervoltageCounterPhaseC) SceneIndex() int  { return -1 }

func (a AverageRmsOvervoltageCounterPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const AverageRmsUndervoltageCounterPhaseCAttr zcl.AttrID = 2579

type AverageRmsUndervoltageCounterPhaseC zcl.Zu16

func (a AverageRmsUndervoltageCounterPhaseC) ID() zcl.AttrID {
	return AverageRmsUndervoltageCounterPhaseCAttr
}
func (a AverageRmsUndervoltageCounterPhaseC) Cluster() zcl.ClusterID                       { return ElectricalMeasurementID }
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

func (a AverageRmsUndervoltageCounterPhaseC) Readable() bool   { return true }
func (a AverageRmsUndervoltageCounterPhaseC) Writable() bool   { return true }
func (a AverageRmsUndervoltageCounterPhaseC) Reportable() bool { return false }
func (a AverageRmsUndervoltageCounterPhaseC) SceneIndex() int  { return -1 }

func (a AverageRmsUndervoltageCounterPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsExtremeOvervoltagePeriodPhaseCAttr zcl.AttrID = 2580

type RmsExtremeOvervoltagePeriodPhaseC zcl.Zu16

func (a RmsExtremeOvervoltagePeriodPhaseC) ID() zcl.AttrID {
	return RmsExtremeOvervoltagePeriodPhaseCAttr
}
func (a RmsExtremeOvervoltagePeriodPhaseC) Cluster() zcl.ClusterID                     { return ElectricalMeasurementID }
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

func (a RmsExtremeOvervoltagePeriodPhaseC) Readable() bool   { return true }
func (a RmsExtremeOvervoltagePeriodPhaseC) Writable() bool   { return true }
func (a RmsExtremeOvervoltagePeriodPhaseC) Reportable() bool { return false }
func (a RmsExtremeOvervoltagePeriodPhaseC) SceneIndex() int  { return -1 }

func (a RmsExtremeOvervoltagePeriodPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsExtremeUndervoltagePeriodPhaseCAttr zcl.AttrID = 2581

type RmsExtremeUndervoltagePeriodPhaseC zcl.Zu16

func (a RmsExtremeUndervoltagePeriodPhaseC) ID() zcl.AttrID {
	return RmsExtremeUndervoltagePeriodPhaseCAttr
}
func (a RmsExtremeUndervoltagePeriodPhaseC) Cluster() zcl.ClusterID                      { return ElectricalMeasurementID }
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

func (a RmsExtremeUndervoltagePeriodPhaseC) Readable() bool   { return true }
func (a RmsExtremeUndervoltagePeriodPhaseC) Writable() bool   { return true }
func (a RmsExtremeUndervoltagePeriodPhaseC) Reportable() bool { return false }
func (a RmsExtremeUndervoltagePeriodPhaseC) SceneIndex() int  { return -1 }

func (a RmsExtremeUndervoltagePeriodPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageSagPeriodPhaseCAttr zcl.AttrID = 2582

type RmsVoltageSagPeriodPhaseC zcl.Zu16

func (a RmsVoltageSagPeriodPhaseC) ID() zcl.AttrID                     { return RmsVoltageSagPeriodPhaseCAttr }
func (a RmsVoltageSagPeriodPhaseC) Cluster() zcl.ClusterID             { return ElectricalMeasurementID }
func (a *RmsVoltageSagPeriodPhaseC) Value() *RmsVoltageSagPeriodPhaseC { return a }
func (a RmsVoltageSagPeriodPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSagPeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSagPeriodPhaseC(*nt)
	return br, err
}

func (a RmsVoltageSagPeriodPhaseC) Readable() bool   { return true }
func (a RmsVoltageSagPeriodPhaseC) Writable() bool   { return true }
func (a RmsVoltageSagPeriodPhaseC) Reportable() bool { return false }
func (a RmsVoltageSagPeriodPhaseC) SceneIndex() int  { return -1 }

func (a RmsVoltageSagPeriodPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const RmsVoltageSwellPeriodPhaseCAttr zcl.AttrID = 2583

type RmsVoltageSwellPeriodPhaseC zcl.Zu16

func (a RmsVoltageSwellPeriodPhaseC) ID() zcl.AttrID                       { return RmsVoltageSwellPeriodPhaseCAttr }
func (a RmsVoltageSwellPeriodPhaseC) Cluster() zcl.ClusterID               { return ElectricalMeasurementID }
func (a *RmsVoltageSwellPeriodPhaseC) Value() *RmsVoltageSwellPeriodPhaseC { return a }
func (a RmsVoltageSwellPeriodPhaseC) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RmsVoltageSwellPeriodPhaseC) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RmsVoltageSwellPeriodPhaseC(*nt)
	return br, err
}

func (a RmsVoltageSwellPeriodPhaseC) Readable() bool   { return true }
func (a RmsVoltageSwellPeriodPhaseC) Writable() bool   { return true }
func (a RmsVoltageSwellPeriodPhaseC) Reportable() bool { return false }
func (a RmsVoltageSwellPeriodPhaseC) SceneIndex() int  { return -1 }

func (a RmsVoltageSwellPeriodPhaseC) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}
