package lighting

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID

type Action zcl.Zenum8

func (Action) Name() string                  { return "Action" }
func (a *Action) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *Action) Value() zcl.Val             { return a }
func (a Action) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *Action) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Action(*nt)
	return br, err
}

func (a *Action) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = Action(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Action) String() string {
	switch a {
	case 0x00:
		return "De-activate color loop"
	case 0x01:
		return "Activate from ColorLoopStartEnhancedHue"
	case 0x02:
		return "Activate from EnhancedCurrentHue"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a Action) IsDeActivateColorLoop() bool                   { return a == 0x00 }
func (a Action) IsActivateFromColorloopstartenhancedhue() bool { return a == 0x01 }
func (a Action) IsActivateFromEnhancedcurrenthue() bool        { return a == 0x02 }
func (a *Action) SetDeActivateColorLoop()                      { *a = 0x00 }
func (a *Action) SetActivateFromColorloopstartenhancedhue()    { *a = 0x01 }
func (a *Action) SetActivateFromEnhancedcurrenthue()           { *a = 0x02 }

func (Action) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "De-activate color loop"},
		{Value: 0x01, Name: "Activate from ColorLoopStartEnhancedHue"},
		{Value: 0x02, Name: "Activate from EnhancedCurrentHue"},
	}
}

// BallastFactorAdjustment specifies the multiplication factor, as a percentage, to be applied
// to the configured light output of the lamps. A typical usage of this
// mechanism is to compensate for reduction in efficiency over the lifetime
// of a lamp. The light output is given by
// Actual light output = configured light output x BallastFactorAdjustment / 100%
// The range for this attribute is manufacturer dependent
type BallastFactorAdjustment zcl.Zu8

const BallastFactorAdjustmentAttr zcl.AttrID = 21

func (BallastFactorAdjustment) ID() zcl.AttrID   { return BallastFactorAdjustmentAttr }
func (BallastFactorAdjustment) Readable() bool   { return true }
func (BallastFactorAdjustment) Writable() bool   { return true }
func (BallastFactorAdjustment) Reportable() bool { return false }
func (BallastFactorAdjustment) SceneIndex() int  { return -1 }

func (BallastFactorAdjustment) Name() string                  { return "Ballast Factor Adjustment" }
func (a *BallastFactorAdjustment) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *BallastFactorAdjustment) Value() zcl.Val             { return a }
func (a BallastFactorAdjustment) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *BallastFactorAdjustment) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = BallastFactorAdjustment(*nt)
	return br, err
}

func (a *BallastFactorAdjustment) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = BallastFactorAdjustment(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BallastFactorAdjustment) String() string {
	return zcl.Percent.Format(float64(a))
}

// BallastStatus It specifies the activity status of the ballast functions. Where a
// function is active, the corresponding bit is 1. Where a function is
// not active, the corresponding bit is 0. This means that bit 0 with
// a value of 0 means the ballast is operational and bit 1 with a
// value of 0 that the lamp(s) is/are in their socket(s)
type BallastStatus zcl.Zbmp8

const BallastStatusAttr zcl.AttrID = 2

func (BallastStatus) ID() zcl.AttrID   { return BallastStatusAttr }
func (BallastStatus) Readable() bool   { return true }
func (BallastStatus) Writable() bool   { return false }
func (BallastStatus) Reportable() bool { return false }
func (BallastStatus) SceneIndex() int  { return -1 }

func (BallastStatus) Name() string                  { return "Ballast Status" }
func (a *BallastStatus) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *BallastStatus) Value() zcl.Val             { return a }
func (a BallastStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *BallastStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = BallastStatus(*nt)
	return br, err
}

func (a *BallastStatus) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = BallastStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a BallastStatus) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Non-operational")
		case 1:
			bstr = append(bstr, "Lamp not in socket")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a BallastStatus) IsNonOperational() bool    { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a BallastStatus) IsLampNotInSocket() bool   { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a *BallastStatus) SetNonOperational(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *BallastStatus) SetLampNotInSocket(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}

func (BallastStatus) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Non-operational"},
		{Value: 1, Name: "Lamp not in socket"},
	}
}

// ColorMode indicates which attributes are currently determining the color of
// the device. This attribute is optional if the device does not implement
// CurrentHue and CurrentSaturation
type ColorMode zcl.Zenum8

const ColorModeAttr zcl.AttrID = 8

func (ColorMode) ID() zcl.AttrID   { return ColorModeAttr }
func (ColorMode) Readable() bool   { return true }
func (ColorMode) Writable() bool   { return false }
func (ColorMode) Reportable() bool { return false }
func (ColorMode) SceneIndex() int  { return -1 }

func (ColorMode) Name() string                  { return "Color Mode" }
func (a *ColorMode) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *ColorMode) Value() zcl.Val             { return a }
func (a ColorMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *ColorMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorMode(*nt)
	return br, err
}

func (a *ColorMode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = ColorMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorMode) String() string {
	switch a {
	case 0x00:
		return "Current hue and current saturation"
	case 0x01:
		return "Current X and Current Y"
	case 0x02:
		return "Color temperature"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a ColorMode) IsCurrentHueAndCurrentSaturation() bool { return a == 0x00 }
func (a ColorMode) IsCurrentXAndCurrentY() bool            { return a == 0x01 }
func (a ColorMode) IsColorTemperature() bool               { return a == 0x02 }
func (a *ColorMode) SetCurrentHueAndCurrentSaturation()    { *a = 0x00 }
func (a *ColorMode) SetCurrentXAndCurrentY()               { *a = 0x01 }
func (a *ColorMode) SetColorTemperature()                  { *a = 0x02 }

func (ColorMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Current hue and current saturation"},
		{Value: 0x01, Name: "Current X and Current Y"},
		{Value: 0x02, Name: "Color temperature"},
	}
}

// ColorPointBlueX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointBlueX zcl.Zu16

const ColorPointBlueXAttr zcl.AttrID = 58

func (ColorPointBlueX) ID() zcl.AttrID   { return ColorPointBlueXAttr }
func (ColorPointBlueX) Readable() bool   { return true }
func (ColorPointBlueX) Writable() bool   { return true }
func (ColorPointBlueX) Reportable() bool { return false }
func (ColorPointBlueX) SceneIndex() int  { return -1 }

func (ColorPointBlueX) Name() string                  { return "Color Point Blue X" }
func (a *ColorPointBlueX) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorPointBlueX) Value() zcl.Val             { return a }
func (a ColorPointBlueX) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorPointBlueX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointBlueX(*nt)
	return br, err
}

func (a *ColorPointBlueX) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorPointBlueX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorPointBlueX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorPointBlueY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointBlueY zcl.Zu16

const ColorPointBlueYAttr zcl.AttrID = 59

func (ColorPointBlueY) ID() zcl.AttrID   { return ColorPointBlueYAttr }
func (ColorPointBlueY) Readable() bool   { return true }
func (ColorPointBlueY) Writable() bool   { return true }
func (ColorPointBlueY) Reportable() bool { return false }
func (ColorPointBlueY) SceneIndex() int  { return -1 }

func (ColorPointBlueY) Name() string                  { return "Color Point Blue Y" }
func (a *ColorPointBlueY) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorPointBlueY) Value() zcl.Val             { return a }
func (a ColorPointBlueY) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorPointBlueY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointBlueY(*nt)
	return br, err
}

func (a *ColorPointBlueY) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorPointBlueY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorPointBlueY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorPointBlueIntensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type ColorPointBlueIntensity zcl.Zu8

const ColorPointBlueIntensityAttr zcl.AttrID = 60

func (ColorPointBlueIntensity) ID() zcl.AttrID   { return ColorPointBlueIntensityAttr }
func (ColorPointBlueIntensity) Readable() bool   { return true }
func (ColorPointBlueIntensity) Writable() bool   { return true }
func (ColorPointBlueIntensity) Reportable() bool { return false }
func (ColorPointBlueIntensity) SceneIndex() int  { return -1 }

func (ColorPointBlueIntensity) Name() string                  { return "Color Point Blue intensity" }
func (a *ColorPointBlueIntensity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *ColorPointBlueIntensity) Value() zcl.Val             { return a }
func (a ColorPointBlueIntensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ColorPointBlueIntensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointBlueIntensity(*nt)
	return br, err
}

func (a *ColorPointBlueIntensity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ColorPointBlueIntensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorPointBlueIntensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// ColorPointGreenX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointGreenX zcl.Zu16

const ColorPointGreenXAttr zcl.AttrID = 54

func (ColorPointGreenX) ID() zcl.AttrID   { return ColorPointGreenXAttr }
func (ColorPointGreenX) Readable() bool   { return true }
func (ColorPointGreenX) Writable() bool   { return true }
func (ColorPointGreenX) Reportable() bool { return false }
func (ColorPointGreenX) SceneIndex() int  { return -1 }

func (ColorPointGreenX) Name() string                  { return "Color Point Green X" }
func (a *ColorPointGreenX) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorPointGreenX) Value() zcl.Val             { return a }
func (a ColorPointGreenX) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorPointGreenX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointGreenX(*nt)
	return br, err
}

func (a *ColorPointGreenX) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorPointGreenX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorPointGreenX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorPointGreenY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointGreenY zcl.Zu16

const ColorPointGreenYAttr zcl.AttrID = 55

func (ColorPointGreenY) ID() zcl.AttrID   { return ColorPointGreenYAttr }
func (ColorPointGreenY) Readable() bool   { return true }
func (ColorPointGreenY) Writable() bool   { return true }
func (ColorPointGreenY) Reportable() bool { return false }
func (ColorPointGreenY) SceneIndex() int  { return -1 }

func (ColorPointGreenY) Name() string                  { return "Color Point Green Y" }
func (a *ColorPointGreenY) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorPointGreenY) Value() zcl.Val             { return a }
func (a ColorPointGreenY) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorPointGreenY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointGreenY(*nt)
	return br, err
}

func (a *ColorPointGreenY) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorPointGreenY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorPointGreenY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorPointGreenIntensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type ColorPointGreenIntensity zcl.Zu8

const ColorPointGreenIntensityAttr zcl.AttrID = 56

