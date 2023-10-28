// Meant to controll HVAC appliances, but may be used for other applications
package hvac

import "hemtjan.st/zcl"

type CommandID = zcl.CommandID
type Frame = zcl.ReceivedZclFrame

const AcCapacityAttr zcl.AttrID = 65

func (AcCapacity) ID() zcl.AttrID   { return AcCapacityAttr }
func (AcCapacity) Readable() bool   { return true }
func (AcCapacity) Writable() bool   { return true }
func (AcCapacity) Reportable() bool { return false }
func (AcCapacity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcCapacity) AttrID() zcl.AttrID   { return v.ID() }
func (v AcCapacity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcCapacity) AttrValue() zcl.Val  { return v.Value() }

func (AcCapacity) Name() string { return `AC Capacity` }
func (AcCapacity) Description() string {
	return `capacity in terms of the format defined by the ACCapacityFormat`
}

// AcCapacity capacity in terms of the format defined by the ACCapacityFormat
type AcCapacity zcl.Zu16

func (v *AcCapacity) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *AcCapacity) Value() zcl.Val     { return v }

func (v AcCapacity) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *AcCapacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcCapacity(*nt)
	return br, err
}

func (v AcCapacity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *AcCapacity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcCapacity(*a)
	return nil
}

func (v *AcCapacity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = AcCapacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcCapacity) String() string {
	return zcl.BTUsPerHour.Format(float64(v))
}

const AcCapacityFormatAttr zcl.AttrID = 71

func (AcCapacityFormat) ID() zcl.AttrID   { return AcCapacityFormatAttr }
func (AcCapacityFormat) Readable() bool   { return true }
func (AcCapacityFormat) Writable() bool   { return true }
func (AcCapacityFormat) Reportable() bool { return false }
func (AcCapacityFormat) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcCapacityFormat) AttrID() zcl.AttrID   { return v.ID() }
func (v AcCapacityFormat) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcCapacityFormat) AttrValue() zcl.Val  { return v.Value() }

func (AcCapacityFormat) Name() string        { return `AC Capacity Format` }
func (AcCapacityFormat) Description() string { return `` }

type AcCapacityFormat zcl.Zenum8

func (v *AcCapacityFormat) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *AcCapacityFormat) Value() zcl.Val     { return v }

func (v AcCapacityFormat) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *AcCapacityFormat) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = AcCapacityFormat(*nt)
	return br, err
}

func (v AcCapacityFormat) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *AcCapacityFormat) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcCapacityFormat(*a)
	return nil
}

func (v *AcCapacityFormat) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = AcCapacityFormat(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcCapacityFormat) String() string {
	switch v {
	case 0x00:
		return "BTUh"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v AcCapacityFormat) IsBtuh() bool { return v == 0x00 }
func (v *AcCapacityFormat) SetBtuh()    { *v = 0x00 }

func (AcCapacityFormat) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "BTUh"},
	}
}

const AcCoilTemperatureAttr zcl.AttrID = 70

func (AcCoilTemperature) ID() zcl.AttrID   { return AcCoilTemperatureAttr }
func (AcCoilTemperature) Readable() bool   { return true }
func (AcCoilTemperature) Writable() bool   { return false }
func (AcCoilTemperature) Reportable() bool { return false }
func (AcCoilTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcCoilTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v AcCoilTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcCoilTemperature) AttrValue() zcl.Val  { return v.Value() }

func (AcCoilTemperature) Name() string        { return `AC Coil Temperature` }
func (AcCoilTemperature) Description() string { return `` }

type AcCoilTemperature zcl.Zs16

func (v *AcCoilTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AcCoilTemperature) Value() zcl.Val     { return v }

func (v AcCoilTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AcCoilTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AcCoilTemperature(*nt)
	return br, err
}

func (v AcCoilTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AcCoilTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcCoilTemperature(*a)
	return nil
}

func (v *AcCoilTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AcCoilTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcCoilTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v AcCoilTemperature) Scaled() float64 {
	return float64(v) / 100
}

const AcCompressorTypeAttr zcl.AttrID = 67

func (AcCompressorType) ID() zcl.AttrID   { return AcCompressorTypeAttr }
func (AcCompressorType) Readable() bool   { return true }
func (AcCompressorType) Writable() bool   { return true }
func (AcCompressorType) Reportable() bool { return false }
func (AcCompressorType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcCompressorType) AttrID() zcl.AttrID   { return v.ID() }
func (v AcCompressorType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcCompressorType) AttrValue() zcl.Val  { return v.Value() }

func (AcCompressorType) Name() string        { return `AC Compressor Type` }
func (AcCompressorType) Description() string { return `` }

type AcCompressorType zcl.Zenum8

func (v *AcCompressorType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *AcCompressorType) Value() zcl.Val     { return v }

func (v AcCompressorType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *AcCompressorType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = AcCompressorType(*nt)
	return br, err
}

func (v AcCompressorType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *AcCompressorType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcCompressorType(*a)
	return nil
}

func (v *AcCompressorType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = AcCompressorType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcCompressorType) String() string {
	switch v {
	case 0x00:
		return "Unknown"
	case 0x01:
		return "T1, Max working ambient 43ºC"
	case 0x02:
		return "T2, Max working ambient 35ºC"
	case 0x03:
		return "T3, Max working ambient 52ºC"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v AcCompressorType) IsUnknown() bool                { return v == 0x00 }
func (v AcCompressorType) IsT1MaxWorkingAmbient43C() bool { return v == 0x01 }
func (v AcCompressorType) IsT2MaxWorkingAmbient35C() bool { return v == 0x02 }
func (v AcCompressorType) IsT3MaxWorkingAmbient52C() bool { return v == 0x03 }
func (v *AcCompressorType) SetUnknown()                   { *v = 0x00 }
func (v *AcCompressorType) SetT1MaxWorkingAmbient43C()    { *v = 0x01 }
func (v *AcCompressorType) SetT2MaxWorkingAmbient35C()    { *v = 0x02 }
func (v *AcCompressorType) SetT3MaxWorkingAmbient52C()    { *v = 0x03 }

func (AcCompressorType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Unknown"},
		{Value: 0x01, Name: "T1, Max working ambient 43ºC"},
		{Value: 0x02, Name: "T2, Max working ambient 35ºC"},
		{Value: 0x03, Name: "T3, Max working ambient 52ºC"},
	}
}

const AcErrorCodeAttr zcl.AttrID = 68

func (AcErrorCode) ID() zcl.AttrID   { return AcErrorCodeAttr }
func (AcErrorCode) Readable() bool   { return true }
func (AcErrorCode) Writable() bool   { return true }
func (AcErrorCode) Reportable() bool { return false }
func (AcErrorCode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcErrorCode) AttrID() zcl.AttrID   { return v.ID() }
func (v AcErrorCode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcErrorCode) AttrValue() zcl.Val  { return v.Value() }

func (AcErrorCode) Name() string        { return `AC Error Code` }
func (AcErrorCode) Description() string { return `` }

type AcErrorCode zcl.Zbmp32

func (v *AcErrorCode) TypeID() zcl.TypeID { return new(zcl.Zbmp32).TypeID() }
func (v *AcErrorCode) Value() zcl.Val     { return v }

func (v AcErrorCode) MarshalZcl() ([]byte, error) { return zcl.Zbmp32(v).MarshalZcl() }

func (v *AcErrorCode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp32)
	br, err := nt.UnmarshalZcl(b)
	*v = AcErrorCode(*nt)
	return br, err
}

func (v AcErrorCode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp32(v))
}

func (v *AcErrorCode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcErrorCode(*a)
	return nil
}

func (v *AcErrorCode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp32); ok {
		*v = AcErrorCode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcErrorCode) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Compressor failure or refrigerant leakage")
		case 1:
			bstr = append(bstr, "Room temperature sensor failure")
		case 2:
			bstr = append(bstr, "Outdoor temperature sensor failure")
		case 3:
			bstr = append(bstr, "Indoor coil temperature sensor failure")
		case 4:
			bstr = append(bstr, "Fan failure")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v AcErrorCode) IsCompressorFailureOrRefrigerantLeakage() bool {
	return zcl.BitmapTest([]byte(v[:]), 0)
}
func (v AcErrorCode) IsRoomTemperatureSensorFailure() bool    { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v AcErrorCode) IsOutdoorTemperatureSensorFailure() bool { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v AcErrorCode) IsIndoorCoilTemperatureSensorFailure() bool {
	return zcl.BitmapTest([]byte(v[:]), 3)
}
func (v AcErrorCode) IsFanFailure() bool { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v *AcErrorCode) SetCompressorFailureOrRefrigerantLeakage(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *AcErrorCode) SetRoomTemperatureSensorFailure(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *AcErrorCode) SetOutdoorTemperatureSensorFailure(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *AcErrorCode) SetIndoorCoilTemperatureSensorFailure(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *AcErrorCode) SetFanFailure(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b)) }

func (AcErrorCode) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Compressor failure or refrigerant leakage"},
		{Value: 1, Name: "Room temperature sensor failure"},
		{Value: 2, Name: "Outdoor temperature sensor failure"},
		{Value: 3, Name: "Indoor coil temperature sensor failure"},
		{Value: 4, Name: "Fan failure"},
	}
}

const AcLouverPositionAttr zcl.AttrID = 69

func (AcLouverPosition) ID() zcl.AttrID   { return AcLouverPositionAttr }
func (AcLouverPosition) Readable() bool   { return true }
func (AcLouverPosition) Writable() bool   { return true }
func (AcLouverPosition) Reportable() bool { return false }
func (AcLouverPosition) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcLouverPosition) AttrID() zcl.AttrID   { return v.ID() }
func (v AcLouverPosition) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcLouverPosition) AttrValue() zcl.Val  { return v.Value() }

func (AcLouverPosition) Name() string        { return `AC Louver Position` }
func (AcLouverPosition) Description() string { return `` }

type AcLouverPosition zcl.Zenum8

func (v *AcLouverPosition) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *AcLouverPosition) Value() zcl.Val     { return v }

func (v AcLouverPosition) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *AcLouverPosition) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = AcLouverPosition(*nt)
	return br, err
}

func (v AcLouverPosition) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *AcLouverPosition) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcLouverPosition(*a)
	return nil
}

func (v *AcLouverPosition) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = AcLouverPosition(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcLouverPosition) String() string {
	switch v {
	case 0x01:
		return "Fully closed"
	case 0x02:
		return "Fully open"
	case 0x03:
		return "Quarter open"
	case 0x04:
		return "Half open"
	case 0x05:
		return "Three quarters open"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v AcLouverPosition) IsFullyClosed() bool       { return v == 0x01 }
func (v AcLouverPosition) IsFullyOpen() bool         { return v == 0x02 }
func (v AcLouverPosition) IsQuarterOpen() bool       { return v == 0x03 }
func (v AcLouverPosition) IsHalfOpen() bool          { return v == 0x04 }
func (v AcLouverPosition) IsThreeQuartersOpen() bool { return v == 0x05 }
func (v *AcLouverPosition) SetFullyClosed()          { *v = 0x01 }
func (v *AcLouverPosition) SetFullyOpen()            { *v = 0x02 }
func (v *AcLouverPosition) SetQuarterOpen()          { *v = 0x03 }
func (v *AcLouverPosition) SetHalfOpen()             { *v = 0x04 }
func (v *AcLouverPosition) SetThreeQuartersOpen()    { *v = 0x05 }

func (AcLouverPosition) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x01, Name: "Fully closed"},
		{Value: 0x02, Name: "Fully open"},
		{Value: 0x03, Name: "Quarter open"},
		{Value: 0x04, Name: "Half open"},
		{Value: 0x05, Name: "Three quarters open"},
	}
}

const AcRefrigerantTypeAttr zcl.AttrID = 66

func (AcRefrigerantType) ID() zcl.AttrID   { return AcRefrigerantTypeAttr }
func (AcRefrigerantType) Readable() bool   { return true }
func (AcRefrigerantType) Writable() bool   { return true }
func (AcRefrigerantType) Reportable() bool { return false }
func (AcRefrigerantType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcRefrigerantType) AttrID() zcl.AttrID   { return v.ID() }
func (v AcRefrigerantType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcRefrigerantType) AttrValue() zcl.Val  { return v.Value() }

func (AcRefrigerantType) Name() string        { return `AC Refrigerant Type` }
func (AcRefrigerantType) Description() string { return `refrigerant used` }

// AcRefrigerantType refrigerant used
type AcRefrigerantType zcl.Zenum8

func (v *AcRefrigerantType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *AcRefrigerantType) Value() zcl.Val     { return v }

func (v AcRefrigerantType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *AcRefrigerantType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = AcRefrigerantType(*nt)
	return br, err
}

func (v AcRefrigerantType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *AcRefrigerantType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcRefrigerantType(*a)
	return nil
}

func (v *AcRefrigerantType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = AcRefrigerantType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcRefrigerantType) String() string {
	switch v {
	case 0x00:
		return "Unknown"
	case 0x01:
		return "R22"
	case 0x02:
		return "R410a"
	case 0x03:
		return "R407c"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v AcRefrigerantType) IsUnknown() bool { return v == 0x00 }
func (v AcRefrigerantType) IsR22() bool     { return v == 0x01 }
func (v AcRefrigerantType) IsR410a() bool   { return v == 0x02 }
func (v AcRefrigerantType) IsR407c() bool   { return v == 0x03 }
func (v *AcRefrigerantType) SetUnknown()    { *v = 0x00 }
func (v *AcRefrigerantType) SetR22()        { *v = 0x01 }
func (v *AcRefrigerantType) SetR410a()      { *v = 0x02 }
func (v *AcRefrigerantType) SetR407c()      { *v = 0x03 }

func (AcRefrigerantType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Unknown"},
		{Value: 0x01, Name: "R22"},
		{Value: 0x02, Name: "R410a"},
		{Value: 0x03, Name: "R407c"},
	}
}

const AcTypeAttr zcl.AttrID = 64

func (AcType) ID() zcl.AttrID   { return AcTypeAttr }
func (AcType) Readable() bool   { return true }
func (AcType) Writable() bool   { return true }
func (AcType) Reportable() bool { return false }
func (AcType) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AcType) AttrID() zcl.AttrID   { return v.ID() }
func (v AcType) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AcType) AttrValue() zcl.Val  { return v.Value() }

func (AcType) Name() string { return `AC Type` }
func (AcType) Description() string {
	return `type of Mini Split depending on how Cooling and Heating condition is
achieved`
}

// AcType type of Mini Split depending on how Cooling and Heating condition is
// achieved
type AcType zcl.Zenum8

func (v *AcType) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *AcType) Value() zcl.Val     { return v }

func (v AcType) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *AcType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = AcType(*nt)
	return br, err
}

func (v AcType) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *AcType) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AcType(*a)
	return nil
}

func (v *AcType) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = AcType(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AcType) String() string {
	switch v {
	case 0x00:
		return "Unknown"
	case 0x01:
		return "Cooling and fixed speed"
	case 0x02:
		return "Heat pump and fixed speed"
	case 0x03:
		return "Cooling and inverter"
	case 0x04:
		return "Heat pump and inverter"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v AcType) IsUnknown() bool               { return v == 0x00 }
func (v AcType) IsCoolingAndFixedSpeed() bool  { return v == 0x01 }
func (v AcType) IsHeatPumpAndFixedSpeed() bool { return v == 0x02 }
func (v AcType) IsCoolingAndInverter() bool    { return v == 0x03 }
func (v AcType) IsHeatPumpAndInverter() bool   { return v == 0x04 }
func (v *AcType) SetUnknown()                  { *v = 0x00 }
func (v *AcType) SetCoolingAndFixedSpeed()     { *v = 0x01 }
func (v *AcType) SetHeatPumpAndFixedSpeed()    { *v = 0x02 }
func (v *AcType) SetCoolingAndInverter()       { *v = 0x03 }
func (v *AcType) SetHeatPumpAndInverter()      { *v = 0x04 }

func (AcType) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Unknown"},
		{Value: 0x01, Name: "Cooling and fixed speed"},
		{Value: 0x02, Name: "Heat pump and fixed speed"},
		{Value: 0x03, Name: "Cooling and inverter"},
		{Value: 0x04, Name: "Heat pump and inverter"},
	}
}

const AbsMaxCoolSetpointLimitAttr zcl.AttrID = 6

func (AbsMaxCoolSetpointLimit) ID() zcl.AttrID   { return AbsMaxCoolSetpointLimitAttr }
func (AbsMaxCoolSetpointLimit) Readable() bool   { return true }
func (AbsMaxCoolSetpointLimit) Writable() bool   { return false }
func (AbsMaxCoolSetpointLimit) Reportable() bool { return false }
func (AbsMaxCoolSetpointLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AbsMaxCoolSetpointLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v AbsMaxCoolSetpointLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AbsMaxCoolSetpointLimit) AttrValue() zcl.Val  { return v.Value() }

func (AbsMaxCoolSetpointLimit) Name() string { return `Abs Max Cool Setpoint Limit` }
func (AbsMaxCoolSetpointLimit) Description() string {
	return `absolute maximum level that the cooling setpoint may be set to. This is
a limitation imposed by the manufacturer. The value is calculated as
described in the LocalTemperature attribute`
}

// AbsMaxCoolSetpointLimit absolute maximum level that the cooling setpoint may be set to. This is
// a limitation imposed by the manufacturer. The value is calculated as
// described in the LocalTemperature attribute
type AbsMaxCoolSetpointLimit zcl.Zs16

func (v *AbsMaxCoolSetpointLimit) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AbsMaxCoolSetpointLimit) Value() zcl.Val     { return v }

func (v AbsMaxCoolSetpointLimit) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AbsMaxCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AbsMaxCoolSetpointLimit(*nt)
	return br, err
}

func (v AbsMaxCoolSetpointLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AbsMaxCoolSetpointLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AbsMaxCoolSetpointLimit(*a)
	return nil
}

func (v *AbsMaxCoolSetpointLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AbsMaxCoolSetpointLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AbsMaxCoolSetpointLimit) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v AbsMaxCoolSetpointLimit) Scaled() float64 {
	return float64(v) / 100
}

const AbsMaxHeatSetpointLimitAttr zcl.AttrID = 4

func (AbsMaxHeatSetpointLimit) ID() zcl.AttrID   { return AbsMaxHeatSetpointLimitAttr }
func (AbsMaxHeatSetpointLimit) Readable() bool   { return true }
func (AbsMaxHeatSetpointLimit) Writable() bool   { return false }
func (AbsMaxHeatSetpointLimit) Reportable() bool { return false }
func (AbsMaxHeatSetpointLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AbsMaxHeatSetpointLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v AbsMaxHeatSetpointLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AbsMaxHeatSetpointLimit) AttrValue() zcl.Val  { return v.Value() }

func (AbsMaxHeatSetpointLimit) Name() string { return `Abs Max Heat Setpoint Limit` }
func (AbsMaxHeatSetpointLimit) Description() string {
	return `absolute maximum level that the heating setpoint may be set to. This is[[s]]
a limitation imposed by the manufacturer. The value is calculated as
described in the LocalTemperature attribute`
}

// AbsMaxHeatSetpointLimit absolute maximum level that the heating setpoint may be set to. This is[[s]]
// a limitation imposed by the manufacturer. The value is calculated as
// described in the LocalTemperature attribute
type AbsMaxHeatSetpointLimit zcl.Zs16

func (v *AbsMaxHeatSetpointLimit) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AbsMaxHeatSetpointLimit) Value() zcl.Val     { return v }

func (v AbsMaxHeatSetpointLimit) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AbsMaxHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AbsMaxHeatSetpointLimit(*nt)
	return br, err
}

func (v AbsMaxHeatSetpointLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AbsMaxHeatSetpointLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AbsMaxHeatSetpointLimit(*a)
	return nil
}

func (v *AbsMaxHeatSetpointLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AbsMaxHeatSetpointLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AbsMaxHeatSetpointLimit) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v AbsMaxHeatSetpointLimit) Scaled() float64 {
	return float64(v) / 100
}

const AbsMinCoolSetpointLimitAttr zcl.AttrID = 5

func (AbsMinCoolSetpointLimit) ID() zcl.AttrID   { return AbsMinCoolSetpointLimitAttr }
func (AbsMinCoolSetpointLimit) Readable() bool   { return true }
func (AbsMinCoolSetpointLimit) Writable() bool   { return false }
func (AbsMinCoolSetpointLimit) Reportable() bool { return false }
func (AbsMinCoolSetpointLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AbsMinCoolSetpointLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v AbsMinCoolSetpointLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AbsMinCoolSetpointLimit) AttrValue() zcl.Val  { return v.Value() }

func (AbsMinCoolSetpointLimit) Name() string { return `Abs Min Cool Setpoint Limit` }
func (AbsMinCoolSetpointLimit) Description() string {
	return `absolute minimum level that the cooling setpoint may be set to. This is
a limitation imposed by the manufacturer. The value is calculated as
described in the LocalTemperature attribute`
}

// AbsMinCoolSetpointLimit absolute minimum level that the cooling setpoint may be set to. This is
// a limitation imposed by the manufacturer. The value is calculated as
// described in the LocalTemperature attribute
type AbsMinCoolSetpointLimit zcl.Zs16

func (v *AbsMinCoolSetpointLimit) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AbsMinCoolSetpointLimit) Value() zcl.Val     { return v }

func (v AbsMinCoolSetpointLimit) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AbsMinCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AbsMinCoolSetpointLimit(*nt)
	return br, err
}

func (v AbsMinCoolSetpointLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AbsMinCoolSetpointLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AbsMinCoolSetpointLimit(*a)
	return nil
}

func (v *AbsMinCoolSetpointLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AbsMinCoolSetpointLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AbsMinCoolSetpointLimit) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v AbsMinCoolSetpointLimit) Scaled() float64 {
	return float64(v) / 100
}

const AbsMinHeatSetpointLimitAttr zcl.AttrID = 3

func (AbsMinHeatSetpointLimit) ID() zcl.AttrID   { return AbsMinHeatSetpointLimitAttr }
func (AbsMinHeatSetpointLimit) Readable() bool   { return true }
func (AbsMinHeatSetpointLimit) Writable() bool   { return false }
func (AbsMinHeatSetpointLimit) Reportable() bool { return false }
func (AbsMinHeatSetpointLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v AbsMinHeatSetpointLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v AbsMinHeatSetpointLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *AbsMinHeatSetpointLimit) AttrValue() zcl.Val  { return v.Value() }

func (AbsMinHeatSetpointLimit) Name() string { return `Abs Min Heat Setpoint Limit` }
func (AbsMinHeatSetpointLimit) Description() string {
	return `absolute minimum level that the heating setpoint may be set to. This is
a limitation imposed by the manufacturer. The value is calculated as
described in the LocalTemperature attribute`
}

// AbsMinHeatSetpointLimit absolute minimum level that the heating setpoint may be set to. This is
// a limitation imposed by the manufacturer. The value is calculated as
// described in the LocalTemperature attribute
type AbsMinHeatSetpointLimit zcl.Zs16

func (v *AbsMinHeatSetpointLimit) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *AbsMinHeatSetpointLimit) Value() zcl.Val     { return v }

func (v AbsMinHeatSetpointLimit) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *AbsMinHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = AbsMinHeatSetpointLimit(*nt)
	return br, err
}

func (v AbsMinHeatSetpointLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *AbsMinHeatSetpointLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = AbsMinHeatSetpointLimit(*a)
	return nil
}

func (v *AbsMinHeatSetpointLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = AbsMinHeatSetpointLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v AbsMinHeatSetpointLimit) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v AbsMinHeatSetpointLimit) Scaled() float64 {
	return float64(v) / 100
}

const CapacityAttr zcl.AttrID = 19

func (Capacity) ID() zcl.AttrID   { return CapacityAttr }
func (Capacity) Readable() bool   { return true }
func (Capacity) Writable() bool   { return false }
func (Capacity) Reportable() bool { return true }
func (Capacity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Capacity) AttrID() zcl.AttrID   { return v.ID() }
func (v Capacity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Capacity) AttrValue() zcl.Val  { return v.Value() }

func (Capacity) Name() string { return `Capacity` }
func (Capacity) Description() string {
	return `actual capacity of the pump as a percentage of the effective maximum
setpoint value. It is updated dynamically as the speed of the pump
changes`
}

// Capacity actual capacity of the pump as a percentage of the effective maximum
// setpoint value. It is updated dynamically as the speed of the pump
// changes
type Capacity zcl.Zs16

func (v *Capacity) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *Capacity) Value() zcl.Val     { return v }

func (v Capacity) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *Capacity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = Capacity(*nt)
	return br, err
}

func (v Capacity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *Capacity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Capacity(*a)
	return nil
}

func (v *Capacity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = Capacity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Capacity) String() string {
	return zcl.Percent.Format(float64(v) / 200)
}

func (v Capacity) Scaled() float64 {
	return float64(v) / 200
}

