// The Binary Input (Basic) cluster provides an interface for reading the value of a binary measurement and accessing various characteristics of that measurement. The cluster is typically used to implement a sensor that measures a two-state physical quantity.
package general

import (
	"neotor.se/zcl/cluster/zcl"
)

// BinaryInputBasic
// The Binary Input (Basic) cluster provides an interface for reading the value of a binary measurement and accessing various characteristics of that measurement. The cluster is typically used to implement a sensor that measures a two-state physical quantity.

func NewBinaryInputBasicServer(profile zcl.ProfileID) *BinaryInputBasicServer {
	return &BinaryInputBasicServer{p: profile}
}
func NewBinaryInputBasicClient(profile zcl.ProfileID) *BinaryInputBasicClient {
	return &BinaryInputBasicClient{p: profile}
}

const BinaryInputBasicCluster zcl.ClusterID = 15

type BinaryInputBasicServer struct {
	p zcl.ProfileID

	BinaryInputOutOfService *BinaryInputOutOfService
	BinaryInputPresentValue *BinaryInputPresentValue
	BinaryInputStatusFlags  *BinaryInputStatusFlags
}

type BinaryInputBasicClient struct {
	p zcl.ProfileID
}

/*
var BinaryInputBasicServer = map[zcl.CommandID]func() zcl.Command{
}

var BinaryInputBasicClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type BinaryInputOutOfService zcl.Zbool

func (a BinaryInputOutOfService) ID() zcl.AttrID                   { return 81 }
func (a BinaryInputOutOfService) Cluster() zcl.ClusterID           { return BinaryInputBasicCluster }
func (a *BinaryInputOutOfService) Value() *BinaryInputOutOfService { return a }
func (a BinaryInputOutOfService) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *BinaryInputOutOfService) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryInputOutOfService(*nt)
	return br, err
}

func (a BinaryInputOutOfService) Readable() bool   { return true }
func (a BinaryInputOutOfService) Writable() bool   { return true }
func (a BinaryInputOutOfService) Reportable() bool { return false }
func (a BinaryInputOutOfService) SceneIndex() int  { return -1 }

func (a BinaryInputOutOfService) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

type BinaryInputPresentValue zcl.Zbool

func (a BinaryInputPresentValue) ID() zcl.AttrID                   { return 85 }
func (a BinaryInputPresentValue) Cluster() zcl.ClusterID           { return BinaryInputBasicCluster }
func (a *BinaryInputPresentValue) Value() *BinaryInputPresentValue { return a }
func (a BinaryInputPresentValue) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *BinaryInputPresentValue) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryInputPresentValue(*nt)
	return br, err
}

func (a BinaryInputPresentValue) Readable() bool   { return true }
func (a BinaryInputPresentValue) Writable() bool   { return true }
func (a BinaryInputPresentValue) Reportable() bool { return false }
func (a BinaryInputPresentValue) SceneIndex() int  { return -1 }

func (a BinaryInputPresentValue) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

type BinaryInputStatusFlags zcl.Zbmp8

func (a BinaryInputStatusFlags) ID() zcl.AttrID                  { return 111 }
func (a BinaryInputStatusFlags) Cluster() zcl.ClusterID          { return BinaryInputBasicCluster }
func (a *BinaryInputStatusFlags) Value() *BinaryInputStatusFlags { return a }
func (a BinaryInputStatusFlags) MarshalZcl() ([]byte, error) {
	return zcl.Zbmp8(a).MarshalZcl()
}
func (a *BinaryInputStatusFlags) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbmp8)
	br, err := nt.UnmarshalZcl(b)
	*a = BinaryInputStatusFlags(*nt)
	return br, err
}

func (a BinaryInputStatusFlags) Readable() bool   { return true }
func (a BinaryInputStatusFlags) Writable() bool   { return false }
func (a BinaryInputStatusFlags) Reportable() bool { return false }
func (a BinaryInputStatusFlags) SceneIndex() int  { return -1 }

func (a BinaryInputStatusFlags) String() string {

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

func (a BinaryInputStatusFlags) IsInAlarm() bool {
	return zcl.BitmapTest([]byte(a), 1)
}
func (a *BinaryInputStatusFlags) SetInAlarm(b bool) {
	*a = BinaryInputStatusFlags(zcl.BitmapSet([]byte(*a), 1, b))
}

func (a BinaryInputStatusFlags) IsFault() bool {
	return zcl.BitmapTest([]byte(a), 2)
}
func (a *BinaryInputStatusFlags) SetFault(b bool) {
	*a = BinaryInputStatusFlags(zcl.BitmapSet([]byte(*a), 2, b))
}

func (a BinaryInputStatusFlags) IsOveridden() bool {
	return zcl.BitmapTest([]byte(a), 4)
}
func (a *BinaryInputStatusFlags) SetOveridden(b bool) {
	*a = BinaryInputStatusFlags(zcl.BitmapSet([]byte(*a), 4, b))
}

func (a BinaryInputStatusFlags) IsOutOfService() bool {
	return zcl.BitmapTest([]byte(a), 8)
}
func (a *BinaryInputStatusFlags) SetOutOfService(b bool) {
	*a = BinaryInputStatusFlags(zcl.BitmapSet([]byte(*a), 8, b))
}
