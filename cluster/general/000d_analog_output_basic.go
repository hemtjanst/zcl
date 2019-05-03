// The Analog Output (Basic) cluster provides an interface for setting the value of an analog output (typically to the environment) and accessing various characteristics of that value.
package general

import (
	"neotor.se/zcl"
)

// AnalogOutputBasic
const AnalogOutputBasicID zcl.ClusterID = 13

var AnalogOutputBasicCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		AnalogOutputDescriptionAttr:       func() zcl.Attr { return new(AnalogOutputDescription) },
		AnalogOutputMaxPresentValueAttr:   func() zcl.Attr { return new(AnalogOutputMaxPresentValue) },
		AnalogOutputMinPresentValueAttr:   func() zcl.Attr { return new(AnalogOutputMinPresentValue) },
		AnalogOutputOutOfServiceAttr:      func() zcl.Attr { return new(AnalogOutputOutOfService) },
		AnalogOutputPresentValueAttr:      func() zcl.Attr { return new(AnalogOutputPresentValue) },
		AnalogOutputPriorityArrayAttr:     func() zcl.Attr { return new(AnalogOutputPriorityArray) },
		AnalogOutputReliabilityAttr:       func() zcl.Attr { return new(AnalogOutputReliability) },
		AnalogOutputRelinquishDefaultAttr: func() zcl.Attr { return new(AnalogOutputRelinquishDefault) },
		AnalogOutputResolutionAttr:        func() zcl.Attr { return new(AnalogOutputResolution) },
		AnalogOutputStatusFlagsAttr:       func() zcl.Attr { return new(AnalogOutputStatusFlags) },
		AnalogOutputEngineeringUnitsAttr:  func() zcl.Attr { return new(AnalogOutputEngineeringUnits) },
		AnalogOutputApplicationTypeAttr:   func() zcl.Attr { return new(AnalogOutputApplicationType) },
		AnalogOutputXiaomi0Xf000Attr:      func() zcl.Attr { return new(AnalogOutputXiaomi0Xf000) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

// AnalogOutputDescription is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputDescription zcl.Zcstring

const AnalogOutputDescriptionAttr zcl.AttrID = 28

func (AnalogOutputDescription) ID() zcl.AttrID                     { return AnalogOutputDescriptionAttr }
func (AnalogOutputDescription) Cluster() zcl.ClusterID             { return AnalogOutputBasicID }
func (AnalogOutputDescription) Name() string                       { return "Analog Output Description" }
func (AnalogOutputDescription) Readable() bool                     { return true }
func (AnalogOutputDescription) Writable() bool                     { return true }
func (AnalogOutputDescription) Reportable() bool                   { return false }
func (AnalogOutputDescription) SceneIndex() int                    { return -1 }
func (a *AnalogOutputDescription) Value() *AnalogOutputDescription { return a }
func (a AnalogOutputDescription) MarshalZcl() ([]byte, error)      { return zcl.Zcstring(a).MarshalZcl() }

func (a *AnalogOutputDescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputDescription(*nt)
	return br, err
}

func (a AnalogOutputDescription) String() string {
	return zcl.Sprintf("%v", zcl.Zcstring(a))
}

// AnalogOutputMaxPresentValue is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputMaxPresentValue zcl.Zfloat

const AnalogOutputMaxPresentValueAttr zcl.AttrID = 65

func (AnalogOutputMaxPresentValue) ID() zcl.AttrID                         { return AnalogOutputMaxPresentValueAttr }
func (AnalogOutputMaxPresentValue) Cluster() zcl.ClusterID                 { return AnalogOutputBasicID }
func (AnalogOutputMaxPresentValue) Name() string                           { return "Analog Output Max Present Value" }
func (AnalogOutputMaxPresentValue) Readable() bool                         { return true }
func (AnalogOutputMaxPresentValue) Writable() bool                         { return true }
func (AnalogOutputMaxPresentValue) Reportable() bool                       { return false }
func (AnalogOutputMaxPresentValue) SceneIndex() int                        { return -1 }
func (a *AnalogOutputMaxPresentValue) Value() *AnalogOutputMaxPresentValue { return a }
func (a AnalogOutputMaxPresentValue) MarshalZcl() ([]byte, error)          { return zcl.Zfloat(a).MarshalZcl() }

func (a *AnalogOutputMaxPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputMaxPresentValue(*nt)
	return br, err
}

func (a AnalogOutputMaxPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(a))
}