const ControlModeAttr zcl.AttrID = 33

func (ControlMode) ID() zcl.AttrID   { return ControlModeAttr }
func (ControlMode) Readable() bool   { return true }
func (ControlMode) Writable() bool   { return true }
func (ControlMode) Reportable() bool { return false }
func (ControlMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ControlMode) AttrID() zcl.AttrID   { return v.ID() }
func (v ControlMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ControlMode) AttrValue() zcl.Val  { return v.Value() }

func (ControlMode) Name() string        { return `Control Mode` }
func (ControlMode) Description() string { return `control mode of the pump` }

// ControlMode control mode of the pump
type ControlMode zcl.Zenum8

func (v *ControlMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *ControlMode) Value() zcl.Val     { return v }

func (v ControlMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *ControlMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = ControlMode(*nt)
	return br, err
}

func (v ControlMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *ControlMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ControlMode(*a)
	return nil
}

func (v *ControlMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = ControlMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ControlMode) String() string {
	switch v {
	case 0x00:
		return "Constant speed"
	case 0x01:
		return "Constant pressure"
	case 0x02:
		return "Proportional pressure"
	case 0x03:
		return "Constant flow"
	case 0x05:
		return "Constant temperature"
	case 0x07:
		return "Automatic"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v ControlMode) IsConstantSpeed() bool        { return v == 0x00 }
func (v ControlMode) IsConstantPressure() bool     { return v == 0x01 }
func (v ControlMode) IsProportionalPressure() bool { return v == 0x02 }
func (v ControlMode) IsConstantFlow() bool         { return v == 0x03 }
func (v ControlMode) IsConstantTemperature() bool  { return v == 0x05 }
func (v ControlMode) IsAutomatic() bool            { return v == 0x07 }
func (v *ControlMode) SetConstantSpeed()           { *v = 0x00 }
func (v *ControlMode) SetConstantPressure()        { *v = 0x01 }
func (v *ControlMode) SetProportionalPressure()    { *v = 0x02 }
func (v *ControlMode) SetConstantFlow()            { *v = 0x03 }
func (v *ControlMode) SetConstantTemperature()     { *v = 0x05 }
func (v *ControlMode) SetAutomatic()               { *v = 0x07 }

func (ControlMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Constant speed"},
		{Value: 0x01, Name: "Constant pressure"},
		{Value: 0x02, Name: "Proportional pressure"},
		{Value: 0x03, Name: "Constant flow"},
		{Value: 0x05, Name: "Constant temperature"},
		{Value: 0x07, Name: "Automatic"},
	}
}

const ControlSequenceOfOperationAttr zcl.AttrID = 27

func (ControlSequenceOfOperation) ID() zcl.AttrID   { return ControlSequenceOfOperationAttr }
func (ControlSequenceOfOperation) Readable() bool   { return true }
func (ControlSequenceOfOperation) Writable() bool   { return true }
func (ControlSequenceOfOperation) Reportable() bool { return false }
func (ControlSequenceOfOperation) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ControlSequenceOfOperation) AttrID() zcl.AttrID   { return v.ID() }
func (v ControlSequenceOfOperation) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ControlSequenceOfOperation) AttrValue() zcl.Val  { return v.Value() }

func (ControlSequenceOfOperation) Name() string { return `Control Sequence Of Operation` }
func (ControlSequenceOfOperation) Description() string {
	return `overall operating environment of the thermostat, and thus the possible
system modes that the thermostat can operate in`
}

// ControlSequenceOfOperation overall operating environment of the thermostat, and thus the possible
// system modes that the thermostat can operate in
type ControlSequenceOfOperation zcl.Zenum8

func (v *ControlSequenceOfOperation) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *ControlSequenceOfOperation) Value() zcl.Val     { return v }

func (v ControlSequenceOfOperation) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *ControlSequenceOfOperation) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = ControlSequenceOfOperation(*nt)
	return br, err
}

func (v ControlSequenceOfOperation) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *ControlSequenceOfOperation) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ControlSequenceOfOperation(*a)
	return nil
}

func (v *ControlSequenceOfOperation) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = ControlSequenceOfOperation(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ControlSequenceOfOperation) String() string {
	switch v {
	case 0x00:
		return "Cooling only"
	case 0x01:
		return "Cooling with reheat"
	case 0x02:
		return "Heating only"
	case 0x03:
		return "Heating with reheat"
	case 0x04:
		return "Cooling and heating 4-pipes"
	case 0x05:
		return "Cooling and heating 4-pipes with reheat"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v ControlSequenceOfOperation) IsCoolingOnly() bool                       { return v == 0x00 }
func (v ControlSequenceOfOperation) IsCoolingWithReheat() bool                 { return v == 0x01 }
func (v ControlSequenceOfOperation) IsHeatingOnly() bool                       { return v == 0x02 }
func (v ControlSequenceOfOperation) IsHeatingWithReheat() bool                 { return v == 0x03 }
func (v ControlSequenceOfOperation) IsCoolingAndHeating4Pipes() bool           { return v == 0x04 }
func (v ControlSequenceOfOperation) IsCoolingAndHeating4PipesWithReheat() bool { return v == 0x05 }
func (v *ControlSequenceOfOperation) SetCoolingOnly()                          { *v = 0x00 }
func (v *ControlSequenceOfOperation) SetCoolingWithReheat()                    { *v = 0x01 }
func (v *ControlSequenceOfOperation) SetHeatingOnly()                          { *v = 0x02 }
func (v *ControlSequenceOfOperation) SetHeatingWithReheat()                    { *v = 0x03 }
func (v *ControlSequenceOfOperation) SetCoolingAndHeating4Pipes()              { *v = 0x04 }
func (v *ControlSequenceOfOperation) SetCoolingAndHeating4PipesWithReheat()    { *v = 0x05 }

func (ControlSequenceOfOperation) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Cooling only"},
		{Value: 0x01, Name: "Cooling with reheat"},
		{Value: 0x02, Name: "Heating only"},
		{Value: 0x03, Name: "Heating with reheat"},
		{Value: 0x04, Name: "Cooling and heating 4-pipes"},
		{Value: 0x05, Name: "Cooling and heating 4-pipes with reheat"},
	}
}

const DehumidificationCoolingAttr zcl.AttrID = 1

func (DehumidificationCooling) ID() zcl.AttrID   { return DehumidificationCoolingAttr }
func (DehumidificationCooling) Readable() bool   { return true }
func (DehumidificationCooling) Writable() bool   { return false }
func (DehumidificationCooling) Reportable() bool { return true }
func (DehumidificationCooling) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DehumidificationCooling) AttrID() zcl.AttrID   { return v.ID() }
func (v DehumidificationCooling) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DehumidificationCooling) AttrValue() zcl.Val  { return v.Value() }

func (DehumidificationCooling) Name() string        { return `Dehumidification Cooling` }
func (DehumidificationCooling) Description() string { return `current dehumidification cooling output` }

// DehumidificationCooling current dehumidification cooling output
type DehumidificationCooling zcl.Zu8

func (v *DehumidificationCooling) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *DehumidificationCooling) Value() zcl.Val     { return v }

func (v DehumidificationCooling) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *DehumidificationCooling) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = DehumidificationCooling(*nt)
	return br, err
}

func (v DehumidificationCooling) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *DehumidificationCooling) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DehumidificationCooling(*a)
	return nil
}

func (v *DehumidificationCooling) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = DehumidificationCooling(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DehumidificationCooling) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const DehumidificationHysteresisAttr zcl.AttrID = 19

func (DehumidificationHysteresis) ID() zcl.AttrID   { return DehumidificationHysteresisAttr }
func (DehumidificationHysteresis) Readable() bool   { return true }
func (DehumidificationHysteresis) Writable() bool   { return true }
func (DehumidificationHysteresis) Reportable() bool { return false }
func (DehumidificationHysteresis) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DehumidificationHysteresis) AttrID() zcl.AttrID   { return v.ID() }
func (v DehumidificationHysteresis) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DehumidificationHysteresis) AttrValue() zcl.Val  { return v.Value() }

func (DehumidificationHysteresis) Name() string        { return `Dehumidification Hysteresis` }
func (DehumidificationHysteresis) Description() string { return `hysteresis associated with RH` }

// DehumidificationHysteresis hysteresis associated with RH
type DehumidificationHysteresis zcl.Zu8

func (v *DehumidificationHysteresis) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *DehumidificationHysteresis) Value() zcl.Val     { return v }

func (v DehumidificationHysteresis) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *DehumidificationHysteresis) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = DehumidificationHysteresis(*nt)
	return br, err
}

func (v DehumidificationHysteresis) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *DehumidificationHysteresis) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DehumidificationHysteresis(*a)
	return nil
}

func (v *DehumidificationHysteresis) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = DehumidificationHysteresis(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DehumidificationHysteresis) String() string {
	return zcl.Percent.Format(float64(v))
}

const DehumidificationLockoutAttr zcl.AttrID = 18

func (DehumidificationLockout) ID() zcl.AttrID   { return DehumidificationLockoutAttr }
func (DehumidificationLockout) Readable() bool   { return true }
func (DehumidificationLockout) Writable() bool   { return true }
func (DehumidificationLockout) Reportable() bool { return false }
func (DehumidificationLockout) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DehumidificationLockout) AttrID() zcl.AttrID   { return v.ID() }
func (v DehumidificationLockout) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DehumidificationLockout) AttrValue() zcl.Val  { return v.Value() }

func (DehumidificationLockout) Name() string { return `Dehumidification Lockout` }
func (DehumidificationLockout) Description() string {
	return `whether dehumidification is allowed or not`
}

// DehumidificationLockout whether dehumidification is allowed or not
type DehumidificationLockout zcl.Zenum8

func (v *DehumidificationLockout) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *DehumidificationLockout) Value() zcl.Val     { return v }

func (v DehumidificationLockout) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *DehumidificationLockout) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = DehumidificationLockout(*nt)
	return br, err
}

func (v DehumidificationLockout) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *DehumidificationLockout) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DehumidificationLockout(*a)
	return nil
}

func (v *DehumidificationLockout) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = DehumidificationLockout(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DehumidificationLockout) String() string {
	switch v {
	case 0x00:
		return "Denied"
	case 0x01:
		return "Allowed"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v DehumidificationLockout) IsDenied() bool  { return v == 0x00 }
func (v DehumidificationLockout) IsAllowed() bool { return v == 0x01 }
func (v *DehumidificationLockout) SetDenied()     { *v = 0x00 }
func (v *DehumidificationLockout) SetAllowed()    { *v = 0x01 }

func (DehumidificationLockout) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Denied"},
		{Value: 0x01, Name: "Allowed"},
	}
}

const DehumidificationMaxCoolAttr zcl.AttrID = 20

func (DehumidificationMaxCool) ID() zcl.AttrID   { return DehumidificationMaxCoolAttr }
func (DehumidificationMaxCool) Readable() bool   { return true }
func (DehumidificationMaxCool) Writable() bool   { return true }
func (DehumidificationMaxCool) Reportable() bool { return false }
func (DehumidificationMaxCool) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v DehumidificationMaxCool) AttrID() zcl.AttrID   { return v.ID() }
func (v DehumidificationMaxCool) AttrType() zcl.TypeID { return v.TypeID() }
func (v *DehumidificationMaxCool) AttrValue() zcl.Val  { return v.Value() }

func (DehumidificationMaxCool) Name() string        { return `Dehumidification Max Cool` }
func (DehumidificationMaxCool) Description() string { return `maximum dehumidification cooling output` }

// DehumidificationMaxCool maximum dehumidification cooling output
type DehumidificationMaxCool zcl.Zu8

func (v *DehumidificationMaxCool) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *DehumidificationMaxCool) Value() zcl.Val     { return v }

func (v DehumidificationMaxCool) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *DehumidificationMaxCool) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = DehumidificationMaxCool(*nt)
	return br, err
}

func (v DehumidificationMaxCool) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *DehumidificationMaxCool) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = DehumidificationMaxCool(*a)
	return nil
}

func (v *DehumidificationMaxCool) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = DehumidificationMaxCool(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v DehumidificationMaxCool) String() string {
	return zcl.Percent.Format(float64(v))
}

const EffectiveControlModeAttr zcl.AttrID = 18

func (EffectiveControlMode) ID() zcl.AttrID   { return EffectiveControlModeAttr }
func (EffectiveControlMode) Readable() bool   { return true }
func (EffectiveControlMode) Writable() bool   { return false }
func (EffectiveControlMode) Reportable() bool { return false }
func (EffectiveControlMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v EffectiveControlMode) AttrID() zcl.AttrID   { return v.ID() }
func (v EffectiveControlMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *EffectiveControlMode) AttrValue() zcl.Val  { return v.Value() }

func (EffectiveControlMode) Name() string { return `Effective Control Mode` }
func (EffectiveControlMode) Description() string {
	return `control mode that currently applies to the pump. It will have the value
of the ControlMode attribute, unless a remote sensor is used as the
sensor for regulation of the pump. In this case, EffectiveControlMode
will display Constant pressure, Constant flow or Constant temperature
if the remote sensor is a pressure sensor, a flow sensor or a
temperature sensor respectively, regardless of the value of the
ControlMode attribute`
}

// EffectiveControlMode control mode that currently applies to the pump. It will have the value
// of the ControlMode attribute, unless a remote sensor is used as the
// sensor for regulation of the pump. In this case, EffectiveControlMode
// will display Constant pressure, Constant flow or Constant temperature
// if the remote sensor is a pressure sensor, a flow sensor or a
// temperature sensor respectively, regardless of the value of the
// ControlMode attribute
type EffectiveControlMode zcl.Zenum8

func (v *EffectiveControlMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *EffectiveControlMode) Value() zcl.Val     { return v }

func (v EffectiveControlMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *EffectiveControlMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = EffectiveControlMode(*nt)
	return br, err
}

func (v EffectiveControlMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *EffectiveControlMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = EffectiveControlMode(*a)
	return nil
}

func (v *EffectiveControlMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = EffectiveControlMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EffectiveControlMode) String() string {
	switch v {
	case 0x00:
		return "Constant speed"
	case 0x01:
		return "Constant pressure"
	case 0x02:
		return "Proportional pressure"
	case 0x03:
		return "Constant flow"
	case 0x05:
		return "Constant temperature"
	case 0x07:
		return "Automatic"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v EffectiveControlMode) IsConstantSpeed() bool        { return v == 0x00 }
func (v EffectiveControlMode) IsConstantPressure() bool     { return v == 0x01 }
func (v EffectiveControlMode) IsProportionalPressure() bool { return v == 0x02 }
func (v EffectiveControlMode) IsConstantFlow() bool         { return v == 0x03 }
func (v EffectiveControlMode) IsConstantTemperature() bool  { return v == 0x05 }
func (v EffectiveControlMode) IsAutomatic() bool            { return v == 0x07 }
func (v *EffectiveControlMode) SetConstantSpeed()           { *v = 0x00 }
func (v *EffectiveControlMode) SetConstantPressure()        { *v = 0x01 }
func (v *EffectiveControlMode) SetProportionalPressure()    { *v = 0x02 }
func (v *EffectiveControlMode) SetConstantFlow()            { *v = 0x03 }
func (v *EffectiveControlMode) SetConstantTemperature()     { *v = 0x05 }
func (v *EffectiveControlMode) SetAutomatic()               { *v = 0x07 }

func (EffectiveControlMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Constant speed"},
		{Value: 0x01, Name: "Constant pressure"},
		{Value: 0x02, Name: "Proportional pressure"},
		{Value: 0x03, Name: "Constant flow"},
		{Value: 0x05, Name: "Constant temperature"},
		{Value: 0x07, Name: "Automatic"},
	}
}

const EffectiveOperationModeAttr zcl.AttrID = 17

func (EffectiveOperationMode) ID() zcl.AttrID   { return EffectiveOperationModeAttr }
func (EffectiveOperationMode) Readable() bool   { return true }
func (EffectiveOperationMode) Writable() bool   { return false }
func (EffectiveOperationMode) Reportable() bool { return false }
func (EffectiveOperationMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v EffectiveOperationMode) AttrID() zcl.AttrID   { return v.ID() }
func (v EffectiveOperationMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *EffectiveOperationMode) AttrValue() zcl.Val  { return v.Value() }

func (EffectiveOperationMode) Name() string { return `Effective Operation Mode` }
func (EffectiveOperationMode) Description() string {
	return `current effective operation mode of the pump. The value of the
EffectiveOperationMode attribute is the same as the OperationMode
attribute of the Pump settings attribute set, except when it is
overridden locally`
}

// EffectiveOperationMode current effective operation mode of the pump. The value of the
// EffectiveOperationMode attribute is the same as the OperationMode
// attribute of the Pump settings attribute set, except when it is
// overridden locally
type EffectiveOperationMode zcl.Zenum8

func (v *EffectiveOperationMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *EffectiveOperationMode) Value() zcl.Val     { return v }

func (v EffectiveOperationMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *EffectiveOperationMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = EffectiveOperationMode(*nt)
	return br, err
}

func (v EffectiveOperationMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *EffectiveOperationMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = EffectiveOperationMode(*a)
	return nil
}

func (v *EffectiveOperationMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = EffectiveOperationMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EffectiveOperationMode) String() string {
	switch v {
	case 0x00:
		return "Normal"
	case 0x01:
		return "Minimum"
	case 0x02:
		return "Maximum"
	case 0x03:
		return "Local"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v EffectiveOperationMode) IsNormal() bool  { return v == 0x00 }
func (v EffectiveOperationMode) IsMinimum() bool { return v == 0x01 }
func (v EffectiveOperationMode) IsMaximum() bool { return v == 0x02 }
func (v EffectiveOperationMode) IsLocal() bool   { return v == 0x03 }
func (v *EffectiveOperationMode) SetNormal()     { *v = 0x00 }
func (v *EffectiveOperationMode) SetMinimum()    { *v = 0x01 }
func (v *EffectiveOperationMode) SetMaximum()    { *v = 0x02 }
func (v *EffectiveOperationMode) SetLocal()      { *v = 0x03 }

func (EffectiveOperationMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Normal"},
		{Value: 0x01, Name: "Minimum"},
		{Value: 0x02, Name: "Maximum"},
		{Value: 0x03, Name: "Local"},
	}
}

const EmergencyHeatDeltaAttr zcl.AttrID = 58

func (EmergencyHeatDelta) ID() zcl.AttrID   { return EmergencyHeatDeltaAttr }
func (EmergencyHeatDelta) Readable() bool   { return true }
func (EmergencyHeatDelta) Writable() bool   { return true }
func (EmergencyHeatDelta) Reportable() bool { return false }
func (EmergencyHeatDelta) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v EmergencyHeatDelta) AttrID() zcl.AttrID   { return v.ID() }
func (v EmergencyHeatDelta) AttrType() zcl.TypeID { return v.TypeID() }
func (v *EmergencyHeatDelta) AttrValue() zcl.Val  { return v.Value() }

func (EmergencyHeatDelta) Name() string { return `Emergency Heat Delta` }
func (EmergencyHeatDelta) Description() string {
	return `degrees between LocalTemperature and the OccupiedHeatingSetpoint or
UnoccupiedHeatingSetpoint attributes at which the Thermostat server
will operate in emergency heat mode`
}

// EmergencyHeatDelta degrees between LocalTemperature and the OccupiedHeatingSetpoint or
// UnoccupiedHeatingSetpoint attributes at which the Thermostat server
// will operate in emergency heat mode
type EmergencyHeatDelta zcl.Zu8

func (v *EmergencyHeatDelta) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *EmergencyHeatDelta) Value() zcl.Val     { return v }

func (v EmergencyHeatDelta) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *EmergencyHeatDelta) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = EmergencyHeatDelta(*nt)
	return br, err
}

func (v EmergencyHeatDelta) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *EmergencyHeatDelta) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = EmergencyHeatDelta(*a)
	return nil
}

func (v *EmergencyHeatDelta) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = EmergencyHeatDelta(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v EmergencyHeatDelta) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v EmergencyHeatDelta) Scaled() float64 {
	return float64(v) / 10
}

const FanModeAttr zcl.AttrID = 0

func (FanMode) ID() zcl.AttrID   { return FanModeAttr }
func (FanMode) Readable() bool   { return true }
func (FanMode) Writable() bool   { return true }
func (FanMode) Reportable() bool { return false }
func (FanMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v FanMode) AttrID() zcl.AttrID   { return v.ID() }
func (v FanMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *FanMode) AttrValue() zcl.Val  { return v.Value() }

func (FanMode) Name() string        { return `Fan Mode` }
func (FanMode) Description() string { return `current speed of the fan` }

// FanMode current speed of the fan
type FanMode zcl.Zenum8

func (v *FanMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *FanMode) Value() zcl.Val     { return v }

func (v FanMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *FanMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = FanMode(*nt)
	return br, err
}

func (v FanMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *FanMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = FanMode(*a)
	return nil
}

func (v *FanMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = FanMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v FanMode) String() string {
	switch v {
	case 0x00:
		return "Off"
	case 0x01:
		return "Low"
	case 0x02:
		return "Medium"
	case 0x03:
		return "High"
	case 0x04:
		return "On"
	case 0x05:
		return "Auto"
	case 0x06:
		return "Smart (based on occupancy)"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v FanMode) IsOff() bool                   { return v == 0x00 }
func (v FanMode) IsLow() bool                   { return v == 0x01 }
func (v FanMode) IsMedium() bool                { return v == 0x02 }
func (v FanMode) IsHigh() bool                  { return v == 0x03 }
func (v FanMode) IsOn() bool                    { return v == 0x04 }
func (v FanMode) IsAuto() bool                  { return v == 0x05 }
func (v FanMode) IsSmartBasedOnOccupancy() bool { return v == 0x06 }
func (v *FanMode) SetOff()                      { *v = 0x00 }
func (v *FanMode) SetLow()                      { *v = 0x01 }
func (v *FanMode) SetMedium()                   { *v = 0x02 }
func (v *FanMode) SetHigh()                     { *v = 0x03 }
func (v *FanMode) SetOn()                       { *v = 0x04 }
func (v *FanMode) SetAuto()                     { *v = 0x05 }
func (v *FanMode) SetSmartBasedOnOccupancy()    { *v = 0x06 }

func (FanMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Off"},
		{Value: 0x01, Name: "Low"},
		{Value: 0x02, Name: "Medium"},
		{Value: 0x03, Name: "High"},
		{Value: 0x04, Name: "On"},
		{Value: 0x05, Name: "Auto"},
		{Value: 0x06, Name: "Smart (based on occupancy)"},
	}
}

const FanModeSequenceAttr zcl.AttrID = 1

func (FanModeSequence) ID() zcl.AttrID   { return FanModeSequenceAttr }
func (FanModeSequence) Readable() bool   { return true }
func (FanModeSequence) Writable() bool   { return true }
func (FanModeSequence) Reportable() bool { return false }
func (FanModeSequence) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v FanModeSequence) AttrID() zcl.AttrID   { return v.ID() }
func (v FanModeSequence) AttrType() zcl.TypeID { return v.TypeID() }
func (v *FanModeSequence) AttrValue() zcl.Val  { return v.Value() }

func (FanModeSequence) Name() string        { return `Fan Mode Sequence` }
func (FanModeSequence) Description() string { return `possible fan speeds that the thermostat can set` }

// FanModeSequence possible fan speeds that the thermostat can set
type FanModeSequence zcl.Zenum8

func (v *FanModeSequence) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *FanModeSequence) Value() zcl.Val     { return v }

func (v FanModeSequence) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *FanModeSequence) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = FanModeSequence(*nt)
	return br, err
}

func (v FanModeSequence) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *FanModeSequence) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = FanModeSequence(*a)
	return nil
}

func (v *FanModeSequence) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = FanModeSequence(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v FanModeSequence) String() string {
	switch v {
	case 0x00:
		return "Low/med/high"
	case 0x01:
		return "Low/high"
	case 0x02:
		return "Low/med/high/auto"
	case 0x03:
		return "Low/high/auto"
	case 0x04:
		return "On/auto"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v FanModeSequence) IsLowMedHigh() bool     { return v == 0x00 }
func (v FanModeSequence) IsLowHigh() bool        { return v == 0x01 }
func (v FanModeSequence) IsLowMedHighAuto() bool { return v == 0x02 }
func (v FanModeSequence) IsLowHighAuto() bool    { return v == 0x03 }
func (v FanModeSequence) IsOnAuto() bool         { return v == 0x04 }
func (v *FanModeSequence) SetLowMedHigh()        { *v = 0x00 }
func (v *FanModeSequence) SetLowHigh()           { *v = 0x01 }
func (v *FanModeSequence) SetLowMedHighAuto()    { *v = 0x02 }
func (v *FanModeSequence) SetLowHighAuto()       { *v = 0x03 }
func (v *FanModeSequence) SetOnAuto()            { *v = 0x04 }

func (FanModeSequence) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Low/med/high"},
		{Value: 0x01, Name: "Low/high"},
		{Value: 0x02, Name: "Low/med/high/auto"},
		{Value: 0x03, Name: "Low/high/auto"},
		{Value: 0x04, Name: "On/auto"},
	}
}

