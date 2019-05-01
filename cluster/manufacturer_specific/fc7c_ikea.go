// IKEA control outlet cluster.
package manufacturer_specific

import (
	"neotor.se/zcl"
)

// Ikea
// Manufacturer code 0x117C
// IKEA control outlet cluster.

func NewIkeaServer(profile zcl.ProfileID) *IkeaServer { return &IkeaServer{p: profile} }
func NewIkeaClient(profile zcl.ProfileID) *IkeaClient { return &IkeaClient{p: profile} }

const IkeaCluster zcl.ClusterID = 64636

type IkeaServer struct {
	p zcl.ProfileID

	Unknown1 *Unknown1
	Unknown2 *Unknown2
}

type IkeaClient struct {
	p zcl.ProfileID
}

/*
var IkeaServer = map[zcl.CommandID]func() zcl.Command{
}

var IkeaClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type Unknown1 zcl.Zu8

func (a Unknown1) ID() zcl.AttrID         { return 16 }
func (a Unknown1) Cluster() zcl.ClusterID { return IkeaCluster }
func (a *Unknown1) Value() *Unknown1      { return a }
func (a Unknown1) MarshalZcl() ([]byte, error) {
	return zcl.Zu8(a).MarshalZcl()
}
func (a *Unknown1) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu8)
	br, err := nt.UnmarshalZcl(b)
	*a = Unknown1(*nt)
	return br, err
}

func (a Unknown1) Readable() bool   { return true }
func (a Unknown1) Writable() bool   { return true }
func (a Unknown1) Reportable() bool { return false }
func (a Unknown1) SceneIndex() int  { return -1 }

func (a Unknown1) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu8(a))
}

type Unknown2 zcl.Zu16

func (a Unknown2) ID() zcl.AttrID         { return 65533 }
func (a Unknown2) Cluster() zcl.ClusterID { return IkeaCluster }
func (a *Unknown2) Value() *Unknown2      { return a }
func (a Unknown2) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *Unknown2) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = Unknown2(*nt)
	return br, err
}

func (a Unknown2) Readable() bool   { return true }
func (a Unknown2) Writable() bool   { return true }
func (a Unknown2) Reportable() bool { return false }
func (a Unknown2) SceneIndex() int  { return -1 }

func (a Unknown2) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu16(a))
}
