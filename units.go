package zcl

import "fmt"

type Unit uint8

type EngineeringUnit uint16

func (u *EngineeringUnit) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(2, buf)
	*u = EngineeringUnit(val)
	return buf, err
}
func (u EngineeringUnit) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(2, uint64(u)) }
func (u *EngineeringUnit) Values() []Val              { return []Val{u} }
func (u EngineeringUnit) Unit() Unit                  { return Unit(u) }
func (u EngineeringUnit) String() string              { return u.Unit().String() }
func (u EngineeringUnit) Format(v float64) string     { return u.Unit().Format(v) }

const (
	SquareMeters                    Unit = 0x00 //  Area
	SquareFeet                      Unit = 0x01 //  Area
	Milliamperes                    Unit = 0x02 //  Electrical
	Amperes                         Unit = 0x03 //  Electrical
	Ohms                            Unit = 0x04 //  Electrical
	Volts                           Unit = 0x05 //  Electrical
	KiloVolts                       Unit = 0x06 //  Electrical
	MegaVolts                       Unit = 0x07 //  Electrical
	VoltAmperes                     Unit = 0x08 //  Electrical
	KiloVoltAmperes                 Unit = 0x09 //  Electrical
	MegaVoltAmperes                 Unit = 0x0A //  Electrical
	VoltAmperesReactive             Unit = 0x0B //  Electrical
	KiloVoltAmperesReactive         Unit = 0x0C //  Electrical
	MegaVoltAmperesReactive         Unit = 0x0D //  Electrical
	DegreesPhase                    Unit = 0x0E //  Electrical
	PowerFactor                     Unit = 0x0F //  Electrical
	Joules                          Unit = 0x10 //  Energy
	Kilojoules                      Unit = 0x11 //  Energy
	WattHours                       Unit = 0x12 //  Energy
	KilowattHours                   Unit = 0x13 //  Energy
	BTUs                            Unit = 0x14 //  Energy
	Therms                          Unit = 0x15 //  Energy
	TonHours                        Unit = 0x16 //  Energy
	JoulesPerKilogramDryAir         Unit = 0x17 //  Enthalpy
	BTUsPerPoundDryAir              Unit = 0x18 //  Enthalpy
	CyclesPerHour                   Unit = 0x19 //  Frequency
	CyclesPerMinute                 Unit = 0x1A //  Frequency
	Hertz                           Unit = 0x1B //  Frequency
	GramsOfWaterPerKilogramDryAir   Unit = 0x1C //  Humidity
	PercentRelativeHumidity         Unit = 0x1D //  Humidity
	Millimeters                     Unit = 0x1E //  Length
	Meters                          Unit = 0x1F //  Length
	Inches                          Unit = 0x20 //  Length
	Feet                            Unit = 0x21 //  Length
	WattsPerSquareFoot              Unit = 0x22 //  Light
	WattsPerSquareMeter             Unit = 0x23 //  Light
	Lumens                          Unit = 0x24 //  Light
	Luxes                           Unit = 0x25 //  Light
	FootCandles                     Unit = 0x26 //  Light
	Kilograms                       Unit = 0x27 //  Mass
	PoundsMass                      Unit = 0x28 //  Mass
	Tons                            Unit = 0x29 //  Mass
	KilogramsPerSecond              Unit = 0x2A //  Mass Flow
	KilogramsPerMinute              Unit = 0x2B //  Mass Flow
	KilogramsPerHour                Unit = 0x2C //  Mass Flow
	PoundsMassPerMinute             Unit = 0x2D //  Mass Flow
	PoundsMassPerHour               Unit = 0x2E //  Mass Flow
	Watts                           Unit = 0x2F //  Power
	Kilowatts                       Unit = 0x30 //  Power
	Megawatts                       Unit = 0x31 //  Power
	BTUsPerHour                     Unit = 0x32 //  Power
	Horsepower                      Unit = 0x33 //  Power
	TonsRefrigeration               Unit = 0x34 //  Power
	Pascals                         Unit = 0x35 //  Pressure
	Kilopascals                     Unit = 0x36 //  Pressure
	Bars                            Unit = 0x37 //  Pressure
	PoundsForcePerSquareInch        Unit = 0x38 //  Pressure
	CentimetersOfWater              Unit = 0x39 //  Pressure
	InchesOfWater                   Unit = 0x3A //  Pressure
	MillimetersOfMercury            Unit = 0x3B //  Pressure
	CentimetersOfMercury            Unit = 0x3C //  Pressure
	InchesOfMercury                 Unit = 0x3D //  Pressure
	DegreesCelsius                  Unit = 0x3E //  Temperature
	DegreesKelvin                   Unit = 0x3F //  Temperature
	DegreesFahrenheit               Unit = 0x40 //  Temperature
	DegreeDaysCelsius               Unit = 0x41 //  Temperature
	DegreeDaysFahrenheit            Unit = 0x42 //  Temperature
	Years                           Unit = 0x43 //  Time
	Months                          Unit = 0x44 //  Time
	Weeks                           Unit = 0x45 //  Time
	Days                            Unit = 0x46 //  Time
	Hours                           Unit = 0x47 //  Time
	Minutes                         Unit = 0x48 //  Time
	Seconds                         Unit = 0x49 //  Time
	MetersPerSecond                 Unit = 0x4A //  Velocity
	KilometersPerHour               Unit = 0x4B //  Velocity
	FeetPerSecond                   Unit = 0x4C //  Velocity
	FeetPerMinute                   Unit = 0x4D //  Velocity
	MilesPerHour                    Unit = 0x4E //  Velocity
	CubicFeet                       Unit = 0x4F //  Volume
	CubicMeters                     Unit = 0x50 //  Volume
	ImperialGallons                 Unit = 0x51 //  Volume
	Liters                          Unit = 0x52 //  Volume
	UsGallons                       Unit = 0x53 //  Volume
	CubicFeetPerMinute              Unit = 0x54 //  Volumetric Flow
	CubicMetersPerSecond            Unit = 0x55 //  Volumetric Flow
	ImperialGallonsPerMinute        Unit = 0x56 //  Volumetric Flow
	LitersPerSecond                 Unit = 0x57 //  Volumetric Flow
	LitersPerMinute                 Unit = 0x58 //  Volumetric Flow
	UsGallonsPerMinute              Unit = 0x59 //  Volumetric Flow
	DegreesAngular                  Unit = 0x5A //  Other
	DegreesCelsiusPerHour           Unit = 0x5B //  Other
	DegreesCelsiusPerMinute         Unit = 0x5C //  Other
	DegreesFahrenheitPerHour        Unit = 0x5D //  Other
	DegreesFahrenheitPerMinute      Unit = 0x5E //  Other
	NoUnits                         Unit = 0x5F //  Other
	PartsPerMillion                 Unit = 0x60 //  Other
	PartsPerBillion                 Unit = 0x61 //  Other
	Percent                         Unit = 0x62 //  Other
	PercentPerSecond                Unit = 0x63 //  Other
	PerMinute                       Unit = 0x64 //  Other
	PerSecond                       Unit = 0x65 //  Other
	PsiPerDegreeFahrenheit          Unit = 0x66 //  Other
	Radians                         Unit = 0x67 //  Other
	RevolutionsPerMinute            Unit = 0x68 //  Other
	Currency1                       Unit = 0x69 //  Currency
	Currency2                       Unit = 0x6A //  Currency
	Currency3                       Unit = 0x6B //  Currency
	Currency4                       Unit = 0x6C //  Currency
	Currency5                       Unit = 0x6D //  Currency
	Currency6                       Unit = 0x6E //  Currency
	Currency7                       Unit = 0x6F //  Currency
	Currency8                       Unit = 0x70 //  Currency
	Currency9                       Unit = 0x71 //  Currency
	Currency10                      Unit = 0x72 //  Currency
	SquareInches                    Unit = 0x73 //  Area
	SquareCentimeters               Unit = 0x74 //  Area
	BTUsPerPound                    Unit = 0x75 //  Enthalpy
	Centimeters                     Unit = 0x76 //  Length
	PoundsMassPerSecond             Unit = 0x77 //  Mass Flow
	DeltaDegreesFahrenheit          Unit = 0x78 //  Temperature
	DeltaDegreesKelvin              Unit = 0x79 //  Temperature
	Kilohms                         Unit = 0x7A //  Electrical
	Megohms                         Unit = 0x7B //  Electrical
	Millivolts                      Unit = 0x7C //  Electrical
	KilojoulesPerKilogram           Unit = 0x7D //  Energy
	Megajoules                      Unit = 0x7E //  Energy
	JoulesPerDegreeKelvin           Unit = 0x7F //  Entrophy
	JoulesPerKilogramDegreeKelvin   Unit = 0x80 //  Entrophy
	Kilohertz                       Unit = 0x81 //  Frequency
	Megahertz                       Unit = 0x82 //  Frequency
	PerHour                         Unit = 0x83 //  Frequency
	Milliwatts                      Unit = 0x84 //  Power
	Hectopascals                    Unit = 0x85 //  Pressure
	Millibars                       Unit = 0x86 //  Pressure
	CubicMetersPerHour              Unit = 0x87 //  Volumetric Flow
	LitersPerHour                   Unit = 0x88 //  Volumetric Flow
	KilowattHoursPerSquareMeter     Unit = 0x89 //  Other
	KilowattHoursPerSquareFoot      Unit = 0x8A //  Other
	MegajoulesPerSquareMeter        Unit = 0x8B //  Other
	MegajoulesPerSquareFoot         Unit = 0x8C //  Other
	WattsPerSquareMeterDegreeKelvin Unit = 0x8D //  Other
	CubicFeetPerSecond              Unit = 0x8E //  Volumetric Flow
	PercentObscurationPerFoot       Unit = 0x8F //  Other
	PercentObscurationPerMeter      Unit = 0x90 //  Other
	Milliohms                       Unit = 0x91 //  Electrical
	MegawattHours                   Unit = 0x92 //  Energy
	KiloBTUs                        Unit = 0x93 //  Energy
	MegaBTUs                        Unit = 0x94 //  Energy
	KilojoulesPerKilogramDryAir     Unit = 0x95 //  Enthalpy
	MegajoulesPerKilogramDryAir     Unit = 0x96 //  Enthalpy
	KilojoulesPerDegreeKelvin       Unit = 0x97 //  Entrophy
	MegajoulesPerDegreeKelvin       Unit = 0x98 //  Entrophy
	Newton                          Unit = 0x99 //  Force
	GramsPerSecond                  Unit = 0x9A //  Mass Flow
	GramsPerMinute                  Unit = 0x9B //  Mass Flow
	TonsPerHour                     Unit = 0x9C //  Mass Flow
	KiloBTUsPerHour                 Unit = 0x9D //  Power
	HundredthsSeconds               Unit = 0x9E //  Time
	Milliseconds                    Unit = 0x9F //  Time
	NewtonMeters                    Unit = 0xA0 //  Torque
	MillimetersPerSecond            Unit = 0xA1 //  Velocity
	MillimetersPerMinute            Unit = 0xA2 //  Velocity
	MetersPerMinute                 Unit = 0xA3 //  Velocity
	MetersPerHour                   Unit = 0xA4 //  Velocity
	CubicMetersPerMinute            Unit = 0xA5 //  Volumetric Flow
	MetersPerSecondPerSecond        Unit = 0xA6 //  Acceleration
	AmperesPerMeter                 Unit = 0xA7 //  Electrical
	AmperesPerSquareMeter           Unit = 0xA8 //  Electrical
	AmpereSquareMeters              Unit = 0xA9 //  Electrical
	Farads                          Unit = 0xAA //  Electrical
	Henrys                          Unit = 0xAB //  Electrical
	OhmMeters                       Unit = 0xAC //  Electrical
	Siemens                         Unit = 0xAD //  Electrical
	SiemensPerMeter                 Unit = 0xAE //  Electrical
	Teslas                          Unit = 0xAF //  Electrical
	VoltsPerDegreeKelvin            Unit = 0xB0 //  Electrical
	VoltsPerMeter                   Unit = 0xB1 //  Electrical
	Webers                          Unit = 0xB2 //  Electrical
	Candelas                        Unit = 0xB3 //  Light
	CandelasPerSquareMeter          Unit = 0xB4 //  Light
	KelvinsPerHour                  Unit = 0xB5 //  Temperature
	KelvinsPerMinute                Unit = 0xB6 //  Temperature
	JouleSeconds                    Unit = 0xB7 //  Other
	SquareMetersPerNewton           Unit = 0xB8 //  Other
	KilogramPerCubicMeter           Unit = 0xB9 //  Other
	NewtonSeconds                   Unit = 0xBA //  Other
	NewtonsPerMeter                 Unit = 0xBB //  Other
	WattsPerMeterPerDegreeKelvin    Unit = 0xBC //  Other
	Other                           Unit = 0xFF
)

