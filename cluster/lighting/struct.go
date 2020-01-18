package lighting

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

func (Action) Name() string        { return `Action` }
func (Action) Description() string { return `` }

type Action zcl.Zenum8

func (v *Action) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *Action) Value() zcl.Val     { return v }

func (v Action) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *Action) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = Action(*nt)
	return br, err
}

func (v Action) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *Action) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Action(*a)
	return nil
}

func (v *Action) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = Action(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Action) String() string {
	switch v {
	case 0x00:
		return "De-activate color loop"
	case 0x01:
		return "Activate from ColorLoopStartEnhancedHue"
	case 0x02:
		return "Activate from EnhancedCurrentHue"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v Action) IsDeActivateColorLoop() bool                   { return v == 0x00 }
func (v Action) IsActivateFromColorloopstartenhancedhue() bool { return v == 0x01 }
func (v Action) IsActivateFromEnhancedcurrenthue() bool        { return v == 0x02 }
func (v *Action) SetDeActivateColorLoop()                      { *v = 0x00 }
func (v *Action) SetActivateFromColorloopstartenhancedhue()    { *v = 0x01 }
func (v *Action) SetActivateFromEnhancedcurrenthue()           { *v = 0x02 }

func (Action) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "De-activate color loop"},
		{Value: 0x01, Name: "Activate from ColorLoopStartEnhancedHue"},
		{Value: 0x02, Name: "Activate from EnhancedCurrentHue"},
	}
}

const BallastFactorAdjustmentAttr zcl.AttrID = 21

func (BallastFactorAdjustment) ID() zcl.AttrID   { return BallastFactorAdjustmentAttr }
func (BallastFactorAdjustment) Readable() bool   { return true }
func (BallastFactorAdjustment) Writable() bool   { return true }
func (BallastFactorAdjustment) Reportable() bool { return false }
func (BallastFactorAdjustment) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BallastFactorAdjustment) AttrID() zcl.AttrID   { return v.ID() }
func (v BallastFactorAdjustment) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BallastFactorAdjustment) AttrValue() zcl.Val  { return v.Value() }

func (BallastFactorAdjustment) Name() string { return `Ballast Factor Adjustment` }
func (BallastFactorAdjustment) Description() string {
	return `specifies the multiplication factor, as a percentage, to be applied
to the configured light output of the lamps. A typical usage of this
mechanism is to compensate for reduction in efficiency over the lifetime
of a lamp. The light output is given by
Actual light output = configured light output x BallastFactorAdjustment / 100%
The range for this attribute is manufacturer dependent`
}

// BallastFactorAdjustment specifies the multiplication factor, as a percentage, to be applied
// to the configured light output of the lamps. A typical usage of this
// mechanism is to compensate for reduction in efficiency over the lifetime
// of a lamp. The light output is given by
// Actual light output = configured light output x BallastFactorAdjustment / 100%
// The range for this attribute is manufacturer dependent
type BallastFactorAdjustment zcl.Zu8

func (v *BallastFactorAdjustment) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *BallastFactorAdjustment) Value() zcl.Val     { return v }

func (v BallastFactorAdjustment) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *BallastFactorAdjustment) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = BallastFactorAdjustment(*nt)
	return br, err
}

func (v BallastFactorAdjustment) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *BallastFactorAdjustment) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BallastFactorAdjustment(*a)
	return nil
}

func (v *BallastFactorAdjustment) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = BallastFactorAdjustment(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BallastFactorAdjustment) String() string {
	return zcl.Percent.Format(float64(v))
}

const BallastStatusAttr zcl.AttrID = 2

func (BallastStatus) ID() zcl.AttrID   { return BallastStatusAttr }
func (BallastStatus) Readable() bool   { return true }
func (BallastStatus) Writable() bool   { return false }
func (BallastStatus) Reportable() bool { return false }
func (BallastStatus) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v BallastStatus) AttrID() zcl.AttrID   { return v.ID() }
func (v BallastStatus) AttrType() zcl.TypeID { return v.TypeID() }
func (v *BallastStatus) AttrValue() zcl.Val  { return v.Value() }

func (BallastStatus) Name() string { return `Ballast Status` }
func (BallastStatus) Description() string {
	return `It specifies the activity status of the ballast functions. Where a
function is active, the corresponding bit is 1. Where a function is
not active, the corresponding bit is 0. This means that bit 0 with
a value of 0 means the ballast is operational and bit 1 with a
value of 0 that the lamp(s) is/are in their socket(s)`
}

// BallastStatus It specifies the activity status of the ballast functions. Where a
// function is active, the corresponding bit is 1. Where a function is
// not active, the corresponding bit is 0. This means that bit 0 with
// a value of 0 means the ballast is operational and bit 1 with a
// value of 0 that the lamp(s) is/are in their socket(s)
type BallastStatus zcl.Zbmp8

func (v *BallastStatus) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *BallastStatus) Value() zcl.Val     { return v }

func (v BallastStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *BallastStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = BallastStatus(*nt)
	return br, err
}

func (v BallastStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *BallastStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = BallastStatus(*a)
	return nil
}

func (v *BallastStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = BallastStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v BallastStatus) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v BallastStatus) IsNonOperational() bool    { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v BallastStatus) IsLampNotInSocket() bool   { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v *BallastStatus) SetNonOperational(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *BallastStatus) SetLampNotInSocket(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}

func (BallastStatus) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Non-operational"},
		{Value: 1, Name: "Lamp not in socket"},
	}
}

const ColorModeAttr zcl.AttrID = 8

func (ColorMode) ID() zcl.AttrID   { return ColorModeAttr }
func (ColorMode) Readable() bool   { return true }
func (ColorMode) Writable() bool   { return false }
func (ColorMode) Reportable() bool { return false }
func (ColorMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorMode) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorMode) AttrValue() zcl.Val  { return v.Value() }

func (ColorMode) Name() string { return `Color Mode` }
func (ColorMode) Description() string {
	return `indicates which attributes are currently determining the color of
the device. This attribute is optional if the device does not implement
CurrentHue and CurrentSaturation`
}

// ColorMode indicates which attributes are currently determining the color of
// the device. This attribute is optional if the device does not implement
// CurrentHue and CurrentSaturation
type ColorMode zcl.Zenum8

func (v *ColorMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *ColorMode) Value() zcl.Val     { return v }

func (v ColorMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *ColorMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorMode(*nt)
	return br, err
}

func (v ColorMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *ColorMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorMode(*a)
	return nil
}

func (v *ColorMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = ColorMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorMode) String() string {
	switch v {
	case 0x00:
		return "Current hue and current saturation"
	case 0x01:
		return "Current X and Current Y"
	case 0x02:
		return "Color temperature"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v ColorMode) IsCurrentHueAndCurrentSaturation() bool { return v == 0x00 }
func (v ColorMode) IsCurrentXAndCurrentY() bool            { return v == 0x01 }
func (v ColorMode) IsColorTemperature() bool               { return v == 0x02 }
func (v *ColorMode) SetCurrentHueAndCurrentSaturation()    { *v = 0x00 }
func (v *ColorMode) SetCurrentXAndCurrentY()               { *v = 0x01 }
func (v *ColorMode) SetColorTemperature()                  { *v = 0x02 }

func (ColorMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Current hue and current saturation"},
		{Value: 0x01, Name: "Current X and Current Y"},
		{Value: 0x02, Name: "Color temperature"},
	}
}

const ColorPointBlueXAttr zcl.AttrID = 58

func (ColorPointBlueX) ID() zcl.AttrID   { return ColorPointBlueXAttr }
func (ColorPointBlueX) Readable() bool   { return true }
func (ColorPointBlueX) Writable() bool   { return true }
func (ColorPointBlueX) Reportable() bool { return false }
func (ColorPointBlueX) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorPointBlueX) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorPointBlueX) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorPointBlueX) AttrValue() zcl.Val  { return v.Value() }

func (ColorPointBlueX) Name() string { return `Color Point Blue X` }
func (ColorPointBlueX) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// ColorPointBlueX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointBlueX zcl.Zu16

func (v *ColorPointBlueX) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorPointBlueX) Value() zcl.Val     { return v }

func (v ColorPointBlueX) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorPointBlueX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorPointBlueX(*nt)
	return br, err
}

func (v ColorPointBlueX) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorPointBlueX) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorPointBlueX(*a)
	return nil
}

func (v *ColorPointBlueX) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorPointBlueX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorPointBlueX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ColorPointBlueYAttr zcl.AttrID = 59

func (ColorPointBlueY) ID() zcl.AttrID   { return ColorPointBlueYAttr }
func (ColorPointBlueY) Readable() bool   { return true }
func (ColorPointBlueY) Writable() bool   { return true }
func (ColorPointBlueY) Reportable() bool { return false }
func (ColorPointBlueY) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorPointBlueY) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorPointBlueY) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorPointBlueY) AttrValue() zcl.Val  { return v.Value() }

func (ColorPointBlueY) Name() string { return `Color Point Blue Y` }
func (ColorPointBlueY) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// ColorPointBlueY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointBlueY zcl.Zu16

func (v *ColorPointBlueY) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorPointBlueY) Value() zcl.Val     { return v }

func (v ColorPointBlueY) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorPointBlueY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorPointBlueY(*nt)
	return br, err
}

func (v ColorPointBlueY) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorPointBlueY) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorPointBlueY(*a)
	return nil
}

func (v *ColorPointBlueY) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorPointBlueY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorPointBlueY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ColorPointBlueIntensityAttr zcl.AttrID = 60

func (ColorPointBlueIntensity) ID() zcl.AttrID   { return ColorPointBlueIntensityAttr }
func (ColorPointBlueIntensity) Readable() bool   { return true }
func (ColorPointBlueIntensity) Writable() bool   { return true }
func (ColorPointBlueIntensity) Reportable() bool { return false }
func (ColorPointBlueIntensity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorPointBlueIntensity) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorPointBlueIntensity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorPointBlueIntensity) AttrValue() zcl.Val  { return v.Value() }

func (ColorPointBlueIntensity) Name() string { return `Color Point Blue intensity` }
func (ColorPointBlueIntensity) Description() string {
	return `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`
}

// ColorPointBlueIntensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type ColorPointBlueIntensity zcl.Zu8

func (v *ColorPointBlueIntensity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ColorPointBlueIntensity) Value() zcl.Val     { return v }

func (v ColorPointBlueIntensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ColorPointBlueIntensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorPointBlueIntensity(*nt)
	return br, err
}

func (v ColorPointBlueIntensity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ColorPointBlueIntensity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorPointBlueIntensity(*a)
	return nil
}

func (v *ColorPointBlueIntensity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ColorPointBlueIntensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorPointBlueIntensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const ColorPointGreenXAttr zcl.AttrID = 54

func (ColorPointGreenX) ID() zcl.AttrID   { return ColorPointGreenXAttr }
func (ColorPointGreenX) Readable() bool   { return true }
func (ColorPointGreenX) Writable() bool   { return true }
func (ColorPointGreenX) Reportable() bool { return false }
func (ColorPointGreenX) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorPointGreenX) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorPointGreenX) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorPointGreenX) AttrValue() zcl.Val  { return v.Value() }

func (ColorPointGreenX) Name() string { return `Color Point Green X` }
func (ColorPointGreenX) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// ColorPointGreenX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointGreenX zcl.Zu16

func (v *ColorPointGreenX) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorPointGreenX) Value() zcl.Val     { return v }

func (v ColorPointGreenX) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorPointGreenX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorPointGreenX(*nt)
	return br, err
}

func (v ColorPointGreenX) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorPointGreenX) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorPointGreenX(*a)
	return nil
}

func (v *ColorPointGreenX) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorPointGreenX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorPointGreenX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ColorPointGreenYAttr zcl.AttrID = 55

func (ColorPointGreenY) ID() zcl.AttrID   { return ColorPointGreenYAttr }
func (ColorPointGreenY) Readable() bool   { return true }
func (ColorPointGreenY) Writable() bool   { return true }
func (ColorPointGreenY) Reportable() bool { return false }
func (ColorPointGreenY) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorPointGreenY) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorPointGreenY) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorPointGreenY) AttrValue() zcl.Val  { return v.Value() }

func (ColorPointGreenY) Name() string { return `Color Point Green Y` }
func (ColorPointGreenY) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// ColorPointGreenY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointGreenY zcl.Zu16

func (v *ColorPointGreenY) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorPointGreenY) Value() zcl.Val     { return v }

func (v ColorPointGreenY) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorPointGreenY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorPointGreenY(*nt)
	return br, err
}

func (v ColorPointGreenY) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorPointGreenY) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorPointGreenY(*a)
	return nil
}

func (v *ColorPointGreenY) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorPointGreenY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorPointGreenY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ColorPointGreenIntensityAttr zcl.AttrID = 56

func (ColorPointGreenIntensity) ID() zcl.AttrID   { return ColorPointGreenIntensityAttr }
func (ColorPointGreenIntensity) Readable() bool   { return true }
func (ColorPointGreenIntensity) Writable() bool   { return true }
func (ColorPointGreenIntensity) Reportable() bool { return false }
func (ColorPointGreenIntensity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorPointGreenIntensity) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorPointGreenIntensity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorPointGreenIntensity) AttrValue() zcl.Val  { return v.Value() }

