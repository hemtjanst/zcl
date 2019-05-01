package general

import (
	"neotor.se/zcl"
)

// MultistateOutputBasic

func NewMultistateOutputBasicServer(profile zcl.ProfileID) *MultistateOutputBasicServer {
	return &MultistateOutputBasicServer{p: profile}
}
func NewMultistateOutputBasicClient(profile zcl.ProfileID) *MultistateOutputBasicClient {
	return &MultistateOutputBasicClient{p: profile}
}

const MultistateOutputBasicCluster zcl.ClusterID = 19

type MultistateOutputBasicServer struct {
	p zcl.ProfileID

	MultiOutputNumberOfStates *MultiOutputNumberOfStates
	MultiOutputOutOfService   *MultiOutputOutOfService
	MultiOutputPresentValue   *MultiOutputPresentValue
	MultiOutputStatusFlags    *MultiOutputStatusFlags
}

type MultistateOutputBasicClient struct {
	p zcl.ProfileID
}

/*
var MultistateOutputBasicServer = map[zcl.CommandID]func() zcl.Command{
}

var MultistateOutputBasicClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type MultiOutputNumberOfStates zcl.Zu16

func (a MultiOutputNumberOfStates) ID() zcl.AttrID                     { return 74 }
func (a MultiOutputNumberOfStates) Cluster() zcl.ClusterID             { return MultistateOutputBasicCluster }
func (a *MultiOutputNumberOfStates) Value() *MultiOutputNumberOfStates { return a }
func (a MultiOutputNumberOfStates) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MultiOutputNumberOfStates) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MultiOutputNumberOfStates(*nt)
	return br, err
}

func (a MultiOutputNumberOfStates) Readable() bool   { return true }
func (a MultiOutputNumberOfStates) Writable() bool   { return true }
func (a MultiOutputNumberOfStates) Reportable() bool { return false }
func (a MultiOutputNumberOfStates) SceneIndex() int  { return -1 }

func (a MultiOutputNumberOfStates) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type MultiOutputOutOfService zcl.Zbool

func (a MultiOutputOutOfService) ID() zcl.AttrID                   { return 81 }
func (a MultiOutputOutOfService) Cluster() zcl.ClusterID           { return MultistateOutputBasicCluster }
func (a *MultiOutputOutOfService) Value() *MultiOutputOutOfService { return a }
func (a MultiOutputOutOfService) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *MultiOutputOutOfService) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = MultiOutputOutOfService(*nt)
	return br, err
}

func (a MultiOutputOutOfService) Readable() bool   { return true }
func (a MultiOutputOutOfService) Writable() bool   { return true }
func (a MultiOutputOutOfService) Reportable() bool { return false }
func (a MultiOutputOutOfService) SceneIndex() int  { return -1 }

func (a MultiOutputOutOfService) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

type MultiOutputPresentValue zcl.Zu16

func (a MultiOutputPresentValue) ID() zcl.AttrID                   { return 85 }
func (a MultiOutputPresentValue) Cluster() zcl.ClusterID           { return MultistateOutputBasicCluster }
func (a *MultiOutputPresentValue) Value() *MultiOutputPresentValue { return a }
func (a MultiOutputPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *MultiOutputPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = MultiOutputPresentValue(*nt)
	return br, err
}

func (a MultiOutputPresentValue) Readable() bool   { return true }
func (a MultiOutputPresentValue) Writable() bool   { return true }
func (a MultiOutputPresentValue) Reportable() bool { return false }
func (a MultiOutputPresentValue) SceneIndex() int  { return -1 }

func (a MultiOutputPresentValue) String() string {

	return zcl.Sprintf("%s", zcl.Zu16(a))
}

type MultiOutputStatusFlags zcl.Zbmp8

func (a MultiOutputStatusFlags) ID() zcl.AttrID                  { return 111 }
func (a MultiOutputStatusFlags) Cluster() zcl.ClusterID          { return MultistateOutputBasicCluster }
func (a *MultiOutputStatusFlags) Value() *MultiOutputStatusFlags { return a }
func (a MultiOutputStatusFlags) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *MultiOutputStatusFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = MultiOutputStatusFlags(*nt)
	return br, err
}

func (a MultiOutputStatusFlags) Readable() bool   { return true }
func (a MultiOutputStatusFlags) Writable() bool   { return false }
func (a MultiOutputStatusFlags) Reportable() bool { return false }
func (a MultiOutputStatusFlags) SceneIndex() int  { return -1 }

func (a MultiOutputStatusFlags) String() string {

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

func (a MultiOutputStatusFlags) IsInAlarm() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *MultiOutputStatusFlags) SetInAlarm(b bool) {
	*a = MultiOutputStatusFlags(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a MultiOutputStatusFlags) IsFault() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *MultiOutputStatusFlags) SetFault(b bool) {
	*a = MultiOutputStatusFlags(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a MultiOutputStatusFlags) IsOveridden() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *MultiOutputStatusFlags) SetOveridden(b bool) {
	*a = MultiOutputStatusFlags(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a MultiOutputStatusFlags) IsOutOfService() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *MultiOutputStatusFlags) SetOutOfService(b bool) {
	*a = MultiOutputStatusFlags(zcl.BitmapSet([]byte(*a), 8, b))
}
