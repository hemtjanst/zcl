package ias

import "hemtjan.st/zcl"

// IasZone
const IasZoneID zcl.ClusterID = 1280

var IasZoneCluster = zcl.Cluster{
	Name: "IAS Zone",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		ZoneEnrollResponseCommand:          func() zcl.Command { return new(ZoneEnrollResponse) },
		InitiateNormalOperationModeCommand: func() zcl.Command { return new(InitiateNormalOperationMode) },
		InitiateTestModeCommand:            func() zcl.Command { return new(InitiateTestMode) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		ZoneStateChangeNotificationCommand: func() zcl.Command { return new(ZoneStateChangeNotification) },
		ZoneEnrollRequestCommand:           func() zcl.Command { return new(ZoneEnrollRequest) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		ZoneStateAttr:        func() zcl.Attr { return new(ZoneState) },
		ZoneTypeAttr:         func() zcl.Attr { return new(ZoneType) },
		ZoneStatusAttr:       func() zcl.Attr { return new(ZoneStatus) },
		IasCieAddressMacAttr: func() zcl.Attr { return new(IasCieAddressMac) },
		ZoneIdAttr:           func() zcl.Attr { return new(ZoneId) },
		NumberOfZoneSensitivityLevelsSupportedAttr: func() zcl.Attr { return new(NumberOfZoneSensitivityLevelsSupported) },
		CurrentZoneSensitivityLevelAttr:            func() zcl.Attr { return new(CurrentZoneSensitivityLevel) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type ZoneEnrollResponse struct {
	EnrollResponseCode EnrollResponseCode
	ZoneId             ZoneId
}

type ZoneEnrollResponseHandler interface {
	HandleZoneEnrollResponse(frame Frame, cmd *ZoneEnrollResponse) error
}

// ZoneEnrollResponseCommand is the Command ID of ZoneEnrollResponse
const ZoneEnrollResponseCommand CommandID = 0x0000

// Values returns all values of ZoneEnrollResponse
func (v *ZoneEnrollResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.EnrollResponseCode,
		&v.ZoneId,
	}
}

// Arguments returns all values of ZoneEnrollResponse
func (v *ZoneEnrollResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "EnrollResponseCode", Argument: &v.EnrollResponseCode},
		{Name: "ZoneId", Argument: &v.ZoneId},
	}
}

// Name of the command
func (ZoneEnrollResponse) Name() string { return `Zone Enroll Response` }

// Description of the command
func (ZoneEnrollResponse) Description() string { return `` }

// ID of the command
func (ZoneEnrollResponse) ID() CommandID { return ZoneEnrollResponseCommand }

// Required
func (ZoneEnrollResponse) Required() bool { return false }

// Cluster ID of the command
func (ZoneEnrollResponse) Cluster() zcl.ClusterID { return IasZoneID }

// Direction of the command
func (ZoneEnrollResponse) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (ZoneEnrollResponse) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ZoneEnrollResponse) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *ZoneEnrollResponse) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ZoneEnrollResponseHandler
	if h, found = handler.(ZoneEnrollResponseHandler); found {
		err = h.HandleZoneEnrollResponse(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of ZoneEnrollResponse
func (v ZoneEnrollResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.EnrollResponseCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ZoneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ZoneEnrollResponse struct
func (v *ZoneEnrollResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.EnrollResponseCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZoneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ZoneEnrollResponse) String() string {
	return zcl.Sprintf(
		"ZoneEnrollResponse{"+zcl.StrJoin([]string{
			"EnrollResponseCode(%v)",
			"ZoneId(%v)",
		}, " ")+"}",
		v.EnrollResponseCode,
		v.ZoneId,
	)
}

type InitiateNormalOperationMode struct {
}

type InitiateNormalOperationModeHandler interface {
	HandleInitiateNormalOperationMode(frame Frame, cmd *InitiateNormalOperationMode) error
}

// InitiateNormalOperationModeCommand is the Command ID of InitiateNormalOperationMode
const InitiateNormalOperationModeCommand CommandID = 0x0001

// Values returns all values of InitiateNormalOperationMode
func (v *InitiateNormalOperationMode) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of InitiateNormalOperationMode
func (v *InitiateNormalOperationMode) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (InitiateNormalOperationMode) Name() string { return `Initiate Normal Operation Mode` }

// Description of the command
func (InitiateNormalOperationMode) Description() string { return `` }

// ID of the command
func (InitiateNormalOperationMode) ID() CommandID { return InitiateNormalOperationModeCommand }

// Required
func (InitiateNormalOperationMode) Required() bool { return false }

// Cluster ID of the command
func (InitiateNormalOperationMode) Cluster() zcl.ClusterID { return IasZoneID }

// Direction of the command
func (InitiateNormalOperationMode) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (InitiateNormalOperationMode) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (InitiateNormalOperationMode) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *InitiateNormalOperationMode) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h InitiateNormalOperationModeHandler
	if h, found = handler.(InitiateNormalOperationModeHandler); found {
		err = h.HandleInitiateNormalOperationMode(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of InitiateNormalOperationMode
func (v InitiateNormalOperationMode) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the InitiateNormalOperationMode struct
func (v *InitiateNormalOperationMode) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v InitiateNormalOperationMode) String() string {
	return zcl.Sprintf(
		"InitiateNormalOperationMode{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type InitiateTestMode struct {
	TestModeDuration            TestModeDuration
	CurrentZoneSensitivityLevel CurrentZoneSensitivityLevel
}

type InitiateTestModeHandler interface {
	HandleInitiateTestMode(frame Frame, cmd *InitiateTestMode) error
}

// InitiateTestModeCommand is the Command ID of InitiateTestMode
const InitiateTestModeCommand CommandID = 0x0002

// Values returns all values of InitiateTestMode
func (v *InitiateTestMode) Values() []zcl.Val {
	return []zcl.Val{
		&v.TestModeDuration,
		&v.CurrentZoneSensitivityLevel,
	}
}

// Arguments returns all values of InitiateTestMode
func (v *InitiateTestMode) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "TestModeDuration", Argument: &v.TestModeDuration},
		{Name: "CurrentZoneSensitivityLevel", Argument: &v.CurrentZoneSensitivityLevel},
	}
}

// Name of the command
func (InitiateTestMode) Name() string { return `Initiate Test Mode` }

// Description of the command
func (InitiateTestMode) Description() string { return `` }

// ID of the command
func (InitiateTestMode) ID() CommandID { return InitiateTestModeCommand }

// Required
func (InitiateTestMode) Required() bool { return false }

// Cluster ID of the command
func (InitiateTestMode) Cluster() zcl.ClusterID { return IasZoneID }

// Direction of the command
func (InitiateTestMode) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (InitiateTestMode) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (InitiateTestMode) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

func (v *InitiateTestMode) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h InitiateTestModeHandler
	if h, found = handler.(InitiateTestModeHandler); found {
		err = h.HandleInitiateTestMode(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of InitiateTestMode
func (v InitiateTestMode) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.TestModeDuration.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.CurrentZoneSensitivityLevel.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the InitiateTestMode struct
func (v *InitiateTestMode) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.TestModeDuration).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.CurrentZoneSensitivityLevel).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v InitiateTestMode) String() string {
	return zcl.Sprintf(
		"InitiateTestMode{"+zcl.StrJoin([]string{
			"TestModeDuration(%v)",
			"CurrentZoneSensitivityLevel(%v)",
		}, " ")+"}",
		v.TestModeDuration,
		v.CurrentZoneSensitivityLevel,
	)
}

