// The door lock cluster provides an interface to a generic way to secure a door.
package closures

import (
	"neotor.se/zcl"
)

// DoorLock
const DoorLockID zcl.ClusterID = 257

var DoorLockCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		LockDoorCommand:                    func() zcl.Command { return new(LockDoor) },
		UnlockDoorCommand:                  func() zcl.Command { return new(UnlockDoor) },
		ToggleLockCommand:                  func() zcl.Command { return new(ToggleLock) },
		GetLogRecordCommand:                func() zcl.Command { return new(GetLogRecord) },
		OperationgEventNotificationCommand: func() zcl.Command { return new(OperationgEventNotification) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{
		GetLogRecordResponseCommand: func() zcl.Command { return new(GetLogRecordResponse) },
	},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		LockStateAttr:                                func() zcl.Attr { return new(LockState) },
		LockTypeAttr:                                 func() zcl.Attr { return new(LockType) },
		ActuatorEnabledAttr:                          func() zcl.Attr { return new(ActuatorEnabled) },
		NumberOfLogRecordsSupportedAttr:              func() zcl.Attr { return new(NumberOfLogRecordsSupported) },
		NumberOfTotalUsersSupportedAttr:              func() zcl.Attr { return new(NumberOfTotalUsersSupported) },
		NumberOfPinUsersSupportedAttr:                func() zcl.Attr { return new(NumberOfPinUsersSupported) },
		NumberOfRfidUsersSupportedAttr:               func() zcl.Attr { return new(NumberOfRfidUsersSupported) },
		NumberOfWeekdaySchedulesSupportedPerUserAttr: func() zcl.Attr { return new(NumberOfWeekdaySchedulesSupportedPerUser) },
		NumberOfYearDaySchedulesSupportedPerUserAttr: func() zcl.Attr { return new(NumberOfYearDaySchedulesSupportedPerUser) },
		NumberOfHolidaySchedulesSupportedAttr:        func() zcl.Attr { return new(NumberOfHolidaySchedulesSupported) },
		EnableLoggingAttr:                            func() zcl.Attr { return new(EnableLogging) },
		ZigbeeSecurityLevelAttr:                      func() zcl.Attr { return new(ZigbeeSecurityLevel) },
		AlarmMaskAttr:                                func() zcl.Attr { return new(AlarmMask) },
		RfOperationEventMaskAttr:                     func() zcl.Attr { return new(RfOperationEventMask) },
		ManualOperationEventMaskAttr:                 func() zcl.Attr { return new(ManualOperationEventMask) },
		EventTypeAttr:                                func() zcl.Attr { return new(EventType) },
		TiltAngleAttr:                                func() zcl.Attr { return new(TiltAngle) },
		VibrationStrengthAttr:                        func() zcl.Attr { return new(VibrationStrength) },
		OrientationAttr:                              func() zcl.Attr { return new(Orientation) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// This command causes the lock device to lock the door.
type LockDoor struct {
}

const LockDoorCommand zcl.CommandID = 0

func (v *LockDoor) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v LockDoor) ID() zcl.CommandID {
	return LockDoorCommand
}

func (v LockDoor) Cluster() zcl.ClusterID {
	return DoorLockID
}

func (v LockDoor) MnfCode() []byte {
	return []byte{}
}

func (v LockDoor) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *LockDoor) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// This command causes the lock device to unlock the door.
type UnlockDoor struct {
}

const UnlockDoorCommand zcl.CommandID = 1

func (v *UnlockDoor) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v UnlockDoor) ID() zcl.CommandID {
	return UnlockDoorCommand
}

func (v UnlockDoor) Cluster() zcl.ClusterID {
	return DoorLockID
}

func (v UnlockDoor) MnfCode() []byte {
	return []byte{}
}

func (v UnlockDoor) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *UnlockDoor) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// This command causes the lock device to toggle.
type ToggleLock struct {
}

const ToggleLockCommand zcl.CommandID = 2

func (v *ToggleLock) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v ToggleLock) ID() zcl.CommandID {
	return ToggleLockCommand
}

func (v ToggleLock) Cluster() zcl.ClusterID {
	return DoorLockID
}

func (v ToggleLock) MnfCode() []byte {
	return []byte{}
}

func (v ToggleLock) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *ToggleLock) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// Request a log record.
type GetLogRecord struct {
	LogIndex zcl.Zu16
}

const GetLogRecordCommand zcl.CommandID = 4

func (v *GetLogRecord) Values() []zcl.Val {
	return []zcl.Val{
		&v.LogIndex,
	}
}

func (v GetLogRecord) ID() zcl.CommandID {
	return GetLogRecordCommand
}

func (v GetLogRecord) Cluster() zcl.ClusterID {
	return DoorLockID
}

func (v GetLogRecord) MnfCode() []byte {
	return []byte{}
}