func (ColorPointGreenIntensity) ID() zcl.AttrID   { return ColorPointGreenIntensityAttr }
func (ColorPointGreenIntensity) Readable() bool   { return true }
func (ColorPointGreenIntensity) Writable() bool   { return true }
func (ColorPointGreenIntensity) Reportable() bool { return false }
func (ColorPointGreenIntensity) SceneIndex() int  { return -1 }

func (ColorPointGreenIntensity) Name() string                  { return "Color Point Green intensity" }
func (a *ColorPointGreenIntensity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *ColorPointGreenIntensity) Value() zcl.Val             { return a }
func (a ColorPointGreenIntensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ColorPointGreenIntensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointGreenIntensity(*nt)
	return br, err
}

func (a *ColorPointGreenIntensity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ColorPointGreenIntensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorPointGreenIntensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// ColorPointRedX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointRedX zcl.Zu16

const ColorPointRedXAttr zcl.AttrID = 50

func (ColorPointRedX) ID() zcl.AttrID   { return ColorPointRedXAttr }
func (ColorPointRedX) Readable() bool   { return true }
func (ColorPointRedX) Writable() bool   { return true }
func (ColorPointRedX) Reportable() bool { return false }
func (ColorPointRedX) SceneIndex() int  { return -1 }

func (ColorPointRedX) Name() string                  { return "Color Point Red X" }
func (a *ColorPointRedX) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorPointRedX) Value() zcl.Val             { return a }
func (a ColorPointRedX) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorPointRedX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointRedX(*nt)
	return br, err
}

func (a *ColorPointRedX) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorPointRedX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorPointRedX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorPointRedY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointRedY zcl.Zu16

const ColorPointRedYAttr zcl.AttrID = 51

func (ColorPointRedY) ID() zcl.AttrID   { return ColorPointRedYAttr }
func (ColorPointRedY) Readable() bool   { return true }
func (ColorPointRedY) Writable() bool   { return true }
func (ColorPointRedY) Reportable() bool { return false }
func (ColorPointRedY) SceneIndex() int  { return -1 }

func (ColorPointRedY) Name() string                  { return "Color Point Red Y" }
func (a *ColorPointRedY) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorPointRedY) Value() zcl.Val             { return a }
func (a ColorPointRedY) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorPointRedY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointRedY(*nt)
	return br, err
}

func (a *ColorPointRedY) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorPointRedY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorPointRedY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorPointRedIntensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type ColorPointRedIntensity zcl.Zu8

const ColorPointRedIntensityAttr zcl.AttrID = 52

func (ColorPointRedIntensity) ID() zcl.AttrID   { return ColorPointRedIntensityAttr }
func (ColorPointRedIntensity) Readable() bool   { return true }
func (ColorPointRedIntensity) Writable() bool   { return true }
func (ColorPointRedIntensity) Reportable() bool { return false }
func (ColorPointRedIntensity) SceneIndex() int  { return -1 }

func (ColorPointRedIntensity) Name() string                  { return "Color Point Red intensity" }
func (a *ColorPointRedIntensity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *ColorPointRedIntensity) Value() zcl.Val             { return a }
func (a ColorPointRedIntensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ColorPointRedIntensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorPointRedIntensity(*nt)
	return br, err
}

func (a *ColorPointRedIntensity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ColorPointRedIntensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorPointRedIntensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type ColorTemperature zcl.Zu16

func (ColorTemperature) Name() string                  { return "Color Temperature" }
func (a *ColorTemperature) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorTemperature) Value() zcl.Val             { return a }
func (a ColorTemperature) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperature(*nt)
	return br, err
}

func (a *ColorTemperature) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorTemperature) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ColorTemperatureMax zcl.Zu16

func (ColorTemperatureMax) Name() string                  { return "Color Temperature max" }
func (a *ColorTemperatureMax) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorTemperatureMax) Value() zcl.Val             { return a }
func (a ColorTemperatureMax) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorTemperatureMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperatureMax(*nt)
	return br, err
}

func (a *ColorTemperatureMax) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorTemperatureMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorTemperatureMax) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ColorTemperatureMaxMireds zcl.Zu16

func (ColorTemperatureMaxMireds) Name() string                  { return "Color Temperature max Mireds" }
func (a *ColorTemperatureMaxMireds) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorTemperatureMaxMireds) Value() zcl.Val             { return a }
func (a ColorTemperatureMaxMireds) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorTemperatureMaxMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperatureMaxMireds(*nt)
	return br, err
}

func (a *ColorTemperatureMaxMireds) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorTemperatureMaxMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorTemperatureMaxMireds) String() string {
	return zcl.Mired.Format(float64(a))
}

type ColorTemperatureMin zcl.Zu16

func (ColorTemperatureMin) Name() string                  { return "Color Temperature min" }
func (a *ColorTemperatureMin) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorTemperatureMin) Value() zcl.Val             { return a }
func (a ColorTemperatureMin) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorTemperatureMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperatureMin(*nt)
	return br, err
}

func (a *ColorTemperatureMin) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorTemperatureMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorTemperatureMin) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type ColorTemperatureMinMireds zcl.Zu16

func (ColorTemperatureMinMireds) Name() string                  { return "Color Temperature min Mireds" }
func (a *ColorTemperatureMinMireds) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorTemperatureMinMireds) Value() zcl.Val             { return a }
func (a ColorTemperatureMinMireds) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorTemperatureMinMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperatureMinMireds(*nt)
	return br, err
}

func (a *ColorTemperatureMinMireds) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorTemperatureMinMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorTemperatureMinMireds) String() string {
	return zcl.Mired.Format(float64(a))
}

// ColorX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorX zcl.Zu16

func (ColorX) Name() string                  { return "Color X" }
func (a *ColorX) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorX) Value() zcl.Val             { return a }
func (a ColorX) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorX(*nt)
	return br, err
}

func (a *ColorX) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorY zcl.Zu16

func (ColorY) Name() string                  { return "Color Y" }
func (a *ColorY) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorY) Value() zcl.Val             { return a }
func (a ColorY) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorY(*nt)
	return br, err
}

func (a *ColorY) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorCapabilities specifies the color capabilities of the device supporting the color
// control cluster
type ColorCapabilities zcl.Zbmp16

const ColorCapabilitiesAttr zcl.AttrID = 16394

func (ColorCapabilities) ID() zcl.AttrID   { return ColorCapabilitiesAttr }
func (ColorCapabilities) Readable() bool   { return true }
func (ColorCapabilities) Writable() bool   { return false }
func (ColorCapabilities) Reportable() bool { return false }
func (ColorCapabilities) SceneIndex() int  { return -1 }

func (ColorCapabilities) Name() string                  { return "Color capabilities" }
func (a *ColorCapabilities) TypeID() zcl.TypeID         { return zcl.Zbmp16(*a).ID() }
func (a *ColorCapabilities) Value() zcl.Val             { return a }
func (a ColorCapabilities) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(a).MarshalZcl() }

func (a *ColorCapabilities) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorCapabilities(*nt)
	return br, err
}

func (a *ColorCapabilities) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp16); ok {
		*a = ColorCapabilities(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorCapabilities) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Hue saturation")
		case 1:
			bstr = append(bstr, "Enhanced Hue saturation")
		case 2:
			bstr = append(bstr, "Color loop")
		case 3:
			bstr = append(bstr, "CIE 1931 XY")
		case 4:
			bstr = append(bstr, "Color temperature")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a ColorCapabilities) IsHueSaturation() bool         { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a ColorCapabilities) IsEnhancedHueSaturation() bool { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a ColorCapabilities) IsColorLoop() bool             { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a ColorCapabilities) IsCie1931Xy() bool             { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a ColorCapabilities) IsColorTemperature() bool      { return zcl.BitmapTest([]byte(a[:]), 4) }
func (a *ColorCapabilities) SetHueSaturation(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b))
}
func (a *ColorCapabilities) SetEnhancedHueSaturation(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b))
}
func (a *ColorCapabilities) SetColorLoop(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b)) }
func (a *ColorCapabilities) SetCie1931Xy(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b)) }
func (a *ColorCapabilities) SetColorTemperature(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 4, b))
}

func (ColorCapabilities) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Hue saturation"},
		{Value: 1, Name: "Enhanced Hue saturation"},
		{Value: 2, Name: "Color loop"},
		{Value: 3, Name: "CIE 1931 XY"},
		{Value: 4, Name: "Color temperature"},
	}
}

// ColorControlOptions is a bitmap that determines the default behavior of some cluster commands
type ColorControlOptions zcl.Zbmp8

const ColorControlOptionsAttr zcl.AttrID = 15

func (ColorControlOptions) ID() zcl.AttrID   { return ColorControlOptionsAttr }
func (ColorControlOptions) Readable() bool   { return true }
func (ColorControlOptions) Writable() bool   { return true }
func (ColorControlOptions) Reportable() bool { return false }
func (ColorControlOptions) SceneIndex() int  { return -1 }

func (ColorControlOptions) Name() string                  { return "Color control options" }
func (a *ColorControlOptions) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *ColorControlOptions) Value() zcl.Val             { return a }
func (a ColorControlOptions) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *ColorControlOptions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorControlOptions(*nt)
	return br, err
}

func (a *ColorControlOptions) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = ColorControlOptions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorControlOptions) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0x00:
			bstr = append(bstr, "Execute if off")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a ColorControlOptions) IsExecuteIfOff() bool { return zcl.BitmapTest([]byte(a[:]), 0x00) }
func (a *ColorControlOptions) SetExecuteIfOff(b bool) {
	copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0x00, b))
}

func (ColorControlOptions) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Execute if off"},
	}
}

// ColorLoopActive specifies the current active status of the color loop. 0x00 means
// inactive, 0x01 means active
type ColorLoopActive zcl.Zu8

const ColorLoopActiveAttr zcl.AttrID = 16386

func (ColorLoopActive) ID() zcl.AttrID   { return ColorLoopActiveAttr }
func (ColorLoopActive) Readable() bool   { return true }
func (ColorLoopActive) Writable() bool   { return false }
func (ColorLoopActive) Reportable() bool { return false }
func (ColorLoopActive) SceneIndex() int  { return 5 }

func (ColorLoopActive) Name() string                  { return "Color loop active" }
func (a *ColorLoopActive) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *ColorLoopActive) Value() zcl.Val             { return a }
func (a ColorLoopActive) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ColorLoopActive) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorLoopActive(*nt)
	return br, err
}

