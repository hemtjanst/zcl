package general

import "hemtjan.st/zcl"

// Location
const LocationID zcl.ClusterID = 11

var LocationCluster = zcl.Cluster{
	Name: "Location",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		SetAbsoluteLocationCommand:    func() zcl.Command { return new(SetAbsoluteLocation) },
		SetDeviceConfigurationCommand: func() zcl.Command { return new(SetDeviceConfiguration) },
		GetDeviceConfigurationCommand: func() zcl.Command { return new(GetDeviceConfiguration) },
		GetLocationDataCommand:        func() zcl.Command { return new(GetLocationData) },
		RssiResponseCommand:           func() zcl.Command { return new(RssiResponse) },
		SendPingsCommand:              func() zcl.Command { return new(SendPings) },
		AnchorNodeAnnounceCommand:     func() zcl.Command { return new(AnchorNodeAnnounce) },
		DistanceMeasureCommand:        func() zcl.Command { return new(DistanceMeasure) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		DeviceConfigurationResponseCommand:     func() zcl.Command { return new(DeviceConfigurationResponse) },
		LocationDataResponseCommand:            func() zcl.Command { return new(LocationDataResponse) },
		LocationDataNotificationCommand:        func() zcl.Command { return new(LocationDataNotification) },
		CompactLocationDataNotificationCommand: func() zcl.Command { return new(CompactLocationDataNotification) },
		RssiPingCommand:                        func() zcl.Command { return new(RssiPing) },
		RssiRequestCommand:                     func() zcl.Command { return new(RssiRequest) },
		ReportRssiMeasurementsCommand:          func() zcl.Command { return new(ReportRssiMeasurements) },
		RequestOwnLocationCommand:              func() zcl.Command { return new(RequestOwnLocation) },
		DistanceMeasureResponseCommand:         func() zcl.Command { return new(DistanceMeasureResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		LocationTypeAttr:           func() zcl.Attr { return new(LocationType) },
		LocationMethodAttr:         func() zcl.Attr { return new(LocationMethod) },
		LocationAgeAttr:            func() zcl.Attr { return new(LocationAge) },
		QualityMeasureAttr:         func() zcl.Attr { return new(QualityMeasure) },
		NumberOfDevicesAttr:        func() zcl.Attr { return new(NumberOfDevices) },
		XCoordinateAttr:            func() zcl.Attr { return new(XCoordinate) },
		YCoordinateAttr:            func() zcl.Attr { return new(YCoordinate) },
		ZCoordinateAttr:            func() zcl.Attr { return new(ZCoordinate) },
		PowerAttr:                  func() zcl.Attr { return new(Power) },
		PathLossExponentAttr:       func() zcl.Attr { return new(PathLossExponent) },
		ReportingPeriodAttr:        func() zcl.Attr { return new(ReportingPeriod) },
		CalculationPeriodAttr:      func() zcl.Attr { return new(CalculationPeriod) },
		NumberRssiMeasurementsAttr: func() zcl.Attr { return new(NumberRssiMeasurements) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type SetAbsoluteLocation struct {
	XCoordinate XCoordinate
	YCoordinate YCoordinate
	ZCoordinate ZCoordinate
	Power       Power
	// PathLossExponent is the rate at which the signal power decays with increasing distance
	PathLossExponent PathLossExponent
}

type SetAbsoluteLocationHandler interface {
	HandleSetAbsoluteLocation(frame Frame, cmd *SetAbsoluteLocation) error
}

// SetAbsoluteLocationCommand is the Command ID of SetAbsoluteLocation
const SetAbsoluteLocationCommand CommandID = 0x0000

// Values returns all values of SetAbsoluteLocation
func (v *SetAbsoluteLocation) Values() []zcl.Val {
	return []zcl.Val{
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
		&v.Power,
		&v.PathLossExponent,
	}
}

// Arguments returns all values of SetAbsoluteLocation
func (v *SetAbsoluteLocation) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "XCoordinate", Argument: &v.XCoordinate},
		{Name: "YCoordinate", Argument: &v.YCoordinate},
		{Name: "ZCoordinate", Argument: &v.ZCoordinate},
		{Name: "Power", Argument: &v.Power},
		{Name: "PathLossExponent", Argument: &v.PathLossExponent},
	}
}

// Name of the command
func (SetAbsoluteLocation) Name() string { return `Set Absolute Location` }

// Description of the command
func (SetAbsoluteLocation) Description() string { return `` }

// ID of the command
func (SetAbsoluteLocation) ID() CommandID { return SetAbsoluteLocationCommand }

// Required
func (SetAbsoluteLocation) Required() bool { return true }

// Cluster ID of the command
func (SetAbsoluteLocation) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (SetAbsoluteLocation) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (SetAbsoluteLocation) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SetAbsoluteLocation) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *SetAbsoluteLocation) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h SetAbsoluteLocationHandler
	if h, found = handler.(SetAbsoluteLocationHandler); found {
		err = h.HandleSetAbsoluteLocation(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of SetAbsoluteLocation
func (v SetAbsoluteLocation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Power.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.PathLossExponent.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the SetAbsoluteLocation struct
func (v *SetAbsoluteLocation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Power).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PathLossExponent).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v SetAbsoluteLocation) String() string {
	return zcl.Sprintf(
		"SetAbsoluteLocation{"+zcl.StrJoin([]string{
			"XCoordinate(%v)",
			"YCoordinate(%v)",
			"ZCoordinate(%v)",
			"Power(%v)",
			"PathLossExponent(%v)",
		}, " ")+"}",
		v.XCoordinate,
		v.YCoordinate,
		v.ZCoordinate,
		v.Power,
		v.PathLossExponent,
	)
}

type SetDeviceConfiguration struct {
	Power Power
	// PathLossExponent is the rate at which the signal power decays with increasing distance
	PathLossExponent  PathLossExponent
	CalculationPeriod CalculationPeriod
	// NumberRssiMeasurements is the number of measurements to use to generate one location estimate
	NumberRssiMeasurements NumberRssiMeasurements
	ReportingPeriod        ReportingPeriod
}

type SetDeviceConfigurationHandler interface {
	HandleSetDeviceConfiguration(frame Frame, cmd *SetDeviceConfiguration) error
}

// SetDeviceConfigurationCommand is the Command ID of SetDeviceConfiguration
const SetDeviceConfigurationCommand CommandID = 0x0001

// Values returns all values of SetDeviceConfiguration
func (v *SetDeviceConfiguration) Values() []zcl.Val {
	return []zcl.Val{
		&v.Power,
		&v.PathLossExponent,
		&v.CalculationPeriod,
		&v.NumberRssiMeasurements,
		&v.ReportingPeriod,
	}
}

// Arguments returns all values of SetDeviceConfiguration
func (v *SetDeviceConfiguration) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Power", Argument: &v.Power},
		{Name: "PathLossExponent", Argument: &v.PathLossExponent},
		{Name: "CalculationPeriod", Argument: &v.CalculationPeriod},
		{Name: "NumberRssiMeasurements", Argument: &v.NumberRssiMeasurements},
		{Name: "ReportingPeriod", Argument: &v.ReportingPeriod},
	}
}

