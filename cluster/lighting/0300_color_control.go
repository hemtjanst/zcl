// Attributes and commands for controlling the color properties of a color-capable light.
package lighting

import (
	"neotor.se/zcl"
)

// ColorControl
const ColorControlID zcl.ClusterID = 768

var ColorControlCluster = zcl.Cluster{
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
		CurrentXInCieXyyAttr:                  func() zcl.Attr { return new(CurrentXInCieXyy) },
		CurrentYInCieXyyAttr:                  func() zcl.Attr { return new(CurrentYInCieXyy) },
		DriftCompensationAttr:                 func() zcl.Attr { return new(DriftCompensation) },
		CompensationTextAttr:                  func() zcl.Attr { return new(CompensationText) },
		ColorTemperatureMiredsAttr:            func() zcl.Attr { return new(ColorTemperatureMireds) },
		ColorModeAttr:                         func() zcl.Attr { return new(ColorMode) },
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
		PowerOnColorTemperatureAttr:           func() zcl.Attr { return new(PowerOnColorTemperature) },
		NumberOfPrimariesAttr:                 func() zcl.Attr { return new(NumberOfPrimaries) },
		Primary1XInCieXyyAttr:                 func() zcl.Attr { return new(Primary1XInCieXyy) },
		Primary1YInCieXyyAttr:                 func() zcl.Attr { return new(Primary1YInCieXyy) },
		Primary1IntensityAttr:                 func() zcl.Attr { return new(Primary1Intensity) },
		Primary2XInCieXyyAttr:                 func() zcl.Attr { return new(Primary2XInCieXyy) },
		Primary2YInCieXyyAttr:                 func() zcl.Attr { return new(Primary2YInCieXyy) },
		Primary2IntensityAttr:                 func() zcl.Attr { return new(Primary2Intensity) },
		Primary3XInCieXyyAttr:                 func() zcl.Attr { return new(Primary3XInCieXyy) },
		Primary3YInCieXyyAttr:                 func() zcl.Attr { return new(Primary3YInCieXyy) },
		Primary3IntensityAttr:                 func() zcl.Attr { return new(Primary3Intensity) },
		Primary4XInCieXyyAttr:                 func() zcl.Attr { return new(Primary4XInCieXyy) },
		Primary4YInCieXyyAttr:                 func() zcl.Attr { return new(Primary4YInCieXyy) },
		Primary4IntensityAttr:                 func() zcl.Attr { return new(Primary4Intensity) },
		Primary5XInCieXyyAttr:                 func() zcl.Attr { return new(Primary5XInCieXyy) },
		Primary5YInCieXyyAttr:                 func() zcl.Attr { return new(Primary5YInCieXyy) },
		Primary5IntensityAttr:                 func() zcl.Attr { return new(Primary5Intensity) },
		Primary6XInCieXyyAttr:                 func() zcl.Attr { return new(Primary6XInCieXyy) },
		Primary6YInCieXyyAttr:                 func() zcl.Attr { return new(Primary6YInCieXyy) },
		Primary6IntensityAttr:                 func() zcl.Attr { return new(Primary6Intensity) },
		WhitePointXInCieXyyAttr:               func() zcl.Attr { return new(WhitePointXInCieXyy) },
		WhitePointYInCieXyyAttr:               func() zcl.Attr { return new(WhitePointYInCieXyy) },
		ColorPointRedXInCieXyyAttr:            func() zcl.Attr { return new(ColorPointRedXInCieXyy) },
		ColorPointRedYInCieXyyAttr:            func() zcl.Attr { return new(ColorPointRedYInCieXyy) },
		ColorPointRedIntensityAttr:            func() zcl.Attr { return new(ColorPointRedIntensity) },
		ColorPointGreenXInCieXyyAttr:          func() zcl.Attr { return new(ColorPointGreenXInCieXyy) },
		ColorPointGreenYInCieXyyAttr:          func() zcl.Attr { return new(ColorPointGreenYInCieXyy) },
		ColorPointGreenIntensityAttr:          func() zcl.Attr { return new(ColorPointGreenIntensity) },
		ColorPointBlueXInCieXyyAttr:           func() zcl.Attr { return new(ColorPointBlueXInCieXyy) },
		ColorPointBlueYInCieXyyAttr:           func() zcl.Attr { return new(ColorPointBlueYInCieXyy) },
		ColorPointBlueIntensityAttr:           func() zcl.Attr { return new(ColorPointBlueIntensity) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr: []zcl.AttrID{
		CurrentXInCieXyyAttr,
		CurrentYInCieXyyAttr,
		EnhancedCurrentHueAttr,
		CurrentSaturationAttr,
		ColorLoopActiveAttr,
		ColorLoopDirectionAttr,
		ColorLoopTimeAttr,
	},
}

type MoveToHue struct {
	Hue       zcl.Zu8
	Direction zcl.Zenum8
	// The transition time in 1/10ths of a second.
	TransitionTime zcl.Zu16
}

const MoveToHueCommand zcl.CommandID = 0

func (v *MoveToHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.Hue,
		&v.Direction,
		&v.TransitionTime,
	}
}

func (v MoveToHue) ID() zcl.CommandID {
	return MoveToHueCommand
}

func (v MoveToHue) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v MoveToHue) MnfCode() []byte {
	return []byte{}
}

func (v MoveToHue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Hue.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Direction.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MoveToHue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

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

type MoveHue struct {
	MoveMode zcl.Zenum8
	Rate     zcl.Zu8
}

const MoveHueCommand zcl.CommandID = 1

func (v *MoveHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.MoveMode,
		&v.Rate,
	}
}

func (v MoveHue) ID() zcl.CommandID {
	return MoveHueCommand
}

func (v MoveHue) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v MoveHue) MnfCode() []byte {
	return []byte{}
}

func (v MoveHue) MarshalZcl() ([]byte, error) {
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

func (v *MoveHue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.MoveMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type StepHue struct {
	StepMode zcl.Zenum8
	StepSize zcl.Zu8
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu8
}

const StepHueCommand zcl.CommandID = 2

func (v *StepHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepMode,
		&v.StepSize,
		&v.TransitionTime,
	}
}

func (v StepHue) ID() zcl.CommandID {
	return StepHueCommand
}

func (v StepHue) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v StepHue) MnfCode() []byte {
	return []byte{}
}

func (v StepHue) MarshalZcl() ([]byte, error) {
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

func (v *StepHue) UnmarshalZcl(b []byte) ([]byte, error) {
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

type MoveToSaturation struct {
	Saturation zcl.Zu8
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu16
}

const MoveToSaturationCommand zcl.CommandID = 3

func (v *MoveToSaturation) Values() []zcl.Val {
	return []zcl.Val{
		&v.Saturation,
		&v.TransitionTime,
	}
}

func (v MoveToSaturation) ID() zcl.CommandID {
	return MoveToSaturationCommand
}

func (v MoveToSaturation) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v MoveToSaturation) MnfCode() []byte {
	return []byte{}
}

func (v MoveToSaturation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Saturation.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MoveToSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.Saturation).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type MoveSaturation struct {
	MoveMode zcl.Zenum8
	// The steps per second.
	Rate zcl.Zu8
}

const MoveSaturationCommand zcl.CommandID = 4

func (v *MoveSaturation) Values() []zcl.Val {
	return []zcl.Val{
		&v.MoveMode,
		&v.Rate,
	}
}

func (v MoveSaturation) ID() zcl.CommandID {
	return MoveSaturationCommand
}

func (v MoveSaturation) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v MoveSaturation) MnfCode() []byte {
	return []byte{}
}

func (v MoveSaturation) MarshalZcl() ([]byte, error) {
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

func (v *MoveSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.MoveMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type StepSaturation struct {
	StepMode zcl.Zenum8
	StepSize zcl.Zu8
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu8
}

const StepSaturationCommand zcl.CommandID = 5

func (v *StepSaturation) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepMode,
		&v.StepSize,
		&v.TransitionTime,
	}
}

func (v StepSaturation) ID() zcl.CommandID {
	return StepSaturationCommand
}

func (v StepSaturation) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v StepSaturation) MnfCode() []byte {
	return []byte{}
}

