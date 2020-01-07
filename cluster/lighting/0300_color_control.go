package lighting

import "hemtjan.st/zcl"

// ColorControl
const ColorControlID zcl.ClusterID = 768

var ColorControlCluster = zcl.Cluster{
	Name: "Color control",
	ServerCmd: map[zcl.CommandID]func() zcl.Command{
		MoveToHueCommand:                      func() zcl.Command { return new(MoveToHue) },
		MoveHueCommand:                        func() zcl.Command { return new(MoveHue) },
		StepHueCommand:                        func() zcl.Command { return new(StepHue) },
		MoveToSaturationCommand:               func() zcl.Command { return new(MoveToSaturation) },
		MoveSaturationCommand:                 func() zcl.Command { return new(MoveSaturation) },
		StepSaturationCommand:                 func() zcl.Command { return new(StepSaturation) },
		MoveToHueAndSaturationCommand:         func() zcl.Command { return new(MoveToHueAndSaturation) },
		MoveToColorCommand:                    func() zcl.Command { return new(MoveToColor) },
		MoveColorCommand:                      func() zcl.Command { return new(MoveColor) },
		StepColorCommand:                      func() zcl.Command { return new(StepColor) },
		MoveToColorTemperatureCommand:         func() zcl.Command { return new(MoveToColorTemperature) },
		EnhancedMoveToHueCommand:              func() zcl.Command { return new(EnhancedMoveToHue) },
		EnhancedMoveHueCommand:                func() zcl.Command { return new(EnhancedMoveHue) },
		EnhancedStepHueCommand:                func() zcl.Command { return new(EnhancedStepHue) },
		EnhancedMoveToHueAndSaturationCommand: func() zcl.Command { return new(EnhancedMoveToHueAndSaturation) },
		ColorLoopSetCommand:                   func() zcl.Command { return new(ColorLoopSet) },
		StopMoveStepCommand:                   func() zcl.Command { return new(StopMoveStep) },
		MoveColorTemperatureCommand:           func() zcl.Command { return new(MoveColorTemperature) },
		StepColorTemperatureCommand:           func() zcl.Command { return new(StepColorTemperature) },
	},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		CurrentHueAttr:                        func() zcl.Attr { return new(CurrentHue) },
		CurrentSaturationAttr:                 func() zcl.Attr { return new(CurrentSaturation) },
		RemainingTimeAttr:                     func() zcl.Attr { return new(RemainingTime) },
		CurrentXAttr:                          func() zcl.Attr { return new(CurrentX) },
		CurrentYAttr:                          func() zcl.Attr { return new(CurrentY) },
		DriftCompensationAttr:                 func() zcl.Attr { return new(DriftCompensation) },
		CompensationTextAttr:                  func() zcl.Attr { return new(CompensationText) },
		ColorTemperatureMiredsAttr:            func() zcl.Attr { return new(ColorTemperatureMireds) },
		ColorModeAttr:                         func() zcl.Attr { return new(ColorMode) },
		ColorControlOptionsAttr:               func() zcl.Attr { return new(ColorControlOptions) },
		EnhancedCurrentHueAttr:                func() zcl.Attr { return new(EnhancedCurrentHue) },
		EnhancedColorModeAttr:                 func() zcl.Attr { return new(EnhancedColorMode) },
		ColorLoopActiveAttr:                   func() zcl.Attr { return new(ColorLoopActive) },
		ColorLoopDirectionAttr:                func() zcl.Attr { return new(ColorLoopDirection) },
		ColorLoopTimeAttr:                     func() zcl.Attr { return new(ColorLoopTime) },
		ColorLoopStartEnhancedHueAttr:         func() zcl.Attr { return new(ColorLoopStartEnhancedHue) },
		ColorLoopStoredEnhancedHueAttr:        func() zcl.Attr { return new(ColorLoopStoredEnhancedHue) },
		ColorCapabilitiesAttr:                 func() zcl.Attr { return new(ColorCapabilities) },
		ColorTemperaturePhysicalMinMiredsAttr: func() zcl.Attr { return new(ColorTemperaturePhysicalMinMireds) },
		ColorTemperaturePhysicalMaxMiredsAttr: func() zcl.Attr { return new(ColorTemperaturePhysicalMaxMireds) },
		CoupleColorTempToLevelMinMiredsAttr:   func() zcl.Attr { return new(CoupleColorTempToLevelMinMireds) },
		PowerOnColorTemperatureAttr:           func() zcl.Attr { return new(PowerOnColorTemperature) },
		NumberOfPrimariesAttr:                 func() zcl.Attr { return new(NumberOfPrimaries) },
		Primary1XAttr:                         func() zcl.Attr { return new(Primary1X) },
		Primary1YAttr:                         func() zcl.Attr { return new(Primary1Y) },
		Primary1IntensityAttr:                 func() zcl.Attr { return new(Primary1Intensity) },
		Primary2XAttr:                         func() zcl.Attr { return new(Primary2X) },
		Primary2YAttr:                         func() zcl.Attr { return new(Primary2Y) },
		Primary2IntensityAttr:                 func() zcl.Attr { return new(Primary2Intensity) },
		Primary3XAttr:                         func() zcl.Attr { return new(Primary3X) },
		Primary3YAttr:                         func() zcl.Attr { return new(Primary3Y) },
		Primary3IntensityAttr:                 func() zcl.Attr { return new(Primary3Intensity) },
		Primary4XAttr:                         func() zcl.Attr { return new(Primary4X) },
		Primary4YAttr:                         func() zcl.Attr { return new(Primary4Y) },
		Primary4IntensityAttr:                 func() zcl.Attr { return new(Primary4Intensity) },
		Primary5XAttr:                         func() zcl.Attr { return new(Primary5X) },
		Primary5YAttr:                         func() zcl.Attr { return new(Primary5Y) },
		Primary5IntensityAttr:                 func() zcl.Attr { return new(Primary5Intensity) },
		Primary6XAttr:                         func() zcl.Attr { return new(Primary6X) },
		Primary6YAttr:                         func() zcl.Attr { return new(Primary6Y) },
		Primary6IntensityAttr:                 func() zcl.Attr { return new(Primary6Intensity) },
		WhitePointXAttr:                       func() zcl.Attr { return new(WhitePointX) },
		WhitePointYAttr:                       func() zcl.Attr { return new(WhitePointY) },
		ColorPointRedXAttr:                    func() zcl.Attr { return new(ColorPointRedX) },
		ColorPointRedYAttr:                    func() zcl.Attr { return new(ColorPointRedY) },
		ColorPointRedIntensityAttr:            func() zcl.Attr { return new(ColorPointRedIntensity) },
		ColorPointGreenXAttr:                  func() zcl.Attr { return new(ColorPointGreenX) },
		ColorPointGreenYAttr:                  func() zcl.Attr { return new(ColorPointGreenY) },
		ColorPointGreenIntensityAttr:          func() zcl.Attr { return new(ColorPointGreenIntensity) },
		ColorPointBlueXAttr:                   func() zcl.Attr { return new(ColorPointBlueX) },
		ColorPointBlueYAttr:                   func() zcl.Attr { return new(ColorPointBlueY) },
		ColorPointBlueIntensityAttr:           func() zcl.Attr { return new(ColorPointBlueIntensity) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr: []zcl.AttrID{
		CurrentXAttr,
		CurrentYAttr,
		EnhancedCurrentHueAttr,
		CurrentSaturationAttr,
		ColorLoopActiveAttr,
		ColorLoopDirectionAttr,
		ColorLoopTimeAttr,
	},
}

type MoveToHue struct {
	Hue       Hue
	Direction Direction
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// MoveToHueCommand is the Command ID of MoveToHue
const MoveToHueCommand CommandID = 0x0000

// Values returns all values of MoveToHue
func (v *MoveToHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.Hue,
		&v.Direction,
		&v.TransitionTime,
	}
}

// Arguments returns all values of MoveToHue
func (v *MoveToHue) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Hue", Argument: &v.Hue},
		{Name: "Direction", Argument: &v.Direction},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (MoveToHue) Name() string { return "Move to hue" }

// ID of the command
func (MoveToHue) ID() CommandID { return MoveToHueCommand }

// Required
func (MoveToHue) Required() bool { return false }

// Cluster ID of the command
func (MoveToHue) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveToHue) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveToHue) MarshalJSON() ([]byte, error) { return []byte("0"), nil }