func (ColorPointGreenIntensity) Name() string { return `Color Point Green intensity` }
func (ColorPointGreenIntensity) Description() string {
	return `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`
}

// ColorPointGreenIntensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type ColorPointGreenIntensity zcl.Zu8

func (v *ColorPointGreenIntensity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ColorPointGreenIntensity) Value() zcl.Val     { return v }

func (v ColorPointGreenIntensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ColorPointGreenIntensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorPointGreenIntensity(*nt)
	return br, err
}

func (v ColorPointGreenIntensity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ColorPointGreenIntensity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorPointGreenIntensity(*a)
	return nil
}

func (v *ColorPointGreenIntensity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ColorPointGreenIntensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorPointGreenIntensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const ColorPointRedXAttr zcl.AttrID = 50

func (ColorPointRedX) ID() zcl.AttrID   { return ColorPointRedXAttr }
func (ColorPointRedX) Readable() bool   { return true }
func (ColorPointRedX) Writable() bool   { return true }
func (ColorPointRedX) Reportable() bool { return false }
func (ColorPointRedX) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorPointRedX) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorPointRedX) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorPointRedX) AttrValue() zcl.Val  { return v.Value() }

func (ColorPointRedX) Name() string { return `Color Point Red X` }
func (ColorPointRedX) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// ColorPointRedX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointRedX zcl.Zu16

func (v *ColorPointRedX) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorPointRedX) Value() zcl.Val     { return v }

func (v ColorPointRedX) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorPointRedX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorPointRedX(*nt)
	return br, err
}

func (v ColorPointRedX) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorPointRedX) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorPointRedX(*a)
	return nil
}

func (v *ColorPointRedX) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorPointRedX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorPointRedX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ColorPointRedYAttr zcl.AttrID = 51

func (ColorPointRedY) ID() zcl.AttrID   { return ColorPointRedYAttr }
func (ColorPointRedY) Readable() bool   { return true }
func (ColorPointRedY) Writable() bool   { return true }
func (ColorPointRedY) Reportable() bool { return false }
func (ColorPointRedY) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorPointRedY) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorPointRedY) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorPointRedY) AttrValue() zcl.Val  { return v.Value() }

func (ColorPointRedY) Name() string { return `Color Point Red Y` }
func (ColorPointRedY) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// ColorPointRedY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorPointRedY zcl.Zu16

func (v *ColorPointRedY) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorPointRedY) Value() zcl.Val     { return v }

func (v ColorPointRedY) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorPointRedY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorPointRedY(*nt)
	return br, err
}

func (v ColorPointRedY) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorPointRedY) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorPointRedY(*a)
	return nil
}

func (v *ColorPointRedY) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorPointRedY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorPointRedY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ColorPointRedIntensityAttr zcl.AttrID = 52

func (ColorPointRedIntensity) ID() zcl.AttrID   { return ColorPointRedIntensityAttr }
func (ColorPointRedIntensity) Readable() bool   { return true }
func (ColorPointRedIntensity) Writable() bool   { return true }
func (ColorPointRedIntensity) Reportable() bool { return false }
func (ColorPointRedIntensity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorPointRedIntensity) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorPointRedIntensity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorPointRedIntensity) AttrValue() zcl.Val  { return v.Value() }

func (ColorPointRedIntensity) Name() string { return `Color Point Red intensity` }
func (ColorPointRedIntensity) Description() string {
	return `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`
}

// ColorPointRedIntensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type ColorPointRedIntensity zcl.Zu8

func (v *ColorPointRedIntensity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ColorPointRedIntensity) Value() zcl.Val     { return v }

func (v ColorPointRedIntensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ColorPointRedIntensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorPointRedIntensity(*nt)
	return br, err
}

func (v ColorPointRedIntensity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ColorPointRedIntensity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorPointRedIntensity(*a)
	return nil
}

func (v *ColorPointRedIntensity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ColorPointRedIntensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorPointRedIntensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (ColorTemperature) Name() string        { return `Color Temperature` }
func (ColorTemperature) Description() string { return `` }

type ColorTemperature zcl.Zu16

func (v *ColorTemperature) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorTemperature) Value() zcl.Val     { return v }

func (v ColorTemperature) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorTemperature(*nt)
	return br, err
}

func (v ColorTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorTemperature(*a)
	return nil
}

func (v *ColorTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorTemperature) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (ColorTemperatureMax) Name() string        { return `Color Temperature max` }
func (ColorTemperatureMax) Description() string { return `` }

type ColorTemperatureMax zcl.Zu16

func (v *ColorTemperatureMax) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorTemperatureMax) Value() zcl.Val     { return v }

func (v ColorTemperatureMax) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorTemperatureMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorTemperatureMax(*nt)
	return br, err
}

func (v ColorTemperatureMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorTemperatureMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorTemperatureMax(*a)
	return nil
}

func (v *ColorTemperatureMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorTemperatureMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorTemperatureMax) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (ColorTemperatureMaxMireds) Name() string        { return `Color Temperature max Mireds` }
func (ColorTemperatureMaxMireds) Description() string { return `` }

type ColorTemperatureMaxMireds zcl.Zu16

func (v *ColorTemperatureMaxMireds) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorTemperatureMaxMireds) Value() zcl.Val     { return v }

func (v ColorTemperatureMaxMireds) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorTemperatureMaxMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorTemperatureMaxMireds(*nt)
	return br, err
}

func (v ColorTemperatureMaxMireds) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorTemperatureMaxMireds) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorTemperatureMaxMireds(*a)
	return nil
}

func (v *ColorTemperatureMaxMireds) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorTemperatureMaxMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorTemperatureMaxMireds) String() string {
	return zcl.Mired.Format(float64(v))
}

func (ColorTemperatureMin) Name() string        { return `Color Temperature min` }
func (ColorTemperatureMin) Description() string { return `` }

type ColorTemperatureMin zcl.Zu16

func (v *ColorTemperatureMin) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorTemperatureMin) Value() zcl.Val     { return v }

func (v ColorTemperatureMin) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorTemperatureMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorTemperatureMin(*nt)
	return br, err
}

func (v ColorTemperatureMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorTemperatureMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorTemperatureMin(*a)
	return nil
}

func (v *ColorTemperatureMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorTemperatureMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorTemperatureMin) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (ColorTemperatureMinMireds) Name() string        { return `Color Temperature min Mireds` }
func (ColorTemperatureMinMireds) Description() string { return `` }

type ColorTemperatureMinMireds zcl.Zu16

func (v *ColorTemperatureMinMireds) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorTemperatureMinMireds) Value() zcl.Val     { return v }

func (v ColorTemperatureMinMireds) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorTemperatureMinMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorTemperatureMinMireds(*nt)
	return br, err
}

func (v ColorTemperatureMinMireds) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorTemperatureMinMireds) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorTemperatureMinMireds(*a)
	return nil
}

func (v *ColorTemperatureMinMireds) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorTemperatureMinMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorTemperatureMinMireds) String() string {
	return zcl.Mired.Format(float64(v))
}

func (ColorX) Name() string { return `Color X` }
func (ColorX) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// ColorX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorX zcl.Zu16

func (v *ColorX) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorX) Value() zcl.Val     { return v }

func (v ColorX) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorX(*nt)
	return br, err
}

func (v ColorX) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorX) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorX(*a)
	return nil
}

func (v *ColorX) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (ColorY) Name() string { return `Color Y` }
func (ColorY) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// ColorY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type ColorY zcl.Zu16

func (v *ColorY) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorY) Value() zcl.Val     { return v }

func (v ColorY) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorY(*nt)
	return br, err
}

func (v ColorY) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorY) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorY(*a)
	return nil
}

func (v *ColorY) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ColorCapabilitiesAttr zcl.AttrID = 16394

func (ColorCapabilities) ID() zcl.AttrID   { return ColorCapabilitiesAttr }
func (ColorCapabilities) Readable() bool   { return true }
func (ColorCapabilities) Writable() bool   { return false }
func (ColorCapabilities) Reportable() bool { return false }
func (ColorCapabilities) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorCapabilities) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorCapabilities) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorCapabilities) AttrValue() zcl.Val  { return v.Value() }

func (ColorCapabilities) Name() string { return `Color capabilities` }
func (ColorCapabilities) Description() string {
	return `specifies the color capabilities of the device supporting the color
control cluster`
}

// ColorCapabilities specifies the color capabilities of the device supporting the color
// control cluster
type ColorCapabilities zcl.Zbmp16

func (v *ColorCapabilities) TypeID() zcl.TypeID { return new(zcl.Zbmp16).TypeID() }
func (v *ColorCapabilities) Value() zcl.Val     { return v }

func (v ColorCapabilities) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(v).MarshalZcl() }

func (v *ColorCapabilities) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorCapabilities(*nt)
	return br, err
}

func (v ColorCapabilities) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp16(v))
}

func (v *ColorCapabilities) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorCapabilities(*a)
	return nil
}

func (v *ColorCapabilities) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp16); ok {
		*v = ColorCapabilities(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorCapabilities) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v ColorCapabilities) IsHueSaturation() bool         { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v ColorCapabilities) IsEnhancedHueSaturation() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v ColorCapabilities) IsColorLoop() bool             { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v ColorCapabilities) IsCie1931Xy() bool             { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v ColorCapabilities) IsColorTemperature() bool      { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v *ColorCapabilities) SetHueSaturation(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *ColorCapabilities) SetEnhancedHueSaturation(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *ColorCapabilities) SetColorLoop(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b)) }
func (v *ColorCapabilities) SetCie1931Xy(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b)) }
func (v *ColorCapabilities) SetColorTemperature(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
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

const ColorControlOptionsAttr zcl.AttrID = 15

func (ColorControlOptions) ID() zcl.AttrID   { return ColorControlOptionsAttr }
func (ColorControlOptions) Readable() bool   { return true }
func (ColorControlOptions) Writable() bool   { return true }
func (ColorControlOptions) Reportable() bool { return false }
func (ColorControlOptions) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorControlOptions) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorControlOptions) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorControlOptions) AttrValue() zcl.Val  { return v.Value() }

func (ColorControlOptions) Name() string { return `Color control options` }
func (ColorControlOptions) Description() string {
	return `is a bitmap that determines the default behavior of some cluster commands`
}

// ColorControlOptions is a bitmap that determines the default behavior of some cluster commands
type ColorControlOptions zcl.Zbmp8

func (v *ColorControlOptions) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *ColorControlOptions) Value() zcl.Val     { return v }

func (v ColorControlOptions) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *ColorControlOptions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorControlOptions(*nt)
	return br, err
}

func (v ColorControlOptions) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *ColorControlOptions) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorControlOptions(*a)
	return nil
}

func (v *ColorControlOptions) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = ColorControlOptions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorControlOptions) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v ColorControlOptions) IsExecuteIfOff() bool { return zcl.BitmapTest([]byte(v[:]), 0x00) }
func (v *ColorControlOptions) SetExecuteIfOff(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0x00, b))
}

func (ColorControlOptions) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Execute if off"},
	}
}

const ColorLoopActiveAttr zcl.AttrID = 16386

func (ColorLoopActive) ID() zcl.AttrID   { return ColorLoopActiveAttr }
func (ColorLoopActive) Readable() bool   { return true }
func (ColorLoopActive) Writable() bool   { return false }
func (ColorLoopActive) Reportable() bool { return false }
func (ColorLoopActive) SceneIndex() int  { return 5 }

// Implements AttrDef/AttrValue interfaces
func (v ColorLoopActive) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorLoopActive) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorLoopActive) AttrValue() zcl.Val  { return v.Value() }

func (ColorLoopActive) Name() string { return `Color loop active` }
func (ColorLoopActive) Description() string {
	return `specifies the current active status of the color loop. 0x00 means
inactive, 0x01 means active`
}

// ColorLoopActive specifies the current active status of the color loop. 0x00 means
// inactive, 0x01 means active
type ColorLoopActive zcl.Zu8

func (v *ColorLoopActive) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ColorLoopActive) Value() zcl.Val     { return v }

func (v ColorLoopActive) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ColorLoopActive) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorLoopActive(*nt)
	return br, err
}

func (v ColorLoopActive) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ColorLoopActive) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorLoopActive(*a)
	return nil
}

func (v *ColorLoopActive) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ColorLoopActive(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorLoopActive) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const ColorLoopDirectionAttr zcl.AttrID = 16387

func (ColorLoopDirection) ID() zcl.AttrID   { return ColorLoopDirectionAttr }
func (ColorLoopDirection) Readable() bool   { return true }
func (ColorLoopDirection) Writable() bool   { return false }
func (ColorLoopDirection) Reportable() bool { return false }
func (ColorLoopDirection) SceneIndex() int  { return 6 }

// Implements AttrDef/AttrValue interfaces
func (v ColorLoopDirection) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorLoopDirection) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorLoopDirection) AttrValue() zcl.Val  { return v.Value() }

func (ColorLoopDirection) Name() string { return `Color loop direction` }
func (ColorLoopDirection) Description() string {
	return `specifies the current direction of the color loop. If this attribute
has the value 0x00, the EnhancedCurrentHue is be decremented. If this
attribute has the value 0x01, the EnhancedCurrentHue is incremented`
}