func (v StepSaturation) MarshalZcl() ([]byte, error) {
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

func (v *StepSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
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

type MoveToHueAndSaturation struct {
	Hue        zcl.Zu8
	Saturation zcl.Zu8
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu16
}

const MoveToHueAndSaturationCommand zcl.CommandID = 6

func (v *MoveToHueAndSaturation) Values() []zcl.Val {
	return []zcl.Val{
		&v.Hue,
		&v.Saturation,
		&v.TransitionTime,
	}
}

func (v MoveToHueAndSaturation) ID() zcl.CommandID {
	return MoveToHueAndSaturationCommand
}

func (v MoveToHueAndSaturation) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v MoveToHueAndSaturation) MnfCode() []byte {
	return []byte{}
}

func (v MoveToHueAndSaturation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.Hue.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Saturation.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MoveToHueAndSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

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

type MoveToColor struct {
	ColorX zcl.Zu16
	ColorY zcl.Zu16
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu16
}

const MoveToColorCommand zcl.CommandID = 7

func (v *MoveToColor) Values() []zcl.Val {
	return []zcl.Val{
		&v.ColorX,
		&v.ColorY,
		&v.TransitionTime,
	}
}

func (v MoveToColor) ID() zcl.CommandID {
	return MoveToColorCommand
}

func (v MoveToColor) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v MoveToColor) MnfCode() []byte {
	return []byte{}
}

func (v MoveToColor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.ColorX.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ColorY.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MoveToColor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

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

type MoveColor struct {
	// The steps per second.
	RateX zcl.Zs16
	// The steps per second.
	RateY zcl.Zs16
}

const MoveColorCommand zcl.CommandID = 8

func (v *MoveColor) Values() []zcl.Val {
	return []zcl.Val{
		&v.RateX,
		&v.RateY,
	}
}

func (v MoveColor) ID() zcl.CommandID {
	return MoveColorCommand
}

func (v MoveColor) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v MoveColor) MnfCode() []byte {
	return []byte{}
}

func (v MoveColor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.RateX.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.RateY.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MoveColor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.RateX).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.RateY).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type StepColor struct {
	StepX zcl.Zs16
	StepY zcl.Zs16
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu16
}

const StepColorCommand zcl.CommandID = 9

func (v *StepColor) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepX,
		&v.StepY,
		&v.TransitionTime,
	}
}

func (v StepColor) ID() zcl.CommandID {
	return StepColorCommand
}

func (v StepColor) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v StepColor) MnfCode() []byte {
	return []byte{}
}

func (v StepColor) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.StepX.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StepY.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *StepColor) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

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

type MoveToColorTemperature struct {
	ColorTemperature zcl.Zu16
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu16
}

const MoveToColorTemperatureCommand zcl.CommandID = 10

func (v *MoveToColorTemperature) Values() []zcl.Val {
	return []zcl.Val{
		&v.ColorTemperature,
		&v.TransitionTime,
	}
}

func (v MoveToColorTemperature) ID() zcl.CommandID {
	return MoveToColorTemperatureCommand
}

func (v MoveToColorTemperature) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v MoveToColorTemperature) MnfCode() []byte {
	return []byte{}
}

func (v MoveToColorTemperature) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.ColorTemperature.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MoveToColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.ColorTemperature).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.TransitionTime).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type EnhancedMoveToHue struct {
	EnhancedHue zcl.Zu16
	Direction   zcl.Zenum8
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu16
}

const EnhancedMoveToHueCommand zcl.CommandID = 64

func (v *EnhancedMoveToHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.EnhancedHue,
		&v.Direction,
		&v.TransitionTime,
	}
}

func (v EnhancedMoveToHue) ID() zcl.CommandID {
	return EnhancedMoveToHueCommand
}

func (v EnhancedMoveToHue) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v EnhancedMoveToHue) MnfCode() []byte {
	return []byte{}
}

func (v EnhancedMoveToHue) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.EnhancedHue.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Direction.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *EnhancedMoveToHue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

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

type EnhancedMoveHue struct {
	MoveMode zcl.Zenum8
	// Steps per second.
	Rate zcl.Zu16
}

const EnhancedMoveHueCommand zcl.CommandID = 65

func (v *EnhancedMoveHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.MoveMode,
		&v.Rate,
	}
}

func (v EnhancedMoveHue) ID() zcl.CommandID {
	return EnhancedMoveHueCommand
}

func (v EnhancedMoveHue) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v EnhancedMoveHue) MnfCode() []byte {
	return []byte{}
}

func (v EnhancedMoveHue) MarshalZcl() ([]byte, error) {
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

func (v *EnhancedMoveHue) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.MoveMode).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Rate).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

type EnhancedStepHue struct {
	StepMode zcl.Zenum8
	StepSize zcl.Zu16
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu16
}

const EnhancedStepHueCommand zcl.CommandID = 66

func (v *EnhancedStepHue) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepMode,
		&v.StepSize,
		&v.TransitionTime,
	}
}

func (v EnhancedStepHue) ID() zcl.CommandID {
	return EnhancedStepHueCommand
}

func (v EnhancedStepHue) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v EnhancedStepHue) MnfCode() []byte {
	return []byte{}
}

func (v EnhancedStepHue) MarshalZcl() ([]byte, error) {
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

func (v *EnhancedStepHue) UnmarshalZcl(b []byte) ([]byte, error) {
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

type EnhancedMoveToHueAndSaturation struct {
	EnhancedHue zcl.Zu16
	Saturation  zcl.Zu8
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu16
}

const EnhancedMoveToHueAndSaturationCommand zcl.CommandID = 67

func (v *EnhancedMoveToHueAndSaturation) Values() []zcl.Val {
	return []zcl.Val{
		&v.EnhancedHue,
		&v.Saturation,
		&v.TransitionTime,
	}
}

func (v EnhancedMoveToHueAndSaturation) ID() zcl.CommandID {
	return EnhancedMoveToHueAndSaturationCommand
}

func (v EnhancedMoveToHueAndSaturation) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v EnhancedMoveToHueAndSaturation) MnfCode() []byte {
	return []byte{}
}

func (v EnhancedMoveToHueAndSaturation) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.EnhancedHue.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Saturation.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.TransitionTime.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *EnhancedMoveToHueAndSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

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

type ColorLoopSet struct {
	UpdateFlags zcl.Zbmp8
	Action      zcl.Zenum8
	Direction   zcl.Zenum8
	// Time in seconds used for a whole color loop.
	Time     zcl.Zu16
	StartHue zcl.Zu16
}

const ColorLoopSetCommand zcl.CommandID = 68

func (v *ColorLoopSet) Values() []zcl.Val {
	return []zcl.Val{
		&v.UpdateFlags,
		&v.Action,
		&v.Direction,
		&v.Time,
		&v.StartHue,
	}
}

func (v ColorLoopSet) ID() zcl.CommandID {
	return ColorLoopSetCommand
}

func (v ColorLoopSet) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v ColorLoopSet) MnfCode() []byte {
	return []byte{}
}

func (v ColorLoopSet) MarshalZcl() ([]byte, error) {
	var data []byte
	var tmp []byte
	var err error

	if tmp, err = v.UpdateFlags.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Action.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Direction.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.Time.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.StartHue.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *ColorLoopSet) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

	if b, err = (&v.UpdateFlags).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Action).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Direction).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.Time).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.StartHue).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

// Stops move to and step commands. It has no effect on a active color loop.
type StopMoveStep struct {
}

const StopMoveStepCommand zcl.CommandID = 71

func (v *StopMoveStep) Values() []zcl.Val {
	return []zcl.Val{}
}

func (v StopMoveStep) ID() zcl.CommandID {
	return StopMoveStepCommand
}

func (v StopMoveStep) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v StopMoveStep) MnfCode() []byte {
	return []byte{}
}

func (v StopMoveStep) MarshalZcl() ([]byte, error) {
	return nil, nil
}

func (v *StopMoveStep) UnmarshalZcl(b []byte) ([]byte, error) {
	return b, nil
}

type MoveColorTemperature struct {
	MoveMode zcl.Zenum8
	// Steps per second.
	Rate zcl.Zu16
	// Specifies a lower bound on the color temperature for the current move operation.
	ColorTemperatureMin zcl.Zu16
	// Specifies a upper bound on the color temperature for the current move operation.
	ColorTemperatureMax zcl.Zu16
}

