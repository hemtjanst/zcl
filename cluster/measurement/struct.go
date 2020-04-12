// provides generic measurement and sensing interfaces
package measurement

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

const IlluminanceLightSensorTypeAttr zcl.AttrID = 1

func (IlluminanceLightSensorType) ID() zcl.AttrID   { return IlluminanceLightSensorTypeAttr }
func (IlluminanceLightSensorType) Readable() bool   { return true }
func (IlluminanceLightSensorType) Writable() bool   { return false }
func (IlluminanceLightSensorType) Reportable() bool { return false }
func (IlluminanceLightSensorType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IlluminanceLightSensorType) AttrID() zcl.AttrID   { return v.ID() }
func (v IlluminanceLightSensorType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IlluminanceLightSensorType) AttrValue() zcl.Val  { return v.Value() }

func (IlluminanceLightSensorType) Name() string { return `Illuminance Light Sensor Type` }
func (IlluminanceLightSensorType) Description() string {
	return `specifies the electronic type of the light sensor. The range 0x40-0xfe is reserved for
manufacturer specific light sensor types`
}

// IlluminanceLightSensorType specifies the electronic type of the light sensor. The range 0x40-0xfe is reserved for
// manufacturer specific light sensor types
type IlluminanceLightSensorType zcl.Zenum8

func (v *IlluminanceLightSensorType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *IlluminanceLightSensorType) Value() zcl.Val     { return v }

func (v IlluminanceLightSensorType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *IlluminanceLightSensorType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = IlluminanceLightSensorType(*nt)
	return br, err
}

func (v IlluminanceLightSensorType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *IlluminanceLightSensorType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IlluminanceLightSensorType(*a)
	return nil
}

func (v *IlluminanceLightSensorType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = IlluminanceLightSensorType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IlluminanceLightSensorType) String() string {
	switch v {
	case 0x00:
		return "Photodiode"
	case 0x01:
		return "CMOS"
	case 0xff:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v IlluminanceLightSensorType) IsPhotodiode() bool { return v == 0x00 }
func (v IlluminanceLightSensorType) IsCmos() bool       { return v == 0x01 }
func (v IlluminanceLightSensorType) IsUnknown() bool    { return v == 0xff }
func (v *IlluminanceLightSensorType) SetPhotodiode()    { *v = 0x00 }
func (v *IlluminanceLightSensorType) SetCmos()          { *v = 0x01 }
func (v *IlluminanceLightSensorType) SetUnknown()       { *v = 0xff }

func (IlluminanceLightSensorType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Photodiode"},
		{Value: 0x01, Name: "CMOS"},
		{Value: 0xff, Name: "Unknown"},
	}
}

const IlluminanceTargetLevelAttr zcl.AttrID = 16

func (IlluminanceTargetLevel) ID() zcl.AttrID   { return IlluminanceTargetLevelAttr }
func (IlluminanceTargetLevel) Readable() bool   { return true }
func (IlluminanceTargetLevel) Writable() bool   { return true }
func (IlluminanceTargetLevel) Reportable() bool { return false }
func (IlluminanceTargetLevel) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v IlluminanceTargetLevel) AttrID() zcl.AttrID   { return v.ID() }
func (v IlluminanceTargetLevel) AttrType() zcl.TypeID { return v.TypeID() }
func (v *IlluminanceTargetLevel) AttrValue() zcl.Val  { return v.Value() }

func (IlluminanceTargetLevel) Name() string { return `Illuminance Target Level` }
func (IlluminanceTargetLevel) Description() string {
	return `specifies the target Illuminance level. This target level is taken as the
centre of a 'dead band', which must be sufficient in width, with hysteresis
bands at both top and bottom, to provide reliable notifications without
'chatter'. IlluminanceTargetLevel represents Illuminance in Lux (symbol lx) as
follows: IlluminanceTargetLevel = 10,000 x log 10 Illuminance`
}

// IlluminanceTargetLevel specifies the target Illuminance level. This target level is taken as the
// centre of a 'dead band', which must be sufficient in width, with hysteresis
// bands at both top and bottom, to provide reliable notifications without
// 'chatter'. IlluminanceTargetLevel represents Illuminance in Lux (symbol lx) as
// follows: IlluminanceTargetLevel = 10,000 x log 10 Illuminance
type IlluminanceTargetLevel zcl.Zu16

func (v *IlluminanceTargetLevel) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *IlluminanceTargetLevel) Value() zcl.Val     { return v }

func (v IlluminanceTargetLevel) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *IlluminanceTargetLevel) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = IlluminanceTargetLevel(*nt)
	return br, err
}

func (v IlluminanceTargetLevel) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *IlluminanceTargetLevel) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = IlluminanceTargetLevel(*a)
	return nil
}

func (v *IlluminanceTargetLevel) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = IlluminanceTargetLevel(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v IlluminanceTargetLevel) String() string {
	return zcl.Luxes.Format(float64(v))
}

const LevelStatusAttr zcl.AttrID = 0

func (LevelStatus) ID() zcl.AttrID   { return LevelStatusAttr }
func (LevelStatus) Readable() bool   { return true }
func (LevelStatus) Writable() bool   { return false }
func (LevelStatus) Reportable() bool { return true }
func (LevelStatus) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LevelStatus) AttrID() zcl.AttrID   { return v.ID() }
func (v LevelStatus) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LevelStatus) AttrValue() zcl.Val  { return v.Value() }

func (LevelStatus) Name() string { return `Level Status` }
func (LevelStatus) Description() string {
	return `indicates whether the measured Flow is above, below, or within a band
around FlowTargetLevel`
}

// LevelStatus indicates whether the measured Flow is above, below, or within a band
// around FlowTargetLevel
type LevelStatus zcl.Zenum8

func (v *LevelStatus) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *LevelStatus) Value() zcl.Val     { return v }

func (v LevelStatus) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *LevelStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = LevelStatus(*nt)
	return br, err
}

func (v LevelStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *LevelStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LevelStatus(*a)
	return nil
}

func (v *LevelStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = LevelStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LevelStatus) String() string {
	switch v {
	case 0x00:
		return "Flow on target"
	case 0x01:
		return "Flow below target"
	case 0x03:
		return "Flow above target"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v LevelStatus) IsFlowOnTarget() bool    { return v == 0x00 }
func (v LevelStatus) IsFlowBelowTarget() bool { return v == 0x01 }
func (v LevelStatus) IsFlowAboveTarget() bool { return v == 0x03 }
func (v *LevelStatus) SetFlowOnTarget()       { *v = 0x00 }
func (v *LevelStatus) SetFlowBelowTarget()    { *v = 0x01 }
func (v *LevelStatus) SetFlowAboveTarget()    { *v = 0x03 }

func (LevelStatus) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Flow on target"},
		{Value: 0x01, Name: "Flow below target"},
		{Value: 0x03, Name: "Flow above target"},
	}
}