// ColorLoopDirection specifies the current direction of the color loop. If this attribute
// has the value 0x00, the EnhancedCurrentHue is be decremented. If this
// attribute has the value 0x01, the EnhancedCurrentHue is incremented
type ColorLoopDirection zcl.Zu8

func (v *ColorLoopDirection) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *ColorLoopDirection) Value() zcl.Val     { return v }

func (v ColorLoopDirection) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *ColorLoopDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorLoopDirection(*nt)
	return br, err
}

func (v ColorLoopDirection) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *ColorLoopDirection) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorLoopDirection(*a)
	return nil
}

func (v *ColorLoopDirection) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = ColorLoopDirection(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorLoopDirection) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const ColorLoopStartEnhancedHueAttr zcl.AttrID = 16389

func (ColorLoopStartEnhancedHue) ID() zcl.AttrID   { return ColorLoopStartEnhancedHueAttr }
func (ColorLoopStartEnhancedHue) Readable() bool   { return true }
func (ColorLoopStartEnhancedHue) Writable() bool   { return false }
func (ColorLoopStartEnhancedHue) Reportable() bool { return false }
func (ColorLoopStartEnhancedHue) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorLoopStartEnhancedHue) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorLoopStartEnhancedHue) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorLoopStartEnhancedHue) AttrValue() zcl.Val  { return v.Value() }

func (ColorLoopStartEnhancedHue) Name() string { return `Color loop start enhanced hue` }
func (ColorLoopStartEnhancedHue) Description() string {
	return `specifies the value of the EnhancedCurrentHue attribute from which
the color loop starts`
}

// ColorLoopStartEnhancedHue specifies the value of the EnhancedCurrentHue attribute from which
// the color loop starts
type ColorLoopStartEnhancedHue zcl.Zu16

func (v *ColorLoopStartEnhancedHue) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorLoopStartEnhancedHue) Value() zcl.Val     { return v }

func (v ColorLoopStartEnhancedHue) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorLoopStartEnhancedHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorLoopStartEnhancedHue(*nt)
	return br, err
}

func (v ColorLoopStartEnhancedHue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorLoopStartEnhancedHue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorLoopStartEnhancedHue(*a)
	return nil
}

func (v *ColorLoopStartEnhancedHue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorLoopStartEnhancedHue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorLoopStartEnhancedHue) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ColorLoopStoredEnhancedHueAttr zcl.AttrID = 16390

func (ColorLoopStoredEnhancedHue) ID() zcl.AttrID   { return ColorLoopStoredEnhancedHueAttr }
func (ColorLoopStoredEnhancedHue) Readable() bool   { return true }
func (ColorLoopStoredEnhancedHue) Writable() bool   { return false }
func (ColorLoopStoredEnhancedHue) Reportable() bool { return false }
func (ColorLoopStoredEnhancedHue) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorLoopStoredEnhancedHue) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorLoopStoredEnhancedHue) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorLoopStoredEnhancedHue) AttrValue() zcl.Val  { return v.Value() }

func (ColorLoopStoredEnhancedHue) Name() string { return `Color loop stored enhanced hue` }
func (ColorLoopStoredEnhancedHue) Description() string {
	return `specifies the value of the EnhancedCurrentHue attribute before the
color loop was started. Once the color loop is complete, It is restored
to this value`
}

// ColorLoopStoredEnhancedHue specifies the value of the EnhancedCurrentHue attribute before the
// color loop was started. Once the color loop is complete, It is restored
// to this value
type ColorLoopStoredEnhancedHue zcl.Zu16

func (v *ColorLoopStoredEnhancedHue) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorLoopStoredEnhancedHue) Value() zcl.Val     { return v }

func (v ColorLoopStoredEnhancedHue) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorLoopStoredEnhancedHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorLoopStoredEnhancedHue(*nt)
	return br, err
}

func (v ColorLoopStoredEnhancedHue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorLoopStoredEnhancedHue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorLoopStoredEnhancedHue(*a)
	return nil
}

func (v *ColorLoopStoredEnhancedHue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorLoopStoredEnhancedHue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorLoopStoredEnhancedHue) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ColorLoopTimeAttr zcl.AttrID = 16388

func (ColorLoopTime) ID() zcl.AttrID   { return ColorLoopTimeAttr }
func (ColorLoopTime) Readable() bool   { return true }
func (ColorLoopTime) Writable() bool   { return false }
func (ColorLoopTime) Reportable() bool { return false }
func (ColorLoopTime) SceneIndex() int  { return 7 }

// Implements AttrDef/AttrValue interfaces
func (v ColorLoopTime) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorLoopTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorLoopTime) AttrValue() zcl.Val  { return v.Value() }

func (ColorLoopTime) Name() string { return `Color loop time` }
func (ColorLoopTime) Description() string {
	return `specifies the number of seconds it takes to perform a full color
loop, i.e., to cycle all values of EnhancedCurrentHue`
}

// ColorLoopTime specifies the number of seconds it takes to perform a full color
// loop, i.e., to cycle all values of EnhancedCurrentHue
type ColorLoopTime zcl.Zu16

func (v *ColorLoopTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorLoopTime) Value() zcl.Val     { return v }

func (v ColorLoopTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorLoopTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorLoopTime(*nt)
	return br, err
}

func (v ColorLoopTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorLoopTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorLoopTime(*a)
	return nil
}

func (v *ColorLoopTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorLoopTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorLoopTime) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ColorTemperatureMiredsAttr zcl.AttrID = 7

func (ColorTemperatureMireds) ID() zcl.AttrID   { return ColorTemperatureMiredsAttr }
func (ColorTemperatureMireds) Readable() bool   { return true }
func (ColorTemperatureMireds) Writable() bool   { return false }
func (ColorTemperatureMireds) Reportable() bool { return true }
func (ColorTemperatureMireds) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorTemperatureMireds) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorTemperatureMireds) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorTemperatureMireds) AttrValue() zcl.Val  { return v.Value() }

func (ColorTemperatureMireds) Name() string { return `Color temperature Mireds` }
func (ColorTemperatureMireds) Description() string {
	return `contains a scaled inverse of the current value of the color
temperature. The unit of ColorTemperatureMireds is the mired
(micro reciprocal degree), a.k.a mirek (micro reciprocal
kelvin). Color temperature in kelvins = 1,000,000 / ColorTemperatureMireds,
where ColorTemperatureMireds is in the range 1 to 65279 mireds inclusive,
giving a color temperature range from 1,000,000 kelvins to 15.32 kelvins`
}

// ColorTemperatureMireds contains a scaled inverse of the current value of the color
// temperature. The unit of ColorTemperatureMireds is the mired
// (micro reciprocal degree), a.k.a mirek (micro reciprocal
// kelvin). Color temperature in kelvins = 1,000,000 / ColorTemperatureMireds,
// where ColorTemperatureMireds is in the range 1 to 65279 mireds inclusive,
// giving a color temperature range from 1,000,000 kelvins to 15.32 kelvins
type ColorTemperatureMireds zcl.Zu16

func (v *ColorTemperatureMireds) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorTemperatureMireds) Value() zcl.Val     { return v }

func (v ColorTemperatureMireds) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ColorTemperatureMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorTemperatureMireds(*nt)
	return br, err
}

func (v ColorTemperatureMireds) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorTemperatureMireds) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorTemperatureMireds(*a)
	return nil
}

func (v *ColorTemperatureMireds) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorTemperatureMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorTemperatureMireds) String() string {
	return zcl.Mired.Format(float64(v))
}

const ColorTemperaturePhysicalMaxMiredsAttr zcl.AttrID = 16396

func (ColorTemperaturePhysicalMaxMireds) ID() zcl.AttrID   { return ColorTemperaturePhysicalMaxMiredsAttr }
func (ColorTemperaturePhysicalMaxMireds) Readable() bool   { return true }
func (ColorTemperaturePhysicalMaxMireds) Writable() bool   { return false }
func (ColorTemperaturePhysicalMaxMireds) Reportable() bool { return false }
func (ColorTemperaturePhysicalMaxMireds) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorTemperaturePhysicalMaxMireds) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorTemperaturePhysicalMaxMireds) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorTemperaturePhysicalMaxMireds) AttrValue() zcl.Val  { return v.Value() }

func (ColorTemperaturePhysicalMaxMireds) Name() string { return `Color temperature physical max Mireds` }
func (ColorTemperaturePhysicalMaxMireds) Description() string {
	return `indicates the maximum mired value supported by the hardware.
ColorTempPhysicalMaxMireds corresponds to the minimum color
temperature in Kelvins supported by the hardware.
ColorTemperatureMireds ≤ ColorTempPhysicalMaxMireds`
}

// ColorTemperaturePhysicalMaxMireds indicates the maximum mired value supported by the hardware.
// ColorTempPhysicalMaxMireds corresponds to the minimum color
// temperature in Kelvins supported by the hardware.
// ColorTemperatureMireds ≤ ColorTempPhysicalMaxMireds
type ColorTemperaturePhysicalMaxMireds zcl.Zu16

func (v *ColorTemperaturePhysicalMaxMireds) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorTemperaturePhysicalMaxMireds) Value() zcl.Val     { return v }

func (v ColorTemperaturePhysicalMaxMireds) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *ColorTemperaturePhysicalMaxMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorTemperaturePhysicalMaxMireds(*nt)
	return br, err
}

func (v ColorTemperaturePhysicalMaxMireds) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorTemperaturePhysicalMaxMireds) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorTemperaturePhysicalMaxMireds(*a)
	return nil
}

func (v *ColorTemperaturePhysicalMaxMireds) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorTemperaturePhysicalMaxMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorTemperaturePhysicalMaxMireds) String() string {
	return zcl.Mired.Format(float64(v))
}

const ColorTemperaturePhysicalMinMiredsAttr zcl.AttrID = 16395

func (ColorTemperaturePhysicalMinMireds) ID() zcl.AttrID   { return ColorTemperaturePhysicalMinMiredsAttr }
func (ColorTemperaturePhysicalMinMireds) Readable() bool   { return true }
func (ColorTemperaturePhysicalMinMireds) Writable() bool   { return false }
func (ColorTemperaturePhysicalMinMireds) Reportable() bool { return false }
func (ColorTemperaturePhysicalMinMireds) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ColorTemperaturePhysicalMinMireds) AttrID() zcl.AttrID   { return v.ID() }
func (v ColorTemperaturePhysicalMinMireds) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ColorTemperaturePhysicalMinMireds) AttrValue() zcl.Val  { return v.Value() }

func (ColorTemperaturePhysicalMinMireds) Name() string { return `Color temperature physical min Mireds` }
func (ColorTemperaturePhysicalMinMireds) Description() string {
	return `indicates the minimum mired value supported by the hardware.
ColorTempPhysicalMinMireds corresponds to the maximum color
temperature in Kelvins supported by the hardware.
ColorTempPhysicalMinMireds ≤ ColorTemperatureMireds`
}

// ColorTemperaturePhysicalMinMireds indicates the minimum mired value supported by the hardware.
// ColorTempPhysicalMinMireds corresponds to the maximum color
// temperature in Kelvins supported by the hardware.
// ColorTempPhysicalMinMireds ≤ ColorTemperatureMireds
type ColorTemperaturePhysicalMinMireds zcl.Zu16

func (v *ColorTemperaturePhysicalMinMireds) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ColorTemperaturePhysicalMinMireds) Value() zcl.Val     { return v }

func (v ColorTemperaturePhysicalMinMireds) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *ColorTemperaturePhysicalMinMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ColorTemperaturePhysicalMinMireds(*nt)
	return br, err
}

func (v ColorTemperaturePhysicalMinMireds) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ColorTemperaturePhysicalMinMireds) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ColorTemperaturePhysicalMinMireds(*a)
	return nil
}

func (v *ColorTemperaturePhysicalMinMireds) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ColorTemperaturePhysicalMinMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ColorTemperaturePhysicalMinMireds) String() string {
	return zcl.Mired.Format(float64(v))
}

const CompensationTextAttr zcl.AttrID = 6

func (CompensationText) ID() zcl.AttrID   { return CompensationTextAttr }
func (CompensationText) Readable() bool   { return true }
func (CompensationText) Writable() bool   { return false }
func (CompensationText) Reportable() bool { return false }
func (CompensationText) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CompensationText) AttrID() zcl.AttrID   { return v.ID() }
func (v CompensationText) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CompensationText) AttrValue() zcl.Val  { return v.Value() }

func (CompensationText) Name() string { return `Compensation Text` }
func (CompensationText) Description() string {
	return `holds a textual indication of what mechanism, if any, is in use to
compensate for color/intensity drift over time`
}

// CompensationText holds a textual indication of what mechanism, if any, is in use to
// compensate for color/intensity drift over time
type CompensationText zcl.Zcstring

func (v *CompensationText) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *CompensationText) Value() zcl.Val     { return v }

func (v CompensationText) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *CompensationText) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = CompensationText(*nt)
	return br, err
}

func (v CompensationText) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *CompensationText) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CompensationText(*a)
	return nil
}

func (v *CompensationText) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = CompensationText(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CompensationText) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const CoupleColorTempToLevelMinMiredsAttr zcl.AttrID = 16397

