package general

import "hemtjan.st/zcl"

// LevelControl
const LevelControlID zcl.ClusterID = 8

var LevelControlCluster = zcl.Cluster{
	Name: "Level Control",
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
		OnOffTransistionTimeAttr: func() zcl.Attr { return new(OnOffTransistionTime) },
		OnLevelAttr:              func() zcl.Attr { return new(OnLevel) },
		OnTransitionTimeAttr:     func() zcl.Attr { return new(OnTransitionTime) },
		OffTransitionTimeAttr:    func() zcl.Attr { return new(OffTransitionTime) },
		DefaultMoveRateAttr:      func() zcl.Attr { return new(DefaultMoveRate) },
		LevelControlOptionsAttr:  func() zcl.Attr { return new(LevelControlOptions) },
		PowerOnLevelAttr:         func() zcl.Attr { return new(PowerOnLevel) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr: []zcl.AttrID{
		CurrentLevelAttr,
	},
}

type MoveToLevel struct {
	Level          Level
	TransitionTime TransitionTime
}

type MoveToLevelHandler interface {
	HandleMoveToLevel(frame Frame, cmd *MoveToLevel) error
}

// MoveToLevelCommand is the Command ID of MoveToLevel
const MoveToLevelCommand CommandID = 0x0000

// Values returns all values of MoveToLevel
func (v *MoveToLevel) Values() []zcl.Val {
	return []zcl.Val{
		&v.Level,
		&v.TransitionTime,
	}
}

// Arguments returns all values of MoveToLevel
func (v *MoveToLevel) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Level", Argument: &v.Level},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (MoveToLevel) Name() string { return `Move to Level` }

// Description of the command
func (MoveToLevel) Description() string { return `` }

// ID of the command
func (MoveToLevel) ID() CommandID { return MoveToLevelCommand }

// Required
func (MoveToLevel) Required() bool { return true }

// Cluster ID of the command
func (MoveToLevel) Cluster() zcl.ClusterID { return LevelControlID }

// Direction of the command
func (MoveToLevel) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveToLevel) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveToLevel) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

func (v *MoveToLevel) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h MoveToLevelHandler
	if h, found = handler.(MoveToLevelHandler); found {
		err = h.HandleMoveToLevel(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of MoveToLevel
func (v MoveToLevel) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Level.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MoveToLevel struct
func (v *MoveToLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Level).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveToLevel) String() string {
	return zcl.Sprintf(
		"MoveToLevel{"+zcl.StrJoin([]string{
			"Level(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.Level,
		v.TransitionTime,
	)
}

type Move struct {
	LevelDirection LevelDirection
	Rate           Rate
}

type MoveHandler interface {
	HandleMove(frame Frame, cmd *Move) error
}

// MoveCommand is the Command ID of Move
const MoveCommand CommandID = 0x0001

// Values returns all values of Move
func (v *Move) Values() []zcl.Val {
	return []zcl.Val{
		&v.LevelDirection,
		&v.Rate,
	}
}

// Arguments returns all values of Move
func (v *Move) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "LevelDirection", Argument: &v.LevelDirection},
		{Name: "Rate", Argument: &v.Rate},
	}
}

// Name of the command
func (Move) Name() string { return `Move` }

// Description of the command
func (Move) Description() string { return `` }

// ID of the command
func (Move) ID() CommandID { return MoveCommand }

// Required
func (Move) Required() bool { return true }

// Cluster ID of the command
func (Move) Cluster() zcl.ClusterID { return LevelControlID }

// Direction of the command
func (Move) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (Move) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (Move) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

func (v *Move) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h MoveHandler
	if h, found = handler.(MoveHandler); found {
		err = h.HandleMove(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of Move
func (v Move) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.LevelDirection.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Rate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the Move struct
func (v *Move) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.LevelDirection).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v Move) String() string {
	return zcl.Sprintf(
		"Move{"+zcl.StrJoin([]string{
			"LevelDirection(%v)",
			"Rate(%v)",
		}, " ")+"}",
		v.LevelDirection,
		v.Rate,
	)
}

type Step struct {
	LevelDirection LevelDirection
	StepSize       StepSize
	TransitionTime TransitionTime
}

type StepHandler interface {
	HandleStep(frame Frame, cmd *Step) error
}

// StepCommand is the Command ID of Step
const StepCommand CommandID = 0x0002

// Values returns all values of Step
func (v *Step) Values() []zcl.Val {
	return []zcl.Val{
		&v.LevelDirection,
		&v.StepSize,
		&v.TransitionTime,
	}
}

// Arguments returns all values of Step
func (v *Step) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "LevelDirection", Argument: &v.LevelDirection},
		{Name: "StepSize", Argument: &v.StepSize},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (Step) Name() string { return `Step` }

// Description of the command
func (Step) Description() string { return `` }

// ID of the command
func (Step) ID() CommandID { return StepCommand }

