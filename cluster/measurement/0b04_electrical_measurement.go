package measurement

import "hemtjan.st/zcl"

// ElectricalMeasurement
const ElectricalMeasurementID zcl.ClusterID = 2820

var ElectricalMeasurementCluster = zcl.Cluster{
	Name: "Electrical Measurement",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		GetProfileInfoCommand:        func() zcl.Command { return new(GetProfileInfo) },
		GetMeasurementProfileCommand: func() zcl.Command { return new(GetMeasurementProfile) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		GetProfileInfoResponseCommand:        func() zcl.Command { return new(GetProfileInfoResponse) },
		GetMeasurementProfileResponseCommand: func() zcl.Command { return new(GetMeasurementProfileResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		ElectricalMeasurementTypeAttr:             func() zcl.Attr { return new(ElectricalMeasurementType) },
		DcVoltageAttr:                             func() zcl.Attr { return new(DcVoltage) },
		DcVoltageMinAttr:                          func() zcl.Attr { return new(DcVoltageMin) },
		DcVoltageMaxAttr:                          func() zcl.Attr { return new(DcVoltageMax) },
		DcCurrentAttr:                             func() zcl.Attr { return new(DcCurrent) },
		DcCurrentMinAttr:                          func() zcl.Attr { return new(DcCurrentMin) },
		DcCurrentMaxAttr:                          func() zcl.Attr { return new(DcCurrentMax) },
		DcPowerAttr:                               func() zcl.Attr { return new(DcPower) },
		DcPowerMinAttr:                            func() zcl.Attr { return new(DcPowerMin) },
		DcPowerMaxAttr:                            func() zcl.Attr { return new(DcPowerMax) },
		DcVoltageMultiplierAttr:                   func() zcl.Attr { return new(DcVoltageMultiplier) },
		DcVoltageDivisorAttr:                      func() zcl.Attr { return new(DcVoltageDivisor) },
		DcCurrentMultiplierAttr:                   func() zcl.Attr { return new(DcCurrentMultiplier) },
		DcCurrentDivisorAttr:                      func() zcl.Attr { return new(DcCurrentDivisor) },
		DcPowerMultiplierAttr:                     func() zcl.Attr { return new(DcPowerMultiplier) },
		DcPowerDivisorAttr:                        func() zcl.Attr { return new(DcPowerDivisor) },
		AcFrequencyAttr:                           func() zcl.Attr { return new(AcFrequency) },
		AcFrequencyMinAttr:                        func() zcl.Attr { return new(AcFrequencyMin) },
		AcFrequencyMaxAttr:                        func() zcl.Attr { return new(AcFrequencyMax) },
		NeutralCurrentAttr:                        func() zcl.Attr { return new(NeutralCurrent) },
		TotalActivePowerAttr:                      func() zcl.Attr { return new(TotalActivePower) },
		TotalReactivePowerAttr:                    func() zcl.Attr { return new(TotalReactivePower) },
		TotalApparentPowerAttr:                    func() zcl.Attr { return new(TotalApparentPower) },
		Measured1StHarmonicCurrentAttr:            func() zcl.Attr { return new(Measured1StHarmonicCurrent) },
		Measured3RdHarmonicCurrentAttr:            func() zcl.Attr { return new(Measured3RdHarmonicCurrent) },
		Measured5ThHarmonicCurrentAttr:            func() zcl.Attr { return new(Measured5ThHarmonicCurrent) },
		Measured7ThHarmonicCurrentAttr:            func() zcl.Attr { return new(Measured7ThHarmonicCurrent) },
		Measured9ThHarmonicCurrentAttr:            func() zcl.Attr { return new(Measured9ThHarmonicCurrent) },
		Measured11ThHarmonicCurrentAttr:           func() zcl.Attr { return new(Measured11ThHarmonicCurrent) },
		MeasuredPhase1StHarmonicCurrentAttr:       func() zcl.Attr { return new(MeasuredPhase1StHarmonicCurrent) },
		MeasuredPhase3RdHarmonicCurrentAttr:       func() zcl.Attr { return new(MeasuredPhase3RdHarmonicCurrent) },
		MeasuredPhase5ThHarmonicCurrentAttr:       func() zcl.Attr { return new(MeasuredPhase5ThHarmonicCurrent) },
		MeasuredPhase7ThHarmonicCurrentAttr:       func() zcl.Attr { return new(MeasuredPhase7ThHarmonicCurrent) },
		MeasuredPhase9ThHarmonicCurrentAttr:       func() zcl.Attr { return new(MeasuredPhase9ThHarmonicCurrent) },
		MeasuredPhase11ThHarmonicCurrentAttr:      func() zcl.Attr { return new(MeasuredPhase11ThHarmonicCurrent) },
		AcFrequencyMultiplierAttr:                 func() zcl.Attr { return new(AcFrequencyMultiplier) },
		AcFrequencyDivisorAttr:                    func() zcl.Attr { return new(AcFrequencyDivisor) },
		PowerMultiplierAttr:                       func() zcl.Attr { return new(PowerMultiplier) },
		PowerDivisorAttr:                          func() zcl.Attr { return new(PowerDivisor) },
		HarmonicCurrentMultiplierAttr:             func() zcl.Attr { return new(HarmonicCurrentMultiplier) },
		PhaseHarmonicCurrentMultiplierAttr:        func() zcl.Attr { return new(PhaseHarmonicCurrentMultiplier) },
		LineCurrentAttr:                           func() zcl.Attr { return new(LineCurrent) },
		ActiveCurrentAttr:                         func() zcl.Attr { return new(ActiveCurrent) },
		ReactiveCurrentAttr:                       func() zcl.Attr { return new(ReactiveCurrent) },
		RmsVoltageAttr:                            func() zcl.Attr { return new(RmsVoltage) },
		RmsVoltageMinAttr:                         func() zcl.Attr { return new(RmsVoltageMin) },
		RmsVoltageMaxAttr:                         func() zcl.Attr { return new(RmsVoltageMax) },
		RmsCurrentAttr:                            func() zcl.Attr { return new(RmsCurrent) },
		RmsCurrentMinAttr:                         func() zcl.Attr { return new(RmsCurrentMin) },
		RmsCurrentMaxAttr:                         func() zcl.Attr { return new(RmsCurrentMax) },
		ActivePowerAttr:                           func() zcl.Attr { return new(ActivePower) },
		ActivePowerMinAttr:                        func() zcl.Attr { return new(ActivePowerMin) },
		ActivePowerMaxAttr:                        func() zcl.Attr { return new(ActivePowerMax) },
		ReactivePowerAttr:                         func() zcl.Attr { return new(ReactivePower) },
		ApparentPowerAttr:                         func() zcl.Attr { return new(ApparentPower) },
		PowerFactorAttr:                           func() zcl.Attr { return new(PowerFactor) },
		AverageRmsVoltageMeasurementPeriodAttr:    func() zcl.Attr { return new(AverageRmsVoltageMeasurementPeriod) },
		AverageRmsOverVoltageCounterAttr:          func() zcl.Attr { return new(AverageRmsOverVoltageCounter) },
		AverageRmsUnderVoltageCounterAttr:         func() zcl.Attr { return new(AverageRmsUnderVoltageCounter) },
		RmsExtremeOverVoltagePeriodAttr:           func() zcl.Attr { return new(RmsExtremeOverVoltagePeriod) },
		RmsExtremeUnderVoltagePeriodAttr:          func() zcl.Attr { return new(RmsExtremeUnderVoltagePeriod) },
		RmsVoltageSagPeriodAttr:                   func() zcl.Attr { return new(RmsVoltageSagPeriod) },
		RmsVoltageSwellPeriodAttr:                 func() zcl.Attr { return new(RmsVoltageSwellPeriod) },
		AcVoltageMultiplierAttr:                   func() zcl.Attr { return new(AcVoltageMultiplier) },
		AcVoltageDivisorAttr:                      func() zcl.Attr { return new(AcVoltageDivisor) },
		AcCurrentMultiplierAttr:                   func() zcl.Attr { return new(AcCurrentMultiplier) },
		AcCurrentDivisorAttr:                      func() zcl.Attr { return new(AcCurrentDivisor) },
		AcPowerMultiplierAttr:                     func() zcl.Attr { return new(AcPowerMultiplier) },
		AcPowerDivisorAttr:                        func() zcl.Attr { return new(AcPowerDivisor) },
		DcOverloadAlarmsMaskAttr:                  func() zcl.Attr { return new(DcOverloadAlarmsMask) },
		DcVoltageOverloadAttr:                     func() zcl.Attr { return new(DcVoltageOverload) },
		DcCurrentOverloadAttr:                     func() zcl.Attr { return new(DcCurrentOverload) },
		AcAlarmsMaskAttr:                          func() zcl.Attr { return new(AcAlarmsMask) },
		AcVoltageOverloadAttr:                     func() zcl.Attr { return new(AcVoltageOverload) },
		AcCurrentOverloadAttr:                     func() zcl.Attr { return new(AcCurrentOverload) },
		AcActivePowerOverloadAttr:                 func() zcl.Attr { return new(AcActivePowerOverload) },
		AcReactivePowerOverloadAttr:               func() zcl.Attr { return new(AcReactivePowerOverload) },
		AverageRmsOverVoltageAttr:                 func() zcl.Attr { return new(AverageRmsOverVoltage) },
		AverageRmsUnderVoltageAttr:                func() zcl.Attr { return new(AverageRmsUnderVoltage) },
		RmsExtremeOverVoltageAttr:                 func() zcl.Attr { return new(RmsExtremeOverVoltage) },
		RmsExtremeUnderVoltageAttr:                func() zcl.Attr { return new(RmsExtremeUnderVoltage) },
		RmsVoltageSagAttr:                         func() zcl.Attr { return new(RmsVoltageSag) },
		RmsVoltageSwellAttr:                       func() zcl.Attr { return new(RmsVoltageSwell) },
		LineCurrentPhBAttr:                        func() zcl.Attr { return new(LineCurrentPhB) },
		ActiveCurrentPhBAttr:                      func() zcl.Attr { return new(ActiveCurrentPhB) },
		ReactiveCurrentPhBAttr:                    func() zcl.Attr { return new(ReactiveCurrentPhB) },
		RmsVoltagePhBAttr:                         func() zcl.Attr { return new(RmsVoltagePhB) },
		RmsVoltageMinPhBAttr:                      func() zcl.Attr { return new(RmsVoltageMinPhB) },
		RmsVoltageMaxPhBAttr:                      func() zcl.Attr { return new(RmsVoltageMaxPhB) },
		RmsCurrentPhBAttr:                         func() zcl.Attr { return new(RmsCurrentPhB) },
		RmsCurrentMinPhBAttr:                      func() zcl.Attr { return new(RmsCurrentMinPhB) },
		RmsCurrentMaxPhBAttr:                      func() zcl.Attr { return new(RmsCurrentMaxPhB) },
		ActivePowerPhBAttr:                        func() zcl.Attr { return new(ActivePowerPhB) },
		ActivePowerMinPhBAttr:                     func() zcl.Attr { return new(ActivePowerMinPhB) },
		ActivePowerMaxPhBAttr:                     func() zcl.Attr { return new(ActivePowerMaxPhB) },
		ReactivePowerPhBAttr:                      func() zcl.Attr { return new(ReactivePowerPhB) },
		ApparentPowerPhBAttr:                      func() zcl.Attr { return new(ApparentPowerPhB) },
		PowerFactorPhBAttr:                        func() zcl.Attr { return new(PowerFactorPhB) },
		AverageRmsVoltageMeasurementPeriodPhBAttr: func() zcl.Attr { return new(AverageRmsVoltageMeasurementPeriodPhB) },
		AverageRmsOverVoltageCounterPhBAttr:       func() zcl.Attr { return new(AverageRmsOverVoltageCounterPhB) },
		AverageRmsUnderVoltageCounterPhBAttr:      func() zcl.Attr { return new(AverageRmsUnderVoltageCounterPhB) },
		RmsExtremeOverVoltagePeriodPhBAttr:        func() zcl.Attr { return new(RmsExtremeOverVoltagePeriodPhB) },
		RmsExtremeUnderVoltagePeriodPhBAttr:       func() zcl.Attr { return new(RmsExtremeUnderVoltagePeriodPhB) },
		RmsVoltageSagPeriodPhBAttr:                func() zcl.Attr { return new(RmsVoltageSagPeriodPhB) },
		RmsVoltageSwellPeriodPhBAttr:              func() zcl.Attr { return new(RmsVoltageSwellPeriodPhB) },
		LineCurrentPhCAttr:                        func() zcl.Attr { return new(LineCurrentPhC) },
		ActiveCurrentPhCAttr:                      func() zcl.Attr { return new(ActiveCurrentPhC) },
		ReactiveCurrentPhCAttr:                    func() zcl.Attr { return new(ReactiveCurrentPhC) },
		RmsVoltagePhCAttr:                         func() zcl.Attr { return new(RmsVoltagePhC) },
		RmsVoltageMinPhCAttr:                      func() zcl.Attr { return new(RmsVoltageMinPhC) },
		RmsVoltageMaxPhCAttr:                      func() zcl.Attr { return new(RmsVoltageMaxPhC) },
		RmsCurrentPhCAttr:                         func() zcl.Attr { return new(RmsCurrentPhC) },
		RmsCurrentMinPhCAttr:                      func() zcl.Attr { return new(RmsCurrentMinPhC) },
		RmsCurrentMaxPhCAttr:                      func() zcl.Attr { return new(RmsCurrentMaxPhC) },
		ActivePowerPhCAttr:                        func() zcl.Attr { return new(ActivePowerPhC) },
		ActivePowerMinPhCAttr:                     func() zcl.Attr { return new(ActivePowerMinPhC) },
		ActivePowerMaxPhCAttr:                     func() zcl.Attr { return new(ActivePowerMaxPhC) },
		ReactivePowerPhCAttr:                      func() zcl.Attr { return new(ReactivePowerPhC) },
		ApparentPowerPhCAttr:                      func() zcl.Attr { return new(ApparentPowerPhC) },
		PowerFactorPhCAttr:                        func() zcl.Attr { return new(PowerFactorPhC) },
		AverageRmsVoltageMeasurementPeriodPhCAttr: func() zcl.Attr { return new(AverageRmsVoltageMeasurementPeriodPhC) },
		AverageRmsOverVoltageCounterPhCAttr:       func() zcl.Attr { return new(AverageRmsOverVoltageCounterPhC) },
		AverageRmsUnderVoltageCounterPhCAttr:      func() zcl.Attr { return new(AverageRmsUnderVoltageCounterPhC) },
		RmsExtremeOverVoltagePeriodPhCAttr:        func() zcl.Attr { return new(RmsExtremeOverVoltagePeriodPhC) },
		RmsExtremeUnderVoltagePeriodPhCAttr:       func() zcl.Attr { return new(RmsExtremeUnderVoltagePeriodPhC) },
		RmsVoltageSagPeriodPhCAttr:                func() zcl.Attr { return new(RmsVoltageSagPeriodPhC) },
		RmsVoltageSwellPeriodPhCAttr:              func() zcl.Attr { return new(RmsVoltageSwellPeriodPhC) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type GetProfileInfo struct {
}

type GetProfileInfoHandler interface {
	HandleGetProfileInfo(frame Frame, cmd *GetProfileInfo) (*GetMeasurementProfileResponse, error)
}

// GetProfileInfoCommand is the Command ID of GetProfileInfo
const GetProfileInfoCommand CommandID = 0x0000

// Values returns all values of GetProfileInfo
func (v *GetProfileInfo) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of GetProfileInfo
func (v *GetProfileInfo) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (GetProfileInfo) Name() string { return `Get Profile Info` }

// Description of the command
func (GetProfileInfo) Description() string { return `` }

// ID of the command
func (GetProfileInfo) ID() CommandID { return GetProfileInfoCommand }

// Required
func (GetProfileInfo) Required() bool { return false }

// Cluster ID of the command
func (GetProfileInfo) Cluster() zcl.ClusterID { return ElectricalMeasurementID }

// Direction of the command
func (GetProfileInfo) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (GetProfileInfo) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetProfileInfo) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *GetProfileInfo) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h GetProfileInfoHandler
	if h, found = handler.(GetProfileInfoHandler); found {
		rsp, err = h.HandleGetProfileInfo(frame, v)
	} else {
		rsp = &zcl.DefaultResponse{Command: v.ID(), Status: zcl.UnsupClusterCommand}
	}
	return
}

// MarshalZcl returns the wire format representation of GetProfileInfo
func (v GetProfileInfo) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the GetProfileInfo struct
func (v *GetProfileInfo) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetProfileInfo) String() string {
	return zcl.Sprintf(
		"GetProfileInfo{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type GetMeasurementProfile struct {
	AttributeId AttributeId
	// StartTime is the end time of the most chronologically recent interval being requested.
	// Example: Data collected from 2:00 PM to 3:00 PM would be specified as a 3:00 PM interval (end time).
	StartTime StartTime
	// NumberOfIntervals is the number of intervals requested or returned
	NumberOfIntervals NumberOfIntervals
}

type GetMeasurementProfileHandler interface {
	HandleGetMeasurementProfile(frame Frame, cmd *GetMeasurementProfile) (*GetMeasurementProfileResponse, error)
}

// GetMeasurementProfileCommand is the Command ID of GetMeasurementProfile
const GetMeasurementProfileCommand CommandID = 0x0001

// Values returns all values of GetMeasurementProfile
func (v *GetMeasurementProfile) Values() []zcl.Val {
	return []zcl.Val{
		&v.AttributeId,
		&v.StartTime,
		&v.NumberOfIntervals,
	}
}

// Arguments returns all values of GetMeasurementProfile
func (v *GetMeasurementProfile) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "AttributeId", Argument: &v.AttributeId},
		{Name: "StartTime", Argument: &v.StartTime},
		{Name: "NumberOfIntervals", Argument: &v.NumberOfIntervals},
	}
}

