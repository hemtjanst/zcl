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

func (v SetpointRaiseLower) ModeString() string {
	switch v.Mode {
	case 0x00:
		return "Heat (adjust Heat Setpoint)"
	case 0x01:
		return "Cool (adjust Cool Setpoint)"
	case 0x02:
		return "Both (adjust Heat Setpoint and Cool Setpoint)"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v.Mode))
}
func (v SetpointRaiseLower) AmountString() string {
	return zcl.Sprintf("%v", zcl.Zs8(v.Amount))
}

func (v SetpointRaiseLower) String() string {
	var str []string
	str = append(str, "Mode["+v.ModeString()+"]")
	str = append(str, "Amount["+v.AmountString()+"]")
	return "SetpointRaiseLower{" + zcl.StrJoin(str, " ") + "}"
}

func (SetpointRaiseLower) Name() string { return "Setpoint Raise/Lower" }

// LocalTemperature is an autogenerated attribute in the Thermostat cluster
type LocalTemperature zcl.Zs16

const LocalTemperatureAttr zcl.AttrID = 0

func (LocalTemperature) ID() zcl.AttrID                { return LocalTemperatureAttr }
func (LocalTemperature) Cluster() zcl.ClusterID        { return ThermostatID }
func (LocalTemperature) Name() string                  { return "Local Temperature" }
func (LocalTemperature) Readable() bool                { return true }
func (LocalTemperature) Writable() bool                { return false }
func (LocalTemperature) Reportable() bool              { return false }
func (LocalTemperature) SceneIndex() int               { return -1 }
func (a *LocalTemperature) Value() *LocalTemperature   { return a }
func (a LocalTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *LocalTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = LocalTemperature(*nt)
	return br, err
}

func (a LocalTemperature) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// OutdoorTemperature is an autogenerated attribute in the Thermostat cluster
type OutdoorTemperature zcl.Zs16

const OutdoorTemperatureAttr zcl.AttrID = 1

func (OutdoorTemperature) ID() zcl.AttrID                { return OutdoorTemperatureAttr }
func (OutdoorTemperature) Cluster() zcl.ClusterID        { return ThermostatID }
func (OutdoorTemperature) Name() string                  { return "Outdoor Temperature" }
func (OutdoorTemperature) Readable() bool                { return true }
func (OutdoorTemperature) Writable() bool                { return false }
func (OutdoorTemperature) Reportable() bool              { return false }
func (OutdoorTemperature) SceneIndex() int               { return -1 }
func (a *OutdoorTemperature) Value() *OutdoorTemperature { return a }
func (a OutdoorTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *OutdoorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = OutdoorTemperature(*nt)
	return br, err
}

func (a OutdoorTemperature) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// Occupancy is an autogenerated attribute in the Thermostat cluster
type Occupancy zcl.Zbmp8

const OccupancyAttr zcl.AttrID = 2

func (Occupancy) ID() zcl.AttrID                { return OccupancyAttr }
func (Occupancy) Cluster() zcl.ClusterID        { return ThermostatID }
func (Occupancy) Name() string                  { return "Occupancy" }
func (Occupancy) Readable() bool                { return true }
func (Occupancy) Writable() bool                { return false }
func (Occupancy) Reportable() bool              { return false }
func (Occupancy) SceneIndex() int               { return -1 }
func (a *Occupancy) Value() *Occupancy          { return a }
func (a Occupancy) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *Occupancy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Occupancy(*nt)
	return br, err
}

func (a Occupancy) String() string {
	return zcl.Sprintf("%v", zcl.Zbmp8(a))
}

// AbsMinHeatSetpointLimit is an autogenerated attribute in the Thermostat cluster
type AbsMinHeatSetpointLimit zcl.Zs16

const AbsMinHeatSetpointLimitAttr zcl.AttrID = 3

func (AbsMinHeatSetpointLimit) ID() zcl.AttrID                     { return AbsMinHeatSetpointLimitAttr }
func (AbsMinHeatSetpointLimit) Cluster() zcl.ClusterID             { return ThermostatID }
func (AbsMinHeatSetpointLimit) Name() string                       { return "Abs Min Heat Setpoint Limit" }
func (AbsMinHeatSetpointLimit) Readable() bool                     { return true }
func (AbsMinHeatSetpointLimit) Writable() bool                     { return false }
func (AbsMinHeatSetpointLimit) Reportable() bool                   { return false }
func (AbsMinHeatSetpointLimit) SceneIndex() int                    { return -1 }
func (a *AbsMinHeatSetpointLimit) Value() *AbsMinHeatSetpointLimit { return a }
func (a AbsMinHeatSetpointLimit) MarshalZcl() ([]byte, error)      { return zcl.Zs16(a).MarshalZcl() }

func (a *AbsMinHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AbsMinHeatSetpointLimit(*nt)
	return br, err
}

func (a AbsMinHeatSetpointLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// AbsMaxHeatSetpointLimit is an autogenerated attribute in the Thermostat cluster
type AbsMaxHeatSetpointLimit zcl.Zs16

const AbsMaxHeatSetpointLimitAttr zcl.AttrID = 4

func (AbsMaxHeatSetpointLimit) ID() zcl.AttrID                     { return AbsMaxHeatSetpointLimitAttr }
func (AbsMaxHeatSetpointLimit) Cluster() zcl.ClusterID             { return ThermostatID }
func (AbsMaxHeatSetpointLimit) Name() string                       { return "Abs Max Heat Setpoint Limit" }
func (AbsMaxHeatSetpointLimit) Readable() bool                     { return true }
func (AbsMaxHeatSetpointLimit) Writable() bool                     { return false }
func (AbsMaxHeatSetpointLimit) Reportable() bool                   { return false }
func (AbsMaxHeatSetpointLimit) SceneIndex() int                    { return -1 }
func (a *AbsMaxHeatSetpointLimit) Value() *AbsMaxHeatSetpointLimit { return a }
func (a AbsMaxHeatSetpointLimit) MarshalZcl() ([]byte, error)      { return zcl.Zs16(a).MarshalZcl() }

func (a *AbsMaxHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AbsMaxHeatSetpointLimit(*nt)
	return br, err
}

func (a AbsMaxHeatSetpointLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// AbsMinCoolSetpointLimit is an autogenerated attribute in the Thermostat cluster
type AbsMinCoolSetpointLimit zcl.Zs16

const AbsMinCoolSetpointLimitAttr zcl.AttrID = 5

func (AbsMinCoolSetpointLimit) ID() zcl.AttrID                     { return AbsMinCoolSetpointLimitAttr }
func (AbsMinCoolSetpointLimit) Cluster() zcl.ClusterID             { return ThermostatID }
func (AbsMinCoolSetpointLimit) Name() string                       { return "Abs Min Cool Setpoint Limit" }
func (AbsMinCoolSetpointLimit) Readable() bool                     { return true }
func (AbsMinCoolSetpointLimit) Writable() bool                     { return false }
func (AbsMinCoolSetpointLimit) Reportable() bool                   { return false }
func (AbsMinCoolSetpointLimit) SceneIndex() int                    { return -1 }
func (a *AbsMinCoolSetpointLimit) Value() *AbsMinCoolSetpointLimit { return a }
func (a AbsMinCoolSetpointLimit) MarshalZcl() ([]byte, error)      { return zcl.Zs16(a).MarshalZcl() }

func (a *AbsMinCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AbsMinCoolSetpointLimit(*nt)
	return br, err
}

func (a AbsMinCoolSetpointLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// AbsMaxCoolSetpointLimit is an autogenerated attribute in the Thermostat cluster
type AbsMaxCoolSetpointLimit zcl.Zs16

const AbsMaxCoolSetpointLimitAttr zcl.AttrID = 6

func (AbsMaxCoolSetpointLimit) ID() zcl.AttrID                     { return AbsMaxCoolSetpointLimitAttr }
func (AbsMaxCoolSetpointLimit) Cluster() zcl.ClusterID             { return ThermostatID }
func (AbsMaxCoolSetpointLimit) Name() string                       { return "Abs Max Cool Setpoint Limit" }
func (AbsMaxCoolSetpointLimit) Readable() bool                     { return true }
func (AbsMaxCoolSetpointLimit) Writable() bool                     { return false }
func (AbsMaxCoolSetpointLimit) Reportable() bool                   { return false }
func (AbsMaxCoolSetpointLimit) SceneIndex() int                    { return -1 }
func (a *AbsMaxCoolSetpointLimit) Value() *AbsMaxCoolSetpointLimit { return a }
func (a AbsMaxCoolSetpointLimit) MarshalZcl() ([]byte, error)      { return zcl.Zs16(a).MarshalZcl() }

func (a *AbsMaxCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = AbsMaxCoolSetpointLimit(*nt)
	return br, err
}

func (a AbsMaxCoolSetpointLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// PiCoolingDemand is an autogenerated attribute in the Thermostat cluster
type PiCoolingDemand zcl.Zu8

const PiCoolingDemandAttr zcl.AttrID = 7

func (PiCoolingDemand) ID() zcl.AttrID                { return PiCoolingDemandAttr }
func (PiCoolingDemand) Cluster() zcl.ClusterID        { return ThermostatID }
func (PiCoolingDemand) Name() string                  { return "PI Cooling Demand" }
func (PiCoolingDemand) Readable() bool                { return true }
func (PiCoolingDemand) Writable() bool                { return false }
func (PiCoolingDemand) Reportable() bool              { return false }
func (PiCoolingDemand) SceneIndex() int               { return -1 }
func (a *PiCoolingDemand) Value() *PiCoolingDemand    { return a }
func (a PiCoolingDemand) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *PiCoolingDemand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PiCoolingDemand(*nt)
	return br, err
}

func (a PiCoolingDemand) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// PiHeatingDemand is an autogenerated attribute in the Thermostat cluster
type PiHeatingDemand zcl.Zu8

const PiHeatingDemandAttr zcl.AttrID = 8

func (PiHeatingDemand) ID() zcl.AttrID                { return PiHeatingDemandAttr }
func (PiHeatingDemand) Cluster() zcl.ClusterID        { return ThermostatID }
func (PiHeatingDemand) Name() string                  { return "PI Heating Demand" }
func (PiHeatingDemand) Readable() bool                { return true }
func (PiHeatingDemand) Writable() bool                { return false }
func (PiHeatingDemand) Reportable() bool              { return false }
func (PiHeatingDemand) SceneIndex() int               { return -1 }
func (a *PiHeatingDemand) Value() *PiHeatingDemand    { return a }
func (a PiHeatingDemand) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *PiHeatingDemand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PiHeatingDemand(*nt)
	return br, err
}

func (a PiHeatingDemand) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// HvacSystemTypeConfiguration is an autogenerated attribute in the Thermostat cluster
type HvacSystemTypeConfiguration zcl.Zbmp8

const HvacSystemTypeConfigurationAttr zcl.AttrID = 9

func (HvacSystemTypeConfiguration) ID() zcl.AttrID                         { return HvacSystemTypeConfigurationAttr }
func (HvacSystemTypeConfiguration) Cluster() zcl.ClusterID                 { return ThermostatID }
func (HvacSystemTypeConfiguration) Name() string                           { return "HVAC System Type Configuration" }
func (HvacSystemTypeConfiguration) Readable() bool                         { return true }
func (HvacSystemTypeConfiguration) Writable() bool                         { return true }
func (HvacSystemTypeConfiguration) Reportable() bool                       { return false }
func (HvacSystemTypeConfiguration) SceneIndex() int                        { return -1 }
func (a *HvacSystemTypeConfiguration) Value() *HvacSystemTypeConfiguration { return a }
func (a HvacSystemTypeConfiguration) MarshalZcl() ([]byte, error)          { return zcl.Zbmp8(a).MarshalZcl() }

func (a *HvacSystemTypeConfiguration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = HvacSystemTypeConfiguration(*nt)
	return br, err
}

func (a HvacSystemTypeConfiguration) String() string {
	return zcl.Sprintf("%v", zcl.Zbmp8(a))
}

// LocalTemperatureCalibration is an autogenerated attribute in the Thermostat cluster
type LocalTemperatureCalibration zcl.Zs8

const LocalTemperatureCalibrationAttr zcl.AttrID = 16

func (LocalTemperatureCalibration) ID() zcl.AttrID                         { return LocalTemperatureCalibrationAttr }
func (LocalTemperatureCalibration) Cluster() zcl.ClusterID                 { return ThermostatID }
func (LocalTemperatureCalibration) Name() string                           { return "Local Temperature Calibration" }
func (LocalTemperatureCalibration) Readable() bool                         { return true }
func (LocalTemperatureCalibration) Writable() bool                         { return true }
func (LocalTemperatureCalibration) Reportable() bool                       { return false }
func (LocalTemperatureCalibration) SceneIndex() int                        { return -1 }
func (a *LocalTemperatureCalibration) Value() *LocalTemperatureCalibration { return a }
func (a LocalTemperatureCalibration) MarshalZcl() ([]byte, error)          { return zcl.Zs8(a).MarshalZcl() }

func (a *LocalTemperatureCalibration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = LocalTemperatureCalibration(*nt)
	return br, err
}

func (a LocalTemperatureCalibration) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(a))
}

