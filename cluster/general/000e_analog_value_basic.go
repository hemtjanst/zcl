// An interface for setting an analog value, typically used as a control system parameter, and accessing various characteristics of that value.
package general

import (
	"neotor.se/zcl"
)

// AnalogValueBasic
const AnalogValueBasicID zcl.ClusterID = 14

var AnalogValueBasicCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		AnalogValueDescriptionAttr:       func() zcl.Attr { return new(AnalogValueDescription) },
		AnalogValueOutOfServiceAttr:      func() zcl.Attr { return new(AnalogValueOutOfService) },
		AnalogValuePresentValueAttr:      func() zcl.Attr { return new(AnalogValuePresentValue) },
		AnalogValuePriorityArrayAttr:     func() zcl.Attr { return new(AnalogValuePriorityArray) },
		AnalogValueReliabilityAttr:       func() zcl.Attr { return new(AnalogValueReliability) },
		AnalogValueRelinquishDefaultAttr: func() zcl.Attr { return new(AnalogValueRelinquishDefault) },
		AnalogValueResolutionAttr:        func() zcl.Attr { return new(AnalogValueResolution) },
		AnalogValueStatusFlagsAttr:       func() zcl.Attr { return new(AnalogValueStatusFlags) },
		AnalogValueEngineeringUnitsAttr:  func() zcl.Attr { return new(AnalogValueEngineeringUnits) },
		AnalogValueApplicationTypeAttr:   func() zcl.Attr { return new(AnalogValueApplicationType) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const AnalogValueDescriptionAttr zcl.AttrID = 28

type AnalogValueDescription zcl.Zcstring

func (a AnalogValueDescription) ID() zcl.AttrID                  { return AnalogValueDescriptionAttr }
func (a AnalogValueDescription) Cluster() zcl.ClusterID          { return AnalogValueBasicID }
func (a *AnalogValueDescription) Value() *AnalogValueDescription { return a }
func (a AnalogValueDescription) MarshalZcl() ([]byte, error) {
	return zcl.Zcstring(a).MarshalZcl()
}
func (a *AnalogValueDescription) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zcstring)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogValueDescription(*nt)
	return br, err
}

func (a AnalogValueDescription) Readable() bool   { return true }
func (a AnalogValueDescription) Writable() bool   { return true }
func (a AnalogValueDescription) Reportable() bool { return false }
func (a AnalogValueDescription) SceneIndex() int  { return -1 }

func (a AnalogValueDescription) String() string {
	return zcl.Sprintf("%s", zcl.Zcstring(a))
}

const AnalogValueOutOfServiceAttr zcl.AttrID = 81

type AnalogValueOutOfService zcl.Zbool

func (a AnalogValueOutOfService) ID() zcl.AttrID                   { return AnalogValueOutOfServiceAttr }
func (a AnalogValueOutOfService) Cluster() zcl.ClusterID           { return AnalogValueBasicID }
func (a *AnalogValueOutOfService) Value() *AnalogValueOutOfService { return a }
func (a AnalogValueOutOfService) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *AnalogValueOutOfService) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogValueOutOfService(*nt)
	return br, err
}

func (a AnalogValueOutOfService) Readable() bool   { return true }
func (a AnalogValueOutOfService) Writable() bool   { return true }
func (a AnalogValueOutOfService) Reportable() bool { return false }
func (a AnalogValueOutOfService) SceneIndex() int  { return -1 }

func (a AnalogValueOutOfService) String() string {
	return zcl.Sprintf("%s", zcl.Zbool(a))
}

const AnalogValuePresentValueAttr zcl.AttrID = 85

type AnalogValuePresentValue zcl.Zfloat

func (a AnalogValuePresentValue) ID() zcl.AttrID                   { return AnalogValuePresentValueAttr }
func (a AnalogValuePresentValue) Cluster() zcl.ClusterID           { return AnalogValueBasicID }
func (a *AnalogValuePresentValue) Value() *AnalogValuePresentValue { return a }
func (a AnalogValuePresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogValuePresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogValuePresentValue(*nt)
	return br, err
}

func (a AnalogValuePresentValue) Readable() bool   { return true }
func (a AnalogValuePresentValue) Writable() bool   { return true }
func (a AnalogValuePresentValue) Reportable() bool { return true }
func (a AnalogValuePresentValue) SceneIndex() int  { return -1 }