func (u Unit) String() string {
	switch u {
	case SquareMeters:
		return "m²"
	case SquareFeet:
		return "ft²"
	case Milliamperes:
		return "mA"
	case Amperes:
		return "A"
	case Ohms:
		return "Ω"
	case Volts:
		return "V"
	case KiloVolts:
		return "kV"
	case MegaVolts:
		return "MV"
	case VoltAmperes:
		return "VA"
	case KiloVoltAmperes:
		return "kVA"
	case MegaVoltAmperes:
		return "MVA"
	case VoltAmperesReactive:
		return "var"
	case KiloVoltAmperesReactive:
		return "kvar"
	case MegaVoltAmperesReactive:
		return "Mvar"
	case DegreesPhase:
		return "°"
	case PowerFactor:
		return "PFC"
	case Joules:
		return "J"
	case Kilojoules:
		return "kJ"
	case WattHours:
		return "Wh"
	case KilowattHours:
		return "kWh"
	case BTUs:
		return "BTU"
	case Therms:
		return "thm"
	case TonHours:
		return "Th"
	case JoulesPerKilogramDryAir:
		return "J/kg dry air"
	case BTUsPerPoundDryAir:
		return "BTU/lbs dry air"
	case CyclesPerHour:
		return "cycles/hr"
	case CyclesPerMinute:
		return "cycles/min"
	case Hertz:
		return "Hz"
	case GramsOfWaterPerKilogramDryAir:
		return "g water/kg dry air"
	case PercentRelativeHumidity:
		return "%% RH"
	case Millimeters:
		return "mm"
	case Meters:
		return "m"
	case Inches:
		return "\""
	case Feet:
		return "'"
	case WattsPerSquareFoot:
		return "W/ft²"
	case WattsPerSquareMeter:
		return "W/m²"
	case Lumens:
		return "lm"
	case Luxes:
		return "lux"
	case FootCandles:
		return "lm/ft²"
	case Kilograms:
		return "kg"
	case PoundsMass:
		return "lbm"
	case Tons:
		return "ton"
	case KilogramsPerSecond:
		return "kg/s"
	case KilogramsPerMinute:
		return "kg/min"
	case KilogramsPerHour:
		return "kg/h"
	case PoundsMassPerMinute:
		return "lbm/min"
	case PoundsMassPerHour:
		return "lbm/h"
	case Watts:
		return "W"
	case Kilowatts:
		return "kW"
	case Megawatts:
		return "MW"
	case BTUsPerHour:
		return "BTU/h"
	case Horsepower:
		return "hp"
	case TonsRefrigeration:
		return "TR"
	case Pascals:
		return "Pa"
	case Kilopascals:
		return "kPa"
	case Bars:
		return "bar"
	case PoundsForcePerSquareInch:
		return "lbf/in²"
	case CentimetersOfWater:
		return "cm h2o"
	case InchesOfWater:
		return "in h2o"
	case MillimetersOfMercury:
		return "mmHg"
	case CentimetersOfMercury:
		return "cmHg"
	case InchesOfMercury:
		return "inHg"
	case DegreesCelsius:
		return "°C"
	case DegreesKelvin:
		return "K"
	case DegreesFahrenheit:
		return "°F"
	case DegreeDaysCelsius:
		return "°C DD"
	case DegreeDaysFahrenheit:
		return "°F DD"
	case Years:
		return "year%s"
	case Months:
		return "month%s"
	case Weeks:
		return "week%s"
	case Days:
		return "day%s"
	case Hours:
		return "hour%s"
	case Minutes:
		return "minute%s"
	case Seconds:
		return "second%s"
	case MetersPerSecond:
		return "m/s"
	case KilometersPerHour:
		return "km/h"
	case FeetPerSecond:
		return "ft/s"
	case FeetPerMinute:
		return "ft/min"
	case MilesPerHour:
		return "mph"
	case CubicFeet:
		return "ft³"
	case CubicMeters:
		return "m³"
	case ImperialGallons:
		return "gal (Imp.)"
	case Liters:
		return "L"
	case UsGallons:
		return "gal (US)"
	case CubicFeetPerMinute:
		return "ft³/min"
	case CubicMetersPerSecond:
		return "m³/s"
	case ImperialGallonsPerMinute:
		return "gal/min (Imp.)"
	case LitersPerSecond:
		return "L/s"
	case LitersPerMinute:
		return "L/min"
	case UsGallonsPerMinute:
		return "gal/min (US)"
	case DegreesAngular:
		return "° angle"
	case DegreesCelsiusPerHour:
		return "°C/h"
	case DegreesCelsiusPerMinute:
		return "°C/min"
	case DegreesFahrenheitPerHour:
		return "°F/h"
	case DegreesFahrenheitPerMinute:
		return "°F/min"
	case NoUnits:
		return ""
	case PartsPerMillion:
		return "ppm"
	case PartsPerBillion:
		return "ppb"
	case Percent:
		return "%%"
	case PercentPerSecond:
		return "%%/s"
	case PerMinute:
		return "/min"
	case PerSecond:
		return "/s"
	case PsiPerDegreeFahrenheit:
		return "psi/°F"
	case Radians:
		return "rad"
	case RevolutionsPerMinute:
		return "RPM"
	case Currency1:
		return "Currency1"
	case Currency2:
		return "Currency2"
	case Currency3:
		return "Currency3"
	case Currency4:
		return "Currency4"
	case Currency5:
		return "Currency5"
	case Currency6:
		return "Currency6"
	case Currency7:
		return "Currency7"
	case Currency8:
		return "Currency8"
	case Currency9:
		return "Currency9"
	case Currency10:
		return "Currency10"
	case SquareInches:
		return "in²"
	case SquareCentimeters:
		return "cm²"
	case BTUsPerPound:
		return "BTU/lbs"
	case Centimeters:
		return "cm"
	case PoundsMassPerSecond:
		return "lbm/s"
	case DeltaDegreesFahrenheit:
		return "delta °F"
	case DeltaDegreesKelvin:
		return "delta-K"
	case Kilohms:
		return "kΩ"
	case Megohms:
		return "MΩ"
	case Millivolts:
		return "mV"
	case KilojoulesPerKilogram:
		return "kJ/kg"
	case Megajoules:
		return "MJ"
	case JoulesPerDegreeKelvin:
		return "J/°K"
	case JoulesPerKilogramDegreeKelvin:
		return "J/kg-K"
	case Kilohertz:
		return "kHz"
	case Megahertz:
		return "MHz"
	case PerHour:
		return "/h"
	case Milliwatts:
		return "mW"
	case Hectopascals:
		return "hpa"
	case Millibars:
		return "mbar"
	case CubicMetersPerHour:
		return "m³/h"
	case LitersPerHour:
		return "L/h"
	case KilowattHoursPerSquareMeter:
		return "kWh/m²"
	case KilowattHoursPerSquareFoot:
		return "kWh/ft²"
	case MegajoulesPerSquareMeter:
		return "MJ/m²"
	case MegajoulesPerSquareFoot:
		return "MJ/ft²"
	case WattsPerSquareMeterDegreeKelvin:
		return "W/m²-K"
	case CubicFeetPerSecond:
		return "ft³/s"
	case PercentObscurationPerFoot:
		return "%% obscuration/ft"
	case PercentObscurationPerMeter:
		return "%% obscuration/m"
	case Milliohms:
		return "mΩ"
	case MegawattHours:
		return "MWh"
	case KiloBTUs:
		return "kBTU"
	case MegaBTUs:
		return "MBTU"
	case KilojoulesPerKilogramDryAir:
		return "KJ/kg dry air"
	case MegajoulesPerKilogramDryAir:
		return "MJ/kg dry air"
	case KilojoulesPerDegreeKelvin:
		return "KJ/deg-K"
	case MegajoulesPerDegreeKelvin:
		return "MJ/deg-K"
	case Newton:
		return "N"
	case GramsPerSecond:
		return "g/s"
	case GramsPerMinute:
		return "g/min"
	case TonsPerHour:
		return "tons/h"
	case KiloBTUsPerHour:
		return "kBTU/h"
	case HundredthsSeconds:
		return "ms"
	case Milliseconds:
		return "ms"
	case NewtonMeters:
		return "Nm"
	case MillimetersPerSecond:
		return "mm/s"
	case MillimetersPerMinute:
		return "mm/min"
	case MetersPerMinute:
		return "m/min"
	case MetersPerHour:
		return "m/h"
	case CubicMetersPerMinute:
		return "m³/min"
	case MetersPerSecondPerSecond:
		return "m/s²"
	case AmperesPerMeter:
		return "A/m"
	case AmperesPerSquareMeter:
		return "A/m²"
	case AmpereSquareMeters:
		return "Am²"
	case Farads:
		return "Farad%s"
	case Henrys:
		return "Henry%s"
	case OhmMeters:
		return "Ωm"
	case Siemens:
		return "Siemens"
	case SiemensPerMeter:
		return "Siemens/m"
	case Teslas:
		return "Tesla%s"
	case VoltsPerDegreeKelvin:
		return "V/deg-K"
	case VoltsPerMeter:
		return "V/m"
	case Webers:
		return "Weber%s"
	case Candelas:
		return "Candela%s"
	case CandelasPerSquareMeter:
		return "Candela/m²"
	case KelvinsPerHour:
		return "K/h"
	case KelvinsPerMinute:
		return "K/min"
	case JouleSeconds:
		return "Js"
	case SquareMetersPerNewton:
		return "m²/N"
	case KilogramPerCubicMeter:
		return "kg/m³"
	case NewtonSeconds:
		return "Ns"
	case NewtonsPerMeter:
		return "N/m"
	case WattsPerMeterPerDegreeKelvin:
		return "W/m/deg-K"
	case Other:
		return "other"
	default:
		return fmt.Sprintf("!Unit(0x%02d)", u)
	}
}