// Name of the command
func (GetMeasurementProfile) Name() string { return `Get Measurement Profile` }

// Description of the command
func (GetMeasurementProfile) Description() string { return `` }

// ID of the command
func (GetMeasurementProfile) ID() CommandID { return GetMeasurementProfileCommand }

// Required
func (GetMeasurementProfile) Required() bool { return false }

// Cluster ID of the command
func (GetMeasurementProfile) Cluster() zcl.ClusterID { return ElectricalMeasurementID }

// Direction of the command
func (GetMeasurementProfile) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (GetMeasurementProfile) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetMeasurementProfile) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *GetMeasurementProfile) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h GetMeasurementProfileHandler
	if h, found = handler.(GetMeasurementProfileHandler); found {
		rsp, err = h.HandleGetMeasurementProfile(frame, v)
	} else {
		rsp = &zcl.DefaultResponse{Command: v.ID(), Status: zcl.UnsupClusterCommand}
	}
	return
}

// MarshalZcl returns the wire format representation of GetMeasurementProfile
func (v GetMeasurementProfile) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.AttributeId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StartTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.NumberOfIntervals.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetMeasurementProfile struct
func (v *GetMeasurementProfile) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

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

// String returns a log-friendly string representation of the struct
func (v GetMeasurementProfile) String() string {
	return zcl.Sprintf(
		"GetMeasurementProfile{"+zcl.StrJoin([]string{
			"AttributeId(%v)",
			"StartTime(%v)",
			"NumberOfIntervals(%v)",
		}, " ")+"}",
		v.AttributeId,
		v.StartTime,
		v.NumberOfIntervals,
	)
}