const LightSensorTypeAttr zcl.AttrID = 4

func (LightSensorType) ID() zcl.AttrID   { return LightSensorTypeAttr }
func (LightSensorType) Readable() bool   { return true }
func (LightSensorType) Writable() bool   { return false }
func (LightSensorType) Reportable() bool { return false }
func (LightSensorType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LightSensorType) AttrID() zcl.AttrID   { return v.ID() }
func (v LightSensorType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LightSensorType) AttrValue() zcl.Val  { return v.Value() }

func (LightSensorType) Name() string { return `Light Sensor Type` }
func (LightSensorType) Description() string {
	return `specifies the electronic type of the light sensor. The range 0x40-0xfe is reserved for
manufacturer specific light sensor types`
}

// LightSensorType specifies the electronic type of the light sensor. The range 0x40-0xfe is reserved for
// manufacturer specific light sensor types
type LightSensorType zcl.Zenum8

func (v *LightSensorType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *LightSensorType) Value() zcl.Val     { return v }

func (v LightSensorType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *LightSensorType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = LightSensorType(*nt)
	return br, err
}

func (v LightSensorType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *LightSensorType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LightSensorType(*a)
	return nil
}

func (v *LightSensorType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = LightSensorType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LightSensorType) String() string {
	switch v {
	case 0x00:
		return "Photodiode"
	case 0x01:
		return "CMOS"
	case 0xff:
		return "Unknown"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v LightSensorType) IsPhotodiode() bool { return v == 0x00 }
func (v LightSensorType) IsCmos() bool       { return v == 0x01 }
func (v LightSensorType) IsUnknown() bool    { return v == 0xff }
func (v *LightSensorType) SetPhotodiode()    { *v = 0x00 }
func (v *LightSensorType) SetCmos()          { *v = 0x01 }
func (v *LightSensorType) SetUnknown()       { *v = 0xff }

func (LightSensorType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Photodiode"},
		{Value: 0x01, Name: "CMOS"},
		{Value: 0xff, Name: "Unknown"},
	}
}

const MaxMeasuredFlowAttr zcl.AttrID = 2

func (MaxMeasuredFlow) ID() zcl.AttrID   { return MaxMeasuredFlowAttr }
func (MaxMeasuredFlow) Readable() bool   { return true }
func (MaxMeasuredFlow) Writable() bool   { return false }
func (MaxMeasuredFlow) Reportable() bool { return false }
func (MaxMeasuredFlow) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxMeasuredFlow) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxMeasuredFlow) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxMeasuredFlow) AttrValue() zcl.Val  { return v.Value() }

func (MaxMeasuredFlow) Name() string { return `Max Measured Flow` }
func (MaxMeasuredFlow) Description() string {
	return `indicates the maximum value of MeasuredFlow that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MaxMeasuredFlow indicates the maximum value of MeasuredFlow that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MaxMeasuredFlow zcl.Zu16

func (v *MaxMeasuredFlow) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxMeasuredFlow) Value() zcl.Val     { return v }

func (v MaxMeasuredFlow) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxMeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxMeasuredFlow(*nt)
	return br, err
}

func (v MaxMeasuredFlow) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxMeasuredFlow) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxMeasuredFlow(*a)
	return nil
}

func (v *MaxMeasuredFlow) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxMeasuredFlow(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxMeasuredFlow) String() string {
	return zcl.CubicMetersPerHour.Format(float64(v) / 10)
}

func (v MaxMeasuredFlow) Scaled() float64 {
	return float64(v) / 10
}

const MaxMeasuredIlluminanceAttr zcl.AttrID = 2

func (MaxMeasuredIlluminance) ID() zcl.AttrID   { return MaxMeasuredIlluminanceAttr }
func (MaxMeasuredIlluminance) Readable() bool   { return true }
func (MaxMeasuredIlluminance) Writable() bool   { return false }
func (MaxMeasuredIlluminance) Reportable() bool { return false }
func (MaxMeasuredIlluminance) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxMeasuredIlluminance) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxMeasuredIlluminance) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxMeasuredIlluminance) AttrValue() zcl.Val  { return v.Value() }

func (MaxMeasuredIlluminance) Name() string { return `Max Measured Illuminance` }
func (MaxMeasuredIlluminance) Description() string {
	return `indicates the maximum value of MeasuredIlluminance that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MaxMeasuredIlluminance indicates the maximum value of MeasuredIlluminance that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MaxMeasuredIlluminance zcl.Zu16

func (v *MaxMeasuredIlluminance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxMeasuredIlluminance) Value() zcl.Val     { return v }

func (v MaxMeasuredIlluminance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxMeasuredIlluminance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxMeasuredIlluminance(*nt)
	return br, err
}

func (v MaxMeasuredIlluminance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxMeasuredIlluminance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxMeasuredIlluminance(*a)
	return nil
}

func (v *MaxMeasuredIlluminance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxMeasuredIlluminance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxMeasuredIlluminance) String() string {
	return zcl.Luxes.Format(float64(v))
}

const MaxMeasuredPressureAttr zcl.AttrID = 2

func (MaxMeasuredPressure) ID() zcl.AttrID   { return MaxMeasuredPressureAttr }
func (MaxMeasuredPressure) Readable() bool   { return true }
func (MaxMeasuredPressure) Writable() bool   { return false }
func (MaxMeasuredPressure) Reportable() bool { return false }
func (MaxMeasuredPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxMeasuredPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxMeasuredPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxMeasuredPressure) AttrValue() zcl.Val  { return v.Value() }

func (MaxMeasuredPressure) Name() string { return `Max Measured Pressure` }
func (MaxMeasuredPressure) Description() string {
	return `indicates the maximum value of MeasuredPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// MaxMeasuredPressure indicates the maximum value of MeasuredPressure that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type MaxMeasuredPressure zcl.Zs16

func (v *MaxMeasuredPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxMeasuredPressure) Value() zcl.Val     { return v }

func (v MaxMeasuredPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxMeasuredPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxMeasuredPressure(*nt)
	return br, err
}

func (v MaxMeasuredPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxMeasuredPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxMeasuredPressure(*a)
	return nil
}

func (v *MaxMeasuredPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxMeasuredPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxMeasuredPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MaxMeasuredPressure) Scaled() float64 {
	return float64(v) / 10
}

const MaxMeasuredRelativeHumidityAttr zcl.AttrID = 2

func (MaxMeasuredRelativeHumidity) ID() zcl.AttrID   { return MaxMeasuredRelativeHumidityAttr }
func (MaxMeasuredRelativeHumidity) Readable() bool   { return true }
func (MaxMeasuredRelativeHumidity) Writable() bool   { return false }
func (MaxMeasuredRelativeHumidity) Reportable() bool { return false }
func (MaxMeasuredRelativeHumidity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxMeasuredRelativeHumidity) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxMeasuredRelativeHumidity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxMeasuredRelativeHumidity) AttrValue() zcl.Val  { return v.Value() }

func (MaxMeasuredRelativeHumidity) Name() string { return `Max Measured Relative Humidity` }
func (MaxMeasuredRelativeHumidity) Description() string {
	return `indicates the maximum value of MeasuredRH that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MaxMeasuredRelativeHumidity indicates the maximum value of MeasuredRH that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MaxMeasuredRelativeHumidity zcl.Zu16

func (v *MaxMeasuredRelativeHumidity) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxMeasuredRelativeHumidity) Value() zcl.Val     { return v }

func (v MaxMeasuredRelativeHumidity) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxMeasuredRelativeHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxMeasuredRelativeHumidity(*nt)
	return br, err
}