// AnalogOutputMinPresentValue is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputMinPresentValue zcl.Zfloat

const AnalogOutputMinPresentValueAttr zcl.AttrID = 69

func (AnalogOutputMinPresentValue) ID() zcl.AttrID                         { return AnalogOutputMinPresentValueAttr }
func (AnalogOutputMinPresentValue) Cluster() zcl.ClusterID                 { return AnalogOutputBasicID }
func (AnalogOutputMinPresentValue) Name() string                           { return "Analog Output Min Present Value" }
func (AnalogOutputMinPresentValue) Readable() bool                         { return true }
func (AnalogOutputMinPresentValue) Writable() bool                         { return true }
func (AnalogOutputMinPresentValue) Reportable() bool                       { return false }
func (AnalogOutputMinPresentValue) SceneIndex() int                        { return -1 }
func (a *AnalogOutputMinPresentValue) Value() *AnalogOutputMinPresentValue { return a }
func (a AnalogOutputMinPresentValue) MarshalZcl() ([]byte, error)          { return zcl.Zfloat(a).MarshalZcl() }

func (a *AnalogOutputMinPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputMinPresentValue(*nt)
	return br, err
}

func (a AnalogOutputMinPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(a))
}

// AnalogOutputOutOfService is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputOutOfService zcl.Zbool

const AnalogOutputOutOfServiceAttr zcl.AttrID = 81

func (AnalogOutputOutOfService) ID() zcl.AttrID                      { return AnalogOutputOutOfServiceAttr }
func (AnalogOutputOutOfService) Cluster() zcl.ClusterID              { return AnalogOutputBasicID }
func (AnalogOutputOutOfService) Name() string                        { return "Analog Output Out of service" }
func (AnalogOutputOutOfService) Readable() bool                      { return true }
func (AnalogOutputOutOfService) Writable() bool                      { return true }
func (AnalogOutputOutOfService) Reportable() bool                    { return false }
func (AnalogOutputOutOfService) SceneIndex() int                     { return -1 }
func (a *AnalogOutputOutOfService) Value() *AnalogOutputOutOfService { return a }
func (a AnalogOutputOutOfService) MarshalZcl() ([]byte, error)       { return zcl.Zbool(a).MarshalZcl() }

func (a *AnalogOutputOutOfService) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputOutOfService(*nt)
	return br, err
}

func (a AnalogOutputOutOfService) String() string {
	return zcl.Sprintf("%v", zcl.Zbool(a))
}

// AnalogOutputPresentValue is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputPresentValue zcl.Zfloat

const AnalogOutputPresentValueAttr zcl.AttrID = 85

func (AnalogOutputPresentValue) ID() zcl.AttrID                      { return AnalogOutputPresentValueAttr }
func (AnalogOutputPresentValue) Cluster() zcl.ClusterID              { return AnalogOutputBasicID }
func (AnalogOutputPresentValue) Name() string                        { return "Analog Output Present value" }
func (AnalogOutputPresentValue) Readable() bool                      { return true }
func (AnalogOutputPresentValue) Writable() bool                      { return true }
func (AnalogOutputPresentValue) Reportable() bool                    { return true }
func (AnalogOutputPresentValue) SceneIndex() int                     { return -1 }
func (a *AnalogOutputPresentValue) Value() *AnalogOutputPresentValue { return a }
func (a AnalogOutputPresentValue) MarshalZcl() ([]byte, error)       { return zcl.Zfloat(a).MarshalZcl() }

