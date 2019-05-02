// Thermostat control cluster attributes and commands.
package hvac

import (
	"neotor.se/zcl"
)

// Thermostat
const ThermostatID zcl.ClusterID = 513

var ThermostatCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		SetpointRaiseLowerCommand: func() zcl.Command { return new(SetpointRaiseLower) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		LocalTemperatureAttr:                   func() zcl.Attr { return new(LocalTemperature) },
		OutdoorTemperatureAttr:                 func() zcl.Attr { return new(OutdoorTemperature) },
		OccupancyAttr:                          func() zcl.Attr { return new(Occupancy) },
		AbsMinHeatSetpointLimitAttr:            func() zcl.Attr { return new(AbsMinHeatSetpointLimit) },
		AbsMaxHeatSetpointLimitAttr:            func() zcl.Attr { return new(AbsMaxHeatSetpointLimit) },
		AbsMinCoolSetpointLimitAttr:            func() zcl.Attr { return new(AbsMinCoolSetpointLimit) },
		AbsMaxCoolSetpointLimitAttr:            func() zcl.Attr { return new(AbsMaxCoolSetpointLimit) },
		PiCoolingDemandAttr:                    func() zcl.Attr { return new(PiCoolingDemand) },
		PiHeatingDemandAttr:                    func() zcl.Attr { return new(PiHeatingDemand) },
		HvacSystemTypeConfigurationAttr:        func() zcl.Attr { return new(HvacSystemTypeConfiguration) },
		LocalTemperatureCalibrationAttr:        func() zcl.Attr { return new(LocalTemperatureCalibration) },
		OccupiedCoolingSetpointAttr:            func() zcl.Attr { return new(OccupiedCoolingSetpoint) },
		OccupiedHeatingSetpointAttr:            func() zcl.Attr { return new(OccupiedHeatingSetpoint) },
		UnoccupiedCoolingSetpointAttr:          func() zcl.Attr { return new(UnoccupiedCoolingSetpoint) },
		UnoccupiedHeatingSetpointAttr:          func() zcl.Attr { return new(UnoccupiedHeatingSetpoint) },
		MinHeatSetpointLimitAttr:               func() zcl.Attr { return new(MinHeatSetpointLimit) },
		MaxHeatSetpointLimitAttr:               func() zcl.Attr { return new(MaxHeatSetpointLimit) },
		MinCoolSetpointLimitAttr:               func() zcl.Attr { return new(MinCoolSetpointLimit) },
		MaxCoolSetpointLimitAttr:               func() zcl.Attr { return new(MaxCoolSetpointLimit) },
		MinSetpointDeadBandAttr:                func() zcl.Attr { return new(MinSetpointDeadBand) },
		RemoteSensingAttr:                      func() zcl.Attr { return new(RemoteSensing) },
		ControlSequenceOfOperationAttr:         func() zcl.Attr { return new(ControlSequenceOfOperation) },
		SystemModeAttr:                         func() zcl.Attr { return new(SystemMode) },
		AlarmMaskAttr:                          func() zcl.Attr { return new(AlarmMask) },
		ThermostatRunningModeAttr:              func() zcl.Attr { return new(ThermostatRunningMode) },
		StartOfWeekAttr:                        func() zcl.Attr { return new(StartOfWeek) },
		NumberOfWeeklyTransitionsAttr:          func() zcl.Attr { return new(NumberOfWeeklyTransitions) },
		NumberOfDailTransitionsAttr:            func() zcl.Attr { return new(NumberOfDailTransitions) },
		TemperatureSetpointHoldAttr:            func() zcl.Attr { return new(TemperatureSetpointHold) },
		TemperatureSetpointHoldDurationAttr:    func() zcl.Attr { return new(TemperatureSetpointHoldDuration) },
		ThermostatProgrammingOperationModeAttr: func() zcl.Attr { return new(ThermostatProgrammingOperationMode) },
		ThermostatRunningStateAttr:             func() zcl.Attr { return new(ThermostatRunningState) },
		TrvModeAttr:                            func() zcl.Attr { return new(TrvMode) },
		SetValvePositionAttr:                   func() zcl.Attr { return new(SetValvePosition) },
		ErrorsAttr:                             func() zcl.Attr { return new(Errors) },
		CurrentTemperatureSetpointAttr:         func() zcl.Attr { return new(CurrentTemperatureSetpoint) },
		HostFlagsAttr:                          func() zcl.Attr { return new(HostFlags) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// This command increases (or decreases) the setpoint(s) by amount, in steps of 0.1°C.
type SetpointRaiseLower struct {
	Mode zcl.Zenum8
	// The amount field is a signed 8-bit integer that specifies the amount the setpoint(s) are to be a increased (or decreased) by, in steps of 0.1°C.
	Amount zcl.Zs8
}

const SetpointRaiseLowerCommand zcl.CommandID = 0

func (v *SetpointRaiseLower) Values() []zcl.Val {
	return []zcl.Val{
		&v.Mode,
		&v.Amount,
	}
}

func (v SetpointRaiseLower) ID() zcl.CommandID {
	return SetpointRaiseLowerCommand
}

func (v SetpointRaiseLower) Cluster() zcl.ClusterID {
	return ThermostatID
}

func (v SetpointRaiseLower) MnfCode() []byte {
	return []byte{}
}

func (v SetpointRaiseLower) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Mode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Amount.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *SetpointRaiseLower) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Mode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Amount).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

const LocalTemperatureAttr zcl.AttrID = 0

type LocalTemperature zcl.Zs16

func (a LocalTemperature) ID() zcl.AttrID            { return LocalTemperatureAttr }
func (a LocalTemperature) Cluster() zcl.ClusterID    { return ThermostatID }
func (a *LocalTemperature) Value() *LocalTemperature { return a }
func (a LocalTemperature) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *LocalTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = LocalTemperature(*nt)
	return br, err
}

func (a LocalTemperature) Readable() bool   { return true }
func (a LocalTemperature) Writable() bool   { return false }
func (a LocalTemperature) Reportable() bool { return false }
func (a LocalTemperature) SceneIndex() int  { return -1 }

func (a LocalTemperature) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const OutdoorTemperatureAttr zcl.AttrID = 1

type OutdoorTemperature zcl.Zs16

func (a OutdoorTemperature) ID() zcl.AttrID              { return OutdoorTemperatureAttr }
func (a OutdoorTemperature) Cluster() zcl.ClusterID      { return ThermostatID }
func (a *OutdoorTemperature) Value() *OutdoorTemperature { return a }
func (a OutdoorTemperature) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *OutdoorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = OutdoorTemperature(*nt)
	return br, err
}

func (a OutdoorTemperature) Readable() bool   { return true }
func (a OutdoorTemperature) Writable() bool   { return false }
func (a OutdoorTemperature) Reportable() bool { return false }
func (a OutdoorTemperature) SceneIndex() int  { return -1 }

func (a OutdoorTemperature) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const OccupancyAttr zcl.AttrID = 2

