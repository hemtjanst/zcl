package hvac

import "hemtjan.st/zcl"

// Thermostat
const ThermostatID zcl.ClusterID = 513

var ThermostatCluster = zcl.Cluster{
	Name: "Thermostat",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		SetpointRaiseLowerCommand:  func() zcl.Command { return new(SetpointRaiseLower) },
		SetWeeklyScheduleCommand:   func() zcl.Command { return new(SetWeeklySchedule) },
		GetWeeklyScheduleCommand:   func() zcl.Command { return new(GetWeeklySchedule) },
		ClearWeeklyScheduleCommand: func() zcl.Command { return new(ClearWeeklySchedule) },
		GetRelayStatusLogCommand:   func() zcl.Command { return new(GetRelayStatusLog) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		GetWeeklyScheduleResponseCommand: func() zcl.Command { return new(GetWeeklyScheduleResponse) },
		GetRelayStatusLogResponseCommand: func() zcl.Command { return new(GetRelayStatusLogResponse) },
	},
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
		ThermostatAlarmMaskAttr:                func() zcl.Attr { return new(ThermostatAlarmMask) },
		ThermostatRunningModeAttr:              func() zcl.Attr { return new(ThermostatRunningMode) },
		StartOfWeekAttr:                        func() zcl.Attr { return new(StartOfWeek) },
		NumberOfWeeklyTransitionsAttr:          func() zcl.Attr { return new(NumberOfWeeklyTransitions) },
		NumberOfDailyTransitionsAttr:           func() zcl.Attr { return new(NumberOfDailyTransitions) },
		TemperatureSetpointHoldAttr:            func() zcl.Attr { return new(TemperatureSetpointHold) },
		TemperatureSetpointHoldDurationAttr:    func() zcl.Attr { return new(TemperatureSetpointHoldDuration) },
		ThermostatProgrammingOperationModeAttr: func() zcl.Attr { return new(ThermostatProgrammingOperationMode) },
		ThermostatRunningStateAttr:             func() zcl.Attr { return new(ThermostatRunningState) },
		SetpointChangeSourceAttr:               func() zcl.Attr { return new(SetpointChangeSource) },
		SetpointChangeAmountAttr:               func() zcl.Attr { return new(SetpointChangeAmount) },
		SetpointChangeSourceTimestampAttr:      func() zcl.Attr { return new(SetpointChangeSourceTimestamp) },
		OccupiedSetbackAttr:                    func() zcl.Attr { return new(OccupiedSetback) },
		OccupiedSetbackMinAttr:                 func() zcl.Attr { return new(OccupiedSetbackMin) },
		OccupiedSetbackMaxAttr:                 func() zcl.Attr { return new(OccupiedSetbackMax) },
		UnoccupiedSetbackAttr:                  func() zcl.Attr { return new(UnoccupiedSetback) },
		UnoccupiedSetbackMinAttr:               func() zcl.Attr { return new(UnoccupiedSetbackMin) },
		UnoccupiedSetbackMaxAttr:               func() zcl.Attr { return new(UnoccupiedSetbackMax) },
		EmergencyHeatDeltaAttr:                 func() zcl.Attr { return new(EmergencyHeatDelta) },
		AcTypeAttr:                             func() zcl.Attr { return new(AcType) },
		AcCapacityAttr:                         func() zcl.Attr { return new(AcCapacity) },
		AcRefrigerantTypeAttr:                  func() zcl.Attr { return new(AcRefrigerantType) },
		AcCompressorTypeAttr:                   func() zcl.Attr { return new(AcCompressorType) },
		AcErrorCodeAttr:                        func() zcl.Attr { return new(AcErrorCode) },
		AcLouverPositionAttr:                   func() zcl.Attr { return new(AcLouverPosition) },
		AcCoilTemperatureAttr:                  func() zcl.Attr { return new(AcCoilTemperature) },
		AcCapacityFormatAttr:                   func() zcl.Attr { return new(AcCapacityFormat) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr: []zcl.AttrID{
		OccupiedCoolingSetpointAttr,
		OccupiedHeatingSetpointAttr,
		SystemModeAttr,
	},
}

type SetpointRaiseLower struct {
	SetpointMode   SetpointMode
	SetpointAmount SetpointAmount
}

type SetpointRaiseLowerHandler interface {
	HandleSetpointRaiseLower(frame Frame, cmd *SetpointRaiseLower) error
}

// SetpointRaiseLowerCommand is the Command ID of SetpointRaiseLower
const SetpointRaiseLowerCommand CommandID = 0x0000

// Values returns all values of SetpointRaiseLower
func (v *SetpointRaiseLower) Values() []zcl.Val {
	return []zcl.Val{
		&v.SetpointMode,
		&v.SetpointAmount,
	}
}

// Arguments returns all values of SetpointRaiseLower
func (v *SetpointRaiseLower) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "SetpointMode", Argument: &v.SetpointMode},
		{Name: "SetpointAmount", Argument: &v.SetpointAmount},
	}
}

