package closures

import "hemtjan.st/zcl"

// WindowCovering
const WindowCoveringID zcl.ClusterID = 258

var WindowCoveringCluster = zcl.Cluster{
	Name: "Window Covering",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		UpOpenCommand:             func() zcl.Command { return new(UpOpen) },
		DownCloseCommand:          func() zcl.Command { return new(DownClose) },
		StopCommand:               func() zcl.Command { return new(Stop) },
		GoToLiftValueCommand:      func() zcl.Command { return new(GoToLiftValue) },
		GoToLiftPercentageCommand: func() zcl.Command { return new(GoToLiftPercentage) },
		GoToTiltValueCommand:      func() zcl.Command { return new(GoToTiltValue) },
		GoToTiltPercentageCommand: func() zcl.Command { return new(GoToTiltPercentage) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		WindowCoveringTypeAttr:            func() zcl.Attr { return new(WindowCoveringType) },
		LiftPhysicalClosedLimitAttr:       func() zcl.Attr { return new(LiftPhysicalClosedLimit) },
		TiltPhysicalClosedLimitAttr:       func() zcl.Attr { return new(TiltPhysicalClosedLimit) },
		LiftCurrentPositionAttr:           func() zcl.Attr { return new(LiftCurrentPosition) },
		TiltCurrentPositionAttr:           func() zcl.Attr { return new(TiltCurrentPosition) },
		LiftNumberOfActuationsAttr:        func() zcl.Attr { return new(LiftNumberOfActuations) },
		TiltNumberOfActuationsAttr:        func() zcl.Attr { return new(TiltNumberOfActuations) },
		ConfigStatusAttr:                  func() zcl.Attr { return new(ConfigStatus) },
		LiftCurrentPositionPercentageAttr: func() zcl.Attr { return new(LiftCurrentPositionPercentage) },
		TiltCurrentPositionPercentageAttr: func() zcl.Attr { return new(TiltCurrentPositionPercentage) },
		LiftInstalledOpenLimitAttr:        func() zcl.Attr { return new(LiftInstalledOpenLimit) },
		LiftInstalledClosedLimitAttr:      func() zcl.Attr { return new(LiftInstalledClosedLimit) },
		TiltAInstalledOpenLimitAttr:       func() zcl.Attr { return new(TiltAInstalledOpenLimit) },
		TiltBInstalledOpenLimitAttr:       func() zcl.Attr { return new(TiltBInstalledOpenLimit) },
		LiftVelocityAttr:                  func() zcl.Attr { return new(LiftVelocity) },
		LiftAccelerationTimeAttr:          func() zcl.Attr { return new(LiftAccelerationTime) },
		LiftDecelerationTimeAttr:          func() zcl.Attr { return new(LiftDecelerationTime) },
		WindowCoveringModeAttr:            func() zcl.Attr { return new(WindowCoveringMode) },
		LiftIntermediateSetpointsAttr:     func() zcl.Attr { return new(LiftIntermediateSetpoints) },
		TiltIntermediateSetpointsAttr:     func() zcl.Attr { return new(TiltIntermediateSetpoints) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

type UpOpen struct {
}

// UpOpenCommand is the Command ID of UpOpen
const UpOpenCommand CommandID = 0x0000

// Values returns all values of UpOpen
func (v *UpOpen) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of UpOpen
func (v *UpOpen) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (UpOpen) Name() string { return "Up / Open" }

// ID of the command
func (UpOpen) ID() CommandID { return UpOpenCommand }

// Required
func (UpOpen) Required() bool { return true }

// Cluster ID of the command
func (UpOpen) Cluster() zcl.ClusterID { return WindowCoveringID }

// MnfCode returns the manufacturer code (if any) of the command
func (UpOpen) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (UpOpen) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

// MarshalZcl returns the wire format representation of UpOpen
func (v UpOpen) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the UpOpen struct
func (v *UpOpen) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v UpOpen) String() string {
	return zcl.Sprintf(
		"UpOpen{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type DownClose struct {
}

// DownCloseCommand is the Command ID of DownClose
const DownCloseCommand CommandID = 0x0001

// Values returns all values of DownClose
func (v *DownClose) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of DownClose
func (v *DownClose) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (DownClose) Name() string { return "Down / Close" }

// ID of the command
func (DownClose) ID() CommandID { return DownCloseCommand }

// Required
func (DownClose) Required() bool { return true }

// Cluster ID of the command
func (DownClose) Cluster() zcl.ClusterID { return WindowCoveringID }

// MnfCode returns the manufacturer code (if any) of the command
func (DownClose) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (DownClose) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

// MarshalZcl returns the wire format representation of DownClose
func (v DownClose) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the DownClose struct
func (v *DownClose) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v DownClose) String() string {
	return zcl.Sprintf(
		"DownClose{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type Stop struct {
}

// StopCommand is the Command ID of Stop
const StopCommand CommandID = 0x0002

// Values returns all values of Stop
func (v *Stop) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of Stop
func (v *Stop) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (Stop) Name() string { return "Stop" }

// ID of the command
func (Stop) ID() CommandID { return StopCommand }

// Required
func (Stop) Required() bool { return true }

// Cluster ID of the command
func (Stop) Cluster() zcl.ClusterID { return WindowCoveringID }

// MnfCode returns the manufacturer code (if any) of the command
func (Stop) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (Stop) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

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

type GoToLiftValue struct {
	Position Position
}

// GoToLiftValueCommand is the Command ID of GoToLiftValue
const GoToLiftValueCommand CommandID = 0x0004

// Values returns all values of GoToLiftValue
func (v *GoToLiftValue) Values() []zcl.Val {
	return []zcl.Val{
		&v.Position,
	}
}

// Arguments returns all values of GoToLiftValue
func (v *GoToLiftValue) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Position", Argument: &v.Position},
	}
}

// Name of the command
func (GoToLiftValue) Name() string { return "Go To Lift Value" }

// ID of the command
func (GoToLiftValue) ID() CommandID { return GoToLiftValueCommand }

// Required
func (GoToLiftValue) Required() bool { return false }

// Cluster ID of the command
func (GoToLiftValue) Cluster() zcl.ClusterID { return WindowCoveringID }

// MnfCode returns the manufacturer code (if any) of the command
func (GoToLiftValue) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GoToLiftValue) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

