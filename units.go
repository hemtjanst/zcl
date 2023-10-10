package zcl

import (
	"fmt"
	"strconv"
	"time"
)

type Unit uint8
type CustomUnit uint8
type EngineeringUnit uint16

const (
	DecibelMilliWatts CustomUnit = iota
	MilliAmpereHours
	Mired
	Bytes
	Concentration
	MicrogramPerCubicMeter
)

func (u CustomUnit) Format(n float64) string {
	v := strconv.FormatFloat(n, 'f', -1, 64)
	v2 := fmt.Sprintf("%.3f", n)
	if len(v2) < len(v) {
		v = v2
	}

	switch u {
	case DecibelMilliWatts:
		return fmt.Sprintf("%sdBm", v)
	case MilliAmpereHours:
		return fmt.Sprintf("%smAh", v)
	case Mired:
		return fmt.Sprintf("%s mired", v)
	case Bytes:
		return fmt.Sprintf("%s bytes", v)
	case MicrogramPerCubicMeter:
		return fmt.Sprintf("%.1f µg/m³")
	case Concentration:
		if n >= 0.001 {
			return fmt.Sprintf("%.1f%%", n*100)
		}
		if n >= 0.000001 {
			return fmt.Sprintf("%.1f ppm", n*1000000)
		}
		return fmt.Sprintf("%.1f ppb", n*1000000000)
	}
	return fmt.Sprintf("%s", v)
}
func (u CustomUnit) String() string {
	switch u {
	case DecibelMilliWatts:
		return "dBm"
	case MilliAmpereHours:
		return "mAh"
	case Mired:
		return "mired"
	case Bytes:
		return "bytes"
	case MicrogramPerCubicMeter:
		return "µg/m³"
	case Concentration:
		return "concentration"
	}
	return fmt.Sprintf("CustomUnit(0x%02x)", uint8(u))
}

func (u *EngineeringUnit) UnmarshalZcl(buf []byte) ([]byte, error) {
	val, buf, err := uintLEUnmarshalZcl(2, buf)
	*u = EngineeringUnit(val)
	return buf, err
}
func (u EngineeringUnit) MarshalZcl() ([]byte, error) { return uintLEMarshalZcl(2, uint64(u)) }
func (EngineeringUnit) TypeID() TypeID                { return Zu16(0).TypeID() }
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
		return "% RH"
	case Millimeters:
		return "mm"
	case Meters:
		return "m"
	case Inches:
		return "in"
	case Feet:
		return "ft"
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
		return "years"
	case Months:
		return "months"
	case Weeks:
		return "weeks"
	case Days:
		return "days"
	case Hours:
		return "hours"
	case Minutes:
		return "minutes"
	case Seconds:
		return "seconds"
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
		return "°"
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
		return "%"
	case PercentPerSecond:
		return "%/s"
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
		return "% obscuration/ft"
	case PercentObscurationPerMeter:
		return "% obscuration/m"
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
		return "Farads"
	case Henrys:
		return "Henrys"
	case OhmMeters:
		return "Ωm"
	case Siemens:
		return "Siemens"
	case SiemensPerMeter:
		return "Siemens/m"
	case Teslas:
		return "Teslas"
	case VoltsPerDegreeKelvin:
		return "V/deg-K"
	case VoltsPerMeter:
		return "V/m"
	case Webers:
		return "Webers"
	case Candelas:
		return "Candelas"
	case CandelasPerSquareMeter:
		return "Candelas/m²"
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