func (a AnalogValuePresentValue) String() string {
	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

const AnalogValuePriorityArrayAttr zcl.AttrID = 87

type AnalogValuePriorityArray zcl.Zarray

func (a AnalogValuePriorityArray) ID() zcl.AttrID                    { return AnalogValuePriorityArrayAttr }
func (a AnalogValuePriorityArray) Cluster() zcl.ClusterID            { return AnalogValueBasicID }
func (a *AnalogValuePriorityArray) Value() *AnalogValuePriorityArray { return a }
func (a AnalogValuePriorityArray) MarshalZcl() ([]byte, error) {
	return zcl.Zarray(a).MarshalZcl()
}
func (a *AnalogValuePriorityArray) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zarray)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogValuePriorityArray(*nt)
	return br, err
}

func (a AnalogValuePriorityArray) Readable() bool   { return true }
func (a AnalogValuePriorityArray) Writable() bool   { return true }
func (a AnalogValuePriorityArray) Reportable() bool { return false }
func (a AnalogValuePriorityArray) SceneIndex() int  { return -1 }

func (a AnalogValuePriorityArray) String() string {
	return zcl.Sprintf("%s", zcl.Zarray(a))
}

const AnalogValueReliabilityAttr zcl.AttrID = 103

type AnalogValueReliability zcl.Zenum8

func (a AnalogValueReliability) ID() zcl.AttrID                  { return AnalogValueReliabilityAttr }
func (a AnalogValueReliability) Cluster() zcl.ClusterID          { return AnalogValueBasicID }
func (a *AnalogValueReliability) Value() *AnalogValueReliability { return a }
func (a AnalogValueReliability) MarshalZcl() ([]byte, error) {
	return zcl.Zenum8(a).MarshalZcl()
}
func (a *AnalogValueReliability) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zenum8)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogValueReliability(*nt)
	return br, err
}

func (a AnalogValueReliability) Readable() bool   { return true }
func (a AnalogValueReliability) Writable() bool   { return true }
func (a AnalogValueReliability) Reportable() bool { return false }
func (a AnalogValueReliability) SceneIndex() int  { return -1 }

func (a AnalogValueReliability) String() string {
	return zcl.Sprintf("%s", zcl.Zenum8(a))
}

const AnalogValueRelinquishDefaultAttr zcl.AttrID = 103

type AnalogValueRelinquishDefault zcl.Zfloat

func (a AnalogValueRelinquishDefault) ID() zcl.AttrID                        { return AnalogValueRelinquishDefaultAttr }
func (a AnalogValueRelinquishDefault) Cluster() zcl.ClusterID                { return AnalogValueBasicID }
func (a *AnalogValueRelinquishDefault) Value() *AnalogValueRelinquishDefault { return a }
func (a AnalogValueRelinquishDefault) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogValueRelinquishDefault) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogValueRelinquishDefault(*nt)
	return br, err
}

func (a AnalogValueRelinquishDefault) Readable() bool   { return true }
func (a AnalogValueRelinquishDefault) Writable() bool   { return true }
func (a AnalogValueRelinquishDefault) Reportable() bool { return false }
func (a AnalogValueRelinquishDefault) SceneIndex() int  { return -1 }

func (a AnalogValueRelinquishDefault) String() string {
	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

const AnalogValueResolutionAttr zcl.AttrID = 106

type AnalogValueResolution zcl.Zfloat

func (a AnalogValueResolution) ID() zcl.AttrID                 { return AnalogValueResolutionAttr }
func (a AnalogValueResolution) Cluster() zcl.ClusterID         { return AnalogValueBasicID }
func (a *AnalogValueResolution) Value() *AnalogValueResolution { return a }
func (a AnalogValueResolution) MarshalZcl() ([]byte, error) {
	return zcl.Zfloat(a).MarshalZcl()
}
func (a *AnalogValueResolution) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zfloat)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogValueResolution(*nt)
	return br, err
}

func (a AnalogValueResolution) Readable() bool   { return true }
func (a AnalogValueResolution) Writable() bool   { return true }
func (a AnalogValueResolution) Reportable() bool { return false }
func (a AnalogValueResolution) SceneIndex() int  { return -1 }