func (v GetLogRecord) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.LogIndex.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetLogRecord) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.LogIndex).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// The door lock server sends out operation event notification when the event is triggered by the various event sources.
type OperationgEventNotification struct {
	OperationEventSource zcl.Zu8
	OperationEventCode   zcl.Zu8
	UserId               zcl.Zu16
	Pin                  zcl.Zu8
	ZigbeeLocalTime      zcl.Zu32
	Data                 zcl.Zcstring
}

const OperationgEventNotificationCommand zcl.CommandID = 32

func (v *OperationgEventNotification) Values() []zcl.Val {
	return []zcl.Val{
		&v.OperationEventSource,
		&v.OperationEventCode,
		&v.UserId,
		&v.Pin,
		&v.ZigbeeLocalTime,
		&v.Data,
	}
}

func (v OperationgEventNotification) ID() zcl.CommandID {
	return OperationgEventNotificationCommand
}

func (v OperationgEventNotification) Cluster() zcl.ClusterID {
	return DoorLockID
}

func (v OperationgEventNotification) MnfCode() []byte {
	return []byte{}
}

func (v OperationgEventNotification) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.OperationEventSource.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.OperationEventCode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.UserId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Pin.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ZigbeeLocalTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Data.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *OperationgEventNotification) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.OperationEventSource).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.OperationEventCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.UserId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Pin).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ZigbeeLocalTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Data).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Returns the specified log record.
type GetLogRecordResponse struct {
	LogEntryId       zcl.Zu16
	Timestamp        zcl.Zu32
	EventType        zcl.Zenum8
	Source           zcl.Zu8
	EventIdAlarmCode zcl.Zu8
	UserId           zcl.Zu16
	Pin              zcl.Zcstring
}

const GetLogRecordResponseCommand zcl.CommandID = 4

func (v *GetLogRecordResponse) Values() []zcl.Val {
	return []zcl.Val{
		&v.LogEntryId,
		&v.Timestamp,
		&v.EventType,
		&v.Source,
		&v.EventIdAlarmCode,
		&v.UserId,
		&v.Pin,
	}
}

func (v GetLogRecordResponse) ID() zcl.CommandID {
	return GetLogRecordResponseCommand
}

func (v GetLogRecordResponse) Cluster() zcl.ClusterID {
	return DoorLockID
}

func (v GetLogRecordResponse) MnfCode() []byte {
	return []byte{}
}

func (v GetLogRecordResponse) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.LogEntryId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Timestamp.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.EventType.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Source.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.EventIdAlarmCode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.UserId.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Pin.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *GetLogRecordResponse) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.LogEntryId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Timestamp).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.EventType).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Source).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.EventIdAlarmCode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.UserId).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Pin).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// LockState is an autogenerated attribute in the DoorLock cluster
type LockState zcl.Zenum8

const LockStateAttr zcl.AttrID = 0

func (a LockState) ID() zcl.AttrID         { return LockStateAttr }
func (a LockState) Cluster() zcl.ClusterID { return DoorLockID }
func (a *LockState) Value() *LockState     { return a }
func (a LockState) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *LockState) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LockState(*nt)
	return br, err
}

func (a LockState) Readable() bool   { return true }
func (a LockState) Writable() bool   { return false }
func (a LockState) Reportable() bool { return false }
func (a LockState) SceneIndex() int  { return -1 }