type Occupancy zcl.Zbmp8

func (a Occupancy) ID() zcl.AttrID         { return OccupancyAttr }
func (a Occupancy) Cluster() zcl.ClusterID { return ThermostatID }
func (a *Occupancy) Value() *Occupancy     { return a }
func (a Occupancy) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *Occupancy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Occupancy(*nt)
	return br, err
}

func (a Occupancy) Readable() bool   { return true }
func (a Occupancy) Writable() bool   { return false }
func (a Occupancy) Reportable() bool { return false }
func (a Occupancy) SceneIndex() int  { return -1 }

func (a Occupancy) String() string {
	return zcl.Sprintf("%s", zcl.Zbmp8(a))
}

const AbsMinHeatSetpointLimitAttr zcl.AttrID = 3

type AbsMinHeatSetpointLimit zcl.Zs16

func (a AbsMinHeatSetpointLimit) ID() zcl.AttrID                   { return AbsMinHeatSetpointLimitAttr }
func (a AbsMinHeatSetpointLimit) Cluster() zcl.ClusterID           { return ThermostatID }
func (a *AbsMinHeatSetpointLimit) Value() *AbsMinHeatSetpointLimit { return a }
func (a AbsMinHeatSetpointLimit) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AbsMinHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AbsMinHeatSetpointLimit(*nt)
	return br, err
}

func (a AbsMinHeatSetpointLimit) Readable() bool   { return true }
func (a AbsMinHeatSetpointLimit) Writable() bool   { return false }
func (a AbsMinHeatSetpointLimit) Reportable() bool { return false }
func (a AbsMinHeatSetpointLimit) SceneIndex() int  { return -1 }

func (a AbsMinHeatSetpointLimit) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const AbsMaxHeatSetpointLimitAttr zcl.AttrID = 4

type AbsMaxHeatSetpointLimit zcl.Zs16

func (a AbsMaxHeatSetpointLimit) ID() zcl.AttrID                   { return AbsMaxHeatSetpointLimitAttr }
func (a AbsMaxHeatSetpointLimit) Cluster() zcl.ClusterID           { return ThermostatID }
func (a *AbsMaxHeatSetpointLimit) Value() *AbsMaxHeatSetpointLimit { return a }
func (a AbsMaxHeatSetpointLimit) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AbsMaxHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AbsMaxHeatSetpointLimit(*nt)
	return br, err
}

func (a AbsMaxHeatSetpointLimit) Readable() bool   { return true }
func (a AbsMaxHeatSetpointLimit) Writable() bool   { return false }
func (a AbsMaxHeatSetpointLimit) Reportable() bool { return false }
func (a AbsMaxHeatSetpointLimit) SceneIndex() int  { return -1 }

func (a AbsMaxHeatSetpointLimit) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const AbsMinCoolSetpointLimitAttr zcl.AttrID = 5

type AbsMinCoolSetpointLimit zcl.Zs16

func (a AbsMinCoolSetpointLimit) ID() zcl.AttrID                   { return AbsMinCoolSetpointLimitAttr }
func (a AbsMinCoolSetpointLimit) Cluster() zcl.ClusterID           { return ThermostatID }
func (a *AbsMinCoolSetpointLimit) Value() *AbsMinCoolSetpointLimit { return a }
func (a AbsMinCoolSetpointLimit) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AbsMinCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AbsMinCoolSetpointLimit(*nt)
	return br, err
}

func (a AbsMinCoolSetpointLimit) Readable() bool   { return true }
func (a AbsMinCoolSetpointLimit) Writable() bool   { return false }
func (a AbsMinCoolSetpointLimit) Reportable() bool { return false }
func (a AbsMinCoolSetpointLimit) SceneIndex() int  { return -1 }

func (a AbsMinCoolSetpointLimit) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const AbsMaxCoolSetpointLimitAttr zcl.AttrID = 6

type AbsMaxCoolSetpointLimit zcl.Zs16

func (a AbsMaxCoolSetpointLimit) ID() zcl.AttrID                   { return AbsMaxCoolSetpointLimitAttr }
func (a AbsMaxCoolSetpointLimit) Cluster() zcl.ClusterID           { return ThermostatID }
func (a *AbsMaxCoolSetpointLimit) Value() *AbsMaxCoolSetpointLimit { return a }
func (a AbsMaxCoolSetpointLimit) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *AbsMaxCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AbsMaxCoolSetpointLimit(*nt)
	return br, err
}

func (a AbsMaxCoolSetpointLimit) Readable() bool   { return true }
func (a AbsMaxCoolSetpointLimit) Writable() bool   { return false }
func (a AbsMaxCoolSetpointLimit) Reportable() bool { return false }
func (a AbsMaxCoolSetpointLimit) SceneIndex() int  { return -1 }

func (a AbsMaxCoolSetpointLimit) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const PiCoolingDemandAttr zcl.AttrID = 7

type PiCoolingDemand zcl.Zu8

func (a PiCoolingDemand) ID() zcl.AttrID           { return PiCoolingDemandAttr }
func (a PiCoolingDemand) Cluster() zcl.ClusterID   { return ThermostatID }
func (a *PiCoolingDemand) Value() *PiCoolingDemand { return a }
func (a PiCoolingDemand) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *PiCoolingDemand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PiCoolingDemand(*nt)
	return br, err
}

func (a PiCoolingDemand) Readable() bool   { return true }
func (a PiCoolingDemand) Writable() bool   { return false }
func (a PiCoolingDemand) Reportable() bool { return false }
func (a PiCoolingDemand) SceneIndex() int  { return -1 }

func (a PiCoolingDemand) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const PiHeatingDemandAttr zcl.AttrID = 8

type PiHeatingDemand zcl.Zu8

func (a PiHeatingDemand) ID() zcl.AttrID           { return PiHeatingDemandAttr }
func (a PiHeatingDemand) Cluster() zcl.ClusterID   { return ThermostatID }
func (a *PiHeatingDemand) Value() *PiHeatingDemand { return a }
func (a PiHeatingDemand) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *PiHeatingDemand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PiHeatingDemand(*nt)
	return br, err
}

func (a PiHeatingDemand) Readable() bool   { return true }
func (a PiHeatingDemand) Writable() bool   { return false }
func (a PiHeatingDemand) Reportable() bool { return false }
func (a PiHeatingDemand) SceneIndex() int  { return -1 }

func (a PiHeatingDemand) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const HvacSystemTypeConfigurationAttr zcl.AttrID = 9

type HvacSystemTypeConfiguration zcl.Zbmp8

func (a HvacSystemTypeConfiguration) ID() zcl.AttrID                       { return HvacSystemTypeConfigurationAttr }
func (a HvacSystemTypeConfiguration) Cluster() zcl.ClusterID               { return ThermostatID }
func (a *HvacSystemTypeConfiguration) Value() *HvacSystemTypeConfiguration { return a }
func (a HvacSystemTypeConfiguration) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *HvacSystemTypeConfiguration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = HvacSystemTypeConfiguration(*nt)
	return br, err
}