const MoveColorTemperatureCommand zcl.CommandID = 75

func (v *MoveColorTemperature) Values() []zcl.Val {
	return []zcl.Val{
		&v.MoveMode,
		&v.Rate,
		&v.ColorTemperatureMin,
		&v.ColorTemperatureMax,
	}
}

func (v MoveColorTemperature) ID() zcl.CommandID {
	return MoveColorTemperatureCommand
}

func (v MoveColorTemperature) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v MoveColorTemperature) MnfCode() []byte {
	return []byte{}
}

func (v MoveColorTemperature) MarshalZcl() ([]byte, error) {
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

	if tmp, err = v.ColorTemperatureMin.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ColorTemperatureMax.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *MoveColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	var err error

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

type StepColorTemperature struct {
	StepMode zcl.Zenum8
	StepSize zcl.Zu16
	// The transitiontime in 1/10 seconds.
	TransitionTime zcl.Zu16
	// Specifies a lower bound on the color temperature for the current step operation.
	ColorTemperatureMinimumMireds zcl.Zu16
	// Specifies a upper bound on the color temperature for the current step operation.
	ColorTemperatureMaximumMireds zcl.Zu16
}

const StepColorTemperatureCommand zcl.CommandID = 76

func (v *StepColorTemperature) Values() []zcl.Val {
	return []zcl.Val{
		&v.StepMode,
		&v.StepSize,
		&v.TransitionTime,
		&v.ColorTemperatureMinimumMireds,
		&v.ColorTemperatureMaximumMireds,
	}
}

func (v StepColorTemperature) ID() zcl.CommandID {
	return StepColorTemperatureCommand
}

func (v StepColorTemperature) Cluster() zcl.ClusterID {
	return ColorControlID
}

func (v StepColorTemperature) MnfCode() []byte {
	return []byte{}
}

func (v StepColorTemperature) MarshalZcl() ([]byte, error) {
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

	if tmp, err = v.ColorTemperatureMinimumMireds.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	if tmp, err = v.ColorTemperatureMaximumMireds.MarshalZcl(); err != nil {
		return nil, err
	}
	data = append(data, tmp...)

	return data, nil
}

func (v *StepColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
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

	if b, err = (&v.ColorTemperatureMinimumMireds).UnmarshalZcl(b); err != nil {
		return b, err
	}

	if b, err = (&v.ColorTemperatureMaximumMireds).UnmarshalZcl(b); err != nil {
		return b, err
	}

	return b, nil
}

const CurrentHueAttr zcl.AttrID = 0

type CurrentHue zcl.Zu8

func (a CurrentHue) ID() zcl.AttrID         { return CurrentHueAttr }
func (a CurrentHue) Cluster() zcl.ClusterID { return ColorControlID }
func (a *CurrentHue) Value() *CurrentHue    { return a }
func (a CurrentHue) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *CurrentHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentHue(*nt)
	return br, err
}

func (a CurrentHue) Readable() bool   { return true }
func (a CurrentHue) Writable() bool   { return false }
func (a CurrentHue) Reportable() bool { return true }
func (a CurrentHue) SceneIndex() int  { return -1 }

func (a CurrentHue) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const CurrentSaturationAttr zcl.AttrID = 1

type CurrentSaturation zcl.Zu8

func (a CurrentSaturation) ID() zcl.AttrID             { return CurrentSaturationAttr }
func (a CurrentSaturation) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *CurrentSaturation) Value() *CurrentSaturation { return a }
func (a CurrentSaturation) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *CurrentSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentSaturation(*nt)
	return br, err
}

func (a CurrentSaturation) Readable() bool   { return true }
func (a CurrentSaturation) Writable() bool   { return false }
func (a CurrentSaturation) Reportable() bool { return true }
func (a CurrentSaturation) SceneIndex() int  { return 4 }

func (a CurrentSaturation) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const RemainingTimeAttr zcl.AttrID = 2

type RemainingTime zcl.Zu16

func (a RemainingTime) ID() zcl.AttrID         { return RemainingTimeAttr }
func (a RemainingTime) Cluster() zcl.ClusterID { return ColorControlID }
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

const CurrentXInCieXyyAttr zcl.AttrID = 3

type CurrentXInCieXyy zcl.Zu16

func (a CurrentXInCieXyy) ID() zcl.AttrID            { return CurrentXInCieXyyAttr }
func (a CurrentXInCieXyy) Cluster() zcl.ClusterID    { return ColorControlID }
func (a *CurrentXInCieXyy) Value() *CurrentXInCieXyy { return a }
func (a CurrentXInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *CurrentXInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentXInCieXyy(*nt)
	return br, err
}

func (a CurrentXInCieXyy) Readable() bool   { return true }
func (a CurrentXInCieXyy) Writable() bool   { return false }
func (a CurrentXInCieXyy) Reportable() bool { return true }
func (a CurrentXInCieXyy) SceneIndex() int  { return 1 }

func (a CurrentXInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const CurrentYInCieXyyAttr zcl.AttrID = 4

type CurrentYInCieXyy zcl.Zu16

func (a CurrentYInCieXyy) ID() zcl.AttrID            { return CurrentYInCieXyyAttr }
func (a CurrentYInCieXyy) Cluster() zcl.ClusterID    { return ColorControlID }
func (a *CurrentYInCieXyy) Value() *CurrentYInCieXyy { return a }
func (a CurrentYInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *CurrentYInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentYInCieXyy(*nt)
	return br, err
}

func (a CurrentYInCieXyy) Readable() bool   { return true }
func (a CurrentYInCieXyy) Writable() bool   { return false }
func (a CurrentYInCieXyy) Reportable() bool { return true }
func (a CurrentYInCieXyy) SceneIndex() int  { return 2 }

func (a CurrentYInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const DriftCompensationAttr zcl.AttrID = 5

type DriftCompensation zcl.Zenum8

func (a DriftCompensation) ID() zcl.AttrID             { return DriftCompensationAttr }
func (a DriftCompensation) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *DriftCompensation) Value() *DriftCompensation { return a }
func (a DriftCompensation) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *DriftCompensation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = DriftCompensation(*nt)
	return br, err
}

func (a DriftCompensation) Readable() bool   { return true }
func (a DriftCompensation) Writable() bool   { return false }
func (a DriftCompensation) Reportable() bool { return false }
func (a DriftCompensation) SceneIndex() int  { return -1 }

func (a DriftCompensation) String() string {
	switch a {
	case 0x00:
		return "None"
	case 0x01:
		return "Other / Unknown"
	case 0x02:
		return "Temperature monitoring"
	case 0x03:
		return "Optical luminance monitoring and feedback"
	case 0x04:
		return "Optical color monitoring and feedback"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsNone checks if DriftCompensation equals the value for None (0x00)
func (a DriftCompensation) IsNone() bool { return a == 0x00 }

// SetNone sets DriftCompensation to None (0x00)
func (a *DriftCompensation) SetNone() { *a = 0x00 }

// IsOtherUnknown checks if DriftCompensation equals the value for Other / Unknown (0x01)
func (a DriftCompensation) IsOtherUnknown() bool { return a == 0x01 }

// SetOtherUnknown sets DriftCompensation to Other / Unknown (0x01)
func (a *DriftCompensation) SetOtherUnknown() { *a = 0x01 }

// IsTemperatureMonitoring checks if DriftCompensation equals the value for Temperature monitoring (0x02)
func (a DriftCompensation) IsTemperatureMonitoring() bool { return a == 0x02 }

// SetTemperatureMonitoring sets DriftCompensation to Temperature monitoring (0x02)
func (a *DriftCompensation) SetTemperatureMonitoring() { *a = 0x02 }

// IsOpticalLuminanceMonitoringAndFeedback checks if DriftCompensation equals the value for Optical luminance monitoring and feedback (0x03)
func (a DriftCompensation) IsOpticalLuminanceMonitoringAndFeedback() bool { return a == 0x03 }

// SetOpticalLuminanceMonitoringAndFeedback sets DriftCompensation to Optical luminance monitoring and feedback (0x03)
func (a *DriftCompensation) SetOpticalLuminanceMonitoringAndFeedback() { *a = 0x03 }

// IsOpticalColorMonitoringAndFeedback checks if DriftCompensation equals the value for Optical color monitoring and feedback (0x04)
func (a DriftCompensation) IsOpticalColorMonitoringAndFeedback() bool { return a == 0x04 }

// SetOpticalColorMonitoringAndFeedback sets DriftCompensation to Optical color monitoring and feedback (0x04)
func (a *DriftCompensation) SetOpticalColorMonitoringAndFeedback() { *a = 0x04 }

const CompensationTextAttr zcl.AttrID = 6

type CompensationText zcl.Zcstring

func (a CompensationText) ID() zcl.AttrID            { return CompensationTextAttr }
func (a CompensationText) Cluster() zcl.ClusterID    { return ColorControlID }
func (a *CompensationText) Value() *CompensationText { return a }
func (a CompensationText) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *CompensationText) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = CompensationText(*nt)
	return br, err
}