func (GetWeeklyDaysToReturn) Name() string        { return `Get Weekly Days To Return` }
func (GetWeeklyDaysToReturn) Description() string { return `` }

type GetWeeklyDaysToReturn zcl.Zbmp8

func (v *GetWeeklyDaysToReturn) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *GetWeeklyDaysToReturn) Value() zcl.Val     { return v }

func (v GetWeeklyDaysToReturn) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *GetWeeklyDaysToReturn) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = GetWeeklyDaysToReturn(*nt)
	return br, err
}

func (v GetWeeklyDaysToReturn) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *GetWeeklyDaysToReturn) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = GetWeeklyDaysToReturn(*a)
	return nil
}

func (v *GetWeeklyDaysToReturn) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = GetWeeklyDaysToReturn(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v GetWeeklyDaysToReturn) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Sunday")
		case 1:
			bstr = append(bstr, "Monday")
		case 2:
			bstr = append(bstr, "Tuesday")
		case 3:
			bstr = append(bstr, "Wednesday")
		case 4:
			bstr = append(bstr, "Thursday")
		case 5:
			bstr = append(bstr, "Friday")
		case 6:
			bstr = append(bstr, "Saturday")
		case 7:
			bstr = append(bstr, "Away or vacation")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v GetWeeklyDaysToReturn) IsSunday() bool         { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v GetWeeklyDaysToReturn) IsMonday() bool         { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v GetWeeklyDaysToReturn) IsTuesday() bool        { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v GetWeeklyDaysToReturn) IsWednesday() bool      { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v GetWeeklyDaysToReturn) IsThursday() bool       { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v GetWeeklyDaysToReturn) IsFriday() bool         { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v GetWeeklyDaysToReturn) IsSaturday() bool       { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v GetWeeklyDaysToReturn) IsAwayOrVacation() bool { return zcl.BitmapTest([]byte(v[:]), 7) }
func (v *GetWeeklyDaysToReturn) SetSunday(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *GetWeeklyDaysToReturn) SetMonday(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *GetWeeklyDaysToReturn) SetTuesday(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *GetWeeklyDaysToReturn) SetWednesday(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *GetWeeklyDaysToReturn) SetThursday(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
}
func (v *GetWeeklyDaysToReturn) SetFriday(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b))
}
func (v *GetWeeklyDaysToReturn) SetSaturday(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b))
}
func (v *GetWeeklyDaysToReturn) SetAwayOrVacation(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 7, b))
}

func (GetWeeklyDaysToReturn) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Sunday"},
		{Value: 1, Name: "Monday"},
		{Value: 2, Name: "Tuesday"},
		{Value: 3, Name: "Wednesday"},
		{Value: 4, Name: "Thursday"},
		{Value: 5, Name: "Friday"},
		{Value: 6, Name: "Saturday"},
		{Value: 7, Name: "Away or vacation"},
	}
}

func (GetWeeklyModeToReturn) Name() string        { return `Get Weekly Mode To Return` }
func (GetWeeklyModeToReturn) Description() string { return `` }

type GetWeeklyModeToReturn zcl.Zbmp8

func (v *GetWeeklyModeToReturn) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *GetWeeklyModeToReturn) Value() zcl.Val     { return v }

func (v GetWeeklyModeToReturn) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *GetWeeklyModeToReturn) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = GetWeeklyModeToReturn(*nt)
	return br, err
}

func (v GetWeeklyModeToReturn) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *GetWeeklyModeToReturn) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = GetWeeklyModeToReturn(*a)
	return nil
}

func (v *GetWeeklyModeToReturn) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = GetWeeklyModeToReturn(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v GetWeeklyModeToReturn) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Heat setpoint")
		case 1:
			bstr = append(bstr, "Cool setpoint")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v GetWeeklyModeToReturn) IsHeatSetpoint() bool { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v GetWeeklyModeToReturn) IsCoolSetpoint() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v *GetWeeklyModeToReturn) SetHeatSetpoint(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *GetWeeklyModeToReturn) SetCoolSetpoint(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}

func (GetWeeklyModeToReturn) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Heat setpoint"},
		{Value: 1, Name: "Cool setpoint"},
	}
}

const HvacSystemTypeConfigurationAttr zcl.AttrID = 9

func (HvacSystemTypeConfiguration) ID() zcl.AttrID   { return HvacSystemTypeConfigurationAttr }
func (HvacSystemTypeConfiguration) Readable() bool   { return false }
func (HvacSystemTypeConfiguration) Writable() bool   { return false }
func (HvacSystemTypeConfiguration) Reportable() bool { return false }
func (HvacSystemTypeConfiguration) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v HvacSystemTypeConfiguration) AttrID() zcl.AttrID   { return v.ID() }
func (v HvacSystemTypeConfiguration) AttrType() zcl.TypeID { return v.TypeID() }
func (v *HvacSystemTypeConfiguration) AttrValue() zcl.Val  { return v.Value() }

func (HvacSystemTypeConfiguration) Name() string { return `HVAC System Type Configuration` }
func (HvacSystemTypeConfiguration) Description() string {
	return `HVAC system type controlled by the thermostat.
Bit | Description
0-1 | Cooling systemn stage
    | 00 Stage 1
    | 01 Stage 2
    | 10 Stage 3
2-3 | Heating system stage
    | 00 Stage 1
    | 01 Stage 2
    | 10 Stage 3
4   | Heating system type
    | 0 Conventional
    | 1 Heat pump
5   | Heating fuel source
    | 0 Electric
    | 1 Gas`
}

// HvacSystemTypeConfiguration HVAC system type controlled by the thermostat.
// Bit | Description
// 0-1 | Cooling systemn stage
// | 00 Stage 1
// | 01 Stage 2
// | 10 Stage 3
// 2-3 | Heating system stage
// | 00 Stage 1
// | 01 Stage 2
// | 10 Stage 3
// 4   | Heating system type
// | 0 Conventional
// | 1 Heat pump
// 5   | Heating fuel source
// | 0 Electric
// | 1 Gas
type HvacSystemTypeConfiguration zcl.Zbmp8

func (v *HvacSystemTypeConfiguration) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *HvacSystemTypeConfiguration) Value() zcl.Val     { return v }

func (v HvacSystemTypeConfiguration) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *HvacSystemTypeConfiguration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = HvacSystemTypeConfiguration(*nt)
	return br, err
}

func (v HvacSystemTypeConfiguration) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *HvacSystemTypeConfiguration) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = HvacSystemTypeConfiguration(*a)
	return nil
}

func (v *HvacSystemTypeConfiguration) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = HvacSystemTypeConfiguration(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v HvacSystemTypeConfiguration) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 1:
			bstr = append(bstr, "Cool stage 2")
		case 2:
			bstr = append(bstr, "Cool stage 3")
		case 3:
			bstr = append(bstr, "Heat stage 2")
		case 4:
			bstr = append(bstr, "Heat stage 3")
		case 5:
			bstr = append(bstr, "Heat pump")
		case 6:
			bstr = append(bstr, "Gas fuel source")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v HvacSystemTypeConfiguration) IsCoolStage2() bool    { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v HvacSystemTypeConfiguration) IsCoolStage3() bool    { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v HvacSystemTypeConfiguration) IsHeatStage2() bool    { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v HvacSystemTypeConfiguration) IsHeatStage3() bool    { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v HvacSystemTypeConfiguration) IsHeatPump() bool      { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v HvacSystemTypeConfiguration) IsGasFuelSource() bool { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v *HvacSystemTypeConfiguration) SetCoolStage2(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *HvacSystemTypeConfiguration) SetCoolStage3(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *HvacSystemTypeConfiguration) SetHeatStage2(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *HvacSystemTypeConfiguration) SetHeatStage3(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
}
func (v *HvacSystemTypeConfiguration) SetHeatPump(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b))
}
func (v *HvacSystemTypeConfiguration) SetGasFuelSource(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b))
}

func (HvacSystemTypeConfiguration) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 1, Name: "Cool stage 2"},
		{Value: 2, Name: "Cool stage 3"},
		{Value: 3, Name: "Heat stage 2"},
		{Value: 4, Name: "Heat stage 3"},
		{Value: 5, Name: "Heat pump"},
		{Value: 6, Name: "Gas fuel source"},
	}
}

const KeypadLockoutAttr zcl.AttrID = 1

func (KeypadLockout) ID() zcl.AttrID   { return KeypadLockoutAttr }
func (KeypadLockout) Readable() bool   { return true }
func (KeypadLockout) Writable() bool   { return true }
func (KeypadLockout) Reportable() bool { return false }
func (KeypadLockout) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v KeypadLockout) AttrID() zcl.AttrID   { return v.ID() }
func (v KeypadLockout) AttrType() zcl.TypeID { return v.TypeID() }
func (v *KeypadLockout) AttrValue() zcl.Val  { return v.Value() }

func (KeypadLockout) Name() string { return `Keypad Lockout` }
func (KeypadLockout) Description() string {
	return `level of functionality that is available to the user via the keypad`
}

// KeypadLockout level of functionality that is available to the user via the keypad
type KeypadLockout zcl.Zenum8

func (v *KeypadLockout) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *KeypadLockout) Value() zcl.Val     { return v }

func (v KeypadLockout) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *KeypadLockout) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = KeypadLockout(*nt)
	return br, err
}

func (v KeypadLockout) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *KeypadLockout) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = KeypadLockout(*a)
	return nil
}

func (v *KeypadLockout) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = KeypadLockout(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v KeypadLockout) String() string {
	switch v {
	case 0x00:
		return "No lockout"
	case 0x01:
		return "Level 1 lockout"
	case 0x02:
		return "Level 2 lockout"
	case 0x03:
		return "Level 3 lockout"
	case 0x04:
		return "Level 4 lockout"
	case 0x05:
		return "Level 5 lockout"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v KeypadLockout) IsNoLockout() bool     { return v == 0x00 }
func (v KeypadLockout) IsLevel1Lockout() bool { return v == 0x01 }
func (v KeypadLockout) IsLevel2Lockout() bool { return v == 0x02 }
func (v KeypadLockout) IsLevel3Lockout() bool { return v == 0x03 }
func (v KeypadLockout) IsLevel4Lockout() bool { return v == 0x04 }
func (v KeypadLockout) IsLevel5Lockout() bool { return v == 0x05 }
func (v *KeypadLockout) SetNoLockout()        { *v = 0x00 }
func (v *KeypadLockout) SetLevel1Lockout()    { *v = 0x01 }
func (v *KeypadLockout) SetLevel2Lockout()    { *v = 0x02 }
func (v *KeypadLockout) SetLevel3Lockout()    { *v = 0x03 }
func (v *KeypadLockout) SetLevel4Lockout()    { *v = 0x04 }
func (v *KeypadLockout) SetLevel5Lockout()    { *v = 0x05 }

func (KeypadLockout) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "No lockout"},
		{Value: 0x01, Name: "Level 1 lockout"},
		{Value: 0x02, Name: "Level 2 lockout"},
		{Value: 0x03, Name: "Level 3 lockout"},
		{Value: 0x04, Name: "Level 4 lockout"},
		{Value: 0x05, Name: "Level 5 lockout"},
	}
}

const LifetimeEnergyConsumedAttr zcl.AttrID = 23

func (LifetimeEnergyConsumed) ID() zcl.AttrID   { return LifetimeEnergyConsumedAttr }
func (LifetimeEnergyConsumed) Readable() bool   { return true }
func (LifetimeEnergyConsumed) Writable() bool   { return false }
func (LifetimeEnergyConsumed) Reportable() bool { return false }
func (LifetimeEnergyConsumed) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LifetimeEnergyConsumed) AttrID() zcl.AttrID   { return v.ID() }
func (v LifetimeEnergyConsumed) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LifetimeEnergyConsumed) AttrValue() zcl.Val  { return v.Value() }

func (LifetimeEnergyConsumed) Name() string { return `Lifetime Energy Consumed` }
func (LifetimeEnergyConsumed) Description() string {
	return `accumulated energy consumption of the pump through the entire lifetime
of the pump in kWh. The value of the LifetimeEnergyConsumed attribute
is updated dynamically as the energy consumption of the pump increases.
If LifetimeEnergyConsumed rises above maximum value it rolls over and
starts at 0 (zero)`
}

// LifetimeEnergyConsumed accumulated energy consumption of the pump through the entire lifetime
// of the pump in kWh. The value of the LifetimeEnergyConsumed attribute
// is updated dynamically as the energy consumption of the pump increases.
// If LifetimeEnergyConsumed rises above maximum value it rolls over and
// starts at 0 (zero)
type LifetimeEnergyConsumed zcl.Zu32

func (v *LifetimeEnergyConsumed) TypeID() zcl.TypeID { return new(zcl.Zu32).TypeID() }
func (v *LifetimeEnergyConsumed) Value() zcl.Val     { return v }

func (v LifetimeEnergyConsumed) MarshalZcl() ([]byte, error) { return zcl.Zu32(v).MarshalZcl() }

func (v *LifetimeEnergyConsumed) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*v = LifetimeEnergyConsumed(*nt)
	return br, err
}

func (v LifetimeEnergyConsumed) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu32(v))
}

func (v *LifetimeEnergyConsumed) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu32)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LifetimeEnergyConsumed(*a)
	return nil
}

func (v *LifetimeEnergyConsumed) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu32); ok {
		*v = LifetimeEnergyConsumed(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LifetimeEnergyConsumed) String() string {
	return zcl.KilowattHours.Format(float64(v))
}

const LifetimeRunningHoursAttr zcl.AttrID = 21

func (LifetimeRunningHours) ID() zcl.AttrID   { return LifetimeRunningHoursAttr }
func (LifetimeRunningHours) Readable() bool   { return true }
func (LifetimeRunningHours) Writable() bool   { return true }
func (LifetimeRunningHours) Reportable() bool { return false }
func (LifetimeRunningHours) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LifetimeRunningHours) AttrID() zcl.AttrID   { return v.ID() }
func (v LifetimeRunningHours) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LifetimeRunningHours) AttrValue() zcl.Val  { return v.Value() }

func (LifetimeRunningHours) Name() string { return `Lifetime Running Hours` }
func (LifetimeRunningHours) Description() string {
	return `ccumulated number of hours that the pump has been powered and the motor
has been running. It is updated dynamically as it increases. It is
preserved over powercycles of the pump. if LifeTimeRunningHours rises
above maximum value it <rolls over= and starts at 0 (zero)`
}

// LifetimeRunningHours ccumulated number of hours that the pump has been powered and the motor
// has been running. It is updated dynamically as it increases. It is
// preserved over powercycles of the pump. if LifeTimeRunningHours rises
// above maximum value it <rolls over= and starts at 0 (zero)
type LifetimeRunningHours zcl.Zu24

func (v *LifetimeRunningHours) TypeID() zcl.TypeID { return new(zcl.Zu24).TypeID() }
func (v *LifetimeRunningHours) Value() zcl.Val     { return v }

func (v LifetimeRunningHours) MarshalZcl() ([]byte, error) { return zcl.Zu24(v).MarshalZcl() }

func (v *LifetimeRunningHours) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*v = LifetimeRunningHours(*nt)
	return br, err
}

func (v LifetimeRunningHours) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu24(v))
}

func (v *LifetimeRunningHours) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LifetimeRunningHours(*a)
	return nil
}

func (v *LifetimeRunningHours) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu24); ok {
		*v = LifetimeRunningHours(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LifetimeRunningHours) String() string {
	return zcl.Hours.Format(float64(v))
}

const LocalTemperatureAttr zcl.AttrID = 0

func (LocalTemperature) ID() zcl.AttrID   { return LocalTemperatureAttr }
func (LocalTemperature) Readable() bool   { return true }
func (LocalTemperature) Writable() bool   { return false }
func (LocalTemperature) Reportable() bool { return true }
func (LocalTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LocalTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v LocalTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LocalTemperature) AttrValue() zcl.Val  { return v.Value() }

func (LocalTemperature) Name() string { return `Local Temperature` }
func (LocalTemperature) Description() string {
	return `temperature in degrees Celsius, as measured locally or remotely (over
the network), including any adjustments applied by
LocalTemperatureCalibration attribute (if any) as follows:
LocalTemperature = 100 x (temperature in degrees Celsius +
LocalTemperatureCalibration)`
}

// LocalTemperature temperature in degrees Celsius, as measured locally or remotely (over
// the network), including any adjustments applied by
// LocalTemperatureCalibration attribute (if any) as follows:
// LocalTemperature = 100 x (temperature in degrees Celsius +
// LocalTemperatureCalibration)
type LocalTemperature zcl.Zs16

func (v *LocalTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *LocalTemperature) Value() zcl.Val     { return v }

func (v LocalTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *LocalTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = LocalTemperature(*nt)
	return br, err
}

func (v LocalTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *LocalTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LocalTemperature(*a)
	return nil
}

func (v *LocalTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = LocalTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LocalTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v LocalTemperature) Scaled() float64 {
	return float64(v) / 100
}

const LocalTemperatureCalibrationAttr zcl.AttrID = 16

func (LocalTemperatureCalibration) ID() zcl.AttrID   { return LocalTemperatureCalibrationAttr }
func (LocalTemperatureCalibration) Readable() bool   { return true }
func (LocalTemperatureCalibration) Writable() bool   { return true }
func (LocalTemperatureCalibration) Reportable() bool { return false }
func (LocalTemperatureCalibration) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v LocalTemperatureCalibration) AttrID() zcl.AttrID   { return v.ID() }
func (v LocalTemperatureCalibration) AttrType() zcl.TypeID { return v.TypeID() }
func (v *LocalTemperatureCalibration) AttrValue() zcl.Val  { return v.Value() }

func (LocalTemperatureCalibration) Name() string { return `Local Temperature Calibration` }
func (LocalTemperatureCalibration) Description() string {
	return `offset the thermostat server shall make to the measured temperature
(locally or remotely) before calculating, displaying, or communicating
the LocalTemperature attribute`
}

// LocalTemperatureCalibration offset the thermostat server shall make to the measured temperature
// (locally or remotely) before calculating, displaying, or communicating
// the LocalTemperature attribute
type LocalTemperatureCalibration zcl.Zs8

func (v *LocalTemperatureCalibration) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *LocalTemperatureCalibration) Value() zcl.Val     { return v }

func (v LocalTemperatureCalibration) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *LocalTemperatureCalibration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = LocalTemperatureCalibration(*nt)
	return br, err
}

func (v LocalTemperatureCalibration) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *LocalTemperatureCalibration) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = LocalTemperatureCalibration(*a)
	return nil
}

func (v *LocalTemperatureCalibration) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = LocalTemperatureCalibration(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v LocalTemperatureCalibration) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v LocalTemperatureCalibration) Scaled() float64 {
	return float64(v) / 10
}

const MaxCompPressureAttr zcl.AttrID = 6

func (MaxCompPressure) ID() zcl.AttrID   { return MaxCompPressureAttr }
func (MaxCompPressure) Readable() bool   { return true }
func (MaxCompPressure) Writable() bool   { return false }
func (MaxCompPressure) Reportable() bool { return false }
func (MaxCompPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxCompPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxCompPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxCompPressure) AttrValue() zcl.Val  { return v.Value() }

func (MaxCompPressure) Name() string { return `Max Comp Pressure` }
func (MaxCompPressure) Description() string {
	return `the maximum compensated pressure the pump can achieve when it is
running and working in control mode Proportional pressure (ControlMode
attribute of the Pump settings attribute set is set to Proportional
pressure)`
}

// MaxCompPressure the maximum compensated pressure the pump can achieve when it is
// running and working in control mode Proportional pressure (ControlMode
// attribute of the Pump settings attribute set is set to Proportional
// pressure)
type MaxCompPressure zcl.Zs16

func (v *MaxCompPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxCompPressure) Value() zcl.Val     { return v }

func (v MaxCompPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxCompPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxCompPressure(*nt)
	return br, err
}

func (v MaxCompPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxCompPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxCompPressure(*a)
	return nil
}

func (v *MaxCompPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxCompPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxCompPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MaxCompPressure) Scaled() float64 {
	return float64(v) / 10
}

const MaxConstFlowAttr zcl.AttrID = 10

func (MaxConstFlow) ID() zcl.AttrID   { return MaxConstFlowAttr }
func (MaxConstFlow) Readable() bool   { return true }
func (MaxConstFlow) Writable() bool   { return false }
func (MaxConstFlow) Reportable() bool { return false }
func (MaxConstFlow) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxConstFlow) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxConstFlow) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxConstFlow) AttrValue() zcl.Val  { return v.Value() }

func (MaxConstFlow) Name() string { return `Max Const Flow` }
func (MaxConstFlow) Description() string {
	return `the maximum flow the pump can achieve when it is running and working in
control mode Constant flow (ControlMode attribute of the Pump settings
attribute set is set to Constant flow)`
}

// MaxConstFlow the maximum flow the pump can achieve when it is running and working in
// control mode Constant flow (ControlMode attribute of the Pump settings
// attribute set is set to Constant flow)
type MaxConstFlow zcl.Zu16

func (v *MaxConstFlow) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxConstFlow) Value() zcl.Val     { return v }

func (v MaxConstFlow) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxConstFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxConstFlow(*nt)
	return br, err
}

func (v MaxConstFlow) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxConstFlow) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxConstFlow(*a)
	return nil
}

func (v *MaxConstFlow) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxConstFlow(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxConstFlow) String() string {
	return zcl.CubicMetersPerHour.Format(float64(v) / 10)
}

func (v MaxConstFlow) Scaled() float64 {
	return float64(v) / 10
}

const MaxConstPressureAttr zcl.AttrID = 4

func (MaxConstPressure) ID() zcl.AttrID   { return MaxConstPressureAttr }
func (MaxConstPressure) Readable() bool   { return true }
func (MaxConstPressure) Writable() bool   { return false }
func (MaxConstPressure) Reportable() bool { return false }
func (MaxConstPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxConstPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxConstPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxConstPressure) AttrValue() zcl.Val  { return v.Value() }

func (MaxConstPressure) Name() string { return `Max Const Pressure` }
func (MaxConstPressure) Description() string {
	return `the maximum pressure the pump can achieve when it is running and
working in control mode constant pressure (ControlMode attribute of the
Pump settings attribute set is set to Constant pressure)`
}

// MaxConstPressure the maximum pressure the pump can achieve when it is running and
// working in control mode constant pressure (ControlMode attribute of the
// Pump settings attribute set is set to Constant pressure)
type MaxConstPressure zcl.Zs16

func (v *MaxConstPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxConstPressure) Value() zcl.Val     { return v }

func (v MaxConstPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxConstPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxConstPressure(*nt)
	return br, err
}

func (v MaxConstPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxConstPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxConstPressure(*a)
	return nil
}

func (v *MaxConstPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxConstPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxConstPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MaxConstPressure) Scaled() float64 {
	return float64(v) / 10
}

const MaxConstSpeedAttr zcl.AttrID = 8

func (MaxConstSpeed) ID() zcl.AttrID   { return MaxConstSpeedAttr }
func (MaxConstSpeed) Readable() bool   { return true }
func (MaxConstSpeed) Writable() bool   { return false }
func (MaxConstSpeed) Reportable() bool { return false }
func (MaxConstSpeed) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxConstSpeed) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxConstSpeed) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxConstSpeed) AttrValue() zcl.Val  { return v.Value() }

func (MaxConstSpeed) Name() string { return `Max Const Speed` }
func (MaxConstSpeed) Description() string {
	return `the maximum speed the pump can achieve when it is running and working
in control mode Constant speed (ControlMode attribute of the Pump
settings attribute set is set to Constant speed)`
}

// MaxConstSpeed the maximum speed the pump can achieve when it is running and working
// in control mode Constant speed (ControlMode attribute of the Pump
// settings attribute set is set to Constant speed)
type MaxConstSpeed zcl.Zu16

func (v *MaxConstSpeed) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxConstSpeed) Value() zcl.Val     { return v }

func (v MaxConstSpeed) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxConstSpeed) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxConstSpeed(*nt)
	return br, err
}

func (v MaxConstSpeed) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxConstSpeed) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxConstSpeed(*a)
	return nil
}

func (v *MaxConstSpeed) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxConstSpeed(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxConstSpeed) String() string {
	return zcl.RevolutionsPerMinute.Format(float64(v))
}

const MaxConstTempAttr zcl.AttrID = 12

func (MaxConstTemp) ID() zcl.AttrID   { return MaxConstTempAttr }
func (MaxConstTemp) Readable() bool   { return true }
func (MaxConstTemp) Writable() bool   { return false }
func (MaxConstTemp) Reportable() bool { return false }
func (MaxConstTemp) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxConstTemp) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxConstTemp) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxConstTemp) AttrValue() zcl.Val  { return v.Value() }