func (a HvacSystemTypeConfiguration) Readable() bool   { return true }
func (a HvacSystemTypeConfiguration) Writable() bool   { return true }
func (a HvacSystemTypeConfiguration) Reportable() bool { return false }
func (a HvacSystemTypeConfiguration) SceneIndex() int  { return -1 }

func (a HvacSystemTypeConfiguration) String() string {
	return zcl.Sprintf("%s", zcl.Zbmp8(a))
}

const LocalTemperatureCalibrationAttr zcl.AttrID = 16

type LocalTemperatureCalibration zcl.Zs8

func (a LocalTemperatureCalibration) ID() zcl.AttrID                       { return LocalTemperatureCalibrationAttr }
func (a LocalTemperatureCalibration) Cluster() zcl.ClusterID               { return ThermostatID }
func (a *LocalTemperatureCalibration) Value() *LocalTemperatureCalibration { return a }
func (a LocalTemperatureCalibration) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *LocalTemperatureCalibration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = LocalTemperatureCalibration(*nt)
	return br, err
}

func (a LocalTemperatureCalibration) Readable() bool   { return true }
func (a LocalTemperatureCalibration) Writable() bool   { return true }
func (a LocalTemperatureCalibration) Reportable() bool { return false }
func (a LocalTemperatureCalibration) SceneIndex() int  { return -1 }

func (a LocalTemperatureCalibration) String() string {
	return zcl.Sprintf("%s", zcl.Zs8(a))
}

const OccupiedCoolingSetpointAttr zcl.AttrID = 17

type OccupiedCoolingSetpoint zcl.Zs16

func (a OccupiedCoolingSetpoint) ID() zcl.AttrID                   { return OccupiedCoolingSetpointAttr }
func (a OccupiedCoolingSetpoint) Cluster() zcl.ClusterID           { return ThermostatID }
func (a *OccupiedCoolingSetpoint) Value() *OccupiedCoolingSetpoint { return a }
func (a OccupiedCoolingSetpoint) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *OccupiedCoolingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = OccupiedCoolingSetpoint(*nt)
	return br, err
}

func (a OccupiedCoolingSetpoint) Readable() bool   { return true }
func (a OccupiedCoolingSetpoint) Writable() bool   { return true }
func (a OccupiedCoolingSetpoint) Reportable() bool { return false }
func (a OccupiedCoolingSetpoint) SceneIndex() int  { return -1 }

func (a OccupiedCoolingSetpoint) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const OccupiedHeatingSetpointAttr zcl.AttrID = 18

type OccupiedHeatingSetpoint zcl.Zs16

func (a OccupiedHeatingSetpoint) ID() zcl.AttrID                   { return OccupiedHeatingSetpointAttr }
func (a OccupiedHeatingSetpoint) Cluster() zcl.ClusterID           { return ThermostatID }
func (a *OccupiedHeatingSetpoint) Value() *OccupiedHeatingSetpoint { return a }
func (a OccupiedHeatingSetpoint) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *OccupiedHeatingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = OccupiedHeatingSetpoint(*nt)
	return br, err
}

func (a OccupiedHeatingSetpoint) Readable() bool   { return true }
func (a OccupiedHeatingSetpoint) Writable() bool   { return true }
func (a OccupiedHeatingSetpoint) Reportable() bool { return false }
func (a OccupiedHeatingSetpoint) SceneIndex() int  { return -1 }

func (a OccupiedHeatingSetpoint) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const UnoccupiedCoolingSetpointAttr zcl.AttrID = 19

type UnoccupiedCoolingSetpoint zcl.Zs16

func (a UnoccupiedCoolingSetpoint) ID() zcl.AttrID                     { return UnoccupiedCoolingSetpointAttr }
func (a UnoccupiedCoolingSetpoint) Cluster() zcl.ClusterID             { return ThermostatID }
func (a *UnoccupiedCoolingSetpoint) Value() *UnoccupiedCoolingSetpoint { return a }
func (a UnoccupiedCoolingSetpoint) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *UnoccupiedCoolingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = UnoccupiedCoolingSetpoint(*nt)
	return br, err
}

func (a UnoccupiedCoolingSetpoint) Readable() bool   { return true }
func (a UnoccupiedCoolingSetpoint) Writable() bool   { return true }
func (a UnoccupiedCoolingSetpoint) Reportable() bool { return false }
func (a UnoccupiedCoolingSetpoint) SceneIndex() int  { return -1 }

func (a UnoccupiedCoolingSetpoint) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const UnoccupiedHeatingSetpointAttr zcl.AttrID = 20

type UnoccupiedHeatingSetpoint zcl.Zs16

func (a UnoccupiedHeatingSetpoint) ID() zcl.AttrID                     { return UnoccupiedHeatingSetpointAttr }
func (a UnoccupiedHeatingSetpoint) Cluster() zcl.ClusterID             { return ThermostatID }
func (a *UnoccupiedHeatingSetpoint) Value() *UnoccupiedHeatingSetpoint { return a }
func (a UnoccupiedHeatingSetpoint) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *UnoccupiedHeatingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = UnoccupiedHeatingSetpoint(*nt)
	return br, err
}

func (a UnoccupiedHeatingSetpoint) Readable() bool   { return true }
func (a UnoccupiedHeatingSetpoint) Writable() bool   { return true }
func (a UnoccupiedHeatingSetpoint) Reportable() bool { return false }
func (a UnoccupiedHeatingSetpoint) SceneIndex() int  { return -1 }

func (a UnoccupiedHeatingSetpoint) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MinHeatSetpointLimitAttr zcl.AttrID = 21

type MinHeatSetpointLimit zcl.Zs16

func (a MinHeatSetpointLimit) ID() zcl.AttrID                { return MinHeatSetpointLimitAttr }
func (a MinHeatSetpointLimit) Cluster() zcl.ClusterID        { return ThermostatID }
func (a *MinHeatSetpointLimit) Value() *MinHeatSetpointLimit { return a }
func (a MinHeatSetpointLimit) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MinHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinHeatSetpointLimit(*nt)
	return br, err
}

func (a MinHeatSetpointLimit) Readable() bool   { return true }
func (a MinHeatSetpointLimit) Writable() bool   { return true }
func (a MinHeatSetpointLimit) Reportable() bool { return false }
func (a MinHeatSetpointLimit) SceneIndex() int  { return -1 }

func (a MinHeatSetpointLimit) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MaxHeatSetpointLimitAttr zcl.AttrID = 22

type MaxHeatSetpointLimit zcl.Zs16

func (a MaxHeatSetpointLimit) ID() zcl.AttrID                { return MaxHeatSetpointLimitAttr }
func (a MaxHeatSetpointLimit) Cluster() zcl.ClusterID        { return ThermostatID }
func (a *MaxHeatSetpointLimit) Value() *MaxHeatSetpointLimit { return a }
func (a MaxHeatSetpointLimit) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MaxHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxHeatSetpointLimit(*nt)
	return br, err
}