// Name of the command
func (SetDeviceConfiguration) Name() string { return `Set Device Configuration` }

// Description of the command
func (SetDeviceConfiguration) Description() string { return `` }

// ID of the command
func (SetDeviceConfiguration) ID() CommandID { return SetDeviceConfigurationCommand }

// Required
func (SetDeviceConfiguration) Required() bool { return true }

// Cluster ID of the command
func (SetDeviceConfiguration) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (SetDeviceConfiguration) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (SetDeviceConfiguration) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SetDeviceConfiguration) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *SetDeviceConfiguration) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h SetDeviceConfigurationHandler
	if h, found = handler.(SetDeviceConfigurationHandler); found {
		err = h.HandleSetDeviceConfiguration(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of SetDeviceConfiguration
func (v SetDeviceConfiguration) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Power.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.PathLossExponent.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.CalculationPeriod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.NumberRssiMeasurements.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ReportingPeriod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the SetDeviceConfiguration struct
func (v *SetDeviceConfiguration) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Power).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PathLossExponent).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.CalculationPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberRssiMeasurements).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ReportingPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v SetDeviceConfiguration) String() string {
	return zcl.Sprintf(
		"SetDeviceConfiguration{"+zcl.StrJoin([]string{
			"Power(%v)",
			"PathLossExponent(%v)",
			"CalculationPeriod(%v)",
			"NumberRssiMeasurements(%v)",
			"ReportingPeriod(%v)",
		}, " ")+"}",
		v.Power,
		v.PathLossExponent,
		v.CalculationPeriod,
		v.NumberRssiMeasurements,
		v.ReportingPeriod,
	)
}

type GetDeviceConfiguration struct {
	TargetAddress Device
}

type GetDeviceConfigurationHandler interface {
	HandleGetDeviceConfiguration(frame Frame, cmd *GetDeviceConfiguration) (*DeviceConfigurationResponse, error)
}

// GetDeviceConfigurationCommand is the Command ID of GetDeviceConfiguration
const GetDeviceConfigurationCommand CommandID = 0x0002

// Values returns all values of GetDeviceConfiguration
func (v *GetDeviceConfiguration) Values() []zcl.Val {
	return []zcl.Val{
		&v.TargetAddress,
	}
}

// Arguments returns all values of GetDeviceConfiguration
func (v *GetDeviceConfiguration) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "TargetAddress", Argument: &v.TargetAddress},
	}
}

// Name of the command
func (GetDeviceConfiguration) Name() string { return `Get Device Configuration` }

// Description of the command
func (GetDeviceConfiguration) Description() string { return `` }

// ID of the command
func (GetDeviceConfiguration) ID() CommandID { return GetDeviceConfigurationCommand }

// Required
func (GetDeviceConfiguration) Required() bool { return true }

// Cluster ID of the command
func (GetDeviceConfiguration) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (GetDeviceConfiguration) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (GetDeviceConfiguration) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetDeviceConfiguration) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

func (v *GetDeviceConfiguration) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h GetDeviceConfigurationHandler
	if h, found = handler.(GetDeviceConfigurationHandler); found {
		rsp, err = h.HandleGetDeviceConfiguration(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of GetDeviceConfiguration
func (v GetDeviceConfiguration) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.TargetAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetDeviceConfiguration struct
func (v *GetDeviceConfiguration) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.TargetAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetDeviceConfiguration) String() string {
	return zcl.Sprintf(
		"GetDeviceConfiguration{"+zcl.StrJoin([]string{
			"TargetAddress(%v)",
		}, " ")+"}",
		v.TargetAddress,
	)
}

type GetLocationData struct {
	LocationFlags   LocationFlags
	NumberResponses NumberResponses
	TargetAddress   Device
}

type GetLocationDataHandler interface {
	HandleGetLocationData(frame Frame, cmd *GetLocationData) (*LocationDataResponse, error)
}

// GetLocationDataCommand is the Command ID of GetLocationData
const GetLocationDataCommand CommandID = 0x0003

// Values returns all values of GetLocationData
func (v *GetLocationData) Values() []zcl.Val {
	return []zcl.Val{
		&v.LocationFlags,
		&v.NumberResponses,
		&v.TargetAddress,
	}
}

// Arguments returns all values of GetLocationData
func (v *GetLocationData) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "LocationFlags", Argument: &v.LocationFlags},
		{Name: "NumberResponses", Argument: &v.NumberResponses},
		{Name: "TargetAddress", Argument: &v.TargetAddress},
	}
}

// Name of the command
func (GetLocationData) Name() string { return `Get Location Data` }

// Description of the command
func (GetLocationData) Description() string { return `` }

// ID of the command
func (GetLocationData) ID() CommandID { return GetLocationDataCommand }

// Required
func (GetLocationData) Required() bool { return true }