// MarshalZcl returns the wire format representation of GoToLiftValue
func (v GoToLiftValue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Position.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GoToLiftValue struct
func (v *GoToLiftValue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Position).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GoToLiftValue) String() string {
	return zcl.Sprintf(
		"GoToLiftValue{"+zcl.StrJoin([]string{
			"Position(%v)",
		}, " ")+"}",
		v.Position,
	)
}

type GoToLiftPercentage struct {
	Percentage Percentage
}

// GoToLiftPercentageCommand is the Command ID of GoToLiftPercentage
const GoToLiftPercentageCommand CommandID = 0x0005

// Values returns all values of GoToLiftPercentage
func (v *GoToLiftPercentage) Values() []zcl.Val {
	return []zcl.Val{
		&v.Percentage,
	}
}

// Arguments returns all values of GoToLiftPercentage
func (v *GoToLiftPercentage) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Percentage", Argument: &v.Percentage},
	}
}

// Name of the command
func (GoToLiftPercentage) Name() string { return "Go to Lift Percentage" }

// ID of the command
func (GoToLiftPercentage) ID() CommandID { return GoToLiftPercentageCommand }

// Required
func (GoToLiftPercentage) Required() bool { return false }

// Cluster ID of the command
func (GoToLiftPercentage) Cluster() zcl.ClusterID { return WindowCoveringID }

// MnfCode returns the manufacturer code (if any) of the command
func (GoToLiftPercentage) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GoToLiftPercentage) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

// MarshalZcl returns the wire format representation of GoToLiftPercentage
func (v GoToLiftPercentage) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Percentage.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GoToLiftPercentage struct
func (v *GoToLiftPercentage) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Percentage).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GoToLiftPercentage) String() string {
	return zcl.Sprintf(
		"GoToLiftPercentage{"+zcl.StrJoin([]string{
			"Percentage(%v)",
		}, " ")+"}",
		v.Percentage,
	)
}

type GoToTiltValue struct {
	Position Position
}

// GoToTiltValueCommand is the Command ID of GoToTiltValue
const GoToTiltValueCommand CommandID = 0x0007

// Values returns all values of GoToTiltValue
func (v *GoToTiltValue) Values() []zcl.Val {
	return []zcl.Val{
		&v.Position,
	}
}

// Arguments returns all values of GoToTiltValue
func (v *GoToTiltValue) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Position", Argument: &v.Position},
	}
}

// Name of the command
func (GoToTiltValue) Name() string { return "Go to Tilt Value" }

// ID of the command
func (GoToTiltValue) ID() CommandID { return GoToTiltValueCommand }

// Required
func (GoToTiltValue) Required() bool { return false }

// Cluster ID of the command
func (GoToTiltValue) Cluster() zcl.ClusterID { return WindowCoveringID }

// MnfCode returns the manufacturer code (if any) of the command
func (GoToTiltValue) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GoToTiltValue) MarshalJSON() ([]byte, error) { return []byte("7"), nil }

// MarshalZcl returns the wire format representation of GoToTiltValue
func (v GoToTiltValue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Position.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GoToTiltValue struct
func (v *GoToTiltValue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Position).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GoToTiltValue) String() string {
	return zcl.Sprintf(
		"GoToTiltValue{"+zcl.StrJoin([]string{
			"Position(%v)",
		}, " ")+"}",
		v.Position,
	)
}

type GoToTiltPercentage struct {
	Percentage Percentage
}

// GoToTiltPercentageCommand is the Command ID of GoToTiltPercentage
const GoToTiltPercentageCommand CommandID = 0x0008

// Values returns all values of GoToTiltPercentage
func (v *GoToTiltPercentage) Values() []zcl.Val {
	return []zcl.Val{
		&v.Percentage,
	}
}

// Arguments returns all values of GoToTiltPercentage
func (v *GoToTiltPercentage) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Percentage", Argument: &v.Percentage},
	}
}

// Name of the command
func (GoToTiltPercentage) Name() string { return "Go to Tilt Percentage" }

// ID of the command
func (GoToTiltPercentage) ID() CommandID { return GoToTiltPercentageCommand }

// Required
func (GoToTiltPercentage) Required() bool { return false }

// Cluster ID of the command
func (GoToTiltPercentage) Cluster() zcl.ClusterID { return WindowCoveringID }

// MnfCode returns the manufacturer code (if any) of the command
func (GoToTiltPercentage) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (GoToTiltPercentage) MarshalJSON() ([]byte, error) { return []byte("8"), nil }

// MarshalZcl returns the wire format representation of GoToTiltPercentage
func (v GoToTiltPercentage) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Percentage.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the GoToTiltPercentage struct
func (v *GoToTiltPercentage) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Percentage).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v GoToTiltPercentage) String() string {
	return zcl.Sprintf(
		"GoToTiltPercentage{"+zcl.StrJoin([]string{
			"Percentage(%v)",
		}, " ")+"}",
		v.Percentage,
	)
}
