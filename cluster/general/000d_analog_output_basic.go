// The Analog Output (Basic) cluster provides an interface for setting the value of an analog output (typically to the environment) and accessing various characteristics of that value.
package general

import (
	"neotor.se/zcl"
)

// AnalogOutputBasic
// The Analog Output (Basic) cluster provides an interface for setting the value of an analog output (typically to the environment) and accessing various characteristics of that value.

func NewAnalogOutputBasicServer(profile zcl.ProfileID) *AnalogOutputBasicServer {
	return &AnalogOutputBasicServer{p: profile}
}
func NewAnalogOutputBasicClient(profile zcl.ProfileID) *AnalogOutputBasicClient {
	return &AnalogOutputBasicClient{p: profile}
}

const AnalogOutputBasicCluster zcl.ClusterID = 13

type AnalogOutputBasicServer struct {
	p zcl.ProfileID

	AnalogOutputDescription       *AnalogOutputDescription
	AnalogOutputMaxPresentValue   *AnalogOutputMaxPresentValue
	AnalogOutputMinPresentValue   *AnalogOutputMinPresentValue
	AnalogOutputOutOfService      *AnalogOutputOutOfService
	AnalogOutputPresentValue      *AnalogOutputPresentValue
	AnalogOutputPriorityArray     *AnalogOutputPriorityArray
	AnalogOutputReliability       *AnalogOutputReliability
	AnalogOutputRelinquishDefault *AnalogOutputRelinquishDefault
	AnalogOutputResolution        *AnalogOutputResolution
	AnalogOutputStatusFlags       *AnalogOutputStatusFlags
	AnalogOutputEngineeringUnits  *AnalogOutputEngineeringUnits
	AnalogOutputApplicationType   *AnalogOutputApplicationType
	AnalogOutputXiaomi0Xf000      *AnalogOutputXiaomi0Xf000
}

type AnalogOutputBasicClient struct {
	p zcl.ProfileID
}