// Cluster ID of the command
func (GetLocationData) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (GetLocationData) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (GetLocationData) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GetLocationData) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

func (v *GetLocationData) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h GetLocationDataHandler
	if h, found = handler.(GetLocationDataHandler); found {
		rsp, err = h.HandleGetLocationData(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of GetLocationData
func (v GetLocationData) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.LocationFlags.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.NumberResponses.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TargetAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetLocationData struct
func (v *GetLocationData) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.LocationFlags).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberResponses).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TargetAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetLocationData) String() string {
	return zcl.Sprintf(
		"GetLocationData{"+zcl.StrJoin([]string{
			"LocationFlags(%v)",
			"NumberResponses(%v)",
			"TargetAddress(%v)",
		}, " ")+"}",
		v.LocationFlags,
		v.NumberResponses,
		v.TargetAddress,
	)
}

type RssiResponse struct {
	Device      Device
	XCoordinate XCoordinate
	YCoordinate YCoordinate
	ZCoordinate ZCoordinate
	Rssi        Rssi
	// NumberRssiMeasurements is the number of measurements to use to generate one location estimate
	NumberRssiMeasurements NumberRssiMeasurements
}

type RssiResponseHandler interface {
	HandleRssiResponse(frame Frame, cmd *RssiResponse) error
}

// RssiResponseCommand is the Command ID of RssiResponse
const RssiResponseCommand CommandID = 0x0004

// Values returns all values of RssiResponse
func (v *RssiResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Device,
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
		&v.Rssi,
		&v.NumberRssiMeasurements,
	}
}

// Arguments returns all values of RssiResponse
func (v *RssiResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Device", Argument: &v.Device},
		{Name: "XCoordinate", Argument: &v.XCoordinate},
		{Name: "YCoordinate", Argument: &v.YCoordinate},
		{Name: "ZCoordinate", Argument: &v.ZCoordinate},
		{Name: "Rssi", Argument: &v.Rssi},
		{Name: "NumberRssiMeasurements", Argument: &v.NumberRssiMeasurements},
	}
}

// Name of the command
func (RssiResponse) Name() string { return `RSSI Response` }

// Description of the command
func (RssiResponse) Description() string { return `` }

// ID of the command
func (RssiResponse) ID() CommandID { return RssiResponseCommand }

// Required
func (RssiResponse) Required() bool { return true }

// Cluster ID of the command
func (RssiResponse) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (RssiResponse) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (RssiResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (RssiResponse) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

func (v *RssiResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h RssiResponseHandler
	if h, found = handler.(RssiResponseHandler); found {
		err = h.HandleRssiResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of RssiResponse
func (v RssiResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Device.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Rssi.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.NumberRssiMeasurements.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RssiResponse struct
func (v *RssiResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Device).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rssi).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberRssiMeasurements).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v RssiResponse) String() string {
	return zcl.Sprintf(
		"RssiResponse{"+zcl.StrJoin([]string{
			"Device(%v)",
			"XCoordinate(%v)",
			"YCoordinate(%v)",
			"ZCoordinate(%v)",
			"Rssi(%v)",
			"NumberRssiMeasurements(%v)",
		}, " ")+"}",
		v.Device,
		v.XCoordinate,
		v.YCoordinate,
		v.ZCoordinate,
		v.Rssi,
		v.NumberRssiMeasurements,
	)
}

type SendPings struct {
	TargetAddress Device
	// NumberRssiMeasurements is the number of measurements to use to generate one location estimate
	NumberRssiMeasurements NumberRssiMeasurements
	CalculationPeriod      CalculationPeriod
}

type SendPingsHandler interface {
	HandleSendPings(frame Frame, cmd *SendPings) (*ReportRssiMeasurements, error)
}

// SendPingsCommand is the Command ID of SendPings
const SendPingsCommand CommandID = 0x0005

// Values returns all values of SendPings
func (v *SendPings) Values() []zcl.Val {
	return []zcl.Val{
		&v.TargetAddress,
		&v.NumberRssiMeasurements,
		&v.CalculationPeriod,
	}
}

// Arguments returns all values of SendPings
func (v *SendPings) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "TargetAddress", Argument: &v.TargetAddress},
		{Name: "NumberRssiMeasurements", Argument: &v.NumberRssiMeasurements},
		{Name: "CalculationPeriod", Argument: &v.CalculationPeriod},
	}
}

// Name of the command
func (SendPings) Name() string { return `Send Pings` }

// Description of the command
func (SendPings) Description() string { return `` }

// ID of the command
func (SendPings) ID() CommandID { return SendPingsCommand }

// Required
func (SendPings) Required() bool { return false }

// Cluster ID of the command
func (SendPings) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (SendPings) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (SendPings) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (SendPings) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

func (v *SendPings) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h SendPingsHandler
	if h, found = handler.(SendPingsHandler); found {
		rsp, err = h.HandleSendPings(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of SendPings
func (v SendPings) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.TargetAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.NumberRssiMeasurements.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.CalculationPeriod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the SendPings struct
func (v *SendPings) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.TargetAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NumberRssiMeasurements).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.CalculationPeriod).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v SendPings) String() string {
	return zcl.Sprintf(
		"SendPings{"+zcl.StrJoin([]string{
			"TargetAddress(%v)",
			"NumberRssiMeasurements(%v)",
			"CalculationPeriod(%v)",
		}, " ")+"}",
		v.TargetAddress,
		v.NumberRssiMeasurements,
		v.CalculationPeriod,
	)
}

type AnchorNodeAnnounce struct {
	Device      Device
	XCoordinate XCoordinate
	YCoordinate YCoordinate
	ZCoordinate ZCoordinate
}

type AnchorNodeAnnounceHandler interface {
	HandleAnchorNodeAnnounce(frame Frame, cmd *AnchorNodeAnnounce) error
}

// AnchorNodeAnnounceCommand is the Command ID of AnchorNodeAnnounce
const AnchorNodeAnnounceCommand CommandID = 0x0006