func (v MaxMeasuredRelativeHumidity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxMeasuredRelativeHumidity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxMeasuredRelativeHumidity(*a)
	return nil
}

func (v *MaxMeasuredRelativeHumidity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxMeasuredRelativeHumidity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxMeasuredRelativeHumidity) String() string {
	return zcl.PercentRelativeHumidity.Format(float64(v) / 100)
}

func (v MaxMeasuredRelativeHumidity) Scaled() float64 {
	return float64(v) / 100
}

const MaxMeasuredTemperatureAttr zcl.AttrID = 2

func (MaxMeasuredTemperature) ID() zcl.AttrID   { return MaxMeasuredTemperatureAttr }
func (MaxMeasuredTemperature) Readable() bool   { return true }
func (MaxMeasuredTemperature) Writable() bool   { return false }
func (MaxMeasuredTemperature) Reportable() bool { return false }
func (MaxMeasuredTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxMeasuredTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxMeasuredTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxMeasuredTemperature) AttrValue() zcl.Val  { return v.Value() }

func (MaxMeasuredTemperature) Name() string { return `Max Measured Temperature` }
func (MaxMeasuredTemperature) Description() string {
	return `indicates the maximum value of MeasuredTemperature that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// MaxMeasuredTemperature indicates the maximum value of MeasuredTemperature that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type MaxMeasuredTemperature zcl.Zs16

func (v *MaxMeasuredTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxMeasuredTemperature) Value() zcl.Val     { return v }

func (v MaxMeasuredTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxMeasuredTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxMeasuredTemperature(*nt)
	return br, err
}

func (v MaxMeasuredTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxMeasuredTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxMeasuredTemperature(*a)
	return nil
}

func (v *MaxMeasuredTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxMeasuredTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxMeasuredTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v MaxMeasuredTemperature) Scaled() float64 {
	return float64(v) / 100
}

const MeasuredFlowAttr zcl.AttrID = 0

func (MeasuredFlow) ID() zcl.AttrID   { return MeasuredFlowAttr }
func (MeasuredFlow) Readable() bool   { return true }
func (MeasuredFlow) Writable() bool   { return false }
func (MeasuredFlow) Reportable() bool { return true }
func (MeasuredFlow) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredFlow) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredFlow) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredFlow) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredFlow) Name() string        { return `Measured Flow` }
func (MeasuredFlow) Description() string { return `represents the flow in m^3/h` }

// MeasuredFlow represents the flow in m^3/h
type MeasuredFlow zcl.Zu16

func (v *MeasuredFlow) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MeasuredFlow) Value() zcl.Val     { return v }

func (v MeasuredFlow) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredFlow(*nt)
	return br, err
}

func (v MeasuredFlow) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MeasuredFlow) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredFlow(*a)
	return nil
}

func (v *MeasuredFlow) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MeasuredFlow(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredFlow) String() string {
	return zcl.CubicMetersPerHour.Format(float64(v) / 10)
}

func (v MeasuredFlow) Scaled() float64 {
	return float64(v) / 10
}

const MeasuredIlluminanceAttr zcl.AttrID = 0

func (MeasuredIlluminance) ID() zcl.AttrID   { return MeasuredIlluminanceAttr }
func (MeasuredIlluminance) Readable() bool   { return true }
func (MeasuredIlluminance) Writable() bool   { return false }
func (MeasuredIlluminance) Reportable() bool { return true }
func (MeasuredIlluminance) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredIlluminance) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredIlluminance) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredIlluminance) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredIlluminance) Name() string { return `Measured Illuminance` }
func (MeasuredIlluminance) Description() string {
	return `represents the Illuminance in Lux (symbol lx) as follows
MeasuredValue = 10,000 x log10 Flow + 1 where 1 lx
<= Flow <=3.576 Mlx, corresponding to a MeasuredValue
in the range 1 to 0xfffe`
}

// MeasuredIlluminance represents the Illuminance in Lux (symbol lx) as follows
// MeasuredValue = 10,000 x log10 Flow + 1 where 1 lx
// <= Flow <=3.576 Mlx, corresponding to a MeasuredValue
// in the range 1 to 0xfffe
type MeasuredIlluminance zcl.Zu16

func (v *MeasuredIlluminance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MeasuredIlluminance) Value() zcl.Val     { return v }

func (v MeasuredIlluminance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MeasuredIlluminance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredIlluminance(*nt)
	return br, err
}

func (v MeasuredIlluminance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MeasuredIlluminance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredIlluminance(*a)
	return nil
}

func (v *MeasuredIlluminance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MeasuredIlluminance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredIlluminance) String() string {
	return zcl.Luxes.Format(float64(v))
}

const MeasuredPressureAttr zcl.AttrID = 0

func (MeasuredPressure) ID() zcl.AttrID   { return MeasuredPressureAttr }
func (MeasuredPressure) Readable() bool   { return true }
func (MeasuredPressure) Writable() bool   { return false }
func (MeasuredPressure) Reportable() bool { return true }
func (MeasuredPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredPressure) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredPressure) Name() string { return `Measured Pressure` }
func (MeasuredPressure) Description() string {
	return `represents the temperature in degrees DegreesCelsius`
}

// MeasuredPressure represents the temperature in degrees DegreesCelsius
type MeasuredPressure zcl.Zs16

func (v *MeasuredPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MeasuredPressure) Value() zcl.Val     { return v }

func (v MeasuredPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MeasuredPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredPressure(*nt)
	return br, err
}

func (v MeasuredPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MeasuredPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredPressure(*a)
	return nil
}

func (v *MeasuredPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MeasuredPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MeasuredPressure) Scaled() float64 {
	return float64(v) / 10
}

const MeasuredRelativeHumidityAttr zcl.AttrID = 0

func (MeasuredRelativeHumidity) ID() zcl.AttrID   { return MeasuredRelativeHumidityAttr }
func (MeasuredRelativeHumidity) Readable() bool   { return true }
func (MeasuredRelativeHumidity) Writable() bool   { return false }
func (MeasuredRelativeHumidity) Reportable() bool { return true }
func (MeasuredRelativeHumidity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredRelativeHumidity) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredRelativeHumidity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredRelativeHumidity) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredRelativeHumidity) Name() string { return `Measured Relative Humidity` }
func (MeasuredRelativeHumidity) Description() string {
	return `represents the relative humidity in percent`
}

// MeasuredRelativeHumidity represents the relative humidity in percent
type MeasuredRelativeHumidity zcl.Zu16

func (v *MeasuredRelativeHumidity) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MeasuredRelativeHumidity) Value() zcl.Val     { return v }

func (v MeasuredRelativeHumidity) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MeasuredRelativeHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredRelativeHumidity(*nt)
	return br, err
}

func (v MeasuredRelativeHumidity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MeasuredRelativeHumidity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredRelativeHumidity(*a)
	return nil
}

func (v *MeasuredRelativeHumidity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MeasuredRelativeHumidity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredRelativeHumidity) String() string {
	return zcl.PercentRelativeHumidity.Format(float64(v) / 100)
}

func (v MeasuredRelativeHumidity) Scaled() float64 {
	return float64(v) / 100
}

const MeasuredTemperatureAttr zcl.AttrID = 0

func (MeasuredTemperature) ID() zcl.AttrID   { return MeasuredTemperatureAttr }
func (MeasuredTemperature) Readable() bool   { return true }
func (MeasuredTemperature) Writable() bool   { return false }
func (MeasuredTemperature) Reportable() bool { return true }
func (MeasuredTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MeasuredTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v MeasuredTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MeasuredTemperature) AttrValue() zcl.Val  { return v.Value() }

func (MeasuredTemperature) Name() string { return `Measured Temperature` }
func (MeasuredTemperature) Description() string {
	return `represents the temperature in degrees DegreesCelsius`
}

// MeasuredTemperature represents the temperature in degrees DegreesCelsius
type MeasuredTemperature zcl.Zs16

func (v *MeasuredTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MeasuredTemperature) Value() zcl.Val     { return v }

func (v MeasuredTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MeasuredTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MeasuredTemperature(*nt)
	return br, err
}

func (v MeasuredTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MeasuredTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MeasuredTemperature(*a)
	return nil
}

func (v *MeasuredTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MeasuredTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MeasuredTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v MeasuredTemperature) Scaled() float64 {
	return float64(v) / 100
}

const MinMeasuredFlowAttr zcl.AttrID = 1

func (MinMeasuredFlow) ID() zcl.AttrID   { return MinMeasuredFlowAttr }
func (MinMeasuredFlow) Readable() bool   { return true }
func (MinMeasuredFlow) Writable() bool   { return false }
func (MinMeasuredFlow) Reportable() bool { return false }
func (MinMeasuredFlow) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinMeasuredFlow) AttrID() zcl.AttrID   { return v.ID() }
func (v MinMeasuredFlow) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinMeasuredFlow) AttrValue() zcl.Val  { return v.Value() }

func (MinMeasuredFlow) Name() string { return `Min Measured Flow` }
func (MinMeasuredFlow) Description() string {
	return `indicates the minimum value of MeasuredFlow that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MinMeasuredFlow indicates the minimum value of MeasuredFlow that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MinMeasuredFlow zcl.Zu16

func (v *MinMeasuredFlow) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MinMeasuredFlow) Value() zcl.Val     { return v }