func (MaxConstTemp) Name() string { return `Max Const Temp` }
func (MaxConstTemp) Description() string {
	return `the maximum temperature the pump can maintain in the system when it is
running and working in control mode Constant temperature (ControlMode
attribute of the Pump settings attribute set is set to Constant
temperature)`
}

// MaxConstTemp the maximum temperature the pump can maintain in the system when it is
// running and working in control mode Constant temperature (ControlMode
// attribute of the Pump settings attribute set is set to Constant
// temperature)
type MaxConstTemp zcl.Zs16

func (v *MaxConstTemp) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxConstTemp) Value() zcl.Val     { return v }

func (v MaxConstTemp) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxConstTemp) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxConstTemp(*nt)
	return br, err
}

func (v MaxConstTemp) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxConstTemp) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxConstTemp(*a)
	return nil
}

func (v *MaxConstTemp) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxConstTemp(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxConstTemp) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v MaxConstTemp) Scaled() float64 {
	return float64(v) / 100
}

const MaxCoolSetpointLimitAttr zcl.AttrID = 24

func (MaxCoolSetpointLimit) ID() zcl.AttrID   { return MaxCoolSetpointLimitAttr }
func (MaxCoolSetpointLimit) Readable() bool   { return true }
func (MaxCoolSetpointLimit) Writable() bool   { return true }
func (MaxCoolSetpointLimit) Reportable() bool { return false }
func (MaxCoolSetpointLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxCoolSetpointLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxCoolSetpointLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxCoolSetpointLimit) AttrValue() zcl.Val  { return v.Value() }

func (MaxCoolSetpointLimit) Name() string { return `Max Cool Setpoint Limit` }
func (MaxCoolSetpointLimit) Description() string {
	return `maximum level that the cooling setpoint may be set to. It must be less
than or equal to AbsMaxCoolSetpointLimit. If this attribute is not
present, it shall be taken as equal to AbsMaxCoolSetpointLimit`
}

// MaxCoolSetpointLimit maximum level that the cooling setpoint may be set to. It must be less
// than or equal to AbsMaxCoolSetpointLimit. If this attribute is not
// present, it shall be taken as equal to AbsMaxCoolSetpointLimit
type MaxCoolSetpointLimit zcl.Zs16

func (v *MaxCoolSetpointLimit) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxCoolSetpointLimit) Value() zcl.Val     { return v }

func (v MaxCoolSetpointLimit) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxCoolSetpointLimit(*nt)
	return br, err
}

func (v MaxCoolSetpointLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxCoolSetpointLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxCoolSetpointLimit(*a)
	return nil
}

func (v *MaxCoolSetpointLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxCoolSetpointLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxCoolSetpointLimit) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v MaxCoolSetpointLimit) Scaled() float64 {
	return float64(v) / 10
}

const MaxFlowAttr zcl.AttrID = 2

func (MaxFlow) ID() zcl.AttrID   { return MaxFlowAttr }
func (MaxFlow) Readable() bool   { return true }
func (MaxFlow) Writable() bool   { return false }
func (MaxFlow) Reportable() bool { return false }
func (MaxFlow) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxFlow) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxFlow) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxFlow) AttrValue() zcl.Val  { return v.Value() }

func (MaxFlow) Name() string { return `Max Flow` }
func (MaxFlow) Description() string {
	return `the maximum flow the pump can achieve. It is a physical limit, and does
not apply to any specific control mode or operation mode`
}

// MaxFlow the maximum flow the pump can achieve. It is a physical limit, and does
// not apply to any specific control mode or operation mode
type MaxFlow zcl.Zu16

func (v *MaxFlow) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxFlow) Value() zcl.Val     { return v }

func (v MaxFlow) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxFlow(*nt)
	return br, err
}

func (v MaxFlow) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxFlow) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxFlow(*a)
	return nil
}

func (v *MaxFlow) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxFlow(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxFlow) String() string {
	return zcl.CubicMetersPerHour.Format(float64(v) / 10)
}

func (v MaxFlow) Scaled() float64 {
	return float64(v) / 10
}

const MaxHeatSetpointLimitAttr zcl.AttrID = 22

func (MaxHeatSetpointLimit) ID() zcl.AttrID   { return MaxHeatSetpointLimitAttr }
func (MaxHeatSetpointLimit) Readable() bool   { return true }
func (MaxHeatSetpointLimit) Writable() bool   { return true }
func (MaxHeatSetpointLimit) Reportable() bool { return false }
func (MaxHeatSetpointLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxHeatSetpointLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxHeatSetpointLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxHeatSetpointLimit) AttrValue() zcl.Val  { return v.Value() }

func (MaxHeatSetpointLimit) Name() string { return `Max Heat Setpoint Limit` }
func (MaxHeatSetpointLimit) Description() string {
	return `maximum level that the heating setpoint MAY be set to. It must be less
than or equal to AbsMaxHeatSetpointLimit. If this attribute is not
present, it shall be taken as equal to AbsMaxHeatSetpointLimit`
}

// MaxHeatSetpointLimit maximum level that the heating setpoint MAY be set to. It must be less
// than or equal to AbsMaxHeatSetpointLimit. If this attribute is not
// present, it shall be taken as equal to AbsMaxHeatSetpointLimit
type MaxHeatSetpointLimit zcl.Zs16

func (v *MaxHeatSetpointLimit) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxHeatSetpointLimit) Value() zcl.Val     { return v }

func (v MaxHeatSetpointLimit) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxHeatSetpointLimit(*nt)
	return br, err
}

func (v MaxHeatSetpointLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxHeatSetpointLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxHeatSetpointLimit(*a)
	return nil
}

func (v *MaxHeatSetpointLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxHeatSetpointLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxHeatSetpointLimit) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v MaxHeatSetpointLimit) Scaled() float64 {
	return float64(v) / 10
}

const MaxPressureAttr zcl.AttrID = 0

func (MaxPressure) ID() zcl.AttrID   { return MaxPressureAttr }
func (MaxPressure) Readable() bool   { return true }
func (MaxPressure) Writable() bool   { return false }
func (MaxPressure) Reportable() bool { return false }
func (MaxPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxPressure) AttrValue() zcl.Val  { return v.Value() }

func (MaxPressure) Name() string { return `Max Pressure` }
func (MaxPressure) Description() string {
	return `the maximum pressure the pump can achieve. It is a physical limit,
and does not apply to any specific control mode or operation mode`
}

// MaxPressure the maximum pressure the pump can achieve. It is a physical limit,
// and does not apply to any specific control mode or operation mode
type MaxPressure zcl.Zs16

func (v *MaxPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MaxPressure) Value() zcl.Val     { return v }

func (v MaxPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MaxPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxPressure(*nt)
	return br, err
}

func (v MaxPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MaxPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxPressure(*a)
	return nil
}

func (v *MaxPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MaxPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MaxPressure) Scaled() float64 {
	return float64(v) / 10
}

const MaxSpeedAttr zcl.AttrID = 1

func (MaxSpeed) ID() zcl.AttrID   { return MaxSpeedAttr }
func (MaxSpeed) Readable() bool   { return true }
func (MaxSpeed) Writable() bool   { return false }
func (MaxSpeed) Reportable() bool { return false }
func (MaxSpeed) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MaxSpeed) AttrID() zcl.AttrID   { return v.ID() }
func (v MaxSpeed) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MaxSpeed) AttrValue() zcl.Val  { return v.Value() }

func (MaxSpeed) Name() string { return `Max Speed` }
func (MaxSpeed) Description() string {
	return `the maximum speed the pump can achieve. It is a physical limit, and
does not apply to any specific control mode or operation mode`
}

// MaxSpeed the maximum speed the pump can achieve. It is a physical limit, and
// does not apply to any specific control mode or operation mode
type MaxSpeed zcl.Zu16

func (v *MaxSpeed) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MaxSpeed) Value() zcl.Val     { return v }

func (v MaxSpeed) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MaxSpeed) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MaxSpeed(*nt)
	return br, err
}

func (v MaxSpeed) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MaxSpeed) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MaxSpeed(*a)
	return nil
}

func (v *MaxSpeed) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MaxSpeed(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MaxSpeed) String() string {
	return zcl.RevolutionsPerMinute.Format(float64(v))
}

const MinCompPressureAttr zcl.AttrID = 5

func (MinCompPressure) ID() zcl.AttrID   { return MinCompPressureAttr }
func (MinCompPressure) Readable() bool   { return true }
func (MinCompPressure) Writable() bool   { return false }
func (MinCompPressure) Reportable() bool { return false }
func (MinCompPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinCompPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MinCompPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinCompPressure) AttrValue() zcl.Val  { return v.Value() }

func (MinCompPressure) Name() string { return `Min Comp Pressure` }
func (MinCompPressure) Description() string {
	return `the minimum compensated pressure the pump can achieve when it is
running and working in control mode Proportional pressure (ControlMode
attribute of the Pump settings attribute set is set to Proportional
pressure)`
}

// MinCompPressure the minimum compensated pressure the pump can achieve when it is
// running and working in control mode Proportional pressure (ControlMode
// attribute of the Pump settings attribute set is set to Proportional
// pressure)
type MinCompPressure zcl.Zs16

func (v *MinCompPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MinCompPressure) Value() zcl.Val     { return v }

func (v MinCompPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MinCompPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinCompPressure(*nt)
	return br, err
}

func (v MinCompPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MinCompPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinCompPressure(*a)
	return nil
}

func (v *MinCompPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MinCompPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinCompPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MinCompPressure) Scaled() float64 {
	return float64(v) / 10
}

const MinConstFlowAttr zcl.AttrID = 9

func (MinConstFlow) ID() zcl.AttrID   { return MinConstFlowAttr }
func (MinConstFlow) Readable() bool   { return true }
func (MinConstFlow) Writable() bool   { return false }
func (MinConstFlow) Reportable() bool { return false }
func (MinConstFlow) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinConstFlow) AttrID() zcl.AttrID   { return v.ID() }
func (v MinConstFlow) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinConstFlow) AttrValue() zcl.Val  { return v.Value() }

func (MinConstFlow) Name() string { return `Min Const Flow` }
func (MinConstFlow) Description() string {
	return `the minimum flow the pump can achieve when it is running and working in
control mode Constant flow (ControlMode attribute of the Pump settings
attribute set is set to Constant flow)`
}

// MinConstFlow the minimum flow the pump can achieve when it is running and working in
// control mode Constant flow (ControlMode attribute of the Pump settings
// attribute set is set to Constant flow)
type MinConstFlow zcl.Zu16

func (v *MinConstFlow) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MinConstFlow) Value() zcl.Val     { return v }

func (v MinConstFlow) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MinConstFlow) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinConstFlow(*nt)
	return br, err
}

func (v MinConstFlow) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MinConstFlow) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinConstFlow(*a)
	return nil
}

func (v *MinConstFlow) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MinConstFlow(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinConstFlow) String() string {
	return zcl.CubicMetersPerHour.Format(float64(v) / 10)
}

func (v MinConstFlow) Scaled() float64 {
	return float64(v) / 10
}

const MinConstPressureAttr zcl.AttrID = 3

func (MinConstPressure) ID() zcl.AttrID   { return MinConstPressureAttr }
func (MinConstPressure) Readable() bool   { return true }
func (MinConstPressure) Writable() bool   { return false }
func (MinConstPressure) Reportable() bool { return false }
func (MinConstPressure) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinConstPressure) AttrID() zcl.AttrID   { return v.ID() }
func (v MinConstPressure) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinConstPressure) AttrValue() zcl.Val  { return v.Value() }

func (MinConstPressure) Name() string { return `Min Const Pressure` }
func (MinConstPressure) Description() string {
	return `the minimum pressure the pump can achieve when it is running and
working in control mode constant pressure (ControlMode attribute of the
Pump settings attribute set is set to Constant pressure)`
}

// MinConstPressure the minimum pressure the pump can achieve when it is running and
// working in control mode constant pressure (ControlMode attribute of the
// Pump settings attribute set is set to Constant pressure)
type MinConstPressure zcl.Zs16

func (v *MinConstPressure) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MinConstPressure) Value() zcl.Val     { return v }

func (v MinConstPressure) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MinConstPressure) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinConstPressure(*nt)
	return br, err
}

func (v MinConstPressure) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MinConstPressure) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinConstPressure(*a)
	return nil
}

func (v *MinConstPressure) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MinConstPressure(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinConstPressure) String() string {
	return zcl.Kilopascals.Format(float64(v) / 10)
}

func (v MinConstPressure) Scaled() float64 {
	return float64(v) / 10
}

const MinConstSpeedAttr zcl.AttrID = 7

func (MinConstSpeed) ID() zcl.AttrID   { return MinConstSpeedAttr }
func (MinConstSpeed) Readable() bool   { return true }
func (MinConstSpeed) Writable() bool   { return false }
func (MinConstSpeed) Reportable() bool { return false }
func (MinConstSpeed) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinConstSpeed) AttrID() zcl.AttrID   { return v.ID() }
func (v MinConstSpeed) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinConstSpeed) AttrValue() zcl.Val  { return v.Value() }

func (MinConstSpeed) Name() string { return `Min Const Speed` }
func (MinConstSpeed) Description() string {
	return `the minimum speed the pump can achieve when it is running and working
in control mode Constant speed (ControlMode attribute of the Pump
settings attribute set is set to Constant speed)`
}

// MinConstSpeed the minimum speed the pump can achieve when it is running and working
// in control mode Constant speed (ControlMode attribute of the Pump
// settings attribute set is set to Constant speed)
type MinConstSpeed zcl.Zu16

func (v *MinConstSpeed) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *MinConstSpeed) Value() zcl.Val     { return v }

func (v MinConstSpeed) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *MinConstSpeed) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinConstSpeed(*nt)
	return br, err
}

func (v MinConstSpeed) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *MinConstSpeed) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinConstSpeed(*a)
	return nil
}

func (v *MinConstSpeed) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = MinConstSpeed(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinConstSpeed) String() string {
	return zcl.RevolutionsPerMinute.Format(float64(v))
}

const MinConstTempAttr zcl.AttrID = 11

func (MinConstTemp) ID() zcl.AttrID   { return MinConstTempAttr }
func (MinConstTemp) Readable() bool   { return true }
func (MinConstTemp) Writable() bool   { return false }
func (MinConstTemp) Reportable() bool { return false }
func (MinConstTemp) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinConstTemp) AttrID() zcl.AttrID   { return v.ID() }
func (v MinConstTemp) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinConstTemp) AttrValue() zcl.Val  { return v.Value() }

func (MinConstTemp) Name() string { return `Min Const Temp` }
func (MinConstTemp) Description() string {
	return `the minimum temperature the pump can maintain in the system when it is
running and working in control mode Constant temperature (ControlMode
attribute of the Pump settings attribute set is set to Constant
temperature)`
}

// MinConstTemp the minimum temperature the pump can maintain in the system when it is
// running and working in control mode Constant temperature (ControlMode
// attribute of the Pump settings attribute set is set to Constant
// temperature)
type MinConstTemp zcl.Zs16

func (v *MinConstTemp) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MinConstTemp) Value() zcl.Val     { return v }

func (v MinConstTemp) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MinConstTemp) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinConstTemp(*nt)
	return br, err
}

func (v MinConstTemp) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MinConstTemp) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinConstTemp(*a)
	return nil
}

func (v *MinConstTemp) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MinConstTemp(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinConstTemp) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v MinConstTemp) Scaled() float64 {
	return float64(v) / 100
}

const MinCoolSetpointLimitAttr zcl.AttrID = 23

func (MinCoolSetpointLimit) ID() zcl.AttrID   { return MinCoolSetpointLimitAttr }
func (MinCoolSetpointLimit) Readable() bool   { return true }
func (MinCoolSetpointLimit) Writable() bool   { return true }
func (MinCoolSetpointLimit) Reportable() bool { return false }
func (MinCoolSetpointLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinCoolSetpointLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v MinCoolSetpointLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinCoolSetpointLimit) AttrValue() zcl.Val  { return v.Value() }

func (MinCoolSetpointLimit) Name() string { return `Min Cool Setpoint Limit` }
func (MinCoolSetpointLimit) Description() string {
	return `minimum level that the cooling setpoint may be set to. It must be
greater than or equal to AbsMinCoolSetpointLimit. If this attribute is
not present, it shall be taken as equal to AbsMinCoolSetpointLimit`
}

// MinCoolSetpointLimit minimum level that the cooling setpoint may be set to. It must be
// greater than or equal to AbsMinCoolSetpointLimit. If this attribute is
// not present, it shall be taken as equal to AbsMinCoolSetpointLimit
type MinCoolSetpointLimit zcl.Zs16

func (v *MinCoolSetpointLimit) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MinCoolSetpointLimit) Value() zcl.Val     { return v }

func (v MinCoolSetpointLimit) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MinCoolSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinCoolSetpointLimit(*nt)
	return br, err
}

func (v MinCoolSetpointLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MinCoolSetpointLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinCoolSetpointLimit(*a)
	return nil
}

func (v *MinCoolSetpointLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MinCoolSetpointLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinCoolSetpointLimit) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v MinCoolSetpointLimit) Scaled() float64 {
	return float64(v) / 10
}

const MinHeatSetpointLimitAttr zcl.AttrID = 21

func (MinHeatSetpointLimit) ID() zcl.AttrID   { return MinHeatSetpointLimitAttr }
func (MinHeatSetpointLimit) Readable() bool   { return true }
func (MinHeatSetpointLimit) Writable() bool   { return true }
func (MinHeatSetpointLimit) Reportable() bool { return false }
func (MinHeatSetpointLimit) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinHeatSetpointLimit) AttrID() zcl.AttrID   { return v.ID() }
func (v MinHeatSetpointLimit) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinHeatSetpointLimit) AttrValue() zcl.Val  { return v.Value() }

func (MinHeatSetpointLimit) Name() string { return `Min Heat Setpoint Limit` }
func (MinHeatSetpointLimit) Description() string {
	return `minimum level that the heating setpoint may be set to. If this
attribute is not present, it shall be taken as equal to
AbsMinHeatSetpointLimit`
}

// MinHeatSetpointLimit minimum level that the heating setpoint may be set to. If this
// attribute is not present, it shall be taken as equal to
// AbsMinHeatSetpointLimit
type MinHeatSetpointLimit zcl.Zs16

func (v *MinHeatSetpointLimit) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *MinHeatSetpointLimit) Value() zcl.Val     { return v }

func (v MinHeatSetpointLimit) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *MinHeatSetpointLimit) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = MinHeatSetpointLimit(*nt)
	return br, err
}

func (v MinHeatSetpointLimit) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *MinHeatSetpointLimit) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinHeatSetpointLimit(*a)
	return nil
}

func (v *MinHeatSetpointLimit) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = MinHeatSetpointLimit(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinHeatSetpointLimit) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v MinHeatSetpointLimit) Scaled() float64 {
	return float64(v) / 10
}

const MinSetpointDeadBandAttr zcl.AttrID = 25

func (MinSetpointDeadBand) ID() zcl.AttrID   { return MinSetpointDeadBandAttr }
func (MinSetpointDeadBand) Readable() bool   { return true }
func (MinSetpointDeadBand) Writable() bool   { return true }
func (MinSetpointDeadBand) Reportable() bool { return false }
func (MinSetpointDeadBand) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v MinSetpointDeadBand) AttrID() zcl.AttrID   { return v.ID() }
func (v MinSetpointDeadBand) AttrType() zcl.TypeID { return v.TypeID() }
func (v *MinSetpointDeadBand) AttrValue() zcl.Val  { return v.Value() }

func (MinSetpointDeadBand) Name() string { return `Min Setpoint Dead Band` }
func (MinSetpointDeadBand) Description() string {
	return `minimum difference between the Heat Setpoint and the Cool SetPoint`
}

// MinSetpointDeadBand minimum difference between the Heat Setpoint and the Cool SetPoint
type MinSetpointDeadBand zcl.Zs8

func (v *MinSetpointDeadBand) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *MinSetpointDeadBand) Value() zcl.Val     { return v }

func (v MinSetpointDeadBand) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *MinSetpointDeadBand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = MinSetpointDeadBand(*nt)
	return br, err
}

func (v MinSetpointDeadBand) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *MinSetpointDeadBand) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = MinSetpointDeadBand(*a)
	return nil
}

func (v *MinSetpointDeadBand) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = MinSetpointDeadBand(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v MinSetpointDeadBand) String() string {
	return zcl.Sprintf("%v", zcl.Zs8(v))
}

const NumberOfDailyTransitionsAttr zcl.AttrID = 34

func (NumberOfDailyTransitions) ID() zcl.AttrID   { return NumberOfDailyTransitionsAttr }
func (NumberOfDailyTransitions) Readable() bool   { return true }
func (NumberOfDailyTransitions) Writable() bool   { return false }
func (NumberOfDailyTransitions) Reportable() bool { return false }
func (NumberOfDailyTransitions) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NumberOfDailyTransitions) AttrID() zcl.AttrID   { return v.ID() }
func (v NumberOfDailyTransitions) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NumberOfDailyTransitions) AttrValue() zcl.Val  { return v.Value() }

func (NumberOfDailyTransitions) Name() string { return `Number Of Daily Transitions` }
func (NumberOfDailyTransitions) Description() string {
	return `how many daily schedule transitions the thermostat is capable of
handling`
}

// NumberOfDailyTransitions how many daily schedule transitions the thermostat is capable of
// handling
type NumberOfDailyTransitions zcl.Zu8

func (v *NumberOfDailyTransitions) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *NumberOfDailyTransitions) Value() zcl.Val     { return v }

func (v NumberOfDailyTransitions) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *NumberOfDailyTransitions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = NumberOfDailyTransitions(*nt)
	return br, err
}

func (v NumberOfDailyTransitions) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *NumberOfDailyTransitions) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NumberOfDailyTransitions(*a)
	return nil
}

func (v *NumberOfDailyTransitions) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = NumberOfDailyTransitions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NumberOfDailyTransitions) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const NumberOfWeeklyTransitionsAttr zcl.AttrID = 33

func (NumberOfWeeklyTransitions) ID() zcl.AttrID   { return NumberOfWeeklyTransitionsAttr }
func (NumberOfWeeklyTransitions) Readable() bool   { return true }
func (NumberOfWeeklyTransitions) Writable() bool   { return false }
func (NumberOfWeeklyTransitions) Reportable() bool { return false }
func (NumberOfWeeklyTransitions) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v NumberOfWeeklyTransitions) AttrID() zcl.AttrID   { return v.ID() }
func (v NumberOfWeeklyTransitions) AttrType() zcl.TypeID { return v.TypeID() }
func (v *NumberOfWeeklyTransitions) AttrValue() zcl.Val  { return v.Value() }

func (NumberOfWeeklyTransitions) Name() string { return `Number Of Weekly Transitions` }
func (NumberOfWeeklyTransitions) Description() string {
	return `how many weekly schedule transitions the thermostat is capable of
handling`
}

// NumberOfWeeklyTransitions how many weekly schedule transitions the thermostat is capable of
// handling
type NumberOfWeeklyTransitions zcl.Zu8

func (v *NumberOfWeeklyTransitions) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *NumberOfWeeklyTransitions) Value() zcl.Val     { return v }

func (v NumberOfWeeklyTransitions) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *NumberOfWeeklyTransitions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = NumberOfWeeklyTransitions(*nt)
	return br, err
}

func (v NumberOfWeeklyTransitions) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *NumberOfWeeklyTransitions) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = NumberOfWeeklyTransitions(*a)
	return nil
}

func (v *NumberOfWeeklyTransitions) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = NumberOfWeeklyTransitions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v NumberOfWeeklyTransitions) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const OccupancyAttr zcl.AttrID = 2

func (Occupancy) ID() zcl.AttrID   { return OccupancyAttr }
func (Occupancy) Readable() bool   { return true }
func (Occupancy) Writable() bool   { return false }
func (Occupancy) Reportable() bool { return false }
func (Occupancy) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Occupancy) AttrID() zcl.AttrID   { return v.ID() }
func (v Occupancy) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Occupancy) AttrValue() zcl.Val  { return v.Value() }

func (Occupancy) Name() string { return `Occupancy` }
func (Occupancy) Description() string {
	return `whether the heated/cooled space is occupied or not, as measured locally
or remotely (over the network)`
}

// Occupancy whether the heated/cooled space is occupied or not, as measured locally
// or remotely (over the network)
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
		case 1:
			bstr = append(bstr, "Occupied")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v Occupancy) IsOccupied() bool    { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v *Occupancy) SetOccupied(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }

func (Occupancy) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 1, Name: "Occupied"},
	}
}

const OccupiedCoolingSetpointAttr zcl.AttrID = 17

