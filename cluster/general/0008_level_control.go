// This cluster provides an interface for controlling a characteristic of a device that can be set to a level, for example the brightness of a light, the degree of closure of a door, or the power output of a heater.
package general

import (
	"neotor.se/zcl"
)

// LevelControl
const LevelControlID zcl.ClusterID = 8

var LevelControlCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		MoveToLevelCommand:          func() zcl.Command { return new(MoveToLevel) },
		MoveCommand:                 func() zcl.Command { return new(Move) },
		StepCommand:                 func() zcl.Command { return new(Step) },
		StopCommand:                 func() zcl.Command { return new(Stop) },
		MoveToLevelWithOnOffCommand: func() zcl.Command { return new(MoveToLevelWithOnOff) },
		MoveWithOnOffCommand:        func() zcl.Command { return new(MoveWithOnOff) },
		StepWithOnOffCommand:        func() zcl.Command { return new(StepWithOnOff) },
		StopWithOnOffCommand:        func() zcl.Command { return new(StopWithOnOff) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		CurrentLevelAttr:         func() zcl.Attr { return new(CurrentLevel) },
		RemainingTimeAttr:        func() zcl.Attr { return new(RemainingTime) },
		UnknownAttr:              func() zcl.Attr { return new(Unknown) },
		OnoffTransistionTimeAttr: func() zcl.Attr { return new(OnoffTransistionTime) },
		OnLevelAttr:              func() zcl.Attr { return new(OnLevel) },
		OnTransitionTimeAttr:     func() zcl.Attr { return new(OnTransitionTime) },
		OffTransitionTimeAttr:    func() zcl.Attr { return new(OffTransitionTime) },
		DefaultMoveRateAttr:      func() zcl.Attr { return new(DefaultMoveRate) },
		PoweronLevelAttr:         func() zcl.Attr { return new(PoweronLevel) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr: []zcl.AttrID{
		CurrentLevelAttr,
	},
}

type MoveToLevel struct {
	Level           zcl.Zu8
	TransistionTime zcl.Zu16
}

const MoveToLevelCommand zcl.CommandID = 0

func (v *MoveToLevel) Values() []zcl.Val {
	return []zcl.Val{
		&v.Level,
		&v.TransistionTime,
	}
}

func (v MoveToLevel) ID() zcl.CommandID {
	return MoveToLevelCommand
}

func (v MoveToLevel) Cluster() zcl.ClusterID {
	return LevelControlID
}

func (v MoveToLevel) MnfCode() []byte {
	return []byte{}
}

func (v MoveToLevel) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Level.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransistionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MoveToLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Level).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransistionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type Move struct {
	MoveMode zcl.Zenum8
	Rate     zcl.Zu8
}

const MoveCommand zcl.CommandID = 1

func (v *Move) Values() []zcl.Val {
	return []zcl.Val{
		&v.MoveMode,
		&v.Rate,
	}
}

func (v Move) ID() zcl.CommandID {
	return MoveCommand
}

func (v Move) Cluster() zcl.ClusterID {
	return LevelControlID
}

func (v Move) MnfCode() []byte {
	return []byte{}
}

func (v Move) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.MoveMode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Rate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *Move) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.MoveMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type Step struct {
	StepMode       zcl.Zenum8
	StepSize       zcl.Zu8
	TransitionTime zcl.Zu16
}

const StepCommand zcl.CommandID = 2

func (v *Step) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepMode,
		&v.StepSize,
		&v.TransitionTime,
	}
}

func (v Step) ID() zcl.CommandID {
	return StepCommand
}

func (v Step) Cluster() zcl.ClusterID {
	return LevelControlID
}

func (v Step) MnfCode() []byte {
	return []byte{}
}

func (v Step) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.StepMode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StepSize.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *Step) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.StepMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StepSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type Stop struct {
}

const StopCommand zcl.CommandID = 3

func (v *Stop) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v Stop) ID() zcl.CommandID {
	return StopCommand
}

func (v Stop) Cluster() zcl.ClusterID {
	return LevelControlID
}

func (v Stop) MnfCode() []byte {
	return []byte{}
}

func (v Stop) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *Stop) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

type MoveToLevelWithOnOff struct {
	Level           zcl.Zu8
	TransistionTime zcl.Zu16
}

const MoveToLevelWithOnOffCommand zcl.CommandID = 4

func (v *MoveToLevelWithOnOff) Values() []zcl.Val {
	return []zcl.Val{
		&v.Level,
		&v.TransistionTime,
	}
}

func (v MoveToLevelWithOnOff) ID() zcl.CommandID {
	return MoveToLevelWithOnOffCommand
}

func (v MoveToLevelWithOnOff) Cluster() zcl.ClusterID {
	return LevelControlID
}

func (v MoveToLevelWithOnOff) MnfCode() []byte {
	return []byte{}
}

func (v MoveToLevelWithOnOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Level.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransistionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MoveToLevelWithOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Level).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransistionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type MoveWithOnOff struct {
	MoveMode zcl.Zenum8
	Rate     zcl.Zu8
}