func (v MinMeasuredFlow) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MinMeasuredFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinMeasuredFlow(*nt)
	return br, err
}

func (v MinMeasuredFlow) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MinMeasuredFlow) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinMeasuredFlow(*a)
	return nil
}

func (v *MinMeasuredFlow) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MinMeasuredFlow(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinMeasuredFlow) String() string {
	return zcl.CubicMetersPerHour.Format(float64(v) / 10)
}

func (v MinMeasuredFlow) Scaled() float64 {
	return float64(v) / 10
}

const MinMeasuredIlluminanceAttr zcl.AttrID = 1

func (MinMeasuredIlluminance) ID() zcl.AttrID   { return MinMeasuredIlluminanceAttr }
func (MinMeasuredIlluminance) Readable() bool   { return true }
func (MinMeasuredIlluminance) Writable() bool   { return false }
func (MinMeasuredIlluminance) Reportable() bool { return false }
func (MinMeasuredIlluminance) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinMeasuredIlluminance) AttrID() zcl.AttrID   { return v.ID() }
func (v MinMeasuredIlluminance) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinMeasuredIlluminance) AttrValue() zcl.Val  { return v.Value() }

func (MinMeasuredIlluminance) Name() string { return `Min Measured Illuminance` }
func (MinMeasuredIlluminance) Description() string {
	return `indicates the minimum value of MeasuredIlluminance that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MinMeasuredIlluminance indicates the minimum value of MeasuredIlluminance that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MinMeasuredIlluminance zcl.Zu16

func (v *MinMeasuredIlluminance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MinMeasuredIlluminance) Value() zcl.Val     { return v }

func (v MinMeasuredIlluminance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MinMeasuredIlluminance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinMeasuredIlluminance(*nt)
	return br, err
}

func (v MinMeasuredIlluminance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MinMeasuredIlluminance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinMeasuredIlluminance(*a)
	return nil
}

func (v *MinMeasuredIlluminance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MinMeasuredIlluminance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinMeasuredIlluminance) String() string {
	return zcl.Luxes.Format(float64(v))
}

const MinMeasuredPressureAttr zcl.AttrID = 1

func (MinMeasuredPressure) ID() zcl.AttrID   { return MinMeasuredPressureAttr }
func (MinMeasuredPressure) Readable() bool   { return true }
func (MinMeasuredPressure) Writable() bool   { return false }
func (MinMeasuredPressure) Reportable() bool { return false }
func (MinMeasuredPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinMeasuredPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MinMeasuredPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinMeasuredPressure) AttrValue() zcl.Val  { return v.Value() }

func (MinMeasuredPressure) Name() string { return `Min Measured Pressure` }
func (MinMeasuredPressure) Description() string {
	return `indicates the minimum value of MeasuredPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// MinMeasuredPressure indicates the minimum value of MeasuredPressure that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type MinMeasuredPressure zcl.Zs16

func (v *MinMeasuredPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MinMeasuredPressure) Value() zcl.Val     { return v }

func (v MinMeasuredPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MinMeasuredPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinMeasuredPressure(*nt)
	return br, err
}

func (v MinMeasuredPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MinMeasuredPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinMeasuredPressure(*a)
	return nil
}

func (v *MinMeasuredPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MinMeasuredPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinMeasuredPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MinMeasuredPressure) Scaled() float64 {
	return float64(v) / 10
}

const MinMeasuredRelativeHumidityAttr zcl.AttrID = 1

func (MinMeasuredRelativeHumidity) ID() zcl.AttrID   { return MinMeasuredRelativeHumidityAttr }
func (MinMeasuredRelativeHumidity) Readable() bool   { return true }
func (MinMeasuredRelativeHumidity) Writable() bool   { return false }
func (MinMeasuredRelativeHumidity) Reportable() bool { return false }
func (MinMeasuredRelativeHumidity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinMeasuredRelativeHumidity) AttrID() zcl.AttrID   { return v.ID() }
func (v MinMeasuredRelativeHumidity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinMeasuredRelativeHumidity) AttrValue() zcl.Val  { return v.Value() }

func (MinMeasuredRelativeHumidity) Name() string { return `Min Measured Relative Humidity` }
func (MinMeasuredRelativeHumidity) Description() string {
	return `indicates the minimum value of MeasuredRH that can be measured. A
value of 0xffff indicates that this attribute is not defined`
}

// MinMeasuredRelativeHumidity indicates the minimum value of MeasuredRH that can be measured. A
// value of 0xffff indicates that this attribute is not defined
type MinMeasuredRelativeHumidity zcl.Zu16

func (v *MinMeasuredRelativeHumidity) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MinMeasuredRelativeHumidity) Value() zcl.Val     { return v }