// OccupiedCoolingSetpoint is an autogenerated attribute in the Thermostat cluster
type OccupiedCoolingSetpoint zcl.Zs16

const OccupiedCoolingSetpointAttr zcl.AttrID = 17

func (OccupiedCoolingSetpoint) ID() zcl.AttrID                     { return OccupiedCoolingSetpointAttr }
func (OccupiedCoolingSetpoint) Cluster() zcl.ClusterID             { return ThermostatID }
func (OccupiedCoolingSetpoint) Name() string                       { return "Occupied Cooling Setpoint" }
func (OccupiedCoolingSetpoint) Readable() bool                     { return true }
func (OccupiedCoolingSetpoint) Writable() bool                     { return true }
func (OccupiedCoolingSetpoint) Reportable() bool                   { return false }
func (OccupiedCoolingSetpoint) SceneIndex() int                    { return -1 }
func (a *OccupiedCoolingSetpoint) Value() *OccupiedCoolingSetpoint { return a }
func (a OccupiedCoolingSetpoint) MarshalZcl() ([]byte, error)      { return zcl.Zs16(a).MarshalZcl() }

func (a *OccupiedCoolingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = OccupiedCoolingSetpoint(*nt)
	return br, err
}

func (a OccupiedCoolingSetpoint) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// OccupiedHeatingSetpoint is an autogenerated attribute in the Thermostat cluster
type OccupiedHeatingSetpoint zcl.Zs16

const OccupiedHeatingSetpointAttr zcl.AttrID = 18

func (OccupiedHeatingSetpoint) ID() zcl.AttrID                     { return OccupiedHeatingSetpointAttr }
func (OccupiedHeatingSetpoint) Cluster() zcl.ClusterID             { return ThermostatID }
func (OccupiedHeatingSetpoint) Name() string                       { return "Occupied Heating Setpoint" }
func (OccupiedHeatingSetpoint) Readable() bool                     { return true }
func (OccupiedHeatingSetpoint) Writable() bool                     { return true }
func (OccupiedHeatingSetpoint) Reportable() bool                   { return false }
func (OccupiedHeatingSetpoint) SceneIndex() int                    { return -1 }
func (a *OccupiedHeatingSetpoint) Value() *OccupiedHeatingSetpoint { return a }
func (a OccupiedHeatingSetpoint) MarshalZcl() ([]byte, error)      { return zcl.Zs16(a).MarshalZcl() }

func (a *OccupiedHeatingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = OccupiedHeatingSetpoint(*nt)
	return br, err
}

func (a OccupiedHeatingSetpoint) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// UnoccupiedCoolingSetpoint is an autogenerated attribute in the Thermostat cluster
type UnoccupiedCoolingSetpoint zcl.Zs16

const UnoccupiedCoolingSetpointAttr zcl.AttrID = 19