func (CoupleColorTempToLevelMinMireds) ID() zcl.AttrID   { return CoupleColorTempToLevelMinMiredsAttr }
func (CoupleColorTempToLevelMinMireds) Readable() bool   { return true }
func (CoupleColorTempToLevelMinMireds) Writable() bool   { return false }
func (CoupleColorTempToLevelMinMireds) Reportable() bool { return false }
func (CoupleColorTempToLevelMinMireds) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CoupleColorTempToLevelMinMireds) AttrID() zcl.AttrID   { return v.ID() }
func (v CoupleColorTempToLevelMinMireds) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CoupleColorTempToLevelMinMireds) AttrValue() zcl.Val  { return v.Value() }

func (CoupleColorTempToLevelMinMireds) Name() string { return `Couple Color Temp to Level Min Mireds` }
func (CoupleColorTempToLevelMinMireds) Description() string {
	return `specifies a lower bound on the value of the ColorTemperatureMireds attribute for the purposes of coupling the
ColorTemperatureMireds attribute to the CurrentLevel attribute when the CoupleColorTempToLevel bit of the
Options attribute of the Level Control cluster is equal to 1`
}

// CoupleColorTempToLevelMinMireds specifies a lower bound on the value of the ColorTemperatureMireds attribute for the purposes of coupling the
// ColorTemperatureMireds attribute to the CurrentLevel attribute when the CoupleColorTempToLevel bit of the
// Options attribute of the Level Control cluster is equal to 1
type CoupleColorTempToLevelMinMireds zcl.Zu16

func (v *CoupleColorTempToLevelMinMireds) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *CoupleColorTempToLevelMinMireds) Value() zcl.Val     { return v }

func (v CoupleColorTempToLevelMinMireds) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *CoupleColorTempToLevelMinMireds) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = CoupleColorTempToLevelMinMireds(*nt)
	return br, err
}

func (v CoupleColorTempToLevelMinMireds) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *CoupleColorTempToLevelMinMireds) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CoupleColorTempToLevelMinMireds(*a)
	return nil
}

func (v *CoupleColorTempToLevelMinMireds) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = CoupleColorTempToLevelMinMireds(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CoupleColorTempToLevelMinMireds) String() string {
	return zcl.Mired.Format(float64(v))
}

const CurrentXAttr zcl.AttrID = 3

func (CurrentX) ID() zcl.AttrID   { return CurrentXAttr }
func (CurrentX) Readable() bool   { return true }
func (CurrentX) Writable() bool   { return false }
func (CurrentX) Reportable() bool { return true }
func (CurrentX) SceneIndex() int  { return 1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentX) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentX) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentX) AttrValue() zcl.Val  { return v.Value() }

func (CurrentX) Name() string { return `Current X` }
func (CurrentX) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// CurrentX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type CurrentX zcl.Zu16

func (v *CurrentX) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *CurrentX) Value() zcl.Val     { return v }

func (v CurrentX) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *CurrentX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentX(*nt)
	return br, err
}

func (v CurrentX) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *CurrentX) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentX(*a)
	return nil
}

func (v *CurrentX) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = CurrentX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const CurrentYAttr zcl.AttrID = 4

func (CurrentY) ID() zcl.AttrID   { return CurrentYAttr }
func (CurrentY) Readable() bool   { return true }
func (CurrentY) Writable() bool   { return false }
func (CurrentY) Reportable() bool { return true }
func (CurrentY) SceneIndex() int  { return 2 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentY) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentY) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentY) AttrValue() zcl.Val  { return v.Value() }

func (CurrentY) Name() string { return `Current Y` }
func (CurrentY) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// CurrentY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type CurrentY zcl.Zu16

func (v *CurrentY) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *CurrentY) Value() zcl.Val     { return v }

func (v CurrentY) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *CurrentY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentY(*nt)
	return br, err
}

func (v CurrentY) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *CurrentY) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentY(*a)
	return nil
}

func (v *CurrentY) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = CurrentY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const CurrentHueAttr zcl.AttrID = 0

func (CurrentHue) ID() zcl.AttrID   { return CurrentHueAttr }
func (CurrentHue) Readable() bool   { return true }
func (CurrentHue) Writable() bool   { return false }
func (CurrentHue) Reportable() bool { return true }
func (CurrentHue) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentHue) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentHue) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentHue) AttrValue() zcl.Val  { return v.Value() }

func (CurrentHue) Name() string { return `Current hue` }
func (CurrentHue) Description() string {
	return `contains the current hue value of the light. Hue = CurrentHue x 360 / 254
(CurrentHue in the range 0 - 254 inclusive)`
}

// CurrentHue contains the current hue value of the light. Hue = CurrentHue x 360 / 254
// (CurrentHue in the range 0 - 254 inclusive)
type CurrentHue zcl.Zu8

func (v *CurrentHue) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *CurrentHue) Value() zcl.Val     { return v }

func (v CurrentHue) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *CurrentHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentHue(*nt)
	return br, err
}

func (v CurrentHue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *CurrentHue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentHue(*a)
	return nil
}

func (v *CurrentHue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = CurrentHue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentHue) String() string {
	return zcl.DegreesAngular.Format(float64(v) / 0.70556)
}

const CurrentSaturationAttr zcl.AttrID = 1

func (CurrentSaturation) ID() zcl.AttrID   { return CurrentSaturationAttr }
func (CurrentSaturation) Readable() bool   { return true }
func (CurrentSaturation) Writable() bool   { return false }
func (CurrentSaturation) Reportable() bool { return true }
func (CurrentSaturation) SceneIndex() int  { return 4 }

// Implements AttrDef/AttrValue interfaces
func (v CurrentSaturation) AttrID() zcl.AttrID   { return v.ID() }
func (v CurrentSaturation) AttrType() zcl.TypeID { return v.TypeID() }
func (v *CurrentSaturation) AttrValue() zcl.Val  { return v.Value() }

func (CurrentSaturation) Name() string { return `Current saturation` }
func (CurrentSaturation) Description() string {
	return `holds the current saturation value of the light.
Saturation = CurrentSaturation/254 (CurrentSaturation in the range
0 - 254 inclusive)`
}

// CurrentSaturation holds the current saturation value of the light.
// Saturation = CurrentSaturation/254 (CurrentSaturation in the range
// 0 - 254 inclusive)
type CurrentSaturation zcl.Zu8

func (v *CurrentSaturation) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *CurrentSaturation) Value() zcl.Val     { return v }

func (v CurrentSaturation) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *CurrentSaturation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = CurrentSaturation(*nt)
	return br, err
}

func (v CurrentSaturation) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *CurrentSaturation) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = CurrentSaturation(*a)
	return nil
}

func (v *CurrentSaturation) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = CurrentSaturation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v CurrentSaturation) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const DriftCompensationAttr zcl.AttrID = 5

func (DriftCompensation) ID() zcl.AttrID   { return DriftCompensationAttr }
func (DriftCompensation) Readable() bool   { return true }
func (DriftCompensation) Writable() bool   { return false }
func (DriftCompensation) Reportable() bool { return false }
func (DriftCompensation) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DriftCompensation) AttrID() zcl.AttrID   { return v.ID() }
func (v DriftCompensation) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DriftCompensation) AttrValue() zcl.Val  { return v.Value() }

func (DriftCompensation) Name() string { return `Drift Compensation` }
func (DriftCompensation) Description() string {
	return `indicates what mechanism, if any, is in use for compensation for
color/intensity drift over time`
}

// DriftCompensation indicates what mechanism, if any, is in use for compensation for
// color/intensity drift over time
type DriftCompensation zcl.Zenum8

func (v *DriftCompensation) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *DriftCompensation) Value() zcl.Val     { return v }

func (v DriftCompensation) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *DriftCompensation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = DriftCompensation(*nt)
	return br, err
}

func (v DriftCompensation) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *DriftCompensation) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DriftCompensation(*a)
	return nil
}

func (v *DriftCompensation) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = DriftCompensation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DriftCompensation) String() string {
	switch v {
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
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v DriftCompensation) IsNone() bool                                  { return v == 0x00 }
func (v DriftCompensation) IsOtherUnknown() bool                          { return v == 0x01 }
func (v DriftCompensation) IsTemperatureMonitoring() bool                 { return v == 0x02 }
func (v DriftCompensation) IsOpticalLuminanceMonitoringAndFeedback() bool { return v == 0x03 }
func (v DriftCompensation) IsOpticalColorMonitoringAndFeedback() bool     { return v == 0x04 }
func (v *DriftCompensation) SetNone()                                     { *v = 0x00 }
func (v *DriftCompensation) SetOtherUnknown()                             { *v = 0x01 }
func (v *DriftCompensation) SetTemperatureMonitoring()                    { *v = 0x02 }
func (v *DriftCompensation) SetOpticalLuminanceMonitoringAndFeedback()    { *v = 0x03 }
func (v *DriftCompensation) SetOpticalColorMonitoringAndFeedback()        { *v = 0x04 }

func (DriftCompensation) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "None"},
		{Value: 0x01, Name: "Other / Unknown"},
		{Value: 0x02, Name: "Temperature monitoring"},
		{Value: 0x03, Name: "Optical luminance monitoring and feedback"},
		{Value: 0x04, Name: "Optical color monitoring and feedback"},
	}
}

func (EnhancedHue) Name() string        { return `Enhanced Hue` }
func (EnhancedHue) Description() string { return `` }

type EnhancedHue zcl.Zu16

func (v *EnhancedHue) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *EnhancedHue) Value() zcl.Val     { return v }

func (v EnhancedHue) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *EnhancedHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = EnhancedHue(*nt)
	return br, err
}

func (v EnhancedHue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *EnhancedHue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = EnhancedHue(*a)
	return nil
}

func (v *EnhancedHue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = EnhancedHue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EnhancedHue) String() string {
	return zcl.DegreesAngular.Format(float64(v) / 0.002756094)
}

const EnhancedColorModeAttr zcl.AttrID = 16385

func (EnhancedColorMode) ID() zcl.AttrID   { return EnhancedColorModeAttr }
func (EnhancedColorMode) Readable() bool   { return true }
func (EnhancedColorMode) Writable() bool   { return false }
func (EnhancedColorMode) Reportable() bool { return false }
func (EnhancedColorMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v EnhancedColorMode) AttrID() zcl.AttrID   { return v.ID() }
func (v EnhancedColorMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *EnhancedColorMode) AttrValue() zcl.Val  { return v.Value() }

func (EnhancedColorMode) Name() string { return `Enhanced color mode` }
func (EnhancedColorMode) Description() string {
	return `specifies which attributes are currently determining the color of
the device`
}

// EnhancedColorMode specifies which attributes are currently determining the color of
// the device
type EnhancedColorMode zcl.Zenum8

func (v *EnhancedColorMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *EnhancedColorMode) Value() zcl.Val     { return v }

func (v EnhancedColorMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *EnhancedColorMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = EnhancedColorMode(*nt)
	return br, err
}

func (v EnhancedColorMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *EnhancedColorMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = EnhancedColorMode(*a)
	return nil
}

func (v *EnhancedColorMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = EnhancedColorMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EnhancedColorMode) String() string {
	switch v {
	case 0x00:
		return "Current hue and current saturation"
	case 0x01:
		return "Current X and Current Y"
	case 0x02:
		return "Color temperature"
	case 0x03:
		return "Enhanced current hue and current saturation"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v EnhancedColorMode) IsCurrentHueAndCurrentSaturation() bool         { return v == 0x00 }
func (v EnhancedColorMode) IsCurrentXAndCurrentY() bool                    { return v == 0x01 }
func (v EnhancedColorMode) IsColorTemperature() bool                       { return v == 0x02 }
func (v EnhancedColorMode) IsEnhancedCurrentHueAndCurrentSaturation() bool { return v == 0x03 }
func (v *EnhancedColorMode) SetCurrentHueAndCurrentSaturation()            { *v = 0x00 }
func (v *EnhancedColorMode) SetCurrentXAndCurrentY()                       { *v = 0x01 }
func (v *EnhancedColorMode) SetColorTemperature()                          { *v = 0x02 }
func (v *EnhancedColorMode) SetEnhancedCurrentHueAndCurrentSaturation()    { *v = 0x03 }

func (EnhancedColorMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Current hue and current saturation"},
		{Value: 0x01, Name: "Current X and Current Y"},
		{Value: 0x02, Name: "Color temperature"},
		{Value: 0x03, Name: "Enhanced current hue and current saturation"},
	}
}

const EnhancedCurrentHueAttr zcl.AttrID = 16384

func (EnhancedCurrentHue) ID() zcl.AttrID   { return EnhancedCurrentHueAttr }
func (EnhancedCurrentHue) Readable() bool   { return true }
func (EnhancedCurrentHue) Writable() bool   { return false }
func (EnhancedCurrentHue) Reportable() bool { return false }
func (EnhancedCurrentHue) SceneIndex() int  { return 3 }

// Implements AttrDef/AttrValue interfaces
func (v EnhancedCurrentHue) AttrID() zcl.AttrID   { return v.ID() }
func (v EnhancedCurrentHue) AttrType() zcl.TypeID { return v.TypeID() }
func (v *EnhancedCurrentHue) AttrValue() zcl.Val  { return v.Value() }