func (v MinMeasuredRelativeHumidity) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MinMeasuredRelativeHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinMeasuredRelativeHumidity(*nt)
	return br, err
}

func (v MinMeasuredRelativeHumidity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MinMeasuredRelativeHumidity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinMeasuredRelativeHumidity(*a)
	return nil
}

func (v *MinMeasuredRelativeHumidity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MinMeasuredRelativeHumidity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinMeasuredRelativeHumidity) String() string {
	return zcl.PercentRelativeHumidity.Format(float64(v) / 100)
}

func (v MinMeasuredRelativeHumidity) Scaled() float64 {
	return float64(v) / 100
}

const MinMeasuredTemperatureAttr zcl.AttrID = 1

func (MinMeasuredTemperature) ID() zcl.AttrID   { return MinMeasuredTemperatureAttr }
func (MinMeasuredTemperature) Readable() bool   { return true }
func (MinMeasuredTemperature) Writable() bool   { return false }
func (MinMeasuredTemperature) Reportable() bool { return false }
func (MinMeasuredTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinMeasuredTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v MinMeasuredTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinMeasuredTemperature) AttrValue() zcl.Val  { return v.Value() }

func (MinMeasuredTemperature) Name() string { return `Min Measured Temperature` }
func (MinMeasuredTemperature) Description() string {
	return `indicates the minimum value of MeasuredTemperature that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// MinMeasuredTemperature indicates the minimum value of MeasuredTemperature that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type MinMeasuredTemperature zcl.Zs16

func (v *MinMeasuredTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MinMeasuredTemperature) Value() zcl.Val     { return v }

func (v MinMeasuredTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MinMeasuredTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinMeasuredTemperature(*nt)
	return br, err
}

func (v MinMeasuredTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MinMeasuredTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinMeasuredTemperature(*a)
	return nil
}

func (v *MinMeasuredTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MinMeasuredTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinMeasuredTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v MinMeasuredTemperature) Scaled() float64 {
	return float64(v) / 100
}

const OccupancyAttr zcl.AttrID = 0

func (Occupancy) ID() zcl.AttrID   { return OccupancyAttr }
func (Occupancy) Readable() bool   { return true }
func (Occupancy) Writable() bool   { return false }
func (Occupancy) Reportable() bool { return true }
func (Occupancy) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Occupancy) AttrID() zcl.AttrID   { return v.ID() }
func (v Occupancy) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Occupancy) AttrValue() zcl.Val  { return v.Value() }

func (Occupancy) Name() string        { return `Occupancy` }
func (Occupancy) Description() string { return `` }

type Occupancy zcl.Zbmp8

func (v *Occupancy) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *Occupancy) Value() zcl.Val     { return v }

func (v Occupancy) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *Occupancy) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = Occupancy(*nt)
	return br, err
}

func (v Occupancy) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *Occupancy) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Occupancy(*a)
	return nil
}

func (v *Occupancy) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = Occupancy(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Occupancy) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Occupied")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v Occupancy) IsOccupied() bool    { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v *Occupancy) SetOccupied(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }

func (Occupancy) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Occupied"},
	}
}

const OccupancyTypeAttr zcl.AttrID = 1

func (OccupancyType) ID() zcl.AttrID   { return OccupancyTypeAttr }
func (OccupancyType) Readable() bool   { return true }
func (OccupancyType) Writable() bool   { return false }
func (OccupancyType) Reportable() bool { return false }
func (OccupancyType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OccupancyType) AttrID() zcl.AttrID   { return v.ID() }
func (v OccupancyType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OccupancyType) AttrValue() zcl.Val  { return v.Value() }

func (OccupancyType) Name() string        { return `Occupancy Type` }
func (OccupancyType) Description() string { return `` }

type OccupancyType zcl.Zenum8

func (v *OccupancyType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *OccupancyType) Value() zcl.Val     { return v }

func (v OccupancyType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *OccupancyType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = OccupancyType(*nt)
	return br, err
}

func (v OccupancyType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *OccupancyType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OccupancyType(*a)
	return nil
}

func (v *OccupancyType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = OccupancyType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OccupancyType) String() string {
	switch v {
	case 0x00:
		return "PIR"
	case 0x01:
		return "Ultrasonic"
	case 0x03:
		return "PIR and ultrasonic"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v OccupancyType) IsPir() bool              { return v == 0x00 }
func (v OccupancyType) IsUltrasonic() bool       { return v == 0x01 }
func (v OccupancyType) IsPirAndUltrasonic() bool { return v == 0x03 }
func (v *OccupancyType) SetPir()                 { *v = 0x00 }
func (v *OccupancyType) SetUltrasonic()          { *v = 0x01 }
func (v *OccupancyType) SetPirAndUltrasonic()    { *v = 0x03 }

func (OccupancyType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "PIR"},
		{Value: 0x01, Name: "Ultrasonic"},
		{Value: 0x03, Name: "PIR and ultrasonic"},
	}
}

const PirOccupiedToUnoccupiedDelayAttr zcl.AttrID = 16

func (PirOccupiedToUnoccupiedDelay) ID() zcl.AttrID   { return PirOccupiedToUnoccupiedDelayAttr }
func (PirOccupiedToUnoccupiedDelay) Readable() bool   { return true }
func (PirOccupiedToUnoccupiedDelay) Writable() bool   { return true }
func (PirOccupiedToUnoccupiedDelay) Reportable() bool { return false }
func (PirOccupiedToUnoccupiedDelay) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PirOccupiedToUnoccupiedDelay) AttrID() zcl.AttrID   { return v.ID() }
func (v PirOccupiedToUnoccupiedDelay) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PirOccupiedToUnoccupiedDelay) AttrValue() zcl.Val  { return v.Value() }

func (PirOccupiedToUnoccupiedDelay) Name() string { return `PIR Occupied To Unoccupied Delay` }
func (PirOccupiedToUnoccupiedDelay) Description() string {
	return `is 16 bits in length and specifies the time delay, in seconds, before the PIR sensor
changes to its unoccupied state after the last detection of movement in the sensed area`
}

// PirOccupiedToUnoccupiedDelay is 16 bits in length and specifies the time delay, in seconds, before the PIR sensor
// changes to its unoccupied state after the last detection of movement in the sensed area
type PirOccupiedToUnoccupiedDelay zcl.Zu16

func (v *PirOccupiedToUnoccupiedDelay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PirOccupiedToUnoccupiedDelay) Value() zcl.Val     { return v }

func (v PirOccupiedToUnoccupiedDelay) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PirOccupiedToUnoccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PirOccupiedToUnoccupiedDelay(*nt)
	return br, err
}

func (v PirOccupiedToUnoccupiedDelay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PirOccupiedToUnoccupiedDelay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PirOccupiedToUnoccupiedDelay(*a)
	return nil
}

func (v *PirOccupiedToUnoccupiedDelay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PirOccupiedToUnoccupiedDelay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PirOccupiedToUnoccupiedDelay) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PirUnoccupiedToOccupiedDelayAttr zcl.AttrID = 17

func (PirUnoccupiedToOccupiedDelay) ID() zcl.AttrID   { return PirUnoccupiedToOccupiedDelayAttr }
func (PirUnoccupiedToOccupiedDelay) Readable() bool   { return true }
func (PirUnoccupiedToOccupiedDelay) Writable() bool   { return true }
func (PirUnoccupiedToOccupiedDelay) Reportable() bool { return false }
func (PirUnoccupiedToOccupiedDelay) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PirUnoccupiedToOccupiedDelay) AttrID() zcl.AttrID   { return v.ID() }
func (v PirUnoccupiedToOccupiedDelay) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PirUnoccupiedToOccupiedDelay) AttrValue() zcl.Val  { return v.Value() }

func (PirUnoccupiedToOccupiedDelay) Name() string { return `PIR Unoccupied To Occupied Delay` }
func (PirUnoccupiedToOccupiedDelay) Description() string {
	return `is 16 bits in length and specifies the time delay, in seconds, before the PIR sensor changes
to its occupied state after the detection of movement in the sensed area. This attribute is
mandatory if the PIRUnoccupiedToOccupiedThreshold attribute is implemented`
}

// PirUnoccupiedToOccupiedDelay is 16 bits in length and specifies the time delay, in seconds, before the PIR sensor changes
// to its occupied state after the detection of movement in the sensed area. This attribute is
// mandatory if the PIRUnoccupiedToOccupiedThreshold attribute is implemented
type PirUnoccupiedToOccupiedDelay zcl.Zu16

func (v *PirUnoccupiedToOccupiedDelay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *PirUnoccupiedToOccupiedDelay) Value() zcl.Val     { return v }

func (v PirUnoccupiedToOccupiedDelay) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *PirUnoccupiedToOccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = PirUnoccupiedToOccupiedDelay(*nt)
	return br, err
}

func (v PirUnoccupiedToOccupiedDelay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *PirUnoccupiedToOccupiedDelay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PirUnoccupiedToOccupiedDelay(*a)
	return nil
}

func (v *PirUnoccupiedToOccupiedDelay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = PirUnoccupiedToOccupiedDelay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PirUnoccupiedToOccupiedDelay) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const PirUnoccupiedToOccupiedThresholdAttr zcl.AttrID = 18

func (PirUnoccupiedToOccupiedThreshold) ID() zcl.AttrID   { return PirUnoccupiedToOccupiedThresholdAttr }
func (PirUnoccupiedToOccupiedThreshold) Readable() bool   { return true }
func (PirUnoccupiedToOccupiedThreshold) Writable() bool   { return true }
func (PirUnoccupiedToOccupiedThreshold) Reportable() bool { return false }
func (PirUnoccupiedToOccupiedThreshold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PirUnoccupiedToOccupiedThreshold) AttrID() zcl.AttrID   { return v.ID() }
func (v PirUnoccupiedToOccupiedThreshold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PirUnoccupiedToOccupiedThreshold) AttrValue() zcl.Val  { return v.Value() }

func (PirUnoccupiedToOccupiedThreshold) Name() string { return `PIR Unoccupied To Occupied Threshold` }
func (PirUnoccupiedToOccupiedThreshold) Description() string {
	return `is 8 bits in length and specifies the number of movement detection events that must occur in the period PIRUnoccupiedToOccupiedDelay, before the PIR sensor changes to its occupied state. This attribute is mandatory if the PIRUnoccupiedToOccupiedDelay attribute is implemented`
}

// PirUnoccupiedToOccupiedThreshold is 8 bits in length and specifies the number of movement detection events that must occur in the period PIRUnoccupiedToOccupiedDelay, before the PIR sensor changes to its occupied state. This attribute is mandatory if the PIRUnoccupiedToOccupiedDelay attribute is implemented
type PirUnoccupiedToOccupiedThreshold zcl.Zu8

func (v *PirUnoccupiedToOccupiedThreshold) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *PirUnoccupiedToOccupiedThreshold) Value() zcl.Val     { return v }

func (v PirUnoccupiedToOccupiedThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(v).MarshalZcl()
}

func (v *PirUnoccupiedToOccupiedThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = PirUnoccupiedToOccupiedThreshold(*nt)
	return br, err
}

func (v PirUnoccupiedToOccupiedThreshold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *PirUnoccupiedToOccupiedThreshold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PirUnoccupiedToOccupiedThreshold(*a)
	return nil
}

func (v *PirUnoccupiedToOccupiedThreshold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = PirUnoccupiedToOccupiedThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PirUnoccupiedToOccupiedThreshold) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const ScaleAttr zcl.AttrID = 20

func (Scale) ID() zcl.AttrID   { return ScaleAttr }
func (Scale) Readable() bool   { return true }
func (Scale) Writable() bool   { return false }
func (Scale) Reportable() bool { return false }
func (Scale) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Scale) AttrID() zcl.AttrID   { return v.ID() }
func (v Scale) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Scale) AttrValue() zcl.Val  { return v.Value() }

func (Scale) Name() string { return `Scale` }
func (Scale) Description() string {
	return `indicates the base 10 exponent used to obtain ScaledPressure`
}

// Scale indicates the base 10 exponent used to obtain ScaledPressure
type Scale zcl.Zs8

func (v *Scale) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *Scale) Value() zcl.Val     { return v }

func (v Scale) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *Scale) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = Scale(*nt)
	return br, err
}

func (v Scale) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *Scale) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Scale(*a)
	return nil
}

func (v *Scale) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = Scale(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Scale) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const ScaledMaxPressureAttr zcl.AttrID = 18

func (ScaledMaxPressure) ID() zcl.AttrID   { return ScaledMaxPressureAttr }
func (ScaledMaxPressure) Readable() bool   { return true }
func (ScaledMaxPressure) Writable() bool   { return false }
func (ScaledMaxPressure) Reportable() bool { return false }
func (ScaledMaxPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ScaledMaxPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v ScaledMaxPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ScaledMaxPressure) AttrValue() zcl.Val  { return v.Value() }

func (ScaledMaxPressure) Name() string { return `Scaled Max Pressure` }
func (ScaledMaxPressure) Description() string {
	return `indicates the maximum value of ScaledPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// ScaledMaxPressure indicates the maximum value of ScaledPressure that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type ScaledMaxPressure zcl.Zs16

func (v *ScaledMaxPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ScaledMaxPressure) Value() zcl.Val     { return v }

func (v ScaledMaxPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ScaledMaxPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ScaledMaxPressure(*nt)
	return br, err
}

func (v ScaledMaxPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ScaledMaxPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ScaledMaxPressure(*a)
	return nil
}

func (v *ScaledMaxPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ScaledMaxPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ScaledMaxPressure) String() string {
	return zcl.Pascals.Format(float64(v))
}

const ScaledMinPressureAttr zcl.AttrID = 17

func (ScaledMinPressure) ID() zcl.AttrID   { return ScaledMinPressureAttr }
func (ScaledMinPressure) Readable() bool   { return true }
func (ScaledMinPressure) Writable() bool   { return false }
func (ScaledMinPressure) Reportable() bool { return false }
func (ScaledMinPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ScaledMinPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v ScaledMinPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ScaledMinPressure) AttrValue() zcl.Val  { return v.Value() }

func (ScaledMinPressure) Name() string { return `Scaled Min Pressure` }
func (ScaledMinPressure) Description() string {
	return `indicates the minimum value of ScaledPressure that can be measured. A
value of 0x8000 indicates that this attribute is not defined`
}

// ScaledMinPressure indicates the minimum value of ScaledPressure that can be measured. A
// value of 0x8000 indicates that this attribute is not defined
type ScaledMinPressure zcl.Zs16

func (v *ScaledMinPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ScaledMinPressure) Value() zcl.Val     { return v }

func (v ScaledMinPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ScaledMinPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ScaledMinPressure(*nt)
	return br, err
}

func (v ScaledMinPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ScaledMinPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ScaledMinPressure(*a)
	return nil
}

func (v *ScaledMinPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ScaledMinPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ScaledMinPressure) String() string {
	return zcl.Pascals.Format(float64(v))
}

const ScaledPressureAttr zcl.AttrID = 16

func (ScaledPressure) ID() zcl.AttrID   { return ScaledPressureAttr }
func (ScaledPressure) Readable() bool   { return true }
func (ScaledPressure) Writable() bool   { return false }
func (ScaledPressure) Reportable() bool { return true }
func (ScaledPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ScaledPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v ScaledPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ScaledPressure) AttrValue() zcl.Val  { return v.Value() }

func (ScaledPressure) Name() string { return `Scaled Pressure` }
func (ScaledPressure) Description() string {
	return `represents the pressure in Pascals as follows:
ScaledPressure = 10^Scale x Pressure in Pa`
}

// ScaledPressure represents the pressure in Pascals as follows:
// ScaledPressure = 10^Scale x Pressure in Pa
type ScaledPressure zcl.Zs16

func (v *ScaledPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *ScaledPressure) Value() zcl.Val     { return v }

func (v ScaledPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *ScaledPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = ScaledPressure(*nt)
	return br, err
}

func (v ScaledPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *ScaledPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ScaledPressure(*a)
	return nil
}

func (v *ScaledPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = ScaledPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ScaledPressure) String() string {
	return zcl.Pascals.Format(float64(v))
}

const ScaledToleranceAttr zcl.AttrID = 19

func (ScaledTolerance) ID() zcl.AttrID   { return ScaledToleranceAttr }
func (ScaledTolerance) Readable() bool   { return true }
func (ScaledTolerance) Writable() bool   { return false }
func (ScaledTolerance) Reportable() bool { return true }
func (ScaledTolerance) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ScaledTolerance) AttrID() zcl.AttrID   { return v.ID() }
func (v ScaledTolerance) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ScaledTolerance) AttrValue() zcl.Val  { return v.Value() }