// Name of the command
func (SetpointRaiseLower) Name() string { return `Setpoint raise/lower` }

// Description of the command
func (SetpointRaiseLower) Description() string { return `` }

// ID of the command
func (SetpointRaiseLower) ID() CommandID { return SetpointRaiseLowerCommand }

// Required
func (SetpointRaiseLower) Required() bool { return true }

// Cluster ID of the command
func (SetpointRaiseLower) Cluster() zcl.ClusterID { return ThermostatID }

// Direction of the command
func (SetpointRaiseLower) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (SetpointRaiseLower) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SetpointRaiseLower) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *SetpointRaiseLower) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h SetpointRaiseLowerHandler
	if h, found = handler.(SetpointRaiseLowerHandler); found {
		err = h.HandleSetpointRaiseLower(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of SetpointRaiseLower
func (v SetpointRaiseLower) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.SetpointMode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SetpointAmount.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the SetpointRaiseLower struct
func (v *SetpointRaiseLower) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.SetpointMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SetpointAmount).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v SetpointRaiseLower) String() string {
	return zcl.Sprintf(
		"SetpointRaiseLower{"+zcl.StrJoin([]string{
			"SetpointMode(%v)",
			"SetpointAmount(%v)",
		}, " ")+"}",
		v.SetpointMode,
		v.SetpointAmount,
	)
}

type SetWeeklySchedule struct {
	// SetWeeklyNumberOfTransitions how many individual transitions to expect for this sequence of
	// commands. If a device supports more than 10 transitions in its
	// schedule they can send this by sending more than 1 Set Weekly Schedule
	// command, each containing the separate information that the device
	// needs to set
	SetWeeklyNumberOfTransitions SetWeeklyNumberOfTransitions
	// SetWeeklyDayOfWeek day of the week at which all the transitions within the payload of the
	// command should be associated to
	SetWeeklyDayOfWeek SetWeeklyDayOfWeek
	// SetWeeklyMode which type of setpoint transition is present in the rest of the
	// command
	SetWeeklyMode             SetWeeklyMode
	SetWeeklyTransitionTime1  SetWeeklyTransitionTime1
	SetWeeklyHeatSetpoint1    SetWeeklyHeatSetpoint1
	SetWeeklyCoolSetpoint1    SetWeeklyCoolSetpoint1
	SetWeeklyTransitionTime10 SetWeeklyTransitionTime10
	SetWeeklyHeatSetpoint10   SetWeeklyHeatSetpoint10
	SetWeeklyCoolSetpoint10   SetWeeklyCoolSetpoint10
}

type SetWeeklyScheduleHandler interface {
	HandleSetWeeklySchedule(frame Frame, cmd *SetWeeklySchedule) error
}

// SetWeeklyScheduleCommand is the Command ID of SetWeeklySchedule
const SetWeeklyScheduleCommand CommandID = 0x0001

// Values returns all values of SetWeeklySchedule
func (v *SetWeeklySchedule) Values() []zcl.Val {
	return []zcl.Val{
		&v.SetWeeklyNumberOfTransitions,
		&v.SetWeeklyDayOfWeek,
		&v.SetWeeklyMode,
		&v.SetWeeklyTransitionTime1,
		&v.SetWeeklyHeatSetpoint1,
		&v.SetWeeklyCoolSetpoint1,
		&v.SetWeeklyTransitionTime10,
		&v.SetWeeklyHeatSetpoint10,
		&v.SetWeeklyCoolSetpoint10,
	}
}

