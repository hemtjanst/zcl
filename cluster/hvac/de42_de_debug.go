// Attributes and commands for debugging purposes.
package hvac

import (
	"neotor.se/zcl"
)

// DeDebug
// Attributes and commands for debugging purposes.

func NewDeDebugServer(profile zcl.ProfileID) *DeDebugServer { return &DeDebugServer{p: profile} }
func NewDeDebugClient(profile zcl.ProfileID) *DeDebugClient { return &DeDebugClient{p: profile} }

const DeDebugCluster zcl.ClusterID = 56898

type DeDebugServer struct {
	p zcl.ProfileID

	DebugEnabled     *DebugEnabled
	DebugDestination *DebugDestination
}

type DeDebugClient struct {
	p zcl.ProfileID
}

/*
var DeDebugServer = map[zcl.CommandID]func() zcl.Command{
}

var DeDebugClient = map[zcl.CommandID]func() zcl.Command{
}
*/

type DebugEnabled zcl.Zbool

func (a DebugEnabled) ID() zcl.AttrID         { return 0 }
func (a DebugEnabled) Cluster() zcl.ClusterID { return DeDebugCluster }
func (a *DebugEnabled) Value() *DebugEnabled  { return a }
func (a DebugEnabled) MarshalZcl() ([]byte, error) {
	return zcl.Zbool(a).MarshalZcl()
}
func (a *DebugEnabled) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zbool)
	br, err := nt.UnmarshalZcl(b)
	*a = DebugEnabled(*nt)
	return br, err
}

func (a DebugEnabled) Readable() bool   { return true }
func (a DebugEnabled) Writable() bool   { return true }
func (a DebugEnabled) Reportable() bool { return false }
func (a DebugEnabled) SceneIndex() int  { return -1 }

func (a DebugEnabled) String() string {

	return zcl.Sprintf("%s", zcl.Zbool(a))
}

type DebugDestination zcl.Zu16

func (a DebugDestination) ID() zcl.AttrID            { return 1 }
func (a DebugDestination) Cluster() zcl.ClusterID    { return DeDebugCluster }
func (a *DebugDestination) Value() *DebugDestination { return a }
func (a DebugDestination) MarshalZcl() ([]byte, error) {
	return zcl.Zu16(a).MarshalZcl()
}
func (a *DebugDestination) UnmarshalZcl(b []byte) ([]byte, error) {
	nt := new(zcl.Zu16)
	br, err := nt.UnmarshalZcl(b)
	*a = DebugDestination(*nt)
	return br, err
}

func (a DebugDestination) Readable() bool   { return true }
func (a DebugDestination) Writable() bool   { return true }
func (a DebugDestination) Reportable() bool { return false }
func (a DebugDestination) SceneIndex() int  { return -1 }

func (a DebugDestination) String() string {

	return zcl.Sprintf("0x%X", zcl.Zu16(a))
}