func (a LockState) String() string {
	switch a {
	case 0x00:
		return "Not fully locked"
	case 0x01:
		return "Locked"
	case 0x02:
		return "Unlocked"
	case 0xFF:
		return "Undefined"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsNotFullyLocked checks if LockState equals the value for Not fully locked (0x00)
func (a LockState) IsNotFullyLocked() bool { return a == 0x00 }

// SetNotFullyLocked sets LockState to Not fully locked (0x00)
func (a *LockState) SetNotFullyLocked() { *a = 0x00 }

// IsLocked checks if LockState equals the value for Locked (0x01)
func (a LockState) IsLocked() bool { return a == 0x01 }

// SetLocked sets LockState to Locked (0x01)
func (a *LockState) SetLocked() { *a = 0x01 }

// IsUnlocked checks if LockState equals the value for Unlocked (0x02)
func (a LockState) IsUnlocked() bool { return a == 0x02 }

// SetUnlocked sets LockState to Unlocked (0x02)
func (a *LockState) SetUnlocked() { *a = 0x02 }

// IsUndefined checks if LockState equals the value for Undefined (0xFF)
func (a LockState) IsUndefined() bool { return a == 0xFF }

// SetUndefined sets LockState to Undefined (0xFF)
func (a *LockState) SetUndefined() { *a = 0xFF }

// LockType is an autogenerated attribute in the DoorLock cluster
type LockType zcl.Zenum8

const LockTypeAttr zcl.AttrID = 1

func (a LockType) ID() zcl.AttrID         { return LockTypeAttr }
func (a LockType) Cluster() zcl.ClusterID { return DoorLockID }
func (a *LockType) Value() *LockType      { return a }
func (a LockType) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *LockType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = LockType(*nt)
	return br, err
}

func (a LockType) Readable() bool   { return true }
func (a LockType) Writable() bool   { return false }
func (a LockType) Reportable() bool { return false }
func (a LockType) SceneIndex() int  { return -1 }

func (a LockType) String() string {
	switch a {
	case 0x00:
		return "Dead bolt"
	case 0x01:
		return "Magnetic"
	case 0x02:
		return "Other"
	case 0x03:
		return "Mortise"
	case 0x04:
		return "Rim"
	case 0x05:
		return "Latch Bolt"
	case 0x06:
		return "Cylindrical Lock"
	case 0x07:
		return "Tubular Lock"
	case 0x08:
		return "Interconnected Lock"
	case 0x09:
		return "Dead Latch"
	case 0x0A:
		return "Door Furniture"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsDeadBolt checks if LockType equals the value for Dead bolt (0x00)
func (a LockType) IsDeadBolt() bool { return a == 0x00 }

// SetDeadBolt sets LockType to Dead bolt (0x00)
func (a *LockType) SetDeadBolt() { *a = 0x00 }

// IsMagnetic checks if LockType equals the value for Magnetic (0x01)
func (a LockType) IsMagnetic() bool { return a == 0x01 }

// SetMagnetic sets LockType to Magnetic (0x01)
func (a *LockType) SetMagnetic() { *a = 0x01 }

// IsOther checks if LockType equals the value for Other (0x02)
func (a LockType) IsOther() bool { return a == 0x02 }

// SetOther sets LockType to Other (0x02)
func (a *LockType) SetOther() { *a = 0x02 }

// IsMortise checks if LockType equals the value for Mortise (0x03)
func (a LockType) IsMortise() bool { return a == 0x03 }

// SetMortise sets LockType to Mortise (0x03)
func (a *LockType) SetMortise() { *a = 0x03 }

// IsRim checks if LockType equals the value for Rim (0x04)
func (a LockType) IsRim() bool { return a == 0x04 }

// SetRim sets LockType to Rim (0x04)
func (a *LockType) SetRim() { *a = 0x04 }

// IsLatchBolt checks if LockType equals the value for Latch Bolt (0x05)
func (a LockType) IsLatchBolt() bool { return a == 0x05 }

// SetLatchBolt sets LockType to Latch Bolt (0x05)
func (a *LockType) SetLatchBolt() { *a = 0x05 }

// IsCylindricalLock checks if LockType equals the value for Cylindrical Lock (0x06)
func (a LockType) IsCylindricalLock() bool { return a == 0x06 }

// SetCylindricalLock sets LockType to Cylindrical Lock (0x06)
func (a *LockType) SetCylindricalLock() { *a = 0x06 }

// IsTubularLock checks if LockType equals the value for Tubular Lock (0x07)
func (a LockType) IsTubularLock() bool { return a == 0x07 }

// SetTubularLock sets LockType to Tubular Lock (0x07)
func (a *LockType) SetTubularLock() { *a = 0x07 }

// IsInterconnectedLock checks if LockType equals the value for Interconnected Lock (0x08)
func (a LockType) IsInterconnectedLock() bool { return a == 0x08 }

// SetInterconnectedLock sets LockType to Interconnected Lock (0x08)
func (a *LockType) SetInterconnectedLock() { *a = 0x08 }

// IsDeadLatch checks if LockType equals the value for Dead Latch (0x09)
func (a LockType) IsDeadLatch() bool { return a == 0x09 }

// SetDeadLatch sets LockType to Dead Latch (0x09)
func (a *LockType) SetDeadLatch() { *a = 0x09 }

// IsDoorFurniture checks if LockType equals the value for Door Furniture (0x0A)
func (a LockType) IsDoorFurniture() bool { return a == 0x0A }

// SetDoorFurniture sets LockType to Door Furniture (0x0A)
func (a *LockType) SetDoorFurniture() { *a = 0x0A }

// ActuatorEnabled is an autogenerated attribute in the DoorLock cluster
type ActuatorEnabled zcl.Zbool

const ActuatorEnabledAttr zcl.AttrID = 2

func (a ActuatorEnabled) ID() zcl.AttrID           { return ActuatorEnabledAttr }
func (a ActuatorEnabled) Cluster() zcl.ClusterID   { return DoorLockID }
func (a *ActuatorEnabled) Value() *ActuatorEnabled { return a }
func (a ActuatorEnabled) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *ActuatorEnabled) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = ActuatorEnabled(*nt)
	return br, err
}

func (a ActuatorEnabled) Readable() bool   { return true }
func (a ActuatorEnabled) Writable() bool   { return false }
func (a ActuatorEnabled) Reportable() bool { return false }
func (a ActuatorEnabled) SceneIndex() int  { return -1 }

func (a ActuatorEnabled) String() string {
	return zcl.Sprintf("%s", zcl.Zbool(a))
}

// NumberOfLogRecordsSupported is an autogenerated attribute in the DoorLock cluster
type NumberOfLogRecordsSupported zcl.Zu16

const NumberOfLogRecordsSupportedAttr zcl.AttrID = 16

func (a NumberOfLogRecordsSupported) ID() zcl.AttrID                       { return NumberOfLogRecordsSupportedAttr }
func (a NumberOfLogRecordsSupported) Cluster() zcl.ClusterID               { return DoorLockID }
func (a *NumberOfLogRecordsSupported) Value() *NumberOfLogRecordsSupported { return a }
func (a NumberOfLogRecordsSupported) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NumberOfLogRecordsSupported) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfLogRecordsSupported(*nt)
	return br, err
}