// MarshalZcl returns the wire format representation of MoveToHue
func (v MoveToHue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Hue.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Direction.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the MoveToHue struct
func (v *MoveToHue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Hue).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Direction).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveToHue) String() string {
	return zcl.Sprintf(
		"MoveToHue{"+zcl.StrJoin([]string{
			"Hue(%v)",
			"Direction(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.Hue,
		v.Direction,
		v.TransitionTime,
	)
}

type MoveHue struct {
	MoveMode MoveMode
	// Rate increment/steps per second
	Rate Rate
}

// MoveHueCommand is the Command ID of MoveHue
const MoveHueCommand CommandID = 0x0001

// Values returns all values of MoveHue
func (v *MoveHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.MoveMode,
		&v.Rate,
	}
}

// Arguments returns all values of MoveHue
func (v *MoveHue) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "MoveMode", Argument: &v.MoveMode},
		{Name: "Rate", Argument: &v.Rate},
	}
}

// Name of the command
func (MoveHue) Name() string { return "Move hue" }

// ID of the command
func (MoveHue) ID() CommandID { return MoveHueCommand }

// Required
func (MoveHue) Required() bool { return false }

// Cluster ID of the command
func (MoveHue) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveHue) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveHue) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

// MarshalZcl returns the wire format representation of MoveHue
func (v MoveHue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.MoveMode.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the MoveHue struct
func (v *MoveHue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.MoveMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveHue) String() string {
	return zcl.Sprintf(
		"MoveHue{"+zcl.StrJoin([]string{
			"MoveMode(%v)",
			"Rate(%v)",
		}, " ")+"}",
		v.MoveMode,
		v.Rate,
	)
}

type StepHue struct {
	StepMode StepMode
	StepSize StepSize
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// StepHueCommand is the Command ID of StepHue
const StepHueCommand CommandID = 0x0002

// Values returns all values of StepHue
func (v *StepHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepMode,
		&v.StepSize,
		&v.TransitionTime,
	}
}

// Arguments returns all values of StepHue
func (v *StepHue) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "StepMode", Argument: &v.StepMode},
		{Name: "StepSize", Argument: &v.StepSize},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (StepHue) Name() string { return "Step hue" }