func (a MaxHeatSetpointLimit) Readable() bool   { return true }
func (a MaxHeatSetpointLimit) Writable() bool   { return true }
func (a MaxHeatSetpointLimit) Reportable() bool { return false }
func (a MaxHeatSetpointLimit) SceneIndex() int  { return -1 }

func (a MaxHeatSetpointLimit) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MinCoolSetpointLimitAttr zcl.AttrID = 23

type MinCoolSetpointLimit zcl.Zs16

func (a MinCoolSetpointLimit) ID() zcl.AttrID                { return MinCoolSetpointLimitAttr }
func (a MinCoolSetpointLimit) Cluster() zcl.ClusterID        { return ThermostatID }
func (a *MinCoolSetpointLimit) Value() *MinCoolSetpointLimit { return a }
func (a MinCoolSetpointLimit) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MinCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinCoolSetpointLimit(*nt)
	return br, err
}

func (a MinCoolSetpointLimit) Readable() bool   { return true }
func (a MinCoolSetpointLimit) Writable() bool   { return true }
func (a MinCoolSetpointLimit) Reportable() bool { return false }
func (a MinCoolSetpointLimit) SceneIndex() int  { return -1 }

func (a MinCoolSetpointLimit) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MaxCoolSetpointLimitAttr zcl.AttrID = 24

type MaxCoolSetpointLimit zcl.Zs16

func (a MaxCoolSetpointLimit) ID() zcl.AttrID                { return MaxCoolSetpointLimitAttr }
func (a MaxCoolSetpointLimit) Cluster() zcl.ClusterID        { return ThermostatID }
func (a *MaxCoolSetpointLimit) Value() *MaxCoolSetpointLimit { return a }
func (a MaxCoolSetpointLimit) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *MaxCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxCoolSetpointLimit(*nt)
	return br, err
}

func (a MaxCoolSetpointLimit) Readable() bool   { return true }
func (a MaxCoolSetpointLimit) Writable() bool   { return true }
func (a MaxCoolSetpointLimit) Reportable() bool { return false }
func (a MaxCoolSetpointLimit) SceneIndex() int  { return -1 }

func (a MaxCoolSetpointLimit) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const MinSetpointDeadBandAttr zcl.AttrID = 25

type MinSetpointDeadBand zcl.Zs8

func (a MinSetpointDeadBand) ID() zcl.AttrID               { return MinSetpointDeadBandAttr }
func (a MinSetpointDeadBand) Cluster() zcl.ClusterID       { return ThermostatID }
func (a *MinSetpointDeadBand) Value() *MinSetpointDeadBand { return a }
func (a MinSetpointDeadBand) MarshalZcl() ([]byte, error) {
	return zcl.Zs8(a).MarshalZcl()
}
func (a *MinSetpointDeadBand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = MinSetpointDeadBand(*nt)
	return br, err
}

func (a MinSetpointDeadBand) Readable() bool   { return true }
func (a MinSetpointDeadBand) Writable() bool   { return true }
func (a MinSetpointDeadBand) Reportable() bool { return false }
func (a MinSetpointDeadBand) SceneIndex() int  { return -1 }

func (a MinSetpointDeadBand) String() string {
	return zcl.Sprintf("%s", zcl.Zs8(a))
}

const RemoteSensingAttr zcl.AttrID = 26

type RemoteSensing zcl.Zbmp8

func (a RemoteSensing) ID() zcl.AttrID         { return RemoteSensingAttr }
func (a RemoteSensing) Cluster() zcl.ClusterID { return ThermostatID }
func (a *RemoteSensing) Value() *RemoteSensing { return a }
func (a RemoteSensing) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *RemoteSensing) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = RemoteSensing(*nt)
	return br, err
}

func (a RemoteSensing) Readable() bool   { return true }
func (a RemoteSensing) Writable() bool   { return true }
func (a RemoteSensing) Reportable() bool { return false }
func (a RemoteSensing) SceneIndex() int  { return -1 }

func (a RemoteSensing) String() string {
	var bstr []string
	if a.IsLocalTemperatureSensedRemotely() {
		bstr = append(bstr, "Local temperature sensed remotely")
	}
	if a.IsOutdoorTemperatureSensedRemotely() {
		bstr = append(bstr, "Outdoor temperature sensed remotely")
	}
	if a.IsOccupancySensedRemotely() {
		bstr = append(bstr, "Occupancy sensed remotely")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a RemoteSensing) IsLocalTemperatureSensedRemotely() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *RemoteSensing) SetLocalTemperatureSensedRemotely(b bool) {
	*a = RemoteSensing(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a RemoteSensing) IsOutdoorTemperatureSensedRemotely() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *RemoteSensing) SetOutdoorTemperatureSensedRemotely(b bool) {
	*a = RemoteSensing(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a RemoteSensing) IsOccupancySensedRemotely() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *RemoteSensing) SetOccupancySensedRemotely(b bool) {
	*a = RemoteSensing(zcl.BitmapSet([]byte(*a), 2, b))
}

const ControlSequenceOfOperationAttr zcl.AttrID = 27

type ControlSequenceOfOperation zcl.Zenum8

func (a ControlSequenceOfOperation) ID() zcl.AttrID                      { return ControlSequenceOfOperationAttr }
func (a ControlSequenceOfOperation) Cluster() zcl.ClusterID              { return ThermostatID }
func (a *ControlSequenceOfOperation) Value() *ControlSequenceOfOperation { return a }
func (a ControlSequenceOfOperation) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *ControlSequenceOfOperation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = ControlSequenceOfOperation(*nt)
	return br, err
}

func (a ControlSequenceOfOperation) Readable() bool   { return true }
func (a ControlSequenceOfOperation) Writable() bool   { return true }
func (a ControlSequenceOfOperation) Reportable() bool { return false }
func (a ControlSequenceOfOperation) SceneIndex() int  { return -1 }