func (OccupiedCoolingSetpoint) ID() zcl.AttrID   { return OccupiedCoolingSetpointAttr }
func (OccupiedCoolingSetpoint) Readable() bool   { return true }
func (OccupiedCoolingSetpoint) Writable() bool   { return true }
func (OccupiedCoolingSetpoint) Reportable() bool { return false }
func (OccupiedCoolingSetpoint) SceneIndex() int  { return 1 }

// Implements AttrDef/AttrValue interfaces
func (v OccupiedCoolingSetpoint) AttrID() zcl.AttrID   { return v.ID() }
func (v OccupiedCoolingSetpoint) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OccupiedCoolingSetpoint) AttrValue() zcl.Val  { return v.Value() }

func (OccupiedCoolingSetpoint) Name() string { return `Occupied Cooling Setpoint` }
func (OccupiedCoolingSetpoint) Description() string {
	return `cooling mode setpoint when the room is occupied. The
OccupiedHeatingSetpoint attribute shall always be below the value
specified by at least MinSetpointDeadband`
}

// OccupiedCoolingSetpoint cooling mode setpoint when the room is occupied. The
// OccupiedHeatingSetpoint attribute shall always be below the value
// specified by at least MinSetpointDeadband
type OccupiedCoolingSetpoint zcl.Zs16

func (v *OccupiedCoolingSetpoint) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *OccupiedCoolingSetpoint) Value() zcl.Val     { return v }

func (v OccupiedCoolingSetpoint) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *OccupiedCoolingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = OccupiedCoolingSetpoint(*nt)
	return br, err
}

func (v OccupiedCoolingSetpoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *OccupiedCoolingSetpoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OccupiedCoolingSetpoint(*a)
	return nil
}

func (v *OccupiedCoolingSetpoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = OccupiedCoolingSetpoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OccupiedCoolingSetpoint) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v OccupiedCoolingSetpoint) Scaled() float64 {
	return float64(v) / 10
}

const OccupiedHeatingSetpointAttr zcl.AttrID = 18

func (OccupiedHeatingSetpoint) ID() zcl.AttrID   { return OccupiedHeatingSetpointAttr }
func (OccupiedHeatingSetpoint) Readable() bool   { return true }
func (OccupiedHeatingSetpoint) Writable() bool   { return true }
func (OccupiedHeatingSetpoint) Reportable() bool { return false }
func (OccupiedHeatingSetpoint) SceneIndex() int  { return 2 }

// Implements AttrDef/AttrValue interfaces
func (v OccupiedHeatingSetpoint) AttrID() zcl.AttrID   { return v.ID() }
func (v OccupiedHeatingSetpoint) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OccupiedHeatingSetpoint) AttrValue() zcl.Val  { return v.Value() }

func (OccupiedHeatingSetpoint) Name() string { return `Occupied Heating Setpoint` }
func (OccupiedHeatingSetpoint) Description() string {
	return `heating mode setpoint when the room is occupied. The
OccupiedCoolingSetpoint attribute shall always be above the value
specified by at least MinSetpointDeadband`
}

// OccupiedHeatingSetpoint heating mode setpoint when the room is occupied. The
// OccupiedCoolingSetpoint attribute shall always be above the value
// specified by at least MinSetpointDeadband
type OccupiedHeatingSetpoint zcl.Zs16

func (v *OccupiedHeatingSetpoint) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *OccupiedHeatingSetpoint) Value() zcl.Val     { return v }

func (v OccupiedHeatingSetpoint) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *OccupiedHeatingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = OccupiedHeatingSetpoint(*nt)
	return br, err
}

func (v OccupiedHeatingSetpoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *OccupiedHeatingSetpoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OccupiedHeatingSetpoint(*a)
	return nil
}

func (v *OccupiedHeatingSetpoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = OccupiedHeatingSetpoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OccupiedHeatingSetpoint) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v OccupiedHeatingSetpoint) Scaled() float64 {
	return float64(v) / 10
}

const OccupiedSetbackAttr zcl.AttrID = 52

func (OccupiedSetback) ID() zcl.AttrID   { return OccupiedSetbackAttr }
func (OccupiedSetback) Readable() bool   { return true }
func (OccupiedSetback) Writable() bool   { return true }
func (OccupiedSetback) Reportable() bool { return false }
func (OccupiedSetback) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OccupiedSetback) AttrID() zcl.AttrID   { return v.ID() }
func (v OccupiedSetback) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OccupiedSetback) AttrValue() zcl.Val  { return v.Value() }

func (OccupiedSetback) Name() string { return `Occupied Setback` }
func (OccupiedSetback) Description() string {
	return `degrees the thermostat will allow the LocalTemperature attribute to
float above the OccupiedCooling setpoint (i.e., OccupiedCooling +
OccupiedSetback) or below the OccupiedHeating setpoint (i.e.,
OccupiedHeating - OccupiedSetback) before initiating a state change to
bring the temperature back to the user's desired setpoint`
}

// OccupiedSetback degrees the thermostat will allow the LocalTemperature attribute to
// float above the OccupiedCooling setpoint (i.e., OccupiedCooling +
// OccupiedSetback) or below the OccupiedHeating setpoint (i.e.,
// OccupiedHeating - OccupiedSetback) before initiating a state change to
// bring the temperature back to the user's desired setpoint
type OccupiedSetback zcl.Zu8

func (v *OccupiedSetback) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *OccupiedSetback) Value() zcl.Val     { return v }

func (v OccupiedSetback) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *OccupiedSetback) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = OccupiedSetback(*nt)
	return br, err
}

func (v OccupiedSetback) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *OccupiedSetback) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OccupiedSetback(*a)
	return nil
}

func (v *OccupiedSetback) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = OccupiedSetback(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OccupiedSetback) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v OccupiedSetback) Scaled() float64 {
	return float64(v) / 10
}

const OccupiedSetbackMaxAttr zcl.AttrID = 54

func (OccupiedSetbackMax) ID() zcl.AttrID   { return OccupiedSetbackMaxAttr }
func (OccupiedSetbackMax) Readable() bool   { return true }
func (OccupiedSetbackMax) Writable() bool   { return false }
func (OccupiedSetbackMax) Reportable() bool { return false }
func (OccupiedSetbackMax) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OccupiedSetbackMax) AttrID() zcl.AttrID   { return v.ID() }
func (v OccupiedSetbackMax) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OccupiedSetbackMax) AttrValue() zcl.Val  { return v.Value() }

func (OccupiedSetbackMax) Name() string { return `Occupied Setback Max` }
func (OccupiedSetbackMax) Description() string {
	return `degrees the thermostat will allow the OccupiedSetback attribute to be
configured by a user`
}

// OccupiedSetbackMax degrees the thermostat will allow the OccupiedSetback attribute to be
// configured by a user
type OccupiedSetbackMax zcl.Zu8

func (v *OccupiedSetbackMax) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *OccupiedSetbackMax) Value() zcl.Val     { return v }

func (v OccupiedSetbackMax) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *OccupiedSetbackMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = OccupiedSetbackMax(*nt)
	return br, err
}

func (v OccupiedSetbackMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *OccupiedSetbackMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OccupiedSetbackMax(*a)
	return nil
}

func (v *OccupiedSetbackMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = OccupiedSetbackMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OccupiedSetbackMax) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v OccupiedSetbackMax) Scaled() float64 {
	return float64(v) / 10
}

const OccupiedSetbackMinAttr zcl.AttrID = 53

func (OccupiedSetbackMin) ID() zcl.AttrID   { return OccupiedSetbackMinAttr }
func (OccupiedSetbackMin) Readable() bool   { return true }
func (OccupiedSetbackMin) Writable() bool   { return false }
func (OccupiedSetbackMin) Reportable() bool { return false }
func (OccupiedSetbackMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OccupiedSetbackMin) AttrID() zcl.AttrID   { return v.ID() }
func (v OccupiedSetbackMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OccupiedSetbackMin) AttrValue() zcl.Val  { return v.Value() }

func (OccupiedSetbackMin) Name() string { return `Occupied Setback Min` }
func (OccupiedSetbackMin) Description() string {
	return `degrees the thermostat will allow the OccupiedSetback attribute to be
configured by a user`
}

// OccupiedSetbackMin degrees the thermostat will allow the OccupiedSetback attribute to be
// configured by a user
type OccupiedSetbackMin zcl.Zu8

func (v *OccupiedSetbackMin) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *OccupiedSetbackMin) Value() zcl.Val     { return v }

func (v OccupiedSetbackMin) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *OccupiedSetbackMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = OccupiedSetbackMin(*nt)
	return br, err
}

func (v OccupiedSetbackMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *OccupiedSetbackMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OccupiedSetbackMin(*a)
	return nil
}

func (v *OccupiedSetbackMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = OccupiedSetbackMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OccupiedSetbackMin) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v OccupiedSetbackMin) Scaled() float64 {
	return float64(v) / 10
}

const OperationModeAttr zcl.AttrID = 32

func (OperationMode) ID() zcl.AttrID   { return OperationModeAttr }
func (OperationMode) Readable() bool   { return true }
func (OperationMode) Writable() bool   { return true }
func (OperationMode) Reportable() bool { return false }
func (OperationMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OperationMode) AttrID() zcl.AttrID   { return v.ID() }
func (v OperationMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OperationMode) AttrValue() zcl.Val  { return v.Value() }

func (OperationMode) Name() string        { return `Operation Mode` }
func (OperationMode) Description() string { return `specifies the operation mode of the pump` }

// OperationMode specifies the operation mode of the pump
type OperationMode zcl.Zenum8

func (v *OperationMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *OperationMode) Value() zcl.Val     { return v }

func (v OperationMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *OperationMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = OperationMode(*nt)
	return br, err
}

func (v OperationMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *OperationMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OperationMode(*a)
	return nil
}

func (v *OperationMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = OperationMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OperationMode) String() string {
	switch v {
	case 0x00:
		return "Normal"
	case 0x01:
		return "Minimum"
	case 0x02:
		return "Maximum"
	case 0x03:
		return "Local"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v OperationMode) IsNormal() bool  { return v == 0x00 }
func (v OperationMode) IsMinimum() bool { return v == 0x01 }
func (v OperationMode) IsMaximum() bool { return v == 0x02 }
func (v OperationMode) IsLocal() bool   { return v == 0x03 }
func (v *OperationMode) SetNormal()     { *v = 0x00 }
func (v *OperationMode) SetMinimum()    { *v = 0x01 }
func (v *OperationMode) SetMaximum()    { *v = 0x02 }
func (v *OperationMode) SetLocal()      { *v = 0x03 }

func (OperationMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Normal"},
		{Value: 0x01, Name: "Minimum"},
		{Value: 0x02, Name: "Maximum"},
		{Value: 0x03, Name: "Local"},
	}
}

const OutdoorTemperatureAttr zcl.AttrID = 1

func (OutdoorTemperature) ID() zcl.AttrID   { return OutdoorTemperatureAttr }
func (OutdoorTemperature) Readable() bool   { return true }
func (OutdoorTemperature) Writable() bool   { return false }
func (OutdoorTemperature) Reportable() bool { return false }
func (OutdoorTemperature) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v OutdoorTemperature) AttrID() zcl.AttrID   { return v.ID() }
func (v OutdoorTemperature) AttrType() zcl.TypeID { return v.TypeID() }
func (v *OutdoorTemperature) AttrValue() zcl.Val  { return v.Value() }

func (OutdoorTemperature) Name() string { return `Outdoor Temperature` }
func (OutdoorTemperature) Description() string {
	return `outdoor temperature in degrees Celsius, as measured locally or remotely
(over the network). It is measured as described for LocalTemperature`
}

// OutdoorTemperature outdoor temperature in degrees Celsius, as measured locally or remotely
// (over the network). It is measured as described for LocalTemperature
type OutdoorTemperature zcl.Zs16

func (v *OutdoorTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *OutdoorTemperature) Value() zcl.Val     { return v }

func (v OutdoorTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *OutdoorTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = OutdoorTemperature(*nt)
	return br, err
}

func (v OutdoorTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *OutdoorTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = OutdoorTemperature(*a)
	return nil
}

func (v *OutdoorTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = OutdoorTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v OutdoorTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v OutdoorTemperature) Scaled() float64 {
	return float64(v) / 100
}

const PiCoolingDemandAttr zcl.AttrID = 7

func (PiCoolingDemand) ID() zcl.AttrID   { return PiCoolingDemandAttr }
func (PiCoolingDemand) Readable() bool   { return true }
func (PiCoolingDemand) Writable() bool   { return false }
func (PiCoolingDemand) Reportable() bool { return true }
func (PiCoolingDemand) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PiCoolingDemand) AttrID() zcl.AttrID   { return v.ID() }
func (v PiCoolingDemand) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PiCoolingDemand) AttrValue() zcl.Val  { return v.Value() }

func (PiCoolingDemand) Name() string { return `PI Cooling Demand` }
func (PiCoolingDemand) Description() string {
	return `specifies the level of cooling demanded by the PI (proportional
integral) control loop in use by the thermostat (if any), in
percent. This value is 0 when the thermostat is in off or heating mode`
}

// PiCoolingDemand specifies the level of cooling demanded by the PI (proportional
// integral) control loop in use by the thermostat (if any), in
// percent. This value is 0 when the thermostat is in off or heating mode
type PiCoolingDemand zcl.Zu8

func (v *PiCoolingDemand) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *PiCoolingDemand) Value() zcl.Val     { return v }

func (v PiCoolingDemand) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *PiCoolingDemand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = PiCoolingDemand(*nt)
	return br, err
}

func (v PiCoolingDemand) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *PiCoolingDemand) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PiCoolingDemand(*a)
	return nil
}

func (v *PiCoolingDemand) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = PiCoolingDemand(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PiCoolingDemand) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const PiHeatingDemandAttr zcl.AttrID = 8

func (PiHeatingDemand) ID() zcl.AttrID   { return PiHeatingDemandAttr }
func (PiHeatingDemand) Readable() bool   { return true }
func (PiHeatingDemand) Writable() bool   { return false }
func (PiHeatingDemand) Reportable() bool { return true }
func (PiHeatingDemand) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PiHeatingDemand) AttrID() zcl.AttrID   { return v.ID() }
func (v PiHeatingDemand) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PiHeatingDemand) AttrValue() zcl.Val  { return v.Value() }

func (PiHeatingDemand) Name() string { return `PI Heating Demand` }
func (PiHeatingDemand) Description() string {
	return `specifies the level of heating demanded by the PI (proportional
integral) control loop in use by the thermostat (if any), in
percent. This value is 0 when the thermostat is in off or cooling mode`
}

// PiHeatingDemand specifies the level of heating demanded by the PI (proportional
// integral) control loop in use by the thermostat (if any), in
// percent. This value is 0 when the thermostat is in off or cooling mode
type PiHeatingDemand zcl.Zu8

func (v *PiHeatingDemand) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *PiHeatingDemand) Value() zcl.Val     { return v }

func (v PiHeatingDemand) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *PiHeatingDemand) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = PiHeatingDemand(*nt)
	return br, err
}

func (v PiHeatingDemand) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *PiHeatingDemand) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PiHeatingDemand(*a)
	return nil
}

func (v *PiHeatingDemand) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = PiHeatingDemand(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PiHeatingDemand) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

const PowerAttr zcl.AttrID = 22

func (Power) ID() zcl.AttrID   { return PowerAttr }
func (Power) Readable() bool   { return true }
func (Power) Writable() bool   { return true }
func (Power) Reportable() bool { return false }
func (Power) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Power) AttrID() zcl.AttrID   { return v.ID() }
func (v Power) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Power) AttrValue() zcl.Val  { return v.Value() }

func (Power) Name() string        { return `Power` }
func (Power) Description() string { return `power consumption of the pump in Watts` }

// Power power consumption of the pump in Watts
type Power zcl.Zu24

func (v *Power) TypeID() zcl.TypeID { return new(zcl.Zu24).TypeID() }
func (v *Power) Value() zcl.Val     { return v }

func (v Power) MarshalZcl() ([]byte, error) { return zcl.Zu24(v).MarshalZcl() }

func (v *Power) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu24)
	br, err := nt.UnmarshalZcl(b)
	*v = Power(*nt)
	return br, err
}

func (v Power) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu24(v))
}

func (v *Power) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu24)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Power(*a)
	return nil
}

func (v *Power) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu24); ok {
		*v = Power(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Power) String() string {
	return zcl.Watts.Format(float64(v))
}

const PumpAlarmMaskAttr zcl.AttrID = 34

func (PumpAlarmMask) ID() zcl.AttrID   { return PumpAlarmMaskAttr }
func (PumpAlarmMask) Readable() bool   { return true }
func (PumpAlarmMask) Writable() bool   { return false }
func (PumpAlarmMask) Reportable() bool { return false }
func (PumpAlarmMask) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PumpAlarmMask) AttrID() zcl.AttrID   { return v.ID() }
func (v PumpAlarmMask) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PumpAlarmMask) AttrValue() zcl.Val  { return v.Value() }

func (PumpAlarmMask) Name() string { return `Pump Alarm Mask` }
func (PumpAlarmMask) Description() string {
	return `whether each of the alarms listed is enabled. When the bit number
corresponding to the alarm code is set to 1, the alarm is enabled, else
it is disabled`
}

// PumpAlarmMask whether each of the alarms listed is enabled. When the bit number
// corresponding to the alarm code is set to 1, the alarm is enabled, else
// it is disabled
type PumpAlarmMask zcl.Zbmp16

func (v *PumpAlarmMask) TypeID() zcl.TypeID { return new(zcl.Zbmp16).TypeID() }
func (v *PumpAlarmMask) Value() zcl.Val     { return v }

func (v PumpAlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(v).MarshalZcl() }

func (v *PumpAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*v = PumpAlarmMask(*nt)
	return br, err
}

func (v PumpAlarmMask) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp16(v))
}

func (v *PumpAlarmMask) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PumpAlarmMask(*a)
	return nil
}

func (v *PumpAlarmMask) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp16); ok {
		*v = PumpAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PumpAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Supply voltage too low")
		case 1:
			bstr = append(bstr, "Supply voltage too high")
		case 10:
			bstr = append(bstr, "Sensor failure")
		case 11:
			bstr = append(bstr, "Electronic non-fatal failure")
		case 12:
			bstr = append(bstr, "Electronic fatal failure")
		case 13:
			bstr = append(bstr, "General fault")
		case 2:
			bstr = append(bstr, "Power missing phase")
		case 3:
			bstr = append(bstr, "System pressure too low")
		case 4:
			bstr = append(bstr, "System pressure too high")
		case 5:
			bstr = append(bstr, "Dry running")
		case 6:
			bstr = append(bstr, "Motor temperature too high")
		case 7:
			bstr = append(bstr, "Pump motor has fatal failure")
		case 8:
			bstr = append(bstr, "Electronic temperature too high")
		case 9:
			bstr = append(bstr, "Pump blocked")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v PumpAlarmMask) IsSupplyVoltageTooLow() bool          { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v PumpAlarmMask) IsSupplyVoltageTooHigh() bool         { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v PumpAlarmMask) IsSensorFailure() bool                { return zcl.BitmapTest([]byte(v[:]), 10) }
func (v PumpAlarmMask) IsElectronicNonFatalFailure() bool    { return zcl.BitmapTest([]byte(v[:]), 11) }
func (v PumpAlarmMask) IsElectronicFatalFailure() bool       { return zcl.BitmapTest([]byte(v[:]), 12) }
func (v PumpAlarmMask) IsGeneralFault() bool                 { return zcl.BitmapTest([]byte(v[:]), 13) }
func (v PumpAlarmMask) IsPowerMissingPhase() bool            { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v PumpAlarmMask) IsSystemPressureTooLow() bool         { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v PumpAlarmMask) IsSystemPressureTooHigh() bool        { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v PumpAlarmMask) IsDryRunning() bool                   { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v PumpAlarmMask) IsMotorTemperatureTooHigh() bool      { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v PumpAlarmMask) IsPumpMotorHasFatalFailure() bool     { return zcl.BitmapTest([]byte(v[:]), 7) }
func (v PumpAlarmMask) IsElectronicTemperatureTooHigh() bool { return zcl.BitmapTest([]byte(v[:]), 8) }
func (v PumpAlarmMask) IsPumpBlocked() bool                  { return zcl.BitmapTest([]byte(v[:]), 9) }
func (v *PumpAlarmMask) SetSupplyVoltageTooLow(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *PumpAlarmMask) SetSupplyVoltageTooHigh(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *PumpAlarmMask) SetSensorFailure(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 10, b))
}
func (v *PumpAlarmMask) SetElectronicNonFatalFailure(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 11, b))
}
func (v *PumpAlarmMask) SetElectronicFatalFailure(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 12, b))
}
func (v *PumpAlarmMask) SetGeneralFault(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 13, b)) }
func (v *PumpAlarmMask) SetPowerMissingPhase(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *PumpAlarmMask) SetSystemPressureTooLow(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *PumpAlarmMask) SetSystemPressureTooHigh(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
}
func (v *PumpAlarmMask) SetDryRunning(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b)) }
func (v *PumpAlarmMask) SetMotorTemperatureTooHigh(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b))
}
func (v *PumpAlarmMask) SetPumpMotorHasFatalFailure(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 7, b))
}
func (v *PumpAlarmMask) SetElectronicTemperatureTooHigh(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 8, b))
}
func (v *PumpAlarmMask) SetPumpBlocked(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 9, b)) }

func (PumpAlarmMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Supply voltage too low"},
		{Value: 1, Name: "Supply voltage too high"},
		{Value: 10, Name: "Sensor failure"},
		{Value: 11, Name: "Electronic non-fatal failure"},
		{Value: 12, Name: "Electronic fatal failure"},
		{Value: 13, Name: "General fault"},
		{Value: 2, Name: "Power missing phase"},
		{Value: 3, Name: "System pressure too low"},
		{Value: 4, Name: "System pressure too high"},
		{Value: 5, Name: "Dry running"},
		{Value: 6, Name: "Motor temperature too high"},
		{Value: 7, Name: "Pump motor has fatal failure"},
		{Value: 8, Name: "Electronic temperature too high"},
		{Value: 9, Name: "Pump blocked"},
	}
}

const PumpStatusAttr zcl.AttrID = 16

func (PumpStatus) ID() zcl.AttrID   { return PumpStatusAttr }
func (PumpStatus) Readable() bool   { return true }
func (PumpStatus) Writable() bool   { return false }
func (PumpStatus) Reportable() bool { return true }
func (PumpStatus) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v PumpStatus) AttrID() zcl.AttrID   { return v.ID() }
func (v PumpStatus) AttrType() zcl.TypeID { return v.TypeID() }
func (v *PumpStatus) AttrValue() zcl.Val  { return v.Value() }

func (PumpStatus) Name() string        { return `Pump Status` }
func (PumpStatus) Description() string { return `` }

type PumpStatus zcl.Zbmp16

func (v *PumpStatus) TypeID() zcl.TypeID { return new(zcl.Zbmp16).TypeID() }
func (v *PumpStatus) Value() zcl.Val     { return v }

func (v PumpStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(v).MarshalZcl() }

func (v *PumpStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*v = PumpStatus(*nt)
	return br, err
}

func (v PumpStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp16(v))
}

func (v *PumpStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = PumpStatus(*a)
	return nil
}

func (v *PumpStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp16); ok {
		*v = PumpStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v PumpStatus) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Device fault")
		case 1:
			bstr = append(bstr, "Supply fault")
		case 2:
			bstr = append(bstr, "Speed low")
		case 3:
			bstr = append(bstr, "Speed high")
		case 4:
			bstr = append(bstr, "Local override")
		case 5:
			bstr = append(bstr, "Running")
		case 6:
			bstr = append(bstr, "Remote pressure")
		case 7:
			bstr = append(bstr, "Remote flow")
		case 8:
			bstr = append(bstr, "Remote temperature")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v PumpStatus) IsDeviceFault() bool       { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v PumpStatus) IsSupplyFault() bool       { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v PumpStatus) IsSpeedLow() bool          { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v PumpStatus) IsSpeedHigh() bool         { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v PumpStatus) IsLocalOverride() bool     { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v PumpStatus) IsRunning() bool           { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v PumpStatus) IsRemotePressure() bool    { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v PumpStatus) IsRemoteFlow() bool        { return zcl.BitmapTest([]byte(v[:]), 7) }
func (v PumpStatus) IsRemoteTemperature() bool { return zcl.BitmapTest([]byte(v[:]), 8) }
func (v *PumpStatus) SetDeviceFault(b bool)    { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *PumpStatus) SetSupplyFault(b bool)    { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }
func (v *PumpStatus) SetSpeedLow(b bool)       { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b)) }
func (v *PumpStatus) SetSpeedHigh(b bool)      { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b)) }
func (v *PumpStatus) SetLocalOverride(b bool)  { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b)) }
func (v *PumpStatus) SetRunning(b bool)        { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b)) }
func (v *PumpStatus) SetRemotePressure(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b)) }
func (v *PumpStatus) SetRemoteFlow(b bool)     { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 7, b)) }
func (v *PumpStatus) SetRemoteTemperature(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 8, b))
}