func (a *AnalogOutputPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputPresentValue(*nt)
	return br, err
}

func (a AnalogOutputPresentValue) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(a))
}

// AnalogOutputPriorityArray is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputPriorityArray zcl.Zarray

const AnalogOutputPriorityArrayAttr zcl.AttrID = 87

func (AnalogOutputPriorityArray) ID() zcl.AttrID                       { return AnalogOutputPriorityArrayAttr }
func (AnalogOutputPriorityArray) Cluster() zcl.ClusterID               { return AnalogOutputBasicID }
func (AnalogOutputPriorityArray) Name() string                         { return "Analog Output Priority Array" }
func (AnalogOutputPriorityArray) Readable() bool                       { return true }
func (AnalogOutputPriorityArray) Writable() bool                       { return true }
func (AnalogOutputPriorityArray) Reportable() bool                     { return false }
func (AnalogOutputPriorityArray) SceneIndex() int                      { return -1 }
func (a *AnalogOutputPriorityArray) Value() *AnalogOutputPriorityArray { return a }
func (a AnalogOutputPriorityArray) MarshalZcl() ([]byte, error)        { return zcl.Zarray(a).MarshalZcl() }

func (a *AnalogOutputPriorityArray) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zarray)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputPriorityArray(*nt)
	return br, err
}

func (a AnalogOutputPriorityArray) String() string {
	return zcl.Sprintf("%v", zcl.Zarray(a))
}

// AnalogOutputReliability is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputReliability zcl.Zenum8

const AnalogOutputReliabilityAttr zcl.AttrID = 103

func (AnalogOutputReliability) ID() zcl.AttrID                     { return AnalogOutputReliabilityAttr }
func (AnalogOutputReliability) Cluster() zcl.ClusterID             { return AnalogOutputBasicID }
func (AnalogOutputReliability) Name() string                       { return "Analog Output Reliability" }
func (AnalogOutputReliability) Readable() bool                     { return true }
func (AnalogOutputReliability) Writable() bool                     { return true }
func (AnalogOutputReliability) Reportable() bool                   { return false }
func (AnalogOutputReliability) SceneIndex() int                    { return -1 }
func (a *AnalogOutputReliability) Value() *AnalogOutputReliability { return a }
func (a AnalogOutputReliability) MarshalZcl() ([]byte, error)      { return zcl.Zenum8(a).MarshalZcl() }

func (a *AnalogOutputReliability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputReliability(*nt)
	return br, err
}

func (a AnalogOutputReliability) String() string {
	switch a {
	case 0x00:
		return "No fault detected"
	case 0x01:
		return "No Sensor"
	case 0x02:
		return "Over Range"
	case 0x03:
		return "Under Range"
	case 0x04:
		return "Open Loop"
	case 0x05:
		return "Shorted Loop"
	case 0x06:
		return "No Output"
	case 0x07:
		return "Unreliable (other)"
	case 0x08:
		return "Process Error"
	case 0x09:
		return "Multi state fault"
	case 0x0A:
		return "Configuration Error"
	}
	return zcl.Sprintf("%v", zcl.Zenum8(a))
}

// IsNoFaultDetected checks if AnalogOutputReliability equals the value for No fault detected (0x00)
func (a AnalogOutputReliability) IsNoFaultDetected() bool { return a == 0x00 }

// SetNoFaultDetected sets AnalogOutputReliability to No fault detected (0x00)
func (a *AnalogOutputReliability) SetNoFaultDetected() { *a = 0x00 }

// IsNoSensor checks if AnalogOutputReliability equals the value for No Sensor (0x01)
func (a AnalogOutputReliability) IsNoSensor() bool { return a == 0x01 }

// SetNoSensor sets AnalogOutputReliability to No Sensor (0x01)
func (a *AnalogOutputReliability) SetNoSensor() { *a = 0x01 }

