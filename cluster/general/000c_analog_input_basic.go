// An interface for reading the value of an analog measurement and accessing various characteristics of that measurement.
package general

import (
	"neotor.se/zcl"
)

// AnalogInputBasic
// An interface for reading the value of an analog measurement and accessing various characteristics of that measurement.

func NewAnalogInputBasicServer(profile zcl.ProfileID) *AnalogInputBasicServer {
	return &AnalogInputBasicServer{p: profile}
}
func NewAnalogInputBasicClient(profile zcl.ProfileID) *AnalogInputBasicClient {
	return &AnalogInputBasicClient{p: profile}
}

const AnalogInputBasicCluster zcl.ClusterID = 12

type AnalogInputBasicServer struct {
	p zcl.ProfileID

	AnalogInputDescription      *AnalogInputDescription
	AnalogInputMaxPresentValue  *AnalogInputMaxPresentValue
	AnalogInputMinPresentValue  *AnalogInputMinPresentValue
	AnalogInputOutOfService     *AnalogInputOutOfService
	AnalogInputPresentValue     *AnalogInputPresentValue
	AnalogInputReliability      *AnalogInputReliability
	AnalogInputResolution       *AnalogInputResolution
	AnalogInputStatusFlags      *AnalogInputStatusFlags
	AnalogInputEngineeringUnits *AnalogInputEngineeringUnits
	AnalogInputApplicationType  *AnalogInputApplicationType
}

type AnalogInputBasicClient struct {
	p zcl.ProfileID
}