func (PumpStatus) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Device fault"},
		{Value: 1, Name: "Supply fault"},
		{Value: 2, Name: "Speed low"},
		{Value: 3, Name: "Speed high"},
		{Value: 4, Name: "Local override"},
		{Value: 5, Name: "Running"},
		{Value: 6, Name: "Remote pressure"},
		{Value: 7, Name: "Remote flow"},
		{Value: 8, Name: "Remote temperature"},
	}
}

const RhDehumidificationSetpointAttr zcl.AttrID = 16

func (RhDehumidificationSetpoint) ID() zcl.AttrID   { return RhDehumidificationSetpointAttr }
func (RhDehumidificationSetpoint) Readable() bool   { return true }
func (RhDehumidificationSetpoint) Writable() bool   { return true }
func (RhDehumidificationSetpoint) Reportable() bool { return false }
func (RhDehumidificationSetpoint) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RhDehumidificationSetpoint) AttrID() zcl.AttrID   { return v.ID() }
func (v RhDehumidificationSetpoint) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RhDehumidificationSetpoint) AttrValue() zcl.Val  { return v.Value() }

func (RhDehumidificationSetpoint) Name() string { return `RH Dehumidification Setpoint` }
func (RhDehumidificationSetpoint) Description() string {
	return `relative humidity at which dehumidification occurs`
}

// RhDehumidificationSetpoint relative humidity at which dehumidification occurs
type RhDehumidificationSetpoint zcl.Zu8

func (v *RhDehumidificationSetpoint) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *RhDehumidificationSetpoint) Value() zcl.Val     { return v }

func (v RhDehumidificationSetpoint) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *RhDehumidificationSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = RhDehumidificationSetpoint(*nt)
	return br, err
}

func (v RhDehumidificationSetpoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *RhDehumidificationSetpoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RhDehumidificationSetpoint(*a)
	return nil
}

func (v *RhDehumidificationSetpoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = RhDehumidificationSetpoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RhDehumidificationSetpoint) String() string {
	return zcl.Percent.Format(float64(v))
}

const RelativeHumidityAttr zcl.AttrID = 0

func (RelativeHumidity) ID() zcl.AttrID   { return RelativeHumidityAttr }
func (RelativeHumidity) Readable() bool   { return true }
func (RelativeHumidity) Writable() bool   { return false }
func (RelativeHumidity) Reportable() bool { return false }
func (RelativeHumidity) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RelativeHumidity) AttrID() zcl.AttrID   { return v.ID() }
func (v RelativeHumidity) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RelativeHumidity) AttrValue() zcl.Val  { return v.Value() }

func (RelativeHumidity) Name() string { return `Relative Humidity` }
func (RelativeHumidity) Description() string {
	return `current relative humidity measured by a local or remote sensor`
}

// RelativeHumidity current relative humidity measured by a local or remote sensor
type RelativeHumidity zcl.Zu8

func (v *RelativeHumidity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *RelativeHumidity) Value() zcl.Val     { return v }

func (v RelativeHumidity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *RelativeHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = RelativeHumidity(*nt)
	return br, err
}

func (v RelativeHumidity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *RelativeHumidity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RelativeHumidity(*a)
	return nil
}

func (v *RelativeHumidity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = RelativeHumidity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RelativeHumidity) String() string {
	return zcl.Percent.Format(float64(v))
}

const RelativeHumidityDisplayAttr zcl.AttrID = 21

func (RelativeHumidityDisplay) ID() zcl.AttrID   { return RelativeHumidityDisplayAttr }
func (RelativeHumidityDisplay) Readable() bool   { return true }
func (RelativeHumidityDisplay) Writable() bool   { return true }
func (RelativeHumidityDisplay) Reportable() bool { return false }
func (RelativeHumidityDisplay) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RelativeHumidityDisplay) AttrID() zcl.AttrID   { return v.ID() }
func (v RelativeHumidityDisplay) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RelativeHumidityDisplay) AttrValue() zcl.Val  { return v.Value() }

func (RelativeHumidityDisplay) Name() string { return `Relative Humidity Display` }
func (RelativeHumidityDisplay) Description() string {
	return `whether the RH value is displayed to the user or not`
}

// RelativeHumidityDisplay whether the RH value is displayed to the user or not
type RelativeHumidityDisplay zcl.Zenum8

func (v *RelativeHumidityDisplay) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *RelativeHumidityDisplay) Value() zcl.Val     { return v }

func (v RelativeHumidityDisplay) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *RelativeHumidityDisplay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = RelativeHumidityDisplay(*nt)
	return br, err
}

func (v RelativeHumidityDisplay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *RelativeHumidityDisplay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RelativeHumidityDisplay(*a)
	return nil
}

func (v *RelativeHumidityDisplay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = RelativeHumidityDisplay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RelativeHumidityDisplay) String() string {
	switch v {
	case 0x00:
		return "Not displayed"
	case 0x01:
		return "Displayed"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v RelativeHumidityDisplay) IsNotDisplayed() bool { return v == 0x00 }
func (v RelativeHumidityDisplay) IsDisplayed() bool    { return v == 0x01 }
func (v *RelativeHumidityDisplay) SetNotDisplayed()    { *v = 0x00 }
func (v *RelativeHumidityDisplay) SetDisplayed()       { *v = 0x01 }

func (RelativeHumidityDisplay) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Not displayed"},
		{Value: 0x01, Name: "Displayed"},
	}
}

const RelativeHumidityModeAttr zcl.AttrID = 17

func (RelativeHumidityMode) ID() zcl.AttrID   { return RelativeHumidityModeAttr }
func (RelativeHumidityMode) Readable() bool   { return true }
func (RelativeHumidityMode) Writable() bool   { return true }
func (RelativeHumidityMode) Reportable() bool { return false }
func (RelativeHumidityMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RelativeHumidityMode) AttrID() zcl.AttrID   { return v.ID() }
func (v RelativeHumidityMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RelativeHumidityMode) AttrValue() zcl.Val  { return v.Value() }

func (RelativeHumidityMode) Name() string { return `Relative Humidity Mode` }
func (RelativeHumidityMode) Description() string {
	return `how the RelativeHumidity value is being updated`
}

// RelativeHumidityMode how the RelativeHumidity value is being updated
type RelativeHumidityMode zcl.Zenum8

func (v *RelativeHumidityMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *RelativeHumidityMode) Value() zcl.Val     { return v }

func (v RelativeHumidityMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *RelativeHumidityMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = RelativeHumidityMode(*nt)
	return br, err
}

func (v RelativeHumidityMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *RelativeHumidityMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RelativeHumidityMode(*a)
	return nil
}

func (v *RelativeHumidityMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = RelativeHumidityMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RelativeHumidityMode) String() string {
	switch v {
	case 0x00:
		return "Locally"
	case 0x01:
		return "Remotely"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v RelativeHumidityMode) IsLocally() bool  { return v == 0x00 }
func (v RelativeHumidityMode) IsRemotely() bool { return v == 0x01 }
func (v *RelativeHumidityMode) SetLocally()     { *v = 0x00 }
func (v *RelativeHumidityMode) SetRemotely()    { *v = 0x01 }

func (RelativeHumidityMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Locally"},
		{Value: 0x01, Name: "Remotely"},
	}
}

func (RelayStatus) Name() string { return `Relay Status` }
func (RelayStatus) Description() string {
	return `status for thermostat when the log is captured. Each bit represents one
relay used by the thermostat. If the bit is on, the associated relay is
on and active. Each thermostat manufacturer can create its own mapping
between the bitmask and the associated relay`
}

// RelayStatus status for thermostat when the log is captured. Each bit represents one
// relay used by the thermostat. If the bit is on, the associated relay is
// on and active. Each thermostat manufacturer can create its own mapping
// between the bitmask and the associated relay
type RelayStatus zcl.Zbmp8

func (v *RelayStatus) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *RelayStatus) Value() zcl.Val     { return v }

func (v RelayStatus) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *RelayStatus) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = RelayStatus(*nt)
	return br, err
}

func (v RelayStatus) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *RelayStatus) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RelayStatus(*a)
	return nil
}

func (v *RelayStatus) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = RelayStatus(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RelayStatus) String() string {
	return zcl.Sprintf("%v", zcl.Zbmp8(v))
}

func (RelayStatusHumidity) Name() string        { return `Relay Status Humidity` }
func (RelayStatusHumidity) Description() string { return `humidity when the log was captured` }

// RelayStatusHumidity humidity when the log was captured
type RelayStatusHumidity zcl.Zu8

func (v *RelayStatusHumidity) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *RelayStatusHumidity) Value() zcl.Val     { return v }

func (v RelayStatusHumidity) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *RelayStatusHumidity) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = RelayStatusHumidity(*nt)
	return br, err
}

func (v RelayStatusHumidity) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *RelayStatusHumidity) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RelayStatusHumidity(*a)
	return nil
}

func (v *RelayStatusHumidity) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = RelayStatusHumidity(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RelayStatusHumidity) String() string {
	return zcl.Percent.Format(float64(v))
}

func (RelayStatusLocalTemperature) Name() string { return `Relay Status Local Temperature` }
func (RelayStatusLocalTemperature) Description() string {
	return `temperature when the log is captured`
}

// RelayStatusLocalTemperature temperature when the log is captured
type RelayStatusLocalTemperature zcl.Zs16

func (v *RelayStatusLocalTemperature) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *RelayStatusLocalTemperature) Value() zcl.Val     { return v }

func (v RelayStatusLocalTemperature) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *RelayStatusLocalTemperature) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = RelayStatusLocalTemperature(*nt)
	return br, err
}

func (v RelayStatusLocalTemperature) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *RelayStatusLocalTemperature) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RelayStatusLocalTemperature(*a)
	return nil
}

func (v *RelayStatusLocalTemperature) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = RelayStatusLocalTemperature(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RelayStatusLocalTemperature) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v RelayStatusLocalTemperature) Scaled() float64 {
	return float64(v) / 100
}

func (RelayStatusLogTimeOfDay) Name() string { return `Relay Status Log Time of Day` }
func (RelayStatusLogTimeOfDay) Description() string {
	return `minutes since midnight when the relay status was captured for this
associated log entry`
}

// RelayStatusLogTimeOfDay minutes since midnight when the relay status was captured for this
// associated log entry
type RelayStatusLogTimeOfDay zcl.Zu16

func (v *RelayStatusLogTimeOfDay) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RelayStatusLogTimeOfDay) Value() zcl.Val     { return v }

func (v RelayStatusLogTimeOfDay) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RelayStatusLogTimeOfDay) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RelayStatusLogTimeOfDay(*nt)
	return br, err
}

func (v RelayStatusLogTimeOfDay) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RelayStatusLogTimeOfDay) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RelayStatusLogTimeOfDay(*a)
	return nil
}

func (v *RelayStatusLogTimeOfDay) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RelayStatusLogTimeOfDay(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RelayStatusLogTimeOfDay) String() string {
	return zcl.Minutes.Format(float64(v))
}

func (RelayStatusSetpoint) Name() string { return `Relay Status Setpoint` }
func (RelayStatusSetpoint) Description() string {
	return `target setpoint temperature when the log is captured`
}

// RelayStatusSetpoint target setpoint temperature when the log is captured
type RelayStatusSetpoint zcl.Zs16

func (v *RelayStatusSetpoint) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *RelayStatusSetpoint) Value() zcl.Val     { return v }

func (v RelayStatusSetpoint) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *RelayStatusSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = RelayStatusSetpoint(*nt)
	return br, err
}

func (v RelayStatusSetpoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *RelayStatusSetpoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RelayStatusSetpoint(*a)
	return nil
}

func (v *RelayStatusSetpoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = RelayStatusSetpoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RelayStatusSetpoint) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v RelayStatusSetpoint) Scaled() float64 {
	return float64(v) / 100
}

func (RelayStatusUnreadEntries) Name() string { return `Relay status Unread Entries` }
func (RelayStatusUnreadEntries) Description() string {
	return `umber of unread entries within the thermostat log system`
}

// RelayStatusUnreadEntries umber of unread entries within the thermostat log system
type RelayStatusUnreadEntries zcl.Zu16

func (v *RelayStatusUnreadEntries) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *RelayStatusUnreadEntries) Value() zcl.Val     { return v }

func (v RelayStatusUnreadEntries) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *RelayStatusUnreadEntries) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = RelayStatusUnreadEntries(*nt)
	return br, err
}

func (v RelayStatusUnreadEntries) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *RelayStatusUnreadEntries) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RelayStatusUnreadEntries(*a)
	return nil
}

func (v *RelayStatusUnreadEntries) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = RelayStatusUnreadEntries(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RelayStatusUnreadEntries) String() string {
	return zcl.Sprintf("%v", zcl.Zu16(v))
}

const RemoteSensingAttr zcl.AttrID = 26

func (RemoteSensing) ID() zcl.AttrID   { return RemoteSensingAttr }
func (RemoteSensing) Readable() bool   { return true }
func (RemoteSensing) Writable() bool   { return true }
func (RemoteSensing) Reportable() bool { return false }
func (RemoteSensing) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v RemoteSensing) AttrID() zcl.AttrID   { return v.ID() }
func (v RemoteSensing) AttrType() zcl.TypeID { return v.TypeID() }
func (v *RemoteSensing) AttrValue() zcl.Val  { return v.Value() }

func (RemoteSensing) Name() string { return `Remote Sensing` }
func (RemoteSensing) Description() string {
	return `specifies whether the local temperature, outdoor temperature and
occupancy are being sensed by internal sensors or remote networked
sensors. When the bit is 0 it means internal, 1 means remote`
}

// RemoteSensing specifies whether the local temperature, outdoor temperature and
// occupancy are being sensed by internal sensors or remote networked
// sensors. When the bit is 0 it means internal, 1 means remote
type RemoteSensing zcl.Zbmp8

func (v *RemoteSensing) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *RemoteSensing) Value() zcl.Val     { return v }

func (v RemoteSensing) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *RemoteSensing) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = RemoteSensing(*nt)
	return br, err
}

func (v RemoteSensing) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *RemoteSensing) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = RemoteSensing(*a)
	return nil
}

func (v *RemoteSensing) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = RemoteSensing(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v RemoteSensing) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Local temperature")
		case 1:
			bstr = append(bstr, "Outdoor temperature")
		case 2:
			bstr = append(bstr, "Occupancy")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v RemoteSensing) IsLocalTemperature() bool   { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v RemoteSensing) IsOutdoorTemperature() bool { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v RemoteSensing) IsOccupancy() bool          { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v *RemoteSensing) SetLocalTemperature(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *RemoteSensing) SetOutdoorTemperature(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *RemoteSensing) SetOccupancy(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b)) }

func (RemoteSensing) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Local temperature"},
		{Value: 1, Name: "Outdoor temperature"},
		{Value: 2, Name: "Occupancy"},
	}
}

const ScheduleProgrammingVisibilityAttr zcl.AttrID = 2

func (ScheduleProgrammingVisibility) ID() zcl.AttrID   { return ScheduleProgrammingVisibilityAttr }
func (ScheduleProgrammingVisibility) Readable() bool   { return true }
func (ScheduleProgrammingVisibility) Writable() bool   { return true }
func (ScheduleProgrammingVisibility) Reportable() bool { return false }
func (ScheduleProgrammingVisibility) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ScheduleProgrammingVisibility) AttrID() zcl.AttrID   { return v.ID() }
func (v ScheduleProgrammingVisibility) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ScheduleProgrammingVisibility) AttrValue() zcl.Val  { return v.Value() }

func (ScheduleProgrammingVisibility) Name() string { return `Schedule Programming Visibility` }
func (ScheduleProgrammingVisibility) Description() string {
	return `hide the weekly schedule programming functionality or menu on a
thermostat from a user to prevent local user programming of the weekly
schedule`
}

// ScheduleProgrammingVisibility hide the weekly schedule programming functionality or menu on a
// thermostat from a user to prevent local user programming of the weekly
// schedule
type ScheduleProgrammingVisibility zcl.Zenum8

func (v *ScheduleProgrammingVisibility) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *ScheduleProgrammingVisibility) Value() zcl.Val     { return v }

func (v ScheduleProgrammingVisibility) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(v).MarshalZcl()
}

func (v *ScheduleProgrammingVisibility) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = ScheduleProgrammingVisibility(*nt)
	return br, err
}

func (v ScheduleProgrammingVisibility) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *ScheduleProgrammingVisibility) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ScheduleProgrammingVisibility(*a)
	return nil
}

func (v *ScheduleProgrammingVisibility) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = ScheduleProgrammingVisibility(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ScheduleProgrammingVisibility) String() string {
	switch v {
	case 0x00:
		return "Not hidden"
	case 0x01:
		return "Hidden"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v ScheduleProgrammingVisibility) IsNotHidden() bool { return v == 0x00 }
func (v ScheduleProgrammingVisibility) IsHidden() bool    { return v == 0x01 }
func (v *ScheduleProgrammingVisibility) SetNotHidden()    { *v = 0x00 }
func (v *ScheduleProgrammingVisibility) SetHidden()       { *v = 0x01 }

func (ScheduleProgrammingVisibility) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Not hidden"},
		{Value: 0x01, Name: "Hidden"},
	}
}

func (SetWeeklyCoolSetpoint1) Name() string        { return `Set Weekly Cool Setpoint 1` }
func (SetWeeklyCoolSetpoint1) Description() string { return `` }

type SetWeeklyCoolSetpoint1 zcl.Zs16

func (v *SetWeeklyCoolSetpoint1) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *SetWeeklyCoolSetpoint1) Value() zcl.Val     { return v }

func (v SetWeeklyCoolSetpoint1) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *SetWeeklyCoolSetpoint1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = SetWeeklyCoolSetpoint1(*nt)
	return br, err
}

func (v SetWeeklyCoolSetpoint1) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *SetWeeklyCoolSetpoint1) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetWeeklyCoolSetpoint1(*a)
	return nil
}

func (v *SetWeeklyCoolSetpoint1) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = SetWeeklyCoolSetpoint1(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetWeeklyCoolSetpoint1) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v SetWeeklyCoolSetpoint1) Scaled() float64 {
	return float64(v) / 100
}

func (SetWeeklyCoolSetpoint10) Name() string        { return `Set Weekly Cool Setpoint 10` }
func (SetWeeklyCoolSetpoint10) Description() string { return `` }

type SetWeeklyCoolSetpoint10 zcl.Zs16

func (v *SetWeeklyCoolSetpoint10) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *SetWeeklyCoolSetpoint10) Value() zcl.Val     { return v }

func (v SetWeeklyCoolSetpoint10) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *SetWeeklyCoolSetpoint10) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = SetWeeklyCoolSetpoint10(*nt)
	return br, err
}

func (v SetWeeklyCoolSetpoint10) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *SetWeeklyCoolSetpoint10) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetWeeklyCoolSetpoint10(*a)
	return nil
}

func (v *SetWeeklyCoolSetpoint10) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = SetWeeklyCoolSetpoint10(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetWeeklyCoolSetpoint10) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v SetWeeklyCoolSetpoint10) Scaled() float64 {
	return float64(v) / 100
}

func (SetWeeklyDayOfWeek) Name() string { return `Set Weekly Day Of Week` }
func (SetWeeklyDayOfWeek) Description() string {
	return `day of the week at which all the transitions within the payload of the
command should be associated to`
}

// SetWeeklyDayOfWeek day of the week at which all the transitions within the payload of the
// command should be associated to
type SetWeeklyDayOfWeek zcl.Zbmp8

func (v *SetWeeklyDayOfWeek) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *SetWeeklyDayOfWeek) Value() zcl.Val     { return v }

func (v SetWeeklyDayOfWeek) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *SetWeeklyDayOfWeek) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = SetWeeklyDayOfWeek(*nt)
	return br, err
}

func (v SetWeeklyDayOfWeek) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *SetWeeklyDayOfWeek) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetWeeklyDayOfWeek(*a)
	return nil
}

func (v *SetWeeklyDayOfWeek) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = SetWeeklyDayOfWeek(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetWeeklyDayOfWeek) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Sunday")
		case 1:
			bstr = append(bstr, "Monday")
		case 2:
			bstr = append(bstr, "Tuesday")
		case 3:
			bstr = append(bstr, "Wednesday")
		case 4:
			bstr = append(bstr, "Thursday")
		case 5:
			bstr = append(bstr, "Friday")
		case 6:
			bstr = append(bstr, "Saturday")
		case 7:
			bstr = append(bstr, "Away or vacation")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v SetWeeklyDayOfWeek) IsSunday() bool         { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v SetWeeklyDayOfWeek) IsMonday() bool         { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v SetWeeklyDayOfWeek) IsTuesday() bool        { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v SetWeeklyDayOfWeek) IsWednesday() bool      { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v SetWeeklyDayOfWeek) IsThursday() bool       { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v SetWeeklyDayOfWeek) IsFriday() bool         { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v SetWeeklyDayOfWeek) IsSaturday() bool       { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v SetWeeklyDayOfWeek) IsAwayOrVacation() bool { return zcl.BitmapTest([]byte(v[:]), 7) }
func (v *SetWeeklyDayOfWeek) SetSunday(b bool)      { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *SetWeeklyDayOfWeek) SetMonday(b bool)      { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }
func (v *SetWeeklyDayOfWeek) SetTuesday(b bool)     { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b)) }
func (v *SetWeeklyDayOfWeek) SetWednesday(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *SetWeeklyDayOfWeek) SetThursday(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b)) }
func (v *SetWeeklyDayOfWeek) SetFriday(b bool)   { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b)) }
func (v *SetWeeklyDayOfWeek) SetSaturday(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b)) }
func (v *SetWeeklyDayOfWeek) SetAwayOrVacation(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 7, b))
}

func (SetWeeklyDayOfWeek) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Sunday"},
		{Value: 1, Name: "Monday"},
		{Value: 2, Name: "Tuesday"},
		{Value: 3, Name: "Wednesday"},
		{Value: 4, Name: "Thursday"},
		{Value: 5, Name: "Friday"},
		{Value: 6, Name: "Saturday"},
		{Value: 7, Name: "Away or vacation"},
	}
}

func (SetWeeklyHeatSetpoint1) Name() string        { return `Set Weekly Heat Setpoint 1` }
func (SetWeeklyHeatSetpoint1) Description() string { return `` }

type SetWeeklyHeatSetpoint1 zcl.Zs16

func (v *SetWeeklyHeatSetpoint1) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *SetWeeklyHeatSetpoint1) Value() zcl.Val     { return v }

func (v SetWeeklyHeatSetpoint1) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *SetWeeklyHeatSetpoint1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = SetWeeklyHeatSetpoint1(*nt)
	return br, err
}

func (v SetWeeklyHeatSetpoint1) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *SetWeeklyHeatSetpoint1) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetWeeklyHeatSetpoint1(*a)
	return nil
}

func (v *SetWeeklyHeatSetpoint1) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = SetWeeklyHeatSetpoint1(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetWeeklyHeatSetpoint1) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v SetWeeklyHeatSetpoint1) Scaled() float64 {
	return float64(v) / 100
}

func (SetWeeklyHeatSetpoint10) Name() string        { return `Set Weekly Heat Setpoint 10` }
func (SetWeeklyHeatSetpoint10) Description() string { return `` }

type SetWeeklyHeatSetpoint10 zcl.Zs16

func (v *SetWeeklyHeatSetpoint10) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *SetWeeklyHeatSetpoint10) Value() zcl.Val     { return v }

func (v SetWeeklyHeatSetpoint10) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *SetWeeklyHeatSetpoint10) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = SetWeeklyHeatSetpoint10(*nt)
	return br, err
}

func (v SetWeeklyHeatSetpoint10) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *SetWeeklyHeatSetpoint10) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetWeeklyHeatSetpoint10(*a)
	return nil
}

func (v *SetWeeklyHeatSetpoint10) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = SetWeeklyHeatSetpoint10(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetWeeklyHeatSetpoint10) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v SetWeeklyHeatSetpoint10) Scaled() float64 {
	return float64(v) / 100
}

func (SetWeeklyMode) Name() string { return `Set Weekly Mode` }
func (SetWeeklyMode) Description() string {
	return `which type of setpoint transition is present in the rest of the
command`
}

// SetWeeklyMode which type of setpoint transition is present in the rest of the
// command
type SetWeeklyMode zcl.Zbmp8

func (v *SetWeeklyMode) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *SetWeeklyMode) Value() zcl.Val     { return v }

func (v SetWeeklyMode) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *SetWeeklyMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = SetWeeklyMode(*nt)
	return br, err
}

func (v SetWeeklyMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *SetWeeklyMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetWeeklyMode(*a)
	return nil
}