// ID of the command
func (StepHue) ID() CommandID { return StepHueCommand }

// Required
func (StepHue) Required() bool { return false }

// Cluster ID of the command
func (StepHue) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (StepHue) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (StepHue) MarshalJSON() ([]byte, error) { return []byte("2"), nil }

// MarshalZcl returns the wire format representation of StepHue
func (v StepHue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.StepMode.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the StepHue struct
func (v *StepHue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

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

// String returns a log-friendly string representation of the struct
func (v StepHue) String() string {
	return zcl.Sprintf(
		"StepHue{"+zcl.StrJoin([]string{
			"StepMode(%v)",
			"StepSize(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.StepMode,
		v.StepSize,
		v.TransitionTime,
	)
}

type MoveToSaturation struct {
	Saturation Saturation
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// MoveToSaturationCommand is the Command ID of MoveToSaturation
const MoveToSaturationCommand CommandID = 0x0003

// Values returns all values of MoveToSaturation
func (v *MoveToSaturation) Values() []zcl.Val {
	return []zcl.Val{
		&v.Saturation,
		&v.TransitionTime,
	}
}

// Arguments returns all values of MoveToSaturation
func (v *MoveToSaturation) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Saturation", Argument: &v.Saturation},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (MoveToSaturation) Name() string { return "Move to saturation" }

// ID of the command
func (MoveToSaturation) ID() CommandID { return MoveToSaturationCommand }

// Required
func (MoveToSaturation) Required() bool { return false }

// Cluster ID of the command
func (MoveToSaturation) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveToSaturation) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveToSaturation) MarshalJSON() ([]byte, error) { return []byte("3"), nil }

// MarshalZcl returns the wire format representation of MoveToSaturation
func (v MoveToSaturation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Saturation.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the MoveToSaturation struct
func (v *MoveToSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Saturation).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveToSaturation) String() string {
	return zcl.Sprintf(
		"MoveToSaturation{"+zcl.StrJoin([]string{
			"Saturation(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.Saturation,
		v.TransitionTime,
	)
}

type MoveSaturation struct {
	MoveMode MoveMode
	// Rate increment/steps per second
	Rate Rate
}

// MoveSaturationCommand is the Command ID of MoveSaturation
const MoveSaturationCommand CommandID = 0x0004

// Values returns all values of MoveSaturation
func (v *MoveSaturation) Values() []zcl.Val {
	return []zcl.Val{
		&v.MoveMode,
		&v.Rate,
	}
}

// Arguments returns all values of MoveSaturation
func (v *MoveSaturation) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "MoveMode", Argument: &v.MoveMode},
		{Name: "Rate", Argument: &v.Rate},
	}
}

// Name of the command
func (MoveSaturation) Name() string { return "Move saturation" }

// ID of the command
func (MoveSaturation) ID() CommandID { return MoveSaturationCommand }

// Required
func (MoveSaturation) Required() bool { return false }

// Cluster ID of the command
func (MoveSaturation) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveSaturation) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveSaturation) MarshalJSON() ([]byte, error) { return []byte("4"), nil }

// MarshalZcl returns the wire format representation of MoveSaturation
func (v MoveSaturation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.MoveMode.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the MoveSaturation struct
func (v *MoveSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.MoveMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveSaturation) String() string {
	return zcl.Sprintf(
		"MoveSaturation{"+zcl.StrJoin([]string{
			"MoveMode(%v)",
			"Rate(%v)",
		}, " ")+"}",
		v.MoveMode,
		v.Rate,
	)
}

type StepSaturation struct {
	StepMode StepMode
	StepSize StepSize
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// StepSaturationCommand is the Command ID of StepSaturation
const StepSaturationCommand CommandID = 0x0005

// Values returns all values of StepSaturation
func (v *StepSaturation) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepMode,
		&v.StepSize,
		&v.TransitionTime,
	}
}

// Arguments returns all values of StepSaturation
func (v *StepSaturation) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "StepMode", Argument: &v.StepMode},
		{Name: "StepSize", Argument: &v.StepSize},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (StepSaturation) Name() string { return "Step saturation" }

// ID of the command
func (StepSaturation) ID() CommandID { return StepSaturationCommand }

