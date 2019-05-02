// Attributes and commands for controlling the color properties of a color-capable
// light.
//
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
		CurrentXAttr:                          func() zcl.Attr { return new(CurrentX) },
		CurrentYAttr:                          func() zcl.Attr { return new(CurrentY) },
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

// Stops move to and step commands. It has no effect on a active
// color loop.
//
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
	// Specifies a lower bound on the color temperature for the
	// current move operation.
	//
	ColorTemperatureMin zcl.Zu16
	// Specifies a upper bound on the color temperature for the
	// current move operation.
	//
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
	// Specifies a lower bound on the color temperature for the
	// current step operation.
	//
	ColorTemperatureMinimumMireds zcl.Zu16
	// Specifies a upper bound on the color temperature for the
	// current step operation.
	//
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

// CurrentHue is an autogenerated attribute in the ColorControl cluster
// It contains the current hue value of the light. Hue = CurrentHue x 360 / 254
// (CurrentHue in the range 0 - 254 inclusive)
//
type CurrentHue zcl.Zu8

const CurrentHueAttr zcl.AttrID = 0

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
	return zcl.DegreesAngular.Format(float64(a))
}

// CurrentSaturation is an autogenerated attribute in the ColorControl cluster
// It holds the current saturation value of the light.
// Saturation = CurrentSaturation/254 (CurrentSaturation in the range
// 0 - 254 inclusive)
//
type CurrentSaturation zcl.Zu8

const CurrentSaturationAttr zcl.AttrID = 1

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

// RemainingTime is an autogenerated attribute in the ColorControl cluster
// It holds the time remaining, in 1/10ths of a second, until the currently
// active command will be complete
//
type RemainingTime zcl.Zu16

const RemainingTimeAttr zcl.AttrID = 2

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
	return zcl.Seconds.Format(float64(a) / 10)
}

// CurrentX is an autogenerated attribute in the ColorControl cluster
// It contains the current value of the normalized chromaticity value x,
// as defined in the CIE xyY Color Space. x = CurrentX / 65536 (CurrentX
// in the range 0 to 65279 inclusive)
//
type CurrentX zcl.Zu16

const CurrentXAttr zcl.AttrID = 3

func (a CurrentX) ID() zcl.AttrID         { return CurrentXAttr }
func (a CurrentX) Cluster() zcl.ClusterID { return ColorControlID }
func (a *CurrentX) Value() *CurrentX      { return a }
func (a CurrentX) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *CurrentX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentX(*nt)
	return br, err
}

func (a CurrentX) Readable() bool   { return true }
func (a CurrentX) Writable() bool   { return false }
func (a CurrentX) Reportable() bool { return true }
func (a CurrentX) SceneIndex() int  { return 1 }

func (a CurrentX) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// CurrentY is an autogenerated attribute in the ColorControl cluster
// It contains the current value of the normalized chromaticity value y,
// as defined in the CIE xyY Color Space. y = CurrentY / 65536 (CurrentY
// in the range 0 to 65279 inclusive)
//
type CurrentY zcl.Zu16

const CurrentYAttr zcl.AttrID = 4

func (a CurrentY) ID() zcl.AttrID         { return CurrentYAttr }
func (a CurrentY) Cluster() zcl.ClusterID { return ColorControlID }
func (a *CurrentY) Value() *CurrentY      { return a }
func (a CurrentY) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *CurrentY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentY(*nt)
	return br, err
}

func (a CurrentY) Readable() bool   { return true }
func (a CurrentY) Writable() bool   { return false }
func (a CurrentY) Reportable() bool { return true }
func (a CurrentY) SceneIndex() int  { return 2 }

func (a CurrentY) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// DriftCompensation is an autogenerated attribute in the ColorControl cluster
// It indicates what mechanism, if any, is in use for compensation for
// color/intensity drift over time
//
type DriftCompensation zcl.Zenum8

const DriftCompensationAttr zcl.AttrID = 5

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

// CompensationText is an autogenerated attribute in the ColorControl cluster
// It holds a textual indication of what mechanism, if any, is in use to
// compensate for color/intensity drift over time
//
type CompensationText zcl.Zcstring