func (EnhancedCurrentHue) Name() string { return `Enhanced current hue` }
func (EnhancedCurrentHue) Description() string {
	return `represents non-equidistant steps along the CIE 1931 color triangle,
and it provides 16-bits precision. The upper 8 bits of this attribute
are used as an index in the implementation specific XY lookup table to
provide the non-equidistance steps. The lower 8 bits are used to
interpolate between these steps in a linear way in order to provide color
zoom for the user. To provide compatibility with standard ZCL, the
CurrentHue attribute contains a hue value in the range 0 to 254,
calculated from the EnhancedCurrentHue attribute`
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

func (v *EnhancedCurrentHue) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *EnhancedCurrentHue) Value() zcl.Val     { return v }

func (v EnhancedCurrentHue) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *EnhancedCurrentHue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = EnhancedCurrentHue(*nt)
	return br, err
}

func (v EnhancedCurrentHue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *EnhancedCurrentHue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = EnhancedCurrentHue(*a)
	return nil
}

func (v *EnhancedCurrentHue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = EnhancedCurrentHue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EnhancedCurrentHue) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

func (Hue) Name() string        { return `Hue` }
func (Hue) Description() string { return `` }

type Hue zcl.Zu8

func (v *Hue) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Hue) Value() zcl.Val     { return v }

func (v Hue) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Hue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Hue(*nt)
	return br, err
}

func (v Hue) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Hue) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Hue(*a)
	return nil
}

func (v *Hue) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Hue(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Hue) String() string {
	return zcl.DegreesAngular.Format(float64(v) / 0.70556)
}

func (HueDirection) Name() string        { return `Hue Direction` }
func (HueDirection) Description() string { return `` }

type HueDirection zcl.Zenum8

func (v *HueDirection) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *HueDirection) Value() zcl.Val     { return v }

func (v HueDirection) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *HueDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = HueDirection(*nt)
	return br, err
}

func (v HueDirection) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *HueDirection) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = HueDirection(*a)
	return nil
}

func (v *HueDirection) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = HueDirection(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v HueDirection) String() string {
	switch v {
	case 0x00:
		return "Decrement hue"
	case 0x01:
		return "Increment hue"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v HueDirection) IsDecrementHue() bool { return v == 0x00 }
func (v HueDirection) IsIncrementHue() bool { return v == 0x01 }
func (v *HueDirection) SetDecrementHue()    { *v = 0x00 }
func (v *HueDirection) SetIncrementHue()    { *v = 0x01 }

func (HueDirection) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Decrement hue"},
		{Value: 0x01, Name: "Increment hue"},
	}
}

const IntrinsicBallastFactorAttr zcl.AttrID = 20

func (IntrinsicBallastFactor) ID() zcl.AttrID   { return IntrinsicBallastFactorAttr }
func (IntrinsicBallastFactor) Readable() bool   { return true }
func (IntrinsicBallastFactor) Writable() bool   { return true }
func (IntrinsicBallastFactor) Reportable() bool { return false }
func (IntrinsicBallastFactor) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IntrinsicBallastFactor) AttrID() zcl.AttrID   { return v.ID() }
func (v IntrinsicBallastFactor) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IntrinsicBallastFactor) AttrValue() zcl.Val  { return v.Value() }

func (IntrinsicBallastFactor) Name() string { return `Intrinsic Ballast Factor` }
func (IntrinsicBallastFactor) Description() string {
	return `specifies, as a percentage, the ballast factor of the ballast/lamp
combination, prior to any adjustment.`
}

// IntrinsicBallastFactor specifies, as a percentage, the ballast factor of the ballast/lamp
// combination, prior to any adjustment.
type IntrinsicBallastFactor zcl.Zu8

func (v *IntrinsicBallastFactor) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *IntrinsicBallastFactor) Value() zcl.Val     { return v }

func (v IntrinsicBallastFactor) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *IntrinsicBallastFactor) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = IntrinsicBallastFactor(*nt)
	return br, err
}

func (v IntrinsicBallastFactor) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *IntrinsicBallastFactor) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IntrinsicBallastFactor(*a)
	return nil
}

func (v *IntrinsicBallastFactor) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = IntrinsicBallastFactor(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IntrinsicBallastFactor) String() string {
	return zcl.Percent.Format(float64(v))
}

const LampAlarmModeAttr zcl.AttrID = 52

func (LampAlarmMode) ID() zcl.AttrID   { return LampAlarmModeAttr }
func (LampAlarmMode) Readable() bool   { return true }
func (LampAlarmMode) Writable() bool   { return true }
func (LampAlarmMode) Reportable() bool { return false }
func (LampAlarmMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LampAlarmMode) AttrID() zcl.AttrID   { return v.ID() }
func (v LampAlarmMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LampAlarmMode) AttrValue() zcl.Val  { return v.Value() }

func (LampAlarmMode) Name() string { return `Lamp Alarm Mode` }
func (LampAlarmMode) Description() string {
	return `specifies which attributes can cause an alarm notification to be
generated`
}

// LampAlarmMode specifies which attributes can cause an alarm notification to be
// generated
type LampAlarmMode zcl.Zbmp8

func (v *LampAlarmMode) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *LampAlarmMode) Value() zcl.Val     { return v }

func (v LampAlarmMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *LampAlarmMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = LampAlarmMode(*nt)
	return br, err
}

func (v LampAlarmMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *LampAlarmMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LampAlarmMode(*a)
	return nil
}

func (v *LampAlarmMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = LampAlarmMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LampAlarmMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v LampAlarmMode) IsLampBurnHours() bool    { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v *LampAlarmMode) SetLampBurnHours(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }

func (LampAlarmMode) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Lamp Burn Hours"},
	}
}

const LampBurnHoursAttr zcl.AttrID = 51

func (LampBurnHours) ID() zcl.AttrID   { return LampBurnHoursAttr }
func (LampBurnHours) Readable() bool   { return true }
func (LampBurnHours) Writable() bool   { return true }
func (LampBurnHours) Reportable() bool { return false }
func (LampBurnHours) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LampBurnHours) AttrID() zcl.AttrID   { return v.ID() }
func (v LampBurnHours) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LampBurnHours) AttrValue() zcl.Val  { return v.Value() }

func (LampBurnHours) Name() string { return `Lamp Burn Hours` }
func (LampBurnHours) Description() string {
	return `specifies the length of time, in hours, the currently connected
lamps have been operated, cumulative since the last re-lamping. Burn
hours are not accumulated if the lamps are off and are reset to 0
when a lamp is changed`
}

// LampBurnHours specifies the length of time, in hours, the currently connected
// lamps have been operated, cumulative since the last re-lamping. Burn
// hours are not accumulated if the lamps are off and are reset to 0
// when a lamp is changed
type LampBurnHours zcl.Zu24

func (v *LampBurnHours) TypeID() zcl.TypeID { return new(zcl.Zu24).TypeID() }
func (v *LampBurnHours) Value() zcl.Val     { return v }

func (v LampBurnHours) MarshalZcl() ([]byte, error) { return zcl.Zu24(v).MarshalZcl() }

func (v *LampBurnHours) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*v = LampBurnHours(*nt)
	return br, err
}

func (v LampBurnHours) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu24(v))
}

func (v *LampBurnHours) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LampBurnHours(*a)
	return nil
}

func (v *LampBurnHours) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu24); ok {
		*v = LampBurnHours(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LampBurnHours) String() string {
	return zcl.Hours.Format(float64(v))
}

const LampBurnHoursTripPointAttr zcl.AttrID = 53

func (LampBurnHoursTripPoint) ID() zcl.AttrID   { return LampBurnHoursTripPointAttr }
func (LampBurnHoursTripPoint) Readable() bool   { return true }
func (LampBurnHoursTripPoint) Writable() bool   { return true }
func (LampBurnHoursTripPoint) Reportable() bool { return false }
func (LampBurnHoursTripPoint) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LampBurnHoursTripPoint) AttrID() zcl.AttrID   { return v.ID() }
func (v LampBurnHoursTripPoint) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LampBurnHoursTripPoint) AttrValue() zcl.Val  { return v.Value() }

func (LampBurnHoursTripPoint) Name() string { return `Lamp Burn Hours Trip Point` }
func (LampBurnHoursTripPoint) Description() string {
	return `specifies the number of hours the LampBurnHours attribute can
reach before an alarm is generated`
}

// LampBurnHoursTripPoint specifies the number of hours the LampBurnHours attribute can
// reach before an alarm is generated
type LampBurnHoursTripPoint zcl.Zu24

func (v *LampBurnHoursTripPoint) TypeID() zcl.TypeID { return new(zcl.Zu24).TypeID() }
func (v *LampBurnHoursTripPoint) Value() zcl.Val     { return v }

func (v LampBurnHoursTripPoint) MarshalZcl() ([]byte, error) { return zcl.Zu24(v).MarshalZcl() }

func (v *LampBurnHoursTripPoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*v = LampBurnHoursTripPoint(*nt)
	return br, err
}

func (v LampBurnHoursTripPoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu24(v))
}

func (v *LampBurnHoursTripPoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LampBurnHoursTripPoint(*a)
	return nil
}

func (v *LampBurnHoursTripPoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu24); ok {
		*v = LampBurnHoursTripPoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LampBurnHoursTripPoint) String() string {
	return zcl.Hours.Format(float64(v))
}

const LampManufacturerAttr zcl.AttrID = 49

func (LampManufacturer) ID() zcl.AttrID   { return LampManufacturerAttr }
func (LampManufacturer) Readable() bool   { return true }
func (LampManufacturer) Writable() bool   { return true }
func (LampManufacturer) Reportable() bool { return false }
func (LampManufacturer) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LampManufacturer) AttrID() zcl.AttrID   { return v.ID() }
func (v LampManufacturer) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LampManufacturer) AttrValue() zcl.Val  { return v.Value() }

func (LampManufacturer) Name() string { return `Lamp Manufacturer` }
func (LampManufacturer) Description() string {
	return `specifies the name of the manufacturer of the currently connected
lamps.`
}

// LampManufacturer specifies the name of the manufacturer of the currently connected
// lamps.
type LampManufacturer zcl.Zcstring

func (v *LampManufacturer) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *LampManufacturer) Value() zcl.Val     { return v }

func (v LampManufacturer) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *LampManufacturer) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = LampManufacturer(*nt)
	return br, err
}

func (v LampManufacturer) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *LampManufacturer) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LampManufacturer(*a)
	return nil
}

func (v *LampManufacturer) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = LampManufacturer(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LampManufacturer) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const LampQuantityAttr zcl.AttrID = 32

func (LampQuantity) ID() zcl.AttrID   { return LampQuantityAttr }
func (LampQuantity) Readable() bool   { return true }
func (LampQuantity) Writable() bool   { return false }
func (LampQuantity) Reportable() bool { return false }
func (LampQuantity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LampQuantity) AttrID() zcl.AttrID   { return v.ID() }
func (v LampQuantity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LampQuantity) AttrValue() zcl.Val  { return v.Value() }

func (LampQuantity) Name() string { return `Lamp Quantity` }
func (LampQuantity) Description() string {
	return `specifies the number of lamps connected to this ballast`
}

// LampQuantity specifies the number of lamps connected to this ballast
type LampQuantity zcl.Zu8

func (v *LampQuantity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *LampQuantity) Value() zcl.Val     { return v }

func (v LampQuantity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *LampQuantity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = LampQuantity(*nt)
	return br, err
}

func (v LampQuantity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *LampQuantity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LampQuantity(*a)
	return nil
}

func (v *LampQuantity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = LampQuantity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LampQuantity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const LampRatedHoursAttr zcl.AttrID = 50

func (LampRatedHours) ID() zcl.AttrID   { return LampRatedHoursAttr }
func (LampRatedHours) Readable() bool   { return true }
func (LampRatedHours) Writable() bool   { return true }
func (LampRatedHours) Reportable() bool { return false }
func (LampRatedHours) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LampRatedHours) AttrID() zcl.AttrID   { return v.ID() }
func (v LampRatedHours) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LampRatedHours) AttrValue() zcl.Val  { return v.Value() }

func (LampRatedHours) Name() string { return `Lamp Rated Hours` }
func (LampRatedHours) Description() string {
	return `specifies the number of hours of use the lamps are rated for by
the manufacturer`
}

// LampRatedHours specifies the number of hours of use the lamps are rated for by
// the manufacturer
type LampRatedHours zcl.Zu24

func (v *LampRatedHours) TypeID() zcl.TypeID { return new(zcl.Zu24).TypeID() }
func (v *LampRatedHours) Value() zcl.Val     { return v }

func (v LampRatedHours) MarshalZcl() ([]byte, error) { return zcl.Zu24(v).MarshalZcl() }

func (v *LampRatedHours) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*v = LampRatedHours(*nt)
	return br, err
}

func (v LampRatedHours) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu24(v))
}

func (v *LampRatedHours) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LampRatedHours(*a)
	return nil
}

func (v *LampRatedHours) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu24); ok {
		*v = LampRatedHours(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LampRatedHours) String() string {
	return zcl.Hours.Format(float64(v))
}

const LampTypeAttr zcl.AttrID = 48

func (LampType) ID() zcl.AttrID   { return LampTypeAttr }
func (LampType) Readable() bool   { return true }
func (LampType) Writable() bool   { return true }
func (LampType) Reportable() bool { return false }
func (LampType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LampType) AttrID() zcl.AttrID   { return v.ID() }
func (v LampType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LampType) AttrValue() zcl.Val  { return v.Value() }

func (LampType) Name() string        { return `Lamp Type` }
func (LampType) Description() string { return `` }

type LampType zcl.Zcstring

func (v *LampType) TypeID() zcl.TypeID { return new(zcl.Zcstring).TypeID() }
func (v *LampType) Value() zcl.Val     { return v }

func (v LampType) MarshalZcl() ([]byte, error) { return zcl.Zcstring(v).MarshalZcl() }

func (v *LampType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*v = LampType(*nt)
	return br, err
}

func (v LampType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zcstring(v))
}