// Required
func (StepSaturation) Required() bool { return false }

// Cluster ID of the command
func (StepSaturation) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (StepSaturation) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (StepSaturation) MarshalJSON() ([]byte, error) { return []byte("5"), nil }

// MarshalZcl returns the wire format representation of StepSaturation
func (v StepSaturation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.StepMode.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the StepSaturation struct
func (v *StepSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

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

// String returns a log-friendly string representation of the struct
func (v StepSaturation) String() string {
	return zcl.Sprintf(
		"StepSaturation{"+zcl.StrJoin([]string{
			"StepMode(%v)",
			"StepSize(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.StepMode,
		v.StepSize,
		v.TransitionTime,
	)
}

type MoveToHueAndSaturation struct {
	Hue        Hue
	Saturation Saturation
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// MoveToHueAndSaturationCommand is the Command ID of MoveToHueAndSaturation
const MoveToHueAndSaturationCommand CommandID = 0x0006

// Values returns all values of MoveToHueAndSaturation
func (v *MoveToHueAndSaturation) Values() []zcl.Val {
	return []zcl.Val{
		&v.Hue,
		&v.Saturation,
		&v.TransitionTime,
	}
}

// Arguments returns all values of MoveToHueAndSaturation
func (v *MoveToHueAndSaturation) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "Hue", Argument: &v.Hue},
		{Name: "Saturation", Argument: &v.Saturation},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (MoveToHueAndSaturation) Name() string { return "Move to hue and saturation" }

// ID of the command
func (MoveToHueAndSaturation) ID() CommandID { return MoveToHueAndSaturationCommand }

// Required
func (MoveToHueAndSaturation) Required() bool { return false }

// Cluster ID of the command
func (MoveToHueAndSaturation) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveToHueAndSaturation) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveToHueAndSaturation) MarshalJSON() ([]byte, error) { return []byte("6"), nil }

// MarshalZcl returns the wire format representation of MoveToHueAndSaturation
func (v MoveToHueAndSaturation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.Hue.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Saturation.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the MoveToHueAndSaturation struct
func (v *MoveToHueAndSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.Hue).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Saturation).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveToHueAndSaturation) String() string {
	return zcl.Sprintf(
		"MoveToHueAndSaturation{"+zcl.StrJoin([]string{
			"Hue(%v)",
			"Saturation(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.Hue,
		v.Saturation,
		v.TransitionTime,
	)
}

type MoveToColor struct {
	// ColorX contains the normalized chromaticity value x for this attribute, as
	// defined in the CIE xyY Color Space. x = value / 65536 (value
	// in the range 0 to 65279 inclusive)
	ColorX ColorX
	// ColorY contains the normalized chromaticity value y for this attribute, as
	// defined in the CIE xyY Color Space. y = value / 65536 (value
	// in the range 0 to 65279 inclusive)
	ColorY ColorY
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// MoveToColorCommand is the Command ID of MoveToColor
const MoveToColorCommand CommandID = 0x0007

// Values returns all values of MoveToColor
func (v *MoveToColor) Values() []zcl.Val {
	return []zcl.Val{
		&v.ColorX,
		&v.ColorY,
		&v.TransitionTime,
	}
}

// Arguments returns all values of MoveToColor
func (v *MoveToColor) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "ColorX", Argument: &v.ColorX},
		{Name: "ColorY", Argument: &v.ColorY},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (MoveToColor) Name() string { return "Move to color" }

// ID of the command
func (MoveToColor) ID() CommandID { return MoveToColorCommand }

// Required
func (MoveToColor) Required() bool { return true }

// Cluster ID of the command
func (MoveToColor) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveToColor) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveToColor) MarshalJSON() ([]byte, error) { return []byte("7"), nil }

// MarshalZcl returns the wire format representation of MoveToColor
func (v MoveToColor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.ColorX.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ColorY.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the MoveToColor struct
func (v *MoveToColor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.ColorX).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ColorY).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveToColor) String() string {
	return zcl.Sprintf(
		"MoveToColor{"+zcl.StrJoin([]string{
			"ColorX(%v)",
			"ColorY(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.ColorX,
		v.ColorY,
		v.TransitionTime,
	)
}

type MoveColor struct {
	// RateX increment/steps per second
	RateX RateX
	// RateY increment/steps per second
	RateY RateY
}

// MoveColorCommand is the Command ID of MoveColor
const MoveColorCommand CommandID = 0x0008

// Values returns all values of MoveColor
func (v *MoveColor) Values() []zcl.Val {
	return []zcl.Val{
		&v.RateX,
		&v.RateY,
	}
}

// Arguments returns all values of MoveColor
func (v *MoveColor) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "RateX", Argument: &v.RateX},
		{Name: "RateY", Argument: &v.RateY},
	}
}

// Name of the command
func (MoveColor) Name() string { return "Move color" }

// ID of the command
func (MoveColor) ID() CommandID { return MoveColorCommand }