func (a NumberOfLogRecordsSupported) Readable() bool   { return true }
func (a NumberOfLogRecordsSupported) Writable() bool   { return false }
func (a NumberOfLogRecordsSupported) Reportable() bool { return false }
func (a NumberOfLogRecordsSupported) SceneIndex() int  { return -1 }

func (a NumberOfLogRecordsSupported) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// NumberOfTotalUsersSupported is an autogenerated attribute in the DoorLock cluster
type NumberOfTotalUsersSupported zcl.Zu16

const NumberOfTotalUsersSupportedAttr zcl.AttrID = 17

func (a NumberOfTotalUsersSupported) ID() zcl.AttrID                       { return NumberOfTotalUsersSupportedAttr }
func (a NumberOfTotalUsersSupported) Cluster() zcl.ClusterID               { return DoorLockID }
func (a *NumberOfTotalUsersSupported) Value() *NumberOfTotalUsersSupported { return a }
func (a NumberOfTotalUsersSupported) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NumberOfTotalUsersSupported) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfTotalUsersSupported(*nt)
	return br, err
}

func (a NumberOfTotalUsersSupported) Readable() bool   { return true }
func (a NumberOfTotalUsersSupported) Writable() bool   { return false }
func (a NumberOfTotalUsersSupported) Reportable() bool { return false }
func (a NumberOfTotalUsersSupported) SceneIndex() int  { return -1 }

func (a NumberOfTotalUsersSupported) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// NumberOfPinUsersSupported is an autogenerated attribute in the DoorLock cluster
type NumberOfPinUsersSupported zcl.Zu16

const NumberOfPinUsersSupportedAttr zcl.AttrID = 18

func (a NumberOfPinUsersSupported) ID() zcl.AttrID                     { return NumberOfPinUsersSupportedAttr }
func (a NumberOfPinUsersSupported) Cluster() zcl.ClusterID             { return DoorLockID }
func (a *NumberOfPinUsersSupported) Value() *NumberOfPinUsersSupported { return a }
func (a NumberOfPinUsersSupported) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NumberOfPinUsersSupported) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfPinUsersSupported(*nt)
	return br, err
}

func (a NumberOfPinUsersSupported) Readable() bool   { return true }
func (a NumberOfPinUsersSupported) Writable() bool   { return false }
func (a NumberOfPinUsersSupported) Reportable() bool { return false }
func (a NumberOfPinUsersSupported) SceneIndex() int  { return -1 }

func (a NumberOfPinUsersSupported) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// NumberOfRfidUsersSupported is an autogenerated attribute in the DoorLock cluster
type NumberOfRfidUsersSupported zcl.Zu16

const NumberOfRfidUsersSupportedAttr zcl.AttrID = 19

func (a NumberOfRfidUsersSupported) ID() zcl.AttrID                      { return NumberOfRfidUsersSupportedAttr }
func (a NumberOfRfidUsersSupported) Cluster() zcl.ClusterID              { return DoorLockID }
func (a *NumberOfRfidUsersSupported) Value() *NumberOfRfidUsersSupported { return a }
func (a NumberOfRfidUsersSupported) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *NumberOfRfidUsersSupported) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfRfidUsersSupported(*nt)
	return br, err
}

func (a NumberOfRfidUsersSupported) Readable() bool   { return true }
func (a NumberOfRfidUsersSupported) Writable() bool   { return false }
func (a NumberOfRfidUsersSupported) Reportable() bool { return false }
func (a NumberOfRfidUsersSupported) SceneIndex() int  { return -1 }

func (a NumberOfRfidUsersSupported) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// NumberOfWeekdaySchedulesSupportedPerUser is an autogenerated attribute in the DoorLock cluster
type NumberOfWeekdaySchedulesSupportedPerUser zcl.Zu8

const NumberOfWeekdaySchedulesSupportedPerUserAttr zcl.AttrID = 20

func (a NumberOfWeekdaySchedulesSupportedPerUser) ID() zcl.AttrID {
	return NumberOfWeekdaySchedulesSupportedPerUserAttr
}
func (a NumberOfWeekdaySchedulesSupportedPerUser) Cluster() zcl.ClusterID { return DoorLockID }
func (a *NumberOfWeekdaySchedulesSupportedPerUser) Value() *NumberOfWeekdaySchedulesSupportedPerUser {
	return a
}
func (a NumberOfWeekdaySchedulesSupportedPerUser) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *NumberOfWeekdaySchedulesSupportedPerUser) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfWeekdaySchedulesSupportedPerUser(*nt)
	return br, err
}