const MoveWithOnOffCommand zcl.CommandID = 5

func (v *MoveWithOnOff) Values() []zcl.Val {
	return []zcl.Val{
		&v.MoveMode,
		&v.Rate,
	}
}

func (v MoveWithOnOff) ID() zcl.CommandID {
	return MoveWithOnOffCommand
}

func (v MoveWithOnOff) Cluster() zcl.ClusterID {
	return LevelControlID
}

func (v MoveWithOnOff) MnfCode() []byte {
	return []byte{}
}

func (v MoveWithOnOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.MoveMode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Rate.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MoveWithOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.MoveMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type StepWithOnOff struct {
	StepMode       zcl.Zenum8
	StepSize       zcl.Zu8
	TransitionTime zcl.Zu16
}

const StepWithOnOffCommand zcl.CommandID = 6

func (v *StepWithOnOff) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepMode,
		&v.StepSize,
		&v.TransitionTime,
	}
}

func (v StepWithOnOff) ID() zcl.CommandID {
	return StepWithOnOffCommand
}

func (v StepWithOnOff) Cluster() zcl.ClusterID {
	return LevelControlID
}

func (v StepWithOnOff) MnfCode() []byte {
	return []byte{}
}

func (v StepWithOnOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.StepMode.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StepSize.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *StepWithOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.StepMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StepSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type StopWithOnOff struct {
}

const StopWithOnOffCommand zcl.CommandID = 7

func (v *StopWithOnOff) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v StopWithOnOff) ID() zcl.CommandID {
	return StopWithOnOffCommand
}

func (v StopWithOnOff) Cluster() zcl.ClusterID {
	return LevelControlID
}

func (v StopWithOnOff) MnfCode() []byte {
	return []byte{}
}

func (v StopWithOnOff) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *StopWithOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// CurrentLevel is an autogenerated attribute in the LevelControl cluster
// The CurrentLevel attribute represents the current level of this device. meaning of 'level' is device dependent.
type CurrentLevel zcl.Zu8

const CurrentLevelAttr zcl.AttrID = 0

func (a CurrentLevel) ID() zcl.AttrID         { return CurrentLevelAttr }
func (a CurrentLevel) Cluster() zcl.ClusterID { return LevelControlID }
func (a *CurrentLevel) Value() *CurrentLevel  { return a }
func (a CurrentLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *CurrentLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentLevel(*nt)
	return br, err
}

func (a CurrentLevel) Readable() bool   { return true }
func (a CurrentLevel) Writable() bool   { return false }
func (a CurrentLevel) Reportable() bool { return true }
func (a CurrentLevel) SceneIndex() int  { return 1 }

func (a CurrentLevel) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

// RemainingTime is an autogenerated attribute in the LevelControl cluster
// The RemainingTime attribute represents the time remaining until the current command is complete - it is specified in 1/10ths of a second.
type RemainingTime zcl.Zu16

const RemainingTimeAttr zcl.AttrID = 1

func (a RemainingTime) ID() zcl.AttrID         { return RemainingTimeAttr }
func (a RemainingTime) Cluster() zcl.ClusterID { return LevelControlID }
func (a *RemainingTime) Value() *RemainingTime { return a }
func (a RemainingTime) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *RemainingTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RemainingTime(*nt)
	return br, err
}

func (a RemainingTime) Readable() bool   { return true }
func (a RemainingTime) Writable() bool   { return false }
func (a RemainingTime) Reportable() bool { return false }
func (a RemainingTime) SceneIndex() int  { return -1 }

func (a RemainingTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

// Unknown is an autogenerated attribute in the LevelControl cluster
// IKEA specific.
type Unknown zcl.Zbmp8

const UnknownAttr zcl.AttrID = 15

func (a Unknown) ID() zcl.AttrID         { return UnknownAttr }
func (a Unknown) Cluster() zcl.ClusterID { return LevelControlID }
func (a *Unknown) Value() *Unknown       { return a }
func (a Unknown) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *Unknown) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = Unknown(*nt)
	return br, err
}

func (a Unknown) Readable() bool   { return true }
func (a Unknown) Writable() bool   { return true }
func (a Unknown) Reportable() bool { return false }
func (a Unknown) SceneIndex() int  { return -1 }

func (a Unknown) String() string {
	return zcl.Sprintf("%s", zcl.Zbmp8(a))
}

// OnoffTransistionTime is an autogenerated attribute in the LevelControl cluster
// The OnOffTransitionTime attribute represents the time taken to move to or from the target level when On of Off commands are received by an On/Off cluster on the same endpoint. It is specified in 1/10ths of a second. The actual time taken should be as close to OnOffTransitionTime as the device is able.
type OnoffTransistionTime zcl.Zu16

const OnoffTransistionTimeAttr zcl.AttrID = 16