// Required
func (MoveColor) Required() bool { return false }

// Cluster ID of the command
func (MoveColor) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveColor) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveColor) MarshalJSON() ([]byte, error) { return []byte("8"), nil }

// MarshalZcl returns the wire format representation of MoveColor
func (v MoveColor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.RateX.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.RateY.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MoveColor struct
func (v *MoveColor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.RateX).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.RateY).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveColor) String() string {
	return zcl.Sprintf(
		"MoveColor{"+zcl.StrJoin([]string{
			"RateX(%v)",
			"RateY(%v)",
		}, " ")+"}",
		v.RateX,
		v.RateY,
	)
}

type StepColor struct {
	StepX StepX
	StepY StepY
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// StepColorCommand is the Command ID of StepColor
const StepColorCommand CommandID = 0x0009

// Values returns all values of StepColor
func (v *StepColor) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepX,
		&v.StepY,
		&v.TransitionTime,
	}
}

// Arguments returns all values of StepColor
func (v *StepColor) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "StepX", Argument: &v.StepX},
		{Name: "StepY", Argument: &v.StepY},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (StepColor) Name() string { return "Step color" }

// ID of the command
func (StepColor) ID() CommandID { return StepColorCommand }

// Required
func (StepColor) Required() bool { return false }

// Cluster ID of the command
func (StepColor) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (StepColor) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (StepColor) MarshalJSON() ([]byte, error) { return []byte("9"), nil }

// MarshalZcl returns the wire format representation of StepColor
func (v StepColor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.StepX.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.StepY.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the StepColor struct
func (v *StepColor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.StepX).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StepY).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v StepColor) String() string {
	return zcl.Sprintf(
		"StepColor{"+zcl.StrJoin([]string{
			"StepX(%v)",
			"StepY(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.StepX,
		v.StepY,
		v.TransitionTime,
	)
}

type MoveToColorTemperature struct {
	ColorTemperature ColorTemperature
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// MoveToColorTemperatureCommand is the Command ID of MoveToColorTemperature
const MoveToColorTemperatureCommand CommandID = 0x000A

// Values returns all values of MoveToColorTemperature
func (v *MoveToColorTemperature) Values() []zcl.Val {
	return []zcl.Val{
		&v.ColorTemperature,
		&v.TransitionTime,
	}
}

// Arguments returns all values of MoveToColorTemperature
func (v *MoveToColorTemperature) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "ColorTemperature", Argument: &v.ColorTemperature},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (MoveToColorTemperature) Name() string { return "Move to color temperature" }

// ID of the command
func (MoveToColorTemperature) ID() CommandID { return MoveToColorTemperatureCommand }

// Required
func (MoveToColorTemperature) Required() bool { return false }

// Cluster ID of the command
func (MoveToColorTemperature) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveToColorTemperature) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveToColorTemperature) MarshalJSON() ([]byte, error) { return []byte("10"), nil }

// MarshalZcl returns the wire format representation of MoveToColorTemperature
func (v MoveToColorTemperature) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.ColorTemperature.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the MoveToColorTemperature struct
func (v *MoveToColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.ColorTemperature).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveToColorTemperature) String() string {
	return zcl.Sprintf(
		"MoveToColorTemperature{"+zcl.StrJoin([]string{
			"ColorTemperature(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.ColorTemperature,
		v.TransitionTime,
	)
}

type EnhancedMoveToHue struct {
	EnhancedHue EnhancedHue
	Direction   Direction
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// EnhancedMoveToHueCommand is the Command ID of EnhancedMoveToHue
const EnhancedMoveToHueCommand CommandID = 0x0040

// Values returns all values of EnhancedMoveToHue
func (v *EnhancedMoveToHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.EnhancedHue,
		&v.Direction,
		&v.TransitionTime,
	}
}

// Arguments returns all values of EnhancedMoveToHue
func (v *EnhancedMoveToHue) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "EnhancedHue", Argument: &v.EnhancedHue},
		{Name: "Direction", Argument: &v.Direction},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (EnhancedMoveToHue) Name() string { return "Enhanced move to hue" }

// ID of the command
func (EnhancedMoveToHue) ID() CommandID { return EnhancedMoveToHueCommand }

// Required
func (EnhancedMoveToHue) Required() bool { return true }

// Cluster ID of the command
func (EnhancedMoveToHue) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (EnhancedMoveToHue) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (EnhancedMoveToHue) MarshalJSON() ([]byte, error) { return []byte("64"), nil }

// MarshalZcl returns the wire format representation of EnhancedMoveToHue
func (v EnhancedMoveToHue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.EnhancedHue.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Direction.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the EnhancedMoveToHue struct
func (v *EnhancedMoveToHue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.EnhancedHue).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Direction).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v EnhancedMoveToHue) String() string {
	return zcl.Sprintf(
		"EnhancedMoveToHue{"+zcl.StrJoin([]string{
			"EnhancedHue(%v)",
			"Direction(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.EnhancedHue,
		v.Direction,
		v.TransitionTime,
	)
}

type EnhancedMoveHue struct {
	MoveMode MoveMode
	// Rate increment/steps per second
	Rate Rate
}

// EnhancedMoveHueCommand is the Command ID of EnhancedMoveHue
const EnhancedMoveHueCommand CommandID = 0x0041

// Values returns all values of EnhancedMoveHue
func (v *EnhancedMoveHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.MoveMode,
		&v.Rate,
	}
}

// Arguments returns all values of EnhancedMoveHue
func (v *EnhancedMoveHue) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "MoveMode", Argument: &v.MoveMode},
		{Name: "Rate", Argument: &v.Rate},
	}
}