func (v *LampType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zcstring)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LampType(*a)
	return nil
}

func (v *LampType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zcstring); ok {
		*v = LampType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LampType) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(v))
}

const MaxLevelAttr zcl.AttrID = 17

func (MaxLevel) ID() zcl.AttrID   { return MaxLevelAttr }
func (MaxLevel) Readable() bool   { return true }
func (MaxLevel) Writable() bool   { return true }
func (MaxLevel) Reportable() bool { return false }
func (MaxLevel) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxLevel) AttrValue() zcl.Val  { return v.Value() }

func (MaxLevel) Name() string { return `Max Level` }
func (MaxLevel) Description() string {
	return `specifies the maximum light level the ballast is permitted to use.
It is specified in the range 0x01 to 0xfe, and specifies the light
output of the ballast according to the dimming light curve. The value
is both less than or equal to PhysicalMaxLevel and greater than or equal
to MinLevel`
}

// MaxLevel specifies the maximum light level the ballast is permitted to use.
// It is specified in the range 0x01 to 0xfe, and specifies the light
// output of the ballast according to the dimming light curve. The value
// is both less than or equal to PhysicalMaxLevel and greater than or equal
// to MinLevel
type MaxLevel zcl.Zu8

func (v *MaxLevel) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *MaxLevel) Value() zcl.Val     { return v }

func (v MaxLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *MaxLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxLevel(*nt)
	return br, err
}

func (v MaxLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *MaxLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxLevel(*a)
	return nil
}

func (v *MaxLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = MaxLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const MinLevelAttr zcl.AttrID = 16

func (MinLevel) ID() zcl.AttrID   { return MinLevelAttr }
func (MinLevel) Readable() bool   { return true }
func (MinLevel) Writable() bool   { return true }
func (MinLevel) Reportable() bool { return false }
func (MinLevel) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v MinLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinLevel) AttrValue() zcl.Val  { return v.Value() }

func (MinLevel) Name() string        { return `Min Level` }
func (MinLevel) Description() string { return `` }

type MinLevel zcl.Zu8

func (v *MinLevel) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *MinLevel) Value() zcl.Val     { return v }

func (v MinLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *MinLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = MinLevel(*nt)
	return br, err
}

func (v MinLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *MinLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinLevel(*a)
	return nil
}

func (v *MinLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = MinLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (MoveDirection) Name() string        { return `Move Direction` }
func (MoveDirection) Description() string { return `` }

type MoveDirection zcl.Zenum8

func (v *MoveDirection) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *MoveDirection) Value() zcl.Val     { return v }

func (v MoveDirection) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *MoveDirection) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = MoveDirection(*nt)
	return br, err
}

func (v MoveDirection) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *MoveDirection) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MoveDirection(*a)
	return nil
}

func (v *MoveDirection) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = MoveDirection(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MoveDirection) String() string {
	switch v {
	case 0x00:
		return "Shortest distance"
	case 0x01:
		return "Longest Distance"
	case 0x02:
		return "Up"
	case 0x03:
		return "Down"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v MoveDirection) IsShortestDistance() bool { return v == 0x00 }
func (v MoveDirection) IsLongestDistance() bool  { return v == 0x01 }
func (v MoveDirection) IsUp() bool               { return v == 0x02 }
func (v MoveDirection) IsDown() bool             { return v == 0x03 }
func (v *MoveDirection) SetShortestDistance()    { *v = 0x00 }
func (v *MoveDirection) SetLongestDistance()     { *v = 0x01 }
func (v *MoveDirection) SetUp()                  { *v = 0x02 }
func (v *MoveDirection) SetDown()                { *v = 0x03 }

func (MoveDirection) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Shortest distance"},
		{Value: 0x01, Name: "Longest Distance"},
		{Value: 0x02, Name: "Up"},
		{Value: 0x03, Name: "Down"},
	}
}

func (MoveMode) Name() string        { return `Move mode` }
func (MoveMode) Description() string { return `` }

type MoveMode zcl.Zenum8

func (v *MoveMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *MoveMode) Value() zcl.Val     { return v }

func (v MoveMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *MoveMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = MoveMode(*nt)
	return br, err
}

func (v MoveMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *MoveMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MoveMode(*a)
	return nil
}

func (v *MoveMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = MoveMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MoveMode) String() string {
	switch v {
	case 0x00:
		return "Stop"
	case 0x01:
		return "Up"
	case 0x03:
		return "Down"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v MoveMode) IsStop() bool { return v == 0x00 }
func (v MoveMode) IsUp() bool   { return v == 0x01 }
func (v MoveMode) IsDown() bool { return v == 0x03 }
func (v *MoveMode) SetStop()    { *v = 0x00 }
func (v *MoveMode) SetUp()      { *v = 0x01 }
func (v *MoveMode) SetDown()    { *v = 0x03 }

func (MoveMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Stop"},
		{Value: 0x01, Name: "Up"},
		{Value: 0x03, Name: "Down"},
	}
}

const NumberOfPrimariesAttr zcl.AttrID = 16

func (NumberOfPrimaries) ID() zcl.AttrID   { return NumberOfPrimariesAttr }
func (NumberOfPrimaries) Readable() bool   { return true }
func (NumberOfPrimaries) Writable() bool   { return false }
func (NumberOfPrimaries) Reportable() bool { return false }
func (NumberOfPrimaries) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NumberOfPrimaries) AttrID() zcl.AttrID   { return v.ID() }
func (v NumberOfPrimaries) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NumberOfPrimaries) AttrValue() zcl.Val  { return v.Value() }

func (NumberOfPrimaries) Name() string { return `Number of primaries` }
func (NumberOfPrimaries) Description() string {
	return `contains the number of color primaries implemented on this device.
A value of 0xff indicates that the number of primaries is unknown`
}

// NumberOfPrimaries contains the number of color primaries implemented on this device.
// A value of 0xff indicates that the number of primaries is unknown
type NumberOfPrimaries zcl.Zu8

func (v *NumberOfPrimaries) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *NumberOfPrimaries) Value() zcl.Val     { return v }

func (v NumberOfPrimaries) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *NumberOfPrimaries) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = NumberOfPrimaries(*nt)
	return br, err
}

func (v NumberOfPrimaries) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *NumberOfPrimaries) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NumberOfPrimaries(*a)
	return nil
}

func (v *NumberOfPrimaries) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = NumberOfPrimaries(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NumberOfPrimaries) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const PhysicalMaxLevelAttr zcl.AttrID = 1

func (PhysicalMaxLevel) ID() zcl.AttrID   { return PhysicalMaxLevelAttr }
func (PhysicalMaxLevel) Readable() bool   { return true }
func (PhysicalMaxLevel) Writable() bool   { return false }
func (PhysicalMaxLevel) Reportable() bool { return false }
func (PhysicalMaxLevel) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PhysicalMaxLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v PhysicalMaxLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PhysicalMaxLevel) AttrValue() zcl.Val  { return v.Value() }

func (PhysicalMaxLevel) Name() string { return `Physical Max Level` }
func (PhysicalMaxLevel) Description() string {
	return `specifies the maximum light level the ballast can achieve. It is
specified in the range 0x01 to 0xfe, and specifies the light output
of the ballast according to the dimming light curve`
}

// PhysicalMaxLevel specifies the maximum light level the ballast can achieve. It is
// specified in the range 0x01 to 0xfe, and specifies the light output
// of the ballast according to the dimming light curve
type PhysicalMaxLevel zcl.Zu8

func (v *PhysicalMaxLevel) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *PhysicalMaxLevel) Value() zcl.Val     { return v }

func (v PhysicalMaxLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *PhysicalMaxLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = PhysicalMaxLevel(*nt)
	return br, err
}

func (v PhysicalMaxLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *PhysicalMaxLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PhysicalMaxLevel(*a)
	return nil
}

func (v *PhysicalMaxLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = PhysicalMaxLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PhysicalMaxLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const PhysicalMinLevelAttr zcl.AttrID = 0

func (PhysicalMinLevel) ID() zcl.AttrID   { return PhysicalMinLevelAttr }
func (PhysicalMinLevel) Readable() bool   { return true }
func (PhysicalMinLevel) Writable() bool   { return false }
func (PhysicalMinLevel) Reportable() bool { return false }
func (PhysicalMinLevel) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PhysicalMinLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v PhysicalMinLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PhysicalMinLevel) AttrValue() zcl.Val  { return v.Value() }

func (PhysicalMinLevel) Name() string { return `Physical Min Level` }
func (PhysicalMinLevel) Description() string {
	return `specifies the minimum light level the ballast can achieve. It is
specified in the range 0x01 to 0xfe, and specifies the light output
of the ballast according to the dimming light curve`
}

// PhysicalMinLevel specifies the minimum light level the ballast can achieve. It is
// specified in the range 0x01 to 0xfe, and specifies the light output
// of the ballast according to the dimming light curve
type PhysicalMinLevel zcl.Zu8

func (v *PhysicalMinLevel) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *PhysicalMinLevel) Value() zcl.Val     { return v }

func (v PhysicalMinLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *PhysicalMinLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = PhysicalMinLevel(*nt)
	return br, err
}

func (v PhysicalMinLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *PhysicalMinLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PhysicalMinLevel(*a)
	return nil
}

func (v *PhysicalMinLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = PhysicalMinLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PhysicalMinLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const PowerOnLevelAttr zcl.AttrID = 18

func (PowerOnLevel) ID() zcl.AttrID   { return PowerOnLevelAttr }
func (PowerOnLevel) Readable() bool   { return true }
func (PowerOnLevel) Writable() bool   { return true }
func (PowerOnLevel) Reportable() bool { return false }
func (PowerOnLevel) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerOnLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerOnLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerOnLevel) AttrValue() zcl.Val  { return v.Value() }

func (PowerOnLevel) Name() string { return `Power On level` }
func (PowerOnLevel) Description() string {
	return `It specifies the light level to which the ballast will go when power
is applied (e.g., when mains power is re-established after a power
failure). It can have a value between 0x00 and 0xfe to set it to
a specific light level, according to the dimming light curve, or
0xff to set it to the value it was before the power failure`
}

// PowerOnLevel It specifies the light level to which the ballast will go when power
// is applied (e.g., when mains power is re-established after a power
// failure). It can have a value between 0x00 and 0xfe to set it to
// a specific light level, according to the dimming light curve, or
// 0xff to set it to the value it was before the power failure
type PowerOnLevel zcl.Zu8

func (v *PowerOnLevel) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *PowerOnLevel) Value() zcl.Val     { return v }

func (v PowerOnLevel) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *PowerOnLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerOnLevel(*nt)
	return br, err
}

func (v PowerOnLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *PowerOnLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerOnLevel(*a)
	return nil
}

func (v *PowerOnLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = PowerOnLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerOnLevel) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const PowerOnTimeAttr zcl.AttrID = 19

func (PowerOnTime) ID() zcl.AttrID   { return PowerOnTimeAttr }
func (PowerOnTime) Readable() bool   { return true }
func (PowerOnTime) Writable() bool   { return true }
func (PowerOnTime) Reportable() bool { return false }
func (PowerOnTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerOnTime) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerOnTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerOnTime) AttrValue() zcl.Val  { return v.Value() }

func (PowerOnTime) Name() string        { return `Power-On Time` }
func (PowerOnTime) Description() string { return `The transition time in 1/10ths of a second.` }

// PowerOnTime The transition time in 1/10ths of a second.
type PowerOnTime zcl.Zu16

func (v *PowerOnTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PowerOnTime) Value() zcl.Val     { return v }

func (v PowerOnTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PowerOnTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerOnTime(*nt)
	return br, err
}

func (v PowerOnTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PowerOnTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerOnTime(*a)
	return nil
}

func (v *PowerOnTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PowerOnTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerOnTime) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

const PowerOnColorTemperatureAttr zcl.AttrID = 16400

func (PowerOnColorTemperature) ID() zcl.AttrID   { return PowerOnColorTemperatureAttr }
func (PowerOnColorTemperature) Readable() bool   { return true }
func (PowerOnColorTemperature) Writable() bool   { return true }
func (PowerOnColorTemperature) Reportable() bool { return false }
func (PowerOnColorTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PowerOnColorTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v PowerOnColorTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PowerOnColorTemperature) AttrValue() zcl.Val  { return v.Value() }

func (PowerOnColorTemperature) Name() string        { return `Power-on color temperature` }
func (PowerOnColorTemperature) Description() string { return `` }

type PowerOnColorTemperature zcl.Zu16

func (v *PowerOnColorTemperature) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PowerOnColorTemperature) Value() zcl.Val     { return v }

func (v PowerOnColorTemperature) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PowerOnColorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PowerOnColorTemperature(*nt)
	return br, err
}

func (v PowerOnColorTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PowerOnColorTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PowerOnColorTemperature(*a)
	return nil
}