func (u Unit) Format(v float64) string {
	suf := ""
	if v != 1 {
		suf = "s"
	}
	switch u {
	case SquareMeters:
		return fmt.Sprintf("%f m²", v)
	case SquareFeet:
		return fmt.Sprintf("%f ft²", v)
	case Milliamperes:
		return fmt.Sprintf("%f mA", v)
	case Amperes:
		return fmt.Sprintf("%f A", v)
	case Ohms:
		return fmt.Sprintf("%f Ω", v)
	case Volts:
		return fmt.Sprintf("%f V", v)
	case KiloVolts:
		return fmt.Sprintf("%f kV", v)
	case MegaVolts:
		return fmt.Sprintf("%f MV", v)
	case VoltAmperes:
		return fmt.Sprintf("%f VA", v)
	case KiloVoltAmperes:
		return fmt.Sprintf("%f kVA", v)
	case MegaVoltAmperes:
		return fmt.Sprintf("%f MVA", v)
	case VoltAmperesReactive:
		return fmt.Sprintf("%f var", v)
	case KiloVoltAmperesReactive:
		return fmt.Sprintf("%f kvar", v)
	case MegaVoltAmperesReactive:
		return fmt.Sprintf("%f Mvar", v)
	case DegreesPhase:
		return fmt.Sprintf("%f °", v)
	case PowerFactor:
		return fmt.Sprintf("%f PFC", v)
	case Joules:
		return fmt.Sprintf("%f J", v)
	case Kilojoules:
		return fmt.Sprintf("%f kJ", v)
	case WattHours:
		return fmt.Sprintf("%f Wh", v)
	case KilowattHours:
		return fmt.Sprintf("%f kWh", v)
	case BTUs:
		return fmt.Sprintf("%f BTU", v)
	case Therms:
		return fmt.Sprintf("%f thm", v) // == 100000 BTU
	case TonHours:
		return fmt.Sprintf("%f Th", v)
	case JoulesPerKilogramDryAir:
		return fmt.Sprintf("%f J/kg dry air", v)
	case BTUsPerPoundDryAir:
		return fmt.Sprintf("%f BTU/lbs dry air", v)
	case CyclesPerHour:
		return fmt.Sprintf("%f cycles/hr", v)
	case CyclesPerMinute:
		return fmt.Sprintf("%f cycles/min", v)
	case Hertz:
		return fmt.Sprintf("%f Hz", v)
	case GramsOfWaterPerKilogramDryAir:
		return fmt.Sprintf("%f g water/kg dry air", v)
	case PercentRelativeHumidity:
		return fmt.Sprintf("%f %% RH", v)
	case Millimeters:
		return fmt.Sprintf("%fmm", v)
	case Meters:
		return fmt.Sprintf("%fm", v)
	case Inches:
		return fmt.Sprintf("%f\"", v)
	case Feet:
		return fmt.Sprintf("%f'", v)
	case WattsPerSquareFoot:
		return fmt.Sprintf("%f W/ft²", v)
	case WattsPerSquareMeter:
		return fmt.Sprintf("%f W/m²", v)
	case Lumens:
		return fmt.Sprintf("%f lm", v)
	case Luxes:
		return fmt.Sprintf("%f lux", v)
	case FootCandles:
		return fmt.Sprintf("%f lm/ft²", v)
	case Kilograms:
		return fmt.Sprintf("%f kg", v)
	case PoundsMass:
		return fmt.Sprintf("%f lbm", v)
	case Tons:
		return fmt.Sprintf("%f ton", v)
	case KilogramsPerSecond:
		return fmt.Sprintf("%f kg/s", v)
	case KilogramsPerMinute:
		return fmt.Sprintf("%f kg/min", v)
	case KilogramsPerHour:
		return fmt.Sprintf("%f kg/h", v)
	case PoundsMassPerMinute:
		return fmt.Sprintf("%f lbm/min", v)
	case PoundsMassPerHour:
		return fmt.Sprintf("%f lbm/h", v)
	case Watts:
		return fmt.Sprintf("%f W", v)
	case Kilowatts:
		return fmt.Sprintf("%f kW", v)
	case Megawatts:
		return fmt.Sprintf("%f MW", v)
	case BTUsPerHour:
		return fmt.Sprintf("%f BTU/h", v)
	case Horsepower:
		return fmt.Sprintf("%f hp", v)
	case TonsRefrigeration:
		return fmt.Sprintf("%f TR", v)
	case Pascals:
		return fmt.Sprintf("%f Pa", v)
	case Kilopascals:
		return fmt.Sprintf("%f kPa", v)
	case Bars:
		return fmt.Sprintf("%f bar", v)
	case PoundsForcePerSquareInch:
		return fmt.Sprintf("%f lbf/in²", v)
	case CentimetersOfWater:
		return fmt.Sprintf("%f cm h2o", v)
	case InchesOfWater:
		return fmt.Sprintf("%f in h2o", v)
	case MillimetersOfMercury:
		return fmt.Sprintf("%f mmHg", v)
	case CentimetersOfMercury:
		return fmt.Sprintf("%f cmHg", v)
	case InchesOfMercury:
		return fmt.Sprintf("%f inHg", v)
	case DegreesCelsius:
		return fmt.Sprintf("%f°C", v)
	case DegreesKelvin:
		return fmt.Sprintf("%fK", v)
	case DegreesFahrenheit:
		return fmt.Sprintf("%f°F", v)
	case DegreeDaysCelsius:
		return fmt.Sprintf("%f°C DD", v)
	case DegreeDaysFahrenheit:
		return fmt.Sprintf("%f°F DD", v)
	case Years:
		return fmt.Sprintf("%f year%s", v, suf)
	case Months:
		return fmt.Sprintf("%f month%s", v, suf)
	case Weeks:
		return fmt.Sprintf("%f week%s", v, suf)
	case Days:
		return fmt.Sprintf("%f day%s", v, suf)
	case Hours:
		return fmt.Sprintf("%f hour%s", v, suf)
	case Minutes:
		return fmt.Sprintf("%f minute%s", v, suf)
	case Seconds:
		return fmt.Sprintf("%f second%s", v, suf)
	case MetersPerSecond:
		return fmt.Sprintf("%f m/s", v)
	case KilometersPerHour:
		return fmt.Sprintf("%f km/h", v)
	case FeetPerSecond:
		return fmt.Sprintf("%f ft/s", v)
	case FeetPerMinute:
		return fmt.Sprintf("%f ft/min", v)
	case MilesPerHour:
		return fmt.Sprintf("%f mph", v)
	case CubicFeet:
		return fmt.Sprintf("%f ft³", v)
	case CubicMeters:
		return fmt.Sprintf("%f m³", v)
	case ImperialGallons:
		return fmt.Sprintf("%f gal (Imp.)", v)
	case Liters:
		return fmt.Sprintf("%f L", v)
	case UsGallons:
		return fmt.Sprintf("%f gal (US)", v)
	case CubicFeetPerMinute:
		return fmt.Sprintf("%f ft³/min", v)
	case CubicMetersPerSecond:
		return fmt.Sprintf("%f m³/s", v)
	case ImperialGallonsPerMinute:
		return fmt.Sprintf("%f gal/min (Imp.)", v)
	case LitersPerSecond:
		return fmt.Sprintf("%f L/s", v)
	case LitersPerMinute:
		return fmt.Sprintf("%f L/min", v)
	case UsGallonsPerMinute:
		return fmt.Sprintf("%f gal/min (US)", v)
	case DegreesAngular:
		return fmt.Sprintf("%f° angle", v)
	case DegreesCelsiusPerHour:
		return fmt.Sprintf("%f°C/h", v)
	case DegreesCelsiusPerMinute:
		return fmt.Sprintf("%f°C/min", v)
	case DegreesFahrenheitPerHour:
		return fmt.Sprintf("%f°F/h", v)
	case DegreesFahrenheitPerMinute:
		return fmt.Sprintf("%f°F/min", v)
	case NoUnits:
		return fmt.Sprintf("%f", v)
	case PartsPerMillion:
		return fmt.Sprintf("%f ppm", v)
	case PartsPerBillion:
		return fmt.Sprintf("%f ppb", v)
	case Percent:
		return fmt.Sprintf("%f %%", v)
	case PercentPerSecond:
		return fmt.Sprintf("%f %%/s", v)
	case PerMinute:
		return fmt.Sprintf("%f/min", v)
	case PerSecond:
		return fmt.Sprintf("%f/s", v)
	case PsiPerDegreeFahrenheit:
		return fmt.Sprintf("%f psi/°F", v)
	case Radians:
		return fmt.Sprintf("%f rad", v)
	case RevolutionsPerMinute:
		return fmt.Sprintf("%f RPM", v)
	case Currency1:
		return fmt.Sprintf("%f", v)
	case Currency2:
		return fmt.Sprintf("%f", v)
	case Currency3:
		return fmt.Sprintf("%f", v)
	case Currency4:
		return fmt.Sprintf("%f", v)
	case Currency5:
		return fmt.Sprintf("%f", v)
	case Currency6:
		return fmt.Sprintf("%f", v)
	case Currency7:
		return fmt.Sprintf("%f", v)
	case Currency8:
		return fmt.Sprintf("%f", v)
	case Currency9:
		return fmt.Sprintf("%f", v)
	case Currency10:
		return fmt.Sprintf("%f", v)
	case SquareInches:
		return fmt.Sprintf("%f in²", v)
	case SquareCentimeters:
		return fmt.Sprintf("%f cm²", v)
	case BTUsPerPound:
		return fmt.Sprintf("%f BTU/lbs", v)
	case Centimeters:
		return fmt.Sprintf("%f cm", v)
	case PoundsMassPerSecond:
		return fmt.Sprintf("%f lbm/s", v)
	case DeltaDegreesFahrenheit:
		return fmt.Sprintf("%f delta °F", v)
	case DeltaDegreesKelvin:
		return fmt.Sprintf("%f delta-K", v)
	case Kilohms:
		return fmt.Sprintf("%f kΩ", v)
	case Megohms:
		return fmt.Sprintf("%f MΩ", v)
	case Millivolts:
		return fmt.Sprintf("%f mV", v)
	case KilojoulesPerKilogram:
		return fmt.Sprintf("%f kJ/kg", v)
	case Megajoules:
		return fmt.Sprintf("%f MJ", v)
	case JoulesPerDegreeKelvin:
		return fmt.Sprintf("%f J/°K", v)
	case JoulesPerKilogramDegreeKelvin:
		return fmt.Sprintf("%f J/kg-K", v)
	case Kilohertz:
		return fmt.Sprintf("%f kHz", v)
	case Megahertz:
		return fmt.Sprintf("%f MHz", v)
	case PerHour:
		return fmt.Sprintf("%f/h", v)
	case Milliwatts:
		return fmt.Sprintf("%f mW", v)
	case Hectopascals:
		return fmt.Sprintf("%f hpa", v)
	case Millibars:
		return fmt.Sprintf("%f mbar", v)
	case CubicMetersPerHour:
		return fmt.Sprintf("%f m³/h", v)
	case LitersPerHour:
		return fmt.Sprintf("%f L/h", v)
	case KilowattHoursPerSquareMeter:
		return fmt.Sprintf("%f kWh/m²", v)
	case KilowattHoursPerSquareFoot:
		return fmt.Sprintf("%f kWh/ft²", v)
	case MegajoulesPerSquareMeter:
		return fmt.Sprintf("%f MJ/m²", v)
	case MegajoulesPerSquareFoot:
		return fmt.Sprintf("%f MJ/ft²", v)
	case WattsPerSquareMeterDegreeKelvin:
		return fmt.Sprintf("%f W/m²-K", v)
	case CubicFeetPerSecond:
		return fmt.Sprintf("%f ft³/s", v)
	case PercentObscurationPerFoot:
		return fmt.Sprintf("%f %% obscuration/ft", v)
	case PercentObscurationPerMeter:
		return fmt.Sprintf("%f %% obscuration/m", v)
	case Milliohms:
		return fmt.Sprintf("%f mΩ", v)
	case MegawattHours:
		return fmt.Sprintf("%f MWh", v)
	case KiloBTUs:
		return fmt.Sprintf("%f kBTU", v)
	case MegaBTUs:
		return fmt.Sprintf("%f MBTU", v)
	case KilojoulesPerKilogramDryAir:
		return fmt.Sprintf("%f KJ/kg dry air", v)
	case MegajoulesPerKilogramDryAir:
		return fmt.Sprintf("%f MJ/kg dry air", v)
	case KilojoulesPerDegreeKelvin:
		return fmt.Sprintf("%f KJ/deg-K", v)
	case MegajoulesPerDegreeKelvin:
		return fmt.Sprintf("%f MJ/deg-K", v)
	case Newton:
		return fmt.Sprintf("%f N", v)
	case GramsPerSecond:
		return fmt.Sprintf("%f g/s", v)
	case GramsPerMinute:
		return fmt.Sprintf("%f g/min", v)
	case TonsPerHour:
		return fmt.Sprintf("%f tons/h", v)
	case KiloBTUsPerHour:
		return fmt.Sprintf("%f kBTU/h", v)
	case HundredthsSeconds:
		return fmt.Sprintf("%f ms", v*10)
	case Milliseconds:
		return fmt.Sprintf("%f ms", v)
	case NewtonMeters:
		return fmt.Sprintf("%f Nm", v)
	case MillimetersPerSecond:
		return fmt.Sprintf("%f mm/s", v)
	case MillimetersPerMinute:
		return fmt.Sprintf("%f mm/min", v)
	case MetersPerMinute:
		return fmt.Sprintf("%f m/min", v)
	case MetersPerHour:
		return fmt.Sprintf("%f m/h", v)
	case CubicMetersPerMinute:
		return fmt.Sprintf("%f m³/min", v)
	case MetersPerSecondPerSecond:
		return fmt.Sprintf("%f m/s²", v)
	case AmperesPerMeter:
		return fmt.Sprintf("%f A/m", v)
	case AmperesPerSquareMeter:
		return fmt.Sprintf("%f A/m²", v)
	case AmpereSquareMeters:
		return fmt.Sprintf("%f Am²", v)
	case Farads:
		return fmt.Sprintf("%f Farad%s", v, suf)
	case Henrys:
		return fmt.Sprintf("%f Henry%s", v, suf)
	case OhmMeters:
		return fmt.Sprintf("%f Ωm", v)
	case Siemens:
		return fmt.Sprintf("%f Siemens", v)
	case SiemensPerMeter:
		return fmt.Sprintf("%f Siemens/m", v)
	case Teslas:
		return fmt.Sprintf("%f Tesla%s", v, suf)
	case VoltsPerDegreeKelvin:
		return fmt.Sprintf("%f V/deg-K", v)
	case VoltsPerMeter:
		return fmt.Sprintf("%f V/m", v)
	case Webers:
		return fmt.Sprintf("%f Weber%s", v, suf)
	case Candelas:
		return fmt.Sprintf("%f Candela%s", v, suf)
	case CandelasPerSquareMeter:
		return fmt.Sprintf("%f Candela/m²", v)
	case KelvinsPerHour:
		return fmt.Sprintf("%f K/h", v)
	case KelvinsPerMinute:
		return fmt.Sprintf("%f K/min", v)
	case JouleSeconds:
		return fmt.Sprintf("%f Js", v)
	case SquareMetersPerNewton:
		return fmt.Sprintf("%f m²/N", v)
	case KilogramPerCubicMeter:
		return fmt.Sprintf("%f kg/m³", v)
	case NewtonSeconds:
		return fmt.Sprintf("%f Ns", v)
	case NewtonsPerMeter:
		return fmt.Sprintf("%f N/m", v)
	case WattsPerMeterPerDegreeKelvin:
		return fmt.Sprintf("%f W/m/deg-K", v)
	default:
		return fmt.Sprintf("%f", v)
	}
}