// Name of the command
func (EnhancedMoveHue) Name() string { return "Enhanced move hue" }

// ID of the command
func (EnhancedMoveHue) ID() CommandID { return EnhancedMoveHueCommand }

// Required
func (EnhancedMoveHue) Required() bool { return true }

// Cluster ID of the command
func (EnhancedMoveHue) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (EnhancedMoveHue) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (EnhancedMoveHue) MarshalJSON() ([]byte, error) { return []byte("65"), nil }

// MarshalZcl returns the wire format representation of EnhancedMoveHue
func (v EnhancedMoveHue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.MoveMode.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the EnhancedMoveHue struct
func (v *EnhancedMoveHue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.MoveMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v EnhancedMoveHue) String() string {
	return zcl.Sprintf(
		"EnhancedMoveHue{"+zcl.StrJoin([]string{
			"MoveMode(%v)",
			"Rate(%v)",
		}, " ")+"}",
		v.MoveMode,
		v.Rate,
	)
}

type EnhancedStepHue struct {
	StepMode StepMode
	StepSize StepSize
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// EnhancedStepHueCommand is the Command ID of EnhancedStepHue
const EnhancedStepHueCommand CommandID = 0x0042

// Values returns all values of EnhancedStepHue
func (v *EnhancedStepHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepMode,
		&v.StepSize,
		&v.TransitionTime,
	}
}

// Arguments returns all values of EnhancedStepHue
func (v *EnhancedStepHue) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "StepMode", Argument: &v.StepMode},
		{Name: "StepSize", Argument: &v.StepSize},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (EnhancedStepHue) Name() string { return "Enhanced step hue" }

// ID of the command
func (EnhancedStepHue) ID() CommandID { return EnhancedStepHueCommand }

// Required
func (EnhancedStepHue) Required() bool { return true }

// Cluster ID of the command
func (EnhancedStepHue) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (EnhancedStepHue) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (EnhancedStepHue) MarshalJSON() ([]byte, error) { return []byte("66"), nil }

// MarshalZcl returns the wire format representation of EnhancedStepHue
func (v EnhancedStepHue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.StepMode.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the EnhancedStepHue struct
func (v *EnhancedStepHue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

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

// String returns a log-friendly string representation of the struct
func (v EnhancedStepHue) String() string {
	return zcl.Sprintf(
		"EnhancedStepHue{"+zcl.StrJoin([]string{
			"StepMode(%v)",
			"StepSize(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.StepMode,
		v.StepSize,
		v.TransitionTime,
	)
}

type EnhancedMoveToHueAndSaturation struct {
	EnhancedHue EnhancedHue
	Saturation  Saturation
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime TransitionTime
}

// EnhancedMoveToHueAndSaturationCommand is the Command ID of EnhancedMoveToHueAndSaturation
const EnhancedMoveToHueAndSaturationCommand CommandID = 0x0043

// Values returns all values of EnhancedMoveToHueAndSaturation
func (v *EnhancedMoveToHueAndSaturation) Values() []zcl.Val {
	return []zcl.Val{
		&v.EnhancedHue,
		&v.Saturation,
		&v.TransitionTime,
	}
}

// Arguments returns all values of EnhancedMoveToHueAndSaturation
func (v *EnhancedMoveToHueAndSaturation) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "EnhancedHue", Argument: &v.EnhancedHue},
		{Name: "Saturation", Argument: &v.Saturation},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
	}
}

// Name of the command
func (EnhancedMoveToHueAndSaturation) Name() string { return "Enhanced move to hue and saturation" }

// ID of the command
func (EnhancedMoveToHueAndSaturation) ID() CommandID { return EnhancedMoveToHueAndSaturationCommand }

// Required
func (EnhancedMoveToHueAndSaturation) Required() bool { return true }

// Cluster ID of the command
func (EnhancedMoveToHueAndSaturation) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (EnhancedMoveToHueAndSaturation) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (EnhancedMoveToHueAndSaturation) MarshalJSON() ([]byte, error) { return []byte("67"), nil }

// MarshalZcl returns the wire format representation of EnhancedMoveToHueAndSaturation
func (v EnhancedMoveToHueAndSaturation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.EnhancedHue.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Saturation.MarshalZcl(); err != nil {
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

// UnmarshalZcl parses the wire format representation into the EnhancedMoveToHueAndSaturation struct
func (v *EnhancedMoveToHueAndSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.EnhancedHue).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Saturation).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v EnhancedMoveToHueAndSaturation) String() string {
	return zcl.Sprintf(
		"EnhancedMoveToHueAndSaturation{"+zcl.StrJoin([]string{
			"EnhancedHue(%v)",
			"Saturation(%v)",
			"TransitionTime(%v)",
		}, " ")+"}",
		v.EnhancedHue,
		v.Saturation,
		v.TransitionTime,
	)
}

type ColorLoopSet struct {
	UpdateFlags  UpdateFlags
	Action       Action
	HueDirection HueDirection
	// Time Time in seconds used for a whole color loop.
	Time        Time
	EnhancedHue EnhancedHue
}

// ColorLoopSetCommand is the Command ID of ColorLoopSet
const ColorLoopSetCommand CommandID = 0x0044

// Values returns all values of ColorLoopSet
func (v *ColorLoopSet) Values() []zcl.Val {
	return []zcl.Val{
		&v.UpdateFlags,
		&v.Action,
		&v.HueDirection,
		&v.Time,
		&v.EnhancedHue,
	}
}

// Arguments returns all values of ColorLoopSet
func (v *ColorLoopSet) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "UpdateFlags", Argument: &v.UpdateFlags},
		{Name: "Action", Argument: &v.Action},
		{Name: "HueDirection", Argument: &v.HueDirection},
		{Name: "Time", Argument: &v.Time},
		{Name: "EnhancedHue", Argument: &v.EnhancedHue},
	}
}