const CompensationTextAttr zcl.AttrID = 6

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

// ColorTemperatureMireds is an autogenerated attribute in the ColorControl cluster
// It contains a scaled inverse of the current value of the color
// temperature. The unit of ColorTemperatureMireds is the mired
// (micro reciprocal degree), a.k.a mirek (micro reciprocal
// kelvin). Color temperature in kelvins = 1,000,000 / ColorTemperatureMireds,
// where ColorTemperatureMireds is in the range 1 to 65279 mireds inclusive,
// giving a color temperature range from 1,000,000 kelvins to 15.32 kelvins
//
type ColorTemperatureMireds zcl.Zu16

const ColorTemperatureMiredsAttr zcl.AttrID = 7

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
	return zcl.Mired.Format(float64(a))
}

// ColorMode is an autogenerated attribute in the ColorControl cluster
// It indicates which attributes are currently determining the color of
// the device. This attribute is optional if the device does not implement
// CurrentHue and CurrentSaturation
//
type ColorMode zcl.Zenum8

const ColorModeAttr zcl.AttrID = 8

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
		return "Current X and Current Y"
	case 0x02:
		return "Color temperature"
	}
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

// IsCurrentHueAndCurrentSaturation checks if ColorMode equals the value for Current hue and current saturation (0x00)
func (a ColorMode) IsCurrentHueAndCurrentSaturation() bool { return a == 0x00 }

// SetCurrentHueAndCurrentSaturation sets ColorMode to Current hue and current saturation (0x00)
func (a *ColorMode) SetCurrentHueAndCurrentSaturation() { *a = 0x00 }

// IsCurrentXAndCurrentY checks if ColorMode equals the value for Current X and Current Y (0x01)
func (a ColorMode) IsCurrentXAndCurrentY() bool { return a == 0x01 }

// SetCurrentXAndCurrentY sets ColorMode to Current X and Current Y (0x01)
func (a *ColorMode) SetCurrentXAndCurrentY() { *a = 0x01 }

// IsColorTemperature checks if ColorMode equals the value for Color temperature (0x02)
func (a ColorMode) IsColorTemperature() bool { return a == 0x02 }

// SetColorTemperature sets ColorMode to Color temperature (0x02)
func (a *ColorMode) SetColorTemperature() { *a = 0x02 }

// EnhancedCurrentHue is an autogenerated attribute in the ColorControl cluster
// It represents non-equidistant steps along the CIE 1931 color triangle,
// and it provides 16-bits precision. The upper 8 bits of this attribute
// are used as an index in the implementation specific XY lookup table to
// provide the non-equidistance steps. The lower 8 bits are used to
// interpolate between these steps in a linear way in order to provide color
// zoom for the user. To provide compatibility with standard ZCL, the
// CurrentHue attribute contains a hue value in the range 0 to 254,
// calculated from the EnhancedCurrentHue attribute
//
type EnhancedCurrentHue zcl.Zu16

const EnhancedCurrentHueAttr zcl.AttrID = 16384

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

// EnhancedColorMode is an autogenerated attribute in the ColorControl cluster
// It specifies which attributes are currently determining the color of
// the device
//
type EnhancedColorMode zcl.Zenum8

const EnhancedColorModeAttr zcl.AttrID = 16385

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
		return "Current X and Current Y"
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

// IsCurrentXAndCurrentY checks if EnhancedColorMode equals the value for Current X and Current Y (0x01)
func (a EnhancedColorMode) IsCurrentXAndCurrentY() bool { return a == 0x01 }

// SetCurrentXAndCurrentY sets EnhancedColorMode to Current X and Current Y (0x01)
func (a *EnhancedColorMode) SetCurrentXAndCurrentY() { *a = 0x01 }

// IsColorTemperature checks if EnhancedColorMode equals the value for Color temperature (0x02)
func (a EnhancedColorMode) IsColorTemperature() bool { return a == 0x02 }

// SetColorTemperature sets EnhancedColorMode to Color temperature (0x02)
func (a *EnhancedColorMode) SetColorTemperature() { *a = 0x02 }