func (UnoccupiedCoolingSetpoint) ID() zcl.AttrID                       { return UnoccupiedCoolingSetpointAttr }
func (UnoccupiedCoolingSetpoint) Cluster() zcl.ClusterID               { return ThermostatID }
func (UnoccupiedCoolingSetpoint) Name() string                         { return "Unoccupied Cooling Setpoint" }
func (UnoccupiedCoolingSetpoint) Readable() bool                       { return true }
func (UnoccupiedCoolingSetpoint) Writable() bool                       { return true }
func (UnoccupiedCoolingSetpoint) Reportable() bool                     { return false }
func (UnoccupiedCoolingSetpoint) SceneIndex() int                      { return -1 }
func (a *UnoccupiedCoolingSetpoint) Value() *UnoccupiedCoolingSetpoint { return a }
func (a UnoccupiedCoolingSetpoint) MarshalZcl() ([]byte, error)        { return zcl.Zs16(a).MarshalZcl() }

func (a *UnoccupiedCoolingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = UnoccupiedCoolingSetpoint(*nt)
	return br, err
}

func (a UnoccupiedCoolingSetpoint) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// UnoccupiedHeatingSetpoint is an autogenerated attribute in the Thermostat cluster
type UnoccupiedHeatingSetpoint zcl.Zs16

const UnoccupiedHeatingSetpointAttr zcl.AttrID = 20

func (UnoccupiedHeatingSetpoint) ID() zcl.AttrID                       { return UnoccupiedHeatingSetpointAttr }
func (UnoccupiedHeatingSetpoint) Cluster() zcl.ClusterID               { return ThermostatID }
func (UnoccupiedHeatingSetpoint) Name() string                         { return "Unoccupied Heating Setpoint" }
func (UnoccupiedHeatingSetpoint) Readable() bool                       { return true }
func (UnoccupiedHeatingSetpoint) Writable() bool                       { return true }
func (UnoccupiedHeatingSetpoint) Reportable() bool                     { return false }
func (UnoccupiedHeatingSetpoint) SceneIndex() int                      { return -1 }
func (a *UnoccupiedHeatingSetpoint) Value() *UnoccupiedHeatingSetpoint { return a }
func (a UnoccupiedHeatingSetpoint) MarshalZcl() ([]byte, error)        { return zcl.Zs16(a).MarshalZcl() }

func (a *UnoccupiedHeatingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = UnoccupiedHeatingSetpoint(*nt)
	return br, err
}

func (a UnoccupiedHeatingSetpoint) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// MinHeatSetpointLimit is an autogenerated attribute in the Thermostat cluster
type MinHeatSetpointLimit zcl.Zs16

const MinHeatSetpointLimitAttr zcl.AttrID = 21

func (MinHeatSetpointLimit) ID() zcl.AttrID                  { return MinHeatSetpointLimitAttr }
func (MinHeatSetpointLimit) Cluster() zcl.ClusterID          { return ThermostatID }
func (MinHeatSetpointLimit) Name() string                    { return "Min Heat Setpoint Limit" }
func (MinHeatSetpointLimit) Readable() bool                  { return true }
func (MinHeatSetpointLimit) Writable() bool                  { return true }
func (MinHeatSetpointLimit) Reportable() bool                { return false }
func (MinHeatSetpointLimit) SceneIndex() int                 { return -1 }
func (a *MinHeatSetpointLimit) Value() *MinHeatSetpointLimit { return a }
func (a MinHeatSetpointLimit) MarshalZcl() ([]byte, error)   { return zcl.Zs16(a).MarshalZcl() }

func (a *MinHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinHeatSetpointLimit(*nt)
	return br, err
}