// Values returns all values of AnchorNodeAnnounce
func (v *AnchorNodeAnnounce) Values() []zcl.Val {
	return []zcl.Val{
		&v.Device,
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
	}
}

// Arguments returns all values of AnchorNodeAnnounce
func (v *AnchorNodeAnnounce) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Device", Argument: &v.Device},
		{Name: "XCoordinate", Argument: &v.XCoordinate},
		{Name: "YCoordinate", Argument: &v.YCoordinate},
		{Name: "ZCoordinate", Argument: &v.ZCoordinate},
	}
}

// Name of the command
func (AnchorNodeAnnounce) Name() string { return `Anchor Node Announce` }

// Description of the command
func (AnchorNodeAnnounce) Description() string { return `` }

// ID of the command
func (AnchorNodeAnnounce) ID() CommandID { return AnchorNodeAnnounceCommand }

// Required
func (AnchorNodeAnnounce) Required() bool { return false }

// Cluster ID of the command
func (AnchorNodeAnnounce) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (AnchorNodeAnnounce) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (AnchorNodeAnnounce) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (AnchorNodeAnnounce) MarshalJSON() ([]byte, error) { return []byte("6"), nil }

func (v *AnchorNodeAnnounce) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h AnchorNodeAnnounceHandler
	if h, found = handler.(AnchorNodeAnnounceHandler); found {
		err = h.HandleAnchorNodeAnnounce(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of AnchorNodeAnnounce
func (v AnchorNodeAnnounce) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Device.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the AnchorNodeAnnounce struct
func (v *AnchorNodeAnnounce) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Device).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v AnchorNodeAnnounce) String() string {
	return zcl.Sprintf(
		"AnchorNodeAnnounce{"+zcl.StrJoin([]string{
			"Device(%v)",
			"XCoordinate(%v)",
			"YCoordinate(%v)",
			"ZCoordinate(%v)",
		}, " ")+"}",
		v.Device,
		v.XCoordinate,
		v.YCoordinate,
		v.ZCoordinate,
	)
}

type DistanceMeasure struct {
	TargetAddress Device
	Resolution    Resolution
}

type DistanceMeasureHandler interface {
	HandleDistanceMeasure(frame Frame, cmd *DistanceMeasure) (*DistanceMeasureResponse, error)
}

// DistanceMeasureCommand is the Command ID of DistanceMeasure
const DistanceMeasureCommand CommandID = 0x0040

// Values returns all values of DistanceMeasure
func (v *DistanceMeasure) Values() []zcl.Val {
	return []zcl.Val{
		&v.TargetAddress,
		&v.Resolution,
	}
}

// Arguments returns all values of DistanceMeasure
func (v *DistanceMeasure) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "TargetAddress", Argument: &v.TargetAddress},
		{Name: "Resolution", Argument: &v.Resolution},
	}
}

// Name of the command
func (DistanceMeasure) Name() string { return `Distance measure` }

// Description of the command
func (DistanceMeasure) Description() string { return `` }

// ID of the command
func (DistanceMeasure) ID() CommandID { return DistanceMeasureCommand }

// Required
func (DistanceMeasure) Required() bool { return true }

// Cluster ID of the command
func (DistanceMeasure) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (DistanceMeasure) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (DistanceMeasure) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (DistanceMeasure) MarshalJSON() ([]byte, error) { return []byte("64"), nil }

func (v *DistanceMeasure) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h DistanceMeasureHandler
	if h, found = handler.(DistanceMeasureHandler); found {
		rsp, err = h.HandleDistanceMeasure(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of DistanceMeasure
func (v DistanceMeasure) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.TargetAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Resolution.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the DistanceMeasure struct
func (v *DistanceMeasure) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.TargetAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Resolution).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v DistanceMeasure) String() string {
	return zcl.Sprintf(
		"DistanceMeasure{"+zcl.StrJoin([]string{
			"TargetAddress(%v)",
			"Resolution(%v)",
		}, " ")+"}",
		v.TargetAddress,
		v.Resolution,
	)
}

type DeviceConfigurationResponse struct {
	Status Status
	Power  Power
	// PathLossExponent is the rate at which the signal power decays with increasing distance
	PathLossExponent  PathLossExponent
	CalculationPeriod CalculationPeriod
	// NumberRssiMeasurements is the number of measurements to use to generate one location estimate
	NumberRssiMeasurements NumberRssiMeasurements
	ReportingPeriod        ReportingPeriod
}

type DeviceConfigurationResponseHandler interface {
	HandleDeviceConfigurationResponse(frame Frame, cmd *DeviceConfigurationResponse) error
}

// DeviceConfigurationResponseCommand is the Command ID of DeviceConfigurationResponse
const DeviceConfigurationResponseCommand CommandID = 0x0000

// Values returns all values of DeviceConfigurationResponse
func (v *DeviceConfigurationResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.Power,
		&v.PathLossExponent,
		&v.CalculationPeriod,
		&v.NumberRssiMeasurements,
		&v.ReportingPeriod,
	}
}

// Arguments returns all values of DeviceConfigurationResponse
func (v *DeviceConfigurationResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "Power", Argument: &v.Power},
		{Name: "PathLossExponent", Argument: &v.PathLossExponent},
		{Name: "CalculationPeriod", Argument: &v.CalculationPeriod},
		{Name: "NumberRssiMeasurements", Argument: &v.NumberRssiMeasurements},
		{Name: "ReportingPeriod", Argument: &v.ReportingPeriod},
	}
}

// Name of the command
func (DeviceConfigurationResponse) Name() string { return `Device Configuration Response` }

// Description of the command
func (DeviceConfigurationResponse) Description() string { return `` }

// ID of the command
func (DeviceConfigurationResponse) ID() CommandID { return DeviceConfigurationResponseCommand }

// Required
func (DeviceConfigurationResponse) Required() bool { return true }