func (ScaledTolerance) Name() string { return `Scaled Tolerance` }
func (ScaledTolerance) Description() string {
	return `indicates the magnitude of the possible error that is associated with ScaledPressure.
The true value is located in the range (ScaledPressure – ScaledTolerance) to
(ScaledPressure + ScaledTolerance)`
}

// ScaledTolerance indicates the magnitude of the possible error that is associated with ScaledPressure.
// The true value is located in the range (ScaledPressure – ScaledTolerance) to
// (ScaledPressure + ScaledTolerance)
type ScaledTolerance zcl.Zu16

func (v *ScaledTolerance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *ScaledTolerance) Value() zcl.Val     { return v }

func (v ScaledTolerance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *ScaledTolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = ScaledTolerance(*nt)
	return br, err
}

func (v ScaledTolerance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *ScaledTolerance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ScaledTolerance(*a)
	return nil
}

func (v *ScaledTolerance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = ScaledTolerance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ScaledTolerance) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const ToleranceAttr zcl.AttrID = 3

func (Tolerance) ID() zcl.AttrID   { return ToleranceAttr }
func (Tolerance) Readable() bool   { return true }
func (Tolerance) Writable() bool   { return false }
func (Tolerance) Reportable() bool { return true }
func (Tolerance) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Tolerance) AttrID() zcl.AttrID   { return v.ID() }
func (v Tolerance) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Tolerance) AttrValue() zcl.Val  { return v.Value() }