func (a AnalogValueResolution) String() string {
	return zcl.Sprintf("%s", zcl.Zfloat(a))
}

const AnalogValueStatusFlagsAttr zcl.AttrID = 111

type AnalogValueStatusFlags zcl.Zbmp8

func (a AnalogValueStatusFlags) ID() zcl.AttrID                  { return AnalogValueStatusFlagsAttr }
func (a AnalogValueStatusFlags) Cluster() zcl.ClusterID          { return AnalogValueBasicID }
func (a *AnalogValueStatusFlags) Value() *AnalogValueStatusFlags { return a }
func (a AnalogValueStatusFlags) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *AnalogValueStatusFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogValueStatusFlags(*nt)
	return br, err
}

func (a AnalogValueStatusFlags) Readable() bool   { return true }
func (a AnalogValueStatusFlags) Writable() bool   { return false }
func (a AnalogValueStatusFlags) Reportable() bool { return true }
func (a AnalogValueStatusFlags) SceneIndex() int  { return -1 }

func (a AnalogValueStatusFlags) String() string {
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

func (a AnalogValueStatusFlags) IsInAlarm() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *AnalogValueStatusFlags) SetInAlarm(b bool) {
	*a = AnalogValueStatusFlags(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a AnalogValueStatusFlags) IsFault() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *AnalogValueStatusFlags) SetFault(b bool) {
	*a = AnalogValueStatusFlags(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a AnalogValueStatusFlags) IsOveridden() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *AnalogValueStatusFlags) SetOveridden(b bool) {
	*a = AnalogValueStatusFlags(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a AnalogValueStatusFlags) IsOutOfService() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *AnalogValueStatusFlags) SetOutOfService(b bool) {
	*a = AnalogValueStatusFlags(zcl.BitmapSet([]byte(*a), 8, b))
}

const AnalogValueEngineeringUnitsAttr zcl.AttrID = 117

type AnalogValueEngineeringUnits zcl.EngineeringUnit

func (a AnalogValueEngineeringUnits) ID() zcl.AttrID                       { return AnalogValueEngineeringUnitsAttr }
func (a AnalogValueEngineeringUnits) Cluster() zcl.ClusterID               { return AnalogValueBasicID }
func (a *AnalogValueEngineeringUnits) Value() *AnalogValueEngineeringUnits { return a }
func (a AnalogValueEngineeringUnits) MarshalZcl() ([]byte, error) {
	return zcl.EngineeringUnit(a).MarshalZcl()
}
func (a *AnalogValueEngineeringUnits) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.EngineeringUnit)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogValueEngineeringUnits(*nt)
	return br, err
}

func (a AnalogValueEngineeringUnits) Readable() bool   { return true }
func (a AnalogValueEngineeringUnits) Writable() bool   { return true }
func (a AnalogValueEngineeringUnits) Reportable() bool { return false }
func (a AnalogValueEngineeringUnits) SceneIndex() int  { return -1 }

func (a AnalogValueEngineeringUnits) String() string {
	return zcl.Sprintf("%s", zcl.EngineeringUnit(a))
}

const AnalogValueApplicationTypeAttr zcl.AttrID = 256

type AnalogValueApplicationType zcl.Zu32

func (a AnalogValueApplicationType) ID() zcl.AttrID                      { return AnalogValueApplicationTypeAttr }
func (a AnalogValueApplicationType) Cluster() zcl.ClusterID              { return AnalogValueBasicID }
func (a *AnalogValueApplicationType) Value() *AnalogValueApplicationType { return a }
func (a AnalogValueApplicationType) MarshalZcl() ([]byte, error) {
	return zcl.Zu32(a).MarshalZcl()
}
func (a *AnalogValueApplicationType) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu32)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogValueApplicationType(*nt)
	return br, err
}

func (a AnalogValueApplicationType) Readable() bool   { return true }
func (a AnalogValueApplicationType) Writable() bool   { return false }
func (a AnalogValueApplicationType) Reportable() bool { return false }
func (a AnalogValueApplicationType) SceneIndex() int  { return -1 }

func (a AnalogValueApplicationType) String() string {
	return zcl.Sprintf("%s", zcl.Zu32(a))
}