func (a MinHeatSetpointLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// MaxHeatSetpointLimit is an autogenerated attribute in the Thermostat cluster
type MaxHeatSetpointLimit zcl.Zs16

const MaxHeatSetpointLimitAttr zcl.AttrID = 22

func (MaxHeatSetpointLimit) ID() zcl.AttrID                  { return MaxHeatSetpointLimitAttr }
func (MaxHeatSetpointLimit) Cluster() zcl.ClusterID          { return ThermostatID }
func (MaxHeatSetpointLimit) Name() string                    { return "Max Heat Setpoint Limit" }
func (MaxHeatSetpointLimit) Readable() bool                  { return true }
func (MaxHeatSetpointLimit) Writable() bool                  { return true }
func (MaxHeatSetpointLimit) Reportable() bool                { return false }
func (MaxHeatSetpointLimit) SceneIndex() int                 { return -1 }
func (a *MaxHeatSetpointLimit) Value() *MaxHeatSetpointLimit { return a }
func (a MaxHeatSetpointLimit) MarshalZcl() ([]byte, error)   { return zcl.Zs16(a).MarshalZcl() }

func (a *MaxHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxHeatSetpointLimit(*nt)
	return br, err
}

func (a MaxHeatSetpointLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// MinCoolSetpointLimit is an autogenerated attribute in the Thermostat cluster
type MinCoolSetpointLimit zcl.Zs16

const MinCoolSetpointLimitAttr zcl.AttrID = 23

func (MinCoolSetpointLimit) ID() zcl.AttrID                  { return MinCoolSetpointLimitAttr }
func (MinCoolSetpointLimit) Cluster() zcl.ClusterID          { return ThermostatID }
func (MinCoolSetpointLimit) Name() string                    { return "Min Cool Setpoint Limit" }
func (MinCoolSetpointLimit) Readable() bool                  { return true }
func (MinCoolSetpointLimit) Writable() bool                  { return true }
func (MinCoolSetpointLimit) Reportable() bool                { return false }
func (MinCoolSetpointLimit) SceneIndex() int                 { return -1 }
func (a *MinCoolSetpointLimit) Value() *MinCoolSetpointLimit { return a }
func (a MinCoolSetpointLimit) MarshalZcl() ([]byte, error)   { return zcl.Zs16(a).MarshalZcl() }

func (a *MinCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MinCoolSetpointLimit(*nt)
	return br, err
}

func (a MinCoolSetpointLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// MaxCoolSetpointLimit is an autogenerated attribute in the Thermostat cluster
type MaxCoolSetpointLimit zcl.Zs16

const MaxCoolSetpointLimitAttr zcl.AttrID = 24

func (MaxCoolSetpointLimit) ID() zcl.AttrID                  { return MaxCoolSetpointLimitAttr }
func (MaxCoolSetpointLimit) Cluster() zcl.ClusterID          { return ThermostatID }
func (MaxCoolSetpointLimit) Name() string                    { return "Max Cool Setpoint Limit" }
func (MaxCoolSetpointLimit) Readable() bool                  { return true }
func (MaxCoolSetpointLimit) Writable() bool                  { return true }
func (MaxCoolSetpointLimit) Reportable() bool                { return false }
func (MaxCoolSetpointLimit) SceneIndex() int                 { return -1 }
func (a *MaxCoolSetpointLimit) Value() *MaxCoolSetpointLimit { return a }
func (a MaxCoolSetpointLimit) MarshalZcl() ([]byte, error)   { return zcl.Zs16(a).MarshalZcl() }

func (a *MaxCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxCoolSetpointLimit(*nt)
	return br, err
}

func (a MaxCoolSetpointLimit) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// MinSetpointDeadBand is an autogenerated attribute in the Thermostat cluster
type MinSetpointDeadBand zcl.Zs8

const MinSetpointDeadBandAttr zcl.AttrID = 25

func (MinSetpointDeadBand) ID() zcl.AttrID                 { return MinSetpointDeadBandAttr }
func (MinSetpointDeadBand) Cluster() zcl.ClusterID         { return ThermostatID }
func (MinSetpointDeadBand) Name() string                   { return "Min Setpoint Dead Band" }
func (MinSetpointDeadBand) Readable() bool                 { return true }
func (MinSetpointDeadBand) Writable() bool                 { return true }
func (MinSetpointDeadBand) Reportable() bool               { return false }
func (MinSetpointDeadBand) SceneIndex() int                { return -1 }
func (a *MinSetpointDeadBand) Value() *MinSetpointDeadBand { return a }
func (a MinSetpointDeadBand) MarshalZcl() ([]byte, error)  { return zcl.Zs8(a).MarshalZcl() }

func (a *MinSetpointDeadBand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*a = MinSetpointDeadBand(*nt)
	return br, err
}

func (a MinSetpointDeadBand) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(a))
}

// RemoteSensing is an autogenerated attribute in the Thermostat cluster
type RemoteSensing zcl.Zbmp8

const RemoteSensingAttr zcl.AttrID = 26

func (RemoteSensing) ID() zcl.AttrID                { return RemoteSensingAttr }
func (RemoteSensing) Cluster() zcl.ClusterID        { return ThermostatID }
func (RemoteSensing) Name() string                  { return "Remote Sensing" }
func (RemoteSensing) Readable() bool                { return true }
func (RemoteSensing) Writable() bool                { return true }
func (RemoteSensing) Reportable() bool              { return false }
func (RemoteSensing) SceneIndex() int               { return -1 }
func (a *RemoteSensing) Value() *RemoteSensing      { return a }
func (a RemoteSensing) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *RemoteSensing) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = RemoteSensing(*nt)
	return br, err
}

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

// ControlSequenceOfOperation is an autogenerated attribute in the Thermostat cluster
type ControlSequenceOfOperation zcl.Zenum8

const ControlSequenceOfOperationAttr zcl.AttrID = 27

func (ControlSequenceOfOperation) ID() zcl.AttrID                        { return ControlSequenceOfOperationAttr }
func (ControlSequenceOfOperation) Cluster() zcl.ClusterID                { return ThermostatID }
func (ControlSequenceOfOperation) Name() string                          { return "Control Sequence Of Operation" }
func (ControlSequenceOfOperation) Readable() bool                        { return true }
func (ControlSequenceOfOperation) Writable() bool                        { return true }
func (ControlSequenceOfOperation) Reportable() bool                      { return false }
func (ControlSequenceOfOperation) SceneIndex() int                       { return -1 }
func (a *ControlSequenceOfOperation) Value() *ControlSequenceOfOperation { return a }
func (a ControlSequenceOfOperation) MarshalZcl() ([]byte, error)         { return zcl.Zenum8(a).MarshalZcl() }

func (a *ControlSequenceOfOperation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = ControlSequenceOfOperation(*nt)
	return br, err
}

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
	return zcl.Sprintf("%v", zcl.Zenum8(a))
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

// SystemMode is an autogenerated attribute in the Thermostat cluster
type SystemMode zcl.Zenum8

const SystemModeAttr zcl.AttrID = 28