// IsOverRange checks if AnalogOutputReliability equals the value for Over Range (0x02)
func (a AnalogOutputReliability) IsOverRange() bool { return a == 0x02 }

// SetOverRange sets AnalogOutputReliability to Over Range (0x02)
func (a *AnalogOutputReliability) SetOverRange() { *a = 0x02 }

// IsUnderRange checks if AnalogOutputReliability equals the value for Under Range (0x03)
func (a AnalogOutputReliability) IsUnderRange() bool { return a == 0x03 }

// SetUnderRange sets AnalogOutputReliability to Under Range (0x03)
func (a *AnalogOutputReliability) SetUnderRange() { *a = 0x03 }

// IsOpenLoop checks if AnalogOutputReliability equals the value for Open Loop (0x04)
func (a AnalogOutputReliability) IsOpenLoop() bool { return a == 0x04 }

// SetOpenLoop sets AnalogOutputReliability to Open Loop (0x04)
func (a *AnalogOutputReliability) SetOpenLoop() { *a = 0x04 }

// IsShortedLoop checks if AnalogOutputReliability equals the value for Shorted Loop (0x05)
func (a AnalogOutputReliability) IsShortedLoop() bool { return a == 0x05 }

// SetShortedLoop sets AnalogOutputReliability to Shorted Loop (0x05)
func (a *AnalogOutputReliability) SetShortedLoop() { *a = 0x05 }

// IsNoOutput checks if AnalogOutputReliability equals the value for No Output (0x06)
func (a AnalogOutputReliability) IsNoOutput() bool { return a == 0x06 }

// SetNoOutput sets AnalogOutputReliability to No Output (0x06)
func (a *AnalogOutputReliability) SetNoOutput() { *a = 0x06 }

// IsUnreliableOther checks if AnalogOutputReliability equals the value for Unreliable (other) (0x07)
func (a AnalogOutputReliability) IsUnreliableOther() bool { return a == 0x07 }

// SetUnreliableOther sets AnalogOutputReliability to Unreliable (other) (0x07)
func (a *AnalogOutputReliability) SetUnreliableOther() { *a = 0x07 }

// IsProcessError checks if AnalogOutputReliability equals the value for Process Error (0x08)
func (a AnalogOutputReliability) IsProcessError() bool { return a == 0x08 }

// SetProcessError sets AnalogOutputReliability to Process Error (0x08)
func (a *AnalogOutputReliability) SetProcessError() { *a = 0x08 }

// IsMultiStateFault checks if AnalogOutputReliability equals the value for Multi state fault (0x09)
func (a AnalogOutputReliability) IsMultiStateFault() bool { return a == 0x09 }

// SetMultiStateFault sets AnalogOutputReliability to Multi state fault (0x09)
func (a *AnalogOutputReliability) SetMultiStateFault() { *a = 0x09 }

// IsConfigurationError checks if AnalogOutputReliability equals the value for Configuration Error (0x0A)
func (a AnalogOutputReliability) IsConfigurationError() bool { return a == 0x0A }

// SetConfigurationError sets AnalogOutputReliability to Configuration Error (0x0A)
func (a *AnalogOutputReliability) SetConfigurationError() { *a = 0x0A }

// AnalogOutputRelinquishDefault is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputRelinquishDefault zcl.Zfloat

const AnalogOutputRelinquishDefaultAttr zcl.AttrID = 104