func (a CompensationText) Readable() bool   { return true }
func (a CompensationText) Writable() bool   { return false }
func (a CompensationText) Reportable() bool { return false }
func (a CompensationText) SceneIndex() int  { return -1 }

func (a CompensationText) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

const ColorTemperatureMiredsAttr zcl.AttrID = 7

type ColorTemperatureMireds zcl.Zu16

func (a ColorTemperatureMireds) ID() zcl.AttrID                  { return ColorTemperatureMiredsAttr }
func (a ColorTemperatureMireds) Cluster() zcl.ClusterID          { return ColorControlID }
func (a *ColorTemperatureMireds) Value() *ColorTemperatureMireds { return a }
func (a ColorTemperatureMireds) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorTemperatureMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperatureMireds(*nt)
	return br, err
}

func (a ColorTemperatureMireds) Readable() bool   { return true }
func (a ColorTemperatureMireds) Writable() bool   { return false }
func (a ColorTemperatureMireds) Reportable() bool { return true }
func (a ColorTemperatureMireds) SceneIndex() int  { return -1 }

func (a ColorTemperatureMireds) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorModeAttr zcl.AttrID = 8

type ColorMode zcl.Zenum8

func (a ColorMode) ID() zcl.AttrID         { return ColorModeAttr }
func (a ColorMode) Cluster() zcl.ClusterID { return ColorControlID }
func (a *ColorMode) Value() *ColorMode     { return a }
func (a ColorMode) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *ColorMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorMode(*nt)
	return br, err
}

func (a ColorMode) Readable() bool   { return true }
func (a ColorMode) Writable() bool   { return false }
func (a ColorMode) Reportable() bool { return false }
func (a ColorMode) SceneIndex() int  { return -1 }

func (a ColorMode) String() string {
	switch a {
	case 0x00:
		return "Current hue and current saturation"
	case 0x01:
		return "Current x and current y in CIE xyY"
	case 0x02:
		return "Color temperature"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsCurrentHueAndCurrentSaturation checks if ColorMode equals the value for Current hue and current saturation (0x00)
func (a ColorMode) IsCurrentHueAndCurrentSaturation() bool { return a == 0x00 }

// SetCurrentHueAndCurrentSaturation sets ColorMode to Current hue and current saturation (0x00)
func (a *ColorMode) SetCurrentHueAndCurrentSaturation() { *a = 0x00 }

// IsCurrentXAndCurrentYInCieXyy checks if ColorMode equals the value for Current x and current y in CIE xyY (0x01)
func (a ColorMode) IsCurrentXAndCurrentYInCieXyy() bool { return a == 0x01 }

// SetCurrentXAndCurrentYInCieXyy sets ColorMode to Current x and current y in CIE xyY (0x01)
func (a *ColorMode) SetCurrentXAndCurrentYInCieXyy() { *a = 0x01 }

// IsColorTemperature checks if ColorMode equals the value for Color temperature (0x02)
func (a ColorMode) IsColorTemperature() bool { return a == 0x02 }

// SetColorTemperature sets ColorMode to Color temperature (0x02)
func (a *ColorMode) SetColorTemperature() { *a = 0x02 }

const EnhancedCurrentHueAttr zcl.AttrID = 16384

type EnhancedCurrentHue zcl.Zu16

func (a EnhancedCurrentHue) ID() zcl.AttrID              { return EnhancedCurrentHueAttr }
func (a EnhancedCurrentHue) Cluster() zcl.ClusterID      { return ColorControlID }
func (a *EnhancedCurrentHue) Value() *EnhancedCurrentHue { return a }
func (a EnhancedCurrentHue) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *EnhancedCurrentHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = EnhancedCurrentHue(*nt)
	return br, err
}

func (a EnhancedCurrentHue) Readable() bool   { return true }
func (a EnhancedCurrentHue) Writable() bool   { return false }
func (a EnhancedCurrentHue) Reportable() bool { return false }
func (a EnhancedCurrentHue) SceneIndex() int  { return 3 }

func (a EnhancedCurrentHue) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const EnhancedColorModeAttr zcl.AttrID = 16385

type EnhancedColorMode zcl.Zenum8

func (a EnhancedColorMode) ID() zcl.AttrID             { return EnhancedColorModeAttr }
func (a EnhancedColorMode) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *EnhancedColorMode) Value() *EnhancedColorMode { return a }
func (a EnhancedColorMode) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *EnhancedColorMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = EnhancedColorMode(*nt)
	return br, err
}

func (a EnhancedColorMode) Readable() bool   { return true }
func (a EnhancedColorMode) Writable() bool   { return false }
func (a EnhancedColorMode) Reportable() bool { return false }
func (a EnhancedColorMode) SceneIndex() int  { return -1 }

func (a EnhancedColorMode) String() string {
	switch a {
	case 0x00:
		return "Current hue and current saturation"
	case 0x01:
		return "Current x and current y in CIE xyY"
	case 0x02:
		return "Color temperature"
	case 0x03:
		return "Enhanced current hue and current saturation"
	}

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsCurrentHueAndCurrentSaturation checks if EnhancedColorMode equals the value for Current hue and current saturation (0x00)
func (a EnhancedColorMode) IsCurrentHueAndCurrentSaturation() bool { return a == 0x00 }

// SetCurrentHueAndCurrentSaturation sets EnhancedColorMode to Current hue and current saturation (0x00)
func (a *EnhancedColorMode) SetCurrentHueAndCurrentSaturation() { *a = 0x00 }

// IsCurrentXAndCurrentYInCieXyy checks if EnhancedColorMode equals the value for Current x and current y in CIE xyY (0x01)
func (a EnhancedColorMode) IsCurrentXAndCurrentYInCieXyy() bool { return a == 0x01 }

// SetCurrentXAndCurrentYInCieXyy sets EnhancedColorMode to Current x and current y in CIE xyY (0x01)
func (a *EnhancedColorMode) SetCurrentXAndCurrentYInCieXyy() { *a = 0x01 }

// IsColorTemperature checks if EnhancedColorMode equals the value for Color temperature (0x02)
func (a EnhancedColorMode) IsColorTemperature() bool { return a == 0x02 }

// SetColorTemperature sets EnhancedColorMode to Color temperature (0x02)
func (a *EnhancedColorMode) SetColorTemperature() { *a = 0x02 }

// IsEnhancedCurrentHueAndCurrentSaturation checks if EnhancedColorMode equals the value for Enhanced current hue and current saturation (0x03)
func (a EnhancedColorMode) IsEnhancedCurrentHueAndCurrentSaturation() bool { return a == 0x03 }

// SetEnhancedCurrentHueAndCurrentSaturation sets EnhancedColorMode to Enhanced current hue and current saturation (0x03)
func (a *EnhancedColorMode) SetEnhancedCurrentHueAndCurrentSaturation() { *a = 0x03 }

const ColorLoopActiveAttr zcl.AttrID = 16386

type ColorLoopActive zcl.Zu8

func (a ColorLoopActive) ID() zcl.AttrID           { return ColorLoopActiveAttr }
func (a ColorLoopActive) Cluster() zcl.ClusterID   { return ColorControlID }
func (a *ColorLoopActive) Value() *ColorLoopActive { return a }
func (a ColorLoopActive) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ColorLoopActive) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorLoopActive(*nt)
	return br, err
}

func (a ColorLoopActive) Readable() bool   { return true }
func (a ColorLoopActive) Writable() bool   { return false }
func (a ColorLoopActive) Reportable() bool { return false }
func (a ColorLoopActive) SceneIndex() int  { return 5 }