// IsEnhancedCurrentHueAndCurrentSaturation checks if EnhancedColorMode equals the value for Enhanced current hue and current saturation (0x03)
func (a EnhancedColorMode) IsEnhancedCurrentHueAndCurrentSaturation() bool { return a == 0x03 }

// SetEnhancedCurrentHueAndCurrentSaturation sets EnhancedColorMode to Enhanced current hue and current saturation (0x03)
func (a *EnhancedColorMode) SetEnhancedCurrentHueAndCurrentSaturation() { *a = 0x03 }

// ColorLoopActive is an autogenerated attribute in the ColorControl cluster
// It specifies the current active status of the color loop. 0x00 means
// inactive, 0x01 means active
//
type ColorLoopActive zcl.Zu8

const ColorLoopActiveAttr zcl.AttrID = 16386

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

// ColorLoopDirection is an autogenerated attribute in the ColorControl cluster
// It specifies the current direction of the color loop. If this attribute
// has the value 0x00, the EnhancedCurrentHue is be decremented. If this
// attribute has the value 0x01, the EnhancedCurrentHue is incremented
//
type ColorLoopDirection zcl.Zu8

const ColorLoopDirectionAttr zcl.AttrID = 16387

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

// ColorLoopTime is an autogenerated attribute in the ColorControl cluster
// It specifies the number of seconds it takes to perform a full color
// loop, i.e., to cycle all values of EnhancedCurrentHue
//
type ColorLoopTime zcl.Zu16

const ColorLoopTimeAttr zcl.AttrID = 16388

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

// ColorLoopStartEnhancedHue is an autogenerated attribute in the ColorControl cluster
// It specifies the value of the EnhancedCurrentHue attribute from which
// the color loop starts
//
type ColorLoopStartEnhancedHue zcl.Zu16

const ColorLoopStartEnhancedHueAttr zcl.AttrID = 16389

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

// ColorLoopStoredEnhancedHue is an autogenerated attribute in the ColorControl cluster
// It specifies the value of the EnhancedCurrentHue attribute before the
// color loop was started. Once the color loop is complete, It is restored
// to this value
//
type ColorLoopStoredEnhancedHue zcl.Zu16

const ColorLoopStoredEnhancedHueAttr zcl.AttrID = 16390

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

// ColorCapabilities is an autogenerated attribute in the ColorControl cluster
// It specifies the color capabilities of the device supporting the color
// control cluster
//
type ColorCapabilities zcl.Zbmp16

const ColorCapabilitiesAttr zcl.AttrID = 16394

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

// ColorTemperaturePhysicalMinMireds is an autogenerated attribute in the ColorControl cluster
// It indicates the minimum mired value supported by the hardware.
// ColorTempPhysicalMinMireds corresponds to the maximum color
// temperature in Kelvins supported by the hardware.
// ColorTempPhysicalMinMireds ≤ ColorTemperatureMireds
//
type ColorTemperaturePhysicalMinMireds zcl.Zu16

const ColorTemperaturePhysicalMinMiredsAttr zcl.AttrID = 16395

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
	return zcl.Mired.Format(float64(a))
}

// ColorTemperaturePhysicalMaxMireds is an autogenerated attribute in the ColorControl cluster
// It indicates the maximum mired value supported by the hardware.
// ColorTempPhysicalMaxMireds corresponds to the minimum color
// temperature in Kelvins supported by the hardware.
// ColorTemperatureMireds ≤ ColorTempPhysicalMaxMireds
//
type ColorTemperaturePhysicalMaxMireds zcl.Zu16

const ColorTemperaturePhysicalMaxMiredsAttr zcl.AttrID = 16396

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
	return zcl.Mired.Format(float64(a))
}

// PowerOnColorTemperature is an autogenerated attribute in the ColorControl cluster
type PowerOnColorTemperature zcl.Zu16

const PowerOnColorTemperatureAttr zcl.AttrID = 16400

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