func (a NumberOfWeekdaySchedulesSupportedPerUser) Readable() bool   { return true }
func (a NumberOfWeekdaySchedulesSupportedPerUser) Writable() bool   { return false }
func (a NumberOfWeekdaySchedulesSupportedPerUser) Reportable() bool { return false }
func (a NumberOfWeekdaySchedulesSupportedPerUser) SceneIndex() int  { return -1 }

func (a NumberOfWeekdaySchedulesSupportedPerUser) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

// NumberOfYearDaySchedulesSupportedPerUser is an autogenerated attribute in the DoorLock cluster
type NumberOfYearDaySchedulesSupportedPerUser zcl.Zu8

const NumberOfYearDaySchedulesSupportedPerUserAttr zcl.AttrID = 21

func (a NumberOfYearDaySchedulesSupportedPerUser) ID() zcl.AttrID {
	return NumberOfYearDaySchedulesSupportedPerUserAttr
}
func (a NumberOfYearDaySchedulesSupportedPerUser) Cluster() zcl.ClusterID { return DoorLockID }
func (a *NumberOfYearDaySchedulesSupportedPerUser) Value() *NumberOfYearDaySchedulesSupportedPerUser {
	return a
}
func (a NumberOfYearDaySchedulesSupportedPerUser) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *NumberOfYearDaySchedulesSupportedPerUser) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfYearDaySchedulesSupportedPerUser(*nt)
	return br, err
}

func (a NumberOfYearDaySchedulesSupportedPerUser) Readable() bool   { return true }
func (a NumberOfYearDaySchedulesSupportedPerUser) Writable() bool   { return false }
func (a NumberOfYearDaySchedulesSupportedPerUser) Reportable() bool { return false }
func (a NumberOfYearDaySchedulesSupportedPerUser) SceneIndex() int  { return -1 }

func (a NumberOfYearDaySchedulesSupportedPerUser) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

// NumberOfHolidaySchedulesSupported is an autogenerated attribute in the DoorLock cluster
type NumberOfHolidaySchedulesSupported zcl.Zu8

const NumberOfHolidaySchedulesSupportedAttr zcl.AttrID = 22

func (a NumberOfHolidaySchedulesSupported) ID() zcl.AttrID {
	return NumberOfHolidaySchedulesSupportedAttr
}
func (a NumberOfHolidaySchedulesSupported) Cluster() zcl.ClusterID                     { return DoorLockID }
func (a *NumberOfHolidaySchedulesSupported) Value() *NumberOfHolidaySchedulesSupported { return a }
func (a NumberOfHolidaySchedulesSupported) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *NumberOfHolidaySchedulesSupported) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfHolidaySchedulesSupported(*nt)
	return br, err
}

func (a NumberOfHolidaySchedulesSupported) Readable() bool   { return true }
func (a NumberOfHolidaySchedulesSupported) Writable() bool   { return false }
func (a NumberOfHolidaySchedulesSupported) Reportable() bool { return false }
func (a NumberOfHolidaySchedulesSupported) SceneIndex() int  { return -1 }

func (a NumberOfHolidaySchedulesSupported) String() string {
	return zcl.Sprintf("%s", zcl.Zu8(a))
}

// EnableLogging is an autogenerated attribute in the DoorLock cluster
type EnableLogging zcl.Zbool

const EnableLoggingAttr zcl.AttrID = 32

func (a EnableLogging) ID() zcl.AttrID         { return EnableLoggingAttr }
func (a EnableLogging) Cluster() zcl.ClusterID { return DoorLockID }
func (a *EnableLogging) Value() *EnableLogging { return a }
func (a EnableLogging) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *EnableLogging) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = EnableLogging(*nt)
	return br, err
}

func (a EnableLogging) Readable() bool   { return true }
func (a EnableLogging) Writable() bool   { return true }
func (a EnableLogging) Reportable() bool { return false }
func (a EnableLogging) SceneIndex() int  { return -1 }

func (a EnableLogging) String() string {
	return zcl.Sprintf("%s", zcl.Zbool(a))
}

// ZigbeeSecurityLevel is an autogenerated attribute in the DoorLock cluster
type ZigbeeSecurityLevel zcl.Zenum8

const ZigbeeSecurityLevelAttr zcl.AttrID = 52

func (a ZigbeeSecurityLevel) ID() zcl.AttrID               { return ZigbeeSecurityLevelAttr }
func (a ZigbeeSecurityLevel) Cluster() zcl.ClusterID       { return DoorLockID }
func (a *ZigbeeSecurityLevel) Value() *ZigbeeSecurityLevel { return a }
func (a ZigbeeSecurityLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *ZigbeeSecurityLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = ZigbeeSecurityLevel(*nt)
	return br, err
}