func (AnalogOutputRelinquishDefault) ID() zcl.AttrID                           { return AnalogOutputRelinquishDefaultAttr }
func (AnalogOutputRelinquishDefault) Cluster() zcl.ClusterID                   { return AnalogOutputBasicID }
func (AnalogOutputRelinquishDefault) Name() string                             { return "Analog Output Relinquish Default" }
func (AnalogOutputRelinquishDefault) Readable() bool                           { return true }
func (AnalogOutputRelinquishDefault) Writable() bool                           { return true }
func (AnalogOutputRelinquishDefault) Reportable() bool                         { return false }
func (AnalogOutputRelinquishDefault) SceneIndex() int                          { return -1 }
func (a *AnalogOutputRelinquishDefault) Value() *AnalogOutputRelinquishDefault { return a }
func (a AnalogOutputRelinquishDefault) MarshalZcl() ([]byte, error)            { return zcl.Zfloat(a).MarshalZcl() }

func (a *AnalogOutputRelinquishDefault) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputRelinquishDefault(*nt)
	return br, err
}

func (a AnalogOutputRelinquishDefault) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(a))
}

// AnalogOutputResolution is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputResolution zcl.Zfloat

const AnalogOutputResolutionAttr zcl.AttrID = 106

func (AnalogOutputResolution) ID() zcl.AttrID                    { return AnalogOutputResolutionAttr }
func (AnalogOutputResolution) Cluster() zcl.ClusterID            { return AnalogOutputBasicID }
func (AnalogOutputResolution) Name() string                      { return "Analog Output Resolution" }
func (AnalogOutputResolution) Readable() bool                    { return true }
func (AnalogOutputResolution) Writable() bool                    { return true }
func (AnalogOutputResolution) Reportable() bool                  { return false }
func (AnalogOutputResolution) SceneIndex() int                   { return -1 }
func (a *AnalogOutputResolution) Value() *AnalogOutputResolution { return a }
func (a AnalogOutputResolution) MarshalZcl() ([]byte, error)     { return zcl.Zfloat(a).MarshalZcl() }

func (a *AnalogOutputResolution) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputResolution(*nt)
	return br, err
}

func (a AnalogOutputResolution) String() string {
	return zcl.Sprintf("%v", zcl.Zfloat(a))
}

// AnalogOutputStatusFlags is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputStatusFlags zcl.Zbmp8

const AnalogOutputStatusFlagsAttr zcl.AttrID = 111

func (AnalogOutputStatusFlags) ID() zcl.AttrID                     { return AnalogOutputStatusFlagsAttr }
func (AnalogOutputStatusFlags) Cluster() zcl.ClusterID             { return AnalogOutputBasicID }
func (AnalogOutputStatusFlags) Name() string                       { return "Analog Output Status flags" }
func (AnalogOutputStatusFlags) Readable() bool                     { return true }
func (AnalogOutputStatusFlags) Writable() bool                     { return false }
func (AnalogOutputStatusFlags) Reportable() bool                   { return true }
func (AnalogOutputStatusFlags) SceneIndex() int                    { return -1 }
func (a *AnalogOutputStatusFlags) Value() *AnalogOutputStatusFlags { return a }
func (a AnalogOutputStatusFlags) MarshalZcl() ([]byte, error)      { return zcl.Zbmp8(a).MarshalZcl() }

func (a *AnalogOutputStatusFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputStatusFlags(*nt)
	return br, err
}

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
	return zcl.BitmapTest([]byte(a), 0)
}
func (a *AnalogOutputStatusFlags) SetInAlarm(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 0, b))
}

func (a AnalogOutputStatusFlags) IsFault() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *AnalogOutputStatusFlags) SetFault(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a AnalogOutputStatusFlags) IsOveridden() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *AnalogOutputStatusFlags) SetOveridden(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a AnalogOutputStatusFlags) IsOutOfService() bool {
	return zcl.BitmapTest([]byte(a), 3)
}
func (a *AnalogOutputStatusFlags) SetOutOfService(b bool) {
	*a = AnalogOutputStatusFlags(zcl.BitmapSet([]byte(*a), 3, b))
}

// AnalogOutputEngineeringUnits is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputEngineeringUnits zcl.EngineeringUnit

const AnalogOutputEngineeringUnitsAttr zcl.AttrID = 117