func (a *ColorLoopActive) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ColorLoopActive(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorLoopActive) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// ColorLoopDirection specifies the current direction of the color loop. If this attribute
// has the value 0x00, the EnhancedCurrentHue is be decremented. If this
// attribute has the value 0x01, the EnhancedCurrentHue is incremented
type ColorLoopDirection zcl.Zu8

const ColorLoopDirectionAttr zcl.AttrID = 16387

func (ColorLoopDirection) ID() zcl.AttrID   { return ColorLoopDirectionAttr }
func (ColorLoopDirection) Readable() bool   { return true }
func (ColorLoopDirection) Writable() bool   { return false }
func (ColorLoopDirection) Reportable() bool { return false }
func (ColorLoopDirection) SceneIndex() int  { return 6 }

func (ColorLoopDirection) Name() string                  { return "Color loop direction" }
func (a *ColorLoopDirection) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *ColorLoopDirection) Value() zcl.Val             { return a }
func (a ColorLoopDirection) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *ColorLoopDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorLoopDirection(*nt)
	return br, err
}

func (a *ColorLoopDirection) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = ColorLoopDirection(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorLoopDirection) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// ColorLoopStartEnhancedHue specifies the value of the EnhancedCurrentHue attribute from which
// the color loop starts
type ColorLoopStartEnhancedHue zcl.Zu16

const ColorLoopStartEnhancedHueAttr zcl.AttrID = 16389

func (ColorLoopStartEnhancedHue) ID() zcl.AttrID   { return ColorLoopStartEnhancedHueAttr }
func (ColorLoopStartEnhancedHue) Readable() bool   { return true }
func (ColorLoopStartEnhancedHue) Writable() bool   { return false }
func (ColorLoopStartEnhancedHue) Reportable() bool { return false }
func (ColorLoopStartEnhancedHue) SceneIndex() int  { return -1 }

func (ColorLoopStartEnhancedHue) Name() string                  { return "Color loop start enhanced hue" }
func (a *ColorLoopStartEnhancedHue) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorLoopStartEnhancedHue) Value() zcl.Val             { return a }
func (a ColorLoopStartEnhancedHue) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorLoopStartEnhancedHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorLoopStartEnhancedHue(*nt)
	return br, err
}

func (a *ColorLoopStartEnhancedHue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorLoopStartEnhancedHue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorLoopStartEnhancedHue) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorLoopStoredEnhancedHue specifies the value of the EnhancedCurrentHue attribute before the
// color loop was started. Once the color loop is complete, It is restored
// to this value
type ColorLoopStoredEnhancedHue zcl.Zu16

const ColorLoopStoredEnhancedHueAttr zcl.AttrID = 16390

func (ColorLoopStoredEnhancedHue) ID() zcl.AttrID   { return ColorLoopStoredEnhancedHueAttr }
func (ColorLoopStoredEnhancedHue) Readable() bool   { return true }
func (ColorLoopStoredEnhancedHue) Writable() bool   { return false }
func (ColorLoopStoredEnhancedHue) Reportable() bool { return false }
func (ColorLoopStoredEnhancedHue) SceneIndex() int  { return -1 }

func (ColorLoopStoredEnhancedHue) Name() string                  { return "Color loop stored enhanced hue" }
func (a *ColorLoopStoredEnhancedHue) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorLoopStoredEnhancedHue) Value() zcl.Val             { return a }
func (a ColorLoopStoredEnhancedHue) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorLoopStoredEnhancedHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorLoopStoredEnhancedHue(*nt)
	return br, err
}

func (a *ColorLoopStoredEnhancedHue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorLoopStoredEnhancedHue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorLoopStoredEnhancedHue) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorLoopTime specifies the number of seconds it takes to perform a full color
// loop, i.e., to cycle all values of EnhancedCurrentHue
type ColorLoopTime zcl.Zu16

const ColorLoopTimeAttr zcl.AttrID = 16388

func (ColorLoopTime) ID() zcl.AttrID   { return ColorLoopTimeAttr }
func (ColorLoopTime) Readable() bool   { return true }
func (ColorLoopTime) Writable() bool   { return false }
func (ColorLoopTime) Reportable() bool { return false }
func (ColorLoopTime) SceneIndex() int  { return 7 }

func (ColorLoopTime) Name() string                  { return "Color loop time" }
func (a *ColorLoopTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorLoopTime) Value() zcl.Val             { return a }
func (a ColorLoopTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorLoopTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorLoopTime(*nt)
	return br, err
}

func (a *ColorLoopTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorLoopTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorLoopTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// ColorTemperatureMireds contains a scaled inverse of the current value of the color
// temperature. The unit of ColorTemperatureMireds is the mired
// (micro reciprocal degree), a.k.a mirek (micro reciprocal
// kelvin). Color temperature in kelvins = 1,000,000 / ColorTemperatureMireds,
// where ColorTemperatureMireds is in the range 1 to 65279 mireds inclusive,
// giving a color temperature range from 1,000,000 kelvins to 15.32 kelvins
type ColorTemperatureMireds zcl.Zu16

const ColorTemperatureMiredsAttr zcl.AttrID = 7

func (ColorTemperatureMireds) ID() zcl.AttrID   { return ColorTemperatureMiredsAttr }
func (ColorTemperatureMireds) Readable() bool   { return true }
func (ColorTemperatureMireds) Writable() bool   { return false }
func (ColorTemperatureMireds) Reportable() bool { return true }
func (ColorTemperatureMireds) SceneIndex() int  { return -1 }

func (ColorTemperatureMireds) Name() string                  { return "Color temperature Mireds" }
func (a *ColorTemperatureMireds) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *ColorTemperatureMireds) Value() zcl.Val             { return a }
func (a ColorTemperatureMireds) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *ColorTemperatureMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperatureMireds(*nt)
	return br, err
}

func (a *ColorTemperatureMireds) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorTemperatureMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorTemperatureMireds) String() string {
	return zcl.Mired.Format(float64(a))
}

// ColorTemperaturePhysicalMaxMireds indicates the maximum mired value supported by the hardware.
// ColorTempPhysicalMaxMireds corresponds to the minimum color
// temperature in Kelvins supported by the hardware.
// ColorTemperatureMireds ≤ ColorTempPhysicalMaxMireds
type ColorTemperaturePhysicalMaxMireds zcl.Zu16

const ColorTemperaturePhysicalMaxMiredsAttr zcl.AttrID = 16396

func (ColorTemperaturePhysicalMaxMireds) ID() zcl.AttrID   { return ColorTemperaturePhysicalMaxMiredsAttr }
func (ColorTemperaturePhysicalMaxMireds) Readable() bool   { return true }
func (ColorTemperaturePhysicalMaxMireds) Writable() bool   { return false }
func (ColorTemperaturePhysicalMaxMireds) Reportable() bool { return false }
func (ColorTemperaturePhysicalMaxMireds) SceneIndex() int  { return -1 }

func (ColorTemperaturePhysicalMaxMireds) Name() string          { return "Color temperature physical max Mireds" }
func (a *ColorTemperaturePhysicalMaxMireds) TypeID() zcl.TypeID { return zcl.Zu16(*a).ID() }
func (a *ColorTemperaturePhysicalMaxMireds) Value() zcl.Val     { return a }
func (a ColorTemperaturePhysicalMaxMireds) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *ColorTemperaturePhysicalMaxMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperaturePhysicalMaxMireds(*nt)
	return br, err
}

func (a *ColorTemperaturePhysicalMaxMireds) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorTemperaturePhysicalMaxMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorTemperaturePhysicalMaxMireds) String() string {
	return zcl.Mired.Format(float64(a))
}

// ColorTemperaturePhysicalMinMireds indicates the minimum mired value supported by the hardware.
// ColorTempPhysicalMinMireds corresponds to the maximum color
// temperature in Kelvins supported by the hardware.
// ColorTempPhysicalMinMireds ≤ ColorTemperatureMireds
type ColorTemperaturePhysicalMinMireds zcl.Zu16

const ColorTemperaturePhysicalMinMiredsAttr zcl.AttrID = 16395

func (ColorTemperaturePhysicalMinMireds) ID() zcl.AttrID   { return ColorTemperaturePhysicalMinMiredsAttr }
func (ColorTemperaturePhysicalMinMireds) Readable() bool   { return true }
func (ColorTemperaturePhysicalMinMireds) Writable() bool   { return false }
func (ColorTemperaturePhysicalMinMireds) Reportable() bool { return false }
func (ColorTemperaturePhysicalMinMireds) SceneIndex() int  { return -1 }

func (ColorTemperaturePhysicalMinMireds) Name() string          { return "Color temperature physical min Mireds" }
func (a *ColorTemperaturePhysicalMinMireds) TypeID() zcl.TypeID { return zcl.Zu16(*a).ID() }
func (a *ColorTemperaturePhysicalMinMireds) Value() zcl.Val     { return a }
func (a ColorTemperaturePhysicalMinMireds) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}

func (a *ColorTemperaturePhysicalMinMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = ColorTemperaturePhysicalMinMireds(*nt)
	return br, err
}

func (a *ColorTemperaturePhysicalMinMireds) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = ColorTemperaturePhysicalMinMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a ColorTemperaturePhysicalMinMireds) String() string {
	return zcl.Mired.Format(float64(a))
}

// CompensationText holds a textual indication of what mechanism, if any, is in use to
// compensate for color/intensity drift over time
type CompensationText zcl.Zcstring

const CompensationTextAttr zcl.AttrID = 6

func (CompensationText) ID() zcl.AttrID   { return CompensationTextAttr }
func (CompensationText) Readable() bool   { return true }
func (CompensationText) Writable() bool   { return false }
func (CompensationText) Reportable() bool { return false }
func (CompensationText) SceneIndex() int  { return -1 }

func (CompensationText) Name() string                  { return "Compensation Text" }
func (a *CompensationText) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *CompensationText) Value() zcl.Val             { return a }
func (a CompensationText) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *CompensationText) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = CompensationText(*nt)
	return br, err
}

func (a *CompensationText) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = CompensationText(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CompensationText) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

// CoupleColorTempToLevelMinMireds specifies a lower bound on the value of the ColorTemperatureMireds attribute for the purposes of coupling the
// ColorTemperatureMireds attribute to the CurrentLevel attribute when the CoupleColorTempToLevel bit of the
// Options attribute of the Level Control cluster is equal to 1
type CoupleColorTempToLevelMinMireds zcl.Zu16

