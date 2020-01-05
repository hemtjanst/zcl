package general

import "hemtjan.st/zcl"

// Alarms
const AlarmsID zcl.ClusterID = 9

var AlarmsCluster = zcl.Cluster{
	Name: "Alarms",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		ResetAlarmCommand:     func() zcl.Command { return new(ResetAlarm) },
		ResetAllAlarmsCommand: func() zcl.Command { return new(ResetAllAlarms) },
		GetAlarmCommand:       func() zcl.Command { return new(GetAlarm) },
		ResetAlarmLogCommand:  func() zcl.Command { return new(ResetAlarmLog) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		AlarmCommand:            func() zcl.Command { return new(Alarm) },
		GetAlarmResponseCommand: func() zcl.Command { return new(GetAlarmResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		AlarmCountAttr: func() zcl.Attr { return new(AlarmCount) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type ResetAlarm struct {
	AlarmCode AlarmCode
	ClusterId ClusterId
}

// ResetAlarmCommand is the Command ID of ResetAlarm
const ResetAlarmCommand CommandID = 0x0000

// Values returns all values of ResetAlarm
func (v *ResetAlarm) Values() []zcl.Val {
	return []zcl.Val{
		&v.AlarmCode,
		&v.ClusterId,
	}
}

// Arguments returns all values of ResetAlarm
func (v *ResetAlarm) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "AlarmCode", Argument: &v.AlarmCode},
		{Name: "ClusterId", Argument: &v.ClusterId},
	}
}

// Name of the command
func (ResetAlarm) Name() string { return "Reset alarm" }

// ID of the command
func (ResetAlarm) ID() CommandID { return ResetAlarmCommand }

// Required
func (ResetAlarm) Required() bool { return true }

// Cluster ID of the command
func (ResetAlarm) Cluster() zcl.ClusterID { return AlarmsID }

// MnfCode returns the manufacturer code (if any) of the command
func (ResetAlarm) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (ResetAlarm) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

// MarshalZcl returns the wire format representation of ResetAlarm
func (v ResetAlarm) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.AlarmCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ClusterId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ResetAlarm struct
func (v *ResetAlarm) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.AlarmCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ClusterId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ResetAlarm) String() string {
	return zcl.Sprintf(
		"ResetAlarm{"+zcl.StrJoin([]string{
			"AlarmCode(%v)",
			"ClusterId(%v)",
		}, " ")+"}",
		v.AlarmCode,
		v.ClusterId,
	)
}

// ResetAllAlarms Resets all alarms, causing triggered alarms to generate new notification
type ResetAllAlarms struct {
}

// ResetAllAlarmsCommand is the Command ID of ResetAllAlarms
const ResetAllAlarmsCommand CommandID = 0x0001

// Values returns all values of ResetAllAlarms
func (v *ResetAllAlarms) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of ResetAllAlarms
func (v *ResetAllAlarms) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (ResetAllAlarms) Name() string { return "Reset all alarms" }

// ID of the command
func (ResetAllAlarms) ID() CommandID { return ResetAllAlarmsCommand }

// Required
func (ResetAllAlarms) Required() bool { return true }

// Cluster ID of the command
func (ResetAllAlarms) Cluster() zcl.ClusterID { return AlarmsID }

// MnfCode returns the manufacturer code (if any) of the command
func (ResetAllAlarms) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (ResetAllAlarms) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

// MarshalZcl returns the wire format representation of ResetAllAlarms
func (v ResetAllAlarms) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the ResetAllAlarms struct
func (v *ResetAllAlarms) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ResetAllAlarms) String() string {
	return zcl.Sprintf(
		"ResetAllAlarms{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

// GetAlarm Retrieves the earliest alarm and removes it from the table
type GetAlarm struct {
}

// GetAlarmCommand is the Command ID of GetAlarm
const GetAlarmCommand CommandID = 0x0002

// Values returns all values of GetAlarm
func (v *GetAlarm) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of GetAlarm
func (v *GetAlarm) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (GetAlarm) Name() string { return "Get Alarm" }

// ID of the command
func (GetAlarm) ID() CommandID { return GetAlarmCommand }

// Required
func (GetAlarm) Required() bool { return false }

// Cluster ID of the command
func (GetAlarm) Cluster() zcl.ClusterID { return AlarmsID }

// MnfCode returns the manufacturer code (if any) of the command
func (GetAlarm) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (GetAlarm) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

// MarshalZcl returns the wire format representation of GetAlarm
func (v GetAlarm) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the GetAlarm struct
func (v *GetAlarm) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetAlarm) String() string {
	return zcl.Sprintf(
		"GetAlarm{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

// ResetAlarmLog Clears the alarm log
type ResetAlarmLog struct {
}

// ResetAlarmLogCommand is the Command ID of ResetAlarmLog
const ResetAlarmLogCommand CommandID = 0x0003

// Values returns all values of ResetAlarmLog
func (v *ResetAlarmLog) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of ResetAlarmLog
func (v *ResetAlarmLog) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (ResetAlarmLog) Name() string { return "Reset alarm log" }

// ID of the command
func (ResetAlarmLog) ID() CommandID { return ResetAlarmLogCommand }

// Required
func (ResetAlarmLog) Required() bool { return false }

// Cluster ID of the command
func (ResetAlarmLog) Cluster() zcl.ClusterID { return AlarmsID }

// MnfCode returns the manufacturer code (if any) of the command
func (ResetAlarmLog) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (ResetAlarmLog) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

// MarshalZcl returns the wire format representation of ResetAlarmLog
func (v ResetAlarmLog) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the ResetAlarmLog struct
func (v *ResetAlarmLog) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ResetAlarmLog) String() string {
	return zcl.Sprintf(
		"ResetAlarmLog{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type Alarm struct {
	AlarmCode AlarmCode
	ClusterId ClusterId
}

// AlarmCommand is the Command ID of Alarm
const AlarmCommand CommandID = 0x0000

// Values returns all values of Alarm
func (v *Alarm) Values() []zcl.Val {
	return []zcl.Val{
		&v.AlarmCode,
		&v.ClusterId,
	}
}

// Arguments returns all values of Alarm
func (v *Alarm) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "AlarmCode", Argument: &v.AlarmCode},
		{Name: "ClusterId", Argument: &v.ClusterId},
	}
}