type GetProfileInfoResponse struct {
	// ProfileCount is the total number of supported profiles
	ProfileCount ProfileCount
	// ProfileIntervalPeriod represents the interval or time frame used to capture parameter for profiling purposes
	ProfileIntervalPeriod ProfileIntervalPeriod
	// MaxNumberOfIntervals represents the maximum number of intervals the device is capable of returning in one command.
	// It is required MaxNumberofIntervals fit within the default Fragmentation ASDU size of 128 bytes,
	// or an optionally agreed upon larger Fragmentation ASDU size supported by both devices.
	MaxNumberOfIntervals MaxNumberOfIntervals
	// ListOfAttributes is the list of attributes being profiled
	ListOfAttributes ListOfAttributes
}

type GetProfileInfoResponseHandler interface {
	HandleGetProfileInfoResponse(frame Frame, cmd *GetProfileInfoResponse) error
}

// GetProfileInfoResponseCommand is the Command ID of GetProfileInfoResponse
const GetProfileInfoResponseCommand CommandID = 0x0000

// Values returns all values of GetProfileInfoResponse
func (v *GetProfileInfoResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.ProfileCount,
		&v.ProfileIntervalPeriod,
		&v.MaxNumberOfIntervals,
		&v.ListOfAttributes,
	}
}