func (a ControlSequenceOfOperation) String() string {
	switch a {
	case 0x00:
		return "Cooling Only"
	case 0x01:
		return "Cooling With Reheat"
	case 0x02:
		return "Heating Only"
	case 0x03:
		return "Heating With Reheat"
	case 0x04:
		return "Cooling and Heating 4-pipes"
	case 0x05:
		return "Cooling and Heating 4-pipes with Reheat"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsCoolingOnly checks if ControlSequenceOfOperation equals the value for Cooling Only (0x00)
func (a ControlSequenceOfOperation) IsCoolingOnly() bool { return a == 0x00 }

// SetCoolingOnly sets ControlSequenceOfOperation to Cooling Only (0x00)
func (a *ControlSequenceOfOperation) SetCoolingOnly() { *a = 0x00 }

// IsCoolingWithReheat checks if ControlSequenceOfOperation equals the value for Cooling With Reheat (0x01)
func (a ControlSequenceOfOperation) IsCoolingWithReheat() bool { return a == 0x01 }

// SetCoolingWithReheat sets ControlSequenceOfOperation to Cooling With Reheat (0x01)
func (a *ControlSequenceOfOperation) SetCoolingWithReheat() { *a = 0x01 }

// IsHeatingOnly checks if ControlSequenceOfOperation equals the value for Heating Only (0x02)
func (a ControlSequenceOfOperation) IsHeatingOnly() bool { return a == 0x02 }

// SetHeatingOnly sets ControlSequenceOfOperation to Heating Only (0x02)
func (a *ControlSequenceOfOperation) SetHeatingOnly() { *a = 0x02 }

// IsHeatingWithReheat checks if ControlSequenceOfOperation equals the value for Heating With Reheat (0x03)
func (a ControlSequenceOfOperation) IsHeatingWithReheat() bool { return a == 0x03 }

// SetHeatingWithReheat sets ControlSequenceOfOperation to Heating With Reheat (0x03)
func (a *ControlSequenceOfOperation) SetHeatingWithReheat() { *a = 0x03 }

// IsCoolingAndHeating4Pipes checks if ControlSequenceOfOperation equals the value for Cooling and Heating 4-pipes (0x04)
func (a ControlSequenceOfOperation) IsCoolingAndHeating4Pipes() bool { return a == 0x04 }

// SetCoolingAndHeating4Pipes sets ControlSequenceOfOperation to Cooling and Heating 4-pipes (0x04)
func (a *ControlSequenceOfOperation) SetCoolingAndHeating4Pipes() { *a = 0x04 }

// IsCoolingAndHeating4PipesWithReheat checks if ControlSequenceOfOperation equals the value for Cooling and Heating 4-pipes with Reheat (0x05)
func (a ControlSequenceOfOperation) IsCoolingAndHeating4PipesWithReheat() bool { return a == 0x05 }

// SetCoolingAndHeating4PipesWithReheat sets ControlSequenceOfOperation to Cooling and Heating 4-pipes with Reheat (0x05)
func (a *ControlSequenceOfOperation) SetCoolingAndHeating4PipesWithReheat() { *a = 0x05 }

const SystemModeAttr zcl.AttrID = 28

type SystemMode zcl.Zenum8

func (a SystemMode) ID() zcl.AttrID         { return SystemModeAttr }
func (a SystemMode) Cluster() zcl.ClusterID { return ThermostatID }
func (a *SystemMode) Value() *SystemMode    { return a }
func (a SystemMode) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *SystemMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = SystemMode(*nt)
	return br, err
}

func (a SystemMode) Readable() bool   { return true }
func (a SystemMode) Writable() bool   { return true }
func (a SystemMode) Reportable() bool { return false }
func (a SystemMode) SceneIndex() int  { return -1 }

func (a SystemMode) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x01:
		return "Auto"
	case 0x03:
		return "Cool"
	case 0x04:
		return "Heat"
	case 0x05:
		return "Emergency heating"
	case 0x06:
		return "Precooling"
	case 0x07:
		return "Fan only"
	case 0x08:
		return "Dry"
	case 0x09:
		return "Sleep"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsOff checks if SystemMode equals the value for Off (0x00)
func (a SystemMode) IsOff() bool { return a == 0x00 }

// SetOff sets SystemMode to Off (0x00)
func (a *SystemMode) SetOff() { *a = 0x00 }

// IsAuto checks if SystemMode equals the value for Auto (0x01)
func (a SystemMode) IsAuto() bool { return a == 0x01 }

// SetAuto sets SystemMode to Auto (0x01)
func (a *SystemMode) SetAuto() { *a = 0x01 }

// IsCool checks if SystemMode equals the value for Cool (0x03)
func (a SystemMode) IsCool() bool { return a == 0x03 }

// SetCool sets SystemMode to Cool (0x03)
func (a *SystemMode) SetCool() { *a = 0x03 }

// IsHeat checks if SystemMode equals the value for Heat (0x04)
func (a SystemMode) IsHeat() bool { return a == 0x04 }

// SetHeat sets SystemMode to Heat (0x04)
func (a *SystemMode) SetHeat() { *a = 0x04 }

// IsEmergencyHeating checks if SystemMode equals the value for Emergency heating (0x05)
func (a SystemMode) IsEmergencyHeating() bool { return a == 0x05 }

// SetEmergencyHeating sets SystemMode to Emergency heating (0x05)
func (a *SystemMode) SetEmergencyHeating() { *a = 0x05 }

// IsPrecooling checks if SystemMode equals the value for Precooling (0x06)
func (a SystemMode) IsPrecooling() bool { return a == 0x06 }

// SetPrecooling sets SystemMode to Precooling (0x06)
func (a *SystemMode) SetPrecooling() { *a = 0x06 }

// IsFanOnly checks if SystemMode equals the value for Fan only (0x07)
func (a SystemMode) IsFanOnly() bool { return a == 0x07 }

// SetFanOnly sets SystemMode to Fan only (0x07)
func (a *SystemMode) SetFanOnly() { *a = 0x07 }

// IsDry checks if SystemMode equals the value for Dry (0x08)
func (a SystemMode) IsDry() bool { return a == 0x08 }

// SetDry sets SystemMode to Dry (0x08)
func (a *SystemMode) SetDry() { *a = 0x08 }

// IsSleep checks if SystemMode equals the value for Sleep (0x09)
func (a SystemMode) IsSleep() bool { return a == 0x09 }

// SetSleep sets SystemMode to Sleep (0x09)
func (a *SystemMode) SetSleep() { *a = 0x09 }

const AlarmMaskAttr zcl.AttrID = 29

type AlarmMask zcl.Zbmp8

func (a AlarmMask) ID() zcl.AttrID         { return AlarmMaskAttr }
func (a AlarmMask) Cluster() zcl.ClusterID { return ThermostatID }
func (a *AlarmMask) Value() *AlarmMask     { return a }
func (a AlarmMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *AlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmMask(*nt)
	return br, err
}

func (a AlarmMask) Readable() bool   { return true }
func (a AlarmMask) Writable() bool   { return false }
func (a AlarmMask) Reportable() bool { return false }
func (a AlarmMask) SceneIndex() int  { return -1 }

func (a AlarmMask) String() string {
	var bstr []string
	if a.IsInitializationFailure() {
		bstr = append(bstr, "Initialization failure")
	}
	if a.IsHardwareFailure() {
		bstr = append(bstr, "Hardware failure")
	}
	if a.IsSelfCalibrationFailure() {
		bstr = append(bstr, "Self-calibration failure")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a AlarmMask) IsInitializationFailure() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *AlarmMask) SetInitializationFailure(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a AlarmMask) IsHardwareFailure() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *AlarmMask) SetHardwareFailure(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a AlarmMask) IsSelfCalibrationFailure() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *AlarmMask) SetSelfCalibrationFailure(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 2, b))
}

const ThermostatRunningModeAttr zcl.AttrID = 30

type ThermostatRunningMode zcl.Zenum8