// Cluster ID of the command
func (DeviceConfigurationResponse) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (DeviceConfigurationResponse) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (DeviceConfigurationResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (DeviceConfigurationResponse) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *DeviceConfigurationResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h DeviceConfigurationResponseHandler
	if h, found = handler.(DeviceConfigurationResponseHandler); found {
		err = h.HandleDeviceConfigurationResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of DeviceConfigurationResponse
func (v DeviceConfigurationResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// Power is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.Power.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// PathLossExponent is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.PathLossExponent.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// CalculationPeriod is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.CalculationPeriod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// NumberRssiMeasurements is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.NumberRssiMeasurements.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ReportingPeriod is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ReportingPeriod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the DeviceConfigurationResponse struct
func (v *DeviceConfigurationResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// Power is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.Power).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// PathLossExponent is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.PathLossExponent).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// CalculationPeriod is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.CalculationPeriod).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// NumberRssiMeasurements is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.NumberRssiMeasurements).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ReportingPeriod is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ReportingPeriod).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v DeviceConfigurationResponse) String() string {
	return zcl.Sprintf(
		"DeviceConfigurationResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"Power(%v)",
			"PathLossExponent(%v)",
			"CalculationPeriod(%v)",
			"NumberRssiMeasurements(%v)",
			"ReportingPeriod(%v)",
		}, " ")+"}",
		v.Status,
		v.Power,
		v.PathLossExponent,
		v.CalculationPeriod,
		v.NumberRssiMeasurements,
		v.ReportingPeriod,
	)
}

type LocationDataResponse struct {
	Status       Status
	LocationType LocationType
	XCoordinate  XCoordinate
	YCoordinate  YCoordinate
	ZCoordinate  ZCoordinate
	Power        Power
	// PathLossExponent is the rate at which the signal power decays with increasing distance
	PathLossExponent PathLossExponent
	LocationMethod   LocationMethod
	QualityMeasure   QualityMeasure
	LocationAge      LocationAge
}

type LocationDataResponseHandler interface {
	HandleLocationDataResponse(frame Frame, cmd *LocationDataResponse) error
}

// LocationDataResponseCommand is the Command ID of LocationDataResponse
const LocationDataResponseCommand CommandID = 0x0001

// Values returns all values of LocationDataResponse
func (v *LocationDataResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.LocationType,
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
		&v.Power,
		&v.PathLossExponent,
		&v.LocationMethod,
		&v.QualityMeasure,
		&v.LocationAge,
	}
}

// Arguments returns all values of LocationDataResponse
func (v *LocationDataResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "LocationType", Argument: &v.LocationType},
		{Name: "XCoordinate", Argument: &v.XCoordinate},
		{Name: "YCoordinate", Argument: &v.YCoordinate},
		{Name: "ZCoordinate", Argument: &v.ZCoordinate},
		{Name: "Power", Argument: &v.Power},
		{Name: "PathLossExponent", Argument: &v.PathLossExponent},
		{Name: "LocationMethod", Argument: &v.LocationMethod},
		{Name: "QualityMeasure", Argument: &v.QualityMeasure},
		{Name: "LocationAge", Argument: &v.LocationAge},
	}
}

// Name of the command
func (LocationDataResponse) Name() string { return `Location Data Response` }

// Description of the command
func (LocationDataResponse) Description() string { return `` }

// ID of the command
func (LocationDataResponse) ID() CommandID { return LocationDataResponseCommand }

// Required
func (LocationDataResponse) Required() bool { return true }