// Arguments returns all values of GetProfileInfoResponse
func (v *GetProfileInfoResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "ProfileCount", Argument: &v.ProfileCount},
		{Name: "ProfileIntervalPeriod", Argument: &v.ProfileIntervalPeriod},
		{Name: "MaxNumberOfIntervals", Argument: &v.MaxNumberOfIntervals},
		{Name: "ListOfAttributes", Argument: &v.ListOfAttributes},
	}
}

// Name of the command
func (GetProfileInfoResponse) Name() string { return `Get Profile Info Response` }

// Description of the command
func (GetProfileInfoResponse) Description() string { return `` }

// ID of the command
func (GetProfileInfoResponse) ID() CommandID { return GetProfileInfoResponseCommand }

// Required
func (GetProfileInfoResponse) Required() bool { return false }

// Cluster ID of the command
func (GetProfileInfoResponse) Cluster() zcl.ClusterID { return ElectricalMeasurementID }

// Direction of the command
func (GetProfileInfoResponse) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (GetProfileInfoResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetProfileInfoResponse) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *GetProfileInfoResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h GetProfileInfoResponseHandler
	if h, found = handler.(GetProfileInfoResponseHandler); found {
		err = h.HandleGetProfileInfoResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of GetProfileInfoResponse
func (v GetProfileInfoResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.ProfileCount.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ProfileIntervalPeriod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.MaxNumberOfIntervals.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ListOfAttributes.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetProfileInfoResponse struct
func (v *GetProfileInfoResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

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

// String returns a log-friendly string representation of the struct
func (v GetProfileInfoResponse) String() string {
	return zcl.Sprintf(
		"GetProfileInfoResponse{"+zcl.StrJoin([]string{
			"ProfileCount(%v)",
			"ProfileIntervalPeriod(%v)",
			"MaxNumberOfIntervals(%v)",
			"ListOfAttributes(%v)",
		}, " ")+"}",
		v.ProfileCount,
		v.ProfileIntervalPeriod,
		v.MaxNumberOfIntervals,
		v.ListOfAttributes,
	)
}

type GetMeasurementProfileResponse struct {
	// StartTime is the end time of the most chronologically recent interval being requested.
	// Example: Data collected from 2:00 PM to 3:00 PM would be specified as a 3:00 PM interval (end time).
	StartTime                 StartTime
	MeasurementResponseStatus MeasurementResponseStatus
	// ProfileIntervalPeriod represents the interval or time frame used to capture parameter for profiling purposes
	ProfileIntervalPeriod ProfileIntervalPeriod
	// NumberOfIntervals is the number of intervals requested or returned
	NumberOfIntervals NumberOfIntervals
	AttributeId       AttributeId
	// Intervals is a series of interval data captured using the period specified by the ProfileIntervalPeriod field. The
	// content of the interval data depend of the type of information requested using the AttributeID field in the Get
	// Measurement Profile Command. Data is organized in a reverse chronological order, the oldest intervals are
	// transmitted first and the newest interval is transmitted last. Invalid intervals should be marked as 0xFFFF.
	// For scaling and data type use the respective attribute set as defined above in attribute sets.
	Intervals Intervals
}

type GetMeasurementProfileResponseHandler interface {
	HandleGetMeasurementProfileResponse(frame Frame, cmd *GetMeasurementProfileResponse) error
}

// GetMeasurementProfileResponseCommand is the Command ID of GetMeasurementProfileResponse
const GetMeasurementProfileResponseCommand CommandID = 0x0001

// Values returns all values of GetMeasurementProfileResponse
func (v *GetMeasurementProfileResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.StartTime,
		&v.MeasurementResponseStatus,
		&v.ProfileIntervalPeriod,
		&v.NumberOfIntervals,
		&v.AttributeId,
		&v.Intervals,
	}
}

// Arguments returns all values of GetMeasurementProfileResponse
func (v *GetMeasurementProfileResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "StartTime", Argument: &v.StartTime},
		{Name: "MeasurementResponseStatus", Argument: &v.MeasurementResponseStatus},
		{Name: "ProfileIntervalPeriod", Argument: &v.ProfileIntervalPeriod},
		{Name: "NumberOfIntervals", Argument: &v.NumberOfIntervals},
		{Name: "AttributeId", Argument: &v.AttributeId},
		{Name: "Intervals", Argument: &v.Intervals},
	}
}