// Name of the command
func (ColorLoopSet) Name() string { return "Color loop set" }

// ID of the command
func (ColorLoopSet) ID() CommandID { return ColorLoopSetCommand }

// Required
func (ColorLoopSet) Required() bool { return true }

// Cluster ID of the command
func (ColorLoopSet) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (ColorLoopSet) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (ColorLoopSet) MarshalJSON() ([]byte, error) { return []byte("68"), nil }

// MarshalZcl returns the wire format representation of ColorLoopSet
func (v ColorLoopSet) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.UpdateFlags.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Action.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.HueDirection.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.Time.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.EnhancedHue.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the ColorLoopSet struct
func (v *ColorLoopSet) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.UpdateFlags).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Action).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.HueDirection).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Time).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.EnhancedHue).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v ColorLoopSet) String() string {
	return zcl.Sprintf(
		"ColorLoopSet{"+zcl.StrJoin([]string{
			"UpdateFlags(%v)",
			"Action(%v)",
			"HueDirection(%v)",
			"Time(%v)",
			"EnhancedHue(%v)",
		}, " ")+"}",
		v.UpdateFlags,
		v.Action,
		v.HueDirection,
		v.Time,
		v.EnhancedHue,
	)
}

// StopMoveStep Stops move to and step commands. It has no effect on a active
// color loop.
type StopMoveStep struct {
}

// StopMoveStepCommand is the Command ID of StopMoveStep
const StopMoveStepCommand CommandID = 0x0047

// Values returns all values of StopMoveStep
func (v *StopMoveStep) Values() []zcl.Val {
	return []zcl.Val{}
}

// Arguments returns all values of StopMoveStep
func (v *StopMoveStep) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{}
}

// Name of the command
func (StopMoveStep) Name() string { return "Stop move step" }

// ID of the command
func (StopMoveStep) ID() CommandID { return StopMoveStepCommand }

// Required
func (StopMoveStep) Required() bool { return true }

// Cluster ID of the command
func (StopMoveStep) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (StopMoveStep) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (StopMoveStep) MarshalJSON() ([]byte, error) { return []byte("71"), nil }

// MarshalZcl returns the wire format representation of StopMoveStep
func (v StopMoveStep) MarshalZcl() ([]byte, error) {
	return nil, nil
}

// UnmarshalZcl parses the wire format representation into the StopMoveStep struct
func (v *StopMoveStep) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v StopMoveStep) String() string {
	return zcl.Sprintf(
		"StopMoveStep{" + zcl.StrJoin([]string{}, " ") + "}",
	)
}

type MoveColorTemperature struct {
	MoveMode MoveMode
	// Rate increment/steps per second
	Rate                Rate
	ColorTemperatureMin ColorTemperatureMin
	ColorTemperatureMax ColorTemperatureMax
}

// MoveColorTemperatureCommand is the Command ID of MoveColorTemperature
const MoveColorTemperatureCommand CommandID = 0x004B

// Values returns all values of MoveColorTemperature
func (v *MoveColorTemperature) Values() []zcl.Val {
	return []zcl.Val{
		&v.MoveMode,
		&v.Rate,
		&v.ColorTemperatureMin,
		&v.ColorTemperatureMax,
	}
}