// Arguments returns all values of SetWeeklySchedule
func (v *SetWeeklySchedule) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "SetWeeklyNumberOfTransitions", Argument: &v.SetWeeklyNumberOfTransitions},
		{Name: "SetWeeklyDayOfWeek", Argument: &v.SetWeeklyDayOfWeek},
		{Name: "SetWeeklyMode", Argument: &v.SetWeeklyMode},
		{Name: "SetWeeklyTransitionTime1", Argument: &v.SetWeeklyTransitionTime1},
		{Name: "SetWeeklyHeatSetpoint1", Argument: &v.SetWeeklyHeatSetpoint1},
		{Name: "SetWeeklyCoolSetpoint1", Argument: &v.SetWeeklyCoolSetpoint1},
		{Name: "SetWeeklyTransitionTime10", Argument: &v.SetWeeklyTransitionTime10},
		{Name: "SetWeeklyHeatSetpoint10", Argument: &v.SetWeeklyHeatSetpoint10},
		{Name: "SetWeeklyCoolSetpoint10", Argument: &v.SetWeeklyCoolSetpoint10},
	}
}

// Name of the command
func (SetWeeklySchedule) Name() string { return `Set weekly schedule` }

// Description of the command
func (SetWeeklySchedule) Description() string { return `` }

// ID of the command
func (SetWeeklySchedule) ID() CommandID { return SetWeeklyScheduleCommand }

// Required
func (SetWeeklySchedule) Required() bool { return false }

// Cluster ID of the command
func (SetWeeklySchedule) Cluster() zcl.ClusterID { return ThermostatID }