// NumberOfPrimaries is an autogenerated attribute in the ColorControl cluster
// It contains the number of color primaries implemented on this device.
// A value of 0xff indicates that the number of primaries is unknown
//
type NumberOfPrimaries zcl.Zu8

const NumberOfPrimariesAttr zcl.AttrID = 16

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

// Primary1X is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value x for this primary, as
// defined in the CIE xyY Color Space. x = PrimaryX / 65536 (PrimaryX
// in the range 0 to 65279 inclusive)
//
type Primary1X zcl.Zu16

const Primary1XAttr zcl.AttrID = 17

func (a Primary1X) ID() zcl.AttrID         { return Primary1XAttr }
func (a Primary1X) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary1X) Value() *Primary1X     { return a }
func (a Primary1X) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary1X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary1X(*nt)
	return br, err
}

func (a Primary1X) Readable() bool   { return true }
func (a Primary1X) Writable() bool   { return false }
func (a Primary1X) Reportable() bool { return false }
func (a Primary1X) SceneIndex() int  { return -1 }

func (a Primary1X) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary1Y is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value y for this primary, as
// defined in the CIE xyY Color Space. y = PrimaryY / 65536 (PrimaryY
// in the range 0 to 65279 inclusive)
//
type Primary1Y zcl.Zu16

const Primary1YAttr zcl.AttrID = 18

func (a Primary1Y) ID() zcl.AttrID         { return Primary1YAttr }
func (a Primary1Y) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary1Y) Value() *Primary1Y     { return a }
func (a Primary1Y) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary1Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary1Y(*nt)
	return br, err
}

func (a Primary1Y) Readable() bool   { return true }
func (a Primary1Y) Writable() bool   { return false }
func (a Primary1Y) Reportable() bool { return false }
func (a Primary1Y) SceneIndex() int  { return -1 }

func (a Primary1Y) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary1Intensity is an autogenerated attribute in the ColorControl cluster
// It contains a representation of the maximum intensity of this primary as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the primary with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this primary is
// not available
//
type Primary1Intensity zcl.Zu8

const Primary1IntensityAttr zcl.AttrID = 19

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

// Primary2X is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value x for this primary, as
// defined in the CIE xyY Color Space. x = PrimaryX / 65536 (PrimaryX
// in the range 0 to 65279 inclusive)
//
type Primary2X zcl.Zu16

const Primary2XAttr zcl.AttrID = 21

func (a Primary2X) ID() zcl.AttrID         { return Primary2XAttr }
func (a Primary2X) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary2X) Value() *Primary2X     { return a }
func (a Primary2X) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary2X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary2X(*nt)
	return br, err
}

func (a Primary2X) Readable() bool   { return true }
func (a Primary2X) Writable() bool   { return false }
func (a Primary2X) Reportable() bool { return false }
func (a Primary2X) SceneIndex() int  { return -1 }

func (a Primary2X) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary2Y is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value y for this primary, as
// defined in the CIE xyY Color Space. y = PrimaryY / 65536 (PrimaryY
// in the range 0 to 65279 inclusive)
//
type Primary2Y zcl.Zu16

const Primary2YAttr zcl.AttrID = 22

func (a Primary2Y) ID() zcl.AttrID         { return Primary2YAttr }
func (a Primary2Y) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary2Y) Value() *Primary2Y     { return a }
func (a Primary2Y) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary2Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary2Y(*nt)
	return br, err
}

func (a Primary2Y) Readable() bool   { return true }
func (a Primary2Y) Writable() bool   { return false }
func (a Primary2Y) Reportable() bool { return false }
func (a Primary2Y) SceneIndex() int  { return -1 }

func (a Primary2Y) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary2Intensity is an autogenerated attribute in the ColorControl cluster
// It contains a representation of the maximum intensity of this primary as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the primary with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this primary is
// not available
//
type Primary2Intensity zcl.Zu8

const Primary2IntensityAttr zcl.AttrID = 23

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

// Primary3X is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value x for this primary, as
// defined in the CIE xyY Color Space. x = PrimaryX / 65536 (PrimaryX
// in the range 0 to 65279 inclusive)
//
type Primary3X zcl.Zu16