const CoupleColorTempToLevelMinMiredsAttr zcl.AttrID = 16397

func (CoupleColorTempToLevelMinMireds) ID() zcl.AttrID   { return CoupleColorTempToLevelMinMiredsAttr }
func (CoupleColorTempToLevelMinMireds) Readable() bool   { return true }
func (CoupleColorTempToLevelMinMireds) Writable() bool   { return false }
func (CoupleColorTempToLevelMinMireds) Reportable() bool { return false }
func (CoupleColorTempToLevelMinMireds) SceneIndex() int  { return -1 }

func (CoupleColorTempToLevelMinMireds) Name() string                  { return "Couple Color Temp to Level Min Mireds" }
func (a *CoupleColorTempToLevelMinMireds) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *CoupleColorTempToLevelMinMireds) Value() zcl.Val             { return a }
func (a CoupleColorTempToLevelMinMireds) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *CoupleColorTempToLevelMinMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CoupleColorTempToLevelMinMireds(*nt)
	return br, err
}

func (a *CoupleColorTempToLevelMinMireds) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = CoupleColorTempToLevelMinMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CoupleColorTempToLevelMinMireds) String() string {
	return zcl.Mired.Format(float64(a))
}

// CurrentX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type CurrentX zcl.Zu16

const CurrentXAttr zcl.AttrID = 3

func (CurrentX) ID() zcl.AttrID   { return CurrentXAttr }
func (CurrentX) Readable() bool   { return true }
func (CurrentX) Writable() bool   { return false }
func (CurrentX) Reportable() bool { return true }
func (CurrentX) SceneIndex() int  { return 1 }

func (CurrentX) Name() string                  { return "Current X" }
func (a *CurrentX) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *CurrentX) Value() zcl.Val             { return a }
func (a CurrentX) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *CurrentX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentX(*nt)
	return br, err
}

func (a *CurrentX) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = CurrentX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// CurrentY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type CurrentY zcl.Zu16

const CurrentYAttr zcl.AttrID = 4

func (CurrentY) ID() zcl.AttrID   { return CurrentYAttr }
func (CurrentY) Readable() bool   { return true }
func (CurrentY) Writable() bool   { return false }
func (CurrentY) Reportable() bool { return true }
func (CurrentY) SceneIndex() int  { return 2 }

func (CurrentY) Name() string                  { return "Current Y" }
func (a *CurrentY) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *CurrentY) Value() zcl.Val             { return a }
func (a CurrentY) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *CurrentY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentY(*nt)
	return br, err
}

func (a *CurrentY) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = CurrentY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// CurrentHue contains the current hue value of the light. Hue = CurrentHue x 360 / 254
// (CurrentHue in the range 0 - 254 inclusive)
type CurrentHue zcl.Zu8

const CurrentHueAttr zcl.AttrID = 0

func (CurrentHue) ID() zcl.AttrID   { return CurrentHueAttr }
func (CurrentHue) Readable() bool   { return true }
func (CurrentHue) Writable() bool   { return false }
func (CurrentHue) Reportable() bool { return true }
func (CurrentHue) SceneIndex() int  { return -1 }

func (CurrentHue) Name() string                  { return "Current hue" }
func (a *CurrentHue) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *CurrentHue) Value() zcl.Val             { return a }
func (a CurrentHue) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *CurrentHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentHue(*nt)
	return br, err
}

func (a *CurrentHue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = CurrentHue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentHue) String() string {
	return zcl.DegreesAngular.Format(float64(a) / 0.70556)
}

// CurrentSaturation holds the current saturation value of the light.
// Saturation = CurrentSaturation/254 (CurrentSaturation in the range
// 0 - 254 inclusive)
type CurrentSaturation zcl.Zu8

const CurrentSaturationAttr zcl.AttrID = 1

func (CurrentSaturation) ID() zcl.AttrID   { return CurrentSaturationAttr }
func (CurrentSaturation) Readable() bool   { return true }
func (CurrentSaturation) Writable() bool   { return false }
func (CurrentSaturation) Reportable() bool { return true }
func (CurrentSaturation) SceneIndex() int  { return 4 }

func (CurrentSaturation) Name() string                  { return "Current saturation" }
func (a *CurrentSaturation) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *CurrentSaturation) Value() zcl.Val             { return a }
func (a CurrentSaturation) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *CurrentSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = CurrentSaturation(*nt)
	return br, err
}

func (a *CurrentSaturation) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = CurrentSaturation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a CurrentSaturation) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type Direction zcl.Zenum8

func (Direction) Name() string                  { return "Direction" }
func (a *Direction) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *Direction) Value() zcl.Val             { return a }
func (a Direction) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *Direction) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = Direction(*nt)
	return br, err
}

func (a *Direction) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = Direction(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Direction) String() string {
	switch a {
	case 0x00:
		return "Shortest distance"
	case 0x01:
		return "Longest Distance"
	case 0x02:
		return "Up"
	case 0x03:
		return "Down"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a Direction) IsShortestDistance() bool { return a == 0x00 }
func (a Direction) IsLongestDistance() bool  { return a == 0x01 }
func (a Direction) IsUp() bool               { return a == 0x02 }
func (a Direction) IsDown() bool             { return a == 0x03 }
func (a *Direction) SetShortestDistance()    { *a = 0x00 }
func (a *Direction) SetLongestDistance()     { *a = 0x01 }
func (a *Direction) SetUp()                  { *a = 0x02 }
func (a *Direction) SetDown()                { *a = 0x03 }

func (Direction) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Shortest distance"},
		{Value: 0x01, Name: "Longest Distance"},
		{Value: 0x02, Name: "Up"},
		{Value: 0x03, Name: "Down"},
	}
}

// DriftCompensation indicates what mechanism, if any, is in use for compensation for
// color/intensity drift over time
type DriftCompensation zcl.Zenum8

const DriftCompensationAttr zcl.AttrID = 5

func (DriftCompensation) ID() zcl.AttrID   { return DriftCompensationAttr }
func (DriftCompensation) Readable() bool   { return true }
func (DriftCompensation) Writable() bool   { return false }
func (DriftCompensation) Reportable() bool { return false }
func (DriftCompensation) SceneIndex() int  { return -1 }

func (DriftCompensation) Name() string                  { return "Drift Compensation" }
func (a *DriftCompensation) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *DriftCompensation) Value() zcl.Val             { return a }
func (a DriftCompensation) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *DriftCompensation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = DriftCompensation(*nt)
	return br, err
}

func (a *DriftCompensation) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = DriftCompensation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

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
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a DriftCompensation) IsNone() bool                                  { return a == 0x00 }
func (a DriftCompensation) IsOtherUnknown() bool                          { return a == 0x01 }
func (a DriftCompensation) IsTemperatureMonitoring() bool                 { return a == 0x02 }
func (a DriftCompensation) IsOpticalLuminanceMonitoringAndFeedback() bool { return a == 0x03 }
func (a DriftCompensation) IsOpticalColorMonitoringAndFeedback() bool     { return a == 0x04 }
func (a *DriftCompensation) SetNone()                                     { *a = 0x00 }
func (a *DriftCompensation) SetOtherUnknown()                             { *a = 0x01 }
func (a *DriftCompensation) SetTemperatureMonitoring()                    { *a = 0x02 }
func (a *DriftCompensation) SetOpticalLuminanceMonitoringAndFeedback()    { *a = 0x03 }
func (a *DriftCompensation) SetOpticalColorMonitoringAndFeedback()        { *a = 0x04 }

func (DriftCompensation) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "None"},
		{Value: 0x01, Name: "Other / Unknown"},
		{Value: 0x02, Name: "Temperature monitoring"},
		{Value: 0x03, Name: "Optical luminance monitoring and feedback"},
		{Value: 0x04, Name: "Optical color monitoring and feedback"},
	}
}

type EnhancedHue zcl.Zu16

func (EnhancedHue) Name() string                  { return "Enhanced Hue" }
func (a *EnhancedHue) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *EnhancedHue) Value() zcl.Val             { return a }
func (a EnhancedHue) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *EnhancedHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = EnhancedHue(*nt)
	return br, err
}

func (a *EnhancedHue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = EnhancedHue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a EnhancedHue) String() string {
	return zcl.DegreesAngular.Format(float64(a) / 0.002756094)
}

// EnhancedColorMode specifies which attributes are currently determining the color of
// the device
type EnhancedColorMode zcl.Zenum8

const EnhancedColorModeAttr zcl.AttrID = 16385

func (EnhancedColorMode) ID() zcl.AttrID   { return EnhancedColorModeAttr }
func (EnhancedColorMode) Readable() bool   { return true }
func (EnhancedColorMode) Writable() bool   { return false }
func (EnhancedColorMode) Reportable() bool { return false }
func (EnhancedColorMode) SceneIndex() int  { return -1 }

func (EnhancedColorMode) Name() string                  { return "Enhanced color mode" }
func (a *EnhancedColorMode) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *EnhancedColorMode) Value() zcl.Val             { return a }
func (a EnhancedColorMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *EnhancedColorMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = EnhancedColorMode(*nt)
	return br, err
}

func (a *EnhancedColorMode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = EnhancedColorMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

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
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a EnhancedColorMode) IsCurrentHueAndCurrentSaturation() bool         { return a == 0x00 }
func (a EnhancedColorMode) IsCurrentXAndCurrentY() bool                    { return a == 0x01 }
func (a EnhancedColorMode) IsColorTemperature() bool                       { return a == 0x02 }
func (a EnhancedColorMode) IsEnhancedCurrentHueAndCurrentSaturation() bool { return a == 0x03 }
func (a *EnhancedColorMode) SetCurrentHueAndCurrentSaturation()            { *a = 0x00 }
func (a *EnhancedColorMode) SetCurrentXAndCurrentY()                       { *a = 0x01 }
func (a *EnhancedColorMode) SetColorTemperature()                          { *a = 0x02 }
func (a *EnhancedColorMode) SetEnhancedCurrentHueAndCurrentSaturation()    { *a = 0x03 }

