// Sending alarm notifications and configuring alarm functionality.
package general

import (
	"hemtjan.st/zcl"
)

// Alarms
const AlarmsID zcl.ClusterID = 9

var AlarmsCluster = zcl.Cluster{
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
	AlarmCode zcl.Zu8
	ClusterId zcl.Zu16
}

const ResetAlarmCommand zcl.CommandID = 0

func (v *ResetAlarm) Values() []zcl.Val {
	return []zcl.Val{
		&v.AlarmCode,
		&v.ClusterId,
	}
}

func (v ResetAlarm) ID() zcl.CommandID {
	return ResetAlarmCommand
}

func (v ResetAlarm) Cluster() zcl.ClusterID {
	return AlarmsID
}

func (v ResetAlarm) MnfCode() []byte {
	return []byte{}
}

func (v ResetAlarm) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.AlarmCode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ClusterId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

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

func (v ResetAlarm) AlarmCodeString() string {
	return zcl.Sprintf("0x%X", zcl.Zu8(v.AlarmCode))
}
func (v ResetAlarm) ClusterIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.ClusterId))
}

func (v ResetAlarm) String() string {
	var str []string
	str = append(str, "AlarmCode["+v.AlarmCodeString()+"]")
	str = append(str, "ClusterId["+v.ClusterIdString()+"]")
	return "ResetAlarm{" + zcl.StrJoin(str, " ") + "}"
}

func (ResetAlarm) Name() string { return "Reset alarm" }

// Resets all alarms, causing triggered alarms to generate new notification
type ResetAllAlarms struct {
}

const ResetAllAlarmsCommand zcl.CommandID = 1

func (v *ResetAllAlarms) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v ResetAllAlarms) ID() zcl.CommandID {
	return ResetAllAlarmsCommand
}

func (v ResetAllAlarms) Cluster() zcl.ClusterID {
	return AlarmsID
}

func (v ResetAllAlarms) MnfCode() []byte {
	return []byte{}
}

func (v ResetAllAlarms) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *ResetAllAlarms) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

func (v ResetAllAlarms) String() string {
	var str []string
	return "ResetAllAlarms{" + zcl.StrJoin(str, " ") + "}"
}

func (ResetAllAlarms) Name() string { return "Reset all alarms" }

// Retrieves the earliest alarm and removes it from the table
type GetAlarm struct {
}

const GetAlarmCommand zcl.CommandID = 2

func (v *GetAlarm) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v GetAlarm) ID() zcl.CommandID {
	return GetAlarmCommand
}

func (v GetAlarm) Cluster() zcl.ClusterID {
	return AlarmsID
}

func (v GetAlarm) MnfCode() []byte {
	return []byte{}
}

func (v GetAlarm) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *GetAlarm) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

func (v GetAlarm) String() string {
	var str []string
	return "GetAlarm{" + zcl.StrJoin(str, " ") + "}"
}

func (GetAlarm) Name() string { return "Get Alarm" }

// Clears the alarm log
type ResetAlarmLog struct {
}

const ResetAlarmLogCommand zcl.CommandID = 3

func (v *ResetAlarmLog) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v ResetAlarmLog) ID() zcl.CommandID {
	return ResetAlarmLogCommand
}

func (v ResetAlarmLog) Cluster() zcl.ClusterID {
	return AlarmsID
}

func (v ResetAlarmLog) MnfCode() []byte {
	return []byte{}
}

func (v ResetAlarmLog) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *ResetAlarmLog) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

func (v ResetAlarmLog) String() string {
	var str []string
	return "ResetAlarmLog{" + zcl.StrJoin(str, " ") + "}"
}

func (ResetAlarmLog) Name() string { return "Reset alarm log" }

type Alarm struct {
	AlarmCode zcl.Zu8
	ClusterId zcl.Zu16
}

const AlarmCommand zcl.CommandID = 0

func (v *Alarm) Values() []zcl.Val {
	return []zcl.Val{
		&v.AlarmCode,
		&v.ClusterId,
	}
}

func (v Alarm) ID() zcl.CommandID {
	return AlarmCommand
}