func (u Unit) Format(n float64) string {
	suf := ""
	if n != 1 {
		suf = "s"
	}
	v := strconv.FormatFloat(n, 'f', -1, 64)
	v2 := fmt.Sprintf("%.3f", n)
	if len(v2) < len(v) {
		v = v2
	}
	switch u {
	case SquareMeters:
		return v + "m²"
	case SquareFeet:
		return v + "ft²"
	case Milliamperes:
		return v + "mA"
	case Amperes:
		return v + "A"
	case Ohms:
		return v + " Ω"
	case Volts:
		return v + "V"
	case KiloVolts:
		return v + "kV"
	case MegaVolts:
		return v + "MV"
	case VoltAmperes:
		return v + "VA"
	case KiloVoltAmperes:
		return v + "kVA"
	case MegaVoltAmperes:
		return v + "MVA"
	case VoltAmperesReactive:
		return v + " var"
	case KiloVoltAmperesReactive:
		return v + " kvar"
	case MegaVoltAmperesReactive:
		return v + " Mvar"
	case DegreesPhase:
		return v + "°"
	case PowerFactor:
		return v + " PFC"
	case Joules:
		return v + "J"
	case Kilojoules:
		return v + "kJ"
	case WattHours:
		return v + "Wh"
	case KilowattHours:
		return v + "kWh"
	case BTUs:
		return v + " BTU"
	case Therms:
		return v + " thm" // == 100000 BTU
	case TonHours:
		return v + " Th"
	case JoulesPerKilogramDryAir:
		return v + " J/kg dry air"
	case BTUsPerPoundDryAir:
		return v + " BTU/lbs dry air"
	case CyclesPerHour:
		return v + " cycles/hr"
	case CyclesPerMinute:
		return v + " cycles/min"
	case Hertz:
		return v + "Hz"
	case GramsOfWaterPerKilogramDryAir:
		return v + "g water/kg dry air"
	case PercentRelativeHumidity:
		return v + "% RH"
	case Millimeters:
		return v + "mm"
	case Meters:
		return v + "m"
	case Inches:
		return v + "in"
	case Feet:
		return v + "ft"
	case WattsPerSquareFoot:
		return v + "W/ft²"
	case WattsPerSquareMeter:
		return v + "W/m²"
	case Lumens:
		return v + "lm"
	case Luxes:
		return v + "lux"
	case FootCandles:
		return v + "lm/ft²"
	case Kilograms:
		return v + "kg"
	case PoundsMass:
		return v + "lbm"
	case Tons:
		return v + "ton"
	case KilogramsPerSecond:
		return v + "kg/s"
	case KilogramsPerMinute:
		return v + "kg/min"
	case KilogramsPerHour:
		return v + "kg/h"
	case PoundsMassPerMinute:
		return v + "lbm/min"
	case PoundsMassPerHour:
		return v + "lbm/h"
	case Watts:
		return v + "W"
	case Kilowatts:
		return v + "kW"
	case Megawatts:
		return v + "MW"
	case BTUsPerHour:
		return v + " BTU/h"
	case Horsepower:
		return v + "hp"
	case TonsRefrigeration:
		return v + " TR"
	case Pascals:
		return v + "Pa"
	case Kilopascals:
		return v + "kPa"
	case Bars:
		return v + " bar"
	case PoundsForcePerSquareInch:
		return v + "lbf/in²"
	case CentimetersOfWater:
		return v + "cm h2o"
	case InchesOfWater:
		return v + "in h2o"
	case MillimetersOfMercury:
		return v + "mmHg"
	case CentimetersOfMercury:
		return v + "cmHg"
	case InchesOfMercury:
		return v + "inHg"
	case DegreesCelsius:
		return v + "°C"
	case DegreesKelvin:
		return v + "K"
	case DegreesFahrenheit:
		return v + "°F"
	case DegreeDaysCelsius:
		return v + "°C DD"
	case DegreeDaysFahrenheit:
		return v + "°F DD"
	case Years:
		return v + " year" + suf
	case Months:
		return v + " month" + suf
	case Weeks:
		return v + " week" + suf
	case Days:
		return v + " day" + suf
	case Hours:
		return time.Duration(n * float64(time.Hour)).String()
	case Minutes:
		return time.Duration(n * float64(time.Minute)).String()
	case Seconds:
		return time.Duration(n * float64(time.Second)).String()
	case MetersPerSecond:
		return v + "m/s"
	case KilometersPerHour:
		return v + "km/h"
	case FeetPerSecond:
		return v + "ft/s"
	case FeetPerMinute:
		return v + "ft/min"
	case MilesPerHour:
		return v + "mph"
	case CubicFeet:
		return v + "ft³"
	case CubicMeters:
		return v + "m³"
	case ImperialGallons:
		return v + "gal (Imp.)"
	case Liters:
		return v + "L"
	case UsGallons:
		return v + "gal (US)"
	case CubicFeetPerMinute:
		return v + "ft³/min"
	case CubicMetersPerSecond:
		return v + "m³/s"
	case ImperialGallonsPerMinute:
		return v + "gal/min (Imp.)"
	case LitersPerSecond:
		return v + "L/s"
	case LitersPerMinute:
		return v + "L/min"
	case UsGallonsPerMinute:
		return v + "gal/min (US)"
	case DegreesAngular:
		return v + "°"
	case DegreesCelsiusPerHour:
		return v + "°C/h"
	case DegreesCelsiusPerMinute:
		return v + "°C/min"
	case DegreesFahrenheitPerHour:
		return v + "°F/h"
	case DegreesFahrenheitPerMinute:
		return v + "°F/min"
	case NoUnits:
		return v
	case PartsPerMillion:
		return v + "ppm"
	case PartsPerBillion:
		return v + "ppb"
	case Percent:
		return v + "%"
	case PercentPerSecond:
		return v + "%/s"
	case PerMinute:
		return v + "/min"
	case PerSecond:
		return v + "/s"
	case PsiPerDegreeFahrenheit:
		return v + "psi/°F"
	case Radians:
		return v + "rad"
	case RevolutionsPerMinute:
		return v + "RPM"
	case Currency1:
		return v
	case Currency2:
		return v
	case Currency3:
		return v
	case Currency4:
		return v
	case Currency5:
		return v
	case Currency6:
		return v
	case Currency7:
		return v
	case Currency8:
		return v
	case Currency9:
		return v
	case Currency10:
		return v
	case SquareInches:
		return v + "in²"
	case SquareCentimeters:
		return v + "cm²"
	case BTUsPerPound:
		return v + " BTU/lbs"
	case Centimeters:
		return v + "cm"
	case PoundsMassPerSecond:
		return v + "lbm/s"
	case DeltaDegreesFahrenheit:
		return v + " delta °F"
	case DeltaDegreesKelvin:
		return v + " delta-K"
	case Kilohms:
		return v + " kΩ"
	case Megohms:
		return v + " MΩ"
	case Millivolts:
		return v + "mV"
	case KilojoulesPerKilogram:
		return v + "kJ/kg"
	case Megajoules:
		return v + " MJ"
	case JoulesPerDegreeKelvin:
		return v + " J/°K"
	case JoulesPerKilogramDegreeKelvin:
		return v + " J/kg-K"
	case Kilohertz:
		return v + "kHz"
	case Megahertz:
		return v + "MHz"
	case PerHour:
		return v + "/h"
	case Milliwatts:
		return v + "mW"
	case Hectopascals:
		return v + " hpa"
	case Millibars:
		return v + " mbar"
	case CubicMetersPerHour:
		return v + "m³/h"
	case LitersPerHour:
		return v + "L/h"
	case KilowattHoursPerSquareMeter:
		return v + "kWh/m²"
	case KilowattHoursPerSquareFoot:
		return v + "kWh/ft²"
	case MegajoulesPerSquareMeter:
		return v + " MJ/m²"
	case MegajoulesPerSquareFoot:
		return v + " MJ/ft²"
	case WattsPerSquareMeterDegreeKelvin:
		return v + "W/m²-K"
	case CubicFeetPerSecond:
		return v + "ft³/s"
	case PercentObscurationPerFoot:
		return v + "% obscuration/ft"
	case PercentObscurationPerMeter:
		return v + "% obscuration/m"
	case Milliohms:
		return v + " mΩ"
	case MegawattHours:
		return v + "MWh"
	case KiloBTUs:
		return v + " kBTU"
	case MegaBTUs:
		return v + " MBTU"
	case KilojoulesPerKilogramDryAir:
		return v + " KJ/kg dry air"
	case MegajoulesPerKilogramDryAir:
		return v + " MJ/kg dry air"
	case KilojoulesPerDegreeKelvin:
		return v + " KJ/deg-K"
	case MegajoulesPerDegreeKelvin:
		return v + " MJ/deg-K"
	case Newton:
		return v + "N"
	case GramsPerSecond:
		return v + "g/s"
	case GramsPerMinute:
		return v + "g/min"
	case TonsPerHour:
		return v + " tons/h"
	case KiloBTUsPerHour:
		return v + " kBTU/h"
	case HundredthsSeconds:
		return time.Duration(n * float64(time.Millisecond) * 10).String()
	case Milliseconds:
		return time.Duration(n * float64(time.Millisecond)).String()
	case NewtonMeters:
		return v + "Nm"
	case MillimetersPerSecond:
		return v + "mm/s"
	case MillimetersPerMinute:
		return v + "mm/min"
	case MetersPerMinute:
		return v + "m/min"
	case MetersPerHour:
		return v + "m/h"
	case CubicMetersPerMinute:
		return v + "m³/min"
	case MetersPerSecondPerSecond:
		return v + "m/s²"
	case AmperesPerMeter:
		return v + "A/m"
	case AmperesPerSquareMeter:
		return v + "A/m²"
	case AmpereSquareMeters:
		return v + "Am²"
	case Farads:
		return v + " Farad" + suf
	case Henrys:
		return v + " Henry" + suf
	case OhmMeters:
		return v + " Ωm"
	case Siemens:
		return v + " Siemens"
	case SiemensPerMeter:
		return v + " Siemens/m"
	case Teslas:
		return v + " Tesla" + suf
	case VoltsPerDegreeKelvin:
		return v + "V/deg-K"
	case VoltsPerMeter:
		return v + "V/m"
	case Webers:
		return v + " Weber" + suf
	case Candelas:
		return v + " Candela" + suf
	case CandelasPerSquareMeter:
		return v + " Candela/m²"
	case KelvinsPerHour:
		return v + "K/h"
	case KelvinsPerMinute:
		return v + "K/min"
	case JouleSeconds:
		return v + " Js"
	case SquareMetersPerNewton:
		return v + "m²/N"
	case KilogramPerCubicMeter:
		return v + "kg/m³"
	case NewtonSeconds:
		return v + "Ns"
	case NewtonsPerMeter:
		return v + "N/m"
	case WattsPerMeterPerDegreeKelvin:
		return v + " W/m/deg-K"
	default:
		return v
	}
}