func (EnhancedColorMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Current hue and current saturation"},
		{Value: 0x01, Name: "Current X and Current Y"},
		{Value: 0x02, Name: "Color temperature"},
		{Value: 0x03, Name: "Enhanced current hue and current saturation"},
	}
}

// EnhancedCurrentHue represents non-equidistant steps along the CIE 1931 color triangle,
// and it provides 16-bits precision. The upper 8 bits of this attribute
// are used as an index in the implementation specific XY lookup table to
// provide the non-equidistance steps. The lower 8 bits are used to
// interpolate between these steps in a linear way in order to provide color
// zoom for the user. To provide compatibility with standard ZCL, the
// CurrentHue attribute contains a hue value in the range 0 to 254,
// calculated from the EnhancedCurrentHue attribute
type EnhancedCurrentHue zcl.Zu16

const EnhancedCurrentHueAttr zcl.AttrID = 16384

func (EnhancedCurrentHue) ID() zcl.AttrID   { return EnhancedCurrentHueAttr }
func (EnhancedCurrentHue) Readable() bool   { return true }
func (EnhancedCurrentHue) Writable() bool   { return false }
func (EnhancedCurrentHue) Reportable() bool { return false }
func (EnhancedCurrentHue) SceneIndex() int  { return 3 }

func (EnhancedCurrentHue) Name() string                  { return "Enhanced current hue" }
func (a *EnhancedCurrentHue) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *EnhancedCurrentHue) Value() zcl.Val             { return a }
func (a EnhancedCurrentHue) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *EnhancedCurrentHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = EnhancedCurrentHue(*nt)
	return br, err
}

func (a *EnhancedCurrentHue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = EnhancedCurrentHue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a EnhancedCurrentHue) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

type Hue zcl.Zu8

func (Hue) Name() string                  { return "Hue" }
func (a *Hue) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Hue) Value() zcl.Val             { return a }
func (a Hue) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Hue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Hue(*nt)
	return br, err
}

func (a *Hue) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Hue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Hue) String() string {
	return zcl.DegreesAngular.Format(float64(a) / 0.70556)
}

type HueDirection zcl.Zenum8

func (HueDirection) Name() string                  { return "Hue Direction" }
func (a *HueDirection) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *HueDirection) Value() zcl.Val             { return a }
func (a HueDirection) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *HueDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = HueDirection(*nt)
	return br, err
}

func (a *HueDirection) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = HueDirection(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a HueDirection) String() string {
	switch a {
	case 0x00:
		return "Decrement hue"
	case 0x01:
		return "Increment hue"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a HueDirection) IsDecrementHue() bool { return a == 0x00 }
func (a HueDirection) IsIncrementHue() bool { return a == 0x01 }
func (a *HueDirection) SetDecrementHue()    { *a = 0x00 }
func (a *HueDirection) SetIncrementHue()    { *a = 0x01 }

func (HueDirection) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Decrement hue"},
		{Value: 0x01, Name: "Increment hue"},
	}
}

// IntrinsicBallastFactor specifies, as a percentage, the ballast factor of the ballast/lamp
// combination, prior to any adjustment.
type IntrinsicBallastFactor zcl.Zu8

const IntrinsicBallastFactorAttr zcl.AttrID = 20

func (IntrinsicBallastFactor) ID() zcl.AttrID   { return IntrinsicBallastFactorAttr }
func (IntrinsicBallastFactor) Readable() bool   { return true }
func (IntrinsicBallastFactor) Writable() bool   { return true }
func (IntrinsicBallastFactor) Reportable() bool { return false }
func (IntrinsicBallastFactor) SceneIndex() int  { return -1 }

func (IntrinsicBallastFactor) Name() string                  { return "Intrinsic Ballast Factor" }
func (a *IntrinsicBallastFactor) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *IntrinsicBallastFactor) Value() zcl.Val             { return a }
func (a IntrinsicBallastFactor) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *IntrinsicBallastFactor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = IntrinsicBallastFactor(*nt)
	return br, err
}

func (a *IntrinsicBallastFactor) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = IntrinsicBallastFactor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a IntrinsicBallastFactor) String() string {
	return zcl.Percent.Format(float64(a))
}

// LampAlarmMode specifies which attributes can cause an alarm notification to be
// generated
type LampAlarmMode zcl.Zbmp8

const LampAlarmModeAttr zcl.AttrID = 52

func (LampAlarmMode) ID() zcl.AttrID   { return LampAlarmModeAttr }
func (LampAlarmMode) Readable() bool   { return true }
func (LampAlarmMode) Writable() bool   { return true }
func (LampAlarmMode) Reportable() bool { return false }
func (LampAlarmMode) SceneIndex() int  { return -1 }

func (LampAlarmMode) Name() string                  { return "Lamp Alarm Mode" }
func (a *LampAlarmMode) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *LampAlarmMode) Value() zcl.Val             { return a }
func (a LampAlarmMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *LampAlarmMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = LampAlarmMode(*nt)
	return br, err
}

func (a *LampAlarmMode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = LampAlarmMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LampAlarmMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Lamp Burn Hours")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a LampAlarmMode) IsLampBurnHours() bool    { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a *LampAlarmMode) SetLampBurnHours(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }

func (LampAlarmMode) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Lamp Burn Hours"},
	}
}

// LampBurnHours specifies the length of time, in hours, the currently connected
// lamps have been operated, cumulative since the last re-lamping. Burn
// hours are not accumulated if the lamps are off and are reset to 0
// when a lamp is changed
type LampBurnHours zcl.Zu24

const LampBurnHoursAttr zcl.AttrID = 51

func (LampBurnHours) ID() zcl.AttrID   { return LampBurnHoursAttr }
func (LampBurnHours) Readable() bool   { return true }
func (LampBurnHours) Writable() bool   { return true }
func (LampBurnHours) Reportable() bool { return false }
func (LampBurnHours) SceneIndex() int  { return -1 }

func (LampBurnHours) Name() string                  { return "Lamp Burn Hours" }
func (a *LampBurnHours) TypeID() zcl.TypeID         { return zcl.Zu24(*a).ID() }
func (a *LampBurnHours) Value() zcl.Val             { return a }
func (a LampBurnHours) MarshalZcl() ([]byte, error) { return zcl.Zu24(a).MarshalZcl() }

func (a *LampBurnHours) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = LampBurnHours(*nt)
	return br, err
}

func (a *LampBurnHours) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu24); ok {
		*a = LampBurnHours(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LampBurnHours) String() string {
	return zcl.Hours.Format(float64(a))
}

// LampBurnHoursTripPoint specifies the number of hours the LampBurnHours attribute can
// reach before an alarm is generated
type LampBurnHoursTripPoint zcl.Zu24

const LampBurnHoursTripPointAttr zcl.AttrID = 53

func (LampBurnHoursTripPoint) ID() zcl.AttrID   { return LampBurnHoursTripPointAttr }
func (LampBurnHoursTripPoint) Readable() bool   { return true }
func (LampBurnHoursTripPoint) Writable() bool   { return true }
func (LampBurnHoursTripPoint) Reportable() bool { return false }
func (LampBurnHoursTripPoint) SceneIndex() int  { return -1 }

func (LampBurnHoursTripPoint) Name() string                  { return "Lamp Burn Hours Trip Point" }
func (a *LampBurnHoursTripPoint) TypeID() zcl.TypeID         { return zcl.Zu24(*a).ID() }
func (a *LampBurnHoursTripPoint) Value() zcl.Val             { return a }
func (a LampBurnHoursTripPoint) MarshalZcl() ([]byte, error) { return zcl.Zu24(a).MarshalZcl() }

func (a *LampBurnHoursTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = LampBurnHoursTripPoint(*nt)
	return br, err
}

func (a *LampBurnHoursTripPoint) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu24); ok {
		*a = LampBurnHoursTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LampBurnHoursTripPoint) String() string {
	return zcl.Hours.Format(float64(a))
}

// LampManufacturer specifies the name of the manufacturer of the currently connected
// lamps.
type LampManufacturer zcl.Zcstring

const LampManufacturerAttr zcl.AttrID = 49

func (LampManufacturer) ID() zcl.AttrID   { return LampManufacturerAttr }
func (LampManufacturer) Readable() bool   { return true }
func (LampManufacturer) Writable() bool   { return true }
func (LampManufacturer) Reportable() bool { return false }
func (LampManufacturer) SceneIndex() int  { return -1 }

func (LampManufacturer) Name() string                  { return "Lamp Manufacturer" }
func (a *LampManufacturer) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *LampManufacturer) Value() zcl.Val             { return a }
func (a LampManufacturer) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *LampManufacturer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = LampManufacturer(*nt)
	return br, err
}

func (a *LampManufacturer) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = LampManufacturer(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LampManufacturer) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

// LampQuantity specifies the number of lamps connected to this ballast
type LampQuantity zcl.Zu8

const LampQuantityAttr zcl.AttrID = 32

func (LampQuantity) ID() zcl.AttrID   { return LampQuantityAttr }
func (LampQuantity) Readable() bool   { return true }
func (LampQuantity) Writable() bool   { return false }
func (LampQuantity) Reportable() bool { return false }
func (LampQuantity) SceneIndex() int  { return -1 }

func (LampQuantity) Name() string                  { return "Lamp Quantity" }
func (a *LampQuantity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *LampQuantity) Value() zcl.Val             { return a }
func (a LampQuantity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *LampQuantity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = LampQuantity(*nt)
	return br, err
}

func (a *LampQuantity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = LampQuantity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LampQuantity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// LampRatedHours specifies the number of hours of use the lamps are rated for by
// the manufacturer
type LampRatedHours zcl.Zu24

const LampRatedHoursAttr zcl.AttrID = 50

func (LampRatedHours) ID() zcl.AttrID   { return LampRatedHoursAttr }
func (LampRatedHours) Readable() bool   { return true }
func (LampRatedHours) Writable() bool   { return true }
func (LampRatedHours) Reportable() bool { return false }
func (LampRatedHours) SceneIndex() int  { return -1 }

func (LampRatedHours) Name() string                  { return "Lamp Rated Hours" }
func (a *LampRatedHours) TypeID() zcl.TypeID         { return zcl.Zu24(*a).ID() }
func (a *LampRatedHours) Value() zcl.Val             { return a }
func (a LampRatedHours) MarshalZcl() ([]byte, error) { return zcl.Zu24(a).MarshalZcl() }

func (a *LampRatedHours) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*a = LampRatedHours(*nt)
	return br, err
}