// Direction of the command
func (SetWeeklySchedule) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (SetWeeklySchedule) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SetWeeklySchedule) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *SetWeeklySchedule) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h SetWeeklyScheduleHandler
	if h, found = handler.(SetWeeklyScheduleHandler); found {
		err = h.HandleSetWeeklySchedule(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of SetWeeklySchedule
func (v SetWeeklySchedule) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.SetWeeklyNumberOfTransitions.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SetWeeklyDayOfWeek.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SetWeeklyMode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SetWeeklyTransitionTime1.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if (v.SetWeeklyMode[0] & 0x01) == 0x01 {
		if tmp, err = v.SetWeeklyHeatSetpoint1.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if (v.SetWeeklyMode[0] & 0x02) == 0x02 {
		if tmp, err = v.SetWeeklyCoolSetpoint1.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SetWeeklyTransitionTime10.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if (v.SetWeeklyMode[0] & 0x01) == 0x01 {
		if tmp, err = v.SetWeeklyHeatSetpoint10.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if (v.SetWeeklyMode[0] & 0x02) == 0x02 {
		if tmp, err = v.SetWeeklyCoolSetpoint10.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the SetWeeklySchedule struct
func (v *SetWeeklySchedule) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.SetWeeklyNumberOfTransitions).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SetWeeklyDayOfWeek).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SetWeeklyMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SetWeeklyTransitionTime1).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if (v.SetWeeklyMode[0] & 0x01) == 0x01 {
		if b, err = (&v.SetWeeklyHeatSetpoint1).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if (v.SetWeeklyMode[0] & 0x02) == 0x02 {
		if b, err = (&v.SetWeeklyCoolSetpoint1).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if b, err = (&v.SetWeeklyTransitionTime10).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if (v.SetWeeklyMode[0] & 0x01) == 0x01 {
		if b, err = (&v.SetWeeklyHeatSetpoint10).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if (v.SetWeeklyMode[0] & 0x02) == 0x02 {
		if b, err = (&v.SetWeeklyCoolSetpoint10).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v SetWeeklySchedule) String() string {
	return zcl.Sprintf(
		"SetWeeklySchedule{"+zcl.StrJoin([]string{
			"SetWeeklyNumberOfTransitions(%v)",
			"SetWeeklyDayOfWeek(%v)",
			"SetWeeklyMode(%v)",
			"SetWeeklyTransitionTime1(%v)",
			"SetWeeklyHeatSetpoint1(%v)",
			"SetWeeklyCoolSetpoint1(%v)",
			"SetWeeklyTransitionTime10(%v)",
			"SetWeeklyHeatSetpoint10(%v)",
			"SetWeeklyCoolSetpoint10(%v)",
		}, " ")+"}",
		v.SetWeeklyNumberOfTransitions,
		v.SetWeeklyDayOfWeek,
		v.SetWeeklyMode,
		v.SetWeeklyTransitionTime1,
		v.SetWeeklyHeatSetpoint1,
		v.SetWeeklyCoolSetpoint1,
		v.SetWeeklyTransitionTime10,
		v.SetWeeklyHeatSetpoint10,
		v.SetWeeklyCoolSetpoint10,
	)
}

type GetWeeklySchedule struct {
	GetWeeklyDaysToReturn GetWeeklyDaysToReturn
	GetWeeklyModeToReturn GetWeeklyModeToReturn
}

type GetWeeklyScheduleHandler interface {
	HandleGetWeeklySchedule(frame Frame, cmd *GetWeeklySchedule) error
}

// GetWeeklyScheduleCommand is the Command ID of GetWeeklySchedule
const GetWeeklyScheduleCommand CommandID = 0x0002

// Values returns all values of GetWeeklySchedule
func (v *GetWeeklySchedule) Values() []zcl.Val {
	return []zcl.Val{
		&v.GetWeeklyDaysToReturn,
		&v.GetWeeklyModeToReturn,
	}
}

// Arguments returns all values of GetWeeklySchedule
func (v *GetWeeklySchedule) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "GetWeeklyDaysToReturn", Argument: &v.GetWeeklyDaysToReturn},
		{Name: "GetWeeklyModeToReturn", Argument: &v.GetWeeklyModeToReturn},
	}
}

// Name of the command
func (GetWeeklySchedule) Name() string { return `Get weekly schedule` }

// Description of the command
func (GetWeeklySchedule) Description() string { return `` }

// ID of the command
func (GetWeeklySchedule) ID() CommandID { return GetWeeklyScheduleCommand }

// Required
func (GetWeeklySchedule) Required() bool { return false }

// Cluster ID of the command
func (GetWeeklySchedule) Cluster() zcl.ClusterID { return ThermostatID }

// Direction of the command
func (GetWeeklySchedule) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (GetWeeklySchedule) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetWeeklySchedule) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

func (v *GetWeeklySchedule) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h GetWeeklyScheduleHandler
	if h, found = handler.(GetWeeklyScheduleHandler); found {
		err = h.HandleGetWeeklySchedule(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of GetWeeklySchedule
func (v GetWeeklySchedule) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.GetWeeklyDaysToReturn.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.GetWeeklyModeToReturn.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetWeeklySchedule struct
func (v *GetWeeklySchedule) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.GetWeeklyDaysToReturn).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.GetWeeklyModeToReturn).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetWeeklySchedule) String() string {
	return zcl.Sprintf(
		"GetWeeklySchedule{"+zcl.StrJoin([]string{
			"GetWeeklyDaysToReturn(%v)",
			"GetWeeklyModeToReturn(%v)",
		}, " ")+"}",
		v.GetWeeklyDaysToReturn,
		v.GetWeeklyModeToReturn,
	)
}

type ClearWeeklySchedule struct {
}

type ClearWeeklyScheduleHandler interface {
	HandleClearWeeklySchedule(frame Frame, cmd *ClearWeeklySchedule) error
}

// ClearWeeklyScheduleCommand is the Command ID of ClearWeeklySchedule
const ClearWeeklyScheduleCommand CommandID = 0x0003

// Values returns all values of ClearWeeklySchedule
func (v *ClearWeeklySchedule) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of ClearWeeklySchedule
func (v *ClearWeeklySchedule) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (ClearWeeklySchedule) Name() string { return `Clear weekly schedule` }

// Description of the command
func (ClearWeeklySchedule) Description() string { return `` }

// ID of the command
func (ClearWeeklySchedule) ID() CommandID { return ClearWeeklyScheduleCommand }

// Required
func (ClearWeeklySchedule) Required() bool { return false }

// Cluster ID of the command
func (ClearWeeklySchedule) Cluster() zcl.ClusterID { return ThermostatID }

// Direction of the command
func (ClearWeeklySchedule) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (ClearWeeklySchedule) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ClearWeeklySchedule) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