func (a ThermostatRunningMode) ID() zcl.AttrID                 { return ThermostatRunningModeAttr }
func (a ThermostatRunningMode) Cluster() zcl.ClusterID         { return ThermostatID }
func (a *ThermostatRunningMode) Value() *ThermostatRunningMode { return a }
func (a ThermostatRunningMode) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *ThermostatRunningMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = ThermostatRunningMode(*nt)
	return br, err
}

func (a ThermostatRunningMode) Readable() bool   { return true }
func (a ThermostatRunningMode) Writable() bool   { return false }
func (a ThermostatRunningMode) Reportable() bool { return false }
func (a ThermostatRunningMode) SceneIndex() int  { return -1 }

func (a ThermostatRunningMode) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x03:
		return "Cool"
	case 0x04:
		return "Heat"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsOff checks if ThermostatRunningMode equals the value for Off (0x00)
func (a ThermostatRunningMode) IsOff() bool { return a == 0x00 }

// SetOff sets ThermostatRunningMode to Off (0x00)
func (a *ThermostatRunningMode) SetOff() { *a = 0x00 }

// IsCool checks if ThermostatRunningMode equals the value for Cool (0x03)
func (a ThermostatRunningMode) IsCool() bool { return a == 0x03 }

// SetCool sets ThermostatRunningMode to Cool (0x03)
func (a *ThermostatRunningMode) SetCool() { *a = 0x03 }

// IsHeat checks if ThermostatRunningMode equals the value for Heat (0x04)
func (a ThermostatRunningMode) IsHeat() bool { return a == 0x04 }

// SetHeat sets ThermostatRunningMode to Heat (0x04)
func (a *ThermostatRunningMode) SetHeat() { *a = 0x04 }

const StartOfWeekAttr zcl.AttrID = 32

type StartOfWeek zcl.Zenum8

func (a StartOfWeek) ID() zcl.AttrID         { return StartOfWeekAttr }
func (a StartOfWeek) Cluster() zcl.ClusterID { return ThermostatID }
func (a *StartOfWeek) Value() *StartOfWeek   { return a }
func (a StartOfWeek) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *StartOfWeek) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = StartOfWeek(*nt)
	return br, err
}

func (a StartOfWeek) Readable() bool   { return true }
func (a StartOfWeek) Writable() bool   { return false }
func (a StartOfWeek) Reportable() bool { return false }
func (a StartOfWeek) SceneIndex() int  { return -1 }

func (a StartOfWeek) String() string {
	switch a {
	case 0x00:
		return "Sunday"
	case 0x01:
		return "Monday"
	case 0x02:
		return "Tuesday"
	case 0x03:
		return "Wednesday"
	case 0x04:
		return "Thursday"
	case 0x05:
		return "Friday"
	case 0x06:
		return "Saturday"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsSunday checks if StartOfWeek equals the value for Sunday (0x00)
func (a StartOfWeek) IsSunday() bool { return a == 0x00 }

// SetSunday sets StartOfWeek to Sunday (0x00)
func (a *StartOfWeek) SetSunday() { *a = 0x00 }

// IsMonday checks if StartOfWeek equals the value for Monday (0x01)
func (a StartOfWeek) IsMonday() bool { return a == 0x01 }

// SetMonday sets StartOfWeek to Monday (0x01)
func (a *StartOfWeek) SetMonday() { *a = 0x01 }

// IsTuesday checks if StartOfWeek equals the value for Tuesday (0x02)
func (a StartOfWeek) IsTuesday() bool { return a == 0x02 }

// SetTuesday sets StartOfWeek to Tuesday (0x02)
func (a *StartOfWeek) SetTuesday() { *a = 0x02 }

// IsWednesday checks if StartOfWeek equals the value for Wednesday (0x03)
func (a StartOfWeek) IsWednesday() bool { return a == 0x03 }

// SetWednesday sets StartOfWeek to Wednesday (0x03)
func (a *StartOfWeek) SetWednesday() { *a = 0x03 }

// IsThursday checks if StartOfWeek equals the value for Thursday (0x04)
func (a StartOfWeek) IsThursday() bool { return a == 0x04 }

// SetThursday sets StartOfWeek to Thursday (0x04)
func (a *StartOfWeek) SetThursday() { *a = 0x04 }

// IsFriday checks if StartOfWeek equals the value for Friday (0x05)
func (a StartOfWeek) IsFriday() bool { return a == 0x05 }

// SetFriday sets StartOfWeek to Friday (0x05)
func (a *StartOfWeek) SetFriday() { *a = 0x05 }

// IsSaturday checks if StartOfWeek equals the value for Saturday (0x06)
func (a StartOfWeek) IsSaturday() bool { return a == 0x06 }

// SetSaturday sets StartOfWeek to Saturday (0x06)
func (a *StartOfWeek) SetSaturday() { *a = 0x06 }

const NumberOfWeeklyTransitionsAttr zcl.AttrID = 33

type NumberOfWeeklyTransitions zcl.Zu8

func (a NumberOfWeeklyTransitions) ID() zcl.AttrID                     { return NumberOfWeeklyTransitionsAttr }
func (a NumberOfWeeklyTransitions) Cluster() zcl.ClusterID             { return ThermostatID }
func (a *NumberOfWeeklyTransitions) Value() *NumberOfWeeklyTransitions { return a }
func (a NumberOfWeeklyTransitions) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *NumberOfWeeklyTransitions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfWeeklyTransitions(*nt)
	return br, err
}

func (a NumberOfWeeklyTransitions) Readable() bool   { return true }
func (a NumberOfWeeklyTransitions) Writable() bool   { return false }
func (a NumberOfWeeklyTransitions) Reportable() bool { return false }
func (a NumberOfWeeklyTransitions) SceneIndex() int  { return -1 }

func (a NumberOfWeeklyTransitions) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const NumberOfDailTransitionsAttr zcl.AttrID = 34

type NumberOfDailTransitions zcl.Zu8

func (a NumberOfDailTransitions) ID() zcl.AttrID                   { return NumberOfDailTransitionsAttr }
func (a NumberOfDailTransitions) Cluster() zcl.ClusterID           { return ThermostatID }
func (a *NumberOfDailTransitions) Value() *NumberOfDailTransitions { return a }
func (a NumberOfDailTransitions) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *NumberOfDailTransitions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfDailTransitions(*nt)
	return br, err
}

func (a NumberOfDailTransitions) Readable() bool   { return true }
func (a NumberOfDailTransitions) Writable() bool   { return false }
func (a NumberOfDailTransitions) Reportable() bool { return false }
func (a NumberOfDailTransitions) SceneIndex() int  { return -1 }

func (a NumberOfDailTransitions) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const TemperatureSetpointHoldAttr zcl.AttrID = 35

type TemperatureSetpointHold zcl.Zenum8

func (a TemperatureSetpointHold) ID() zcl.AttrID                   { return TemperatureSetpointHoldAttr }
func (a TemperatureSetpointHold) Cluster() zcl.ClusterID           { return ThermostatID }
func (a *TemperatureSetpointHold) Value() *TemperatureSetpointHold { return a }
func (a TemperatureSetpointHold) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *TemperatureSetpointHold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = TemperatureSetpointHold(*nt)
	return br, err
}