func (a *LampRatedHours) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu24); ok {
		*a = LampRatedHours(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LampRatedHours) String() string {
	return zcl.Hours.Format(float64(a))
}

type LampType zcl.Zcstring

const LampTypeAttr zcl.AttrID = 48

func (LampType) ID() zcl.AttrID   { return LampTypeAttr }
func (LampType) Readable() bool   { return true }
func (LampType) Writable() bool   { return true }
func (LampType) Reportable() bool { return false }
func (LampType) SceneIndex() int  { return -1 }

func (LampType) Name() string                  { return "Lamp Type" }
func (a *LampType) TypeID() zcl.TypeID         { return zcl.Zcstring(*a).ID() }
func (a *LampType) Value() zcl.Val             { return a }
func (a LampType) MarshalZcl() ([]byte, error) { return zcl.Zcstring(a).MarshalZcl() }

func (a *LampType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = LampType(*nt)
	return br, err
}

func (a *LampType) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zcstring); ok {
		*a = LampType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a LampType) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

// MaxLevel specifies the maximum light level the ballast is permitted to use.
// It is specified in the range 0x01 to 0xfe, and specifies the light
// output of the ballast according to the dimming light curve. The value
// is both less than or equal to PhysicalMaxLevel and greater than or equal
// to MinLevel
type MaxLevel zcl.Zu8

const MaxLevelAttr zcl.AttrID = 17

func (MaxLevel) ID() zcl.AttrID   { return MaxLevelAttr }
func (MaxLevel) Readable() bool   { return true }
func (MaxLevel) Writable() bool   { return true }
func (MaxLevel) Reportable() bool { return false }
func (MaxLevel) SceneIndex() int  { return -1 }

func (MaxLevel) Name() string                  { return "Max Level" }
func (a *MaxLevel) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *MaxLevel) Value() zcl.Val             { return a }
func (a MaxLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *MaxLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = MaxLevel(*nt)
	return br, err
}

func (a *MaxLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = MaxLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MaxLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type MinLevel zcl.Zu8

const MinLevelAttr zcl.AttrID = 16

func (MinLevel) ID() zcl.AttrID   { return MinLevelAttr }
func (MinLevel) Readable() bool   { return true }
func (MinLevel) Writable() bool   { return true }
func (MinLevel) Reportable() bool { return false }
func (MinLevel) SceneIndex() int  { return -1 }

func (MinLevel) Name() string                  { return "Min Level" }
func (a *MinLevel) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *MinLevel) Value() zcl.Val             { return a }
func (a MinLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *MinLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = MinLevel(*nt)
	return br, err
}

func (a *MinLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = MinLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MinLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type MoveMode zcl.Zenum8

func (MoveMode) Name() string                  { return "Move mode" }
func (a *MoveMode) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *MoveMode) Value() zcl.Val             { return a }
func (a MoveMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *MoveMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = MoveMode(*nt)
	return br, err
}

func (a *MoveMode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = MoveMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a MoveMode) String() string {
	switch a {
	case 0x00:
		return "Stop"
	case 0x01:
		return "Up"
	case 0x03:
		return "Down"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a MoveMode) IsStop() bool { return a == 0x00 }
func (a MoveMode) IsUp() bool   { return a == 0x01 }
func (a MoveMode) IsDown() bool { return a == 0x03 }
func (a *MoveMode) SetStop()    { *a = 0x00 }
func (a *MoveMode) SetUp()      { *a = 0x01 }
func (a *MoveMode) SetDown()    { *a = 0x03 }

func (MoveMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Stop"},
		{Value: 0x01, Name: "Up"},
		{Value: 0x03, Name: "Down"},
	}
}

// NumberOfPrimaries contains the number of color primaries implemented on this device.
// A value of 0xff indicates that the number of primaries is unknown
type NumberOfPrimaries zcl.Zu8

const NumberOfPrimariesAttr zcl.AttrID = 16

func (NumberOfPrimaries) ID() zcl.AttrID   { return NumberOfPrimariesAttr }
func (NumberOfPrimaries) Readable() bool   { return true }
func (NumberOfPrimaries) Writable() bool   { return false }
func (NumberOfPrimaries) Reportable() bool { return false }
func (NumberOfPrimaries) SceneIndex() int  { return -1 }

func (NumberOfPrimaries) Name() string                  { return "Number of primaries" }
func (a *NumberOfPrimaries) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *NumberOfPrimaries) Value() zcl.Val             { return a }
func (a NumberOfPrimaries) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *NumberOfPrimaries) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = NumberOfPrimaries(*nt)
	return br, err
}

func (a *NumberOfPrimaries) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = NumberOfPrimaries(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a NumberOfPrimaries) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// PhysicalMaxLevel specifies the maximum light level the ballast can achieve. It is
// specified in the range 0x01 to 0xfe, and specifies the light output
// of the ballast according to the dimming light curve
type PhysicalMaxLevel zcl.Zu8

const PhysicalMaxLevelAttr zcl.AttrID = 1

func (PhysicalMaxLevel) ID() zcl.AttrID   { return PhysicalMaxLevelAttr }
func (PhysicalMaxLevel) Readable() bool   { return true }
func (PhysicalMaxLevel) Writable() bool   { return false }
func (PhysicalMaxLevel) Reportable() bool { return false }
func (PhysicalMaxLevel) SceneIndex() int  { return -1 }

func (PhysicalMaxLevel) Name() string                  { return "Physical Max Level" }
func (a *PhysicalMaxLevel) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *PhysicalMaxLevel) Value() zcl.Val             { return a }
func (a PhysicalMaxLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *PhysicalMaxLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PhysicalMaxLevel(*nt)
	return br, err
}

func (a *PhysicalMaxLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = PhysicalMaxLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PhysicalMaxLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// PhysicalMinLevel specifies the minimum light level the ballast can achieve. It is
// specified in the range 0x01 to 0xfe, and specifies the light output
// of the ballast according to the dimming light curve
type PhysicalMinLevel zcl.Zu8

const PhysicalMinLevelAttr zcl.AttrID = 0

func (PhysicalMinLevel) ID() zcl.AttrID   { return PhysicalMinLevelAttr }
func (PhysicalMinLevel) Readable() bool   { return true }
func (PhysicalMinLevel) Writable() bool   { return false }
func (PhysicalMinLevel) Reportable() bool { return false }
func (PhysicalMinLevel) SceneIndex() int  { return -1 }

func (PhysicalMinLevel) Name() string                  { return "Physical Min Level" }
func (a *PhysicalMinLevel) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *PhysicalMinLevel) Value() zcl.Val             { return a }
func (a PhysicalMinLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *PhysicalMinLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PhysicalMinLevel(*nt)
	return br, err
}

func (a *PhysicalMinLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = PhysicalMinLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PhysicalMinLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// PowerOnLevel It specifies the light level to which the ballast will go when power
// is applied (e.g., when mains power is re-established after a power
// failure). It can have a value between 0x00 and 0xfe to set it to
// a specific light level, according to the dimming light curve, or
// 0xff to set it to the value it was before the power failure
type PowerOnLevel zcl.Zu8

const PowerOnLevelAttr zcl.AttrID = 18

func (PowerOnLevel) ID() zcl.AttrID   { return PowerOnLevelAttr }
func (PowerOnLevel) Readable() bool   { return true }
func (PowerOnLevel) Writable() bool   { return true }
func (PowerOnLevel) Reportable() bool { return false }
func (PowerOnLevel) SceneIndex() int  { return -1 }

func (PowerOnLevel) Name() string                  { return "Power On level" }
func (a *PowerOnLevel) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *PowerOnLevel) Value() zcl.Val             { return a }
func (a PowerOnLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *PowerOnLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerOnLevel(*nt)
	return br, err
}

func (a *PowerOnLevel) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = PowerOnLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerOnLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// PowerOnTime The transition time in 1/10ths of a second.
type PowerOnTime zcl.Zu16

const PowerOnTimeAttr zcl.AttrID = 19

func (PowerOnTime) ID() zcl.AttrID   { return PowerOnTimeAttr }
func (PowerOnTime) Readable() bool   { return true }
func (PowerOnTime) Writable() bool   { return true }
func (PowerOnTime) Reportable() bool { return false }
func (PowerOnTime) SceneIndex() int  { return -1 }

func (PowerOnTime) Name() string                  { return "Power-On Time" }
func (a *PowerOnTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *PowerOnTime) Value() zcl.Val             { return a }
func (a PowerOnTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *PowerOnTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerOnTime(*nt)
	return br, err
}

func (a *PowerOnTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = PowerOnTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerOnTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type PowerOnColorTemperature zcl.Zu16

const PowerOnColorTemperatureAttr zcl.AttrID = 16400

func (PowerOnColorTemperature) ID() zcl.AttrID   { return PowerOnColorTemperatureAttr }
func (PowerOnColorTemperature) Readable() bool   { return true }
func (PowerOnColorTemperature) Writable() bool   { return true }
func (PowerOnColorTemperature) Reportable() bool { return false }
func (PowerOnColorTemperature) SceneIndex() int  { return -1 }

func (PowerOnColorTemperature) Name() string                  { return "Power-on color temperature" }
func (a *PowerOnColorTemperature) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *PowerOnColorTemperature) Value() zcl.Val             { return a }
func (a PowerOnColorTemperature) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *PowerOnColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = PowerOnColorTemperature(*nt)
	return br, err
}

func (a *PowerOnColorTemperature) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = PowerOnColorTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a PowerOnColorTemperature) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary1X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary1X zcl.Zu16

const Primary1XAttr zcl.AttrID = 17

func (Primary1X) ID() zcl.AttrID   { return Primary1XAttr }
func (Primary1X) Readable() bool   { return true }
func (Primary1X) Writable() bool   { return false }
func (Primary1X) Reportable() bool { return false }
func (Primary1X) SceneIndex() int  { return -1 }

func (Primary1X) Name() string                  { return "Primary1 X" }
func (a *Primary1X) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary1X) Value() zcl.Val             { return a }
func (a Primary1X) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary1X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary1X(*nt)
	return br, err
}

func (a *Primary1X) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary1X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary1X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary1Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary1Y zcl.Zu16