func (SystemMode) ID() zcl.AttrID                { return SystemModeAttr }
func (SystemMode) Cluster() zcl.ClusterID        { return ThermostatID }
func (SystemMode) Name() string                  { return "System Mode" }
func (SystemMode) Readable() bool                { return true }
func (SystemMode) Writable() bool                { return true }
func (SystemMode) Reportable() bool              { return false }
func (SystemMode) SceneIndex() int               { return -1 }
func (a *SystemMode) Value() *SystemMode         { return a }
func (a SystemMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *SystemMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = SystemMode(*nt)
	return br, err
}

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
	return zcl.Sprintf("%v", zcl.Zenum8(a))
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

// AlarmMask is an autogenerated attribute in the Thermostat cluster
type AlarmMask zcl.Zbmp8

const AlarmMaskAttr zcl.AttrID = 29

func (AlarmMask) ID() zcl.AttrID                { return AlarmMaskAttr }
func (AlarmMask) Cluster() zcl.ClusterID        { return ThermostatID }
func (AlarmMask) Name() string                  { return "Alarm Mask" }
func (AlarmMask) Readable() bool                { return true }
func (AlarmMask) Writable() bool                { return false }
func (AlarmMask) Reportable() bool              { return false }
func (AlarmMask) SceneIndex() int               { return -1 }
func (a *AlarmMask) Value() *AlarmMask          { return a }
func (a AlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *AlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmMask(*nt)
	return br, err
}

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

// ThermostatRunningMode is an autogenerated attribute in the Thermostat cluster
type ThermostatRunningMode zcl.Zenum8

const ThermostatRunningModeAttr zcl.AttrID = 30

func (ThermostatRunningMode) ID() zcl.AttrID                   { return ThermostatRunningModeAttr }
func (ThermostatRunningMode) Cluster() zcl.ClusterID           { return ThermostatID }
func (ThermostatRunningMode) Name() string                     { return "Thermostat Running Mode" }
func (ThermostatRunningMode) Readable() bool                   { return true }
func (ThermostatRunningMode) Writable() bool                   { return false }
func (ThermostatRunningMode) Reportable() bool                 { return false }
func (ThermostatRunningMode) SceneIndex() int                  { return -1 }
func (a *ThermostatRunningMode) Value() *ThermostatRunningMode { return a }
func (a ThermostatRunningMode) MarshalZcl() ([]byte, error)    { return zcl.Zenum8(a).MarshalZcl() }

func (a *ThermostatRunningMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = ThermostatRunningMode(*nt)
	return br, err
}

func (a ThermostatRunningMode) String() string {
	switch a {
	case 0x00:
		return "Off"
	case 0x03:
		return "Cool"
	case 0x04:
		return "Heat"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
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

// StartOfWeek is an autogenerated attribute in the Thermostat cluster
type StartOfWeek zcl.Zenum8

const StartOfWeekAttr zcl.AttrID = 32

func (StartOfWeek) ID() zcl.AttrID                { return StartOfWeekAttr }
func (StartOfWeek) Cluster() zcl.ClusterID        { return ThermostatID }
func (StartOfWeek) Name() string                  { return "Start of Week" }
func (StartOfWeek) Readable() bool                { return true }
func (StartOfWeek) Writable() bool                { return false }
func (StartOfWeek) Reportable() bool              { return false }
func (StartOfWeek) SceneIndex() int               { return -1 }
func (a *StartOfWeek) Value() *StartOfWeek        { return a }
func (a StartOfWeek) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *StartOfWeek) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = StartOfWeek(*nt)
	return br, err
}

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
	return zcl.Sprintf("%v", zcl.Zenum8(a))
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

// NumberOfWeeklyTransitions is an autogenerated attribute in the Thermostat cluster
type NumberOfWeeklyTransitions zcl.Zu8

const NumberOfWeeklyTransitionsAttr zcl.AttrID = 33

func (NumberOfWeeklyTransitions) ID() zcl.AttrID                       { return NumberOfWeeklyTransitionsAttr }
func (NumberOfWeeklyTransitions) Cluster() zcl.ClusterID               { return ThermostatID }
func (NumberOfWeeklyTransitions) Name() string                         { return "Number of Weekly Transitions" }
func (NumberOfWeeklyTransitions) Readable() bool                       { return true }
func (NumberOfWeeklyTransitions) Writable() bool                       { return false }
func (NumberOfWeeklyTransitions) Reportable() bool                     { return false }
func (NumberOfWeeklyTransitions) SceneIndex() int                      { return -1 }
func (a *NumberOfWeeklyTransitions) Value() *NumberOfWeeklyTransitions { return a }
func (a NumberOfWeeklyTransitions) MarshalZcl() ([]byte, error)        { return zcl.Zu8(a).MarshalZcl() }

func (a *NumberOfWeeklyTransitions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfWeeklyTransitions(*nt)
	return br, err
}

func (a NumberOfWeeklyTransitions) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// NumberOfDailTransitions is an autogenerated attribute in the Thermostat cluster
type NumberOfDailTransitions zcl.Zu8

const NumberOfDailTransitionsAttr zcl.AttrID = 34

func (NumberOfDailTransitions) ID() zcl.AttrID                     { return NumberOfDailTransitionsAttr }
func (NumberOfDailTransitions) Cluster() zcl.ClusterID             { return ThermostatID }
func (NumberOfDailTransitions) Name() string                       { return "Number of Dail Transitions" }
func (NumberOfDailTransitions) Readable() bool                     { return true }
func (NumberOfDailTransitions) Writable() bool                     { return false }
func (NumberOfDailTransitions) Reportable() bool                   { return false }
func (NumberOfDailTransitions) SceneIndex() int                    { return -1 }
func (a *NumberOfDailTransitions) Value() *NumberOfDailTransitions { return a }
func (a NumberOfDailTransitions) MarshalZcl() ([]byte, error)      { return zcl.Zu8(a).MarshalZcl() }