// Cluster ID of the command
func (LocationDataResponse) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (LocationDataResponse) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (LocationDataResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (LocationDataResponse) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *LocationDataResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h LocationDataResponseHandler
	if h, found = handler.(LocationDataResponseHandler); found {
		err = h.HandleLocationDataResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of LocationDataResponse
func (v LocationDataResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// LocationType is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.LocationType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// XCoordinate is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// YCoordinate is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ZCoordinate is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// Power is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.Power.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// PathLossExponent is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.PathLossExponent.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// LocationMethod is only included if successful and location type is calculated (i.e. absolute bit not set)
	if v.Status == 0x00 && (v.LocationType&0x0001) == 0x00 {
		if tmp, err = v.LocationMethod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// QualityMeasure is only included if successful and location type is calculated (i.e. absolute bit not set)
	if v.Status == 0x00 && (v.LocationType&0x0001) == 0x00 {
		if tmp, err = v.QualityMeasure.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// LocationAge is only included if successful and location type is calculated (i.e. absolute bit not set)
	if v.Status == 0x00 && (v.LocationType&0x0001) == 0x00 {
		if tmp, err = v.LocationAge.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the LocationDataResponse struct
func (v *LocationDataResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// LocationType is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.LocationType).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// XCoordinate is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// YCoordinate is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ZCoordinate is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// Power is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.Power).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// PathLossExponent is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.PathLossExponent).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// LocationMethod is only included if successful and location type is calculated (i.e. absolute bit not set)
	if v.Status == 0x00 && (v.LocationType&0x0001) == 0x00 {
		if b, err = (&v.LocationMethod).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// QualityMeasure is only included if successful and location type is calculated (i.e. absolute bit not set)
	if v.Status == 0x00 && (v.LocationType&0x0001) == 0x00 {
		if b, err = (&v.QualityMeasure).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// LocationAge is only included if successful and location type is calculated (i.e. absolute bit not set)
	if v.Status == 0x00 && (v.LocationType&0x0001) == 0x00 {
		if b, err = (&v.LocationAge).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v LocationDataResponse) String() string {
	return zcl.Sprintf(
		"LocationDataResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"LocationType(%v)",
			"XCoordinate(%v)",
			"YCoordinate(%v)",
			"ZCoordinate(%v)",
			"Power(%v)",
			"PathLossExponent(%v)",
			"LocationMethod(%v)",
			"QualityMeasure(%v)",
			"LocationAge(%v)",
		}, " ")+"}",
		v.Status,
		v.LocationType,
		v.XCoordinate,
		v.YCoordinate,
		v.ZCoordinate,
		v.Power,
		v.PathLossExponent,
		v.LocationMethod,
		v.QualityMeasure,
		v.LocationAge,
	)
}

type LocationDataNotification struct {
	LocationType LocationType
	XCoordinate  XCoordinate
	YCoordinate  YCoordinate
	ZCoordinate  ZCoordinate
	Power        Power
	// PathLossExponent is the rate at which the signal power decays with increasing distance
	PathLossExponent PathLossExponent
	LocationMethod   LocationMethod
	QualityMeasure   QualityMeasure
	LocationAge      LocationAge
}

type LocationDataNotificationHandler interface {
	HandleLocationDataNotification(frame Frame, cmd *LocationDataNotification) error
}

// LocationDataNotificationCommand is the Command ID of LocationDataNotification
const LocationDataNotificationCommand CommandID = 0x0002

// Values returns all values of LocationDataNotification
func (v *LocationDataNotification) Values() []zcl.Val {
	return []zcl.Val{
		&v.LocationType,
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
		&v.Power,
		&v.PathLossExponent,
		&v.LocationMethod,
		&v.QualityMeasure,
		&v.LocationAge,
	}
}

// Arguments returns all values of LocationDataNotification
func (v *LocationDataNotification) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "LocationType", Argument: &v.LocationType},
		{Name: "XCoordinate", Argument: &v.XCoordinate},
		{Name: "YCoordinate", Argument: &v.YCoordinate},
		{Name: "ZCoordinate", Argument: &v.ZCoordinate},
		{Name: "Power", Argument: &v.Power},
		{Name: "PathLossExponent", Argument: &v.PathLossExponent},
		{Name: "LocationMethod", Argument: &v.LocationMethod},
		{Name: "QualityMeasure", Argument: &v.QualityMeasure},
		{Name: "LocationAge", Argument: &v.LocationAge},
	}
}

// Name of the command
func (LocationDataNotification) Name() string { return `Location Data Notification` }

// Description of the command
func (LocationDataNotification) Description() string { return `` }

// ID of the command
func (LocationDataNotification) ID() CommandID { return LocationDataNotificationCommand }

// Required
func (LocationDataNotification) Required() bool { return true }

// Cluster ID of the command
func (LocationDataNotification) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (LocationDataNotification) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (LocationDataNotification) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (LocationDataNotification) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

func (v *LocationDataNotification) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h LocationDataNotificationHandler
	if h, found = handler.(LocationDataNotificationHandler); found {
		err = h.HandleLocationDataNotification(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of LocationDataNotification
func (v LocationDataNotification) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.LocationType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ZCoordinate is only included if location type is 3d (either absolute or calculated)
	if (v.LocationType & 0x02) == 0x00 {
		if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Power.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.PathLossExponent.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// LocationMethod is only included if location type is calculated (i.e. absolute bit not set)
	if (v.LocationType & 0x0001) == 0x00 {
		if tmp, err = v.LocationMethod.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// QualityMeasure is only included if location type is calculated (i.e. absolute bit not set)
	if (v.LocationType & 0x0001) == 0x00 {
		if tmp, err = v.QualityMeasure.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// LocationAge is only included if location type is calculated (i.e. absolute bit not set)
	if (v.LocationType & 0x0001) == 0x00 {
		if tmp, err = v.LocationAge.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the LocationDataNotification struct
func (v *LocationDataNotification) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.LocationType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// ZCoordinate is only included if location type is 3d (either absolute or calculated)
	if (v.LocationType & 0x02) == 0x00 {
		if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if b, err = (&v.Power).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.PathLossExponent).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// LocationMethod is only included if location type is calculated (i.e. absolute bit not set)
	if (v.LocationType & 0x0001) == 0x00 {
		if b, err = (&v.LocationMethod).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// QualityMeasure is only included if location type is calculated (i.e. absolute bit not set)
	if (v.LocationType & 0x0001) == 0x00 {
		if b, err = (&v.QualityMeasure).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// LocationAge is only included if location type is calculated (i.e. absolute bit not set)
	if (v.LocationType & 0x0001) == 0x00 {
		if b, err = (&v.LocationAge).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v LocationDataNotification) String() string {
	return zcl.Sprintf(
		"LocationDataNotification{"+zcl.StrJoin([]string{
			"LocationType(%v)",
			"XCoordinate(%v)",
			"YCoordinate(%v)",
			"ZCoordinate(%v)",
			"Power(%v)",
			"PathLossExponent(%v)",
			"LocationMethod(%v)",
			"QualityMeasure(%v)",
			"LocationAge(%v)",
		}, " ")+"}",
		v.LocationType,
		v.XCoordinate,
		v.YCoordinate,
		v.ZCoordinate,
		v.Power,
		v.PathLossExponent,
		v.LocationMethod,
		v.QualityMeasure,
		v.LocationAge,
	)
}

type CompactLocationDataNotification struct {
	LocationType   LocationType
	XCoordinate    XCoordinate
	YCoordinate    YCoordinate
	ZCoordinate    ZCoordinate
	QualityMeasure QualityMeasure
	LocationAge    LocationAge
}

type CompactLocationDataNotificationHandler interface {
	HandleCompactLocationDataNotification(frame Frame, cmd *CompactLocationDataNotification) error
}

// CompactLocationDataNotificationCommand is the Command ID of CompactLocationDataNotification
const CompactLocationDataNotificationCommand CommandID = 0x0003

// Values returns all values of CompactLocationDataNotification
func (v *CompactLocationDataNotification) Values() []zcl.Val {
	return []zcl.Val{
		&v.LocationType,
		&v.XCoordinate,
		&v.YCoordinate,
		&v.ZCoordinate,
		&v.QualityMeasure,
		&v.LocationAge,
	}
}

// Arguments returns all values of CompactLocationDataNotification
func (v *CompactLocationDataNotification) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "LocationType", Argument: &v.LocationType},
		{Name: "XCoordinate", Argument: &v.XCoordinate},
		{Name: "YCoordinate", Argument: &v.YCoordinate},
		{Name: "ZCoordinate", Argument: &v.ZCoordinate},
		{Name: "QualityMeasure", Argument: &v.QualityMeasure},
		{Name: "LocationAge", Argument: &v.LocationAge},
	}
}

// Name of the command
func (CompactLocationDataNotification) Name() string { return `Compact Location Data Notification` }

// Description of the command
func (CompactLocationDataNotification) Description() string { return `` }

// ID of the command
func (CompactLocationDataNotification) ID() CommandID { return CompactLocationDataNotificationCommand }

// Required
func (CompactLocationDataNotification) Required() bool { return true }

// Cluster ID of the command
func (CompactLocationDataNotification) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (CompactLocationDataNotification) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (CompactLocationDataNotification) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (CompactLocationDataNotification) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

func (v *CompactLocationDataNotification) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h CompactLocationDataNotificationHandler
	if h, found = handler.(CompactLocationDataNotificationHandler); found {
		err = h.HandleCompactLocationDataNotification(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of CompactLocationDataNotification
func (v CompactLocationDataNotification) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.LocationType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.XCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.YCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ZCoordinate is only included if location type is 3d (either absolute or calculated)
	if (v.LocationType & 0x02) == 0x00 {
		if tmp, err = v.ZCoordinate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// QualityMeasure is only included if location type is calculated (i.e. absolute bit not set)
	if (v.LocationType & 0x0001) == 0x00 {
		if tmp, err = v.QualityMeasure.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// LocationAge is only included if location type is calculated (i.e. absolute bit not set)
	if (v.LocationType & 0x0001) == 0x00 {
		if tmp, err = v.LocationAge.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the CompactLocationDataNotification struct
func (v *CompactLocationDataNotification) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.LocationType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.XCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.YCoordinate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// ZCoordinate is only included if location type is 3d (either absolute or calculated)
	if (v.LocationType & 0x02) == 0x00 {
		if b, err = (&v.ZCoordinate).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// QualityMeasure is only included if location type is calculated (i.e. absolute bit not set)
	if (v.LocationType & 0x0001) == 0x00 {
		if b, err = (&v.QualityMeasure).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// LocationAge is only included if location type is calculated (i.e. absolute bit not set)
	if (v.LocationType & 0x0001) == 0x00 {
		if b, err = (&v.LocationAge).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v CompactLocationDataNotification) String() string {
	return zcl.Sprintf(
		"CompactLocationDataNotification{"+zcl.StrJoin([]string{
			"LocationType(%v)",
			"XCoordinate(%v)",
			"YCoordinate(%v)",
			"ZCoordinate(%v)",
			"QualityMeasure(%v)",
			"LocationAge(%v)",
		}, " ")+"}",
		v.LocationType,
		v.XCoordinate,
		v.YCoordinate,
		v.ZCoordinate,
		v.QualityMeasure,
		v.LocationAge,
	)
}

type RssiPing struct {
	LocationType LocationType
}

type RssiPingHandler interface {
	HandleRssiPing(frame Frame, cmd *RssiPing) error
}

// RssiPingCommand is the Command ID of RssiPing
const RssiPingCommand CommandID = 0x0004

// Values returns all values of RssiPing
func (v *RssiPing) Values() []zcl.Val {
	return []zcl.Val{
		&v.LocationType,
	}
}

// Arguments returns all values of RssiPing
func (v *RssiPing) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "LocationType", Argument: &v.LocationType},
	}
}

// Name of the command
func (RssiPing) Name() string { return `RSSI Ping` }

// Description of the command
func (RssiPing) Description() string { return `` }

// ID of the command
func (RssiPing) ID() CommandID { return RssiPingCommand }

// Required
func (RssiPing) Required() bool { return true }

// Cluster ID of the command
func (RssiPing) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (RssiPing) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (RssiPing) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (RssiPing) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

func (v *RssiPing) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h RssiPingHandler
	if h, found = handler.(RssiPingHandler); found {
		err = h.HandleRssiPing(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of RssiPing
func (v RssiPing) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.LocationType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RssiPing struct
func (v *RssiPing) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.LocationType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v RssiPing) String() string {
	return zcl.Sprintf(
		"RssiPing{"+zcl.StrJoin([]string{
			"LocationType(%v)",
		}, " ")+"}",
		v.LocationType,
	)
}

type RssiRequest struct {
}

type RssiRequestHandler interface {
	HandleRssiRequest(frame Frame, cmd *RssiRequest) (*RssiResponse, error)
}

// RssiRequestCommand is the Command ID of RssiRequest
const RssiRequestCommand CommandID = 0x0005

// Values returns all values of RssiRequest
func (v *RssiRequest) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of RssiRequest
func (v *RssiRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (RssiRequest) Name() string { return `RSSI Request` }

// Description of the command
func (RssiRequest) Description() string { return `` }

// ID of the command
func (RssiRequest) ID() CommandID { return RssiRequestCommand }

// Required
func (RssiRequest) Required() bool { return false }

// Cluster ID of the command
func (RssiRequest) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (RssiRequest) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (RssiRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (RssiRequest) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

func (v *RssiRequest) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h RssiRequestHandler
	if h, found = handler.(RssiRequestHandler); found {
		rsp, err = h.HandleRssiRequest(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of RssiRequest
func (v RssiRequest) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the RssiRequest struct
func (v *RssiRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v RssiRequest) String() string {
	return zcl.Sprintf(
		"RssiRequest{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type ReportRssiMeasurements struct {
	Device            Device
	NeighborsInfoList NeighborsInfoList
}

type ReportRssiMeasurementsHandler interface {
	HandleReportRssiMeasurements(frame Frame, cmd *ReportRssiMeasurements) error
}

// ReportRssiMeasurementsCommand is the Command ID of ReportRssiMeasurements
const ReportRssiMeasurementsCommand CommandID = 0x0006

// Values returns all values of ReportRssiMeasurements
func (v *ReportRssiMeasurements) Values() []zcl.Val {
	return []zcl.Val{
		&v.Device,
		&v.NeighborsInfoList,
	}
}

// Arguments returns all values of ReportRssiMeasurements
func (v *ReportRssiMeasurements) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Device", Argument: &v.Device},
		{Name: "NeighborsInfoList", Argument: &v.NeighborsInfoList},
	}
}

// Name of the command
func (ReportRssiMeasurements) Name() string { return `Report RSSI Measurements` }

// Description of the command
func (ReportRssiMeasurements) Description() string { return `` }

// ID of the command
func (ReportRssiMeasurements) ID() CommandID { return ReportRssiMeasurementsCommand }

// Required
func (ReportRssiMeasurements) Required() bool { return false }

// Cluster ID of the command
func (ReportRssiMeasurements) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (ReportRssiMeasurements) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (ReportRssiMeasurements) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ReportRssiMeasurements) MarshalJSON() ([]byte, error) { return []byte("6"), nil }

func (v *ReportRssiMeasurements) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ReportRssiMeasurementsHandler
	if h, found = handler.(ReportRssiMeasurementsHandler); found {
		err = h.HandleReportRssiMeasurements(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of ReportRssiMeasurements
func (v ReportRssiMeasurements) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Device.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.NeighborsInfoList.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ReportRssiMeasurements struct
func (v *ReportRssiMeasurements) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Device).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.NeighborsInfoList).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ReportRssiMeasurements) String() string {
	return zcl.Sprintf(
		"ReportRssiMeasurements{"+zcl.StrJoin([]string{
			"Device(%v)",
			"NeighborsInfoList(%v)",
		}, " ")+"}",
		v.Device,
		v.NeighborsInfoList,
	)
}

type RequestOwnLocation struct {
	BlindNodeAddress Device
}

type RequestOwnLocationHandler interface {
	HandleRequestOwnLocation(frame Frame, cmd *RequestOwnLocation) (*SetAbsoluteLocation, error)
}

// RequestOwnLocationCommand is the Command ID of RequestOwnLocation
const RequestOwnLocationCommand CommandID = 0x0007

// Values returns all values of RequestOwnLocation
func (v *RequestOwnLocation) Values() []zcl.Val {
	return []zcl.Val{
		&v.BlindNodeAddress,
	}
}

// Arguments returns all values of RequestOwnLocation
func (v *RequestOwnLocation) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "BlindNodeAddress", Argument: &v.BlindNodeAddress},
	}
}

// Name of the command
func (RequestOwnLocation) Name() string { return `Request Own Location` }

// Description of the command
func (RequestOwnLocation) Description() string { return `` }

// ID of the command
func (RequestOwnLocation) ID() CommandID { return RequestOwnLocationCommand }

// Required
func (RequestOwnLocation) Required() bool { return false }

// Cluster ID of the command
func (RequestOwnLocation) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (RequestOwnLocation) Direction() zcl.Direction { return zcl.ServerToClient }

// MnfCode returns the manufacturer code (if any) of the command
func (RequestOwnLocation) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (RequestOwnLocation) MarshalJSON() ([]byte, error) { return []byte("7"), nil }

func (v *RequestOwnLocation) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h RequestOwnLocationHandler
	if h, found = handler.(RequestOwnLocationHandler); found {
		rsp, err = h.HandleRequestOwnLocation(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of RequestOwnLocation
func (v RequestOwnLocation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.BlindNodeAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the RequestOwnLocation struct
func (v *RequestOwnLocation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.BlindNodeAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v RequestOwnLocation) String() string {
	return zcl.Sprintf(
		"RequestOwnLocation{"+zcl.StrJoin([]string{
			"BlindNodeAddress(%v)",
		}, " ")+"}",
		v.BlindNodeAddress,
	)
}

// DistanceMeasureResponse Returns the result of a distance measure.
type DistanceMeasureResponse struct {
	TargetAddress Device
	Distance      Distance
	QualityIndex  QualityIndex
}

type DistanceMeasureResponseHandler interface {
	HandleDistanceMeasureResponse(frame Frame, cmd *DistanceMeasureResponse) error
}

// DistanceMeasureResponseCommand is the Command ID of DistanceMeasureResponse
const DistanceMeasureResponseCommand CommandID = 0x0040

// Values returns all values of DistanceMeasureResponse
func (v *DistanceMeasureResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.TargetAddress,
		&v.Distance,
		&v.QualityIndex,
	}
}

// Arguments returns all values of DistanceMeasureResponse
func (v *DistanceMeasureResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "TargetAddress", Argument: &v.TargetAddress},
		{Name: "Distance", Argument: &v.Distance},
		{Name: "QualityIndex", Argument: &v.QualityIndex},
	}
}

// Name of the command
func (DistanceMeasureResponse) Name() string { return `Distance measure response` }

// Description of the command
func (DistanceMeasureResponse) Description() string {
	return `Returns the result of a distance measure.`
}

// ID of the command
func (DistanceMeasureResponse) ID() CommandID { return DistanceMeasureResponseCommand }

// Required
func (DistanceMeasureResponse) Required() bool { return true }

// Cluster ID of the command
func (DistanceMeasureResponse) Cluster() zcl.ClusterID { return LocationID }

// Direction of the command
func (DistanceMeasureResponse) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (DistanceMeasureResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (DistanceMeasureResponse) MarshalJSON() ([]byte, error) { return []byte("64"), nil }

func (v *DistanceMeasureResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h DistanceMeasureResponseHandler
	if h, found = handler.(DistanceMeasureResponseHandler); found {
		err = h.HandleDistanceMeasureResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of DistanceMeasureResponse
func (v DistanceMeasureResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.TargetAddress.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Distance.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.QualityIndex.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the DistanceMeasureResponse struct
func (v *DistanceMeasureResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.TargetAddress).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Distance).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.QualityIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v DistanceMeasureResponse) String() string {
	return zcl.Sprintf(
		"DistanceMeasureResponse{"+zcl.StrJoin([]string{
			"TargetAddress(%v)",
			"Distance(%v)",
			"QualityIndex(%v)",
		}, " ")+"}",
		v.TargetAddress,
		v.Distance,
		v.QualityIndex,
	)
}