// Arguments returns all values of MoveColorTemperature
func (v *MoveColorTemperature) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "MoveMode", Argument: &v.MoveMode},
		{Name: "Rate", Argument: &v.Rate},
		{Name: "ColorTemperatureMin", Argument: &v.ColorTemperatureMin},
		{Name: "ColorTemperatureMax", Argument: &v.ColorTemperatureMax},
	}
}

// Name of the command
func (MoveColorTemperature) Name() string { return "Move color temperature" }

// ID of the command
func (MoveColorTemperature) ID() CommandID { return MoveColorTemperatureCommand }

// Required
func (MoveColorTemperature) Required() bool { return true }

// Cluster ID of the command
func (MoveColorTemperature) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (MoveColorTemperature) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (MoveColorTemperature) MarshalJSON() ([]byte, error) { return []byte("75"), nil }

// MarshalZcl returns the wire format representation of MoveColorTemperature
func (v MoveColorTemperature) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.MoveMode.MarshalZcl(); err != nil {
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
	{
		if tmp, err = v.ColorTemperatureMin.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ColorTemperatureMax.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the MoveColorTemperature struct
func (v *MoveColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.MoveMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ColorTemperatureMin).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ColorTemperatureMax).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v MoveColorTemperature) String() string {
	return zcl.Sprintf(
		"MoveColorTemperature{"+zcl.StrJoin([]string{
			"MoveMode(%v)",
			"Rate(%v)",
			"ColorTemperatureMin(%v)",
			"ColorTemperatureMax(%v)",
		}, " ")+"}",
		v.MoveMode,
		v.Rate,
		v.ColorTemperatureMin,
		v.ColorTemperatureMax,
	)
}

type StepColorTemperature struct {
	StepMode StepMode
	StepSize StepSize
	// TransitionTime The transition time in 1/10ths of a second.
	TransitionTime            TransitionTime
	ColorTemperatureMinMireds ColorTemperatureMinMireds
	ColorTemperatureMaxMireds ColorTemperatureMaxMireds
}

// StepColorTemperatureCommand is the Command ID of StepColorTemperature
const StepColorTemperatureCommand CommandID = 0x004C

// Values returns all values of StepColorTemperature
func (v *StepColorTemperature) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepMode,
		&v.StepSize,
		&v.TransitionTime,
		&v.ColorTemperatureMinMireds,
		&v.ColorTemperatureMaxMireds,
	}
}

// Arguments returns all values of StepColorTemperature
func (v *StepColorTemperature) Arguments() []zcl.ArgDesc {
	return []zcl.ArgDesc{
		{Name: "StepMode", Argument: &v.StepMode},
		{Name: "StepSize", Argument: &v.StepSize},
		{Name: "TransitionTime", Argument: &v.TransitionTime},
		{Name: "ColorTemperatureMinMireds", Argument: &v.ColorTemperatureMinMireds},
		{Name: "ColorTemperatureMaxMireds", Argument: &v.ColorTemperatureMaxMireds},
	}
}

// Name of the command
func (StepColorTemperature) Name() string { return "Step color temperature" }

// ID of the command
func (StepColorTemperature) ID() CommandID { return StepColorTemperatureCommand }

// Required
func (StepColorTemperature) Required() bool { return true }

// Cluster ID of the command
func (StepColorTemperature) Cluster() zcl.ClusterID { return ColorControlID }

// MnfCode returns the manufacturer code (if any) of the command
func (StepColorTemperature) MnfCode() []byte { return []byte{} }

// MarshalJSON is a helper that returns the command as an uint wrapped in a byte-array
// func (StepColorTemperature) MarshalJSON() ([]byte, error) { return []byte("76"), nil }

// MarshalZcl returns the wire format representation of StepColorTemperature
func (v StepColorTemperature) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	tmp2 := uint32(0)
	_ = tmp2
	var err error

	{
		if tmp, err = v.StepMode.MarshalZcl(); err != nil {
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
	{
		if tmp, err = v.ColorTemperatureMinMireds.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	{
		if tmp, err = v.ColorTemperatureMaxMireds.MarshalZcl(); err != nil {
			return nil, err
		}
		data = append(data, tmp...)
	}
	return data, nil
}

// UnmarshalZcl parses the wire format representation into the StepColorTemperature struct
func (v *StepColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error
	tmp2 := uint32(0)
	_ = tmp2

	if b, err = (&v.StepMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StepSize).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ColorTemperatureMinMireds).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ColorTemperatureMaxMireds).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// String returns a log-friendly string representation of the struct
func (v StepColorTemperature) String() string {
	return zcl.Sprintf(
		"StepColorTemperature{"+zcl.StrJoin([]string{
			"StepMode(%v)",
			"StepSize(%v)",
			"TransitionTime(%v)",
			"ColorTemperatureMinMireds(%v)",
			"ColorTemperatureMaxMireds(%v)",
		}, " ")+"}",
		v.StepMode,
		v.StepSize,
		v.TransitionTime,
		v.ColorTemperatureMinMireds,
		v.ColorTemperatureMaxMireds,
	)
}