func (a *NumberOfDailTransitions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfDailTransitions(*nt)
	return br, err
}

func (a NumberOfDailTransitions) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// TemperatureSetpointHold is an autogenerated attribute in the Thermostat cluster
type TemperatureSetpointHold zcl.Zenum8

const TemperatureSetpointHoldAttr zcl.AttrID = 35

func (TemperatureSetpointHold) ID() zcl.AttrID                     { return TemperatureSetpointHoldAttr }
func (TemperatureSetpointHold) Cluster() zcl.ClusterID             { return ThermostatID }
func (TemperatureSetpointHold) Name() string                       { return "Temperature Setpoint Hold" }
func (TemperatureSetpointHold) Readable() bool                     { return true }
func (TemperatureSetpointHold) Writable() bool                     { return false }
func (TemperatureSetpointHold) Reportable() bool                   { return false }
func (TemperatureSetpointHold) SceneIndex() int                    { return -1 }
func (a *TemperatureSetpointHold) Value() *TemperatureSetpointHold { return a }
func (a TemperatureSetpointHold) MarshalZcl() ([]byte, error)      { return zcl.Zenum8(a).MarshalZcl() }

func (a *TemperatureSetpointHold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = TemperatureSetpointHold(*nt)
	return br, err
}

func (a TemperatureSetpointHold) String() string {
	switch a {
	case 0x00:
		return "Setpoint Hold Off"
	case 0x01:
		return "Setpoint Hold On"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

// IsSetpointHoldOff checks if TemperatureSetpointHold equals the value for Setpoint Hold Off (0x00)
func (a TemperatureSetpointHold) IsSetpointHoldOff() bool { return a == 0x00 }

// SetSetpointHoldOff sets TemperatureSetpointHold to Setpoint Hold Off (0x00)
func (a *TemperatureSetpointHold) SetSetpointHoldOff() { *a = 0x00 }

// IsSetpointHoldOn checks if TemperatureSetpointHold equals the value for Setpoint Hold On (0x01)
func (a TemperatureSetpointHold) IsSetpointHoldOn() bool { return a == 0x01 }

// SetSetpointHoldOn sets TemperatureSetpointHold to Setpoint Hold On (0x01)
func (a *TemperatureSetpointHold) SetSetpointHoldOn() { *a = 0x01 }

// TemperatureSetpointHoldDuration is an autogenerated attribute in the Thermostat cluster
type TemperatureSetpointHoldDuration zcl.Zu8

const TemperatureSetpointHoldDurationAttr zcl.AttrID = 36

func (TemperatureSetpointHoldDuration) ID() zcl.AttrID                             { return TemperatureSetpointHoldDurationAttr }
func (TemperatureSetpointHoldDuration) Cluster() zcl.ClusterID                     { return ThermostatID }
func (TemperatureSetpointHoldDuration) Name() string                               { return "Temperature Setpoint Hold Duration" }
func (TemperatureSetpointHoldDuration) Readable() bool                             { return true }
func (TemperatureSetpointHoldDuration) Writable() bool                             { return true }
func (TemperatureSetpointHoldDuration) Reportable() bool                           { return false }
func (TemperatureSetpointHoldDuration) SceneIndex() int                            { return -1 }
func (a *TemperatureSetpointHoldDuration) Value() *TemperatureSetpointHoldDuration { return a }
func (a TemperatureSetpointHoldDuration) MarshalZcl() ([]byte, error)              { return zcl.Zu8(a).MarshalZcl() }

func (a *TemperatureSetpointHoldDuration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = TemperatureSetpointHoldDuration(*nt)
	return br, err
}

func (a TemperatureSetpointHoldDuration) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// ThermostatProgrammingOperationMode is an autogenerated attribute in the Thermostat cluster
type ThermostatProgrammingOperationMode zcl.Zbmp8

const ThermostatProgrammingOperationModeAttr zcl.AttrID = 37

func (ThermostatProgrammingOperationMode) ID() zcl.AttrID {
	return ThermostatProgrammingOperationModeAttr
}
func (ThermostatProgrammingOperationMode) Cluster() zcl.ClusterID { return ThermostatID }
func (ThermostatProgrammingOperationMode) Name() string {
	return "Thermostat Programming Operation Mode"
}
func (ThermostatProgrammingOperationMode) Readable() bool                                { return true }
func (ThermostatProgrammingOperationMode) Writable() bool                                { return true }
func (ThermostatProgrammingOperationMode) Reportable() bool                              { return false }
func (ThermostatProgrammingOperationMode) SceneIndex() int                               { return -1 }
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

// ThermostatRunningState is an autogenerated attribute in the Thermostat cluster
type ThermostatRunningState zcl.Zbmp16

const ThermostatRunningStateAttr zcl.AttrID = 41

func (ThermostatRunningState) ID() zcl.AttrID                    { return ThermostatRunningStateAttr }
func (ThermostatRunningState) Cluster() zcl.ClusterID            { return ThermostatID }
func (ThermostatRunningState) Name() string                      { return "Thermostat Running State" }
func (ThermostatRunningState) Readable() bool                    { return true }
func (ThermostatRunningState) Writable() bool                    { return false }
func (ThermostatRunningState) Reportable() bool                  { return false }
func (ThermostatRunningState) SceneIndex() int                   { return -1 }
func (a *ThermostatRunningState) Value() *ThermostatRunningState { return a }
func (a ThermostatRunningState) MarshalZcl() ([]byte, error)     { return zcl.Zbmp16(a).MarshalZcl() }

func (a *ThermostatRunningState) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = ThermostatRunningState(*nt)
	return br, err
}

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

// TrvMode is an autogenerated attribute in the Thermostat cluster
type TrvMode zcl.Zenum8

const TrvModeAttr zcl.AttrID = 16384

func (TrvMode) ID() zcl.AttrID                { return TrvModeAttr }
func (TrvMode) Cluster() zcl.ClusterID        { return ThermostatID }
func (TrvMode) Name() string                  { return "TRV Mode" }
func (TrvMode) Readable() bool                { return true }
func (TrvMode) Writable() bool                { return true }
func (TrvMode) Reportable() bool              { return false }
func (TrvMode) SceneIndex() int               { return -1 }
func (a *TrvMode) Value() *TrvMode            { return a }
func (a TrvMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *TrvMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = TrvMode(*nt)
	return br, err
}

func (a TrvMode) String() string {
	switch a {
	case 0x00:
		return "Unknown 1"
	case 0x01:
		return "Unknown 2"
	case 0x02:
		return "manual"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
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

// SetValvePosition is an autogenerated attribute in the Thermostat cluster
type SetValvePosition zcl.Zu8

const SetValvePositionAttr zcl.AttrID = 16385

func (SetValvePosition) ID() zcl.AttrID                { return SetValvePositionAttr }
func (SetValvePosition) Cluster() zcl.ClusterID        { return ThermostatID }
func (SetValvePosition) Name() string                  { return "Set Valve Position" }
func (SetValvePosition) Readable() bool                { return true }
func (SetValvePosition) Writable() bool                { return true }
func (SetValvePosition) Reportable() bool              { return false }
func (SetValvePosition) SceneIndex() int               { return -1 }
func (a *SetValvePosition) Value() *SetValvePosition   { return a }
func (a SetValvePosition) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *SetValvePosition) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = SetValvePosition(*nt)
	return br, err
}

func (a SetValvePosition) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// Errors is an autogenerated attribute in the Thermostat cluster
type Errors zcl.Zu8

const ErrorsAttr zcl.AttrID = 16386

func (Errors) ID() zcl.AttrID                { return ErrorsAttr }
func (Errors) Cluster() zcl.ClusterID        { return ThermostatID }
func (Errors) Name() string                  { return "Errors" }
func (Errors) Readable() bool                { return true }
func (Errors) Writable() bool                { return false }
func (Errors) Reportable() bool              { return false }
func (Errors) SceneIndex() int               { return -1 }
func (a *Errors) Value() *Errors             { return a }
func (a Errors) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Errors) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Errors(*nt)
	return br, err
}