type ZoneStateChangeNotification struct {
	ZoneStatus     ZoneStatus
	ExtendedStatus ExtendedStatus
	ZoneId         ZoneId
	Delay          Delay
}

type ZoneStateChangeNotificationHandler interface {
	HandleZoneStateChangeNotification(frame Frame, cmd *ZoneStateChangeNotification) error
}

// ZoneStateChangeNotificationCommand is the Command ID of ZoneStateChangeNotification
const ZoneStateChangeNotificationCommand CommandID = 0x0000

// Values returns all values of ZoneStateChangeNotification
func (v *ZoneStateChangeNotification) Values() []zcl.Val {
	return []zcl.Val{
		&v.ZoneStatus,
		&v.ExtendedStatus,
		&v.ZoneId,
		&v.Delay,
	}
}

// Arguments returns all values of ZoneStateChangeNotification
func (v *ZoneStateChangeNotification) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "ZoneStatus", Argument: &v.ZoneStatus},
		{Name: "ExtendedStatus", Argument: &v.ExtendedStatus},
		{Name: "ZoneId", Argument: &v.ZoneId},
		{Name: "Delay", Argument: &v.Delay},
	}
}

// Name of the command
func (ZoneStateChangeNotification) Name() string { return `Zone State Change Notification` }

// Description of the command
func (ZoneStateChangeNotification) Description() string { return `` }