const Primary1YAttr zcl.AttrID = 18

func (Primary1Y) ID() zcl.AttrID   { return Primary1YAttr }
func (Primary1Y) Readable() bool   { return true }
func (Primary1Y) Writable() bool   { return false }
func (Primary1Y) Reportable() bool { return false }
func (Primary1Y) SceneIndex() int  { return -1 }

func (Primary1Y) Name() string                  { return "Primary1 Y" }
func (a *Primary1Y) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary1Y) Value() zcl.Val             { return a }
func (a Primary1Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary1Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary1Y(*nt)
	return br, err
}

func (a *Primary1Y) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary1Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary1Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary1Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary1Intensity zcl.Zu8

const Primary1IntensityAttr zcl.AttrID = 19

func (Primary1Intensity) ID() zcl.AttrID   { return Primary1IntensityAttr }
func (Primary1Intensity) Readable() bool   { return true }
func (Primary1Intensity) Writable() bool   { return false }
func (Primary1Intensity) Reportable() bool { return false }
func (Primary1Intensity) SceneIndex() int  { return -1 }

func (Primary1Intensity) Name() string                  { return "Primary1 intensity" }
func (a *Primary1Intensity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Primary1Intensity) Value() zcl.Val             { return a }
func (a Primary1Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Primary1Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary1Intensity(*nt)
	return br, err
}

func (a *Primary1Intensity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Primary1Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary1Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// Primary2X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary2X zcl.Zu16

const Primary2XAttr zcl.AttrID = 21

func (Primary2X) ID() zcl.AttrID   { return Primary2XAttr }
func (Primary2X) Readable() bool   { return true }
func (Primary2X) Writable() bool   { return false }
func (Primary2X) Reportable() bool { return false }
func (Primary2X) SceneIndex() int  { return -1 }

func (Primary2X) Name() string                  { return "Primary2 X" }
func (a *Primary2X) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary2X) Value() zcl.Val             { return a }
func (a Primary2X) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary2X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary2X(*nt)
	return br, err
}

func (a *Primary2X) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary2X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary2X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary2Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary2Y zcl.Zu16

const Primary2YAttr zcl.AttrID = 22

func (Primary2Y) ID() zcl.AttrID   { return Primary2YAttr }
func (Primary2Y) Readable() bool   { return true }
func (Primary2Y) Writable() bool   { return false }
func (Primary2Y) Reportable() bool { return false }
func (Primary2Y) SceneIndex() int  { return -1 }

func (Primary2Y) Name() string                  { return "Primary2 Y" }
func (a *Primary2Y) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary2Y) Value() zcl.Val             { return a }
func (a Primary2Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary2Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary2Y(*nt)
	return br, err
}

func (a *Primary2Y) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary2Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary2Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary2Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary2Intensity zcl.Zu8

const Primary2IntensityAttr zcl.AttrID = 23

func (Primary2Intensity) ID() zcl.AttrID   { return Primary2IntensityAttr }
func (Primary2Intensity) Readable() bool   { return true }
func (Primary2Intensity) Writable() bool   { return false }
func (Primary2Intensity) Reportable() bool { return false }
func (Primary2Intensity) SceneIndex() int  { return -1 }

func (Primary2Intensity) Name() string                  { return "Primary2 intensity" }
func (a *Primary2Intensity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Primary2Intensity) Value() zcl.Val             { return a }
func (a Primary2Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Primary2Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary2Intensity(*nt)
	return br, err
}

func (a *Primary2Intensity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Primary2Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary2Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// Primary3X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary3X zcl.Zu16

const Primary3XAttr zcl.AttrID = 25

func (Primary3X) ID() zcl.AttrID   { return Primary3XAttr }
func (Primary3X) Readable() bool   { return true }
func (Primary3X) Writable() bool   { return false }
func (Primary3X) Reportable() bool { return false }
func (Primary3X) SceneIndex() int  { return -1 }

func (Primary3X) Name() string                  { return "Primary3 X" }
func (a *Primary3X) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary3X) Value() zcl.Val             { return a }
func (a Primary3X) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary3X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary3X(*nt)
	return br, err
}

func (a *Primary3X) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary3X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary3X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary3Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary3Y zcl.Zu16

const Primary3YAttr zcl.AttrID = 26

func (Primary3Y) ID() zcl.AttrID   { return Primary3YAttr }
func (Primary3Y) Readable() bool   { return true }
func (Primary3Y) Writable() bool   { return false }
func (Primary3Y) Reportable() bool { return false }
func (Primary3Y) SceneIndex() int  { return -1 }

func (Primary3Y) Name() string                  { return "Primary3 Y" }
func (a *Primary3Y) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary3Y) Value() zcl.Val             { return a }
func (a Primary3Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary3Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary3Y(*nt)
	return br, err
}

func (a *Primary3Y) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary3Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary3Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary3Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary3Intensity zcl.Zu8

const Primary3IntensityAttr zcl.AttrID = 27

func (Primary3Intensity) ID() zcl.AttrID   { return Primary3IntensityAttr }
func (Primary3Intensity) Readable() bool   { return true }
func (Primary3Intensity) Writable() bool   { return false }
func (Primary3Intensity) Reportable() bool { return false }
func (Primary3Intensity) SceneIndex() int  { return -1 }

func (Primary3Intensity) Name() string                  { return "Primary3 intensity" }
func (a *Primary3Intensity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Primary3Intensity) Value() zcl.Val             { return a }
func (a Primary3Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Primary3Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary3Intensity(*nt)
	return br, err
}

func (a *Primary3Intensity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Primary3Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary3Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// Primary4X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary4X zcl.Zu16

const Primary4XAttr zcl.AttrID = 32

func (Primary4X) ID() zcl.AttrID   { return Primary4XAttr }
func (Primary4X) Readable() bool   { return true }
func (Primary4X) Writable() bool   { return false }
func (Primary4X) Reportable() bool { return false }
func (Primary4X) SceneIndex() int  { return -1 }

func (Primary4X) Name() string                  { return "Primary4 X" }
func (a *Primary4X) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary4X) Value() zcl.Val             { return a }
func (a Primary4X) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary4X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary4X(*nt)
	return br, err
}

func (a *Primary4X) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary4X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary4X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary4Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary4Y zcl.Zu16

const Primary4YAttr zcl.AttrID = 33

func (Primary4Y) ID() zcl.AttrID   { return Primary4YAttr }
func (Primary4Y) Readable() bool   { return true }
func (Primary4Y) Writable() bool   { return false }
func (Primary4Y) Reportable() bool { return false }
func (Primary4Y) SceneIndex() int  { return -1 }

func (Primary4Y) Name() string                  { return "Primary4 Y" }
func (a *Primary4Y) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary4Y) Value() zcl.Val             { return a }
func (a Primary4Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary4Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary4Y(*nt)
	return br, err
}

func (a *Primary4Y) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary4Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary4Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary4Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary4Intensity zcl.Zu8

const Primary4IntensityAttr zcl.AttrID = 34

func (Primary4Intensity) ID() zcl.AttrID   { return Primary4IntensityAttr }
func (Primary4Intensity) Readable() bool   { return true }
func (Primary4Intensity) Writable() bool   { return false }
func (Primary4Intensity) Reportable() bool { return false }
func (Primary4Intensity) SceneIndex() int  { return -1 }

func (Primary4Intensity) Name() string                  { return "Primary4 intensity" }
func (a *Primary4Intensity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Primary4Intensity) Value() zcl.Val             { return a }
func (a Primary4Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Primary4Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary4Intensity(*nt)
	return br, err
}

func (a *Primary4Intensity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Primary4Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary4Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// Primary5X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary5X zcl.Zu16

const Primary5XAttr zcl.AttrID = 36

func (Primary5X) ID() zcl.AttrID   { return Primary5XAttr }
func (Primary5X) Readable() bool   { return true }
func (Primary5X) Writable() bool   { return false }
func (Primary5X) Reportable() bool { return false }
func (Primary5X) SceneIndex() int  { return -1 }

func (Primary5X) Name() string                  { return "Primary5 X" }
func (a *Primary5X) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary5X) Value() zcl.Val             { return a }
func (a Primary5X) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary5X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary5X(*nt)
	return br, err
}

func (a *Primary5X) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary5X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary5X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary5Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary5Y zcl.Zu16

const Primary5YAttr zcl.AttrID = 37

func (Primary5Y) ID() zcl.AttrID   { return Primary5YAttr }
func (Primary5Y) Readable() bool   { return true }
func (Primary5Y) Writable() bool   { return false }
func (Primary5Y) Reportable() bool { return false }
func (Primary5Y) SceneIndex() int  { return -1 }

func (Primary5Y) Name() string                  { return "Primary5 Y" }
func (a *Primary5Y) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary5Y) Value() zcl.Val             { return a }
func (a Primary5Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary5Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary5Y(*nt)
	return br, err
}

func (a *Primary5Y) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary5Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary5Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary5Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary5Intensity zcl.Zu8

const Primary5IntensityAttr zcl.AttrID = 38

func (Primary5Intensity) ID() zcl.AttrID   { return Primary5IntensityAttr }
func (Primary5Intensity) Readable() bool   { return true }
func (Primary5Intensity) Writable() bool   { return false }
func (Primary5Intensity) Reportable() bool { return false }
func (Primary5Intensity) SceneIndex() int  { return -1 }

func (Primary5Intensity) Name() string                  { return "Primary5 intensity" }
func (a *Primary5Intensity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Primary5Intensity) Value() zcl.Val             { return a }
func (a Primary5Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Primary5Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary5Intensity(*nt)
	return br, err
}

func (a *Primary5Intensity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Primary5Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary5Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// Primary6X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary6X zcl.Zu16

const Primary6XAttr zcl.AttrID = 40

func (Primary6X) ID() zcl.AttrID   { return Primary6XAttr }
func (Primary6X) Readable() bool   { return true }
func (Primary6X) Writable() bool   { return false }
func (Primary6X) Reportable() bool { return false }
func (Primary6X) SceneIndex() int  { return -1 }

func (Primary6X) Name() string                  { return "Primary6 X" }
func (a *Primary6X) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary6X) Value() zcl.Val             { return a }
func (a Primary6X) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary6X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary6X(*nt)
	return br, err
}