func (Tolerance) Name() string { return `Tolerance` }
func (Tolerance) Description() string {
	return `indicates the magnitude of the possible error that is associated with MeasuredValue.
The true value is located in the range (MeasuredValue – Tolerance) to
(MeasuredValue + Tolerance)`
}

// Tolerance indicates the magnitude of the possible error that is associated with MeasuredValue.
// The true value is located in the range (MeasuredValue – Tolerance) to
// (MeasuredValue + Tolerance)
type Tolerance zcl.Zu16

func (v *Tolerance) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Tolerance) Value() zcl.Val     { return v }

func (v Tolerance) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Tolerance) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Tolerance(*nt)
	return br, err
}

func (v Tolerance) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Tolerance) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Tolerance(*a)
	return nil
}

func (v *Tolerance) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Tolerance(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Tolerance) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const UltrasonicOccupiedToUnoccupiedDelayAttr zcl.AttrID = 32

func (UltrasonicOccupiedToUnoccupiedDelay) ID() zcl.AttrID {
	return UltrasonicOccupiedToUnoccupiedDelayAttr
}
func (UltrasonicOccupiedToUnoccupiedDelay) Readable() bool   { return true }
func (UltrasonicOccupiedToUnoccupiedDelay) Writable() bool   { return true }
func (UltrasonicOccupiedToUnoccupiedDelay) Reportable() bool { return false }
func (UltrasonicOccupiedToUnoccupiedDelay) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UltrasonicOccupiedToUnoccupiedDelay) AttrID() zcl.AttrID   { return v.ID() }
func (v UltrasonicOccupiedToUnoccupiedDelay) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UltrasonicOccupiedToUnoccupiedDelay) AttrValue() zcl.Val  { return v.Value() }