// ID of the command
func (ZoneStateChangeNotification) ID() CommandID { return ZoneStateChangeNotificationCommand }

// Required
func (ZoneStateChangeNotification) Required() bool { return false }

// Cluster ID of the command
func (ZoneStateChangeNotification) Cluster() zcl.ClusterID { return IasZoneID }

// Direction of the command
func (ZoneStateChangeNotification) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (ZoneStateChangeNotification) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ZoneStateChangeNotification) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *ZoneStateChangeNotification) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ZoneStateChangeNotificationHandler
	if h, found = handler.(ZoneStateChangeNotificationHandler); found {
		err = h.HandleZoneStateChangeNotification(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of ZoneStateChangeNotification
func (v ZoneStateChangeNotification) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.ZoneStatus.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ExtendedStatus.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ZoneId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Delay.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ZoneStateChangeNotification struct
func (v *ZoneStateChangeNotification) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.ZoneStatus).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ExtendedStatus).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZoneId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Delay).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ZoneStateChangeNotification) String() string {
	return zcl.Sprintf(
		"ZoneStateChangeNotification{"+zcl.StrJoin([]string{
			"ZoneStatus(%v)",
			"ExtendedStatus(%v)",
			"ZoneId(%v)",
			"Delay(%v)",
		}, " ")+"}",
		v.ZoneStatus,
		v.ExtendedStatus,
		v.ZoneId,
		v.Delay,
	)
}

type ZoneEnrollRequest struct {
	ZoneType         ZoneType
	ManufacturerCode ManufacturerCode
}

type ZoneEnrollRequestHandler interface {
	HandleZoneEnrollRequest(frame Frame, cmd *ZoneEnrollRequest) (*ZoneEnrollResponse, error)
}

// ZoneEnrollRequestCommand is the Command ID of ZoneEnrollRequest
const ZoneEnrollRequestCommand CommandID = 0x0001

// Values returns all values of ZoneEnrollRequest
func (v *ZoneEnrollRequest) Values() []zcl.Val {
	return []zcl.Val{
		&v.ZoneType,
		&v.ManufacturerCode,
	}
}

// Arguments returns all values of ZoneEnrollRequest
func (v *ZoneEnrollRequest) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "ZoneType", Argument: &v.ZoneType},
		{Name: "ManufacturerCode", Argument: &v.ManufacturerCode},
	}
}

// Name of the command
func (ZoneEnrollRequest) Name() string { return `Zone Enroll Request` }

// Description of the command
func (ZoneEnrollRequest) Description() string { return `` }

// ID of the command
func (ZoneEnrollRequest) ID() CommandID { return ZoneEnrollRequestCommand }

// Required
func (ZoneEnrollRequest) Required() bool { return false }

// Cluster ID of the command
func (ZoneEnrollRequest) Cluster() zcl.ClusterID { return IasZoneID }

// Direction of the command
func (ZoneEnrollRequest) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (ZoneEnrollRequest) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ZoneEnrollRequest) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *ZoneEnrollRequest) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h ZoneEnrollRequestHandler
	if h, found = handler.(ZoneEnrollRequestHandler); found {
		rsp, err = h.HandleZoneEnrollRequest(frame, v)
	} else {
		rsp = &zcl.DefaultResponse{Command: v.ID(), Status: zcl.UnsupClusterCommand}
	}
	return
}

// MarshalZcl returns the wire format representation of ZoneEnrollRequest
func (v ZoneEnrollRequest) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.ZoneType.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ManufacturerCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ZoneEnrollRequest struct
func (v *ZoneEnrollRequest) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.ZoneType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ManufacturerCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ZoneEnrollRequest) String() string {
	return zcl.Sprintf(
		"ZoneEnrollRequest{"+zcl.StrJoin([]string{
			"ZoneType(%v)",
			"ManufacturerCode(%v)",
		}, " ")+"}",
		v.ZoneType,
		v.ManufacturerCode,
	)
}