func (a OnoffTransistionTime) ID() zcl.AttrID                { return OnoffTransistionTimeAttr }
func (a OnoffTransistionTime) Cluster() zcl.ClusterID        { return LevelControlID }
func (a *OnoffTransistionTime) Value() *OnoffTransistionTime { return a }
func (a OnoffTransistionTime) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *OnoffTransistionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OnoffTransistionTime(*nt)
	return br, err
}

func (a OnoffTransistionTime) Readable() bool   { return true }
func (a OnoffTransistionTime) Writable() bool   { return true }
func (a OnoffTransistionTime) Reportable() bool { return false }
func (a OnoffTransistionTime) SceneIndex() int  { return -1 }

func (a OnoffTransistionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

// OnLevel is an autogenerated attribute in the LevelControl cluster
// The OnLevel attribute determines the value that the CurrentLevel attribute is set to when the OnOff attribute of an On/Off cluster on the same endpoint is set to On. If the OnLevel attribute is not implemented, or is set to 0xff, it has no effect.
type OnLevel zcl.Zu8

const OnLevelAttr zcl.AttrID = 17

func (a OnLevel) ID() zcl.AttrID         { return OnLevelAttr }
func (a OnLevel) Cluster() zcl.ClusterID { return LevelControlID }
func (a *OnLevel) Value() *OnLevel       { return a }
func (a OnLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *OnLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = OnLevel(*nt)
	return br, err
}

func (a OnLevel) Readable() bool   { return true }
func (a OnLevel) Writable() bool   { return true }
func (a OnLevel) Reportable() bool { return false }
func (a OnLevel) SceneIndex() int  { return -1 }

func (a OnLevel) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}

// OnTransitionTime is an autogenerated attribute in the LevelControl cluster
type OnTransitionTime zcl.Zu16

const OnTransitionTimeAttr zcl.AttrID = 18

func (a OnTransitionTime) ID() zcl.AttrID            { return OnTransitionTimeAttr }
func (a OnTransitionTime) Cluster() zcl.ClusterID    { return LevelControlID }
func (a *OnTransitionTime) Value() *OnTransitionTime { return a }
func (a OnTransitionTime) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *OnTransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OnTransitionTime(*nt)
	return br, err
}

func (a OnTransitionTime) Readable() bool   { return true }
func (a OnTransitionTime) Writable() bool   { return true }
func (a OnTransitionTime) Reportable() bool { return false }
func (a OnTransitionTime) SceneIndex() int  { return -1 }

func (a OnTransitionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

// OffTransitionTime is an autogenerated attribute in the LevelControl cluster
type OffTransitionTime zcl.Zu16

const OffTransitionTimeAttr zcl.AttrID = 19

func (a OffTransitionTime) ID() zcl.AttrID             { return OffTransitionTimeAttr }
func (a OffTransitionTime) Cluster() zcl.ClusterID     { return LevelControlID }
func (a *OffTransitionTime) Value() *OffTransitionTime { return a }
func (a OffTransitionTime) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *OffTransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = OffTransitionTime(*nt)
	return br, err
}

func (a OffTransitionTime) Readable() bool   { return true }
func (a OffTransitionTime) Writable() bool   { return true }
func (a OffTransitionTime) Reportable() bool { return false }
func (a OffTransitionTime) SceneIndex() int  { return -1 }

func (a OffTransitionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

// DefaultMoveRate is an autogenerated attribute in the LevelControl cluster
type DefaultMoveRate zcl.Zu8

const DefaultMoveRateAttr zcl.AttrID = 20

func (a DefaultMoveRate) ID() zcl.AttrID           { return DefaultMoveRateAttr }
func (a DefaultMoveRate) Cluster() zcl.ClusterID   { return LevelControlID }
func (a *DefaultMoveRate) Value() *DefaultMoveRate { return a }
func (a DefaultMoveRate) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *DefaultMoveRate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = DefaultMoveRate(*nt)
	return br, err
}

func (a DefaultMoveRate) Readable() bool   { return true }
func (a DefaultMoveRate) Writable() bool   { return true }
func (a DefaultMoveRate) Reportable() bool { return false }
func (a DefaultMoveRate) SceneIndex() int  { return -1 }

func (a DefaultMoveRate) String() string {
	return zcl.PercentPerSecond.Format(float64(a) / 2.54)
}

// PoweronLevel is an autogenerated attribute in the LevelControl cluster
type PoweronLevel zcl.Zu8

const PoweronLevelAttr zcl.AttrID = 16384

func (a PoweronLevel) ID() zcl.AttrID         { return PoweronLevelAttr }
func (a PoweronLevel) Cluster() zcl.ClusterID { return LevelControlID }
func (a *PoweronLevel) Value() *PoweronLevel  { return a }
func (a PoweronLevel) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *PoweronLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PoweronLevel(*nt)
	return br, err
}

func (a PoweronLevel) Readable() bool   { return true }
func (a PoweronLevel) Writable() bool   { return true }
func (a PoweronLevel) Reportable() bool { return false }
func (a PoweronLevel) SceneIndex() int  { return -1 }

func (a PoweronLevel) String() string {
	return zcl.Percent.Format(float64(a) / 2.54)
}