func (a ColorLoopActive) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const ColorLoopDirectionAttr zcl.AttrID = 16387

type ColorLoopDirection zcl.Zu8

func (a ColorLoopDirection) ID() zcl.AttrID              { return ColorLoopDirectionAttr }
func (a ColorLoopDirection) Cluster() zcl.ClusterID      { return ColorControlID }
func (a *ColorLoopDirection) Value() *ColorLoopDirection { return a }
func (a ColorLoopDirection) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ColorLoopDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorLoopDirection(*nt)
	return br, err
}

func (a ColorLoopDirection) Readable() bool   { return true }
func (a ColorLoopDirection) Writable() bool   { return false }
func (a ColorLoopDirection) Reportable() bool { return false }
func (a ColorLoopDirection) SceneIndex() int  { return 6 }

func (a ColorLoopDirection) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const ColorLoopTimeAttr zcl.AttrID = 16388

type ColorLoopTime zcl.Zu16

func (a ColorLoopTime) ID() zcl.AttrID         { return ColorLoopTimeAttr }
func (a ColorLoopTime) Cluster() zcl.ClusterID { return ColorControlID }
func (a *ColorLoopTime) Value() *ColorLoopTime { return a }
func (a ColorLoopTime) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorLoopTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorLoopTime(*nt)
	return br, err
}

func (a ColorLoopTime) Readable() bool   { return true }
func (a ColorLoopTime) Writable() bool   { return false }
func (a ColorLoopTime) Reportable() bool { return false }
func (a ColorLoopTime) SceneIndex() int  { return 7 }

func (a ColorLoopTime) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorLoopStartEnhancedHueAttr zcl.AttrID = 16389

type ColorLoopStartEnhancedHue zcl.Zu16

func (a ColorLoopStartEnhancedHue) ID() zcl.AttrID                     { return ColorLoopStartEnhancedHueAttr }
func (a ColorLoopStartEnhancedHue) Cluster() zcl.ClusterID             { return ColorControlID }
func (a *ColorLoopStartEnhancedHue) Value() *ColorLoopStartEnhancedHue { return a }
func (a ColorLoopStartEnhancedHue) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorLoopStartEnhancedHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorLoopStartEnhancedHue(*nt)
	return br, err
}

func (a ColorLoopStartEnhancedHue) Readable() bool   { return true }
func (a ColorLoopStartEnhancedHue) Writable() bool   { return false }
func (a ColorLoopStartEnhancedHue) Reportable() bool { return false }
func (a ColorLoopStartEnhancedHue) SceneIndex() int  { return -1 }

func (a ColorLoopStartEnhancedHue) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorLoopStoredEnhancedHueAttr zcl.AttrID = 16390

type ColorLoopStoredEnhancedHue zcl.Zu16

func (a ColorLoopStoredEnhancedHue) ID() zcl.AttrID                      { return ColorLoopStoredEnhancedHueAttr }
func (a ColorLoopStoredEnhancedHue) Cluster() zcl.ClusterID              { return ColorControlID }
func (a *ColorLoopStoredEnhancedHue) Value() *ColorLoopStoredEnhancedHue { return a }
func (a ColorLoopStoredEnhancedHue) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorLoopStoredEnhancedHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorLoopStoredEnhancedHue(*nt)
	return br, err
}

func (a ColorLoopStoredEnhancedHue) Readable() bool   { return true }
func (a ColorLoopStoredEnhancedHue) Writable() bool   { return false }
func (a ColorLoopStoredEnhancedHue) Reportable() bool { return false }
func (a ColorLoopStoredEnhancedHue) SceneIndex() int  { return -1 }

func (a ColorLoopStoredEnhancedHue) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorCapabilitiesAttr zcl.AttrID = 16394

type ColorCapabilities zcl.Zbmp16

func (a ColorCapabilities) ID() zcl.AttrID             { return ColorCapabilitiesAttr }
func (a ColorCapabilities) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *ColorCapabilities) Value() *ColorCapabilities { return a }
func (a ColorCapabilities) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp16(a).MarshalZcl()
}
func (a *ColorCapabilities) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorCapabilities(*nt)
	return br, err
}

func (a ColorCapabilities) Readable() bool   { return true }
func (a ColorCapabilities) Writable() bool   { return false }
func (a ColorCapabilities) Reportable() bool { return false }
func (a ColorCapabilities) SceneIndex() int  { return -1 }

func (a ColorCapabilities) String() string {

	var bstr []string
	if a.IsHueSaturation() {
		bstr = append(bstr, "Hue saturation")
	}
	if a.IsEnhancedHueSaturation() {
		bstr = append(bstr, "Enhanced Hue saturation")
	}
	if a.IsColorLoop() {
		bstr = append(bstr, "Color loop")
	}
	if a.IsCie1931Xy() {
		bstr = append(bstr, "CIE 1931 XY")
	}
	if a.IsColorTemperature() {
		bstr = append(bstr, "Color temperature")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a ColorCapabilities) IsHueSaturation() bool {
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *ColorCapabilities) SetHueSaturation(b bool) {
	*a = ColorCapabilities(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a ColorCapabilities) IsEnhancedHueSaturation() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *ColorCapabilities) SetEnhancedHueSaturation(b bool) {
	*a = ColorCapabilities(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a ColorCapabilities) IsColorLoop() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *ColorCapabilities) SetColorLoop(b bool) {
	*a = ColorCapabilities(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a ColorCapabilities) IsCie1931Xy() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *ColorCapabilities) SetCie1931Xy(b bool) {
	*a = ColorCapabilities(zcl.BitmapSet([]byte(*a), 3, b))
}

func (a ColorCapabilities) IsColorTemperature() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *ColorCapabilities) SetColorTemperature(b bool) {
	*a = ColorCapabilities(zcl.BitmapSet([]byte(*a), 4, b))
}

const ColorTemperaturePhysicalMinMiredsAttr zcl.AttrID = 16395

type ColorTemperaturePhysicalMinMireds zcl.Zu16

func (a ColorTemperaturePhysicalMinMireds) ID() zcl.AttrID {
	return ColorTemperaturePhysicalMinMiredsAttr
}
func (a ColorTemperaturePhysicalMinMireds) Cluster() zcl.ClusterID                     { return ColorControlID }
func (a *ColorTemperaturePhysicalMinMireds) Value() *ColorTemperaturePhysicalMinMireds { return a }
func (a ColorTemperaturePhysicalMinMireds) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorTemperaturePhysicalMinMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperaturePhysicalMinMireds(*nt)
	return br, err
}

func (a ColorTemperaturePhysicalMinMireds) Readable() bool   { return true }
func (a ColorTemperaturePhysicalMinMireds) Writable() bool   { return false }
func (a ColorTemperaturePhysicalMinMireds) Reportable() bool { return false }
func (a ColorTemperaturePhysicalMinMireds) SceneIndex() int  { return -1 }

func (a ColorTemperaturePhysicalMinMireds) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorTemperaturePhysicalMaxMiredsAttr zcl.AttrID = 16396

type ColorTemperaturePhysicalMaxMireds zcl.Zu16

func (a ColorTemperaturePhysicalMaxMireds) ID() zcl.AttrID {
	return ColorTemperaturePhysicalMaxMiredsAttr
}
func (a ColorTemperaturePhysicalMaxMireds) Cluster() zcl.ClusterID                     { return ColorControlID }
func (a *ColorTemperaturePhysicalMaxMireds) Value() *ColorTemperaturePhysicalMaxMireds { return a }
func (a ColorTemperaturePhysicalMaxMireds) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorTemperaturePhysicalMaxMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperaturePhysicalMaxMireds(*nt)
	return br, err
}

func (a ColorTemperaturePhysicalMaxMireds) Readable() bool   { return true }
func (a ColorTemperaturePhysicalMaxMireds) Writable() bool   { return false }
func (a ColorTemperaturePhysicalMaxMireds) Reportable() bool { return false }
func (a ColorTemperaturePhysicalMaxMireds) SceneIndex() int  { return -1 }

func (a ColorTemperaturePhysicalMaxMireds) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const PowerOnColorTemperatureAttr zcl.AttrID = 16400

