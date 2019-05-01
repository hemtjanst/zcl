// This cluster provides an interface for controlling a characteristic of a device that can be set to a level, for example the brightness of a light, the degree of closure of a door, or the power output of a heater.
package general

import (
	"neotor.se/zcl/cluster/zcl"
)

// LevelControl
// This cluster provides an interface for controlling a characteristic of a device that can be set to a level, for example the brightness of a light, the degree of closure of a door, or the power output of a heater.

func NewLevelControlServer(profile zcl.ProfileID) *LevelControlServer {
	return &LevelControlServer{p: profile}
}
func NewLevelControlClient(profile zcl.ProfileID) *LevelControlClient {
	return &LevelControlClient{p: profile}
}

const LevelControlCluster zcl.ClusterID = 8

type LevelControlServer struct {
	p zcl.ProfileID

	CurrentLevel         *CurrentLevel
	RemainingTime        *RemainingTime
	Unknown              *Unknown
	OnoffTransistionTime *OnoffTransistionTime
	OnLevel              *OnLevel
	OnTransitionTime     *OnTransitionTime
	OffTransitionTime    *OffTransitionTime
	DefaultMoveRate      *DefaultMoveRate
	PoweronLevel         *PoweronLevel
}

func (s *LevelControlServer) MoveToLevel() *MoveToLevel { return new(MoveToLevel) }
func (s *LevelControlServer) Move() *Move               { return new(Move) }
func (s *LevelControlServer) Step() *Step               { return new(Step) }
func (s *LevelControlServer) Stop() *Stop               { return new(Stop) }
func (s *LevelControlServer) MoveToLevelWithOnOff() *MoveToLevelWithOnOff {
	return new(MoveToLevelWithOnOff)
}
func (s *LevelControlServer) MoveWithOnOff() *MoveWithOnOff { return new(MoveWithOnOff) }
func (s *LevelControlServer) StepWithOnOff() *StepWithOnOff { return new(StepWithOnOff) }
func (s *LevelControlServer) StopWithOnOff() *StopWithOnOff { return new(StopWithOnOff) }

type LevelControlClient struct {
	p zcl.ProfileID
}

/*
var LevelControlServer = map[zcl.CommandID]func() zcl.Command{
    MoveToLevelID: func() zcl.Command { return new(MoveToLevel) },
    MoveID: func() zcl.Command { return new(Move) },
    StepID: func() zcl.Command { return new(Step) },
    StopID: func() zcl.Command { return new(Stop) },
    MoveToLevelWithOnOffID: func() zcl.Command { return new(MoveToLevelWithOnOff) },
    MoveWithOnOffID: func() zcl.Command { return new(MoveWithOnOff) },
    StepWithOnOffID: func() zcl.Command { return new(StepWithOnOff) },
    StopWithOnOffID: func() zcl.Command { return new(StopWithOnOff) },
}

var LevelControlClient = map[zcl.CommandID]func() zcl.Command{
}
*/

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
	return LevelControlCluster
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
	return LevelControlCluster
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
	return LevelControlCluster
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
	return LevelControlCluster
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
	return LevelControlCluster
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
	return LevelControlCluster
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
	return LevelControlCluster
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
	return LevelControlCluster
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

type CurrentLevel zcl.Zu8

func (a CurrentLevel) ID() zcl.AttrID         { return 0 }
func (a CurrentLevel) Cluster() zcl.ClusterID { return LevelControlCluster }
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

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type RemainingTime zcl.Zu16

func (a RemainingTime) ID() zcl.AttrID         { return 1 }
func (a RemainingTime) Cluster() zcl.ClusterID { return LevelControlCluster }
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

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type Unknown zcl.Zbmp8

func (a Unknown) ID() zcl.AttrID         { return 15 }
func (a Unknown) Cluster() zcl.ClusterID { return LevelControlCluster }
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

type OnoffTransistionTime zcl.Zu16

func (a OnoffTransistionTime) ID() zcl.AttrID                { return 16 }
func (a OnoffTransistionTime) Cluster() zcl.ClusterID        { return LevelControlCluster }
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

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type OnLevel zcl.Zu8

func (a OnLevel) ID() zcl.AttrID         { return 17 }
func (a OnLevel) Cluster() zcl.ClusterID { return LevelControlCluster }
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

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type OnTransitionTime zcl.Zu16

func (a OnTransitionTime) ID() zcl.AttrID            { return 18 }
func (a OnTransitionTime) Cluster() zcl.ClusterID    { return LevelControlCluster }
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

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type OffTransitionTime zcl.Zu16

func (a OffTransitionTime) ID() zcl.AttrID             { return 19 }
func (a OffTransitionTime) Cluster() zcl.ClusterID     { return LevelControlCluster }
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

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type DefaultMoveRate zcl.Zu8

func (a DefaultMoveRate) ID() zcl.AttrID           { return 20 }
func (a DefaultMoveRate) Cluster() zcl.ClusterID   { return LevelControlCluster }
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

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

type PoweronLevel zcl.Zu8

func (a PoweronLevel) ID() zcl.AttrID         { return 16384 }
func (a PoweronLevel) Cluster() zcl.ClusterID { return LevelControlCluster }
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

	return zcl.Sprintf("%s", zcl.Zu8(a))
}