func (a ZigbeeSecurityLevel) Readable() bool   { return true }
func (a ZigbeeSecurityLevel) Writable() bool   { return false }
func (a ZigbeeSecurityLevel) Reportable() bool { return false }
func (a ZigbeeSecurityLevel) SceneIndex() int  { return -1 }

func (a ZigbeeSecurityLevel) String() string {
	switch a {
	case 0x00:
		return "Network Security"
	case 0x01:
		return "APS Security"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsNetworkSecurity checks if ZigbeeSecurityLevel equals the value for Network Security (0x00)
func (a ZigbeeSecurityLevel) IsNetworkSecurity() bool { return a == 0x00 }

// SetNetworkSecurity sets ZigbeeSecurityLevel to Network Security (0x00)
func (a *ZigbeeSecurityLevel) SetNetworkSecurity() { *a = 0x00 }

// IsApsSecurity checks if ZigbeeSecurityLevel equals the value for APS Security (0x01)
func (a ZigbeeSecurityLevel) IsApsSecurity() bool { return a == 0x01 }

// SetApsSecurity sets ZigbeeSecurityLevel to APS Security (0x01)
func (a *ZigbeeSecurityLevel) SetApsSecurity() { *a = 0x01 }

// AlarmMask is an autogenerated attribute in the DoorLock cluster
type AlarmMask zcl.Zbmp16

const AlarmMaskAttr zcl.AttrID = 64

func (a AlarmMask) ID() zcl.AttrID         { return AlarmMaskAttr }
func (a AlarmMask) Cluster() zcl.ClusterID { return DoorLockID }
func (a *AlarmMask) Value() *AlarmMask     { return a }
func (a AlarmMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp16(a).MarshalZcl()
}
func (a *AlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = AlarmMask(*nt)
	return br, err
}

func (a AlarmMask) Readable() bool   { return true }
func (a AlarmMask) Writable() bool   { return true }
func (a AlarmMask) Reportable() bool { return false }
func (a AlarmMask) SceneIndex() int  { return -1 }

func (a AlarmMask) String() string {
	var bstr []string
	if a.IsDeadboltJammed() {
		bstr = append(bstr, "Deadbolt Jammed")
	}
	if a.IsLockResetToFactoryDefaults() {
		bstr = append(bstr, "Lock Reset to Factory Defaults")
	}
	if a.IsReserved() {
		bstr = append(bstr, "Reserved")
	}
	if a.IsRfModulePowerCycled() {
		bstr = append(bstr, "RF Module Power Cycled")
	}
	if a.IsTamperAlarmWrongCodeEntryLimit() {
		bstr = append(bstr, "Tamper Alarm - wrong code entry limit")
	}
	if a.IsTamperAlarmFrontEscutcheonRemovedFromMain() {
		bstr = append(bstr, "Tamper Alarm - front escutcheon removed from main")
	}
	if a.IsForcedDoorOpenUnderDoorLockedCondition() {
		bstr = append(bstr, "Forced Door Open under Door Locked Condition")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a AlarmMask) IsDeadboltJammed() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *AlarmMask) SetDeadboltJammed(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a AlarmMask) IsLockResetToFactoryDefaults() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *AlarmMask) SetLockResetToFactoryDefaults(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a AlarmMask) IsReserved() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *AlarmMask) SetReserved(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a AlarmMask) IsRfModulePowerCycled() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *AlarmMask) SetRfModulePowerCycled(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a AlarmMask) IsTamperAlarmWrongCodeEntryLimit() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *AlarmMask) SetTamperAlarmWrongCodeEntryLimit(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a AlarmMask) IsTamperAlarmFrontEscutcheonRemovedFromMain() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *AlarmMask) SetTamperAlarmFrontEscutcheonRemovedFromMain(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a AlarmMask) IsForcedDoorOpenUnderDoorLockedCondition() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *AlarmMask) SetForcedDoorOpenUnderDoorLockedCondition(b bool) {
	*a = AlarmMask(zcl.BitmapSet([]byte(*a), 6, b))
}

// RfOperationEventMask is an autogenerated attribute in the DoorLock cluster
type RfOperationEventMask zcl.Zbmp16

const RfOperationEventMaskAttr zcl.AttrID = 66

func (a RfOperationEventMask) ID() zcl.AttrID                { return RfOperationEventMaskAttr }
func (a RfOperationEventMask) Cluster() zcl.ClusterID        { return DoorLockID }
func (a *RfOperationEventMask) Value() *RfOperationEventMask { return a }
func (a RfOperationEventMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp16(a).MarshalZcl()
}
func (a *RfOperationEventMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = RfOperationEventMask(*nt)
	return br, err
}

func (a RfOperationEventMask) Readable() bool   { return true }
func (a RfOperationEventMask) Writable() bool   { return true }
func (a RfOperationEventMask) Reportable() bool { return false }
func (a RfOperationEventMask) SceneIndex() int  { return -1 }