/*
var AnalogInputBasicServer = map[zcl.CommandID]func() zcl.Command{
}

var AnalogInputBasicClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type AnalogInputDescription zcl.Zcstring

func (a AnalogInputDescription) ID() zcl.AttrID                  { return 28 }
func (a AnalogInputDescription) Cluster() zcl.ClusterID          { return AnalogInputBasicCluster }
func (a *AnalogInputDescription) Value() *AnalogInputDescription { return a }
func (a AnalogInputDescription) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *AnalogInputDescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputDescription(*nt)
	return br, err
}

func (a AnalogInputDescription) Readable() bool   { return true }
func (a AnalogInputDescription) Writable() bool   { return true }
func (a AnalogInputDescription) Reportable() bool { return false }
func (a AnalogInputDescription) SceneIndex() int  { return -1 }

func (a AnalogInputDescription) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

type AnalogInputMaxPresentValue zcl.Zfloat

func (a AnalogInputMaxPresentValue) ID() zcl.AttrID                      { return 65 }
func (a AnalogInputMaxPresentValue) Cluster() zcl.ClusterID              { return AnalogInputBasicCluster }
func (a *AnalogInputMaxPresentValue) Value() *AnalogInputMaxPresentValue { return a }
func (a AnalogInputMaxPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogInputMaxPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputMaxPresentValue(*nt)
	return br, err
}

func (a AnalogInputMaxPresentValue) Readable() bool   { return true }
func (a AnalogInputMaxPresentValue) Writable() bool   { return true }
func (a AnalogInputMaxPresentValue) Reportable() bool { return false }
func (a AnalogInputMaxPresentValue) SceneIndex() int  { return -1 }

func (a AnalogInputMaxPresentValue) String() string {

	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

type AnalogInputMinPresentValue zcl.Zfloat

func (a AnalogInputMinPresentValue) ID() zcl.AttrID                      { return 66 }
func (a AnalogInputMinPresentValue) Cluster() zcl.ClusterID              { return AnalogInputBasicCluster }
func (a *AnalogInputMinPresentValue) Value() *AnalogInputMinPresentValue { return a }
func (a AnalogInputMinPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogInputMinPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputMinPresentValue(*nt)
	return br, err
}

func (a AnalogInputMinPresentValue) Readable() bool   { return true }
func (a AnalogInputMinPresentValue) Writable() bool   { return true }
func (a AnalogInputMinPresentValue) Reportable() bool { return false }
func (a AnalogInputMinPresentValue) SceneIndex() int  { return -1 }

func (a AnalogInputMinPresentValue) String() string {

	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

type AnalogInputOutOfService zcl.Zbool

func (a AnalogInputOutOfService) ID() zcl.AttrID                   { return 81 }
func (a AnalogInputOutOfService) Cluster() zcl.ClusterID           { return AnalogInputBasicCluster }
func (a *AnalogInputOutOfService) Value() *AnalogInputOutOfService { return a }
func (a AnalogInputOutOfService) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *AnalogInputOutOfService) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputOutOfService(*nt)
	return br, err
}

func (a AnalogInputOutOfService) Readable() bool   { return true }
func (a AnalogInputOutOfService) Writable() bool   { return true }
func (a AnalogInputOutOfService) Reportable() bool { return false }
func (a AnalogInputOutOfService) SceneIndex() int  { return -1 }

func (a AnalogInputOutOfService) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

type AnalogInputPresentValue zcl.Zfloat

func (a AnalogInputPresentValue) ID() zcl.AttrID                   { return 85 }
func (a AnalogInputPresentValue) Cluster() zcl.ClusterID           { return AnalogInputBasicCluster }
func (a *AnalogInputPresentValue) Value() *AnalogInputPresentValue { return a }
func (a AnalogInputPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogInputPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputPresentValue(*nt)
	return br, err
}

func (a AnalogInputPresentValue) Readable() bool   { return true }
func (a AnalogInputPresentValue) Writable() bool   { return true }
func (a AnalogInputPresentValue) Reportable() bool { return true }
func (a AnalogInputPresentValue) SceneIndex() int  { return -1 }

func (a AnalogInputPresentValue) String() string {

	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

type AnalogInputReliability zcl.Zenum8

func (a AnalogInputReliability) ID() zcl.AttrID                  { return 103 }
func (a AnalogInputReliability) Cluster() zcl.ClusterID          { return AnalogInputBasicCluster }
func (a *AnalogInputReliability) Value() *AnalogInputReliability { return a }
func (a AnalogInputReliability) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *AnalogInputReliability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputReliability(*nt)
	return br, err
}

func (a AnalogInputReliability) Readable() bool   { return true }
func (a AnalogInputReliability) Writable() bool   { return true }
func (a AnalogInputReliability) Reportable() bool { return false }
func (a AnalogInputReliability) SceneIndex() int  { return -1 }

func (a AnalogInputReliability) String() string {

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

type AnalogInputResolution zcl.Zfloat

func (a AnalogInputResolution) ID() zcl.AttrID                 { return 106 }
func (a AnalogInputResolution) Cluster() zcl.ClusterID         { return AnalogInputBasicCluster }
func (a *AnalogInputResolution) Value() *AnalogInputResolution { return a }
func (a AnalogInputResolution) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogInputResolution) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputResolution(*nt)
	return br, err
}

func (a AnalogInputResolution) Readable() bool   { return true }
func (a AnalogInputResolution) Writable() bool   { return true }
func (a AnalogInputResolution) Reportable() bool { return false }
func (a AnalogInputResolution) SceneIndex() int  { return -1 }

func (a AnalogInputResolution) String() string {

	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

type AnalogInputStatusFlags zcl.Zbmp8

func (a AnalogInputStatusFlags) ID() zcl.AttrID                  { return 111 }
func (a AnalogInputStatusFlags) Cluster() zcl.ClusterID          { return AnalogInputBasicCluster }
func (a *AnalogInputStatusFlags) Value() *AnalogInputStatusFlags { return a }
func (a AnalogInputStatusFlags) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *AnalogInputStatusFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputStatusFlags(*nt)
	return br, err
}

func (a AnalogInputStatusFlags) Readable() bool   { return true }
func (a AnalogInputStatusFlags) Writable() bool   { return false }
func (a AnalogInputStatusFlags) Reportable() bool { return true }
func (a AnalogInputStatusFlags) SceneIndex() int  { return -1 }

func (a AnalogInputStatusFlags) String() string {

	var bstr []string
	if a.IsInAlarm() {
		bstr = append(bstr, "In Alarm")
	}
	if a.IsFault() {
		bstr = append(bstr, "Fault")
	}
	if a.IsOveridden() {
		bstr = append(bstr, "Overidden")
	}
	if a.IsOutOfService() {
		bstr = append(bstr, "Out of Service")
	}
	return zcl.StrJoin(bstr, ", ")

}

func (a AnalogInputStatusFlags) IsInAlarm() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *AnalogInputStatusFlags) SetInAlarm(b bool) {
	*a = AnalogInputStatusFlags(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a AnalogInputStatusFlags) IsFault() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *AnalogInputStatusFlags) SetFault(b bool) {
	*a = AnalogInputStatusFlags(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a AnalogInputStatusFlags) IsOveridden() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *AnalogInputStatusFlags) SetOveridden(b bool) {
	*a = AnalogInputStatusFlags(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a AnalogInputStatusFlags) IsOutOfService() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *AnalogInputStatusFlags) SetOutOfService(b bool) {
	*a = AnalogInputStatusFlags(zcl.BitmapSet([]byte(*a), 8, b))
}

type AnalogInputEngineeringUnits zcl.Zenum16

func (a AnalogInputEngineeringUnits) ID() zcl.AttrID                       { return 117 }
func (a AnalogInputEngineeringUnits) Cluster() zcl.ClusterID               { return AnalogInputBasicCluster }
func (a *AnalogInputEngineeringUnits) Value() *AnalogInputEngineeringUnits { return a }
func (a AnalogInputEngineeringUnits) MarshalZcl() ([]byte, error) {
	return zcl.Zenum16(a).MarshalZcl()
}
func (a *AnalogInputEngineeringUnits) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputEngineeringUnits(*nt)
	return br, err
}

func (a AnalogInputEngineeringUnits) Readable() bool   { return true }
func (a AnalogInputEngineeringUnits) Writable() bool   { return true }
func (a AnalogInputEngineeringUnits) Reportable() bool { return false }
func (a AnalogInputEngineeringUnits) SceneIndex() int  { return -1 }

func (a AnalogInputEngineeringUnits) String() string {
	switch a {
	case 0x0000:
		return "Square-meters"
	case 0x0001:
		return "Square-feet"
	case 0x0002:
		return "Milliamperes"
	case 0x0003:
		return "Amperes"
	case 0x0004:
		return "Ohms"
	case 0x0005:
		return "Volts"
	case 0x0006:
		return "Kilo-volts"
	case 0x0007:
		return "Mega-volts"
	case 0x0008:
		return "Volt-amperes"
	case 0x0009:
		return "Kilo-volt-amperes"
	case 0x000A:
		return "Mega-volt-amperes"
	case 0x000B:
		return "Volt-amperes-reactive"
	case 0x000C:
		return "Kilo-volt-amperes-reactive"
	case 0x000D:
		return "Mega-volt-amperes-reactive"
	case 0x000E:
		return "Degrees-phase"
	case 0x000F:
		return "Power-factor"
	case 0x0010:
		return "Joules"
	case 0x0011:
		return "Kilojoules"
	case 0x0012:
		return "Watt-hours"
	case 0x0013:
		return "Kilowatt-hours"
	case 0x0014:
		return "BTUs"
	case 0x0015:
		return "Therms"
	case 0x0016:
		return "Ton-hours"
	case 0x0017:
		return "Joules-per-kilogram-dry-air"
	case 0x0018:
		return "BTUs-per-pound-dry-air"
	case 0x0019:
		return "Cycles-per-hour"
	case 0x001A:
		return "Cycles-per-minute"
	case 0x001B:
		return "Hertz"
	case 0x001C:
		return "Grams-of-water-per-kilogram-dry-air"
	case 0x001D:
		return "Percent-relative-humidity"
	case 0x001E:
		return "Millimeters"
	case 0x001F:
		return "Meters"
	case 0x0020:
		return "Inches"
	case 0x0021:
		return "Feet"
	case 0x0022:
		return "Watts-per-square-foot"
	case 0x0023:
		return "Watts-per-square-meter"
	case 0x0024:
		return "Lumens"
	case 0x0025:
		return "Luxes"
	case 0x0026:
		return "Foot-candles"
	case 0x0027:
		return "Kilograms"
	case 0x0028:
		return "Pounds-mass"
	case 0x0029:
		return "Tons"
	case 0x002A:
		return "Kilograms-per-second"
	case 0x002B:
		return "Kilograms-per-minute"
	case 0x002C:
		return "Kilograms-per-hour"
	case 0x002D:
		return "Pounds-mass-per-minute"
	case 0x002E:
		return "Pounds-mass-per-hour"
	case 0x002F:
		return "Watts"
	case 0x0030:
		return "Kilowatts"
	case 0x0031:
		return "Megawatts"
	case 0x0032:
		return "BTUs-per-hour"
	case 0x0033:
		return "Horsepower"
	case 0x0034:
		return "Tons-refrigeration"
	case 0x0035:
		return "Pascals"
	case 0x0036:
		return "Kilopascals"
	case 0x0037:
		return "Bars"
	case 0x0038:
		return "Pounds-force-per-square-inch"
	case 0x0039:
		return "Centimeters-of-water"
	case 0x003A:
		return "Inches-of-water"
	case 0x003B:
		return "Millimeters-of-mercury"
	case 0x003C:
		return "Centimeters-of-mercury"
	case 0x003D:
		return "Inches-of-mercury"
	case 0x003E:
		return "Degrees-Celsius"
	case 0x003F:
		return "Degrees-Kelvin"
	case 0x0040:
		return "Degrees-Fahrenheit"
	case 0x0041:
		return "Degree-days-Celsius"
	case 0x0042:
		return "Degree-days-Fahrenheit"
	case 0x0043:
		return "Years"
	case 0x0044:
		return "Months"
	case 0x0045:
		return "Weeks"
	case 0x0046:
		return "Days"
	case 0x0047:
		return "Hours"
	case 0x0048:
		return "Minutes"
	case 0x0049:
		return "Seconds"
	case 0x004A:
		return "Meters-per-second"
	case 0x004B:
		return "Kilometers-per-hour"
	case 0x004C:
		return "Feet-per-second"
	case 0x004D:
		return "Feet-per-minute"
	case 0x004E:
		return "Miles-per-hour"
	case 0x004F:
		return "Cubic-feet"
	case 0x0050:
		return "Cubic-meters"
	case 0x0051:
		return "Imperial-gallons"
	case 0x0052:
		return "Liters"
	case 0x0053:
		return "Us-gallons"
	case 0x0054:
		return "Cubic-feet-per-minute"
	case 0x0055:
		return "Cubic-meters-per-second"
	case 0x0056:
		return "Imperial-gallons-per-minute"
	case 0x0057:
		return "Liters-per-second"
	case 0x0058:
		return "Liters-per-minute"
	case 0x0059:
		return "Us-gallons-per-minute"
	case 0x005A:
		return "Degrees-angular"
	case 0x005B:
		return "Degrees-Celsius-per-hour"
	case 0x005C:
		return "Degrees-Celsius-per-minute"
	case 0x005D:
		return "Degrees-Fahrenheit-per-hour"
	case 0x005E:
		return "Degrees-Fahrenheit-per-minute"
	case 0x005F:
		return "No-units"
	case 0x0060:
		return "Parts-per-million"
	case 0x0061:
		return "Parts-per-billion"
	case 0x0062:
		return "Percent"
	case 0x0063:
		return "Percent-per-second"
	case 0x0064:
		return "Per-minute"
	case 0x0065:
		return "Per-second"
	case 0x0066:
		return "Psi-per-Degree-Fahrenheit"
	case 0x0067:
		return "Radians"
	case 0x0068:
		return "Revolutions-per-minute"
	case 0x0069:
		return "Currency1"
	case 0x006A:
		return "Currency2"
	case 0x006B:
		return "Currency3"
	case 0x006C:
		return "Currency4"
	case 0x006D:
		return "Currency5"
	case 0x006E:
		return "Currency6"
	case 0x006F:
		return "Currency7"
	case 0x0070:
		return "Currency8"
	case 0x0071:
		return "Currency9"
	case 0x0072:
		return "Currency10"
	case 0x0073:
		return "Square-inches"
	case 0x0074:
		return "Square-centimeters"
	case 0x0075:
		return "BTUs-per-pound"
	case 0x0076:
		return "Centimeters"
	case 0x0077:
		return "Pounds-mass-per-second"
	case 0x0078:
		return "Delta-Degrees-Fahrenheit"
	case 0x0079:
		return "Delta-Degrees-Kelvin"
	case 0x007A:
		return "Kilohms"
	case 0x007B:
		return "Megohms"
	case 0x007C:
		return "Millivolts"
	case 0x007D:
		return "Kilojoules-per-kilogram"
	case 0x007E:
		return "Megajoules"
	case 0x007F:
		return "Joules-per-degree-Kelvin"
	case 0x0080:
		return "Joules-per-kilogram-degree-Kelvin"
	case 0x0081:
		return "Kilohertz"
	case 0x0082:
		return "Megahertz"
	case 0x0083:
		return "Per-hour"
	case 0x0084:
		return "Milliwatts"
	case 0x0085:
		return "Hectopascals"
	case 0x0086:
		return "Millibars"
	case 0x0087:
		return "Cubic-meters-per-hour"
	case 0x0088:
		return "Liters-per-hour"
	case 0x0089:
		return "Kilowatt-hours-per-square-meter"
	case 0x008A:
		return "Kilowatt-hours-per-square-foot"
	case 0x008B:
		return "Megajoules-per-square-meter"
	case 0x008C:
		return "Megajoules-per-square-foot"
	case 0x008D:
		return "Watts-per-square-meter-Degree-Kelvin"
	case 0x008E:
		return "Cubic-feet-per-second"
	case 0x008F:
		return "Percent-obscuration-per-foot"
	case 0x0090:
		return "Percent-obscuration-per-meter"
	case 0x0091:
		return "Milliohms"
	case 0x0092:
		return "Megawatt-hours"
	case 0x0093:
		return "Kilo-BTUs"
	case 0x0094:
		return "Mega-BTUs"
	case 0x0095:
		return "Kilojoules-per-kilogram-dry-air"
	case 0x0096:
		return "Megajoules-per-kilogram-dry-air"
	case 0x0097:
		return "Kilojoules-per-degree-Kelvin"
	case 0x0098:
		return "Megajoules-per-degree-Kelvin"
	case 0x0099:
		return "Newton"
	case 0x009A:
		return "Grams-per-second"
	case 0x009B:
		return "Grams-per-minute"
	case 0x009C:
		return "Tons-per-hour"
	case 0x009D:
		return "Kilo-BTUs-per-hour"
	case 0x009E:
		return "Hundredths-seconds"
	case 0x009F:
		return "Milliseconds"
	case 0x00A0:
		return "Newton-meters"
	case 0x00A1:
		return "Millimeters-per-second"
	case 0x00A2:
		return "Millimeters-per-minute"
	case 0x00A3:
		return "Meters-per-minute"
	case 0x00A4:
		return "Meters-per-hour"
	case 0x00A5:
		return "Cubic-meters-per-minute"
	case 0x00A6:
		return "Meters-per-second-per-second"
	case 0x00A7:
		return "Amperes-per-meter"
	case 0x00A8:
		return "Amperes-per-square-meter"
	case 0x00A9:
		return "Ampere-square-meters"
	case 0x00AA:
		return "Farads"
	case 0x00AB:
		return "Henrys"
	case 0x00AC:
		return "Ohm-meters"
	case 0x00AD:
		return "Siemens"
	case 0x00AE:
		return "Siemens-per-meter"
	case 0x00AF:
		return "Teslas"
	case 0x00B0:
		return "Volts-per-degree-Kelvin"
	case 0x00B1:
		return "Volts-per-meter"
	case 0x00B2:
		return "Webers"
	case 0x00B3:
		return "Candelas"
	case 0x00B4:
		return "Candelas-per-square-meter"
	case 0x00B5:
		return "Kelvins-per-hour"
	case 0x00B6:
		return "Kelvins-per-minute"
	case 0x00B7:
		return "Joule-seconds"
	case 0x00B8:
		return "Square-meters-per-Newton"
	case 0x00B9:
		return "Kilogram-per-cubic-meter"
	case 0x00BA:
		return "Newton-seconds"
	case 0x00BB:
		return "Newtons-per-meter"
	case 0x00BC:
		return "Watts-per-meter-per-degree-Kelvin"
	case 0x00FF:
		return "Other"
	}

	return zcl.Sprintf("%s", zcl.Zenum16(a))
}

// IsSquareMeters checks if AnalogInputEngineeringUnits equals the value for Square-meters (0x0000)
func (a AnalogInputEngineeringUnits) IsSquareMeters() bool { return a == 0x0000 }

// SetSquareMeters sets AnalogInputEngineeringUnits to Square-meters (0x0000)
func (a *AnalogInputEngineeringUnits) SetSquareMeters() { *a = 0x0000 }

// IsSquareFeet checks if AnalogInputEngineeringUnits equals the value for Square-feet (0x0001)
func (a AnalogInputEngineeringUnits) IsSquareFeet() bool { return a == 0x0001 }

// SetSquareFeet sets AnalogInputEngineeringUnits to Square-feet (0x0001)
func (a *AnalogInputEngineeringUnits) SetSquareFeet() { *a = 0x0001 }

// IsMilliamperes checks if AnalogInputEngineeringUnits equals the value for Milliamperes (0x0002)
func (a AnalogInputEngineeringUnits) IsMilliamperes() bool { return a == 0x0002 }

// SetMilliamperes sets AnalogInputEngineeringUnits to Milliamperes (0x0002)
func (a *AnalogInputEngineeringUnits) SetMilliamperes() { *a = 0x0002 }

// IsAmperes checks if AnalogInputEngineeringUnits equals the value for Amperes (0x0003)
func (a AnalogInputEngineeringUnits) IsAmperes() bool { return a == 0x0003 }

// SetAmperes sets AnalogInputEngineeringUnits to Amperes (0x0003)
func (a *AnalogInputEngineeringUnits) SetAmperes() { *a = 0x0003 }

// IsOhms checks if AnalogInputEngineeringUnits equals the value for Ohms (0x0004)
func (a AnalogInputEngineeringUnits) IsOhms() bool { return a == 0x0004 }

// SetOhms sets AnalogInputEngineeringUnits to Ohms (0x0004)
func (a *AnalogInputEngineeringUnits) SetOhms() { *a = 0x0004 }

// IsVolts checks if AnalogInputEngineeringUnits equals the value for Volts (0x0005)
func (a AnalogInputEngineeringUnits) IsVolts() bool { return a == 0x0005 }

// SetVolts sets AnalogInputEngineeringUnits to Volts (0x0005)
func (a *AnalogInputEngineeringUnits) SetVolts() { *a = 0x0005 }

// IsKiloVolts checks if AnalogInputEngineeringUnits equals the value for Kilo-volts (0x0006)
func (a AnalogInputEngineeringUnits) IsKiloVolts() bool { return a == 0x0006 }

// SetKiloVolts sets AnalogInputEngineeringUnits to Kilo-volts (0x0006)
func (a *AnalogInputEngineeringUnits) SetKiloVolts() { *a = 0x0006 }

// IsMegaVolts checks if AnalogInputEngineeringUnits equals the value for Mega-volts (0x0007)
func (a AnalogInputEngineeringUnits) IsMegaVolts() bool { return a == 0x0007 }

// SetMegaVolts sets AnalogInputEngineeringUnits to Mega-volts (0x0007)
func (a *AnalogInputEngineeringUnits) SetMegaVolts() { *a = 0x0007 }

// IsVoltAmperes checks if AnalogInputEngineeringUnits equals the value for Volt-amperes (0x0008)
func (a AnalogInputEngineeringUnits) IsVoltAmperes() bool { return a == 0x0008 }

// SetVoltAmperes sets AnalogInputEngineeringUnits to Volt-amperes (0x0008)
func (a *AnalogInputEngineeringUnits) SetVoltAmperes() { *a = 0x0008 }

// IsKiloVoltAmperes checks if AnalogInputEngineeringUnits equals the value for Kilo-volt-amperes (0x0009)
func (a AnalogInputEngineeringUnits) IsKiloVoltAmperes() bool { return a == 0x0009 }

// SetKiloVoltAmperes sets AnalogInputEngineeringUnits to Kilo-volt-amperes (0x0009)
func (a *AnalogInputEngineeringUnits) SetKiloVoltAmperes() { *a = 0x0009 }

// IsMegaVoltAmperes checks if AnalogInputEngineeringUnits equals the value for Mega-volt-amperes (0x000A)
func (a AnalogInputEngineeringUnits) IsMegaVoltAmperes() bool { return a == 0x000A }

// SetMegaVoltAmperes sets AnalogInputEngineeringUnits to Mega-volt-amperes (0x000A)
func (a *AnalogInputEngineeringUnits) SetMegaVoltAmperes() { *a = 0x000A }

// IsVoltAmperesReactive checks if AnalogInputEngineeringUnits equals the value for Volt-amperes-reactive (0x000B)
func (a AnalogInputEngineeringUnits) IsVoltAmperesReactive() bool { return a == 0x000B }

// SetVoltAmperesReactive sets AnalogInputEngineeringUnits to Volt-amperes-reactive (0x000B)
func (a *AnalogInputEngineeringUnits) SetVoltAmperesReactive() { *a = 0x000B }

// IsKiloVoltAmperesReactive checks if AnalogInputEngineeringUnits equals the value for Kilo-volt-amperes-reactive (0x000C)
func (a AnalogInputEngineeringUnits) IsKiloVoltAmperesReactive() bool { return a == 0x000C }

// SetKiloVoltAmperesReactive sets AnalogInputEngineeringUnits to Kilo-volt-amperes-reactive (0x000C)
func (a *AnalogInputEngineeringUnits) SetKiloVoltAmperesReactive() { *a = 0x000C }

// IsMegaVoltAmperesReactive checks if AnalogInputEngineeringUnits equals the value for Mega-volt-amperes-reactive (0x000D)
func (a AnalogInputEngineeringUnits) IsMegaVoltAmperesReactive() bool { return a == 0x000D }

// SetMegaVoltAmperesReactive sets AnalogInputEngineeringUnits to Mega-volt-amperes-reactive (0x000D)
func (a *AnalogInputEngineeringUnits) SetMegaVoltAmperesReactive() { *a = 0x000D }

// IsDegreesPhase checks if AnalogInputEngineeringUnits equals the value for Degrees-phase (0x000E)
func (a AnalogInputEngineeringUnits) IsDegreesPhase() bool { return a == 0x000E }

// SetDegreesPhase sets AnalogInputEngineeringUnits to Degrees-phase (0x000E)
func (a *AnalogInputEngineeringUnits) SetDegreesPhase() { *a = 0x000E }

// IsPowerFactor checks if AnalogInputEngineeringUnits equals the value for Power-factor (0x000F)
func (a AnalogInputEngineeringUnits) IsPowerFactor() bool { return a == 0x000F }

// SetPowerFactor sets AnalogInputEngineeringUnits to Power-factor (0x000F)
func (a *AnalogInputEngineeringUnits) SetPowerFactor() { *a = 0x000F }

// IsJoules checks if AnalogInputEngineeringUnits equals the value for Joules (0x0010)
func (a AnalogInputEngineeringUnits) IsJoules() bool { return a == 0x0010 }

// SetJoules sets AnalogInputEngineeringUnits to Joules (0x0010)
func (a *AnalogInputEngineeringUnits) SetJoules() { *a = 0x0010 }

// IsKilojoules checks if AnalogInputEngineeringUnits equals the value for Kilojoules (0x0011)
func (a AnalogInputEngineeringUnits) IsKilojoules() bool { return a == 0x0011 }

// SetKilojoules sets AnalogInputEngineeringUnits to Kilojoules (0x0011)
func (a *AnalogInputEngineeringUnits) SetKilojoules() { *a = 0x0011 }

// IsWattHours checks if AnalogInputEngineeringUnits equals the value for Watt-hours (0x0012)
func (a AnalogInputEngineeringUnits) IsWattHours() bool { return a == 0x0012 }

// SetWattHours sets AnalogInputEngineeringUnits to Watt-hours (0x0012)
func (a *AnalogInputEngineeringUnits) SetWattHours() { *a = 0x0012 }

// IsKilowattHours checks if AnalogInputEngineeringUnits equals the value for Kilowatt-hours (0x0013)
func (a AnalogInputEngineeringUnits) IsKilowattHours() bool { return a == 0x0013 }

// SetKilowattHours sets AnalogInputEngineeringUnits to Kilowatt-hours (0x0013)
func (a *AnalogInputEngineeringUnits) SetKilowattHours() { *a = 0x0013 }

// IsBtus checks if AnalogInputEngineeringUnits equals the value for BTUs (0x0014)
func (a AnalogInputEngineeringUnits) IsBtus() bool { return a == 0x0014 }

// SetBtus sets AnalogInputEngineeringUnits to BTUs (0x0014)
func (a *AnalogInputEngineeringUnits) SetBtus() { *a = 0x0014 }

// IsTherms checks if AnalogInputEngineeringUnits equals the value for Therms (0x0015)
func (a AnalogInputEngineeringUnits) IsTherms() bool { return a == 0x0015 }

// SetTherms sets AnalogInputEngineeringUnits to Therms (0x0015)
func (a *AnalogInputEngineeringUnits) SetTherms() { *a = 0x0015 }

// IsTonHours checks if AnalogInputEngineeringUnits equals the value for Ton-hours (0x0016)
func (a AnalogInputEngineeringUnits) IsTonHours() bool { return a == 0x0016 }

// SetTonHours sets AnalogInputEngineeringUnits to Ton-hours (0x0016)
func (a *AnalogInputEngineeringUnits) SetTonHours() { *a = 0x0016 }

// IsJoulesPerKilogramDryAir checks if AnalogInputEngineeringUnits equals the value for Joules-per-kilogram-dry-air (0x0017)
func (a AnalogInputEngineeringUnits) IsJoulesPerKilogramDryAir() bool { return a == 0x0017 }

// SetJoulesPerKilogramDryAir sets AnalogInputEngineeringUnits to Joules-per-kilogram-dry-air (0x0017)
func (a *AnalogInputEngineeringUnits) SetJoulesPerKilogramDryAir() { *a = 0x0017 }

// IsBtusPerPoundDryAir checks if AnalogInputEngineeringUnits equals the value for BTUs-per-pound-dry-air (0x0018)
func (a AnalogInputEngineeringUnits) IsBtusPerPoundDryAir() bool { return a == 0x0018 }

// SetBtusPerPoundDryAir sets AnalogInputEngineeringUnits to BTUs-per-pound-dry-air (0x0018)
func (a *AnalogInputEngineeringUnits) SetBtusPerPoundDryAir() { *a = 0x0018 }

// IsCyclesPerHour checks if AnalogInputEngineeringUnits equals the value for Cycles-per-hour (0x0019)
func (a AnalogInputEngineeringUnits) IsCyclesPerHour() bool { return a == 0x0019 }

// SetCyclesPerHour sets AnalogInputEngineeringUnits to Cycles-per-hour (0x0019)
func (a *AnalogInputEngineeringUnits) SetCyclesPerHour() { *a = 0x0019 }

// IsCyclesPerMinute checks if AnalogInputEngineeringUnits equals the value for Cycles-per-minute (0x001A)
func (a AnalogInputEngineeringUnits) IsCyclesPerMinute() bool { return a == 0x001A }

// SetCyclesPerMinute sets AnalogInputEngineeringUnits to Cycles-per-minute (0x001A)
func (a *AnalogInputEngineeringUnits) SetCyclesPerMinute() { *a = 0x001A }

// IsHertz checks if AnalogInputEngineeringUnits equals the value for Hertz (0x001B)
func (a AnalogInputEngineeringUnits) IsHertz() bool { return a == 0x001B }

// SetHertz sets AnalogInputEngineeringUnits to Hertz (0x001B)
func (a *AnalogInputEngineeringUnits) SetHertz() { *a = 0x001B }

// IsGramsOfWaterPerKilogramDryAir checks if AnalogInputEngineeringUnits equals the value for Grams-of-water-per-kilogram-dry-air (0x001C)
func (a AnalogInputEngineeringUnits) IsGramsOfWaterPerKilogramDryAir() bool { return a == 0x001C }

// SetGramsOfWaterPerKilogramDryAir sets AnalogInputEngineeringUnits to Grams-of-water-per-kilogram-dry-air (0x001C)
func (a *AnalogInputEngineeringUnits) SetGramsOfWaterPerKilogramDryAir() { *a = 0x001C }

// IsPercentRelativeHumidity checks if AnalogInputEngineeringUnits equals the value for Percent-relative-humidity (0x001D)
func (a AnalogInputEngineeringUnits) IsPercentRelativeHumidity() bool { return a == 0x001D }

// SetPercentRelativeHumidity sets AnalogInputEngineeringUnits to Percent-relative-humidity (0x001D)
func (a *AnalogInputEngineeringUnits) SetPercentRelativeHumidity() { *a = 0x001D }

// IsMillimeters checks if AnalogInputEngineeringUnits equals the value for Millimeters (0x001E)
func (a AnalogInputEngineeringUnits) IsMillimeters() bool { return a == 0x001E }

// SetMillimeters sets AnalogInputEngineeringUnits to Millimeters (0x001E)
func (a *AnalogInputEngineeringUnits) SetMillimeters() { *a = 0x001E }

// IsMeters checks if AnalogInputEngineeringUnits equals the value for Meters (0x001F)
func (a AnalogInputEngineeringUnits) IsMeters() bool { return a == 0x001F }

// SetMeters sets AnalogInputEngineeringUnits to Meters (0x001F)
func (a *AnalogInputEngineeringUnits) SetMeters() { *a = 0x001F }

// IsInches checks if AnalogInputEngineeringUnits equals the value for Inches (0x0020)
func (a AnalogInputEngineeringUnits) IsInches() bool { return a == 0x0020 }

// SetInches sets AnalogInputEngineeringUnits to Inches (0x0020)
func (a *AnalogInputEngineeringUnits) SetInches() { *a = 0x0020 }

// IsFeet checks if AnalogInputEngineeringUnits equals the value for Feet (0x0021)
func (a AnalogInputEngineeringUnits) IsFeet() bool { return a == 0x0021 }

// SetFeet sets AnalogInputEngineeringUnits to Feet (0x0021)
func (a *AnalogInputEngineeringUnits) SetFeet() { *a = 0x0021 }

// IsWattsPerSquareFoot checks if AnalogInputEngineeringUnits equals the value for Watts-per-square-foot (0x0022)
func (a AnalogInputEngineeringUnits) IsWattsPerSquareFoot() bool { return a == 0x0022 }

// SetWattsPerSquareFoot sets AnalogInputEngineeringUnits to Watts-per-square-foot (0x0022)
func (a *AnalogInputEngineeringUnits) SetWattsPerSquareFoot() { *a = 0x0022 }

// IsWattsPerSquareMeter checks if AnalogInputEngineeringUnits equals the value for Watts-per-square-meter (0x0023)
func (a AnalogInputEngineeringUnits) IsWattsPerSquareMeter() bool { return a == 0x0023 }

// SetWattsPerSquareMeter sets AnalogInputEngineeringUnits to Watts-per-square-meter (0x0023)
func (a *AnalogInputEngineeringUnits) SetWattsPerSquareMeter() { *a = 0x0023 }

// IsLumens checks if AnalogInputEngineeringUnits equals the value for Lumens (0x0024)
func (a AnalogInputEngineeringUnits) IsLumens() bool { return a == 0x0024 }

// SetLumens sets AnalogInputEngineeringUnits to Lumens (0x0024)
func (a *AnalogInputEngineeringUnits) SetLumens() { *a = 0x0024 }

// IsLuxes checks if AnalogInputEngineeringUnits equals the value for Luxes (0x0025)
func (a AnalogInputEngineeringUnits) IsLuxes() bool { return a == 0x0025 }

// SetLuxes sets AnalogInputEngineeringUnits to Luxes (0x0025)
func (a *AnalogInputEngineeringUnits) SetLuxes() { *a = 0x0025 }

// IsFootCandles checks if AnalogInputEngineeringUnits equals the value for Foot-candles (0x0026)
func (a AnalogInputEngineeringUnits) IsFootCandles() bool { return a == 0x0026 }

// SetFootCandles sets AnalogInputEngineeringUnits to Foot-candles (0x0026)
func (a *AnalogInputEngineeringUnits) SetFootCandles() { *a = 0x0026 }

// IsKilograms checks if AnalogInputEngineeringUnits equals the value for Kilograms (0x0027)
func (a AnalogInputEngineeringUnits) IsKilograms() bool { return a == 0x0027 }

// SetKilograms sets AnalogInputEngineeringUnits to Kilograms (0x0027)
func (a *AnalogInputEngineeringUnits) SetKilograms() { *a = 0x0027 }

// IsPoundsMass checks if AnalogInputEngineeringUnits equals the value for Pounds-mass (0x0028)
func (a AnalogInputEngineeringUnits) IsPoundsMass() bool { return a == 0x0028 }

// SetPoundsMass sets AnalogInputEngineeringUnits to Pounds-mass (0x0028)
func (a *AnalogInputEngineeringUnits) SetPoundsMass() { *a = 0x0028 }

// IsTons checks if AnalogInputEngineeringUnits equals the value for Tons (0x0029)
func (a AnalogInputEngineeringUnits) IsTons() bool { return a == 0x0029 }

// SetTons sets AnalogInputEngineeringUnits to Tons (0x0029)
func (a *AnalogInputEngineeringUnits) SetTons() { *a = 0x0029 }

// IsKilogramsPerSecond checks if AnalogInputEngineeringUnits equals the value for Kilograms-per-second (0x002A)
func (a AnalogInputEngineeringUnits) IsKilogramsPerSecond() bool { return a == 0x002A }

// SetKilogramsPerSecond sets AnalogInputEngineeringUnits to Kilograms-per-second (0x002A)
func (a *AnalogInputEngineeringUnits) SetKilogramsPerSecond() { *a = 0x002A }

// IsKilogramsPerMinute checks if AnalogInputEngineeringUnits equals the value for Kilograms-per-minute (0x002B)
func (a AnalogInputEngineeringUnits) IsKilogramsPerMinute() bool { return a == 0x002B }

// SetKilogramsPerMinute sets AnalogInputEngineeringUnits to Kilograms-per-minute (0x002B)
func (a *AnalogInputEngineeringUnits) SetKilogramsPerMinute() { *a = 0x002B }

// IsKilogramsPerHour checks if AnalogInputEngineeringUnits equals the value for Kilograms-per-hour (0x002C)
func (a AnalogInputEngineeringUnits) IsKilogramsPerHour() bool { return a == 0x002C }

// SetKilogramsPerHour sets AnalogInputEngineeringUnits to Kilograms-per-hour (0x002C)
func (a *AnalogInputEngineeringUnits) SetKilogramsPerHour() { *a = 0x002C }

// IsPoundsMassPerMinute checks if AnalogInputEngineeringUnits equals the value for Pounds-mass-per-minute (0x002D)
func (a AnalogInputEngineeringUnits) IsPoundsMassPerMinute() bool { return a == 0x002D }

// SetPoundsMassPerMinute sets AnalogInputEngineeringUnits to Pounds-mass-per-minute (0x002D)
func (a *AnalogInputEngineeringUnits) SetPoundsMassPerMinute() { *a = 0x002D }

// IsPoundsMassPerHour checks if AnalogInputEngineeringUnits equals the value for Pounds-mass-per-hour (0x002E)
func (a AnalogInputEngineeringUnits) IsPoundsMassPerHour() bool { return a == 0x002E }

// SetPoundsMassPerHour sets AnalogInputEngineeringUnits to Pounds-mass-per-hour (0x002E)
func (a *AnalogInputEngineeringUnits) SetPoundsMassPerHour() { *a = 0x002E }

// IsWatts checks if AnalogInputEngineeringUnits equals the value for Watts (0x002F)
func (a AnalogInputEngineeringUnits) IsWatts() bool { return a == 0x002F }

// SetWatts sets AnalogInputEngineeringUnits to Watts (0x002F)
func (a *AnalogInputEngineeringUnits) SetWatts() { *a = 0x002F }

// IsKilowatts checks if AnalogInputEngineeringUnits equals the value for Kilowatts (0x0030)
func (a AnalogInputEngineeringUnits) IsKilowatts() bool { return a == 0x0030 }

// SetKilowatts sets AnalogInputEngineeringUnits to Kilowatts (0x0030)
func (a *AnalogInputEngineeringUnits) SetKilowatts() { *a = 0x0030 }

// IsMegawatts checks if AnalogInputEngineeringUnits equals the value for Megawatts (0x0031)
func (a AnalogInputEngineeringUnits) IsMegawatts() bool { return a == 0x0031 }

// SetMegawatts sets AnalogInputEngineeringUnits to Megawatts (0x0031)
func (a *AnalogInputEngineeringUnits) SetMegawatts() { *a = 0x0031 }

// IsBtusPerHour checks if AnalogInputEngineeringUnits equals the value for BTUs-per-hour (0x0032)
func (a AnalogInputEngineeringUnits) IsBtusPerHour() bool { return a == 0x0032 }

// SetBtusPerHour sets AnalogInputEngineeringUnits to BTUs-per-hour (0x0032)
func (a *AnalogInputEngineeringUnits) SetBtusPerHour() { *a = 0x0032 }

// IsHorsepower checks if AnalogInputEngineeringUnits equals the value for Horsepower (0x0033)
func (a AnalogInputEngineeringUnits) IsHorsepower() bool { return a == 0x0033 }

// SetHorsepower sets AnalogInputEngineeringUnits to Horsepower (0x0033)
func (a *AnalogInputEngineeringUnits) SetHorsepower() { *a = 0x0033 }

// IsTonsRefrigeration checks if AnalogInputEngineeringUnits equals the value for Tons-refrigeration (0x0034)
func (a AnalogInputEngineeringUnits) IsTonsRefrigeration() bool { return a == 0x0034 }

// SetTonsRefrigeration sets AnalogInputEngineeringUnits to Tons-refrigeration (0x0034)
func (a *AnalogInputEngineeringUnits) SetTonsRefrigeration() { *a = 0x0034 }

// IsPascals checks if AnalogInputEngineeringUnits equals the value for Pascals (0x0035)
func (a AnalogInputEngineeringUnits) IsPascals() bool { return a == 0x0035 }

// SetPascals sets AnalogInputEngineeringUnits to Pascals (0x0035)
func (a *AnalogInputEngineeringUnits) SetPascals() { *a = 0x0035 }

// IsKilopascals checks if AnalogInputEngineeringUnits equals the value for Kilopascals (0x0036)
func (a AnalogInputEngineeringUnits) IsKilopascals() bool { return a == 0x0036 }

// SetKilopascals sets AnalogInputEngineeringUnits to Kilopascals (0x0036)
func (a *AnalogInputEngineeringUnits) SetKilopascals() { *a = 0x0036 }

// IsBars checks if AnalogInputEngineeringUnits equals the value for Bars (0x0037)
func (a AnalogInputEngineeringUnits) IsBars() bool { return a == 0x0037 }

// SetBars sets AnalogInputEngineeringUnits to Bars (0x0037)
func (a *AnalogInputEngineeringUnits) SetBars() { *a = 0x0037 }

// IsPoundsForcePerSquareInch checks if AnalogInputEngineeringUnits equals the value for Pounds-force-per-square-inch (0x0038)
func (a AnalogInputEngineeringUnits) IsPoundsForcePerSquareInch() bool { return a == 0x0038 }

// SetPoundsForcePerSquareInch sets AnalogInputEngineeringUnits to Pounds-force-per-square-inch (0x0038)
func (a *AnalogInputEngineeringUnits) SetPoundsForcePerSquareInch() { *a = 0x0038 }

// IsCentimetersOfWater checks if AnalogInputEngineeringUnits equals the value for Centimeters-of-water (0x0039)
func (a AnalogInputEngineeringUnits) IsCentimetersOfWater() bool { return a == 0x0039 }

// SetCentimetersOfWater sets AnalogInputEngineeringUnits to Centimeters-of-water (0x0039)
func (a *AnalogInputEngineeringUnits) SetCentimetersOfWater() { *a = 0x0039 }

// IsInchesOfWater checks if AnalogInputEngineeringUnits equals the value for Inches-of-water (0x003A)
func (a AnalogInputEngineeringUnits) IsInchesOfWater() bool { return a == 0x003A }

// SetInchesOfWater sets AnalogInputEngineeringUnits to Inches-of-water (0x003A)
func (a *AnalogInputEngineeringUnits) SetInchesOfWater() { *a = 0x003A }

// IsMillimetersOfMercury checks if AnalogInputEngineeringUnits equals the value for Millimeters-of-mercury (0x003B)
func (a AnalogInputEngineeringUnits) IsMillimetersOfMercury() bool { return a == 0x003B }

// SetMillimetersOfMercury sets AnalogInputEngineeringUnits to Millimeters-of-mercury (0x003B)
func (a *AnalogInputEngineeringUnits) SetMillimetersOfMercury() { *a = 0x003B }

// IsCentimetersOfMercury checks if AnalogInputEngineeringUnits equals the value for Centimeters-of-mercury (0x003C)
func (a AnalogInputEngineeringUnits) IsCentimetersOfMercury() bool { return a == 0x003C }

// SetCentimetersOfMercury sets AnalogInputEngineeringUnits to Centimeters-of-mercury (0x003C)
func (a *AnalogInputEngineeringUnits) SetCentimetersOfMercury() { *a = 0x003C }

// IsInchesOfMercury checks if AnalogInputEngineeringUnits equals the value for Inches-of-mercury (0x003D)
func (a AnalogInputEngineeringUnits) IsInchesOfMercury() bool { return a == 0x003D }

// SetInchesOfMercury sets AnalogInputEngineeringUnits to Inches-of-mercury (0x003D)
func (a *AnalogInputEngineeringUnits) SetInchesOfMercury() { *a = 0x003D }

// IsDegreesCelsius checks if AnalogInputEngineeringUnits equals the value for Degrees-Celsius (0x003E)
func (a AnalogInputEngineeringUnits) IsDegreesCelsius() bool { return a == 0x003E }

// SetDegreesCelsius sets AnalogInputEngineeringUnits to Degrees-Celsius (0x003E)
func (a *AnalogInputEngineeringUnits) SetDegreesCelsius() { *a = 0x003E }

// IsDegreesKelvin checks if AnalogInputEngineeringUnits equals the value for Degrees-Kelvin (0x003F)
func (a AnalogInputEngineeringUnits) IsDegreesKelvin() bool { return a == 0x003F }

// SetDegreesKelvin sets AnalogInputEngineeringUnits to Degrees-Kelvin (0x003F)
func (a *AnalogInputEngineeringUnits) SetDegreesKelvin() { *a = 0x003F }

// IsDegreesFahrenheit checks if AnalogInputEngineeringUnits equals the value for Degrees-Fahrenheit (0x0040)
func (a AnalogInputEngineeringUnits) IsDegreesFahrenheit() bool { return a == 0x0040 }

// SetDegreesFahrenheit sets AnalogInputEngineeringUnits to Degrees-Fahrenheit (0x0040)
func (a *AnalogInputEngineeringUnits) SetDegreesFahrenheit() { *a = 0x0040 }

// IsDegreeDaysCelsius checks if AnalogInputEngineeringUnits equals the value for Degree-days-Celsius (0x0041)
func (a AnalogInputEngineeringUnits) IsDegreeDaysCelsius() bool { return a == 0x0041 }

// SetDegreeDaysCelsius sets AnalogInputEngineeringUnits to Degree-days-Celsius (0x0041)
func (a *AnalogInputEngineeringUnits) SetDegreeDaysCelsius() { *a = 0x0041 }

// IsDegreeDaysFahrenheit checks if AnalogInputEngineeringUnits equals the value for Degree-days-Fahrenheit (0x0042)
func (a AnalogInputEngineeringUnits) IsDegreeDaysFahrenheit() bool { return a == 0x0042 }

// SetDegreeDaysFahrenheit sets AnalogInputEngineeringUnits to Degree-days-Fahrenheit (0x0042)
func (a *AnalogInputEngineeringUnits) SetDegreeDaysFahrenheit() { *a = 0x0042 }

// IsYears checks if AnalogInputEngineeringUnits equals the value for Years (0x0043)
func (a AnalogInputEngineeringUnits) IsYears() bool { return a == 0x0043 }

// SetYears sets AnalogInputEngineeringUnits to Years (0x0043)
func (a *AnalogInputEngineeringUnits) SetYears() { *a = 0x0043 }

// IsMonths checks if AnalogInputEngineeringUnits equals the value for Months (0x0044)
func (a AnalogInputEngineeringUnits) IsMonths() bool { return a == 0x0044 }

// SetMonths sets AnalogInputEngineeringUnits to Months (0x0044)
func (a *AnalogInputEngineeringUnits) SetMonths() { *a = 0x0044 }

// IsWeeks checks if AnalogInputEngineeringUnits equals the value for Weeks (0x0045)
func (a AnalogInputEngineeringUnits) IsWeeks() bool { return a == 0x0045 }

// SetWeeks sets AnalogInputEngineeringUnits to Weeks (0x0045)
func (a *AnalogInputEngineeringUnits) SetWeeks() { *a = 0x0045 }

// IsDays checks if AnalogInputEngineeringUnits equals the value for Days (0x0046)
func (a AnalogInputEngineeringUnits) IsDays() bool { return a == 0x0046 }

// SetDays sets AnalogInputEngineeringUnits to Days (0x0046)
func (a *AnalogInputEngineeringUnits) SetDays() { *a = 0x0046 }

// IsHours checks if AnalogInputEngineeringUnits equals the value for Hours (0x0047)
func (a AnalogInputEngineeringUnits) IsHours() bool { return a == 0x0047 }

// SetHours sets AnalogInputEngineeringUnits to Hours (0x0047)
func (a *AnalogInputEngineeringUnits) SetHours() { *a = 0x0047 }

// IsMinutes checks if AnalogInputEngineeringUnits equals the value for Minutes (0x0048)
func (a AnalogInputEngineeringUnits) IsMinutes() bool { return a == 0x0048 }

// SetMinutes sets AnalogInputEngineeringUnits to Minutes (0x0048)
func (a *AnalogInputEngineeringUnits) SetMinutes() { *a = 0x0048 }

// IsSeconds checks if AnalogInputEngineeringUnits equals the value for Seconds (0x0049)
func (a AnalogInputEngineeringUnits) IsSeconds() bool { return a == 0x0049 }

// SetSeconds sets AnalogInputEngineeringUnits to Seconds (0x0049)
func (a *AnalogInputEngineeringUnits) SetSeconds() { *a = 0x0049 }

// IsMetersPerSecond checks if AnalogInputEngineeringUnits equals the value for Meters-per-second (0x004A)
func (a AnalogInputEngineeringUnits) IsMetersPerSecond() bool { return a == 0x004A }

// SetMetersPerSecond sets AnalogInputEngineeringUnits to Meters-per-second (0x004A)
func (a *AnalogInputEngineeringUnits) SetMetersPerSecond() { *a = 0x004A }

// IsKilometersPerHour checks if AnalogInputEngineeringUnits equals the value for Kilometers-per-hour (0x004B)
func (a AnalogInputEngineeringUnits) IsKilometersPerHour() bool { return a == 0x004B }

// SetKilometersPerHour sets AnalogInputEngineeringUnits to Kilometers-per-hour (0x004B)
func (a *AnalogInputEngineeringUnits) SetKilometersPerHour() { *a = 0x004B }

// IsFeetPerSecond checks if AnalogInputEngineeringUnits equals the value for Feet-per-second (0x004C)
func (a AnalogInputEngineeringUnits) IsFeetPerSecond() bool { return a == 0x004C }

// SetFeetPerSecond sets AnalogInputEngineeringUnits to Feet-per-second (0x004C)
func (a *AnalogInputEngineeringUnits) SetFeetPerSecond() { *a = 0x004C }

// IsFeetPerMinute checks if AnalogInputEngineeringUnits equals the value for Feet-per-minute (0x004D)
func (a AnalogInputEngineeringUnits) IsFeetPerMinute() bool { return a == 0x004D }

// SetFeetPerMinute sets AnalogInputEngineeringUnits to Feet-per-minute (0x004D)
func (a *AnalogInputEngineeringUnits) SetFeetPerMinute() { *a = 0x004D }

// IsMilesPerHour checks if AnalogInputEngineeringUnits equals the value for Miles-per-hour (0x004E)
func (a AnalogInputEngineeringUnits) IsMilesPerHour() bool { return a == 0x004E }

// SetMilesPerHour sets AnalogInputEngineeringUnits to Miles-per-hour (0x004E)
func (a *AnalogInputEngineeringUnits) SetMilesPerHour() { *a = 0x004E }

// IsCubicFeet checks if AnalogInputEngineeringUnits equals the value for Cubic-feet (0x004F)
func (a AnalogInputEngineeringUnits) IsCubicFeet() bool { return a == 0x004F }

// SetCubicFeet sets AnalogInputEngineeringUnits to Cubic-feet (0x004F)
func (a *AnalogInputEngineeringUnits) SetCubicFeet() { *a = 0x004F }

// IsCubicMeters checks if AnalogInputEngineeringUnits equals the value for Cubic-meters (0x0050)
func (a AnalogInputEngineeringUnits) IsCubicMeters() bool { return a == 0x0050 }

// SetCubicMeters sets AnalogInputEngineeringUnits to Cubic-meters (0x0050)
func (a *AnalogInputEngineeringUnits) SetCubicMeters() { *a = 0x0050 }

// IsImperialGallons checks if AnalogInputEngineeringUnits equals the value for Imperial-gallons (0x0051)
func (a AnalogInputEngineeringUnits) IsImperialGallons() bool { return a == 0x0051 }

// SetImperialGallons sets AnalogInputEngineeringUnits to Imperial-gallons (0x0051)
func (a *AnalogInputEngineeringUnits) SetImperialGallons() { *a = 0x0051 }

// IsLiters checks if AnalogInputEngineeringUnits equals the value for Liters (0x0052)
func (a AnalogInputEngineeringUnits) IsLiters() bool { return a == 0x0052 }

// SetLiters sets AnalogInputEngineeringUnits to Liters (0x0052)
func (a *AnalogInputEngineeringUnits) SetLiters() { *a = 0x0052 }

// IsUsGallons checks if AnalogInputEngineeringUnits equals the value for Us-gallons (0x0053)
func (a AnalogInputEngineeringUnits) IsUsGallons() bool { return a == 0x0053 }

// SetUsGallons sets AnalogInputEngineeringUnits to Us-gallons (0x0053)
func (a *AnalogInputEngineeringUnits) SetUsGallons() { *a = 0x0053 }

// IsCubicFeetPerMinute checks if AnalogInputEngineeringUnits equals the value for Cubic-feet-per-minute (0x0054)
func (a AnalogInputEngineeringUnits) IsCubicFeetPerMinute() bool { return a == 0x0054 }

// SetCubicFeetPerMinute sets AnalogInputEngineeringUnits to Cubic-feet-per-minute (0x0054)
func (a *AnalogInputEngineeringUnits) SetCubicFeetPerMinute() { *a = 0x0054 }

// IsCubicMetersPerSecond checks if AnalogInputEngineeringUnits equals the value for Cubic-meters-per-second (0x0055)
func (a AnalogInputEngineeringUnits) IsCubicMetersPerSecond() bool { return a == 0x0055 }

// SetCubicMetersPerSecond sets AnalogInputEngineeringUnits to Cubic-meters-per-second (0x0055)
func (a *AnalogInputEngineeringUnits) SetCubicMetersPerSecond() { *a = 0x0055 }

// IsImperialGallonsPerMinute checks if AnalogInputEngineeringUnits equals the value for Imperial-gallons-per-minute (0x0056)
func (a AnalogInputEngineeringUnits) IsImperialGallonsPerMinute() bool { return a == 0x0056 }

// SetImperialGallonsPerMinute sets AnalogInputEngineeringUnits to Imperial-gallons-per-minute (0x0056)
func (a *AnalogInputEngineeringUnits) SetImperialGallonsPerMinute() { *a = 0x0056 }

// IsLitersPerSecond checks if AnalogInputEngineeringUnits equals the value for Liters-per-second (0x0057)
func (a AnalogInputEngineeringUnits) IsLitersPerSecond() bool { return a == 0x0057 }

// SetLitersPerSecond sets AnalogInputEngineeringUnits to Liters-per-second (0x0057)
func (a *AnalogInputEngineeringUnits) SetLitersPerSecond() { *a = 0x0057 }

// IsLitersPerMinute checks if AnalogInputEngineeringUnits equals the value for Liters-per-minute (0x0058)
func (a AnalogInputEngineeringUnits) IsLitersPerMinute() bool { return a == 0x0058 }

// SetLitersPerMinute sets AnalogInputEngineeringUnits to Liters-per-minute (0x0058)
func (a *AnalogInputEngineeringUnits) SetLitersPerMinute() { *a = 0x0058 }

// IsUsGallonsPerMinute checks if AnalogInputEngineeringUnits equals the value for Us-gallons-per-minute (0x0059)
func (a AnalogInputEngineeringUnits) IsUsGallonsPerMinute() bool { return a == 0x0059 }

// SetUsGallonsPerMinute sets AnalogInputEngineeringUnits to Us-gallons-per-minute (0x0059)
func (a *AnalogInputEngineeringUnits) SetUsGallonsPerMinute() { *a = 0x0059 }

// IsDegreesAngular checks if AnalogInputEngineeringUnits equals the value for Degrees-angular (0x005A)
func (a AnalogInputEngineeringUnits) IsDegreesAngular() bool { return a == 0x005A }

// SetDegreesAngular sets AnalogInputEngineeringUnits to Degrees-angular (0x005A)
func (a *AnalogInputEngineeringUnits) SetDegreesAngular() { *a = 0x005A }

// IsDegreesCelsiusPerHour checks if AnalogInputEngineeringUnits equals the value for Degrees-Celsius-per-hour (0x005B)
func (a AnalogInputEngineeringUnits) IsDegreesCelsiusPerHour() bool { return a == 0x005B }

// SetDegreesCelsiusPerHour sets AnalogInputEngineeringUnits to Degrees-Celsius-per-hour (0x005B)
func (a *AnalogInputEngineeringUnits) SetDegreesCelsiusPerHour() { *a = 0x005B }

// IsDegreesCelsiusPerMinute checks if AnalogInputEngineeringUnits equals the value for Degrees-Celsius-per-minute (0x005C)
func (a AnalogInputEngineeringUnits) IsDegreesCelsiusPerMinute() bool { return a == 0x005C }

// SetDegreesCelsiusPerMinute sets AnalogInputEngineeringUnits to Degrees-Celsius-per-minute (0x005C)
func (a *AnalogInputEngineeringUnits) SetDegreesCelsiusPerMinute() { *a = 0x005C }

// IsDegreesFahrenheitPerHour checks if AnalogInputEngineeringUnits equals the value for Degrees-Fahrenheit-per-hour (0x005D)
func (a AnalogInputEngineeringUnits) IsDegreesFahrenheitPerHour() bool { return a == 0x005D }

// SetDegreesFahrenheitPerHour sets AnalogInputEngineeringUnits to Degrees-Fahrenheit-per-hour (0x005D)
func (a *AnalogInputEngineeringUnits) SetDegreesFahrenheitPerHour() { *a = 0x005D }

// IsDegreesFahrenheitPerMinute checks if AnalogInputEngineeringUnits equals the value for Degrees-Fahrenheit-per-minute (0x005E)
func (a AnalogInputEngineeringUnits) IsDegreesFahrenheitPerMinute() bool { return a == 0x005E }

// SetDegreesFahrenheitPerMinute sets AnalogInputEngineeringUnits to Degrees-Fahrenheit-per-minute (0x005E)
func (a *AnalogInputEngineeringUnits) SetDegreesFahrenheitPerMinute() { *a = 0x005E }

// IsNoUnits checks if AnalogInputEngineeringUnits equals the value for No-units (0x005F)
func (a AnalogInputEngineeringUnits) IsNoUnits() bool { return a == 0x005F }

// SetNoUnits sets AnalogInputEngineeringUnits to No-units (0x005F)
func (a *AnalogInputEngineeringUnits) SetNoUnits() { *a = 0x005F }

// IsPartsPerMillion checks if AnalogInputEngineeringUnits equals the value for Parts-per-million (0x0060)
func (a AnalogInputEngineeringUnits) IsPartsPerMillion() bool { return a == 0x0060 }

// SetPartsPerMillion sets AnalogInputEngineeringUnits to Parts-per-million (0x0060)
func (a *AnalogInputEngineeringUnits) SetPartsPerMillion() { *a = 0x0060 }

// IsPartsPerBillion checks if AnalogInputEngineeringUnits equals the value for Parts-per-billion (0x0061)
func (a AnalogInputEngineeringUnits) IsPartsPerBillion() bool { return a == 0x0061 }

// SetPartsPerBillion sets AnalogInputEngineeringUnits to Parts-per-billion (0x0061)
func (a *AnalogInputEngineeringUnits) SetPartsPerBillion() { *a = 0x0061 }

// IsPercent checks if AnalogInputEngineeringUnits equals the value for Percent (0x0062)
func (a AnalogInputEngineeringUnits) IsPercent() bool { return a == 0x0062 }

// SetPercent sets AnalogInputEngineeringUnits to Percent (0x0062)
func (a *AnalogInputEngineeringUnits) SetPercent() { *a = 0x0062 }

// IsPercentPerSecond checks if AnalogInputEngineeringUnits equals the value for Percent-per-second (0x0063)
func (a AnalogInputEngineeringUnits) IsPercentPerSecond() bool { return a == 0x0063 }

// SetPercentPerSecond sets AnalogInputEngineeringUnits to Percent-per-second (0x0063)
func (a *AnalogInputEngineeringUnits) SetPercentPerSecond() { *a = 0x0063 }

// IsPerMinute checks if AnalogInputEngineeringUnits equals the value for Per-minute (0x0064)
func (a AnalogInputEngineeringUnits) IsPerMinute() bool { return a == 0x0064 }

// SetPerMinute sets AnalogInputEngineeringUnits to Per-minute (0x0064)
func (a *AnalogInputEngineeringUnits) SetPerMinute() { *a = 0x0064 }

// IsPerSecond checks if AnalogInputEngineeringUnits equals the value for Per-second (0x0065)
func (a AnalogInputEngineeringUnits) IsPerSecond() bool { return a == 0x0065 }

// SetPerSecond sets AnalogInputEngineeringUnits to Per-second (0x0065)
func (a *AnalogInputEngineeringUnits) SetPerSecond() { *a = 0x0065 }

// IsPsiPerDegreeFahrenheit checks if AnalogInputEngineeringUnits equals the value for Psi-per-Degree-Fahrenheit (0x0066)
func (a AnalogInputEngineeringUnits) IsPsiPerDegreeFahrenheit() bool { return a == 0x0066 }

// SetPsiPerDegreeFahrenheit sets AnalogInputEngineeringUnits to Psi-per-Degree-Fahrenheit (0x0066)
func (a *AnalogInputEngineeringUnits) SetPsiPerDegreeFahrenheit() { *a = 0x0066 }

// IsRadians checks if AnalogInputEngineeringUnits equals the value for Radians (0x0067)
func (a AnalogInputEngineeringUnits) IsRadians() bool { return a == 0x0067 }

// SetRadians sets AnalogInputEngineeringUnits to Radians (0x0067)
func (a *AnalogInputEngineeringUnits) SetRadians() { *a = 0x0067 }

// IsRevolutionsPerMinute checks if AnalogInputEngineeringUnits equals the value for Revolutions-per-minute (0x0068)
func (a AnalogInputEngineeringUnits) IsRevolutionsPerMinute() bool { return a == 0x0068 }

// SetRevolutionsPerMinute sets AnalogInputEngineeringUnits to Revolutions-per-minute (0x0068)
func (a *AnalogInputEngineeringUnits) SetRevolutionsPerMinute() { *a = 0x0068 }

// IsCurrency1 checks if AnalogInputEngineeringUnits equals the value for Currency1 (0x0069)
func (a AnalogInputEngineeringUnits) IsCurrency1() bool { return a == 0x0069 }

// SetCurrency1 sets AnalogInputEngineeringUnits to Currency1 (0x0069)
func (a *AnalogInputEngineeringUnits) SetCurrency1() { *a = 0x0069 }

// IsCurrency2 checks if AnalogInputEngineeringUnits equals the value for Currency2 (0x006A)
func (a AnalogInputEngineeringUnits) IsCurrency2() bool { return a == 0x006A }

// SetCurrency2 sets AnalogInputEngineeringUnits to Currency2 (0x006A)
func (a *AnalogInputEngineeringUnits) SetCurrency2() { *a = 0x006A }

// IsCurrency3 checks if AnalogInputEngineeringUnits equals the value for Currency3 (0x006B)
func (a AnalogInputEngineeringUnits) IsCurrency3() bool { return a == 0x006B }

// SetCurrency3 sets AnalogInputEngineeringUnits to Currency3 (0x006B)
func (a *AnalogInputEngineeringUnits) SetCurrency3() { *a = 0x006B }

// IsCurrency4 checks if AnalogInputEngineeringUnits equals the value for Currency4 (0x006C)
func (a AnalogInputEngineeringUnits) IsCurrency4() bool { return a == 0x006C }

// SetCurrency4 sets AnalogInputEngineeringUnits to Currency4 (0x006C)
func (a *AnalogInputEngineeringUnits) SetCurrency4() { *a = 0x006C }

// IsCurrency5 checks if AnalogInputEngineeringUnits equals the value for Currency5 (0x006D)
func (a AnalogInputEngineeringUnits) IsCurrency5() bool { return a == 0x006D }

// SetCurrency5 sets AnalogInputEngineeringUnits to Currency5 (0x006D)
func (a *AnalogInputEngineeringUnits) SetCurrency5() { *a = 0x006D }

// IsCurrency6 checks if AnalogInputEngineeringUnits equals the value for Currency6 (0x006E)
func (a AnalogInputEngineeringUnits) IsCurrency6() bool { return a == 0x006E }

// SetCurrency6 sets AnalogInputEngineeringUnits to Currency6 (0x006E)
func (a *AnalogInputEngineeringUnits) SetCurrency6() { *a = 0x006E }

// IsCurrency7 checks if AnalogInputEngineeringUnits equals the value for Currency7 (0x006F)
func (a AnalogInputEngineeringUnits) IsCurrency7() bool { return a == 0x006F }

// SetCurrency7 sets AnalogInputEngineeringUnits to Currency7 (0x006F)
func (a *AnalogInputEngineeringUnits) SetCurrency7() { *a = 0x006F }

// IsCurrency8 checks if AnalogInputEngineeringUnits equals the value for Currency8 (0x0070)
func (a AnalogInputEngineeringUnits) IsCurrency8() bool { return a == 0x0070 }

// SetCurrency8 sets AnalogInputEngineeringUnits to Currency8 (0x0070)
func (a *AnalogInputEngineeringUnits) SetCurrency8() { *a = 0x0070 }

// IsCurrency9 checks if AnalogInputEngineeringUnits equals the value for Currency9 (0x0071)
func (a AnalogInputEngineeringUnits) IsCurrency9() bool { return a == 0x0071 }

// SetCurrency9 sets AnalogInputEngineeringUnits to Currency9 (0x0071)
func (a *AnalogInputEngineeringUnits) SetCurrency9() { *a = 0x0071 }

// IsCurrency10 checks if AnalogInputEngineeringUnits equals the value for Currency10 (0x0072)
func (a AnalogInputEngineeringUnits) IsCurrency10() bool { return a == 0x0072 }

// SetCurrency10 sets AnalogInputEngineeringUnits to Currency10 (0x0072)
func (a *AnalogInputEngineeringUnits) SetCurrency10() { *a = 0x0072 }

// IsSquareInches checks if AnalogInputEngineeringUnits equals the value for Square-inches (0x0073)
func (a AnalogInputEngineeringUnits) IsSquareInches() bool { return a == 0x0073 }

// SetSquareInches sets AnalogInputEngineeringUnits to Square-inches (0x0073)
func (a *AnalogInputEngineeringUnits) SetSquareInches() { *a = 0x0073 }

// IsSquareCentimeters checks if AnalogInputEngineeringUnits equals the value for Square-centimeters (0x0074)
func (a AnalogInputEngineeringUnits) IsSquareCentimeters() bool { return a == 0x0074 }

// SetSquareCentimeters sets AnalogInputEngineeringUnits to Square-centimeters (0x0074)
func (a *AnalogInputEngineeringUnits) SetSquareCentimeters() { *a = 0x0074 }

// IsBtusPerPound checks if AnalogInputEngineeringUnits equals the value for BTUs-per-pound (0x0075)
func (a AnalogInputEngineeringUnits) IsBtusPerPound() bool { return a == 0x0075 }

// SetBtusPerPound sets AnalogInputEngineeringUnits to BTUs-per-pound (0x0075)
func (a *AnalogInputEngineeringUnits) SetBtusPerPound() { *a = 0x0075 }

// IsCentimeters checks if AnalogInputEngineeringUnits equals the value for Centimeters (0x0076)
func (a AnalogInputEngineeringUnits) IsCentimeters() bool { return a == 0x0076 }

// SetCentimeters sets AnalogInputEngineeringUnits to Centimeters (0x0076)
func (a *AnalogInputEngineeringUnits) SetCentimeters() { *a = 0x0076 }

// IsPoundsMassPerSecond checks if AnalogInputEngineeringUnits equals the value for Pounds-mass-per-second (0x0077)
func (a AnalogInputEngineeringUnits) IsPoundsMassPerSecond() bool { return a == 0x0077 }

// SetPoundsMassPerSecond sets AnalogInputEngineeringUnits to Pounds-mass-per-second (0x0077)
func (a *AnalogInputEngineeringUnits) SetPoundsMassPerSecond() { *a = 0x0077 }

// IsDeltaDegreesFahrenheit checks if AnalogInputEngineeringUnits equals the value for Delta-Degrees-Fahrenheit (0x0078)
func (a AnalogInputEngineeringUnits) IsDeltaDegreesFahrenheit() bool { return a == 0x0078 }

// SetDeltaDegreesFahrenheit sets AnalogInputEngineeringUnits to Delta-Degrees-Fahrenheit (0x0078)
func (a *AnalogInputEngineeringUnits) SetDeltaDegreesFahrenheit() { *a = 0x0078 }

// IsDeltaDegreesKelvin checks if AnalogInputEngineeringUnits equals the value for Delta-Degrees-Kelvin (0x0079)
func (a AnalogInputEngineeringUnits) IsDeltaDegreesKelvin() bool { return a == 0x0079 }

// SetDeltaDegreesKelvin sets AnalogInputEngineeringUnits to Delta-Degrees-Kelvin (0x0079)
func (a *AnalogInputEngineeringUnits) SetDeltaDegreesKelvin() { *a = 0x0079 }

// IsKilohms checks if AnalogInputEngineeringUnits equals the value for Kilohms (0x007A)
func (a AnalogInputEngineeringUnits) IsKilohms() bool { return a == 0x007A }

// SetKilohms sets AnalogInputEngineeringUnits to Kilohms (0x007A)
func (a *AnalogInputEngineeringUnits) SetKilohms() { *a = 0x007A }

// IsMegohms checks if AnalogInputEngineeringUnits equals the value for Megohms (0x007B)
func (a AnalogInputEngineeringUnits) IsMegohms() bool { return a == 0x007B }

// SetMegohms sets AnalogInputEngineeringUnits to Megohms (0x007B)
func (a *AnalogInputEngineeringUnits) SetMegohms() { *a = 0x007B }

// IsMillivolts checks if AnalogInputEngineeringUnits equals the value for Millivolts (0x007C)
func (a AnalogInputEngineeringUnits) IsMillivolts() bool { return a == 0x007C }

// SetMillivolts sets AnalogInputEngineeringUnits to Millivolts (0x007C)
func (a *AnalogInputEngineeringUnits) SetMillivolts() { *a = 0x007C }

// IsKilojoulesPerKilogram checks if AnalogInputEngineeringUnits equals the value for Kilojoules-per-kilogram (0x007D)
func (a AnalogInputEngineeringUnits) IsKilojoulesPerKilogram() bool { return a == 0x007D }

// SetKilojoulesPerKilogram sets AnalogInputEngineeringUnits to Kilojoules-per-kilogram (0x007D)
func (a *AnalogInputEngineeringUnits) SetKilojoulesPerKilogram() { *a = 0x007D }

// IsMegajoules checks if AnalogInputEngineeringUnits equals the value for Megajoules (0x007E)
func (a AnalogInputEngineeringUnits) IsMegajoules() bool { return a == 0x007E }

// SetMegajoules sets AnalogInputEngineeringUnits to Megajoules (0x007E)
func (a *AnalogInputEngineeringUnits) SetMegajoules() { *a = 0x007E }

// IsJoulesPerDegreeKelvin checks if AnalogInputEngineeringUnits equals the value for Joules-per-degree-Kelvin (0x007F)
func (a AnalogInputEngineeringUnits) IsJoulesPerDegreeKelvin() bool { return a == 0x007F }

// SetJoulesPerDegreeKelvin sets AnalogInputEngineeringUnits to Joules-per-degree-Kelvin (0x007F)
func (a *AnalogInputEngineeringUnits) SetJoulesPerDegreeKelvin() { *a = 0x007F }

// IsJoulesPerKilogramDegreeKelvin checks if AnalogInputEngineeringUnits equals the value for Joules-per-kilogram-degree-Kelvin (0x0080)
func (a AnalogInputEngineeringUnits) IsJoulesPerKilogramDegreeKelvin() bool { return a == 0x0080 }

// SetJoulesPerKilogramDegreeKelvin sets AnalogInputEngineeringUnits to Joules-per-kilogram-degree-Kelvin (0x0080)
func (a *AnalogInputEngineeringUnits) SetJoulesPerKilogramDegreeKelvin() { *a = 0x0080 }

// IsKilohertz checks if AnalogInputEngineeringUnits equals the value for Kilohertz (0x0081)
func (a AnalogInputEngineeringUnits) IsKilohertz() bool { return a == 0x0081 }

// SetKilohertz sets AnalogInputEngineeringUnits to Kilohertz (0x0081)
func (a *AnalogInputEngineeringUnits) SetKilohertz() { *a = 0x0081 }

// IsMegahertz checks if AnalogInputEngineeringUnits equals the value for Megahertz (0x0082)
func (a AnalogInputEngineeringUnits) IsMegahertz() bool { return a == 0x0082 }

// SetMegahertz sets AnalogInputEngineeringUnits to Megahertz (0x0082)
func (a *AnalogInputEngineeringUnits) SetMegahertz() { *a = 0x0082 }

// IsPerHour checks if AnalogInputEngineeringUnits equals the value for Per-hour (0x0083)
func (a AnalogInputEngineeringUnits) IsPerHour() bool { return a == 0x0083 }

// SetPerHour sets AnalogInputEngineeringUnits to Per-hour (0x0083)
func (a *AnalogInputEngineeringUnits) SetPerHour() { *a = 0x0083 }

// IsMilliwatts checks if AnalogInputEngineeringUnits equals the value for Milliwatts (0x0084)
func (a AnalogInputEngineeringUnits) IsMilliwatts() bool { return a == 0x0084 }

// SetMilliwatts sets AnalogInputEngineeringUnits to Milliwatts (0x0084)
func (a *AnalogInputEngineeringUnits) SetMilliwatts() { *a = 0x0084 }

// IsHectopascals checks if AnalogInputEngineeringUnits equals the value for Hectopascals (0x0085)
func (a AnalogInputEngineeringUnits) IsHectopascals() bool { return a == 0x0085 }

// SetHectopascals sets AnalogInputEngineeringUnits to Hectopascals (0x0085)
func (a *AnalogInputEngineeringUnits) SetHectopascals() { *a = 0x0085 }

// IsMillibars checks if AnalogInputEngineeringUnits equals the value for Millibars (0x0086)
func (a AnalogInputEngineeringUnits) IsMillibars() bool { return a == 0x0086 }

// SetMillibars sets AnalogInputEngineeringUnits to Millibars (0x0086)
func (a *AnalogInputEngineeringUnits) SetMillibars() { *a = 0x0086 }

// IsCubicMetersPerHour checks if AnalogInputEngineeringUnits equals the value for Cubic-meters-per-hour (0x0087)
func (a AnalogInputEngineeringUnits) IsCubicMetersPerHour() bool { return a == 0x0087 }

// SetCubicMetersPerHour sets AnalogInputEngineeringUnits to Cubic-meters-per-hour (0x0087)
func (a *AnalogInputEngineeringUnits) SetCubicMetersPerHour() { *a = 0x0087 }

// IsLitersPerHour checks if AnalogInputEngineeringUnits equals the value for Liters-per-hour (0x0088)
func (a AnalogInputEngineeringUnits) IsLitersPerHour() bool { return a == 0x0088 }

// SetLitersPerHour sets AnalogInputEngineeringUnits to Liters-per-hour (0x0088)
func (a *AnalogInputEngineeringUnits) SetLitersPerHour() { *a = 0x0088 }

// IsKilowattHoursPerSquareMeter checks if AnalogInputEngineeringUnits equals the value for Kilowatt-hours-per-square-meter (0x0089)
func (a AnalogInputEngineeringUnits) IsKilowattHoursPerSquareMeter() bool { return a == 0x0089 }

// SetKilowattHoursPerSquareMeter sets AnalogInputEngineeringUnits to Kilowatt-hours-per-square-meter (0x0089)
func (a *AnalogInputEngineeringUnits) SetKilowattHoursPerSquareMeter() { *a = 0x0089 }

// IsKilowattHoursPerSquareFoot checks if AnalogInputEngineeringUnits equals the value for Kilowatt-hours-per-square-foot (0x008A)
func (a AnalogInputEngineeringUnits) IsKilowattHoursPerSquareFoot() bool { return a == 0x008A }

// SetKilowattHoursPerSquareFoot sets AnalogInputEngineeringUnits to Kilowatt-hours-per-square-foot (0x008A)
func (a *AnalogInputEngineeringUnits) SetKilowattHoursPerSquareFoot() { *a = 0x008A }

// IsMegajoulesPerSquareMeter checks if AnalogInputEngineeringUnits equals the value for Megajoules-per-square-meter (0x008B)
func (a AnalogInputEngineeringUnits) IsMegajoulesPerSquareMeter() bool { return a == 0x008B }

// SetMegajoulesPerSquareMeter sets AnalogInputEngineeringUnits to Megajoules-per-square-meter (0x008B)
func (a *AnalogInputEngineeringUnits) SetMegajoulesPerSquareMeter() { *a = 0x008B }

// IsMegajoulesPerSquareFoot checks if AnalogInputEngineeringUnits equals the value for Megajoules-per-square-foot (0x008C)
func (a AnalogInputEngineeringUnits) IsMegajoulesPerSquareFoot() bool { return a == 0x008C }

// SetMegajoulesPerSquareFoot sets AnalogInputEngineeringUnits to Megajoules-per-square-foot (0x008C)
func (a *AnalogInputEngineeringUnits) SetMegajoulesPerSquareFoot() { *a = 0x008C }

// IsWattsPerSquareMeterDegreeKelvin checks if AnalogInputEngineeringUnits equals the value for Watts-per-square-meter-Degree-Kelvin (0x008D)
func (a AnalogInputEngineeringUnits) IsWattsPerSquareMeterDegreeKelvin() bool { return a == 0x008D }

// SetWattsPerSquareMeterDegreeKelvin sets AnalogInputEngineeringUnits to Watts-per-square-meter-Degree-Kelvin (0x008D)
func (a *AnalogInputEngineeringUnits) SetWattsPerSquareMeterDegreeKelvin() { *a = 0x008D }

// IsCubicFeetPerSecond checks if AnalogInputEngineeringUnits equals the value for Cubic-feet-per-second (0x008E)
func (a AnalogInputEngineeringUnits) IsCubicFeetPerSecond() bool { return a == 0x008E }

// SetCubicFeetPerSecond sets AnalogInputEngineeringUnits to Cubic-feet-per-second (0x008E)
func (a *AnalogInputEngineeringUnits) SetCubicFeetPerSecond() { *a = 0x008E }

// IsPercentObscurationPerFoot checks if AnalogInputEngineeringUnits equals the value for Percent-obscuration-per-foot (0x008F)
func (a AnalogInputEngineeringUnits) IsPercentObscurationPerFoot() bool { return a == 0x008F }

// SetPercentObscurationPerFoot sets AnalogInputEngineeringUnits to Percent-obscuration-per-foot (0x008F)
func (a *AnalogInputEngineeringUnits) SetPercentObscurationPerFoot() { *a = 0x008F }

// IsPercentObscurationPerMeter checks if AnalogInputEngineeringUnits equals the value for Percent-obscuration-per-meter (0x0090)
func (a AnalogInputEngineeringUnits) IsPercentObscurationPerMeter() bool { return a == 0x0090 }

// SetPercentObscurationPerMeter sets AnalogInputEngineeringUnits to Percent-obscuration-per-meter (0x0090)
func (a *AnalogInputEngineeringUnits) SetPercentObscurationPerMeter() { *a = 0x0090 }

// IsMilliohms checks if AnalogInputEngineeringUnits equals the value for Milliohms (0x0091)
func (a AnalogInputEngineeringUnits) IsMilliohms() bool { return a == 0x0091 }

// SetMilliohms sets AnalogInputEngineeringUnits to Milliohms (0x0091)
func (a *AnalogInputEngineeringUnits) SetMilliohms() { *a = 0x0091 }

// IsMegawattHours checks if AnalogInputEngineeringUnits equals the value for Megawatt-hours (0x0092)
func (a AnalogInputEngineeringUnits) IsMegawattHours() bool { return a == 0x0092 }

// SetMegawattHours sets AnalogInputEngineeringUnits to Megawatt-hours (0x0092)
func (a *AnalogInputEngineeringUnits) SetMegawattHours() { *a = 0x0092 }

// IsKiloBtus checks if AnalogInputEngineeringUnits equals the value for Kilo-BTUs (0x0093)
func (a AnalogInputEngineeringUnits) IsKiloBtus() bool { return a == 0x0093 }

// SetKiloBtus sets AnalogInputEngineeringUnits to Kilo-BTUs (0x0093)
func (a *AnalogInputEngineeringUnits) SetKiloBtus() { *a = 0x0093 }

// IsMegaBtus checks if AnalogInputEngineeringUnits equals the value for Mega-BTUs (0x0094)
func (a AnalogInputEngineeringUnits) IsMegaBtus() bool { return a == 0x0094 }

// SetMegaBtus sets AnalogInputEngineeringUnits to Mega-BTUs (0x0094)
func (a *AnalogInputEngineeringUnits) SetMegaBtus() { *a = 0x0094 }

// IsKilojoulesPerKilogramDryAir checks if AnalogInputEngineeringUnits equals the value for Kilojoules-per-kilogram-dry-air (0x0095)
func (a AnalogInputEngineeringUnits) IsKilojoulesPerKilogramDryAir() bool { return a == 0x0095 }

// SetKilojoulesPerKilogramDryAir sets AnalogInputEngineeringUnits to Kilojoules-per-kilogram-dry-air (0x0095)
func (a *AnalogInputEngineeringUnits) SetKilojoulesPerKilogramDryAir() { *a = 0x0095 }

// IsMegajoulesPerKilogramDryAir checks if AnalogInputEngineeringUnits equals the value for Megajoules-per-kilogram-dry-air (0x0096)
func (a AnalogInputEngineeringUnits) IsMegajoulesPerKilogramDryAir() bool { return a == 0x0096 }

// SetMegajoulesPerKilogramDryAir sets AnalogInputEngineeringUnits to Megajoules-per-kilogram-dry-air (0x0096)
func (a *AnalogInputEngineeringUnits) SetMegajoulesPerKilogramDryAir() { *a = 0x0096 }

// IsKilojoulesPerDegreeKelvin checks if AnalogInputEngineeringUnits equals the value for Kilojoules-per-degree-Kelvin (0x0097)
func (a AnalogInputEngineeringUnits) IsKilojoulesPerDegreeKelvin() bool { return a == 0x0097 }

// SetKilojoulesPerDegreeKelvin sets AnalogInputEngineeringUnits to Kilojoules-per-degree-Kelvin (0x0097)
func (a *AnalogInputEngineeringUnits) SetKilojoulesPerDegreeKelvin() { *a = 0x0097 }

// IsMegajoulesPerDegreeKelvin checks if AnalogInputEngineeringUnits equals the value for Megajoules-per-degree-Kelvin (0x0098)
func (a AnalogInputEngineeringUnits) IsMegajoulesPerDegreeKelvin() bool { return a == 0x0098 }

// SetMegajoulesPerDegreeKelvin sets AnalogInputEngineeringUnits to Megajoules-per-degree-Kelvin (0x0098)
func (a *AnalogInputEngineeringUnits) SetMegajoulesPerDegreeKelvin() { *a = 0x0098 }

// IsNewton checks if AnalogInputEngineeringUnits equals the value for Newton (0x0099)
func (a AnalogInputEngineeringUnits) IsNewton() bool { return a == 0x0099 }

// SetNewton sets AnalogInputEngineeringUnits to Newton (0x0099)
func (a *AnalogInputEngineeringUnits) SetNewton() { *a = 0x0099 }

// IsGramsPerSecond checks if AnalogInputEngineeringUnits equals the value for Grams-per-second (0x009A)
func (a AnalogInputEngineeringUnits) IsGramsPerSecond() bool { return a == 0x009A }

// SetGramsPerSecond sets AnalogInputEngineeringUnits to Grams-per-second (0x009A)
func (a *AnalogInputEngineeringUnits) SetGramsPerSecond() { *a = 0x009A }

// IsGramsPerMinute checks if AnalogInputEngineeringUnits equals the value for Grams-per-minute (0x009B)
func (a AnalogInputEngineeringUnits) IsGramsPerMinute() bool { return a == 0x009B }

// SetGramsPerMinute sets AnalogInputEngineeringUnits to Grams-per-minute (0x009B)
func (a *AnalogInputEngineeringUnits) SetGramsPerMinute() { *a = 0x009B }

// IsTonsPerHour checks if AnalogInputEngineeringUnits equals the value for Tons-per-hour (0x009C)
func (a AnalogInputEngineeringUnits) IsTonsPerHour() bool { return a == 0x009C }

// SetTonsPerHour sets AnalogInputEngineeringUnits to Tons-per-hour (0x009C)
func (a *AnalogInputEngineeringUnits) SetTonsPerHour() { *a = 0x009C }

// IsKiloBtusPerHour checks if AnalogInputEngineeringUnits equals the value for Kilo-BTUs-per-hour (0x009D)
func (a AnalogInputEngineeringUnits) IsKiloBtusPerHour() bool { return a == 0x009D }

// SetKiloBtusPerHour sets AnalogInputEngineeringUnits to Kilo-BTUs-per-hour (0x009D)
func (a *AnalogInputEngineeringUnits) SetKiloBtusPerHour() { *a = 0x009D }

// IsHundredthsSeconds checks if AnalogInputEngineeringUnits equals the value for Hundredths-seconds (0x009E)
func (a AnalogInputEngineeringUnits) IsHundredthsSeconds() bool { return a == 0x009E }

// SetHundredthsSeconds sets AnalogInputEngineeringUnits to Hundredths-seconds (0x009E)
func (a *AnalogInputEngineeringUnits) SetHundredthsSeconds() { *a = 0x009E }

// IsMilliseconds checks if AnalogInputEngineeringUnits equals the value for Milliseconds (0x009F)
func (a AnalogInputEngineeringUnits) IsMilliseconds() bool { return a == 0x009F }

// SetMilliseconds sets AnalogInputEngineeringUnits to Milliseconds (0x009F)
func (a *AnalogInputEngineeringUnits) SetMilliseconds() { *a = 0x009F }

// IsNewtonMeters checks if AnalogInputEngineeringUnits equals the value for Newton-meters (0x00A0)
func (a AnalogInputEngineeringUnits) IsNewtonMeters() bool { return a == 0x00A0 }

// SetNewtonMeters sets AnalogInputEngineeringUnits to Newton-meters (0x00A0)
func (a *AnalogInputEngineeringUnits) SetNewtonMeters() { *a = 0x00A0 }

// IsMillimetersPerSecond checks if AnalogInputEngineeringUnits equals the value for Millimeters-per-second (0x00A1)
func (a AnalogInputEngineeringUnits) IsMillimetersPerSecond() bool { return a == 0x00A1 }

// SetMillimetersPerSecond sets AnalogInputEngineeringUnits to Millimeters-per-second (0x00A1)
func (a *AnalogInputEngineeringUnits) SetMillimetersPerSecond() { *a = 0x00A1 }

// IsMillimetersPerMinute checks if AnalogInputEngineeringUnits equals the value for Millimeters-per-minute (0x00A2)
func (a AnalogInputEngineeringUnits) IsMillimetersPerMinute() bool { return a == 0x00A2 }

// SetMillimetersPerMinute sets AnalogInputEngineeringUnits to Millimeters-per-minute (0x00A2)
func (a *AnalogInputEngineeringUnits) SetMillimetersPerMinute() { *a = 0x00A2 }

// IsMetersPerMinute checks if AnalogInputEngineeringUnits equals the value for Meters-per-minute (0x00A3)
func (a AnalogInputEngineeringUnits) IsMetersPerMinute() bool { return a == 0x00A3 }

// SetMetersPerMinute sets AnalogInputEngineeringUnits to Meters-per-minute (0x00A3)
func (a *AnalogInputEngineeringUnits) SetMetersPerMinute() { *a = 0x00A3 }

// IsMetersPerHour checks if AnalogInputEngineeringUnits equals the value for Meters-per-hour (0x00A4)
func (a AnalogInputEngineeringUnits) IsMetersPerHour() bool { return a == 0x00A4 }

// SetMetersPerHour sets AnalogInputEngineeringUnits to Meters-per-hour (0x00A4)
func (a *AnalogInputEngineeringUnits) SetMetersPerHour() { *a = 0x00A4 }

// IsCubicMetersPerMinute checks if AnalogInputEngineeringUnits equals the value for Cubic-meters-per-minute (0x00A5)
func (a AnalogInputEngineeringUnits) IsCubicMetersPerMinute() bool { return a == 0x00A5 }

// SetCubicMetersPerMinute sets AnalogInputEngineeringUnits to Cubic-meters-per-minute (0x00A5)
func (a *AnalogInputEngineeringUnits) SetCubicMetersPerMinute() { *a = 0x00A5 }

// IsMetersPerSecondPerSecond checks if AnalogInputEngineeringUnits equals the value for Meters-per-second-per-second (0x00A6)
func (a AnalogInputEngineeringUnits) IsMetersPerSecondPerSecond() bool { return a == 0x00A6 }

// SetMetersPerSecondPerSecond sets AnalogInputEngineeringUnits to Meters-per-second-per-second (0x00A6)
func (a *AnalogInputEngineeringUnits) SetMetersPerSecondPerSecond() { *a = 0x00A6 }

// IsAmperesPerMeter checks if AnalogInputEngineeringUnits equals the value for Amperes-per-meter (0x00A7)
func (a AnalogInputEngineeringUnits) IsAmperesPerMeter() bool { return a == 0x00A7 }

// SetAmperesPerMeter sets AnalogInputEngineeringUnits to Amperes-per-meter (0x00A7)
func (a *AnalogInputEngineeringUnits) SetAmperesPerMeter() { *a = 0x00A7 }

// IsAmperesPerSquareMeter checks if AnalogInputEngineeringUnits equals the value for Amperes-per-square-meter (0x00A8)
func (a AnalogInputEngineeringUnits) IsAmperesPerSquareMeter() bool { return a == 0x00A8 }

// SetAmperesPerSquareMeter sets AnalogInputEngineeringUnits to Amperes-per-square-meter (0x00A8)
func (a *AnalogInputEngineeringUnits) SetAmperesPerSquareMeter() { *a = 0x00A8 }

// IsAmpereSquareMeters checks if AnalogInputEngineeringUnits equals the value for Ampere-square-meters (0x00A9)
func (a AnalogInputEngineeringUnits) IsAmpereSquareMeters() bool { return a == 0x00A9 }

// SetAmpereSquareMeters sets AnalogInputEngineeringUnits to Ampere-square-meters (0x00A9)
func (a *AnalogInputEngineeringUnits) SetAmpereSquareMeters() { *a = 0x00A9 }

// IsFarads checks if AnalogInputEngineeringUnits equals the value for Farads (0x00AA)
func (a AnalogInputEngineeringUnits) IsFarads() bool { return a == 0x00AA }

// SetFarads sets AnalogInputEngineeringUnits to Farads (0x00AA)
func (a *AnalogInputEngineeringUnits) SetFarads() { *a = 0x00AA }

// IsHenrys checks if AnalogInputEngineeringUnits equals the value for Henrys (0x00AB)
func (a AnalogInputEngineeringUnits) IsHenrys() bool { return a == 0x00AB }

// SetHenrys sets AnalogInputEngineeringUnits to Henrys (0x00AB)
func (a *AnalogInputEngineeringUnits) SetHenrys() { *a = 0x00AB }

// IsOhmMeters checks if AnalogInputEngineeringUnits equals the value for Ohm-meters (0x00AC)
func (a AnalogInputEngineeringUnits) IsOhmMeters() bool { return a == 0x00AC }

// SetOhmMeters sets AnalogInputEngineeringUnits to Ohm-meters (0x00AC)
func (a *AnalogInputEngineeringUnits) SetOhmMeters() { *a = 0x00AC }

// IsSiemens checks if AnalogInputEngineeringUnits equals the value for Siemens (0x00AD)
func (a AnalogInputEngineeringUnits) IsSiemens() bool { return a == 0x00AD }

// SetSiemens sets AnalogInputEngineeringUnits to Siemens (0x00AD)
func (a *AnalogInputEngineeringUnits) SetSiemens() { *a = 0x00AD }

// IsSiemensPerMeter checks if AnalogInputEngineeringUnits equals the value for Siemens-per-meter (0x00AE)
func (a AnalogInputEngineeringUnits) IsSiemensPerMeter() bool { return a == 0x00AE }

// SetSiemensPerMeter sets AnalogInputEngineeringUnits to Siemens-per-meter (0x00AE)
func (a *AnalogInputEngineeringUnits) SetSiemensPerMeter() { *a = 0x00AE }

// IsTeslas checks if AnalogInputEngineeringUnits equals the value for Teslas (0x00AF)
func (a AnalogInputEngineeringUnits) IsTeslas() bool { return a == 0x00AF }

// SetTeslas sets AnalogInputEngineeringUnits to Teslas (0x00AF)
func (a *AnalogInputEngineeringUnits) SetTeslas() { *a = 0x00AF }

// IsVoltsPerDegreeKelvin checks if AnalogInputEngineeringUnits equals the value for Volts-per-degree-Kelvin (0x00B0)
func (a AnalogInputEngineeringUnits) IsVoltsPerDegreeKelvin() bool { return a == 0x00B0 }

// SetVoltsPerDegreeKelvin sets AnalogInputEngineeringUnits to Volts-per-degree-Kelvin (0x00B0)
func (a *AnalogInputEngineeringUnits) SetVoltsPerDegreeKelvin() { *a = 0x00B0 }

// IsVoltsPerMeter checks if AnalogInputEngineeringUnits equals the value for Volts-per-meter (0x00B1)
func (a AnalogInputEngineeringUnits) IsVoltsPerMeter() bool { return a == 0x00B1 }

// SetVoltsPerMeter sets AnalogInputEngineeringUnits to Volts-per-meter (0x00B1)
func (a *AnalogInputEngineeringUnits) SetVoltsPerMeter() { *a = 0x00B1 }

// IsWebers checks if AnalogInputEngineeringUnits equals the value for Webers (0x00B2)
func (a AnalogInputEngineeringUnits) IsWebers() bool { return a == 0x00B2 }

// SetWebers sets AnalogInputEngineeringUnits to Webers (0x00B2)
func (a *AnalogInputEngineeringUnits) SetWebers() { *a = 0x00B2 }

// IsCandelas checks if AnalogInputEngineeringUnits equals the value for Candelas (0x00B3)
func (a AnalogInputEngineeringUnits) IsCandelas() bool { return a == 0x00B3 }

// SetCandelas sets AnalogInputEngineeringUnits to Candelas (0x00B3)
func (a *AnalogInputEngineeringUnits) SetCandelas() { *a = 0x00B3 }

// IsCandelasPerSquareMeter checks if AnalogInputEngineeringUnits equals the value for Candelas-per-square-meter (0x00B4)
func (a AnalogInputEngineeringUnits) IsCandelasPerSquareMeter() bool { return a == 0x00B4 }

// SetCandelasPerSquareMeter sets AnalogInputEngineeringUnits to Candelas-per-square-meter (0x00B4)
func (a *AnalogInputEngineeringUnits) SetCandelasPerSquareMeter() { *a = 0x00B4 }

// IsKelvinsPerHour checks if AnalogInputEngineeringUnits equals the value for Kelvins-per-hour (0x00B5)
func (a AnalogInputEngineeringUnits) IsKelvinsPerHour() bool { return a == 0x00B5 }

// SetKelvinsPerHour sets AnalogInputEngineeringUnits to Kelvins-per-hour (0x00B5)
func (a *AnalogInputEngineeringUnits) SetKelvinsPerHour() { *a = 0x00B5 }

// IsKelvinsPerMinute checks if AnalogInputEngineeringUnits equals the value for Kelvins-per-minute (0x00B6)
func (a AnalogInputEngineeringUnits) IsKelvinsPerMinute() bool { return a == 0x00B6 }

// SetKelvinsPerMinute sets AnalogInputEngineeringUnits to Kelvins-per-minute (0x00B6)
func (a *AnalogInputEngineeringUnits) SetKelvinsPerMinute() { *a = 0x00B6 }

// IsJouleSeconds checks if AnalogInputEngineeringUnits equals the value for Joule-seconds (0x00B7)
func (a AnalogInputEngineeringUnits) IsJouleSeconds() bool { return a == 0x00B7 }

// SetJouleSeconds sets AnalogInputEngineeringUnits to Joule-seconds (0x00B7)
func (a *AnalogInputEngineeringUnits) SetJouleSeconds() { *a = 0x00B7 }

// IsSquareMetersPerNewton checks if AnalogInputEngineeringUnits equals the value for Square-meters-per-Newton (0x00B8)
func (a AnalogInputEngineeringUnits) IsSquareMetersPerNewton() bool { return a == 0x00B8 }

// SetSquareMetersPerNewton sets AnalogInputEngineeringUnits to Square-meters-per-Newton (0x00B8)
func (a *AnalogInputEngineeringUnits) SetSquareMetersPerNewton() { *a = 0x00B8 }

// IsKilogramPerCubicMeter checks if AnalogInputEngineeringUnits equals the value for Kilogram-per-cubic-meter (0x00B9)
func (a AnalogInputEngineeringUnits) IsKilogramPerCubicMeter() bool { return a == 0x00B9 }

// SetKilogramPerCubicMeter sets AnalogInputEngineeringUnits to Kilogram-per-cubic-meter (0x00B9)
func (a *AnalogInputEngineeringUnits) SetKilogramPerCubicMeter() { *a = 0x00B9 }

// IsNewtonSeconds checks if AnalogInputEngineeringUnits equals the value for Newton-seconds (0x00BA)
func (a AnalogInputEngineeringUnits) IsNewtonSeconds() bool { return a == 0x00BA }

// SetNewtonSeconds sets AnalogInputEngineeringUnits to Newton-seconds (0x00BA)
func (a *AnalogInputEngineeringUnits) SetNewtonSeconds() { *a = 0x00BA }

// IsNewtonsPerMeter checks if AnalogInputEngineeringUnits equals the value for Newtons-per-meter (0x00BB)
func (a AnalogInputEngineeringUnits) IsNewtonsPerMeter() bool { return a == 0x00BB }

// SetNewtonsPerMeter sets AnalogInputEngineeringUnits to Newtons-per-meter (0x00BB)
func (a *AnalogInputEngineeringUnits) SetNewtonsPerMeter() { *a = 0x00BB }

// IsWattsPerMeterPerDegreeKelvin checks if AnalogInputEngineeringUnits equals the value for Watts-per-meter-per-degree-Kelvin (0x00BC)
func (a AnalogInputEngineeringUnits) IsWattsPerMeterPerDegreeKelvin() bool { return a == 0x00BC }

// SetWattsPerMeterPerDegreeKelvin sets AnalogInputEngineeringUnits to Watts-per-meter-per-degree-Kelvin (0x00BC)
func (a *AnalogInputEngineeringUnits) SetWattsPerMeterPerDegreeKelvin() { *a = 0x00BC }

// IsOther checks if AnalogInputEngineeringUnits equals the value for Other (0x00FF)
func (a AnalogInputEngineeringUnits) IsOther() bool { return a == 0x00FF }

// SetOther sets AnalogInputEngineeringUnits to Other (0x00FF)
func (a *AnalogInputEngineeringUnits) SetOther() { *a = 0x00FF }

type AnalogInputApplicationType zcl.Zu32

func (a AnalogInputApplicationType) ID() zcl.AttrID                      { return 256 }
func (a AnalogInputApplicationType) Cluster() zcl.ClusterID              { return AnalogInputBasicCluster }
func (a *AnalogInputApplicationType) Value() *AnalogInputApplicationType { return a }
func (a AnalogInputApplicationType) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *AnalogInputApplicationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputApplicationType(*nt)
	return br, err
}

func (a AnalogInputApplicationType) Readable() bool   { return true }
func (a AnalogInputApplicationType) Writable() bool   { return false }
func (a AnalogInputApplicationType) Reportable() bool { return false }
func (a AnalogInputApplicationType) SceneIndex() int  { return -1 }

func (a AnalogInputApplicationType) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}