const Primary3XAttr zcl.AttrID = 25

func (a Primary3X) ID() zcl.AttrID         { return Primary3XAttr }
func (a Primary3X) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary3X) Value() *Primary3X     { return a }
func (a Primary3X) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary3X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary3X(*nt)
	return br, err
}

func (a Primary3X) Readable() bool   { return true }
func (a Primary3X) Writable() bool   { return false }
func (a Primary3X) Reportable() bool { return false }
func (a Primary3X) SceneIndex() int  { return -1 }

func (a Primary3X) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary3Y is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value y for this primary, as
// defined in the CIE xyY Color Space. y = PrimaryY / 65536 (PrimaryY
// in the range 0 to 65279 inclusive)
//
type Primary3Y zcl.Zu16

const Primary3YAttr zcl.AttrID = 26

func (a Primary3Y) ID() zcl.AttrID         { return Primary3YAttr }
func (a Primary3Y) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary3Y) Value() *Primary3Y     { return a }
func (a Primary3Y) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary3Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary3Y(*nt)
	return br, err
}

func (a Primary3Y) Readable() bool   { return true }
func (a Primary3Y) Writable() bool   { return false }
func (a Primary3Y) Reportable() bool { return false }
func (a Primary3Y) SceneIndex() int  { return -1 }

func (a Primary3Y) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary3Intensity is an autogenerated attribute in the ColorControl cluster
// It contains a representation of the maximum intensity of this primary as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the primary with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this primary is
// not available
//
type Primary3Intensity zcl.Zu8

const Primary3IntensityAttr zcl.AttrID = 27

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

// Primary4X is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value x for this primary, as
// defined in the CIE xyY Color Space. x = PrimaryX / 65536 (PrimaryX
// in the range 0 to 65279 inclusive)
//
type Primary4X zcl.Zu16

const Primary4XAttr zcl.AttrID = 32

func (a Primary4X) ID() zcl.AttrID         { return Primary4XAttr }
func (a Primary4X) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary4X) Value() *Primary4X     { return a }
func (a Primary4X) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary4X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary4X(*nt)
	return br, err
}

func (a Primary4X) Readable() bool   { return true }
func (a Primary4X) Writable() bool   { return false }
func (a Primary4X) Reportable() bool { return false }
func (a Primary4X) SceneIndex() int  { return -1 }

func (a Primary4X) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary4Y is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value y for this primary, as
// defined in the CIE xyY Color Space. y = PrimaryY / 65536 (PrimaryY
// in the range 0 to 65279 inclusive)
//
type Primary4Y zcl.Zu16

const Primary4YAttr zcl.AttrID = 33

func (a Primary4Y) ID() zcl.AttrID         { return Primary4YAttr }
func (a Primary4Y) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary4Y) Value() *Primary4Y     { return a }
func (a Primary4Y) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary4Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary4Y(*nt)
	return br, err
}

func (a Primary4Y) Readable() bool   { return true }
func (a Primary4Y) Writable() bool   { return false }
func (a Primary4Y) Reportable() bool { return false }
func (a Primary4Y) SceneIndex() int  { return -1 }

func (a Primary4Y) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary4Intensity is an autogenerated attribute in the ColorControl cluster
// It contains a representation of the maximum intensity of this primary as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the primary with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this primary is
// not available
//
type Primary4Intensity zcl.Zu8

const Primary4IntensityAttr zcl.AttrID = 34

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

// Primary5X is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value x for this primary, as
// defined in the CIE xyY Color Space. x = PrimaryX / 65536 (PrimaryX
// in the range 0 to 65279 inclusive)
//
type Primary5X zcl.Zu16

const Primary5XAttr zcl.AttrID = 36

func (a Primary5X) ID() zcl.AttrID         { return Primary5XAttr }
func (a Primary5X) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary5X) Value() *Primary5X     { return a }
func (a Primary5X) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary5X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary5X(*nt)
	return br, err
}

