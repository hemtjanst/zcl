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

// MoveToLevelCommand is the Command ID of MoveToLevel
const MoveToLevelCommand CommandID = 0x0000

// Values returns all values of MoveToLevel
func (v *MoveToLevel) Values() []zcl.Val {
	return []zcl.Val{
		&v.Level,
		&v.TransitionTime,
	}
}

// Name of the command (needed to fulfill interface)
func (MoveToLevel) Name() string { return "Move to Level" }

// ID of the command (needed to fulfill interface)
func (MoveToLevel) ID() CommandID { return MoveToLevelCommand }

// Cluster ID of the command (needed to fulfill interface)
func (MoveToLevel) Cluster() zcl.ClusterID { return LevelControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveToLevel) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of MoveToLevel
func (v MoveToLevel) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
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

// MoveCommand is the Command ID of Move
const MoveCommand CommandID = 0x0001

// Values returns all values of Move
func (v *Move) Values() []zcl.Val {
	return []zcl.Val{
		&v.LevelDirection,
		&v.Rate,
	}
}

// Name of the command (needed to fulfill interface)
func (Move) Name() string { return "Move" }

// ID of the command (needed to fulfill interface)
func (Move) ID() CommandID { return MoveCommand }

// Cluster ID of the command (needed to fulfill interface)
func (Move) Cluster() zcl.ClusterID { return LevelControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (Move) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of Move
func (v Move) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
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

// Name of the command (needed to fulfill interface)
func (Step) Name() string { return "Step" }

// ID of the command (needed to fulfill interface)
func (Step) ID() CommandID { return StepCommand }

// Cluster ID of the command (needed to fulfill interface)
func (Step) Cluster() zcl.ClusterID { return LevelControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (Step) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of Step
func (v Step) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
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

// StopCommand is the Command ID of Stop
const StopCommand CommandID = 0x0003

// Values returns all values of Stop
func (v *Stop) Values() []zcl.Val {
	return []zcl.Val{}
}

// Name of the command (needed to fulfill interface)
func (Stop) Name() string { return "Stop" }

// ID of the command (needed to fulfill interface)
func (Stop) ID() CommandID { return StopCommand }

// Cluster ID of the command (needed to fulfill interface)
func (Stop) Cluster() zcl.ClusterID { return LevelControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (Stop) MnfCode() []byte { return []byte{} }

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

// MoveToLevelWithOnOffCommand is the Command ID of MoveToLevelWithOnOff
const MoveToLevelWithOnOffCommand CommandID = 0x0004

// Values returns all values of MoveToLevelWithOnOff
func (v *MoveToLevelWithOnOff) Values() []zcl.Val {
	return []zcl.Val{
		&v.Level,
		&v.TransitionTime,
	}
}

// Name of the command (needed to fulfill interface)
func (MoveToLevelWithOnOff) Name() string { return "Move to Level (with On/Off)" }

// ID of the command (needed to fulfill interface)
func (MoveToLevelWithOnOff) ID() CommandID { return MoveToLevelWithOnOffCommand }

// Cluster ID of the command (needed to fulfill interface)
func (MoveToLevelWithOnOff) Cluster() zcl.ClusterID { return LevelControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveToLevelWithOnOff) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of MoveToLevelWithOnOff
func (v MoveToLevelWithOnOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
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

// MoveWithOnOffCommand is the Command ID of MoveWithOnOff
const MoveWithOnOffCommand CommandID = 0x0005

// Values returns all values of MoveWithOnOff
func (v *MoveWithOnOff) Values() []zcl.Val {
	return []zcl.Val{
		&v.LevelDirection,
		&v.Rate,
	}
}

// Name of the command (needed to fulfill interface)
func (MoveWithOnOff) Name() string { return "Move (with On/Off)" }

// ID of the command (needed to fulfill interface)
func (MoveWithOnOff) ID() CommandID { return MoveWithOnOffCommand }

// Cluster ID of the command (needed to fulfill interface)
func (MoveWithOnOff) Cluster() zcl.ClusterID { return LevelControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveWithOnOff) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of MoveWithOnOff
func (v MoveWithOnOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
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

// Name of the command (needed to fulfill interface)
func (StepWithOnOff) Name() string { return "Step (with On/Off)" }

// ID of the command (needed to fulfill interface)
func (StepWithOnOff) ID() CommandID { return StepWithOnOffCommand }

// Cluster ID of the command (needed to fulfill interface)
func (StepWithOnOff) Cluster() zcl.ClusterID { return LevelControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (StepWithOnOff) MnfCode() []byte { return []byte{} }

// MarshalZcl returns the wire format representation of StepWithOnOff
func (v StepWithOnOff) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
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

// StopWithOnOffCommand is the Command ID of StopWithOnOff
const StopWithOnOffCommand CommandID = 0x0007

// Values returns all values of StopWithOnOff
func (v *StopWithOnOff) Values() []zcl.Val {
	return []zcl.Val{}
}

// Name of the command (needed to fulfill interface)
func (StopWithOnOff) Name() string { return "Stop (with On/Off)" }

// ID of the command (needed to fulfill interface)
func (StopWithOnOff) ID() CommandID { return StopWithOnOffCommand }

// Cluster ID of the command (needed to fulfill interface)
func (StopWithOnOff) Cluster() zcl.ClusterID { return LevelControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (StopWithOnOff) MnfCode() []byte { return []byte{} }

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