func (a RfOperationEventMask) String() string {
	var bstr []string
	if a.IsUnknownOrManufacturerSpecificRfOperationEvent() {
		bstr = append(bstr, "Unknown or manufacturer-specific RF operation event")
	}
	if a.IsLockSourceRf() {
		bstr = append(bstr, "Lock, source: RF")
	}
	if a.IsUnlockSourceRf() {
		bstr = append(bstr, "Unlock, source: RF")
	}
	if a.IsLockSourceRfErrorInvalidCode() {
		bstr = append(bstr, "Lock, source: RF, error: invalid code")
	}
	if a.IsLockSourceRfErrorInvalidSchedule() {
		bstr = append(bstr, "Lock, source: RF, error: invalid schedule")
	}
	if a.IsUnlockSourceRfErrorInvalidCode() {
		bstr = append(bstr, "Unlock, source: RF, error: invalid code")
	}
	if a.IsUnlockSourceRfErrorInvalidSchedule() {
		bstr = append(bstr, "Unlock, source: RF, error: invalid schedule")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a RfOperationEventMask) IsUnknownOrManufacturerSpecificRfOperationEvent() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *RfOperationEventMask) SetUnknownOrManufacturerSpecificRfOperationEvent(b bool) {
	*a = RfOperationEventMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a RfOperationEventMask) IsLockSourceRf() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *RfOperationEventMask) SetLockSourceRf(b bool) {
	*a = RfOperationEventMask(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a RfOperationEventMask) IsUnlockSourceRf() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *RfOperationEventMask) SetUnlockSourceRf(b bool) {
	*a = RfOperationEventMask(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a RfOperationEventMask) IsLockSourceRfErrorInvalidCode() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *RfOperationEventMask) SetLockSourceRfErrorInvalidCode(b bool) {
	*a = RfOperationEventMask(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a RfOperationEventMask) IsLockSourceRfErrorInvalidSchedule() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *RfOperationEventMask) SetLockSourceRfErrorInvalidSchedule(b bool) {
	*a = RfOperationEventMask(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a RfOperationEventMask) IsUnlockSourceRfErrorInvalidCode() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *RfOperationEventMask) SetUnlockSourceRfErrorInvalidCode(b bool) {
	*a = RfOperationEventMask(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a RfOperationEventMask) IsUnlockSourceRfErrorInvalidSchedule() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *RfOperationEventMask) SetUnlockSourceRfErrorInvalidSchedule(b bool) {
	*a = RfOperationEventMask(zcl.BitmapSet([]byte(*a), 6, b))
}

// ManualOperationEventMask is an autogenerated attribute in the DoorLock cluster
type ManualOperationEventMask zcl.Zbmp16

const ManualOperationEventMaskAttr zcl.AttrID = 67

func (a ManualOperationEventMask) ID() zcl.AttrID                    { return ManualOperationEventMaskAttr }
func (a ManualOperationEventMask) Cluster() zcl.ClusterID            { return DoorLockID }
func (a *ManualOperationEventMask) Value() *ManualOperationEventMask { return a }
func (a ManualOperationEventMask) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp16(a).MarshalZcl()
}
func (a *ManualOperationEventMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = ManualOperationEventMask(*nt)
	return br, err
}

func (a ManualOperationEventMask) Readable() bool   { return true }
func (a ManualOperationEventMask) Writable() bool   { return true }
func (a ManualOperationEventMask) Reportable() bool { return false }
func (a ManualOperationEventMask) SceneIndex() int  { return -1 }

func (a ManualOperationEventMask) String() string {
	var bstr []string
	if a.IsUnknownOrManufacturerSpecificManualOperationEvent() {
		bstr = append(bstr, "Unknown or manufacturer-specific manual operation event")
	}
	if a.IsThumbturnLock() {
		bstr = append(bstr, "Thumbturn Lock")
	}
	if a.IsManualUnlockKeyOrThumbturn() {
		bstr = append(bstr, "Manual Unlock (Key or Thumbturn)")
	}
	if a.IsThumbturnUnlock() {
		bstr = append(bstr, "Thumbturn Unlock")
	}
	if a.IsOneTouchLock() {
		bstr = append(bstr, "One touch lock")
	}
	if a.IsKeyLock() {
		bstr = append(bstr, "Key Lock")
	}
	if a.IsKeyUnlock() {
		bstr = append(bstr, "Key Unlock")
	}
	if a.IsAutoLock() {
		bstr = append(bstr, "Auto lock")
	}
	if a.IsScheduleLock() {
		bstr = append(bstr, "Schedule Lock")
	}
	if a.IsScheduleUnlock() {
		bstr = append(bstr, "Schedule Unlock")
	}
	if a.IsManualLockKeyOrThumbturn() {
		bstr = append(bstr, "Manual Lock (Key or Thumbturn)")
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a ManualOperationEventMask) IsUnknownOrManufacturerSpecificManualOperationEvent() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *ManualOperationEventMask) SetUnknownOrManufacturerSpecificManualOperationEvent(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a ManualOperationEventMask) IsThumbturnLock() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *ManualOperationEventMask) SetThumbturnLock(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a ManualOperationEventMask) IsManualUnlockKeyOrThumbturn() bool {
	return zcl.BitmapTest([]byte(a), 10)
}
func (a *ManualOperationEventMask) SetManualUnlockKeyOrThumbturn(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 10, b))
}