func (a Primary5X) Readable() bool   { return true }
func (a Primary5X) Writable() bool   { return false }
func (a Primary5X) Reportable() bool { return false }
func (a Primary5X) SceneIndex() int  { return -1 }

func (a Primary5X) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary5Y is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value y for this primary, as
// defined in the CIE xyY Color Space. y = PrimaryY / 65536 (PrimaryY
// in the range 0 to 65279 inclusive)
//
type Primary5Y zcl.Zu16

const Primary5YAttr zcl.AttrID = 37

func (a Primary5Y) ID() zcl.AttrID         { return Primary5YAttr }
func (a Primary5Y) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary5Y) Value() *Primary5Y     { return a }
func (a Primary5Y) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary5Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary5Y(*nt)
	return br, err
}

func (a Primary5Y) Readable() bool   { return true }
func (a Primary5Y) Writable() bool   { return false }
func (a Primary5Y) Reportable() bool { return false }
func (a Primary5Y) SceneIndex() int  { return -1 }

func (a Primary5Y) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary5Intensity is an autogenerated attribute in the ColorControl cluster
// It contains a representation of the maximum intensity of this primary as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the primary with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this primary is
// not available
//
type Primary5Intensity zcl.Zu8

const Primary5IntensityAttr zcl.AttrID = 38

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

// Primary6X is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value x for this primary, as
// defined in the CIE xyY Color Space. x = PrimaryX / 65536 (PrimaryX
// in the range 0 to 65279 inclusive)
//
type Primary6X zcl.Zu16

const Primary6XAttr zcl.AttrID = 40

func (a Primary6X) ID() zcl.AttrID         { return Primary6XAttr }
func (a Primary6X) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary6X) Value() *Primary6X     { return a }
func (a Primary6X) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary6X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary6X(*nt)
	return br, err
}

func (a Primary6X) Readable() bool   { return true }
func (a Primary6X) Writable() bool   { return false }
func (a Primary6X) Reportable() bool { return false }
func (a Primary6X) SceneIndex() int  { return -1 }

func (a Primary6X) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary6Y is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value y for this primary, as
// defined in the CIE xyY Color Space. y = PrimaryY / 65536 (PrimaryY
// in the range 0 to 65279 inclusive)
//
type Primary6Y zcl.Zu16

const Primary6YAttr zcl.AttrID = 41

func (a Primary6Y) ID() zcl.AttrID         { return Primary6YAttr }
func (a Primary6Y) Cluster() zcl.ClusterID { return ColorControlID }
func (a *Primary6Y) Value() *Primary6Y     { return a }
func (a Primary6Y) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Primary6Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary6Y(*nt)
	return br, err
}

func (a Primary6Y) Readable() bool   { return true }
func (a Primary6Y) Writable() bool   { return false }
func (a Primary6Y) Reportable() bool { return false }
func (a Primary6Y) SceneIndex() int  { return -1 }

func (a Primary6Y) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// Primary6Intensity is an autogenerated attribute in the ColorControl cluster
// It contains a representation of the maximum intensity of this primary as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the primary with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this primary is
// not available
//
type Primary6Intensity zcl.Zu8

const Primary6IntensityAttr zcl.AttrID = 42

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

// WhitePointX is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value x, as defined in the
// CIE xyY Color Space, of the current white point of the device.
// x = WhitePointX / 65536 (WhitePointX in the range 0 to 65279 inclusive)
//
type WhitePointX zcl.Zu16

const WhitePointXAttr zcl.AttrID = 48

func (a WhitePointX) ID() zcl.AttrID         { return WhitePointXAttr }
func (a WhitePointX) Cluster() zcl.ClusterID { return ColorControlID }
func (a *WhitePointX) Value() *WhitePointX   { return a }
func (a WhitePointX) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *WhitePointX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = WhitePointX(*nt)
	return br, err
}

func (a WhitePointX) Readable() bool   { return true }
func (a WhitePointX) Writable() bool   { return true }
func (a WhitePointX) Reportable() bool { return false }
func (a WhitePointX) SceneIndex() int  { return -1 }

