// An interface for reading the value of an analog measurement and accessing various characteristics of that measurement.
package general

import (
	"neotor.se/zcl"
)

// AnalogInputBasic
const AnalogInputBasicID zcl.ClusterID = 12

var AnalogInputBasicCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		AnalogInputDescriptionAttr:      func() zcl.Attr { return new(AnalogInputDescription) },
		AnalogInputMaxPresentValueAttr:  func() zcl.Attr { return new(AnalogInputMaxPresentValue) },
		AnalogInputMinPresentValueAttr:  func() zcl.Attr { return new(AnalogInputMinPresentValue) },
		AnalogInputOutOfServiceAttr:     func() zcl.Attr { return new(AnalogInputOutOfService) },
		AnalogInputPresentValueAttr:     func() zcl.Attr { return new(AnalogInputPresentValue) },
		AnalogInputReliabilityAttr:      func() zcl.Attr { return new(AnalogInputReliability) },
		AnalogInputResolutionAttr:       func() zcl.Attr { return new(AnalogInputResolution) },
		AnalogInputStatusFlagsAttr:      func() zcl.Attr { return new(AnalogInputStatusFlags) },
		AnalogInputEngineeringUnitsAttr: func() zcl.Attr { return new(AnalogInputEngineeringUnits) },
		AnalogInputApplicationTypeAttr:  func() zcl.Attr { return new(AnalogInputApplicationType) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const AnalogInputDescriptionAttr zcl.AttrID = 28

type AnalogInputDescription zcl.Zcstring

func (a AnalogInputDescription) ID() zcl.AttrID                  { return AnalogInputDescriptionAttr }
func (a AnalogInputDescription) Cluster() zcl.ClusterID          { return AnalogInputBasicID }
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

const AnalogInputMaxPresentValueAttr zcl.AttrID = 65

type AnalogInputMaxPresentValue zcl.Zfloat

func (a AnalogInputMaxPresentValue) ID() zcl.AttrID                      { return AnalogInputMaxPresentValueAttr }
func (a AnalogInputMaxPresentValue) Cluster() zcl.ClusterID              { return AnalogInputBasicID }
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

const AnalogInputMinPresentValueAttr zcl.AttrID = 66

type AnalogInputMinPresentValue zcl.Zfloat

func (a AnalogInputMinPresentValue) ID() zcl.AttrID                      { return AnalogInputMinPresentValueAttr }
func (a AnalogInputMinPresentValue) Cluster() zcl.ClusterID              { return AnalogInputBasicID }
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

const AnalogInputOutOfServiceAttr zcl.AttrID = 81

type AnalogInputOutOfService zcl.Zbool

func (a AnalogInputOutOfService) ID() zcl.AttrID                   { return AnalogInputOutOfServiceAttr }
func (a AnalogInputOutOfService) Cluster() zcl.ClusterID           { return AnalogInputBasicID }
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

const AnalogInputPresentValueAttr zcl.AttrID = 85

type AnalogInputPresentValue zcl.Zfloat

func (a AnalogInputPresentValue) ID() zcl.AttrID                   { return AnalogInputPresentValueAttr }
func (a AnalogInputPresentValue) Cluster() zcl.ClusterID           { return AnalogInputBasicID }
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

const AnalogInputReliabilityAttr zcl.AttrID = 103

type AnalogInputReliability zcl.Zenum8

func (a AnalogInputReliability) ID() zcl.AttrID                  { return AnalogInputReliabilityAttr }
func (a AnalogInputReliability) Cluster() zcl.ClusterID          { return AnalogInputBasicID }
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

const AnalogInputResolutionAttr zcl.AttrID = 106

type AnalogInputResolution zcl.Zfloat

func (a AnalogInputResolution) ID() zcl.AttrID                 { return AnalogInputResolutionAttr }
func (a AnalogInputResolution) Cluster() zcl.ClusterID         { return AnalogInputBasicID }
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

const AnalogInputStatusFlagsAttr zcl.AttrID = 111

type AnalogInputStatusFlags zcl.Zbmp8

func (a AnalogInputStatusFlags) ID() zcl.AttrID                  { return AnalogInputStatusFlagsAttr }
func (a AnalogInputStatusFlags) Cluster() zcl.ClusterID          { return AnalogInputBasicID }
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

const AnalogInputEngineeringUnitsAttr zcl.AttrID = 117

type AnalogInputEngineeringUnits zcl.EngineeringUnit

func (a AnalogInputEngineeringUnits) ID() zcl.AttrID                       { return AnalogInputEngineeringUnitsAttr }
func (a AnalogInputEngineeringUnits) Cluster() zcl.ClusterID               { return AnalogInputBasicID }
func (a *AnalogInputEngineeringUnits) Value() *AnalogInputEngineeringUnits { return a }
func (a AnalogInputEngineeringUnits) MarshalZcl() ([]byte, error) {
	return zcl.EngineeringUnit(a).MarshalZcl()
}
func (a *AnalogInputEngineeringUnits) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.EngineeringUnit)
	br, err := nt.UnmarshalZcl(b)
	*a = AnalogInputEngineeringUnits(*nt)
	return br, err
}

func (a AnalogInputEngineeringUnits) Readable() bool   { return true }
func (a AnalogInputEngineeringUnits) Writable() bool   { return true }
func (a AnalogInputEngineeringUnits) Reportable() bool { return false }
func (a AnalogInputEngineeringUnits) SceneIndex() int  { return -1 }

func (a AnalogInputEngineeringUnits) String() string {
	return zcl.Sprintf("%s", zcl.EngineeringUnit(a))
}

const AnalogInputApplicationTypeAttr zcl.AttrID = 256

type AnalogInputApplicationType zcl.Zu32

func (a AnalogInputApplicationType) ID() zcl.AttrID                      { return AnalogInputApplicationTypeAttr }
func (a AnalogInputApplicationType) Cluster() zcl.ClusterID              { return AnalogInputBasicID }
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