/*
var AnalogOutputBasicServer = map[zcl.CommandID]func() zcl.Command{
}

var AnalogOutputBasicClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type AnalogOutputDescription zcl.Zcstring

func (a AnalogOutputDescription) ID() zcl.AttrID                   { return 28 }
func (a AnalogOutputDescription) Cluster() zcl.ClusterID           { return AnalogOutputBasicCluster }
func (a *AnalogOutputDescription) Value() *AnalogOutputDescription { return a }
func (a AnalogOutputDescription) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *AnalogOutputDescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputDescription(*nt)
	return br, err
}

func (a AnalogOutputDescription) Readable() bool   { return true }
func (a AnalogOutputDescription) Writable() bool   { return true }
func (a AnalogOutputDescription) Reportable() bool { return false }
func (a AnalogOutputDescription) SceneIndex() int  { return -1 }

func (a AnalogOutputDescription) String() string {

	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

type AnalogOutputMaxPresentValue zcl.Zfloat

func (a AnalogOutputMaxPresentValue) ID() zcl.AttrID                       { return 65 }
func (a AnalogOutputMaxPresentValue) Cluster() zcl.ClusterID               { return AnalogOutputBasicCluster }
func (a *AnalogOutputMaxPresentValue) Value() *AnalogOutputMaxPresentValue { return a }
func (a AnalogOutputMaxPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogOutputMaxPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputMaxPresentValue(*nt)
	return br, err
}

func (a AnalogOutputMaxPresentValue) Readable() bool   { return true }
func (a AnalogOutputMaxPresentValue) Writable() bool   { return true }
func (a AnalogOutputMaxPresentValue) Reportable() bool { return false }
func (a AnalogOutputMaxPresentValue) SceneIndex() int  { return -1 }

func (a AnalogOutputMaxPresentValue) String() string {

	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

type AnalogOutputMinPresentValue zcl.Zfloat

func (a AnalogOutputMinPresentValue) ID() zcl.AttrID                       { return 69 }
func (a AnalogOutputMinPresentValue) Cluster() zcl.ClusterID               { return AnalogOutputBasicCluster }
func (a *AnalogOutputMinPresentValue) Value() *AnalogOutputMinPresentValue { return a }
func (a AnalogOutputMinPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogOutputMinPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputMinPresentValue(*nt)
	return br, err
}

func (a AnalogOutputMinPresentValue) Readable() bool   { return true }
func (a AnalogOutputMinPresentValue) Writable() bool   { return true }
func (a AnalogOutputMinPresentValue) Reportable() bool { return false }
func (a AnalogOutputMinPresentValue) SceneIndex() int  { return -1 }

func (a AnalogOutputMinPresentValue) String() string {

	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

type AnalogOutputOutOfService zcl.Zbool

func (a AnalogOutputOutOfService) ID() zcl.AttrID                    { return 81 }
func (a AnalogOutputOutOfService) Cluster() zcl.ClusterID            { return AnalogOutputBasicCluster }
func (a *AnalogOutputOutOfService) Value() *AnalogOutputOutOfService { return a }
func (a AnalogOutputOutOfService) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *AnalogOutputOutOfService) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputOutOfService(*nt)
	return br, err
}

func (a AnalogOutputOutOfService) Readable() bool   { return true }
func (a AnalogOutputOutOfService) Writable() bool   { return true }
func (a AnalogOutputOutOfService) Reportable() bool { return false }
func (a AnalogOutputOutOfService) SceneIndex() int  { return -1 }

func (a AnalogOutputOutOfService) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

type AnalogOutputPresentValue zcl.Zfloat

func (a AnalogOutputPresentValue) ID() zcl.AttrID                    { return 85 }
func (a AnalogOutputPresentValue) Cluster() zcl.ClusterID            { return AnalogOutputBasicCluster }
func (a *AnalogOutputPresentValue) Value() *AnalogOutputPresentValue { return a }
func (a AnalogOutputPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogOutputPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputPresentValue(*nt)
	return br, err
}

func (a AnalogOutputPresentValue) Readable() bool   { return true }
func (a AnalogOutputPresentValue) Writable() bool   { return true }
func (a AnalogOutputPresentValue) Reportable() bool { return false }
func (a AnalogOutputPresentValue) SceneIndex() int  { return -1 }

func (a AnalogOutputPresentValue) String() string {

	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

type AnalogOutputPriorityArray zcl.Zarray

func (a AnalogOutputPriorityArray) ID() zcl.AttrID                     { return 87 }
func (a AnalogOutputPriorityArray) Cluster() zcl.ClusterID             { return AnalogOutputBasicCluster }
func (a *AnalogOutputPriorityArray) Value() *AnalogOutputPriorityArray { return a }
func (a AnalogOutputPriorityArray) MarshalZcl() ([]byte, error) {
	return zcl.Zarray(a).MarshalZcl()
}
func (a *AnalogOutputPriorityArray) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zarray)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputPriorityArray(*nt)
	return br, err
}

func (a AnalogOutputPriorityArray) Readable() bool   { return true }
func (a AnalogOutputPriorityArray) Writable() bool   { return true }
func (a AnalogOutputPriorityArray) Reportable() bool { return false }
func (a AnalogOutputPriorityArray) SceneIndex() int  { return -1 }

func (a AnalogOutputPriorityArray) String() string {

	return zcl.Sprintf("%s", zcl.Zarray(a))
}

type AnalogOutputReliability zcl.Zenum8

func (a AnalogOutputReliability) ID() zcl.AttrID                   { return 103 }
func (a AnalogOutputReliability) Cluster() zcl.ClusterID           { return AnalogOutputBasicCluster }
func (a *AnalogOutputReliability) Value() *AnalogOutputReliability { return a }
func (a AnalogOutputReliability) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *AnalogOutputReliability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputReliability(*nt)
	return br, err
}

func (a AnalogOutputReliability) Readable() bool   { return true }
func (a AnalogOutputReliability) Writable() bool   { return true }
func (a AnalogOutputReliability) Reportable() bool { return false }
func (a AnalogOutputReliability) SceneIndex() int  { return -1 }

func (a AnalogOutputReliability) String() string {

	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

type AnalogOutputRelinquishDefault zcl.Zfloat

func (a AnalogOutputRelinquishDefault) ID() zcl.AttrID                         { return 103 }
func (a AnalogOutputRelinquishDefault) Cluster() zcl.ClusterID                 { return AnalogOutputBasicCluster }
func (a *AnalogOutputRelinquishDefault) Value() *AnalogOutputRelinquishDefault { return a }
func (a AnalogOutputRelinquishDefault) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogOutputRelinquishDefault) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputRelinquishDefault(*nt)
	return br, err
}

func (a AnalogOutputRelinquishDefault) Readable() bool   { return true }
func (a AnalogOutputRelinquishDefault) Writable() bool   { return true }
func (a AnalogOutputRelinquishDefault) Reportable() bool { return false }
func (a AnalogOutputRelinquishDefault) SceneIndex() int  { return -1 }

func (a AnalogOutputRelinquishDefault) String() string {

	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

type AnalogOutputResolution zcl.Zfloat

func (a AnalogOutputResolution) ID() zcl.AttrID                  { return 106 }
func (a AnalogOutputResolution) Cluster() zcl.ClusterID          { return AnalogOutputBasicCluster }
func (a *AnalogOutputResolution) Value() *AnalogOutputResolution { return a }
func (a AnalogOutputResolution) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogOutputResolution) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputResolution(*nt)
	return br, err
}

func (a AnalogOutputResolution) Readable() bool   { return true }
func (a AnalogOutputResolution) Writable() bool   { return true }
func (a AnalogOutputResolution) Reportable() bool { return false }
func (a AnalogOutputResolution) SceneIndex() int  { return -1 }

func (a AnalogOutputResolution) String() string {

	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

type AnalogOutputStatusFlags zcl.Zbmp8

func (a AnalogOutputStatusFlags) ID() zcl.AttrID                   { return 111 }
func (a AnalogOutputStatusFlags) Cluster() zcl.ClusterID           { return AnalogOutputBasicCluster }
func (a *AnalogOutputStatusFlags) Value() *AnalogOutputStatusFlags { return a }
func (a AnalogOutputStatusFlags) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *AnalogOutputStatusFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputStatusFlags(*nt)
	return br, err
}

func (a AnalogOutputStatusFlags) Readable() bool   { return true }
func (a AnalogOutputStatusFlags) Writable() bool   { return false }
func (a AnalogOutputStatusFlags) Reportable() bool { return true }
func (a AnalogOutputStatusFlags) SceneIndex() int  { return -1 }

func (a AnalogOutputStatusFlags) String() string {

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

func (a AnalogOutputStatusFlags) IsInAlarm() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *AnalogOutputStatusFlags) SetInAlarm(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a AnalogOutputStatusFlags) IsFault() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *AnalogOutputStatusFlags) SetFault(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a AnalogOutputStatusFlags) IsOveridden() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *AnalogOutputStatusFlags) SetOveridden(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a AnalogOutputStatusFlags) IsOutOfService() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *AnalogOutputStatusFlags) SetOutOfService(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 8, b))
}

type AnalogOutputEngineeringUnits zcl.Zenum16

func (a AnalogOutputEngineeringUnits) ID() zcl.AttrID                        { return 117 }
func (a AnalogOutputEngineeringUnits) Cluster() zcl.ClusterID                { return AnalogOutputBasicCluster }
func (a *AnalogOutputEngineeringUnits) Value() *AnalogOutputEngineeringUnits { return a }
func (a AnalogOutputEngineeringUnits) MarshalZcl() ([]byte, error) {
	return zcl.Zenum16(a).MarshalZcl()
}
func (a *AnalogOutputEngineeringUnits) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum16)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputEngineeringUnits(*nt)
	return br, err
}

func (a AnalogOutputEngineeringUnits) Readable() bool   { return true }
func (a AnalogOutputEngineeringUnits) Writable() bool   { return true }
func (a AnalogOutputEngineeringUnits) Reportable() bool { return false }
func (a AnalogOutputEngineeringUnits) SceneIndex() int  { return -1 }

func (a AnalogOutputEngineeringUnits) String() string {
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

// IsSquareMeters checks if AnalogOutputEngineeringUnits equals the value for Square-meters (0x0000)
func (a AnalogOutputEngineeringUnits) IsSquareMeters() bool { return a == 0x0000 }

// SetSquareMeters sets AnalogOutputEngineeringUnits to Square-meters (0x0000)
func (a *AnalogOutputEngineeringUnits) SetSquareMeters() { *a = 0x0000 }

// IsSquareFeet checks if AnalogOutputEngineeringUnits equals the value for Square-feet (0x0001)
func (a AnalogOutputEngineeringUnits) IsSquareFeet() bool { return a == 0x0001 }

// SetSquareFeet sets AnalogOutputEngineeringUnits to Square-feet (0x0001)
func (a *AnalogOutputEngineeringUnits) SetSquareFeet() { *a = 0x0001 }

// IsMilliamperes checks if AnalogOutputEngineeringUnits equals the value for Milliamperes (0x0002)
func (a AnalogOutputEngineeringUnits) IsMilliamperes() bool { return a == 0x0002 }

// SetMilliamperes sets AnalogOutputEngineeringUnits to Milliamperes (0x0002)
func (a *AnalogOutputEngineeringUnits) SetMilliamperes() { *a = 0x0002 }

// IsAmperes checks if AnalogOutputEngineeringUnits equals the value for Amperes (0x0003)
func (a AnalogOutputEngineeringUnits) IsAmperes() bool { return a == 0x0003 }

// SetAmperes sets AnalogOutputEngineeringUnits to Amperes (0x0003)
func (a *AnalogOutputEngineeringUnits) SetAmperes() { *a = 0x0003 }

// IsOhms checks if AnalogOutputEngineeringUnits equals the value for Ohms (0x0004)
func (a AnalogOutputEngineeringUnits) IsOhms() bool { return a == 0x0004 }

// SetOhms sets AnalogOutputEngineeringUnits to Ohms (0x0004)
func (a *AnalogOutputEngineeringUnits) SetOhms() { *a = 0x0004 }

// IsVolts checks if AnalogOutputEngineeringUnits equals the value for Volts (0x0005)
func (a AnalogOutputEngineeringUnits) IsVolts() bool { return a == 0x0005 }

// SetVolts sets AnalogOutputEngineeringUnits to Volts (0x0005)
func (a *AnalogOutputEngineeringUnits) SetVolts() { *a = 0x0005 }

// IsKiloVolts checks if AnalogOutputEngineeringUnits equals the value for Kilo-volts (0x0006)
func (a AnalogOutputEngineeringUnits) IsKiloVolts() bool { return a == 0x0006 }

// SetKiloVolts sets AnalogOutputEngineeringUnits to Kilo-volts (0x0006)
func (a *AnalogOutputEngineeringUnits) SetKiloVolts() { *a = 0x0006 }

// IsMegaVolts checks if AnalogOutputEngineeringUnits equals the value for Mega-volts (0x0007)
func (a AnalogOutputEngineeringUnits) IsMegaVolts() bool { return a == 0x0007 }

// SetMegaVolts sets AnalogOutputEngineeringUnits to Mega-volts (0x0007)
func (a *AnalogOutputEngineeringUnits) SetMegaVolts() { *a = 0x0007 }

// IsVoltAmperes checks if AnalogOutputEngineeringUnits equals the value for Volt-amperes (0x0008)
func (a AnalogOutputEngineeringUnits) IsVoltAmperes() bool { return a == 0x0008 }

// SetVoltAmperes sets AnalogOutputEngineeringUnits to Volt-amperes (0x0008)
func (a *AnalogOutputEngineeringUnits) SetVoltAmperes() { *a = 0x0008 }

// IsKiloVoltAmperes checks if AnalogOutputEngineeringUnits equals the value for Kilo-volt-amperes (0x0009)
func (a AnalogOutputEngineeringUnits) IsKiloVoltAmperes() bool { return a == 0x0009 }

// SetKiloVoltAmperes sets AnalogOutputEngineeringUnits to Kilo-volt-amperes (0x0009)
func (a *AnalogOutputEngineeringUnits) SetKiloVoltAmperes() { *a = 0x0009 }

// IsMegaVoltAmperes checks if AnalogOutputEngineeringUnits equals the value for Mega-volt-amperes (0x000A)
func (a AnalogOutputEngineeringUnits) IsMegaVoltAmperes() bool { return a == 0x000A }

// SetMegaVoltAmperes sets AnalogOutputEngineeringUnits to Mega-volt-amperes (0x000A)
func (a *AnalogOutputEngineeringUnits) SetMegaVoltAmperes() { *a = 0x000A }

// IsVoltAmperesReactive checks if AnalogOutputEngineeringUnits equals the value for Volt-amperes-reactive (0x000B)
func (a AnalogOutputEngineeringUnits) IsVoltAmperesReactive() bool { return a == 0x000B }

// SetVoltAmperesReactive sets AnalogOutputEngineeringUnits to Volt-amperes-reactive (0x000B)
func (a *AnalogOutputEngineeringUnits) SetVoltAmperesReactive() { *a = 0x000B }

// IsKiloVoltAmperesReactive checks if AnalogOutputEngineeringUnits equals the value for Kilo-volt-amperes-reactive (0x000C)
func (a AnalogOutputEngineeringUnits) IsKiloVoltAmperesReactive() bool { return a == 0x000C }

// SetKiloVoltAmperesReactive sets AnalogOutputEngineeringUnits to Kilo-volt-amperes-reactive (0x000C)
func (a *AnalogOutputEngineeringUnits) SetKiloVoltAmperesReactive() { *a = 0x000C }

// IsMegaVoltAmperesReactive checks if AnalogOutputEngineeringUnits equals the value for Mega-volt-amperes-reactive (0x000D)
func (a AnalogOutputEngineeringUnits) IsMegaVoltAmperesReactive() bool { return a == 0x000D }

// SetMegaVoltAmperesReactive sets AnalogOutputEngineeringUnits to Mega-volt-amperes-reactive (0x000D)
func (a *AnalogOutputEngineeringUnits) SetMegaVoltAmperesReactive() { *a = 0x000D }

// IsDegreesPhase checks if AnalogOutputEngineeringUnits equals the value for Degrees-phase (0x000E)
func (a AnalogOutputEngineeringUnits) IsDegreesPhase() bool { return a == 0x000E }

// SetDegreesPhase sets AnalogOutputEngineeringUnits to Degrees-phase (0x000E)
func (a *AnalogOutputEngineeringUnits) SetDegreesPhase() { *a = 0x000E }

// IsPowerFactor checks if AnalogOutputEngineeringUnits equals the value for Power-factor (0x000F)
func (a AnalogOutputEngineeringUnits) IsPowerFactor() bool { return a == 0x000F }

// SetPowerFactor sets AnalogOutputEngineeringUnits to Power-factor (0x000F)
func (a *AnalogOutputEngineeringUnits) SetPowerFactor() { *a = 0x000F }

// IsJoules checks if AnalogOutputEngineeringUnits equals the value for Joules (0x0010)
func (a AnalogOutputEngineeringUnits) IsJoules() bool { return a == 0x0010 }

// SetJoules sets AnalogOutputEngineeringUnits to Joules (0x0010)
func (a *AnalogOutputEngineeringUnits) SetJoules() { *a = 0x0010 }

// IsKilojoules checks if AnalogOutputEngineeringUnits equals the value for Kilojoules (0x0011)
func (a AnalogOutputEngineeringUnits) IsKilojoules() bool { return a == 0x0011 }

// SetKilojoules sets AnalogOutputEngineeringUnits to Kilojoules (0x0011)
func (a *AnalogOutputEngineeringUnits) SetKilojoules() { *a = 0x0011 }

// IsWattHours checks if AnalogOutputEngineeringUnits equals the value for Watt-hours (0x0012)
func (a AnalogOutputEngineeringUnits) IsWattHours() bool { return a == 0x0012 }

// SetWattHours sets AnalogOutputEngineeringUnits to Watt-hours (0x0012)
func (a *AnalogOutputEngineeringUnits) SetWattHours() { *a = 0x0012 }

// IsKilowattHours checks if AnalogOutputEngineeringUnits equals the value for Kilowatt-hours (0x0013)
func (a AnalogOutputEngineeringUnits) IsKilowattHours() bool { return a == 0x0013 }

// SetKilowattHours sets AnalogOutputEngineeringUnits to Kilowatt-hours (0x0013)
func (a *AnalogOutputEngineeringUnits) SetKilowattHours() { *a = 0x0013 }

// IsBtus checks if AnalogOutputEngineeringUnits equals the value for BTUs (0x0014)
func (a AnalogOutputEngineeringUnits) IsBtus() bool { return a == 0x0014 }

// SetBtus sets AnalogOutputEngineeringUnits to BTUs (0x0014)
func (a *AnalogOutputEngineeringUnits) SetBtus() { *a = 0x0014 }

// IsTherms checks if AnalogOutputEngineeringUnits equals the value for Therms (0x0015)
func (a AnalogOutputEngineeringUnits) IsTherms() bool { return a == 0x0015 }

// SetTherms sets AnalogOutputEngineeringUnits to Therms (0x0015)
func (a *AnalogOutputEngineeringUnits) SetTherms() { *a = 0x0015 }

// IsTonHours checks if AnalogOutputEngineeringUnits equals the value for Ton-hours (0x0016)
func (a AnalogOutputEngineeringUnits) IsTonHours() bool { return a == 0x0016 }

// SetTonHours sets AnalogOutputEngineeringUnits to Ton-hours (0x0016)
func (a *AnalogOutputEngineeringUnits) SetTonHours() { *a = 0x0016 }

// IsJoulesPerKilogramDryAir checks if AnalogOutputEngineeringUnits equals the value for Joules-per-kilogram-dry-air (0x0017)
func (a AnalogOutputEngineeringUnits) IsJoulesPerKilogramDryAir() bool { return a == 0x0017 }

// SetJoulesPerKilogramDryAir sets AnalogOutputEngineeringUnits to Joules-per-kilogram-dry-air (0x0017)
func (a *AnalogOutputEngineeringUnits) SetJoulesPerKilogramDryAir() { *a = 0x0017 }

// IsBtusPerPoundDryAir checks if AnalogOutputEngineeringUnits equals the value for BTUs-per-pound-dry-air (0x0018)
func (a AnalogOutputEngineeringUnits) IsBtusPerPoundDryAir() bool { return a == 0x0018 }

// SetBtusPerPoundDryAir sets AnalogOutputEngineeringUnits to BTUs-per-pound-dry-air (0x0018)
func (a *AnalogOutputEngineeringUnits) SetBtusPerPoundDryAir() { *a = 0x0018 }

// IsCyclesPerHour checks if AnalogOutputEngineeringUnits equals the value for Cycles-per-hour (0x0019)
func (a AnalogOutputEngineeringUnits) IsCyclesPerHour() bool { return a == 0x0019 }

// SetCyclesPerHour sets AnalogOutputEngineeringUnits to Cycles-per-hour (0x0019)
func (a *AnalogOutputEngineeringUnits) SetCyclesPerHour() { *a = 0x0019 }

// IsCyclesPerMinute checks if AnalogOutputEngineeringUnits equals the value for Cycles-per-minute (0x001A)
func (a AnalogOutputEngineeringUnits) IsCyclesPerMinute() bool { return a == 0x001A }

// SetCyclesPerMinute sets AnalogOutputEngineeringUnits to Cycles-per-minute (0x001A)
func (a *AnalogOutputEngineeringUnits) SetCyclesPerMinute() { *a = 0x001A }

// IsHertz checks if AnalogOutputEngineeringUnits equals the value for Hertz (0x001B)
func (a AnalogOutputEngineeringUnits) IsHertz() bool { return a == 0x001B }

// SetHertz sets AnalogOutputEngineeringUnits to Hertz (0x001B)
func (a *AnalogOutputEngineeringUnits) SetHertz() { *a = 0x001B }

// IsGramsOfWaterPerKilogramDryAir checks if AnalogOutputEngineeringUnits equals the value for Grams-of-water-per-kilogram-dry-air (0x001C)
func (a AnalogOutputEngineeringUnits) IsGramsOfWaterPerKilogramDryAir() bool { return a == 0x001C }

// SetGramsOfWaterPerKilogramDryAir sets AnalogOutputEngineeringUnits to Grams-of-water-per-kilogram-dry-air (0x001C)
func (a *AnalogOutputEngineeringUnits) SetGramsOfWaterPerKilogramDryAir() { *a = 0x001C }

// IsPercentRelativeHumidity checks if AnalogOutputEngineeringUnits equals the value for Percent-relative-humidity (0x001D)
func (a AnalogOutputEngineeringUnits) IsPercentRelativeHumidity() bool { return a == 0x001D }

// SetPercentRelativeHumidity sets AnalogOutputEngineeringUnits to Percent-relative-humidity (0x001D)
func (a *AnalogOutputEngineeringUnits) SetPercentRelativeHumidity() { *a = 0x001D }

// IsMillimeters checks if AnalogOutputEngineeringUnits equals the value for Millimeters (0x001E)
func (a AnalogOutputEngineeringUnits) IsMillimeters() bool { return a == 0x001E }

// SetMillimeters sets AnalogOutputEngineeringUnits to Millimeters (0x001E)
func (a *AnalogOutputEngineeringUnits) SetMillimeters() { *a = 0x001E }

// IsMeters checks if AnalogOutputEngineeringUnits equals the value for Meters (0x001F)
func (a AnalogOutputEngineeringUnits) IsMeters() bool { return a == 0x001F }

// SetMeters sets AnalogOutputEngineeringUnits to Meters (0x001F)
func (a *AnalogOutputEngineeringUnits) SetMeters() { *a = 0x001F }

// IsInches checks if AnalogOutputEngineeringUnits equals the value for Inches (0x0020)
func (a AnalogOutputEngineeringUnits) IsInches() bool { return a == 0x0020 }

// SetInches sets AnalogOutputEngineeringUnits to Inches (0x0020)
func (a *AnalogOutputEngineeringUnits) SetInches() { *a = 0x0020 }

// IsFeet checks if AnalogOutputEngineeringUnits equals the value for Feet (0x0021)
func (a AnalogOutputEngineeringUnits) IsFeet() bool { return a == 0x0021 }

// SetFeet sets AnalogOutputEngineeringUnits to Feet (0x0021)
func (a *AnalogOutputEngineeringUnits) SetFeet() { *a = 0x0021 }

// IsWattsPerSquareFoot checks if AnalogOutputEngineeringUnits equals the value for Watts-per-square-foot (0x0022)
func (a AnalogOutputEngineeringUnits) IsWattsPerSquareFoot() bool { return a == 0x0022 }

// SetWattsPerSquareFoot sets AnalogOutputEngineeringUnits to Watts-per-square-foot (0x0022)
func (a *AnalogOutputEngineeringUnits) SetWattsPerSquareFoot() { *a = 0x0022 }

// IsWattsPerSquareMeter checks if AnalogOutputEngineeringUnits equals the value for Watts-per-square-meter (0x0023)
func (a AnalogOutputEngineeringUnits) IsWattsPerSquareMeter() bool { return a == 0x0023 }

// SetWattsPerSquareMeter sets AnalogOutputEngineeringUnits to Watts-per-square-meter (0x0023)
func (a *AnalogOutputEngineeringUnits) SetWattsPerSquareMeter() { *a = 0x0023 }

// IsLumens checks if AnalogOutputEngineeringUnits equals the value for Lumens (0x0024)
func (a AnalogOutputEngineeringUnits) IsLumens() bool { return a == 0x0024 }

// SetLumens sets AnalogOutputEngineeringUnits to Lumens (0x0024)
func (a *AnalogOutputEngineeringUnits) SetLumens() { *a = 0x0024 }

// IsLuxes checks if AnalogOutputEngineeringUnits equals the value for Luxes (0x0025)
func (a AnalogOutputEngineeringUnits) IsLuxes() bool { return a == 0x0025 }

// SetLuxes sets AnalogOutputEngineeringUnits to Luxes (0x0025)
func (a *AnalogOutputEngineeringUnits) SetLuxes() { *a = 0x0025 }

// IsFootCandles checks if AnalogOutputEngineeringUnits equals the value for Foot-candles (0x0026)
func (a AnalogOutputEngineeringUnits) IsFootCandles() bool { return a == 0x0026 }

// SetFootCandles sets AnalogOutputEngineeringUnits to Foot-candles (0x0026)
func (a *AnalogOutputEngineeringUnits) SetFootCandles() { *a = 0x0026 }

// IsKilograms checks if AnalogOutputEngineeringUnits equals the value for Kilograms (0x0027)
func (a AnalogOutputEngineeringUnits) IsKilograms() bool { return a == 0x0027 }

// SetKilograms sets AnalogOutputEngineeringUnits to Kilograms (0x0027)
func (a *AnalogOutputEngineeringUnits) SetKilograms() { *a = 0x0027 }

// IsPoundsMass checks if AnalogOutputEngineeringUnits equals the value for Pounds-mass (0x0028)
func (a AnalogOutputEngineeringUnits) IsPoundsMass() bool { return a == 0x0028 }

// SetPoundsMass sets AnalogOutputEngineeringUnits to Pounds-mass (0x0028)
func (a *AnalogOutputEngineeringUnits) SetPoundsMass() { *a = 0x0028 }

// IsTons checks if AnalogOutputEngineeringUnits equals the value for Tons (0x0029)
func (a AnalogOutputEngineeringUnits) IsTons() bool { return a == 0x0029 }

// SetTons sets AnalogOutputEngineeringUnits to Tons (0x0029)
func (a *AnalogOutputEngineeringUnits) SetTons() { *a = 0x0029 }

// IsKilogramsPerSecond checks if AnalogOutputEngineeringUnits equals the value for Kilograms-per-second (0x002A)
func (a AnalogOutputEngineeringUnits) IsKilogramsPerSecond() bool { return a == 0x002A }

// SetKilogramsPerSecond sets AnalogOutputEngineeringUnits to Kilograms-per-second (0x002A)
func (a *AnalogOutputEngineeringUnits) SetKilogramsPerSecond() { *a = 0x002A }

// IsKilogramsPerMinute checks if AnalogOutputEngineeringUnits equals the value for Kilograms-per-minute (0x002B)
func (a AnalogOutputEngineeringUnits) IsKilogramsPerMinute() bool { return a == 0x002B }

// SetKilogramsPerMinute sets AnalogOutputEngineeringUnits to Kilograms-per-minute (0x002B)
func (a *AnalogOutputEngineeringUnits) SetKilogramsPerMinute() { *a = 0x002B }

// IsKilogramsPerHour checks if AnalogOutputEngineeringUnits equals the value for Kilograms-per-hour (0x002C)
func (a AnalogOutputEngineeringUnits) IsKilogramsPerHour() bool { return a == 0x002C }

// SetKilogramsPerHour sets AnalogOutputEngineeringUnits to Kilograms-per-hour (0x002C)
func (a *AnalogOutputEngineeringUnits) SetKilogramsPerHour() { *a = 0x002C }

// IsPoundsMassPerMinute checks if AnalogOutputEngineeringUnits equals the value for Pounds-mass-per-minute (0x002D)
func (a AnalogOutputEngineeringUnits) IsPoundsMassPerMinute() bool { return a == 0x002D }

// SetPoundsMassPerMinute sets AnalogOutputEngineeringUnits to Pounds-mass-per-minute (0x002D)
func (a *AnalogOutputEngineeringUnits) SetPoundsMassPerMinute() { *a = 0x002D }

// IsPoundsMassPerHour checks if AnalogOutputEngineeringUnits equals the value for Pounds-mass-per-hour (0x002E)
func (a AnalogOutputEngineeringUnits) IsPoundsMassPerHour() bool { return a == 0x002E }

// SetPoundsMassPerHour sets AnalogOutputEngineeringUnits to Pounds-mass-per-hour (0x002E)
func (a *AnalogOutputEngineeringUnits) SetPoundsMassPerHour() { *a = 0x002E }

// IsWatts checks if AnalogOutputEngineeringUnits equals the value for Watts (0x002F)
func (a AnalogOutputEngineeringUnits) IsWatts() bool { return a == 0x002F }

// SetWatts sets AnalogOutputEngineeringUnits to Watts (0x002F)
func (a *AnalogOutputEngineeringUnits) SetWatts() { *a = 0x002F }

// IsKilowatts checks if AnalogOutputEngineeringUnits equals the value for Kilowatts (0x0030)
func (a AnalogOutputEngineeringUnits) IsKilowatts() bool { return a == 0x0030 }

// SetKilowatts sets AnalogOutputEngineeringUnits to Kilowatts (0x0030)
func (a *AnalogOutputEngineeringUnits) SetKilowatts() { *a = 0x0030 }

// IsMegawatts checks if AnalogOutputEngineeringUnits equals the value for Megawatts (0x0031)
func (a AnalogOutputEngineeringUnits) IsMegawatts() bool { return a == 0x0031 }

// SetMegawatts sets AnalogOutputEngineeringUnits to Megawatts (0x0031)
func (a *AnalogOutputEngineeringUnits) SetMegawatts() { *a = 0x0031 }

// IsBtusPerHour checks if AnalogOutputEngineeringUnits equals the value for BTUs-per-hour (0x0032)
func (a AnalogOutputEngineeringUnits) IsBtusPerHour() bool { return a == 0x0032 }

// SetBtusPerHour sets AnalogOutputEngineeringUnits to BTUs-per-hour (0x0032)
func (a *AnalogOutputEngineeringUnits) SetBtusPerHour() { *a = 0x0032 }

// IsHorsepower checks if AnalogOutputEngineeringUnits equals the value for Horsepower (0x0033)
func (a AnalogOutputEngineeringUnits) IsHorsepower() bool { return a == 0x0033 }

// SetHorsepower sets AnalogOutputEngineeringUnits to Horsepower (0x0033)
func (a *AnalogOutputEngineeringUnits) SetHorsepower() { *a = 0x0033 }

// IsTonsRefrigeration checks if AnalogOutputEngineeringUnits equals the value for Tons-refrigeration (0x0034)
func (a AnalogOutputEngineeringUnits) IsTonsRefrigeration() bool { return a == 0x0034 }

// SetTonsRefrigeration sets AnalogOutputEngineeringUnits to Tons-refrigeration (0x0034)
func (a *AnalogOutputEngineeringUnits) SetTonsRefrigeration() { *a = 0x0034 }

// IsPascals checks if AnalogOutputEngineeringUnits equals the value for Pascals (0x0035)
func (a AnalogOutputEngineeringUnits) IsPascals() bool { return a == 0x0035 }

// SetPascals sets AnalogOutputEngineeringUnits to Pascals (0x0035)
func (a *AnalogOutputEngineeringUnits) SetPascals() { *a = 0x0035 }

// IsKilopascals checks if AnalogOutputEngineeringUnits equals the value for Kilopascals (0x0036)
func (a AnalogOutputEngineeringUnits) IsKilopascals() bool { return a == 0x0036 }

// SetKilopascals sets AnalogOutputEngineeringUnits to Kilopascals (0x0036)
func (a *AnalogOutputEngineeringUnits) SetKilopascals() { *a = 0x0036 }

// IsBars checks if AnalogOutputEngineeringUnits equals the value for Bars (0x0037)
func (a AnalogOutputEngineeringUnits) IsBars() bool { return a == 0x0037 }

// SetBars sets AnalogOutputEngineeringUnits to Bars (0x0037)
func (a *AnalogOutputEngineeringUnits) SetBars() { *a = 0x0037 }

// IsPoundsForcePerSquareInch checks if AnalogOutputEngineeringUnits equals the value for Pounds-force-per-square-inch (0x0038)
func (a AnalogOutputEngineeringUnits) IsPoundsForcePerSquareInch() bool { return a == 0x0038 }

// SetPoundsForcePerSquareInch sets AnalogOutputEngineeringUnits to Pounds-force-per-square-inch (0x0038)
func (a *AnalogOutputEngineeringUnits) SetPoundsForcePerSquareInch() { *a = 0x0038 }

// IsCentimetersOfWater checks if AnalogOutputEngineeringUnits equals the value for Centimeters-of-water (0x0039)
func (a AnalogOutputEngineeringUnits) IsCentimetersOfWater() bool { return a == 0x0039 }

// SetCentimetersOfWater sets AnalogOutputEngineeringUnits to Centimeters-of-water (0x0039)
func (a *AnalogOutputEngineeringUnits) SetCentimetersOfWater() { *a = 0x0039 }

// IsInchesOfWater checks if AnalogOutputEngineeringUnits equals the value for Inches-of-water (0x003A)
func (a AnalogOutputEngineeringUnits) IsInchesOfWater() bool { return a == 0x003A }

// SetInchesOfWater sets AnalogOutputEngineeringUnits to Inches-of-water (0x003A)
func (a *AnalogOutputEngineeringUnits) SetInchesOfWater() { *a = 0x003A }

// IsMillimetersOfMercury checks if AnalogOutputEngineeringUnits equals the value for Millimeters-of-mercury (0x003B)
func (a AnalogOutputEngineeringUnits) IsMillimetersOfMercury() bool { return a == 0x003B }

// SetMillimetersOfMercury sets AnalogOutputEngineeringUnits to Millimeters-of-mercury (0x003B)
func (a *AnalogOutputEngineeringUnits) SetMillimetersOfMercury() { *a = 0x003B }

// IsCentimetersOfMercury checks if AnalogOutputEngineeringUnits equals the value for Centimeters-of-mercury (0x003C)
func (a AnalogOutputEngineeringUnits) IsCentimetersOfMercury() bool { return a == 0x003C }

// SetCentimetersOfMercury sets AnalogOutputEngineeringUnits to Centimeters-of-mercury (0x003C)
func (a *AnalogOutputEngineeringUnits) SetCentimetersOfMercury() { *a = 0x003C }

// IsInchesOfMercury checks if AnalogOutputEngineeringUnits equals the value for Inches-of-mercury (0x003D)
func (a AnalogOutputEngineeringUnits) IsInchesOfMercury() bool { return a == 0x003D }

// SetInchesOfMercury sets AnalogOutputEngineeringUnits to Inches-of-mercury (0x003D)
func (a *AnalogOutputEngineeringUnits) SetInchesOfMercury() { *a = 0x003D }

// IsDegreesCelsius checks if AnalogOutputEngineeringUnits equals the value for Degrees-Celsius (0x003E)
func (a AnalogOutputEngineeringUnits) IsDegreesCelsius() bool { return a == 0x003E }

// SetDegreesCelsius sets AnalogOutputEngineeringUnits to Degrees-Celsius (0x003E)
func (a *AnalogOutputEngineeringUnits) SetDegreesCelsius() { *a = 0x003E }

// IsDegreesKelvin checks if AnalogOutputEngineeringUnits equals the value for Degrees-Kelvin (0x003F)
func (a AnalogOutputEngineeringUnits) IsDegreesKelvin() bool { return a == 0x003F }

// SetDegreesKelvin sets AnalogOutputEngineeringUnits to Degrees-Kelvin (0x003F)
func (a *AnalogOutputEngineeringUnits) SetDegreesKelvin() { *a = 0x003F }

// IsDegreesFahrenheit checks if AnalogOutputEngineeringUnits equals the value for Degrees-Fahrenheit (0x0040)
func (a AnalogOutputEngineeringUnits) IsDegreesFahrenheit() bool { return a == 0x0040 }

// SetDegreesFahrenheit sets AnalogOutputEngineeringUnits to Degrees-Fahrenheit (0x0040)
func (a *AnalogOutputEngineeringUnits) SetDegreesFahrenheit() { *a = 0x0040 }

// IsDegreeDaysCelsius checks if AnalogOutputEngineeringUnits equals the value for Degree-days-Celsius (0x0041)
func (a AnalogOutputEngineeringUnits) IsDegreeDaysCelsius() bool { return a == 0x0041 }

// SetDegreeDaysCelsius sets AnalogOutputEngineeringUnits to Degree-days-Celsius (0x0041)
func (a *AnalogOutputEngineeringUnits) SetDegreeDaysCelsius() { *a = 0x0041 }

// IsDegreeDaysFahrenheit checks if AnalogOutputEngineeringUnits equals the value for Degree-days-Fahrenheit (0x0042)
func (a AnalogOutputEngineeringUnits) IsDegreeDaysFahrenheit() bool { return a == 0x0042 }

// SetDegreeDaysFahrenheit sets AnalogOutputEngineeringUnits to Degree-days-Fahrenheit (0x0042)
func (a *AnalogOutputEngineeringUnits) SetDegreeDaysFahrenheit() { *a = 0x0042 }

// IsYears checks if AnalogOutputEngineeringUnits equals the value for Years (0x0043)
func (a AnalogOutputEngineeringUnits) IsYears() bool { return a == 0x0043 }

// SetYears sets AnalogOutputEngineeringUnits to Years (0x0043)
func (a *AnalogOutputEngineeringUnits) SetYears() { *a = 0x0043 }

// IsMonths checks if AnalogOutputEngineeringUnits equals the value for Months (0x0044)
func (a AnalogOutputEngineeringUnits) IsMonths() bool { return a == 0x0044 }

// SetMonths sets AnalogOutputEngineeringUnits to Months (0x0044)
func (a *AnalogOutputEngineeringUnits) SetMonths() { *a = 0x0044 }

// IsWeeks checks if AnalogOutputEngineeringUnits equals the value for Weeks (0x0045)
func (a AnalogOutputEngineeringUnits) IsWeeks() bool { return a == 0x0045 }

// SetWeeks sets AnalogOutputEngineeringUnits to Weeks (0x0045)
func (a *AnalogOutputEngineeringUnits) SetWeeks() { *a = 0x0045 }

// IsDays checks if AnalogOutputEngineeringUnits equals the value for Days (0x0046)
func (a AnalogOutputEngineeringUnits) IsDays() bool { return a == 0x0046 }

// SetDays sets AnalogOutputEngineeringUnits to Days (0x0046)
func (a *AnalogOutputEngineeringUnits) SetDays() { *a = 0x0046 }

// IsHours checks if AnalogOutputEngineeringUnits equals the value for Hours (0x0047)
func (a AnalogOutputEngineeringUnits) IsHours() bool { return a == 0x0047 }

// SetHours sets AnalogOutputEngineeringUnits to Hours (0x0047)
func (a *AnalogOutputEngineeringUnits) SetHours() { *a = 0x0047 }

// IsMinutes checks if AnalogOutputEngineeringUnits equals the value for Minutes (0x0048)
func (a AnalogOutputEngineeringUnits) IsMinutes() bool { return a == 0x0048 }

// SetMinutes sets AnalogOutputEngineeringUnits to Minutes (0x0048)
func (a *AnalogOutputEngineeringUnits) SetMinutes() { *a = 0x0048 }

// IsSeconds checks if AnalogOutputEngineeringUnits equals the value for Seconds (0x0049)
func (a AnalogOutputEngineeringUnits) IsSeconds() bool { return a == 0x0049 }

// SetSeconds sets AnalogOutputEngineeringUnits to Seconds (0x0049)
func (a *AnalogOutputEngineeringUnits) SetSeconds() { *a = 0x0049 }

// IsMetersPerSecond checks if AnalogOutputEngineeringUnits equals the value for Meters-per-second (0x004A)
func (a AnalogOutputEngineeringUnits) IsMetersPerSecond() bool { return a == 0x004A }

// SetMetersPerSecond sets AnalogOutputEngineeringUnits to Meters-per-second (0x004A)
func (a *AnalogOutputEngineeringUnits) SetMetersPerSecond() { *a = 0x004A }

// IsKilometersPerHour checks if AnalogOutputEngineeringUnits equals the value for Kilometers-per-hour (0x004B)
func (a AnalogOutputEngineeringUnits) IsKilometersPerHour() bool { return a == 0x004B }

// SetKilometersPerHour sets AnalogOutputEngineeringUnits to Kilometers-per-hour (0x004B)
func (a *AnalogOutputEngineeringUnits) SetKilometersPerHour() { *a = 0x004B }

// IsFeetPerSecond checks if AnalogOutputEngineeringUnits equals the value for Feet-per-second (0x004C)
func (a AnalogOutputEngineeringUnits) IsFeetPerSecond() bool { return a == 0x004C }

// SetFeetPerSecond sets AnalogOutputEngineeringUnits to Feet-per-second (0x004C)
func (a *AnalogOutputEngineeringUnits) SetFeetPerSecond() { *a = 0x004C }

// IsFeetPerMinute checks if AnalogOutputEngineeringUnits equals the value for Feet-per-minute (0x004D)
func (a AnalogOutputEngineeringUnits) IsFeetPerMinute() bool { return a == 0x004D }

// SetFeetPerMinute sets AnalogOutputEngineeringUnits to Feet-per-minute (0x004D)
func (a *AnalogOutputEngineeringUnits) SetFeetPerMinute() { *a = 0x004D }

// IsMilesPerHour checks if AnalogOutputEngineeringUnits equals the value for Miles-per-hour (0x004E)
func (a AnalogOutputEngineeringUnits) IsMilesPerHour() bool { return a == 0x004E }

// SetMilesPerHour sets AnalogOutputEngineeringUnits to Miles-per-hour (0x004E)
func (a *AnalogOutputEngineeringUnits) SetMilesPerHour() { *a = 0x004E }

// IsCubicFeet checks if AnalogOutputEngineeringUnits equals the value for Cubic-feet (0x004F)
func (a AnalogOutputEngineeringUnits) IsCubicFeet() bool { return a == 0x004F }

// SetCubicFeet sets AnalogOutputEngineeringUnits to Cubic-feet (0x004F)
func (a *AnalogOutputEngineeringUnits) SetCubicFeet() { *a = 0x004F }

// IsCubicMeters checks if AnalogOutputEngineeringUnits equals the value for Cubic-meters (0x0050)
func (a AnalogOutputEngineeringUnits) IsCubicMeters() bool { return a == 0x0050 }

// SetCubicMeters sets AnalogOutputEngineeringUnits to Cubic-meters (0x0050)
func (a *AnalogOutputEngineeringUnits) SetCubicMeters() { *a = 0x0050 }

// IsImperialGallons checks if AnalogOutputEngineeringUnits equals the value for Imperial-gallons (0x0051)
func (a AnalogOutputEngineeringUnits) IsImperialGallons() bool { return a == 0x0051 }

// SetImperialGallons sets AnalogOutputEngineeringUnits to Imperial-gallons (0x0051)
func (a *AnalogOutputEngineeringUnits) SetImperialGallons() { *a = 0x0051 }

// IsLiters checks if AnalogOutputEngineeringUnits equals the value for Liters (0x0052)
func (a AnalogOutputEngineeringUnits) IsLiters() bool { return a == 0x0052 }

// SetLiters sets AnalogOutputEngineeringUnits to Liters (0x0052)
func (a *AnalogOutputEngineeringUnits) SetLiters() { *a = 0x0052 }

// IsUsGallons checks if AnalogOutputEngineeringUnits equals the value for Us-gallons (0x0053)
func (a AnalogOutputEngineeringUnits) IsUsGallons() bool { return a == 0x0053 }

// SetUsGallons sets AnalogOutputEngineeringUnits to Us-gallons (0x0053)
func (a *AnalogOutputEngineeringUnits) SetUsGallons() { *a = 0x0053 }

// IsCubicFeetPerMinute checks if AnalogOutputEngineeringUnits equals the value for Cubic-feet-per-minute (0x0054)
func (a AnalogOutputEngineeringUnits) IsCubicFeetPerMinute() bool { return a == 0x0054 }

// SetCubicFeetPerMinute sets AnalogOutputEngineeringUnits to Cubic-feet-per-minute (0x0054)
func (a *AnalogOutputEngineeringUnits) SetCubicFeetPerMinute() { *a = 0x0054 }

// IsCubicMetersPerSecond checks if AnalogOutputEngineeringUnits equals the value for Cubic-meters-per-second (0x0055)
func (a AnalogOutputEngineeringUnits) IsCubicMetersPerSecond() bool { return a == 0x0055 }

// SetCubicMetersPerSecond sets AnalogOutputEngineeringUnits to Cubic-meters-per-second (0x0055)
func (a *AnalogOutputEngineeringUnits) SetCubicMetersPerSecond() { *a = 0x0055 }

// IsImperialGallonsPerMinute checks if AnalogOutputEngineeringUnits equals the value for Imperial-gallons-per-minute (0x0056)
func (a AnalogOutputEngineeringUnits) IsImperialGallonsPerMinute() bool { return a == 0x0056 }

// SetImperialGallonsPerMinute sets AnalogOutputEngineeringUnits to Imperial-gallons-per-minute (0x0056)
func (a *AnalogOutputEngineeringUnits) SetImperialGallonsPerMinute() { *a = 0x0056 }

// IsLitersPerSecond checks if AnalogOutputEngineeringUnits equals the value for Liters-per-second (0x0057)
func (a AnalogOutputEngineeringUnits) IsLitersPerSecond() bool { return a == 0x0057 }

// SetLitersPerSecond sets AnalogOutputEngineeringUnits to Liters-per-second (0x0057)
func (a *AnalogOutputEngineeringUnits) SetLitersPerSecond() { *a = 0x0057 }

// IsLitersPerMinute checks if AnalogOutputEngineeringUnits equals the value for Liters-per-minute (0x0058)
func (a AnalogOutputEngineeringUnits) IsLitersPerMinute() bool { return a == 0x0058 }

// SetLitersPerMinute sets AnalogOutputEngineeringUnits to Liters-per-minute (0x0058)
func (a *AnalogOutputEngineeringUnits) SetLitersPerMinute() { *a = 0x0058 }

// IsUsGallonsPerMinute checks if AnalogOutputEngineeringUnits equals the value for Us-gallons-per-minute (0x0059)
func (a AnalogOutputEngineeringUnits) IsUsGallonsPerMinute() bool { return a == 0x0059 }

// SetUsGallonsPerMinute sets AnalogOutputEngineeringUnits to Us-gallons-per-minute (0x0059)
func (a *AnalogOutputEngineeringUnits) SetUsGallonsPerMinute() { *a = 0x0059 }

// IsDegreesAngular checks if AnalogOutputEngineeringUnits equals the value for Degrees-angular (0x005A)
func (a AnalogOutputEngineeringUnits) IsDegreesAngular() bool { return a == 0x005A }

// SetDegreesAngular sets AnalogOutputEngineeringUnits to Degrees-angular (0x005A)
func (a *AnalogOutputEngineeringUnits) SetDegreesAngular() { *a = 0x005A }

// IsDegreesCelsiusPerHour checks if AnalogOutputEngineeringUnits equals the value for Degrees-Celsius-per-hour (0x005B)
func (a AnalogOutputEngineeringUnits) IsDegreesCelsiusPerHour() bool { return a == 0x005B }

// SetDegreesCelsiusPerHour sets AnalogOutputEngineeringUnits to Degrees-Celsius-per-hour (0x005B)
func (a *AnalogOutputEngineeringUnits) SetDegreesCelsiusPerHour() { *a = 0x005B }

// IsDegreesCelsiusPerMinute checks if AnalogOutputEngineeringUnits equals the value for Degrees-Celsius-per-minute (0x005C)
func (a AnalogOutputEngineeringUnits) IsDegreesCelsiusPerMinute() bool { return a == 0x005C }

// SetDegreesCelsiusPerMinute sets AnalogOutputEngineeringUnits to Degrees-Celsius-per-minute (0x005C)
func (a *AnalogOutputEngineeringUnits) SetDegreesCelsiusPerMinute() { *a = 0x005C }

// IsDegreesFahrenheitPerHour checks if AnalogOutputEngineeringUnits equals the value for Degrees-Fahrenheit-per-hour (0x005D)
func (a AnalogOutputEngineeringUnits) IsDegreesFahrenheitPerHour() bool { return a == 0x005D }

// SetDegreesFahrenheitPerHour sets AnalogOutputEngineeringUnits to Degrees-Fahrenheit-per-hour (0x005D)
func (a *AnalogOutputEngineeringUnits) SetDegreesFahrenheitPerHour() { *a = 0x005D }

// IsDegreesFahrenheitPerMinute checks if AnalogOutputEngineeringUnits equals the value for Degrees-Fahrenheit-per-minute (0x005E)
func (a AnalogOutputEngineeringUnits) IsDegreesFahrenheitPerMinute() bool { return a == 0x005E }

// SetDegreesFahrenheitPerMinute sets AnalogOutputEngineeringUnits to Degrees-Fahrenheit-per-minute (0x005E)
func (a *AnalogOutputEngineeringUnits) SetDegreesFahrenheitPerMinute() { *a = 0x005E }

// IsNoUnits checks if AnalogOutputEngineeringUnits equals the value for No-units (0x005F)
func (a AnalogOutputEngineeringUnits) IsNoUnits() bool { return a == 0x005F }

// SetNoUnits sets AnalogOutputEngineeringUnits to No-units (0x005F)
func (a *AnalogOutputEngineeringUnits) SetNoUnits() { *a = 0x005F }

// IsPartsPerMillion checks if AnalogOutputEngineeringUnits equals the value for Parts-per-million (0x0060)
func (a AnalogOutputEngineeringUnits) IsPartsPerMillion() bool { return a == 0x0060 }

// SetPartsPerMillion sets AnalogOutputEngineeringUnits to Parts-per-million (0x0060)
func (a *AnalogOutputEngineeringUnits) SetPartsPerMillion() { *a = 0x0060 }

// IsPartsPerBillion checks if AnalogOutputEngineeringUnits equals the value for Parts-per-billion (0x0061)
func (a AnalogOutputEngineeringUnits) IsPartsPerBillion() bool { return a == 0x0061 }

// SetPartsPerBillion sets AnalogOutputEngineeringUnits to Parts-per-billion (0x0061)
func (a *AnalogOutputEngineeringUnits) SetPartsPerBillion() { *a = 0x0061 }

// IsPercent checks if AnalogOutputEngineeringUnits equals the value for Percent (0x0062)
func (a AnalogOutputEngineeringUnits) IsPercent() bool { return a == 0x0062 }

// SetPercent sets AnalogOutputEngineeringUnits to Percent (0x0062)
func (a *AnalogOutputEngineeringUnits) SetPercent() { *a = 0x0062 }

// IsPercentPerSecond checks if AnalogOutputEngineeringUnits equals the value for Percent-per-second (0x0063)
func (a AnalogOutputEngineeringUnits) IsPercentPerSecond() bool { return a == 0x0063 }

// SetPercentPerSecond sets AnalogOutputEngineeringUnits to Percent-per-second (0x0063)
func (a *AnalogOutputEngineeringUnits) SetPercentPerSecond() { *a = 0x0063 }

// IsPerMinute checks if AnalogOutputEngineeringUnits equals the value for Per-minute (0x0064)
func (a AnalogOutputEngineeringUnits) IsPerMinute() bool { return a == 0x0064 }

// SetPerMinute sets AnalogOutputEngineeringUnits to Per-minute (0x0064)
func (a *AnalogOutputEngineeringUnits) SetPerMinute() { *a = 0x0064 }

// IsPerSecond checks if AnalogOutputEngineeringUnits equals the value for Per-second (0x0065)
func (a AnalogOutputEngineeringUnits) IsPerSecond() bool { return a == 0x0065 }

// SetPerSecond sets AnalogOutputEngineeringUnits to Per-second (0x0065)
func (a *AnalogOutputEngineeringUnits) SetPerSecond() { *a = 0x0065 }

// IsPsiPerDegreeFahrenheit checks if AnalogOutputEngineeringUnits equals the value for Psi-per-Degree-Fahrenheit (0x0066)
func (a AnalogOutputEngineeringUnits) IsPsiPerDegreeFahrenheit() bool { return a == 0x0066 }

// SetPsiPerDegreeFahrenheit sets AnalogOutputEngineeringUnits to Psi-per-Degree-Fahrenheit (0x0066)
func (a *AnalogOutputEngineeringUnits) SetPsiPerDegreeFahrenheit() { *a = 0x0066 }

// IsRadians checks if AnalogOutputEngineeringUnits equals the value for Radians (0x0067)
func (a AnalogOutputEngineeringUnits) IsRadians() bool { return a == 0x0067 }

// SetRadians sets AnalogOutputEngineeringUnits to Radians (0x0067)
func (a *AnalogOutputEngineeringUnits) SetRadians() { *a = 0x0067 }

// IsRevolutionsPerMinute checks if AnalogOutputEngineeringUnits equals the value for Revolutions-per-minute (0x0068)
func (a AnalogOutputEngineeringUnits) IsRevolutionsPerMinute() bool { return a == 0x0068 }

// SetRevolutionsPerMinute sets AnalogOutputEngineeringUnits to Revolutions-per-minute (0x0068)
func (a *AnalogOutputEngineeringUnits) SetRevolutionsPerMinute() { *a = 0x0068 }

// IsCurrency1 checks if AnalogOutputEngineeringUnits equals the value for Currency1 (0x0069)
func (a AnalogOutputEngineeringUnits) IsCurrency1() bool { return a == 0x0069 }

// SetCurrency1 sets AnalogOutputEngineeringUnits to Currency1 (0x0069)
func (a *AnalogOutputEngineeringUnits) SetCurrency1() { *a = 0x0069 }

// IsCurrency2 checks if AnalogOutputEngineeringUnits equals the value for Currency2 (0x006A)
func (a AnalogOutputEngineeringUnits) IsCurrency2() bool { return a == 0x006A }

// SetCurrency2 sets AnalogOutputEngineeringUnits to Currency2 (0x006A)
func (a *AnalogOutputEngineeringUnits) SetCurrency2() { *a = 0x006A }

// IsCurrency3 checks if AnalogOutputEngineeringUnits equals the value for Currency3 (0x006B)
func (a AnalogOutputEngineeringUnits) IsCurrency3() bool { return a == 0x006B }

// SetCurrency3 sets AnalogOutputEngineeringUnits to Currency3 (0x006B)
func (a *AnalogOutputEngineeringUnits) SetCurrency3() { *a = 0x006B }

// IsCurrency4 checks if AnalogOutputEngineeringUnits equals the value for Currency4 (0x006C)
func (a AnalogOutputEngineeringUnits) IsCurrency4() bool { return a == 0x006C }

// SetCurrency4 sets AnalogOutputEngineeringUnits to Currency4 (0x006C)
func (a *AnalogOutputEngineeringUnits) SetCurrency4() { *a = 0x006C }

// IsCurrency5 checks if AnalogOutputEngineeringUnits equals the value for Currency5 (0x006D)
func (a AnalogOutputEngineeringUnits) IsCurrency5() bool { return a == 0x006D }

// SetCurrency5 sets AnalogOutputEngineeringUnits to Currency5 (0x006D)
func (a *AnalogOutputEngineeringUnits) SetCurrency5() { *a = 0x006D }

// IsCurrency6 checks if AnalogOutputEngineeringUnits equals the value for Currency6 (0x006E)
func (a AnalogOutputEngineeringUnits) IsCurrency6() bool { return a == 0x006E }

// SetCurrency6 sets AnalogOutputEngineeringUnits to Currency6 (0x006E)
func (a *AnalogOutputEngineeringUnits) SetCurrency6() { *a = 0x006E }

// IsCurrency7 checks if AnalogOutputEngineeringUnits equals the value for Currency7 (0x006F)
func (a AnalogOutputEngineeringUnits) IsCurrency7() bool { return a == 0x006F }

// SetCurrency7 sets AnalogOutputEngineeringUnits to Currency7 (0x006F)
func (a *AnalogOutputEngineeringUnits) SetCurrency7() { *a = 0x006F }

// IsCurrency8 checks if AnalogOutputEngineeringUnits equals the value for Currency8 (0x0070)
func (a AnalogOutputEngineeringUnits) IsCurrency8() bool { return a == 0x0070 }

// SetCurrency8 sets AnalogOutputEngineeringUnits to Currency8 (0x0070)
func (a *AnalogOutputEngineeringUnits) SetCurrency8() { *a = 0x0070 }

// IsCurrency9 checks if AnalogOutputEngineeringUnits equals the value for Currency9 (0x0071)
func (a AnalogOutputEngineeringUnits) IsCurrency9() bool { return a == 0x0071 }

// SetCurrency9 sets AnalogOutputEngineeringUnits to Currency9 (0x0071)
func (a *AnalogOutputEngineeringUnits) SetCurrency9() { *a = 0x0071 }

// IsCurrency10 checks if AnalogOutputEngineeringUnits equals the value for Currency10 (0x0072)
func (a AnalogOutputEngineeringUnits) IsCurrency10() bool { return a == 0x0072 }

// SetCurrency10 sets AnalogOutputEngineeringUnits to Currency10 (0x0072)
func (a *AnalogOutputEngineeringUnits) SetCurrency10() { *a = 0x0072 }

// IsSquareInches checks if AnalogOutputEngineeringUnits equals the value for Square-inches (0x0073)
func (a AnalogOutputEngineeringUnits) IsSquareInches() bool { return a == 0x0073 }

// SetSquareInches sets AnalogOutputEngineeringUnits to Square-inches (0x0073)
func (a *AnalogOutputEngineeringUnits) SetSquareInches() { *a = 0x0073 }

// IsSquareCentimeters checks if AnalogOutputEngineeringUnits equals the value for Square-centimeters (0x0074)
func (a AnalogOutputEngineeringUnits) IsSquareCentimeters() bool { return a == 0x0074 }

// SetSquareCentimeters sets AnalogOutputEngineeringUnits to Square-centimeters (0x0074)
func (a *AnalogOutputEngineeringUnits) SetSquareCentimeters() { *a = 0x0074 }

// IsBtusPerPound checks if AnalogOutputEngineeringUnits equals the value for BTUs-per-pound (0x0075)
func (a AnalogOutputEngineeringUnits) IsBtusPerPound() bool { return a == 0x0075 }

// SetBtusPerPound sets AnalogOutputEngineeringUnits to BTUs-per-pound (0x0075)
func (a *AnalogOutputEngineeringUnits) SetBtusPerPound() { *a = 0x0075 }

// IsCentimeters checks if AnalogOutputEngineeringUnits equals the value for Centimeters (0x0076)
func (a AnalogOutputEngineeringUnits) IsCentimeters() bool { return a == 0x0076 }

// SetCentimeters sets AnalogOutputEngineeringUnits to Centimeters (0x0076)
func (a *AnalogOutputEngineeringUnits) SetCentimeters() { *a = 0x0076 }

// IsPoundsMassPerSecond checks if AnalogOutputEngineeringUnits equals the value for Pounds-mass-per-second (0x0077)
func (a AnalogOutputEngineeringUnits) IsPoundsMassPerSecond() bool { return a == 0x0077 }

// SetPoundsMassPerSecond sets AnalogOutputEngineeringUnits to Pounds-mass-per-second (0x0077)
func (a *AnalogOutputEngineeringUnits) SetPoundsMassPerSecond() { *a = 0x0077 }

// IsDeltaDegreesFahrenheit checks if AnalogOutputEngineeringUnits equals the value for Delta-Degrees-Fahrenheit (0x0078)
func (a AnalogOutputEngineeringUnits) IsDeltaDegreesFahrenheit() bool { return a == 0x0078 }

// SetDeltaDegreesFahrenheit sets AnalogOutputEngineeringUnits to Delta-Degrees-Fahrenheit (0x0078)
func (a *AnalogOutputEngineeringUnits) SetDeltaDegreesFahrenheit() { *a = 0x0078 }

// IsDeltaDegreesKelvin checks if AnalogOutputEngineeringUnits equals the value for Delta-Degrees-Kelvin (0x0079)
func (a AnalogOutputEngineeringUnits) IsDeltaDegreesKelvin() bool { return a == 0x0079 }

// SetDeltaDegreesKelvin sets AnalogOutputEngineeringUnits to Delta-Degrees-Kelvin (0x0079)
func (a *AnalogOutputEngineeringUnits) SetDeltaDegreesKelvin() { *a = 0x0079 }

// IsKilohms checks if AnalogOutputEngineeringUnits equals the value for Kilohms (0x007A)
func (a AnalogOutputEngineeringUnits) IsKilohms() bool { return a == 0x007A }

// SetKilohms sets AnalogOutputEngineeringUnits to Kilohms (0x007A)
func (a *AnalogOutputEngineeringUnits) SetKilohms() { *a = 0x007A }

// IsMegohms checks if AnalogOutputEngineeringUnits equals the value for Megohms (0x007B)
func (a AnalogOutputEngineeringUnits) IsMegohms() bool { return a == 0x007B }

// SetMegohms sets AnalogOutputEngineeringUnits to Megohms (0x007B)
func (a *AnalogOutputEngineeringUnits) SetMegohms() { *a = 0x007B }

// IsMillivolts checks if AnalogOutputEngineeringUnits equals the value for Millivolts (0x007C)
func (a AnalogOutputEngineeringUnits) IsMillivolts() bool { return a == 0x007C }

// SetMillivolts sets AnalogOutputEngineeringUnits to Millivolts (0x007C)
func (a *AnalogOutputEngineeringUnits) SetMillivolts() { *a = 0x007C }

// IsKilojoulesPerKilogram checks if AnalogOutputEngineeringUnits equals the value for Kilojoules-per-kilogram (0x007D)
func (a AnalogOutputEngineeringUnits) IsKilojoulesPerKilogram() bool { return a == 0x007D }

// SetKilojoulesPerKilogram sets AnalogOutputEngineeringUnits to Kilojoules-per-kilogram (0x007D)
func (a *AnalogOutputEngineeringUnits) SetKilojoulesPerKilogram() { *a = 0x007D }

// IsMegajoules checks if AnalogOutputEngineeringUnits equals the value for Megajoules (0x007E)
func (a AnalogOutputEngineeringUnits) IsMegajoules() bool { return a == 0x007E }

// SetMegajoules sets AnalogOutputEngineeringUnits to Megajoules (0x007E)
func (a *AnalogOutputEngineeringUnits) SetMegajoules() { *a = 0x007E }

// IsJoulesPerDegreeKelvin checks if AnalogOutputEngineeringUnits equals the value for Joules-per-degree-Kelvin (0x007F)
func (a AnalogOutputEngineeringUnits) IsJoulesPerDegreeKelvin() bool { return a == 0x007F }

// SetJoulesPerDegreeKelvin sets AnalogOutputEngineeringUnits to Joules-per-degree-Kelvin (0x007F)
func (a *AnalogOutputEngineeringUnits) SetJoulesPerDegreeKelvin() { *a = 0x007F }

// IsJoulesPerKilogramDegreeKelvin checks if AnalogOutputEngineeringUnits equals the value for Joules-per-kilogram-degree-Kelvin (0x0080)
func (a AnalogOutputEngineeringUnits) IsJoulesPerKilogramDegreeKelvin() bool { return a == 0x0080 }

// SetJoulesPerKilogramDegreeKelvin sets AnalogOutputEngineeringUnits to Joules-per-kilogram-degree-Kelvin (0x0080)
func (a *AnalogOutputEngineeringUnits) SetJoulesPerKilogramDegreeKelvin() { *a = 0x0080 }

// IsKilohertz checks if AnalogOutputEngineeringUnits equals the value for Kilohertz (0x0081)
func (a AnalogOutputEngineeringUnits) IsKilohertz() bool { return a == 0x0081 }

// SetKilohertz sets AnalogOutputEngineeringUnits to Kilohertz (0x0081)
func (a *AnalogOutputEngineeringUnits) SetKilohertz() { *a = 0x0081 }

// IsMegahertz checks if AnalogOutputEngineeringUnits equals the value for Megahertz (0x0082)
func (a AnalogOutputEngineeringUnits) IsMegahertz() bool { return a == 0x0082 }

// SetMegahertz sets AnalogOutputEngineeringUnits to Megahertz (0x0082)
func (a *AnalogOutputEngineeringUnits) SetMegahertz() { *a = 0x0082 }

// IsPerHour checks if AnalogOutputEngineeringUnits equals the value for Per-hour (0x0083)
func (a AnalogOutputEngineeringUnits) IsPerHour() bool { return a == 0x0083 }

// SetPerHour sets AnalogOutputEngineeringUnits to Per-hour (0x0083)
func (a *AnalogOutputEngineeringUnits) SetPerHour() { *a = 0x0083 }

// IsMilliwatts checks if AnalogOutputEngineeringUnits equals the value for Milliwatts (0x0084)
func (a AnalogOutputEngineeringUnits) IsMilliwatts() bool { return a == 0x0084 }

// SetMilliwatts sets AnalogOutputEngineeringUnits to Milliwatts (0x0084)
func (a *AnalogOutputEngineeringUnits) SetMilliwatts() { *a = 0x0084 }

// IsHectopascals checks if AnalogOutputEngineeringUnits equals the value for Hectopascals (0x0085)
func (a AnalogOutputEngineeringUnits) IsHectopascals() bool { return a == 0x0085 }

// SetHectopascals sets AnalogOutputEngineeringUnits to Hectopascals (0x0085)
func (a *AnalogOutputEngineeringUnits) SetHectopascals() { *a = 0x0085 }

// IsMillibars checks if AnalogOutputEngineeringUnits equals the value for Millibars (0x0086)
func (a AnalogOutputEngineeringUnits) IsMillibars() bool { return a == 0x0086 }

// SetMillibars sets AnalogOutputEngineeringUnits to Millibars (0x0086)
func (a *AnalogOutputEngineeringUnits) SetMillibars() { *a = 0x0086 }

// IsCubicMetersPerHour checks if AnalogOutputEngineeringUnits equals the value for Cubic-meters-per-hour (0x0087)
func (a AnalogOutputEngineeringUnits) IsCubicMetersPerHour() bool { return a == 0x0087 }

// SetCubicMetersPerHour sets AnalogOutputEngineeringUnits to Cubic-meters-per-hour (0x0087)
func (a *AnalogOutputEngineeringUnits) SetCubicMetersPerHour() { *a = 0x0087 }

// IsLitersPerHour checks if AnalogOutputEngineeringUnits equals the value for Liters-per-hour (0x0088)
func (a AnalogOutputEngineeringUnits) IsLitersPerHour() bool { return a == 0x0088 }

// SetLitersPerHour sets AnalogOutputEngineeringUnits to Liters-per-hour (0x0088)
func (a *AnalogOutputEngineeringUnits) SetLitersPerHour() { *a = 0x0088 }

// IsKilowattHoursPerSquareMeter checks if AnalogOutputEngineeringUnits equals the value for Kilowatt-hours-per-square-meter (0x0089)
func (a AnalogOutputEngineeringUnits) IsKilowattHoursPerSquareMeter() bool { return a == 0x0089 }

// SetKilowattHoursPerSquareMeter sets AnalogOutputEngineeringUnits to Kilowatt-hours-per-square-meter (0x0089)
func (a *AnalogOutputEngineeringUnits) SetKilowattHoursPerSquareMeter() { *a = 0x0089 }

// IsKilowattHoursPerSquareFoot checks if AnalogOutputEngineeringUnits equals the value for Kilowatt-hours-per-square-foot (0x008A)
func (a AnalogOutputEngineeringUnits) IsKilowattHoursPerSquareFoot() bool { return a == 0x008A }

// SetKilowattHoursPerSquareFoot sets AnalogOutputEngineeringUnits to Kilowatt-hours-per-square-foot (0x008A)
func (a *AnalogOutputEngineeringUnits) SetKilowattHoursPerSquareFoot() { *a = 0x008A }

// IsMegajoulesPerSquareMeter checks if AnalogOutputEngineeringUnits equals the value for Megajoules-per-square-meter (0x008B)
func (a AnalogOutputEngineeringUnits) IsMegajoulesPerSquareMeter() bool { return a == 0x008B }

// SetMegajoulesPerSquareMeter sets AnalogOutputEngineeringUnits to Megajoules-per-square-meter (0x008B)
func (a *AnalogOutputEngineeringUnits) SetMegajoulesPerSquareMeter() { *a = 0x008B }

// IsMegajoulesPerSquareFoot checks if AnalogOutputEngineeringUnits equals the value for Megajoules-per-square-foot (0x008C)
func (a AnalogOutputEngineeringUnits) IsMegajoulesPerSquareFoot() bool { return a == 0x008C }

// SetMegajoulesPerSquareFoot sets AnalogOutputEngineeringUnits to Megajoules-per-square-foot (0x008C)
func (a *AnalogOutputEngineeringUnits) SetMegajoulesPerSquareFoot() { *a = 0x008C }

// IsWattsPerSquareMeterDegreeKelvin checks if AnalogOutputEngineeringUnits equals the value for Watts-per-square-meter-Degree-Kelvin (0x008D)
func (a AnalogOutputEngineeringUnits) IsWattsPerSquareMeterDegreeKelvin() bool { return a == 0x008D }

// SetWattsPerSquareMeterDegreeKelvin sets AnalogOutputEngineeringUnits to Watts-per-square-meter-Degree-Kelvin (0x008D)
func (a *AnalogOutputEngineeringUnits) SetWattsPerSquareMeterDegreeKelvin() { *a = 0x008D }

// IsCubicFeetPerSecond checks if AnalogOutputEngineeringUnits equals the value for Cubic-feet-per-second (0x008E)
func (a AnalogOutputEngineeringUnits) IsCubicFeetPerSecond() bool { return a == 0x008E }

// SetCubicFeetPerSecond sets AnalogOutputEngineeringUnits to Cubic-feet-per-second (0x008E)
func (a *AnalogOutputEngineeringUnits) SetCubicFeetPerSecond() { *a = 0x008E }

// IsPercentObscurationPerFoot checks if AnalogOutputEngineeringUnits equals the value for Percent-obscuration-per-foot (0x008F)
func (a AnalogOutputEngineeringUnits) IsPercentObscurationPerFoot() bool { return a == 0x008F }

// SetPercentObscurationPerFoot sets AnalogOutputEngineeringUnits to Percent-obscuration-per-foot (0x008F)
func (a *AnalogOutputEngineeringUnits) SetPercentObscurationPerFoot() { *a = 0x008F }

// IsPercentObscurationPerMeter checks if AnalogOutputEngineeringUnits equals the value for Percent-obscuration-per-meter (0x0090)
func (a AnalogOutputEngineeringUnits) IsPercentObscurationPerMeter() bool { return a == 0x0090 }

// SetPercentObscurationPerMeter sets AnalogOutputEngineeringUnits to Percent-obscuration-per-meter (0x0090)
func (a *AnalogOutputEngineeringUnits) SetPercentObscurationPerMeter() { *a = 0x0090 }

// IsMilliohms checks if AnalogOutputEngineeringUnits equals the value for Milliohms (0x0091)
func (a AnalogOutputEngineeringUnits) IsMilliohms() bool { return a == 0x0091 }

// SetMilliohms sets AnalogOutputEngineeringUnits to Milliohms (0x0091)
func (a *AnalogOutputEngineeringUnits) SetMilliohms() { *a = 0x0091 }

// IsMegawattHours checks if AnalogOutputEngineeringUnits equals the value for Megawatt-hours (0x0092)
func (a AnalogOutputEngineeringUnits) IsMegawattHours() bool { return a == 0x0092 }

// SetMegawattHours sets AnalogOutputEngineeringUnits to Megawatt-hours (0x0092)
func (a *AnalogOutputEngineeringUnits) SetMegawattHours() { *a = 0x0092 }

// IsKiloBtus checks if AnalogOutputEngineeringUnits equals the value for Kilo-BTUs (0x0093)
func (a AnalogOutputEngineeringUnits) IsKiloBtus() bool { return a == 0x0093 }

// SetKiloBtus sets AnalogOutputEngineeringUnits to Kilo-BTUs (0x0093)
func (a *AnalogOutputEngineeringUnits) SetKiloBtus() { *a = 0x0093 }

// IsMegaBtus checks if AnalogOutputEngineeringUnits equals the value for Mega-BTUs (0x0094)
func (a AnalogOutputEngineeringUnits) IsMegaBtus() bool { return a == 0x0094 }

// SetMegaBtus sets AnalogOutputEngineeringUnits to Mega-BTUs (0x0094)
func (a *AnalogOutputEngineeringUnits) SetMegaBtus() { *a = 0x0094 }

// IsKilojoulesPerKilogramDryAir checks if AnalogOutputEngineeringUnits equals the value for Kilojoules-per-kilogram-dry-air (0x0095)
func (a AnalogOutputEngineeringUnits) IsKilojoulesPerKilogramDryAir() bool { return a == 0x0095 }

// SetKilojoulesPerKilogramDryAir sets AnalogOutputEngineeringUnits to Kilojoules-per-kilogram-dry-air (0x0095)
func (a *AnalogOutputEngineeringUnits) SetKilojoulesPerKilogramDryAir() { *a = 0x0095 }

// IsMegajoulesPerKilogramDryAir checks if AnalogOutputEngineeringUnits equals the value for Megajoules-per-kilogram-dry-air (0x0096)
func (a AnalogOutputEngineeringUnits) IsMegajoulesPerKilogramDryAir() bool { return a == 0x0096 }

// SetMegajoulesPerKilogramDryAir sets AnalogOutputEngineeringUnits to Megajoules-per-kilogram-dry-air (0x0096)
func (a *AnalogOutputEngineeringUnits) SetMegajoulesPerKilogramDryAir() { *a = 0x0096 }

// IsKilojoulesPerDegreeKelvin checks if AnalogOutputEngineeringUnits equals the value for Kilojoules-per-degree-Kelvin (0x0097)
func (a AnalogOutputEngineeringUnits) IsKilojoulesPerDegreeKelvin() bool { return a == 0x0097 }

// SetKilojoulesPerDegreeKelvin sets AnalogOutputEngineeringUnits to Kilojoules-per-degree-Kelvin (0x0097)
func (a *AnalogOutputEngineeringUnits) SetKilojoulesPerDegreeKelvin() { *a = 0x0097 }

// IsMegajoulesPerDegreeKelvin checks if AnalogOutputEngineeringUnits equals the value for Megajoules-per-degree-Kelvin (0x0098)
func (a AnalogOutputEngineeringUnits) IsMegajoulesPerDegreeKelvin() bool { return a == 0x0098 }

// SetMegajoulesPerDegreeKelvin sets AnalogOutputEngineeringUnits to Megajoules-per-degree-Kelvin (0x0098)
func (a *AnalogOutputEngineeringUnits) SetMegajoulesPerDegreeKelvin() { *a = 0x0098 }

// IsNewton checks if AnalogOutputEngineeringUnits equals the value for Newton (0x0099)
func (a AnalogOutputEngineeringUnits) IsNewton() bool { return a == 0x0099 }

// SetNewton sets AnalogOutputEngineeringUnits to Newton (0x0099)
func (a *AnalogOutputEngineeringUnits) SetNewton() { *a = 0x0099 }

// IsGramsPerSecond checks if AnalogOutputEngineeringUnits equals the value for Grams-per-second (0x009A)
func (a AnalogOutputEngineeringUnits) IsGramsPerSecond() bool { return a == 0x009A }

// SetGramsPerSecond sets AnalogOutputEngineeringUnits to Grams-per-second (0x009A)
func (a *AnalogOutputEngineeringUnits) SetGramsPerSecond() { *a = 0x009A }

// IsGramsPerMinute checks if AnalogOutputEngineeringUnits equals the value for Grams-per-minute (0x009B)
func (a AnalogOutputEngineeringUnits) IsGramsPerMinute() bool { return a == 0x009B }

// SetGramsPerMinute sets AnalogOutputEngineeringUnits to Grams-per-minute (0x009B)
func (a *AnalogOutputEngineeringUnits) SetGramsPerMinute() { *a = 0x009B }

// IsTonsPerHour checks if AnalogOutputEngineeringUnits equals the value for Tons-per-hour (0x009C)
func (a AnalogOutputEngineeringUnits) IsTonsPerHour() bool { return a == 0x009C }

// SetTonsPerHour sets AnalogOutputEngineeringUnits to Tons-per-hour (0x009C)
func (a *AnalogOutputEngineeringUnits) SetTonsPerHour() { *a = 0x009C }

// IsKiloBtusPerHour checks if AnalogOutputEngineeringUnits equals the value for Kilo-BTUs-per-hour (0x009D)
func (a AnalogOutputEngineeringUnits) IsKiloBtusPerHour() bool { return a == 0x009D }

// SetKiloBtusPerHour sets AnalogOutputEngineeringUnits to Kilo-BTUs-per-hour (0x009D)
func (a *AnalogOutputEngineeringUnits) SetKiloBtusPerHour() { *a = 0x009D }

// IsHundredthsSeconds checks if AnalogOutputEngineeringUnits equals the value for Hundredths-seconds (0x009E)
func (a AnalogOutputEngineeringUnits) IsHundredthsSeconds() bool { return a == 0x009E }

// SetHundredthsSeconds sets AnalogOutputEngineeringUnits to Hundredths-seconds (0x009E)
func (a *AnalogOutputEngineeringUnits) SetHundredthsSeconds() { *a = 0x009E }

// IsMilliseconds checks if AnalogOutputEngineeringUnits equals the value for Milliseconds (0x009F)
func (a AnalogOutputEngineeringUnits) IsMilliseconds() bool { return a == 0x009F }

// SetMilliseconds sets AnalogOutputEngineeringUnits to Milliseconds (0x009F)
func (a *AnalogOutputEngineeringUnits) SetMilliseconds() { *a = 0x009F }

// IsNewtonMeters checks if AnalogOutputEngineeringUnits equals the value for Newton-meters (0x00A0)
func (a AnalogOutputEngineeringUnits) IsNewtonMeters() bool { return a == 0x00A0 }

// SetNewtonMeters sets AnalogOutputEngineeringUnits to Newton-meters (0x00A0)
func (a *AnalogOutputEngineeringUnits) SetNewtonMeters() { *a = 0x00A0 }

// IsMillimetersPerSecond checks if AnalogOutputEngineeringUnits equals the value for Millimeters-per-second (0x00A1)
func (a AnalogOutputEngineeringUnits) IsMillimetersPerSecond() bool { return a == 0x00A1 }

// SetMillimetersPerSecond sets AnalogOutputEngineeringUnits to Millimeters-per-second (0x00A1)
func (a *AnalogOutputEngineeringUnits) SetMillimetersPerSecond() { *a = 0x00A1 }

// IsMillimetersPerMinute checks if AnalogOutputEngineeringUnits equals the value for Millimeters-per-minute (0x00A2)
func (a AnalogOutputEngineeringUnits) IsMillimetersPerMinute() bool { return a == 0x00A2 }

// SetMillimetersPerMinute sets AnalogOutputEngineeringUnits to Millimeters-per-minute (0x00A2)
func (a *AnalogOutputEngineeringUnits) SetMillimetersPerMinute() { *a = 0x00A2 }

// IsMetersPerMinute checks if AnalogOutputEngineeringUnits equals the value for Meters-per-minute (0x00A3)
func (a AnalogOutputEngineeringUnits) IsMetersPerMinute() bool { return a == 0x00A3 }

// SetMetersPerMinute sets AnalogOutputEngineeringUnits to Meters-per-minute (0x00A3)
func (a *AnalogOutputEngineeringUnits) SetMetersPerMinute() { *a = 0x00A3 }

// IsMetersPerHour checks if AnalogOutputEngineeringUnits equals the value for Meters-per-hour (0x00A4)
func (a AnalogOutputEngineeringUnits) IsMetersPerHour() bool { return a == 0x00A4 }

// SetMetersPerHour sets AnalogOutputEngineeringUnits to Meters-per-hour (0x00A4)
func (a *AnalogOutputEngineeringUnits) SetMetersPerHour() { *a = 0x00A4 }

// IsCubicMetersPerMinute checks if AnalogOutputEngineeringUnits equals the value for Cubic-meters-per-minute (0x00A5)
func (a AnalogOutputEngineeringUnits) IsCubicMetersPerMinute() bool { return a == 0x00A5 }

// SetCubicMetersPerMinute sets AnalogOutputEngineeringUnits to Cubic-meters-per-minute (0x00A5)
func (a *AnalogOutputEngineeringUnits) SetCubicMetersPerMinute() { *a = 0x00A5 }

// IsMetersPerSecondPerSecond checks if AnalogOutputEngineeringUnits equals the value for Meters-per-second-per-second (0x00A6)
func (a AnalogOutputEngineeringUnits) IsMetersPerSecondPerSecond() bool { return a == 0x00A6 }

// SetMetersPerSecondPerSecond sets AnalogOutputEngineeringUnits to Meters-per-second-per-second (0x00A6)
func (a *AnalogOutputEngineeringUnits) SetMetersPerSecondPerSecond() { *a = 0x00A6 }

// IsAmperesPerMeter checks if AnalogOutputEngineeringUnits equals the value for Amperes-per-meter (0x00A7)
func (a AnalogOutputEngineeringUnits) IsAmperesPerMeter() bool { return a == 0x00A7 }

// SetAmperesPerMeter sets AnalogOutputEngineeringUnits to Amperes-per-meter (0x00A7)
func (a *AnalogOutputEngineeringUnits) SetAmperesPerMeter() { *a = 0x00A7 }

// IsAmperesPerSquareMeter checks if AnalogOutputEngineeringUnits equals the value for Amperes-per-square-meter (0x00A8)
func (a AnalogOutputEngineeringUnits) IsAmperesPerSquareMeter() bool { return a == 0x00A8 }

// SetAmperesPerSquareMeter sets AnalogOutputEngineeringUnits to Amperes-per-square-meter (0x00A8)
func (a *AnalogOutputEngineeringUnits) SetAmperesPerSquareMeter() { *a = 0x00A8 }

// IsAmpereSquareMeters checks if AnalogOutputEngineeringUnits equals the value for Ampere-square-meters (0x00A9)
func (a AnalogOutputEngineeringUnits) IsAmpereSquareMeters() bool { return a == 0x00A9 }

// SetAmpereSquareMeters sets AnalogOutputEngineeringUnits to Ampere-square-meters (0x00A9)
func (a *AnalogOutputEngineeringUnits) SetAmpereSquareMeters() { *a = 0x00A9 }

// IsFarads checks if AnalogOutputEngineeringUnits equals the value for Farads (0x00AA)
func (a AnalogOutputEngineeringUnits) IsFarads() bool { return a == 0x00AA }

// SetFarads sets AnalogOutputEngineeringUnits to Farads (0x00AA)
func (a *AnalogOutputEngineeringUnits) SetFarads() { *a = 0x00AA }

// IsHenrys checks if AnalogOutputEngineeringUnits equals the value for Henrys (0x00AB)
func (a AnalogOutputEngineeringUnits) IsHenrys() bool { return a == 0x00AB }

// SetHenrys sets AnalogOutputEngineeringUnits to Henrys (0x00AB)
func (a *AnalogOutputEngineeringUnits) SetHenrys() { *a = 0x00AB }

// IsOhmMeters checks if AnalogOutputEngineeringUnits equals the value for Ohm-meters (0x00AC)
func (a AnalogOutputEngineeringUnits) IsOhmMeters() bool { return a == 0x00AC }

// SetOhmMeters sets AnalogOutputEngineeringUnits to Ohm-meters (0x00AC)
func (a *AnalogOutputEngineeringUnits) SetOhmMeters() { *a = 0x00AC }

// IsSiemens checks if AnalogOutputEngineeringUnits equals the value for Siemens (0x00AD)
func (a AnalogOutputEngineeringUnits) IsSiemens() bool { return a == 0x00AD }

// SetSiemens sets AnalogOutputEngineeringUnits to Siemens (0x00AD)
func (a *AnalogOutputEngineeringUnits) SetSiemens() { *a = 0x00AD }

// IsSiemensPerMeter checks if AnalogOutputEngineeringUnits equals the value for Siemens-per-meter (0x00AE)
func (a AnalogOutputEngineeringUnits) IsSiemensPerMeter() bool { return a == 0x00AE }

// SetSiemensPerMeter sets AnalogOutputEngineeringUnits to Siemens-per-meter (0x00AE)
func (a *AnalogOutputEngineeringUnits) SetSiemensPerMeter() { *a = 0x00AE }

// IsTeslas checks if AnalogOutputEngineeringUnits equals the value for Teslas (0x00AF)
func (a AnalogOutputEngineeringUnits) IsTeslas() bool { return a == 0x00AF }

// SetTeslas sets AnalogOutputEngineeringUnits to Teslas (0x00AF)
func (a *AnalogOutputEngineeringUnits) SetTeslas() { *a = 0x00AF }

// IsVoltsPerDegreeKelvin checks if AnalogOutputEngineeringUnits equals the value for Volts-per-degree-Kelvin (0x00B0)
func (a AnalogOutputEngineeringUnits) IsVoltsPerDegreeKelvin() bool { return a == 0x00B0 }

// SetVoltsPerDegreeKelvin sets AnalogOutputEngineeringUnits to Volts-per-degree-Kelvin (0x00B0)
func (a *AnalogOutputEngineeringUnits) SetVoltsPerDegreeKelvin() { *a = 0x00B0 }

// IsVoltsPerMeter checks if AnalogOutputEngineeringUnits equals the value for Volts-per-meter (0x00B1)
func (a AnalogOutputEngineeringUnits) IsVoltsPerMeter() bool { return a == 0x00B1 }

// SetVoltsPerMeter sets AnalogOutputEngineeringUnits to Volts-per-meter (0x00B1)
func (a *AnalogOutputEngineeringUnits) SetVoltsPerMeter() { *a = 0x00B1 }

// IsWebers checks if AnalogOutputEngineeringUnits equals the value for Webers (0x00B2)
func (a AnalogOutputEngineeringUnits) IsWebers() bool { return a == 0x00B2 }

// SetWebers sets AnalogOutputEngineeringUnits to Webers (0x00B2)
func (a *AnalogOutputEngineeringUnits) SetWebers() { *a = 0x00B2 }

// IsCandelas checks if AnalogOutputEngineeringUnits equals the value for Candelas (0x00B3)
func (a AnalogOutputEngineeringUnits) IsCandelas() bool { return a == 0x00B3 }

// SetCandelas sets AnalogOutputEngineeringUnits to Candelas (0x00B3)
func (a *AnalogOutputEngineeringUnits) SetCandelas() { *a = 0x00B3 }

// IsCandelasPerSquareMeter checks if AnalogOutputEngineeringUnits equals the value for Candelas-per-square-meter (0x00B4)
func (a AnalogOutputEngineeringUnits) IsCandelasPerSquareMeter() bool { return a == 0x00B4 }

// SetCandelasPerSquareMeter sets AnalogOutputEngineeringUnits to Candelas-per-square-meter (0x00B4)
func (a *AnalogOutputEngineeringUnits) SetCandelasPerSquareMeter() { *a = 0x00B4 }

// IsKelvinsPerHour checks if AnalogOutputEngineeringUnits equals the value for Kelvins-per-hour (0x00B5)
func (a AnalogOutputEngineeringUnits) IsKelvinsPerHour() bool { return a == 0x00B5 }

// SetKelvinsPerHour sets AnalogOutputEngineeringUnits to Kelvins-per-hour (0x00B5)
func (a *AnalogOutputEngineeringUnits) SetKelvinsPerHour() { *a = 0x00B5 }

// IsKelvinsPerMinute checks if AnalogOutputEngineeringUnits equals the value for Kelvins-per-minute (0x00B6)
func (a AnalogOutputEngineeringUnits) IsKelvinsPerMinute() bool { return a == 0x00B6 }

// SetKelvinsPerMinute sets AnalogOutputEngineeringUnits to Kelvins-per-minute (0x00B6)
func (a *AnalogOutputEngineeringUnits) SetKelvinsPerMinute() { *a = 0x00B6 }

// IsJouleSeconds checks if AnalogOutputEngineeringUnits equals the value for Joule-seconds (0x00B7)
func (a AnalogOutputEngineeringUnits) IsJouleSeconds() bool { return a == 0x00B7 }

// SetJouleSeconds sets AnalogOutputEngineeringUnits to Joule-seconds (0x00B7)
func (a *AnalogOutputEngineeringUnits) SetJouleSeconds() { *a = 0x00B7 }

// IsSquareMetersPerNewton checks if AnalogOutputEngineeringUnits equals the value for Square-meters-per-Newton (0x00B8)
func (a AnalogOutputEngineeringUnits) IsSquareMetersPerNewton() bool { return a == 0x00B8 }

// SetSquareMetersPerNewton sets AnalogOutputEngineeringUnits to Square-meters-per-Newton (0x00B8)
func (a *AnalogOutputEngineeringUnits) SetSquareMetersPerNewton() { *a = 0x00B8 }

// IsKilogramPerCubicMeter checks if AnalogOutputEngineeringUnits equals the value for Kilogram-per-cubic-meter (0x00B9)
func (a AnalogOutputEngineeringUnits) IsKilogramPerCubicMeter() bool { return a == 0x00B9 }

// SetKilogramPerCubicMeter sets AnalogOutputEngineeringUnits to Kilogram-per-cubic-meter (0x00B9)
func (a *AnalogOutputEngineeringUnits) SetKilogramPerCubicMeter() { *a = 0x00B9 }

// IsNewtonSeconds checks if AnalogOutputEngineeringUnits equals the value for Newton-seconds (0x00BA)
func (a AnalogOutputEngineeringUnits) IsNewtonSeconds() bool { return a == 0x00BA }

// SetNewtonSeconds sets AnalogOutputEngineeringUnits to Newton-seconds (0x00BA)
func (a *AnalogOutputEngineeringUnits) SetNewtonSeconds() { *a = 0x00BA }

// IsNewtonsPerMeter checks if AnalogOutputEngineeringUnits equals the value for Newtons-per-meter (0x00BB)
func (a AnalogOutputEngineeringUnits) IsNewtonsPerMeter() bool { return a == 0x00BB }

// SetNewtonsPerMeter sets AnalogOutputEngineeringUnits to Newtons-per-meter (0x00BB)
func (a *AnalogOutputEngineeringUnits) SetNewtonsPerMeter() { *a = 0x00BB }

// IsWattsPerMeterPerDegreeKelvin checks if AnalogOutputEngineeringUnits equals the value for Watts-per-meter-per-degree-Kelvin (0x00BC)
func (a AnalogOutputEngineeringUnits) IsWattsPerMeterPerDegreeKelvin() bool { return a == 0x00BC }

// SetWattsPerMeterPerDegreeKelvin sets AnalogOutputEngineeringUnits to Watts-per-meter-per-degree-Kelvin (0x00BC)
func (a *AnalogOutputEngineeringUnits) SetWattsPerMeterPerDegreeKelvin() { *a = 0x00BC }

// IsOther checks if AnalogOutputEngineeringUnits equals the value for Other (0x00FF)
func (a AnalogOutputEngineeringUnits) IsOther() bool { return a == 0x00FF }

// SetOther sets AnalogOutputEngineeringUnits to Other (0x00FF)
func (a *AnalogOutputEngineeringUnits) SetOther() { *a = 0x00FF }

type AnalogOutputApplicationType zcl.Zu32

func (a AnalogOutputApplicationType) ID() zcl.AttrID                       { return 256 }
func (a AnalogOutputApplicationType) Cluster() zcl.ClusterID               { return AnalogOutputBasicCluster }
func (a *AnalogOutputApplicationType) Value() *AnalogOutputApplicationType { return a }
func (a AnalogOutputApplicationType) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *AnalogOutputApplicationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputApplicationType(*nt)
	return br, err
}

func (a AnalogOutputApplicationType) Readable() bool   { return true }
func (a AnalogOutputApplicationType) Writable() bool   { return false }
func (a AnalogOutputApplicationType) Reportable() bool { return false }
func (a AnalogOutputApplicationType) SceneIndex() int  { return -1 }

func (a AnalogOutputApplicationType) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}

type AnalogOutputXiaomi0Xf000 zcl.Zu32

func (a AnalogOutputXiaomi0Xf000) ID() zcl.AttrID                    { return 61440 }
func (a AnalogOutputXiaomi0Xf000) Cluster() zcl.ClusterID            { return AnalogOutputBasicCluster }
func (a *AnalogOutputXiaomi0Xf000) Value() *AnalogOutputXiaomi0Xf000 { return a }
func (a AnalogOutputXiaomi0Xf000) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *AnalogOutputXiaomi0Xf000) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputXiaomi0Xf000(*nt)
	return br, err
}

func (a AnalogOutputXiaomi0Xf000) Readable() bool   { return true }
func (a AnalogOutputXiaomi0Xf000) Writable() bool   { return false }
func (a AnalogOutputXiaomi0Xf000) Reportable() bool { return false }
func (a AnalogOutputXiaomi0Xf000) SceneIndex() int  { return -1 }

func (a AnalogOutputXiaomi0Xf000) String() string {

	return zcl.Sprintf("%s", zcl.Zu32(a))
}