func (v *ClearWeeklySchedule) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ClearWeeklyScheduleHandler
	if h, found = handler.(ClearWeeklyScheduleHandler); found {
		err = h.HandleClearWeeklySchedule(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of ClearWeeklySchedule
func (v ClearWeeklySchedule) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the ClearWeeklySchedule struct
func (v *ClearWeeklySchedule) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ClearWeeklySchedule) String() string {
	return zcl.Sprintf(
		"ClearWeeklySchedule{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type GetRelayStatusLog struct {
}

type GetRelayStatusLogHandler interface {
	HandleGetRelayStatusLog(frame Frame, cmd *GetRelayStatusLog) error
}

// GetRelayStatusLogCommand is the Command ID of GetRelayStatusLog
const GetRelayStatusLogCommand CommandID = 0x0004

// Values returns all values of GetRelayStatusLog
func (v *GetRelayStatusLog) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of GetRelayStatusLog
func (v *GetRelayStatusLog) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (GetRelayStatusLog) Name() string { return `Get relay status log` }

// Description of the command
func (GetRelayStatusLog) Description() string { return `` }

// ID of the command
func (GetRelayStatusLog) ID() CommandID { return GetRelayStatusLogCommand }

// Required
func (GetRelayStatusLog) Required() bool { return false }

// Cluster ID of the command
func (GetRelayStatusLog) Cluster() zcl.ClusterID { return ThermostatID }

// Direction of the command
func (GetRelayStatusLog) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (GetRelayStatusLog) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetRelayStatusLog) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

func (v *GetRelayStatusLog) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h GetRelayStatusLogHandler
	if h, found = handler.(GetRelayStatusLogHandler); found {
		err = h.HandleGetRelayStatusLog(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of GetRelayStatusLog
func (v GetRelayStatusLog) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the GetRelayStatusLog struct
func (v *GetRelayStatusLog) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetRelayStatusLog) String() string {
	return zcl.Sprintf(
		"GetRelayStatusLog{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type GetWeeklyScheduleResponse struct {
	// SetWeeklyNumberOfTransitions how many individual transitions to expect for this sequence of
	// commands. If a device supports more than 10 transitions in its
	// schedule they can send this by sending more than 1 Set Weekly Schedule
	// command, each containing the separate information that the device
	// needs to set
	SetWeeklyNumberOfTransitions SetWeeklyNumberOfTransitions
	// SetWeeklyDayOfWeek day of the week at which all the transitions within the payload of the
	// command should be associated to
	SetWeeklyDayOfWeek SetWeeklyDayOfWeek
	// SetWeeklyMode which type of setpoint transition is present in the rest of the
	// command
	SetWeeklyMode             SetWeeklyMode
	SetWeeklyTransitionTime1  SetWeeklyTransitionTime1
	SetWeeklyHeatSetpoint1    SetWeeklyHeatSetpoint1
	SetWeeklyCoolSetpoint1    SetWeeklyCoolSetpoint1
	SetWeeklyTransitionTime10 SetWeeklyTransitionTime10
	SetWeeklyHeatSetpoint10   SetWeeklyHeatSetpoint10
	SetWeeklyCoolSetpoint10   SetWeeklyCoolSetpoint10
}

type GetWeeklyScheduleResponseHandler interface {
	HandleGetWeeklyScheduleResponse(frame Frame, cmd *GetWeeklyScheduleResponse) error
}

// GetWeeklyScheduleResponseCommand is the Command ID of GetWeeklyScheduleResponse
const GetWeeklyScheduleResponseCommand CommandID = 0x0000

// Values returns all values of GetWeeklyScheduleResponse
func (v *GetWeeklyScheduleResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.SetWeeklyNumberOfTransitions,
		&v.SetWeeklyDayOfWeek,
		&v.SetWeeklyMode,
		&v.SetWeeklyTransitionTime1,
		&v.SetWeeklyHeatSetpoint1,
		&v.SetWeeklyCoolSetpoint1,
		&v.SetWeeklyTransitionTime10,
		&v.SetWeeklyHeatSetpoint10,
		&v.SetWeeklyCoolSetpoint10,
	}
}

// Arguments returns all values of GetWeeklyScheduleResponse
func (v *GetWeeklyScheduleResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "SetWeeklyNumberOfTransitions", Argument: &v.SetWeeklyNumberOfTransitions},
		{Name: "SetWeeklyDayOfWeek", Argument: &v.SetWeeklyDayOfWeek},
		{Name: "SetWeeklyMode", Argument: &v.SetWeeklyMode},
		{Name: "SetWeeklyTransitionTime1", Argument: &v.SetWeeklyTransitionTime1},
		{Name: "SetWeeklyHeatSetpoint1", Argument: &v.SetWeeklyHeatSetpoint1},
		{Name: "SetWeeklyCoolSetpoint1", Argument: &v.SetWeeklyCoolSetpoint1},
		{Name: "SetWeeklyTransitionTime10", Argument: &v.SetWeeklyTransitionTime10},
		{Name: "SetWeeklyHeatSetpoint10", Argument: &v.SetWeeklyHeatSetpoint10},
		{Name: "SetWeeklyCoolSetpoint10", Argument: &v.SetWeeklyCoolSetpoint10},
	}
}