func (v *PowerOnColorTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PowerOnColorTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PowerOnColorTemperature) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary1XAttr zcl.AttrID = 17

func (Primary1X) ID() zcl.AttrID   { return Primary1XAttr }
func (Primary1X) Readable() bool   { return true }
func (Primary1X) Writable() bool   { return false }
func (Primary1X) Reportable() bool { return false }
func (Primary1X) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary1X) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary1X) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary1X) AttrValue() zcl.Val  { return v.Value() }

func (Primary1X) Name() string { return `Primary1 X` }
func (Primary1X) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary1X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary1X zcl.Zu16

func (v *Primary1X) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary1X) Value() zcl.Val     { return v }

func (v Primary1X) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary1X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary1X(*nt)
	return br, err
}

func (v Primary1X) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary1X) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary1X(*a)
	return nil
}

func (v *Primary1X) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary1X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary1X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary1YAttr zcl.AttrID = 18

func (Primary1Y) ID() zcl.AttrID   { return Primary1YAttr }
func (Primary1Y) Readable() bool   { return true }
func (Primary1Y) Writable() bool   { return false }
func (Primary1Y) Reportable() bool { return false }
func (Primary1Y) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary1Y) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary1Y) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary1Y) AttrValue() zcl.Val  { return v.Value() }

func (Primary1Y) Name() string { return `Primary1 Y` }
func (Primary1Y) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary1Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary1Y zcl.Zu16

func (v *Primary1Y) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary1Y) Value() zcl.Val     { return v }

func (v Primary1Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary1Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary1Y(*nt)
	return br, err
}

func (v Primary1Y) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary1Y) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary1Y(*a)
	return nil
}

func (v *Primary1Y) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary1Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary1Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary1IntensityAttr zcl.AttrID = 19

func (Primary1Intensity) ID() zcl.AttrID   { return Primary1IntensityAttr }
func (Primary1Intensity) Readable() bool   { return true }
func (Primary1Intensity) Writable() bool   { return false }
func (Primary1Intensity) Reportable() bool { return false }
func (Primary1Intensity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary1Intensity) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary1Intensity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary1Intensity) AttrValue() zcl.Val  { return v.Value() }

func (Primary1Intensity) Name() string { return `Primary1 intensity` }
func (Primary1Intensity) Description() string {
	return `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`
}

// Primary1Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary1Intensity zcl.Zu8

func (v *Primary1Intensity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Primary1Intensity) Value() zcl.Val     { return v }

func (v Primary1Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Primary1Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary1Intensity(*nt)
	return br, err
}

func (v Primary1Intensity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Primary1Intensity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary1Intensity(*a)
	return nil
}

func (v *Primary1Intensity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Primary1Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary1Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const Primary2XAttr zcl.AttrID = 21

func (Primary2X) ID() zcl.AttrID   { return Primary2XAttr }
func (Primary2X) Readable() bool   { return true }
func (Primary2X) Writable() bool   { return false }
func (Primary2X) Reportable() bool { return false }
func (Primary2X) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary2X) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary2X) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary2X) AttrValue() zcl.Val  { return v.Value() }

func (Primary2X) Name() string { return `Primary2 X` }
func (Primary2X) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary2X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary2X zcl.Zu16

func (v *Primary2X) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary2X) Value() zcl.Val     { return v }

func (v Primary2X) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary2X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary2X(*nt)
	return br, err
}

func (v Primary2X) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary2X) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary2X(*a)
	return nil
}

func (v *Primary2X) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary2X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary2X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary2YAttr zcl.AttrID = 22

func (Primary2Y) ID() zcl.AttrID   { return Primary2YAttr }
func (Primary2Y) Readable() bool   { return true }
func (Primary2Y) Writable() bool   { return false }
func (Primary2Y) Reportable() bool { return false }
func (Primary2Y) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary2Y) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary2Y) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary2Y) AttrValue() zcl.Val  { return v.Value() }

func (Primary2Y) Name() string { return `Primary2 Y` }
func (Primary2Y) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary2Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary2Y zcl.Zu16

func (v *Primary2Y) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary2Y) Value() zcl.Val     { return v }

func (v Primary2Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary2Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary2Y(*nt)
	return br, err
}

func (v Primary2Y) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary2Y) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary2Y(*a)
	return nil
}

func (v *Primary2Y) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary2Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary2Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary2IntensityAttr zcl.AttrID = 23

func (Primary2Intensity) ID() zcl.AttrID   { return Primary2IntensityAttr }
func (Primary2Intensity) Readable() bool   { return true }
func (Primary2Intensity) Writable() bool   { return false }
func (Primary2Intensity) Reportable() bool { return false }
func (Primary2Intensity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary2Intensity) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary2Intensity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary2Intensity) AttrValue() zcl.Val  { return v.Value() }

func (Primary2Intensity) Name() string { return `Primary2 intensity` }
func (Primary2Intensity) Description() string {
	return `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`
}

// Primary2Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary2Intensity zcl.Zu8

func (v *Primary2Intensity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Primary2Intensity) Value() zcl.Val     { return v }

func (v Primary2Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Primary2Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary2Intensity(*nt)
	return br, err
}

func (v Primary2Intensity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Primary2Intensity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary2Intensity(*a)
	return nil
}

func (v *Primary2Intensity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Primary2Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary2Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const Primary3XAttr zcl.AttrID = 25

func (Primary3X) ID() zcl.AttrID   { return Primary3XAttr }
func (Primary3X) Readable() bool   { return true }
func (Primary3X) Writable() bool   { return false }
func (Primary3X) Reportable() bool { return false }
func (Primary3X) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary3X) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary3X) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary3X) AttrValue() zcl.Val  { return v.Value() }

func (Primary3X) Name() string { return `Primary3 X` }
func (Primary3X) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary3X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary3X zcl.Zu16

func (v *Primary3X) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary3X) Value() zcl.Val     { return v }

func (v Primary3X) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary3X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary3X(*nt)
	return br, err
}

func (v Primary3X) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary3X) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary3X(*a)
	return nil
}

func (v *Primary3X) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary3X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary3X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary3YAttr zcl.AttrID = 26

func (Primary3Y) ID() zcl.AttrID   { return Primary3YAttr }
func (Primary3Y) Readable() bool   { return true }
func (Primary3Y) Writable() bool   { return false }
func (Primary3Y) Reportable() bool { return false }
func (Primary3Y) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary3Y) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary3Y) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary3Y) AttrValue() zcl.Val  { return v.Value() }

func (Primary3Y) Name() string { return `Primary3 Y` }
func (Primary3Y) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary3Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary3Y zcl.Zu16

func (v *Primary3Y) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary3Y) Value() zcl.Val     { return v }

func (v Primary3Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary3Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary3Y(*nt)
	return br, err
}

func (v Primary3Y) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary3Y) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary3Y(*a)
	return nil
}

func (v *Primary3Y) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary3Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary3Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary3IntensityAttr zcl.AttrID = 27

func (Primary3Intensity) ID() zcl.AttrID   { return Primary3IntensityAttr }
func (Primary3Intensity) Readable() bool   { return true }
func (Primary3Intensity) Writable() bool   { return false }
func (Primary3Intensity) Reportable() bool { return false }
func (Primary3Intensity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary3Intensity) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary3Intensity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary3Intensity) AttrValue() zcl.Val  { return v.Value() }

func (Primary3Intensity) Name() string { return `Primary3 intensity` }
func (Primary3Intensity) Description() string {
	return `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`
}

// Primary3Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary3Intensity zcl.Zu8

func (v *Primary3Intensity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Primary3Intensity) Value() zcl.Val     { return v }

func (v Primary3Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Primary3Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary3Intensity(*nt)
	return br, err
}

func (v Primary3Intensity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Primary3Intensity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary3Intensity(*a)
	return nil
}

func (v *Primary3Intensity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Primary3Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary3Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const Primary4XAttr zcl.AttrID = 32

func (Primary4X) ID() zcl.AttrID   { return Primary4XAttr }
func (Primary4X) Readable() bool   { return true }
func (Primary4X) Writable() bool   { return false }
func (Primary4X) Reportable() bool { return false }
func (Primary4X) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary4X) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary4X) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary4X) AttrValue() zcl.Val  { return v.Value() }

func (Primary4X) Name() string { return `Primary4 X` }
func (Primary4X) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary4X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary4X zcl.Zu16

func (v *Primary4X) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary4X) Value() zcl.Val     { return v }

func (v Primary4X) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary4X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary4X(*nt)
	return br, err
}

func (v Primary4X) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary4X) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary4X(*a)
	return nil
}

func (v *Primary4X) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary4X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary4X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary4YAttr zcl.AttrID = 33

func (Primary4Y) ID() zcl.AttrID   { return Primary4YAttr }
func (Primary4Y) Readable() bool   { return true }
func (Primary4Y) Writable() bool   { return false }
func (Primary4Y) Reportable() bool { return false }
func (Primary4Y) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary4Y) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary4Y) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary4Y) AttrValue() zcl.Val  { return v.Value() }

func (Primary4Y) Name() string { return `Primary4 Y` }
func (Primary4Y) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary4Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary4Y zcl.Zu16

func (v *Primary4Y) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary4Y) Value() zcl.Val     { return v }

func (v Primary4Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary4Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary4Y(*nt)
	return br, err
}

func (v Primary4Y) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary4Y) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary4Y(*a)
	return nil
}

func (v *Primary4Y) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary4Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary4Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary4IntensityAttr zcl.AttrID = 34

func (Primary4Intensity) ID() zcl.AttrID   { return Primary4IntensityAttr }
func (Primary4Intensity) Readable() bool   { return true }
func (Primary4Intensity) Writable() bool   { return false }
func (Primary4Intensity) Reportable() bool { return false }
func (Primary4Intensity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary4Intensity) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary4Intensity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary4Intensity) AttrValue() zcl.Val  { return v.Value() }

func (Primary4Intensity) Name() string { return `Primary4 intensity` }
func (Primary4Intensity) Description() string {
	return `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`
}

// Primary4Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary4Intensity zcl.Zu8

func (v *Primary4Intensity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Primary4Intensity) Value() zcl.Val     { return v }

func (v Primary4Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Primary4Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary4Intensity(*nt)
	return br, err
}

func (v Primary4Intensity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Primary4Intensity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary4Intensity(*a)
	return nil
}

func (v *Primary4Intensity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Primary4Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary4Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const Primary5XAttr zcl.AttrID = 36

func (Primary5X) ID() zcl.AttrID   { return Primary5XAttr }
func (Primary5X) Readable() bool   { return true }
func (Primary5X) Writable() bool   { return false }
func (Primary5X) Reportable() bool { return false }
func (Primary5X) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary5X) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary5X) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary5X) AttrValue() zcl.Val  { return v.Value() }

func (Primary5X) Name() string { return `Primary5 X` }
func (Primary5X) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary5X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary5X zcl.Zu16

func (v *Primary5X) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary5X) Value() zcl.Val     { return v }

func (v Primary5X) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary5X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary5X(*nt)
	return br, err
}

func (v Primary5X) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary5X) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary5X(*a)
	return nil
}

func (v *Primary5X) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary5X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary5X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary5YAttr zcl.AttrID = 37

func (Primary5Y) ID() zcl.AttrID   { return Primary5YAttr }
func (Primary5Y) Readable() bool   { return true }
func (Primary5Y) Writable() bool   { return false }
func (Primary5Y) Reportable() bool { return false }
func (Primary5Y) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary5Y) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary5Y) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary5Y) AttrValue() zcl.Val  { return v.Value() }

func (Primary5Y) Name() string { return `Primary5 Y` }
func (Primary5Y) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary5Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary5Y zcl.Zu16

func (v *Primary5Y) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary5Y) Value() zcl.Val     { return v }

func (v Primary5Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary5Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary5Y(*nt)
	return br, err
}

func (v Primary5Y) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary5Y) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary5Y(*a)
	return nil
}

func (v *Primary5Y) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary5Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary5Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary5IntensityAttr zcl.AttrID = 38

func (Primary5Intensity) ID() zcl.AttrID   { return Primary5IntensityAttr }
func (Primary5Intensity) Readable() bool   { return true }
func (Primary5Intensity) Writable() bool   { return false }
func (Primary5Intensity) Reportable() bool { return false }
func (Primary5Intensity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary5Intensity) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary5Intensity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary5Intensity) AttrValue() zcl.Val  { return v.Value() }

func (Primary5Intensity) Name() string { return `Primary5 intensity` }
func (Primary5Intensity) Description() string {
	return `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`
}

// Primary5Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary5Intensity zcl.Zu8

func (v *Primary5Intensity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Primary5Intensity) Value() zcl.Val     { return v }

func (v Primary5Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Primary5Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary5Intensity(*nt)
	return br, err
}

func (v Primary5Intensity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Primary5Intensity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary5Intensity(*a)
	return nil
}