func (AnalogOutputEngineeringUnits) ID() zcl.AttrID                          { return AnalogOutputEngineeringUnitsAttr }
func (AnalogOutputEngineeringUnits) Cluster() zcl.ClusterID                  { return AnalogOutputBasicID }
func (AnalogOutputEngineeringUnits) Name() string                            { return "Analog Output Engineering Units" }
func (AnalogOutputEngineeringUnits) Readable() bool                          { return true }
func (AnalogOutputEngineeringUnits) Writable() bool                          { return true }
func (AnalogOutputEngineeringUnits) Reportable() bool                        { return false }
func (AnalogOutputEngineeringUnits) SceneIndex() int                         { return -1 }
func (a *AnalogOutputEngineeringUnits) Value() *AnalogOutputEngineeringUnits { return a }
func (a AnalogOutputEngineeringUnits) MarshalZcl() ([]byte, error) {
	return zcl.EngineeringUnit(a).MarshalZcl()
}

func (a *AnalogOutputEngineeringUnits) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.EngineeringUnit)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputEngineeringUnits(*nt)
	return br, err
}

func (a AnalogOutputEngineeringUnits) String() string {
	return zcl.Sprintf("%v", zcl.EngineeringUnit(a))
}

// AnalogOutputApplicationType is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputApplicationType zcl.Zu32

const AnalogOutputApplicationTypeAttr zcl.AttrID = 256

func (AnalogOutputApplicationType) ID() zcl.AttrID                         { return AnalogOutputApplicationTypeAttr }
func (AnalogOutputApplicationType) Cluster() zcl.ClusterID                 { return AnalogOutputBasicID }
func (AnalogOutputApplicationType) Name() string                           { return "Analog Output Application Type" }
func (AnalogOutputApplicationType) Readable() bool                         { return true }
func (AnalogOutputApplicationType) Writable() bool                         { return false }
func (AnalogOutputApplicationType) Reportable() bool                       { return false }
func (AnalogOutputApplicationType) SceneIndex() int                        { return -1 }
func (a *AnalogOutputApplicationType) Value() *AnalogOutputApplicationType { return a }
func (a AnalogOutputApplicationType) MarshalZcl() ([]byte, error)          { return zcl.Zu32(a).MarshalZcl() }

func (a *AnalogOutputApplicationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputApplicationType(*nt)
	return br, err
}

func (a AnalogOutputApplicationType) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}

// AnalogOutputXiaomi0Xf000 is an autogenerated attribute in the AnalogOutputBasic cluster
type AnalogOutputXiaomi0Xf000 zcl.Zu32

const AnalogOutputXiaomi0Xf000Attr zcl.AttrID = 61440

func (AnalogOutputXiaomi0Xf000) ID() zcl.AttrID                      { return AnalogOutputXiaomi0Xf000Attr }
func (AnalogOutputXiaomi0Xf000) Cluster() zcl.ClusterID              { return AnalogOutputBasicID }
func (AnalogOutputXiaomi0Xf000) Name() string                        { return "Analog Output Xiaomi 0xf000" }
func (AnalogOutputXiaomi0Xf000) Readable() bool                      { return true }
func (AnalogOutputXiaomi0Xf000) Writable() bool                      { return false }
func (AnalogOutputXiaomi0Xf000) Reportable() bool                    { return false }
func (AnalogOutputXiaomi0Xf000) SceneIndex() int                     { return -1 }
func (a *AnalogOutputXiaomi0Xf000) Value() *AnalogOutputXiaomi0Xf000 { return a }
func (a AnalogOutputXiaomi0Xf000) MarshalZcl() ([]byte, error)       { return zcl.Zu32(a).MarshalZcl() }

func (a *AnalogOutputXiaomi0Xf000) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogOutputXiaomi0Xf000(*nt)
	return br, err
}

func (a AnalogOutputXiaomi0Xf000) String() string {
	return zcl.Sprintf("%v", zcl.Zu32(a))
}