// Name of the command
func (GetWeeklyScheduleResponse) Name() string { return `Get weekly schedule response` }

// Description of the command
func (GetWeeklyScheduleResponse) Description() string { return `` }

// ID of the command
func (GetWeeklyScheduleResponse) ID() CommandID { return GetWeeklyScheduleResponseCommand }

// Required
func (GetWeeklyScheduleResponse) Required() bool { return false }

// Cluster ID of the command
func (GetWeeklyScheduleResponse) Cluster() zcl.ClusterID { return ThermostatID }

// Direction of the command
func (GetWeeklyScheduleResponse) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (GetWeeklyScheduleResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetWeeklyScheduleResponse) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *GetWeeklyScheduleResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h GetWeeklyScheduleResponseHandler
	if h, found = handler.(GetWeeklyScheduleResponseHandler); found {
		err = h.HandleGetWeeklyScheduleResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of GetWeeklyScheduleResponse
func (v GetWeeklyScheduleResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.SetWeeklyNumberOfTransitions.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SetWeeklyDayOfWeek.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SetWeeklyMode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SetWeeklyTransitionTime1.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if (v.SetWeeklyMode[0] & 0x01) == 0x01 {
		if tmp, err = v.SetWeeklyHeatSetpoint1.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if (v.SetWeeklyMode[0] & 0x02) == 0x02 {
		if tmp, err = v.SetWeeklyCoolSetpoint1.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.SetWeeklyTransitionTime10.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if (v.SetWeeklyMode[0] & 0x01) == 0x01 {
		if tmp, err = v.SetWeeklyHeatSetpoint10.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	if (v.SetWeeklyMode[0] & 0x02) == 0x02 {
		if tmp, err = v.SetWeeklyCoolSetpoint10.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetWeeklyScheduleResponse struct
func (v *GetWeeklyScheduleResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.SetWeeklyNumberOfTransitions).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SetWeeklyDayOfWeek).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SetWeeklyMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.SetWeeklyTransitionTime1).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if (v.SetWeeklyMode[0] & 0x01) == 0x01 {
		if b, err = (&v.SetWeeklyHeatSetpoint1).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if (v.SetWeeklyMode[0] & 0x02) == 0x02 {
		if b, err = (&v.SetWeeklyCoolSetpoint1).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if b, err = (&v.SetWeeklyTransitionTime10).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if (v.SetWeeklyMode[0] & 0x01) == 0x01 {
		if b, err = (&v.SetWeeklyHeatSetpoint10).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if (v.SetWeeklyMode[0] & 0x02) == 0x02 {
		if b, err = (&v.SetWeeklyCoolSetpoint10).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetWeeklyScheduleResponse) String() string {
	return zcl.Sprintf(
		"GetWeeklyScheduleResponse{"+zcl.StrJoin([]string{
			"SetWeeklyNumberOfTransitions(%v)",
			"SetWeeklyDayOfWeek(%v)",
			"SetWeeklyMode(%v)",
			"SetWeeklyTransitionTime1(%v)",
			"SetWeeklyHeatSetpoint1(%v)",
			"SetWeeklyCoolSetpoint1(%v)",
			"SetWeeklyTransitionTime10(%v)",
			"SetWeeklyHeatSetpoint10(%v)",
			"SetWeeklyCoolSetpoint10(%v)",
		}, " ")+"}",
		v.SetWeeklyNumberOfTransitions,
		v.SetWeeklyDayOfWeek,
		v.SetWeeklyMode,
		v.SetWeeklyTransitionTime1,
		v.SetWeeklyHeatSetpoint1,
		v.SetWeeklyCoolSetpoint1,
		v.SetWeeklyTransitionTime10,
		v.SetWeeklyHeatSetpoint10,
		v.SetWeeklyCoolSetpoint10,
	)
}

type GetRelayStatusLogResponse struct {
	// RelayStatusLogTimeOfDay minutes since midnight when the relay status was captured for this
	// associated log entry
	RelayStatusLogTimeOfDay RelayStatusLogTimeOfDay
	// RelayStatus status for thermostat when the log is captured. Each bit represents one
	// relay used by the thermostat. If the bit is on, the associated relay is
	// on and active. Each thermostat manufacturer can create its own mapping
	// between the bitmask and the associated relay
	RelayStatus RelayStatus
	// RelayStatusLocalTemperature temperature when the log is captured
	RelayStatusLocalTemperature RelayStatusLocalTemperature
	// RelayStatusHumidity humidity when the log was captured
	RelayStatusHumidity RelayStatusHumidity
	// RelayStatusSetpoint target setpoint temperature when the log is captured
	RelayStatusSetpoint RelayStatusSetpoint
	// RelayStatusUnreadEntries umber of unread entries within the thermostat log system
	RelayStatusUnreadEntries RelayStatusUnreadEntries
}

type GetRelayStatusLogResponseHandler interface {
	HandleGetRelayStatusLogResponse(frame Frame, cmd *GetRelayStatusLogResponse) error
}

// GetRelayStatusLogResponseCommand is the Command ID of GetRelayStatusLogResponse
const GetRelayStatusLogResponseCommand CommandID = 0x0001

// Values returns all values of GetRelayStatusLogResponse
func (v *GetRelayStatusLogResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.RelayStatusLogTimeOfDay,
		&v.RelayStatus,
		&v.RelayStatusLocalTemperature,
		&v.RelayStatusHumidity,
		&v.RelayStatusSetpoint,
		&v.RelayStatusUnreadEntries,
	}
}

// Arguments returns all values of GetRelayStatusLogResponse
func (v *GetRelayStatusLogResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "RelayStatusLogTimeOfDay", Argument: &v.RelayStatusLogTimeOfDay},
		{Name: "RelayStatus", Argument: &v.RelayStatus},
		{Name: "RelayStatusLocalTemperature", Argument: &v.RelayStatusLocalTemperature},
		{Name: "RelayStatusHumidity", Argument: &v.RelayStatusHumidity},
		{Name: "RelayStatusSetpoint", Argument: &v.RelayStatusSetpoint},
		{Name: "RelayStatusUnreadEntries", Argument: &v.RelayStatusUnreadEntries},
	}
}

// Name of the command
func (GetRelayStatusLogResponse) Name() string { return `Get relay status log response` }

// Description of the command
func (GetRelayStatusLogResponse) Description() string { return `` }

// ID of the command
func (GetRelayStatusLogResponse) ID() CommandID { return GetRelayStatusLogResponseCommand }

// Required
func (GetRelayStatusLogResponse) Required() bool { return false }

// Cluster ID of the command
func (GetRelayStatusLogResponse) Cluster() zcl.ClusterID { return ThermostatID }

// Direction of the command
func (GetRelayStatusLogResponse) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (GetRelayStatusLogResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetRelayStatusLogResponse) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *GetRelayStatusLogResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h GetRelayStatusLogResponseHandler
	if h, found = handler.(GetRelayStatusLogResponseHandler); found {
		err = h.HandleGetRelayStatusLogResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of GetRelayStatusLogResponse
func (v GetRelayStatusLogResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.RelayStatusLogTimeOfDay.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.RelayStatus.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.RelayStatusLocalTemperature.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.RelayStatusHumidity.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.RelayStatusSetpoint.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.RelayStatusUnreadEntries.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetRelayStatusLogResponse struct
func (v *GetRelayStatusLogResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.RelayStatusLogTimeOfDay).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.RelayStatus).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.RelayStatusLocalTemperature).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.RelayStatusHumidity).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.RelayStatusSetpoint).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.RelayStatusUnreadEntries).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetRelayStatusLogResponse) String() string {
	return zcl.Sprintf(
		"GetRelayStatusLogResponse{"+zcl.StrJoin([]string{
			"RelayStatusLogTimeOfDay(%v)",
			"RelayStatus(%v)",
			"RelayStatusLocalTemperature(%v)",
			"RelayStatusHumidity(%v)",
			"RelayStatusSetpoint(%v)",
			"RelayStatusUnreadEntries(%v)",
		}, " ")+"}",
		v.RelayStatusLogTimeOfDay,
		v.RelayStatus,
		v.RelayStatusLocalTemperature,
		v.RelayStatusHumidity,
		v.RelayStatusSetpoint,
		v.RelayStatusUnreadEntries,
	)
}
