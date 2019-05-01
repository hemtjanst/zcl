// Sending alarm notifications and configuring alarm functionality.
package general

import (
	"neotor.se/zcl"
)

// Alarms
// Sending alarm notifications and configuring alarm functionality.

func NewAlarmsServer(profile zcl.ProfileID) *AlarmsServer { return &AlarmsServer{p: profile} }
func NewAlarmsClient(profile zcl.ProfileID) *AlarmsClient { return &AlarmsClient{p: profile} }

const AlarmsCluster zcl.ClusterID = 9

type AlarmsServer struct {
	p zcl.ProfileID

	AlarmCount *AlarmCount
}

func (s *AlarmsServer) ResetAlarm() *ResetAlarm         { return new(ResetAlarm) }
func (s *AlarmsServer) ResetAllAlarms() *ResetAllAlarms { return new(ResetAllAlarms) }
func (s *AlarmsServer) GetAlarm() *GetAlarm             { return new(GetAlarm) }
func (s *AlarmsServer) ResetAlarmLog() *ResetAlarmLog   { return new(ResetAlarmLog) }

type AlarmsClient struct {
	p zcl.ProfileID
}

func (s *AlarmsClient) Alarm() *Alarm                       { return new(Alarm) }
func (s *AlarmsClient) GetAlarmResponse() *GetAlarmResponse { return new(GetAlarmResponse) }

/*
var AlarmsServer = map[zcl.CommandID]func() zcl.Command{
    ResetAlarmID: func() zcl.Command { return new(ResetAlarm) },
    ResetAllAlarmsID: func() zcl.Command { return new(ResetAllAlarms) },
    GetAlarmID: func() zcl.Command { return new(GetAlarm) },
    ResetAlarmLogID: func() zcl.Command { return new(ResetAlarmLog) },
}

var AlarmsClient = map[zcl.CommandID]func() zcl.Command{
    AlarmID: func() zcl.Command { return new(Alarm) },
    GetAlarmResponseID: func() zcl.Command { return new(GetAlarmResponse) },
}
*/

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
	return AlarmsCluster
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
	return AlarmsCluster
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
	return AlarmsCluster
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
	return AlarmsCluster
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
	return AlarmsCluster
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
	return AlarmsCluster
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

type AlarmCount zcl.Zu16

func (a AlarmCount) ID() zcl.AttrID         { return 0 }
func (a AlarmCount) Cluster() zcl.ClusterID { return AlarmsCluster }
func (a *AlarmCount) Value() *AlarmCount    { return a }
func (a AlarmCount) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *AlarmCount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmCount(*nt)
	return br, err
}

func (a AlarmCount) Readable() bool   { return true }
func (a AlarmCount) Writable() bool   { return true }
func (a AlarmCount) Reportable() bool { return false }
func (a AlarmCount) SceneIndex() int  { return -1 }

func (a AlarmCount) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}