// Required
func (Step) Required() bool { return true }

// Cluster ID of the command
func (Step) Cluster() zcl.ClusterID { return LevelControlID }

// Direction of the command
func (Step) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (Step) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (Step) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

func (v *Step) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h StepHandler
	if h, found = handler.(StepHandler); found {
		err = h.HandleStep(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of Step
func (v Step) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.LevelDirection.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StepSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the Step struct
func (v *Step) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.LevelDirection).UnmarshalZcl(b); err != nil {
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

// String returns a log-friendly string representation of the struct
func (v Step) String() string {
	return zcl.Sprintf(
		"Step{"+zcl.StrJoin([]string{
			"LevelDirection(%v)",
			"StepSize(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.LevelDirection,
		v.StepSize,
		v.TransitionTime,
	)
}

type Stop struct {
}

type StopHandler interface {
	HandleStop(frame Frame, cmd *Stop) error
}

// StopCommand is the Command ID of Stop
const StopCommand CommandID = 0x0003

// Values returns all values of Stop
func (v *Stop) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of Stop
func (v *Stop) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (Stop) Name() string { return `Stop` }

// Description of the command
func (Stop) Description() string { return `` }

// ID of the command
func (Stop) ID() CommandID { return StopCommand }

// Required
func (Stop) Required() bool { return true }

// Cluster ID of the command
func (Stop) Cluster() zcl.ClusterID { return LevelControlID }

// Direction of the command
func (Stop) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (Stop) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (Stop) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

func (v *Stop) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h StopHandler
	if h, found = handler.(StopHandler); found {
		err = h.HandleStop(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of Stop
func (v Stop) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the Stop struct
func (v *Stop) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v Stop) String() string {
	return zcl.Sprintf(
		"Stop{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type MoveToLevelWithOnOff struct {
	Level          Level
	TransitionTime TransitionTime
}

type MoveToLevelWithOnOffHandler interface {
	HandleMoveToLevelWithOnOff(frame Frame, cmd *MoveToLevelWithOnOff) error
}

// MoveToLevelWithOnOffCommand is the Command ID of MoveToLevelWithOnOff
const MoveToLevelWithOnOffCommand CommandID = 0x0004

// Values returns all values of MoveToLevelWithOnOff
func (v *MoveToLevelWithOnOff) Values() []zcl.Val {
	return []zcl.Val{
		&v.Level,
		&v.TransitionTime,
	}
}

// Arguments returns all values of MoveToLevelWithOnOff
func (v *MoveToLevelWithOnOff) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Level", Argument: &v.Level},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (MoveToLevelWithOnOff) Name() string { return `Move to Level (with On/Off)` }

// Description of the command
func (MoveToLevelWithOnOff) Description() string { return `` }

// ID of the command
func (MoveToLevelWithOnOff) ID() CommandID { return MoveToLevelWithOnOffCommand }

// Required
func (MoveToLevelWithOnOff) Required() bool { return true }

// Cluster ID of the command
func (MoveToLevelWithOnOff) Cluster() zcl.ClusterID { return LevelControlID }

// Direction of the command
func (MoveToLevelWithOnOff) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveToLevelWithOnOff) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveToLevelWithOnOff) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

func (v *MoveToLevelWithOnOff) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h MoveToLevelWithOnOffHandler
	if h, found = handler.(MoveToLevelWithOnOffHandler); found {
		err = h.HandleMoveToLevelWithOnOff(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of MoveToLevelWithOnOff
func (v MoveToLevelWithOnOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Level.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MoveToLevelWithOnOff struct
func (v *MoveToLevelWithOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Level).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveToLevelWithOnOff) String() string {
	return zcl.Sprintf(
		"MoveToLevelWithOnOff{"+zcl.StrJoin([]string{
			"Level(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.Level,
		v.TransitionTime,
	)
}

type MoveWithOnOff struct {
	LevelDirection LevelDirection
	Rate           Rate
}

type MoveWithOnOffHandler interface {
	HandleMoveWithOnOff(frame Frame, cmd *MoveWithOnOff) error
}

// MoveWithOnOffCommand is the Command ID of MoveWithOnOff
const MoveWithOnOffCommand CommandID = 0x0005

// Values returns all values of MoveWithOnOff
func (v *MoveWithOnOff) Values() []zcl.Val {
	return []zcl.Val{
		&v.LevelDirection,
		&v.Rate,
	}
}

// Arguments returns all values of MoveWithOnOff
func (v *MoveWithOnOff) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "LevelDirection", Argument: &v.LevelDirection},
		{Name: "Rate", Argument: &v.Rate},
	}
}

// Name of the command
func (MoveWithOnOff) Name() string { return `Move (with On/Off)` }

// Description of the command
func (MoveWithOnOff) Description() string { return `` }

// ID of the command
func (MoveWithOnOff) ID() CommandID { return MoveWithOnOffCommand }

// Required
func (MoveWithOnOff) Required() bool { return true }

// Cluster ID of the command
func (MoveWithOnOff) Cluster() zcl.ClusterID { return LevelControlID }

// Direction of the command
func (MoveWithOnOff) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveWithOnOff) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveWithOnOff) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

func (v *MoveWithOnOff) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h MoveWithOnOffHandler
	if h, found = handler.(MoveWithOnOffHandler); found {
		err = h.HandleMoveWithOnOff(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of MoveWithOnOff
func (v MoveWithOnOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.LevelDirection.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Rate.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MoveWithOnOff struct
func (v *MoveWithOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.LevelDirection).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveWithOnOff) String() string {
	return zcl.Sprintf(
		"MoveWithOnOff{"+zcl.StrJoin([]string{
			"LevelDirection(%v)",
			"Rate(%v)",
		}, " ")+"}",
		v.LevelDirection,
		v.Rate,
	)
}

type StepWithOnOff struct {
	LevelDirection LevelDirection
	StepSize       StepSize
	TransitionTime TransitionTime
}

type StepWithOnOffHandler interface {
	HandleStepWithOnOff(frame Frame, cmd *StepWithOnOff) error
}

// StepWithOnOffCommand is the Command ID of StepWithOnOff
const StepWithOnOffCommand CommandID = 0x0006

// Values returns all values of StepWithOnOff
func (v *StepWithOnOff) Values() []zcl.Val {
	return []zcl.Val{
		&v.LevelDirection,
		&v.StepSize,
		&v.TransitionTime,
	}
}

// Arguments returns all values of StepWithOnOff
func (v *StepWithOnOff) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "LevelDirection", Argument: &v.LevelDirection},
		{Name: "StepSize", Argument: &v.StepSize},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (StepWithOnOff) Name() string { return `Step (with On/Off)` }

// Description of the command
func (StepWithOnOff) Description() string { return `` }

// ID of the command
func (StepWithOnOff) ID() CommandID { return StepWithOnOffCommand }

// Required
func (StepWithOnOff) Required() bool { return true }

// Cluster ID of the command
func (StepWithOnOff) Cluster() zcl.ClusterID { return LevelControlID }

// Direction of the command
func (StepWithOnOff) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (StepWithOnOff) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (StepWithOnOff) MarshalJSON() ([]byte, error) { return []byte("6"), nil }

func (v *StepWithOnOff) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h StepWithOnOffHandler
	if h, found = handler.(StepWithOnOffHandler); found {
		err = h.HandleStepWithOnOff(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of StepWithOnOff
func (v StepWithOnOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.LevelDirection.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StepSize.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the StepWithOnOff struct
func (v *StepWithOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.LevelDirection).UnmarshalZcl(b); err != nil {
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

// String returns a log-friendly string representation of the struct
func (v StepWithOnOff) String() string {
	return zcl.Sprintf(
		"StepWithOnOff{"+zcl.StrJoin([]string{
			"LevelDirection(%v)",
			"StepSize(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.LevelDirection,
		v.StepSize,
		v.TransitionTime,
	)
}

type StopWithOnOff struct {
}

type StopWithOnOffHandler interface {
	HandleStopWithOnOff(frame Frame, cmd *StopWithOnOff) error
}

// StopWithOnOffCommand is the Command ID of StopWithOnOff
const StopWithOnOffCommand CommandID = 0x0007

// Values returns all values of StopWithOnOff
func (v *StopWithOnOff) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of StopWithOnOff
func (v *StopWithOnOff) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (StopWithOnOff) Name() string { return `Stop (with On/Off)` }

// Description of the command
func (StopWithOnOff) Description() string { return `` }

// ID of the command
func (StopWithOnOff) ID() CommandID { return StopWithOnOffCommand }

// Required
func (StopWithOnOff) Required() bool { return true }

// Cluster ID of the command
func (StopWithOnOff) Cluster() zcl.ClusterID { return LevelControlID }

// Direction of the command
func (StopWithOnOff) Direction() zcl.Direction { return zcl.ClientToServer }

// MnfCode returns the manufacturer code (if any) of the command
func (StopWithOnOff) MnfCode() uint16 { return 0 }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (StopWithOnOff) MarshalJSON() ([]byte, error) { return []byte("7"), nil }

func (v *StopWithOnOff) Handle(frame Frame, handler interface{}) (rsp zcl.General, found bool, err error) {
	var h StopWithOnOffHandler
	if h, found = handler.(StopWithOnOffHandler); found {
		err = h.HandleStopWithOnOff(frame, v)
	}
	return
}

// MarshalZcl returns the wire format representation of StopWithOnOff
func (v StopWithOnOff) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the StopWithOnOff struct
func (v *StopWithOnOff) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v StopWithOnOff) String() string {
	return zcl.Sprintf(
		"StopWithOnOff{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}