func (v *SetWeeklyMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = SetWeeklyMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetWeeklyMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Heat setpoint")
		case 1:
			bstr = append(bstr, "Cool setpoint")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v SetWeeklyMode) IsHeatSetpoint() bool    { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v SetWeeklyMode) IsCoolSetpoint() bool    { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v *SetWeeklyMode) SetHeatSetpoint(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b)) }
func (v *SetWeeklyMode) SetCoolSetpoint(b bool) { copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b)) }

func (SetWeeklyMode) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Heat setpoint"},
		{Value: 1, Name: "Cool setpoint"},
	}
}

func (SetWeeklyNumberOfTransitions) Name() string { return `Set Weekly Number Of Transitions` }
func (SetWeeklyNumberOfTransitions) Description() string {
	return `how many individual transitions to expect for this sequence of
commands. If a device supports more than 10 transitions in its
schedule they can send this by sending more than 1 Set Weekly Schedule
command, each containing the separate information that the device
needs to set`
}

// SetWeeklyNumberOfTransitions how many individual transitions to expect for this sequence of
// commands. If a device supports more than 10 transitions in its
// schedule they can send this by sending more than 1 Set Weekly Schedule
// command, each containing the separate information that the device
// needs to set
type SetWeeklyNumberOfTransitions zcl.Zu8

func (v *SetWeeklyNumberOfTransitions) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *SetWeeklyNumberOfTransitions) Value() zcl.Val     { return v }

func (v SetWeeklyNumberOfTransitions) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *SetWeeklyNumberOfTransitions) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = SetWeeklyNumberOfTransitions(*nt)
	return br, err
}

func (v SetWeeklyNumberOfTransitions) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *SetWeeklyNumberOfTransitions) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetWeeklyNumberOfTransitions(*a)
	return nil
}

func (v *SetWeeklyNumberOfTransitions) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = SetWeeklyNumberOfTransitions(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetWeeklyNumberOfTransitions) String() string {
	return zcl.Sprintf("%v", zcl.Zu8(v))
}

func (SetWeeklyTransitionTime1) Name() string        { return `Set Weekly Transition Time 1` }
func (SetWeeklyTransitionTime1) Description() string { return `` }

type SetWeeklyTransitionTime1 zcl.Zu16

func (v *SetWeeklyTransitionTime1) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *SetWeeklyTransitionTime1) Value() zcl.Val     { return v }

func (v SetWeeklyTransitionTime1) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *SetWeeklyTransitionTime1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = SetWeeklyTransitionTime1(*nt)
	return br, err
}

func (v SetWeeklyTransitionTime1) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *SetWeeklyTransitionTime1) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetWeeklyTransitionTime1(*a)
	return nil
}

func (v *SetWeeklyTransitionTime1) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = SetWeeklyTransitionTime1(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetWeeklyTransitionTime1) String() string {
	return zcl.Minutes.Format(float64(v))
}

func (SetWeeklyTransitionTime10) Name() string        { return `Set Weekly Transition Time 10` }
func (SetWeeklyTransitionTime10) Description() string { return `` }

type SetWeeklyTransitionTime10 zcl.Zu16

func (v *SetWeeklyTransitionTime10) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *SetWeeklyTransitionTime10) Value() zcl.Val     { return v }

func (v SetWeeklyTransitionTime10) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *SetWeeklyTransitionTime10) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = SetWeeklyTransitionTime10(*nt)
	return br, err
}

func (v SetWeeklyTransitionTime10) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *SetWeeklyTransitionTime10) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetWeeklyTransitionTime10(*a)
	return nil
}

func (v *SetWeeklyTransitionTime10) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = SetWeeklyTransitionTime10(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetWeeklyTransitionTime10) String() string {
	return zcl.Minutes.Format(float64(v))
}

func (SetpointAmount) Name() string        { return `Setpoint Amount` }
func (SetpointAmount) Description() string { return `` }

type SetpointAmount zcl.Zs8

func (v *SetpointAmount) TypeID() zcl.TypeID { return new(zcl.Zs8).TypeID() }
func (v *SetpointAmount) Value() zcl.Val     { return v }

func (v SetpointAmount) MarshalZcl() ([]byte, error) { return zcl.Zs8(v).MarshalZcl() }

func (v *SetpointAmount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs8)
	br, err := nt.UnmarshalZcl(b)
	*v = SetpointAmount(*nt)
	return br, err
}

func (v SetpointAmount) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs8(v))
}

func (v *SetpointAmount) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetpointAmount(*a)
	return nil
}

func (v *SetpointAmount) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs8); ok {
		*v = SetpointAmount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetpointAmount) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v SetpointAmount) Scaled() float64 {
	return float64(v) / 10
}

const SetpointChangeAmountAttr zcl.AttrID = 49

func (SetpointChangeAmount) ID() zcl.AttrID   { return SetpointChangeAmountAttr }
func (SetpointChangeAmount) Readable() bool   { return true }
func (SetpointChangeAmount) Writable() bool   { return false }
func (SetpointChangeAmount) Reportable() bool { return false }
func (SetpointChangeAmount) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SetpointChangeAmount) AttrID() zcl.AttrID   { return v.ID() }
func (v SetpointChangeAmount) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SetpointChangeAmount) AttrValue() zcl.Val  { return v.Value() }

func (SetpointChangeAmount) Name() string { return `Setpoint Change Amount` }
func (SetpointChangeAmount) Description() string {
	return `delta between the current active OccupiedCoolingSetpoint or
OccupiedHeatingSetpoint and the previous active setpoint`
}

// SetpointChangeAmount delta between the current active OccupiedCoolingSetpoint or
// OccupiedHeatingSetpoint and the previous active setpoint
type SetpointChangeAmount zcl.Zs16

func (v *SetpointChangeAmount) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *SetpointChangeAmount) Value() zcl.Val     { return v }

func (v SetpointChangeAmount) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *SetpointChangeAmount) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = SetpointChangeAmount(*nt)
	return br, err
}

func (v SetpointChangeAmount) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *SetpointChangeAmount) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetpointChangeAmount(*a)
	return nil
}

func (v *SetpointChangeAmount) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = SetpointChangeAmount(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetpointChangeAmount) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 100)
}

func (v SetpointChangeAmount) Scaled() float64 {
	return float64(v) / 100
}

const SetpointChangeSourceAttr zcl.AttrID = 48

func (SetpointChangeSource) ID() zcl.AttrID   { return SetpointChangeSourceAttr }
func (SetpointChangeSource) Readable() bool   { return true }
func (SetpointChangeSource) Writable() bool   { return false }
func (SetpointChangeSource) Reportable() bool { return false }
func (SetpointChangeSource) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SetpointChangeSource) AttrID() zcl.AttrID   { return v.ID() }
func (v SetpointChangeSource) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SetpointChangeSource) AttrValue() zcl.Val  { return v.Value() }

func (SetpointChangeSource) Name() string { return `Setpoint Change Source` }
func (SetpointChangeSource) Description() string {
	return `source of the current active OccupiedCoolingSetpoint or
OccupiedHeatingSetpoint (i.e., who or what determined the current
setpoint)`
}

// SetpointChangeSource source of the current active OccupiedCoolingSetpoint or
// OccupiedHeatingSetpoint (i.e., who or what determined the current
// setpoint)
type SetpointChangeSource zcl.Zenum8

func (v *SetpointChangeSource) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *SetpointChangeSource) Value() zcl.Val     { return v }

func (v SetpointChangeSource) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *SetpointChangeSource) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = SetpointChangeSource(*nt)
	return br, err
}

func (v SetpointChangeSource) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *SetpointChangeSource) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetpointChangeSource(*a)
	return nil
}

func (v *SetpointChangeSource) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = SetpointChangeSource(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetpointChangeSource) String() string {
	switch v {
	case 0x00:
		return "Manual/user initiated"
	case 0x01:
		return "Scheduling/programming initiated"
	case 0x02:
		return "Externally initiated by cluster command or attribute write"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v SetpointChangeSource) IsManualUserInitiated() bool            { return v == 0x00 }
func (v SetpointChangeSource) IsSchedulingProgrammingInitiated() bool { return v == 0x01 }
func (v SetpointChangeSource) IsExternallyInitiatedByClusterCommandOrAttributeWrite() bool {
	return v == 0x02
}
func (v *SetpointChangeSource) SetManualUserInitiated()                                 { *v = 0x00 }
func (v *SetpointChangeSource) SetSchedulingProgrammingInitiated()                      { *v = 0x01 }
func (v *SetpointChangeSource) SetExternallyInitiatedByClusterCommandOrAttributeWrite() { *v = 0x02 }

func (SetpointChangeSource) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Manual/user initiated"},
		{Value: 0x01, Name: "Scheduling/programming initiated"},
		{Value: 0x02, Name: "Externally initiated by cluster command or attribute write"},
	}
}

const SetpointChangeSourceTimestampAttr zcl.AttrID = 50

func (SetpointChangeSourceTimestamp) ID() zcl.AttrID   { return SetpointChangeSourceTimestampAttr }
func (SetpointChangeSourceTimestamp) Readable() bool   { return true }
func (SetpointChangeSourceTimestamp) Writable() bool   { return false }
func (SetpointChangeSourceTimestamp) Reportable() bool { return false }
func (SetpointChangeSourceTimestamp) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v SetpointChangeSourceTimestamp) AttrID() zcl.AttrID   { return v.ID() }
func (v SetpointChangeSourceTimestamp) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SetpointChangeSourceTimestamp) AttrValue() zcl.Val  { return v.Value() }

func (SetpointChangeSourceTimestamp) Name() string { return `Setpoint Change Source Timestamp` }
func (SetpointChangeSourceTimestamp) Description() string {
	return `time in UTC at which the SetpointChangeSourceAmount attribute change
was recorded`
}

// SetpointChangeSourceTimestamp time in UTC at which the SetpointChangeSourceAmount attribute change
// was recorded
type SetpointChangeSourceTimestamp zcl.Zutc

func (v *SetpointChangeSourceTimestamp) TypeID() zcl.TypeID { return new(zcl.Zutc).TypeID() }
func (v *SetpointChangeSourceTimestamp) Value() zcl.Val     { return v }

func (v SetpointChangeSourceTimestamp) MarshalZcl() ([]byte, error) { return zcl.Zutc(v).MarshalZcl() }

func (v *SetpointChangeSourceTimestamp) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zutc)
	br, err := nt.UnmarshalZcl(b)
	*v = SetpointChangeSourceTimestamp(*nt)
	return br, err
}

func (v SetpointChangeSourceTimestamp) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zutc(v))
}

func (v *SetpointChangeSourceTimestamp) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zutc)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetpointChangeSourceTimestamp(*a)
	return nil
}

func (v *SetpointChangeSourceTimestamp) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zutc); ok {
		*v = SetpointChangeSourceTimestamp(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetpointChangeSourceTimestamp) String() string {
	return zcl.Sprintf("%v", zcl.Zutc(v))
}

func (SetpointMode) Name() string        { return `Setpoint Mode` }
func (SetpointMode) Description() string { return `` }

type SetpointMode zcl.Zenum8

func (v *SetpointMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *SetpointMode) Value() zcl.Val     { return v }

func (v SetpointMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *SetpointMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = SetpointMode(*nt)
	return br, err
}

func (v SetpointMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *SetpointMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SetpointMode(*a)
	return nil
}

func (v *SetpointMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = SetpointMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SetpointMode) String() string {
	switch v {
	case 0x00:
		return "Adjust heat setpoint"
	case 0x01:
		return "Adjust cool setpoint"
	case 0x02:
		return "Adjust heat and cool setpoint"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v SetpointMode) IsAdjustHeatSetpoint() bool        { return v == 0x00 }
func (v SetpointMode) IsAdjustCoolSetpoint() bool        { return v == 0x01 }
func (v SetpointMode) IsAdjustHeatAndCoolSetpoint() bool { return v == 0x02 }
func (v *SetpointMode) SetAdjustHeatSetpoint()           { *v = 0x00 }
func (v *SetpointMode) SetAdjustCoolSetpoint()           { *v = 0x01 }
func (v *SetpointMode) SetAdjustHeatAndCoolSetpoint()    { *v = 0x02 }

func (SetpointMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Adjust heat setpoint"},
		{Value: 0x01, Name: "Adjust cool setpoint"},
		{Value: 0x02, Name: "Adjust heat and cool setpoint"},
	}
}

const SpeedAttr zcl.AttrID = 20

func (Speed) ID() zcl.AttrID   { return SpeedAttr }
func (Speed) Readable() bool   { return true }
func (Speed) Writable() bool   { return false }
func (Speed) Reportable() bool { return false }
func (Speed) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v Speed) AttrID() zcl.AttrID   { return v.ID() }
func (v Speed) AttrType() zcl.TypeID { return v.TypeID() }
func (v *Speed) AttrValue() zcl.Val  { return v.Value() }

func (Speed) Name() string { return `Speed` }
func (Speed) Description() string {
	return `actual speed of the pump measured in RPM. It is updated dynamically as
the speed of the pump changes`
}

// Speed actual speed of the pump measured in RPM. It is updated dynamically as
// the speed of the pump changes
type Speed zcl.Zu16

func (v *Speed) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *Speed) Value() zcl.Val     { return v }

func (v Speed) MarshalZcl() ([]byte, error) { return zcl.Zu16(v).MarshalZcl() }

func (v *Speed) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = Speed(*nt)
	return br, err
}

func (v Speed) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *Speed) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = Speed(*a)
	return nil
}

func (v *Speed) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = Speed(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v Speed) String() string {
	return zcl.RevolutionsPerMinute.Format(float64(v))
}

const StartOfWeekAttr zcl.AttrID = 32

func (StartOfWeek) ID() zcl.AttrID   { return StartOfWeekAttr }
func (StartOfWeek) Readable() bool   { return true }
func (StartOfWeek) Writable() bool   { return false }
func (StartOfWeek) Reportable() bool { return false }
func (StartOfWeek) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v StartOfWeek) AttrID() zcl.AttrID   { return v.ID() }
func (v StartOfWeek) AttrType() zcl.TypeID { return v.TypeID() }
func (v *StartOfWeek) AttrValue() zcl.Val  { return v.Value() }

func (StartOfWeek) Name() string { return `Start Of Week` }
func (StartOfWeek) Description() string {
	return `day of the week that this thermostat considers to be the start of week
for weekly set point scheduling. This attribute may be able to be used
as the base to determine if the device supports weekly scheduling by
reading the attribute. Successful response means that the weekly
scheduling is supported`
}

// StartOfWeek day of the week that this thermostat considers to be the start of week
// for weekly set point scheduling. This attribute may be able to be used
// as the base to determine if the device supports weekly scheduling by
// reading the attribute. Successful response means that the weekly
// scheduling is supported
type StartOfWeek zcl.Zenum8

func (v *StartOfWeek) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *StartOfWeek) Value() zcl.Val     { return v }

func (v StartOfWeek) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *StartOfWeek) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = StartOfWeek(*nt)
	return br, err
}

func (v StartOfWeek) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *StartOfWeek) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = StartOfWeek(*a)
	return nil
}

func (v *StartOfWeek) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = StartOfWeek(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v StartOfWeek) String() string {
	switch v {
	case 0x00:
		return "Sunday"
	case 0x01:
		return "Monday"
	case 0x02:
		return "Tuesday"
	case 0x03:
		return "Wednesday"
	case 0x04:
		return "Thursday"
	case 0x05:
		return "Friday"
	case 0x06:
		return "Saturday"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v StartOfWeek) IsSunday() bool    { return v == 0x00 }
func (v StartOfWeek) IsMonday() bool    { return v == 0x01 }
func (v StartOfWeek) IsTuesday() bool   { return v == 0x02 }
func (v StartOfWeek) IsWednesday() bool { return v == 0x03 }
func (v StartOfWeek) IsThursday() bool  { return v == 0x04 }
func (v StartOfWeek) IsFriday() bool    { return v == 0x05 }
func (v StartOfWeek) IsSaturday() bool  { return v == 0x06 }
func (v *StartOfWeek) SetSunday()       { *v = 0x00 }
func (v *StartOfWeek) SetMonday()       { *v = 0x01 }
func (v *StartOfWeek) SetTuesday()      { *v = 0x02 }
func (v *StartOfWeek) SetWednesday()    { *v = 0x03 }
func (v *StartOfWeek) SetThursday()     { *v = 0x04 }
func (v *StartOfWeek) SetFriday()       { *v = 0x05 }
func (v *StartOfWeek) SetSaturday()     { *v = 0x06 }

func (StartOfWeek) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Sunday"},
		{Value: 0x01, Name: "Monday"},
		{Value: 0x02, Name: "Tuesday"},
		{Value: 0x03, Name: "Wednesday"},
		{Value: 0x04, Name: "Thursday"},
		{Value: 0x05, Name: "Friday"},
		{Value: 0x06, Name: "Saturday"},
	}
}

const SystemModeAttr zcl.AttrID = 28

func (SystemMode) ID() zcl.AttrID   { return SystemModeAttr }
func (SystemMode) Readable() bool   { return true }
func (SystemMode) Writable() bool   { return true }
func (SystemMode) Reportable() bool { return false }
func (SystemMode) SceneIndex() int  { return 3 }

// Implements AttrDef/AttrValue interfaces
func (v SystemMode) AttrID() zcl.AttrID   { return v.ID() }
func (v SystemMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *SystemMode) AttrValue() zcl.Val  { return v.Value() }

func (SystemMode) Name() string { return `System Mode` }
func (SystemMode) Description() string {
	return `specifies the current operating mode of the thermostat`
}

// SystemMode specifies the current operating mode of the thermostat
type SystemMode zcl.Zenum8

func (v *SystemMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *SystemMode) Value() zcl.Val     { return v }

func (v SystemMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *SystemMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = SystemMode(*nt)
	return br, err
}

func (v SystemMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *SystemMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = SystemMode(*a)
	return nil
}

func (v *SystemMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = SystemMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v SystemMode) String() string {
	switch v {
	case 0x00:
		return "Off"
	case 0x01:
		return "Auto"
	case 0x03:
		return "Cool"
	case 0x04:
		return "Heat"
	case 0x05:
		return "Emergency heating"
	case 0x06:
		return "Precooling"
	case 0x07:
		return "Fan only"
	case 0x08:
		return "Dry"
	case 0x09:
		return "Sleep"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v SystemMode) IsOff() bool              { return v == 0x00 }
func (v SystemMode) IsAuto() bool             { return v == 0x01 }
func (v SystemMode) IsCool() bool             { return v == 0x03 }
func (v SystemMode) IsHeat() bool             { return v == 0x04 }
func (v SystemMode) IsEmergencyHeating() bool { return v == 0x05 }
func (v SystemMode) IsPrecooling() bool       { return v == 0x06 }
func (v SystemMode) IsFanOnly() bool          { return v == 0x07 }
func (v SystemMode) IsDry() bool              { return v == 0x08 }
func (v SystemMode) IsSleep() bool            { return v == 0x09 }
func (v *SystemMode) SetOff()                 { *v = 0x00 }
func (v *SystemMode) SetAuto()                { *v = 0x01 }
func (v *SystemMode) SetCool()                { *v = 0x03 }
func (v *SystemMode) SetHeat()                { *v = 0x04 }
func (v *SystemMode) SetEmergencyHeating()    { *v = 0x05 }
func (v *SystemMode) SetPrecooling()          { *v = 0x06 }
func (v *SystemMode) SetFanOnly()             { *v = 0x07 }
func (v *SystemMode) SetDry()                 { *v = 0x08 }
func (v *SystemMode) SetSleep()               { *v = 0x09 }

func (SystemMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Off"},
		{Value: 0x01, Name: "Auto"},
		{Value: 0x03, Name: "Cool"},
		{Value: 0x04, Name: "Heat"},
		{Value: 0x05, Name: "Emergency heating"},
		{Value: 0x06, Name: "Precooling"},
		{Value: 0x07, Name: "Fan only"},
		{Value: 0x08, Name: "Dry"},
		{Value: 0x09, Name: "Sleep"},
	}
}

const TemperatureDisplayModeAttr zcl.AttrID = 0

func (TemperatureDisplayMode) ID() zcl.AttrID   { return TemperatureDisplayModeAttr }
func (TemperatureDisplayMode) Readable() bool   { return true }
func (TemperatureDisplayMode) Writable() bool   { return true }
func (TemperatureDisplayMode) Reportable() bool { return false }
func (TemperatureDisplayMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TemperatureDisplayMode) AttrID() zcl.AttrID   { return v.ID() }
func (v TemperatureDisplayMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TemperatureDisplayMode) AttrValue() zcl.Val  { return v.Value() }

func (TemperatureDisplayMode) Name() string { return `Temperature Display Mode` }
func (TemperatureDisplayMode) Description() string {
	return `units of the temperature displayed on the thermostat screen`
}

// TemperatureDisplayMode units of the temperature displayed on the thermostat screen
type TemperatureDisplayMode zcl.Zenum8

func (v *TemperatureDisplayMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *TemperatureDisplayMode) Value() zcl.Val     { return v }

func (v TemperatureDisplayMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *TemperatureDisplayMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = TemperatureDisplayMode(*nt)
	return br, err
}

func (v TemperatureDisplayMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *TemperatureDisplayMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TemperatureDisplayMode(*a)
	return nil
}

func (v *TemperatureDisplayMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = TemperatureDisplayMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TemperatureDisplayMode) String() string {
	switch v {
	case 0x00:
		return "Temperature in Celsius"
	case 0x01:
		return "Temperature in Fahrenheit"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v TemperatureDisplayMode) IsTemperatureInCelsius() bool    { return v == 0x00 }
func (v TemperatureDisplayMode) IsTemperatureInFahrenheit() bool { return v == 0x01 }
func (v *TemperatureDisplayMode) SetTemperatureInCelsius()       { *v = 0x00 }
func (v *TemperatureDisplayMode) SetTemperatureInFahrenheit()    { *v = 0x01 }

func (TemperatureDisplayMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Temperature in Celsius"},
		{Value: 0x01, Name: "Temperature in Fahrenheit"},
	}
}

const TemperatureSetpointHoldAttr zcl.AttrID = 35

func (TemperatureSetpointHold) ID() zcl.AttrID   { return TemperatureSetpointHoldAttr }
func (TemperatureSetpointHold) Readable() bool   { return true }
func (TemperatureSetpointHold) Writable() bool   { return true }
func (TemperatureSetpointHold) Reportable() bool { return false }
func (TemperatureSetpointHold) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TemperatureSetpointHold) AttrID() zcl.AttrID   { return v.ID() }
func (v TemperatureSetpointHold) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TemperatureSetpointHold) AttrValue() zcl.Val  { return v.Value() }

func (TemperatureSetpointHold) Name() string { return `Temperature Setpoint Hold` }
func (TemperatureSetpointHold) Description() string {
	return `temperature hold status on the thermostat. If hold status is on, the
thermostat should maintain the temperature set point for the current
mode until a system mode change. If hold status is off, the thermostat
should follow the setpoint transitions specified by its internal
scheduling program. If the thermostat supports setpoint hold for a
specific duration, it should also implement the
TemperatureSetpointHoldDuration attribute`
}

// TemperatureSetpointHold temperature hold status on the thermostat. If hold status is on, the
// thermostat should maintain the temperature set point for the current
// mode until a system mode change. If hold status is off, the thermostat
// should follow the setpoint transitions specified by its internal
// scheduling program. If the thermostat supports setpoint hold for a
// specific duration, it should also implement the
// TemperatureSetpointHoldDuration attribute
type TemperatureSetpointHold zcl.Zenum8

func (v *TemperatureSetpointHold) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *TemperatureSetpointHold) Value() zcl.Val     { return v }

func (v TemperatureSetpointHold) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *TemperatureSetpointHold) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = TemperatureSetpointHold(*nt)
	return br, err
}

func (v TemperatureSetpointHold) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *TemperatureSetpointHold) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TemperatureSetpointHold(*a)
	return nil
}

func (v *TemperatureSetpointHold) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = TemperatureSetpointHold(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TemperatureSetpointHold) String() string {
	switch v {
	case 0x00:
		return "Off"
	case 0x01:
		return "On"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v TemperatureSetpointHold) IsOff() bool { return v == 0x00 }
func (v TemperatureSetpointHold) IsOn() bool  { return v == 0x01 }
func (v *TemperatureSetpointHold) SetOff()    { *v = 0x00 }
func (v *TemperatureSetpointHold) SetOn()     { *v = 0x01 }

func (TemperatureSetpointHold) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Off"},
		{Value: 0x01, Name: "On"},
	}
}

const TemperatureSetpointHoldDurationAttr zcl.AttrID = 36

func (TemperatureSetpointHoldDuration) ID() zcl.AttrID   { return TemperatureSetpointHoldDurationAttr }
func (TemperatureSetpointHoldDuration) Readable() bool   { return true }
func (TemperatureSetpointHoldDuration) Writable() bool   { return true }
func (TemperatureSetpointHoldDuration) Reportable() bool { return false }
func (TemperatureSetpointHoldDuration) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v TemperatureSetpointHoldDuration) AttrID() zcl.AttrID   { return v.ID() }
func (v TemperatureSetpointHoldDuration) AttrType() zcl.TypeID { return v.TypeID() }
func (v *TemperatureSetpointHoldDuration) AttrValue() zcl.Val  { return v.Value() }