// Name of the command
func (GetMeasurementProfileResponse) Name() string { return `Get Measurement Profile Response` }

// Description of the command
func (GetMeasurementProfileResponse) Description() string { return `` }

// ID of the command
func (GetMeasurementProfileResponse) ID() CommandID { return GetMeasurementProfileResponseCommand }

// Required
func (GetMeasurementProfileResponse) Required() bool { return false }

// Cluster ID of the command
func (GetMeasurementProfileResponse) Cluster() zcl.ClusterID { return ElectricalMeasurementID }

// Direction of the command
func (GetMeasurementProfileResponse) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (GetMeasurementProfileResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetMeasurementProfileResponse) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *GetMeasurementProfileResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h GetMeasurementProfileResponseHandler
	if h, found = handler.(GetMeasurementProfileResponseHandler); found {
		err = h.HandleGetMeasurementProfileResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of GetMeasurementProfileResponse
func (v GetMeasurementProfileResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.StartTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.MeasurementResponseStatus.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ProfileIntervalPeriod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.NumberOfIntervals.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.AttributeId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Intervals.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetMeasurementProfileResponse struct
func (v *GetMeasurementProfileResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.StartTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.MeasurementResponseStatus).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ProfileIntervalPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberOfIntervals).UnmarshalZcl(b); err != nil {
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

// String returns a log-friendly string representation of the struct
func (v GetMeasurementProfileResponse) String() string {
	return zcl.Sprintf(
		"GetMeasurementProfileResponse{"+zcl.StrJoin([]string{
			"StartTime(%v)",
			"MeasurementResponseStatus(%v)",
			"ProfileIntervalPeriod(%v)",
			"NumberOfIntervals(%v)",
			"AttributeId(%v)",
			"Intervals(%v)",
		}, " ")+"}",
		v.StartTime,
		v.MeasurementResponseStatus,
		v.ProfileIntervalPeriod,
		v.NumberOfIntervals,
		v.AttributeId,
		v.Intervals,
	)
}