type PowerOnColorTemperature zcl.Zu16

func (a PowerOnColorTemperature) ID() zcl.AttrID                   { return PowerOnColorTemperatureAttr }
func (a PowerOnColorTemperature) Cluster() zcl.ClusterID           { return ColorControlID }
func (a *PowerOnColorTemperature) Value() *PowerOnColorTemperature { return a }
func (a PowerOnColorTemperature) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *PowerOnColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerOnColorTemperature(*nt)
	return br, err
}

func (a PowerOnColorTemperature) Readable() bool   { return true }
func (a PowerOnColorTemperature) Writable() bool   { return true }
func (a PowerOnColorTemperature) Reportable() bool { return false }
func (a PowerOnColorTemperature) SceneIndex() int  { return -1 }

func (a PowerOnColorTemperature) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const NumberOfPrimariesAttr zcl.AttrID = 16

type NumberOfPrimaries zcl.Zu8

func (a NumberOfPrimaries) ID() zcl.AttrID             { return NumberOfPrimariesAttr }
func (a NumberOfPrimaries) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *NumberOfPrimaries) Value() *NumberOfPrimaries { return a }
func (a NumberOfPrimaries) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *NumberOfPrimaries) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfPrimaries(*nt)
	return br, err
}

func (a NumberOfPrimaries) Readable() bool   { return true }
func (a NumberOfPrimaries) Writable() bool   { return false }
func (a NumberOfPrimaries) Reportable() bool { return false }
func (a NumberOfPrimaries) SceneIndex() int  { return -1 }

func (a NumberOfPrimaries) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const Primary1XInCieXyyAttr zcl.AttrID = 17

type Primary1XInCieXyy zcl.Zu16

func (a Primary1XInCieXyy) ID() zcl.AttrID             { return Primary1XInCieXyyAttr }
func (a Primary1XInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary1XInCieXyy) Value() *Primary1XInCieXyy { return a }
func (a Primary1XInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary1XInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary1XInCieXyy(*nt)
	return br, err
}

func (a Primary1XInCieXyy) Readable() bool   { return true }
func (a Primary1XInCieXyy) Writable() bool   { return false }
func (a Primary1XInCieXyy) Reportable() bool { return false }
func (a Primary1XInCieXyy) SceneIndex() int  { return -1 }