func (a ManualOperationEventMask) IsThumbturnUnlock() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *ManualOperationEventMask) SetThumbturnUnlock(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a ManualOperationEventMask) IsOneTouchLock() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *ManualOperationEventMask) SetOneTouchLock(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a ManualOperationEventMask) IsKeyLock() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *ManualOperationEventMask) SetKeyLock(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a ManualOperationEventMask) IsKeyUnlock() bool {
	return zcl.BitmapTest([]byte(a), 5)
}
func (a *ManualOperationEventMask) SetKeyUnlock(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 5, b))
}

func (a ManualOperationEventMask) IsAutoLock() bool {
	return zcl.BitmapTest([]byte(a), 6)
}
func (a *ManualOperationEventMask) SetAutoLock(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 6, b))
}

func (a ManualOperationEventMask) IsScheduleLock() bool {
	return zcl.BitmapTest([]byte(a), 7)
}
func (a *ManualOperationEventMask) SetScheduleLock(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 7, b))
}

func (a ManualOperationEventMask) IsScheduleUnlock() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *ManualOperationEventMask) SetScheduleUnlock(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 8, b))
}

func (a ManualOperationEventMask) IsManualLockKeyOrThumbturn() bool {
	return zcl.BitmapTest([]byte(a), 9)
}
func (a *ManualOperationEventMask) SetManualLockKeyOrThumbturn(b bool) {
	*a = ManualOperationEventMask(zcl.BitmapSet([]byte(*a), 9, b))
}

// EventType is an autogenerated attribute in the DoorLock cluster
type EventType zcl.Zu16

const EventTypeAttr zcl.AttrID = 85

func (a EventType) ID() zcl.AttrID         { return EventTypeAttr }
func (a EventType) Cluster() zcl.ClusterID { return DoorLockID }
func (a *EventType) Value() *EventType     { return a }
func (a EventType) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *EventType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = EventType(*nt)
	return br, err
}

func (a EventType) Readable() bool   { return true }
func (a EventType) Writable() bool   { return false }
func (a EventType) Reportable() bool { return false }
func (a EventType) SceneIndex() int  { return -1 }

func (a EventType) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// TiltAngle is an autogenerated attribute in the DoorLock cluster
type TiltAngle zcl.Zu16

const TiltAngleAttr zcl.AttrID = 1283

func (a TiltAngle) ID() zcl.AttrID         { return TiltAngleAttr }
func (a TiltAngle) Cluster() zcl.ClusterID { return DoorLockID }
func (a *TiltAngle) Value() *TiltAngle     { return a }
func (a TiltAngle) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *TiltAngle) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TiltAngle(*nt)
	return br, err
}

func (a TiltAngle) Readable() bool   { return true }
func (a TiltAngle) Writable() bool   { return false }
func (a TiltAngle) Reportable() bool { return false }
func (a TiltAngle) SceneIndex() int  { return -1 }

func (a TiltAngle) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// VibrationStrength is an autogenerated attribute in the DoorLock cluster
type VibrationStrength zcl.Zu32

const VibrationStrengthAttr zcl.AttrID = 1285

func (a VibrationStrength) ID() zcl.AttrID             { return VibrationStrengthAttr }
func (a VibrationStrength) Cluster() zcl.ClusterID     { return DoorLockID }
func (a *VibrationStrength) Value() *VibrationStrength { return a }
func (a VibrationStrength) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *VibrationStrength) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = VibrationStrength(*nt)
	return br, err
}

func (a VibrationStrength) Readable() bool   { return true }
func (a VibrationStrength) Writable() bool   { return false }
func (a VibrationStrength) Reportable() bool { return false }
func (a VibrationStrength) SceneIndex() int  { return -1 }

func (a VibrationStrength) String() string {
	return zcl.Sprintf("0x%X", zcl.Zu32(a))
}

// Orientation is an autogenerated attribute in the DoorLock cluster
type Orientation zcl.Zu48

const OrientationAttr zcl.AttrID = 1288

func (a Orientation) ID() zcl.AttrID         { return OrientationAttr }
func (a Orientation) Cluster() zcl.ClusterID { return DoorLockID }
func (a *Orientation) Value() *Orientation   { return a }
func (a Orientation) MarshalZcl() ([]byte, error) {
	return zcl.Zu48(a).MarshalZcl()
}
func (a *Orientation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu48)
	br, err := nt.UnmarshalZcl(b)
	*a = Orientation(*nt)
	return br, err
}

func (a Orientation) Readable() bool   { return true }
func (a Orientation) Writable() bool   { return false }
func (a Orientation) Reportable() bool { return false }
func (a Orientation) SceneIndex() int  { return -1 }

func (a Orientation) String() string {
	return zcl.Sprintf("0x%X", zcl.Zu48(a))
}
