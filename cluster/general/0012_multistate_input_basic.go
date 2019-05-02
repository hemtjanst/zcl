// Provides an interface for reading the value of a multistate measurement and accessing various characteristics of that measurement. The cluster is typically used to implement a sensor that measures a physical quantity that can take on one of a number of discrete states.
package general

import (
	"neotor.se/zcl"
)

// MultistateInputBasic
const MultistateInputBasicID zcl.ClusterID = 18

var MultistateInputBasicCluster = zcl.Cluster{
	ServerCmd: map[zcl.CommandID]func() zcl.Command{},
	ClientCmd: map[zcl.CommandID]func() zcl.Command{},
	ServerAttr: map[zcl.AttrID]func() zcl.Attr{
		MultiInputNumberOfStatesAttr: func() zcl.Attr { return new(MultiInputNumberOfStates) },
		MultiInputOutOfServiceAttr:   func() zcl.Attr { return new(MultiInputOutOfService) },
		MultiInputPresentValueAttr:   func() zcl.Attr { return new(MultiInputPresentValue) },
		MultiInputStatusFlagsAttr:    func() zcl.Attr { return new(MultiInputStatusFlags) },
	},
	ClientAttr: map[zcl.AttrID]func() zcl.Attr{},
	SceneAttr:  []zcl.AttrID{},
}

const MultiInputNumberOfStatesAttr zcl.AttrID = 74

type MultiInputNumberOfStates zcl.Zu16

func (a MultiInputNumberOfStates) ID() zcl.AttrID                    { return MultiInputNumberOfStatesAttr }
func (a MultiInputNumberOfStates) Cluster() zcl.ClusterID            { return MultistateInputBasicID }
func (a *MultiInputNumberOfStates) Value() *MultiInputNumberOfStates { return a }
func (a MultiInputNumberOfStates) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MultiInputNumberOfStates) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MultiInputNumberOfStates(*nt)
	return br, err
}

func (a MultiInputNumberOfStates) Readable() bool   { return true }
func (a MultiInputNumberOfStates) Writable() bool   { return true }
func (a MultiInputNumberOfStates) Reportable() bool { return false }
func (a MultiInputNumberOfStates) SceneIndex() int  { return -1 }

func (a MultiInputNumberOfStates) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MultiInputOutOfServiceAttr zcl.AttrID = 81

type MultiInputOutOfService zcl.Zbool

func (a MultiInputOutOfService) ID() zcl.AttrID                  { return MultiInputOutOfServiceAttr }
func (a MultiInputOutOfService) Cluster() zcl.ClusterID          { return MultistateInputBasicID }
func (a *MultiInputOutOfService) Value() *MultiInputOutOfService { return a }
func (a MultiInputOutOfService) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *MultiInputOutOfService) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = MultiInputOutOfService(*nt)
	return br, err
}

func (a MultiInputOutOfService) Readable() bool   { return true }
func (a MultiInputOutOfService) Writable() bool   { return true }
func (a MultiInputOutOfService) Reportable() bool { return false }
func (a MultiInputOutOfService) SceneIndex() int  { return -1 }

func (a MultiInputOutOfService) String() string {
	return zcl.Sprintf("%s", zcl.Zbool(a))
}

const MultiInputPresentValueAttr zcl.AttrID = 85

type MultiInputPresentValue zcl.Zu16

func (a MultiInputPresentValue) ID() zcl.AttrID                  { return MultiInputPresentValueAttr }
func (a MultiInputPresentValue) Cluster() zcl.ClusterID          { return MultistateInputBasicID }
func (a *MultiInputPresentValue) Value() *MultiInputPresentValue { return a }
func (a MultiInputPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MultiInputPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MultiInputPresentValue(*nt)
	return br, err
}

func (a MultiInputPresentValue) Readable() bool   { return true }
func (a MultiInputPresentValue) Writable() bool   { return true }
func (a MultiInputPresentValue) Reportable() bool { return false }
func (a MultiInputPresentValue) SceneIndex() int  { return -1 }

func (a MultiInputPresentValue) String() string {
	return zcl.Sprintf("%s", zcl.Zu16(a))
}

const MultiInputStatusFlagsAttr zcl.AttrID = 111

type MultiInputStatusFlags zcl.Zbmp8

func (a MultiInputStatusFlags) ID() zcl.AttrID                 { return MultiInputStatusFlagsAttr }
func (a MultiInputStatusFlags) Cluster() zcl.ClusterID         { return MultistateInputBasicID }
func (a *MultiInputStatusFlags) Value() *MultiInputStatusFlags { return a }
func (a MultiInputStatusFlags) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *MultiInputStatusFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = MultiInputStatusFlags(*nt)
	return br, err
}

func (a MultiInputStatusFlags) Readable() bool   { return true }
func (a MultiInputStatusFlags) Writable() bool   { return false }
func (a MultiInputStatusFlags) Reportable() bool { return false }
func (a MultiInputStatusFlags) SceneIndex() int  { return -1 }

func (a MultiInputStatusFlags) String() string {
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

func (a MultiInputStatusFlags) IsInAlarm() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *MultiInputStatusFlags) SetInAlarm(b bool) {
	*a = MultiInputStatusFlags(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a MultiInputStatusFlags) IsFault() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *MultiInputStatusFlags) SetFault(b bool) {
	*a = MultiInputStatusFlags(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a MultiInputStatusFlags) IsOveridden() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *MultiInputStatusFlags) SetOveridden(b bool) {
	*a = MultiInputStatusFlags(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a MultiInputStatusFlags) IsOutOfService() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *MultiInputStatusFlags) SetOutOfService(b bool) {
	*a = MultiInputStatusFlags(zcl.BitmapSet([]byte(*a), 8, b))
}