func (a *Primary6X) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary6X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary6X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary6Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary6Y zcl.Zu16

const Primary6YAttr zcl.AttrID = 41

func (Primary6Y) ID() zcl.AttrID   { return Primary6YAttr }
func (Primary6Y) Readable() bool   { return true }
func (Primary6Y) Writable() bool   { return false }
func (Primary6Y) Reportable() bool { return false }
func (Primary6Y) SceneIndex() int  { return -1 }

func (Primary6Y) Name() string                  { return "Primary6 Y" }
func (a *Primary6Y) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Primary6Y) Value() zcl.Val             { return a }
func (a Primary6Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Primary6Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary6Y(*nt)
	return br, err
}

func (a *Primary6Y) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Primary6Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary6Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// Primary6Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary6Intensity zcl.Zu8

const Primary6IntensityAttr zcl.AttrID = 42

func (Primary6Intensity) ID() zcl.AttrID   { return Primary6IntensityAttr }
func (Primary6Intensity) Readable() bool   { return true }
func (Primary6Intensity) Writable() bool   { return false }
func (Primary6Intensity) Reportable() bool { return false }
func (Primary6Intensity) SceneIndex() int  { return -1 }

func (Primary6Intensity) Name() string                  { return "Primary6 intensity" }
func (a *Primary6Intensity) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Primary6Intensity) Value() zcl.Val             { return a }
func (a Primary6Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Primary6Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Primary6Intensity(*nt)
	return br, err
}

func (a *Primary6Intensity) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Primary6Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Primary6Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// Rate increment/steps per second
type Rate zcl.Zu8

func (Rate) Name() string                  { return "Rate" }
func (a *Rate) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Rate) Value() zcl.Val             { return a }
func (a Rate) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Rate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Rate(*nt)
	return br, err
}

func (a *Rate) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Rate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Rate) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// RateX increment/steps per second
type RateX zcl.Zs16

func (RateX) Name() string                  { return "Rate X" }
func (a *RateX) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *RateX) Value() zcl.Val             { return a }
func (a RateX) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *RateX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RateX(*nt)
	return br, err
}

func (a *RateX) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = RateX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a RateX) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// RateY increment/steps per second
type RateY zcl.Zs16

func (RateY) Name() string                  { return "Rate Y" }
func (a *RateY) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *RateY) Value() zcl.Val             { return a }
func (a RateY) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *RateY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = RateY(*nt)
	return br, err
}

func (a *RateY) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = RateY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a RateY) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

// RemainingTime holds the time remaining, in 1/10ths of a second, until the currently
// active command will be complete
type RemainingTime zcl.Zu16

const RemainingTimeAttr zcl.AttrID = 2

func (RemainingTime) ID() zcl.AttrID   { return RemainingTimeAttr }
func (RemainingTime) Readable() bool   { return true }
func (RemainingTime) Writable() bool   { return false }
func (RemainingTime) Reportable() bool { return false }
func (RemainingTime) SceneIndex() int  { return -1 }

func (RemainingTime) Name() string                  { return "Remaining time" }
func (a *RemainingTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *RemainingTime) Value() zcl.Val             { return a }
func (a RemainingTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *RemainingTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = RemainingTime(*nt)
	return br, err
}

func (a *RemainingTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = RemainingTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a RemainingTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type Saturation zcl.Zu8

func (Saturation) Name() string                  { return "Saturation" }
func (a *Saturation) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *Saturation) Value() zcl.Val             { return a }
func (a Saturation) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *Saturation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Saturation(*nt)
	return br, err
}

func (a *Saturation) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = Saturation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Saturation) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

type StepX zcl.Zs16

func (StepX) Name() string                  { return "Step X" }
func (a *StepX) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *StepX) Value() zcl.Val             { return a }
func (a StepX) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *StepX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = StepX(*nt)
	return br, err
}

func (a *StepX) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = StepX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StepX) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

type StepY zcl.Zs16

func (StepY) Name() string                  { return "Step Y" }
func (a *StepY) TypeID() zcl.TypeID         { return zcl.Zs16(*a).ID() }
func (a *StepY) Value() zcl.Val             { return a }
func (a StepY) MarshalZcl() ([]byte, error) { return zcl.Zs16(a).MarshalZcl() }

func (a *StepY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*a = StepY(*nt)
	return br, err
}

func (a *StepY) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zs16); ok {
		*a = StepY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StepY) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(a))
}

type StepMode zcl.Zenum8

func (StepMode) Name() string                  { return "Step mode" }
func (a *StepMode) TypeID() zcl.TypeID         { return zcl.Zenum8(*a).ID() }
func (a *StepMode) Value() zcl.Val             { return a }
func (a StepMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(a).MarshalZcl() }

func (a *StepMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = StepMode(*nt)
	return br, err
}

func (a *StepMode) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zenum8); ok {
		*a = StepMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StepMode) String() string {
	switch a {
	case 0x01:
		return "Up"
	case 0x03:
		return "Down"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

func (a StepMode) IsUp() bool   { return a == 0x01 }
func (a StepMode) IsDown() bool { return a == 0x03 }
func (a *StepMode) SetUp()      { *a = 0x01 }
func (a *StepMode) SetDown()    { *a = 0x03 }

func (StepMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x01, Name: "Up"},
		{Value: 0x03, Name: "Down"},
	}
}

type StepSize zcl.Zu8

func (StepSize) Name() string                  { return "Step size" }
func (a *StepSize) TypeID() zcl.TypeID         { return zcl.Zu8(*a).ID() }
func (a *StepSize) Value() zcl.Val             { return a }
func (a StepSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(a).MarshalZcl() }

func (a *StepSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = StepSize(*nt)
	return br, err
}

func (a *StepSize) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu8); ok {
		*a = StepSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a StepSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(a))
}

// Time Time in seconds used for a whole color loop.
type Time zcl.Zu16

func (Time) Name() string                  { return "Time" }
func (a *Time) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *Time) Value() zcl.Val             { return a }
func (a Time) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *Time) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Time(*nt)
	return br, err
}

func (a *Time) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = Time(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a Time) String() string {
	return zcl.Seconds.Format(float64(a))
}

// TransitionTime The transition time in 1/10ths of a second.
type TransitionTime zcl.Zu16

func (TransitionTime) Name() string                  { return "Transition time" }
func (a *TransitionTime) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *TransitionTime) Value() zcl.Val             { return a }
func (a TransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *TransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = TransitionTime(*nt)
	return br, err
}

func (a *TransitionTime) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = TransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a TransitionTime) String() string {
	return zcl.Seconds.Format(float64(a) / 10)
}

type UpdateFlags zcl.Zbmp8

func (UpdateFlags) Name() string                  { return "Update flags" }
func (a *UpdateFlags) TypeID() zcl.TypeID         { return zcl.Zbmp8(*a).ID() }
func (a *UpdateFlags) Value() zcl.Val             { return a }
func (a UpdateFlags) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(a).MarshalZcl() }

func (a *UpdateFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = UpdateFlags(*nt)
	return br, err
}

func (a *UpdateFlags) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zbmp8); ok {
		*a = UpdateFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a UpdateFlags) String() string {
	var bstr []string
	bits := zcl.BitmapList(a[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Update action")
		case 1:
			bstr = append(bstr, "Update direction")
		case 2:
			bstr = append(bstr, "Update time")
		case 3:
			bstr = append(bstr, "Update start hue")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (a UpdateFlags) IsUpdateAction() bool       { return zcl.BitmapTest([]byte(a[:]), 0) }
func (a UpdateFlags) IsUpdateDirection() bool    { return zcl.BitmapTest([]byte(a[:]), 1) }
func (a UpdateFlags) IsUpdateTime() bool         { return zcl.BitmapTest([]byte(a[:]), 2) }
func (a UpdateFlags) IsUpdateStartHue() bool     { return zcl.BitmapTest([]byte(a[:]), 3) }
func (a *UpdateFlags) SetUpdateAction(b bool)    { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 0, b)) }
func (a *UpdateFlags) SetUpdateDirection(b bool) { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 1, b)) }
func (a *UpdateFlags) SetUpdateTime(b bool)      { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 2, b)) }
func (a *UpdateFlags) SetUpdateStartHue(b bool)  { copy((*a)[:], zcl.BitmapSet([]byte((*a)[:]), 3, b)) }

func (UpdateFlags) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Update action"},
		{Value: 1, Name: "Update direction"},
		{Value: 2, Name: "Update time"},
		{Value: 3, Name: "Update start hue"},
	}
}

// WhitePointX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type WhitePointX zcl.Zu16

const WhitePointXAttr zcl.AttrID = 48

func (WhitePointX) ID() zcl.AttrID   { return WhitePointXAttr }
func (WhitePointX) Readable() bool   { return true }
func (WhitePointX) Writable() bool   { return true }
func (WhitePointX) Reportable() bool { return false }
func (WhitePointX) SceneIndex() int  { return -1 }

func (WhitePointX) Name() string                  { return "White Point X" }
func (a *WhitePointX) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *WhitePointX) Value() zcl.Val             { return a }
func (a WhitePointX) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *WhitePointX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = WhitePointX(*nt)
	return br, err
}

func (a *WhitePointX) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = WhitePointX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a WhitePointX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}

// WhitePointY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type WhitePointY zcl.Zu16

const WhitePointYAttr zcl.AttrID = 49

func (WhitePointY) ID() zcl.AttrID   { return WhitePointYAttr }
func (WhitePointY) Readable() bool   { return true }
func (WhitePointY) Writable() bool   { return true }
func (WhitePointY) Reportable() bool { return false }
func (WhitePointY) SceneIndex() int  { return -1 }

func (WhitePointY) Name() string                  { return "White Point Y" }
func (a *WhitePointY) TypeID() zcl.TypeID         { return zcl.Zu16(*a).ID() }
func (a *WhitePointY) Value() zcl.Val             { return a }
func (a WhitePointY) MarshalZcl() ([]byte, error) { return zcl.Zu16(a).MarshalZcl() }

func (a *WhitePointY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = WhitePointY(*nt)
	return br, err
}

func (a *WhitePointY) SetValue(v zcl.Val) error {
	if nv, ok := v.(*zcl.Zu16); ok {
		*a = WhitePointY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (a WhitePointY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(a))
}