func (UltrasonicOccupiedToUnoccupiedDelay) Name() string {
	return `Ultrasonic Occupied To Unoccupied Delay`
}
func (UltrasonicOccupiedToUnoccupiedDelay) Description() string {
	return `is 16 bits in length and specifies the time delay, in seconds, before the Ultrasonic sensor
changes to its unoccupied state after the last detection of movement in the sensed area`
}

// UltrasonicOccupiedToUnoccupiedDelay is 16 bits in length and specifies the time delay, in seconds, before the Ultrasonic sensor
// changes to its unoccupied state after the last detection of movement in the sensed area
type UltrasonicOccupiedToUnoccupiedDelay zcl.Zu16

func (v *UltrasonicOccupiedToUnoccupiedDelay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *UltrasonicOccupiedToUnoccupiedDelay) Value() zcl.Val     { return v }

func (v UltrasonicOccupiedToUnoccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *UltrasonicOccupiedToUnoccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = UltrasonicOccupiedToUnoccupiedDelay(*nt)
	return br, err
}

func (v UltrasonicOccupiedToUnoccupiedDelay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *UltrasonicOccupiedToUnoccupiedDelay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UltrasonicOccupiedToUnoccupiedDelay(*a)
	return nil
}

func (v *UltrasonicOccupiedToUnoccupiedDelay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = UltrasonicOccupiedToUnoccupiedDelay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UltrasonicOccupiedToUnoccupiedDelay) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const UltrasonicUnoccupiedToOccupiedDelayAttr zcl.AttrID = 33

func (UltrasonicUnoccupiedToOccupiedDelay) ID() zcl.AttrID {
	return UltrasonicUnoccupiedToOccupiedDelayAttr
}
func (UltrasonicUnoccupiedToOccupiedDelay) Readable() bool   { return true }
func (UltrasonicUnoccupiedToOccupiedDelay) Writable() bool   { return true }
func (UltrasonicUnoccupiedToOccupiedDelay) Reportable() bool { return false }
func (UltrasonicUnoccupiedToOccupiedDelay) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UltrasonicUnoccupiedToOccupiedDelay) AttrID() zcl.AttrID   { return v.ID() }
func (v UltrasonicUnoccupiedToOccupiedDelay) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UltrasonicUnoccupiedToOccupiedDelay) AttrValue() zcl.Val  { return v.Value() }

func (UltrasonicUnoccupiedToOccupiedDelay) Name() string {
	return `Ultrasonic Unoccupied To Occupied Delay`
}
func (UltrasonicUnoccupiedToOccupiedDelay) Description() string {
	return `is 16 bits in length and specifies the time delay, in seconds, before the Ultrasonic sensor changes
to its occupied state after the detection of movement in the sensed area. This attribute is
mandatory if the UltrasonicUnoccupiedToOccupiedThreshold attribute is implemented`
}

// UltrasonicUnoccupiedToOccupiedDelay is 16 bits in length and specifies the time delay, in seconds, before the Ultrasonic sensor changes
// to its occupied state after the detection of movement in the sensed area. This attribute is
// mandatory if the UltrasonicUnoccupiedToOccupiedThreshold attribute is implemented
type UltrasonicUnoccupiedToOccupiedDelay zcl.Zu16

func (v *UltrasonicUnoccupiedToOccupiedDelay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *UltrasonicUnoccupiedToOccupiedDelay) Value() zcl.Val     { return v }

func (v UltrasonicUnoccupiedToOccupiedDelay) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *UltrasonicUnoccupiedToOccupiedDelay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = UltrasonicUnoccupiedToOccupiedDelay(*nt)
	return br, err
}

func (v UltrasonicUnoccupiedToOccupiedDelay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *UltrasonicUnoccupiedToOccupiedDelay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UltrasonicUnoccupiedToOccupiedDelay(*a)
	return nil
}

func (v *UltrasonicUnoccupiedToOccupiedDelay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = UltrasonicUnoccupiedToOccupiedDelay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UltrasonicUnoccupiedToOccupiedDelay) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const UltrasonicUnoccupiedToOccupiedThresholdAttr zcl.AttrID = 34

func (UltrasonicUnoccupiedToOccupiedThreshold) ID() zcl.AttrID {
	return UltrasonicUnoccupiedToOccupiedThresholdAttr
}
func (UltrasonicUnoccupiedToOccupiedThreshold) Readable() bool   { return true }
func (UltrasonicUnoccupiedToOccupiedThreshold) Writable() bool   { return true }
func (UltrasonicUnoccupiedToOccupiedThreshold) Reportable() bool { return false }
func (UltrasonicUnoccupiedToOccupiedThreshold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UltrasonicUnoccupiedToOccupiedThreshold) AttrID() zcl.AttrID   { return v.ID() }
func (v UltrasonicUnoccupiedToOccupiedThreshold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UltrasonicUnoccupiedToOccupiedThreshold) AttrValue() zcl.Val  { return v.Value() }

func (UltrasonicUnoccupiedToOccupiedThreshold) Name() string {
	return `Ultrasonic Unoccupied To Occupied Threshold`
}
func (UltrasonicUnoccupiedToOccupiedThreshold) Description() string {
	return `is 8 bits in length and specifies the number of movement detection events that must occur in the period UltrasonicUnoccupiedToOccupiedDelay, before the Ultrasonic sensor changes to its occupied state. This attribute is mandatory if the UltrasonicUnoccupiedToOccupiedDelay attribute is implemented`
}

// UltrasonicUnoccupiedToOccupiedThreshold is 8 bits in length and specifies the number of movement detection events that must occur in the period UltrasonicUnoccupiedToOccupiedDelay, before the Ultrasonic sensor changes to its occupied state. This attribute is mandatory if the UltrasonicUnoccupiedToOccupiedDelay attribute is implemented
type UltrasonicUnoccupiedToOccupiedThreshold zcl.Zu8

func (v *UltrasonicUnoccupiedToOccupiedThreshold) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *UltrasonicUnoccupiedToOccupiedThreshold) Value() zcl.Val     { return v }

func (v UltrasonicUnoccupiedToOccupiedThreshold) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(v).MarshalZcl()
}

func (v *UltrasonicUnoccupiedToOccupiedThreshold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = UltrasonicUnoccupiedToOccupiedThreshold(*nt)
	return br, err
}

func (v UltrasonicUnoccupiedToOccupiedThreshold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *UltrasonicUnoccupiedToOccupiedThreshold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UltrasonicUnoccupiedToOccupiedThreshold(*a)
	return nil
}

func (v *UltrasonicUnoccupiedToOccupiedThreshold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = UltrasonicUnoccupiedToOccupiedThreshold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UltrasonicUnoccupiedToOccupiedThreshold) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}