func (a Primary1XInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary1YInCieXyyAttr zcl.AttrID = 18

type Primary1YInCieXyy zcl.Zu16

func (a Primary1YInCieXyy) ID() zcl.AttrID             { return Primary1YInCieXyyAttr }
func (a Primary1YInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary1YInCieXyy) Value() *Primary1YInCieXyy { return a }
func (a Primary1YInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary1YInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary1YInCieXyy(*nt)
	return br, err
}

func (a Primary1YInCieXyy) Readable() bool   { return true }
func (a Primary1YInCieXyy) Writable() bool   { return false }
func (a Primary1YInCieXyy) Reportable() bool { return false }
func (a Primary1YInCieXyy) SceneIndex() int  { return -1 }

func (a Primary1YInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary1IntensityAttr zcl.AttrID = 19

type Primary1Intensity zcl.Zu8

func (a Primary1Intensity) ID() zcl.AttrID             { return Primary1IntensityAttr }
func (a Primary1Intensity) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary1Intensity) Value() *Primary1Intensity { return a }
func (a Primary1Intensity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Primary1Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary1Intensity(*nt)
	return br, err
}

func (a Primary1Intensity) Readable() bool   { return true }
func (a Primary1Intensity) Writable() bool   { return false }
func (a Primary1Intensity) Reportable() bool { return false }
func (a Primary1Intensity) SceneIndex() int  { return -1 }

func (a Primary1Intensity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const Primary2XInCieXyyAttr zcl.AttrID = 21

type Primary2XInCieXyy zcl.Zu16

func (a Primary2XInCieXyy) ID() zcl.AttrID             { return Primary2XInCieXyyAttr }
func (a Primary2XInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary2XInCieXyy) Value() *Primary2XInCieXyy { return a }
func (a Primary2XInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary2XInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary2XInCieXyy(*nt)
	return br, err
}

func (a Primary2XInCieXyy) Readable() bool   { return true }
func (a Primary2XInCieXyy) Writable() bool   { return false }
func (a Primary2XInCieXyy) Reportable() bool { return false }
func (a Primary2XInCieXyy) SceneIndex() int  { return -1 }

func (a Primary2XInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary2YInCieXyyAttr zcl.AttrID = 22

type Primary2YInCieXyy zcl.Zu16

func (a Primary2YInCieXyy) ID() zcl.AttrID             { return Primary2YInCieXyyAttr }
func (a Primary2YInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary2YInCieXyy) Value() *Primary2YInCieXyy { return a }
func (a Primary2YInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary2YInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary2YInCieXyy(*nt)
	return br, err
}

func (a Primary2YInCieXyy) Readable() bool   { return true }
func (a Primary2YInCieXyy) Writable() bool   { return false }
func (a Primary2YInCieXyy) Reportable() bool { return false }
func (a Primary2YInCieXyy) SceneIndex() int  { return -1 }

func (a Primary2YInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary2IntensityAttr zcl.AttrID = 23

type Primary2Intensity zcl.Zu8

func (a Primary2Intensity) ID() zcl.AttrID             { return Primary2IntensityAttr }
func (a Primary2Intensity) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary2Intensity) Value() *Primary2Intensity { return a }
func (a Primary2Intensity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Primary2Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary2Intensity(*nt)
	return br, err
}

func (a Primary2Intensity) Readable() bool   { return true }
func (a Primary2Intensity) Writable() bool   { return false }
func (a Primary2Intensity) Reportable() bool { return false }
func (a Primary2Intensity) SceneIndex() int  { return -1 }

func (a Primary2Intensity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const Primary3XInCieXyyAttr zcl.AttrID = 25

type Primary3XInCieXyy zcl.Zu16

func (a Primary3XInCieXyy) ID() zcl.AttrID             { return Primary3XInCieXyyAttr }
func (a Primary3XInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary3XInCieXyy) Value() *Primary3XInCieXyy { return a }
func (a Primary3XInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary3XInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary3XInCieXyy(*nt)
	return br, err
}

func (a Primary3XInCieXyy) Readable() bool   { return true }
func (a Primary3XInCieXyy) Writable() bool   { return false }
func (a Primary3XInCieXyy) Reportable() bool { return false }
func (a Primary3XInCieXyy) SceneIndex() int  { return -1 }

func (a Primary3XInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary3YInCieXyyAttr zcl.AttrID = 26

type Primary3YInCieXyy zcl.Zu16

func (a Primary3YInCieXyy) ID() zcl.AttrID             { return Primary3YInCieXyyAttr }
func (a Primary3YInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary3YInCieXyy) Value() *Primary3YInCieXyy { return a }
func (a Primary3YInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary3YInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary3YInCieXyy(*nt)
	return br, err
}

func (a Primary3YInCieXyy) Readable() bool   { return true }
func (a Primary3YInCieXyy) Writable() bool   { return false }
func (a Primary3YInCieXyy) Reportable() bool { return false }
func (a Primary3YInCieXyy) SceneIndex() int  { return -1 }

func (a Primary3YInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary3IntensityAttr zcl.AttrID = 27

type Primary3Intensity zcl.Zu8

func (a Primary3Intensity) ID() zcl.AttrID             { return Primary3IntensityAttr }
func (a Primary3Intensity) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary3Intensity) Value() *Primary3Intensity { return a }
func (a Primary3Intensity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Primary3Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary3Intensity(*nt)
	return br, err
}

func (a Primary3Intensity) Readable() bool   { return true }
func (a Primary3Intensity) Writable() bool   { return false }
func (a Primary3Intensity) Reportable() bool { return false }
func (a Primary3Intensity) SceneIndex() int  { return -1 }

func (a Primary3Intensity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const Primary4XInCieXyyAttr zcl.AttrID = 32

type Primary4XInCieXyy zcl.Zu16

func (a Primary4XInCieXyy) ID() zcl.AttrID             { return Primary4XInCieXyyAttr }
func (a Primary4XInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary4XInCieXyy) Value() *Primary4XInCieXyy { return a }
func (a Primary4XInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary4XInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary4XInCieXyy(*nt)
	return br, err
}

func (a Primary4XInCieXyy) Readable() bool   { return true }
func (a Primary4XInCieXyy) Writable() bool   { return false }
func (a Primary4XInCieXyy) Reportable() bool { return false }
func (a Primary4XInCieXyy) SceneIndex() int  { return -1 }

func (a Primary4XInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary4YInCieXyyAttr zcl.AttrID = 33

type Primary4YInCieXyy zcl.Zu16

func (a Primary4YInCieXyy) ID() zcl.AttrID             { return Primary4YInCieXyyAttr }
func (a Primary4YInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary4YInCieXyy) Value() *Primary4YInCieXyy { return a }
func (a Primary4YInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary4YInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary4YInCieXyy(*nt)
	return br, err
}

func (a Primary4YInCieXyy) Readable() bool   { return true }
func (a Primary4YInCieXyy) Writable() bool   { return false }
func (a Primary4YInCieXyy) Reportable() bool { return false }
func (a Primary4YInCieXyy) SceneIndex() int  { return -1 }

func (a Primary4YInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary4IntensityAttr zcl.AttrID = 34

type Primary4Intensity zcl.Zu8

func (a Primary4Intensity) ID() zcl.AttrID             { return Primary4IntensityAttr }
func (a Primary4Intensity) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary4Intensity) Value() *Primary4Intensity { return a }
func (a Primary4Intensity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Primary4Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary4Intensity(*nt)
	return br, err
}

func (a Primary4Intensity) Readable() bool   { return true }
func (a Primary4Intensity) Writable() bool   { return false }
func (a Primary4Intensity) Reportable() bool { return false }
func (a Primary4Intensity) SceneIndex() int  { return -1 }

func (a Primary4Intensity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const Primary5XInCieXyyAttr zcl.AttrID = 36

type Primary5XInCieXyy zcl.Zu16

func (a Primary5XInCieXyy) ID() zcl.AttrID             { return Primary5XInCieXyyAttr }
func (a Primary5XInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary5XInCieXyy) Value() *Primary5XInCieXyy { return a }
func (a Primary5XInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary5XInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary5XInCieXyy(*nt)
	return br, err
}

func (a Primary5XInCieXyy) Readable() bool   { return true }
func (a Primary5XInCieXyy) Writable() bool   { return false }
func (a Primary5XInCieXyy) Reportable() bool { return false }
func (a Primary5XInCieXyy) SceneIndex() int  { return -1 }

func (a Primary5XInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary5YInCieXyyAttr zcl.AttrID = 37

type Primary5YInCieXyy zcl.Zu16

func (a Primary5YInCieXyy) ID() zcl.AttrID             { return Primary5YInCieXyyAttr }
func (a Primary5YInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary5YInCieXyy) Value() *Primary5YInCieXyy { return a }
func (a Primary5YInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary5YInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary5YInCieXyy(*nt)
	return br, err
}

func (a Primary5YInCieXyy) Readable() bool   { return true }
func (a Primary5YInCieXyy) Writable() bool   { return false }
func (a Primary5YInCieXyy) Reportable() bool { return false }
func (a Primary5YInCieXyy) SceneIndex() int  { return -1 }

func (a Primary5YInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary5IntensityAttr zcl.AttrID = 38

type Primary5Intensity zcl.Zu8

func (a Primary5Intensity) ID() zcl.AttrID             { return Primary5IntensityAttr }
func (a Primary5Intensity) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary5Intensity) Value() *Primary5Intensity { return a }
func (a Primary5Intensity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Primary5Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary5Intensity(*nt)
	return br, err
}

func (a Primary5Intensity) Readable() bool   { return true }
func (a Primary5Intensity) Writable() bool   { return false }
func (a Primary5Intensity) Reportable() bool { return false }
func (a Primary5Intensity) SceneIndex() int  { return -1 }

func (a Primary5Intensity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const Primary6XInCieXyyAttr zcl.AttrID = 40

type Primary6XInCieXyy zcl.Zu16

func (a Primary6XInCieXyy) ID() zcl.AttrID             { return Primary6XInCieXyyAttr }
func (a Primary6XInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary6XInCieXyy) Value() *Primary6XInCieXyy { return a }
func (a Primary6XInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary6XInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary6XInCieXyy(*nt)
	return br, err
}

func (a Primary6XInCieXyy) Readable() bool   { return true }
func (a Primary6XInCieXyy) Writable() bool   { return false }
func (a Primary6XInCieXyy) Reportable() bool { return false }
func (a Primary6XInCieXyy) SceneIndex() int  { return -1 }

func (a Primary6XInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary6YInCieXyyAttr zcl.AttrID = 41

type Primary6YInCieXyy zcl.Zu16

func (a Primary6YInCieXyy) ID() zcl.AttrID             { return Primary6YInCieXyyAttr }
func (a Primary6YInCieXyy) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary6YInCieXyy) Value() *Primary6YInCieXyy { return a }
func (a Primary6YInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary6YInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary6YInCieXyy(*nt)
	return br, err
}

func (a Primary6YInCieXyy) Readable() bool   { return true }
func (a Primary6YInCieXyy) Writable() bool   { return false }
func (a Primary6YInCieXyy) Reportable() bool { return false }
func (a Primary6YInCieXyy) SceneIndex() int  { return -1 }

func (a Primary6YInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const Primary6IntensityAttr zcl.AttrID = 42

type Primary6Intensity zcl.Zu8

func (a Primary6Intensity) ID() zcl.AttrID             { return Primary6IntensityAttr }
func (a Primary6Intensity) Cluster() zcl.ClusterID     { return ColorControlID }
func (a *Primary6Intensity) Value() *Primary6Intensity { return a }
func (a Primary6Intensity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Primary6Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary6Intensity(*nt)
	return br, err
}

func (a Primary6Intensity) Readable() bool   { return true }
func (a Primary6Intensity) Writable() bool   { return false }
func (a Primary6Intensity) Reportable() bool { return false }
func (a Primary6Intensity) SceneIndex() int  { return -1 }

func (a Primary6Intensity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const WhitePointXInCieXyyAttr zcl.AttrID = 48

type WhitePointXInCieXyy zcl.Zu16

func (a WhitePointXInCieXyy) ID() zcl.AttrID               { return WhitePointXInCieXyyAttr }
func (a WhitePointXInCieXyy) Cluster() zcl.ClusterID       { return ColorControlID }
func (a *WhitePointXInCieXyy) Value() *WhitePointXInCieXyy { return a }
func (a WhitePointXInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *WhitePointXInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = WhitePointXInCieXyy(*nt)
	return br, err
}

func (a WhitePointXInCieXyy) Readable() bool   { return true }
func (a WhitePointXInCieXyy) Writable() bool   { return true }
func (a WhitePointXInCieXyy) Reportable() bool { return false }
func (a WhitePointXInCieXyy) SceneIndex() int  { return -1 }

func (a WhitePointXInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const WhitePointYInCieXyyAttr zcl.AttrID = 49

type WhitePointYInCieXyy zcl.Zu16

func (a WhitePointYInCieXyy) ID() zcl.AttrID               { return WhitePointYInCieXyyAttr }
func (a WhitePointYInCieXyy) Cluster() zcl.ClusterID       { return ColorControlID }
func (a *WhitePointYInCieXyy) Value() *WhitePointYInCieXyy { return a }
func (a WhitePointYInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *WhitePointYInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = WhitePointYInCieXyy(*nt)
	return br, err
}

func (a WhitePointYInCieXyy) Readable() bool   { return true }
func (a WhitePointYInCieXyy) Writable() bool   { return true }
func (a WhitePointYInCieXyy) Reportable() bool { return false }
func (a WhitePointYInCieXyy) SceneIndex() int  { return -1 }

func (a WhitePointYInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorPointRedXInCieXyyAttr zcl.AttrID = 50

type ColorPointRedXInCieXyy zcl.Zu16

func (a ColorPointRedXInCieXyy) ID() zcl.AttrID                  { return ColorPointRedXInCieXyyAttr }
func (a ColorPointRedXInCieXyy) Cluster() zcl.ClusterID          { return ColorControlID }
func (a *ColorPointRedXInCieXyy) Value() *ColorPointRedXInCieXyy { return a }
func (a ColorPointRedXInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointRedXInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointRedXInCieXyy(*nt)
	return br, err
}

func (a ColorPointRedXInCieXyy) Readable() bool   { return true }
func (a ColorPointRedXInCieXyy) Writable() bool   { return true }
func (a ColorPointRedXInCieXyy) Reportable() bool { return false }
func (a ColorPointRedXInCieXyy) SceneIndex() int  { return -1 }

func (a ColorPointRedXInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorPointRedYInCieXyyAttr zcl.AttrID = 51

type ColorPointRedYInCieXyy zcl.Zu16

func (a ColorPointRedYInCieXyy) ID() zcl.AttrID                  { return ColorPointRedYInCieXyyAttr }
func (a ColorPointRedYInCieXyy) Cluster() zcl.ClusterID          { return ColorControlID }
func (a *ColorPointRedYInCieXyy) Value() *ColorPointRedYInCieXyy { return a }
func (a ColorPointRedYInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointRedYInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointRedYInCieXyy(*nt)
	return br, err
}

func (a ColorPointRedYInCieXyy) Readable() bool   { return true }
func (a ColorPointRedYInCieXyy) Writable() bool   { return true }
func (a ColorPointRedYInCieXyy) Reportable() bool { return false }
func (a ColorPointRedYInCieXyy) SceneIndex() int  { return -1 }

func (a ColorPointRedYInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorPointRedIntensityAttr zcl.AttrID = 52

type ColorPointRedIntensity zcl.Zu8

func (a ColorPointRedIntensity) ID() zcl.AttrID                  { return ColorPointRedIntensityAttr }
func (a ColorPointRedIntensity) Cluster() zcl.ClusterID          { return ColorControlID }
func (a *ColorPointRedIntensity) Value() *ColorPointRedIntensity { return a }
func (a ColorPointRedIntensity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ColorPointRedIntensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointRedIntensity(*nt)
	return br, err
}

func (a ColorPointRedIntensity) Readable() bool   { return true }
func (a ColorPointRedIntensity) Writable() bool   { return true }
func (a ColorPointRedIntensity) Reportable() bool { return false }
func (a ColorPointRedIntensity) SceneIndex() int  { return -1 }

func (a ColorPointRedIntensity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const ColorPointGreenXInCieXyyAttr zcl.AttrID = 54

type ColorPointGreenXInCieXyy zcl.Zu16

func (a ColorPointGreenXInCieXyy) ID() zcl.AttrID                    { return ColorPointGreenXInCieXyyAttr }
func (a ColorPointGreenXInCieXyy) Cluster() zcl.ClusterID            { return ColorControlID }
func (a *ColorPointGreenXInCieXyy) Value() *ColorPointGreenXInCieXyy { return a }
func (a ColorPointGreenXInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointGreenXInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointGreenXInCieXyy(*nt)
	return br, err
}

func (a ColorPointGreenXInCieXyy) Readable() bool   { return true }
func (a ColorPointGreenXInCieXyy) Writable() bool   { return true }
func (a ColorPointGreenXInCieXyy) Reportable() bool { return false }
func (a ColorPointGreenXInCieXyy) SceneIndex() int  { return -1 }

func (a ColorPointGreenXInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorPointGreenYInCieXyyAttr zcl.AttrID = 55

type ColorPointGreenYInCieXyy zcl.Zu16

func (a ColorPointGreenYInCieXyy) ID() zcl.AttrID                    { return ColorPointGreenYInCieXyyAttr }
func (a ColorPointGreenYInCieXyy) Cluster() zcl.ClusterID            { return ColorControlID }
func (a *ColorPointGreenYInCieXyy) Value() *ColorPointGreenYInCieXyy { return a }
func (a ColorPointGreenYInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointGreenYInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointGreenYInCieXyy(*nt)
	return br, err
}

func (a ColorPointGreenYInCieXyy) Readable() bool   { return true }
func (a ColorPointGreenYInCieXyy) Writable() bool   { return true }
func (a ColorPointGreenYInCieXyy) Reportable() bool { return false }
func (a ColorPointGreenYInCieXyy) SceneIndex() int  { return -1 }

func (a ColorPointGreenYInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorPointGreenIntensityAttr zcl.AttrID = 56

type ColorPointGreenIntensity zcl.Zu8

func (a ColorPointGreenIntensity) ID() zcl.AttrID                    { return ColorPointGreenIntensityAttr }
func (a ColorPointGreenIntensity) Cluster() zcl.ClusterID            { return ColorControlID }
func (a *ColorPointGreenIntensity) Value() *ColorPointGreenIntensity { return a }
func (a ColorPointGreenIntensity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ColorPointGreenIntensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointGreenIntensity(*nt)
	return br, err
}

func (a ColorPointGreenIntensity) Readable() bool   { return true }
func (a ColorPointGreenIntensity) Writable() bool   { return true }
func (a ColorPointGreenIntensity) Reportable() bool { return false }
func (a ColorPointGreenIntensity) SceneIndex() int  { return -1 }

func (a ColorPointGreenIntensity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}

const ColorPointBlueXInCieXyyAttr zcl.AttrID = 58

type ColorPointBlueXInCieXyy zcl.Zu16

func (a ColorPointBlueXInCieXyy) ID() zcl.AttrID                   { return ColorPointBlueXInCieXyyAttr }
func (a ColorPointBlueXInCieXyy) Cluster() zcl.ClusterID           { return ColorControlID }
func (a *ColorPointBlueXInCieXyy) Value() *ColorPointBlueXInCieXyy { return a }
func (a ColorPointBlueXInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointBlueXInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointBlueXInCieXyy(*nt)
	return br, err
}

func (a ColorPointBlueXInCieXyy) Readable() bool   { return true }
func (a ColorPointBlueXInCieXyy) Writable() bool   { return true }
func (a ColorPointBlueXInCieXyy) Reportable() bool { return false }
func (a ColorPointBlueXInCieXyy) SceneIndex() int  { return -1 }

func (a ColorPointBlueXInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorPointBlueYInCieXyyAttr zcl.AttrID = 59

type ColorPointBlueYInCieXyy zcl.Zu16

func (a ColorPointBlueYInCieXyy) ID() zcl.AttrID                   { return ColorPointBlueYInCieXyyAttr }
func (a ColorPointBlueYInCieXyy) Cluster() zcl.ClusterID           { return ColorControlID }
func (a *ColorPointBlueYInCieXyy) Value() *ColorPointBlueYInCieXyy { return a }
func (a ColorPointBlueYInCieXyy) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointBlueYInCieXyy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointBlueYInCieXyy(*nt)
	return br, err
}

func (a ColorPointBlueYInCieXyy) Readable() bool   { return true }
func (a ColorPointBlueYInCieXyy) Writable() bool   { return true }
func (a ColorPointBlueYInCieXyy) Reportable() bool { return false }
func (a ColorPointBlueYInCieXyy) SceneIndex() int  { return -1 }

func (a ColorPointBlueYInCieXyy) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const ColorPointBlueIntensityAttr zcl.AttrID = 60

type ColorPointBlueIntensity zcl.Zu8

func (a ColorPointBlueIntensity) ID() zcl.AttrID                   { return ColorPointBlueIntensityAttr }
func (a ColorPointBlueIntensity) Cluster() zcl.ClusterID           { return ColorControlID }
func (a *ColorPointBlueIntensity) Value() *ColorPointBlueIntensity { return a }
func (a ColorPointBlueIntensity) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *ColorPointBlueIntensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointBlueIntensity(*nt)
	return br, err
}

func (a ColorPointBlueIntensity) Readable() bool   { return true }
func (a ColorPointBlueIntensity) Writable() bool   { return true }
func (a ColorPointBlueIntensity) Reportable() bool { return false }
func (a ColorPointBlueIntensity) SceneIndex() int  { return -1 }

func (a ColorPointBlueIntensity) String() string {

	return zcl.Sprintf("%s", zcl.Zu8(a))
}