// Name of the command
func (Alarm) Name() string { return "Alarm" }

// ID of the command
func (Alarm) ID() CommandID { return AlarmCommand }

// Required
func (Alarm) Required() bool { return true }

// Cluster ID of the command
func (Alarm) Cluster() zcl.ClusterID { return AlarmsID }

// MnfCode returns the manufacturer code (if any) of the command
func (Alarm) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (Alarm) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

// MarshalZcl returns the wire format representation of Alarm
func (v Alarm) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.AlarmCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ClusterId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the Alarm struct
func (v *Alarm) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.AlarmCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ClusterId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v Alarm) String() string {
	return zcl.Sprintf(
		"Alarm{"+zcl.StrJoin([]string{
			"AlarmCode(%v)",
			"ClusterId(%v)",
		}, " ")+"}",
		v.AlarmCode,
		v.ClusterId,
	)
}

type GetAlarmResponse struct {
	Status    Status
	AlarmCode AlarmCode
	ClusterId ClusterId
	Time      Time
}

// GetAlarmResponseCommand is the Command ID of GetAlarmResponse
const GetAlarmResponseCommand CommandID = 0x0001

// Values returns all values of GetAlarmResponse
func (v *GetAlarmResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.AlarmCode,
		&v.ClusterId,
		&v.Time,
	}
}

// Arguments returns all values of GetAlarmResponse
func (v *GetAlarmResponse) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Status", Argument: &v.Status},
		{Name: "AlarmCode", Argument: &v.AlarmCode},
		{Name: "ClusterId", Argument: &v.ClusterId},
		{Name: "Time", Argument: &v.Time},
	}
}

// Name of the command
func (GetAlarmResponse) Name() string { return "Get alarm response" }

// ID of the command
func (GetAlarmResponse) ID() CommandID { return GetAlarmResponseCommand }

// Required
func (GetAlarmResponse) Required() bool { return false }

// Cluster ID of the command
func (GetAlarmResponse) Cluster() zcl.ClusterID { return AlarmsID }

// MnfCode returns the manufacturer code (if any) of the command
func (GetAlarmResponse) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
func (GetAlarmResponse) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

// MarshalZcl returns the wire format representation of GetAlarmResponse
func (v GetAlarmResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	{
		if tmp, err = v.Status.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// AlarmCode is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.AlarmCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// ClusterId is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.ClusterId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	// Time is only included if successful
	if v.Status == 0x00 {
		if tmp, err = v.Time.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GetAlarmResponse struct
func (v *GetAlarmResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	// AlarmCode is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.AlarmCode).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// ClusterId is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.ClusterId).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	// Time is only included if successful
	if v.Status == 0x00 {
		if b, err = (&v.Time).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GetAlarmResponse) String() string {
	return zcl.Sprintf(
		"GetAlarmResponse{"+zcl.StrJoin([]string{
			"Status(%v)",
			"AlarmCode(%v)",
			"ClusterId(%v)",
			"Time(%v)",
		}, " ")+"}",
		v.Status,
		v.AlarmCode,
		v.ClusterId,
		v.Time,
	)
}