func (TemperatureSetpointHoldDuration) Name() string { return `Temperature Setpoint Hold Duration` }
func (TemperatureSetpointHoldDuration) Description() string {
	return `period in minutes for which a setpoint hold is active`
}

// TemperatureSetpointHoldDuration period in minutes for which a setpoint hold is active
type TemperatureSetpointHoldDuration zcl.Zu16

func (v *TemperatureSetpointHoldDuration) TypeID() zcl.TypeID { return new(zcl.Zu16).TypeID() }
func (v *TemperatureSetpointHoldDuration) Value() zcl.Val     { return v }

func (v TemperatureSetpointHoldDuration) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(v).MarshalZcl()
}

func (v *TemperatureSetpointHoldDuration) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*v = TemperatureSetpointHoldDuration(*nt)
	return br, err
}

func (v TemperatureSetpointHoldDuration) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu16(v))
}

func (v *TemperatureSetpointHoldDuration) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = TemperatureSetpointHoldDuration(*a)
	return nil
}

func (v *TemperatureSetpointHoldDuration) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu16); ok {
		*v = TemperatureSetpointHoldDuration(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v TemperatureSetpointHoldDuration) String() string {
	return zcl.Minutes.Format(float64(v))
}

const ThermostatAlarmMaskAttr zcl.AttrID = 29

func (ThermostatAlarmMask) ID() zcl.AttrID   { return ThermostatAlarmMaskAttr }
func (ThermostatAlarmMask) Readable() bool   { return true }
func (ThermostatAlarmMask) Writable() bool   { return false }
func (ThermostatAlarmMask) Reportable() bool { return false }
func (ThermostatAlarmMask) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ThermostatAlarmMask) AttrID() zcl.AttrID   { return v.ID() }
func (v ThermostatAlarmMask) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ThermostatAlarmMask) AttrValue() zcl.Val  { return v.Value() }

func (ThermostatAlarmMask) Name() string { return `Thermostat Alarm Mask` }
func (ThermostatAlarmMask) Description() string {
	return `specifies whether each of the alarms listed is enabled. When the bit
number corresponding to the alarm code is set to 1, the alarm is
enabled, else it is disabled`
}

// ThermostatAlarmMask specifies whether each of the alarms listed is enabled. When the bit
// number corresponding to the alarm code is set to 1, the alarm is
// enabled, else it is disabled
type ThermostatAlarmMask zcl.Zbmp8

func (v *ThermostatAlarmMask) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *ThermostatAlarmMask) Value() zcl.Val     { return v }

func (v ThermostatAlarmMask) MarshalZcl() ([]byte, error) { return zcl.Zbmp8(v).MarshalZcl() }

func (v *ThermostatAlarmMask) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = ThermostatAlarmMask(*nt)
	return br, err
}

func (v ThermostatAlarmMask) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *ThermostatAlarmMask) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ThermostatAlarmMask(*a)
	return nil
}

func (v *ThermostatAlarmMask) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = ThermostatAlarmMask(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ThermostatAlarmMask) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Initialisation failure")
		case 1:
			bstr = append(bstr, "Hardware failure")
		case 2:
			bstr = append(bstr, "Self-calibration failure")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v ThermostatAlarmMask) IsInitialisationFailure() bool  { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v ThermostatAlarmMask) IsHardwareFailure() bool        { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v ThermostatAlarmMask) IsSelfCalibrationFailure() bool { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v *ThermostatAlarmMask) SetInitialisationFailure(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *ThermostatAlarmMask) SetHardwareFailure(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *ThermostatAlarmMask) SetSelfCalibrationFailure(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}

func (ThermostatAlarmMask) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Initialisation failure"},
		{Value: 1, Name: "Hardware failure"},
		{Value: 2, Name: "Self-calibration failure"},
	}
}

const ThermostatProgrammingOperationModeAttr zcl.AttrID = 37

func (ThermostatProgrammingOperationMode) ID() zcl.AttrID {
	return ThermostatProgrammingOperationModeAttr
}
func (ThermostatProgrammingOperationMode) Readable() bool   { return true }
func (ThermostatProgrammingOperationMode) Writable() bool   { return true }
func (ThermostatProgrammingOperationMode) Reportable() bool { return true }
func (ThermostatProgrammingOperationMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ThermostatProgrammingOperationMode) AttrID() zcl.AttrID   { return v.ID() }
func (v ThermostatProgrammingOperationMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ThermostatProgrammingOperationMode) AttrValue() zcl.Val  { return v.Value() }

func (ThermostatProgrammingOperationMode) Name() string {
	return `Thermostat Programming Operation Mode`
}
func (ThermostatProgrammingOperationMode) Description() string {
	return `operational state of the thermostat's programming. The thermostat shall
modify its programming operation when this attribute is modified by a
client and update this attribute when its programming operation is
modified locally by a user. When a bit is 0 it means off, 1 means on.
For the scheduling mode bit, 0 means the thermostate is manually
controlled, whereas 1 means it is following a programmed weekly
schedule`
}

// ThermostatProgrammingOperationMode operational state of the thermostat's programming. The thermostat shall
// modify its programming operation when this attribute is modified by a
// client and update this attribute when its programming operation is
// modified locally by a user. When a bit is 0 it means off, 1 means on.
// For the scheduling mode bit, 0 means the thermostate is manually
// controlled, whereas 1 means it is following a programmed weekly
// schedule
type ThermostatProgrammingOperationMode zcl.Zbmp8

func (v *ThermostatProgrammingOperationMode) TypeID() zcl.TypeID { return new(zcl.Zbmp8).TypeID() }
func (v *ThermostatProgrammingOperationMode) Value() zcl.Val     { return v }

func (v ThermostatProgrammingOperationMode) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(v).MarshalZcl()
}

func (v *ThermostatProgrammingOperationMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*v = ThermostatProgrammingOperationMode(*nt)
	return br, err
}

func (v ThermostatProgrammingOperationMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp8(v))
}

func (v *ThermostatProgrammingOperationMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ThermostatProgrammingOperationMode(*a)
	return nil
}

func (v *ThermostatProgrammingOperationMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp8); ok {
		*v = ThermostatProgrammingOperationMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ThermostatProgrammingOperationMode) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Scheduling mode")
		case 1:
			bstr = append(bstr, "Auto/recovery mode")
		case 2:
			bstr = append(bstr, "Economy mode")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v ThermostatProgrammingOperationMode) IsSchedulingMode() bool {
	return zcl.BitmapTest([]byte(v[:]), 0)
}
func (v ThermostatProgrammingOperationMode) IsAutoRecoveryMode() bool {
	return zcl.BitmapTest([]byte(v[:]), 1)
}
func (v ThermostatProgrammingOperationMode) IsEconomyMode() bool {
	return zcl.BitmapTest([]byte(v[:]), 2)
}
func (v *ThermostatProgrammingOperationMode) SetSchedulingMode(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *ThermostatProgrammingOperationMode) SetAutoRecoveryMode(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *ThermostatProgrammingOperationMode) SetEconomyMode(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}

func (ThermostatProgrammingOperationMode) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Scheduling mode"},
		{Value: 1, Name: "Auto/recovery mode"},
		{Value: 2, Name: "Economy mode"},
	}
}

const ThermostatRunningModeAttr zcl.AttrID = 30

func (ThermostatRunningMode) ID() zcl.AttrID   { return ThermostatRunningModeAttr }
func (ThermostatRunningMode) Readable() bool   { return true }
func (ThermostatRunningMode) Writable() bool   { return false }
func (ThermostatRunningMode) Reportable() bool { return false }
func (ThermostatRunningMode) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ThermostatRunningMode) AttrID() zcl.AttrID   { return v.ID() }
func (v ThermostatRunningMode) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ThermostatRunningMode) AttrValue() zcl.Val  { return v.Value() }

func (ThermostatRunningMode) Name() string { return `Thermostat Running Mode` }
func (ThermostatRunningMode) Description() string {
	return `represents the running mode of the thermostat`
}

// ThermostatRunningMode represents the running mode of the thermostat
type ThermostatRunningMode zcl.Zenum8

func (v *ThermostatRunningMode) TypeID() zcl.TypeID { return new(zcl.Zenum8).TypeID() }
func (v *ThermostatRunningMode) Value() zcl.Val     { return v }

func (v ThermostatRunningMode) MarshalZcl() ([]byte, error) { return zcl.Zenum8(v).MarshalZcl() }

func (v *ThermostatRunningMode) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*v = ThermostatRunningMode(*nt)
	return br, err
}

func (v ThermostatRunningMode) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zenum8(v))
}

func (v *ThermostatRunningMode) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zenum8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ThermostatRunningMode(*a)
	return nil
}

func (v *ThermostatRunningMode) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zenum8); ok {
		*v = ThermostatRunningMode(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ThermostatRunningMode) String() string {
	switch v {
	case 0x00:
		return "Off"
	case 0x03:
		return "Cool"
	case 0x04:
		return "Heat"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(v))
}

func (v ThermostatRunningMode) IsOff() bool  { return v == 0x00 }
func (v ThermostatRunningMode) IsCool() bool { return v == 0x03 }
func (v ThermostatRunningMode) IsHeat() bool { return v == 0x04 }
func (v *ThermostatRunningMode) SetOff()     { *v = 0x00 }
func (v *ThermostatRunningMode) SetCool()    { *v = 0x03 }
func (v *ThermostatRunningMode) SetHeat()    { *v = 0x04 }

func (ThermostatRunningMode) SingleOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0x00, Name: "Off"},
		{Value: 0x03, Name: "Cool"},
		{Value: 0x04, Name: "Heat"},
	}
}

const ThermostatRunningStateAttr zcl.AttrID = 41

func (ThermostatRunningState) ID() zcl.AttrID   { return ThermostatRunningStateAttr }
func (ThermostatRunningState) Readable() bool   { return true }
func (ThermostatRunningState) Writable() bool   { return false }
func (ThermostatRunningState) Reportable() bool { return false }
func (ThermostatRunningState) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v ThermostatRunningState) AttrID() zcl.AttrID   { return v.ID() }
func (v ThermostatRunningState) AttrType() zcl.TypeID { return v.TypeID() }
func (v *ThermostatRunningState) AttrValue() zcl.Val  { return v.Value() }

func (ThermostatRunningState) Name() string        { return `Thermostat Running State` }
func (ThermostatRunningState) Description() string { return `` }

type ThermostatRunningState zcl.Zbmp16

func (v *ThermostatRunningState) TypeID() zcl.TypeID { return new(zcl.Zbmp16).TypeID() }
func (v *ThermostatRunningState) Value() zcl.Val     { return v }

func (v ThermostatRunningState) MarshalZcl() ([]byte, error) { return zcl.Zbmp16(v).MarshalZcl() }

func (v *ThermostatRunningState) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp16)
	br, err := nt.UnmarshalZcl(b)
	*v = ThermostatRunningState(*nt)
	return br, err
}

func (v ThermostatRunningState) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zbmp16(v))
}

func (v *ThermostatRunningState) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zbmp16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = ThermostatRunningState(*a)
	return nil
}

func (v *ThermostatRunningState) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zbmp16); ok {
		*v = ThermostatRunningState(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v ThermostatRunningState) String() string {
	var bstr []string
	bits := zcl.BitmapList(v[:])
	for _, bit := range bits {
		switch bit {
		case 0:
			bstr = append(bstr, "Heat stage on")
		case 1:
			bstr = append(bstr, "Cool stage on")
		case 2:
			bstr = append(bstr, "Fan stage on")
		case 3:
			bstr = append(bstr, "Heat second stage on")
		case 4:
			bstr = append(bstr, "Cool second stage on")
		case 5:
			bstr = append(bstr, "Fan second stage on")
		case 6:
			bstr = append(bstr, "Fan third stage on")
		default:
			bstr = append(bstr, zcl.Sprintf("Unknown(%d)", bit))
		}
	}
	return zcl.StrJoin(bstr, ", ")
}

func (v ThermostatRunningState) IsHeatStageOn() bool       { return zcl.BitmapTest([]byte(v[:]), 0) }
func (v ThermostatRunningState) IsCoolStageOn() bool       { return zcl.BitmapTest([]byte(v[:]), 1) }
func (v ThermostatRunningState) IsFanStageOn() bool        { return zcl.BitmapTest([]byte(v[:]), 2) }
func (v ThermostatRunningState) IsHeatSecondStageOn() bool { return zcl.BitmapTest([]byte(v[:]), 3) }
func (v ThermostatRunningState) IsCoolSecondStageOn() bool { return zcl.BitmapTest([]byte(v[:]), 4) }
func (v ThermostatRunningState) IsFanSecondStageOn() bool  { return zcl.BitmapTest([]byte(v[:]), 5) }
func (v ThermostatRunningState) IsFanThirdStageOn() bool   { return zcl.BitmapTest([]byte(v[:]), 6) }
func (v *ThermostatRunningState) SetHeatStageOn(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 0, b))
}
func (v *ThermostatRunningState) SetCoolStageOn(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 1, b))
}
func (v *ThermostatRunningState) SetFanStageOn(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 2, b))
}
func (v *ThermostatRunningState) SetHeatSecondStageOn(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 3, b))
}
func (v *ThermostatRunningState) SetCoolSecondStageOn(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 4, b))
}
func (v *ThermostatRunningState) SetFanSecondStageOn(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 5, b))
}
func (v *ThermostatRunningState) SetFanThirdStageOn(b bool) {
	copy((*v)[:], zcl.BitmapSet([]byte((*v)[:]), 6, b))
}

func (ThermostatRunningState) MultiOptions() []zcl.Option {
	return []zcl.Option{
		{Value: 0, Name: "Heat stage on"},
		{Value: 1, Name: "Cool stage on"},
		{Value: 2, Name: "Fan stage on"},
		{Value: 3, Name: "Heat second stage on"},
		{Value: 4, Name: "Cool second stage on"},
		{Value: 5, Name: "Fan second stage on"},
		{Value: 6, Name: "Fan third stage on"},
	}
}

const UnoccupiedCoolingSetpointAttr zcl.AttrID = 19

func (UnoccupiedCoolingSetpoint) ID() zcl.AttrID   { return UnoccupiedCoolingSetpointAttr }
func (UnoccupiedCoolingSetpoint) Readable() bool   { return true }
func (UnoccupiedCoolingSetpoint) Writable() bool   { return true }
func (UnoccupiedCoolingSetpoint) Reportable() bool { return false }
func (UnoccupiedCoolingSetpoint) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UnoccupiedCoolingSetpoint) AttrID() zcl.AttrID   { return v.ID() }
func (v UnoccupiedCoolingSetpoint) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UnoccupiedCoolingSetpoint) AttrValue() zcl.Val  { return v.Value() }

func (UnoccupiedCoolingSetpoint) Name() string { return `Unoccupied Cooling Setpoint` }
func (UnoccupiedCoolingSetpoint) Description() string {
	return `cooling mode setpoint when the room is unoccupied. The
UnoccupiedHeatingSetpoint attribute shall always be below the value
specified by at least MinSetpointDeadband`
}

// UnoccupiedCoolingSetpoint cooling mode setpoint when the room is unoccupied. The
// UnoccupiedHeatingSetpoint attribute shall always be below the value
// specified by at least MinSetpointDeadband
type UnoccupiedCoolingSetpoint zcl.Zs16

func (v *UnoccupiedCoolingSetpoint) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *UnoccupiedCoolingSetpoint) Value() zcl.Val     { return v }

func (v UnoccupiedCoolingSetpoint) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *UnoccupiedCoolingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = UnoccupiedCoolingSetpoint(*nt)
	return br, err
}

func (v UnoccupiedCoolingSetpoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *UnoccupiedCoolingSetpoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UnoccupiedCoolingSetpoint(*a)
	return nil
}

func (v *UnoccupiedCoolingSetpoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = UnoccupiedCoolingSetpoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UnoccupiedCoolingSetpoint) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v UnoccupiedCoolingSetpoint) Scaled() float64 {
	return float64(v) / 10
}

const UnoccupiedHeatingSetpointAttr zcl.AttrID = 20

func (UnoccupiedHeatingSetpoint) ID() zcl.AttrID   { return UnoccupiedHeatingSetpointAttr }
func (UnoccupiedHeatingSetpoint) Readable() bool   { return true }
func (UnoccupiedHeatingSetpoint) Writable() bool   { return true }
func (UnoccupiedHeatingSetpoint) Reportable() bool { return false }
func (UnoccupiedHeatingSetpoint) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UnoccupiedHeatingSetpoint) AttrID() zcl.AttrID   { return v.ID() }
func (v UnoccupiedHeatingSetpoint) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UnoccupiedHeatingSetpoint) AttrValue() zcl.Val  { return v.Value() }

func (UnoccupiedHeatingSetpoint) Name() string { return `Unoccupied Heating Setpoint` }
func (UnoccupiedHeatingSetpoint) Description() string {
	return `heating mode setpoint when the room is unoccupied. The
UnoccupiedCoolingSetpoint attribute shall always be above the value
specified by at least MinSetpointDeadband`
}

// UnoccupiedHeatingSetpoint heating mode setpoint when the room is unoccupied. The
// UnoccupiedCoolingSetpoint attribute shall always be above the value
// specified by at least MinSetpointDeadband
type UnoccupiedHeatingSetpoint zcl.Zs16

func (v *UnoccupiedHeatingSetpoint) TypeID() zcl.TypeID { return new(zcl.Zs16).TypeID() }
func (v *UnoccupiedHeatingSetpoint) Value() zcl.Val     { return v }

func (v UnoccupiedHeatingSetpoint) MarshalZcl() ([]byte, error) { return zcl.Zs16(v).MarshalZcl() }

func (v *UnoccupiedHeatingSetpoint) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zs16)
	br, err := nt.UnmarshalZcl(b)
	*v = UnoccupiedHeatingSetpoint(*nt)
	return br, err
}

func (v UnoccupiedHeatingSetpoint) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zs16(v))
}

func (v *UnoccupiedHeatingSetpoint) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zs16)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UnoccupiedHeatingSetpoint(*a)
	return nil
}

func (v *UnoccupiedHeatingSetpoint) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zs16); ok {
		*v = UnoccupiedHeatingSetpoint(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UnoccupiedHeatingSetpoint) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v UnoccupiedHeatingSetpoint) Scaled() float64 {
	return float64(v) / 10
}

const UnoccupiedSetbackAttr zcl.AttrID = 55

func (UnoccupiedSetback) ID() zcl.AttrID   { return UnoccupiedSetbackAttr }
func (UnoccupiedSetback) Readable() bool   { return true }
func (UnoccupiedSetback) Writable() bool   { return true }
func (UnoccupiedSetback) Reportable() bool { return false }
func (UnoccupiedSetback) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UnoccupiedSetback) AttrID() zcl.AttrID   { return v.ID() }
func (v UnoccupiedSetback) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UnoccupiedSetback) AttrValue() zcl.Val  { return v.Value() }

func (UnoccupiedSetback) Name() string { return `Unoccupied Setback` }
func (UnoccupiedSetback) Description() string {
	return `degrees the thermostat will allow the LocalTemperature attribute to
float above the UnoccupiedCooling setpoint (i.e., UnoccupiedCooling +
UnoccupiedSetback) or below the UnoccupiedHeating setpoint (i.e.,
UnoccupiedHeating - UnoccupiedSetback) before initiating a state change
to bring the temperature back to the user's desired setpoint`
}

// UnoccupiedSetback degrees the thermostat will allow the LocalTemperature attribute to
// float above the UnoccupiedCooling setpoint (i.e., UnoccupiedCooling +
// UnoccupiedSetback) or below the UnoccupiedHeating setpoint (i.e.,
// UnoccupiedHeating - UnoccupiedSetback) before initiating a state change
// to bring the temperature back to the user's desired setpoint
type UnoccupiedSetback zcl.Zu8

func (v *UnoccupiedSetback) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *UnoccupiedSetback) Value() zcl.Val     { return v }

func (v UnoccupiedSetback) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *UnoccupiedSetback) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = UnoccupiedSetback(*nt)
	return br, err
}

func (v UnoccupiedSetback) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *UnoccupiedSetback) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UnoccupiedSetback(*a)
	return nil
}

func (v *UnoccupiedSetback) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = UnoccupiedSetback(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UnoccupiedSetback) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v UnoccupiedSetback) Scaled() float64 {
	return float64(v) / 10
}

const UnoccupiedSetbackMaxAttr zcl.AttrID = 57

func (UnoccupiedSetbackMax) ID() zcl.AttrID   { return UnoccupiedSetbackMaxAttr }
func (UnoccupiedSetbackMax) Readable() bool   { return true }
func (UnoccupiedSetbackMax) Writable() bool   { return false }
func (UnoccupiedSetbackMax) Reportable() bool { return false }
func (UnoccupiedSetbackMax) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UnoccupiedSetbackMax) AttrID() zcl.AttrID   { return v.ID() }
func (v UnoccupiedSetbackMax) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UnoccupiedSetbackMax) AttrValue() zcl.Val  { return v.Value() }

func (UnoccupiedSetbackMax) Name() string { return `Unoccupied Setback Max` }
func (UnoccupiedSetbackMax) Description() string {
	return `degrees the thermostat will allow the UnoccupiedSetback attribute to be
configured by a user`
}

// UnoccupiedSetbackMax degrees the thermostat will allow the UnoccupiedSetback attribute to be
// configured by a user
type UnoccupiedSetbackMax zcl.Zu8

func (v *UnoccupiedSetbackMax) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *UnoccupiedSetbackMax) Value() zcl.Val     { return v }

func (v UnoccupiedSetbackMax) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *UnoccupiedSetbackMax) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = UnoccupiedSetbackMax(*nt)
	return br, err
}

func (v UnoccupiedSetbackMax) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *UnoccupiedSetbackMax) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UnoccupiedSetbackMax(*a)
	return nil
}

func (v *UnoccupiedSetbackMax) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = UnoccupiedSetbackMax(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UnoccupiedSetbackMax) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v UnoccupiedSetbackMax) Scaled() float64 {
	return float64(v) / 10
}

const UnoccupiedSetbackMinAttr zcl.AttrID = 56

func (UnoccupiedSetbackMin) ID() zcl.AttrID   { return UnoccupiedSetbackMinAttr }
func (UnoccupiedSetbackMin) Readable() bool   { return true }
func (UnoccupiedSetbackMin) Writable() bool   { return false }
func (UnoccupiedSetbackMin) Reportable() bool { return false }
func (UnoccupiedSetbackMin) SceneIndex() int  { return -1 }

// Implements AttrDef/AttrValue interfaces
func (v UnoccupiedSetbackMin) AttrID() zcl.AttrID   { return v.ID() }
func (v UnoccupiedSetbackMin) AttrType() zcl.TypeID { return v.TypeID() }
func (v *UnoccupiedSetbackMin) AttrValue() zcl.Val  { return v.Value() }

func (UnoccupiedSetbackMin) Name() string { return `Unoccupied Setback Min` }
func (UnoccupiedSetbackMin) Description() string {
	return `degrees the thermostat will allow the UnoccupiedSetback attribute to be
configured by a user`
}

// UnoccupiedSetbackMin degrees the thermostat will allow the UnoccupiedSetback attribute to be
// configured by a user
type UnoccupiedSetbackMin zcl.Zu8

func (v *UnoccupiedSetbackMin) TypeID() zcl.TypeID { return new(zcl.Zu8).TypeID() }
func (v *UnoccupiedSetbackMin) Value() zcl.Val     { return v }

func (v UnoccupiedSetbackMin) MarshalZcl() ([]byte, error) { return zcl.Zu8(v).MarshalZcl() }

func (v *UnoccupiedSetbackMin) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*v = UnoccupiedSetbackMin(*nt)
	return br, err
}

func (v UnoccupiedSetbackMin) MarshalJSON() ([]byte, error) {
	return zcl.ToJson(zcl.Zu8(v))
}

func (v *UnoccupiedSetbackMin) UnmarshalJSON(b []byte) error {
	a := new(zcl.Zu8)
	if err := zcl.ParseJson(b, a); err != nil {
		return err
	}
	*v = UnoccupiedSetbackMin(*a)
	return nil
}

func (v *UnoccupiedSetbackMin) SetValue(a zcl.Val) error {
	if nv, ok := a.(*zcl.Zu8); ok {
		*v = UnoccupiedSetbackMin(*nv)
		return nil
	}
	return zcl.ErrInvalidType
}

func (v UnoccupiedSetbackMin) String() string {
	return zcl.DegreesCelsius.Format(float64(v) / 10)
}

func (v UnoccupiedSetbackMin) Scaled() float64 {
	return float64(v) / 10
}