func (v *Primary5Intensity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Primary5Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary5Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const Primary6XAttr zcl.AttrID = 40

func (Primary6X) ID() zcl.AttrID   { return Primary6XAttr }
func (Primary6X) Readable() bool   { return true }
func (Primary6X) Writable() bool   { return false }
func (Primary6X) Reportable() bool { return false }
func (Primary6X) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary6X) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary6X) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary6X) AttrValue() zcl.Val  { return v.Value() }

func (Primary6X) Name() string { return `Primary6 X` }
func (Primary6X) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary6X contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary6X zcl.Zu16

func (v *Primary6X) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary6X) Value() zcl.Val     { return v }

func (v Primary6X) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary6X) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary6X(*nt)
	return br, err
}

func (v Primary6X) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary6X) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary6X(*a)
	return nil
}

func (v *Primary6X) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary6X(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary6X) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary6YAttr zcl.AttrID = 41

func (Primary6Y) ID() zcl.AttrID   { return Primary6YAttr }
func (Primary6Y) Readable() bool   { return true }
func (Primary6Y) Writable() bool   { return false }
func (Primary6Y) Reportable() bool { return false }
func (Primary6Y) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary6Y) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary6Y) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary6Y) AttrValue() zcl.Val  { return v.Value() }

func (Primary6Y) Name() string { return `Primary6 Y` }
func (Primary6Y) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// Primary6Y contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type Primary6Y zcl.Zu16

func (v *Primary6Y) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Primary6Y) Value() zcl.Val     { return v }

func (v Primary6Y) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Primary6Y) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary6Y(*nt)
	return br, err
}

func (v Primary6Y) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Primary6Y) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary6Y(*a)
	return nil
}

func (v *Primary6Y) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Primary6Y(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary6Y) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const Primary6IntensityAttr zcl.AttrID = 42

func (Primary6Intensity) ID() zcl.AttrID   { return Primary6IntensityAttr }
func (Primary6Intensity) Readable() bool   { return true }
func (Primary6Intensity) Writable() bool   { return false }
func (Primary6Intensity) Reportable() bool { return false }
func (Primary6Intensity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Primary6Intensity) AttrID() zcl.AttrID   { return v.ID() }
func (v Primary6Intensity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Primary6Intensity) AttrValue() zcl.Val  { return v.Value() }

func (Primary6Intensity) Name() string { return `Primary6 intensity` }
func (Primary6Intensity) Description() string {
	return `contains a representation of the maximum intensity of this attribute as
defined in the Dimming Light Curve in the Ballast Configuration cluster,
normalized such that the attribute with the highest maximum intensity
contains the value 0xfe. A value of 0xff indicates that this attribute is
not available`
}

// Primary6Intensity contains a representation of the maximum intensity of this attribute as
// defined in the Dimming Light Curve in the Ballast Configuration cluster,
// normalized such that the attribute with the highest maximum intensity
// contains the value 0xfe. A value of 0xff indicates that this attribute is
// not available
type Primary6Intensity zcl.Zu8

func (v *Primary6Intensity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Primary6Intensity) Value() zcl.Val     { return v }

func (v Primary6Intensity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Primary6Intensity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Primary6Intensity(*nt)
	return br, err
}

func (v Primary6Intensity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Primary6Intensity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Primary6Intensity(*a)
	return nil
}

func (v *Primary6Intensity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Primary6Intensity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Primary6Intensity) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (Rate) Name() string        { return `Rate` }
func (Rate) Description() string { return `increment/steps per second` }

// Rate increment/steps per second
type Rate zcl.Zu8

func (v *Rate) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Rate) Value() zcl.Val     { return v }

func (v Rate) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Rate) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Rate(*nt)
	return br, err
}

func (v Rate) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Rate) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Rate(*a)
	return nil
}

func (v *Rate) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Rate(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Rate) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (RateX) Name() string        { return `Rate X` }
func (RateX) Description() string { return `increment/steps per second` }

// RateX increment/steps per second
type RateX zcl.Zs16

func (v *RateX) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *RateX) Value() zcl.Val     { return v }

func (v RateX) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *RateX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = RateX(*nt)
	return br, err
}

func (v RateX) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *RateX) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RateX(*a)
	return nil
}

func (v *RateX) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = RateX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RateX) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

func (RateY) Name() string        { return `Rate Y` }
func (RateY) Description() string { return `increment/steps per second` }

// RateY increment/steps per second
type RateY zcl.Zs16

func (v *RateY) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *RateY) Value() zcl.Val     { return v }

func (v RateY) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *RateY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = RateY(*nt)
	return br, err
}

func (v RateY) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *RateY) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RateY(*a)
	return nil
}

func (v *RateY) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = RateY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RateY) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

const RemainingTimeAttr zcl.AttrID = 2

func (RemainingTime) ID() zcl.AttrID   { return RemainingTimeAttr }
func (RemainingTime) Readable() bool   { return true }
func (RemainingTime) Writable() bool   { return false }
func (RemainingTime) Reportable() bool { return false }
func (RemainingTime) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RemainingTime) AttrID() zcl.AttrID   { return v.ID() }
func (v RemainingTime) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RemainingTime) AttrValue() zcl.Val  { return v.Value() }

func (RemainingTime) Name() string { return `Remaining time` }
func (RemainingTime) Description() string {
	return `holds the time remaining, in 1/10ths of a second, until the currently
active command will be complete`
}

// RemainingTime holds the time remaining, in 1/10ths of a second, until the currently
// active command will be complete
type RemainingTime zcl.Zu16

func (v *RemainingTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RemainingTime) Value() zcl.Val     { return v }

func (v RemainingTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RemainingTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RemainingTime(*nt)
	return br, err
}

func (v RemainingTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RemainingTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RemainingTime(*a)
	return nil
}

func (v *RemainingTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RemainingTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RemainingTime) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

func (Saturation) Name() string        { return `Saturation` }
func (Saturation) Description() string { return `` }

type Saturation zcl.Zu8

func (v *Saturation) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *Saturation) Value() zcl.Val     { return v }

func (v Saturation) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *Saturation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = Saturation(*nt)
	return br, err
}

func (v Saturation) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *Saturation) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Saturation(*a)
	return nil
}

func (v *Saturation) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = Saturation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Saturation) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (StepX) Name() string        { return `Step X` }
func (StepX) Description() string { return `` }

type StepX zcl.Zs16

func (v *StepX) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *StepX) Value() zcl.Val     { return v }

func (v StepX) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *StepX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = StepX(*nt)
	return br, err
}

func (v StepX) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *StepX) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StepX(*a)
	return nil
}

func (v *StepX) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = StepX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StepX) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

func (StepY) Name() string        { return `Step Y` }
func (StepY) Description() string { return `` }

type StepY zcl.Zs16

func (v *StepY) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *StepY) Value() zcl.Val     { return v }

func (v StepY) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *StepY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = StepY(*nt)
	return br, err
}

func (v StepY) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *StepY) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StepY(*a)
	return nil
}

func (v *StepY) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = StepY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StepY) String() string {
	return zcl.Sprintf("%v", zcl.Zs16(v))
}

func (StepMode) Name() string        { return `Step mode` }
func (StepMode) Description() string { return `` }

type StepMode zcl.Zenum8

func (v *StepMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *StepMode) Value() zcl.Val     { return v }

func (v StepMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *StepMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = StepMode(*nt)
	return br, err
}

func (v StepMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *StepMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StepMode(*a)
	return nil
}

func (v *StepMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = StepMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StepMode) String() string {
	switch v {
	case 0x01:
		return "Up"
	case 0x03:
		return "Down"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v StepMode) IsUp() bool   { return v == 0x01 }
func (v StepMode) IsDown() bool { return v == 0x03 }
func (v *StepMode) SetUp()      { *v = 0x01 }
func (v *StepMode) SetDown()    { *v = 0x03 }

func (StepMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x01, Name: "Up"},
		{Value: 0x03, Name: "Down"},
	}
}

func (StepSize) Name() string        { return `Step size` }
func (StepSize) Description() string { return `` }

type StepSize zcl.Zu8

func (v *StepSize) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *StepSize) Value() zcl.Val     { return v }

func (v StepSize) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *StepSize) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = StepSize(*nt)
	return br, err
}

func (v StepSize) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *StepSize) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StepSize(*a)
	return nil
}

func (v *StepSize) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = StepSize(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StepSize) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (Time) Name() string        { return `Time` }
func (Time) Description() string { return `Time in seconds used for a whole color loop.` }

// Time Time in seconds used for a whole color loop.
type Time zcl.Zu16

func (v *Time) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Time) Value() zcl.Val     { return v }

func (v Time) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Time) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Time(*nt)
	return br, err
}

func (v Time) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Time) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Time(*a)
	return nil
}

func (v *Time) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Time(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Time) String() string {
	return zcl.Seconds.Format(float64(v))
}

func (TransitionTime) Name() string        { return `Transition time` }
func (TransitionTime) Description() string { return `The transition time in 1/10ths of a second.` }

// TransitionTime The transition time in 1/10ths of a second.
type TransitionTime zcl.Zu16

func (v *TransitionTime) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TransitionTime) Value() zcl.Val     { return v }

func (v TransitionTime) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *TransitionTime) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TransitionTime(*nt)
	return br, err
}

func (v TransitionTime) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TransitionTime) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TransitionTime(*a)
	return nil
}

func (v *TransitionTime) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TransitionTime(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TransitionTime) String() string {
	return zcl.Seconds.Format(float64(v) / 10)
}

func (UpdateFlags) Name() string        { return `Update flags` }
func (UpdateFlags) Description() string { return `` }

type UpdateFlags zcl.Zbmp8

func (v *UpdateFlags) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *UpdateFlags) Value() zcl.Val     { return v }

func (v UpdateFlags) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *UpdateFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = UpdateFlags(*nt)
	return br, err
}

func (v UpdateFlags) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *UpdateFlags) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UpdateFlags(*a)
	return nil
}

func (v *UpdateFlags) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = UpdateFlags(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UpdateFlags) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
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

func (v UpdateFlags) IsUpdateAction() bool       { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v UpdateFlags) IsUpdateDirection() bool    { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v UpdateFlags) IsUpdateTime() bool         { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v UpdateFlags) IsUpdateStartHue() bool     { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v *UpdateFlags) SetUpdateAction(b bool)    { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *UpdateFlags) SetUpdateDirection(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }
func (v *UpdateFlags) SetUpdateTime(b bool)      { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b)) }
func (v *UpdateFlags) SetUpdateStartHue(b bool)  { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b)) }

func (UpdateFlags) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Update action"},
		{Value: 1, Name: "Update direction"},
		{Value: 2, Name: "Update time"},
		{Value: 3, Name: "Update start hue"},
	}
}

const WhitePointXAttr zcl.AttrID = 48

func (WhitePointX) ID() zcl.AttrID   { return WhitePointXAttr }
func (WhitePointX) Readable() bool   { return true }
func (WhitePointX) Writable() bool   { return true }
func (WhitePointX) Reportable() bool { return false }
func (WhitePointX) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v WhitePointX) AttrID() zcl.AttrID   { return v.ID() }
func (v WhitePointX) AttrType() zcl.TypeID { return v.TypeID() }
func (v *WhitePointX) AttrValue() zcl.Val  { return v.Value() }

func (WhitePointX) Name() string { return `White Point X` }
func (WhitePointX) Description() string {
	return `contains the normalized chromaticity value x for this attribute, as
defined in the CIE xyY Color Space. x = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// WhitePointX contains the normalized chromaticity value x for this attribute, as
// defined in the CIE xyY Color Space. x = value / 65536 (value
// in the range 0 to 65279 inclusive)
type WhitePointX zcl.Zu16

func (v *WhitePointX) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *WhitePointX) Value() zcl.Val     { return v }

func (v WhitePointX) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *WhitePointX) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = WhitePointX(*nt)
	return br, err
}

func (v WhitePointX) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *WhitePointX) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = WhitePointX(*a)
	return nil
}

func (v *WhitePointX) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = WhitePointX(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v WhitePointX) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const WhitePointYAttr zcl.AttrID = 49

func (WhitePointY) ID() zcl.AttrID   { return WhitePointYAttr }
func (WhitePointY) Readable() bool   { return true }
func (WhitePointY) Writable() bool   { return true }
func (WhitePointY) Reportable() bool { return false }
func (WhitePointY) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v WhitePointY) AttrID() zcl.AttrID   { return v.ID() }
func (v WhitePointY) AttrType() zcl.TypeID { return v.TypeID() }
func (v *WhitePointY) AttrValue() zcl.Val  { return v.Value() }

func (WhitePointY) Name() string { return `White Point Y` }
func (WhitePointY) Description() string {
	return `contains the normalized chromaticity value y for this attribute, as
defined in the CIE xyY Color Space. y = value / 65536 (value
in the range 0 to 65279 inclusive)`
}

// WhitePointY contains the normalized chromaticity value y for this attribute, as
// defined in the CIE xyY Color Space. y = value / 65536 (value
// in the range 0 to 65279 inclusive)
type WhitePointY zcl.Zu16

func (v *WhitePointY) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *WhitePointY) Value() zcl.Val     { return v }

func (v WhitePointY) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *WhitePointY) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = WhitePointY(*nt)
	return br, err
}

func (v WhitePointY) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *WhitePointY) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = WhitePointY(*a)
	return nil
}

func (v *WhitePointY) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = WhitePointY(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v WhitePointY) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}