func (a TemperatureSetpointHold) Readable() bool   { return true }
func (a TemperatureSetpointHold) Writable() bool   { return false }
func (a TemperatureSetpointHold) Reportable() bool { return false }
func (a TemperatureSetpointHold) SceneIndex() int  { return -1 }

func (a TemperatureSetpointHold) String() string {
	switch a {
	case 0x00:
		return "Setpoint Hold Off"
	case 0x01:
		return "Setpoint Hold On"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsSetpointHoldOff checks if TemperatureSetpointHold equals the value for Setpoint Hold Off (0x00)
func (a TemperatureSetpointHold) IsSetpointHoldOff() bool { return a == 0x00 }

// SetSetpointHoldOff sets TemperatureSetpointHold to Setpoint Hold Off (0x00)
func (a *TemperatureSetpointHold) SetSetpointHoldOff() { *a = 0x00 }

// IsSetpointHoldOn checks if TemperatureSetpointHold equals the value for Setpoint Hold On (0x01)
func (a TemperatureSetpointHold) IsSetpointHoldOn() bool { return a == 0x01 }

// SetSetpointHoldOn sets TemperatureSetpointHold to Setpoint Hold On (0x01)
func (a *TemperatureSetpointHold) SetSetpointHoldOn() { *a = 0x01 }

const TemperatureSetpointHoldDurationAttr zcl.AttrID = 36

type TemperatureSetpointHoldDuration zcl.Zu8

func (a TemperatureSetpointHoldDuration) ID() zcl.AttrID                           { return TemperatureSetpointHoldDurationAttr }
func (a TemperatureSetpointHoldDuration) Cluster() zcl.ClusterID                   { return ThermostatID }
func (a *TemperatureSetpointHoldDuration) Value() *TemperatureSetpointHoldDuration { return a }
func (a TemperatureSetpointHoldDuration) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *TemperatureSetpointHoldDuration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = TemperatureSetpointHoldDuration(*nt)
	return br, err
}

func (a TemperatureSetpointHoldDuration) Readable() bool   { return true }
func (a TemperatureSetpointHoldDuration) Writable() bool   { return true }
func (a TemperatureSetpointHoldDuration) Reportable() bool { return false }
func (a TemperatureSetpointHoldDuration) SceneIndex() int  { return -1 }

func (a TemperatureSetpointHoldDuration) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const ThermostatProgrammingOperationModeAttr zcl.AttrID = 37

type ThermostatProgrammingOperationMode zcl.Zbmp8

func (a ThermostatProgrammingOperationMode) ID() zcl.AttrID {
	return ThermostatProgrammingOperationModeAttr
}
func (a ThermostatProgrammingOperationMode) Cluster() zcl.ClusterID                      { return ThermostatID }
func (a *ThermostatProgrammingOperationMode) Value() *ThermostatProgrammingOperationMode { return a }
func (a ThermostatProgrammingOperationMode) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *ThermostatProgrammingOperationMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = ThermostatProgrammingOperationMode(*nt)
	return br, err
}

func (a ThermostatProgrammingOperationMode) Readable() bool   { return true }
func (a ThermostatProgrammingOperationMode) Writable() bool   { return true }
func (a ThermostatProgrammingOperationMode) Reportable() bool { return false }
func (a ThermostatProgrammingOperationMode) SceneIndex() int  { return -1 }

func (a ThermostatProgrammingOperationMode) String() string {
	var bstr []string
	if a.IsSimpleSetpointMode() {
		bstr = append(bstr, "Simple/setpoint mode")
	}
	if a.IsAutoRecoveryMode() {
		bstr = append(bstr, "Auto/recovery mode")
	}
	if a.IsEconomyEnergystarMode() {
		bstr = append(bstr, "Economy/EnergyStar mode")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a ThermostatProgrammingOperationMode) IsSimpleSetpointMode() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *ThermostatProgrammingOperationMode) SetSimpleSetpointMode(b bool) {
	*a = ThermostatProgrammingOperationMode(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a ThermostatProgrammingOperationMode) IsAutoRecoveryMode() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *ThermostatProgrammingOperationMode) SetAutoRecoveryMode(b bool) {
	*a = ThermostatProgrammingOperationMode(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a ThermostatProgrammingOperationMode) IsEconomyEnergystarMode() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *ThermostatProgrammingOperationMode) SetEconomyEnergystarMode(b bool) {
	*a = ThermostatProgrammingOperationMode(zcl.BitmapSet([]byte(*a), 2, b))
}

const ThermostatRunningStateAttr zcl.AttrID = 41

type ThermostatRunningState zcl.Zbmp16

func (a ThermostatRunningState) ID() zcl.AttrID                  { return ThermostatRunningStateAttr }
func (a ThermostatRunningState) Cluster() zcl.ClusterID          { return ThermostatID }
func (a *ThermostatRunningState) Value() *ThermostatRunningState { return a }
func (a ThermostatRunningState) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp16(a).MarshalZcl()
}
func (a *ThermostatRunningState) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = ThermostatRunningState(*nt)
	return br, err
}

func (a ThermostatRunningState) Readable() bool   { return true }
func (a ThermostatRunningState) Writable() bool   { return false }
func (a ThermostatRunningState) Reportable() bool { return false }
func (a ThermostatRunningState) SceneIndex() int  { return -1 }

func (a ThermostatRunningState) String() string {
	var bstr []string
	if a.IsHeatStateOn() {
		bstr = append(bstr, "Heat State On")
	}
	if a.IsCoolStateOn() {
		bstr = append(bstr, "Cool State On")
	}
	if a.IsFanStateOn() {
		bstr = append(bstr, "Fan State On")
	}
	if a.IsHeat2NdStageStateOn() {
		bstr = append(bstr, "Heat 2nd Stage State On")
	}
	if a.IsCool2NdStageStateOn() {
		bstr = append(bstr, "Cool 2nd Stage State On")
	}
	if a.IsFan2NdStageStateOn() {
		bstr = append(bstr, "Fan 2nd Stage State On")
	}
	if a.IsFan3RdStageStateOn() {
		bstr = append(bstr, "Fan 3rd Stage State On")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a ThermostatRunningState) IsHeatStateOn() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *ThermostatRunningState) SetHeatStateOn(b bool) {
	*a = ThermostatRunningState(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a ThermostatRunningState) IsCoolStateOn() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *ThermostatRunningState) SetCoolStateOn(b bool) {
	*a = ThermostatRunningState(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a ThermostatRunningState) IsFanStateOn() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *ThermostatRunningState) SetFanStateOn(b bool) {
	*a = ThermostatRunningState(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a ThermostatRunningState) IsHeat2NdStageStateOn() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *ThermostatRunningState) SetHeat2NdStageStateOn(b bool) {
	*a = ThermostatRunningState(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a ThermostatRunningState) IsCool2NdStageStateOn() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *ThermostatRunningState) SetCool2NdStageStateOn(b bool) {
	*a = ThermostatRunningState(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a ThermostatRunningState) IsFan2NdStageStateOn() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *ThermostatRunningState) SetFan2NdStageStateOn(b bool) {
	*a = ThermostatRunningState(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a ThermostatRunningState) IsFan3RdStageStateOn() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *ThermostatRunningState) SetFan3RdStageStateOn(b bool) {
	*a = ThermostatRunningState(zcl.BitmapSet([]byte(*a), 6, b))
}