func (a WhitePointX) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// WhitePointY is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value y, as defined in the
// CIE xyY Color Space, of the current white point of the device.
// y = WhitePointY / 65536 (WhitePointY in the range 0 to 65279 inclusive)
//
type WhitePointY zcl.Zu16

const WhitePointYAttr zcl.AttrID = 49

func (a WhitePointY) ID() zcl.AttrID         { return WhitePointYAttr }
func (a WhitePointY) Cluster() zcl.ClusterID { return ColorControlID }
func (a *WhitePointY) Value() *WhitePointY   { return a }
func (a WhitePointY) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *WhitePointY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = WhitePointY(*nt)
	return br, err
}

func (a WhitePointY) Readable() bool   { return true }
func (a WhitePointY) Writable() bool   { return true }
func (a WhitePointY) Reportable() bool { return false }
func (a WhitePointY) SceneIndex() int  { return -1 }

func (a WhitePointY) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// ColorPointRedX is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value x, as defined in the
// CIE xyY Color Space, of the red color point of the device.
// x = ColorPointRX / 65536 (ColorPointRX in the range 0 to 65279 inclusive)
//
type ColorPointRedX zcl.Zu16

const ColorPointRedXAttr zcl.AttrID = 50

func (a ColorPointRedX) ID() zcl.AttrID          { return ColorPointRedXAttr }
func (a ColorPointRedX) Cluster() zcl.ClusterID  { return ColorControlID }
func (a *ColorPointRedX) Value() *ColorPointRedX { return a }
func (a ColorPointRedX) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointRedX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointRedX(*nt)
	return br, err
}

func (a ColorPointRedX) Readable() bool   { return true }
func (a ColorPointRedX) Writable() bool   { return true }
func (a ColorPointRedX) Reportable() bool { return false }
func (a ColorPointRedX) SceneIndex() int  { return -1 }

func (a ColorPointRedX) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// ColorPointRedY is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value y, as defined in the
// CIE xyY Color Space, of the red color point of the device.
// y = ColorPointRY / 65536 (ColorPointRY in the range 0 to 65279 inclusive)
//
type ColorPointRedY zcl.Zu16

const ColorPointRedYAttr zcl.AttrID = 51

func (a ColorPointRedY) ID() zcl.AttrID          { return ColorPointRedYAttr }
func (a ColorPointRedY) Cluster() zcl.ClusterID  { return ColorControlID }
func (a *ColorPointRedY) Value() *ColorPointRedY { return a }
func (a ColorPointRedY) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointRedY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointRedY(*nt)
	return br, err
}

func (a ColorPointRedY) Readable() bool   { return true }
func (a ColorPointRedY) Writable() bool   { return true }
func (a ColorPointRedY) Reportable() bool { return false }
func (a ColorPointRedY) SceneIndex() int  { return -1 }

func (a ColorPointRedY) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// ColorPointRedIntensity is an autogenerated attribute in the ColorControl cluster
// It contains a representation of the relative intensity of the red color
// point as defined in the Dimming Light Curve in the Ballast Configuration
// cluster, normalized such that the color point with the highest relative
// intensity contains the value 0xfe
//
type ColorPointRedIntensity zcl.Zu8

const ColorPointRedIntensityAttr zcl.AttrID = 52

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

// ColorPointGreenX is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value x, as defined in the
// CIE xyY Color Space, of the green color point of the device.
// x = ColorPointGX / 65536 (ColorPointGX in the range 0 to 65279 inclusive)
//
type ColorPointGreenX zcl.Zu16

const ColorPointGreenXAttr zcl.AttrID = 54

func (a ColorPointGreenX) ID() zcl.AttrID            { return ColorPointGreenXAttr }
func (a ColorPointGreenX) Cluster() zcl.ClusterID    { return ColorControlID }
func (a *ColorPointGreenX) Value() *ColorPointGreenX { return a }
func (a ColorPointGreenX) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointGreenX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointGreenX(*nt)
	return br, err
}

func (a ColorPointGreenX) Readable() bool   { return true }
func (a ColorPointGreenX) Writable() bool   { return true }
func (a ColorPointGreenX) Reportable() bool { return false }
func (a ColorPointGreenX) SceneIndex() int  { return -1 }