func (v Alarm) Cluster() zcl.ClusterID {
	return AlarmsID
}

func (v Alarm) MnfCode() []byte {
	return []byte{}
}

func (v Alarm) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.AlarmCode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ClusterId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

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

func (v Alarm) AlarmCodeString() string {
	return zcl.Sprintf("0x%X", zcl.Zu8(v.AlarmCode))
}
func (v Alarm) ClusterIdString() string {
	return zcl.Sprintf("0x%X", zcl.Zu16(v.ClusterId))
}

func (v Alarm) String() string {
	var str []string
	str = append(str, "AlarmCode["+v.AlarmCodeString()+"]")
	str = append(str, "ClusterId["+v.ClusterIdString()+"]")
	return "Alarm{" + zcl.StrJoin(str, " ") + "}"
}

func (Alarm) Name() string { return "Alarm" }

type GetAlarmResponse struct {
	Status    zcl.Status
	AlarmCode zcl.Zenum8
	ClusterId zcl.Zu16
	Timestamp zcl.Zutc
}

const GetAlarmResponseCommand zcl.CommandID = 1

func (v *GetAlarmResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.Status,
		&v.AlarmCode,
		&v.ClusterId,
		&v.Timestamp,
	}
}

func (v GetAlarmResponse) ID() zcl.CommandID {
	return GetAlarmResponseCommand
}

func (v GetAlarmResponse) Cluster() zcl.ClusterID {
	return AlarmsID
}

func (v GetAlarmResponse) MnfCode() []byte {
	return []byte{}
}

func (v GetAlarmResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Status.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if v.Status == 0x00 {
		if tmp, err = v.AlarmCode.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 {
		if tmp, err = v.ClusterId.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	if v.Status == 0x00 {
		if tmp, err = v.Timestamp.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}

	return data, nil
}

func (v *GetAlarmResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Status).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if v.Status == 0x00 {
		if b, err = (&v.AlarmCode).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 {
		if b, err = (&v.ClusterId).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	if v.Status == 0x00 {
		if b, err = (&v.Timestamp).UnmarshalZcl(b); err != nil {
			return b, err
		}
	}

	return b, nil
}

func (v GetAlarmResponse) StatusString() string {
	return zcl.Sprintf("%v", zcl.Status(v.Status))
}
func (v GetAlarmResponse) AlarmCodeString() string {
	return zcl.Sprintf("%v", zcl.Zenum8(v.AlarmCode))
}
func (v GetAlarmResponse) ClusterIdString() string {
	return zcl.Sprintf("%v", zcl.Zu16(v.ClusterId))
}
func (v GetAlarmResponse) TimestampString() string {
	return zcl.Sprintf("%v", zcl.Zutc(v.Timestamp))
}

func (v GetAlarmResponse) String() string {
	var str []string
	str = append(str, "Status["+v.StatusString()+"]")
	str = append(str, "AlarmCode["+v.AlarmCodeString()+"]")
	str = append(str, "ClusterId["+v.ClusterIdString()+"]")
	str = append(str, "Timestamp["+v.TimestampString()+"]")
	return "GetAlarmResponse{" + zcl.StrJoin(str, " ") + "}"
}

func (GetAlarmResponse) Name() string { return "Get alarm response" }

// AlarmCount is an autogenerated attribute in the Alarms cluster
// Number of alarms currently defined
type AlarmCount zcl.Zu16

const AlarmCountAttr zcl.AttrID = 0

func (AlarmCount) ID() zcl.AttrID                { return AlarmCountAttr }
func (AlarmCount) Cluster() zcl.ClusterID        { return AlarmsID }
func (AlarmCount) Name() string                  { return "Alarm Count" }
func (AlarmCount) Readable() bool                { return true }
func (AlarmCount) Writable() bool                { return true }
func (AlarmCount) Reportable() bool              { return false }
func (AlarmCount) SceneIndex() int               { return -1 }
func (a *AlarmCount) Value() *AlarmCount         { return a }
func (a AlarmCount) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *AlarmCount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmCount(*nt)
	return br, err
}

func (a AlarmCount) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}