const TrvModeAttr zcl.AttrID = 16384

type TrvMode zcl.Zenum8

func (a TrvMode) ID() zcl.AttrID         { return TrvModeAttr }
func (a TrvMode) Cluster() zcl.ClusterID { return ThermostatID }
func (a *TrvMode) Value() *TrvMode       { return a }
func (a TrvMode) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *TrvMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = TrvMode(*nt)
	return br, err
}

func (a TrvMode) Readable() bool   { return true }
func (a TrvMode) Writable() bool   { return true }
func (a TrvMode) Reportable() bool { return false }
func (a TrvMode) SceneIndex() int  { return -1 }

func (a TrvMode) String() string {
	switch a {
	case 0x00:
		return "Unknown 1"
	case 0x01:
		return "Unknown 2"
	case 0x02:
		return "manual"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsUnknown1 checks if TrvMode equals the value for Unknown 1 (0x00)
func (a TrvMode) IsUnknown1() bool { return a == 0x00 }

// SetUnknown1 sets TrvMode to Unknown 1 (0x00)
func (a *TrvMode) SetUnknown1() { *a = 0x00 }

// IsUnknown2 checks if TrvMode equals the value for Unknown 2 (0x01)
func (a TrvMode) IsUnknown2() bool { return a == 0x01 }

// SetUnknown2 sets TrvMode to Unknown 2 (0x01)
func (a *TrvMode) SetUnknown2() { *a = 0x01 }

// IsManual checks if TrvMode equals the value for manual (0x02)
func (a TrvMode) IsManual() bool { return a == 0x02 }

// SetManual sets TrvMode to manual (0x02)
func (a *TrvMode) SetManual() { *a = 0x02 }

const SetValvePositionAttr zcl.AttrID = 16385

type SetValvePosition zcl.Zu8

func (a SetValvePosition) ID() zcl.AttrID            { return SetValvePositionAttr }
func (a SetValvePosition) Cluster() zcl.ClusterID    { return ThermostatID }
func (a *SetValvePosition) Value() *SetValvePosition { return a }
func (a SetValvePosition) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *SetValvePosition) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SetValvePosition(*nt)
	return br, err
}

func (a SetValvePosition) Readable() bool   { return true }
func (a SetValvePosition) Writable() bool   { return true }
func (a SetValvePosition) Reportable() bool { return false }
func (a SetValvePosition) SceneIndex() int  { return -1 }

func (a SetValvePosition) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const ErrorsAttr zcl.AttrID = 16386

type Errors zcl.Zu8

func (a Errors) ID() zcl.AttrID         { return ErrorsAttr }
func (a Errors) Cluster() zcl.ClusterID { return ThermostatID }
func (a *Errors) Value() *Errors        { return a }
func (a Errors) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Errors) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Errors(*nt)
	return br, err
}

func (a Errors) Readable() bool   { return true }
func (a Errors) Writable() bool   { return false }
func (a Errors) Reportable() bool { return false }
func (a Errors) SceneIndex() int  { return -1 }

func (a Errors) String() string {
	switch a {
	case 0x03:
		return "Valve Adaption Failed (E1)"
	case 0x04:
		return "Valve Movement too slow (E2)"
	case 0x05:
		return "Valve not Moving (E3)"
	}
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

// IsValveAdaptionFailedE1 checks if Errors equals the value for Valve Adaption Failed (E1) (0x03)
func (a Errors) IsValveAdaptionFailedE1() bool { return a == 0x03 }

// SetValveAdaptionFailedE1 sets Errors to Valve Adaption Failed (E1) (0x03)
func (a *Errors) SetValveAdaptionFailedE1() { *a = 0x03 }

// IsValveMovementTooSlowE2 checks if Errors equals the value for Valve Movement too slow (E2) (0x04)
func (a Errors) IsValveMovementTooSlowE2() bool { return a == 0x04 }

// SetValveMovementTooSlowE2 sets Errors to Valve Movement too slow (E2) (0x04)
func (a *Errors) SetValveMovementTooSlowE2() { *a = 0x04 }

// IsValveNotMovingE3 checks if Errors equals the value for Valve not Moving (E3) (0x05)
func (a Errors) IsValveNotMovingE3() bool { return a == 0x05 }

// SetValveNotMovingE3 sets Errors to Valve not Moving (E3) (0x05)
func (a *Errors) SetValveNotMovingE3() { *a = 0x05 }

const CurrentTemperatureSetpointAttr zcl.AttrID = 16387

type CurrentTemperatureSetpoint zcl.Zs16

func (a CurrentTemperatureSetpoint) ID() zcl.AttrID                      { return CurrentTemperatureSetpointAttr }
func (a CurrentTemperatureSetpoint) Cluster() zcl.ClusterID              { return ThermostatID }
func (a *CurrentTemperatureSetpoint) Value() *CurrentTemperatureSetpoint { return a }
func (a CurrentTemperatureSetpoint) MarshalZcl() ([]byte, error) {
	return zcl.Zs16(a).MarshalZcl()
}
func (a *CurrentTemperatureSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentTemperatureSetpoint(*nt)
	return br, err
}

func (a CurrentTemperatureSetpoint) Readable() bool   { return true }
func (a CurrentTemperatureSetpoint) Writable() bool   { return true }
func (a CurrentTemperatureSetpoint) Reportable() bool { return false }
func (a CurrentTemperatureSetpoint) SceneIndex() int  { return -1 }

func (a CurrentTemperatureSetpoint) String() string {
	return zcl.Sprintf("%s", zcl.Zs16(a))
}

const HostFlagsAttr zcl.AttrID = 16392

type HostFlags zcl.Zu24

func (a HostFlags) ID() zcl.AttrID         { return HostFlagsAttr }
func (a HostFlags) Cluster() zcl.ClusterID { return ThermostatID }
func (a *HostFlags) Value() *HostFlags     { return a }
func (a HostFlags) MarshalZcl() ([]byte, error) {
	return zcl.Zu24(a).MarshalZcl()
}
func (a *HostFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = HostFlags(*nt)
	return br, err
}

func (a HostFlags) Readable() bool   { return true }
func (a HostFlags) Writable() bool   { return true }
func (a HostFlags) Reportable() bool { return false }
func (a HostFlags) SceneIndex() int  { return -1 }

func (a HostFlags) String() string {
	return zcl.Sprintf("0x%X", zcl.Zu24(a))
}