func (a ColorPointGreenX) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// ColorPointGreenY is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value y, as defined in the
// CIE xyY Color Space, of the green color point of the device.
// y = ColorPointGY / 65536 (ColorPointGY in the range 0 to 65279 inclusive)
//
type ColorPointGreenY zcl.Zu16

const ColorPointGreenYAttr zcl.AttrID = 55

func (a ColorPointGreenY) ID() zcl.AttrID            { return ColorPointGreenYAttr }
func (a ColorPointGreenY) Cluster() zcl.ClusterID    { return ColorControlID }
func (a *ColorPointGreenY) Value() *ColorPointGreenY { return a }
func (a ColorPointGreenY) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointGreenY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointGreenY(*nt)
	return br, err
}

func (a ColorPointGreenY) Readable() bool   { return true }
func (a ColorPointGreenY) Writable() bool   { return true }
func (a ColorPointGreenY) Reportable() bool { return false }
func (a ColorPointGreenY) SceneIndex() int  { return -1 }

func (a ColorPointGreenY) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// ColorPointGreenIntensity is an autogenerated attribute in the ColorControl cluster
// It contains a representation of the relative intensity of the green
// color point as defined in the Dimming Light Curve in the Ballast Configuration
// cluster, normalized such that the color point with the highest relative
// intensity contains the value 0xfe
//
type ColorPointGreenIntensity zcl.Zu8

const ColorPointGreenIntensityAttr zcl.AttrID = 56

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

// ColorPointBlueX is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value x, as defined in the
// CIE xyY Color Space, of the blue color point of the device.
// x = ColorPointBX / 65536 (ColorPointBX in the range 0 to 65279 inclusive)
//
type ColorPointBlueX zcl.Zu16

const ColorPointBlueXAttr zcl.AttrID = 58

func (a ColorPointBlueX) ID() zcl.AttrID           { return ColorPointBlueXAttr }
func (a ColorPointBlueX) Cluster() zcl.ClusterID   { return ColorControlID }
func (a *ColorPointBlueX) Value() *ColorPointBlueX { return a }
func (a ColorPointBlueX) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointBlueX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointBlueX(*nt)
	return br, err
}

func (a ColorPointBlueX) Readable() bool   { return true }
func (a ColorPointBlueX) Writable() bool   { return true }
func (a ColorPointBlueX) Reportable() bool { return false }
func (a ColorPointBlueX) SceneIndex() int  { return -1 }

func (a ColorPointBlueX) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// ColorPointBlueY is an autogenerated attribute in the ColorControl cluster
// It contains the normalized chromaticity value y, as defined in the
// CIE xyY Color Space, of the blue color point of the device.
// y = ColorPointBY / 65536 (ColorPointBY in the range 0 to 65279 inclusive)
//
type ColorPointBlueY zcl.Zu16

const ColorPointBlueYAttr zcl.AttrID = 59

func (a ColorPointBlueY) ID() zcl.AttrID           { return ColorPointBlueYAttr }
func (a ColorPointBlueY) Cluster() zcl.ClusterID   { return ColorControlID }
func (a *ColorPointBlueY) Value() *ColorPointBlueY { return a }
func (a ColorPointBlueY) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *ColorPointBlueY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointBlueY(*nt)
	return br, err
}

func (a ColorPointBlueY) Readable() bool   { return true }
func (a ColorPointBlueY) Writable() bool   { return true }
func (a ColorPointBlueY) Reportable() bool { return false }
func (a ColorPointBlueY) SceneIndex() int  { return -1 }

func (a ColorPointBlueY) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

// ColorPointBlueIntensity is an autogenerated attribute in the ColorControl cluster
// It contains a representation of the relative intensity of the blue
// color point as defined in the Dimming Light Curve in the Ballast Configuration
// cluster, normalized such that the color point with the highest relative
// intensity contains the value 0xfe
//
type ColorPointBlueIntensity zcl.Zu8

const ColorPointBlueIntensityAttr zcl.AttrID = 60

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