func (a Errors) String() string {
	switch a {
	case 0x03:
		return "Valve Adaption Failed (E1)"
	case 0x04:
		return "Valve Movement too slow (E2)"
	case 0x05:
		return "Valve not Moving (E3)"
	}
	return zcl.Sprintf("%v", zcl.Zu8(a))
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

// CurrentTemperatureSetpoint is an autogenerated attribute in the Thermostat cluster
type CurrentTemperatureSetpoint zcl.Zs16

const CurrentTemperatureSetpointAttr zcl.AttrID = 16387

func (CurrentTemperatureSetpoint) ID() zcl.AttrID                        { return CurrentTemperatureSetpointAttr }
func (CurrentTemperatureSetpoint) Cluster() zcl.ClusterID                { return ThermostatID }
func (CurrentTemperatureSetpoint) Name() string                          { return "Current Temperature Setpoint" }
func (CurrentTemperatureSetpoint) Readable() bool                        { return true }
func (CurrentTemperatureSetpoint) Writable() bool                        { return true }
func (CurrentTemperatureSetpoint) Reportable() bool                      { return false }
func (CurrentTemperatureSetpoint) SceneIndex() int                       { return -1 }
func (a *CurrentTemperatureSetpoint) Value() *CurrentTemperatureSetpoint { return a }
func (a CurrentTemperatureSetpoint) MarshalZcl() ([]byte, error)         { return zcl.Zs16(a).MarshalZcl() }

func (a *CurrentTemperatureSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentTemperatureSetpoint(*nt)
	return br, err
}

func (a CurrentTemperatureSetpoint) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// HostFlags is an autogenerated attribute in the Thermostat cluster
type HostFlags zcl.Zu24

const HostFlagsAttr zcl.AttrID = 16392

func (HostFlags) ID() zcl.AttrID                { return HostFlagsAttr }
func (HostFlags) Cluster() zcl.ClusterID        { return ThermostatID }
func (HostFlags) Name() string                  { return "Host Flags" }
func (HostFlags) Readable() bool                { return true }
func (HostFlags) Writable() bool                { return true }
func (HostFlags) Reportable() bool              { return false }
func (HostFlags) SceneIndex() int               { return -1 }
func (a *HostFlags) Value() *HostFlags          { return a }
func (a HostFlags) MarshalZcl() ([]byte, error) { return zcl.Zu24(a).MarshalZcl() }

func (a *HostFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = HostFlags(*nt)
	return br, err
}

func (a HostFlags) String() string {
	return zcl.Sprintf("%v", zcl.Zu24(a))
}
